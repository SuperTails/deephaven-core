package io.deephaven.engine.table.impl.tuplesource.generated;

import io.deephaven.chunk.CharChunk;
import io.deephaven.chunk.Chunk;
import io.deephaven.chunk.IntChunk;
import io.deephaven.chunk.ObjectChunk;
import io.deephaven.chunk.WritableChunk;
import io.deephaven.chunk.WritableObjectChunk;
import io.deephaven.chunk.attributes.Values;
import io.deephaven.datastructures.util.SmartKey;
import io.deephaven.engine.table.ColumnSource;
import io.deephaven.engine.table.TupleSource;
import io.deephaven.engine.table.WritableColumnSource;
import io.deephaven.engine.table.impl.tuplesource.AbstractTupleSource;
import io.deephaven.engine.table.impl.tuplesource.ThreeColumnTupleSourceFactory;
import io.deephaven.time.DateTime;
import io.deephaven.time.DateTimeUtils;
import io.deephaven.tuple.generated.IntCharLongTuple;
import io.deephaven.util.type.TypeUtils;
import org.jetbrains.annotations.NotNull;


/**
 * <p>{@link TupleSource} that produces key column values from {@link ColumnSource} types Integer, Character, and DateTime.
 * <p>Generated by io.deephaven.replicators.TupleSourceCodeGenerator.
 */
@SuppressWarnings({"unused", "WeakerAccess"})
public class IntegerCharacterDateTimeColumnTupleSource extends AbstractTupleSource<IntCharLongTuple> {

    /** {@link ThreeColumnTupleSourceFactory} instance to create instances of {@link IntegerCharacterDateTimeColumnTupleSource}. **/
    public static final ThreeColumnTupleSourceFactory<IntCharLongTuple, Integer, Character, DateTime> FACTORY = new Factory();

    private final ColumnSource<Integer> columnSource1;
    private final ColumnSource<Character> columnSource2;
    private final ColumnSource<DateTime> columnSource3;

    public IntegerCharacterDateTimeColumnTupleSource(
            @NotNull final ColumnSource<Integer> columnSource1,
            @NotNull final ColumnSource<Character> columnSource2,
            @NotNull final ColumnSource<DateTime> columnSource3
    ) {
        super(columnSource1, columnSource2, columnSource3);
        this.columnSource1 = columnSource1;
        this.columnSource2 = columnSource2;
        this.columnSource3 = columnSource3;
    }

    @Override
    public final IntCharLongTuple createTuple(final long rowKey) {
        return new IntCharLongTuple(
                columnSource1.getInt(rowKey),
                columnSource2.getChar(rowKey),
                DateTimeUtils.nanos(columnSource3.get(rowKey))
        );
    }

    @Override
    public final IntCharLongTuple createPreviousTuple(final long rowKey) {
        return new IntCharLongTuple(
                columnSource1.getPrevInt(rowKey),
                columnSource2.getPrevChar(rowKey),
                DateTimeUtils.nanos(columnSource3.getPrev(rowKey))
        );
    }

    @Override
    public final IntCharLongTuple createTupleFromValues(@NotNull final Object... values) {
        return new IntCharLongTuple(
                TypeUtils.unbox((Integer)values[0]),
                TypeUtils.unbox((Character)values[1]),
                DateTimeUtils.nanos((DateTime)values[2])
        );
    }

    @Override
    public final IntCharLongTuple createTupleFromReinterpretedValues(@NotNull final Object... values) {
        return new IntCharLongTuple(
                TypeUtils.unbox((Integer)values[0]),
                TypeUtils.unbox((Character)values[1]),
                DateTimeUtils.nanos((DateTime)values[2])
        );
    }

    @SuppressWarnings("unchecked")
    @Override
    public final <ELEMENT_TYPE> void exportElement(@NotNull final IntCharLongTuple tuple, final int elementIndex, @NotNull final WritableColumnSource<ELEMENT_TYPE> writableSource, final long destinationRowKey) {
        if (elementIndex == 0) {
            writableSource.set(destinationRowKey, tuple.getFirstElement());
            return;
        }
        if (elementIndex == 1) {
            writableSource.set(destinationRowKey, tuple.getSecondElement());
            return;
        }
        if (elementIndex == 2) {
            writableSource.set(destinationRowKey, (ELEMENT_TYPE) DateTimeUtils.nanosToTime(tuple.getThirdElement()));
            return;
        }
        throw new IndexOutOfBoundsException("Invalid element index " + elementIndex + " for export");
    }

    @Override
    public final Object exportToExternalKey(@NotNull final IntCharLongTuple tuple) {
        return new SmartKey(
                TypeUtils.box(tuple.getFirstElement()),
                TypeUtils.box(tuple.getSecondElement()),
                DateTimeUtils.nanosToTime(tuple.getThirdElement())
        );
    }

    @Override
    public final Object exportElement(@NotNull final IntCharLongTuple tuple, int elementIndex) {
        if (elementIndex == 0) {
            return TypeUtils.box(tuple.getFirstElement());
        }
        if (elementIndex == 1) {
            return TypeUtils.box(tuple.getSecondElement());
        }
        if (elementIndex == 2) {
            return DateTimeUtils.nanosToTime(tuple.getThirdElement());
        }
        throw new IllegalArgumentException("Bad elementIndex for 3 element tuple: " + elementIndex);
    }

    @Override
    public final Object exportElementReinterpreted(@NotNull final IntCharLongTuple tuple, int elementIndex) {
        if (elementIndex == 0) {
            return TypeUtils.box(tuple.getFirstElement());
        }
        if (elementIndex == 1) {
            return TypeUtils.box(tuple.getSecondElement());
        }
        if (elementIndex == 2) {
            return DateTimeUtils.nanosToTime(tuple.getThirdElement());
        }
        throw new IllegalArgumentException("Bad elementIndex for 3 element tuple: " + elementIndex);
    }

    @Override
    protected void convertChunks(@NotNull WritableChunk<? super Values> destination, int chunkSize, Chunk<Values> [] chunks) {
        WritableObjectChunk<IntCharLongTuple, ? super Values> destinationObjectChunk = destination.asWritableObjectChunk();
        IntChunk<Values> chunk1 = chunks[0].asIntChunk();
        CharChunk<Values> chunk2 = chunks[1].asCharChunk();
        ObjectChunk<DateTime, Values> chunk3 = chunks[2].asObjectChunk();
        for (int ii = 0; ii < chunkSize; ++ii) {
            destinationObjectChunk.set(ii, new IntCharLongTuple(chunk1.get(ii), chunk2.get(ii), DateTimeUtils.nanos(chunk3.get(ii))));
        }
        destinationObjectChunk.setSize(chunkSize);
    }

    /** {@link ThreeColumnTupleSourceFactory} for instances of {@link IntegerCharacterDateTimeColumnTupleSource}. **/
    private static final class Factory implements ThreeColumnTupleSourceFactory<IntCharLongTuple, Integer, Character, DateTime> {

        private Factory() {
        }

        @Override
        public TupleSource<IntCharLongTuple> create(
                @NotNull final ColumnSource<Integer> columnSource1,
                @NotNull final ColumnSource<Character> columnSource2,
                @NotNull final ColumnSource<DateTime> columnSource3
        ) {
            return new IntegerCharacterDateTimeColumnTupleSource(
                    columnSource1,
                    columnSource2,
                    columnSource3
            );
        }
    }
}