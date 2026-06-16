package cross_call_master

/*
#include "resource.h"
#cgo noescape ResourceHandleCreate
#cgo noescape ResourceHandleCreateDefault
#cgo noescape ResourceHandleDestroy
#cgo noescape ResourceHandleGetId
#cgo noescape ResourceHandleGetName
#cgo noescape ResourceHandleSetName
#cgo noescape ResourceHandleIncrementCounter
#cgo noescape ResourceHandleGetCounter
#cgo noescape ResourceHandleAddData
#cgo noescape ResourceHandleGetData
#cgo noescape ResourceHandleGetAliveCount
#cgo noescape ResourceHandleGetTotalCreated
#cgo noescape ResourceHandleGetTotalDestroyed
*/
import "C"
import (
	"errors"
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

// Generated from cross_call_master (group: resource)

var P_ResourceHandleCreate = func(id int32, name string) uintptr {
	var __retVal uintptr
	__id := C.int32_t(id)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.ResourceHandleCreate(__id, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// ResourceHandleCreate 
func ResourceHandleCreate(id int32, name string) uintptr {
	defer plugify.Scope("cross_call_master::ResourceHandleCreate", buildInfo, 3)()
	return P_ResourceHandleCreate(id, name)
}

var P_ResourceHandleCreateDefault = func() uintptr {
	__retVal := uintptr(C.ResourceHandleCreateDefault())
	return __retVal
}

// ResourceHandleCreateDefault 
func ResourceHandleCreateDefault() uintptr {
	defer plugify.Scope("cross_call_master::ResourceHandleCreateDefault", buildInfo, 3)()
	return P_ResourceHandleCreateDefault()
}

var P_ResourceHandleDestroy = func(handle uintptr) {
	__handle := C.uintptr_t(handle)
	C.ResourceHandleDestroy(__handle)
}

// ResourceHandleDestroy 
func ResourceHandleDestroy(handle uintptr) {
	defer plugify.Scope("cross_call_master::ResourceHandleDestroy", buildInfo, 3)()
	P_ResourceHandleDestroy(handle)
}

var P_ResourceHandleGetId = func(handle uintptr) int32 {
	var __retVal int32
	__handle := C.uintptr_t(handle)
	__retVal = int32(C.ResourceHandleGetId(__handle))
	return __retVal
}

// ResourceHandleGetId 
func ResourceHandleGetId(handle uintptr) int32 {
	defer plugify.Scope("cross_call_master::ResourceHandleGetId", buildInfo, 3)()
	return P_ResourceHandleGetId(handle)
}

var P_ResourceHandleGetName = func(handle uintptr) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__handle := C.uintptr_t(handle)
	plugify.Block {
		Try: func() {
			__native := C.ResourceHandleGetName(__handle)
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

// ResourceHandleGetName 
func ResourceHandleGetName(handle uintptr) string {
	defer plugify.Scope("cross_call_master::ResourceHandleGetName", buildInfo, 3)()
	return P_ResourceHandleGetName(handle)
}

var P_ResourceHandleSetName = func(handle uintptr, name string) {
	__handle := C.uintptr_t(handle)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			C.ResourceHandleSetName(__handle, (*C.String)(unsafe.Pointer(&__name)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// ResourceHandleSetName 
func ResourceHandleSetName(handle uintptr, name string) {
	defer plugify.Scope("cross_call_master::ResourceHandleSetName", buildInfo, 3)()
	P_ResourceHandleSetName(handle, name)
}

var P_ResourceHandleIncrementCounter = func(handle uintptr) {
	__handle := C.uintptr_t(handle)
	C.ResourceHandleIncrementCounter(__handle)
}

// ResourceHandleIncrementCounter 
func ResourceHandleIncrementCounter(handle uintptr) {
	defer plugify.Scope("cross_call_master::ResourceHandleIncrementCounter", buildInfo, 3)()
	P_ResourceHandleIncrementCounter(handle)
}

var P_ResourceHandleGetCounter = func(handle uintptr) int32 {
	var __retVal int32
	__handle := C.uintptr_t(handle)
	__retVal = int32(C.ResourceHandleGetCounter(__handle))
	return __retVal
}

// ResourceHandleGetCounter 
func ResourceHandleGetCounter(handle uintptr) int32 {
	defer plugify.Scope("cross_call_master::ResourceHandleGetCounter", buildInfo, 3)()
	return P_ResourceHandleGetCounter(handle)
}

var P_ResourceHandleAddData = func(handle uintptr, value float32) {
	__handle := C.uintptr_t(handle)
	__value := C.float(value)
	C.ResourceHandleAddData(__handle, __value)
}

// ResourceHandleAddData 
func ResourceHandleAddData(handle uintptr, value float32) {
	defer plugify.Scope("cross_call_master::ResourceHandleAddData", buildInfo, 3)()
	P_ResourceHandleAddData(handle, value)
}

var P_ResourceHandleGetData = func(handle uintptr) []float32 {
	var __retVal []float32
	var __retVal_native plugify.PlgVector
	__handle := C.uintptr_t(handle)
	plugify.Block {
		Try: func() {
			__native := C.ResourceHandleGetData(__handle)
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

// ResourceHandleGetData 
func ResourceHandleGetData(handle uintptr) []float32 {
	defer plugify.Scope("cross_call_master::ResourceHandleGetData", buildInfo, 3)()
	return P_ResourceHandleGetData(handle)
}

var P_ResourceHandleGetAliveCount = func() int32 {
	__retVal := int32(C.ResourceHandleGetAliveCount())
	return __retVal
}

// ResourceHandleGetAliveCount 
func ResourceHandleGetAliveCount() int32 {
	defer plugify.Scope("cross_call_master::ResourceHandleGetAliveCount", buildInfo, 3)()
	return P_ResourceHandleGetAliveCount()
}

var P_ResourceHandleGetTotalCreated = func() int32 {
	__retVal := int32(C.ResourceHandleGetTotalCreated())
	return __retVal
}

// ResourceHandleGetTotalCreated 
func ResourceHandleGetTotalCreated() int32 {
	defer plugify.Scope("cross_call_master::ResourceHandleGetTotalCreated", buildInfo, 3)()
	return P_ResourceHandleGetTotalCreated()
}

var P_ResourceHandleGetTotalDestroyed = func() int32 {
	__retVal := int32(C.ResourceHandleGetTotalDestroyed())
	return __retVal
}

// ResourceHandleGetTotalDestroyed 
func ResourceHandleGetTotalDestroyed() int32 {
	defer plugify.Scope("cross_call_master::ResourceHandleGetTotalDestroyed", buildInfo, 3)()
	return P_ResourceHandleGetTotalDestroyed()
}

