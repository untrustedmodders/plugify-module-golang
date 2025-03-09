package main

// #include "autoexports.h"
import "C"
import (
	"github.com/untrustedmodders/go-plugify"
	"reflect"
	"unsafe"
)

// Exported methods

//export __NoParamReturnVoid
func __NoParamReturnVoid() {
	NoParamReturnVoid()
}

//export __NoParamReturnBool
func __NoParamReturnBool() bool {
	__result := NoParamReturnBool()
	return __result
}

//export __NoParamReturnChar8
func __NoParamReturnChar8() int8 {
	__result := NoParamReturnChar8()
	return __result
}

//export __NoParamReturnChar16
func __NoParamReturnChar16() uint16 {
	__result := NoParamReturnChar16()
	return __result
}

//export __NoParamReturnInt8
func __NoParamReturnInt8() int8 {
	__result := NoParamReturnInt8()
	return __result
}

//export __NoParamReturnInt16
func __NoParamReturnInt16() int16 {
	__result := NoParamReturnInt16()
	return __result
}

//export __NoParamReturnInt32
func __NoParamReturnInt32() int32 {
	__result := NoParamReturnInt32()
	return __result
}

//export __NoParamReturnInt64
func __NoParamReturnInt64() int64 {
	__result := NoParamReturnInt64()
	return __result
}

//export __NoParamReturnUInt8
func __NoParamReturnUInt8() uint8 {
	__result := NoParamReturnUInt8()
	return __result
}

//export __NoParamReturnUInt16
func __NoParamReturnUInt16() uint16 {
	__result := NoParamReturnUInt16()
	return __result
}

//export __NoParamReturnUInt32
func __NoParamReturnUInt32() uint32 {
	__result := NoParamReturnUInt32()
	return __result
}

//export __NoParamReturnUInt64
func __NoParamReturnUInt64() uint64 {
	__result := NoParamReturnUInt64()
	return __result
}

//export __NoParamReturnPointer
func __NoParamReturnPointer() uintptr {
	__result := NoParamReturnPointer()
	return __result
}

//export __NoParamReturnFloat
func __NoParamReturnFloat() float32 {
	__result := NoParamReturnFloat()
	return __result
}

//export __NoParamReturnDouble
func __NoParamReturnDouble() float64 {
	__result := NoParamReturnDouble()
	return __result
}

//export __NoParamReturnFunction
func __NoParamReturnFunction() unsafe.Pointer {
	__result := NoParamReturnFunction()
	return plugify.GetFunctionPointerForDelegate(__result)
}

