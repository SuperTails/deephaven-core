/* ---------------------------------------------------------------------------------------------------------------------
 * AUTO-GENERATED CLASS - DO NOT EDIT MANUALLY - for any changes edit TestCharPrimitives and regenerate
 * ------------------------------------------------------------------------------------------------------------------ */
/*
 * Copyright (c) 2016-2021 Deephaven Data Labs and Patent Pending
 */

package io.deephaven.libs.primitives;

import io.deephaven.base.testing.BaseArrayTestCase;
import io.deephaven.db.tables.dbarrays.DbIntArray;
import io.deephaven.db.tables.dbarrays.DbIntArrayDirect;

import static io.deephaven.libs.primitives.IntegerPrimitives.*;
import static io.deephaven.util.QueryConstants.NULL_INT;
import static io.deephaven.util.QueryConstants.NULL_LONG;

public class TestIntPrimitives extends BaseArrayTestCase {

    public void testUnbox(){
        assertNull(unbox((Integer[])null));
        assertEquals(new int[]{1, NULL_INT, 3, NULL_INT}, unbox((int)1, null, (int)3, NULL_INT));
    }

    public void testIsNull(){
        assertFalse(isNull((int)3));
        assertTrue(isNull(NULL_INT));
    }

    public void testNullToValueScalar() {
        assertEquals((int) 3, nullToValue((int) 3, (int) 7));
        assertEquals((int) 7, nullToValue(NULL_INT, (int) 7));
    }

    public void testNullToValueArray() {
        assertEquals(new int[]{(int) 3, (int) 7, (int) 11}, nullToValue(new DbIntArrayDirect(new int[]{(int) 3, NULL_INT, (int) 11}), (int) 7));

        assertEquals(new int[]{(int) 3, (int) 7, (int) 11}, nullToValue(new int[]{(int) 3, NULL_INT, (int) 11}, (int) 7));
    }

    public void testCount(){
        assertEquals(0, count((DbIntArray)null));
        assertEquals(3,count(new DbIntArrayDirect(new int[]{40,50,60})));
        assertEquals(0,count(new DbIntArrayDirect()));
        assertEquals(0,count(new DbIntArrayDirect(NULL_INT)));
        assertEquals(2,count(new DbIntArrayDirect(new int[]{5,NULL_INT,15})));
    }

    public void testLast(){
        assertTrue(Math.abs(60-last(new DbIntArrayDirect(new int[]{40,50,60})))==0.0);
        assertEquals(NULL_INT,last(new DbIntArrayDirect()));
        assertEquals(NULL_INT,last(new DbIntArrayDirect(NULL_INT)));
        assertTrue(Math.abs(15-last(new DbIntArrayDirect(new int[]{5,NULL_INT,15})))==0.0);
        assertTrue(Math.abs(40-last(new DbIntArrayDirect((int)40)))==0.0);

        assertTrue(Math.abs(60-last(new int[]{40,50,60}))==0.0);
        assertEquals(NULL_INT,last(new int[]{}));
        assertEquals(NULL_INT,last(new int[]{NULL_INT}));
        assertTrue(Math.abs(15-last(new int[]{5,NULL_INT,15}))==0.0);
        assertTrue(Math.abs(40-last(new int[]{(int)40}))==0.0);
    }

    public void testFirst(){
        assertTrue(Math.abs(40-first(new DbIntArrayDirect(new int[]{40,50,60})))==0.0);
        assertEquals(NULL_INT,first(new DbIntArrayDirect()));
        assertEquals(NULL_INT,first(new DbIntArrayDirect(NULL_INT)));
        assertTrue(Math.abs(5-first(new DbIntArrayDirect(new int[]{5,NULL_INT,15})))==0.0);
        assertTrue(Math.abs(40-first(new DbIntArrayDirect((int)40)))==0.0);

        assertTrue(Math.abs(40-first(new int[]{40,50,60}))==0.0);
        assertEquals(NULL_INT,first(new int[]{}));
        assertEquals(NULL_INT,first(new int[]{NULL_INT}));
        assertTrue(Math.abs(5-first(new int[]{5,NULL_INT,15}))==0.0);
        assertTrue(Math.abs(40-first(new int[]{(int)40}))==0.0);
    }

