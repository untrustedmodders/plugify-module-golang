package main

/*
typedef struct { float x, y; } Vector2;
typedef struct { float x, y, z; } Vector3;
typedef struct { float x, y, z, w; } Vector4;
typedef struct { float m00, m01, m02, m03, m10, m11, m12, m13, m20, m21, m22, m23, m30, m31, m32, m33; } Matrix4x4;
*/
import "C"
import (
	"fmt"
	"math"
	"reflect"
	"github.com/untrustedmodders/go-plugify"
	"plugify-plugin/example_cpp_plugin"
)

func assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func init() {
	plugify.OnPluginStart(func() {
		fmt.Println("Go: OnPluginStart")

		fmt.Println("=================================Go=======================================")

		example_cpp_plugin.NoParamReturnVoid()
		assert(example_cpp_plugin.NoParamReturnBool() == true, fmt.Sprintf("Expected NoParamReturnBool() to return true, but got %t", example_cpp_plugin.NoParamReturnBool()))
		assert(example_cpp_plugin.NoParamReturnChar8() == math.MaxInt8, fmt.Sprintf("Expected NoParamReturnChar8() to return %d, but got %d", math.MaxInt8, example_cpp_plugin.NoParamReturnChar8()))
		assert(example_cpp_plugin.NoParamReturnChar16() == math.MaxUint16, fmt.Sprintf("Expected NoParamReturnChar16() to return %d, but got %d", math.MaxUint16, example_cpp_plugin.NoParamReturnChar16()))
		assert(example_cpp_plugin.NoParamReturnInt8() == math.MaxInt8, fmt.Sprintf("Expected NoParamReturnInt8() to return %d, but got %d", math.MaxInt8, example_cpp_plugin.NoParamReturnInt8()))
		assert(example_cpp_plugin.NoParamReturnInt16() == math.MaxInt16, fmt.Sprintf("Expected NoParamReturnInt16() to return %d, but got %d", math.MaxInt16, example_cpp_plugin.NoParamReturnInt16()))
		assert(example_cpp_plugin.NoParamReturnInt32() == math.MaxInt32, fmt.Sprintf("Expected NoParamReturnInt32() to return %d, but got %d", math.MaxInt32, example_cpp_plugin.NoParamReturnInt32()))
		assert(example_cpp_plugin.NoParamReturnInt64() == math.MaxInt64, fmt.Sprintf("Expected NoParamReturnInt64() to return %d, but got %d", math.MaxInt64, example_cpp_plugin.NoParamReturnInt64()))
		assert(example_cpp_plugin.NoParamReturnUInt8() == math.MaxUint8, fmt.Sprintf("Expected NoParamReturnUInt8() to return %d, but got %d", math.MaxUint8, example_cpp_plugin.NoParamReturnUInt8()))
		assert(example_cpp_plugin.NoParamReturnUInt16() == math.MaxUint16, fmt.Sprintf("Expected NoParamReturnUInt16() to return %d, but got %d", math.MaxUint16, example_cpp_plugin.NoParamReturnUInt16()))
		assert(example_cpp_plugin.NoParamReturnUInt32() == math.MaxUint32, fmt.Sprintf("Expected NoParamReturnUInt32() to return %d, but got %d", math.MaxUint32, example_cpp_plugin.NoParamReturnUInt32()))
		assert(example_cpp_plugin.NoParamReturnUInt64() == math.MaxUint64, fmt.Sprintf("Expected NoParamReturnUInt64() to return math.MaxUint64, but got %d", example_cpp_plugin.NoParamReturnUInt64()))

		assert(example_cpp_plugin.NoParamReturnPtr64() == 1, fmt.Sprintf("Expected NoParamReturnPtr64() to return 1, but got %d", example_cpp_plugin.NoParamReturnPtr64()))

		epsilon := 0.001

		assert(math.Abs(float64(example_cpp_plugin.NoParamReturnFloat()-math.MaxFloat32)) < epsilon, fmt.Sprintf("Expected NoParamReturnFloat() to return %f, but got %f", math.MaxFloat32, example_cpp_plugin.NoParamReturnFloat()))
		assert(math.Abs(example_cpp_plugin.NoParamReturnDouble()-math.MaxFloat64) < epsilon, fmt.Sprintf("Expected NoParamReturnDouble() to return %f, but got %f", math.MaxFloat64, example_cpp_plugin.NoParamReturnDouble()))

		assert(example_cpp_plugin.NoParamReturnFunction() == 0, fmt.Sprintf("Expected NoParamReturnFunction() to return nil, but got %v", example_cpp_plugin.NoParamReturnFunction()))

		assert(example_cpp_plugin.NoParamReturnString() == "Hello World", fmt.Sprintf("Expected NoParamReturnString() to return 'Hello World', but got %s", example_cpp_plugin.NoParamReturnString()))

		expectedBoolSlice := []bool{true, false}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayBool(), expectedBoolSlice), fmt.Sprintf("Expected NoParamReturnArrayBool() to return %v, but got %v", expectedBoolSlice, example_cpp_plugin.NoParamReturnArrayBool()))

		expectedChar8Slice := []int8{'a', 'b', 'c'}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayChar8(), expectedChar8Slice), fmt.Sprintf("Expected NoParamReturnArrayChar8() to return %v, but got %v", expectedChar8Slice, example_cpp_plugin.NoParamReturnArrayChar8()))

		expectedChar16Slice := []uint16{'a', 'b', 'c', 'd'}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayChar16(), expectedChar16Slice), fmt.Sprintf("Expected NoParamReturnArrayChar16() to return %v, but got %v", expectedChar16Slice, example_cpp_plugin.NoParamReturnArrayChar16()))

		expectedInt8Slice := []int8{-3, -2, -1, 0, 1}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayInt8(), expectedInt8Slice), fmt.Sprintf("Expected NoParamReturnArrayInt8() to return %v, but got %v", expectedInt8Slice, example_cpp_plugin.NoParamReturnArrayInt8()))

		expectedInt16Slice := []int16{-4, -3, -2, -1, 0, 1}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayInt16(), expectedInt16Slice), fmt.Sprintf("Expected NoParamReturnArrayInt16() to return %v, but got %v", expectedInt16Slice, example_cpp_plugin.NoParamReturnArrayInt16()))

		expectedInt32Slice := []int32{-5, -4, -3, -2, -1, 0, 1}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayInt32(), expectedInt32Slice), fmt.Sprintf("Expected NoParamReturnArrayInt32() to return %v, but got %v", expectedInt32Slice, example_cpp_plugin.NoParamReturnArrayInt32()))

		expectedInt64Slice := []int64{-6, -5, -4, -3, -2, -1, 0, 1}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayInt64(), expectedInt64Slice), fmt.Sprintf("Expected NoParamReturnArrayInt64() to return %v, but got %v", expectedInt64Slice, example_cpp_plugin.NoParamReturnArrayInt64()))

		expectedUInt8Slice := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayUInt8(), expectedUInt8Slice), fmt.Sprintf("Expected NoParamReturnArrayUInt8() to return %v, but got %v", expectedUInt8Slice, example_cpp_plugin.NoParamReturnArrayUInt8()))

		expectedUInt16Slice := []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayUInt16(), expectedUInt16Slice), fmt.Sprintf("Expected NoParamReturnArrayUInt16() to return %v, but got %v", expectedUInt16Slice, example_cpp_plugin.NoParamReturnArrayUInt16()))

		expectedUInt32Slice := []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayUInt32(), expectedUInt32Slice), fmt.Sprintf("Expected NoParamReturnArrayUInt32() to return %v, but got %v", expectedUInt32Slice, example_cpp_plugin.NoParamReturnArrayUInt32()))

		expectedUInt64Slice := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayUInt64(), expectedUInt64Slice), fmt.Sprintf("Expected NoParamReturnArrayUInt64() to return %v, but got %v", expectedUInt64Slice, example_cpp_plugin.NoParamReturnArrayUInt64()))

		expectedIntPtrSlice := []uintptr{0, 1, 2, 3}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayPtr64(), expectedIntPtrSlice), fmt.Sprintf("Expected NoParamReturnArrayPtr64() to return %v, but got %v", expectedIntPtrSlice, example_cpp_plugin.NoParamReturnArrayPtr64()))

		expectedFloatSlice := []float32{-12.34, 0.0, 12.34}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayFloat(), expectedFloatSlice), fmt.Sprintf("Expected NoParamReturnArrayFloat() to return %v, but got %v", expectedFloatSlice, example_cpp_plugin.NoParamReturnArrayFloat()))

		expectedDoubleSlice := []float64{-12.345, 0.0, 12.345}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayDouble(), expectedDoubleSlice), fmt.Sprintf("Expected NoParamReturnArrayDouble() to return %v, but got %v", expectedDoubleSlice, example_cpp_plugin.NoParamReturnArrayDouble()))

		expectedStringSlice := []string{"1st string", "2nd string", "3rd element string (Should be big enough to avoid small string optimization)"}
		assert(reflect.DeepEqual(example_cpp_plugin.NoParamReturnArrayString(), expectedStringSlice), fmt.Sprintf("Expected NoParamReturnArrayString() to return %v, but got %v", expectedStringSlice, example_cpp_plugin.NoParamReturnArrayString()))

		//Assertion for vector and matrix types could be added once the appropriate types are defined or imported.
		//assert(example_cpp_plugin.NoParamReturnVector2() == new Vector2(1, 2), "NoParamReturnVector2() failed.");
		//assert(example_cpp_plugin.NoParamReturnVector3() == new Vector3(1, 2, 3), "NoParamReturnVector3() failed.");
		//assert(example_cpp_plugin.NoParamReturnVector4() == new Vector4(1, 2, 3, 4), "NoParamReturnVector4() failed.");
		//assert(example_cpp_plugin.NoParamReturnMatrix4x4() == new Matrix4x4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16), "NoParamReturnMatrix4x4() failed.");
		// Vector tests are not included as there is no equivalent in Go

		intValue := int32(42)
		floatValue := float32(3.14)
		doubleValue := 6.28
		vector4Value := example_cpp_plugin.Vector4{1.0, 2.0, 3.0, 4.0}
		longListValue := []int64{100, 200, 300}
		charValue := int8('A')
		stringValue := "Hello"
		floatValue2 := float32(2.71)
		shortValue := int16(10)
		var ptrValue uintptr // Provide appropriate value for pointer

		// Params (no refs)
		example_cpp_plugin.Param1(intValue)
		example_cpp_plugin.Param2(intValue, floatValue)
		example_cpp_plugin.Param3(intValue, floatValue, doubleValue)
		example_cpp_plugin.Param4(intValue, floatValue, doubleValue, vector4Value)
		example_cpp_plugin.Param5(intValue, floatValue, doubleValue, vector4Value, longListValue)
		example_cpp_plugin.Param6(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue)
		example_cpp_plugin.Param7(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue)
		example_cpp_plugin.Param8(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue, floatValue2)
		example_cpp_plugin.Param9(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue, floatValue2, shortValue)
		example_cpp_plugin.Param10(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue, floatValue2, shortValue, ptrValue)

		// Params (with refs)

		// ParamRef1
		example_cpp_plugin.ParamRef1(&intValue)
		assert(intValue == 42, fmt.Sprintf("Expected intValue to be 42, but got %d", intValue))

		// ParamRef2
		example_cpp_plugin.ParamRef2(&intValue, &floatValue)
		assert(intValue == 10, fmt.Sprintf("Expected intValue to be 10, but got %d", intValue))
		assert(math.Abs(float64(floatValue)-3.14) < epsilon, fmt.Sprintf("Expected floatValue to be approximately 3.14, but got %f", floatValue))

		// ParamRef3
		example_cpp_plugin.ParamRef3(&intValue, &floatValue, &doubleValue)
		assert(intValue == -20, fmt.Sprintf("Expected intValue to be -20, but got %d", intValue))
		assert(math.Abs(float64(floatValue)-2.718) < epsilon, fmt.Sprintf("Expected floatValue to be approximately 2.718, but got %f", floatValue))
		assert(math.Abs(doubleValue-3.14159) < epsilon, fmt.Sprintf("Expected doubleValue to be approximately 3.14159, but got %f", doubleValue))

		// ParamRef4
		example_cpp_plugin.ParamRef4(&intValue, &floatValue, &doubleValue, &vector4Value)
		assert(intValue == 100, fmt.Sprintf("Expected intValue to be 100, but got %d", intValue))
		assert(math.Abs(float64(floatValue)+5.55) < epsilon, fmt.Sprintf("Expected floatValue to be approximately -5.55, but got %f", floatValue))
		assert(math.Abs(doubleValue-1.618) < epsilon, fmt.Sprintf("Expected doubleValue to be approximately 1.618, but got %f", doubleValue))
		assert(vector4Value == example_cpp_plugin.Vector4{1, 2, 3, 4}, fmt.Sprintf("Expected vector4Value to be (1, 2, 3, 4), but got %v", vector4Value))

		// ParamRef5
		example_cpp_plugin.ParamRef5(&intValue, &floatValue, &doubleValue, &vector4Value, &longListValue)
		assert(intValue == 500, fmt.Sprintf("Expected intValue to be 500, but got %d", intValue))
		assert(math.Abs(float64(floatValue)+10.5) < epsilon, fmt.Sprintf("Expected floatValue to be approximately -10.5, but got %f", floatValue))
		assert(math.Abs(doubleValue-2.71828) < epsilon, fmt.Sprintf("Expected doubleValue to be approximately 2.71828, but got %f", doubleValue))
		assert(vector4Value == example_cpp_plugin.Vector4{-1, -2, -3, -4}, fmt.Sprintf("Expected vector4Value to be (-1, -2, -3, -4), but got %v", vector4Value))
		assert(reflect.DeepEqual(longListValue, []int64{-6, -5, -4, -3, -2, -1, 0, 1}), fmt.Sprintf("Expected longListValue to be (-6, -5, -4, -3, -2, -1), but got %v", longListValue))

		// ParamRef6
		example_cpp_plugin.ParamRef6(&intValue, &floatValue, &doubleValue, &vector4Value, &longListValue, &charValue)
		assert(intValue == 750, fmt.Sprintf("Expected intValue to be 750, but got %d", intValue))
		assert(math.Abs(float64(floatValue)-20.0) < epsilon, fmt.Sprintf("Expected floatValue to be approximately 20.0, but got %f", floatValue))
		assert(math.Abs(doubleValue-1.23456) < epsilon, fmt.Sprintf("Expected doubleValue to be approximately 1.23456, but got %f", doubleValue))
		assert(vector4Value == example_cpp_plugin.Vector4{10, 20, 30, 40}, fmt.Sprintf("Expected vector4Value to be (10, 20, 30, 40), but got %v", vector4Value))
		assert(reflect.DeepEqual(longListValue, []int64{-6, -5, -4}), fmt.Sprintf("Expected longListValue to be (-6, -5, -4), but got %v", longListValue))
		assert(charValue == 'Z', fmt.Sprintf("Expected charValue to be 'Z', but got %c", charValue))

		// ParamRef7
		example_cpp_plugin.ParamRef7(&intValue, &floatValue, &doubleValue, &vector4Value, &longListValue, &charValue, &stringValue)
		assert(intValue == -1000, fmt.Sprintf("Expected intValue to be -1000, but got %d", intValue))
		assert(math.Abs(float64(floatValue)-3.0) < epsilon, fmt.Sprintf("Expected floatValue to be approximately 3.0, but got %f", floatValue))
		assert(doubleValue == -1, fmt.Sprintf("Expected doubleValue to be -1, but got %f", doubleValue))
		assert(vector4Value == example_cpp_plugin.Vector4{100, 200, 300, 400}, fmt.Sprintf("Expected vector4Value to be (100, 200, 300, 400), but got %v", vector4Value))
		assert(reflect.DeepEqual(longListValue, []int64{-6, -5, -4, -3}), fmt.Sprintf("Expected longListValue to be (-6, -5, -4, -3), but got %v", longListValue))
		assert(charValue == 'X', fmt.Sprintf("Expected charValue to be 'X', but got %c", charValue))
		assert(stringValue == "Hello, World!", fmt.Sprintf("Expected stringValue to be 'Hello, World!', but got %s", stringValue))

		// ParamRef8
		example_cpp_plugin.ParamRef8(&intValue, &floatValue, &doubleValue, &vector4Value, &longListValue, &charValue, &stringValue, &floatValue2)
		assert(intValue == 999, fmt.Sprintf("Expected intValue to be 999, but got %d", intValue))
		assert(math.Abs(float64(floatValue)+7.5) < epsilon, fmt.Sprintf("Expected floatValue to be approximately -7.5, but got %f", floatValue))
		assert(math.Abs(doubleValue-0.123456) < 0.000001, fmt.Sprintf("Expected doubleValue to be approximately 0.123456, but got %f", doubleValue))
		assert(vector4Value == example_cpp_plugin.Vector4{-100, -200, -300, -400}, fmt.Sprintf("Expected vector4Value to be (-100, -200, -300, -400), but got %v", vector4Value))
		assert(reflect.DeepEqual(longListValue, []int64{-6, -5, -4, -3, -2, -1}), fmt.Sprintf("Expected longListValue to be (-6, -5, -4, -3, -2, -1), but got %v", longListValue))
		assert(charValue == 'Y', fmt.Sprintf("Expected charValue to be 'Y', but got %c", charValue))
		assert(stringValue == "Goodbye, World!", fmt.Sprintf("Expected stringValue to be 'Goodbye, World!', but got %s", stringValue))
		assert(math.Abs(float64(floatValue2)-99.99) < epsilon, fmt.Sprintf("Expected floatValue2 to be approximately 99.99, but got %f", floatValue2))

		// ParamRef9
		example_cpp_plugin.ParamRef9(&intValue, &floatValue, &doubleValue, &vector4Value, &longListValue, &charValue, &stringValue, &floatValue2, &shortValue)
		assert(intValue == -1234, fmt.Sprintf("Expected intValue to be -1234, but got %d", intValue))
		assert(math.Abs(float64(floatValue)-123.45) < epsilon, fmt.Sprintf("Expected floatValue to be approximately 123.45, but got %f", floatValue))
		assert(math.Abs(doubleValue+678.9) < epsilon, fmt.Sprintf("Expected doubleValue to be approximately -678.9, but got %f", doubleValue))
		assert(vector4Value == example_cpp_plugin.Vector4{987.65, 432.1, 123.456, 789.123}, fmt.Sprintf("Expected vector4Value to be (987.65 , 432.1, 123.456, 789.123), but got %v", vector4Value))
		assert(reflect.DeepEqual(longListValue, []int64{-6, -5, -4, -3, -2, -1, 0, 1, 5, 9}), fmt.Sprintf("Expected longListValue to be (-6, -5, -4, -3, -2, -1, 0, 1, 5, 9), but got %v", longListValue))
		assert(charValue == 'A', fmt.Sprintf("Expected charValue to be 'A', but got %c", charValue))
		assert(stringValue == "Testing, 1 2 3", fmt.Sprintf("Expected stringValue to be 'Testing, 1 2 3', but got %s", stringValue))
		assert(math.Abs(float64(floatValue2)+987.654) < epsilon, fmt.Sprintf("Expected floatValue2 to be approximately -987.654, but got %f", floatValue2))
		assert(shortValue == 42, fmt.Sprintf("Expected shortValue to be 42, but got %d", shortValue))

		// ParamRef10
		example_cpp_plugin.ParamRef10(&intValue, &floatValue, &doubleValue, &vector4Value, &longListValue, &charValue, &stringValue, &floatValue2, &shortValue, &ptrValue)
		assert(intValue == 987, fmt.Sprintf("Expected intValue to be 987, but got %d", intValue))
		assert(math.Abs(float64(floatValue)+0.123) < epsilon, fmt.Sprintf("Expected floatValue to be approximately -0.123, but got %f", floatValue))
		assert(math.Abs(doubleValue-456.789) < 0.000001, fmt.Sprintf("Expected doubleValue to be approximately 456.789, but got %f", doubleValue))
		assert(vector4Value == example_cpp_plugin.Vector4{-123.456, 0.987, 654.321, -789.123}, fmt.Sprintf("Expected vector4Value to be (-123.456, 0.987, 654.321, -789.123), but got %v", vector4Value))
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
		example_cpp_plugin.ParamRefVectors(&boolArray, &char8Array, &char16Array, &sbyteArray, &shortArray, &intArray, &longArray,
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

		returnValue := example_cpp_plugin.ParamAllPrimitives(
			boolArray[0], char16Array[0], sbyteArray[0], shortArray[0], intArray[0], longArray[0],
			byteArray[0], ushortArray[0], uintArray[0], ulongArray[0], intPtrArray[0], floatArray[0], doubleArray[0],
		)

		assert(returnValue == 56, fmt.Sprintf("Expected return value to be 56, but got %d", returnValue))

		fmt.Println("Go: OnPluginStart - done")
	})

	plugify.OnPluginEnd(func() {
		fmt.Println("Go: OnPluginEnd")
	})
}

