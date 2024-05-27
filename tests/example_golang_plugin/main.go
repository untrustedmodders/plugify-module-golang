package main

import "C"
import (
	"fmt"
	"plugify-plugin/cpptest"
	"plugify-plugin/plugify"
	"math"
	"reflect"
)

func assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func init() {
	plugify.OnPluginStart(func() {
		fmt.Println("Go:OnPluginStart")

		fmt.Println("=================================Go=======================================")

        cpptest.NoParamReturnVoid()
        assert(cpptest.NoParamReturnBool() == true, fmt.Sprintf("Expected NoParamReturnBool() to return true, but got %t", cpptest.NoParamReturnBool()))
        assert(cpptest.NoParamReturnChar8() == math.MaxInt8, fmt.Sprintf("Expected NoParamReturnChar8() to return %d, but got %d", math.MaxInt8, cpptest.NoParamReturnChar8()))
        assert(cpptest.NoParamReturnChar16() == math.MaxUint16, fmt.Sprintf("Expected NoParamReturnChar16() to return %d, but got %d", math.MaxUint16, cpptest.NoParamReturnChar16()))
        assert(cpptest.NoParamReturnInt8() == math.MaxInt8, fmt.Sprintf("Expected NoParamReturnInt8() to return %d, but got %d", math.MaxInt8, cpptest.NoParamReturnInt8()))
        assert(cpptest.NoParamReturnInt16() == math.MaxInt16, fmt.Sprintf("Expected NoParamReturnInt16() to return %d, but got %d", math.MaxInt16, cpptest.NoParamReturnInt16()))
        assert(cpptest.NoParamReturnInt32() == math.MaxInt32, fmt.Sprintf("Expected NoParamReturnInt32() to return %d, but got %d", math.MaxInt32, cpptest.NoParamReturnInt32()))
        assert(cpptest.NoParamReturnInt64() == math.MaxInt64, fmt.Sprintf("Expected NoParamReturnInt64() to return %d, but got %d", math.MaxInt64, cpptest.NoParamReturnInt64()))
        assert(cpptest.NoParamReturnUInt8() == math.MaxUint8, fmt.Sprintf("Expected NoParamReturnUInt8() to return %d, but got %d", math.MaxUint8, cpptest.NoParamReturnUInt8()))
        assert(cpptest.NoParamReturnUInt16() == math.MaxUint16, fmt.Sprintf("Expected NoParamReturnUInt16() to return %d, but got %d", math.MaxUint16, cpptest.NoParamReturnUInt16()))
        assert(cpptest.NoParamReturnUInt32() == math.MaxUint32, fmt.Sprintf("Expected NoParamReturnUInt32() to return %d, but got %d", math.MaxUint32, cpptest.NoParamReturnUInt32()))
        assert(cpptest.NoParamReturnUInt64() == math.MaxUint64, fmt.Sprintf("Expected NoParamReturnUInt64() to return math.MaxUint64, but got %d", cpptest.NoParamReturnUInt64()))

        assert(cpptest.NoParamReturnPtr64() == 1, fmt.Sprintf("Expected NoParamReturnPtr64() to return 1, but got %d", cpptest.NoParamReturnPtr64()))

        epsilon := 0.001

        assert(math.Abs(float64(cpptest.NoParamReturnFloat()-math.MaxFloat32)) < epsilon, fmt.Sprintf("Expected NoParamReturnFloat() to return %f, but got %f", math.MaxFloat32, cpptest.NoParamReturnFloat()))
        assert(math.Abs(cpptest.NoParamReturnDouble()-math.MaxFloat64) < epsilon, fmt.Sprintf("Expected NoParamReturnDouble() to return %f, but got %f", math.MaxFloat64, cpptest.NoParamReturnDouble()))

        assert(cpptest.NoParamReturnFunction() == 0, fmt.Sprintf("Expected NoParamReturnFunction() to return nil, but got %v", cpptest.NoParamReturnFunction()))

        assert(cpptest.NoParamReturnString() == "Hello World", fmt.Sprintf("Expected NoParamReturnString() to return 'Hello World', but got %s", cpptest.NoParamReturnString()))

        //expectedBoolSlice := []bool{true, false}
        //assert(reflect.DeepEqual(cpptest.NoParamReturnArrayBool(), expectedBoolSlice), fmt.Sprintf("Expected NoParamReturnArrayBool() to return %v, but got %v", expectedBoolSlice, cpptest.NoParamReturnArrayBool()))

        //expectedChar8Slice := []int8{'a', 'b', 'c'}
        //assert(reflect.DeepEqual(cpptest.NoParamReturnArrayChar8(), expectedChar8Slice), fmt.Sprintf("Expected NoParamReturnArrayChar8() to return %v, but got %v", expectedChar8Slice, cpptest.NoParamReturnArrayChar8()))

        //expectedChar16Slice := []uint16{'a', 'b', 'c', 'd'}
        //assert(reflect.DeepEqual(cpptest.NoParamReturnArrayChar16(), expectedChar16Slice), fmt.Sprintf("Expected NoParamReturnArrayChar16() to return %v, but got %v", expectedChar16Slice, cpptest.NoParamReturnArrayChar16()))

        expectedInt8Slice := []int8{-3, -2, -1, 0, 1}
        assert(reflect.DeepEqual(cpptest.NoParamReturnArrayInt8(), expectedInt8Slice), fmt.Sprintf("Expected NoParamReturnArrayInt8() to return %v, but got %v", expectedInt8Slice, cpptest.NoParamReturnArrayInt8()))

        expectedInt16Slice := []int16{-4, -3, -2, -1, 0, 1}
        assert(reflect.DeepEqual(cpptest.NoParamReturnArrayInt16(), expectedInt16Slice), fmt.Sprintf("Expected NoParamReturnArrayInt16() to return %v, but got %v", expectedInt16Slice, cpptest.NoParamReturnArrayInt16()))

        expectedInt32Slice := []int32{-5, -4, -3, -2, -1, 0, 1}
        assert(reflect.DeepEqual(cpptest.NoParamReturnArrayInt32(), expectedInt32Slice), fmt.Sprintf("Expected NoParamReturnArrayInt32() to return %v, but got %v", expectedInt32Slice, cpptest.NoParamReturnArrayInt32()))

        expectedInt64Slice := []int64{-6, -5, -4, -3, -2, -1, 0, 1}
        assert(reflect.DeepEqual(cpptest.NoParamReturnArrayInt64(), expectedInt64Slice), fmt.Sprintf("Expected NoParamReturnArrayInt64() to return %v, but got %v", expectedInt64Slice, cpptest.NoParamReturnArrayInt64()))

        expectedUInt8Slice := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8}
        assert(reflect.DeepEqual(cpptest.NoParamReturnArrayUInt8(), expectedUInt8Slice), fmt.Sprintf("Expected NoParamReturnArrayUInt8() to return %v, but got %v", expectedUInt8Slice, cpptest.NoParamReturnArrayUInt8()))

        expectedUInt16Slice := []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
        assert(reflect.DeepEqual(cpptest.NoParamReturnArrayUInt16(), expectedUInt16Slice), fmt.Sprintf("Expected NoParamReturnArrayUInt16() to return %v, but got %v", expectedUInt16Slice, cpptest.NoParamReturnArrayUInt16()))

        expectedUInt32Slice := []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
        assert(reflect.DeepEqual(cpptest.NoParamReturnArrayUInt32(), expectedUInt32Slice), fmt.Sprintf("Expected NoParamReturnArrayUInt32() to return %v, but got %v", expectedUInt32Slice, cpptest.NoParamReturnArrayUInt32()))

        expectedUInt64Slice := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
        assert(reflect.DeepEqual(cpptest.NoParamReturnArrayUInt64(), expectedUInt64Slice), fmt.Sprintf("Expected NoParamReturnArrayUInt64() to return %v, but got %v", expectedUInt64Slice, cpptest.NoParamReturnArrayUInt64()))

        expectedIntPtrSlice := []uintptr{0, 1, 2, 3}
        assert(reflect.DeepEqual(cpptest.NoParamReturnArrayPtr64(), expectedIntPtrSlice), fmt.Sprintf("Expected NoParamReturnArrayPtr64() to return %v, but got %v", expectedIntPtrSlice, cpptest.NoParamReturnArrayPtr64()))

        expectedFloatSlice := []float32{-12.34, 0.0, 12.34}
        assert(reflect.DeepEqual(cpptest.NoParamReturnArrayFloat(), expectedFloatSlice), fmt.Sprintf("Expected NoParamReturnArrayFloat() to return %v, but got %v", expectedFloatSlice, cpptest.NoParamReturnArrayFloat()))

        expectedDoubleSlice := []float64{-12.345, 0.0, 12.345}
        assert(reflect.DeepEqual(cpptest.NoParamReturnArrayDouble(), expectedDoubleSlice), fmt.Sprintf("Expected NoParamReturnArrayDouble() to return %v, but got %v", expectedDoubleSlice, cpptest.NoParamReturnArrayDouble()))

        expectedStringSlice := []string{"1st string", "2nd string", "3rd element string (Should be big enough to avoid small string optimization)"}
        assert(reflect.DeepEqual(cpptest.NoParamReturnArrayString(), expectedStringSlice), fmt.Sprintf("Expected NoParamReturnArrayString() to return %v, but got %v", expectedStringSlice, cpptest.NoParamReturnArrayString()))

        //Assertion for vector and matrix types could be added once the appropriate types are defined or imported.
        //assert(cpptest.NoParamReturnVector2() == new Vector2(1, 2), "NoParamReturnVector2() failed.");
        //assert(cpptest.NoParamReturnVector3() == new Vector3(1, 2, 3), "NoParamReturnVector3() failed.");
        //assert(cpptest.NoParamReturnVector4() == new Vector4(1, 2, 3, 4), "NoParamReturnVector4() failed.");
        //assert(cpptest.NoParamReturnMatrix4x4() == new Matrix4x4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16), "NoParamReturnMatrix4x4() failed.");
        // Vector tests are not included as there is no equivalent in Go

        intValue := int32(42)
        floatValue := float32(3.14)
        doubleValue := 6.28
        vector4Value := cpptest.Vector4{1.0, 2.0, 3.0, 4.0}
        longListValue := []int64{100, 200, 300}
        charValue := int8('A')
        stringValue := "Hello"
        floatValue2 := float32(2.71)
        shortValue := int16(10)
        var ptrValue uintptr // Provide appropriate value for pointer

        // Params (no refs)
        cpptest.Param1(intValue)
        cpptest.Param2(intValue, floatValue)
        cpptest.Param3(intValue, floatValue, doubleValue)
        cpptest.Param4(intValue, floatValue, doubleValue, vector4Value)
        cpptest.Param5(intValue, floatValue, doubleValue, vector4Value, longListValue)
        cpptest.Param6(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue)
        cpptest.Param7(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue)
        cpptest.Param8(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue, floatValue2)
        cpptest.Param9(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue, floatValue2, shortValue)
        cpptest.Param10(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue, floatValue2, shortValue, ptrValue)

        // Params (with refs)

        // ParamRef1
        cpptest.ParamRef1(&intValue)
        assert(intValue == 42, fmt.Sprintf("Expected intValue to be 42, but got %d", intValue))

        // ParamRef2
        cpptest.ParamRef2(&intValue, &floatValue)
        assert(intValue == 10, fmt.Sprintf("Expected intValue to be 10, but got %d", intValue))
        assert(math.Abs(float64(floatValue)-3.14) < epsilon, fmt.Sprintf("Expected floatValue to be approximately 3.14, but got %f", floatValue))

        // ParamRef3
        cpptest.ParamRef3(&intValue, &floatValue, &doubleValue)
        assert(intValue == -20, fmt.Sprintf("Expected intValue to be -20, but got %d", intValue))
        assert(math.Abs(float64(floatValue)-2.718) < epsilon, fmt.Sprintf("Expected floatValue to be approximately 2.718, but got %f", floatValue))
        assert(math.Abs(doubleValue-3.14159) < epsilon, fmt.Sprintf("Expected doubleValue to be approximately 3.14159, but got %f", doubleValue))

        // ParamRef4
        cpptest.ParamRef4(&intValue, &floatValue, &doubleValue, &vector4Value)
        assert(intValue == 100, fmt.Sprintf("Expected intValue to be 100, but got %d", intValue))
        assert(math.Abs(float64(floatValue)+5.55) < epsilon, fmt.Sprintf("Expected floatValue to be approximately -5.55, but got %f", floatValue))
        assert(math.Abs(doubleValue-1.618) < epsilon, fmt.Sprintf("Expected doubleValue to be approximately 1.618, but got %f", doubleValue))
        assert(vector4Value == cpptest.Vector4{1, 2, 3, 4}, fmt.Sprintf("Expected vector4Value to be (1, 2, 3, 4), but got %v", vector4Value))

        // ParamRef5
        cpptest.ParamRef5(&intValue, &floatValue, &doubleValue, &vector4Value, &longListValue)
        assert(intValue == 500, fmt.Sprintf("Expected intValue to be 500, but got %d", intValue))
        assert(math.Abs(float64(floatValue)+10.5) < epsilon, fmt.Sprintf("Expected floatValue to be approximately -10.5, but got %f", floatValue))
        assert(math.Abs(doubleValue-2.71828) < epsilon, fmt.Sprintf("Expected doubleValue to be approximately 2.71828, but got %f", doubleValue))
        assert(vector4Value == cpptest.Vector4{-1, -2, -3, -4}, fmt.Sprintf("Expected vector4Value to be (-1, -2, -3, -4), but got %v", vector4Value))
        assert(reflect.DeepEqual(longListValue, []int64{-6, -5, -4, -3, -2, -1, 0, 1}), fmt.Sprintf("Expected longListValue to be (-6, -5, -4, -3, -2, -1), but got %v", longListValue))

        // ParamRef6
        cpptest.ParamRef6(&intValue, &floatValue, &doubleValue, &vector4Value, &longListValue, &charValue)
        assert(intValue == 750, fmt.Sprintf("Expected intValue to be 750, but got %d", intValue))
        assert(math.Abs(float64(floatValue)-20.0) < epsilon, fmt.Sprintf("Expected floatValue to be approximately 20.0, but got %f", floatValue))
        assert(math.Abs(doubleValue-1.23456) < epsilon, fmt.Sprintf("Expected doubleValue to be approximately 1.23456, but got %f", doubleValue))
        assert(vector4Value == cpptest.Vector4{10, 20, 30, 40}, fmt.Sprintf("Expected vector4Value to be (10, 20, 30, 40), but got %v", vector4Value))
        assert(reflect.DeepEqual(longListValue, []int64{-6, -5, -4}), fmt.Sprintf("Expected longListValue to be (-6, -5, -4), but got %v", longListValue))
        assert(charValue == 'Z', fmt.Sprintf("Expected charValue to be 'Z', but got %c", charValue))

        // ParamRef7
        cpptest.ParamRef7(&intValue, &floatValue, &doubleValue, &vector4Value, &longListValue, &charValue, &stringValue)
        assert(intValue == -1000, fmt.Sprintf("Expected intValue to be -1000, but got %d", intValue))
        assert(math.Abs(float64(floatValue)-3.0) < epsilon, fmt.Sprintf("Expected floatValue to be approximately 3.0, but got %f", floatValue))
        assert(doubleValue == -1, fmt.Sprintf("Expected doubleValue to be -1, but got %f", doubleValue))
        assert(vector4Value == cpptest.Vector4{100, 200, 300, 400}, fmt.Sprintf("Expected vector4Value to be (100, 200, 300, 400), but got %v", vector4Value))
        assert(reflect.DeepEqual(longListValue, []int64{-6, -5, -4, -3}), fmt.Sprintf("Expected longListValue to be (-6, -5, -4, -3), but got %v", longListValue))
        assert(charValue == 'X', fmt.Sprintf("Expected charValue to be 'X', but got %c", charValue))
        assert(stringValue == "Hello, World!", fmt.Sprintf("Expected stringValue to be 'Hello, World!', but got %s", stringValue))

        // ParamRef8
        cpptest.ParamRef8(&intValue, &floatValue, &doubleValue, &vector4Value, &longListValue, &charValue, &stringValue, &floatValue2)
        assert(intValue == 999, fmt.Sprintf("Expected intValue to be 999, but got %d", intValue))
        assert(math.Abs(float64(floatValue)+7.5) < epsilon, fmt.Sprintf("Expected floatValue to be approximately -7.5, but got %f", floatValue))
        assert(math.Abs(doubleValue-0.123456) < 0.000001, fmt.Sprintf("Expected doubleValue to be approximately 0.123456, but got %f", doubleValue))
        assert(vector4Value == cpptest.Vector4{-100, -200, -300, -400}, fmt.Sprintf("Expected vector4Value to be (-100, -200, -300, -400), but got %v", vector4Value))
        assert(reflect.DeepEqual(longListValue, []int64{-6, -5, -4, -3, -2, -1}), fmt.Sprintf("Expected longListValue to be (-6, -5, -4, -3, -2, -1), but got %v", longListValue))
        assert(charValue == 'Y', fmt.Sprintf("Expected charValue to be 'Y', but got %c", charValue))
        assert(stringValue == "Goodbye, World!", fmt.Sprintf("Expected stringValue to be 'Goodbye, World!', but got %s", stringValue))
        assert(math.Abs(float64(floatValue2)-99.99) < epsilon, fmt.Sprintf("Expected floatValue2 to be approximately 99.99, but got %f", floatValue2))

        // ParamRef9
        cpptest.ParamRef9(&intValue, &floatValue, &doubleValue, &vector4Value, &longListValue, &charValue, &stringValue, &floatValue2, &shortValue)
        assert(intValue == -1234, fmt.Sprintf("Expected intValue to be -1234, but got %d", intValue))
        assert(math.Abs(float64(floatValue)-123.45) < epsilon, fmt.Sprintf("Expected floatValue to be approximately 123.45, but got %f", floatValue))
        assert(math.Abs(doubleValue+678.9) < epsilon, fmt.Sprintf("Expected doubleValue to be approximately -678.9, but got %f", doubleValue))
        assert(vector4Value == cpptest.Vector4{987.65, 432.1, 123.456, 789.123}, fmt.Sprintf("Expected vector4Value to be (987.65 , 432.1, 123.456, 789.123), but got %v", vector4Value))
        assert(reflect.DeepEqual(longListValue, []int64{-6, -5, -4, -3, -2, -1, 0, 1, 5, 9}), fmt.Sprintf("Expected longListValue to be (-6, -5, -4, -3, -2, -1, 0, 1, 5, 9), but got %v", longListValue))
        assert(charValue == 'A', fmt.Sprintf("Expected charValue to be 'A', but got %c", charValue))
        assert(stringValue == "Testing, 1 2 3", fmt.Sprintf("Expected stringValue to be 'Testing, 1 2 3', but got %s", stringValue))
        assert(math.Abs(float64(floatValue2)+987.654) < epsilon, fmt.Sprintf("Expected floatValue2 to be approximately -987.654, but got %f", floatValue2))
        assert(shortValue == 42, fmt.Sprintf("Expected shortValue to be 42, but got %d", shortValue))

        // ParamRef10
        cpptest.ParamRef10(&intValue, &floatValue, &doubleValue, &vector4Value, &longListValue, &charValue, &stringValue, &floatValue2, &shortValue, &ptrValue)
        assert(intValue == 987, fmt.Sprintf("Expected intValue to be 987, but got %d", intValue))
        assert(math.Abs(float64(floatValue)+0.123) < epsilon, fmt.Sprintf("Expected floatValue to be approximately -0.123, but got %f", floatValue))
        assert(math.Abs(doubleValue-456.789) < 0.000001, fmt.Sprintf("Expected doubleValue to be approximately 456.789, but got %f", doubleValue))
        assert(vector4Value == cpptest.Vector4{-123.456, 0.987, 654.321, -789.123}, fmt.Sprintf("Expected vector4Value to be (-123.456, 0.987, 654.321, -789.123), but got %v", vector4Value))
        assert(reflect.DeepEqual(longListValue, []int64{-6, -5, -4, -3, -2, -1, 0, 1, 5, 9, 4, -7}), fmt.Sprintf("Expected longListValue to be (-6, -5, -4, -3, -2, -1, 0, 1, 5, 9, 4, -7), but got %v", longListValue))
        assert(charValue == 'B', fmt.Sprintf("Expected charValue to be 'B', but got %c", charValue))
        assert(stringValue == "Another string", fmt.Sprintf("Expected stringValue to be 'Another string', but got %s", stringValue))
        assert(math.Abs(float64(floatValue2)-3.141592) < 0.000001, fmt.Sprintf("Expected floatValue2 to be approximately 3.141592, but got %f", floatValue2))
        assert(shortValue == -32768, fmt.Sprintf("Expected shortValue to be -32768, but got %d", shortValue))
        assert(ptrValue == uintptr(0), fmt.Sprintf("Expected ptrValue to be 0, but got %d", ptrValue))

        // Initialize arrays
        boolArray := []bool{true, false, true}
        char8Array := []int8{'A', 'B', 'C'}
        char16Array := []uint16{'A', 'B', 'C'}
        sbyteArray := []int8{-1, -2, -3}
        shortArray := []int16{10, 20, 30}
        intArray := []int32{100, 200, 300}
        longArray := []int64{1000, 2000, 3000}
        byteArray := []uint8{1, 2, 3}
        ushortArray := []uint16{1000, 2000, 3000}
        uintArray := []uint32{10000, 20000, 30000}
        ulongArray := []uint64{100000, 200000, 300000}
        intPtrArray := []uintptr{0, 1, 2}
        floatArray := []float32{1.1, 2.2, 3.3}
        doubleArray := []float64{1.1, 2.2, 3.3}
        stringArray := []string{"Hello", "World", "!"}

        // Call function ParamRefVectors and print how values change
        cpptest.ParamRefVectors(&boolArray, &char8Array, &char16Array, &sbyteArray, &shortArray, &intArray, &longArray,
            &byteArray, &ushortArray, &uintArray, &ulongArray, &intPtrArray, &floatArray, &doubleArray, &stringArray)

        assert(reflect.DeepEqual(boolArray, []bool{true}), fmt.Sprintf("Expected boolArray to be (true), but got %v", boolArray))
        assert(reflect.DeepEqual(char8Array, []int8{'a', 'b', 'c'}), fmt.Sprintf("Expected charArray to be ('a', 'b', 'c'), but got %v", char8Array))
        assert(reflect.DeepEqual(char16Array, []uint16{'a', 'b', 'c'}), fmt.Sprintf("Expected char16Array to be ('a', 'b', 'c'), but got %v", char16Array))
        assert(reflect.DeepEqual(sbyteArray, []int8{-3, -2, -1, 0, 1, 2, 3}), fmt.Sprintf("Expected sbyteArray to be (-3, -2, -1, 0, 1, 2, 3), but got %v", sbyteArray))
        assert(reflect.DeepEqual(shortArray, []int16{-4, -3, -2, -1, 0, 1, 2, 3, 4}), fmt.Sprintf("Expected shortArray to be (-4, -3, -2, -1, 0, 1, 2, 3, 4), but got %v", shortArray))
        assert(reflect.DeepEqual(intArray, []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}), fmt.Sprintf("Expected intArray to be (-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5), but got %v", intArray))
        assert(reflect.DeepEqual(longArray, []int64{-6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6}), fmt.Sprintf("Expected longArray to be (-6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6), but got %v", longArray))
        assert(reflect.DeepEqual(byteArray, []uint8{0, 1, 2, 3, 4, 5, 6, 7}), fmt.Sprintf("Expected byteArray to be (0, 1, 2, 3, 4, 5, 6, 7), but got %v", byteArray))
        assert(reflect.DeepEqual(ushortArray, []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8}), fmt.Sprintf("Expected ushortArray to be (0, 1, 2, 3, 4, 5, 6, 7, 8), but got %v", ushortArray))
        assert(reflect.DeepEqual(uintArray, []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), fmt.Sprintf("Expected uintArray to be (0, 1, 2, 3, 4, 5, 6, 7, 8, 9), but got %v", uintArray))
        assert(reflect.DeepEqual(ulongArray, []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), fmt.Sprintf("Expected ulongArray to be (0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10), but got %v", ulongArray))
        assert(reflect.DeepEqual(intPtrArray, []uintptr{0, 1}), fmt.Sprintf("Expected intPtrArray to be (IntPtr.Zero, IntPtr.Zero + 1), but got %v", intPtrArray))
        assert(reflect.DeepEqual(floatArray, []float32{-12.34, 0.0, 12.34}), fmt.Sprintf("Expected floatArray to be (-12.34, 0.0, 12.34), but got %v", floatArray))
        assert(reflect.DeepEqual(doubleArray, []float64{-12.345, 0.0, 12.345}), fmt.Sprintf("Expected doubleArray to be (-12.345, 0.0, 12.345), but got %v", doubleArray))
        assert(reflect.DeepEqual(stringArray, []string{"Hello", "World", "OpenAI"}), fmt.Sprintf("Expected stringArray to be ('Hello', 'World', 'OpenAI'), but got %v", stringArray))

        returnValue := cpptest.ParamAllPrimitives(
            boolArray[0], char16Array[0], sbyteArray[0], shortArray[0], intArray[0], longArray[0],
            byteArray[0], ushortArray[0], uintArray[0], ulongArray[0], intPtrArray[0], floatArray[0], doubleArray[0],
        )

        assert(returnValue == 56, fmt.Sprintf("Expected return value to be 56, but got %d", returnValue))

        fmt.Println("Go:OnPluginStart - done")
	})

	plugify.OnPluginEnd(func() {
		fmt.Println("Go:OnPluginEnd")
	})
}

//export SayHello
func SayHello(count int64) bool {
	if count > 10 {
		return false
	}

	fmt.Printf("Hello %d times!\n", count)

	return true
}

func main() {}
