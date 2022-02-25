/*
 * Copyright (c) 2016-2021 Deephaven Data Labs and Patent Pending
 */

package io.deephaven.engine.table.impl.sources;

import io.deephaven.engine.table.ColumnSource;
import io.deephaven.engine.table.impl.MutableColumnSourceGetDefaults;

import static io.deephaven.util.type.TypeUtils.box;
import static io.deephaven.util.type.TypeUtils.unbox;

/**
 * Simple array source for Long.
 */
public class LongSparseArraySource extends AbstractSparseLongArraySource<Long> implements MutableColumnSourceGetDefaults.ForLong {
    public LongSparseArraySource() {
        super(long.class);
    }

    @Override
    public void set(long key, Long value) {
        set(key, unbox(value));
    }
}