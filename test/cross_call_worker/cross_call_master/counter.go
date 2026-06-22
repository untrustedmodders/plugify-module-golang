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
var _ = plugify.ApiVersion

// Generated from cross_call_master (group: counter)

var _CounterCreate = func(initialValue int64) uintptr {
	var __retVal uintptr
	__initialValue := C.int64_t(initialValue)
	__retVal = uintptr(C.CounterCreate(__initialValue))
	return __retVal
}

// CounterCreate 
func CounterCreate(initialValue int64) uintptr {
	defer plugify.Scope("cross_call_master::CounterCreate", ModuleName, 3)()
	return _CounterCreate(initialValue)
}

var _CounterCreateZero = func() uintptr {
	__retVal := uintptr(C.CounterCreateZero())
	return __retVal
}

// CounterCreateZero 
func CounterCreateZero() uintptr {
	defer plugify.Scope("cross_call_master::CounterCreateZero", ModuleName, 3)()
	return _CounterCreateZero()
}

var _CounterGetValue = func(counter uintptr) int64 {
	var __retVal int64
	__counter := C.uintptr_t(counter)
	__retVal = int64(C.CounterGetValue(__counter))
	return __retVal
}

// CounterGetValue 
func CounterGetValue(counter uintptr) int64 {
	defer plugify.Scope("cross_call_master::CounterGetValue", ModuleName, 3)()
	return _CounterGetValue(counter)
}

var _CounterSetValue = func(counter uintptr, value int64) {
	__counter := C.uintptr_t(counter)
	__value := C.int64_t(value)
	C.CounterSetValue(__counter, __value)
}

// CounterSetValue 
func CounterSetValue(counter uintptr, value int64) {
	defer plugify.Scope("cross_call_master::CounterSetValue", ModuleName, 3)()
	_CounterSetValue(counter, value)
}

var _CounterIncrement = func(counter uintptr) {
	__counter := C.uintptr_t(counter)
	C.CounterIncrement(__counter)
}

// CounterIncrement 
func CounterIncrement(counter uintptr) {
	defer plugify.Scope("cross_call_master::CounterIncrement", ModuleName, 3)()
	_CounterIncrement(counter)
}

var _CounterDecrement = func(counter uintptr) {
	__counter := C.uintptr_t(counter)
	C.CounterDecrement(__counter)
}

// CounterDecrement 
func CounterDecrement(counter uintptr) {
	defer plugify.Scope("cross_call_master::CounterDecrement", ModuleName, 3)()
	_CounterDecrement(counter)
}

var _CounterAdd = func(counter uintptr, amount int64) {
	__counter := C.uintptr_t(counter)
	__amount := C.int64_t(amount)
	C.CounterAdd(__counter, __amount)
}

// CounterAdd 
func CounterAdd(counter uintptr, amount int64) {
	defer plugify.Scope("cross_call_master::CounterAdd", ModuleName, 3)()
	_CounterAdd(counter, amount)
}

var _CounterReset = func(counter uintptr) {
	__counter := C.uintptr_t(counter)
	C.CounterReset(__counter)
}

// CounterReset 
func CounterReset(counter uintptr) {
	defer plugify.Scope("cross_call_master::CounterReset", ModuleName, 3)()
	_CounterReset(counter)
}

var _CounterIsPositive = func(counter uintptr) bool {
	var __retVal bool
	__counter := C.uintptr_t(counter)
	__retVal = bool(C.CounterIsPositive(__counter))
	return __retVal
}

// CounterIsPositive 
func CounterIsPositive(counter uintptr) bool {
	defer plugify.Scope("cross_call_master::CounterIsPositive", ModuleName, 3)()
	return _CounterIsPositive(counter)
}

var _CounterCompare = func(value1 int64, value2 int64) int32 {
	var __retVal int32
	__value1 := C.int64_t(value1)
	__value2 := C.int64_t(value2)
	__retVal = int32(C.CounterCompare(__value1, __value2))
	return __retVal
}

// CounterCompare 
func CounterCompare(value1 int64, value2 int64) int32 {
	defer plugify.Scope("cross_call_master::CounterCompare", ModuleName, 3)()
	return _CounterCompare(value1, value2)
}

var _CounterSum = func(values []int64) int64 {
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

// CounterSum 
func CounterSum(values []int64) int64 {
	defer plugify.Scope("cross_call_master::CounterSum", ModuleName, 3)()
	return _CounterSum(values)
}

