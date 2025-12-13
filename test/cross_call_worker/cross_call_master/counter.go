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

func CounterCreate(initialValue int64) uintptr {
	var __retVal uintptr
	__initialValue := C.int64_t(initialValue)
	__retVal = uintptr(C.CounterCreate(__initialValue))
	return __retVal
}

func CounterCreateZero() uintptr {
	__retVal := uintptr(C.CounterCreateZero())
	return __retVal
}

func CounterGetValue(counter uintptr) int64 {
	var __retVal int64
	__counter := C.uintptr_t(counter)
	__retVal = int64(C.CounterGetValue(__counter))
	return __retVal
}

func CounterSetValue(counter uintptr, value int64) {
	__counter := C.uintptr_t(counter)
	__value := C.int64_t(value)
	C.CounterSetValue(__counter, __value)
}

func CounterIncrement(counter uintptr) {
	__counter := C.uintptr_t(counter)
	C.CounterIncrement(__counter)
}

func CounterDecrement(counter uintptr) {
	__counter := C.uintptr_t(counter)
	C.CounterDecrement(__counter)
}

func CounterAdd(counter uintptr, amount int64) {
	__counter := C.uintptr_t(counter)
	__amount := C.int64_t(amount)
	C.CounterAdd(__counter, __amount)
}

func CounterReset(counter uintptr) {
	__counter := C.uintptr_t(counter)
	C.CounterReset(__counter)
}

func CounterIsPositive(counter uintptr) bool {
	var __retVal bool
	__counter := C.uintptr_t(counter)
	__retVal = bool(C.CounterIsPositive(__counter))
	return __retVal
}

func CounterCompare(value1 int64, value2 int64) int32 {
	var __retVal int32
	__value1 := C.int64_t(value1)
	__value2 := C.int64_t(value2)
	__retVal = int32(C.CounterCompare(__value1, __value2))
	return __retVal
}

func CounterSum(values []int64) int64 {
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
