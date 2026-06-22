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

var _ResourceHandleCreate = func(id int32, name string) uintptr {
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
	defer plugify.Scope("cross_call_master::ResourceHandleCreate", ModuleName, 3)()
	return _ResourceHandleCreate(id, name)
}

var _ResourceHandleCreateDefault = func() uintptr {
	__retVal := uintptr(C.ResourceHandleCreateDefault())
	return __retVal
}

// ResourceHandleCreateDefault 
func ResourceHandleCreateDefault() uintptr {
	defer plugify.Scope("cross_call_master::ResourceHandleCreateDefault", ModuleName, 3)()
	return _ResourceHandleCreateDefault()
}

var _ResourceHandleDestroy = func(handle uintptr) {
	__handle := C.uintptr_t(handle)
	C.ResourceHandleDestroy(__handle)
}

// ResourceHandleDestroy 
func ResourceHandleDestroy(handle uintptr) {
	defer plugify.Scope("cross_call_master::ResourceHandleDestroy", ModuleName, 3)()
	_ResourceHandleDestroy(handle)
}

var _ResourceHandleGetId = func(handle uintptr) int32 {
	var __retVal int32
	__handle := C.uintptr_t(handle)
	__retVal = int32(C.ResourceHandleGetId(__handle))
	return __retVal
}

// ResourceHandleGetId 
func ResourceHandleGetId(handle uintptr) int32 {
	defer plugify.Scope("cross_call_master::ResourceHandleGetId", ModuleName, 3)()
	return _ResourceHandleGetId(handle)
}

var _ResourceHandleGetName = func(handle uintptr) string {
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
	defer plugify.Scope("cross_call_master::ResourceHandleGetName", ModuleName, 3)()
	return _ResourceHandleGetName(handle)
}

var _ResourceHandleSetName = func(handle uintptr, name string) {
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
	defer plugify.Scope("cross_call_master::ResourceHandleSetName", ModuleName, 3)()
	_ResourceHandleSetName(handle, name)
}

var _ResourceHandleIncrementCounter = func(handle uintptr) {
	__handle := C.uintptr_t(handle)
	C.ResourceHandleIncrementCounter(__handle)
}

// ResourceHandleIncrementCounter 
func ResourceHandleIncrementCounter(handle uintptr) {
	defer plugify.Scope("cross_call_master::ResourceHandleIncrementCounter", ModuleName, 3)()
	_ResourceHandleIncrementCounter(handle)
}

var _ResourceHandleGetCounter = func(handle uintptr) int32 {
	var __retVal int32
	__handle := C.uintptr_t(handle)
	__retVal = int32(C.ResourceHandleGetCounter(__handle))
	return __retVal
}

// ResourceHandleGetCounter 
func ResourceHandleGetCounter(handle uintptr) int32 {
	defer plugify.Scope("cross_call_master::ResourceHandleGetCounter", ModuleName, 3)()
	return _ResourceHandleGetCounter(handle)
}

var _ResourceHandleAddData = func(handle uintptr, value float32) {
	__handle := C.uintptr_t(handle)
	__value := C.float(value)
	C.ResourceHandleAddData(__handle, __value)
}

// ResourceHandleAddData 
func ResourceHandleAddData(handle uintptr, value float32) {
	defer plugify.Scope("cross_call_master::ResourceHandleAddData", ModuleName, 3)()
	_ResourceHandleAddData(handle, value)
}

var _ResourceHandleGetData = func(handle uintptr) []float32 {
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
	defer plugify.Scope("cross_call_master::ResourceHandleGetData", ModuleName, 3)()
	return _ResourceHandleGetData(handle)
}

var _ResourceHandleGetAliveCount = func() int32 {
	__retVal := int32(C.ResourceHandleGetAliveCount())
	return __retVal
}

// ResourceHandleGetAliveCount 
func ResourceHandleGetAliveCount() int32 {
	defer plugify.Scope("cross_call_master::ResourceHandleGetAliveCount", ModuleName, 3)()
	return _ResourceHandleGetAliveCount()
}

var _ResourceHandleGetTotalCreated = func() int32 {
	__retVal := int32(C.ResourceHandleGetTotalCreated())
	return __retVal
}

// ResourceHandleGetTotalCreated 
func ResourceHandleGetTotalCreated() int32 {
	defer plugify.Scope("cross_call_master::ResourceHandleGetTotalCreated", ModuleName, 3)()
	return _ResourceHandleGetTotalCreated()
}

var _ResourceHandleGetTotalDestroyed = func() int32 {
	__retVal := int32(C.ResourceHandleGetTotalDestroyed())
	return __retVal
}

// ResourceHandleGetTotalDestroyed 
func ResourceHandleGetTotalDestroyed() int32 {
	defer plugify.Scope("cross_call_master::ResourceHandleGetTotalDestroyed", ModuleName, 3)()
	return _ResourceHandleGetTotalDestroyed()
}