    public void testNth(){
        assertEquals(NULL_INT, nth(-1,new DbIntArrayDirect(new int[]{40,50,60})));
        assertEquals((int)40, nth(0,new DbIntArrayDirect(new int[]{40,50,60})));
        assertEquals((int)50, nth(1,new DbIntArrayDirect(new int[]{40,50,60})));
        assertEquals((int)60, nth(2,new DbIntArrayDirect(new int[]{40,50,60})));
        assertEquals(NULL_INT, nth(10,new DbIntArrayDirect(new int[]{40,50,60})));

        assertEquals(NULL_INT, nth(-1,new int[]{40,50,60}));
        assertEquals((int)40, nth(0,new int[]{40,50,60}));
        assertEquals((int)50, nth(1,new int[]{40,50,60}));
        assertEquals((int)60, nth(2,new int[]{40,50,60}));
        assertEquals(NULL_INT, nth(10,new int[]{40,50,60}));
    }

    public void testCountDistinct() {
        assertEquals(NULL_LONG, countDistinct((DbIntArrayDirect)null));
        assertEquals(NULL_LONG, countDistinct((DbIntArrayDirect)null,true));
        assertEquals(0, countDistinct(new DbIntArrayDirect(new int[]{})));
        assertEquals(0, countDistinct(new DbIntArrayDirect(new int[]{NULL_INT})));
        assertEquals(1, countDistinct(new DbIntArrayDirect(new int[]{1})));
        assertEquals(2, countDistinct(new DbIntArrayDirect(new int[]{1,2,1,NULL_INT,NULL_INT})));
        assertEquals(2, countDistinct(new DbIntArrayDirect(new int[]{1,2,1,NULL_INT,NULL_INT}), false));
        assertEquals(3, countDistinct(new DbIntArrayDirect(new int[]{1,2,1,NULL_INT,NULL_INT}), true));

        assertEquals(NULL_LONG, countDistinct((int[])null));
        assertEquals(NULL_LONG, countDistinct((int[])null,true));
        assertEquals(0, countDistinct(new int[]{}));
        assertEquals(0, countDistinct(new int[]{NULL_INT}));
        assertEquals(1, countDistinct(new int[]{1}));
        assertEquals(2, countDistinct(new int[]{1,2,1,NULL_INT,NULL_INT}));
        assertEquals(2, countDistinct(new int[]{1,2,1,NULL_INT,NULL_INT}, false));
        assertEquals(3, countDistinct(new int[]{1,2,1,NULL_INT,NULL_INT}, true));
    }

