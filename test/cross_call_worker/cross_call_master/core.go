package cross_call_master

/*
#include "core.h"
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
	"errors"
	"github.com/untrustedmodders/go-plugify"
	"reflect"
	"runtime"
	"unsafe"
)

var _ = errors.New("")
var _ = reflect.TypeOf(0)
var _ = runtime.GOOS
var _ = unsafe.Sizeof(0)
var _ = plugify.Plugin.Loaded

// Generated from cross_call_master (group: core)

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

func NoParamReturnVoidCallback() {
	C.NoParamReturnVoidCallback()
}

func NoParamReturnBoolCallback() bool {
	__retVal := bool(C.NoParamReturnBoolCallback())
	return __retVal
}

func NoParamReturnChar8Callback() int8 {
	__retVal := int8(C.NoParamReturnChar8Callback())
	return __retVal
}

func NoParamReturnChar16Callback() uint16 {
	__retVal := uint16(C.NoParamReturnChar16Callback())
	return __retVal
}

func NoParamReturnInt8Callback() int8 {
	__retVal := int8(C.NoParamReturnInt8Callback())
	return __retVal
}

func NoParamReturnInt16Callback() int16 {
	__retVal := int16(C.NoParamReturnInt16Callback())
	return __retVal
}

func NoParamReturnInt32Callback() int32 {
	__retVal := int32(C.NoParamReturnInt32Callback())
	return __retVal
}

func NoParamReturnInt64Callback() int64 {
	__retVal := int64(C.NoParamReturnInt64Callback())
	return __retVal
}

func NoParamReturnUInt8Callback() uint8 {
	__retVal := uint8(C.NoParamReturnUInt8Callback())
	return __retVal
}

func NoParamReturnUInt16Callback() uint16 {
	__retVal := uint16(C.NoParamReturnUInt16Callback())
	return __retVal
}

func NoParamReturnUInt32Callback() uint32 {
	__retVal := uint32(C.NoParamReturnUInt32Callback())
	return __retVal
}

func NoParamReturnUInt64Callback() uint64 {
	__retVal := uint64(C.NoParamReturnUInt64Callback())
	return __retVal
}

func NoParamReturnPointerCallback() uintptr {
	__retVal := uintptr(C.NoParamReturnPointerCallback())
	return __retVal
}

func NoParamReturnFloatCallback() float32 {
	__retVal := float32(C.NoParamReturnFloatCallback())
	return __retVal
}

func NoParamReturnDoubleCallback() float64 {
	__retVal := float64(C.NoParamReturnDoubleCallback())
	return __retVal
}

func NoParamReturnFunctionCallback() NoParamReturnFunctionCallbackFunc {
	__retVal := plugify.GetDelegateForFunctionPointer(C.NoParamReturnFunctionCallback(), reflect.TypeOf(NoParamReturnFunctionCallbackFunc(nil))).(NoParamReturnFunctionCallbackFunc)
	return __retVal
}

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

func NoParamReturnAnyCallback() any {
	var __retVal any
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

func NoParamReturnArrayAnyCallback() []any {
	var __retVal []any
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

func NoParamReturnVector2Callback() plugify.Vector2 {
	__native := C.NoParamReturnVector2Callback()
	__retVal := *(*plugify.Vector2)(unsafe.Pointer(&__native))
	return __retVal
}

func NoParamReturnVector3Callback() plugify.Vector3 {
	__native := C.NoParamReturnVector3Callback()
	__retVal := *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

func NoParamReturnVector4Callback() plugify.Vector4 {
	__native := C.NoParamReturnVector4Callback()
	__retVal := *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

func NoParamReturnMatrix4x4Callback() plugify.Matrix4x4 {
	__native := C.NoParamReturnMatrix4x4Callback()
	__retVal := *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

func Param1Callback(a int32) {
	__a := C.int32_t(a)
	C.Param1Callback(__a)
}

func Param2Callback(a int32, b float32) {
	__a := C.int32_t(a)
	__b := C.float(b)
	C.Param2Callback(__a, __b)
}

func Param3Callback(a int32, b float32, c float64) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	C.Param3Callback(__a, __b, __c)
}

func Param4Callback(a int32, b float32, c float64, d plugify.Vector4) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	C.Param4Callback(__a, __b, __c, &__d)
}

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

func ParamRef1Callback(a *int32) {
	__a := C.int32_t(*a)
	C.ParamRef1Callback(&__a)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
}

func ParamRef2Callback(a *int32, b *float32) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	C.ParamRef2Callback(&__a, &__b)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
	*b = float32(__b)
}

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

func ParamVariantCallback(p1 any, p2 []any) {
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

func ParamVariantRefCallback(p1 *any, p2 *[]any) {
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

func CallFuncVoidCallback(func_ FuncVoid) {
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	C.CallFuncVoidCallback(__func_)
}

func CallFuncBoolCallback(func_ FuncBool) bool {
	var __retVal bool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = bool(C.CallFuncBoolCallback(__func_))
	return __retVal
}

func CallFuncChar8Callback(func_ FuncChar8) int8 {
	var __retVal int8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int8(C.CallFuncChar8Callback(__func_))
	return __retVal
}

func CallFuncChar16Callback(func_ FuncChar16) uint16 {
	var __retVal uint16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint16(C.CallFuncChar16Callback(__func_))
	return __retVal
}

func CallFuncInt8Callback(func_ FuncInt8) int8 {
	var __retVal int8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int8(C.CallFuncInt8Callback(__func_))
	return __retVal
}

func CallFuncInt16Callback(func_ FuncInt16) int16 {
	var __retVal int16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int16(C.CallFuncInt16Callback(__func_))
	return __retVal
}

func CallFuncInt32Callback(func_ FuncInt32) int32 {
	var __retVal int32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int32(C.CallFuncInt32Callback(__func_))
	return __retVal
}

func CallFuncInt64Callback(func_ FuncInt64) int64 {
	var __retVal int64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int64(C.CallFuncInt64Callback(__func_))
	return __retVal
}

func CallFuncUInt8Callback(func_ FuncUInt8) uint8 {
	var __retVal uint8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint8(C.CallFuncUInt8Callback(__func_))
	return __retVal
}

func CallFuncUInt16Callback(func_ FuncUInt16) uint16 {
	var __retVal uint16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint16(C.CallFuncUInt16Callback(__func_))
	return __retVal
}

func CallFuncUInt32Callback(func_ FuncUInt32) uint32 {
	var __retVal uint32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint32(C.CallFuncUInt32Callback(__func_))
	return __retVal
}

func CallFuncUInt64Callback(func_ FuncUInt64) uint64 {
	var __retVal uint64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint64(C.CallFuncUInt64Callback(__func_))
	return __retVal
}

func CallFuncPtrCallback(func_ FuncPtr) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFuncPtrCallback(__func_))
	return __retVal
}

func CallFuncFloatCallback(func_ FuncFloat) float32 {
	var __retVal float32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = float32(C.CallFuncFloatCallback(__func_))
	return __retVal
}

func CallFuncDoubleCallback(func_ FuncDouble) float64 {
	var __retVal float64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = float64(C.CallFuncDoubleCallback(__func_))
	return __retVal
}

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

func CallFuncAnyCallback(func_ FuncAny) any {
	var __retVal any
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

func CallFuncFunctionCallback(func_ FuncFunction) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFuncFunctionCallback(__func_))
	return __retVal
}

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

func CallFuncAnyVectorCallback(func_ FuncAnyVector) []any {
	var __retVal []any
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

func CallFuncVec2Callback(func_ FuncVec2) plugify.Vector2 {
	var __retVal plugify.Vector2
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncVec2Callback(__func_)
	__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
	return __retVal
}

func CallFuncVec3Callback(func_ FuncVec3) plugify.Vector3 {
	var __retVal plugify.Vector3
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncVec3Callback(__func_)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

func CallFuncVec4Callback(func_ FuncVec4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncVec4Callback(__func_)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

func CallFuncMat4x4Callback(func_ FuncMat4x4) plugify.Matrix4x4 {
	var __retVal plugify.Matrix4x4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncMat4x4Callback(__func_)
	__retVal = *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

func CallFunc1Callback(func_ Func1) int32 {
	var __retVal int32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int32(C.CallFunc1Callback(__func_))
	return __retVal
}

func CallFunc2Callback(func_ Func2) int8 {
	var __retVal int8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int8(C.CallFunc2Callback(__func_))
	return __retVal
}

func CallFunc3Callback(func_ Func3) {
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	C.CallFunc3Callback(__func_)
}

func CallFunc4Callback(func_ Func4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFunc4Callback(__func_)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

func CallFunc5Callback(func_ Func5) bool {
	var __retVal bool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = bool(C.CallFunc5Callback(__func_))
	return __retVal
}

func CallFunc6Callback(func_ Func6) int64 {
	var __retVal int64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int64(C.CallFunc6Callback(__func_))
	return __retVal
}

func CallFunc7Callback(func_ Func7) float64 {
	var __retVal float64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = float64(C.CallFunc7Callback(__func_))
	return __retVal
}

func CallFunc8Callback(func_ Func8) plugify.Matrix4x4 {
	var __retVal plugify.Matrix4x4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFunc8Callback(__func_)
	__retVal = *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

func CallFunc9Callback(func_ Func9) {
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	C.CallFunc9Callback(__func_)
}

func CallFunc10Callback(func_ Func10) uint32 {
	var __retVal uint32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint32(C.CallFunc10Callback(__func_))
	return __retVal
}

func CallFunc11Callback(func_ Func11) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFunc11Callback(__func_))
	return __retVal
}

func CallFunc12Callback(func_ Func12) bool {
	var __retVal bool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = bool(C.CallFunc12Callback(__func_))
	return __retVal
}

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

func CallFunc15Callback(func_ Func15) int16 {
	var __retVal int16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int16(C.CallFunc15Callback(__func_))
	return __retVal
}

func CallFunc16Callback(func_ Func16) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFunc16Callback(__func_))
	return __retVal
}

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

var (
	ResourceHandleErrEmptyHandle = errors.New("ResourceHandle: empty handle")
)

// ResourceHandle - RAII wrapper for ResourceHandle pointer
type ResourceHandle struct {
	handle    uintptr
	cleanup   runtime.Cleanup
	ownership ownership
	noCopy    noCopy
}

func NewResourceHandleResourceHandleCreate(id int32, name string) *ResourceHandle {
	return NewResourceHandleOwned(ResourceHandleCreate(id, name))
}

func NewResourceHandleResourceHandleCreateDefault() *ResourceHandle {
	return NewResourceHandleOwned(ResourceHandleCreateDefault())
}

// NewResourceHandleBorrowed creates a ResourceHandle from a borrowed handle
func NewResourceHandleBorrowed(handle uintptr) *ResourceHandle {
	if handle == 0 {
		return &ResourceHandle{}
	}
	return &ResourceHandle{
		handle:    handle,
		ownership: Borrowed,
	}
}

// NewResourceHandleOwned creates a ResourceHandle from an owned handle
func NewResourceHandleOwned(handle uintptr) *ResourceHandle {
	if handle == 0 {
		return &ResourceHandle{}
	}
	w := &ResourceHandle{
		handle:    handle,
		ownership: Owned,
	}
	w.cleanup = runtime.AddCleanup(w, w.finalize, struct{}{})
	return w
}

// finalize is the finalizer function (like C++ destructor)
func (w *ResourceHandle) finalize(_ struct{}) {
	if plugify.Plugin.Loaded {
		w.destroy()
	}
}

// destroy cleans up owned handles
func (w *ResourceHandle) destroy() {
	if w.handle != 0 && w.ownership == Owned {
		ResourceHandleDestroy(w.handle)
	}
}

// nullify resets the handle
func (w *ResourceHandle) nullify() {
	w.handle = 0
	w.ownership = Borrowed
}

// Close explicitly destroys the handle (like C++ destructor, but manual)
func (w *ResourceHandle) Close() {
	w.Reset()
}

// Get returns the underlying handle
func (w *ResourceHandle) Get() uintptr {
	return w.handle
}

// Release releases ownership and returns the handle
func (w *ResourceHandle) Release() uintptr {
	if w.ownership == Owned {
		w.cleanup.Stop()
	}
	handle := w.handle
	w.nullify()
	return handle
}

// Reset destroys and resets the handle
func (w *ResourceHandle) Reset() {
	if w.ownership == Owned {
		w.cleanup.Stop()
	}
	w.destroy()
	w.nullify()
}

// IsValid returns true if handle is not nil
func (w *ResourceHandle) IsValid() bool {
	return w.handle != 0
}

func (w *ResourceHandle) GetId() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, ResourceHandleErrEmptyHandle
	}
	return ResourceHandleGetId(w.handle), nil
}

func (w *ResourceHandle) GetName() (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, ResourceHandleErrEmptyHandle
	}
	return ResourceHandleGetName(w.handle), nil
}

func (w *ResourceHandle) SetName(name string) error {
	if w.handle == 0 {
		return ResourceHandleErrEmptyHandle
	}
	ResourceHandleSetName(w.handle, name)
	return nil
}

func (w *ResourceHandle) IncrementCounter() error {
	if w.handle == 0 {
		return ResourceHandleErrEmptyHandle
	}
	ResourceHandleIncrementCounter(w.handle)
	return nil
}

func (w *ResourceHandle) GetCounter() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, ResourceHandleErrEmptyHandle
	}
	return ResourceHandleGetCounter(w.handle), nil
}

func (w *ResourceHandle) AddData(value float32) error {
	if w.handle == 0 {
		return ResourceHandleErrEmptyHandle
	}
	ResourceHandleAddData(w.handle, value)
	return nil
}

func (w *ResourceHandle) GetData() ([]float32, error) {
	if w.handle == 0 {
		var zero []float32
		return zero, ResourceHandleErrEmptyHandle
	}
	return ResourceHandleGetData(w.handle), nil
}

func (w *ResourceHandle) GetAliveCount() int32 {
	return ResourceHandleGetAliveCount()
}

func (w *ResourceHandle) GetTotalCreated() int32 {
	return ResourceHandleGetTotalCreated()
}

func (w *ResourceHandle) GetTotalDestroyed() int32 {
	return ResourceHandleGetTotalDestroyed()
}

var (
	CounterErrEmptyHandle = errors.New("Counter: empty handle")
)

type Counter struct {
	handle uintptr
}

func NewCounterCounterCreate(initialValue int64) *Counter {
	return &Counter{
		handle: CounterCreate(initialValue),
	}
}

func NewCounterCounterCreateZero() *Counter {
	return &Counter{
		handle: CounterCreateZero(),
	}
}

// NewCounter creates a Counter from a handle
func NewCounter(handle uintptr) *Counter {
	return &Counter{
		handle: handle,
	}
}

// Get returns the underlying handle
func (w *Counter) Get() uintptr {
	return w.handle
}

// Release releases ownership and returns the handle
func (w *Counter) Release() uintptr {
	handle := w.handle
	w.handle = 0
	return handle
}

// Reset destroys and resets the handle
func (w *Counter) Reset() {
	w.handle = 0
}

// IsValid returns true if handle is not nil
func (w *Counter) IsValid() bool {
	return w.handle != 0
}

func (w *Counter) GetValue() (int64, error) {
	if w.handle == 0 {
		var zero int64
		return zero, CounterErrEmptyHandle
	}
	return CounterGetValue(w.handle), nil
}

func (w *Counter) SetValue(value int64) error {
	if w.handle == 0 {
		return CounterErrEmptyHandle
	}
	CounterSetValue(w.handle, value)
	return nil
}

func (w *Counter) Increment() error {
	if w.handle == 0 {
		return CounterErrEmptyHandle
	}
	CounterIncrement(w.handle)
	return nil
}

func (w *Counter) Decrement() error {
	if w.handle == 0 {
		return CounterErrEmptyHandle
	}
	CounterDecrement(w.handle)
	return nil
}

func (w *Counter) Add(amount int64) error {
	if w.handle == 0 {
		return CounterErrEmptyHandle
	}
	CounterAdd(w.handle, amount)
	return nil
}

func (w *Counter) Reset2() error {
	if w.handle == 0 {
		return CounterErrEmptyHandle
	}
	CounterReset(w.handle)
	return nil
}

func (w *Counter) IsPositive() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CounterErrEmptyHandle
	}
	return CounterIsPositive(w.handle), nil
}

func (w *Counter) Compare(value1 int64, value2 int64) int32 {
	return CounterCompare(value1, value2)
}

func (w *Counter) Sum(values []int64) int64 {
	return CounterSum(values)
}
