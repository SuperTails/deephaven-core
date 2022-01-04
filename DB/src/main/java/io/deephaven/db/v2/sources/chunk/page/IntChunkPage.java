/* ---------------------------------------------------------------------------------------------------------------------
 * AUTO-GENERATED CLASS - DO NOT EDIT MANUALLY - for any changes edit CharChunkPage and regenerate
 * ------------------------------------------------------------------------------------------------------------------ */
package io.deephaven.db.v2.sources.chunk.page;

import io.deephaven.base.verify.Require;
import io.deephaven.db.v2.sources.chunk.*;
import io.deephaven.db.v2.utils.OrderedKeys;
import org.jetbrains.annotations.NotNull;

public class IntChunkPage<ATTR extends Attributes.Any> extends IntChunk<ATTR> implements ChunkPage<ATTR> {

    private final long mask;
    private final long firstRow;

    public static <ATTR extends Attributes.Any> IntChunkPage<ATTR> pageWrap(long beginRow, int[] data, int offset, int capacity, long mask) {
        return new IntChunkPage<>(beginRow, data, offset, capacity, mask);
    }

    public static <ATTR extends Attributes.Any> IntChunkPage<ATTR> pageWrap(long beginRow, int[] data, long mask) {
        return new IntChunkPage<>(beginRow, data, 0, data.length, mask);
    }

    private IntChunkPage(long firstRow, int[] data, int offset, int capacity, long mask) {
        super(data, offset, Require.lt(capacity, "capacity", Integer.MAX_VALUE, "INT_MAX"));
        this.mask = mask;
        this.firstRow = Require.inRange(firstRow, "firstRow", mask, "mask");
    }

    @Override
    public final void fillChunkAppend(@NotNull FillContext context, @NotNull WritableChunk<? super ATTR> destination, @NotNull OrderedKeys orderedKeys) {
        WritableIntChunk<? super ATTR> to = destination.asWritableIntChunk();

        if (orderedKeys.getAverageRunLengthEstimate() >= Chunk.SYSTEM_ARRAYCOPY_THRESHOLD) {
            orderedKeys.forAllLongRanges((final long rangeStartKey, final long rangeEndKey) ->
                    to.appendTypedChunk(this, getChunkOffset(rangeStartKey), (int) (rangeEndKey - rangeStartKey + 1)));
        } else {
            orderedKeys.forEachLong((final long key) -> {
                to.add(get(getChunkOffset(key)));
                return true;
            });
        }
    }

    @Override
    public final long firstRowOffset() {
        return firstRow;
    }

    @Override
    public final long mask() {
        return mask;
    }
}