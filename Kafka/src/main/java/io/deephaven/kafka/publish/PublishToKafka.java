package io.deephaven.kafka.publish;

import io.deephaven.base.verify.Assert;
import io.deephaven.configuration.Configuration;
import io.deephaven.datastructures.util.CollectionUtil;
import io.deephaven.db.tables.Table;
import io.deephaven.db.tables.live.LiveTableMonitor;
import io.deephaven.db.util.liveness.LivenessArtifact;
import io.deephaven.db.util.liveness.LivenessScope;
import io.deephaven.db.v2.*;
import io.deephaven.db.v2.sources.chunk.Attributes;
import io.deephaven.db.v2.sources.chunk.ObjectChunk;
import io.deephaven.db.v2.utils.Index;
import io.deephaven.db.v2.utils.OrderedKeys;
import io.deephaven.db.v2.utils.ReadOnlyIndex;
import io.deephaven.util.SafeCloseable;
import io.deephaven.util.annotations.ReferentialIntegrity;
import org.apache.kafka.clients.producer.Callback;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.clients.producer.RecordMetadata;
import org.jetbrains.annotations.NotNull;

import java.util.Properties;
import java.util.concurrent.atomic.AtomicLong;
import java.util.concurrent.atomic.AtomicReference;

public class PublishToKafka<K, V> extends LivenessArtifact {

    public static final int CHUNK_SIZE =
            Configuration.getInstance().getIntegerForClassWithDefault(PublishToKafka.class, "chunkSize", 2048);

    private final Table table;
    private final KafkaProducer<K, V> producer;
    private final String topic;
    private final KeyOrValueSerializer<K> keySerializer;
    private final KeyOrValueSerializer<V> valueSerializer;

    @ReferentialIntegrity
    private final PublishListener publishListener;

    /**
     * <p>
     * Construct a publisher for {@code table} according the to Kafka {@code props} for the supplied {@code topic}.
     * <p>
     * The new publisher will produce records for existing {@code table} data at construction.
     * <p>
     * If {@code table} is a dynamic, refreshing table ({@link Table#isLive()}), the calling thread must block the
     * {@link LiveTableMonitor#DEFAULT LiveTableMonitor} by holding either its {@link LiveTableMonitor#exclusiveLock()
     * exclusive lock} or its {@link LiveTableMonitor#sharedLock() shared lock}. The publisher will install a listener
     * in order to produce new records as updates become available. Callers must be sure to maintain a reference to the
     * publisher and ensure that it remains {@link io.deephaven.db.util.liveness.LivenessReferent live}. The easiest way
     * to do this may be to construct the publisher enclosed by a {@link io.deephaven.db.util.liveness.LivenessScope
     * liveness scope} with {@code enforceStrongReachability} specified as {@code true}, and
     * {@link LivenessScope#release() release} the scope when publication is no longer needed. For example:
     * 
     * <pre>
     *     // To initiate publication:
     *     final LivenessScope publisherScope = new LivenessScope(true);
     *     try (final SafeCloseable ignored = LivenessScopeStack.open(publisherScope, false)) {
     *         new PublishToKafka(...);
     *     }
     *     // To cease publication:
     *     publisherScope.release();
     * </pre>
     *
     * @param props The Kafka {@link Properties}
     * @param table The source {@link Table}
     * @param topic The destination topic
     * @param keySerializerFactory Optional factory for a {@link KeyOrValueSerializer} to produce Kafka record keys
     * @param valueSerializerFactory Optional factory for a {@link KeyOrValueSerializer} to produce Kafka record values
     * @param collapseByKeyColumns Whether to publish only the last record for each unique key. Ignored when
     *        {@code keySerializerFactory} is {@code null}. If
     *        {@code keySerializerFactory != null && !collapseByKeyColumns}, it is expected that {@code table} will not
     *        produce any row shifts; that is, the publisher expects keyed tables to be streams, add-only, or
     *        aggregated.
     */
    public PublishToKafka(final Properties props,
            Table table,
            final String topic,
            final KeyOrValueSerializer.Factory<K> keySerializerFactory,
            final KeyOrValueSerializer.Factory<V> valueSerializerFactory,
            final boolean collapseByKeyColumns) {
        if (keySerializerFactory != null) {
            keySerializerFactory.validateColumns(table.getDefinition());
        }
        if (valueSerializerFactory != null) {
            valueSerializerFactory.validateColumns(table.getDefinition());
        }
        if (table.isLive()
                && !LiveTableMonitor.DEFAULT.exclusiveLock().isHeldByCurrentThread()
                && !LiveTableMonitor.DEFAULT.sharedLock().isHeldByCurrentThread()) {
            throw new KafkaPublisherException(
                    "Calling thread must hold an exclusive or shared LiveTableMonitor lock to publish live sources");
        }

        this.table = table = (keySerializerFactory != null && collapseByKeyColumns)
                ? table.lastBy(keySerializerFactory.sourceColumnNames(table.getDefinition()))
                : table.coalesce();
        this.producer = new KafkaProducer<>(props);
        this.topic = topic;
        this.keySerializer = keySerializerFactory == null ? null : keySerializerFactory.create(table);
        this.valueSerializer = valueSerializerFactory == null ? null : valueSerializerFactory.create(table);

        // Publish the initial table state
        try (final PublicationGuard guard = new PublicationGuard()) {
            publishMessages(table.getIndex(), false, true, guard);
        }

        // Install a listener to publish subsequent updates
        if (table.isLive()) {
            ((DynamicTable) table).listenForUpdates(publishListener = new PublishListener(
                    getModifiedColumnSet(table, keySerializerFactory),
                    getModifiedColumnSet(table, valueSerializerFactory)));
            manage(publishListener);
        } else {
            publishListener = null;
            producer.close();
        }
    }

