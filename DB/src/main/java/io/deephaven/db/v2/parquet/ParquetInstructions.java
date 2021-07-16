package io.deephaven.db.v2.parquet;

import io.deephaven.db.v2.ColumnToCodecMappings;
import io.deephaven.hash.KeyedObjectHashMap;
import io.deephaven.hash.KeyedObjectKey;
import io.deephaven.util.annotations.VisibleForTesting;
import org.apache.parquet.hadoop.metadata.CompressionCodecName;
import org.jetbrains.annotations.NotNull;

import java.util.Collections;
import java.util.Objects;
import java.util.Set;
import java.util.function.Function;

/**
 * This class provides instructions intended for read and write parquet operations (which take
 * it as an optional argument) specifying desired transformations.  Examples are
 * mapping column names and use of specific codecs during (de)serialization.
 */
public abstract class ParquetInstructions implements ColumnToCodecMappings {
    private static volatile String defaultCompressionCodecName = CompressionCodecName.SNAPPY.toString();
    public static void setDefaultCompressionCodecName(final String name) {
        defaultCompressionCodecName = name;
    }

    public ParquetInstructions() {
    }

    public final String getColumnNameFromParquetColumnNameOrDefault(final String parquetColumnName) {
        final String mapped = getColumnNameFromParquetColumnName(parquetColumnName);
        return (mapped != null) ? mapped : parquetColumnName;
    }
    public abstract String getParquetColumnNameFromColumnNameOrDefault(final String columnName);
    public abstract String getColumnNameFromParquetColumnName(final String parquetColumnName);
    @Override public abstract String getCodecName(final String columnName);
    @Override public abstract String getCodecArgs(final String columnName);
    public abstract String getCompressionCodecName();
    public abstract boolean isLegacyParquet();

    @VisibleForTesting
    public static boolean sameColumnNamesAndCodecMappings(final ParquetInstructions i1, final ParquetInstructions i2) {
        if (i1 == EMPTY) {
            if (i2 == EMPTY) {
                return true;
            }
            return ((ReadOnly) i2).columnNameToInstructions.size() == 0;
        }
        if (i2 == EMPTY) {
            return ((ReadOnly) i1).columnNameToInstructions.size() == 0;
        }
        return ReadOnly.sameCodecMappings((ReadOnly) i1, (ReadOnly) i2);
    }

    public static final ParquetInstructions EMPTY = new ParquetInstructions() {
        @Override
        public String getParquetColumnNameFromColumnNameOrDefault(final String columnName) {
            return columnName;
        }
        @Override
        public String getColumnNameFromParquetColumnName(final String parquetColumnName) {
            return null;
        }
        @Override
        public String getCodecName(final String columnName) {
            return null;
        }
        @Override
        public String getCodecArgs(final String columnName) {
            return null;
        }

        @Override
        public String getCompressionCodecName() {
            return defaultCompressionCodecName;
        }

        @Override
        public boolean isLegacyParquet() {
            return false;
        }
    };

    private static class ColumnInstructions {
        private final String columnName;
        private String parquetColumnName;
        private String codecName;
        private String codecArgs;

        public ColumnInstructions(final String columnName) {
            this.columnName = columnName;
        }

        public String getColumnName() {
            return columnName;
        }

        public String getParquetColumnName() {
            return parquetColumnName != null ? parquetColumnName : columnName;
        }
        public ColumnInstructions setParquetColumnName(final String parquetColumnName) {
            this.parquetColumnName = parquetColumnName;
            return this;
        }

        public String getCodecName() {
            return codecName;
        }
        public ColumnInstructions setCodecName(final String codecName) {
            this.codecName = codecName;
            return this;
        }

        public String getCodecArgs() {
            return codecArgs;
        }
        public ColumnInstructions setCodecArgs(final String codecArgs) {
            this.codecArgs = codecArgs;
            return this;
        }
    }

    private static final class ReadOnly extends ParquetInstructions {
        private final KeyedObjectHashMap<String, ColumnInstructions> columnNameToInstructions;
        /**
         * Note parquetColumnNameToInstructions may be null while columnNameToInstructions is not null;
         * We only store entries in parquetColumnNameToInstructions when the parquetColumnName is
         * different than the columnName (ie, the column name mapping is not the default mapping)
         */
        private final KeyedObjectHashMap<String, ColumnInstructions> parquetColumnNameToInstructions;
        private final String compressionCodecName;
        private final boolean isLegacyParquet;

