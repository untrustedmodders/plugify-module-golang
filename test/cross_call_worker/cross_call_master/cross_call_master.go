package cross_call_master

/*
#include "cross_call_master.h"
#cgo noescape ReverseReturn
#cgo noescape NoParamReturnVoidCallback
#cgo noescape NoParamReturnBoolCallback
#cgo noescape NoParamReturnChar8Callback
#cgo noescape NoParamReturnChar16Callback
#cgo noescape NoParamReturnInt8Callback
#cgo noescape NoParamReturnInt16Callback
#cgo noescape NoParamReturnInt32Callback
#cgo noescape NoParamReturnInt64Callback
#cgo noescape NoParamReturnUInt8Callback
#cgo noescape NoParamReturnUInt16Callback
#cgo noescape NoParamReturnUInt32Callback
#cgo noescape NoParamReturnUInt64Callback
#cgo noescape NoParamReturnPointerCallback
#cgo noescape NoParamReturnFloatCallback
#cgo noescape NoParamReturnDoubleCallback
#cgo noescape NoParamReturnFunctionCallback
#cgo noescape NoParamReturnStringCallback
#cgo noescape NoParamReturnAnyCallback
#cgo noescape NoParamReturnArrayBoolCallback
#cgo noescape NoParamReturnArrayChar8Callback
#cgo noescape NoParamReturnArrayChar16Callback
#cgo noescape NoParamReturnArrayInt8Callback
#cgo noescape NoParamReturnArrayInt16Callback
#cgo noescape NoParamReturnArrayInt32Callback
#cgo noescape NoParamReturnArrayInt64Callback
#cgo noescape NoParamReturnArrayUInt8Callback
#cgo noescape NoParamReturnArrayUInt16Callback
#cgo noescape NoParamReturnArrayUInt32Callback
#cgo noescape NoParamReturnArrayUInt64Callback
#cgo noescape NoParamReturnArrayPointerCallback
#cgo noescape NoParamReturnArrayFloatCallback
#cgo noescape NoParamReturnArrayDoubleCallback
#cgo noescape NoParamReturnArrayStringCallback
#cgo noescape NoParamReturnArrayAnyCallback
#cgo noescape NoParamReturnArrayVector2Callback
#cgo noescape NoParamReturnArrayVector3Callback
#cgo noescape NoParamReturnArrayVector4Callback
#cgo noescape NoParamReturnArrayMatrix4x4Callback
#cgo noescape NoParamReturnVector2Callback
#cgo noescape NoParamReturnVector3Callback
#cgo noescape NoParamReturnVector4Callback
#cgo noescape NoParamReturnMatrix4x4Callback
#cgo noescape Param1Callback
#cgo noescape Param2Callback
#cgo noescape Param3Callback
#cgo noescape Param4Callback
#cgo noescape Param5Callback
#cgo noescape Param6Callback
#cgo noescape Param7Callback
#cgo noescape Param8Callback
#cgo noescape Param9Callback
#cgo noescape Param10Callback
#cgo noescape ParamRef1Callback
#cgo noescape ParamRef2Callback
#cgo noescape ParamRef3Callback
#cgo noescape ParamRef4Callback
#cgo noescape ParamRef5Callback
#cgo noescape ParamRef6Callback
#cgo noescape ParamRef7Callback
#cgo noescape ParamRef8Callback
#cgo noescape ParamRef9Callback
#cgo noescape ParamRef10Callback
#cgo noescape ParamRefVectorsCallback
#cgo noescape ParamAllPrimitivesCallback
#cgo noescape ParamEnumCallback
#cgo noescape ParamEnumRefCallback
#cgo noescape ParamVariantCallback
#cgo noescape ParamVariantRefCallback
#cgo noescape CallFuncVoidCallback
#cgo noescape CallFuncBoolCallback
#cgo noescape CallFuncChar8Callback
#cgo noescape CallFuncChar16Callback
#cgo noescape CallFuncInt8Callback
#cgo noescape CallFuncInt16Callback
#cgo noescape CallFuncInt32Callback
#cgo noescape CallFuncInt64Callback
#cgo noescape CallFuncUInt8Callback
#cgo noescape CallFuncUInt16Callback
#cgo noescape CallFuncUInt32Callback
#cgo noescape CallFuncUInt64Callback
#cgo noescape CallFuncPtrCallback
#cgo noescape CallFuncFloatCallback
#cgo noescape CallFuncDoubleCallback
#cgo noescape CallFuncStringCallback
#cgo noescape CallFuncAnyCallback
#cgo noescape CallFuncFunctionCallback
#cgo noescape CallFuncBoolVectorCallback
#cgo noescape CallFuncChar8VectorCallback
#cgo noescape CallFuncChar16VectorCallback
#cgo noescape CallFuncInt8VectorCallback
#cgo noescape CallFuncInt16VectorCallback
#cgo noescape CallFuncInt32VectorCallback
#cgo noescape CallFuncInt64VectorCallback
#cgo noescape CallFuncUInt8VectorCallback
#cgo noescape CallFuncUInt16VectorCallback
#cgo noescape CallFuncUInt32VectorCallback
#cgo noescape CallFuncUInt64VectorCallback
#cgo noescape CallFuncPtrVectorCallback
#cgo noescape CallFuncFloatVectorCallback
#cgo noescape CallFuncDoubleVectorCallback
#cgo noescape CallFuncStringVectorCallback
#cgo noescape CallFuncAnyVectorCallback
#cgo noescape CallFuncVec2VectorCallback
#cgo noescape CallFuncVec3VectorCallback
#cgo noescape CallFuncVec4VectorCallback
#cgo noescape CallFuncMat4x4VectorCallback
#cgo noescape CallFuncVec2Callback
#cgo noescape CallFuncVec3Callback
#cgo noescape CallFuncVec4Callback
#cgo noescape CallFuncMat4x4Callback
#cgo noescape CallFunc1Callback
#cgo noescape CallFunc2Callback
#cgo noescape CallFunc3Callback
#cgo noescape CallFunc4Callback
#cgo noescape CallFunc5Callback
#cgo noescape CallFunc6Callback
#cgo noescape CallFunc7Callback
#cgo noescape CallFunc8Callback
#cgo noescape CallFunc9Callback
#cgo noescape CallFunc10Callback
#cgo noescape CallFunc11Callback
#cgo noescape CallFunc12Callback
#cgo noescape CallFunc13Callback
#cgo noescape CallFunc14Callback
#cgo noescape CallFunc15Callback
#cgo noescape CallFunc16Callback
#cgo noescape CallFunc17Callback
#cgo noescape CallFunc18Callback
#cgo noescape CallFunc19Callback
#cgo noescape CallFunc20Callback
#cgo noescape CallFunc21Callback
#cgo noescape CallFunc22Callback
#cgo noescape CallFunc23Callback
#cgo noescape CallFunc24Callback
#cgo noescape CallFunc25Callback
#cgo noescape CallFunc26Callback
#cgo noescape CallFunc27Callback
#cgo noescape CallFunc28Callback
#cgo noescape CallFunc29Callback
#cgo noescape CallFunc30Callback
#cgo noescape CallFunc31Callback
#cgo noescape CallFunc32Callback
#cgo noescape CallFunc33Callback
#cgo noescape CallFuncEnumCallback
*/
import "C"
import (
	"github.com/untrustedmodders/go-plugify"
	"reflect"
	"unsafe"
)

