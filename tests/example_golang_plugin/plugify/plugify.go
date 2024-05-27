package plugify

/*
#cgo LDFLAGS: -L${SRCDIR}/libplugify.a
#include <plugify.h>
*/
import "C"
import "unsafe"

type PluginStartCallback func()
type PluginEndCallback func()

type Plugify struct {
	fnPluginStartCallback PluginStartCallback
	fnPluginEndCallback   PluginEndCallback
}

var plugify Plugify = Plugify {
	fnPluginStartCallback: func() {},
	fnPluginEndCallback:   func() {},
}

const kApiVersion = 1

//export Plugify_Init
func Plugify_Init(api []uintptr, version int32) int32 {
	if version < kApiVersion {
		return kApiVersion
	}
	C.Plugify_SetGetMethodPtr(unsafe.Pointer(api[0]))
    C.Plugify_SetCreateStringEmpty(unsafe.Pointer(api[1]))
    C.Plugify_SetCreateString(unsafe.Pointer(api[2]))
    C.Plugify_SetGetString(unsafe.Pointer(api[3]))
    C.Plugify_SetDeleteString(unsafe.Pointer(api[4]))
    C.Plugify_SetCreateVectorBool(unsafe.Pointer(api[5]))
    C.Plugify_SetCreateVectorChar8(unsafe.Pointer(api[6]))
    C.Plugify_SetCreateVectorChar16(unsafe.Pointer(api[7]))
    C.Plugify_SetCreateVectorInt8(unsafe.Pointer(api[8]))
    C.Plugify_SetCreateVectorInt16(unsafe.Pointer(api[9]))
    C.Plugify_SetCreateVectorInt32(unsafe.Pointer(api[10]))
    C.Plugify_SetCreateVectorInt64(unsafe.Pointer(api[11]))
    C.Plugify_SetCreateVectorUInt8(unsafe.Pointer(api[12]))
    C.Plugify_SetCreateVectorUInt16(unsafe.Pointer(api[13]))
    C.Plugify_SetCreateVectorUInt32(unsafe.Pointer(api[14]))
    C.Plugify_SetCreateVectorUInt64(unsafe.Pointer(api[15]))
    C.Plugify_SetCreateVectorUIntPtr(unsafe.Pointer(api[16]))
    C.Plugify_SetCreateVectorFloat(unsafe.Pointer(api[17]))
    C.Plugify_SetCreateVectorDouble(unsafe.Pointer(api[18]))
    C.Plugify_SetCreateVectorString(unsafe.Pointer(api[19]))
    C.Plugify_SetCreateVectorBoolE(unsafe.Pointer(api[20]))
    C.Plugify_SetCreateVectorChar8E(unsafe.Pointer(api[21]))
    C.Plugify_SetCreateVectorChar16E(unsafe.Pointer(api[22]))
    C.Plugify_SetCreateVectorInt8E(unsafe.Pointer(api[23]))
    C.Plugify_SetCreateVectorInt16E(unsafe.Pointer(api[24]))
    C.Plugify_SetCreateVectorInt32E(unsafe.Pointer(api[25]))
    C.Plugify_SetCreateVectorInt64E(unsafe.Pointer(api[26]))
    C.Plugify_SetCreateVectorUInt8E(unsafe.Pointer(api[27]))
    C.Plugify_SetCreateVectorUInt16E(unsafe.Pointer(api[28]))
    C.Plugify_SetCreateVectorUInt32E(unsafe.Pointer(api[29]))
    C.Plugify_SetCreateVectorUInt64E(unsafe.Pointer(api[30]))
    C.Plugify_SetCreateVectorUIntPtrE(unsafe.Pointer(api[31]))
    C.Plugify_SetCreateVectorFloatE(unsafe.Pointer(api[32]))
    C.Plugify_SetCreateVectorDoubleE(unsafe.Pointer(api[33]))
    C.Plugify_SetCreateVectorStringE(unsafe.Pointer(api[34]))
    C.Plugify_SetDeleteVectorBool(unsafe.Pointer(api[35]))
    C.Plugify_SetDeleteVectorChar8(unsafe.Pointer(api[36]))
    C.Plugify_SetDeleteVectorChar16(unsafe.Pointer(api[37]))
    C.Plugify_SetDeleteVectorInt8(unsafe.Pointer(api[38]))
    C.Plugify_SetDeleteVectorInt16(unsafe.Pointer(api[39]))
    C.Plugify_SetDeleteVectorInt32(unsafe.Pointer(api[40]))
    C.Plugify_SetDeleteVectorInt64(unsafe.Pointer(api[41]))
    C.Plugify_SetDeleteVectorUInt8(unsafe.Pointer(api[42]))
    C.Plugify_SetDeleteVectorUInt16(unsafe.Pointer(api[43]))
    C.Plugify_SetDeleteVectorUInt32(unsafe.Pointer(api[44]))
    C.Plugify_SetDeleteVectorUInt64(unsafe.Pointer(api[45]))
    C.Plugify_SetDeleteVectorUIntPtr(unsafe.Pointer(api[46]))
    C.Plugify_SetDeleteVectorFloat(unsafe.Pointer(api[47]))
    C.Plugify_SetDeleteVectorDouble(unsafe.Pointer(api[48]))
    C.Plugify_SetDeleteVectorString(unsafe.Pointer(api[49]))
    C.Plugify_SetGetVectorSizeBool(unsafe.Pointer(api[50]))
    C.Plugify_SetGetVectorSizeChar8(unsafe.Pointer(api[51]))
    C.Plugify_SetGetVectorSizeChar16(unsafe.Pointer(api[52]))
    C.Plugify_SetGetVectorSizeInt8(unsafe.Pointer(api[53]))
    C.Plugify_SetGetVectorSizeInt16(unsafe.Pointer(api[54]))
    C.Plugify_SetGetVectorSizeInt32(unsafe.Pointer(api[55]))
    C.Plugify_SetGetVectorSizeInt64(unsafe.Pointer(api[56]))
    C.Plugify_SetGetVectorSizeUInt8(unsafe.Pointer(api[57]))
    C.Plugify_SetGetVectorSizeUInt16(unsafe.Pointer(api[58]))
    C.Plugify_SetGetVectorSizeUInt32(unsafe.Pointer(api[59]))
    C.Plugify_SetGetVectorSizeUInt64(unsafe.Pointer(api[60]))
    C.Plugify_SetGetVectorSizeUIntPtr(unsafe.Pointer(api[61]))
    C.Plugify_SetGetVectorSizeFloat(unsafe.Pointer(api[62]))
    C.Plugify_SetGetVectorSizeDouble(unsafe.Pointer(api[63]))
    C.Plugify_SetGetVectorSizeString(unsafe.Pointer(api[64]))
    C.Plugify_SetGetVectorData(unsafe.Pointer(api[65]))
    C.Plugify_SetGetVectorDataB(unsafe.Pointer(api[66]))
    C.Plugify_SetDeleteVectorDataB(unsafe.Pointer(api[67]))
    C.Plugify_SetGetVectorDataS(unsafe.Pointer(api[68]))
    C.Plugify_SetDeleteVectorDataS(unsafe.Pointer(api[69]))
	return 0
}

//export Plugify_PluginStart
func Plugify_PluginStart() {
	plugify.fnPluginStartCallback()
}

func OnPluginStart(fn PluginStartCallback) {
	plugify.fnPluginStartCallback = fn
}

//export Plugify_PluginEnd
func Plugify_PluginEnd() {
	plugify.fnPluginEndCallback()
}

func OnPluginEnd(fn PluginEndCallback) {
	plugify.fnPluginEndCallback = fn
}

type Vector2 struct {
    x float32
    y float32
}
type Vector3 struct {
    x float32
    y float32
    z float32
}
type Vector4 struct {
    x float32
    y float32
    z float32
    w float32
}