//export NoParamReturnVoid
func NoParamReturnVoid() {
	fmt.Println("Go: NoParamReturnVoid")
}

//export NoParamReturnBool
func NoParamReturnBool() bool {
	fmt.Println("Go: NoParamReturnBool")
	return true
}

//export NoParamReturnChar8
func NoParamReturnChar8() int8 {
	fmt.Println("Go: NoParamReturnChar16")
	return math.MaxInt8
}

//export NoParamReturnChar16
func NoParamReturnChar16() uint16 {
	fmt.Println("Go: NoParamReturnChar16")
	return math.MaxUint16
}

//export NoParamReturnInt8
func NoParamReturnInt8() int8 {
	fmt.Println("Go: NoParamReturnInt8")
	return math.MaxInt8
}

//export NoParamReturnInt16
func NoParamReturnInt16() int16 {
	fmt.Println("Go: NoParamReturnInt16")
	return math.MaxInt16
}

//export NoParamReturnInt32
func NoParamReturnInt32() int32 {
	fmt.Println("Go: NoParamReturnInt32")
	return math.MaxInt32
}

//export NoParamReturnInt64
func NoParamReturnInt64() int64 {
	fmt.Println("Go: NoParamReturnInt64")
	return math.MaxInt64
}

//export NoParamReturnUInt8
func NoParamReturnUInt8() uint8 {
	fmt.Println("Go: NoParamReturnUInt8")
	return math.MaxUint8
}