// Generated with https://github.com/untrustedmodders/plugify-module-golang/blob/main/generator/generator.py from cross_call_master

type Example = int32

const (
	First  Example = 1
	Second Example = 2
	Third  Example = 3
	Forth  Example = 4
)

type NoParamReturnFunctionCallbackFunc func() int32
type FuncVoid func()
type FuncBool func() bool
type FuncChar8 func() int8
type FuncChar16 func() uint16
type FuncInt8 func() int8
type FuncInt16 func() int16
type FuncInt32 func() int32
type FuncInt64 func() int64
type FuncUInt8 func() uint8
type FuncUInt16 func() uint16
type FuncUInt32 func() uint32
type FuncUInt64 func() uint64
type FuncPtr func() uintptr
type FuncFloat func() float32
type FuncDouble func() float64
type FuncString func() string
type FuncAny func() interface{}
type FuncFunction func() uintptr
type FuncBoolVector func() []bool
type FuncChar8Vector func() []int8
type FuncChar16Vector func() []uint16
type FuncInt8Vector func() []int8
type FuncInt16Vector func() []int16
type FuncInt32Vector func() []int32
type FuncInt64Vector func() []int64
type FuncUInt8Vector func() []uint8
type FuncUInt16Vector func() []uint16
type FuncUInt32Vector func() []uint32
type FuncUInt64Vector func() []uint64
type FuncPtrVector func() []uintptr
type FuncFloatVector func() []float32
type FuncDoubleVector func() []float64
type FuncStringVector func() []string
type FuncAnyVector func() []interface{}
type FuncVec2Vector func() []plugify.Vector2
type FuncVec3Vector func() []plugify.Vector3
type FuncVec4Vector func() []plugify.Vector4
type FuncMat4x4Vector func() []plugify.Matrix4x4
type FuncVec2 func() plugify.Vector2
type FuncVec3 func() plugify.Vector3
type FuncVec4 func() plugify.Vector4
type FuncMat4x4 func() plugify.Matrix4x4
type Func1 func(a plugify.Vector3) int32
type Func2 func(a float32, b int64) int8
type Func3 func(a uintptr, b plugify.Vector4, c string)
type Func4 func(a bool, b int32, c uint16, d plugify.Matrix4x4) plugify.Vector4
type Func5 func(a int8, b plugify.Vector2, c uintptr, d float64, e []uint64) bool
type Func6 func(a string, b float32, c []float32, d int16, e []uint8, f uintptr) int64
type Func7 func(vecC []int8, u16 uint16, ch16 uint16, vecU32 []uint32, vec4 plugify.Vector4, b bool, u64 uint64) float64
type Func8 func(vec3 plugify.Vector3, vecU32 []uint32, i16 int16, b bool, vec4 plugify.Vector4, vecC16 []uint16, ch16 uint16, i32 int32) plugify.Matrix4x4
type Func9 func(f float32, vec2 plugify.Vector2, vecI8 []int8, u64 uint64, b bool, str string, vec4 plugify.Vector4, i16 int16, ptr uintptr)
type Func10 func(vec4 plugify.Vector4, mat plugify.Matrix4x4, vecU32 []uint32, u64 uint64, vecC []int8, i32 int32, b bool, vec2 plugify.Vector2, i64 int64, d float64) uint32
type Func11 func(vecB []bool, ch16 uint16, u8 uint8, d float64, vec3 plugify.Vector3, vecI8 []int8, i64 int64, u16 uint16, f float32, vec2 plugify.Vector2, u32 uint32) uintptr
type Func12 func(ptr uintptr, vecD []float64, u32 uint32, d float64, b bool, i32 int32, i8 int8, u64 uint64, f float32, vecPtr []uintptr, i64 int64, ch int8) bool
type Func13 func(i64 int64, vecC []int8, d uint16, f float32, b []bool, vec4 plugify.Vector4, str string, int32_ int32, vec3 plugify.Vector3, ptr uintptr, vec2 plugify.Vector2, arr []uint8, i16 int16) string
type Func14 func(vecC []int8, vecU32 []uint32, mat plugify.Matrix4x4, b bool, ch16 uint16, i32 int32, vecF []float32, u16 uint16, vecU8 []uint8, i8 int8, vec3 plugify.Vector3, vec4 plugify.Vector4, d float64, ptr uintptr) []string
type Func15 func(vecI16 []int16, mat plugify.Matrix4x4, vec4 plugify.Vector4, ptr uintptr, u64 uint64, vecU32 []uint32, b bool, f float32, vecC16 []uint16, u8 uint8, i32 int32, vec2 plugify.Vector2, u16 uint16, d float64, vecU8 []uint8) int16
type Func16 func(vecB []bool, i16 int16, vecI8 []int8, vec4 plugify.Vector4, mat plugify.Matrix4x4, vec2 plugify.Vector2, vecU64 []uint64, vecC []int8, str string, i64 int64, vecU32 []uint32, vec3 plugify.Vector3, f float32, d float64, i8 int8, u16 uint16) uintptr
type Func17 func(i32 *int32)
type Func18 func(i8 *int8, i16 *int16) plugify.Vector2
type Func19 func(u32 *uint32, vec3 *plugify.Vector3, vecU32 *[]uint32)
type Func20 func(ch16 *uint16, vec4 *plugify.Vector4, vecU64 *[]uint64, ch *int8) int32
type Func21 func(mat *plugify.Matrix4x4, vecI32 *[]int32, vec2 *plugify.Vector2, b *bool, extraParam *float64) float32
type Func22 func(ptr64Ref *uintptr, uint32Ref *uint32, vectorDoubleRef *[]float64, int16Ref *int16, plgStringRef *string, plgVector4Ref *plugify.Vector4) uint64
type Func23 func(uint64Ref *uint64, plgVector2Ref *plugify.Vector2, vectorInt16Ref *[]int16, char16Ref *uint16, floatRef *float32, int8Ref *int8, vectorUInt8Ref *[]uint8)
type Func24 func(vectorCharRef *[]int8, int64Ref *int64, vectorUInt8Ref *[]uint8, plgVector4Ref *plugify.Vector4, uint64Ref *uint64, vectorptr64Ref *[]uintptr, doubleRef *float64, vectorptr64Ref2 *[]uintptr) plugify.Matrix4x4
type Func25 func(int32Ref *int32, vectorptr64Ref *[]uintptr, boolRef *bool, uint8Ref *uint8, plgStringRef *string, plgVector3Ref *plugify.Vector3, int64Ref *int64, plgVector4Ref *plugify.Vector4, uint16Ref *uint16) float64
type Func26 func(char16Ref *uint16, plgVector2Ref *plugify.Vector2, plgMatrix4x4Ref *plugify.Matrix4x4, vectorFloatRef *[]float32, int16Ref *int16, uint64Ref *uint64, uint32Ref *uint32, vectorUInt16Ref *[]uint16, ptr64Ref *uintptr, boolRef *bool) int8
type Func27 func(floatRef *float32, plgVector3Ref *plugify.Vector3, ptr64Ref *uintptr, plgVector2Ref *plugify.Vector2, vectorInt16Ref *[]int16, plgMatrix4x4Ref *plugify.Matrix4x4, boolRef *bool, plgVector4Ref *plugify.Vector4, int8Ref *int8, int32Ref *int32, vectorUInt8Ref *[]uint8) uint8
type Func28 func(ptr64Ref *uintptr, uint16Ref *uint16, vectorUInt32Ref *[]uint32, plgMatrix4x4Ref *plugify.Matrix4x4, floatRef *float32, plgVector4Ref *plugify.Vector4, plgStringRef *string, vectorUInt64Ref *[]uint64, int64Ref *int64, boolRef *bool, plgVector3Ref *plugify.Vector3, vectorFloatRef *[]float32) string
type Func29 func(plgVector4Ref *plugify.Vector4, int32Ref *int32, vectorInt8Ref *[]int8, doubleRef *float64, boolRef *bool, int8Ref *int8, vectorUInt16Ref *[]uint16, floatRef *float32, plgStringRef *string, plgMatrix4x4Ref *plugify.Matrix4x4, uint64Ref *uint64, plgVector3Ref *plugify.Vector3, vectorInt64Ref *[]int64) []string
type Func30 func(ptr64Ref *uintptr, plgVector4Ref *plugify.Vector4, int64Ref *int64, vectorUInt32Ref *[]uint32, boolRef *bool, plgStringRef *string, plgVector3Ref *plugify.Vector3, vectorUInt8Ref *[]uint8, floatRef *float32, plgVector2Ref *plugify.Vector2, plgMatrix4x4Ref *plugify.Matrix4x4, int8Ref *int8, vectorFloatRef *[]float32, doubleRef *float64) int32
type Func31 func(charRef *int8, uint32Ref *uint32, vectorUInt64Ref *[]uint64, plgVector4Ref *plugify.Vector4, plgStringRef *string, boolRef *bool, int64Ref *int64, vec2Ref *plugify.Vector2, int8Ref *int8, uint16Ref *uint16, vectorInt16Ref *[]int16, mat4x4Ref *plugify.Matrix4x4, vec3Ref *plugify.Vector3, floatRef *float32, vectorDoubleRef *[]float64) plugify.Vector3
type Func32 func(p1 *int32, p2 *uint16, p3 *[]int8, p4 *plugify.Vector4, p5 *uintptr, p6 *[]uint32, p7 *plugify.Matrix4x4, p8 *uint64, p9 *string, p10 *int64, p11 *plugify.Vector2, p12 *[]int8, p13 *bool, p14 *plugify.Vector3, p15 *uint8, p16 *[]uint16) float64
type Func33 func(variant *interface{})
type FuncEnum func(p1 Example, p2 *[]Example) []Example

