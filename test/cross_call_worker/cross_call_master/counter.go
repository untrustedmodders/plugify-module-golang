package cross_call_master

/*
#include "counter.h"
#cgo noescape CounterCreate
#cgo noescape CounterCreateZero
#cgo noescape CounterGetValue
#cgo noescape CounterSetValue
#cgo noescape CounterIncrement
#cgo noescape CounterDecrement
#cgo noescape CounterAdd
#cgo noescape CounterReset
#cgo noescape CounterIsPositive
#cgo noescape CounterCompare
#cgo noescape CounterSum
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

// Generated from cross_call_master (group: counter)

// CounterCreate 
func CounterCreate(initialValue int64) uintptr {
	plugify.Log("cross_call_master::CounterCreate", plugify.Trace, 2)
	var __retVal uintptr
	__initialValue := C.int64_t(initialValue)
	__retVal = uintptr(C.CounterCreate(__initialValue))
	return __retVal
}

// CounterCreateZero 
func CounterCreateZero() uintptr {
	plugify.Log("cross_call_master::CounterCreateZero", plugify.Trace, 2)
	__retVal := uintptr(C.CounterCreateZero())
	return __retVal
}

// CounterGetValue 
func CounterGetValue(counter uintptr) int64 {
	plugify.Log("cross_call_master::CounterGetValue", plugify.Trace, 2)
	var __retVal int64
	__counter := C.uintptr_t(counter)
	__retVal = int64(C.CounterGetValue(__counter))
	return __retVal
}

// CounterSetValue 
func CounterSetValue(counter uintptr, value int64) {
	plugify.Log("cross_call_master::CounterSetValue", plugify.Trace, 2)
	__counter := C.uintptr_t(counter)
	__value := C.int64_t(value)
	C.CounterSetValue(__counter, __value)
}

// CounterIncrement 
func CounterIncrement(counter uintptr) {
	plugify.Log("cross_call_master::CounterIncrement", plugify.Trace, 2)
	__counter := C.uintptr_t(counter)
	C.CounterIncrement(__counter)
}

// CounterDecrement 
func CounterDecrement(counter uintptr) {
	plugify.Log("cross_call_master::CounterDecrement", plugify.Trace, 2)
	__counter := C.uintptr_t(counter)
	C.CounterDecrement(__counter)
}

// CounterAdd 
func CounterAdd(counter uintptr, amount int64) {
	plugify.Log("cross_call_master::CounterAdd", plugify.Trace, 2)
	__counter := C.uintptr_t(counter)
	__amount := C.int64_t(amount)
	C.CounterAdd(__counter, __amount)
}

// CounterReset 
func CounterReset(counter uintptr) {
	plugify.Log("cross_call_master::CounterReset", plugify.Trace, 2)
	__counter := C.uintptr_t(counter)
	C.CounterReset(__counter)
}

// CounterIsPositive 
func CounterIsPositive(counter uintptr) bool {
	plugify.Log("cross_call_master::CounterIsPositive", plugify.Trace, 2)
	var __retVal bool
	__counter := C.uintptr_t(counter)
	__retVal = bool(C.CounterIsPositive(__counter))
	return __retVal
}

// CounterCompare 
func CounterCompare(value1 int64, value2 int64) int32 {
	plugify.Log("cross_call_master::CounterCompare", plugify.Trace, 2)
	var __retVal int32
	__value1 := C.int64_t(value1)
	__value2 := C.int64_t(value2)
	__retVal = int32(C.CounterCompare(__value1, __value2))
	return __retVal
}

// CounterSum 
func CounterSum(values []int64) int64 {
	plugify.Log("cross_call_master::CounterSum", plugify.Trace, 2)
	var __retVal int64
	__values := plugify.ConstructVectorInt64(values)
	plugify.Block {
		Try: func() {
			__retVal = int64(C.CounterSum((*C.Vector)(unsafe.Pointer(&__values))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt64(&__values)
		},
	}.Do()
	return __retVal
}