    public void testDistinct() {
        assertEquals(null, distinct((DbIntArrayDirect)null));
        assertEquals(null, distinct((DbIntArrayDirect)null, true, true));
        assertEquals(new DbIntArrayDirect(), distinct(new DbIntArrayDirect(new int[]{})));
        assertEquals(new DbIntArrayDirect(), distinct(new DbIntArrayDirect(new int[]{NULL_INT})));
        assertEquals(new DbIntArrayDirect(new int[]{1}), distinct(new DbIntArrayDirect(new int[]{1})));
        assertEquals(new DbIntArrayDirect(new int[]{1,2}), distinct(new DbIntArrayDirect(new int[]{1,2,1,NULL_INT,NULL_INT})));
        assertEquals(new DbIntArrayDirect(new int[]{1,2}), distinct(new DbIntArrayDirect(new int[]{1,2,1,NULL_INT,NULL_INT}), false, false));
        assertEquals(new DbIntArrayDirect(new int[]{1,2,NULL_INT}), distinct(new DbIntArrayDirect(new int[]{1,2,1,NULL_INT,NULL_INT}), true, false));
        assertEquals(new DbIntArrayDirect(new int[]{1,2,3}), distinct(new DbIntArrayDirect(new int[]{3,1,2,1,NULL_INT,NULL_INT}), false, true));
        assertEquals(new DbIntArrayDirect(new int[]{1,2,3,4}), distinct(new DbIntArrayDirect(new int[]{3,1,2,4,1,NULL_INT,NULL_INT}), false, true));
        assertEquals(new DbIntArrayDirect(new int[]{NULL_INT,1,2,3,4}), distinct(new DbIntArrayDirect(new int[]{3,1,2,4,1,NULL_INT,NULL_INT}), true, true));

        assertEquals(null, distinct((int[])null));
        assertEquals(null, distinct((int[])null, true, true));
        assertEquals(new int[]{}, distinct(new int[]{}));
        assertEquals(new int[]{}, distinct(new int[]{NULL_INT}));
        assertEquals(new int[]{1}, distinct(new int[]{1}));
        assertEquals(new int[]{1,2}, distinct(new int[]{1,2,1,NULL_INT,NULL_INT}));
        assertEquals(new int[]{1,2}, distinct(new int[]{1,2,1,NULL_INT,NULL_INT}, false, false));
        assertEquals(new int[]{1,2,NULL_INT}, distinct(new int[]{1,2,1,NULL_INT,NULL_INT}, true, false));
        assertEquals(new int[]{1,2,3}, distinct(new int[]{3,1,2,1,NULL_INT,NULL_INT}, false, true));
        assertEquals(new int[]{1,2,3,4}, distinct(new int[]{3,1,2,4,1,NULL_INT,NULL_INT}, false, true));
        assertEquals(new int[]{NULL_INT,1,2,3,4}, distinct(new int[]{3,1,2,4,1,NULL_INT,NULL_INT}, true, true));
    }

    public void testVec(){
        assertEquals(new int[]{(int)1,(int)3,(int)5}, vec(new DbIntArrayDirect((int)1,(int)3,(int)5)));
    }

    public void testArray(){
        assertEquals(new DbIntArrayDirect((int)1,(int)3,(int)5), array(new int[]{(int)1,(int)3,(int)5}));
    }

    public void testIn(){
        assertTrue(in((int)1,(int)1,(int)2,(int)3));
        assertFalse(in((int)5,(int)1,(int)2,(int)3));
        assertFalse(in(NULL_INT,(int)1,(int)2,(int)3));
        assertTrue(in(NULL_INT,(int)1,(int)2,NULL_INT,(int)3));
    }

    public void testInRange(){
        assertTrue(inRange((int)2,(int)1,(int)3));
        assertTrue(inRange((int)1,(int)1,(int)3));
        assertFalse(inRange(NULL_INT,(int)1,(int)3));
        assertTrue(inRange((int)3,(int)1,(int)3));
        assertFalse(inRange((int)4,(int)1,(int)3));
    }

    public void testRepeat() {
        assertEquals(new int[]{5,5,5}, repeat((int) 5, 3));
        assertEquals(new int[]{}, repeat((int) 5, -3));
    }

    public void testEnlist() {
        assertEquals(new int[]{1, 11, 6}, enlist((int)1, (int)11, (int)6));
        assertEquals(new int[]{}, enlist((int[])(null)));
    }

    public void testConcat() {
        assertEquals(new int[]{}, concat((int[][])null));
        assertEquals(new int[]{1,2,3,4,5,6}, concat(new int[]{1,2}, new int[]{3}, new int[]{4,5,6}));
        assertEquals(new int[]{}, concat((int[])(null)));

        assertEquals(new int[]{}, concat((DbIntArray[])null));
        assertEquals(new int[]{1,2,3,4,5,6}, concat(new DbIntArrayDirect(new int[]{1,2}), new DbIntArrayDirect(new int[]{3}), new DbIntArrayDirect(new int[]{4,5,6})));
        assertEquals(new int[]{}, concat((DbIntArray) (null)));
    }

    public void testReverse() {
        assertEquals(new int[]{3,2,1}, reverse((int)1,(int)2,(int)3));
        assertEquals(null, reverse((int[])(null)));

        assertEquals(new int[]{3,2,1}, reverse(new DbIntArrayDirect(new int[]{1,2,3})));
        assertEquals(null, reverse((DbIntArray) (null)));
    }
}
