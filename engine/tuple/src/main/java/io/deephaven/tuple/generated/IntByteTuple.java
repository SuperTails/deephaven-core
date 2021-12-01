package io.deephaven.tuple.generated;

import gnu.trove.map.TIntObjectMap;
import io.deephaven.tuple.CanonicalizableTuple;
import io.deephaven.tuple.serialization.SerializationUtils;
import io.deephaven.tuple.serialization.StreamingExternalizable;
import io.deephaven.util.compare.ByteComparisons;
import io.deephaven.util.compare.IntComparisons;
import org.jetbrains.annotations.NotNull;

import java.io.Externalizable;
import java.io.IOException;
import java.io.ObjectInput;
import java.io.ObjectOutput;
import java.util.function.UnaryOperator;

/**
 * <p>2-Tuple (double) key class composed of int and byte elements.
 * <p>Generated by io.deephaven.replicators.TupleCodeGenerator.
 */
public class IntByteTuple implements Comparable<IntByteTuple>, Externalizable, StreamingExternalizable, CanonicalizableTuple<IntByteTuple> {

    private static final long serialVersionUID = 1L;

    private int element1;
    private byte element2;

    private transient int cachedHashCode;

    public IntByteTuple(
            final int element1,
            final byte element2
    ) {
        initialize(
                element1,
                element2
        );
    }

    /** Public no-arg constructor for {@link Externalizable} support only. <em>Application code should not use this!</em> **/
    public IntByteTuple() {
    }

    private void initialize(
            final int element1,
            final byte element2
    ) {
        this.element1 = element1;
        this.element2 = element2;
        cachedHashCode = (31 +
                Integer.hashCode(element1)) * 31 +
                Byte.hashCode(element2);
    }

    public final int getFirstElement() {
        return element1;
    }

    public final byte getSecondElement() {
        return element2;
    }

    @Override
    public final int hashCode() {
        return cachedHashCode;
    }

    @Override
    public final boolean equals(final Object other) {
        if (this == other) {
            return true;
        }
        if (other == null || getClass() != other.getClass()) {
            return false;
        }
        final IntByteTuple typedOther = (IntByteTuple) other;
        // @formatter:off
        return element1 == typedOther.element1 &&
               element2 == typedOther.element2;
        // @formatter:on
    }

    @Override
    public final int compareTo(@NotNull final IntByteTuple other) {
        if (this == other) {
            return 0;
        }
        int comparison;
        // @formatter:off
        return 0 != (comparison = IntComparisons.compare(element1, other.element1)) ? comparison :
               ByteComparisons.compare(element2, other.element2);
        // @formatter:on
    }

    @Override
    public void writeExternal(@NotNull final ObjectOutput out) throws IOException {
        out.writeInt(element1);
        out.writeByte(element2);
    }

    @Override
    public void readExternal(@NotNull final ObjectInput in) throws IOException, ClassNotFoundException {
        initialize(
                in.readInt(),
                in.readByte()
        );
    }

    @Override
    public void writeExternalStreaming(@NotNull final ObjectOutput out, @NotNull final TIntObjectMap<SerializationUtils.Writer> cachedWriters) throws IOException {
        out.writeInt(element1);
        out.writeByte(element2);
    }

    @Override
    public void readExternalStreaming(@NotNull final ObjectInput in, @NotNull final TIntObjectMap<SerializationUtils.Reader> cachedReaders) throws Exception {
        initialize(
                in.readInt(),
                in.readByte()
        );
    }

    @Override
    public String toString() {
        return "IntByteTuple{" +
                element1 + ", " +
                element2 + '}';
    }

    @Override
    public IntByteTuple canonicalize(@NotNull final UnaryOperator<Object> canonicalizer) {
        return this;
    }
}