//export NoParamReturnUInt16
func NoParamReturnUInt16() uint16 {
	fmt.Println("Go: NoParamReturnUInt16")
	return math.MaxUint16
}

//export NoParamReturnUInt32
func NoParamReturnUInt32() uint32 {
	fmt.Println("Go: NoParamReturnUInt32")
	return math.MaxUint32
}

//export NoParamReturnUInt64
func NoParamReturnUInt64() uint64 {
	fmt.Println("Go: NoParamReturnUInt64")
	return math.MaxUint64
}

//export NoParamReturnPtr64
func NoParamReturnPtr64() uintptr {
	fmt.Println("Go: NoParamReturnPtr64")
	return uintptr(1)
}

//export NoParamReturnFloat
func NoParamReturnFloat() float32 {
	fmt.Println("Go: NoParamReturnFloat")
	return math.MaxFloat32
}

//export NoParamReturnDouble
func NoParamReturnDouble() float64 {
	fmt.Println("Go: NoParamReturnDouble")
	return math.MaxFloat64
}

//export NoParamReturnFunction
func NoParamReturnFunction() uintptr {
	fmt.Println("Go: NoParamReturnFunction")
	return uintptr(0)
}

//export NoParamReturnString
func NoParamReturnString() string {
	fmt.Println("Go: NoParamReturnString")
	return "Hello World"
}