//export __NoParamReturnString
func __NoParamReturnString() C.String {
	__result := NoParamReturnString()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __NoParamReturnAny
func __NoParamReturnAny() C.Variant {
	__result := NoParamReturnAny()
	__return := plugify.ConstructVariant(__result)
	return *(*C.Variant)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayBool
func __NoParamReturnArrayBool() C.Vector {
	__result := NoParamReturnArrayBool()
	__return := plugify.ConstructVectorBool(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayChar8
func __NoParamReturnArrayChar8() C.Vector {
	__result := NoParamReturnArrayChar8()
	__return := plugify.ConstructVectorChar8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayChar16
func __NoParamReturnArrayChar16() C.Vector {
	__result := NoParamReturnArrayChar16()
	__return := plugify.ConstructVectorChar16(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayInt8
func __NoParamReturnArrayInt8() C.Vector {
	__result := NoParamReturnArrayInt8()
	__return := plugify.ConstructVectorInt8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayInt16
func __NoParamReturnArrayInt16() C.Vector {
	__result := NoParamReturnArrayInt16()
	__return := plugify.ConstructVectorInt16(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayInt32
func __NoParamReturnArrayInt32() C.Vector {
	__result := NoParamReturnArrayInt32()
	__return := plugify.ConstructVectorInt32(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayInt64
func __NoParamReturnArrayInt64() C.Vector {
	__result := NoParamReturnArrayInt64()
	__return := plugify.ConstructVectorInt64(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayUInt8
func __NoParamReturnArrayUInt8() C.Vector {
	__result := NoParamReturnArrayUInt8()
	__return := plugify.ConstructVectorUInt8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayUInt16
func __NoParamReturnArrayUInt16() C.Vector {
	__result := NoParamReturnArrayUInt16()
	__return := plugify.ConstructVectorUInt16(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayUInt32
func __NoParamReturnArrayUInt32() C.Vector {
	__result := NoParamReturnArrayUInt32()
	__return := plugify.ConstructVectorUInt32(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayUInt64
func __NoParamReturnArrayUInt64() C.Vector {
	__result := NoParamReturnArrayUInt64()
	__return := plugify.ConstructVectorUInt64(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayPointer
func __NoParamReturnArrayPointer() C.Vector {
	__result := NoParamReturnArrayPointer()
	__return := plugify.ConstructVectorPointer(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayFloat
func __NoParamReturnArrayFloat() C.Vector {
	__result := NoParamReturnArrayFloat()
	__return := plugify.ConstructVectorFloat(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayDouble
func __NoParamReturnArrayDouble() C.Vector {
	__result := NoParamReturnArrayDouble()
	__return := plugify.ConstructVectorDouble(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayString
func __NoParamReturnArrayString() C.Vector {
	__result := NoParamReturnArrayString()
	__return := plugify.ConstructVectorString(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayAny
func __NoParamReturnArrayAny() C.Vector {
	__result := NoParamReturnArrayAny()
	__return := plugify.ConstructVectorVariant(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayVector2
func __NoParamReturnArrayVector2() C.Vector {
	__result := NoParamReturnArrayVector2()
	__return := plugify.ConstructVectorVector2(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayVector3
func __NoParamReturnArrayVector3() C.Vector {
	__result := NoParamReturnArrayVector3()
	__return := plugify.ConstructVectorVector3(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayVector4
func __NoParamReturnArrayVector4() C.Vector {
	__result := NoParamReturnArrayVector4()
	__return := plugify.ConstructVectorVector4(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayMatrix4x4
func __NoParamReturnArrayMatrix4x4() C.Vector {
	__result := NoParamReturnArrayMatrix4x4()
	__return := plugify.ConstructVectorMatrix4x4(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnVector2
func __NoParamReturnVector2() C.Vector2 {
	__result := NoParamReturnVector2()
	return *(*C.Vector2)(unsafe.Pointer(&__result))
}

//export __NoParamReturnVector3
func __NoParamReturnVector3() C.Vector3 {
	__result := NoParamReturnVector3()
	return *(*C.Vector3)(unsafe.Pointer(&__result))
}

//export __NoParamReturnVector4
func __NoParamReturnVector4() C.Vector4 {
	__result := NoParamReturnVector4()
	return *(*C.Vector4)(unsafe.Pointer(&__result))
}

//export __NoParamReturnMatrix4x4
func __NoParamReturnMatrix4x4() C.Matrix4x4 {
	__result := NoParamReturnMatrix4x4()
	return *(*C.Matrix4x4)(unsafe.Pointer(&__result))
}

//export __Param1
func __Param1(a int32) {
	Param1(a)
}

//export __Param2
func __Param2(a int32, b float32) {
	Param2(a, b)
}

//export __Param3
func __Param3(a int32, b float32, c float64) {
	Param3(a, b, c)
}

//export __Param4
func __Param4(a int32, b float32, c float64, d *C.Vector4) {
	Param4(a, b, c, *(*plugify.Vector4)(unsafe.Pointer(d)))
}

//export __Param5
func __Param5(a int32, b float32, c float64, d *C.Vector4, e *C.Vector) {
	Param5(a, b, c, *(*plugify.Vector4)(unsafe.Pointer(d)), plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(e))))
}

//export __Param6
func __Param6(a int32, b float32, c float64, d *C.Vector4, e *C.Vector, f int8) {
	Param6(a, b, c, *(*plugify.Vector4)(unsafe.Pointer(d)), plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(e))), f)
}

//export __Param7
func __Param7(a int32, b float32, c float64, d *C.Vector4, e *C.Vector, f int8, g *C.String) {
	Param7(a, b, c, *(*plugify.Vector4)(unsafe.Pointer(d)), plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(e))), f, plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(g))))
}

//export __Param8
func __Param8(a int32, b float32, c float64, d *C.Vector4, e *C.Vector, f int8, g *C.String, h uint16) {
	Param8(a, b, c, *(*plugify.Vector4)(unsafe.Pointer(d)), plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(e))), f, plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(g))), h)
}

//export __Param9
func __Param9(a int32, b float32, c float64, d *C.Vector4, e *C.Vector, f int8, g *C.String, h uint16, k int16) {
	Param9(a, b, c, *(*plugify.Vector4)(unsafe.Pointer(d)), plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(e))), f, plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(g))), h, k)
}

//export __Param10
func __Param10(a int32, b float32, c float64, d *C.Vector4, e *C.Vector, f int8, g *C.String, h uint16, k int16, l uintptr) {
	Param10(a, b, c, *(*plugify.Vector4)(unsafe.Pointer(d)), plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(e))), f, plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(g))), h, k, l)
}

//export __ParamRef1
func __ParamRef1(a *int32) {
	ParamRef1(a)
}

//export __ParamRef2
func __ParamRef2(a *int32, b *float32) {
	ParamRef2(a, b)
}

//export __ParamRef3
func __ParamRef3(a *int32, b *float32, c *float64) {
	ParamRef3(a, b, c)
}

//export __ParamRef4
func __ParamRef4(a *int32, b *float32, c *float64, d *C.Vector4) {
	ParamRef4(a, b, c, (*plugify.Vector4)(unsafe.Pointer(d)))
}

//export __ParamRef5
func __ParamRef5(a *int32, b *float32, c *float64, d *C.Vector4, e *C.Vector) {
	_e := plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(e)))
	ParamRef5(a, b, c, (*plugify.Vector4)(unsafe.Pointer(d)), &_e)
	plugify.AssignVectorInt64((*plugify.PlgVector)(unsafe.Pointer(e)), _e)
}

//export __ParamRef6
func __ParamRef6(a *int32, b *float32, c *float64, d *C.Vector4, e *C.Vector, f *int8) {
	_e := plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(e)))
	ParamRef6(a, b, c, (*plugify.Vector4)(unsafe.Pointer(d)), &_e, f)
	plugify.AssignVectorInt64((*plugify.PlgVector)(unsafe.Pointer(e)), _e)
}

//export __ParamRef7
func __ParamRef7(a *int32, b *float32, c *float64, d *C.Vector4, e *C.Vector, f *int8, g *C.String) {
	_e := plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(e)))
	_g := plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(g)))
	ParamRef7(a, b, c, (*plugify.Vector4)(unsafe.Pointer(d)), &_e, f, &_g)
	plugify.AssignVectorInt64((*plugify.PlgVector)(unsafe.Pointer(e)), _e)
	plugify.AssignString((*plugify.PlgString)(unsafe.Pointer(g)), _g)
}

//export __ParamRef8
func __ParamRef8(a *int32, b *float32, c *float64, d *C.Vector4, e *C.Vector, f *int8, g *C.String, h *uint16) {
	_e := plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(e)))
	_g := plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(g)))
	ParamRef8(a, b, c, (*plugify.Vector4)(unsafe.Pointer(d)), &_e, f, &_g, h)
	plugify.AssignVectorInt64((*plugify.PlgVector)(unsafe.Pointer(e)), _e)
	plugify.AssignString((*plugify.PlgString)(unsafe.Pointer(g)), _g)
}

//export __ParamRef9
func __ParamRef9(a *int32, b *float32, c *float64, d *C.Vector4, e *C.Vector, f *int8, g *C.String, h *uint16, k *int16) {
	_e := plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(e)))
	_g := plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(g)))
	ParamRef9(a, b, c, (*plugify.Vector4)(unsafe.Pointer(d)), &_e, f, &_g, h, k)
	plugify.AssignVectorInt64((*plugify.PlgVector)(unsafe.Pointer(e)), _e)
	plugify.AssignString((*plugify.PlgString)(unsafe.Pointer(g)), _g)
}

//export __ParamRef10
func __ParamRef10(a *int32, b *float32, c *float64, d *C.Vector4, e *C.Vector, f *int8, g *C.String, h *uint16, k *int16, l *uintptr) {
	_e := plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(e)))
	_g := plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(g)))
	ParamRef10(a, b, c, (*plugify.Vector4)(unsafe.Pointer(d)), &_e, f, &_g, h, k, l)
	plugify.AssignVectorInt64((*plugify.PlgVector)(unsafe.Pointer(e)), _e)
	plugify.AssignString((*plugify.PlgString)(unsafe.Pointer(g)), _g)
}