    private static ModifiedColumnSet getModifiedColumnSet(@NotNull final Table table,
            final KeyOrValueSerializer.Factory<?> serializerFactory) {
        return serializerFactory == null
                ? ModifiedColumnSet.EMPTY
                : ((BaseTable) table).newModifiedColumnSet(
                        serializerFactory.sourceColumnNames(table.getDefinition())
                                .toArray(CollectionUtil.ZERO_LENGTH_STRING_ARRAY));
    }

    private void publishMessages(@NotNull final ReadOnlyIndex rowsToPublish, final boolean usePrevious,
            final boolean publishValues, @NotNull final Callback callback) {
        if (rowsToPublish.isEmpty()) {
            return;
        }

        final int chunkSize = (int) Math.min(CHUNK_SIZE, rowsToPublish.size());
        try (final OrderedKeys.Iterator rowsIterator = rowsToPublish.getOrderedKeysIterator();
                final KeyOrValueSerializer.Context keyContext =
                        keySerializer != null ? keySerializer.makeContext(chunkSize) : null;
                final KeyOrValueSerializer.Context valueContext =
                        publishValues && valueSerializer != null ? valueSerializer.makeContext(chunkSize) : null) {
            while (rowsIterator.hasMore()) {
                final OrderedKeys chunkRowKeys = rowsIterator.getNextOrderedKeysWithLength(chunkSize);

                final ObjectChunk<K, Attributes.Values> keyChunk;
                if (keyContext != null) {
                    keyChunk = keySerializer.handleChunk(keyContext, chunkRowKeys, usePrevious);
                } else {
                    keyChunk = null;
                }

                final ObjectChunk<V, Attributes.Values> valueChunk;
                if (valueContext != null) {
                    valueChunk = valueSerializer.handleChunk(valueContext, chunkRowKeys, usePrevious);
                } else {
                    valueChunk = null;
                }

                for (int ii = 0; ii < chunkRowKeys.intSize(); ++ii) {
                    final ProducerRecord<K, V> record = new ProducerRecord<>(topic,
                            keyChunk != null ? keyChunk.get(ii) : null, valueChunk != null ? valueChunk.get(ii) : null);
                    producer.send(record, callback);
                }
            }
        }
    }

    /**
     * Re-usable, {@link SafeCloseable} {@link Callback} used to bracket multiple calls to
     * {@link KafkaProducer#send(ProducerRecord, Callback) send} and ensure correct completion. Used in the following
     * pattern:
     * 
     * <pre>
     * final PublicationGuard guard = new PublicationGuard();
     * try (final Closeable ignored = guard) {
     *     // Call producer.send(record, guard) 0 or more times
     * }
     * </pre>
     */
    private class PublicationGuard implements Callback, SafeCloseable {

        private long sentCount;
        private final AtomicLong completedCount = new AtomicLong();
        private final AtomicReference<Exception> sendException = new AtomicReference<>();

        private void reset() {
            sentCount = 0;
            completedCount.set(0);
            sendException.set(null);
        }

        @Override
        public void onCompletion(@NotNull final RecordMetadata metadata, final Exception exception) {
            completedCount.getAndIncrement();
            if (exception != null) {
                sendException.compareAndSet(null, exception);
            }
        }

        @Override
        public void close() {
            try {
                if (sentCount == 0) {
                    return;
                }
                try {
                    producer.flush();
                } catch (Exception e) {
                    throw new KafkaPublisherException("KafkaProducer reported flush failure", e);
                }
                final Exception localSendException = sendException.get();
                if (localSendException != null) {
                    throw new KafkaPublisherException("KafkaProducer reported send failure", localSendException);
                }
                final long localCompletedCount = completedCount.get();
                if (sentCount != localCompletedCount) {
                    throw new KafkaPublisherException(String.format("Sent count %d does not match completed count %d",
                            sentCount, localCompletedCount));
                }
            } finally {
                reset();
            }
        }
    }

    private class PublishListener extends InstrumentedShiftAwareListenerAdapter {

        private final ModifiedColumnSet keysModified;
        private final ModifiedColumnSet valuesModified;
        private final boolean isStream;

        private final PublicationGuard guard = new PublicationGuard();

        private PublishListener(
                @NotNull final ModifiedColumnSet keysModified,
                @NotNull final ModifiedColumnSet valuesModified) {
            super("PublishToKafka", (DynamicTable) table, false);
            this.keysModified = keysModified;
            this.valuesModified = valuesModified;
            this.isStream = StreamTableTools.isStream(table);
        }

        @Override
        public void onUpdate(Update upstream) {
            if (keySerializer != null || isStream) {
                Assert.assertion(upstream.shifted.empty(), "upstream.shifted.empty()");
            }

            try (final SafeCloseable ignored = guard) {
                if (isStream) {
                    Assert.assertion(upstream.modified.empty(), "upstream.modified.empty()");
                    // We always ignore removes on streams, and expect no modifies or shifts
                    publishMessages(upstream.added, false, true, guard);
                    return;
                }

                // Regular table, either keyless, add-only, or aggregated
                publishMessages(upstream.removed, true, false, guard);
                if (keysModified.containsAny(upstream.modifiedColumnSet)
                        || valuesModified.containsAny(upstream.modifiedColumnSet)) {
                    try (final Index addedAndModified = upstream.added.union(upstream.modified)) {
                        publishMessages(addedAndModified, false, true, guard);
                    }
                } else {
                    publishMessages(upstream.added, false, true, guard);
                }
            }
        }
    }

    @Override
    protected void destroy() {
        super.destroy();
        producer.close();
    }
}