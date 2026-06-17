package cross_call_worker

/*
#include "core.h"
#cgo noescape NoParamReturnVoid
#cgo noescape NoParamReturnBool
#cgo noescape NoParamReturnChar8
#cgo noescape NoParamReturnChar16
#cgo noescape NoParamReturnInt8
#cgo noescape NoParamReturnInt16
#cgo noescape NoParamReturnInt32
#cgo noescape NoParamReturnInt64
#cgo noescape NoParamReturnUInt8
#cgo noescape NoParamReturnUInt16
#cgo noescape NoParamReturnUInt32
#cgo noescape NoParamReturnUInt64
#cgo noescape NoParamReturnPointer
#cgo noescape NoParamReturnFloat
#cgo noescape NoParamReturnDouble
#cgo noescape NoParamReturnFunction
#cgo noescape NoParamReturnString
#cgo noescape NoParamReturnAny
#cgo noescape NoParamReturnArrayBool
#cgo noescape NoParamReturnArrayChar8
#cgo noescape NoParamReturnArrayChar16
#cgo noescape NoParamReturnArrayInt8
#cgo noescape NoParamReturnArrayInt16
#cgo noescape NoParamReturnArrayInt32
#cgo noescape NoParamReturnArrayInt64
#cgo noescape NoParamReturnArrayUInt8
#cgo noescape NoParamReturnArrayUInt16
#cgo noescape NoParamReturnArrayUInt32
#cgo noescape NoParamReturnArrayUInt64
#cgo noescape NoParamReturnArrayPointer
#cgo noescape NoParamReturnArrayFloat
#cgo noescape NoParamReturnArrayDouble
#cgo noescape NoParamReturnArrayString
#cgo noescape NoParamReturnArrayAny
#cgo noescape NoParamReturnArrayVector2
#cgo noescape NoParamReturnArrayVector3
#cgo noescape NoParamReturnArrayVector4
#cgo noescape NoParamReturnArrayMatrix4x4
#cgo noescape NoParamReturnVector2
#cgo noescape NoParamReturnVector3
#cgo noescape NoParamReturnVector4
#cgo noescape NoParamReturnMatrix4x4
#cgo noescape Param1
#cgo noescape Param2
#cgo noescape Param3
#cgo noescape Param4
#cgo noescape Param5
#cgo noescape Param6
#cgo noescape Param7
#cgo noescape Param8
#cgo noescape Param9
#cgo noescape Param10
#cgo noescape ParamRef1
#cgo noescape ParamRef2
#cgo noescape ParamRef3
#cgo noescape ParamRef4
#cgo noescape ParamRef5
#cgo noescape ParamRef6
#cgo noescape ParamRef7
#cgo noescape ParamRef8
#cgo noescape ParamRef9
#cgo noescape ParamRef10
#cgo noescape ParamRefVectors
#cgo noescape ParamAllPrimitives
#cgo noescape ParamAllAliases
#cgo noescape ParamAllRefAliases
#cgo noescape ParamVariant
#cgo noescape ParamEnum
#cgo noescape ParamEnumRef
#cgo noescape ParamVariantRef
#cgo noescape CallFuncVoid
#cgo noescape CallFuncBool
#cgo noescape CallFuncChar8
#cgo noescape CallFuncChar16
#cgo noescape CallFuncInt8
#cgo noescape CallFuncInt16
#cgo noescape CallFuncInt32
#cgo noescape CallFuncInt64
#cgo noescape CallFuncUInt8
#cgo noescape CallFuncUInt16
#cgo noescape CallFuncUInt32
#cgo noescape CallFuncUInt64
#cgo noescape CallFuncPtr
#cgo noescape CallFuncFloat
#cgo noescape CallFuncDouble
#cgo noescape CallFuncString
#cgo noescape CallFuncAny
#cgo noescape CallFuncFunction
#cgo noescape CallFuncBoolVector
#cgo noescape CallFuncChar8Vector
#cgo noescape CallFuncChar16Vector
#cgo noescape CallFuncInt8Vector
#cgo noescape CallFuncInt16Vector
#cgo noescape CallFuncInt32Vector
#cgo noescape CallFuncInt64Vector
#cgo noescape CallFuncUInt8Vector
#cgo noescape CallFuncUInt16Vector
#cgo noescape CallFuncUInt32Vector
#cgo noescape CallFuncUInt64Vector
#cgo noescape CallFuncPtrVector
#cgo noescape CallFuncFloatVector
#cgo noescape CallFuncDoubleVector
#cgo noescape CallFuncStringVector
#cgo noescape CallFuncAnyVector
#cgo noescape CallFuncVec2Vector
#cgo noescape CallFuncVec3Vector
#cgo noescape CallFuncVec4Vector
#cgo noescape CallFuncMat4x4Vector
#cgo noescape CallFuncVec2
#cgo noescape CallFuncVec3
#cgo noescape CallFuncVec4
#cgo noescape CallFuncMat4x4
#cgo noescape CallFuncAliasBool
#cgo noescape CallFuncAliasChar8
#cgo noescape CallFuncAliasChar16
#cgo noescape CallFuncAliasInt8
#cgo noescape CallFuncAliasInt16
#cgo noescape CallFuncAliasInt32
#cgo noescape CallFuncAliasInt64
#cgo noescape CallFuncAliasUInt8
#cgo noescape CallFuncAliasUInt16
#cgo noescape CallFuncAliasUInt32
#cgo noescape CallFuncAliasUInt64
#cgo noescape CallFuncAliasPtr
#cgo noescape CallFuncAliasFloat
#cgo noescape CallFuncAliasDouble
#cgo noescape CallFuncAliasString
#cgo noescape CallFuncAliasAny
#cgo noescape CallFuncAliasFunction
#cgo noescape CallFuncAliasBoolVector
#cgo noescape CallFuncAliasChar8Vector
#cgo noescape CallFuncAliasChar16Vector
#cgo noescape CallFuncAliasInt8Vector
#cgo noescape CallFuncAliasInt16Vector
#cgo noescape CallFuncAliasInt32Vector
#cgo noescape CallFuncAliasInt64Vector
#cgo noescape CallFuncAliasUInt8Vector
#cgo noescape CallFuncAliasUInt16Vector
#cgo noescape CallFuncAliasUInt32Vector
#cgo noescape CallFuncAliasUInt64Vector
#cgo noescape CallFuncAliasPtrVector
#cgo noescape CallFuncAliasFloatVector
#cgo noescape CallFuncAliasDoubleVector
#cgo noescape CallFuncAliasStringVector
#cgo noescape CallFuncAliasAnyVector
#cgo noescape CallFuncAliasVec2Vector
#cgo noescape CallFuncAliasVec3Vector
#cgo noescape CallFuncAliasVec4Vector
#cgo noescape CallFuncAliasMat4x4Vector
#cgo noescape CallFuncAliasVec2
#cgo noescape CallFuncAliasVec3
#cgo noescape CallFuncAliasVec4
#cgo noescape CallFuncAliasMat4x4
#cgo noescape CallFuncAliasAll
#cgo noescape CallFunc1
#cgo noescape CallFunc2
#cgo noescape CallFunc3
#cgo noescape CallFunc4
#cgo noescape CallFunc5
#cgo noescape CallFunc6
#cgo noescape CallFunc7
#cgo noescape CallFunc8
#cgo noescape CallFunc9
#cgo noescape CallFunc10
#cgo noescape CallFunc11
#cgo noescape CallFunc12
#cgo noescape CallFunc13
#cgo noescape CallFunc14
#cgo noescape CallFunc15
#cgo noescape CallFunc16
#cgo noescape CallFunc17
#cgo noescape CallFunc18
#cgo noescape CallFunc19
#cgo noescape CallFunc20
#cgo noescape CallFunc21
#cgo noescape CallFunc22
#cgo noescape CallFunc23
#cgo noescape CallFunc24
#cgo noescape CallFunc25
#cgo noescape CallFunc26
#cgo noescape CallFunc27
#cgo noescape CallFunc28
#cgo noescape CallFunc29
#cgo noescape CallFunc30
#cgo noescape CallFunc31
#cgo noescape CallFunc32
#cgo noescape CallFunc33
#cgo noescape CallFuncEnum
#cgo noescape ReverseCall
*/
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/untrustedmodders/go-plugify"
)

var _ = errors.New("")
var _ = reflect.TypeOf(0)
var _ = runtime.GOOS
var _ = unsafe.Sizeof(0)
var _ = plugify.ApiVersion

// Generated from cross_call_worker (group: core)

var P_NoParamReturnVoid = func() {
	fmt.Println("P_NoParamReturnVoid Callling with marshal")
	C.NoParamReturnVoid()
}

// NoParamReturnVoid
func NoParamReturnVoid() {
	defer plugify.Scope("cross_call_worker::NoParamReturnVoid", ModuleName, 3)()
	P_NoParamReturnVoid()
}

var P_NoParamReturnBool = func() bool {
	__retVal := bool(C.NoParamReturnBool())
	return __retVal
}

// NoParamReturnBool
func NoParamReturnBool() bool {
	defer plugify.Scope("cross_call_worker::NoParamReturnBool", ModuleName, 3)()
	return P_NoParamReturnBool()
}

var P_NoParamReturnChar8 = func() int8 {
	__retVal := int8(C.NoParamReturnChar8())
	return __retVal
}

// NoParamReturnChar8
func NoParamReturnChar8() int8 {
	defer plugify.Scope("cross_call_worker::NoParamReturnChar8", ModuleName, 3)()
	return P_NoParamReturnChar8()
}

var P_NoParamReturnChar16 = func() uint16 {
	__retVal := uint16(C.NoParamReturnChar16())
	return __retVal
}

// NoParamReturnChar16
func NoParamReturnChar16() uint16 {
	defer plugify.Scope("cross_call_worker::NoParamReturnChar16", ModuleName, 3)()
	return P_NoParamReturnChar16()
}

var P_NoParamReturnInt8 = func() int8 {
	__retVal := int8(C.NoParamReturnInt8())
	return __retVal
}

// NoParamReturnInt8
func NoParamReturnInt8() int8 {
	defer plugify.Scope("cross_call_worker::NoParamReturnInt8", ModuleName, 3)()
	return P_NoParamReturnInt8()
}

var P_NoParamReturnInt16 = func() int16 {
	__retVal := int16(C.NoParamReturnInt16())
	return __retVal
}

// NoParamReturnInt16
func NoParamReturnInt16() int16 {
	defer plugify.Scope("cross_call_worker::NoParamReturnInt16", ModuleName, 3)()
	return P_NoParamReturnInt16()
}

var P_NoParamReturnInt32 = func() int32 {
	__retVal := int32(C.NoParamReturnInt32())
	return __retVal
}

// NoParamReturnInt32
func NoParamReturnInt32() int32 {
	defer plugify.Scope("cross_call_worker::NoParamReturnInt32", ModuleName, 3)()
	return P_NoParamReturnInt32()
}

var P_NoParamReturnInt64 = func() int64 {
	__retVal := int64(C.NoParamReturnInt64())
	return __retVal
}

// NoParamReturnInt64
func NoParamReturnInt64() int64 {
	defer plugify.Scope("cross_call_worker::NoParamReturnInt64", ModuleName, 3)()
	return P_NoParamReturnInt64()
}

var P_NoParamReturnUInt8 = func() uint8 {
	__retVal := uint8(C.NoParamReturnUInt8())
	return __retVal
}

// NoParamReturnUInt8
func NoParamReturnUInt8() uint8 {
	defer plugify.Scope("cross_call_worker::NoParamReturnUInt8", ModuleName, 3)()
	return P_NoParamReturnUInt8()
}