//export __ParamRefVectors
func __ParamRefVectors(p1 *C.Vector, p2 *C.Vector, p3 *C.Vector, p4 *C.Vector, p5 *C.Vector, p6 *C.Vector, p7 *C.Vector, p8 *C.Vector, p9 *C.Vector, p10 *C.Vector, p11 *C.Vector, p12 *C.Vector, p13 *C.Vector, p14 *C.Vector, p15 *C.Vector) {
	_p1 := plugify.GetVectorDataBool((*plugify.PlgVector)(unsafe.Pointer(p1)))
	_p2 := plugify.GetVectorDataChar8((*plugify.PlgVector)(unsafe.Pointer(p2)))
	_p3 := plugify.GetVectorDataChar16((*plugify.PlgVector)(unsafe.Pointer(p3)))
	_p4 := plugify.GetVectorDataInt8((*plugify.PlgVector)(unsafe.Pointer(p4)))
	_p5 := plugify.GetVectorDataInt16((*plugify.PlgVector)(unsafe.Pointer(p5)))
	_p6 := plugify.GetVectorDataInt32((*plugify.PlgVector)(unsafe.Pointer(p6)))
	_p7 := plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(p7)))
	_p8 := plugify.GetVectorDataUInt8((*plugify.PlgVector)(unsafe.Pointer(p8)))
	_p9 := plugify.GetVectorDataUInt16((*plugify.PlgVector)(unsafe.Pointer(p9)))
	_p10 := plugify.GetVectorDataUInt32((*plugify.PlgVector)(unsafe.Pointer(p10)))
	_p11 := plugify.GetVectorDataUInt64((*plugify.PlgVector)(unsafe.Pointer(p11)))
	_p12 := plugify.GetVectorDataPointer((*plugify.PlgVector)(unsafe.Pointer(p12)))
	_p13 := plugify.GetVectorDataFloat((*plugify.PlgVector)(unsafe.Pointer(p13)))
	_p14 := plugify.GetVectorDataDouble((*plugify.PlgVector)(unsafe.Pointer(p14)))
	_p15 := plugify.GetVectorDataString((*plugify.PlgVector)(unsafe.Pointer(p15)))
	ParamRefVectors(&_p1, &_p2, &_p3, &_p4, &_p5, &_p6, &_p7, &_p8, &_p9, &_p10, &_p11, &_p12, &_p13, &_p14, &_p15)
	plugify.AssignVectorBool((*plugify.PlgVector)(unsafe.Pointer(p1)), _p1)
	plugify.AssignVectorChar8((*plugify.PlgVector)(unsafe.Pointer(p2)), _p2)
	plugify.AssignVectorChar16((*plugify.PlgVector)(unsafe.Pointer(p3)), _p3)
	plugify.AssignVectorInt8((*plugify.PlgVector)(unsafe.Pointer(p4)), _p4)
	plugify.AssignVectorInt16((*plugify.PlgVector)(unsafe.Pointer(p5)), _p5)
	plugify.AssignVectorInt32((*plugify.PlgVector)(unsafe.Pointer(p6)), _p6)
	plugify.AssignVectorInt64((*plugify.PlgVector)(unsafe.Pointer(p7)), _p7)
	plugify.AssignVectorUInt8((*plugify.PlgVector)(unsafe.Pointer(p8)), _p8)
	plugify.AssignVectorUInt16((*plugify.PlgVector)(unsafe.Pointer(p9)), _p9)
	plugify.AssignVectorUInt32((*plugify.PlgVector)(unsafe.Pointer(p10)), _p10)
	plugify.AssignVectorUInt64((*plugify.PlgVector)(unsafe.Pointer(p11)), _p11)
	plugify.AssignVectorPointer((*plugify.PlgVector)(unsafe.Pointer(p12)), _p12)
	plugify.AssignVectorFloat((*plugify.PlgVector)(unsafe.Pointer(p13)), _p13)
	plugify.AssignVectorDouble((*plugify.PlgVector)(unsafe.Pointer(p14)), _p14)
	plugify.AssignVectorString((*plugify.PlgVector)(unsafe.Pointer(p15)), _p15)
}