// Declare global variables
var globalBool []bool = []bool{true, false}
var globalChar8 []int8 = []int8{'a', 'b', 'c', 'd'}
var globalChar16 []uint16 = []uint16{'a', 'b', 'c', 'd'}
var globalInt8 []int8 = []int8{-3, -2, -1, 0, 1}
var globalInt16 []int16 = []int16{-4, -3, -2, -1, 0, 1}
var globalInt32 []int32 = []int32{-5, -4, -3, -2, -1, 0, 1}
var globalInt64 []int64 = []int64{-6, -5, -4, -3, -2, -1, 0, 1}
var globalUInt8 []uint8 = []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8}
var globalUInt16 []uint16 = []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var globalUInt32 []uint32 = []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var globalUInt64 []uint64 = []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
var globalPtr64 []uintptr = []uintptr{0, 1, 2, 3}
var globalFloat []float32 = []float32{-12.34, 0.0, 12.34}
var globalDouble []float64 = []float64{-12.345, 0.0, 12.345}
var globalString []string= []string{
	"1st string", "2nd string",
	"3rd element string (Should be big enough to avoid small string optimization)",
}

//export NoParamReturnArrayBool
func NoParamReturnArrayBool() []bool {
	fmt.Println("Go: NoParamReturnArrayBool")
	return globalBool
}

