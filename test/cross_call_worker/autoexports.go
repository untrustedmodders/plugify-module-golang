package main

// #include "autoexports.h"
import "C"
import (
	"github.com/untrustedmodders/go-plugify"
	"plugify-plugin/cross_call_master"
	"reflect"
	"unsafe"
)

var _ = reflect.TypeOf(0)
var _ = unsafe.Sizeof(0)
var _ = plugify.Plugin.Loaded

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
	__return := plugify.ConstructVectorInt8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __NoParamReturnArrayChar16
func __NoParamReturnArrayChar16() C.Vector {
	__result := NoParamReturnArrayChar16()
	__return := plugify.ConstructVectorUInt16(__result)
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
	_p2 := plugify.GetVectorDataInt8((*plugify.PlgVector)(unsafe.Pointer(p2)))
	_p3 := plugify.GetVectorDataUInt16((*plugify.PlgVector)(unsafe.Pointer(p3)))
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
	plugify.AssignVectorInt8((*plugify.PlgVector)(unsafe.Pointer(p2)), _p2)
	plugify.AssignVectorUInt16((*plugify.PlgVector)(unsafe.Pointer(p3)), _p3)
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

//export __ParamAllAliases
func __ParamAllAliases(aBool bool, aChar8 int8, aChar16 uint16, aInt8 int8, aInt16 int16, aInt32 int32, aInt64 int64, aPtr uintptr, aFloat float32, aDouble float64, aString *C.String, aAny *C.Variant, aVec2 *C.Vector2, aVec3 *C.Vector3, aVec4 *C.Vector4, aMat4x4 *C.Matrix4x4, aBoolVec *C.Vector, aChar8Vec *C.Vector, aChar16Vec *C.Vector, aInt8Vec *C.Vector, aInt16Vec *C.Vector, aInt32Vec *C.Vector, aInt64Vec *C.Vector, aPtrVec *C.Vector, aFloatVec *C.Vector, aDoubleVec *C.Vector, aStringVec *C.Vector, aAnyVec *C.Vector, aVec2Vec *C.Vector, aVec3Vec *C.Vector, aVec4Vec *C.Vector) int32 {
	__result := ParamAllAliases(aBool, aChar8, aChar16, aInt8, aInt16, aInt32, aInt64, aPtr, aFloat, aDouble, plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(aString))), plugify.GetVariantData((*plugify.PlgVariant)(unsafe.Pointer(aAny))), *(*plugify.Vector2)(unsafe.Pointer(aVec2)), *(*plugify.Vector3)(unsafe.Pointer(aVec3)), *(*plugify.Vector4)(unsafe.Pointer(aVec4)), *(*plugify.Matrix4x4)(unsafe.Pointer(aMat4x4)), plugify.GetVectorDataBool((*plugify.PlgVector)(unsafe.Pointer(aBoolVec))), plugify.GetVectorDataInt8((*plugify.PlgVector)(unsafe.Pointer(aChar8Vec))), plugify.GetVectorDataUInt16((*plugify.PlgVector)(unsafe.Pointer(aChar16Vec))), plugify.GetVectorDataInt8((*plugify.PlgVector)(unsafe.Pointer(aInt8Vec))), plugify.GetVectorDataInt16((*plugify.PlgVector)(unsafe.Pointer(aInt16Vec))), plugify.GetVectorDataInt32((*plugify.PlgVector)(unsafe.Pointer(aInt32Vec))), plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(aInt64Vec))), plugify.GetVectorDataPointer((*plugify.PlgVector)(unsafe.Pointer(aPtrVec))), plugify.GetVectorDataFloat((*plugify.PlgVector)(unsafe.Pointer(aFloatVec))), plugify.GetVectorDataDouble((*plugify.PlgVector)(unsafe.Pointer(aDoubleVec))), plugify.GetVectorDataString((*plugify.PlgVector)(unsafe.Pointer(aStringVec))), plugify.GetVectorDataVariant((*plugify.PlgVector)(unsafe.Pointer(aAnyVec))), plugify.GetVectorDataVector2((*plugify.PlgVector)(unsafe.Pointer(aVec2Vec))), plugify.GetVectorDataVector3((*plugify.PlgVector)(unsafe.Pointer(aVec3Vec))), plugify.GetVectorDataVector4((*plugify.PlgVector)(unsafe.Pointer(aVec4Vec))))
	return __result
}