//export __ParamAllPrimitives
func __ParamAllPrimitives(p1 bool, p2 int8, p3 uint16, p4 int8, p5 int16, p6 int32, p7 int64, p8 uint8, p9 uint16, p10 uint32, p11 uint64, p12 uintptr, p13 float32, p14 float64) int64 {
	__result := ParamAllPrimitives(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14)
	return __result
}

//export __ParamVariant
func __ParamVariant(p1 *C.Variant, p2 *C.Vector) {
	ParamVariant(plugify.GetVariantData((*plugify.PlgVariant)(unsafe.Pointer(p1))), plugify.GetVectorDataVariant((*plugify.PlgVector)(unsafe.Pointer(p2))))
}

//export __ParamEnum
func __ParamEnum(p1 int32, p2 *C.Vector) int32 {
	__result := ParamEnum(Example(p1), plugify.GetVectorDataInt32T[Example]((*plugify.PlgVector)(unsafe.Pointer(p2))))
	return __result
}

//export __ParamEnumRef
func __ParamEnumRef(p1 *int32, p2 *C.Vector) int32 {
	_p2 := plugify.GetVectorDataInt32T[Example]((*plugify.PlgVector)(unsafe.Pointer(p2)))
	__result := ParamEnumRef((*Example)(p1), &_p2)
	plugify.AssignVectorInt32((*plugify.PlgVector)(unsafe.Pointer(p2)), _p2)
	return __result
}