var P_NoParamReturnUInt16 = func() uint16 {
	__retVal := uint16(C.NoParamReturnUInt16())
	return __retVal
}

// NoParamReturnUInt16
func NoParamReturnUInt16() uint16 {
	defer plugify.Scope("cross_call_worker::NoParamReturnUInt16", ModuleName, 3)()
	return P_NoParamReturnUInt16()
}

var P_NoParamReturnUInt32 = func() uint32 {
	__retVal := uint32(C.NoParamReturnUInt32())
	return __retVal
}

// NoParamReturnUInt32
func NoParamReturnUInt32() uint32 {
	defer plugify.Scope("cross_call_worker::NoParamReturnUInt32", ModuleName, 3)()
	return P_NoParamReturnUInt32()
}

var P_NoParamReturnUInt64 = func() uint64 {
	__retVal := uint64(C.NoParamReturnUInt64())
	return __retVal
}

// NoParamReturnUInt64
func NoParamReturnUInt64() uint64 {
	defer plugify.Scope("cross_call_worker::NoParamReturnUInt64", ModuleName, 3)()
	return P_NoParamReturnUInt64()
}

var P_NoParamReturnPointer = func() uintptr {
	__retVal := uintptr(C.NoParamReturnPointer())
	return __retVal
}

// NoParamReturnPointer
func NoParamReturnPointer() uintptr {
	defer plugify.Scope("cross_call_worker::NoParamReturnPointer", ModuleName, 3)()
	return P_NoParamReturnPointer()
}

var P_NoParamReturnFloat = func() float32 {
	__retVal := float32(C.NoParamReturnFloat())
	return __retVal
}

// NoParamReturnFloat
func NoParamReturnFloat() float32 {
	defer plugify.Scope("cross_call_worker::NoParamReturnFloat", ModuleName, 3)()
	return P_NoParamReturnFloat()
}

var P_NoParamReturnDouble = func() float64 {
	__retVal := float64(C.NoParamReturnDouble())
	return __retVal
}

// NoParamReturnDouble
func NoParamReturnDouble() float64 {
	defer plugify.Scope("cross_call_worker::NoParamReturnDouble", ModuleName, 3)()
	return P_NoParamReturnDouble()
}

var P_NoParamReturnFunction = func() NoParamReturnFunctionFunc {
	__retVal := plugify.GetDelegateForFunctionPointer(C.NoParamReturnFunction(), reflect.TypeOf(NoParamReturnFunctionFunc(nil))).(NoParamReturnFunctionFunc)
	return __retVal
}

// NoParamReturnFunction
func NoParamReturnFunction() NoParamReturnFunctionFunc {
	defer plugify.Scope("cross_call_worker::NoParamReturnFunction", ModuleName, 3)()
	return P_NoParamReturnFunction()
}

