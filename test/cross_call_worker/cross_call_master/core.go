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
#cgo noescape ParamAllAliasesCallback
#cgo noescape ParamAllRefAliasesCallback
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
#cgo noescape CallFuncAliasBoolCallback
#cgo noescape CallFuncAliasChar8Callback
#cgo noescape CallFuncAliasChar16Callback
#cgo noescape CallFuncAliasInt8Callback
#cgo noescape CallFuncAliasInt16Callback
#cgo noescape CallFuncAliasInt32Callback
#cgo noescape CallFuncAliasInt64Callback
#cgo noescape CallFuncAliasUInt8Callback
#cgo noescape CallFuncAliasUInt16Callback
#cgo noescape CallFuncAliasUInt32Callback
#cgo noescape CallFuncAliasUInt64Callback
#cgo noescape CallFuncAliasPtrCallback
#cgo noescape CallFuncAliasFloatCallback
#cgo noescape CallFuncAliasDoubleCallback
#cgo noescape CallFuncAliasStringCallback
#cgo noescape CallFuncAliasAnyCallback
#cgo noescape CallFuncAliasFunctionCallback
#cgo noescape CallFuncAliasBoolVectorCallback
#cgo noescape CallFuncAliasChar8VectorCallback
#cgo noescape CallFuncAliasChar16VectorCallback
#cgo noescape CallFuncAliasInt8VectorCallback
#cgo noescape CallFuncAliasInt16VectorCallback
#cgo noescape CallFuncAliasInt32VectorCallback
#cgo noescape CallFuncAliasInt64VectorCallback
#cgo noescape CallFuncAliasUInt8VectorCallback
#cgo noescape CallFuncAliasUInt16VectorCallback
#cgo noescape CallFuncAliasUInt32VectorCallback
#cgo noescape CallFuncAliasUInt64VectorCallback
#cgo noescape CallFuncAliasPtrVectorCallback
#cgo noescape CallFuncAliasFloatVectorCallback
#cgo noescape CallFuncAliasDoubleVectorCallback
#cgo noescape CallFuncAliasStringVectorCallback
#cgo noescape CallFuncAliasAnyVectorCallback
#cgo noescape CallFuncAliasVec2VectorCallback
#cgo noescape CallFuncAliasVec3VectorCallback
#cgo noescape CallFuncAliasVec4VectorCallback
#cgo noescape CallFuncAliasMat4x4VectorCallback
#cgo noescape CallFuncAliasVec2Callback
#cgo noescape CallFuncAliasVec3Callback
#cgo noescape CallFuncAliasVec4Callback
#cgo noescape CallFuncAliasMat4x4Callback
#cgo noescape CallFuncAliasAllCallback
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
var _ = plugify.ApiVersion

// Generated from cross_call_master (group: core)