//export __ParamVariantRef
func __ParamVariantRef(p1 *C.Variant, p2 *C.Vector) {
	_p1 := plugify.GetVariantData((*plugify.PlgVariant)(unsafe.Pointer(p1)))
	_p2 := plugify.GetVectorDataVariant((*plugify.PlgVector)(unsafe.Pointer(p2)))
	ParamVariantRef(&_p1, &_p2)
	plugify.AssignVariant((*plugify.PlgVariant)(unsafe.Pointer(p1)), _p1)
	plugify.AssignVectorVariant((*plugify.PlgVector)(unsafe.Pointer(p2)), _p2)
}

//export __CallFuncVoid
func __CallFuncVoid(func_ unsafe.Pointer) {
	CallFuncVoid(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncVoid(nil))).(FuncVoid))
}

//export __CallFuncBool
func __CallFuncBool(func_ unsafe.Pointer) bool {
	__result := CallFuncBool(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncBool(nil))).(FuncBool))
	return __result
}

//export __CallFuncChar8
func __CallFuncChar8(func_ unsafe.Pointer) int8 {
	__result := CallFuncChar8(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncChar8(nil))).(FuncChar8))
	return __result
}

//export __CallFuncChar16
func __CallFuncChar16(func_ unsafe.Pointer) uint16 {
	__result := CallFuncChar16(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncChar16(nil))).(FuncChar16))
	return __result
}

//export __CallFuncInt8
func __CallFuncInt8(func_ unsafe.Pointer) int8 {
	__result := CallFuncInt8(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncInt8(nil))).(FuncInt8))
	return __result
}

//export __CallFuncInt16
func __CallFuncInt16(func_ unsafe.Pointer) int16 {
	__result := CallFuncInt16(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncInt16(nil))).(FuncInt16))
	return __result
}

//export __CallFuncInt32
func __CallFuncInt32(func_ unsafe.Pointer) int32 {
	__result := CallFuncInt32(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncInt32(nil))).(FuncInt32))
	return __result
}

//export __CallFuncInt64
func __CallFuncInt64(func_ unsafe.Pointer) int64 {
	__result := CallFuncInt64(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncInt64(nil))).(FuncInt64))
	return __result
}