// ReverseReturn - No description provided.
// @param returnString: No description available.
func ReverseReturn(returnString string) {
	__returnString := plugify.ConstructString(returnString)
	plugify.Block{
		Try: func() {
			C.ReverseReturn((*C.String)(unsafe.Pointer(&__returnString)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__returnString)
		},
	}.Do()
}

// NoParamReturnVoidCallback - No description provided.
func NoParamReturnVoidCallback() {
	C.NoParamReturnVoidCallback()
}

// NoParamReturnBoolCallback - No description provided.
// @return No description available.
func NoParamReturnBoolCallback() bool {
	__retVal := bool(C.NoParamReturnBoolCallback())
	return __retVal
}

// NoParamReturnChar8Callback - No description provided.
// @return No description available.
func NoParamReturnChar8Callback() int8 {
	__retVal := int8(C.NoParamReturnChar8Callback())
	return __retVal
}

// NoParamReturnChar16Callback - No description provided.
// @return No description available.
func NoParamReturnChar16Callback() uint16 {
	__retVal := uint16(C.NoParamReturnChar16Callback())
	return __retVal
}

// NoParamReturnInt8Callback - No description provided.
// @return No description available.
func NoParamReturnInt8Callback() int8 {
	__retVal := int8(C.NoParamReturnInt8Callback())
	return __retVal
}

// NoParamReturnInt16Callback - No description provided.
// @return No description available.
func NoParamReturnInt16Callback() int16 {
	__retVal := int16(C.NoParamReturnInt16Callback())
	return __retVal
}

// NoParamReturnInt32Callback - No description provided.
// @return No description available.
func NoParamReturnInt32Callback() int32 {
	__retVal := int32(C.NoParamReturnInt32Callback())
	return __retVal
}

// NoParamReturnInt64Callback - No description provided.
// @return No description available.
func NoParamReturnInt64Callback() int64 {
	__retVal := int64(C.NoParamReturnInt64Callback())
	return __retVal
}

// NoParamReturnUInt8Callback - No description provided.
// @return No description available.
func NoParamReturnUInt8Callback() uint8 {
	__retVal := uint8(C.NoParamReturnUInt8Callback())
	return __retVal
}

// NoParamReturnUInt16Callback - No description provided.
// @return No description available.
func NoParamReturnUInt16Callback() uint16 {
	__retVal := uint16(C.NoParamReturnUInt16Callback())
	return __retVal
}

// NoParamReturnUInt32Callback - No description provided.
// @return No description available.
func NoParamReturnUInt32Callback() uint32 {
	__retVal := uint32(C.NoParamReturnUInt32Callback())
	return __retVal
}

// NoParamReturnUInt64Callback - No description provided.
// @return No description available.
func NoParamReturnUInt64Callback() uint64 {
	__retVal := uint64(C.NoParamReturnUInt64Callback())
	return __retVal
}

// NoParamReturnPointerCallback - No description provided.
// @return No description available.
func NoParamReturnPointerCallback() uintptr {
	__retVal := uintptr(C.NoParamReturnPointerCallback())
	return __retVal
}

// NoParamReturnFloatCallback - No description provided.
// @return No description available.
func NoParamReturnFloatCallback() float32 {
	__retVal := float32(C.NoParamReturnFloatCallback())
	return __retVal
}

// NoParamReturnDoubleCallback - No description provided.
// @return No description available.
func NoParamReturnDoubleCallback() float64 {
	__retVal := float64(C.NoParamReturnDoubleCallback())
	return __retVal
}

// NoParamReturnFunctionCallback - No description provided.
// @return No description available.
func NoParamReturnFunctionCallback() NoParamReturnFunctionCallbackFunc {
	__retVal := plugify.GetDelegateForFunctionPointer(C.NoParamReturnFunctionCallback(), reflect.TypeOf(NoParamReturnFunctionCallbackFunc(nil))).(NoParamReturnFunctionCallbackFunc)
	return __retVal
}

// NoParamReturnStringCallback - No description provided.
// @return No description available.
func NoParamReturnStringCallback() string {
	var __retVal string
	var __retVal_native plugify.PlgString
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnStringCallback()
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnAnyCallback - No description provided.
// @return No description available.
func NoParamReturnAnyCallback() interface{} {
	var __retVal interface{}
	var __retVal_native plugify.PlgVariant
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnAnyCallback()
			__retVal_native = *(*plugify.PlgVariant)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVariantData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayBoolCallback - No description provided.
// @return No description available.
func NoParamReturnArrayBoolCallback() []bool {
	var __retVal []bool
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayBoolCallback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataBool(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorBool(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayChar8Callback - No description provided.
// @return No description available.
func NoParamReturnArrayChar8Callback() []int8 {
	var __retVal []int8
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayChar8Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataChar8(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorChar8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayChar16Callback - No description provided.
// @return No description available.
func NoParamReturnArrayChar16Callback() []uint16 {
	var __retVal []uint16
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayChar16Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataChar16(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorChar16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayInt8Callback - No description provided.
// @return No description available.
func NoParamReturnArrayInt8Callback() []int8 {
	var __retVal []int8
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayInt8Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt8(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayInt16Callback - No description provided.
// @return No description available.
func NoParamReturnArrayInt16Callback() []int16 {
	var __retVal []int16
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayInt16Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt16(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayInt32Callback - No description provided.
// @return No description available.
func NoParamReturnArrayInt32Callback() []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayInt32Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt32(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayInt64Callback - No description provided.
// @return No description available.
func NoParamReturnArrayInt64Callback() []int64 {
	var __retVal []int64
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayInt64Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt64(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayUInt8Callback - No description provided.
// @return No description available.
func NoParamReturnArrayUInt8Callback() []uint8 {
	var __retVal []uint8
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayUInt8Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt8(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayUInt16Callback - No description provided.
// @return No description available.
func NoParamReturnArrayUInt16Callback() []uint16 {
	var __retVal []uint16
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayUInt16Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt16(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayUInt32Callback - No description provided.
// @return No description available.
func NoParamReturnArrayUInt32Callback() []uint32 {
	var __retVal []uint32
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayUInt32Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt32(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayUInt64Callback - No description provided.
// @return No description available.
func NoParamReturnArrayUInt64Callback() []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayUInt64Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt64(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayPointerCallback - No description provided.
// @return No description available.
func NoParamReturnArrayPointerCallback() []uintptr {
	var __retVal []uintptr
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayPointerCallback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataPointer(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorPointer(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayFloatCallback - No description provided.
// @return No description available.
func NoParamReturnArrayFloatCallback() []float32 {
	var __retVal []float32
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayFloatCallback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataFloat(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorFloat(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayDoubleCallback - No description provided.
// @return No description available.
func NoParamReturnArrayDoubleCallback() []float64 {
	var __retVal []float64
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayDoubleCallback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataDouble(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorDouble(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayStringCallback - No description provided.
// @return No description available.
func NoParamReturnArrayStringCallback() []string {
	var __retVal []string
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayStringCallback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataString(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayAnyCallback - No description provided.
// @return No description available.
func NoParamReturnArrayAnyCallback() []interface{} {
	var __retVal []interface{}
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayAnyCallback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVariant(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVariant(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayVector2Callback - No description provided.
// @return No description available.
func NoParamReturnArrayVector2Callback() []plugify.Vector2 {
	var __retVal []plugify.Vector2
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayVector2Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector2(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector2(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayVector3Callback - No description provided.
// @return No description available.
func NoParamReturnArrayVector3Callback() []plugify.Vector3 {
	var __retVal []plugify.Vector3
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayVector3Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector3(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector3(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayVector4Callback - No description provided.
// @return No description available.
func NoParamReturnArrayVector4Callback() []plugify.Vector4 {
	var __retVal []plugify.Vector4
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayVector4Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector4(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector4(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayMatrix4x4Callback - No description provided.
// @return No description available.
func NoParamReturnArrayMatrix4x4Callback() []plugify.Matrix4x4 {
	var __retVal []plugify.Matrix4x4
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayMatrix4x4Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataMatrix4x4(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorMatrix4x4(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnVector2Callback - No description provided.
// @return No description available.
func NoParamReturnVector2Callback() plugify.Vector2 {
	__native := C.NoParamReturnVector2Callback()
	__retVal := *(*plugify.Vector2)(unsafe.Pointer(&__native))
	return __retVal
}

// NoParamReturnVector3Callback - No description provided.
// @return No description available.
func NoParamReturnVector3Callback() plugify.Vector3 {
	__native := C.NoParamReturnVector3Callback()
	__retVal := *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// NoParamReturnVector4Callback - No description provided.
// @return No description available.
func NoParamReturnVector4Callback() plugify.Vector4 {
	__native := C.NoParamReturnVector4Callback()
	__retVal := *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// NoParamReturnMatrix4x4Callback - No description provided.
// @return No description available.
func NoParamReturnMatrix4x4Callback() plugify.Matrix4x4 {
	__native := C.NoParamReturnMatrix4x4Callback()
	__retVal := *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// Param1Callback - No description provided.
// @param a: No description available.
func Param1Callback(a int32) {
	__a := C.int32_t(a)
	C.Param1Callback(__a)
}

// Param2Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
func Param2Callback(a int32, b float32) {
	__a := C.int32_t(a)
	__b := C.float(b)
	C.Param2Callback(__a, __b)
}

// Param3Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
func Param3Callback(a int32, b float32, c float64) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	C.Param3Callback(__a, __b, __c)
}

// Param4Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
func Param4Callback(a int32, b float32, c float64, d plugify.Vector4) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	C.Param4Callback(__a, __b, __c, &__d)
}

// Param5Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
// @param e: No description available.
func Param5Callback(a int32, b float32, c float64, d plugify.Vector4, e []int64) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	__e := plugify.ConstructVectorInt64(e)
	plugify.Block{
		Try: func() {
			C.Param5Callback(__a, __b, __c, &__d, (*C.Vector)(unsafe.Pointer(&__e)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
		},
	}.Do()
}

// Param6Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
// @param e: No description available.
// @param f: No description available.
func Param6Callback(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	__e := plugify.ConstructVectorInt64(e)
	__f := C.int8_t(f)
	plugify.Block{
		Try: func() {
			C.Param6Callback(__a, __b, __c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), __f)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
		},
	}.Do()
}

// Param7Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
// @param e: No description available.
// @param f: No description available.
// @param g: No description available.
func Param7Callback(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	__e := plugify.ConstructVectorInt64(e)
	__f := C.int8_t(f)
	__g := plugify.ConstructString(g)
	plugify.Block{
		Try: func() {
			C.Param7Callback(__a, __b, __c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), __f, (*C.String)(unsafe.Pointer(&__g)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// Param8Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
// @param e: No description available.
// @param f: No description available.
// @param g: No description available.
// @param h: No description available.
func Param8Callback(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	__e := plugify.ConstructVectorInt64(e)
	__f := C.int8_t(f)
	__g := plugify.ConstructString(g)
	__h := C.uint16_t(h)
	plugify.Block{
		Try: func() {
			C.Param8Callback(__a, __b, __c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), __f, (*C.String)(unsafe.Pointer(&__g)), __h)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// Param9Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
// @param e: No description available.
// @param f: No description available.
// @param g: No description available.
// @param h: No description available.
// @param k: No description available.
func Param9Callback(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16, k int16) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	__e := plugify.ConstructVectorInt64(e)
	__f := C.int8_t(f)
	__g := plugify.ConstructString(g)
	__h := C.uint16_t(h)
	__k := C.int16_t(k)
	plugify.Block{
		Try: func() {
			C.Param9Callback(__a, __b, __c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), __f, (*C.String)(unsafe.Pointer(&__g)), __h, __k)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// Param10Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
// @param e: No description available.
// @param f: No description available.
// @param g: No description available.
// @param h: No description available.
// @param k: No description available.
// @param l: No description available.
func Param10Callback(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16, k int16, l uintptr) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	__e := plugify.ConstructVectorInt64(e)
	__f := C.int8_t(f)
	__g := plugify.ConstructString(g)
	__h := C.uint16_t(h)
	__k := C.int16_t(k)
	__l := C.uintptr_t(l)
	plugify.Block{
		Try: func() {
			C.Param10Callback(__a, __b, __c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), __f, (*C.String)(unsafe.Pointer(&__g)), __h, __k, __l)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// ParamRef1Callback - No description provided.
// @param a: No description available.
func ParamRef1Callback(a *int32) {
	__a := C.int32_t(*a)
	C.ParamRef1Callback(&__a)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
}

// ParamRef2Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
func ParamRef2Callback(a *int32, b *float32) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	C.ParamRef2Callback(&__a, &__b)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
	*b = float32(__b)
}

// ParamRef3Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
func ParamRef3Callback(a *int32, b *float32, c *float64) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	C.ParamRef3Callback(&__a, &__b, &__c)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
	*b = float32(__b)
	*c = float64(__c)
}

// ParamRef4Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
func ParamRef4Callback(a *int32, b *float32, c *float64, d *plugify.Vector4) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	__d := *(*C.Vector4)(unsafe.Pointer(d))
	C.ParamRef4Callback(&__a, &__b, &__c, &__d)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
	*b = float32(__b)
	*c = float64(__c)
	*d = *(*plugify.Vector4)(unsafe.Pointer(&__d))
}

// ParamRef5Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
// @param e: No description available.
func ParamRef5Callback(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	__d := *(*C.Vector4)(unsafe.Pointer(d))
	__e := plugify.ConstructVectorInt64(*e)
	plugify.Block{
		Try: func() {
			C.ParamRef5Callback(&__a, &__b, &__c, &__d, (*C.Vector)(unsafe.Pointer(&__e)))
			// Unmarshal - Convert native data to managed data.
			*a = int32(__a)
			*b = float32(__b)
			*c = float64(__c)
			*d = *(*plugify.Vector4)(unsafe.Pointer(&__d))
			plugify.GetVectorDataInt64To(&__e, e)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
		},
	}.Do()
}

// ParamRef6Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
// @param e: No description available.
// @param f: No description available.
func ParamRef6Callback(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	__d := *(*C.Vector4)(unsafe.Pointer(d))
	__e := plugify.ConstructVectorInt64(*e)
	__f := C.int8_t(*f)
	plugify.Block{
		Try: func() {
			C.ParamRef6Callback(&__a, &__b, &__c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), &__f)
			// Unmarshal - Convert native data to managed data.
			*a = int32(__a)
			*b = float32(__b)
			*c = float64(__c)
			*d = *(*plugify.Vector4)(unsafe.Pointer(&__d))
			plugify.GetVectorDataInt64To(&__e, e)
			*f = int8(__f)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
		},
	}.Do()
}

// ParamRef7Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
// @param e: No description available.
// @param f: No description available.
// @param g: No description available.
func ParamRef7Callback(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	__d := *(*C.Vector4)(unsafe.Pointer(d))
	__e := plugify.ConstructVectorInt64(*e)
	__f := C.int8_t(*f)
	__g := plugify.ConstructString(*g)
	plugify.Block{
		Try: func() {
			C.ParamRef7Callback(&__a, &__b, &__c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), &__f, (*C.String)(unsafe.Pointer(&__g)))
			// Unmarshal - Convert native data to managed data.
			*a = int32(__a)
			*b = float32(__b)
			*c = float64(__c)
			*d = *(*plugify.Vector4)(unsafe.Pointer(&__d))
			plugify.GetVectorDataInt64To(&__e, e)
			*f = int8(__f)
			*g = plugify.GetStringData(&__g)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// ParamRef8Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
// @param e: No description available.
// @param f: No description available.
// @param g: No description available.
// @param h: No description available.
func ParamRef8Callback(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	__d := *(*C.Vector4)(unsafe.Pointer(d))
	__e := plugify.ConstructVectorInt64(*e)
	__f := C.int8_t(*f)
	__g := plugify.ConstructString(*g)
	__h := C.uint16_t(*h)
	plugify.Block{
		Try: func() {
			C.ParamRef8Callback(&__a, &__b, &__c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), &__f, (*C.String)(unsafe.Pointer(&__g)), &__h)
			// Unmarshal - Convert native data to managed data.
			*a = int32(__a)
			*b = float32(__b)
			*c = float64(__c)
			*d = *(*plugify.Vector4)(unsafe.Pointer(&__d))
			plugify.GetVectorDataInt64To(&__e, e)
			*f = int8(__f)
			*g = plugify.GetStringData(&__g)
			*h = uint16(__h)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// ParamRef9Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
// @param e: No description available.
// @param f: No description available.
// @param g: No description available.
// @param h: No description available.
// @param k: No description available.
func ParamRef9Callback(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	__d := *(*C.Vector4)(unsafe.Pointer(d))
	__e := plugify.ConstructVectorInt64(*e)
	__f := C.int8_t(*f)
	__g := plugify.ConstructString(*g)
	__h := C.uint16_t(*h)
	__k := C.int16_t(*k)
	plugify.Block{
		Try: func() {
			C.ParamRef9Callback(&__a, &__b, &__c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), &__f, (*C.String)(unsafe.Pointer(&__g)), &__h, &__k)
			// Unmarshal - Convert native data to managed data.
			*a = int32(__a)
			*b = float32(__b)
			*c = float64(__c)
			*d = *(*plugify.Vector4)(unsafe.Pointer(&__d))
			plugify.GetVectorDataInt64To(&__e, e)
			*f = int8(__f)
			*g = plugify.GetStringData(&__g)
			*h = uint16(__h)
			*k = int16(__k)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// ParamRef10Callback - No description provided.
// @param a: No description available.
// @param b: No description available.
// @param c: No description available.
// @param d: No description available.
// @param e: No description available.
// @param f: No description available.
// @param g: No description available.
// @param h: No description available.
// @param k: No description available.
// @param l: No description available.
func ParamRef10Callback(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16, l *uintptr) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	__d := *(*C.Vector4)(unsafe.Pointer(d))
	__e := plugify.ConstructVectorInt64(*e)
	__f := C.int8_t(*f)
	__g := plugify.ConstructString(*g)
	__h := C.uint16_t(*h)
	__k := C.int16_t(*k)
	__l := C.uintptr_t(*l)
	plugify.Block{
		Try: func() {
			C.ParamRef10Callback(&__a, &__b, &__c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), &__f, (*C.String)(unsafe.Pointer(&__g)), &__h, &__k, &__l)
			// Unmarshal - Convert native data to managed data.
			*a = int32(__a)
			*b = float32(__b)
			*c = float64(__c)
			*d = *(*plugify.Vector4)(unsafe.Pointer(&__d))
			plugify.GetVectorDataInt64To(&__e, e)
			*f = int8(__f)
			*g = plugify.GetStringData(&__g)
			*h = uint16(__h)
			*k = int16(__k)
			*l = uintptr(__l)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// ParamRefVectorsCallback - No description provided.
// @param p1: No description available.
// @param p2: No description available.
// @param p3: No description available.
// @param p4: No description available.
// @param p5: No description available.
// @param p6: No description available.
// @param p7: No description available.
// @param p8: No description available.
// @param p9: No description available.
// @param p10: No description available.
// @param p11: No description available.
// @param p12: No description available.
// @param p13: No description available.
// @param p14: No description available.
// @param p15: No description available.
func ParamRefVectorsCallback(p1 *[]bool, p2 *[]int8, p3 *[]uint16, p4 *[]int8, p5 *[]int16, p6 *[]int32, p7 *[]int64, p8 *[]uint8, p9 *[]uint16, p10 *[]uint32, p11 *[]uint64, p12 *[]uintptr, p13 *[]float32, p14 *[]float64, p15 *[]string) {
	__p1 := plugify.ConstructVectorBool(*p1)
	__p2 := plugify.ConstructVectorChar8(*p2)
	__p3 := plugify.ConstructVectorChar16(*p3)
	__p4 := plugify.ConstructVectorInt8(*p4)
	__p5 := plugify.ConstructVectorInt16(*p5)
	__p6 := plugify.ConstructVectorInt32(*p6)
	__p7 := plugify.ConstructVectorInt64(*p7)
	__p8 := plugify.ConstructVectorUInt8(*p8)
	__p9 := plugify.ConstructVectorUInt16(*p9)
	__p10 := plugify.ConstructVectorUInt32(*p10)
	__p11 := plugify.ConstructVectorUInt64(*p11)
	__p12 := plugify.ConstructVectorPointer(*p12)
	__p13 := plugify.ConstructVectorFloat(*p13)
	__p14 := plugify.ConstructVectorDouble(*p14)
	__p15 := plugify.ConstructVectorString(*p15)
	plugify.Block{
		Try: func() {
			C.ParamRefVectorsCallback((*C.Vector)(unsafe.Pointer(&__p1)), (*C.Vector)(unsafe.Pointer(&__p2)), (*C.Vector)(unsafe.Pointer(&__p3)), (*C.Vector)(unsafe.Pointer(&__p4)), (*C.Vector)(unsafe.Pointer(&__p5)), (*C.Vector)(unsafe.Pointer(&__p6)), (*C.Vector)(unsafe.Pointer(&__p7)), (*C.Vector)(unsafe.Pointer(&__p8)), (*C.Vector)(unsafe.Pointer(&__p9)), (*C.Vector)(unsafe.Pointer(&__p10)), (*C.Vector)(unsafe.Pointer(&__p11)), (*C.Vector)(unsafe.Pointer(&__p12)), (*C.Vector)(unsafe.Pointer(&__p13)), (*C.Vector)(unsafe.Pointer(&__p14)), (*C.Vector)(unsafe.Pointer(&__p15)))
			// Unmarshal - Convert native data to managed data.
			plugify.GetVectorDataBoolTo(&__p1, p1)
			plugify.GetVectorDataChar8To(&__p2, p2)
			plugify.GetVectorDataChar16To(&__p3, p3)
			plugify.GetVectorDataInt8To(&__p4, p4)
			plugify.GetVectorDataInt16To(&__p5, p5)
			plugify.GetVectorDataInt32To(&__p6, p6)
			plugify.GetVectorDataInt64To(&__p7, p7)
			plugify.GetVectorDataUInt8To(&__p8, p8)
			plugify.GetVectorDataUInt16To(&__p9, p9)
			plugify.GetVectorDataUInt32To(&__p10, p10)
			plugify.GetVectorDataUInt64To(&__p11, p11)
			plugify.GetVectorDataPointerTo(&__p12, p12)
			plugify.GetVectorDataFloatTo(&__p13, p13)
			plugify.GetVectorDataDoubleTo(&__p14, p14)
			plugify.GetVectorDataStringTo(&__p15, p15)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorBool(&__p1)
			plugify.DestroyVectorChar8(&__p2)
			plugify.DestroyVectorChar16(&__p3)
			plugify.DestroyVectorInt8(&__p4)
			plugify.DestroyVectorInt16(&__p5)
			plugify.DestroyVectorInt32(&__p6)
			plugify.DestroyVectorInt64(&__p7)
			plugify.DestroyVectorUInt8(&__p8)
			plugify.DestroyVectorUInt16(&__p9)
			plugify.DestroyVectorUInt32(&__p10)
			plugify.DestroyVectorUInt64(&__p11)
			plugify.DestroyVectorPointer(&__p12)
			plugify.DestroyVectorFloat(&__p13)
			plugify.DestroyVectorDouble(&__p14)
			plugify.DestroyVectorString(&__p15)
		},
	}.Do()
}

// ParamAllPrimitivesCallback - No description provided.
// @param p1: No description available.
// @param p2: No description available.
// @param p3: No description available.
// @param p4: No description available.
// @param p5: No description available.
// @param p6: No description available.
// @param p7: No description available.
// @param p8: No description available.
// @param p9: No description available.
// @param p10: No description available.
// @param p11: No description available.
// @param p12: No description available.
// @param p13: No description available.
// @param p14: No description available.
// @return No description available.
func ParamAllPrimitivesCallback(p1 bool, p2 int8, p3 uint16, p4 int8, p5 int16, p6 int32, p7 int64, p8 uint8, p9 uint16, p10 uint32, p11 uint64, p12 uintptr, p13 float32, p14 float64) int64 {
	var __retVal int64
	__p1 := C.bool(p1)
	__p2 := C.int8_t(p2)
	__p3 := C.uint16_t(p3)
	__p4 := C.int8_t(p4)
	__p5 := C.int16_t(p5)
	__p6 := C.int32_t(p6)
	__p7 := C.int64_t(p7)
	__p8 := C.uint8_t(p8)
	__p9 := C.uint16_t(p9)
	__p10 := C.uint32_t(p10)
	__p11 := C.uint64_t(p11)
	__p12 := C.uintptr_t(p12)
	__p13 := C.float(p13)
	__p14 := C.double(p14)
	__retVal = int64(C.ParamAllPrimitivesCallback(__p1, __p2, __p3, __p4, __p5, __p6, __p7, __p8, __p9, __p10, __p11, __p12, __p13, __p14))
	return __retVal
}

// ParamEnumCallback - No description provided.
// @param p1: No description available.
// @param p2: No description available.
// @return No description available.
func ParamEnumCallback(p1 Example, p2 []Example) int32 {
	var __retVal int32
	__p1 := C.int32_t(p1)
	__p2 := plugify.ConstructVectorInt32(p2)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.ParamEnumCallback(__p1, (*C.Vector)(unsafe.Pointer(&__p2))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__p2)
		},
	}.Do()
	return __retVal
}

// ParamEnumRefCallback - No description provided.
// @param p1: No description available.
// @param p2: No description available.
// @return No description available.
func ParamEnumRefCallback(p1 *Example, p2 *[]Example) int32 {
	var __retVal int32
	__p1 := C.int32_t(*p1)
	__p2 := plugify.ConstructVectorInt32(*p2)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.ParamEnumRefCallback(&__p1, (*C.Vector)(unsafe.Pointer(&__p2))))
			// Unmarshal - Convert native data to managed data.
			*p1 = Example(__p1)
			plugify.GetVectorDataInt32To(&__p2, p2)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__p2)
		},
	}.Do()
	return __retVal
}

// ParamVariantCallback - No description provided.
// @param p1: No description available.
// @param p2: No description available.
func ParamVariantCallback(p1 interface{}, p2 []interface{}) {
	__p1 := plugify.ConstructVariant(p1)
	__p2 := plugify.ConstructVectorVariant(p2)
	plugify.Block{
		Try: func() {
			C.ParamVariantCallback((*C.Variant)(unsafe.Pointer(&__p1)), (*C.Vector)(unsafe.Pointer(&__p2)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__p1)
			plugify.DestroyVectorVariant(&__p2)
		},
	}.Do()
}

// ParamVariantRefCallback - No description provided.
// @param p1: No description available.
// @param p2: No description available.
func ParamVariantRefCallback(p1 *interface{}, p2 *[]interface{}) {
	__p1 := plugify.ConstructVariant(*p1)
	__p2 := plugify.ConstructVectorVariant(*p2)
	plugify.Block{
		Try: func() {
			C.ParamVariantRefCallback((*C.Variant)(unsafe.Pointer(&__p1)), (*C.Vector)(unsafe.Pointer(&__p2)))
			// Unmarshal - Convert native data to managed data.
			*p1 = plugify.GetVariantData(&__p1)
			plugify.GetVectorDataVariantTo(&__p2, p2)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__p1)
			plugify.DestroyVectorVariant(&__p2)
		},
	}.Do()
}

// CallFuncVoidCallback - No description provided.
// @param func: No description available.
func CallFuncVoidCallback(func_ FuncVoid) {
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	C.CallFuncVoidCallback(__func_)
}

// CallFuncBoolCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncBoolCallback(func_ FuncBool) bool {
	var __retVal bool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = bool(C.CallFuncBoolCallback(__func_))
	return __retVal
}

// CallFuncChar8Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncChar8Callback(func_ FuncChar8) int8 {
	var __retVal int8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int8(C.CallFuncChar8Callback(__func_))
	return __retVal
}

// CallFuncChar16Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncChar16Callback(func_ FuncChar16) uint16 {
	var __retVal uint16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint16(C.CallFuncChar16Callback(__func_))
	return __retVal
}

// CallFuncInt8Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncInt8Callback(func_ FuncInt8) int8 {
	var __retVal int8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int8(C.CallFuncInt8Callback(__func_))
	return __retVal
}

// CallFuncInt16Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncInt16Callback(func_ FuncInt16) int16 {
	var __retVal int16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int16(C.CallFuncInt16Callback(__func_))
	return __retVal
}

// CallFuncInt32Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncInt32Callback(func_ FuncInt32) int32 {
	var __retVal int32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int32(C.CallFuncInt32Callback(__func_))
	return __retVal
}

// CallFuncInt64Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncInt64Callback(func_ FuncInt64) int64 {
	var __retVal int64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int64(C.CallFuncInt64Callback(__func_))
	return __retVal
}

// CallFuncUInt8Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncUInt8Callback(func_ FuncUInt8) uint8 {
	var __retVal uint8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint8(C.CallFuncUInt8Callback(__func_))
	return __retVal
}

// CallFuncUInt16Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncUInt16Callback(func_ FuncUInt16) uint16 {
	var __retVal uint16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint16(C.CallFuncUInt16Callback(__func_))
	return __retVal
}

// CallFuncUInt32Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncUInt32Callback(func_ FuncUInt32) uint32 {
	var __retVal uint32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint32(C.CallFuncUInt32Callback(__func_))
	return __retVal
}

// CallFuncUInt64Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncUInt64Callback(func_ FuncUInt64) uint64 {
	var __retVal uint64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint64(C.CallFuncUInt64Callback(__func_))
	return __retVal
}

// CallFuncPtrCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncPtrCallback(func_ FuncPtr) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFuncPtrCallback(__func_))
	return __retVal
}

// CallFuncFloatCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncFloatCallback(func_ FuncFloat) float32 {
	var __retVal float32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = float32(C.CallFuncFloatCallback(__func_))
	return __retVal
}

// CallFuncDoubleCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncDoubleCallback(func_ FuncDouble) float64 {
	var __retVal float64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = float64(C.CallFuncDoubleCallback(__func_))
	return __retVal
}

// CallFuncStringCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncStringCallback(func_ FuncString) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncStringCallback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAnyCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncAnyCallback(func_ FuncAny) interface{} {
	var __retVal interface{}
	var __retVal_native plugify.PlgVariant
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAnyCallback(__func_)
			__retVal_native = *(*plugify.PlgVariant)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVariantData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncFunctionCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncFunctionCallback(func_ FuncFunction) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFuncFunctionCallback(__func_))
	return __retVal
}

// CallFuncBoolVectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncBoolVectorCallback(func_ FuncBoolVector) []bool {
	var __retVal []bool
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncBoolVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataBool(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorBool(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncChar8VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncChar8VectorCallback(func_ FuncChar8Vector) []int8 {
	var __retVal []int8
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncChar8VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataChar8(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorChar8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncChar16VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncChar16VectorCallback(func_ FuncChar16Vector) []uint16 {
	var __retVal []uint16
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncChar16VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataChar16(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorChar16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncInt8VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncInt8VectorCallback(func_ FuncInt8Vector) []int8 {
	var __retVal []int8
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncInt8VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt8(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncInt16VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncInt16VectorCallback(func_ FuncInt16Vector) []int16 {
	var __retVal []int16
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncInt16VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt16(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncInt32VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncInt32VectorCallback(func_ FuncInt32Vector) []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncInt32VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt32(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncInt64VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncInt64VectorCallback(func_ FuncInt64Vector) []int64 {
	var __retVal []int64
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncInt64VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt64(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncUInt8VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncUInt8VectorCallback(func_ FuncUInt8Vector) []uint8 {
	var __retVal []uint8
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncUInt8VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt8(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncUInt16VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncUInt16VectorCallback(func_ FuncUInt16Vector) []uint16 {
	var __retVal []uint16
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncUInt16VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt16(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncUInt32VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncUInt32VectorCallback(func_ FuncUInt32Vector) []uint32 {
	var __retVal []uint32
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncUInt32VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt32(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncUInt64VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncUInt64VectorCallback(func_ FuncUInt64Vector) []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncUInt64VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt64(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncPtrVectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncPtrVectorCallback(func_ FuncPtrVector) []uintptr {
	var __retVal []uintptr
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncPtrVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataPointer(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorPointer(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncFloatVectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncFloatVectorCallback(func_ FuncFloatVector) []float32 {
	var __retVal []float32
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncFloatVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataFloat(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorFloat(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncDoubleVectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncDoubleVectorCallback(func_ FuncDoubleVector) []float64 {
	var __retVal []float64
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncDoubleVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataDouble(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorDouble(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncStringVectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncStringVectorCallback(func_ FuncStringVector) []string {
	var __retVal []string
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncStringVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataString(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAnyVectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncAnyVectorCallback(func_ FuncAnyVector) []interface{} {
	var __retVal []interface{}
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAnyVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVariant(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVariant(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncVec2VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncVec2VectorCallback(func_ FuncVec2Vector) []plugify.Vector2 {
	var __retVal []plugify.Vector2
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncVec2VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector2(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector2(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncVec3VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncVec3VectorCallback(func_ FuncVec3Vector) []plugify.Vector3 {
	var __retVal []plugify.Vector3
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncVec3VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector3(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector3(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncVec4VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncVec4VectorCallback(func_ FuncVec4Vector) []plugify.Vector4 {
	var __retVal []plugify.Vector4
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncVec4VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector4(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector4(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncMat4x4VectorCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncMat4x4VectorCallback(func_ FuncMat4x4Vector) []plugify.Matrix4x4 {
	var __retVal []plugify.Matrix4x4
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncMat4x4VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataMatrix4x4(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorMatrix4x4(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncVec2Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncVec2Callback(func_ FuncVec2) plugify.Vector2 {
	var __retVal plugify.Vector2
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncVec2Callback(__func_)
	__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncVec3Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncVec3Callback(func_ FuncVec3) plugify.Vector3 {
	var __retVal plugify.Vector3
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncVec3Callback(__func_)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncVec4Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncVec4Callback(func_ FuncVec4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncVec4Callback(__func_)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncMat4x4Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncMat4x4Callback(func_ FuncMat4x4) plugify.Matrix4x4 {
	var __retVal plugify.Matrix4x4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncMat4x4Callback(__func_)
	__retVal = *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFunc1Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc1Callback(func_ Func1) int32 {
	var __retVal int32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int32(C.CallFunc1Callback(__func_))
	return __retVal
}

// CallFunc2Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc2Callback(func_ Func2) int8 {
	var __retVal int8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int8(C.CallFunc2Callback(__func_))
	return __retVal
}

// CallFunc3Callback - No description provided.
// @param func: No description available.
func CallFunc3Callback(func_ Func3) {
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	C.CallFunc3Callback(__func_)
}

// CallFunc4Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc4Callback(func_ Func4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFunc4Callback(__func_)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFunc5Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc5Callback(func_ Func5) bool {
	var __retVal bool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = bool(C.CallFunc5Callback(__func_))
	return __retVal
}

// CallFunc6Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc6Callback(func_ Func6) int64 {
	var __retVal int64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int64(C.CallFunc6Callback(__func_))
	return __retVal
}

// CallFunc7Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc7Callback(func_ Func7) float64 {
	var __retVal float64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = float64(C.CallFunc7Callback(__func_))
	return __retVal
}

// CallFunc8Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc8Callback(func_ Func8) plugify.Matrix4x4 {
	var __retVal plugify.Matrix4x4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFunc8Callback(__func_)
	__retVal = *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFunc9Callback - No description provided.
// @param func: No description available.
func CallFunc9Callback(func_ Func9) {
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	C.CallFunc9Callback(__func_)
}

// CallFunc10Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc10Callback(func_ Func10) uint32 {
	var __retVal uint32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint32(C.CallFunc10Callback(__func_))
	return __retVal
}

// CallFunc11Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc11Callback(func_ Func11) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFunc11Callback(__func_))
	return __retVal
}

// CallFunc12Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc12Callback(func_ Func12) bool {
	var __retVal bool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = bool(C.CallFunc12Callback(__func_))
	return __retVal
}

// CallFunc13Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc13Callback(func_ Func13) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc13Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc14Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc14Callback(func_ Func14) []string {
	var __retVal []string
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc14Callback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataString(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc15Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc15Callback(func_ Func15) int16 {
	var __retVal int16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int16(C.CallFunc15Callback(__func_))
	return __retVal
}

// CallFunc16Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc16Callback(func_ Func16) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFunc16Callback(__func_))
	return __retVal
}

// CallFunc17Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc17Callback(func_ Func17) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc17Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc18Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc18Callback(func_ Func18) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc18Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc19Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc19Callback(func_ Func19) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc19Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc20Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc20Callback(func_ Func20) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc20Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc21Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc21Callback(func_ Func21) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc21Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc22Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc22Callback(func_ Func22) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc22Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc23Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc23Callback(func_ Func23) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc23Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc24Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc24Callback(func_ Func24) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc24Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc25Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc25Callback(func_ Func25) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc25Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc26Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc26Callback(func_ Func26) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc26Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc27Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc27Callback(func_ Func27) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc27Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc28Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc28Callback(func_ Func28) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc28Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc29Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc29Callback(func_ Func29) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc29Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc30Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc30Callback(func_ Func30) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc30Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc31Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc31Callback(func_ Func31) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc31Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc32Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc32Callback(func_ Func32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc32Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc33Callback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFunc33Callback(func_ Func33) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc33Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncEnumCallback - No description provided.
// @param func: No description available.
// @return No description available.
func CallFuncEnumCallback(func_ FuncEnum) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncEnumCallback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}