//export __ParamAllRefAliases
func __ParamAllRefAliases(aBool *bool, aChar8 *int8, aChar16 *uint16, aInt8 *int8, aInt16 *int16, aInt32 *int32, aInt64 *int64, aPtr *uintptr, aFloat *float32, aDouble *float64, aString *C.String, aAny *C.Variant, aVec2 *C.Vector2, aVec3 *C.Vector3, aVec4 *C.Vector4, aMat4x4 *C.Matrix4x4, aBoolVec *C.Vector, aChar8Vec *C.Vector, aChar16Vec *C.Vector, aInt8Vec *C.Vector, aInt16Vec *C.Vector, aInt32Vec *C.Vector, aInt64Vec *C.Vector, aPtrVec *C.Vector, aFloatVec *C.Vector, aDoubleVec *C.Vector, aStringVec *C.Vector, aAnyVec *C.Vector, aVec2Vec *C.Vector, aVec3Vec *C.Vector, aVec4Vec *C.Vector) int64 {
	_aString := plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(aString)))
	_aAny := plugify.GetVariantData((*plugify.PlgVariant)(unsafe.Pointer(aAny)))
	_aBoolVec := plugify.GetVectorDataBool((*plugify.PlgVector)(unsafe.Pointer(aBoolVec)))
	_aChar8Vec := plugify.GetVectorDataInt8((*plugify.PlgVector)(unsafe.Pointer(aChar8Vec)))
	_aChar16Vec := plugify.GetVectorDataUInt16((*plugify.PlgVector)(unsafe.Pointer(aChar16Vec)))
	_aInt8Vec := plugify.GetVectorDataInt8((*plugify.PlgVector)(unsafe.Pointer(aInt8Vec)))
	_aInt16Vec := plugify.GetVectorDataInt16((*plugify.PlgVector)(unsafe.Pointer(aInt16Vec)))
	_aInt32Vec := plugify.GetVectorDataInt32((*plugify.PlgVector)(unsafe.Pointer(aInt32Vec)))
	_aInt64Vec := plugify.GetVectorDataInt64((*plugify.PlgVector)(unsafe.Pointer(aInt64Vec)))
	_aPtrVec := plugify.GetVectorDataPointer((*plugify.PlgVector)(unsafe.Pointer(aPtrVec)))
	_aFloatVec := plugify.GetVectorDataFloat((*plugify.PlgVector)(unsafe.Pointer(aFloatVec)))
	_aDoubleVec := plugify.GetVectorDataDouble((*plugify.PlgVector)(unsafe.Pointer(aDoubleVec)))
	_aStringVec := plugify.GetVectorDataString((*plugify.PlgVector)(unsafe.Pointer(aStringVec)))
	_aAnyVec := plugify.GetVectorDataVariant((*plugify.PlgVector)(unsafe.Pointer(aAnyVec)))
	_aVec2Vec := plugify.GetVectorDataVector2((*plugify.PlgVector)(unsafe.Pointer(aVec2Vec)))
	_aVec3Vec := plugify.GetVectorDataVector3((*plugify.PlgVector)(unsafe.Pointer(aVec3Vec)))
	_aVec4Vec := plugify.GetVectorDataVector4((*plugify.PlgVector)(unsafe.Pointer(aVec4Vec)))
	__result := ParamAllRefAliases(aBool, aChar8, aChar16, aInt8, aInt16, aInt32, aInt64, aPtr, aFloat, aDouble, &_aString, &_aAny, (*plugify.Vector2)(unsafe.Pointer(aVec2)), (*plugify.Vector3)(unsafe.Pointer(aVec3)), (*plugify.Vector4)(unsafe.Pointer(aVec4)), (*plugify.Matrix4x4)(unsafe.Pointer(aMat4x4)), &_aBoolVec, &_aChar8Vec, &_aChar16Vec, &_aInt8Vec, &_aInt16Vec, &_aInt32Vec, &_aInt64Vec, &_aPtrVec, &_aFloatVec, &_aDoubleVec, &_aStringVec, &_aAnyVec, &_aVec2Vec, &_aVec3Vec, &_aVec4Vec)
	plugify.AssignString((*plugify.PlgString)(unsafe.Pointer(aString)), _aString)
	plugify.AssignVariant((*plugify.PlgVariant)(unsafe.Pointer(aAny)), _aAny)
	plugify.AssignVectorBool((*plugify.PlgVector)(unsafe.Pointer(aBoolVec)), _aBoolVec)
	plugify.AssignVectorInt8((*plugify.PlgVector)(unsafe.Pointer(aChar8Vec)), _aChar8Vec)
	plugify.AssignVectorUInt16((*plugify.PlgVector)(unsafe.Pointer(aChar16Vec)), _aChar16Vec)
	plugify.AssignVectorInt8((*plugify.PlgVector)(unsafe.Pointer(aInt8Vec)), _aInt8Vec)
	plugify.AssignVectorInt16((*plugify.PlgVector)(unsafe.Pointer(aInt16Vec)), _aInt16Vec)
	plugify.AssignVectorInt32((*plugify.PlgVector)(unsafe.Pointer(aInt32Vec)), _aInt32Vec)
	plugify.AssignVectorInt64((*plugify.PlgVector)(unsafe.Pointer(aInt64Vec)), _aInt64Vec)
	plugify.AssignVectorPointer((*plugify.PlgVector)(unsafe.Pointer(aPtrVec)), _aPtrVec)
	plugify.AssignVectorFloat((*plugify.PlgVector)(unsafe.Pointer(aFloatVec)), _aFloatVec)
	plugify.AssignVectorDouble((*plugify.PlgVector)(unsafe.Pointer(aDoubleVec)), _aDoubleVec)
	plugify.AssignVectorString((*plugify.PlgVector)(unsafe.Pointer(aStringVec)), _aStringVec)
	plugify.AssignVectorVariant((*plugify.PlgVector)(unsafe.Pointer(aAnyVec)), _aAnyVec)
	plugify.AssignVectorVector2((*plugify.PlgVector)(unsafe.Pointer(aVec2Vec)), _aVec2Vec)
	plugify.AssignVectorVector3((*plugify.PlgVector)(unsafe.Pointer(aVec3Vec)), _aVec3Vec)
	plugify.AssignVectorVector4((*plugify.PlgVector)(unsafe.Pointer(aVec4Vec)), _aVec4Vec)
	return __result
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

//export __ParamVariant
func __ParamVariant(p1 *C.Variant, p2 *C.Vector) {
	ParamVariant(plugify.GetVariantData((*plugify.PlgVariant)(unsafe.Pointer(p1))), plugify.GetVectorDataVariant((*plugify.PlgVector)(unsafe.Pointer(p2))))
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
	CallFuncVoid(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncVoid(nil))).(cross_call_master.FuncVoid))
}

//export __CallFuncBool
func __CallFuncBool(func_ unsafe.Pointer) bool {
	__result := CallFuncBool(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncBool(nil))).(cross_call_master.FuncBool))
	return __result
}

//export __CallFuncChar8
func __CallFuncChar8(func_ unsafe.Pointer) int8 {
	__result := CallFuncChar8(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncChar8(nil))).(cross_call_master.FuncChar8))
	return __result
}

//export __CallFuncChar16
func __CallFuncChar16(func_ unsafe.Pointer) uint16 {
	__result := CallFuncChar16(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncChar16(nil))).(cross_call_master.FuncChar16))
	return __result
}

//export __CallFuncInt8
func __CallFuncInt8(func_ unsafe.Pointer) int8 {
	__result := CallFuncInt8(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncInt8(nil))).(cross_call_master.FuncInt8))
	return __result
}

//export __CallFuncInt16
func __CallFuncInt16(func_ unsafe.Pointer) int16 {
	__result := CallFuncInt16(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncInt16(nil))).(cross_call_master.FuncInt16))
	return __result
}

//export __CallFuncInt32
func __CallFuncInt32(func_ unsafe.Pointer) int32 {
	__result := CallFuncInt32(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncInt32(nil))).(cross_call_master.FuncInt32))
	return __result
}

//export __CallFuncInt64
func __CallFuncInt64(func_ unsafe.Pointer) int64 {
	__result := CallFuncInt64(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncInt64(nil))).(cross_call_master.FuncInt64))
	return __result
}

//export __CallFuncUInt8
func __CallFuncUInt8(func_ unsafe.Pointer) uint8 {
	__result := CallFuncUInt8(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncUInt8(nil))).(cross_call_master.FuncUInt8))
	return __result
}

//export __CallFuncUInt16
func __CallFuncUInt16(func_ unsafe.Pointer) uint16 {
	__result := CallFuncUInt16(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncUInt16(nil))).(cross_call_master.FuncUInt16))
	return __result
}

//export __CallFuncUInt32
func __CallFuncUInt32(func_ unsafe.Pointer) uint32 {
	__result := CallFuncUInt32(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncUInt32(nil))).(cross_call_master.FuncUInt32))
	return __result
}

//export __CallFuncUInt64
func __CallFuncUInt64(func_ unsafe.Pointer) uint64 {
	__result := CallFuncUInt64(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncUInt64(nil))).(cross_call_master.FuncUInt64))
	return __result
}

//export __CallFuncPtr
func __CallFuncPtr(func_ unsafe.Pointer) uintptr {
	__result := CallFuncPtr(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncPtr(nil))).(cross_call_master.FuncPtr))
	return __result
}

//export __CallFuncFloat
func __CallFuncFloat(func_ unsafe.Pointer) float32 {
	__result := CallFuncFloat(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncFloat(nil))).(cross_call_master.FuncFloat))
	return __result
}

//export __CallFuncDouble
func __CallFuncDouble(func_ unsafe.Pointer) float64 {
	__result := CallFuncDouble(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncDouble(nil))).(cross_call_master.FuncDouble))
	return __result
}