//export __CallFuncUInt8
func __CallFuncUInt8(func_ unsafe.Pointer) uint8 {
	__result := CallFuncUInt8(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncUInt8(nil))).(FuncUInt8))
	return __result
}

//export __CallFuncUInt16
func __CallFuncUInt16(func_ unsafe.Pointer) uint16 {
	__result := CallFuncUInt16(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncUInt16(nil))).(FuncUInt16))
	return __result
}

//export __CallFuncUInt32
func __CallFuncUInt32(func_ unsafe.Pointer) uint32 {
	__result := CallFuncUInt32(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncUInt32(nil))).(FuncUInt32))
	return __result
}

//export __CallFuncUInt64
func __CallFuncUInt64(func_ unsafe.Pointer) uint64 {
	__result := CallFuncUInt64(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncUInt64(nil))).(FuncUInt64))
	return __result
}

//export __CallFuncPtr
func __CallFuncPtr(func_ unsafe.Pointer) uintptr {
	__result := CallFuncPtr(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncPtr(nil))).(FuncPtr))
	return __result
}

//export __CallFuncFloat
func __CallFuncFloat(func_ unsafe.Pointer) float32 {
	__result := CallFuncFloat(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncFloat(nil))).(FuncFloat))
	return __result
}

//export __CallFuncDouble
func __CallFuncDouble(func_ unsafe.Pointer) float64 {
	__result := CallFuncDouble(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncDouble(nil))).(FuncDouble))
	return __result
}

//export __CallFuncString
func __CallFuncString(func_ unsafe.Pointer) C.String {
	__result := CallFuncString(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncString(nil))).(FuncString))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFuncAny
func __CallFuncAny(func_ unsafe.Pointer) C.Variant {
	__result := CallFuncAny(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncAny(nil))).(FuncAny))
	__return := plugify.ConstructVariant(__result)
	return *(*C.Variant)(unsafe.Pointer(&__return))
}

//export __CallFuncFunction
func __CallFuncFunction(func_ unsafe.Pointer) uintptr {
	__result := CallFuncFunction(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncFunction(nil))).(FuncFunction))
	return __result
}

//export __CallFuncBoolVector
func __CallFuncBoolVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncBoolVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncBoolVector(nil))).(FuncBoolVector))
	__return := plugify.ConstructVectorBool(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncChar8Vector
func __CallFuncChar8Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncChar8Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncChar8Vector(nil))).(FuncChar8Vector))
	__return := plugify.ConstructVectorChar8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncChar16Vector
func __CallFuncChar16Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncChar16Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncChar16Vector(nil))).(FuncChar16Vector))
	__return := plugify.ConstructVectorChar16(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncInt8Vector
func __CallFuncInt8Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncInt8Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncInt8Vector(nil))).(FuncInt8Vector))
	__return := plugify.ConstructVectorInt8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncInt16Vector
func __CallFuncInt16Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncInt16Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncInt16Vector(nil))).(FuncInt16Vector))
	__return := plugify.ConstructVectorInt16(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncInt32Vector
func __CallFuncInt32Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncInt32Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncInt32Vector(nil))).(FuncInt32Vector))
	__return := plugify.ConstructVectorInt32(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncInt64Vector
func __CallFuncInt64Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncInt64Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncInt64Vector(nil))).(FuncInt64Vector))
	__return := plugify.ConstructVectorInt64(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncUInt8Vector
func __CallFuncUInt8Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncUInt8Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncUInt8Vector(nil))).(FuncUInt8Vector))
	__return := plugify.ConstructVectorUInt8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncUInt16Vector
