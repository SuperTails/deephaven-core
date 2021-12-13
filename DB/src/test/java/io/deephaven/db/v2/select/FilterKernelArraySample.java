package io.deephaven.db.v2.select;
// QueryLibrary internal version number: DEFAULT
import java.lang.*;
import java.util.*;
import io.deephaven.base.string.cache.CompressedString;
import io.deephaven.db.tables.DataColumn;
import io.deephaven.db.tables.Table;
import io.deephaven.db.tables.select.Param;
import io.deephaven.db.tables.utils.ArrayUtils;
import io.deephaven.db.tables.utils.DBDateTime;
import io.deephaven.db.tables.utils.DBPeriod;
import io.deephaven.db.tables.utils.DBTimeUtils;
import io.deephaven.db.util.LongSizedDataStructure;
import static io.deephaven.db.v2.select.ConditionFilter.FilterKernel;
import io.deephaven.db.v2.sources.ColumnSource;
import io.deephaven.db.v2.sources.chunk.Attributes;
import io.deephaven.db.v2.sources.chunk.ByteChunk;
import io.deephaven.db.v2.sources.chunk.CharChunk;
import io.deephaven.db.v2.sources.chunk.Chunk;
import io.deephaven.db.v2.sources.chunk.Context;
import io.deephaven.db.v2.sources.chunk.DoubleChunk;
import io.deephaven.db.v2.sources.chunk.FloatChunk;
import io.deephaven.db.v2.sources.chunk.IntChunk;
import io.deephaven.db.v2.sources.chunk.LongChunk;
import io.deephaven.db.v2.sources.chunk.ObjectChunk;
import io.deephaven.db.v2.sources.chunk.ShortChunk;
import io.deephaven.db.v2.sources.chunk.WritableByteChunk;
import io.deephaven.db.v2.sources.chunk.WritableCharChunk;
import io.deephaven.db.v2.sources.chunk.WritableChunk;
import io.deephaven.db.v2.sources.chunk.WritableDoubleChunk;
import io.deephaven.db.v2.sources.chunk.WritableFloatChunk;
import io.deephaven.db.v2.sources.chunk.WritableIntChunk;
import io.deephaven.db.v2.sources.chunk.WritableLongChunk;
import io.deephaven.db.v2.sources.chunk.WritableObjectChunk;
import io.deephaven.db.v2.sources.chunk.WritableShortChunk;
import io.deephaven.db.v2.utils.Index;
import static io.deephaven.db.v2.utils.Index.SequentialBuilder;
import io.deephaven.db.v2.utils.IndexBuilder;
import io.deephaven.db.v2.utils.OrderedKeys;
import io.deephaven.util.type.TypeUtils;
import java.lang.reflect.Array;
import java.util.concurrent.ConcurrentHashMap;
import org.joda.time.LocalTime;
import static io.deephaven.base.string.cache.CompressedString.*;
import static io.deephaven.db.tables.lang.DBLanguageFunctionUtil.*;
import static io.deephaven.db.tables.utils.DBTimeUtils.*;
import static io.deephaven.db.tables.utils.DBTimeZone.*;
import static io.deephaven.db.tables.utils.WhereClause.*;
import static io.deephaven.db.tables.verify.TableAssertions.*;
import static io.deephaven.db.util.DBColorUtilImpl.*;
import static io.deephaven.db.v2.sources.chunk.Attributes.*;
import static io.deephaven.gui.color.Color.*;
import static io.deephaven.libs.primitives.BinSearch.*;
import static io.deephaven.libs.primitives.BooleanPrimitives.*;
import static io.deephaven.libs.primitives.ByteNumericPrimitives.*;
import static io.deephaven.libs.primitives.BytePrimitives.*;
import static io.deephaven.libs.primitives.Casting.*;
import static io.deephaven.libs.primitives.CharacterPrimitives.*;
import static io.deephaven.libs.primitives.ComparePrimitives.*;
import static io.deephaven.libs.primitives.DoubleFpPrimitives.*;
import static io.deephaven.libs.primitives.DoubleNumericPrimitives.*;
import static io.deephaven.libs.primitives.DoublePrimitives.*;
import static io.deephaven.libs.primitives.FloatFpPrimitives.*;
import static io.deephaven.libs.primitives.FloatNumericPrimitives.*;
import static io.deephaven.libs.primitives.FloatPrimitives.*;
import static io.deephaven.libs.primitives.IntegerNumericPrimitives.*;
import static io.deephaven.libs.primitives.IntegerPrimitives.*;
import static io.deephaven.libs.primitives.LongNumericPrimitives.*;
import static io.deephaven.libs.primitives.LongPrimitives.*;
import static io.deephaven.libs.primitives.ObjectPrimitives.*;
import static io.deephaven.libs.primitives.PrimitiveParseUtil.*;
import static io.deephaven.libs.primitives.ShortNumericPrimitives.*;
import static io.deephaven.libs.primitives.ShortPrimitives.*;
import static io.deephaven.util.QueryConstants.*;
import static io.deephaven.util.calendar.StaticCalendarMethods.*;

public class FilterKernelArraySample implements io.deephaven.db.v2.select.ConditionFilter.FilterKernel<FilterKernel.Context>{


    // Array Column Variables
    private final io.deephaven.db.tables.dbarrays.DbDoubleArray v2_;
    private final io.deephaven.db.tables.dbarrays.DbShortArray v1_;


    public FilterKernelArraySample(Table table, Index fullSet, Param... params) {

        // Array Column Variables
        v2_ = new io.deephaven.db.v2.dbarrays.DbDoubleArrayColumnWrapper(table.getColumnSource("v2"), fullSet);
        v1_ = new io.deephaven.db.v2.dbarrays.DbShortArrayColumnWrapper(table.getColumnSource("v1"), fullSet);
    }
    @Override
    public Context getContext(int maxChunkSize) {
        return new Context(maxChunkSize);
    }
    
    @Override
    public LongChunk<OrderedKeyIndices> filter(Context context, LongChunk<OrderedKeyIndices> indices, Chunk... inputChunks) {
        final int size = indices.size();
        context.resultChunk.setSize(0);
        for (int __my_i__ = 0; __my_i__ < size; __my_i__++) {
            if (eq(v1_.size(), v2_.size())) {
                context.resultChunk.add(indices.get(__my_i__));
            }
        }
        return context.resultChunk;
    }
}