//export __CallFuncFunction
func __CallFuncFunction(func_ unsafe.Pointer) unsafe.Pointer {
	__result := CallFuncFunction(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncFunction(nil))).(cross_call_master.FuncFunction))
	return plugify.GetFunctionPointerForDelegate(__result)
}

//export __CallFuncString
func __CallFuncString(func_ unsafe.Pointer) C.String {
	__result := CallFuncString(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncString(nil))).(cross_call_master.FuncString))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFuncAny
func __CallFuncAny(func_ unsafe.Pointer) C.Variant {
	__result := CallFuncAny(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAny(nil))).(cross_call_master.FuncAny))
	__return := plugify.ConstructVariant(__result)
	return *(*C.Variant)(unsafe.Pointer(&__return))
}

//export __CallFuncBoolVector
func __CallFuncBoolVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncBoolVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncBoolVector(nil))).(cross_call_master.FuncBoolVector))
	__return := plugify.ConstructVectorBool(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncChar8Vector
func __CallFuncChar8Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncChar8Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncChar8Vector(nil))).(cross_call_master.FuncChar8Vector))
	__return := plugify.ConstructVectorInt8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncChar16Vector
func __CallFuncChar16Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncChar16Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncChar16Vector(nil))).(cross_call_master.FuncChar16Vector))
	__return := plugify.ConstructVectorUInt16(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncInt8Vector
func __CallFuncInt8Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncInt8Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncInt8Vector(nil))).(cross_call_master.FuncInt8Vector))
	__return := plugify.ConstructVectorInt8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncInt16Vector
func __CallFuncInt16Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncInt16Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncInt16Vector(nil))).(cross_call_master.FuncInt16Vector))
	__return := plugify.ConstructVectorInt16(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncInt32Vector
func __CallFuncInt32Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncInt32Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncInt32Vector(nil))).(cross_call_master.FuncInt32Vector))
	__return := plugify.ConstructVectorInt32(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncInt64Vector
func __CallFuncInt64Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncInt64Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncInt64Vector(nil))).(cross_call_master.FuncInt64Vector))
	__return := plugify.ConstructVectorInt64(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncUInt8Vector
func __CallFuncUInt8Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncUInt8Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncUInt8Vector(nil))).(cross_call_master.FuncUInt8Vector))
	__return := plugify.ConstructVectorUInt8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncUInt16Vector
func __CallFuncUInt16Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncUInt16Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncUInt16Vector(nil))).(cross_call_master.FuncUInt16Vector))
	__return := plugify.ConstructVectorUInt16(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncUInt32Vector
func __CallFuncUInt32Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncUInt32Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncUInt32Vector(nil))).(cross_call_master.FuncUInt32Vector))
	__return := plugify.ConstructVectorUInt32(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncUInt64Vector
func __CallFuncUInt64Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncUInt64Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncUInt64Vector(nil))).(cross_call_master.FuncUInt64Vector))
	__return := plugify.ConstructVectorUInt64(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncPtrVector
func __CallFuncPtrVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncPtrVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncPtrVector(nil))).(cross_call_master.FuncPtrVector))
	__return := plugify.ConstructVectorPointer(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncFloatVector
func __CallFuncFloatVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncFloatVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncFloatVector(nil))).(cross_call_master.FuncFloatVector))
	__return := plugify.ConstructVectorFloat(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncDoubleVector
func __CallFuncDoubleVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncDoubleVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncDoubleVector(nil))).(cross_call_master.FuncDoubleVector))
	__return := plugify.ConstructVectorDouble(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncStringVector
func __CallFuncStringVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncStringVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncStringVector(nil))).(cross_call_master.FuncStringVector))
	__return := plugify.ConstructVectorString(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAnyVector
func __CallFuncAnyVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAnyVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAnyVector(nil))).(cross_call_master.FuncAnyVector))
	__return := plugify.ConstructVectorVariant(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncVec2Vector
func __CallFuncVec2Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncVec2Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncVec2Vector(nil))).(cross_call_master.FuncVec2Vector))
	__return := plugify.ConstructVectorVector2(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncVec3Vector
func __CallFuncVec3Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncVec3Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncVec3Vector(nil))).(cross_call_master.FuncVec3Vector))
	__return := plugify.ConstructVectorVector3(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncVec4Vector
func __CallFuncVec4Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncVec4Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncVec4Vector(nil))).(cross_call_master.FuncVec4Vector))
	__return := plugify.ConstructVectorVector4(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncMat4x4Vector
func __CallFuncMat4x4Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncMat4x4Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncMat4x4Vector(nil))).(cross_call_master.FuncMat4x4Vector))
	__return := plugify.ConstructVectorMatrix4x4(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncVec2
func __CallFuncVec2(func_ unsafe.Pointer) C.Vector2 {
	__result := CallFuncVec2(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncVec2(nil))).(cross_call_master.FuncVec2))
	return *(*C.Vector2)(unsafe.Pointer(&__result))
}

//export __CallFuncVec3
func __CallFuncVec3(func_ unsafe.Pointer) C.Vector3 {
	__result := CallFuncVec3(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncVec3(nil))).(cross_call_master.FuncVec3))
	return *(*C.Vector3)(unsafe.Pointer(&__result))
}

//export __CallFuncVec4
func __CallFuncVec4(func_ unsafe.Pointer) C.Vector4 {
	__result := CallFuncVec4(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncVec4(nil))).(cross_call_master.FuncVec4))
	return *(*C.Vector4)(unsafe.Pointer(&__result))
}

//export __CallFuncMat4x4
func __CallFuncMat4x4(func_ unsafe.Pointer) C.Matrix4x4 {
	__result := CallFuncMat4x4(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncMat4x4(nil))).(cross_call_master.FuncMat4x4))
	return *(*C.Matrix4x4)(unsafe.Pointer(&__result))
}

//export __CallFunc1
func __CallFunc1(func_ unsafe.Pointer) int32 {
	__result := CallFunc1(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func1(nil))).(cross_call_master.Func1))
	return __result
}

//export __CallFunc2
func __CallFunc2(func_ unsafe.Pointer) int8 {
	__result := CallFunc2(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func2(nil))).(cross_call_master.Func2))
	return __result
}

//export __CallFunc3
func __CallFunc3(func_ unsafe.Pointer) {
	CallFunc3(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func3(nil))).(cross_call_master.Func3))
}

//export __CallFunc4
func __CallFunc4(func_ unsafe.Pointer) C.Vector4 {
	__result := CallFunc4(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func4(nil))).(cross_call_master.Func4))
	return *(*C.Vector4)(unsafe.Pointer(&__result))
}

//export __CallFunc5
func __CallFunc5(func_ unsafe.Pointer) bool {
	__result := CallFunc5(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func5(nil))).(cross_call_master.Func5))
	return __result
}

//export __CallFunc6
func __CallFunc6(func_ unsafe.Pointer) int64 {
	__result := CallFunc6(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func6(nil))).(cross_call_master.Func6))
	return __result
}

//export __CallFunc7
func __CallFunc7(func_ unsafe.Pointer) float64 {
	__result := CallFunc7(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func7(nil))).(cross_call_master.Func7))
	return __result
}