func __CallFuncUInt16Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncUInt16Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncUInt16Vector(nil))).(FuncUInt16Vector))
	__return := plugify.ConstructVectorUInt16(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncUInt32Vector
func __CallFuncUInt32Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncUInt32Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncUInt32Vector(nil))).(FuncUInt32Vector))
	__return := plugify.ConstructVectorUInt32(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncUInt64Vector
func __CallFuncUInt64Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncUInt64Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncUInt64Vector(nil))).(FuncUInt64Vector))
	__return := plugify.ConstructVectorUInt64(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncPtrVector
func __CallFuncPtrVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncPtrVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncPtrVector(nil))).(FuncPtrVector))
	__return := plugify.ConstructVectorPointer(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncFloatVector
func __CallFuncFloatVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncFloatVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncFloatVector(nil))).(FuncFloatVector))
	__return := plugify.ConstructVectorFloat(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncDoubleVector
func __CallFuncDoubleVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncDoubleVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncDoubleVector(nil))).(FuncDoubleVector))
	__return := plugify.ConstructVectorDouble(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncStringVector
func __CallFuncStringVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncStringVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncStringVector(nil))).(FuncStringVector))
	__return := plugify.ConstructVectorString(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAnyVector
func __CallFuncAnyVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAnyVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncAnyVector(nil))).(FuncAnyVector))
	__return := plugify.ConstructVectorVariant(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncVec2Vector
func __CallFuncVec2Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncVec2Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncVec2Vector(nil))).(FuncVec2Vector))
	__return := plugify.ConstructVectorVector2(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncVec3Vector
func __CallFuncVec3Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncVec3Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncVec3Vector(nil))).(FuncVec3Vector))
	__return := plugify.ConstructVectorVector3(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncVec4Vector
func __CallFuncVec4Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncVec4Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncVec4Vector(nil))).(FuncVec4Vector))
	__return := plugify.ConstructVectorVector4(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncMat4x4Vector
func __CallFuncMat4x4Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncMat4x4Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncMat4x4Vector(nil))).(FuncMat4x4Vector))
	__return := plugify.ConstructVectorMatrix4x4(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncVec2
func __CallFuncVec2(func_ unsafe.Pointer) C.Vector2 {
	__result := CallFuncVec2(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncVec2(nil))).(FuncVec2))
	return *(*C.Vector2)(unsafe.Pointer(&__result))
}

//export __CallFuncVec3
func __CallFuncVec3(func_ unsafe.Pointer) C.Vector3 {
	__result := CallFuncVec3(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncVec3(nil))).(FuncVec3))
	return *(*C.Vector3)(unsafe.Pointer(&__result))
}

//export __CallFuncVec4
func __CallFuncVec4(func_ unsafe.Pointer) C.Vector4 {
	__result := CallFuncVec4(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncVec4(nil))).(FuncVec4))
	return *(*C.Vector4)(unsafe.Pointer(&__result))
}

//export __CallFuncMat4x4
func __CallFuncMat4x4(func_ unsafe.Pointer) C.Matrix4x4 {
	__result := CallFuncMat4x4(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncMat4x4(nil))).(FuncMat4x4))
	return *(*C.Matrix4x4)(unsafe.Pointer(&__result))
}

//export __CallFunc1
func __CallFunc1(func_ unsafe.Pointer) int32 {
	__result := CallFunc1(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func1(nil))).(Func1))
	return __result
}

//export __CallFunc2
func __CallFunc2(func_ unsafe.Pointer) int8 {
	__result := CallFunc2(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func2(nil))).(Func2))
	return __result
}

//export __CallFunc3
func __CallFunc3(func_ unsafe.Pointer) {
	CallFunc3(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func3(nil))).(Func3))
}

//export __CallFunc4
func __CallFunc4(func_ unsafe.Pointer) C.Vector4 {
	__result := CallFunc4(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func4(nil))).(Func4))
	return *(*C.Vector4)(unsafe.Pointer(&__result))
}

//export __CallFunc5
func __CallFunc5(func_ unsafe.Pointer) bool {
	__result := CallFunc5(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func5(nil))).(Func5))
	return __result
}

//export __CallFunc6
func __CallFunc6(func_ unsafe.Pointer) int64 {
	__result := CallFunc6(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func6(nil))).(Func6))
	return __result
}

//export __CallFunc7
func __CallFunc7(func_ unsafe.Pointer) float64 {
	__result := CallFunc7(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func7(nil))).(Func7))
	return __result
}

