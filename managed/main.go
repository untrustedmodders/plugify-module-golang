package main

/*
#include "main.h"
*/
import "C"
import (
	"fmt"
	"plugin"
	"reflect"
	"unsafe"

	"github.com/untrustedmodders/go-plugify"
)

var _ = plugify.ApiVersion

var (
	nextID    int
	plugins   = map[int]*plugin.Plugin{}
	lastError string
)

//export Open
func Open(path string) int {
	pl, err := plugin.Open(path)
	if err != nil {
		lastError = err.Error()
		return 0
	}

	for i, p := range plugins {
		if p == pl {
			return i
		}
	}

	nextID++

	id := nextID
	plugins[id] = pl
	return id
}

//export Call
func Call(id int, method string) bool {
	pl, ok := plugins[id]
	if !ok {
		return false
	}
	sym, err := pl.Lookup(method)
	if err != nil {
		lastError = err.Error()
		return false
	}

	sym.(func())()
	return true
}

//export Find
func Find(id int, method string) bool {
	pl, ok := plugins[id]
	if !ok {
		return false
	}
	sym, err := pl.Lookup(method)
	if err != nil {
		lastError = err.Error()
		return false
	}

	return sym != nil
}

//export Bind
func Bind(from int, to int, fromMethod string, toVar string) bool {
	fromPl, ok := plugins[from]
	if !ok {
		return false
	}
	fromSym, err := fromPl.Lookup(fromMethod)
	if err != nil {
		lastError = err.Error()
		return false
	}

	toPl, ok := plugins[to]
	if !ok {
		return false
	}
	toSym, err := toPl.Lookup(toVar)
	if err != nil {
		lastError = err.Error()
		return false
	}

	// plugin symbols are always pointers — unwrap from
	fromVal := reflect.ValueOf(fromSym)
	if fromVal.Kind() == reflect.Ptr {
		fromVal = fromVal.Elem()
	}
	if fromVal.Kind() != reflect.Func {
		lastError = fmt.Sprintf("from symbol is not a function but a %v", fromVal.Kind())
		return false
	}

	// to must be a pointer to a func variable, OR a pointer to a pointer to a func variable
	toVal := reflect.ValueOf(toSym)
	if toVal.Kind() != reflect.Ptr {
		lastError = fmt.Sprintf("to symbol must be a pointer (or pointer-to-pointer) to a func variable, but got %v", toVal.Type())
		return false
	}

	// unwrap one level: now toVal should point at either a func or another pointer
	elem := toVal.Elem()

	switch elem.Kind() {
	case reflect.Func:
		// toSym was *Delegate — toVal already points at the func slot we assign into
	case reflect.Ptr:
		// toSym was **Delegate — descend one more level
		if elem.IsNil() {
			// need a non-nil *Delegate to assign into; allocate one
			newPtr := reflect.New(elem.Type().Elem())
			elem.Set(newPtr)
		}
		toVal = elem // now toVal.Elem() will be the actual Delegate slot
		elem = toVal.Elem()
		if elem.Kind() != reflect.Func {
			lastError = fmt.Sprintf("to symbol must ultimately point to a func variable, but got %v", elem.Type())
			return false
		}
	default:
		lastError = fmt.Sprintf("to symbol must be a pointer (or pointer-to-pointer) to a func variable, but got %v", toVal.Type())
		return false
	}

	fromType := fromVal.Type()
	toType := elem.Type()

	// check compatibility before attempting any assignment or bridge
	if !compatible(fromType, toType) {
		lastError = fmt.Sprintf("incompatible function signatures: %v vs %v", fromType, toType)
		return false
	}

	// fast path — types are identical, direct assign
	if fromType == toType {
		toVal.Elem().Set(fromVal)
		return true
	}

	// Go handles it natively — direct assign via Convert
	if fromType.AssignableTo(toType) {
		toVal.Elem().Set(fromVal)
		return true
	}
	if fromType.ConvertibleTo(toType) {
		toVal.Elem().Set(fromVal.Convert(toType))
		return true
	}

	// slow path — structurally compatible but nominally different
	// named types in signatures cross plugin boundary, build bridge
	toVal.Elem().Set(bridgeFunc(fromVal, toType))

	return true
}

//export Error
func Error() C.String {
	buffer := &C.error_buffer[0]
	slice := unsafe.Slice((*byte)(unsafe.Pointer(buffer)), C.sizeof_error_buffer)
	length := copy(slice, lastError)
	return C.String{
		p: buffer,
		n: C.ptrdiff_t(length),
	}
}