//export __CallFunc8
func __CallFunc8(func_ unsafe.Pointer) C.Matrix4x4 {
	__result := CallFunc8(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func8(nil))).(cross_call_master.Func8))
	return *(*C.Matrix4x4)(unsafe.Pointer(&__result))
}

//export __CallFunc9
func __CallFunc9(func_ unsafe.Pointer) {
	CallFunc9(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func9(nil))).(cross_call_master.Func9))
}

//export __CallFunc10
func __CallFunc10(func_ unsafe.Pointer) uint32 {
	__result := CallFunc10(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func10(nil))).(cross_call_master.Func10))
	return __result
}

//export __CallFunc11
func __CallFunc11(func_ unsafe.Pointer) uintptr {
	__result := CallFunc11(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func11(nil))).(cross_call_master.Func11))
	return __result
}

//export __CallFunc12
func __CallFunc12(func_ unsafe.Pointer) bool {
	__result := CallFunc12(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func12(nil))).(cross_call_master.Func12))
	return __result
}

//export __CallFunc13
func __CallFunc13(func_ unsafe.Pointer) C.String {
	__result := CallFunc13(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func13(nil))).(cross_call_master.Func13))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc14
func __CallFunc14(func_ unsafe.Pointer) C.Vector {
	__result := CallFunc14(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func14(nil))).(cross_call_master.Func14))
	__return := plugify.ConstructVectorString(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFunc15
func __CallFunc15(func_ unsafe.Pointer) int16 {
	__result := CallFunc15(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func15(nil))).(cross_call_master.Func15))
	return __result
}

//export __CallFunc16
func __CallFunc16(func_ unsafe.Pointer) uintptr {
	__result := CallFunc16(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func16(nil))).(cross_call_master.Func16))
	return __result
}

//export __CallFunc17
func __CallFunc17(func_ unsafe.Pointer) C.String {
	__result := CallFunc17(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func17(nil))).(cross_call_master.Func17))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc18
func __CallFunc18(func_ unsafe.Pointer) C.String {
	__result := CallFunc18(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func18(nil))).(cross_call_master.Func18))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc19
func __CallFunc19(func_ unsafe.Pointer) C.String {
	__result := CallFunc19(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func19(nil))).(cross_call_master.Func19))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc20
func __CallFunc20(func_ unsafe.Pointer) C.String {
	__result := CallFunc20(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func20(nil))).(cross_call_master.Func20))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc21
func __CallFunc21(func_ unsafe.Pointer) C.String {
	__result := CallFunc21(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func21(nil))).(cross_call_master.Func21))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc22
func __CallFunc22(func_ unsafe.Pointer) C.String {
	__result := CallFunc22(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func22(nil))).(cross_call_master.Func22))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc23
func __CallFunc23(func_ unsafe.Pointer) C.String {
	__result := CallFunc23(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func23(nil))).(cross_call_master.Func23))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc24
func __CallFunc24(func_ unsafe.Pointer) C.String {
	__result := CallFunc24(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func24(nil))).(cross_call_master.Func24))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc25
func __CallFunc25(func_ unsafe.Pointer) C.String {
	__result := CallFunc25(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func25(nil))).(cross_call_master.Func25))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc26
func __CallFunc26(func_ unsafe.Pointer) C.String {
	__result := CallFunc26(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func26(nil))).(cross_call_master.Func26))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc27
func __CallFunc27(func_ unsafe.Pointer) C.String {
	__result := CallFunc27(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func27(nil))).(cross_call_master.Func27))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc28
func __CallFunc28(func_ unsafe.Pointer) C.String {
	__result := CallFunc28(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func28(nil))).(cross_call_master.Func28))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc29
func __CallFunc29(func_ unsafe.Pointer) C.String {
	__result := CallFunc29(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func29(nil))).(cross_call_master.Func29))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc30
func __CallFunc30(func_ unsafe.Pointer) C.String {
	__result := CallFunc30(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func30(nil))).(cross_call_master.Func30))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc31
func __CallFunc31(func_ unsafe.Pointer) C.String {
	__result := CallFunc31(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func31(nil))).(cross_call_master.Func31))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc32
func __CallFunc32(func_ unsafe.Pointer) C.String {
	__result := CallFunc32(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func32(nil))).(cross_call_master.Func32))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFunc33
func __CallFunc33(func_ unsafe.Pointer) C.String {
	__result := CallFunc33(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.Func33(nil))).(cross_call_master.Func33))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFuncEnum
func __CallFuncEnum(func_ unsafe.Pointer) C.String {
	__result := CallFuncEnum(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncEnum(nil))).(cross_call_master.FuncEnum))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasBool
func __CallFuncAliasBool(func_ unsafe.Pointer) bool {
	__result := CallFuncAliasBool(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasBool(nil))).(cross_call_master.FuncAliasBool))
	return __result
}

//export __CallFuncAliasChar8
func __CallFuncAliasChar8(func_ unsafe.Pointer) int8 {
	__result := CallFuncAliasChar8(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasChar8(nil))).(cross_call_master.FuncAliasChar8))
	return __result
}

//export __CallFuncAliasChar16
func __CallFuncAliasChar16(func_ unsafe.Pointer) uint16 {
	__result := CallFuncAliasChar16(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasChar16(nil))).(cross_call_master.FuncAliasChar16))
	return __result
}

//export __CallFuncAliasInt8
func __CallFuncAliasInt8(func_ unsafe.Pointer) int8 {
	__result := CallFuncAliasInt8(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasInt8(nil))).(cross_call_master.FuncAliasInt8))
	return __result
}

//export __CallFuncAliasInt16
func __CallFuncAliasInt16(func_ unsafe.Pointer) int16 {
	__result := CallFuncAliasInt16(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasInt16(nil))).(cross_call_master.FuncAliasInt16))
	return __result
}

//export __CallFuncAliasInt32
func __CallFuncAliasInt32(func_ unsafe.Pointer) int32 {
	__result := CallFuncAliasInt32(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasInt32(nil))).(cross_call_master.FuncAliasInt32))
	return __result
}

//export __CallFuncAliasInt64
func __CallFuncAliasInt64(func_ unsafe.Pointer) int64 {
	__result := CallFuncAliasInt64(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasInt64(nil))).(cross_call_master.FuncAliasInt64))
	return __result
}

//export __CallFuncAliasUInt8
func __CallFuncAliasUInt8(func_ unsafe.Pointer) uint8 {
	__result := CallFuncAliasUInt8(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasUInt8(nil))).(cross_call_master.FuncAliasUInt8))
	return __result
}

//export __CallFuncAliasUInt16
func __CallFuncAliasUInt16(func_ unsafe.Pointer) uint16 {
	__result := CallFuncAliasUInt16(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasUInt16(nil))).(cross_call_master.FuncAliasUInt16))
	return __result
}

//export __CallFuncAliasUInt32
func __CallFuncAliasUInt32(func_ unsafe.Pointer) uint32 {
	__result := CallFuncAliasUInt32(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasUInt32(nil))).(cross_call_master.FuncAliasUInt32))
	return __result
}