        protected ReadOnly(
                final KeyedObjectHashMap<String, ColumnInstructions> columnNameToInstructions,
                final KeyedObjectHashMap<String, ColumnInstructions> parquetColumnNameToColumnName,
                final String compressionCodecName,
                final boolean isLegacyParquet) {
            this.columnNameToInstructions = columnNameToInstructions;
            this.parquetColumnNameToInstructions = parquetColumnNameToColumnName;
            this.compressionCodecName = compressionCodecName;
            this.isLegacyParquet = isLegacyParquet;
        }

        private String getOrDefault(final String columnName, final String defaultValue, final Function<ColumnInstructions, String> fun) {
            if (columnNameToInstructions == null) {
                return defaultValue;
            }
            final ColumnInstructions ci = columnNameToInstructions.get(columnName);
            if (ci == null) {
                return defaultValue;
            }
            return fun.apply(ci);
        }

        @Override
        public String getParquetColumnNameFromColumnNameOrDefault(final String columnName) {
            return getOrDefault(columnName, columnName, ColumnInstructions::getParquetColumnName);
        }

        @Override
        public String getColumnNameFromParquetColumnName(final String parquetColumnName) {
            if (parquetColumnNameToInstructions == null) {
                return null;
            }
            final ColumnInstructions ci = parquetColumnNameToInstructions.get(parquetColumnName);
            if (ci == null) {
                return null;
            }
            return ci.getColumnName();
        }

        @Override
        public String getCodecName(final String columnName) {
            return getOrDefault(columnName, null, ColumnInstructions::getCodecName);
        }

        @Override
        public String getCodecArgs(final String columnName) {
            return getOrDefault(columnName, null, ColumnInstructions::getCodecArgs);
        }

        @Override
        public String getCompressionCodecName() {
            return compressionCodecName;
        }

        @Override
        public boolean isLegacyParquet() {
            return isLegacyParquet;
        }

        KeyedObjectHashMap<String, ColumnInstructions> copyColumnNameToInstructions() {
            // noinspection unchecked
            return (columnNameToInstructions == null)
                    ? null
                    : (KeyedObjectHashMap<String, ColumnInstructions>) columnNameToInstructions.clone()
                    ;
        }

        KeyedObjectHashMap<String, ColumnInstructions> copyParquetColumnNameToInstructions() {
            // noinspection unchecked
            return (parquetColumnNameToInstructions == null)
                    ? null
                    : (KeyedObjectHashMap<String, ColumnInstructions>) parquetColumnNameToInstructions.clone()
                    ;
        }

        private static boolean sameCodecMappings(final ReadOnly r1, final ReadOnly r2) {
            final Set<String> r1ColumnNames = r1.columnNameToInstructions.keySet();
            if (r2.columnNameToInstructions.size() != r1ColumnNames.size()) {
                return false;
            }
            for (String colName : r1ColumnNames) {
                if (!r2.columnNameToInstructions.containsKey(colName)) {
                    return false;
                }
                final String r1CodecName = r1.getCodecName(colName);
                final String r2CodecName = r2.getCodecName(colName);
                if (!Objects.equals(r1CodecName, r2CodecName)) {
                    return false;
                }
                final String r1CodecArgs = r1.getCodecArgs(colName);
                final String r2CodecArgs = r2.getCodecArgs(colName);
                if (!Objects.equals(r1CodecArgs, r2CodecArgs)) {
                    return false;
                }
            }
            return true;
        }
    }

    public static class Builder {
        private KeyedObjectHashMap<String, ColumnInstructions> columnNameToInstructions;
        // Note parquetColumnNameToInstructions may be null while columnNameToInstructions is not null;
        // We only store entries in parquetColumnNameToInstructions when the parquetColumnName is
        // different than the columnName (ie, the column name mapping is not the default mapping)
        private KeyedObjectHashMap<String, ColumnInstructions> parquetColumnNameToInstructions;
        private String compressionCodecName = defaultCompressionCodecName;
        private boolean isLegacyParquet;

        public Builder() {
        }

        public Builder(final ParquetInstructions parquetInstructions) {
            if (parquetInstructions == EMPTY) {
                return;
            }
            final ReadOnly readOnlyParquetInstructions = (ReadOnly) parquetInstructions;
            columnNameToInstructions = readOnlyParquetInstructions.copyColumnNameToInstructions();
            parquetColumnNameToInstructions = readOnlyParquetInstructions.copyParquetColumnNameToInstructions();
        }

