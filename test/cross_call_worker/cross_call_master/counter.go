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

// Generated from cross_call_master (group: counter)

var P_CounterCreate = func(initialValue int64) uintptr {
	var __retVal uintptr
	__initialValue := C.int64_t(initialValue)
	__retVal = uintptr(C.CounterCreate(__initialValue))
	return __retVal
}

// CounterCreate
func CounterCreate(initialValue int64) uintptr {
	defer plugify.Scope("cross_call_master::CounterCreate", ModuleName, 3)()
	return P_CounterCreate(initialValue)
}

var P_CounterCreateZero = func() uintptr {
	__retVal := uintptr(C.CounterCreateZero())
	return __retVal
}

// CounterCreateZero
func CounterCreateZero() uintptr {
	defer plugify.Scope("cross_call_master::CounterCreateZero", ModuleName, 3)()
	return P_CounterCreateZero()
}

var P_CounterGetValue = func(counter uintptr) int64 {
	var __retVal int64
	__counter := C.uintptr_t(counter)
	__retVal = int64(C.CounterGetValue(__counter))
	return __retVal
}

// CounterGetValue
func CounterGetValue(counter uintptr) int64 {
	defer plugify.Scope("cross_call_master::CounterGetValue", ModuleName, 3)()
	return P_CounterGetValue(counter)
}

var P_CounterSetValue = func(counter uintptr, value int64) {
	__counter := C.uintptr_t(counter)
	__value := C.int64_t(value)
	C.CounterSetValue(__counter, __value)
}

// CounterSetValue
func CounterSetValue(counter uintptr, value int64) {
	defer plugify.Scope("cross_call_master::CounterSetValue", ModuleName, 3)()
	P_CounterSetValue(counter, value)
}

var P_CounterIncrement = func(counter uintptr) {
	__counter := C.uintptr_t(counter)
	C.CounterIncrement(__counter)
}

// CounterIncrement
func CounterIncrement(counter uintptr) {
	defer plugify.Scope("cross_call_master::CounterIncrement", ModuleName, 3)()
	P_CounterIncrement(counter)
}

var P_CounterDecrement = func(counter uintptr) {
	__counter := C.uintptr_t(counter)
	C.CounterDecrement(__counter)
}

// CounterDecrement
func CounterDecrement(counter uintptr) {
	defer plugify.Scope("cross_call_master::CounterDecrement", ModuleName, 3)()
	P_CounterDecrement(counter)
}

var P_CounterAdd = func(counter uintptr, amount int64) {
	__counter := C.uintptr_t(counter)
	__amount := C.int64_t(amount)
	C.CounterAdd(__counter, __amount)
}

// CounterAdd
func CounterAdd(counter uintptr, amount int64) {
	defer plugify.Scope("cross_call_master::CounterAdd", ModuleName, 3)()
	P_CounterAdd(counter, amount)
}

var P_CounterReset = func(counter uintptr) {
	__counter := C.uintptr_t(counter)
	C.CounterReset(__counter)
}

// CounterReset
func CounterReset(counter uintptr) {
	defer plugify.Scope("cross_call_master::CounterReset", ModuleName, 3)()
	P_CounterReset(counter)
}

var P_CounterIsPositive = func(counter uintptr) bool {
	var __retVal bool
	__counter := C.uintptr_t(counter)
	__retVal = bool(C.CounterIsPositive(__counter))
	return __retVal
}

// CounterIsPositive
func CounterIsPositive(counter uintptr) bool {
	defer plugify.Scope("cross_call_master::CounterIsPositive", ModuleName, 3)()
	return P_CounterIsPositive(counter)
}

var P_CounterCompare = func(value1 int64, value2 int64) int32 {
	var __retVal int32
	__value1 := C.int64_t(value1)
	__value2 := C.int64_t(value2)
	__retVal = int32(C.CounterCompare(__value1, __value2))
	return __retVal
}

// CounterCompare
func CounterCompare(value1 int64, value2 int64) int32 {
	defer plugify.Scope("cross_call_master::CounterCompare", ModuleName, 3)()
	return P_CounterCompare(value1, value2)
}

var P_CounterSum = func(values []int64) int64 {
	var __retVal int64
	__values := plugify.ConstructVectorInt64(values)
	plugify.Block{
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
	return P_CounterSum(values)
}