//export __CallFunc8
func __CallFunc8(func_ unsafe.Pointer) C.Matrix4x4 {
	__result := CallFunc8(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func8(nil))).(Func8))
	return *(*C.Matrix4x4)(unsafe.Pointer(&__result))
}

//export __CallFunc9
func __CallFunc9(func_ unsafe.Pointer) {
	CallFunc9(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func9(nil))).(Func9))
}

//export __CallFunc10
func __CallFunc10(func_ unsafe.Pointer) uint32 {
	__result := CallFunc10(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func10(nil))).(Func10))
	return __result
}

//export __CallFunc11
func __CallFunc11(func_ unsafe.Pointer) uintptr {
	__result := CallFunc11(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func11(nil))).(Func11))
	return __result
}

//export __CallFunc12
func __CallFunc12(func_ unsafe.Pointer) bool {
	__result := CallFunc12(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func12(nil))).(Func12))
	return __result
}

//export __CallFunc13
func __CallFunc13(func_ unsafe.Pointer) C.String {
	__result := CallFunc13(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func13(nil))).(Func13))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc14
func __CallFunc14(func_ unsafe.Pointer) C.Vector {
	__result := CallFunc14(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func14(nil))).(Func14))
	__return := plugify.ConstructVectorString(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFunc15
func __CallFunc15(func_ unsafe.Pointer) int16 {
	__result := CallFunc15(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func15(nil))).(Func15))
	return __result
}

//export __CallFunc16
func __CallFunc16(func_ unsafe.Pointer) uintptr {
	__result := CallFunc16(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func16(nil))).(Func16))
	return __result
}

//export __CallFunc17
func __CallFunc17(func_ unsafe.Pointer) C.String {
	__result := CallFunc17(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func17(nil))).(Func17))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc18
func __CallFunc18(func_ unsafe.Pointer) C.String {
	__result := CallFunc18(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func18(nil))).(Func18))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc19
func __CallFunc19(func_ unsafe.Pointer) C.String {
	__result := CallFunc19(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func19(nil))).(Func19))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc20
func __CallFunc20(func_ unsafe.Pointer) C.String {
	__result := CallFunc20(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func20(nil))).(Func20))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc21
func __CallFunc21(func_ unsafe.Pointer) C.String {
	__result := CallFunc21(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func21(nil))).(Func21))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc22
func __CallFunc22(func_ unsafe.Pointer) C.String {
	__result := CallFunc22(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func22(nil))).(Func22))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc23
func __CallFunc23(func_ unsafe.Pointer) C.String {
	__result := CallFunc23(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func23(nil))).(Func23))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc24
func __CallFunc24(func_ unsafe.Pointer) C.String {
	__result := CallFunc24(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func24(nil))).(Func24))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc25
func __CallFunc25(func_ unsafe.Pointer) C.String {
	__result := CallFunc25(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func25(nil))).(Func25))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc26
func __CallFunc26(func_ unsafe.Pointer) C.String {
	__result := CallFunc26(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func26(nil))).(Func26))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc27
func __CallFunc27(func_ unsafe.Pointer) C.String {
	__result := CallFunc27(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func27(nil))).(Func27))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc28
func __CallFunc28(func_ unsafe.Pointer) C.String {
	__result := CallFunc28(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func28(nil))).(Func28))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc29
func __CallFunc29(func_ unsafe.Pointer) C.String {
	__result := CallFunc29(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func29(nil))).(Func29))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc30
func __CallFunc30(func_ unsafe.Pointer) C.String {
	__result := CallFunc30(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func30(nil))).(Func30))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc31
func __CallFunc31(func_ unsafe.Pointer) C.String {
	__result := CallFunc31(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func31(nil))).(Func31))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc32
func __CallFunc32(func_ unsafe.Pointer) C.String {
	__result := CallFunc32(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func32(nil))).(Func32))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc33
func __CallFunc33(func_ unsafe.Pointer) C.String {
	__result := CallFunc33(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(Func33(nil))).(Func33))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFuncEnum
func __CallFuncEnum(func_ unsafe.Pointer) C.String {
	__result := CallFuncEnum(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(FuncEnum(nil))).(FuncEnum))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCall
func __ReverseCall(test *C.String) {
	ReverseCall(plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(test))))
}
