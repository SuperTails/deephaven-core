/**
 * Copyright (c) 2016-2022 Deephaven Data Labs and Patent Pending
 */
package io.deephaven.parquet.base;

import java.io.IOException;
import java.nio.IntBuffer;

public interface ColumnWriter {
    void addPageNoNulls(Object pageData, int valuesCount) throws IOException;

    void addDictionaryPage(Object dictionaryValues, int valuesCount) throws IOException;

    void addPage(Object pageData, Object nullValues, int valuesCount) throws IOException;

    void addVectorPage(Object pageData, IntBuffer repeatCount, int valuesCount, Object nullValue) throws IOException;

    void close();
}