//export NoParamReturnArrayChar8
func NoParamReturnArrayChar8() []int8 {
	fmt.Println("Go: NoParamReturnArrayChar8")
	return globalChar8
}

//export NoParamReturnArrayChar16
func NoParamReturnArrayChar16() []uint16 {
	fmt.Println("Go: NoParamReturnArrayChar16")
	return globalChar16
}

//export NoParamReturnArrayInt8
func NoParamReturnArrayInt8() []int8 {
	fmt.Println("Go: NoParamReturnArrayInt8")
	return globalInt8
}

//export NoParamReturnArrayInt16
func NoParamReturnArrayInt16() []int16 {
	fmt.Println("Go: NoParamReturnArrayInt16")
	return globalInt16
}

//export NoParamReturnArrayInt32
func NoParamReturnArrayInt32() []int32 {
	fmt.Println("Go: NoParamReturnArrayInt32")
	return globalInt32
}

//export NoParamReturnArrayInt64
func NoParamReturnArrayInt64() []int64 {
	fmt.Println("Go: NoParamReturnArrayInt64")
	return globalInt64
}

//export NoParamReturnArrayUInt8
func NoParamReturnArrayUInt8() []uint8 {
	fmt.Println("Go: NoParamReturnArrayUInt8")
	return globalUInt8
}