var P_NoParamReturnString = func() string {
	var __retVal string
	var __retVal_native plugify.PlgString
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnString()
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

// NoParamReturnString
func NoParamReturnString() string {
	defer plugify.Scope("cross_call_worker::NoParamReturnString", ModuleName, 3)()
	return P_NoParamReturnString()
}

var P_NoParamReturnAny = func() any {
	var __retVal any
	var __retVal_native plugify.PlgVariant
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnAny()
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

// NoParamReturnAny
func NoParamReturnAny() any {
	defer plugify.Scope("cross_call_worker::NoParamReturnAny", ModuleName, 3)()
	return P_NoParamReturnAny()
}

var P_NoParamReturnArrayBool = func() []bool {
	var __retVal []bool
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayBool()
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

// NoParamReturnArrayBool
func NoParamReturnArrayBool() []bool {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayBool", ModuleName, 3)()
	return P_NoParamReturnArrayBool()
}

var P_NoParamReturnArrayChar8 = func() []int8 {
	var __retVal []int8
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayChar8()
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

// NoParamReturnArrayChar8
func NoParamReturnArrayChar8() []int8 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayChar8", ModuleName, 3)()
	return P_NoParamReturnArrayChar8()
}

var P_NoParamReturnArrayChar16 = func() []uint16 {
	var __retVal []uint16
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayChar16()
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

// NoParamReturnArrayChar16
func NoParamReturnArrayChar16() []uint16 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayChar16", ModuleName, 3)()
	return P_NoParamReturnArrayChar16()
}

var P_NoParamReturnArrayInt8 = func() []int8 {
	var __retVal []int8
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayInt8()
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

// NoParamReturnArrayInt8
func NoParamReturnArrayInt8() []int8 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayInt8", ModuleName, 3)()
	return P_NoParamReturnArrayInt8()
}

var P_NoParamReturnArrayInt16 = func() []int16 {
	var __retVal []int16
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayInt16()
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

// NoParamReturnArrayInt16
func NoParamReturnArrayInt16() []int16 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayInt16", ModuleName, 3)()
	return P_NoParamReturnArrayInt16()
}

var P_NoParamReturnArrayInt32 = func() []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayInt32()
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

// NoParamReturnArrayInt32
func NoParamReturnArrayInt32() []int32 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayInt32", ModuleName, 3)()
	return P_NoParamReturnArrayInt32()
}

var P_NoParamReturnArrayInt64 = func() []int64 {
	var __retVal []int64
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayInt64()
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

// NoParamReturnArrayInt64
func NoParamReturnArrayInt64() []int64 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayInt64", ModuleName, 3)()
	return P_NoParamReturnArrayInt64()
}

var P_NoParamReturnArrayUInt8 = func() []uint8 {
	var __retVal []uint8
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayUInt8()
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

// NoParamReturnArrayUInt8
func NoParamReturnArrayUInt8() []uint8 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayUInt8", ModuleName, 3)()
	return P_NoParamReturnArrayUInt8()
}

var P_NoParamReturnArrayUInt16 = func() []uint16 {
	var __retVal []uint16
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayUInt16()
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

// NoParamReturnArrayUInt16
func NoParamReturnArrayUInt16() []uint16 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayUInt16", ModuleName, 3)()
	return P_NoParamReturnArrayUInt16()
}

var P_NoParamReturnArrayUInt32 = func() []uint32 {
	var __retVal []uint32
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayUInt32()
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

// NoParamReturnArrayUInt32
func NoParamReturnArrayUInt32() []uint32 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayUInt32", ModuleName, 3)()
	return P_NoParamReturnArrayUInt32()
}

var P_NoParamReturnArrayUInt64 = func() []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayUInt64()
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

// NoParamReturnArrayUInt64
func NoParamReturnArrayUInt64() []uint64 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayUInt64", ModuleName, 3)()
	return P_NoParamReturnArrayUInt64()
}

var P_NoParamReturnArrayPointer = func() []uintptr {
	var __retVal []uintptr
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayPointer()
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

// NoParamReturnArrayPointer
func NoParamReturnArrayPointer() []uintptr {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayPointer", ModuleName, 3)()
	return P_NoParamReturnArrayPointer()
}

var P_NoParamReturnArrayFloat = func() []float32 {
	var __retVal []float32
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayFloat()
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

// NoParamReturnArrayFloat
func NoParamReturnArrayFloat() []float32 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayFloat", ModuleName, 3)()
	return P_NoParamReturnArrayFloat()
}

var P_NoParamReturnArrayDouble = func() []float64 {
	var __retVal []float64
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayDouble()
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

// NoParamReturnArrayDouble
func NoParamReturnArrayDouble() []float64 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayDouble", ModuleName, 3)()
	return P_NoParamReturnArrayDouble()
}

var P_NoParamReturnArrayString = func() []string {
	var __retVal []string
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayString()
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

// NoParamReturnArrayString
func NoParamReturnArrayString() []string {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayString", ModuleName, 3)()
	return P_NoParamReturnArrayString()
}

var P_NoParamReturnArrayAny = func() []any {
	var __retVal []any
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayAny()
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

// NoParamReturnArrayAny
func NoParamReturnArrayAny() []any {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayAny", ModuleName, 3)()
	return P_NoParamReturnArrayAny()
}

var P_NoParamReturnArrayVector2 = func() []plugify.Vector2 {
	var __retVal []plugify.Vector2
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayVector2()
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

// NoParamReturnArrayVector2
func NoParamReturnArrayVector2() []plugify.Vector2 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayVector2", ModuleName, 3)()
	return P_NoParamReturnArrayVector2()
}

var P_NoParamReturnArrayVector3 = func() []plugify.Vector3 {
	var __retVal []plugify.Vector3
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayVector3()
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

// NoParamReturnArrayVector3
func NoParamReturnArrayVector3() []plugify.Vector3 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayVector3", ModuleName, 3)()
	return P_NoParamReturnArrayVector3()
}

var P_NoParamReturnArrayVector4 = func() []plugify.Vector4 {
	var __retVal []plugify.Vector4
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayVector4()
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

// NoParamReturnArrayVector4
func NoParamReturnArrayVector4() []plugify.Vector4 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayVector4", ModuleName, 3)()
	return P_NoParamReturnArrayVector4()
}

var P_NoParamReturnArrayMatrix4x4 = func() []plugify.Matrix4x4 {
	var __retVal []plugify.Matrix4x4
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.NoParamReturnArrayMatrix4x4()
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

// NoParamReturnArrayMatrix4x4
func NoParamReturnArrayMatrix4x4() []plugify.Matrix4x4 {
	defer plugify.Scope("cross_call_worker::NoParamReturnArrayMatrix4x4", ModuleName, 3)()
	return P_NoParamReturnArrayMatrix4x4()
}

var P_NoParamReturnVector2 = func() plugify.Vector2 {
	__native := C.NoParamReturnVector2()
	__retVal := *(*plugify.Vector2)(unsafe.Pointer(&__native))
	return __retVal
}

// NoParamReturnVector2
func NoParamReturnVector2() plugify.Vector2 {
	defer plugify.Scope("cross_call_worker::NoParamReturnVector2", ModuleName, 3)()
	return P_NoParamReturnVector2()
}

var P_NoParamReturnVector3 = func() plugify.Vector3 {
	__native := C.NoParamReturnVector3()
	__retVal := *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// NoParamReturnVector3
func NoParamReturnVector3() plugify.Vector3 {
	defer plugify.Scope("cross_call_worker::NoParamReturnVector3", ModuleName, 3)()
	return P_NoParamReturnVector3()
}

var P_NoParamReturnVector4 = func() plugify.Vector4 {
	__native := C.NoParamReturnVector4()
	__retVal := *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// NoParamReturnVector4
func NoParamReturnVector4() plugify.Vector4 {
	defer plugify.Scope("cross_call_worker::NoParamReturnVector4", ModuleName, 3)()
	return P_NoParamReturnVector4()
}

var P_NoParamReturnMatrix4x4 = func() plugify.Matrix4x4 {
	__native := C.NoParamReturnMatrix4x4()
	__retVal := *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// NoParamReturnMatrix4x4
func NoParamReturnMatrix4x4() plugify.Matrix4x4 {
	defer plugify.Scope("cross_call_worker::NoParamReturnMatrix4x4", ModuleName, 3)()
	return P_NoParamReturnMatrix4x4()
}

var P_Param1 = func(a int32) {
	__a := C.int32_t(a)
	C.Param1(__a)
}

// Param1
func Param1(a int32) {
	defer plugify.Scope("cross_call_worker::Param1", ModuleName, 3)()
	P_Param1(a)
}

var P_Param2 = func(a int32, b float32) {
	__a := C.int32_t(a)
	__b := C.float(b)
	C.Param2(__a, __b)
}

// Param2
func Param2(a int32, b float32) {
	defer plugify.Scope("cross_call_worker::Param2", ModuleName, 3)()
	P_Param2(a, b)
}

var P_Param3 = func(a int32, b float32, c float64) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	C.Param3(__a, __b, __c)
}

// Param3
func Param3(a int32, b float32, c float64) {
	defer plugify.Scope("cross_call_worker::Param3", ModuleName, 3)()
	P_Param3(a, b, c)
}

var P_Param4 = func(a int32, b float32, c float64, d plugify.Vector4) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	C.Param4(__a, __b, __c, &__d)
}

// Param4
func Param4(a int32, b float32, c float64, d plugify.Vector4) {
	defer plugify.Scope("cross_call_worker::Param4", ModuleName, 3)()
	P_Param4(a, b, c, d)
}

var P_Param5 = func(a int32, b float32, c float64, d plugify.Vector4, e []int64) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	__e := plugify.ConstructVectorInt64(e)
	plugify.Block{
		Try: func() {
			C.Param5(__a, __b, __c, &__d, (*C.Vector)(unsafe.Pointer(&__e)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
		},
	}.Do()
}

// Param5
func Param5(a int32, b float32, c float64, d plugify.Vector4, e []int64) {
	defer plugify.Scope("cross_call_worker::Param5", ModuleName, 3)()
	P_Param5(a, b, c, d, e)
}

var P_Param6 = func(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	__e := plugify.ConstructVectorInt64(e)
	__f := C.int8_t(f)
	plugify.Block{
		Try: func() {
			C.Param6(__a, __b, __c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), __f)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
		},
	}.Do()
}

// Param6
func Param6(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8) {
	defer plugify.Scope("cross_call_worker::Param6", ModuleName, 3)()
	P_Param6(a, b, c, d, e, f)
}

var P_Param7 = func(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string) {
	__a := C.int32_t(a)
	__b := C.float(b)
	__c := C.double(c)
	__d := *(*C.Vector4)(unsafe.Pointer(&d))
	__e := plugify.ConstructVectorInt64(e)
	__f := C.int8_t(f)
	__g := plugify.ConstructString(g)
	plugify.Block{
		Try: func() {
			C.Param7(__a, __b, __c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), __f, (*C.String)(unsafe.Pointer(&__g)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// Param7
func Param7(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string) {
	defer plugify.Scope("cross_call_worker::Param7", ModuleName, 3)()
	P_Param7(a, b, c, d, e, f, g)
}

var P_Param8 = func(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16) {
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
			C.Param8(__a, __b, __c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), __f, (*C.String)(unsafe.Pointer(&__g)), __h)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// Param8
func Param8(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16) {
	defer plugify.Scope("cross_call_worker::Param8", ModuleName, 3)()
	P_Param8(a, b, c, d, e, f, g, h)
}

var P_Param9 = func(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16, k int16) {
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
			C.Param9(__a, __b, __c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), __f, (*C.String)(unsafe.Pointer(&__g)), __h, __k)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// Param9
func Param9(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16, k int16) {
	defer plugify.Scope("cross_call_worker::Param9", ModuleName, 3)()
	P_Param9(a, b, c, d, e, f, g, h, k)
}

var P_Param10 = func(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16, k int16, l uintptr) {
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
			C.Param10(__a, __b, __c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), __f, (*C.String)(unsafe.Pointer(&__g)), __h, __k, __l)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__e)
			plugify.DestroyString(&__g)
		},
	}.Do()
}

// Param10
func Param10(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16, k int16, l uintptr) {
	defer plugify.Scope("cross_call_worker::Param10", ModuleName, 3)()
	P_Param10(a, b, c, d, e, f, g, h, k, l)
}

var P_ParamRef1 = func(a *int32) {
	__a := C.int32_t(*a)
	C.ParamRef1(&__a)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
}

// ParamRef1
func ParamRef1(a *int32) {
	defer plugify.Scope("cross_call_worker::ParamRef1", ModuleName, 3)()
	P_ParamRef1(a)
}

var P_ParamRef2 = func(a *int32, b *float32) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	C.ParamRef2(&__a, &__b)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
	*b = float32(__b)
}

// ParamRef2
func ParamRef2(a *int32, b *float32) {
	defer plugify.Scope("cross_call_worker::ParamRef2", ModuleName, 3)()
	P_ParamRef2(a, b)
}

var P_ParamRef3 = func(a *int32, b *float32, c *float64) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	C.ParamRef3(&__a, &__b, &__c)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
	*b = float32(__b)
	*c = float64(__c)
}

// ParamRef3
func ParamRef3(a *int32, b *float32, c *float64) {
	defer plugify.Scope("cross_call_worker::ParamRef3", ModuleName, 3)()
	P_ParamRef3(a, b, c)
}

var P_ParamRef4 = func(a *int32, b *float32, c *float64, d *plugify.Vector4) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	__d := *(*C.Vector4)(unsafe.Pointer(d))
	C.ParamRef4(&__a, &__b, &__c, &__d)
	// Unmarshal - Convert native data to managed data.
	*a = int32(__a)
	*b = float32(__b)
	*c = float64(__c)
	*d = *(*plugify.Vector4)(unsafe.Pointer(&__d))
}

// ParamRef4
func ParamRef4(a *int32, b *float32, c *float64, d *plugify.Vector4) {
	defer plugify.Scope("cross_call_worker::ParamRef4", ModuleName, 3)()
	P_ParamRef4(a, b, c, d)
}

var P_ParamRef5 = func(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	__d := *(*C.Vector4)(unsafe.Pointer(d))
	__e := plugify.ConstructVectorInt64(*e)
	plugify.Block{
		Try: func() {
			C.ParamRef5(&__a, &__b, &__c, &__d, (*C.Vector)(unsafe.Pointer(&__e)))
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

// ParamRef5
func ParamRef5(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64) {
	defer plugify.Scope("cross_call_worker::ParamRef5", ModuleName, 3)()
	P_ParamRef5(a, b, c, d, e)
}

var P_ParamRef6 = func(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	__d := *(*C.Vector4)(unsafe.Pointer(d))
	__e := plugify.ConstructVectorInt64(*e)
	__f := C.int8_t(*f)
	plugify.Block{
		Try: func() {
			C.ParamRef6(&__a, &__b, &__c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), &__f)
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

// ParamRef6
func ParamRef6(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8) {
	defer plugify.Scope("cross_call_worker::ParamRef6", ModuleName, 3)()
	P_ParamRef6(a, b, c, d, e, f)
}

var P_ParamRef7 = func(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string) {
	__a := C.int32_t(*a)
	__b := C.float(*b)
	__c := C.double(*c)
	__d := *(*C.Vector4)(unsafe.Pointer(d))
	__e := plugify.ConstructVectorInt64(*e)
	__f := C.int8_t(*f)
	__g := plugify.ConstructString(*g)
	plugify.Block{
		Try: func() {
			C.ParamRef7(&__a, &__b, &__c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), &__f, (*C.String)(unsafe.Pointer(&__g)))
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

// ParamRef7
func ParamRef7(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string) {
	defer plugify.Scope("cross_call_worker::ParamRef7", ModuleName, 3)()
	P_ParamRef7(a, b, c, d, e, f, g)
}

var P_ParamRef8 = func(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16) {
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
			C.ParamRef8(&__a, &__b, &__c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), &__f, (*C.String)(unsafe.Pointer(&__g)), &__h)
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

// ParamRef8
func ParamRef8(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16) {
	defer plugify.Scope("cross_call_worker::ParamRef8", ModuleName, 3)()
	P_ParamRef8(a, b, c, d, e, f, g, h)
}

var P_ParamRef9 = func(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16) {
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
			C.ParamRef9(&__a, &__b, &__c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), &__f, (*C.String)(unsafe.Pointer(&__g)), &__h, &__k)
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

// ParamRef9
func ParamRef9(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16) {
	defer plugify.Scope("cross_call_worker::ParamRef9", ModuleName, 3)()
	P_ParamRef9(a, b, c, d, e, f, g, h, k)
}

var P_ParamRef10 = func(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16, l *uintptr) {
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
			C.ParamRef10(&__a, &__b, &__c, &__d, (*C.Vector)(unsafe.Pointer(&__e)), &__f, (*C.String)(unsafe.Pointer(&__g)), &__h, &__k, &__l)
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

// ParamRef10
func ParamRef10(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16, l *uintptr) {
	defer plugify.Scope("cross_call_worker::ParamRef10", ModuleName, 3)()
	P_ParamRef10(a, b, c, d, e, f, g, h, k, l)
}

var P_ParamRefVectors = func(p1 *[]bool, p2 *[]int8, p3 *[]uint16, p4 *[]int8, p5 *[]int16, p6 *[]int32, p7 *[]int64, p8 *[]uint8, p9 *[]uint16, p10 *[]uint32, p11 *[]uint64, p12 *[]uintptr, p13 *[]float32, p14 *[]float64, p15 *[]string) {
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
			C.ParamRefVectors((*C.Vector)(unsafe.Pointer(&__p1)), (*C.Vector)(unsafe.Pointer(&__p2)), (*C.Vector)(unsafe.Pointer(&__p3)), (*C.Vector)(unsafe.Pointer(&__p4)), (*C.Vector)(unsafe.Pointer(&__p5)), (*C.Vector)(unsafe.Pointer(&__p6)), (*C.Vector)(unsafe.Pointer(&__p7)), (*C.Vector)(unsafe.Pointer(&__p8)), (*C.Vector)(unsafe.Pointer(&__p9)), (*C.Vector)(unsafe.Pointer(&__p10)), (*C.Vector)(unsafe.Pointer(&__p11)), (*C.Vector)(unsafe.Pointer(&__p12)), (*C.Vector)(unsafe.Pointer(&__p13)), (*C.Vector)(unsafe.Pointer(&__p14)), (*C.Vector)(unsafe.Pointer(&__p15)))
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

// ParamRefVectors
func ParamRefVectors(p1 *[]bool, p2 *[]int8, p3 *[]uint16, p4 *[]int8, p5 *[]int16, p6 *[]int32, p7 *[]int64, p8 *[]uint8, p9 *[]uint16, p10 *[]uint32, p11 *[]uint64, p12 *[]uintptr, p13 *[]float32, p14 *[]float64, p15 *[]string) {
	defer plugify.Scope("cross_call_worker::ParamRefVectors", ModuleName, 3)()
	P_ParamRefVectors(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15)
}

var P_ParamAllPrimitives = func(p1 bool, p2 int8, p3 uint16, p4 int8, p5 int16, p6 int32, p7 int64, p8 uint8, p9 uint16, p10 uint32, p11 uint64, p12 uintptr, p13 float32, p14 float64) int64 {
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
	__retVal = int64(C.ParamAllPrimitives(__p1, __p2, __p3, __p4, __p5, __p6, __p7, __p8, __p9, __p10, __p11, __p12, __p13, __p14))
	return __retVal
}

// ParamAllPrimitives
func ParamAllPrimitives(p1 bool, p2 int8, p3 uint16, p4 int8, p5 int16, p6 int32, p7 int64, p8 uint8, p9 uint16, p10 uint32, p11 uint64, p12 uintptr, p13 float32, p14 float64) int64 {
	defer plugify.Scope("cross_call_worker::ParamAllPrimitives", ModuleName, 3)()
	return P_ParamAllPrimitives(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14)
}

var P_ParamAllAliases = func(aBool AliasBool, aChar8 AliasChar8, aChar16 AliasChar16, aInt8 AliasInt8, aInt16 AliasInt16, aInt32 AliasInt32, aInt64 AliasInt64, aPtr AliasPtr, aFloat AliasFloat, aDouble AliasDouble, aString AliasString, aAny AliasAny, aVec2 AliasVec2, aVec3 AliasVec3, aVec4 AliasVec4, aMat4x4 AliasMat4x4, aBoolVec AliasBoolVector, aChar8Vec AliasChar8Vector, aChar16Vec AliasChar16Vector, aInt8Vec AliasInt8Vector, aInt16Vec AliasInt16Vector, aInt32Vec AliasInt32Vector, aInt64Vec AliasInt64Vector, aPtrVec AliasPtrVector, aFloatVec AliasFloatVector, aDoubleVec AliasDoubleVector, aStringVec AliasStringVector, aAnyVec AliasAnyVector, aVec2Vec AliasVec2Vector, aVec3Vec AliasVec3Vector, aVec4Vec AliasVec4Vector) int32 {
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
			__retVal = int32(C.ParamAllAliases(__aBool, __aChar8, __aChar16, __aInt8, __aInt16, __aInt32, __aInt64, __aPtr, __aFloat, __aDouble, (*C.String)(unsafe.Pointer(&__aString)), (*C.Variant)(unsafe.Pointer(&__aAny)), &__aVec2, &__aVec3, &__aVec4, &__aMat4x4, (*C.Vector)(unsafe.Pointer(&__aBoolVec)), (*C.Vector)(unsafe.Pointer(&__aChar8Vec)), (*C.Vector)(unsafe.Pointer(&__aChar16Vec)), (*C.Vector)(unsafe.Pointer(&__aInt8Vec)), (*C.Vector)(unsafe.Pointer(&__aInt16Vec)), (*C.Vector)(unsafe.Pointer(&__aInt32Vec)), (*C.Vector)(unsafe.Pointer(&__aInt64Vec)), (*C.Vector)(unsafe.Pointer(&__aPtrVec)), (*C.Vector)(unsafe.Pointer(&__aFloatVec)), (*C.Vector)(unsafe.Pointer(&__aDoubleVec)), (*C.Vector)(unsafe.Pointer(&__aStringVec)), (*C.Vector)(unsafe.Pointer(&__aAnyVec)), (*C.Vector)(unsafe.Pointer(&__aVec2Vec)), (*C.Vector)(unsafe.Pointer(&__aVec3Vec)), (*C.Vector)(unsafe.Pointer(&__aVec4Vec))))
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

// ParamAllAliases
func ParamAllAliases(aBool AliasBool, aChar8 AliasChar8, aChar16 AliasChar16, aInt8 AliasInt8, aInt16 AliasInt16, aInt32 AliasInt32, aInt64 AliasInt64, aPtr AliasPtr, aFloat AliasFloat, aDouble AliasDouble, aString AliasString, aAny AliasAny, aVec2 AliasVec2, aVec3 AliasVec3, aVec4 AliasVec4, aMat4x4 AliasMat4x4, aBoolVec AliasBoolVector, aChar8Vec AliasChar8Vector, aChar16Vec AliasChar16Vector, aInt8Vec AliasInt8Vector, aInt16Vec AliasInt16Vector, aInt32Vec AliasInt32Vector, aInt64Vec AliasInt64Vector, aPtrVec AliasPtrVector, aFloatVec AliasFloatVector, aDoubleVec AliasDoubleVector, aStringVec AliasStringVector, aAnyVec AliasAnyVector, aVec2Vec AliasVec2Vector, aVec3Vec AliasVec3Vector, aVec4Vec AliasVec4Vector) int32 {
	defer plugify.Scope("cross_call_worker::ParamAllAliases", ModuleName, 3)()
	return P_ParamAllAliases(aBool, aChar8, aChar16, aInt8, aInt16, aInt32, aInt64, aPtr, aFloat, aDouble, aString, aAny, aVec2, aVec3, aVec4, aMat4x4, aBoolVec, aChar8Vec, aChar16Vec, aInt8Vec, aInt16Vec, aInt32Vec, aInt64Vec, aPtrVec, aFloatVec, aDoubleVec, aStringVec, aAnyVec, aVec2Vec, aVec3Vec, aVec4Vec)
}

var P_ParamAllRefAliases = func(aBool *AliasBool, aChar8 *AliasChar8, aChar16 *AliasChar16, aInt8 *AliasInt8, aInt16 *AliasInt16, aInt32 *AliasInt32, aInt64 *AliasInt64, aPtr *AliasPtr, aFloat *AliasFloat, aDouble *AliasDouble, aString *AliasString, aAny *AliasAny, aVec2 *AliasVec2, aVec3 *AliasVec3, aVec4 *AliasVec4, aMat4x4 *AliasMat4x4, aBoolVec *AliasBoolVector, aChar8Vec *AliasChar8Vector, aChar16Vec *AliasChar16Vector, aInt8Vec *AliasInt8Vector, aInt16Vec *AliasInt16Vector, aInt32Vec *AliasInt32Vector, aInt64Vec *AliasInt64Vector, aPtrVec *AliasPtrVector, aFloatVec *AliasFloatVector, aDoubleVec *AliasDoubleVector, aStringVec *AliasStringVector, aAnyVec *AliasAnyVector, aVec2Vec *AliasVec2Vector, aVec3Vec *AliasVec3Vector, aVec4Vec *AliasVec4Vector) int64 {
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
			__retVal = int64(C.ParamAllRefAliases(&__aBool, &__aChar8, &__aChar16, &__aInt8, &__aInt16, &__aInt32, &__aInt64, &__aPtr, &__aFloat, &__aDouble, (*C.String)(unsafe.Pointer(&__aString)), (*C.Variant)(unsafe.Pointer(&__aAny)), &__aVec2, &__aVec3, &__aVec4, &__aMat4x4, (*C.Vector)(unsafe.Pointer(&__aBoolVec)), (*C.Vector)(unsafe.Pointer(&__aChar8Vec)), (*C.Vector)(unsafe.Pointer(&__aChar16Vec)), (*C.Vector)(unsafe.Pointer(&__aInt8Vec)), (*C.Vector)(unsafe.Pointer(&__aInt16Vec)), (*C.Vector)(unsafe.Pointer(&__aInt32Vec)), (*C.Vector)(unsafe.Pointer(&__aInt64Vec)), (*C.Vector)(unsafe.Pointer(&__aPtrVec)), (*C.Vector)(unsafe.Pointer(&__aFloatVec)), (*C.Vector)(unsafe.Pointer(&__aDoubleVec)), (*C.Vector)(unsafe.Pointer(&__aStringVec)), (*C.Vector)(unsafe.Pointer(&__aAnyVec)), (*C.Vector)(unsafe.Pointer(&__aVec2Vec)), (*C.Vector)(unsafe.Pointer(&__aVec3Vec)), (*C.Vector)(unsafe.Pointer(&__aVec4Vec))))
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

// ParamAllRefAliases
func ParamAllRefAliases(aBool *AliasBool, aChar8 *AliasChar8, aChar16 *AliasChar16, aInt8 *AliasInt8, aInt16 *AliasInt16, aInt32 *AliasInt32, aInt64 *AliasInt64, aPtr *AliasPtr, aFloat *AliasFloat, aDouble *AliasDouble, aString *AliasString, aAny *AliasAny, aVec2 *AliasVec2, aVec3 *AliasVec3, aVec4 *AliasVec4, aMat4x4 *AliasMat4x4, aBoolVec *AliasBoolVector, aChar8Vec *AliasChar8Vector, aChar16Vec *AliasChar16Vector, aInt8Vec *AliasInt8Vector, aInt16Vec *AliasInt16Vector, aInt32Vec *AliasInt32Vector, aInt64Vec *AliasInt64Vector, aPtrVec *AliasPtrVector, aFloatVec *AliasFloatVector, aDoubleVec *AliasDoubleVector, aStringVec *AliasStringVector, aAnyVec *AliasAnyVector, aVec2Vec *AliasVec2Vector, aVec3Vec *AliasVec3Vector, aVec4Vec *AliasVec4Vector) int64 {
	defer plugify.Scope("cross_call_worker::ParamAllRefAliases", ModuleName, 3)()
	return P_ParamAllRefAliases(aBool, aChar8, aChar16, aInt8, aInt16, aInt32, aInt64, aPtr, aFloat, aDouble, aString, aAny, aVec2, aVec3, aVec4, aMat4x4, aBoolVec, aChar8Vec, aChar16Vec, aInt8Vec, aInt16Vec, aInt32Vec, aInt64Vec, aPtrVec, aFloatVec, aDoubleVec, aStringVec, aAnyVec, aVec2Vec, aVec3Vec, aVec4Vec)
}

var P_ParamVariant = func(p1 any, p2 []any) {
	__p1 := plugify.ConstructVariant(p1)
	__p2 := plugify.ConstructVectorVariant(p2)
	plugify.Block{
		Try: func() {
			C.ParamVariant((*C.Variant)(unsafe.Pointer(&__p1)), (*C.Vector)(unsafe.Pointer(&__p2)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__p1)
			plugify.DestroyVectorVariant(&__p2)
		},
	}.Do()
}

// ParamVariant
func ParamVariant(p1 any, p2 []any) {
	defer plugify.Scope("cross_call_worker::ParamVariant", ModuleName, 3)()
	P_ParamVariant(p1, p2)
}

var P_ParamEnum = func(p1 Example, p2 []Example) int32 {
	var __retVal int32
	__p1 := C.int32_t(p1)
	__p2 := plugify.ConstructVectorInt32(p2)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.ParamEnum(__p1, (*C.Vector)(unsafe.Pointer(&__p2))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__p2)
		},
	}.Do()
	return __retVal
}

// ParamEnum
func ParamEnum(p1 Example, p2 []Example) int32 {
	defer plugify.Scope("cross_call_worker::ParamEnum", ModuleName, 3)()
	return P_ParamEnum(p1, p2)
}

var P_ParamEnumRef = func(p1 *Example, p2 *[]Example) int32 {
	var __retVal int32
	__p1 := C.int32_t(*p1)
	__p2 := plugify.ConstructVectorInt32(*p2)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.ParamEnumRef(&__p1, (*C.Vector)(unsafe.Pointer(&__p2))))
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

// ParamEnumRef
func ParamEnumRef(p1 *Example, p2 *[]Example) int32 {
	defer plugify.Scope("cross_call_worker::ParamEnumRef", ModuleName, 3)()
	return P_ParamEnumRef(p1, p2)
}

var P_ParamVariantRef = func(p1 *any, p2 *[]any) {
	__p1 := plugify.ConstructVariant(*p1)
	__p2 := plugify.ConstructVectorVariant(*p2)
	plugify.Block{
		Try: func() {
			C.ParamVariantRef((*C.Variant)(unsafe.Pointer(&__p1)), (*C.Vector)(unsafe.Pointer(&__p2)))
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

// ParamVariantRef
func ParamVariantRef(p1 *any, p2 *[]any) {
	defer plugify.Scope("cross_call_worker::ParamVariantRef", ModuleName, 3)()
	P_ParamVariantRef(p1, p2)
}

var P_CallFuncVoid = func(func_ FuncVoid) {
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	C.CallFuncVoid(__func_)
}

// CallFuncVoid
func CallFuncVoid(func_ FuncVoid) {
	defer plugify.Scope("cross_call_worker::CallFuncVoid", ModuleName, 3)()
	P_CallFuncVoid(func_)
}

var P_CallFuncBool = func(func_ FuncBool) bool {
	var __retVal bool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = bool(C.CallFuncBool(__func_))
	return __retVal
}

// CallFuncBool
func CallFuncBool(func_ FuncBool) bool {
	defer plugify.Scope("cross_call_worker::CallFuncBool", ModuleName, 3)()
	return P_CallFuncBool(func_)
}

var P_CallFuncChar8 = func(func_ FuncChar8) int8 {
	var __retVal int8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int8(C.CallFuncChar8(__func_))
	return __retVal
}

// CallFuncChar8
func CallFuncChar8(func_ FuncChar8) int8 {
	defer plugify.Scope("cross_call_worker::CallFuncChar8", ModuleName, 3)()
	return P_CallFuncChar8(func_)
}

var P_CallFuncChar16 = func(func_ FuncChar16) uint16 {
	var __retVal uint16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint16(C.CallFuncChar16(__func_))
	return __retVal
}

// CallFuncChar16
func CallFuncChar16(func_ FuncChar16) uint16 {
	defer plugify.Scope("cross_call_worker::CallFuncChar16", ModuleName, 3)()
	return P_CallFuncChar16(func_)
}

var P_CallFuncInt8 = func(func_ FuncInt8) int8 {
	var __retVal int8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int8(C.CallFuncInt8(__func_))
	return __retVal
}

// CallFuncInt8
func CallFuncInt8(func_ FuncInt8) int8 {
	defer plugify.Scope("cross_call_worker::CallFuncInt8", ModuleName, 3)()
	return P_CallFuncInt8(func_)
}

var P_CallFuncInt16 = func(func_ FuncInt16) int16 {
	var __retVal int16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int16(C.CallFuncInt16(__func_))
	return __retVal
}

// CallFuncInt16
func CallFuncInt16(func_ FuncInt16) int16 {
	defer plugify.Scope("cross_call_worker::CallFuncInt16", ModuleName, 3)()
	return P_CallFuncInt16(func_)
}

var P_CallFuncInt32 = func(func_ FuncInt32) int32 {
	var __retVal int32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int32(C.CallFuncInt32(__func_))
	return __retVal
}

// CallFuncInt32
func CallFuncInt32(func_ FuncInt32) int32 {
	defer plugify.Scope("cross_call_worker::CallFuncInt32", ModuleName, 3)()
	return P_CallFuncInt32(func_)
}

var P_CallFuncInt64 = func(func_ FuncInt64) int64 {
	var __retVal int64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int64(C.CallFuncInt64(__func_))
	return __retVal
}

// CallFuncInt64
func CallFuncInt64(func_ FuncInt64) int64 {
	defer plugify.Scope("cross_call_worker::CallFuncInt64", ModuleName, 3)()
	return P_CallFuncInt64(func_)
}

var P_CallFuncUInt8 = func(func_ FuncUInt8) uint8 {
	var __retVal uint8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint8(C.CallFuncUInt8(__func_))
	return __retVal
}

// CallFuncUInt8
func CallFuncUInt8(func_ FuncUInt8) uint8 {
	defer plugify.Scope("cross_call_worker::CallFuncUInt8", ModuleName, 3)()
	return P_CallFuncUInt8(func_)
}

var P_CallFuncUInt16 = func(func_ FuncUInt16) uint16 {
	var __retVal uint16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint16(C.CallFuncUInt16(__func_))
	return __retVal
}

// CallFuncUInt16
func CallFuncUInt16(func_ FuncUInt16) uint16 {
	defer plugify.Scope("cross_call_worker::CallFuncUInt16", ModuleName, 3)()
	return P_CallFuncUInt16(func_)
}

var P_CallFuncUInt32 = func(func_ FuncUInt32) uint32 {
	var __retVal uint32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint32(C.CallFuncUInt32(__func_))
	return __retVal
}

// CallFuncUInt32
func CallFuncUInt32(func_ FuncUInt32) uint32 {
	defer plugify.Scope("cross_call_worker::CallFuncUInt32", ModuleName, 3)()
	return P_CallFuncUInt32(func_)
}

var P_CallFuncUInt64 = func(func_ FuncUInt64) uint64 {
	var __retVal uint64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint64(C.CallFuncUInt64(__func_))
	return __retVal
}

// CallFuncUInt64
func CallFuncUInt64(func_ FuncUInt64) uint64 {
	defer plugify.Scope("cross_call_worker::CallFuncUInt64", ModuleName, 3)()
	return P_CallFuncUInt64(func_)
}

var P_CallFuncPtr = func(func_ FuncPtr) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFuncPtr(__func_))
	return __retVal
}

// CallFuncPtr
func CallFuncPtr(func_ FuncPtr) uintptr {
	defer plugify.Scope("cross_call_worker::CallFuncPtr", ModuleName, 3)()
	return P_CallFuncPtr(func_)
}

var P_CallFuncFloat = func(func_ FuncFloat) float32 {
	var __retVal float32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = float32(C.CallFuncFloat(__func_))
	return __retVal
}

// CallFuncFloat
func CallFuncFloat(func_ FuncFloat) float32 {
	defer plugify.Scope("cross_call_worker::CallFuncFloat", ModuleName, 3)()
	return P_CallFuncFloat(func_)
}

var P_CallFuncDouble = func(func_ FuncDouble) float64 {
	var __retVal float64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = float64(C.CallFuncDouble(__func_))
	return __retVal
}

// CallFuncDouble
func CallFuncDouble(func_ FuncDouble) float64 {
	defer plugify.Scope("cross_call_worker::CallFuncDouble", ModuleName, 3)()
	return P_CallFuncDouble(func_)
}

var P_CallFuncString = func(func_ FuncString) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncString(__func_)
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

// CallFuncString
func CallFuncString(func_ FuncString) string {
	defer plugify.Scope("cross_call_worker::CallFuncString", ModuleName, 3)()
	return P_CallFuncString(func_)
}

var P_CallFuncAny = func(func_ FuncAny) any {
	var __retVal any
	var __retVal_native plugify.PlgVariant
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAny(__func_)
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

// CallFuncAny
func CallFuncAny(func_ FuncAny) any {
	defer plugify.Scope("cross_call_worker::CallFuncAny", ModuleName, 3)()
	return P_CallFuncAny(func_)
}

var P_CallFuncFunction = func(func_ FuncFunction) FuncFunctionInner {
	var __retVal FuncFunctionInner
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = plugify.GetDelegateForFunctionPointer(C.CallFuncFunction(__func_), reflect.TypeOf(FuncFunctionInner(nil))).(FuncFunctionInner)
	return __retVal
}

// CallFuncFunction
func CallFuncFunction(func_ FuncFunction) FuncFunctionInner {
	defer plugify.Scope("cross_call_worker::CallFuncFunction", ModuleName, 3)()
	return P_CallFuncFunction(func_)
}

var P_CallFuncBoolVector = func(func_ FuncBoolVector) []bool {
	var __retVal []bool
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncBoolVector(__func_)
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

// CallFuncBoolVector
func CallFuncBoolVector(func_ FuncBoolVector) []bool {
	defer plugify.Scope("cross_call_worker::CallFuncBoolVector", ModuleName, 3)()
	return P_CallFuncBoolVector(func_)
}

var P_CallFuncChar8Vector = func(func_ FuncChar8Vector) []int8 {
	var __retVal []int8
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncChar8Vector(__func_)
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

// CallFuncChar8Vector
func CallFuncChar8Vector(func_ FuncChar8Vector) []int8 {
	defer plugify.Scope("cross_call_worker::CallFuncChar8Vector", ModuleName, 3)()
	return P_CallFuncChar8Vector(func_)
}

var P_CallFuncChar16Vector = func(func_ FuncChar16Vector) []uint16 {
	var __retVal []uint16
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncChar16Vector(__func_)
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

// CallFuncChar16Vector
func CallFuncChar16Vector(func_ FuncChar16Vector) []uint16 {
	defer plugify.Scope("cross_call_worker::CallFuncChar16Vector", ModuleName, 3)()
	return P_CallFuncChar16Vector(func_)
}

var P_CallFuncInt8Vector = func(func_ FuncInt8Vector) []int8 {
	var __retVal []int8
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncInt8Vector(__func_)
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

// CallFuncInt8Vector
func CallFuncInt8Vector(func_ FuncInt8Vector) []int8 {
	defer plugify.Scope("cross_call_worker::CallFuncInt8Vector", ModuleName, 3)()
	return P_CallFuncInt8Vector(func_)
}

var P_CallFuncInt16Vector = func(func_ FuncInt16Vector) []int16 {
	var __retVal []int16
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncInt16Vector(__func_)
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

// CallFuncInt16Vector
func CallFuncInt16Vector(func_ FuncInt16Vector) []int16 {
	defer plugify.Scope("cross_call_worker::CallFuncInt16Vector", ModuleName, 3)()
	return P_CallFuncInt16Vector(func_)
}

var P_CallFuncInt32Vector = func(func_ FuncInt32Vector) []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncInt32Vector(__func_)
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

// CallFuncInt32Vector
func CallFuncInt32Vector(func_ FuncInt32Vector) []int32 {
	defer plugify.Scope("cross_call_worker::CallFuncInt32Vector", ModuleName, 3)()
	return P_CallFuncInt32Vector(func_)
}

var P_CallFuncInt64Vector = func(func_ FuncInt64Vector) []int64 {
	var __retVal []int64
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncInt64Vector(__func_)
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

// CallFuncInt64Vector
func CallFuncInt64Vector(func_ FuncInt64Vector) []int64 {
	defer plugify.Scope("cross_call_worker::CallFuncInt64Vector", ModuleName, 3)()
	return P_CallFuncInt64Vector(func_)
}

var P_CallFuncUInt8Vector = func(func_ FuncUInt8Vector) []uint8 {
	var __retVal []uint8
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncUInt8Vector(__func_)
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

// CallFuncUInt8Vector
func CallFuncUInt8Vector(func_ FuncUInt8Vector) []uint8 {
	defer plugify.Scope("cross_call_worker::CallFuncUInt8Vector", ModuleName, 3)()
	return P_CallFuncUInt8Vector(func_)
}

var P_CallFuncUInt16Vector = func(func_ FuncUInt16Vector) []uint16 {
	var __retVal []uint16
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncUInt16Vector(__func_)
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

// CallFuncUInt16Vector
func CallFuncUInt16Vector(func_ FuncUInt16Vector) []uint16 {
	defer plugify.Scope("cross_call_worker::CallFuncUInt16Vector", ModuleName, 3)()
	return P_CallFuncUInt16Vector(func_)
}

var P_CallFuncUInt32Vector = func(func_ FuncUInt32Vector) []uint32 {
	var __retVal []uint32
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncUInt32Vector(__func_)
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

// CallFuncUInt32Vector
func CallFuncUInt32Vector(func_ FuncUInt32Vector) []uint32 {
	defer plugify.Scope("cross_call_worker::CallFuncUInt32Vector", ModuleName, 3)()
	return P_CallFuncUInt32Vector(func_)
}

var P_CallFuncUInt64Vector = func(func_ FuncUInt64Vector) []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncUInt64Vector(__func_)
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

// CallFuncUInt64Vector
func CallFuncUInt64Vector(func_ FuncUInt64Vector) []uint64 {
	defer plugify.Scope("cross_call_worker::CallFuncUInt64Vector", ModuleName, 3)()
	return P_CallFuncUInt64Vector(func_)
}

var P_CallFuncPtrVector = func(func_ FuncPtrVector) []uintptr {
	var __retVal []uintptr
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncPtrVector(__func_)
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

// CallFuncPtrVector
func CallFuncPtrVector(func_ FuncPtrVector) []uintptr {
	defer plugify.Scope("cross_call_worker::CallFuncPtrVector", ModuleName, 3)()
	return P_CallFuncPtrVector(func_)
}

var P_CallFuncFloatVector = func(func_ FuncFloatVector) []float32 {
	var __retVal []float32
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncFloatVector(__func_)
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

// CallFuncFloatVector
func CallFuncFloatVector(func_ FuncFloatVector) []float32 {
	defer plugify.Scope("cross_call_worker::CallFuncFloatVector", ModuleName, 3)()
	return P_CallFuncFloatVector(func_)
}

var P_CallFuncDoubleVector = func(func_ FuncDoubleVector) []float64 {
	var __retVal []float64
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncDoubleVector(__func_)
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

// CallFuncDoubleVector
func CallFuncDoubleVector(func_ FuncDoubleVector) []float64 {
	defer plugify.Scope("cross_call_worker::CallFuncDoubleVector", ModuleName, 3)()
	return P_CallFuncDoubleVector(func_)
}

var P_CallFuncStringVector = func(func_ FuncStringVector) []string {
	var __retVal []string
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncStringVector(__func_)
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

// CallFuncStringVector
func CallFuncStringVector(func_ FuncStringVector) []string {
	defer plugify.Scope("cross_call_worker::CallFuncStringVector", ModuleName, 3)()
	return P_CallFuncStringVector(func_)
}

var P_CallFuncAnyVector = func(func_ FuncAnyVector) []any {
	var __retVal []any
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAnyVector(__func_)
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

// CallFuncAnyVector
func CallFuncAnyVector(func_ FuncAnyVector) []any {
	defer plugify.Scope("cross_call_worker::CallFuncAnyVector", ModuleName, 3)()
	return P_CallFuncAnyVector(func_)
}

var P_CallFuncVec2Vector = func(func_ FuncVec2Vector) []plugify.Vector2 {
	var __retVal []plugify.Vector2
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncVec2Vector(__func_)
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

// CallFuncVec2Vector
func CallFuncVec2Vector(func_ FuncVec2Vector) []plugify.Vector2 {
	defer plugify.Scope("cross_call_worker::CallFuncVec2Vector", ModuleName, 3)()
	return P_CallFuncVec2Vector(func_)
}

var P_CallFuncVec3Vector = func(func_ FuncVec3Vector) []plugify.Vector3 {
	var __retVal []plugify.Vector3
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncVec3Vector(__func_)
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

// CallFuncVec3Vector
func CallFuncVec3Vector(func_ FuncVec3Vector) []plugify.Vector3 {
	defer plugify.Scope("cross_call_worker::CallFuncVec3Vector", ModuleName, 3)()
	return P_CallFuncVec3Vector(func_)
}

var P_CallFuncVec4Vector = func(func_ FuncVec4Vector) []plugify.Vector4 {
	var __retVal []plugify.Vector4
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncVec4Vector(__func_)
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

// CallFuncVec4Vector
func CallFuncVec4Vector(func_ FuncVec4Vector) []plugify.Vector4 {
	defer plugify.Scope("cross_call_worker::CallFuncVec4Vector", ModuleName, 3)()
	return P_CallFuncVec4Vector(func_)
}

var P_CallFuncMat4x4Vector = func(func_ FuncMat4x4Vector) []plugify.Matrix4x4 {
	var __retVal []plugify.Matrix4x4
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncMat4x4Vector(__func_)
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

// CallFuncMat4x4Vector
func CallFuncMat4x4Vector(func_ FuncMat4x4Vector) []plugify.Matrix4x4 {
	defer plugify.Scope("cross_call_worker::CallFuncMat4x4Vector", ModuleName, 3)()
	return P_CallFuncMat4x4Vector(func_)
}

var P_CallFuncVec2 = func(func_ FuncVec2) plugify.Vector2 {
	var __retVal plugify.Vector2
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncVec2(__func_)
	__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncVec2
func CallFuncVec2(func_ FuncVec2) plugify.Vector2 {
	defer plugify.Scope("cross_call_worker::CallFuncVec2", ModuleName, 3)()
	return P_CallFuncVec2(func_)
}

var P_CallFuncVec3 = func(func_ FuncVec3) plugify.Vector3 {
	var __retVal plugify.Vector3
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncVec3(__func_)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncVec3
func CallFuncVec3(func_ FuncVec3) plugify.Vector3 {
	defer plugify.Scope("cross_call_worker::CallFuncVec3", ModuleName, 3)()
	return P_CallFuncVec3(func_)
}

var P_CallFuncVec4 = func(func_ FuncVec4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncVec4(__func_)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncVec4
func CallFuncVec4(func_ FuncVec4) plugify.Vector4 {
	defer plugify.Scope("cross_call_worker::CallFuncVec4", ModuleName, 3)()
	return P_CallFuncVec4(func_)
}

var P_CallFuncMat4x4 = func(func_ FuncMat4x4) plugify.Matrix4x4 {
	var __retVal plugify.Matrix4x4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncMat4x4(__func_)
	__retVal = *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncMat4x4
func CallFuncMat4x4(func_ FuncMat4x4) plugify.Matrix4x4 {
	defer plugify.Scope("cross_call_worker::CallFuncMat4x4", ModuleName, 3)()
	return P_CallFuncMat4x4(func_)
}

var P_CallFuncAliasBool = func(func_ FuncAliasBool) AliasBool {
	var __retVal AliasBool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasBool(C.CallFuncAliasBool(__func_))
	return __retVal
}

// CallFuncAliasBool
func CallFuncAliasBool(func_ FuncAliasBool) AliasBool {
	defer plugify.Scope("cross_call_worker::CallFuncAliasBool", ModuleName, 3)()
	return P_CallFuncAliasBool(func_)
}

var P_CallFuncAliasChar8 = func(func_ FuncAliasChar8) AliasChar8 {
	var __retVal AliasChar8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasChar8(C.CallFuncAliasChar8(__func_))
	return __retVal
}

// CallFuncAliasChar8
func CallFuncAliasChar8(func_ FuncAliasChar8) AliasChar8 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasChar8", ModuleName, 3)()
	return P_CallFuncAliasChar8(func_)
}

var P_CallFuncAliasChar16 = func(func_ FuncAliasChar16) AliasChar16 {
	var __retVal AliasChar16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasChar16(C.CallFuncAliasChar16(__func_))
	return __retVal
}

// CallFuncAliasChar16
func CallFuncAliasChar16(func_ FuncAliasChar16) AliasChar16 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasChar16", ModuleName, 3)()
	return P_CallFuncAliasChar16(func_)
}

var P_CallFuncAliasInt8 = func(func_ FuncAliasInt8) AliasInt8 {
	var __retVal AliasInt8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasInt8(C.CallFuncAliasInt8(__func_))
	return __retVal
}

// CallFuncAliasInt8
func CallFuncAliasInt8(func_ FuncAliasInt8) AliasInt8 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasInt8", ModuleName, 3)()
	return P_CallFuncAliasInt8(func_)
}

var P_CallFuncAliasInt16 = func(func_ FuncAliasInt16) AliasInt16 {
	var __retVal AliasInt16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasInt16(C.CallFuncAliasInt16(__func_))
	return __retVal
}

// CallFuncAliasInt16
func CallFuncAliasInt16(func_ FuncAliasInt16) AliasInt16 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasInt16", ModuleName, 3)()
	return P_CallFuncAliasInt16(func_)
}

var P_CallFuncAliasInt32 = func(func_ FuncAliasInt32) AliasInt32 {
	var __retVal AliasInt32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasInt32(C.CallFuncAliasInt32(__func_))
	return __retVal
}

// CallFuncAliasInt32
func CallFuncAliasInt32(func_ FuncAliasInt32) AliasInt32 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasInt32", ModuleName, 3)()
	return P_CallFuncAliasInt32(func_)
}

var P_CallFuncAliasInt64 = func(func_ FuncAliasInt64) AliasInt64 {
	var __retVal AliasInt64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasInt64(C.CallFuncAliasInt64(__func_))
	return __retVal
}

// CallFuncAliasInt64
func CallFuncAliasInt64(func_ FuncAliasInt64) AliasInt64 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasInt64", ModuleName, 3)()
	return P_CallFuncAliasInt64(func_)
}

var P_CallFuncAliasUInt8 = func(func_ FuncAliasUInt8) AliasUInt8 {
	var __retVal AliasUInt8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasUInt8(C.CallFuncAliasUInt8(__func_))
	return __retVal
}

// CallFuncAliasUInt8
func CallFuncAliasUInt8(func_ FuncAliasUInt8) AliasUInt8 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasUInt8", ModuleName, 3)()
	return P_CallFuncAliasUInt8(func_)
}

var P_CallFuncAliasUInt16 = func(func_ FuncAliasUInt16) AliasUInt16 {
	var __retVal AliasUInt16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasUInt16(C.CallFuncAliasUInt16(__func_))
	return __retVal
}

// CallFuncAliasUInt16
func CallFuncAliasUInt16(func_ FuncAliasUInt16) AliasUInt16 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasUInt16", ModuleName, 3)()
	return P_CallFuncAliasUInt16(func_)
}

var P_CallFuncAliasUInt32 = func(func_ FuncAliasUInt32) AliasUInt32 {
	var __retVal AliasUInt32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasUInt32(C.CallFuncAliasUInt32(__func_))
	return __retVal
}

// CallFuncAliasUInt32
func CallFuncAliasUInt32(func_ FuncAliasUInt32) AliasUInt32 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasUInt32", ModuleName, 3)()
	return P_CallFuncAliasUInt32(func_)
}

var P_CallFuncAliasUInt64 = func(func_ FuncAliasUInt64) AliasUInt64 {
	var __retVal AliasUInt64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasUInt64(C.CallFuncAliasUInt64(__func_))
	return __retVal
}

// CallFuncAliasUInt64
func CallFuncAliasUInt64(func_ FuncAliasUInt64) AliasUInt64 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasUInt64", ModuleName, 3)()
	return P_CallFuncAliasUInt64(func_)
}

var P_CallFuncAliasPtr = func(func_ FuncAliasPtr) AliasPtr {
	var __retVal AliasPtr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasPtr(C.CallFuncAliasPtr(__func_))
	return __retVal
}

// CallFuncAliasPtr
func CallFuncAliasPtr(func_ FuncAliasPtr) AliasPtr {
	defer plugify.Scope("cross_call_worker::CallFuncAliasPtr", ModuleName, 3)()
	return P_CallFuncAliasPtr(func_)
}

var P_CallFuncAliasFloat = func(func_ FuncAliasFloat) AliasFloat {
	var __retVal AliasFloat
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasFloat(C.CallFuncAliasFloat(__func_))
	return __retVal
}

// CallFuncAliasFloat
func CallFuncAliasFloat(func_ FuncAliasFloat) AliasFloat {
	defer plugify.Scope("cross_call_worker::CallFuncAliasFloat", ModuleName, 3)()
	return P_CallFuncAliasFloat(func_)
}

var P_CallFuncAliasDouble = func(func_ FuncAliasDouble) AliasDouble {
	var __retVal AliasDouble
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = AliasDouble(C.CallFuncAliasDouble(__func_))
	return __retVal
}

// CallFuncAliasDouble
func CallFuncAliasDouble(func_ FuncAliasDouble) AliasDouble {
	defer plugify.Scope("cross_call_worker::CallFuncAliasDouble", ModuleName, 3)()
	return P_CallFuncAliasDouble(func_)
}

var P_CallFuncAliasString = func(func_ FuncAliasString) AliasString {
	var __retVal AliasString
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasString(__func_)
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

// CallFuncAliasString
func CallFuncAliasString(func_ FuncAliasString) AliasString {
	defer plugify.Scope("cross_call_worker::CallFuncAliasString", ModuleName, 3)()
	return P_CallFuncAliasString(func_)
}

var P_CallFuncAliasAny = func(func_ FuncAliasAny) AliasAny {
	var __retVal AliasAny
	var __retVal_native plugify.PlgVariant
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasAny(__func_)
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

// CallFuncAliasAny
func CallFuncAliasAny(func_ FuncAliasAny) AliasAny {
	defer plugify.Scope("cross_call_worker::CallFuncAliasAny", ModuleName, 3)()
	return P_CallFuncAliasAny(func_)
}

var P_CallFuncAliasFunction = func(func_ FuncAliasFunction) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFuncAliasFunction(__func_))
	return __retVal
}

// CallFuncAliasFunction
func CallFuncAliasFunction(func_ FuncAliasFunction) uintptr {
	defer plugify.Scope("cross_call_worker::CallFuncAliasFunction", ModuleName, 3)()
	return P_CallFuncAliasFunction(func_)
}

var P_CallFuncAliasBoolVector = func(func_ FuncAliasBoolVector) AliasBoolVector {
	var __retVal AliasBoolVector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasBoolVector(__func_)
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

// CallFuncAliasBoolVector
func CallFuncAliasBoolVector(func_ FuncAliasBoolVector) AliasBoolVector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasBoolVector", ModuleName, 3)()
	return P_CallFuncAliasBoolVector(func_)
}

var P_CallFuncAliasChar8Vector = func(func_ FuncAliasChar8Vector) AliasChar8Vector {
	var __retVal AliasChar8Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasChar8Vector(__func_)
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

// CallFuncAliasChar8Vector
func CallFuncAliasChar8Vector(func_ FuncAliasChar8Vector) AliasChar8Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasChar8Vector", ModuleName, 3)()
	return P_CallFuncAliasChar8Vector(func_)
}

var P_CallFuncAliasChar16Vector = func(func_ FuncAliasChar16Vector) AliasChar16Vector {
	var __retVal AliasChar16Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasChar16Vector(__func_)
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

// CallFuncAliasChar16Vector
func CallFuncAliasChar16Vector(func_ FuncAliasChar16Vector) AliasChar16Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasChar16Vector", ModuleName, 3)()
	return P_CallFuncAliasChar16Vector(func_)
}

var P_CallFuncAliasInt8Vector = func(func_ FuncAliasInt8Vector) AliasInt8Vector {
	var __retVal AliasInt8Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasInt8Vector(__func_)
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

// CallFuncAliasInt8Vector
func CallFuncAliasInt8Vector(func_ FuncAliasInt8Vector) AliasInt8Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasInt8Vector", ModuleName, 3)()
	return P_CallFuncAliasInt8Vector(func_)
}

var P_CallFuncAliasInt16Vector = func(func_ FuncAliasInt16Vector) AliasInt16Vector {
	var __retVal AliasInt16Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasInt16Vector(__func_)
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

// CallFuncAliasInt16Vector
func CallFuncAliasInt16Vector(func_ FuncAliasInt16Vector) AliasInt16Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasInt16Vector", ModuleName, 3)()
	return P_CallFuncAliasInt16Vector(func_)
}

var P_CallFuncAliasInt32Vector = func(func_ FuncAliasInt32Vector) AliasInt32Vector {
	var __retVal AliasInt32Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasInt32Vector(__func_)
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

// CallFuncAliasInt32Vector
func CallFuncAliasInt32Vector(func_ FuncAliasInt32Vector) AliasInt32Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasInt32Vector", ModuleName, 3)()
	return P_CallFuncAliasInt32Vector(func_)
}

var P_CallFuncAliasInt64Vector = func(func_ FuncAliasInt64Vector) AliasInt64Vector {
	var __retVal AliasInt64Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasInt64Vector(__func_)
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

// CallFuncAliasInt64Vector
func CallFuncAliasInt64Vector(func_ FuncAliasInt64Vector) AliasInt64Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasInt64Vector", ModuleName, 3)()
	return P_CallFuncAliasInt64Vector(func_)
}

var P_CallFuncAliasUInt8Vector = func(func_ FuncAliasUInt8Vector) AliasUInt8Vector {
	var __retVal AliasUInt8Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasUInt8Vector(__func_)
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

// CallFuncAliasUInt8Vector
func CallFuncAliasUInt8Vector(func_ FuncAliasUInt8Vector) AliasUInt8Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasUInt8Vector", ModuleName, 3)()
	return P_CallFuncAliasUInt8Vector(func_)
}

var P_CallFuncAliasUInt16Vector = func(func_ FuncAliasUInt16Vector) AliasUInt16Vector {
	var __retVal AliasUInt16Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasUInt16Vector(__func_)
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

// CallFuncAliasUInt16Vector
func CallFuncAliasUInt16Vector(func_ FuncAliasUInt16Vector) AliasUInt16Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasUInt16Vector", ModuleName, 3)()
	return P_CallFuncAliasUInt16Vector(func_)
}

var P_CallFuncAliasUInt32Vector = func(func_ FuncAliasUInt32Vector) AliasUInt32Vector {
	var __retVal AliasUInt32Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasUInt32Vector(__func_)
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

// CallFuncAliasUInt32Vector
func CallFuncAliasUInt32Vector(func_ FuncAliasUInt32Vector) AliasUInt32Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasUInt32Vector", ModuleName, 3)()
	return P_CallFuncAliasUInt32Vector(func_)
}

var P_CallFuncAliasUInt64Vector = func(func_ FuncAliasUInt64Vector) AliasUInt64Vector {
	var __retVal AliasUInt64Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasUInt64Vector(__func_)
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

// CallFuncAliasUInt64Vector
func CallFuncAliasUInt64Vector(func_ FuncAliasUInt64Vector) AliasUInt64Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasUInt64Vector", ModuleName, 3)()
	return P_CallFuncAliasUInt64Vector(func_)
}

var P_CallFuncAliasPtrVector = func(func_ FuncAliasPtrVector) AliasPtrVector {
	var __retVal AliasPtrVector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasPtrVector(__func_)
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

// CallFuncAliasPtrVector
func CallFuncAliasPtrVector(func_ FuncAliasPtrVector) AliasPtrVector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasPtrVector", ModuleName, 3)()
	return P_CallFuncAliasPtrVector(func_)
}

var P_CallFuncAliasFloatVector = func(func_ FuncAliasFloatVector) AliasFloatVector {
	var __retVal AliasFloatVector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasFloatVector(__func_)
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

// CallFuncAliasFloatVector
func CallFuncAliasFloatVector(func_ FuncAliasFloatVector) AliasFloatVector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasFloatVector", ModuleName, 3)()
	return P_CallFuncAliasFloatVector(func_)
}

var P_CallFuncAliasDoubleVector = func(func_ FuncAliasDoubleVector) AliasDoubleVector {
	var __retVal AliasDoubleVector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasDoubleVector(__func_)
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

// CallFuncAliasDoubleVector
func CallFuncAliasDoubleVector(func_ FuncAliasDoubleVector) AliasDoubleVector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasDoubleVector", ModuleName, 3)()
	return P_CallFuncAliasDoubleVector(func_)
}

var P_CallFuncAliasStringVector = func(func_ FuncAliasStringVector) AliasStringVector {
	var __retVal AliasStringVector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasStringVector(__func_)
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

// CallFuncAliasStringVector
func CallFuncAliasStringVector(func_ FuncAliasStringVector) AliasStringVector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasStringVector", ModuleName, 3)()
	return P_CallFuncAliasStringVector(func_)
}

var P_CallFuncAliasAnyVector = func(func_ FuncAliasAnyVector) AliasAnyVector {
	var __retVal AliasAnyVector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasAnyVector(__func_)
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

// CallFuncAliasAnyVector
func CallFuncAliasAnyVector(func_ FuncAliasAnyVector) AliasAnyVector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasAnyVector", ModuleName, 3)()
	return P_CallFuncAliasAnyVector(func_)
}

var P_CallFuncAliasVec2Vector = func(func_ FuncAliasVec2Vector) AliasVec2Vector {
	var __retVal AliasVec2Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasVec2Vector(__func_)
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

// CallFuncAliasVec2Vector
func CallFuncAliasVec2Vector(func_ FuncAliasVec2Vector) AliasVec2Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasVec2Vector", ModuleName, 3)()
	return P_CallFuncAliasVec2Vector(func_)
}

var P_CallFuncAliasVec3Vector = func(func_ FuncAliasVec3Vector) AliasVec3Vector {
	var __retVal AliasVec3Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasVec3Vector(__func_)
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

// CallFuncAliasVec3Vector
func CallFuncAliasVec3Vector(func_ FuncAliasVec3Vector) AliasVec3Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasVec3Vector", ModuleName, 3)()
	return P_CallFuncAliasVec3Vector(func_)
}

var P_CallFuncAliasVec4Vector = func(func_ FuncAliasVec4Vector) AliasVec4Vector {
	var __retVal AliasVec4Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasVec4Vector(__func_)
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

// CallFuncAliasVec4Vector
func CallFuncAliasVec4Vector(func_ FuncAliasVec4Vector) AliasVec4Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasVec4Vector", ModuleName, 3)()
	return P_CallFuncAliasVec4Vector(func_)
}

var P_CallFuncAliasMat4x4Vector = func(func_ FuncAliasMat4x4Vector) AliasMat4x4Vector {
	var __retVal AliasMat4x4Vector
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasMat4x4Vector(__func_)
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

// CallFuncAliasMat4x4Vector
func CallFuncAliasMat4x4Vector(func_ FuncAliasMat4x4Vector) AliasMat4x4Vector {
	defer plugify.Scope("cross_call_worker::CallFuncAliasMat4x4Vector", ModuleName, 3)()
	return P_CallFuncAliasMat4x4Vector(func_)
}

var P_CallFuncAliasVec2 = func(func_ FuncAliasVec2) AliasVec2 {
	var __retVal AliasVec2
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncAliasVec2(__func_)
	__retVal = *(*AliasVec2)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncAliasVec2
func CallFuncAliasVec2(func_ FuncAliasVec2) AliasVec2 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasVec2", ModuleName, 3)()
	return P_CallFuncAliasVec2(func_)
}

var P_CallFuncAliasVec3 = func(func_ FuncAliasVec3) AliasVec3 {
	var __retVal AliasVec3
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncAliasVec3(__func_)
	__retVal = *(*AliasVec3)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncAliasVec3
func CallFuncAliasVec3(func_ FuncAliasVec3) AliasVec3 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasVec3", ModuleName, 3)()
	return P_CallFuncAliasVec3(func_)
}

var P_CallFuncAliasVec4 = func(func_ FuncAliasVec4) AliasVec4 {
	var __retVal AliasVec4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncAliasVec4(__func_)
	__retVal = *(*AliasVec4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncAliasVec4
func CallFuncAliasVec4(func_ FuncAliasVec4) AliasVec4 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasVec4", ModuleName, 3)()
	return P_CallFuncAliasVec4(func_)
}

var P_CallFuncAliasMat4x4 = func(func_ FuncAliasMat4x4) AliasMat4x4 {
	var __retVal AliasMat4x4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFuncAliasMat4x4(__func_)
	__retVal = *(*AliasMat4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFuncAliasMat4x4
func CallFuncAliasMat4x4(func_ FuncAliasMat4x4) AliasMat4x4 {
	defer plugify.Scope("cross_call_worker::CallFuncAliasMat4x4", ModuleName, 3)()
	return P_CallFuncAliasMat4x4(func_)
}

var P_CallFuncAliasAll = func(func_ FuncAliasAll) AliasString {
	var __retVal AliasString
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncAliasAll(__func_)
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

// CallFuncAliasAll
func CallFuncAliasAll(func_ FuncAliasAll) AliasString {
	defer plugify.Scope("cross_call_worker::CallFuncAliasAll", ModuleName, 3)()
	return P_CallFuncAliasAll(func_)
}

var P_CallFunc1 = func(func_ Func1) int32 {
	var __retVal int32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int32(C.CallFunc1(__func_))
	return __retVal
}

// CallFunc1
func CallFunc1(func_ Func1) int32 {
	defer plugify.Scope("cross_call_worker::CallFunc1", ModuleName, 3)()
	return P_CallFunc1(func_)
}

var P_CallFunc2 = func(func_ Func2) int8 {
	var __retVal int8
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int8(C.CallFunc2(__func_))
	return __retVal
}

// CallFunc2
func CallFunc2(func_ Func2) int8 {
	defer plugify.Scope("cross_call_worker::CallFunc2", ModuleName, 3)()
	return P_CallFunc2(func_)
}

var P_CallFunc3 = func(func_ Func3) {
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	C.CallFunc3(__func_)
}

// CallFunc3
func CallFunc3(func_ Func3) {
	defer plugify.Scope("cross_call_worker::CallFunc3", ModuleName, 3)()
	P_CallFunc3(func_)
}

var P_CallFunc4 = func(func_ Func4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFunc4(__func_)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFunc4
func CallFunc4(func_ Func4) plugify.Vector4 {
	defer plugify.Scope("cross_call_worker::CallFunc4", ModuleName, 3)()
	return P_CallFunc4(func_)
}

var P_CallFunc5 = func(func_ Func5) bool {
	var __retVal bool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = bool(C.CallFunc5(__func_))
	return __retVal
}

// CallFunc5
func CallFunc5(func_ Func5) bool {
	defer plugify.Scope("cross_call_worker::CallFunc5", ModuleName, 3)()
	return P_CallFunc5(func_)
}

var P_CallFunc6 = func(func_ Func6) int64 {
	var __retVal int64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int64(C.CallFunc6(__func_))
	return __retVal
}

// CallFunc6
func CallFunc6(func_ Func6) int64 {
	defer plugify.Scope("cross_call_worker::CallFunc6", ModuleName, 3)()
	return P_CallFunc6(func_)
}

var P_CallFunc7 = func(func_ Func7) float64 {
	var __retVal float64
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = float64(C.CallFunc7(__func_))
	return __retVal
}

// CallFunc7
func CallFunc7(func_ Func7) float64 {
	defer plugify.Scope("cross_call_worker::CallFunc7", ModuleName, 3)()
	return P_CallFunc7(func_)
}

var P_CallFunc8 = func(func_ Func8) plugify.Matrix4x4 {
	var __retVal plugify.Matrix4x4
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__native := C.CallFunc8(__func_)
	__retVal = *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// CallFunc8
func CallFunc8(func_ Func8) plugify.Matrix4x4 {
	defer plugify.Scope("cross_call_worker::CallFunc8", ModuleName, 3)()
	return P_CallFunc8(func_)
}

var P_CallFunc9 = func(func_ Func9) {
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	C.CallFunc9(__func_)
}

// CallFunc9
func CallFunc9(func_ Func9) {
	defer plugify.Scope("cross_call_worker::CallFunc9", ModuleName, 3)()
	P_CallFunc9(func_)
}

var P_CallFunc10 = func(func_ Func10) uint32 {
	var __retVal uint32
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uint32(C.CallFunc10(__func_))
	return __retVal
}

// CallFunc10
func CallFunc10(func_ Func10) uint32 {
	defer plugify.Scope("cross_call_worker::CallFunc10", ModuleName, 3)()
	return P_CallFunc10(func_)
}

var P_CallFunc11 = func(func_ Func11) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFunc11(__func_))
	return __retVal
}

// CallFunc11
func CallFunc11(func_ Func11) uintptr {
	defer plugify.Scope("cross_call_worker::CallFunc11", ModuleName, 3)()
	return P_CallFunc11(func_)
}

var P_CallFunc12 = func(func_ Func12) bool {
	var __retVal bool
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = bool(C.CallFunc12(__func_))
	return __retVal
}

// CallFunc12
func CallFunc12(func_ Func12) bool {
	defer plugify.Scope("cross_call_worker::CallFunc12", ModuleName, 3)()
	return P_CallFunc12(func_)
}

var P_CallFunc13 = func(func_ Func13) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc13(__func_)
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

// CallFunc13
func CallFunc13(func_ Func13) string {
	defer plugify.Scope("cross_call_worker::CallFunc13", ModuleName, 3)()
	return P_CallFunc13(func_)
}

var P_CallFunc14 = func(func_ Func14) []string {
	var __retVal []string
	var __retVal_native plugify.PlgVector
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc14(__func_)
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

// CallFunc14
func CallFunc14(func_ Func14) []string {
	defer plugify.Scope("cross_call_worker::CallFunc14", ModuleName, 3)()
	return P_CallFunc14(func_)
}

var P_CallFunc15 = func(func_ Func15) int16 {
	var __retVal int16
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = int16(C.CallFunc15(__func_))
	return __retVal
}

// CallFunc15
func CallFunc15(func_ Func15) int16 {
	defer plugify.Scope("cross_call_worker::CallFunc15", ModuleName, 3)()
	return P_CallFunc15(func_)
}

var P_CallFunc16 = func(func_ Func16) uintptr {
	var __retVal uintptr
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	__retVal = uintptr(C.CallFunc16(__func_))
	return __retVal
}

// CallFunc16
func CallFunc16(func_ Func16) uintptr {
	defer plugify.Scope("cross_call_worker::CallFunc16", ModuleName, 3)()
	return P_CallFunc16(func_)
}

var P_CallFunc17 = func(func_ Func17) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc17(__func_)
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

// CallFunc17
func CallFunc17(func_ Func17) string {
	defer plugify.Scope("cross_call_worker::CallFunc17", ModuleName, 3)()
	return P_CallFunc17(func_)
}

var P_CallFunc18 = func(func_ Func18) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc18(__func_)
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

// CallFunc18
func CallFunc18(func_ Func18) string {
	defer plugify.Scope("cross_call_worker::CallFunc18", ModuleName, 3)()
	return P_CallFunc18(func_)
}

var P_CallFunc19 = func(func_ Func19) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc19(__func_)
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

// CallFunc19
func CallFunc19(func_ Func19) string {
	defer plugify.Scope("cross_call_worker::CallFunc19", ModuleName, 3)()
	return P_CallFunc19(func_)
}

var P_CallFunc20 = func(func_ Func20) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc20(__func_)
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

// CallFunc20
func CallFunc20(func_ Func20) string {
	defer plugify.Scope("cross_call_worker::CallFunc20", ModuleName, 3)()
	return P_CallFunc20(func_)
}

var P_CallFunc21 = func(func_ Func21) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc21(__func_)
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

// CallFunc21
func CallFunc21(func_ Func21) string {
	defer plugify.Scope("cross_call_worker::CallFunc21", ModuleName, 3)()
	return P_CallFunc21(func_)
}

var P_CallFunc22 = func(func_ Func22) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc22(__func_)
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

// CallFunc22
func CallFunc22(func_ Func22) string {
	defer plugify.Scope("cross_call_worker::CallFunc22", ModuleName, 3)()
	return P_CallFunc22(func_)
}

var P_CallFunc23 = func(func_ Func23) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc23(__func_)
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

// CallFunc23
func CallFunc23(func_ Func23) string {
	defer plugify.Scope("cross_call_worker::CallFunc23", ModuleName, 3)()
	return P_CallFunc23(func_)
}

var P_CallFunc24 = func(func_ Func24) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc24(__func_)
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

// CallFunc24
func CallFunc24(func_ Func24) string {
	defer plugify.Scope("cross_call_worker::CallFunc24", ModuleName, 3)()
	return P_CallFunc24(func_)
}

var P_CallFunc25 = func(func_ Func25) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc25(__func_)
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

// CallFunc25
func CallFunc25(func_ Func25) string {
	defer plugify.Scope("cross_call_worker::CallFunc25", ModuleName, 3)()
	return P_CallFunc25(func_)
}

var P_CallFunc26 = func(func_ Func26) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc26(__func_)
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

// CallFunc26
func CallFunc26(func_ Func26) string {
	defer plugify.Scope("cross_call_worker::CallFunc26", ModuleName, 3)()
	return P_CallFunc26(func_)
}

var P_CallFunc27 = func(func_ Func27) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc27(__func_)
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

// CallFunc27
func CallFunc27(func_ Func27) string {
	defer plugify.Scope("cross_call_worker::CallFunc27", ModuleName, 3)()
	return P_CallFunc27(func_)
}

var P_CallFunc28 = func(func_ Func28) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc28(__func_)
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

// CallFunc28
func CallFunc28(func_ Func28) string {
	defer plugify.Scope("cross_call_worker::CallFunc28", ModuleName, 3)()
	return P_CallFunc28(func_)
}

var P_CallFunc29 = func(func_ Func29) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc29(__func_)
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

// CallFunc29
func CallFunc29(func_ Func29) string {
	defer plugify.Scope("cross_call_worker::CallFunc29", ModuleName, 3)()
	return P_CallFunc29(func_)
}

var P_CallFunc30 = func(func_ Func30) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc30(__func_)
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

// CallFunc30
func CallFunc30(func_ Func30) string {
	defer plugify.Scope("cross_call_worker::CallFunc30", ModuleName, 3)()
	return P_CallFunc30(func_)
}

var P_CallFunc31 = func(func_ Func31) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc31(__func_)
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

// CallFunc31
func CallFunc31(func_ Func31) string {
	defer plugify.Scope("cross_call_worker::CallFunc31", ModuleName, 3)()
	return P_CallFunc31(func_)
}

var P_CallFunc32 = func(func_ Func32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc32(__func_)
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

// CallFunc32
func CallFunc32(func_ Func32) string {
	defer plugify.Scope("cross_call_worker::CallFunc32", ModuleName, 3)()
	return P_CallFunc32(func_)
}

var P_CallFunc33 = func(func_ Func33) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFunc33(__func_)
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

// CallFunc33
func CallFunc33(func_ Func33) string {
	defer plugify.Scope("cross_call_worker::CallFunc33", ModuleName, 3)()
	return P_CallFunc33(func_)
}

var P_CallFuncEnum = func(func_ FuncEnum) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	fmt.Println("P_CallFuncEnum Callling with marshal")
	__func_ := plugify.GetFunctionPointerForDelegate(func_)
	plugify.Block{
		Try: func() {
			__native := C.CallFuncEnum(__func_)
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

// CallFuncEnum
func CallFuncEnum(func_ FuncEnum) string {
	defer plugify.Scope("cross_call_worker::CallFuncEnum", ModuleName, 3)()
	return P_CallFuncEnum(func_)
}

var P_ReverseCall = func(test string) {
	__test := plugify.ConstructString(test)
	plugify.Block{
		Try: func() {
			C.ReverseCall((*C.String)(unsafe.Pointer(&__test)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__test)
		},
	}.Do()
}

// ReverseCall
func ReverseCall(test string) {
	defer plugify.Scope("cross_call_worker::ReverseCall", ModuleName, 3)()
	P_ReverseCall(test)
}