//export __CallFuncAliasUInt64
func __CallFuncAliasUInt64(func_ unsafe.Pointer) uint64 {
	__result := CallFuncAliasUInt64(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasUInt64(nil))).(cross_call_master.FuncAliasUInt64))
	return __result
}

//export __CallFuncAliasPtr
func __CallFuncAliasPtr(func_ unsafe.Pointer) uintptr {
	__result := CallFuncAliasPtr(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasPtr(nil))).(cross_call_master.FuncAliasPtr))
	return __result
}

//export __CallFuncAliasFloat
func __CallFuncAliasFloat(func_ unsafe.Pointer) float32 {
	__result := CallFuncAliasFloat(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasFloat(nil))).(cross_call_master.FuncAliasFloat))
	return __result
}

//export __CallFuncAliasDouble
func __CallFuncAliasDouble(func_ unsafe.Pointer) float64 {
	__result := CallFuncAliasDouble(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasDouble(nil))).(cross_call_master.FuncAliasDouble))
	return __result
}

//export __CallFuncAliasFunction
func __CallFuncAliasFunction(func_ unsafe.Pointer) uintptr {
	__result := CallFuncAliasFunction(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasFunction(nil))).(cross_call_master.FuncAliasFunction))
	return __result
}

//export __CallFuncAliasString
func __CallFuncAliasString(func_ unsafe.Pointer) C.String {
	__result := CallFuncAliasString(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasString(nil))).(cross_call_master.FuncAliasString))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasAny
func __CallFuncAliasAny(func_ unsafe.Pointer) C.Variant {
	__result := CallFuncAliasAny(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasAny(nil))).(cross_call_master.FuncAliasAny))
	__return := plugify.ConstructVariant(__result)
	return *(*C.Variant)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasBoolVector
func __CallFuncAliasBoolVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasBoolVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasBoolVector(nil))).(cross_call_master.FuncAliasBoolVector))
	__return := plugify.ConstructVectorBool(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasChar8Vector
func __CallFuncAliasChar8Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasChar8Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasChar8Vector(nil))).(cross_call_master.FuncAliasChar8Vector))
	__return := plugify.ConstructVectorInt8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasChar16Vector
func __CallFuncAliasChar16Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasChar16Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasChar16Vector(nil))).(cross_call_master.FuncAliasChar16Vector))
	__return := plugify.ConstructVectorUInt16(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasInt8Vector
func __CallFuncAliasInt8Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasInt8Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasInt8Vector(nil))).(cross_call_master.FuncAliasInt8Vector))
	__return := plugify.ConstructVectorInt8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasInt16Vector
func __CallFuncAliasInt16Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasInt16Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasInt16Vector(nil))).(cross_call_master.FuncAliasInt16Vector))
	__return := plugify.ConstructVectorInt16(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasInt32Vector
func __CallFuncAliasInt32Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasInt32Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasInt32Vector(nil))).(cross_call_master.FuncAliasInt32Vector))
	__return := plugify.ConstructVectorInt32(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasInt64Vector
func __CallFuncAliasInt64Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasInt64Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasInt64Vector(nil))).(cross_call_master.FuncAliasInt64Vector))
	__return := plugify.ConstructVectorInt64(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasUInt8Vector
func __CallFuncAliasUInt8Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasUInt8Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasUInt8Vector(nil))).(cross_call_master.FuncAliasUInt8Vector))
	__return := plugify.ConstructVectorUInt8(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasUInt16Vector
func __CallFuncAliasUInt16Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasUInt16Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasUInt16Vector(nil))).(cross_call_master.FuncAliasUInt16Vector))
	__return := plugify.ConstructVectorUInt16(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasUInt32Vector
func __CallFuncAliasUInt32Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasUInt32Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasUInt32Vector(nil))).(cross_call_master.FuncAliasUInt32Vector))
	__return := plugify.ConstructVectorUInt32(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasUInt64Vector
func __CallFuncAliasUInt64Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasUInt64Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasUInt64Vector(nil))).(cross_call_master.FuncAliasUInt64Vector))
	__return := plugify.ConstructVectorUInt64(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasPtrVector
func __CallFuncAliasPtrVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasPtrVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasPtrVector(nil))).(cross_call_master.FuncAliasPtrVector))
	__return := plugify.ConstructVectorPointer(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasFloatVector
func __CallFuncAliasFloatVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasFloatVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasFloatVector(nil))).(cross_call_master.FuncAliasFloatVector))
	__return := plugify.ConstructVectorFloat(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasDoubleVector
func __CallFuncAliasDoubleVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasDoubleVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasDoubleVector(nil))).(cross_call_master.FuncAliasDoubleVector))
	__return := plugify.ConstructVectorDouble(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasStringVector
func __CallFuncAliasStringVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasStringVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasStringVector(nil))).(cross_call_master.FuncAliasStringVector))
	__return := plugify.ConstructVectorString(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasAnyVector
func __CallFuncAliasAnyVector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasAnyVector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasAnyVector(nil))).(cross_call_master.FuncAliasAnyVector))
	__return := plugify.ConstructVectorVariant(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasVec2Vector
func __CallFuncAliasVec2Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasVec2Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasVec2Vector(nil))).(cross_call_master.FuncAliasVec2Vector))
	__return := plugify.ConstructVectorVector2(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasVec3Vector
func __CallFuncAliasVec3Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasVec3Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasVec3Vector(nil))).(cross_call_master.FuncAliasVec3Vector))
	__return := plugify.ConstructVectorVector3(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasVec4Vector
func __CallFuncAliasVec4Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasVec4Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasVec4Vector(nil))).(cross_call_master.FuncAliasVec4Vector))
	__return := plugify.ConstructVectorVector4(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasMat4x4Vector
func __CallFuncAliasMat4x4Vector(func_ unsafe.Pointer) C.Vector {
	__result := CallFuncAliasMat4x4Vector(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasMat4x4Vector(nil))).(cross_call_master.FuncAliasMat4x4Vector))
	__return := plugify.ConstructVectorMatrix4x4(__result)
	return *(*C.Vector)(unsafe.Pointer(&__return))
}

//export __CallFuncAliasVec2
func __CallFuncAliasVec2(func_ unsafe.Pointer) C.Vector2 {
	__result := CallFuncAliasVec2(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasVec2(nil))).(cross_call_master.FuncAliasVec2))
	return *(*C.Vector2)(unsafe.Pointer(&__result))
}

//export __CallFuncAliasVec3
func __CallFuncAliasVec3(func_ unsafe.Pointer) C.Vector3 {
	__result := CallFuncAliasVec3(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasVec3(nil))).(cross_call_master.FuncAliasVec3))
	return *(*C.Vector3)(unsafe.Pointer(&__result))
}

//export __CallFuncAliasVec4
func __CallFuncAliasVec4(func_ unsafe.Pointer) C.Vector4 {
	__result := CallFuncAliasVec4(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasVec4(nil))).(cross_call_master.FuncAliasVec4))
	return *(*C.Vector4)(unsafe.Pointer(&__result))
}