//export NoParamReturnArrayUInt16
func NoParamReturnArrayUInt16() []uint16 {
	fmt.Println("Go: NoParamReturnArrayUInt16")
	return globalUInt16
}

//export NoParamReturnArrayUInt32
func NoParamReturnArrayUInt32() []uint32 {
	fmt.Println("Go: NoParamReturnArrayUInt32")
	return globalUInt32
}

//export NoParamReturnArrayUInt64
func NoParamReturnArrayUInt64() []uint64 {
	fmt.Println("Go: NoParamReturnArrayUInt64")
	return globalUInt64
}

//export NoParamReturnArrayPtr64
func NoParamReturnArrayPtr64() []uintptr {
	fmt.Println("Go: NoParamReturnArrayPtr64")
	return globalPtr64
}

//export NoParamReturnArrayFloat
func NoParamReturnArrayFloat() []float32 {
	fmt.Println("Go: NoParamReturnArrayFloat")
	return globalFloat
}

//export NoParamReturnArrayDouble
func NoParamReturnArrayDouble() []float64 {
	fmt.Println("Go: NoParamReturnArrayDouble")
	return globalDouble
}

//export NoParamReturnArrayString
func NoParamReturnArrayString() []string {
	fmt.Println("Go: NoParamReturnArrayString")
	return globalString
}