        private void newColumnNameToInstructionsMap() {
            columnNameToInstructions = new KeyedObjectHashMap<>(new KeyedObjectKey.Basic<String, ColumnInstructions>() {
                @Override
                public String getKey(@NotNull final ColumnInstructions value) {
                    return value.getColumnName();
                }
            });
        }

        private void newParquetColumnNameToInstructionsMap() {
            parquetColumnNameToInstructions = new KeyedObjectHashMap<>(new KeyedObjectKey.Basic<String, ColumnInstructions>() {
                @Override
                public String getKey(@NotNull final ColumnInstructions value) {
                    return value.getParquetColumnName();
                }
            });
        }

        public Builder addColumnNameMapping(final String parquetColumnName, final String columnName) {
            if (parquetColumnName.equals(columnName)) {
                return this;
            }
            if (columnNameToInstructions == null) {
                newColumnNameToInstructionsMap();
                final ColumnInstructions ci = new ColumnInstructions(columnName);
                ci.setParquetColumnName(parquetColumnName);
                columnNameToInstructions.put(columnName, ci);
                newParquetColumnNameToInstructionsMap();
                parquetColumnNameToInstructions.put(parquetColumnName, ci);
                return this;
            }

            ColumnInstructions ci = columnNameToInstructions.get(columnName);
            if (ci != null) {
                if (ci.parquetColumnName != null) {
                    if (ci.parquetColumnName.equals(parquetColumnName)) {
                        return this;
                    }
                    throw new IllegalArgumentException(
                            "Cannot add a mapping from parquetColumnName=" + parquetColumnName
                            + ": columnName=" + columnName + " already mapped to parquetColumnName=" + ci.parquetColumnName);
                }
            } else {
                ci = new ColumnInstructions(columnName);
                columnNameToInstructions.put(columnName, ci);
            }

            if (parquetColumnNameToInstructions == null) {
                newParquetColumnNameToInstructionsMap();
                parquetColumnNameToInstructions.put(parquetColumnName, ci);
                return this;
            }

            final ColumnInstructions fromParquetColumnNameInstructions = parquetColumnNameToInstructions.get(parquetColumnName);
            if (fromParquetColumnNameInstructions != null) {
                if (fromParquetColumnNameInstructions == ci) {
                    return this;
                }
                throw new IllegalArgumentException(
                        "Cannot add new mapping from parquetColumnName=" + parquetColumnName + " to columnName=" + columnName
                                + ": already mapped to columnName=" + fromParquetColumnNameInstructions.getColumnName());
            }
            ci.setParquetColumnName(parquetColumnName);
            parquetColumnNameToInstructions.put(parquetColumnName, ci);
            return this;
        }

        public Set<String> getTakenNames() {
            return (columnNameToInstructions == null) ? Collections.emptySet() : columnNameToInstructions.keySet();
        }

        public Builder addColumnCodec(final String columnName, final String codecName) {
            return addColumnCodec(columnName, codecName, null);
        }

        public Builder addColumnCodec(final String columnName, final String codecName, final String codecArgs) {
            final ColumnInstructions ci;
            if (columnNameToInstructions == null) {
                newColumnNameToInstructionsMap();
                ci = new ColumnInstructions(columnName);
                columnNameToInstructions.put(columnName, ci);
            } else {
                ci = columnNameToInstructions.putIfAbsent(columnName, ColumnInstructions::new);
            }
            ci.codecName = codecName;
            ci.codecArgs = codecArgs;
            return this;
        }

        public Builder setCompressionCodecName(final String compressionCodecName) {
            this.compressionCodecName = compressionCodecName;
            return this;
        }

        public Builder setIsLegacyParquet(final boolean isLegacyParquet) {
            this.isLegacyParquet = isLegacyParquet;
            return this;
        }

        public ParquetInstructions build() {
            final KeyedObjectHashMap<String, ColumnInstructions> columnNameToInstructionsOut = columnNameToInstructions;
            columnNameToInstructions = null;
            final KeyedObjectHashMap<String, ColumnInstructions> parquetColumnNameToColumnNameOut = parquetColumnNameToInstructions;
            parquetColumnNameToInstructions = null;
            return new ReadOnly(columnNameToInstructionsOut, parquetColumnNameToColumnNameOut, compressionCodecName, isLegacyParquet);
        }
    }

    public static Builder builder() {
        return new Builder();
    }
}