//export __CallFuncAliasMat4x4
func __CallFuncAliasMat4x4(func_ unsafe.Pointer) C.Matrix4x4 {
	__result := CallFuncAliasMat4x4(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasMat4x4(nil))).(cross_call_master.FuncAliasMat4x4))
	return *(*C.Matrix4x4)(unsafe.Pointer(&__result))
}

//export __CallFuncAliasAll
func __CallFuncAliasAll(func_ unsafe.Pointer) C.String {
	__result := CallFuncAliasAll(plugify.GetDelegateForFunctionPointer(func_, reflect.TypeOf(cross_call_master.FuncAliasAll(nil))).(cross_call_master.FuncAliasAll))
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnVoid
func __ReverseNoParamReturnVoid() C.String {
	__result := ReverseNoParamReturnVoid()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnBool
func __ReverseNoParamReturnBool() C.String {
	__result := ReverseNoParamReturnBool()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnChar8
func __ReverseNoParamReturnChar8() C.String {
	__result := ReverseNoParamReturnChar8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnChar16
func __ReverseNoParamReturnChar16() C.String {
	__result := ReverseNoParamReturnChar16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnInt8
func __ReverseNoParamReturnInt8() C.String {
	__result := ReverseNoParamReturnInt8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnInt16
func __ReverseNoParamReturnInt16() C.String {
	__result := ReverseNoParamReturnInt16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnInt32
func __ReverseNoParamReturnInt32() C.String {
	__result := ReverseNoParamReturnInt32()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnInt64
func __ReverseNoParamReturnInt64() C.String {
	__result := ReverseNoParamReturnInt64()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnUInt8
func __ReverseNoParamReturnUInt8() C.String {
	__result := ReverseNoParamReturnUInt8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnUInt16
func __ReverseNoParamReturnUInt16() C.String {
	__result := ReverseNoParamReturnUInt16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnUInt32
func __ReverseNoParamReturnUInt32() C.String {
	__result := ReverseNoParamReturnUInt32()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnUInt64
func __ReverseNoParamReturnUInt64() C.String {
	__result := ReverseNoParamReturnUInt64()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnPointer
func __ReverseNoParamReturnPointer() C.String {
	__result := ReverseNoParamReturnPointer()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnFloat
func __ReverseNoParamReturnFloat() C.String {
	__result := ReverseNoParamReturnFloat()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnDouble
func __ReverseNoParamReturnDouble() C.String {
	__result := ReverseNoParamReturnDouble()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnFunction
func __ReverseNoParamReturnFunction() C.String {
	__result := ReverseNoParamReturnFunction()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnString
func __ReverseNoParamReturnString() C.String {
	__result := ReverseNoParamReturnString()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnAny
func __ReverseNoParamReturnAny() C.String {
	__result := ReverseNoParamReturnAny()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayBool
func __ReverseNoParamReturnArrayBool() C.String {
	__result := ReverseNoParamReturnArrayBool()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayChar8
func __ReverseNoParamReturnArrayChar8() C.String {
	__result := ReverseNoParamReturnArrayChar8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayChar16
func __ReverseNoParamReturnArrayChar16() C.String {
	__result := ReverseNoParamReturnArrayChar16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayInt8
func __ReverseNoParamReturnArrayInt8() C.String {
	__result := ReverseNoParamReturnArrayInt8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayInt16
func __ReverseNoParamReturnArrayInt16() C.String {
	__result := ReverseNoParamReturnArrayInt16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayInt32
func __ReverseNoParamReturnArrayInt32() C.String {
	__result := ReverseNoParamReturnArrayInt32()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayInt64
func __ReverseNoParamReturnArrayInt64() C.String {
	__result := ReverseNoParamReturnArrayInt64()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayUInt8
func __ReverseNoParamReturnArrayUInt8() C.String {
	__result := ReverseNoParamReturnArrayUInt8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayUInt16
func __ReverseNoParamReturnArrayUInt16() C.String {
	__result := ReverseNoParamReturnArrayUInt16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayUInt32
func __ReverseNoParamReturnArrayUInt32() C.String {
	__result := ReverseNoParamReturnArrayUInt32()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayUInt64
func __ReverseNoParamReturnArrayUInt64() C.String {
	__result := ReverseNoParamReturnArrayUInt64()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayPointer
func __ReverseNoParamReturnArrayPointer() C.String {
	__result := ReverseNoParamReturnArrayPointer()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayFloat
func __ReverseNoParamReturnArrayFloat() C.String {
	__result := ReverseNoParamReturnArrayFloat()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayDouble
func __ReverseNoParamReturnArrayDouble() C.String {
	__result := ReverseNoParamReturnArrayDouble()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayString
func __ReverseNoParamReturnArrayString() C.String {
	__result := ReverseNoParamReturnArrayString()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnArrayAny
func __ReverseNoParamReturnArrayAny() C.String {
	__result := ReverseNoParamReturnArrayAny()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnVector2
func __ReverseNoParamReturnVector2() C.String {
	__result := ReverseNoParamReturnVector2()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnVector3
func __ReverseNoParamReturnVector3() C.String {
	__result := ReverseNoParamReturnVector3()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnVector4
func __ReverseNoParamReturnVector4() C.String {
	__result := ReverseNoParamReturnVector4()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseNoParamReturnMatrix4x4
func __ReverseNoParamReturnMatrix4x4() C.String {
	__result := ReverseNoParamReturnMatrix4x4()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParam1
func __ReverseParam1() C.String {
	__result := ReverseParam1()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParam2
func __ReverseParam2() C.String {
	__result := ReverseParam2()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParam3
func __ReverseParam3() C.String {
	__result := ReverseParam3()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParam4
func __ReverseParam4() C.String {
	__result := ReverseParam4()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParam5
func __ReverseParam5() C.String {
	__result := ReverseParam5()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParam6
func __ReverseParam6() C.String {
	__result := ReverseParam6()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParam7
func __ReverseParam7() C.String {
	__result := ReverseParam7()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParam8
func __ReverseParam8() C.String {
	__result := ReverseParam8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParam9
func __ReverseParam9() C.String {
	__result := ReverseParam9()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParam10
func __ReverseParam10() C.String {
	__result := ReverseParam10()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamRef1
func __ReverseParamRef1() C.String {
	__result := ReverseParamRef1()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamRef2
func __ReverseParamRef2() C.String {
	__result := ReverseParamRef2()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamRef3
func __ReverseParamRef3() C.String {
	__result := ReverseParamRef3()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamRef4
func __ReverseParamRef4() C.String {
	__result := ReverseParamRef4()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamRef5
func __ReverseParamRef5() C.String {
	__result := ReverseParamRef5()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamRef6
func __ReverseParamRef6() C.String {
	__result := ReverseParamRef6()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamRef7
func __ReverseParamRef7() C.String {
	__result := ReverseParamRef7()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamRef8
func __ReverseParamRef8() C.String {
	__result := ReverseParamRef8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamRef9
func __ReverseParamRef9() C.String {
	__result := ReverseParamRef9()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamRef10
func __ReverseParamRef10() C.String {
	__result := ReverseParamRef10()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamRefVectors
func __ReverseParamRefVectors() C.String {
	__result := ReverseParamRefVectors()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamAllPrimitives
func __ReverseParamAllPrimitives() C.String {
	__result := ReverseParamAllPrimitives()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamAllAliases
func __ReverseParamAllAliases() C.String {
	__result := ReverseParamAllAliases()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamAllRefAliases
func __ReverseParamAllRefAliases() C.String {
	__result := ReverseParamAllRefAliases()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamEnum
func __ReverseParamEnum() C.String {
	__result := ReverseParamEnum()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamEnumRef
func __ReverseParamEnumRef() C.String {
	__result := ReverseParamEnumRef()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamVariant
func __ReverseParamVariant() C.String {
	__result := ReverseParamVariant()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseParamVariantRef
func __ReverseParamVariantRef() C.String {
	__result := ReverseParamVariantRef()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncVoid
func __ReverseCallFuncVoid() C.String {
	__result := ReverseCallFuncVoid()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncBool
func __ReverseCallFuncBool() C.String {
	__result := ReverseCallFuncBool()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncChar8
func __ReverseCallFuncChar8() C.String {
	__result := ReverseCallFuncChar8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncChar16
func __ReverseCallFuncChar16() C.String {
	__result := ReverseCallFuncChar16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncInt8
func __ReverseCallFuncInt8() C.String {
	__result := ReverseCallFuncInt8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncInt16
func __ReverseCallFuncInt16() C.String {
	__result := ReverseCallFuncInt16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncInt32
func __ReverseCallFuncInt32() C.String {
	__result := ReverseCallFuncInt32()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncInt64
func __ReverseCallFuncInt64() C.String {
	__result := ReverseCallFuncInt64()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncUInt8
func __ReverseCallFuncUInt8() C.String {
	__result := ReverseCallFuncUInt8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncUInt16
func __ReverseCallFuncUInt16() C.String {
	__result := ReverseCallFuncUInt16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncUInt32
func __ReverseCallFuncUInt32() C.String {
	__result := ReverseCallFuncUInt32()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncUInt64
func __ReverseCallFuncUInt64() C.String {
	__result := ReverseCallFuncUInt64()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncPtr
func __ReverseCallFuncPtr() C.String {
	__result := ReverseCallFuncPtr()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncFloat
func __ReverseCallFuncFloat() C.String {
	__result := ReverseCallFuncFloat()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncDouble
func __ReverseCallFuncDouble() C.String {
	__result := ReverseCallFuncDouble()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncString
func __ReverseCallFuncString() C.String {
	__result := ReverseCallFuncString()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAny
func __ReverseCallFuncAny() C.String {
	__result := ReverseCallFuncAny()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncBoolVector
func __ReverseCallFuncBoolVector() C.String {
	__result := ReverseCallFuncBoolVector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncChar8Vector
func __ReverseCallFuncChar8Vector() C.String {
	__result := ReverseCallFuncChar8Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncChar16Vector
func __ReverseCallFuncChar16Vector() C.String {
	__result := ReverseCallFuncChar16Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncInt8Vector
func __ReverseCallFuncInt8Vector() C.String {
	__result := ReverseCallFuncInt8Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncInt16Vector
func __ReverseCallFuncInt16Vector() C.String {
	__result := ReverseCallFuncInt16Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncInt32Vector
func __ReverseCallFuncInt32Vector() C.String {
	__result := ReverseCallFuncInt32Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncInt64Vector
func __ReverseCallFuncInt64Vector() C.String {
	__result := ReverseCallFuncInt64Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncUInt8Vector
func __ReverseCallFuncUInt8Vector() C.String {
	__result := ReverseCallFuncUInt8Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncUInt16Vector
func __ReverseCallFuncUInt16Vector() C.String {
	__result := ReverseCallFuncUInt16Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncUInt32Vector
func __ReverseCallFuncUInt32Vector() C.String {
	__result := ReverseCallFuncUInt32Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncUInt64Vector
func __ReverseCallFuncUInt64Vector() C.String {
	__result := ReverseCallFuncUInt64Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncPtrVector
func __ReverseCallFuncPtrVector() C.String {
	__result := ReverseCallFuncPtrVector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncFloatVector
func __ReverseCallFuncFloatVector() C.String {
	__result := ReverseCallFuncFloatVector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncDoubleVector
func __ReverseCallFuncDoubleVector() C.String {
	__result := ReverseCallFuncDoubleVector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncStringVector
func __ReverseCallFuncStringVector() C.String {
	__result := ReverseCallFuncStringVector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAnyVector
func __ReverseCallFuncAnyVector() C.String {
	__result := ReverseCallFuncAnyVector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncVec2Vector
func __ReverseCallFuncVec2Vector() C.String {
	__result := ReverseCallFuncVec2Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncVec3Vector
func __ReverseCallFuncVec3Vector() C.String {
	__result := ReverseCallFuncVec3Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncVec4Vector
func __ReverseCallFuncVec4Vector() C.String {
	__result := ReverseCallFuncVec4Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncMat4x4Vector
func __ReverseCallFuncMat4x4Vector() C.String {
	__result := ReverseCallFuncMat4x4Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncVec2
func __ReverseCallFuncVec2() C.String {
	__result := ReverseCallFuncVec2()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncVec3
func __ReverseCallFuncVec3() C.String {
	__result := ReverseCallFuncVec3()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncVec4
func __ReverseCallFuncVec4() C.String {
	__result := ReverseCallFuncVec4()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncMat4x4
func __ReverseCallFuncMat4x4() C.String {
	__result := ReverseCallFuncMat4x4()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasBool
func __ReverseCallFuncAliasBool() C.String {
	__result := ReverseCallFuncAliasBool()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasChar8
func __ReverseCallFuncAliasChar8() C.String {
	__result := ReverseCallFuncAliasChar8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasChar16
func __ReverseCallFuncAliasChar16() C.String {
	__result := ReverseCallFuncAliasChar16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasInt8
func __ReverseCallFuncAliasInt8() C.String {
	__result := ReverseCallFuncAliasInt8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasInt16
func __ReverseCallFuncAliasInt16() C.String {
	__result := ReverseCallFuncAliasInt16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasInt32
func __ReverseCallFuncAliasInt32() C.String {
	__result := ReverseCallFuncAliasInt32()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasInt64
func __ReverseCallFuncAliasInt64() C.String {
	__result := ReverseCallFuncAliasInt64()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasUInt8
func __ReverseCallFuncAliasUInt8() C.String {
	__result := ReverseCallFuncAliasUInt8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasUInt16
func __ReverseCallFuncAliasUInt16() C.String {
	__result := ReverseCallFuncAliasUInt16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasUInt32
func __ReverseCallFuncAliasUInt32() C.String {
	__result := ReverseCallFuncAliasUInt32()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasUInt64
func __ReverseCallFuncAliasUInt64() C.String {
	__result := ReverseCallFuncAliasUInt64()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasPtr
func __ReverseCallFuncAliasPtr() C.String {
	__result := ReverseCallFuncAliasPtr()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasFloat
func __ReverseCallFuncAliasFloat() C.String {
	__result := ReverseCallFuncAliasFloat()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasDouble
func __ReverseCallFuncAliasDouble() C.String {
	__result := ReverseCallFuncAliasDouble()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasString
func __ReverseCallFuncAliasString() C.String {
	__result := ReverseCallFuncAliasString()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasAny
func __ReverseCallFuncAliasAny() C.String {
	__result := ReverseCallFuncAliasAny()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasBoolVector
func __ReverseCallFuncAliasBoolVector() C.String {
	__result := ReverseCallFuncAliasBoolVector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasChar8Vector
func __ReverseCallFuncAliasChar8Vector() C.String {
	__result := ReverseCallFuncAliasChar8Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasChar16Vector
func __ReverseCallFuncAliasChar16Vector() C.String {
	__result := ReverseCallFuncAliasChar16Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasInt8Vector
func __ReverseCallFuncAliasInt8Vector() C.String {
	__result := ReverseCallFuncAliasInt8Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasInt16Vector
func __ReverseCallFuncAliasInt16Vector() C.String {
	__result := ReverseCallFuncAliasInt16Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasInt32Vector
func __ReverseCallFuncAliasInt32Vector() C.String {
	__result := ReverseCallFuncAliasInt32Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasInt64Vector
func __ReverseCallFuncAliasInt64Vector() C.String {
	__result := ReverseCallFuncAliasInt64Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasUInt8Vector
func __ReverseCallFuncAliasUInt8Vector() C.String {
	__result := ReverseCallFuncAliasUInt8Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasUInt16Vector
func __ReverseCallFuncAliasUInt16Vector() C.String {
	__result := ReverseCallFuncAliasUInt16Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasUInt32Vector
func __ReverseCallFuncAliasUInt32Vector() C.String {
	__result := ReverseCallFuncAliasUInt32Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasUInt64Vector
func __ReverseCallFuncAliasUInt64Vector() C.String {
	__result := ReverseCallFuncAliasUInt64Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasPtrVector
func __ReverseCallFuncAliasPtrVector() C.String {
	__result := ReverseCallFuncAliasPtrVector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasFloatVector
func __ReverseCallFuncAliasFloatVector() C.String {
	__result := ReverseCallFuncAliasFloatVector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasDoubleVector
func __ReverseCallFuncAliasDoubleVector() C.String {
	__result := ReverseCallFuncAliasDoubleVector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasStringVector
func __ReverseCallFuncAliasStringVector() C.String {
	__result := ReverseCallFuncAliasStringVector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasAnyVector
func __ReverseCallFuncAliasAnyVector() C.String {
	__result := ReverseCallFuncAliasAnyVector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasVec2Vector
func __ReverseCallFuncAliasVec2Vector() C.String {
	__result := ReverseCallFuncAliasVec2Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasVec3Vector
func __ReverseCallFuncAliasVec3Vector() C.String {
	__result := ReverseCallFuncAliasVec3Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasVec4Vector
func __ReverseCallFuncAliasVec4Vector() C.String {
	__result := ReverseCallFuncAliasVec4Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasMat4x4Vector
func __ReverseCallFuncAliasMat4x4Vector() C.String {
	__result := ReverseCallFuncAliasMat4x4Vector()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasVec2
func __ReverseCallFuncAliasVec2() C.String {
	__result := ReverseCallFuncAliasVec2()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasVec3
func __ReverseCallFuncAliasVec3() C.String {
	__result := ReverseCallFuncAliasVec3()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasVec4
func __ReverseCallFuncAliasVec4() C.String {
	__result := ReverseCallFuncAliasVec4()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasMat4x4
func __ReverseCallFuncAliasMat4x4() C.String {
	__result := ReverseCallFuncAliasMat4x4()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncAliasAll
func __ReverseCallFuncAliasAll() C.String {
	__result := ReverseCallFuncAliasAll()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc1
func __ReverseCallFunc1() C.String {
	__result := ReverseCallFunc1()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc2
func __ReverseCallFunc2() C.String {
	__result := ReverseCallFunc2()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc3
func __ReverseCallFunc3() C.String {
	__result := ReverseCallFunc3()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc4
func __ReverseCallFunc4() C.String {
	__result := ReverseCallFunc4()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc5
func __ReverseCallFunc5() C.String {
	__result := ReverseCallFunc5()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc6
func __ReverseCallFunc6() C.String {
	__result := ReverseCallFunc6()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc7
func __ReverseCallFunc7() C.String {
	__result := ReverseCallFunc7()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc8
func __ReverseCallFunc8() C.String {
	__result := ReverseCallFunc8()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc9
func __ReverseCallFunc9() C.String {
	__result := ReverseCallFunc9()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc10
func __ReverseCallFunc10() C.String {
	__result := ReverseCallFunc10()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc11
func __ReverseCallFunc11() C.String {
	__result := ReverseCallFunc11()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc12
func __ReverseCallFunc12() C.String {
	__result := ReverseCallFunc12()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc13
func __ReverseCallFunc13() C.String {
	__result := ReverseCallFunc13()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc14
func __ReverseCallFunc14() C.String {
	__result := ReverseCallFunc14()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc15
func __ReverseCallFunc15() C.String {
	__result := ReverseCallFunc15()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc16
func __ReverseCallFunc16() C.String {
	__result := ReverseCallFunc16()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc17
func __ReverseCallFunc17() C.String {
	__result := ReverseCallFunc17()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc18
func __ReverseCallFunc18() C.String {
	__result := ReverseCallFunc18()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc19
func __ReverseCallFunc19() C.String {
	__result := ReverseCallFunc19()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc20
func __ReverseCallFunc20() C.String {
	__result := ReverseCallFunc20()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc21
func __ReverseCallFunc21() C.String {
	__result := ReverseCallFunc21()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc22
func __ReverseCallFunc22() C.String {
	__result := ReverseCallFunc22()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc23
func __ReverseCallFunc23() C.String {
	__result := ReverseCallFunc23()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc24
func __ReverseCallFunc24() C.String {
	__result := ReverseCallFunc24()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc25
func __ReverseCallFunc25() C.String {
	__result := ReverseCallFunc25()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc26
func __ReverseCallFunc26() C.String {
	__result := ReverseCallFunc26()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc27
func __ReverseCallFunc27() C.String {
	__result := ReverseCallFunc27()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc28
func __ReverseCallFunc28() C.String {
	__result := ReverseCallFunc28()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc29
func __ReverseCallFunc29() C.String {
	__result := ReverseCallFunc29()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc30
func __ReverseCallFunc30() C.String {
	__result := ReverseCallFunc30()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc31
func __ReverseCallFunc31() C.String {
	__result := ReverseCallFunc31()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc32
func __ReverseCallFunc32() C.String {
	__result := ReverseCallFunc32()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFunc33
func __ReverseCallFunc33() C.String {
	__result := ReverseCallFunc33()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCallFuncEnum
func __ReverseCallFuncEnum() C.String {
	__result := ReverseCallFuncEnum()
	__return := plugify.ConstructString(__result)
	return *(*C.String)(unsafe.Pointer(&__return))
}

//export __ReverseCall
func __ReverseCall(test *C.String) {
	ReverseCall(plugify.GetStringData((*plugify.PlgString)(unsafe.Pointer(test))))
}