// compatible checks if two types can be bridged.
// Defers to Go's own rules first, falls back to structural
// check only for func types with named params across plugin boundary.
func compatible(a, b reflect.Type) bool {
	if a == b {
		return true
	}

	if a.AssignableTo(b) || a.ConvertibleTo(b) {
		return true
	}

	switch a.Kind() {
	case reflect.Func:
		if b.Kind() != reflect.Func {
			return false
		}
		return funcSignaturesMatch(a, b)
	case reflect.Ptr:
		if b.Kind() != reflect.Ptr {
			return false
		}
		return compatible(a.Elem(), b.Elem())
	case reflect.Slice:
		if b.Kind() != reflect.Slice {
			return false
		}
		return compatible(a.Elem(), b.Elem())
	case reflect.Map:
		if b.Kind() != reflect.Map {
			return false
		}
		return compatible(a.Key(), b.Key()) && compatible(a.Elem(), b.Elem())
	case reflect.Struct:
		if b.Kind() != reflect.Struct || a.NumField() != b.NumField() {
			return false
		}
		for i := 0; i < a.NumField(); i++ {
			if !compatible(a.Field(i).Type, b.Field(i).Type) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// funcSignaturesMatch structurally compares two func types,
// recursing into param/return types to handle nested named funcs
// and named types that cross plugin boundaries.
func funcSignaturesMatch(a, b reflect.Type) bool {
	if a.NumIn() != b.NumIn() || a.NumOut() != b.NumOut() {
		return false
	}
	if a.IsVariadic() != b.IsVariadic() {
		return false
	}
	for i := 0; i < a.NumIn(); i++ {
		if !compatible(a.In(i), b.In(i)) {
			return false
		}
	}
	for i := 0; i < a.NumOut(); i++ {
		if !compatible(a.Out(i), b.Out(i)) {
			return false
		}
	}
	return true
}

// bridgeFunc wraps from into a new func value of targetType,
// converting all arguments and return values recursively.
func bridgeFunc(from reflect.Value, targetType reflect.Type) reflect.Value {
	fromType := from.Type()
	return reflect.MakeFunc(targetType, func(args []reflect.Value) []reflect.Value {
		convertedArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			convertedArgs[i] = convertValue(arg, fromType.In(i))
		}

		results := from.Call(convertedArgs)

		// convertValue allocates a *new* object for pointer args (different
		// backing memory than the caller's pointer), so writes the callee makes
		// through that pointer (e.g. *[]T mutation/append, *struct field writes)
		// only land in convertedArgs[i] and are lost unless copied back here.
		for i, arg := range args {
			if arg.Kind() == reflect.Ptr && !arg.IsNil() &&
				!convertedArgs[i].IsNil() && arg.Type() != convertedArgs[i].Type() {
				arg.Elem().Set(convertValue(convertedArgs[i].Elem(), arg.Type().Elem()))
			}
		}

		convertedResults := make([]reflect.Value, len(results))
		for i, r := range results {
			convertedResults[i] = convertValue(r, targetType.Out(i))
		}
		return convertedResults
	})
}

// convertValue converts v to targetType, recursively bridging
// func types and struct fields that cross plugin boundaries.
func convertValue(v reflect.Value, target reflect.Type) reflect.Value {
	if v.Type() == target {
		return v
	}

	// Go handles it natively
	if v.Type().AssignableTo(target) {
		return v
	}
	if v.Type().ConvertibleTo(target) {
		return v.Convert(target)
	}

	switch v.Kind() {
	case reflect.Func:
		if v.IsNil() {
			return reflect.Zero(target)
		}
		return bridgeFunc(v, target)

	case reflect.Ptr:
		if v.IsNil() {
			return reflect.Zero(target)
		}
		// allocate new pointer of target type, recursively convert pointee
		out := reflect.New(target.Elem())
		out.Elem().Set(convertValue(v.Elem(), target.Elem()))
		return out

	case reflect.Struct:
		if v.Type().AssignableTo(target) {
			return v
		}
		out := reflect.New(target).Elem()
		for i := 0; i < v.NumField(); i++ {
			out.Field(i).Set(convertValue(v.Field(i), target.Field(i).Type))
		}
		return out

	case reflect.Slice:
		if v.IsNil() {
			return reflect.Zero(target)
		}
		out := reflect.MakeSlice(target, v.Len(), v.Cap())
		for i := 0; i < v.Len(); i++ {
			out.Index(i).Set(convertValue(v.Index(i), target.Elem()))
		}
		return out

	case reflect.Map:
		if v.IsNil() {
			return reflect.Zero(target)
		}
		out := reflect.MakeMap(target)
		for _, key := range v.MapKeys() {
			convertedKey := convertValue(key, target.Key())
			convertedVal := convertValue(v.MapIndex(key), target.Elem())
			out.SetMapIndex(convertedKey, convertedVal)
		}
		return out
	default:
		panic("convertValue: incompatible types " + v.Type().String() + " -> " + target.String())
	}
}

func main() {
}

func init() {
}