//export NoParamReturnVector2
func NoParamReturnVector2() C.Vector2 {
	fmt.Println("Go: NoParamReturnVector2")
	return C.Vector2{1, 2}
}
//export NoParamReturnVector3
func NoParamReturnVector3() C.Vector3 {
	fmt.Println("Go: NoParamReturnVector3")
	return C.Vector3{1, 2, 3}
}
//export NoParamReturnVector4
func NoParamReturnVector4() C.Vector4 {
	fmt.Println("Go: NoParamReturnVector4")
	return C.Vector4{1, 2, 3, 4}
}
//export NoParamReturnMatrix4x4
func NoParamReturnMatrix4x4() C.Matrix4x4 {
	fmt.Println("Go: NoParamReturnMatrix4x4")
	return C.Matrix4x4{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16};
}

//export Param1
func Param1(a int32) {
	fmt.Printf("Param1: a = %d\n", a)
}

//export Param2
func Param2(a int32, b float32) {
	fmt.Printf("Param2: a = %d, b = %f\n", a, b)
}

//export Param3
func Param3(a int32, b float32, c float64) {
	fmt.Printf("Param3: a = %d, b = %f, c = %f\n", a, b, c)
}

//export Param4
func Param4(a int32, b float32, c float64, d C.Vector4) {
	fmt.Printf("Param4: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f]\n", a, b, c, d.x, d.y, d.z, d.w)
}

//export Param5
func Param5(a int32, b float32, c float64, d C.Vector4, e []int64) {
	fmt.Printf("Param5: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Println("]")
}

//export Param6
func Param6(a int32, b float32, c float64, d C.Vector4, e []int64, f uint16) {
	fmt.Printf("Param6: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c\n", f)
}

//export Param7
func Param7(a int32, b float32, c float64, d C.Vector4, e []int64, f uint16, g string) {
	fmt.Printf("Param7: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s\n", f, g)
}

//export Param8
func Param8(a int32, b float32, c float64, d C.Vector4, e []int64, f uint16, g string, h float32) {
	fmt.Printf("Param8: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s, h = %f\n", f, g, h)
}

//export Param9
func Param9(a int32, b float32, c float64, d C.Vector4, e []int64, f uint16, g string, h float32, k int16) {
	fmt.Printf("Param9: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s, h = %f, k = %d\n", f, g, h, k)
}

//export Param10
func Param10(a int32, b float32, c float64, d C.Vector4, e []int64, f uint16, g string, h float32, k int16, l uintptr) {
	fmt.Printf("Param10: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s, h = %f, k = %d, l = %v\n", f, g, h, k, l)
}

//export ParamRef1
func ParamRef1(a *int32) {
	*a = 42
}

//export ParamRef2
func ParamRef2(a *int32, b *float32) {
	*a = 10
	*b = 3.14;
}

//export ParamRef3
func ParamRef3(a *int32, b *float32, c *float64) {
	*a = -20
	*b = 2.718
	*c = 3.14159
}

