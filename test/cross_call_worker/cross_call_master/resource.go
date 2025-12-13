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
var _ = plugify.Plugin.Loaded

// Generated from cross_call_master (group: resource)

func ResourceHandleCreate(id int32, name string) uintptr {
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

func ResourceHandleCreateDefault() uintptr {
	__retVal := uintptr(C.ResourceHandleCreateDefault())
	return __retVal
}

func ResourceHandleDestroy(handle uintptr) {
	__handle := C.uintptr_t(handle)
	C.ResourceHandleDestroy(__handle)
}

func ResourceHandleGetId(handle uintptr) int32 {
	var __retVal int32
	__handle := C.uintptr_t(handle)
	__retVal = int32(C.ResourceHandleGetId(__handle))
	return __retVal
}

func ResourceHandleGetName(handle uintptr) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__handle := C.uintptr_t(handle)
	plugify.Block {
		Try: func() {
			__native := C.ResourceHandleGetName(__handle)
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

func ResourceHandleSetName(handle uintptr, name string) {
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

func ResourceHandleIncrementCounter(handle uintptr) {
	__handle := C.uintptr_t(handle)
	C.ResourceHandleIncrementCounter(__handle)
}

func ResourceHandleGetCounter(handle uintptr) int32 {
	var __retVal int32
	__handle := C.uintptr_t(handle)
	__retVal = int32(C.ResourceHandleGetCounter(__handle))
	return __retVal
}

func ResourceHandleAddData(handle uintptr, value float32) {
	__handle := C.uintptr_t(handle)
	__value := C.float(value)
	C.ResourceHandleAddData(__handle, __value)
}

func ResourceHandleGetData(handle uintptr) []float32 {
	var __retVal []float32
	var __retVal_native plugify.PlgVector
	__handle := C.uintptr_t(handle)
	plugify.Block {
		Try: func() {
			__native := C.ResourceHandleGetData(__handle)
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

func ResourceHandleGetAliveCount() int32 {
	__retVal := int32(C.ResourceHandleGetAliveCount())
	return __retVal
}

func ResourceHandleGetTotalCreated() int32 {
	__retVal := int32(C.ResourceHandleGetTotalCreated())
	return __retVal
}

func ResourceHandleGetTotalDestroyed() int32 {
	__retVal := int32(C.ResourceHandleGetTotalDestroyed())
	return __retVal
}