var _ReverseReturn = func(returnString string) {
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

// ReverseReturn
func ReverseReturn(returnString string) {
	defer plugify.Scope("cross_call_master::ReverseReturn", ModuleName, 3)()
	_ReverseReturn(returnString)
}

var _NoParamReturnVoidCallback = func() {
	C.NoParamReturnVoidCallback()
}

// NoParamReturnVoidCallback
func NoParamReturnVoidCallback() {
	defer plugify.Scope("cross_call_master::NoParamReturnVoidCallback", ModuleName, 3)()
	_NoParamReturnVoidCallback()
}

var _NoParamReturnBoolCallback = func() bool {
	__retVal := bool(C.NoParamReturnBoolCallback())
	return __retVal
}

// NoParamReturnBoolCallback
func NoParamReturnBoolCallback() bool {
	defer plugify.Scope("cross_call_master::NoParamReturnBoolCallback", ModuleName, 3)()
	return _NoParamReturnBoolCallback()
}

var _NoParamReturnChar8Callback = func() int8 {
	__retVal := int8(C.NoParamReturnChar8Callback())
	return __retVal
}

// NoParamReturnChar8Callback
func NoParamReturnChar8Callback() int8 {
	defer plugify.Scope("cross_call_master::NoParamReturnChar8Callback", ModuleName, 3)()
	return _NoParamReturnChar8Callback()
}

var _NoParamReturnChar16Callback = func() uint16 {
	__retVal := uint16(C.NoParamReturnChar16Callback())
	return __retVal
}

// NoParamReturnChar16Callback
func NoParamReturnChar16Callback() uint16 {
	defer plugify.Scope("cross_call_master::NoParamReturnChar16Callback", ModuleName, 3)()
	return _NoParamReturnChar16Callback()
}

var _NoParamReturnInt8Callback = func() int8 {
	__retVal := int8(C.NoParamReturnInt8Callback())
	return __retVal
}

// NoParamReturnInt8Callback
func NoParamReturnInt8Callback() int8 {
	defer plugify.Scope("cross_call_master::NoParamReturnInt8Callback", ModuleName, 3)()
	return _NoParamReturnInt8Callback()
}

var _NoParamReturnInt16Callback = func() int16 {
	__retVal := int16(C.NoParamReturnInt16Callback())
	return __retVal
}

// NoParamReturnInt16Callback
func NoParamReturnInt16Callback() int16 {
	defer plugify.Scope("cross_call_master::NoParamReturnInt16Callback", ModuleName, 3)()
	return _NoParamReturnInt16Callback()
}

var _NoParamReturnInt32Callback = func() int32 {
	__retVal := int32(C.NoParamReturnInt32Callback())
	return __retVal
}

// NoParamReturnInt32Callback
func NoParamReturnInt32Callback() int32 {
	defer plugify.Scope("cross_call_master::NoParamReturnInt32Callback", ModuleName, 3)()
	return _NoParamReturnInt32Callback()
}

var _NoParamReturnInt64Callback = func() int64 {
	__retVal := int64(C.NoParamReturnInt64Callback())
	return __retVal
}

// NoParamReturnInt64Callback
func NoParamReturnInt64Callback() int64 {
	defer plugify.Scope("cross_call_master::NoParamReturnInt64Callback", ModuleName, 3)()
	return _NoParamReturnInt64Callback()
}

var _NoParamReturnUInt8Callback = func() uint8 {
	__retVal := uint8(C.NoParamReturnUInt8Callback())
	return __retVal
}

// NoParamReturnUInt8Callback
func NoParamReturnUInt8Callback() uint8 {
	defer plugify.Scope("cross_call_master::NoParamReturnUInt8Callback", ModuleName, 3)()
	return _NoParamReturnUInt8Callback()
}

var _NoParamReturnUInt16Callback = func() uint16 {
	__retVal := uint16(C.NoParamReturnUInt16Callback())
	return __retVal
}

// NoParamReturnUInt16Callback
func NoParamReturnUInt16Callback() uint16 {
	defer plugify.Scope("cross_call_master::NoParamReturnUInt16Callback", ModuleName, 3)()
	return _NoParamReturnUInt16Callback()
}

var _NoParamReturnUInt32Callback = func() uint32 {
	__retVal := uint32(C.NoParamReturnUInt32Callback())
	return __retVal
}

// NoParamReturnUInt32Callback
func NoParamReturnUInt32Callback() uint32 {
	defer plugify.Scope("cross_call_master::NoParamReturnUInt32Callback", ModuleName, 3)()
	return _NoParamReturnUInt32Callback()
}

var _NoParamReturnUInt64Callback = func() uint64 {
	__retVal := uint64(C.NoParamReturnUInt64Callback())
	return __retVal
}

// NoParamReturnUInt64Callback
func NoParamReturnUInt64Callback() uint64 {
	defer plugify.Scope("cross_call_master::NoParamReturnUInt64Callback", ModuleName, 3)()
	return _NoParamReturnUInt64Callback()
}

var _NoParamReturnPointerCallback = func() uintptr {
	__retVal := uintptr(C.NoParamReturnPointerCallback())
	return __retVal
}

// NoParamReturnPointerCallback
func NoParamReturnPointerCallback() uintptr {
	defer plugify.Scope("cross_call_master::NoParamReturnPointerCallback", ModuleName, 3)()
	return _NoParamReturnPointerCallback()
}

var _NoParamReturnFloatCallback = func() float32 {
	__retVal := float32(C.NoParamReturnFloatCallback())
	return __retVal
}

// NoParamReturnFloatCallback
func NoParamReturnFloatCallback() float32 {
	defer plugify.Scope("cross_call_master::NoParamReturnFloatCallback", ModuleName, 3)()
	return _NoParamReturnFloatCallback()
}

var _NoParamReturnDoubleCallback = func() float64 {
	__retVal := float64(C.NoParamReturnDoubleCallback())
	return __retVal
}

// NoParamReturnDoubleCallback
func NoParamReturnDoubleCallback() float64 {
	defer plugify.Scope("cross_call_master::NoParamReturnDoubleCallback", ModuleName, 3)()
	return _NoParamReturnDoubleCallback()
}

var _NoParamReturnFunctionCallback = func() NoParamReturnFunctionCallbackFunc {
	__retVal := plugify.GetDelegateForFunctionPointer(C.NoParamReturnFunctionCallback(), reflect.TypeOf(NoParamReturnFunctionCallbackFunc(nil))).(NoParamReturnFunctionCallbackFunc)
	return __retVal
}

// NoParamReturnFunctionCallback
func NoParamReturnFunctionCallback() NoParamReturnFunctionCallbackFunc {
	defer plugify.Scope("cross_call_master::NoParamReturnFunctionCallback", ModuleName, 3)()
	return _NoParamReturnFunctionCallback()
}

var _NoParamReturnStringCallback = func() string {
	var __retVal string
	var __retVal_native plugify.PlgString
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnStringCallback()
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnStringCallback
func NoParamReturnStringCallback() string {
	defer plugify.Scope("cross_call_master::NoParamReturnStringCallback", ModuleName, 3)()
	return _NoParamReturnStringCallback()
}

var _NoParamReturnAnyCallback = func() any {
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

// NoParamReturnAnyCallback
func NoParamReturnAnyCallback() any {
	defer plugify.Scope("cross_call_master::NoParamReturnAnyCallback", ModuleName, 3)()
	return _NoParamReturnAnyCallback()
}

var _NoParamReturnArrayBoolCallback = func() []bool {
	var __retVal []bool
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayBoolCallback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataBool[bool](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorBool(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayBoolCallback
func NoParamReturnArrayBoolCallback() []bool {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayBoolCallback", ModuleName, 3)()
	return _NoParamReturnArrayBoolCallback()
}

var _NoParamReturnArrayChar8Callback = func() []int8 {
	var __retVal []int8
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayChar8Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataChar8[int8](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorChar8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayChar8Callback
func NoParamReturnArrayChar8Callback() []int8 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayChar8Callback", ModuleName, 3)()
	return _NoParamReturnArrayChar8Callback()
}

var _NoParamReturnArrayChar16Callback = func() []uint16 {
	var __retVal []uint16
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayChar16Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataChar16[uint16](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorChar16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayChar16Callback
func NoParamReturnArrayChar16Callback() []uint16 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayChar16Callback", ModuleName, 3)()
	return _NoParamReturnArrayChar16Callback()
}

var _NoParamReturnArrayInt8Callback = func() []int8 {
	var __retVal []int8
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayInt8Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt8[int8](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayInt8Callback
func NoParamReturnArrayInt8Callback() []int8 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayInt8Callback", ModuleName, 3)()
	return _NoParamReturnArrayInt8Callback()
}

var _NoParamReturnArrayInt16Callback = func() []int16 {
	var __retVal []int16
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayInt16Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt16[int16](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayInt16Callback
func NoParamReturnArrayInt16Callback() []int16 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayInt16Callback", ModuleName, 3)()
	return _NoParamReturnArrayInt16Callback()
}

var _NoParamReturnArrayInt32Callback = func() []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayInt32Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt32[int32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayInt32Callback
func NoParamReturnArrayInt32Callback() []int32 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayInt32Callback", ModuleName, 3)()
	return _NoParamReturnArrayInt32Callback()
}

var _NoParamReturnArrayInt64Callback = func() []int64 {
	var __retVal []int64
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayInt64Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt64[int64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayInt64Callback
func NoParamReturnArrayInt64Callback() []int64 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayInt64Callback", ModuleName, 3)()
	return _NoParamReturnArrayInt64Callback()
}

var _NoParamReturnArrayUInt8Callback = func() []uint8 {
	var __retVal []uint8
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayUInt8Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt8[uint8](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayUInt8Callback
func NoParamReturnArrayUInt8Callback() []uint8 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayUInt8Callback", ModuleName, 3)()
	return _NoParamReturnArrayUInt8Callback()
}

var _NoParamReturnArrayUInt16Callback = func() []uint16 {
	var __retVal []uint16
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayUInt16Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt16[uint16](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayUInt16Callback
func NoParamReturnArrayUInt16Callback() []uint16 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayUInt16Callback", ModuleName, 3)()
	return _NoParamReturnArrayUInt16Callback()
}

var _NoParamReturnArrayUInt32Callback = func() []uint32 {
	var __retVal []uint32
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayUInt32Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt32[uint32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayUInt32Callback
func NoParamReturnArrayUInt32Callback() []uint32 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayUInt32Callback", ModuleName, 3)()
	return _NoParamReturnArrayUInt32Callback()
}

var _NoParamReturnArrayUInt64Callback = func() []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayUInt64Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt64[uint64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayUInt64Callback
func NoParamReturnArrayUInt64Callback() []uint64 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayUInt64Callback", ModuleName, 3)()
	return _NoParamReturnArrayUInt64Callback()
}

var _NoParamReturnArrayPointerCallback = func() []uintptr {
	var __retVal []uintptr
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayPointerCallback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataPointer[uintptr](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorPointer(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayPointerCallback
func NoParamReturnArrayPointerCallback() []uintptr {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayPointerCallback", ModuleName, 3)()
	return _NoParamReturnArrayPointerCallback()
}

var _NoParamReturnArrayFloatCallback = func() []float32 {
	var __retVal []float32
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayFloatCallback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataFloat[float32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorFloat(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayFloatCallback
func NoParamReturnArrayFloatCallback() []float32 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayFloatCallback", ModuleName, 3)()
	return _NoParamReturnArrayFloatCallback()
}

var _NoParamReturnArrayDoubleCallback = func() []float64 {
	var __retVal []float64
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayDoubleCallback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataDouble[float64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorDouble(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayDoubleCallback
func NoParamReturnArrayDoubleCallback() []float64 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayDoubleCallback", ModuleName, 3)()
	return _NoParamReturnArrayDoubleCallback()
}

var _NoParamReturnArrayStringCallback = func() []string {
	var __retVal []string
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayStringCallback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataString[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayStringCallback
func NoParamReturnArrayStringCallback() []string {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayStringCallback", ModuleName, 3)()
	return _NoParamReturnArrayStringCallback()
}

var _NoParamReturnArrayAnyCallback = func() []any {
	var __retVal []any
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayAnyCallback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVariant[any, []any](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVariant(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayAnyCallback
func NoParamReturnArrayAnyCallback() []any {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayAnyCallback", ModuleName, 3)()
	return _NoParamReturnArrayAnyCallback()
}

var _NoParamReturnArrayVector2Callback = func() []plugify.Vector2 {
	var __retVal []plugify.Vector2
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayVector2Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector2[plugify.Vector2](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector2(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayVector2Callback
func NoParamReturnArrayVector2Callback() []plugify.Vector2 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayVector2Callback", ModuleName, 3)()
	return _NoParamReturnArrayVector2Callback()
}

var _NoParamReturnArrayVector3Callback = func() []plugify.Vector3 {
	var __retVal []plugify.Vector3
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayVector3Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector3[plugify.Vector3](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector3(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayVector3Callback
func NoParamReturnArrayVector3Callback() []plugify.Vector3 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayVector3Callback", ModuleName, 3)()
	return _NoParamReturnArrayVector3Callback()
}

var _NoParamReturnArrayVector4Callback = func() []plugify.Vector4 {
	var __retVal []plugify.Vector4
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayVector4Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector4[plugify.Vector4](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector4(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayVector4Callback
func NoParamReturnArrayVector4Callback() []plugify.Vector4 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayVector4Callback", ModuleName, 3)()
	return _NoParamReturnArrayVector4Callback()
}

var _NoParamReturnArrayMatrix4x4Callback = func() []plugify.Matrix4x4 {
	var __retVal []plugify.Matrix4x4
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayMatrix4x4Callback()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataMatrix4x4[plugify.Matrix4x4](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorMatrix4x4(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// NoParamReturnArrayMatrix4x4Callback
func NoParamReturnArrayMatrix4x4Callback() []plugify.Matrix4x4 {
	defer plugify.Scope("cross_call_master::NoParamReturnArrayMatrix4x4Callback", ModuleName, 3)()
	return _NoParamReturnArrayMatrix4x4Callback()
}

var _NoParamReturnVector2Callback = func() plugify.Vector2 {
	__native := C.NoParamReturnVector2Callback()
	__retVal := *(*plugify.Vector2)(unsafe.Pointer(&__native))
	return __retVal
}

// NoParamReturnVector2Callback
func NoParamReturnVector2Callback() plugify.Vector2 {
	defer plugify.Scope("cross_call_master::NoParamReturnVector2Callback", ModuleName, 3)()
	return _NoParamReturnVector2Callback()
}

var _NoParamReturnVector3Callback = func() plugify.Vector3 {
	__native := C.NoParamReturnVector3Callback()
	__retVal := *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// NoParamReturnVector3Callback
func NoParamReturnVector3Callback() plugify.Vector3 {
	defer plugify.Scope("cross_call_master::NoParamReturnVector3Callback", ModuleName, 3)()
	return _NoParamReturnVector3Callback()
}

var _NoParamReturnVector4Callback = func() plugify.Vector4 {
	__native := C.NoParamReturnVector4Callback()
	__retVal := *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// NoParamReturnVector4Callback
func NoParamReturnVector4Callback() plugify.Vector4 {
	defer plugify.Scope("cross_call_master::NoParamReturnVector4Callback", ModuleName, 3)()
	return _NoParamReturnVector4Callback()
}

var _NoParamReturnMatrix4x4Callback = func() plugify.Matrix4x4 {
	__native := C.NoParamReturnMatrix4x4Callback()
	__retVal := *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// NoParamReturnMatrix4x4Callback
func NoParamReturnMatrix4x4Callback() plugify.Matrix4x4 {
	defer plugify.Scope("cross_call_master::NoParamReturnMatrix4x4Callback", ModuleName, 3)()
	return _NoParamReturnMatrix4x4Callback()
}

var _Param1Callback = func(a int32) {
	__a := C.int32_t(a)
	C.Param1Callback(__a)
}

// Param1Callback
func Param1Callback(a int32) {
	defer plugify.Scope("cross_call_master::Param1Callback", ModuleName, 3)()
	_Param1Callback(a)
}

var _Param2Callback = func(a int32, b float32) {
	__a := C.int32_t(a)
	__b := C.float(b)
	C.Param2Callback(__a, __b)
}

// Param2Callback
func Param2Callback(a int32, b float32) {
	defer plugify.Scope("cross_call_master::Param2Callback", ModuleName, 3)()
	_Param2Callback(a, b)
}

var _Param3Callback = func(a int32, b float32, c float64) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	C.Param3Callback(__a, __b, __c)
}

// Param3Callback
func Param3Callback(a int32, b float32, c float64) {
	defer plugify.Scope("cross_call_master::Param3Callback", ModuleName, 3)()
	_Param3Callback(a, b, c)
}

var _Param4Callback = func(a int32, b float32, c float64, d plugify.Vector4) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	C.Param4Callback(__a, __b, __c, &__d)
}

// Param4Callback
func Param4Callback(a int32, b float32, c float64, d plugify.Vector4) {
	defer plugify.Scope("cross_call_master::Param4Callback", ModuleName, 3)()
	_Param4Callback(a, b, c, d)
}

var _Param5Callback = func(a int32, b float32, c float64, d plugify.Vector4, e []int64) {
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

// Param5Callback
func Param5Callback(a int32, b float32, c float64, d plugify.Vector4, e []int64) {
	defer plugify.Scope("cross_call_master::Param5Callback", ModuleName, 3)()
	_Param5Callback(a, b, c, d, e)
}

var _Param6Callback = func(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8) {
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

// Param6Callback
func Param6Callback(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8) {
	defer plugify.Scope("cross_call_master::Param6Callback", ModuleName, 3)()
	_Param6Callback(a, b, c, d, e, f)
}

var _Param7Callback = func(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string) {
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

// Param7Callback
func Param7Callback(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string) {
	defer plugify.Scope("cross_call_master::Param7Callback", ModuleName, 3)()
	_Param7Callback(a, b, c, d, e, f, g)
}

var _Param8Callback = func(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16) {
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

// Param8Callback
func Param8Callback(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16) {
	defer plugify.Scope("cross_call_master::Param8Callback", ModuleName, 3)()
	_Param8Callback(a, b, c, d, e, f, g, h)
}

var _Param9Callback = func(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16, k int16) {
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

// Param9Callback
func Param9Callback(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16, k int16) {
	defer plugify.Scope("cross_call_master::Param9Callback", ModuleName, 3)()
	_Param9Callback(a, b, c, d, e, f, g, h, k)
}

var _Param10Callback = func(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16, k int16, l uintptr) {
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

// Param10Callback
func Param10Callback(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16, k int16, l uintptr) {
	defer plugify.Scope("cross_call_master::Param10Callback", ModuleName, 3)()
	_Param10Callback(a, b, c, d, e, f, g, h, k, l)
}

var _ParamRef1Callback = func(a *int32) {
	__a := C.int32_t(*a)
	C.ParamRef1Callback(&__a)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
}

// ParamRef1Callback
func ParamRef1Callback(a *int32) {
	defer plugify.Scope("cross_call_master::ParamRef1Callback", ModuleName, 3)()
	_ParamRef1Callback(a)
}

var _ParamRef2Callback = func(a *int32, b *float32) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	C.ParamRef2Callback(&__a, &__b)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
	*b = float32(__b)
}

// ParamRef2Callback
func ParamRef2Callback(a *int32, b *float32) {
	defer plugify.Scope("cross_call_master::ParamRef2Callback", ModuleName, 3)()
	_ParamRef2Callback(a, b)
}

var _ParamRef3Callback = func(a *int32, b *float32, c *float64) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	C.ParamRef3Callback(&__a, &__b, &__c)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
	*b = float32(__b)
	*c = float64(__c)
}

// ParamRef3Callback
func ParamRef3Callback(a *int32, b *float32, c *float64) {
	defer plugify.Scope("cross_call_master::ParamRef3Callback", ModuleName, 3)()
	_ParamRef3Callback(a, b, c)
}

var _ParamRef4Callback = func(a *int32, b *float32, c *float64, d *plugify.Vector4) {
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

// ParamRef4Callback
func ParamRef4Callback(a *int32, b *float32, c *float64, d *plugify.Vector4) {
	defer plugify.Scope("cross_call_master::ParamRef4Callback", ModuleName, 3)()
	_ParamRef4Callback(a, b, c, d)
}

var _ParamRef5Callback = func(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64) {
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

// ParamRef5Callback
func ParamRef5Callback(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64) {
	defer plugify.Scope("cross_call_master::ParamRef5Callback", ModuleName, 3)()
	_ParamRef5Callback(a, b, c, d, e)
}

var _ParamRef6Callback = func(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8) {
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

// ParamRef6Callback
func ParamRef6Callback(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8) {
	defer plugify.Scope("cross_call_master::ParamRef6Callback", ModuleName, 3)()
	_ParamRef6Callback(a, b, c, d, e, f)
}

var _ParamRef7Callback = func(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string) {
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
			*g = plugify.GetStringData[string](&__g)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// ParamRef7Callback
func ParamRef7Callback(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string) {
	defer plugify.Scope("cross_call_master::ParamRef7Callback", ModuleName, 3)()
	_ParamRef7Callback(a, b, c, d, e, f, g)
}

var _ParamRef8Callback = func(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16) {
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
			*g = plugify.GetStringData[string](&__g)
			*h = uint16(__h)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// ParamRef8Callback
func ParamRef8Callback(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16) {
	defer plugify.Scope("cross_call_master::ParamRef8Callback", ModuleName, 3)()
	_ParamRef8Callback(a, b, c, d, e, f, g, h)
}

var _ParamRef9Callback = func(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16) {
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
			*g = plugify.GetStringData[string](&__g)
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

// ParamRef9Callback
func ParamRef9Callback(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16) {
	defer plugify.Scope("cross_call_master::ParamRef9Callback", ModuleName, 3)()
	_ParamRef9Callback(a, b, c, d, e, f, g, h, k)
}

var _ParamRef10Callback = func(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16, l *uintptr) {
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
			*g = plugify.GetStringData[string](&__g)
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

// ParamRef10Callback
func ParamRef10Callback(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16, l *uintptr) {
	defer plugify.Scope("cross_call_master::ParamRef10Callback", ModuleName, 3)()
	_ParamRef10Callback(a, b, c, d, e, f, g, h, k, l)
}

var _ParamRefVectorsCallback = func(p1 *[]bool, p2 *[]int8, p3 *[]uint16, p4 *[]int8, p5 *[]int16, p6 *[]int32, p7 *[]int64, p8 *[]uint8, p9 *[]uint16, p10 *[]uint32, p11 *[]uint64, p12 *[]uintptr, p13 *[]float32, p14 *[]float64, p15 *[]string) {
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

// ParamRefVectorsCallback
func ParamRefVectorsCallback(p1 *[]bool, p2 *[]int8, p3 *[]uint16, p4 *[]int8, p5 *[]int16, p6 *[]int32, p7 *[]int64, p8 *[]uint8, p9 *[]uint16, p10 *[]uint32, p11 *[]uint64, p12 *[]uintptr, p13 *[]float32, p14 *[]float64, p15 *[]string) {
	defer plugify.Scope("cross_call_master::ParamRefVectorsCallback", ModuleName, 3)()
	_ParamRefVectorsCallback(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15)
}

var _ParamAllPrimitivesCallback = func(p1 bool, p2 int8, p3 uint16, p4 int8, p5 int16, p6 int32, p7 int64, p8 uint8, p9 uint16, p10 uint32, p11 uint64, p12 uintptr, p13 float32, p14 float64) int64 {
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

// ParamAllPrimitivesCallback
func ParamAllPrimitivesCallback(p1 bool, p2 int8, p3 uint16, p4 int8, p5 int16, p6 int32, p7 int64, p8 uint8, p9 uint16, p10 uint32, p11 uint64, p12 uintptr, p13 float32, p14 float64) int64 {
	defer plugify.Scope("cross_call_master::ParamAllPrimitivesCallback", ModuleName, 3)()
	return _ParamAllPrimitivesCallback(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14)
}

var _ParamAllAliasesCallback = func(aBool AliasBool, aChar8 AliasChar8, aChar16 AliasChar16, aInt8 AliasInt8, aInt16 AliasInt16, aInt32 AliasInt32, aInt64 AliasInt64, aPtr AliasPtr, aFloat AliasFloat, aDouble AliasDouble, aString AliasString, aAny AliasAny, aVec2 AliasVec2, aVec3 AliasVec3, aVec4 AliasVec4, aMat4x4 AliasMat4x4, aBoolVec AliasBoolVector, aChar8Vec AliasChar8Vector, aChar16Vec AliasChar16Vector, aInt8Vec AliasInt8Vector, aInt16Vec AliasInt16Vector, aInt32Vec AliasInt32Vector, aInt64Vec AliasInt64Vector, aPtrVec AliasPtrVector, aFloatVec AliasFloatVector, aDoubleVec AliasDoubleVector, aStringVec AliasStringVector, aAnyVec AliasAnyVector, aVec2Vec AliasVec2Vector, aVec3Vec AliasVec3Vector, aVec4Vec AliasVec4Vector) int32 {
	var __retVal int32
	__aBool := C.bool(aBool)
	__aChar8 := C.int8_t(aChar8)
	__aChar16 := C.uint16_t(aChar16)
	__aInt8 := C.int8_t(aInt8)
	__aInt16 := C.int16_t(aInt16)
	__aInt32 := C.int32_t(aInt32)
	__aInt64 := C.int64_t(aInt64)
	__aPtr := C.uintptr_t(aPtr)
	__aFloat := C.float(aFloat)
	__aDouble := C.double(aDouble)
	__aString := plugify.ConstructString(aString)
	__aAny := plugify.ConstructVariant(aAny)
	__aVec2 := *(*C.Vector2)(unsafe.Pointer(&aVec2))
	__aVec3 := *(*C.Vector3)(unsafe.Pointer(&aVec3))
	__aVec4 := *(*C.Vector4)(unsafe.Pointer(&aVec4))
	__aMat4x4 := *(*C.Matrix4x4)(unsafe.Pointer(&aMat4x4))
	__aBoolVec := plugify.ConstructVectorBool(aBoolVec)
	__aChar8Vec := plugify.ConstructVectorChar8(aChar8Vec)
	__aChar16Vec := plugify.ConstructVectorChar16(aChar16Vec)
	__aInt8Vec := plugify.ConstructVectorInt8(aInt8Vec)
	__aInt16Vec := plugify.ConstructVectorInt16(aInt16Vec)
	__aInt32Vec := plugify.ConstructVectorInt32(aInt32Vec)
	__aInt64Vec := plugify.ConstructVectorInt64(aInt64Vec)
	__aPtrVec := plugify.ConstructVectorPointer(aPtrVec)
	__aFloatVec := plugify.ConstructVectorFloat(aFloatVec)
	__aDoubleVec := plugify.ConstructVectorDouble(aDoubleVec)
	__aStringVec := plugify.ConstructVectorString(aStringVec)
	__aAnyVec := plugify.ConstructVectorVariant(aAnyVec)
	__aVec2Vec := plugify.ConstructVectorVector2(aVec2Vec)
	__aVec3Vec := plugify.ConstructVectorVector3(aVec3Vec)
	__aVec4Vec := plugify.ConstructVectorVector4(aVec4Vec)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.ParamAllAliasesCallback(__aBool, __aChar8, __aChar16, __aInt8, __aInt16, __aInt32, __aInt64, __aPtr, __aFloat, __aDouble, (*C.String)(unsafe.Pointer(&__aString)), (*C.Variant)(unsafe.Pointer(&__aAny)), &__aVec2, &__aVec3, &__aVec4, &__aMat4x4, (*C.Vector)(unsafe.Pointer(&__aBoolVec)), (*C.Vector)(unsafe.Pointer(&__aChar8Vec)), (*C.Vector)(unsafe.Pointer(&__aChar16Vec)), (*C.Vector)(unsafe.Pointer(&__aInt8Vec)), (*C.Vector)(unsafe.Pointer(&__aInt16Vec)), (*C.Vector)(unsafe.Pointer(&__aInt32Vec)), (*C.Vector)(unsafe.Pointer(&__aInt64Vec)), (*C.Vector)(unsafe.Pointer(&__aPtrVec)), (*C.Vector)(unsafe.Pointer(&__aFloatVec)), (*C.Vector)(unsafe.Pointer(&__aDoubleVec)), (*C.Vector)(unsafe.Pointer(&__aStringVec)), (*C.Vector)(unsafe.Pointer(&__aAnyVec)), (*C.Vector)(unsafe.Pointer(&__aVec2Vec)), (*C.Vector)(unsafe.Pointer(&__aVec3Vec)), (*C.Vector)(unsafe.Pointer(&__aVec4Vec))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__aString)
			plugify.DestroyVariant(&__aAny)
			plugify.DestroyVectorBool(&__aBoolVec)
			plugify.DestroyVectorChar8(&__aChar8Vec)
			plugify.DestroyVectorChar16(&__aChar16Vec)
			plugify.DestroyVectorInt8(&__aInt8Vec)
			plugify.DestroyVectorInt16(&__aInt16Vec)
			plugify.DestroyVectorInt32(&__aInt32Vec)
			plugify.DestroyVectorInt64(&__aInt64Vec)
			plugify.DestroyVectorPointer(&__aPtrVec)
			plugify.DestroyVectorFloat(&__aFloatVec)
			plugify.DestroyVectorDouble(&__aDoubleVec)
			plugify.DestroyVectorString(&__aStringVec)
			plugify.DestroyVectorVariant(&__aAnyVec)
			plugify.DestroyVectorVector2(&__aVec2Vec)
			plugify.DestroyVectorVector3(&__aVec3Vec)
			plugify.DestroyVectorVector4(&__aVec4Vec)
		},
	}.Do()
	return __retVal
}

// ParamAllAliasesCallback
func ParamAllAliasesCallback(aBool AliasBool, aChar8 AliasChar8, aChar16 AliasChar16, aInt8 AliasInt8, aInt16 AliasInt16, aInt32 AliasInt32, aInt64 AliasInt64, aPtr AliasPtr, aFloat AliasFloat, aDouble AliasDouble, aString AliasString, aAny AliasAny, aVec2 AliasVec2, aVec3 AliasVec3, aVec4 AliasVec4, aMat4x4 AliasMat4x4, aBoolVec AliasBoolVector, aChar8Vec AliasChar8Vector, aChar16Vec AliasChar16Vector, aInt8Vec AliasInt8Vector, aInt16Vec AliasInt16Vector, aInt32Vec AliasInt32Vector, aInt64Vec AliasInt64Vector, aPtrVec AliasPtrVector, aFloatVec AliasFloatVector, aDoubleVec AliasDoubleVector, aStringVec AliasStringVector, aAnyVec AliasAnyVector, aVec2Vec AliasVec2Vector, aVec3Vec AliasVec3Vector, aVec4Vec AliasVec4Vector) int32 {
	defer plugify.Scope("cross_call_master::ParamAllAliasesCallback", ModuleName, 3)()
	return _ParamAllAliasesCallback(aBool, aChar8, aChar16, aInt8, aInt16, aInt32, aInt64, aPtr, aFloat, aDouble, aString, aAny, aVec2, aVec3, aVec4, aMat4x4, aBoolVec, aChar8Vec, aChar16Vec, aInt8Vec, aInt16Vec, aInt32Vec, aInt64Vec, aPtrVec, aFloatVec, aDoubleVec, aStringVec, aAnyVec, aVec2Vec, aVec3Vec, aVec4Vec)
}

var _ParamAllRefAliasesCallback = func(aBool *AliasBool, aChar8 *AliasChar8, aChar16 *AliasChar16, aInt8 *AliasInt8, aInt16 *AliasInt16, aInt32 *AliasInt32, aInt64 *AliasInt64, aPtr *AliasPtr, aFloat *AliasFloat, aDouble *AliasDouble, aString *AliasString, aAny *AliasAny, aVec2 *AliasVec2, aVec3 *AliasVec3, aVec4 *AliasVec4, aMat4x4 *AliasMat4x4, aBoolVec *AliasBoolVector, aChar8Vec *AliasChar8Vector, aChar16Vec *AliasChar16Vector, aInt8Vec *AliasInt8Vector, aInt16Vec *AliasInt16Vector, aInt32Vec *AliasInt32Vector, aInt64Vec *AliasInt64Vector, aPtrVec *AliasPtrVector, aFloatVec *AliasFloatVector, aDoubleVec *AliasDoubleVector, aStringVec *AliasStringVector, aAnyVec *AliasAnyVector, aVec2Vec *AliasVec2Vector, aVec3Vec *AliasVec3Vector, aVec4Vec *AliasVec4Vector) int64 {
	var __retVal int64
	__aBool := C.bool(*aBool)
	__aChar8 := C.int8_t(*aChar8)
	__aChar16 := C.uint16_t(*aChar16)
	__aInt8 := C.int8_t(*aInt8)
	__aInt16 := C.int16_t(*aInt16)
	__aInt32 := C.int32_t(*aInt32)
	__aInt64 := C.int64_t(*aInt64)
	__aPtr := C.uintptr_t(*aPtr)
	__aFloat := C.float(*aFloat)
	__aDouble := C.double(*aDouble)
	__aString := plugify.ConstructString(*aString)
	__aAny := plugify.ConstructVariant(*aAny)
	__aVec2 := *(*C.Vector2)(unsafe.Pointer(aVec2))
	__aVec3 := *(*C.Vector3)(unsafe.Pointer(aVec3))
	__aVec4 := *(*C.Vector4)(unsafe.Pointer(aVec4))
	__aMat4x4 := *(*C.Matrix4x4)(unsafe.Pointer(aMat4x4))
	__aBoolVec := plugify.ConstructVectorBool(*aBoolVec)
	__aChar8Vec := plugify.ConstructVectorChar8(*aChar8Vec)
	__aChar16Vec := plugify.ConstructVectorChar16(*aChar16Vec)
	__aInt8Vec := plugify.ConstructVectorInt8(*aInt8Vec)
	__aInt16Vec := plugify.ConstructVectorInt16(*aInt16Vec)
	__aInt32Vec := plugify.ConstructVectorInt32(*aInt32Vec)
	__aInt64Vec := plugify.ConstructVectorInt64(*aInt64Vec)
	__aPtrVec := plugify.ConstructVectorPointer(*aPtrVec)
	__aFloatVec := plugify.ConstructVectorFloat(*aFloatVec)
	__aDoubleVec := plugify.ConstructVectorDouble(*aDoubleVec)
	__aStringVec := plugify.ConstructVectorString(*aStringVec)
	__aAnyVec := plugify.ConstructVectorVariant(*aAnyVec)
	__aVec2Vec := plugify.ConstructVectorVector2(*aVec2Vec)
	__aVec3Vec := plugify.ConstructVectorVector3(*aVec3Vec)
	__aVec4Vec := plugify.ConstructVectorVector4(*aVec4Vec)
	plugify.Block{
		Try: func() {
			__retVal = int64(C.ParamAllRefAliasesCallback(&__aBool, &__aChar8, &__aChar16, &__aInt8, &__aInt16, &__aInt32, &__aInt64, &__aPtr, &__aFloat, &__aDouble, (*C.String)(unsafe.Pointer(&__aString)), (*C.Variant)(unsafe.Pointer(&__aAny)), &__aVec2, &__aVec3, &__aVec4, &__aMat4x4, (*C.Vector)(unsafe.Pointer(&__aBoolVec)), (*C.Vector)(unsafe.Pointer(&__aChar8Vec)), (*C.Vector)(unsafe.Pointer(&__aChar16Vec)), (*C.Vector)(unsafe.Pointer(&__aInt8Vec)), (*C.Vector)(unsafe.Pointer(&__aInt16Vec)), (*C.Vector)(unsafe.Pointer(&__aInt32Vec)), (*C.Vector)(unsafe.Pointer(&__aInt64Vec)), (*C.Vector)(unsafe.Pointer(&__aPtrVec)), (*C.Vector)(unsafe.Pointer(&__aFloatVec)), (*C.Vector)(unsafe.Pointer(&__aDoubleVec)), (*C.Vector)(unsafe.Pointer(&__aStringVec)), (*C.Vector)(unsafe.Pointer(&__aAnyVec)), (*C.Vector)(unsafe.Pointer(&__aVec2Vec)), (*C.Vector)(unsafe.Pointer(&__aVec3Vec)), (*C.Vector)(unsafe.Pointer(&__aVec4Vec))))
			// Unmarshal - Convert native data to managed data.
			*aBool = AliasBool(__aBool)
			*aChar8 = AliasChar8(__aChar8)
			*aChar16 = AliasChar16(__aChar16)
			*aInt8 = AliasInt8(__aInt8)
			*aInt16 = AliasInt16(__aInt16)
			*aInt32 = AliasInt32(__aInt32)
			*aInt64 = AliasInt64(__aInt64)
			*aPtr = AliasPtr(__aPtr)
			*aFloat = AliasFloat(__aFloat)
			*aDouble = AliasDouble(__aDouble)
			*aString = plugify.GetStringData[AliasString](&__aString)
			*aAny = plugify.GetVariantData(&__aAny)
			*aVec2 = *(*AliasVec2)(unsafe.Pointer(&__aVec2))
			*aVec3 = *(*AliasVec3)(unsafe.Pointer(&__aVec3))
			*aVec4 = *(*AliasVec4)(unsafe.Pointer(&__aVec4))
			*aMat4x4 = *(*AliasMat4x4)(unsafe.Pointer(&__aMat4x4))
			plugify.GetVectorDataBoolTo(&__aBoolVec, aBoolVec)
			plugify.GetVectorDataChar8To(&__aChar8Vec, aChar8Vec)
			plugify.GetVectorDataChar16To(&__aChar16Vec, aChar16Vec)
			plugify.GetVectorDataInt8To(&__aInt8Vec, aInt8Vec)
			plugify.GetVectorDataInt16To(&__aInt16Vec, aInt16Vec)
			plugify.GetVectorDataInt32To(&__aInt32Vec, aInt32Vec)
			plugify.GetVectorDataInt64To(&__aInt64Vec, aInt64Vec)
			plugify.GetVectorDataPointerTo(&__aPtrVec, aPtrVec)
			plugify.GetVectorDataFloatTo(&__aFloatVec, aFloatVec)
			plugify.GetVectorDataDoubleTo(&__aDoubleVec, aDoubleVec)
			plugify.GetVectorDataStringTo(&__aStringVec, aStringVec)
			plugify.GetVectorDataVariantTo(&__aAnyVec, aAnyVec)
			plugify.GetVectorDataVector2To(&__aVec2Vec, aVec2Vec)
			plugify.GetVectorDataVector3To(&__aVec3Vec, aVec3Vec)
			plugify.GetVectorDataVector4To(&__aVec4Vec, aVec4Vec)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__aString)
			plugify.DestroyVariant(&__aAny)
			plugify.DestroyVectorBool(&__aBoolVec)
			plugify.DestroyVectorChar8(&__aChar8Vec)
			plugify.DestroyVectorChar16(&__aChar16Vec)
			plugify.DestroyVectorInt8(&__aInt8Vec)
			plugify.DestroyVectorInt16(&__aInt16Vec)
			plugify.DestroyVectorInt32(&__aInt32Vec)
			plugify.DestroyVectorInt64(&__aInt64Vec)
			plugify.DestroyVectorPointer(&__aPtrVec)
			plugify.DestroyVectorFloat(&__aFloatVec)
			plugify.DestroyVectorDouble(&__aDoubleVec)
			plugify.DestroyVectorString(&__aStringVec)
			plugify.DestroyVectorVariant(&__aAnyVec)
			plugify.DestroyVectorVector2(&__aVec2Vec)
			plugify.DestroyVectorVector3(&__aVec3Vec)
			plugify.DestroyVectorVector4(&__aVec4Vec)
		},
	}.Do()
	return __retVal
}

// ParamAllRefAliasesCallback
func ParamAllRefAliasesCallback(aBool *AliasBool, aChar8 *AliasChar8, aChar16 *AliasChar16, aInt8 *AliasInt8, aInt16 *AliasInt16, aInt32 *AliasInt32, aInt64 *AliasInt64, aPtr *AliasPtr, aFloat *AliasFloat, aDouble *AliasDouble, aString *AliasString, aAny *AliasAny, aVec2 *AliasVec2, aVec3 *AliasVec3, aVec4 *AliasVec4, aMat4x4 *AliasMat4x4, aBoolVec *AliasBoolVector, aChar8Vec *AliasChar8Vector, aChar16Vec *AliasChar16Vector, aInt8Vec *AliasInt8Vector, aInt16Vec *AliasInt16Vector, aInt32Vec *AliasInt32Vector, aInt64Vec *AliasInt64Vector, aPtrVec *AliasPtrVector, aFloatVec *AliasFloatVector, aDoubleVec *AliasDoubleVector, aStringVec *AliasStringVector, aAnyVec *AliasAnyVector, aVec2Vec *AliasVec2Vector, aVec3Vec *AliasVec3Vector, aVec4Vec *AliasVec4Vector) int64 {
	defer plugify.Scope("cross_call_master::ParamAllRefAliasesCallback", ModuleName, 3)()
	return _ParamAllRefAliasesCallback(aBool, aChar8, aChar16, aInt8, aInt16, aInt32, aInt64, aPtr, aFloat, aDouble, aString, aAny, aVec2, aVec3, aVec4, aMat4x4, aBoolVec, aChar8Vec, aChar16Vec, aInt8Vec, aInt16Vec, aInt32Vec, aInt64Vec, aPtrVec, aFloatVec, aDoubleVec, aStringVec, aAnyVec, aVec2Vec, aVec3Vec, aVec4Vec)
}

var _ParamEnumCallback = func(p1 Example, p2 []Example) int32 {
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

// ParamEnumCallback
func ParamEnumCallback(p1 Example, p2 []Example) int32 {
	defer plugify.Scope("cross_call_master::ParamEnumCallback", ModuleName, 3)()
	return _ParamEnumCallback(p1, p2)
}

var _ParamEnumRefCallback = func(p1 *Example, p2 *[]Example) int32 {
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

// ParamEnumRefCallback
func ParamEnumRefCallback(p1 *Example, p2 *[]Example) int32 {
	defer plugify.Scope("cross_call_master::ParamEnumRefCallback", ModuleName, 3)()
	return _ParamEnumRefCallback(p1, p2)
}

var _ParamVariantCallback = func(p1 any, p2 []any) {
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

// ParamVariantCallback
func ParamVariantCallback(p1 any, p2 []any) {
	defer plugify.Scope("cross_call_master::ParamVariantCallback", ModuleName, 3)()
	_ParamVariantCallback(p1, p2)
}

var _ParamVariantRefCallback = func(p1 *any, p2 *[]any) {
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

// ParamVariantRefCallback
func ParamVariantRefCallback(p1 *any, p2 *[]any) {
	defer plugify.Scope("cross_call_master::ParamVariantRefCallback", ModuleName, 3)()
	_ParamVariantRefCallback(p1, p2)
}

var _CallFuncVoidCallback = func(func_ FuncVoid) {
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	C.CallFuncVoidCallback(__func_)
}

// CallFuncVoidCallback
func CallFuncVoidCallback(func_ FuncVoid) {
	defer plugify.Scope("cross_call_master::CallFuncVoidCallback", ModuleName, 3)()
	_CallFuncVoidCallback(func_)
}

var _CallFuncBoolCallback = func(func_ FuncBool) bool {
	var __retVal bool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = bool(C.CallFuncBoolCallback(__func_))
	return __retVal
}

// CallFuncBoolCallback
func CallFuncBoolCallback(func_ FuncBool) bool {
	defer plugify.Scope("cross_call_master::CallFuncBoolCallback", ModuleName, 3)()
	return _CallFuncBoolCallback(func_)
}

var _CallFuncChar8Callback = func(func_ FuncChar8) int8 {
	var __retVal int8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int8(C.CallFuncChar8Callback(__func_))
	return __retVal
}

// CallFuncChar8Callback
func CallFuncChar8Callback(func_ FuncChar8) int8 {
	defer plugify.Scope("cross_call_master::CallFuncChar8Callback", ModuleName, 3)()
	return _CallFuncChar8Callback(func_)
}

var _CallFuncChar16Callback = func(func_ FuncChar16) uint16 {
	var __retVal uint16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint16(C.CallFuncChar16Callback(__func_))
	return __retVal
}

// CallFuncChar16Callback
func CallFuncChar16Callback(func_ FuncChar16) uint16 {
	defer plugify.Scope("cross_call_master::CallFuncChar16Callback", ModuleName, 3)()
	return _CallFuncChar16Callback(func_)
}

var _CallFuncInt8Callback = func(func_ FuncInt8) int8 {
	var __retVal int8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int8(C.CallFuncInt8Callback(__func_))
	return __retVal
}

// CallFuncInt8Callback
func CallFuncInt8Callback(func_ FuncInt8) int8 {
	defer plugify.Scope("cross_call_master::CallFuncInt8Callback", ModuleName, 3)()
	return _CallFuncInt8Callback(func_)
}

var _CallFuncInt16Callback = func(func_ FuncInt16) int16 {
	var __retVal int16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int16(C.CallFuncInt16Callback(__func_))
	return __retVal
}

// CallFuncInt16Callback
func CallFuncInt16Callback(func_ FuncInt16) int16 {
	defer plugify.Scope("cross_call_master::CallFuncInt16Callback", ModuleName, 3)()
	return _CallFuncInt16Callback(func_)
}

var _CallFuncInt32Callback = func(func_ FuncInt32) int32 {
	var __retVal int32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int32(C.CallFuncInt32Callback(__func_))
	return __retVal
}

// CallFuncInt32Callback
func CallFuncInt32Callback(func_ FuncInt32) int32 {
	defer plugify.Scope("cross_call_master::CallFuncInt32Callback", ModuleName, 3)()
	return _CallFuncInt32Callback(func_)
}

var _CallFuncInt64Callback = func(func_ FuncInt64) int64 {
	var __retVal int64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int64(C.CallFuncInt64Callback(__func_))
	return __retVal
}

// CallFuncInt64Callback
func CallFuncInt64Callback(func_ FuncInt64) int64 {
	defer plugify.Scope("cross_call_master::CallFuncInt64Callback", ModuleName, 3)()
	return _CallFuncInt64Callback(func_)
}

var _CallFuncUInt8Callback = func(func_ FuncUInt8) uint8 {
	var __retVal uint8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint8(C.CallFuncUInt8Callback(__func_))
	return __retVal
}

// CallFuncUInt8Callback
func CallFuncUInt8Callback(func_ FuncUInt8) uint8 {
	defer plugify.Scope("cross_call_master::CallFuncUInt8Callback", ModuleName, 3)()
	return _CallFuncUInt8Callback(func_)
}

var _CallFuncUInt16Callback = func(func_ FuncUInt16) uint16 {
	var __retVal uint16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint16(C.CallFuncUInt16Callback(__func_))
	return __retVal
}

// CallFuncUInt16Callback
func CallFuncUInt16Callback(func_ FuncUInt16) uint16 {
	defer plugify.Scope("cross_call_master::CallFuncUInt16Callback", ModuleName, 3)()
	return _CallFuncUInt16Callback(func_)
}

var _CallFuncUInt32Callback = func(func_ FuncUInt32) uint32 {
	var __retVal uint32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint32(C.CallFuncUInt32Callback(__func_))
	return __retVal
}

// CallFuncUInt32Callback
func CallFuncUInt32Callback(func_ FuncUInt32) uint32 {
	defer plugify.Scope("cross_call_master::CallFuncUInt32Callback", ModuleName, 3)()
	return _CallFuncUInt32Callback(func_)
}

var _CallFuncUInt64Callback = func(func_ FuncUInt64) uint64 {
	var __retVal uint64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint64(C.CallFuncUInt64Callback(__func_))
	return __retVal
}

// CallFuncUInt64Callback
func CallFuncUInt64Callback(func_ FuncUInt64) uint64 {
	defer plugify.Scope("cross_call_master::CallFuncUInt64Callback", ModuleName, 3)()
	return _CallFuncUInt64Callback(func_)
}

var _CallFuncPtrCallback = func(func_ FuncPtr) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFuncPtrCallback(__func_))
	return __retVal
}

// CallFuncPtrCallback
func CallFuncPtrCallback(func_ FuncPtr) uintptr {
	defer plugify.Scope("cross_call_master::CallFuncPtrCallback", ModuleName, 3)()
	return _CallFuncPtrCallback(func_)
}

var _CallFuncFloatCallback = func(func_ FuncFloat) float32 {
	var __retVal float32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = float32(C.CallFuncFloatCallback(__func_))
	return __retVal
}

// CallFuncFloatCallback
func CallFuncFloatCallback(func_ FuncFloat) float32 {
	defer plugify.Scope("cross_call_master::CallFuncFloatCallback", ModuleName, 3)()
	return _CallFuncFloatCallback(func_)
}

var _CallFuncDoubleCallback = func(func_ FuncDouble) float64 {
	var __retVal float64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = float64(C.CallFuncDoubleCallback(__func_))
	return __retVal
}

// CallFuncDoubleCallback
func CallFuncDoubleCallback(func_ FuncDouble) float64 {
	defer plugify.Scope("cross_call_master::CallFuncDoubleCallback", ModuleName, 3)()
	return _CallFuncDoubleCallback(func_)
}

var _CallFuncStringCallback = func(func_ FuncString) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncStringCallback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncStringCallback
func CallFuncStringCallback(func_ FuncString) string {
	defer plugify.Scope("cross_call_master::CallFuncStringCallback", ModuleName, 3)()
	return _CallFuncStringCallback(func_)
}

var _CallFuncAnyCallback = func(func_ FuncAny) any {
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

// CallFuncAnyCallback
func CallFuncAnyCallback(func_ FuncAny) any {
	defer plugify.Scope("cross_call_master::CallFuncAnyCallback", ModuleName, 3)()
	return _CallFuncAnyCallback(func_)
}

var _CallFuncFunctionCallback = func(func_ FuncFunction) FuncFunctionInner {
	var __retVal FuncFunctionInner
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = plugify.GetDelegateForFunctionPointer(C.CallFuncFunctionCallback(__func_), reflect.TypeOf(FuncFunctionInner(nil))).(FuncFunctionInner)
	return __retVal
}

// CallFuncFunctionCallback
func CallFuncFunctionCallback(func_ FuncFunction) FuncFunctionInner {
	defer plugify.Scope("cross_call_master::CallFuncFunctionCallback", ModuleName, 3)()
	return _CallFuncFunctionCallback(func_)
}

var _CallFuncBoolVectorCallback = func(func_ FuncBoolVector) []bool {
	var __retVal []bool
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncBoolVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataBool[bool](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorBool(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncBoolVectorCallback
func CallFuncBoolVectorCallback(func_ FuncBoolVector) []bool {
	defer plugify.Scope("cross_call_master::CallFuncBoolVectorCallback", ModuleName, 3)()
	return _CallFuncBoolVectorCallback(func_)
}

var _CallFuncChar8VectorCallback = func(func_ FuncChar8Vector) []int8 {
	var __retVal []int8
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncChar8VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataChar8[int8](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorChar8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncChar8VectorCallback
func CallFuncChar8VectorCallback(func_ FuncChar8Vector) []int8 {
	defer plugify.Scope("cross_call_master::CallFuncChar8VectorCallback", ModuleName, 3)()
	return _CallFuncChar8VectorCallback(func_)
}

var _CallFuncChar16VectorCallback = func(func_ FuncChar16Vector) []uint16 {
	var __retVal []uint16
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncChar16VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataChar16[uint16](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorChar16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncChar16VectorCallback
func CallFuncChar16VectorCallback(func_ FuncChar16Vector) []uint16 {
	defer plugify.Scope("cross_call_master::CallFuncChar16VectorCallback", ModuleName, 3)()
	return _CallFuncChar16VectorCallback(func_)
}

var _CallFuncInt8VectorCallback = func(func_ FuncInt8Vector) []int8 {
	var __retVal []int8
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncInt8VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt8[int8](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncInt8VectorCallback
func CallFuncInt8VectorCallback(func_ FuncInt8Vector) []int8 {
	defer plugify.Scope("cross_call_master::CallFuncInt8VectorCallback", ModuleName, 3)()
	return _CallFuncInt8VectorCallback(func_)
}

var _CallFuncInt16VectorCallback = func(func_ FuncInt16Vector) []int16 {
	var __retVal []int16
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncInt16VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt16[int16](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncInt16VectorCallback
func CallFuncInt16VectorCallback(func_ FuncInt16Vector) []int16 {
	defer plugify.Scope("cross_call_master::CallFuncInt16VectorCallback", ModuleName, 3)()
	return _CallFuncInt16VectorCallback(func_)
}

var _CallFuncInt32VectorCallback = func(func_ FuncInt32Vector) []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncInt32VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt32[int32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncInt32VectorCallback
func CallFuncInt32VectorCallback(func_ FuncInt32Vector) []int32 {
	defer plugify.Scope("cross_call_master::CallFuncInt32VectorCallback", ModuleName, 3)()
	return _CallFuncInt32VectorCallback(func_)
}

var _CallFuncInt64VectorCallback = func(func_ FuncInt64Vector) []int64 {
	var __retVal []int64
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncInt64VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt64[int64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncInt64VectorCallback
func CallFuncInt64VectorCallback(func_ FuncInt64Vector) []int64 {
	defer plugify.Scope("cross_call_master::CallFuncInt64VectorCallback", ModuleName, 3)()
	return _CallFuncInt64VectorCallback(func_)
}

var _CallFuncUInt8VectorCallback = func(func_ FuncUInt8Vector) []uint8 {
	var __retVal []uint8
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncUInt8VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt8[uint8](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncUInt8VectorCallback
func CallFuncUInt8VectorCallback(func_ FuncUInt8Vector) []uint8 {
	defer plugify.Scope("cross_call_master::CallFuncUInt8VectorCallback", ModuleName, 3)()
	return _CallFuncUInt8VectorCallback(func_)
}

var _CallFuncUInt16VectorCallback = func(func_ FuncUInt16Vector) []uint16 {
	var __retVal []uint16
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncUInt16VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt16[uint16](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncUInt16VectorCallback
func CallFuncUInt16VectorCallback(func_ FuncUInt16Vector) []uint16 {
	defer plugify.Scope("cross_call_master::CallFuncUInt16VectorCallback", ModuleName, 3)()
	return _CallFuncUInt16VectorCallback(func_)
}

var _CallFuncUInt32VectorCallback = func(func_ FuncUInt32Vector) []uint32 {
	var __retVal []uint32
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncUInt32VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt32[uint32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncUInt32VectorCallback
func CallFuncUInt32VectorCallback(func_ FuncUInt32Vector) []uint32 {
	defer plugify.Scope("cross_call_master::CallFuncUInt32VectorCallback", ModuleName, 3)()
	return _CallFuncUInt32VectorCallback(func_)
}

var _CallFuncUInt64VectorCallback = func(func_ FuncUInt64Vector) []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncUInt64VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt64[uint64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncUInt64VectorCallback
func CallFuncUInt64VectorCallback(func_ FuncUInt64Vector) []uint64 {
	defer plugify.Scope("cross_call_master::CallFuncUInt64VectorCallback", ModuleName, 3)()
	return _CallFuncUInt64VectorCallback(func_)
}

var _CallFuncPtrVectorCallback = func(func_ FuncPtrVector) []uintptr {
	var __retVal []uintptr
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncPtrVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataPointer[uintptr](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorPointer(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncPtrVectorCallback
func CallFuncPtrVectorCallback(func_ FuncPtrVector) []uintptr {
	defer plugify.Scope("cross_call_master::CallFuncPtrVectorCallback", ModuleName, 3)()
	return _CallFuncPtrVectorCallback(func_)
}

var _CallFuncFloatVectorCallback = func(func_ FuncFloatVector) []float32 {
	var __retVal []float32
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncFloatVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataFloat[float32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorFloat(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncFloatVectorCallback
func CallFuncFloatVectorCallback(func_ FuncFloatVector) []float32 {
	defer plugify.Scope("cross_call_master::CallFuncFloatVectorCallback", ModuleName, 3)()
	return _CallFuncFloatVectorCallback(func_)
}

var _CallFuncDoubleVectorCallback = func(func_ FuncDoubleVector) []float64 {
	var __retVal []float64
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncDoubleVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataDouble[float64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorDouble(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncDoubleVectorCallback
func CallFuncDoubleVectorCallback(func_ FuncDoubleVector) []float64 {
	defer plugify.Scope("cross_call_master::CallFuncDoubleVectorCallback", ModuleName, 3)()
	return _CallFuncDoubleVectorCallback(func_)
}

var _CallFuncStringVectorCallback = func(func_ FuncStringVector) []string {
	var __retVal []string
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncStringVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataString[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncStringVectorCallback
func CallFuncStringVectorCallback(func_ FuncStringVector) []string {
	defer plugify.Scope("cross_call_master::CallFuncStringVectorCallback", ModuleName, 3)()
	return _CallFuncStringVectorCallback(func_)
}

var _CallFuncAnyVectorCallback = func(func_ FuncAnyVector) []any {
	var __retVal []any
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAnyVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVariant[any, []any](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVariant(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAnyVectorCallback
func CallFuncAnyVectorCallback(func_ FuncAnyVector) []any {
	defer plugify.Scope("cross_call_master::CallFuncAnyVectorCallback", ModuleName, 3)()
	return _CallFuncAnyVectorCallback(func_)
}

var _CallFuncVec2VectorCallback = func(func_ FuncVec2Vector) []plugify.Vector2 {
	var __retVal []plugify.Vector2
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncVec2VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector2[plugify.Vector2](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector2(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncVec2VectorCallback
func CallFuncVec2VectorCallback(func_ FuncVec2Vector) []plugify.Vector2 {
	defer plugify.Scope("cross_call_master::CallFuncVec2VectorCallback", ModuleName, 3)()
	return _CallFuncVec2VectorCallback(func_)
}

var _CallFuncVec3VectorCallback = func(func_ FuncVec3Vector) []plugify.Vector3 {
	var __retVal []plugify.Vector3
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncVec3VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector3[plugify.Vector3](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector3(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncVec3VectorCallback
func CallFuncVec3VectorCallback(func_ FuncVec3Vector) []plugify.Vector3 {
	defer plugify.Scope("cross_call_master::CallFuncVec3VectorCallback", ModuleName, 3)()
	return _CallFuncVec3VectorCallback(func_)
}

var _CallFuncVec4VectorCallback = func(func_ FuncVec4Vector) []plugify.Vector4 {
	var __retVal []plugify.Vector4
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncVec4VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector4[plugify.Vector4](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector4(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncVec4VectorCallback
func CallFuncVec4VectorCallback(func_ FuncVec4Vector) []plugify.Vector4 {
	defer plugify.Scope("cross_call_master::CallFuncVec4VectorCallback", ModuleName, 3)()
	return _CallFuncVec4VectorCallback(func_)
}

var _CallFuncMat4x4VectorCallback = func(func_ FuncMat4x4Vector) []plugify.Matrix4x4 {
	var __retVal []plugify.Matrix4x4
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncMat4x4VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataMatrix4x4[plugify.Matrix4x4](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorMatrix4x4(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncMat4x4VectorCallback
func CallFuncMat4x4VectorCallback(func_ FuncMat4x4Vector) []plugify.Matrix4x4 {
	defer plugify.Scope("cross_call_master::CallFuncMat4x4VectorCallback", ModuleName, 3)()
	return _CallFuncMat4x4VectorCallback(func_)
}

var _CallFuncVec2Callback = func(func_ FuncVec2) plugify.Vector2 {
	var __retVal plugify.Vector2
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncVec2Callback(__func_)
	__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncVec2Callback
func CallFuncVec2Callback(func_ FuncVec2) plugify.Vector2 {
	defer plugify.Scope("cross_call_master::CallFuncVec2Callback", ModuleName, 3)()
	return _CallFuncVec2Callback(func_)
}

var _CallFuncVec3Callback = func(func_ FuncVec3) plugify.Vector3 {
	var __retVal plugify.Vector3
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncVec3Callback(__func_)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncVec3Callback
func CallFuncVec3Callback(func_ FuncVec3) plugify.Vector3 {
	defer plugify.Scope("cross_call_master::CallFuncVec3Callback", ModuleName, 3)()
	return _CallFuncVec3Callback(func_)
}

var _CallFuncVec4Callback = func(func_ FuncVec4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncVec4Callback(__func_)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncVec4Callback
func CallFuncVec4Callback(func_ FuncVec4) plugify.Vector4 {
	defer plugify.Scope("cross_call_master::CallFuncVec4Callback", ModuleName, 3)()
	return _CallFuncVec4Callback(func_)
}

var _CallFuncMat4x4Callback = func(func_ FuncMat4x4) plugify.Matrix4x4 {
	var __retVal plugify.Matrix4x4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncMat4x4Callback(__func_)
	__retVal = *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncMat4x4Callback
func CallFuncMat4x4Callback(func_ FuncMat4x4) plugify.Matrix4x4 {
	defer plugify.Scope("cross_call_master::CallFuncMat4x4Callback", ModuleName, 3)()
	return _CallFuncMat4x4Callback(func_)
}

var _CallFuncAliasBoolCallback = func(func_ FuncAliasBool) AliasBool {
	var __retVal AliasBool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasBool(C.CallFuncAliasBoolCallback(__func_))
	return __retVal
}

// CallFuncAliasBoolCallback
func CallFuncAliasBoolCallback(func_ FuncAliasBool) AliasBool {
	defer plugify.Scope("cross_call_master::CallFuncAliasBoolCallback", ModuleName, 3)()
	return _CallFuncAliasBoolCallback(func_)
}

var _CallFuncAliasChar8Callback = func(func_ FuncAliasChar8) AliasChar8 {
	var __retVal AliasChar8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasChar8(C.CallFuncAliasChar8Callback(__func_))
	return __retVal
}

// CallFuncAliasChar8Callback
func CallFuncAliasChar8Callback(func_ FuncAliasChar8) AliasChar8 {
	defer plugify.Scope("cross_call_master::CallFuncAliasChar8Callback", ModuleName, 3)()
	return _CallFuncAliasChar8Callback(func_)
}

var _CallFuncAliasChar16Callback = func(func_ FuncAliasChar16) AliasChar16 {
	var __retVal AliasChar16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasChar16(C.CallFuncAliasChar16Callback(__func_))
	return __retVal
}

// CallFuncAliasChar16Callback
func CallFuncAliasChar16Callback(func_ FuncAliasChar16) AliasChar16 {
	defer plugify.Scope("cross_call_master::CallFuncAliasChar16Callback", ModuleName, 3)()
	return _CallFuncAliasChar16Callback(func_)
}

var _CallFuncAliasInt8Callback = func(func_ FuncAliasInt8) AliasInt8 {
	var __retVal AliasInt8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasInt8(C.CallFuncAliasInt8Callback(__func_))
	return __retVal
}

// CallFuncAliasInt8Callback
func CallFuncAliasInt8Callback(func_ FuncAliasInt8) AliasInt8 {
	defer plugify.Scope("cross_call_master::CallFuncAliasInt8Callback", ModuleName, 3)()
	return _CallFuncAliasInt8Callback(func_)
}

var _CallFuncAliasInt16Callback = func(func_ FuncAliasInt16) AliasInt16 {
	var __retVal AliasInt16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasInt16(C.CallFuncAliasInt16Callback(__func_))
	return __retVal
}

// CallFuncAliasInt16Callback
func CallFuncAliasInt16Callback(func_ FuncAliasInt16) AliasInt16 {
	defer plugify.Scope("cross_call_master::CallFuncAliasInt16Callback", ModuleName, 3)()
	return _CallFuncAliasInt16Callback(func_)
}

var _CallFuncAliasInt32Callback = func(func_ FuncAliasInt32) AliasInt32 {
	var __retVal AliasInt32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasInt32(C.CallFuncAliasInt32Callback(__func_))
	return __retVal
}

// CallFuncAliasInt32Callback
func CallFuncAliasInt32Callback(func_ FuncAliasInt32) AliasInt32 {
	defer plugify.Scope("cross_call_master::CallFuncAliasInt32Callback", ModuleName, 3)()
	return _CallFuncAliasInt32Callback(func_)
}

var _CallFuncAliasInt64Callback = func(func_ FuncAliasInt64) AliasInt64 {
	var __retVal AliasInt64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasInt64(C.CallFuncAliasInt64Callback(__func_))
	return __retVal
}

// CallFuncAliasInt64Callback
func CallFuncAliasInt64Callback(func_ FuncAliasInt64) AliasInt64 {
	defer plugify.Scope("cross_call_master::CallFuncAliasInt64Callback", ModuleName, 3)()
	return _CallFuncAliasInt64Callback(func_)
}

var _CallFuncAliasUInt8Callback = func(func_ FuncAliasUInt8) AliasUInt8 {
	var __retVal AliasUInt8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasUInt8(C.CallFuncAliasUInt8Callback(__func_))
	return __retVal
}

// CallFuncAliasUInt8Callback
func CallFuncAliasUInt8Callback(func_ FuncAliasUInt8) AliasUInt8 {
	defer plugify.Scope("cross_call_master::CallFuncAliasUInt8Callback", ModuleName, 3)()
	return _CallFuncAliasUInt8Callback(func_)
}

var _CallFuncAliasUInt16Callback = func(func_ FuncAliasUInt16) AliasUInt16 {
	var __retVal AliasUInt16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasUInt16(C.CallFuncAliasUInt16Callback(__func_))
	return __retVal
}

// CallFuncAliasUInt16Callback
func CallFuncAliasUInt16Callback(func_ FuncAliasUInt16) AliasUInt16 {
	defer plugify.Scope("cross_call_master::CallFuncAliasUInt16Callback", ModuleName, 3)()
	return _CallFuncAliasUInt16Callback(func_)
}

var _CallFuncAliasUInt32Callback = func(func_ FuncAliasUInt32) AliasUInt32 {
	var __retVal AliasUInt32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasUInt32(C.CallFuncAliasUInt32Callback(__func_))
	return __retVal
}

// CallFuncAliasUInt32Callback
func CallFuncAliasUInt32Callback(func_ FuncAliasUInt32) AliasUInt32 {
	defer plugify.Scope("cross_call_master::CallFuncAliasUInt32Callback", ModuleName, 3)()
	return _CallFuncAliasUInt32Callback(func_)
}

var _CallFuncAliasUInt64Callback = func(func_ FuncAliasUInt64) AliasUInt64 {
	var __retVal AliasUInt64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasUInt64(C.CallFuncAliasUInt64Callback(__func_))
	return __retVal
}

// CallFuncAliasUInt64Callback
func CallFuncAliasUInt64Callback(func_ FuncAliasUInt64) AliasUInt64 {
	defer plugify.Scope("cross_call_master::CallFuncAliasUInt64Callback", ModuleName, 3)()
	return _CallFuncAliasUInt64Callback(func_)
}

var _CallFuncAliasPtrCallback = func(func_ FuncAliasPtr) AliasPtr {
	var __retVal AliasPtr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasPtr(C.CallFuncAliasPtrCallback(__func_))
	return __retVal
}

// CallFuncAliasPtrCallback
func CallFuncAliasPtrCallback(func_ FuncAliasPtr) AliasPtr {
	defer plugify.Scope("cross_call_master::CallFuncAliasPtrCallback", ModuleName, 3)()
	return _CallFuncAliasPtrCallback(func_)
}

var _CallFuncAliasFloatCallback = func(func_ FuncAliasFloat) AliasFloat {
	var __retVal AliasFloat
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasFloat(C.CallFuncAliasFloatCallback(__func_))
	return __retVal
}

// CallFuncAliasFloatCallback
func CallFuncAliasFloatCallback(func_ FuncAliasFloat) AliasFloat {
	defer plugify.Scope("cross_call_master::CallFuncAliasFloatCallback", ModuleName, 3)()
	return _CallFuncAliasFloatCallback(func_)
}

var _CallFuncAliasDoubleCallback = func(func_ FuncAliasDouble) AliasDouble {
	var __retVal AliasDouble
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasDouble(C.CallFuncAliasDoubleCallback(__func_))
	return __retVal
}

// CallFuncAliasDoubleCallback
func CallFuncAliasDoubleCallback(func_ FuncAliasDouble) AliasDouble {
	defer plugify.Scope("cross_call_master::CallFuncAliasDoubleCallback", ModuleName, 3)()
	return _CallFuncAliasDoubleCallback(func_)
}

var _CallFuncAliasStringCallback = func(func_ FuncAliasString) AliasString {
	var __retVal AliasString
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasStringCallback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[AliasString](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasStringCallback
func CallFuncAliasStringCallback(func_ FuncAliasString) AliasString {
	defer plugify.Scope("cross_call_master::CallFuncAliasStringCallback", ModuleName, 3)()
	return _CallFuncAliasStringCallback(func_)
}

var _CallFuncAliasAnyCallback = func(func_ FuncAliasAny) AliasAny {
	var __retVal AliasAny
	var __retVal_native plugify.PlgVariant
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasAnyCallback(__func_)
			__retVal_native = *(*plugify.PlgVariant)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = AliasAny(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasAnyCallback
func CallFuncAliasAnyCallback(func_ FuncAliasAny) AliasAny {
	defer plugify.Scope("cross_call_master::CallFuncAliasAnyCallback", ModuleName, 3)()
	return _CallFuncAliasAnyCallback(func_)
}

var _CallFuncAliasFunctionCallback = func(func_ FuncAliasFunction) AliasFunction {
	var __retVal AliasFunction
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = plugify.GetDelegateForFunctionPointer(C.CallFuncAliasFunctionCallback(__func_), reflect.TypeOf(AliasFunction(nil))).(AliasFunction)
	return __retVal
}

// CallFuncAliasFunctionCallback
func CallFuncAliasFunctionCallback(func_ FuncAliasFunction) AliasFunction {
	defer plugify.Scope("cross_call_master::CallFuncAliasFunctionCallback", ModuleName, 3)()
	return _CallFuncAliasFunctionCallback(func_)
}

var _CallFuncAliasBoolVectorCallback = func(func_ FuncAliasBoolVector) AliasBoolVector {
	var __retVal AliasBoolVector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasBoolVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataBool[bool](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorBool(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasBoolVectorCallback
func CallFuncAliasBoolVectorCallback(func_ FuncAliasBoolVector) AliasBoolVector {
	defer plugify.Scope("cross_call_master::CallFuncAliasBoolVectorCallback", ModuleName, 3)()
	return _CallFuncAliasBoolVectorCallback(func_)
}

var _CallFuncAliasChar8VectorCallback = func(func_ FuncAliasChar8Vector) AliasChar8Vector {
	var __retVal AliasChar8Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasChar8VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataChar8[int8](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorChar8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasChar8VectorCallback
func CallFuncAliasChar8VectorCallback(func_ FuncAliasChar8Vector) AliasChar8Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasChar8VectorCallback", ModuleName, 3)()
	return _CallFuncAliasChar8VectorCallback(func_)
}

var _CallFuncAliasChar16VectorCallback = func(func_ FuncAliasChar16Vector) AliasChar16Vector {
	var __retVal AliasChar16Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasChar16VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataChar16[uint16](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorChar16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasChar16VectorCallback
func CallFuncAliasChar16VectorCallback(func_ FuncAliasChar16Vector) AliasChar16Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasChar16VectorCallback", ModuleName, 3)()
	return _CallFuncAliasChar16VectorCallback(func_)
}

var _CallFuncAliasInt8VectorCallback = func(func_ FuncAliasInt8Vector) AliasInt8Vector {
	var __retVal AliasInt8Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasInt8VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt8[int8](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasInt8VectorCallback
func CallFuncAliasInt8VectorCallback(func_ FuncAliasInt8Vector) AliasInt8Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasInt8VectorCallback", ModuleName, 3)()
	return _CallFuncAliasInt8VectorCallback(func_)
}

var _CallFuncAliasInt16VectorCallback = func(func_ FuncAliasInt16Vector) AliasInt16Vector {
	var __retVal AliasInt16Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasInt16VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt16[int16](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasInt16VectorCallback
func CallFuncAliasInt16VectorCallback(func_ FuncAliasInt16Vector) AliasInt16Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasInt16VectorCallback", ModuleName, 3)()
	return _CallFuncAliasInt16VectorCallback(func_)
}

var _CallFuncAliasInt32VectorCallback = func(func_ FuncAliasInt32Vector) AliasInt32Vector {
	var __retVal AliasInt32Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasInt32VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt32[int32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasInt32VectorCallback
func CallFuncAliasInt32VectorCallback(func_ FuncAliasInt32Vector) AliasInt32Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasInt32VectorCallback", ModuleName, 3)()
	return _CallFuncAliasInt32VectorCallback(func_)
}

var _CallFuncAliasInt64VectorCallback = func(func_ FuncAliasInt64Vector) AliasInt64Vector {
	var __retVal AliasInt64Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasInt64VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt64[int64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasInt64VectorCallback
func CallFuncAliasInt64VectorCallback(func_ FuncAliasInt64Vector) AliasInt64Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasInt64VectorCallback", ModuleName, 3)()
	return _CallFuncAliasInt64VectorCallback(func_)
}

var _CallFuncAliasUInt8VectorCallback = func(func_ FuncAliasUInt8Vector) AliasUInt8Vector {
	var __retVal AliasUInt8Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasUInt8VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt8[uint8](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasUInt8VectorCallback
func CallFuncAliasUInt8VectorCallback(func_ FuncAliasUInt8Vector) AliasUInt8Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasUInt8VectorCallback", ModuleName, 3)()
	return _CallFuncAliasUInt8VectorCallback(func_)
}

var _CallFuncAliasUInt16VectorCallback = func(func_ FuncAliasUInt16Vector) AliasUInt16Vector {
	var __retVal AliasUInt16Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasUInt16VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt16[uint16](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt16(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasUInt16VectorCallback
func CallFuncAliasUInt16VectorCallback(func_ FuncAliasUInt16Vector) AliasUInt16Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasUInt16VectorCallback", ModuleName, 3)()
	return _CallFuncAliasUInt16VectorCallback(func_)
}

var _CallFuncAliasUInt32VectorCallback = func(func_ FuncAliasUInt32Vector) AliasUInt32Vector {
	var __retVal AliasUInt32Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasUInt32VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt32[uint32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasUInt32VectorCallback
func CallFuncAliasUInt32VectorCallback(func_ FuncAliasUInt32Vector) AliasUInt32Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasUInt32VectorCallback", ModuleName, 3)()
	return _CallFuncAliasUInt32VectorCallback(func_)
}

var _CallFuncAliasUInt64VectorCallback = func(func_ FuncAliasUInt64Vector) AliasUInt64Vector {
	var __retVal AliasUInt64Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasUInt64VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt64[uint64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasUInt64VectorCallback
func CallFuncAliasUInt64VectorCallback(func_ FuncAliasUInt64Vector) AliasUInt64Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasUInt64VectorCallback", ModuleName, 3)()
	return _CallFuncAliasUInt64VectorCallback(func_)
}

var _CallFuncAliasPtrVectorCallback = func(func_ FuncAliasPtrVector) AliasPtrVector {
	var __retVal AliasPtrVector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasPtrVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataPointer[uintptr](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorPointer(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasPtrVectorCallback
func CallFuncAliasPtrVectorCallback(func_ FuncAliasPtrVector) AliasPtrVector {
	defer plugify.Scope("cross_call_master::CallFuncAliasPtrVectorCallback", ModuleName, 3)()
	return _CallFuncAliasPtrVectorCallback(func_)
}

var _CallFuncAliasFloatVectorCallback = func(func_ FuncAliasFloatVector) AliasFloatVector {
	var __retVal AliasFloatVector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasFloatVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataFloat[float32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorFloat(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasFloatVectorCallback
func CallFuncAliasFloatVectorCallback(func_ FuncAliasFloatVector) AliasFloatVector {
	defer plugify.Scope("cross_call_master::CallFuncAliasFloatVectorCallback", ModuleName, 3)()
	return _CallFuncAliasFloatVectorCallback(func_)
}

var _CallFuncAliasDoubleVectorCallback = func(func_ FuncAliasDoubleVector) AliasDoubleVector {
	var __retVal AliasDoubleVector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasDoubleVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataDouble[float64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorDouble(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasDoubleVectorCallback
func CallFuncAliasDoubleVectorCallback(func_ FuncAliasDoubleVector) AliasDoubleVector {
	defer plugify.Scope("cross_call_master::CallFuncAliasDoubleVectorCallback", ModuleName, 3)()
	return _CallFuncAliasDoubleVectorCallback(func_)
}

var _CallFuncAliasStringVectorCallback = func(func_ FuncAliasStringVector) AliasStringVector {
	var __retVal AliasStringVector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasStringVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataString[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasStringVectorCallback
func CallFuncAliasStringVectorCallback(func_ FuncAliasStringVector) AliasStringVector {
	defer plugify.Scope("cross_call_master::CallFuncAliasStringVectorCallback", ModuleName, 3)()
	return _CallFuncAliasStringVectorCallback(func_)
}

var _CallFuncAliasAnyVectorCallback = func(func_ FuncAliasAnyVector) AliasAnyVector {
	var __retVal AliasAnyVector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasAnyVectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVariant[any, AliasAnyVector](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVariant(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasAnyVectorCallback
func CallFuncAliasAnyVectorCallback(func_ FuncAliasAnyVector) AliasAnyVector {
	defer plugify.Scope("cross_call_master::CallFuncAliasAnyVectorCallback", ModuleName, 3)()
	return _CallFuncAliasAnyVectorCallback(func_)
}

var _CallFuncAliasVec2VectorCallback = func(func_ FuncAliasVec2Vector) AliasVec2Vector {
	var __retVal AliasVec2Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasVec2VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector2[plugify.Vector2](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector2(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasVec2VectorCallback
func CallFuncAliasVec2VectorCallback(func_ FuncAliasVec2Vector) AliasVec2Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasVec2VectorCallback", ModuleName, 3)()
	return _CallFuncAliasVec2VectorCallback(func_)
}

var _CallFuncAliasVec3VectorCallback = func(func_ FuncAliasVec3Vector) AliasVec3Vector {
	var __retVal AliasVec3Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasVec3VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector3[plugify.Vector3](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector3(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasVec3VectorCallback
func CallFuncAliasVec3VectorCallback(func_ FuncAliasVec3Vector) AliasVec3Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasVec3VectorCallback", ModuleName, 3)()
	return _CallFuncAliasVec3VectorCallback(func_)
}

var _CallFuncAliasVec4VectorCallback = func(func_ FuncAliasVec4Vector) AliasVec4Vector {
	var __retVal AliasVec4Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasVec4VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataVector4[plugify.Vector4](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVector4(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasVec4VectorCallback
func CallFuncAliasVec4VectorCallback(func_ FuncAliasVec4Vector) AliasVec4Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasVec4VectorCallback", ModuleName, 3)()
	return _CallFuncAliasVec4VectorCallback(func_)
}

var _CallFuncAliasMat4x4VectorCallback = func(func_ FuncAliasMat4x4Vector) AliasMat4x4Vector {
	var __retVal AliasMat4x4Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasMat4x4VectorCallback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataMatrix4x4[plugify.Matrix4x4](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorMatrix4x4(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasMat4x4VectorCallback
func CallFuncAliasMat4x4VectorCallback(func_ FuncAliasMat4x4Vector) AliasMat4x4Vector {
	defer plugify.Scope("cross_call_master::CallFuncAliasMat4x4VectorCallback", ModuleName, 3)()
	return _CallFuncAliasMat4x4VectorCallback(func_)
}

var _CallFuncAliasVec2Callback = func(func_ FuncAliasVec2) AliasVec2 {
	var __retVal AliasVec2
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncAliasVec2Callback(__func_)
	__retVal = *(*AliasVec2)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncAliasVec2Callback
func CallFuncAliasVec2Callback(func_ FuncAliasVec2) AliasVec2 {
	defer plugify.Scope("cross_call_master::CallFuncAliasVec2Callback", ModuleName, 3)()
	return _CallFuncAliasVec2Callback(func_)
}

var _CallFuncAliasVec3Callback = func(func_ FuncAliasVec3) AliasVec3 {
	var __retVal AliasVec3
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncAliasVec3Callback(__func_)
	__retVal = *(*AliasVec3)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncAliasVec3Callback
func CallFuncAliasVec3Callback(func_ FuncAliasVec3) AliasVec3 {
	defer plugify.Scope("cross_call_master::CallFuncAliasVec3Callback", ModuleName, 3)()
	return _CallFuncAliasVec3Callback(func_)
}

var _CallFuncAliasVec4Callback = func(func_ FuncAliasVec4) AliasVec4 {
	var __retVal AliasVec4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncAliasVec4Callback(__func_)
	__retVal = *(*AliasVec4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncAliasVec4Callback
func CallFuncAliasVec4Callback(func_ FuncAliasVec4) AliasVec4 {
	defer plugify.Scope("cross_call_master::CallFuncAliasVec4Callback", ModuleName, 3)()
	return _CallFuncAliasVec4Callback(func_)
}

var _CallFuncAliasMat4x4Callback = func(func_ FuncAliasMat4x4) AliasMat4x4 {
	var __retVal AliasMat4x4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncAliasMat4x4Callback(__func_)
	__retVal = *(*AliasMat4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncAliasMat4x4Callback
func CallFuncAliasMat4x4Callback(func_ FuncAliasMat4x4) AliasMat4x4 {
	defer plugify.Scope("cross_call_master::CallFuncAliasMat4x4Callback", ModuleName, 3)()
	return _CallFuncAliasMat4x4Callback(func_)
}

var _CallFuncAliasAllCallback = func(func_ FuncAliasAll) AliasString {
	var __retVal AliasString
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasAllCallback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[AliasString](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncAliasAllCallback
func CallFuncAliasAllCallback(func_ FuncAliasAll) AliasString {
	defer plugify.Scope("cross_call_master::CallFuncAliasAllCallback", ModuleName, 3)()
	return _CallFuncAliasAllCallback(func_)
}

var _CallFunc1Callback = func(func_ Func1) int32 {
	var __retVal int32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int32(C.CallFunc1Callback(__func_))
	return __retVal
}

// CallFunc1Callback
func CallFunc1Callback(func_ Func1) int32 {
	defer plugify.Scope("cross_call_master::CallFunc1Callback", ModuleName, 3)()
	return _CallFunc1Callback(func_)
}

var _CallFunc2Callback = func(func_ Func2) int8 {
	var __retVal int8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int8(C.CallFunc2Callback(__func_))
	return __retVal
}

// CallFunc2Callback
func CallFunc2Callback(func_ Func2) int8 {
	defer plugify.Scope("cross_call_master::CallFunc2Callback", ModuleName, 3)()
	return _CallFunc2Callback(func_)
}

var _CallFunc3Callback = func(func_ Func3) {
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	C.CallFunc3Callback(__func_)
}

// CallFunc3Callback
func CallFunc3Callback(func_ Func3) {
	defer plugify.Scope("cross_call_master::CallFunc3Callback", ModuleName, 3)()
	_CallFunc3Callback(func_)
}

var _CallFunc4Callback = func(func_ Func4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFunc4Callback(__func_)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFunc4Callback
func CallFunc4Callback(func_ Func4) plugify.Vector4 {
	defer plugify.Scope("cross_call_master::CallFunc4Callback", ModuleName, 3)()
	return _CallFunc4Callback(func_)
}

var _CallFunc5Callback = func(func_ Func5) bool {
	var __retVal bool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = bool(C.CallFunc5Callback(__func_))
	return __retVal
}

// CallFunc5Callback
func CallFunc5Callback(func_ Func5) bool {
	defer plugify.Scope("cross_call_master::CallFunc5Callback", ModuleName, 3)()
	return _CallFunc5Callback(func_)
}

var _CallFunc6Callback = func(func_ Func6) int64 {
	var __retVal int64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int64(C.CallFunc6Callback(__func_))
	return __retVal
}

// CallFunc6Callback
func CallFunc6Callback(func_ Func6) int64 {
	defer plugify.Scope("cross_call_master::CallFunc6Callback", ModuleName, 3)()
	return _CallFunc6Callback(func_)
}

var _CallFunc7Callback = func(func_ Func7) float64 {
	var __retVal float64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = float64(C.CallFunc7Callback(__func_))
	return __retVal
}

// CallFunc7Callback
func CallFunc7Callback(func_ Func7) float64 {
	defer plugify.Scope("cross_call_master::CallFunc7Callback", ModuleName, 3)()
	return _CallFunc7Callback(func_)
}

var _CallFunc8Callback = func(func_ Func8) plugify.Matrix4x4 {
	var __retVal plugify.Matrix4x4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFunc8Callback(__func_)
	__retVal = *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFunc8Callback
func CallFunc8Callback(func_ Func8) plugify.Matrix4x4 {
	defer plugify.Scope("cross_call_master::CallFunc8Callback", ModuleName, 3)()
	return _CallFunc8Callback(func_)
}

var _CallFunc9Callback = func(func_ Func9) {
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	C.CallFunc9Callback(__func_)
}

// CallFunc9Callback
func CallFunc9Callback(func_ Func9) {
	defer plugify.Scope("cross_call_master::CallFunc9Callback", ModuleName, 3)()
	_CallFunc9Callback(func_)
}

var _CallFunc10Callback = func(func_ Func10) uint32 {
	var __retVal uint32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint32(C.CallFunc10Callback(__func_))
	return __retVal
}

// CallFunc10Callback
func CallFunc10Callback(func_ Func10) uint32 {
	defer plugify.Scope("cross_call_master::CallFunc10Callback", ModuleName, 3)()
	return _CallFunc10Callback(func_)
}

var _CallFunc11Callback = func(func_ Func11) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFunc11Callback(__func_))
	return __retVal
}

// CallFunc11Callback
func CallFunc11Callback(func_ Func11) uintptr {
	defer plugify.Scope("cross_call_master::CallFunc11Callback", ModuleName, 3)()
	return _CallFunc11Callback(func_)
}

var _CallFunc12Callback = func(func_ Func12) bool {
	var __retVal bool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = bool(C.CallFunc12Callback(__func_))
	return __retVal
}

// CallFunc12Callback
func CallFunc12Callback(func_ Func12) bool {
	defer plugify.Scope("cross_call_master::CallFunc12Callback", ModuleName, 3)()
	return _CallFunc12Callback(func_)
}

var _CallFunc13Callback = func(func_ Func13) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc13Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc13Callback
func CallFunc13Callback(func_ Func13) string {
	defer plugify.Scope("cross_call_master::CallFunc13Callback", ModuleName, 3)()
	return _CallFunc13Callback(func_)
}

var _CallFunc14Callback = func(func_ Func14) []string {
	var __retVal []string
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc14Callback(__func_)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataString[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc14Callback
func CallFunc14Callback(func_ Func14) []string {
	defer plugify.Scope("cross_call_master::CallFunc14Callback", ModuleName, 3)()
	return _CallFunc14Callback(func_)
}

var _CallFunc15Callback = func(func_ Func15) int16 {
	var __retVal int16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int16(C.CallFunc15Callback(__func_))
	return __retVal
}

// CallFunc15Callback
func CallFunc15Callback(func_ Func15) int16 {
	defer plugify.Scope("cross_call_master::CallFunc15Callback", ModuleName, 3)()
	return _CallFunc15Callback(func_)
}

var _CallFunc16Callback = func(func_ Func16) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFunc16Callback(__func_))
	return __retVal
}

// CallFunc16Callback
func CallFunc16Callback(func_ Func16) uintptr {
	defer plugify.Scope("cross_call_master::CallFunc16Callback", ModuleName, 3)()
	return _CallFunc16Callback(func_)
}

var _CallFunc17Callback = func(func_ Func17) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc17Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc17Callback
func CallFunc17Callback(func_ Func17) string {
	defer plugify.Scope("cross_call_master::CallFunc17Callback", ModuleName, 3)()
	return _CallFunc17Callback(func_)
}

var _CallFunc18Callback = func(func_ Func18) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc18Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc18Callback
func CallFunc18Callback(func_ Func18) string {
	defer plugify.Scope("cross_call_master::CallFunc18Callback", ModuleName, 3)()
	return _CallFunc18Callback(func_)
}

var _CallFunc19Callback = func(func_ Func19) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc19Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc19Callback
func CallFunc19Callback(func_ Func19) string {
	defer plugify.Scope("cross_call_master::CallFunc19Callback", ModuleName, 3)()
	return _CallFunc19Callback(func_)
}

var _CallFunc20Callback = func(func_ Func20) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc20Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc20Callback
func CallFunc20Callback(func_ Func20) string {
	defer plugify.Scope("cross_call_master::CallFunc20Callback", ModuleName, 3)()
	return _CallFunc20Callback(func_)
}

var _CallFunc21Callback = func(func_ Func21) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc21Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc21Callback
func CallFunc21Callback(func_ Func21) string {
	defer plugify.Scope("cross_call_master::CallFunc21Callback", ModuleName, 3)()
	return _CallFunc21Callback(func_)
}

var _CallFunc22Callback = func(func_ Func22) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc22Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc22Callback
func CallFunc22Callback(func_ Func22) string {
	defer plugify.Scope("cross_call_master::CallFunc22Callback", ModuleName, 3)()
	return _CallFunc22Callback(func_)
}

var _CallFunc23Callback = func(func_ Func23) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc23Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc23Callback
func CallFunc23Callback(func_ Func23) string {
	defer plugify.Scope("cross_call_master::CallFunc23Callback", ModuleName, 3)()
	return _CallFunc23Callback(func_)
}

var _CallFunc24Callback = func(func_ Func24) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc24Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc24Callback
func CallFunc24Callback(func_ Func24) string {
	defer plugify.Scope("cross_call_master::CallFunc24Callback", ModuleName, 3)()
	return _CallFunc24Callback(func_)
}

var _CallFunc25Callback = func(func_ Func25) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc25Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc25Callback
func CallFunc25Callback(func_ Func25) string {
	defer plugify.Scope("cross_call_master::CallFunc25Callback", ModuleName, 3)()
	return _CallFunc25Callback(func_)
}

var _CallFunc26Callback = func(func_ Func26) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc26Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc26Callback
func CallFunc26Callback(func_ Func26) string {
	defer plugify.Scope("cross_call_master::CallFunc26Callback", ModuleName, 3)()
	return _CallFunc26Callback(func_)
}

var _CallFunc27Callback = func(func_ Func27) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc27Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc27Callback
func CallFunc27Callback(func_ Func27) string {
	defer plugify.Scope("cross_call_master::CallFunc27Callback", ModuleName, 3)()
	return _CallFunc27Callback(func_)
}

var _CallFunc28Callback = func(func_ Func28) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc28Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc28Callback
func CallFunc28Callback(func_ Func28) string {
	defer plugify.Scope("cross_call_master::CallFunc28Callback", ModuleName, 3)()
	return _CallFunc28Callback(func_)
}

var _CallFunc29Callback = func(func_ Func29) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc29Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc29Callback
func CallFunc29Callback(func_ Func29) string {
	defer plugify.Scope("cross_call_master::CallFunc29Callback", ModuleName, 3)()
	return _CallFunc29Callback(func_)
}

var _CallFunc30Callback = func(func_ Func30) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc30Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc30Callback
func CallFunc30Callback(func_ Func30) string {
	defer plugify.Scope("cross_call_master::CallFunc30Callback", ModuleName, 3)()
	return _CallFunc30Callback(func_)
}

var _CallFunc31Callback = func(func_ Func31) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc31Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc31Callback
func CallFunc31Callback(func_ Func31) string {
	defer plugify.Scope("cross_call_master::CallFunc31Callback", ModuleName, 3)()
	return _CallFunc31Callback(func_)
}

var _CallFunc32Callback = func(func_ Func32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc32Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc32Callback
func CallFunc32Callback(func_ Func32) string {
	defer plugify.Scope("cross_call_master::CallFunc32Callback", ModuleName, 3)()
	return _CallFunc32Callback(func_)
}

var _CallFunc33Callback = func(func_ Func33) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc33Callback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFunc33Callback
func CallFunc33Callback(func_ Func33) string {
	defer plugify.Scope("cross_call_master::CallFunc33Callback", ModuleName, 3)()
	return _CallFunc33Callback(func_)
}

var _CallFuncEnumCallback = func(func_ FuncEnum) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncEnumCallback(__func_)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// CallFuncEnumCallback
func CallFuncEnumCallback(func_ FuncEnum) string {
	defer plugify.Scope("cross_call_master::CallFuncEnumCallback", ModuleName, 3)()
	return _CallFuncEnumCallback(func_)
}

var (
	ResourceHandleErrEmptyHandle = errors.New("ResourceHandle: empty handle")
)

// @brief RAII wrapper for ResourceHandle pointer
type ResourceHandle struct {
	handle    uintptr
	cleanup   runtime.Cleanup
	ownership ownership
	noCopy    noCopy
}

// NewResourceHandleResourceHandleCreate
func NewResourceHandleResourceHandleCreate(id int32, name string) *ResourceHandle {
	return NewResourceHandleOwned(ResourceHandleCreate(id, name))
}

// NewResourceHandleResourceHandleCreateDefault
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
	w.cleanup = runtime.AddCleanup(w, destroyResourceHandleHandle, handle)
	return w
}

// destroyResourceHandleHandle destroys an owned handle.
func destroyResourceHandleHandle(handle uintptr) {
	if handle != 0 {
		ResourceHandleDestroy(handle)
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
	if w.ownership == Owned && w.handle != 0 {
		w.cleanup.Stop()
	}
	handle := w.handle
	w.nullify()
	return handle
}

// Reset destroys and resets the handle
func (w *ResourceHandle) Reset() {
	if w.ownership == Owned && w.handle != 0 {
		w.cleanup.Stop()
	}
	w.destroy()
	w.nullify()
}

// IsValid returns true if handle is not nil
func (w *ResourceHandle) IsValid() bool {
	return w.handle != 0
}

// GetId
func (w *ResourceHandle) GetId() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, ResourceHandleErrEmptyHandle
	}
	return ResourceHandleGetId(w.handle), nil
}

// GetName
func (w *ResourceHandle) GetName() (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, ResourceHandleErrEmptyHandle
	}
	return ResourceHandleGetName(w.handle), nil
}

// SetName
func (w *ResourceHandle) SetName(name string) error {
	if w.handle == 0 {
		return ResourceHandleErrEmptyHandle
	}
	ResourceHandleSetName(w.handle, name)
	return nil
}

// IncrementCounter
func (w *ResourceHandle) IncrementCounter() error {
	if w.handle == 0 {
		return ResourceHandleErrEmptyHandle
	}
	ResourceHandleIncrementCounter(w.handle)
	return nil
}

// GetCounter
func (w *ResourceHandle) GetCounter() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, ResourceHandleErrEmptyHandle
	}
	return ResourceHandleGetCounter(w.handle), nil
}

// AddData
func (w *ResourceHandle) AddData(value float32) error {
	if w.handle == 0 {
		return ResourceHandleErrEmptyHandle
	}
	ResourceHandleAddData(w.handle, value)
	return nil
}

// GetData
func (w *ResourceHandle) GetData() ([]float32, error) {
	if w.handle == 0 {
		var zero []float32
		return zero, ResourceHandleErrEmptyHandle
	}
	return ResourceHandleGetData(w.handle), nil
}

// GetAliveCount
func (w *ResourceHandle) GetAliveCount() int32 {
	return ResourceHandleGetAliveCount()
}

// GetTotalCreated
func (w *ResourceHandle) GetTotalCreated() int32 {
	return ResourceHandleGetTotalCreated()
}

// GetTotalDestroyed
func (w *ResourceHandle) GetTotalDestroyed() int32 {
	return ResourceHandleGetTotalDestroyed()
}

var (
	CounterErrEmptyHandle = errors.New("Counter: empty handle")
)

type Counter struct {
	handle uintptr
}

// NewCounterCounterCreate
func NewCounterCounterCreate(initialValue int64) *Counter {
	return &Counter{
		handle: CounterCreate(initialValue),
	}
}

// NewCounterCounterCreateZero
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

// GetValue
func (w *Counter) GetValue() (int64, error) {
	if w.handle == 0 {
		var zero int64
		return zero, CounterErrEmptyHandle
	}
	return CounterGetValue(w.handle), nil
}

// SetValue
func (w *Counter) SetValue(value int64) error {
	if w.handle == 0 {
		return CounterErrEmptyHandle
	}
	CounterSetValue(w.handle, value)
	return nil
}

// Increment
func (w *Counter) Increment() error {
	if w.handle == 0 {
		return CounterErrEmptyHandle
	}
	CounterIncrement(w.handle)
	return nil
}

// Decrement
func (w *Counter) Decrement() error {
	if w.handle == 0 {
		return CounterErrEmptyHandle
	}
	CounterDecrement(w.handle)
	return nil
}

// Add
func (w *Counter) Add(amount int64) error {
	if w.handle == 0 {
		return CounterErrEmptyHandle
	}
	CounterAdd(w.handle, amount)
	return nil
}

// Reset
func (w *Counter) Reset2() error {
	if w.handle == 0 {
		return CounterErrEmptyHandle
	}
	CounterReset(w.handle)
	return nil
}

// IsPositive
func (w *Counter) IsPositive() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CounterErrEmptyHandle
	}
	return CounterIsPositive(w.handle), nil
}

// Compare
func (w *Counter) Compare(value1 int64, value2 int64) int32 {
	return CounterCompare(value1, value2)
}

// Sum
func (w *Counter) Sum(values []int64) int64 {
	return CounterSum(values)
}