//export ParamRef4
func ParamRef4(a *int32, b *float32, c *float64, d *C.Vector4) {
	*a = 100
	*b = -5.55
	*c = 1.618
	*d = C.Vector4{1.0, 2.0, 3.0, 4.0}
}

//export ParamRef5
func ParamRef5(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64) {
	*a = 500
	*b = -10.5
	*c = 2.71828
	*d = C.Vector4{-1.0, -2.0, -3.0, -4.0}
	*e = []int64{-6, -5, -4, -3, -2, -1, 0, 1}
}

//export ParamRef6
func ParamRef6(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *uint16) {
	*a = 750
	*b = 20.0
	*c = 1.23456
	*d = C.Vector4{10.0, 20.0, 30.0, 40.0}
	*e = []int64{-6, -5, -4}
	*f = 'Z'
}

//export ParamRef7
func ParamRef7(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *uint16, g *string) {
	*a = -1000
	*b = 3.0
	*c = -1.0
	*d = C.Vector4{100.0, 200.0, 300.0, 400.0}
	*e = []int64{-6, -5, -4, -3}
	*f = uint16('X')
	*g = "Hello, World!"
}

//export ParamRef8
func ParamRef8(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *uint16, g *string, h *float32) {
	*a = 999
	*b = -7.5
	*c = 0.123456
	*d = C.Vector4{-100.0, -200.0, -300.0, -400.0}
	*e = []int64{-6, -5, -4, -3, -2, -1}
	*f = uint16('Y')
	*g = "Goodbye, World!"
	*h = 99.99
}

//export ParamRef9
func ParamRef9(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *uint16, g *string, h *float32, k *int16) {
	*a = -1234
	*b = 123.45
	*c = -678.9
	*d = C.Vector4{987.65, 432.1, 123.456, 789.123}
	*e = []int64{-6, -5, -4, -3, -2, -1, 0, 1, 5, 9}
	*f = uint16('A')
	*g = "Testing, 1 2 3"
	*h = -987.654
	*k = 42
}

//export ParamRef10
func ParamRef10(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *uint16, g *string, h *float32, k *int16, l *uintptr) {
	*a = 987
	*b = -0.123
	*c = 456.789
	*d = C.Vector4{-123.456, 0.987, 654.321, -789.123}
	*e = []int64{-6, -5, -4, -3, -2, -1, 0, 1, 5, 9, 4, -7}
	*f = uint16('B')
	*g = "Another string"
	*h = 3.141592
	*k = -32768
	*l = 0 // Null pointer
}

//export ParamRefVectors
func ParamRefVectors(p1 *[]bool, p2 *[]int8, p3 *[]uint16, p4 *[]int8, p5 *[]int16, p6 *[]int32, p7 *[]int64, p8 *[]uint8, p9 *[]uint16, p10 *[]uint32, p11 *[]uint64, p12 *[]uintptr, p13 *[]float32, p14 *[]float64, p15 *[]string) {
	*p1 = []bool{true}
	*p2 = []int8{'a', 'b', 'c'}
	*p3 = []uint16{'a', 'b', 'c'}
	*p4 = []int8{-3, -2, -1, 0, 1, 2, 3}
	*p5 = []int16{-4, -3, -2, -1, 0, 1, 2, 3, 4}
	*p6 = []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	*p7 = []int64{-6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6}
	*p8 = []uint8{0, 1, 2, 3, 4, 5, 6, 7}
	*p9 = []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8}
	*p10 = []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	*p11 = []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	*p12 = []uintptr{0, 1, 2}
	*p13 = []float32{-12.34, 0.0, 12.34}
	*p14 = []float64{-12.345, 0.0, 12.345}
	*p15 = []string{"1", "12", "123", "1234", "12345", "123456"}
}

//export ParamAllPrimitives
func ParamAllPrimitives(p1 bool, p2 uint16, p3 int8, p4 int16, p5 int32, p6 int64, p7 uint8, p8 uint16, p9 uint32, p10 uint64, p11 uintptr, p12 float32, p13 float64) int64 {
	sum := int64(0)
	if p1 {
		sum += int64(1)
	}
	sum += int64(p2)
	sum += int64(p3)
	sum += int64(p4)
	sum += int64(p5)
	sum += p6
	sum += int64(p7)
	sum += int64(p8)
	sum += int64(p9)
	sum += int64(p10)
	sum += int64(p11)
	sum += int64(p12)
	sum += int64(p13)
	return sum
}

func main() {}
