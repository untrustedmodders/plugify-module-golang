#pragma once

#include "shared.h"

extern void (*__cross_call_worker_NoParamReturnVoid)();

static void NoParamReturnVoid() {
	__cross_call_worker_NoParamReturnVoid();
}

extern bool (*__cross_call_worker_NoParamReturnBool)();

static bool NoParamReturnBool() {
	return __cross_call_worker_NoParamReturnBool();
}

extern int8_t (*__cross_call_worker_NoParamReturnChar8)();

static int8_t NoParamReturnChar8() {
	return __cross_call_worker_NoParamReturnChar8();
}

extern uint16_t (*__cross_call_worker_NoParamReturnChar16)();

static uint16_t NoParamReturnChar16() {
	return __cross_call_worker_NoParamReturnChar16();
}

extern int8_t (*__cross_call_worker_NoParamReturnInt8)();

static int8_t NoParamReturnInt8() {
	return __cross_call_worker_NoParamReturnInt8();
}

extern int16_t (*__cross_call_worker_NoParamReturnInt16)();

static int16_t NoParamReturnInt16() {
	return __cross_call_worker_NoParamReturnInt16();
}

extern int32_t (*__cross_call_worker_NoParamReturnInt32)();

static int32_t NoParamReturnInt32() {
	return __cross_call_worker_NoParamReturnInt32();
}

extern int64_t (*__cross_call_worker_NoParamReturnInt64)();

static int64_t NoParamReturnInt64() {
	return __cross_call_worker_NoParamReturnInt64();
}

extern uint8_t (*__cross_call_worker_NoParamReturnUInt8)();

static uint8_t NoParamReturnUInt8() {
	return __cross_call_worker_NoParamReturnUInt8();
}

extern uint16_t (*__cross_call_worker_NoParamReturnUInt16)();

static uint16_t NoParamReturnUInt16() {
	return __cross_call_worker_NoParamReturnUInt16();
}

extern uint32_t (*__cross_call_worker_NoParamReturnUInt32)();

static uint32_t NoParamReturnUInt32() {
	return __cross_call_worker_NoParamReturnUInt32();
}

extern uint64_t (*__cross_call_worker_NoParamReturnUInt64)();

static uint64_t NoParamReturnUInt64() {
	return __cross_call_worker_NoParamReturnUInt64();
}

extern uintptr_t (*__cross_call_worker_NoParamReturnPointer)();

static uintptr_t NoParamReturnPointer() {
	return __cross_call_worker_NoParamReturnPointer();
}

extern float (*__cross_call_worker_NoParamReturnFloat)();

static float NoParamReturnFloat() {
	return __cross_call_worker_NoParamReturnFloat();
}

extern double (*__cross_call_worker_NoParamReturnDouble)();

static double NoParamReturnDouble() {
	return __cross_call_worker_NoParamReturnDouble();
}

extern void* (*__cross_call_worker_NoParamReturnFunction)();

static void* NoParamReturnFunction() {
	return __cross_call_worker_NoParamReturnFunction();
}

extern String (*__cross_call_worker_NoParamReturnString)();

static String NoParamReturnString() {
	return __cross_call_worker_NoParamReturnString();
}

extern Variant (*__cross_call_worker_NoParamReturnAny)();

static Variant NoParamReturnAny() {
	return __cross_call_worker_NoParamReturnAny();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayBool)();

static Vector NoParamReturnArrayBool() {
	return __cross_call_worker_NoParamReturnArrayBool();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayChar8)();

static Vector NoParamReturnArrayChar8() {
	return __cross_call_worker_NoParamReturnArrayChar8();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayChar16)();

static Vector NoParamReturnArrayChar16() {
	return __cross_call_worker_NoParamReturnArrayChar16();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayInt8)();

static Vector NoParamReturnArrayInt8() {
	return __cross_call_worker_NoParamReturnArrayInt8();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayInt16)();

static Vector NoParamReturnArrayInt16() {
	return __cross_call_worker_NoParamReturnArrayInt16();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayInt32)();

static Vector NoParamReturnArrayInt32() {
	return __cross_call_worker_NoParamReturnArrayInt32();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayInt64)();

static Vector NoParamReturnArrayInt64() {
	return __cross_call_worker_NoParamReturnArrayInt64();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayUInt8)();

static Vector NoParamReturnArrayUInt8() {
	return __cross_call_worker_NoParamReturnArrayUInt8();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayUInt16)();

static Vector NoParamReturnArrayUInt16() {
	return __cross_call_worker_NoParamReturnArrayUInt16();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayUInt32)();

static Vector NoParamReturnArrayUInt32() {
	return __cross_call_worker_NoParamReturnArrayUInt32();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayUInt64)();

static Vector NoParamReturnArrayUInt64() {
	return __cross_call_worker_NoParamReturnArrayUInt64();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayPointer)();

static Vector NoParamReturnArrayPointer() {
	return __cross_call_worker_NoParamReturnArrayPointer();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayFloat)();

static Vector NoParamReturnArrayFloat() {
	return __cross_call_worker_NoParamReturnArrayFloat();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayDouble)();

static Vector NoParamReturnArrayDouble() {
	return __cross_call_worker_NoParamReturnArrayDouble();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayString)();

static Vector NoParamReturnArrayString() {
	return __cross_call_worker_NoParamReturnArrayString();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayAny)();

static Vector NoParamReturnArrayAny() {
	return __cross_call_worker_NoParamReturnArrayAny();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayVector2)();

static Vector NoParamReturnArrayVector2() {
	return __cross_call_worker_NoParamReturnArrayVector2();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayVector3)();

static Vector NoParamReturnArrayVector3() {
	return __cross_call_worker_NoParamReturnArrayVector3();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayVector4)();

static Vector NoParamReturnArrayVector4() {
	return __cross_call_worker_NoParamReturnArrayVector4();
}

extern Vector (*__cross_call_worker_NoParamReturnArrayMatrix4x4)();

static Vector NoParamReturnArrayMatrix4x4() {
	return __cross_call_worker_NoParamReturnArrayMatrix4x4();
}

extern Vector2 (*__cross_call_worker_NoParamReturnVector2)();

static Vector2 NoParamReturnVector2() {
	return __cross_call_worker_NoParamReturnVector2();
}

extern Vector3 (*__cross_call_worker_NoParamReturnVector3)();

static Vector3 NoParamReturnVector3() {
	return __cross_call_worker_NoParamReturnVector3();
}

extern Vector4 (*__cross_call_worker_NoParamReturnVector4)();

static Vector4 NoParamReturnVector4() {
	return __cross_call_worker_NoParamReturnVector4();
}

extern Matrix4x4 (*__cross_call_worker_NoParamReturnMatrix4x4)();

static Matrix4x4 NoParamReturnMatrix4x4() {
	return __cross_call_worker_NoParamReturnMatrix4x4();
}

extern void (*__cross_call_worker_Param1)(int32_t);

static void Param1(int32_t a) {
	__cross_call_worker_Param1(a);
}

extern void (*__cross_call_worker_Param2)(int32_t, float);

static void Param2(int32_t a, float b) {
	__cross_call_worker_Param2(a, b);
}

extern void (*__cross_call_worker_Param3)(int32_t, float, double);

static void Param3(int32_t a, float b, double c) {
	__cross_call_worker_Param3(a, b, c);
}

extern void (*__cross_call_worker_Param4)(int32_t, float, double, Vector4*);

static void Param4(int32_t a, float b, double c, Vector4* d) {
	__cross_call_worker_Param4(a, b, c, d);
}

extern void (*__cross_call_worker_Param5)(int32_t, float, double, Vector4*, Vector*);

static void Param5(int32_t a, float b, double c, Vector4* d, Vector* e) {
	__cross_call_worker_Param5(a, b, c, d, e);
}

extern void (*__cross_call_worker_Param6)(int32_t, float, double, Vector4*, Vector*, int8_t);

static void Param6(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f) {
	__cross_call_worker_Param6(a, b, c, d, e, f);
}

extern void (*__cross_call_worker_Param7)(int32_t, float, double, Vector4*, Vector*, int8_t, String*);

static void Param7(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f, String* g) {
	__cross_call_worker_Param7(a, b, c, d, e, f, g);
}

extern void (*__cross_call_worker_Param8)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t);

static void Param8(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f, String* g, uint16_t h) {
	__cross_call_worker_Param8(a, b, c, d, e, f, g, h);
}

extern void (*__cross_call_worker_Param9)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t, int16_t);

static void Param9(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f, String* g, uint16_t h, int16_t k) {
	__cross_call_worker_Param9(a, b, c, d, e, f, g, h, k);
}

extern void (*__cross_call_worker_Param10)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t, int16_t, uintptr_t);

static void Param10(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f, String* g, uint16_t h, int16_t k, uintptr_t l) {
	__cross_call_worker_Param10(a, b, c, d, e, f, g, h, k, l);
}

extern void (*__cross_call_worker_ParamRef1)(int32_t*);

static void ParamRef1(int32_t* a) {
	__cross_call_worker_ParamRef1(a);
}

extern void (*__cross_call_worker_ParamRef2)(int32_t*, float*);

static void ParamRef2(int32_t* a, float* b) {
	__cross_call_worker_ParamRef2(a, b);
}

extern void (*__cross_call_worker_ParamRef3)(int32_t*, float*, double*);

static void ParamRef3(int32_t* a, float* b, double* c) {
	__cross_call_worker_ParamRef3(a, b, c);
}

extern void (*__cross_call_worker_ParamRef4)(int32_t*, float*, double*, Vector4*);

static void ParamRef4(int32_t* a, float* b, double* c, Vector4* d) {
	__cross_call_worker_ParamRef4(a, b, c, d);
}

extern void (*__cross_call_worker_ParamRef5)(int32_t*, float*, double*, Vector4*, Vector*);

static void ParamRef5(int32_t* a, float* b, double* c, Vector4* d, Vector* e) {
	__cross_call_worker_ParamRef5(a, b, c, d, e);
}

extern void (*__cross_call_worker_ParamRef6)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*);

static void ParamRef6(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f) {
	__cross_call_worker_ParamRef6(a, b, c, d, e, f);
}

extern void (*__cross_call_worker_ParamRef7)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*);

static void ParamRef7(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f, String* g) {
	__cross_call_worker_ParamRef7(a, b, c, d, e, f, g);
}

extern void (*__cross_call_worker_ParamRef8)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*);

static void ParamRef8(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f, String* g, uint16_t* h) {
	__cross_call_worker_ParamRef8(a, b, c, d, e, f, g, h);
}

extern void (*__cross_call_worker_ParamRef9)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*, int16_t*);

static void ParamRef9(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f, String* g, uint16_t* h, int16_t* k) {
	__cross_call_worker_ParamRef9(a, b, c, d, e, f, g, h, k);
}

extern void (*__cross_call_worker_ParamRef10)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*, int16_t*, uintptr_t*);

static void ParamRef10(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f, String* g, uint16_t* h, int16_t* k, uintptr_t* l) {
	__cross_call_worker_ParamRef10(a, b, c, d, e, f, g, h, k, l);
}

extern void (*__cross_call_worker_ParamRefVectors)(Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*);

static void ParamRefVectors(Vector* p1, Vector* p2, Vector* p3, Vector* p4, Vector* p5, Vector* p6, Vector* p7, Vector* p8, Vector* p9, Vector* p10, Vector* p11, Vector* p12, Vector* p13, Vector* p14, Vector* p15) {
	__cross_call_worker_ParamRefVectors(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15);
}

extern int64_t (*__cross_call_worker_ParamAllPrimitives)(bool, int8_t, uint16_t, int8_t, int16_t, int32_t, int64_t, uint8_t, uint16_t, uint32_t, uint64_t, uintptr_t, float, double);

static int64_t ParamAllPrimitives(bool p1, int8_t p2, uint16_t p3, int8_t p4, int16_t p5, int32_t p6, int64_t p7, uint8_t p8, uint16_t p9, uint32_t p10, uint64_t p11, uintptr_t p12, float p13, double p14) {
	return __cross_call_worker_ParamAllPrimitives(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14);
}

extern int32_t (*__cross_call_worker_ParamAllAliases)(bool, int8_t, uint16_t, int8_t, int16_t, int32_t, int64_t, uintptr_t, float, double, String*, Variant*, Vector2*, Vector3*, Vector4*, Matrix4x4*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*);

static int32_t ParamAllAliases(bool aBool, int8_t aChar8, uint16_t aChar16, int8_t aInt8, int16_t aInt16, int32_t aInt32, int64_t aInt64, uintptr_t aPtr, float aFloat, double aDouble, String* aString, Variant* aAny, Vector2* aVec2, Vector3* aVec3, Vector4* aVec4, Matrix4x4* aMat4x4, Vector* aBoolVec, Vector* aChar8Vec, Vector* aChar16Vec, Vector* aInt8Vec, Vector* aInt16Vec, Vector* aInt32Vec, Vector* aInt64Vec, Vector* aPtrVec, Vector* aFloatVec, Vector* aDoubleVec, Vector* aStringVec, Vector* aAnyVec, Vector* aVec2Vec, Vector* aVec3Vec, Vector* aVec4Vec) {
	return __cross_call_worker_ParamAllAliases(aBool, aChar8, aChar16, aInt8, aInt16, aInt32, aInt64, aPtr, aFloat, aDouble, aString, aAny, aVec2, aVec3, aVec4, aMat4x4, aBoolVec, aChar8Vec, aChar16Vec, aInt8Vec, aInt16Vec, aInt32Vec, aInt64Vec, aPtrVec, aFloatVec, aDoubleVec, aStringVec, aAnyVec, aVec2Vec, aVec3Vec, aVec4Vec);
}

extern int64_t (*__cross_call_worker_ParamAllRefAliases)(bool*, int8_t*, uint16_t*, int8_t*, int16_t*, int32_t*, int64_t*, uintptr_t*, float*, double*, String*, Variant*, Vector2*, Vector3*, Vector4*, Matrix4x4*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*);

static int64_t ParamAllRefAliases(bool* aBool, int8_t* aChar8, uint16_t* aChar16, int8_t* aInt8, int16_t* aInt16, int32_t* aInt32, int64_t* aInt64, uintptr_t* aPtr, float* aFloat, double* aDouble, String* aString, Variant* aAny, Vector2* aVec2, Vector3* aVec3, Vector4* aVec4, Matrix4x4* aMat4x4, Vector* aBoolVec, Vector* aChar8Vec, Vector* aChar16Vec, Vector* aInt8Vec, Vector* aInt16Vec, Vector* aInt32Vec, Vector* aInt64Vec, Vector* aPtrVec, Vector* aFloatVec, Vector* aDoubleVec, Vector* aStringVec, Vector* aAnyVec, Vector* aVec2Vec, Vector* aVec3Vec, Vector* aVec4Vec) {
	return __cross_call_worker_ParamAllRefAliases(aBool, aChar8, aChar16, aInt8, aInt16, aInt32, aInt64, aPtr, aFloat, aDouble, aString, aAny, aVec2, aVec3, aVec4, aMat4x4, aBoolVec, aChar8Vec, aChar16Vec, aInt8Vec, aInt16Vec, aInt32Vec, aInt64Vec, aPtrVec, aFloatVec, aDoubleVec, aStringVec, aAnyVec, aVec2Vec, aVec3Vec, aVec4Vec);
}

extern void (*__cross_call_worker_ParamVariant)(Variant*, Vector*);

static void ParamVariant(Variant* p1, Vector* p2) {
	__cross_call_worker_ParamVariant(p1, p2);
}

extern int32_t (*__cross_call_worker_ParamEnum)(int32_t, Vector*);

static int32_t ParamEnum(int32_t p1, Vector* p2) {
	return __cross_call_worker_ParamEnum(p1, p2);
}

extern int32_t (*__cross_call_worker_ParamEnumRef)(int32_t*, Vector*);

static int32_t ParamEnumRef(int32_t* p1, Vector* p2) {
	return __cross_call_worker_ParamEnumRef(p1, p2);
}

extern void (*__cross_call_worker_ParamVariantRef)(Variant*, Vector*);

static void ParamVariantRef(Variant* p1, Vector* p2) {
	__cross_call_worker_ParamVariantRef(p1, p2);
}

extern void (*__cross_call_worker_CallFuncVoid)(void*);

static void CallFuncVoid(void* func_) {
	__cross_call_worker_CallFuncVoid(func_);
}

extern bool (*__cross_call_worker_CallFuncBool)(void*);

static bool CallFuncBool(void* func_) {
	return __cross_call_worker_CallFuncBool(func_);
}

extern int8_t (*__cross_call_worker_CallFuncChar8)(void*);

static int8_t CallFuncChar8(void* func_) {
	return __cross_call_worker_CallFuncChar8(func_);
}

extern uint16_t (*__cross_call_worker_CallFuncChar16)(void*);

static uint16_t CallFuncChar16(void* func_) {
	return __cross_call_worker_CallFuncChar16(func_);
}

extern int8_t (*__cross_call_worker_CallFuncInt8)(void*);

static int8_t CallFuncInt8(void* func_) {
	return __cross_call_worker_CallFuncInt8(func_);
}

extern int16_t (*__cross_call_worker_CallFuncInt16)(void*);

static int16_t CallFuncInt16(void* func_) {
	return __cross_call_worker_CallFuncInt16(func_);
}

extern int32_t (*__cross_call_worker_CallFuncInt32)(void*);

static int32_t CallFuncInt32(void* func_) {
	return __cross_call_worker_CallFuncInt32(func_);
}

extern int64_t (*__cross_call_worker_CallFuncInt64)(void*);

static int64_t CallFuncInt64(void* func_) {
	return __cross_call_worker_CallFuncInt64(func_);
}

extern uint8_t (*__cross_call_worker_CallFuncUInt8)(void*);

static uint8_t CallFuncUInt8(void* func_) {
	return __cross_call_worker_CallFuncUInt8(func_);
}

extern uint16_t (*__cross_call_worker_CallFuncUInt16)(void*);

static uint16_t CallFuncUInt16(void* func_) {
	return __cross_call_worker_CallFuncUInt16(func_);
}

extern uint32_t (*__cross_call_worker_CallFuncUInt32)(void*);

static uint32_t CallFuncUInt32(void* func_) {
	return __cross_call_worker_CallFuncUInt32(func_);
}

extern uint64_t (*__cross_call_worker_CallFuncUInt64)(void*);

static uint64_t CallFuncUInt64(void* func_) {
	return __cross_call_worker_CallFuncUInt64(func_);
}

extern uintptr_t (*__cross_call_worker_CallFuncPtr)(void*);

static uintptr_t CallFuncPtr(void* func_) {
	return __cross_call_worker_CallFuncPtr(func_);
}

extern float (*__cross_call_worker_CallFuncFloat)(void*);

static float CallFuncFloat(void* func_) {
	return __cross_call_worker_CallFuncFloat(func_);
}

extern double (*__cross_call_worker_CallFuncDouble)(void*);

static double CallFuncDouble(void* func_) {
	return __cross_call_worker_CallFuncDouble(func_);
}

extern String (*__cross_call_worker_CallFuncString)(void*);

static String CallFuncString(void* func_) {
	return __cross_call_worker_CallFuncString(func_);
}

extern Variant (*__cross_call_worker_CallFuncAny)(void*);

static Variant CallFuncAny(void* func_) {
	return __cross_call_worker_CallFuncAny(func_);
}

extern uintptr_t (*__cross_call_worker_CallFuncFunction)(void*);

static uintptr_t CallFuncFunction(void* func_) {
	return __cross_call_worker_CallFuncFunction(func_);
}

extern Vector (*__cross_call_worker_CallFuncBoolVector)(void*);

static Vector CallFuncBoolVector(void* func_) {
	return __cross_call_worker_CallFuncBoolVector(func_);
}

extern Vector (*__cross_call_worker_CallFuncChar8Vector)(void*);

static Vector CallFuncChar8Vector(void* func_) {
	return __cross_call_worker_CallFuncChar8Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncChar16Vector)(void*);

static Vector CallFuncChar16Vector(void* func_) {
	return __cross_call_worker_CallFuncChar16Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncInt8Vector)(void*);

static Vector CallFuncInt8Vector(void* func_) {
	return __cross_call_worker_CallFuncInt8Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncInt16Vector)(void*);

static Vector CallFuncInt16Vector(void* func_) {
	return __cross_call_worker_CallFuncInt16Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncInt32Vector)(void*);

static Vector CallFuncInt32Vector(void* func_) {
	return __cross_call_worker_CallFuncInt32Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncInt64Vector)(void*);

static Vector CallFuncInt64Vector(void* func_) {
	return __cross_call_worker_CallFuncInt64Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncUInt8Vector)(void*);

static Vector CallFuncUInt8Vector(void* func_) {
	return __cross_call_worker_CallFuncUInt8Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncUInt16Vector)(void*);

static Vector CallFuncUInt16Vector(void* func_) {
	return __cross_call_worker_CallFuncUInt16Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncUInt32Vector)(void*);

static Vector CallFuncUInt32Vector(void* func_) {
	return __cross_call_worker_CallFuncUInt32Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncUInt64Vector)(void*);

static Vector CallFuncUInt64Vector(void* func_) {
	return __cross_call_worker_CallFuncUInt64Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncPtrVector)(void*);

static Vector CallFuncPtrVector(void* func_) {
	return __cross_call_worker_CallFuncPtrVector(func_);
}

extern Vector (*__cross_call_worker_CallFuncFloatVector)(void*);

static Vector CallFuncFloatVector(void* func_) {
	return __cross_call_worker_CallFuncFloatVector(func_);
}

extern Vector (*__cross_call_worker_CallFuncDoubleVector)(void*);

static Vector CallFuncDoubleVector(void* func_) {
	return __cross_call_worker_CallFuncDoubleVector(func_);
}

extern Vector (*__cross_call_worker_CallFuncStringVector)(void*);

static Vector CallFuncStringVector(void* func_) {
	return __cross_call_worker_CallFuncStringVector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAnyVector)(void*);

static Vector CallFuncAnyVector(void* func_) {
	return __cross_call_worker_CallFuncAnyVector(func_);
}

extern Vector (*__cross_call_worker_CallFuncVec2Vector)(void*);

static Vector CallFuncVec2Vector(void* func_) {
	return __cross_call_worker_CallFuncVec2Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncVec3Vector)(void*);

static Vector CallFuncVec3Vector(void* func_) {
	return __cross_call_worker_CallFuncVec3Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncVec4Vector)(void*);

static Vector CallFuncVec4Vector(void* func_) {
	return __cross_call_worker_CallFuncVec4Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncMat4x4Vector)(void*);

static Vector CallFuncMat4x4Vector(void* func_) {
	return __cross_call_worker_CallFuncMat4x4Vector(func_);
}

extern Vector2 (*__cross_call_worker_CallFuncVec2)(void*);

static Vector2 CallFuncVec2(void* func_) {
	return __cross_call_worker_CallFuncVec2(func_);
}

extern Vector3 (*__cross_call_worker_CallFuncVec3)(void*);

static Vector3 CallFuncVec3(void* func_) {
	return __cross_call_worker_CallFuncVec3(func_);
}

extern Vector4 (*__cross_call_worker_CallFuncVec4)(void*);

static Vector4 CallFuncVec4(void* func_) {
	return __cross_call_worker_CallFuncVec4(func_);
}

extern Matrix4x4 (*__cross_call_worker_CallFuncMat4x4)(void*);

static Matrix4x4 CallFuncMat4x4(void* func_) {
	return __cross_call_worker_CallFuncMat4x4(func_);
}

extern bool (*__cross_call_worker_CallFuncAliasBool)(void*);

static bool CallFuncAliasBool(void* func_) {
	return __cross_call_worker_CallFuncAliasBool(func_);
}

extern int8_t (*__cross_call_worker_CallFuncAliasChar8)(void*);

static int8_t CallFuncAliasChar8(void* func_) {
	return __cross_call_worker_CallFuncAliasChar8(func_);
}

extern uint16_t (*__cross_call_worker_CallFuncAliasChar16)(void*);

static uint16_t CallFuncAliasChar16(void* func_) {
	return __cross_call_worker_CallFuncAliasChar16(func_);
}

extern int8_t (*__cross_call_worker_CallFuncAliasInt8)(void*);

static int8_t CallFuncAliasInt8(void* func_) {
	return __cross_call_worker_CallFuncAliasInt8(func_);
}

extern int16_t (*__cross_call_worker_CallFuncAliasInt16)(void*);

static int16_t CallFuncAliasInt16(void* func_) {
	return __cross_call_worker_CallFuncAliasInt16(func_);
}

extern int32_t (*__cross_call_worker_CallFuncAliasInt32)(void*);

static int32_t CallFuncAliasInt32(void* func_) {
	return __cross_call_worker_CallFuncAliasInt32(func_);
}

extern int64_t (*__cross_call_worker_CallFuncAliasInt64)(void*);

static int64_t CallFuncAliasInt64(void* func_) {
	return __cross_call_worker_CallFuncAliasInt64(func_);
}

extern uint8_t (*__cross_call_worker_CallFuncAliasUInt8)(void*);

static uint8_t CallFuncAliasUInt8(void* func_) {
	return __cross_call_worker_CallFuncAliasUInt8(func_);
}

extern uint16_t (*__cross_call_worker_CallFuncAliasUInt16)(void*);

static uint16_t CallFuncAliasUInt16(void* func_) {
	return __cross_call_worker_CallFuncAliasUInt16(func_);
}

extern uint32_t (*__cross_call_worker_CallFuncAliasUInt32)(void*);

static uint32_t CallFuncAliasUInt32(void* func_) {
	return __cross_call_worker_CallFuncAliasUInt32(func_);
}

extern uint64_t (*__cross_call_worker_CallFuncAliasUInt64)(void*);

static uint64_t CallFuncAliasUInt64(void* func_) {
	return __cross_call_worker_CallFuncAliasUInt64(func_);
}

extern uintptr_t (*__cross_call_worker_CallFuncAliasPtr)(void*);

static uintptr_t CallFuncAliasPtr(void* func_) {
	return __cross_call_worker_CallFuncAliasPtr(func_);
}

extern float (*__cross_call_worker_CallFuncAliasFloat)(void*);

static float CallFuncAliasFloat(void* func_) {
	return __cross_call_worker_CallFuncAliasFloat(func_);
}

extern double (*__cross_call_worker_CallFuncAliasDouble)(void*);

static double CallFuncAliasDouble(void* func_) {
	return __cross_call_worker_CallFuncAliasDouble(func_);
}

extern String (*__cross_call_worker_CallFuncAliasString)(void*);

static String CallFuncAliasString(void* func_) {
	return __cross_call_worker_CallFuncAliasString(func_);
}

extern Variant (*__cross_call_worker_CallFuncAliasAny)(void*);

static Variant CallFuncAliasAny(void* func_) {
	return __cross_call_worker_CallFuncAliasAny(func_);
}

extern uintptr_t (*__cross_call_worker_CallFuncAliasFunction)(void*);

static uintptr_t CallFuncAliasFunction(void* func_) {
	return __cross_call_worker_CallFuncAliasFunction(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasBoolVector)(void*);

static Vector CallFuncAliasBoolVector(void* func_) {
	return __cross_call_worker_CallFuncAliasBoolVector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasChar8Vector)(void*);

static Vector CallFuncAliasChar8Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasChar8Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasChar16Vector)(void*);

static Vector CallFuncAliasChar16Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasChar16Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasInt8Vector)(void*);

static Vector CallFuncAliasInt8Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasInt8Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasInt16Vector)(void*);

static Vector CallFuncAliasInt16Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasInt16Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasInt32Vector)(void*);

static Vector CallFuncAliasInt32Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasInt32Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasInt64Vector)(void*);

static Vector CallFuncAliasInt64Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasInt64Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasUInt8Vector)(void*);

static Vector CallFuncAliasUInt8Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasUInt8Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasUInt16Vector)(void*);

static Vector CallFuncAliasUInt16Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasUInt16Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasUInt32Vector)(void*);

static Vector CallFuncAliasUInt32Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasUInt32Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasUInt64Vector)(void*);

static Vector CallFuncAliasUInt64Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasUInt64Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasPtrVector)(void*);

static Vector CallFuncAliasPtrVector(void* func_) {
	return __cross_call_worker_CallFuncAliasPtrVector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasFloatVector)(void*);

static Vector CallFuncAliasFloatVector(void* func_) {
	return __cross_call_worker_CallFuncAliasFloatVector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasDoubleVector)(void*);

static Vector CallFuncAliasDoubleVector(void* func_) {
	return __cross_call_worker_CallFuncAliasDoubleVector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasStringVector)(void*);

static Vector CallFuncAliasStringVector(void* func_) {
	return __cross_call_worker_CallFuncAliasStringVector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasAnyVector)(void*);

static Vector CallFuncAliasAnyVector(void* func_) {
	return __cross_call_worker_CallFuncAliasAnyVector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasVec2Vector)(void*);

static Vector CallFuncAliasVec2Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasVec2Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasVec3Vector)(void*);

static Vector CallFuncAliasVec3Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasVec3Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasVec4Vector)(void*);

static Vector CallFuncAliasVec4Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasVec4Vector(func_);
}

extern Vector (*__cross_call_worker_CallFuncAliasMat4x4Vector)(void*);

static Vector CallFuncAliasMat4x4Vector(void* func_) {
	return __cross_call_worker_CallFuncAliasMat4x4Vector(func_);
}

extern Vector2 (*__cross_call_worker_CallFuncAliasVec2)(void*);

static Vector2 CallFuncAliasVec2(void* func_) {
	return __cross_call_worker_CallFuncAliasVec2(func_);
}

extern Vector3 (*__cross_call_worker_CallFuncAliasVec3)(void*);

static Vector3 CallFuncAliasVec3(void* func_) {
	return __cross_call_worker_CallFuncAliasVec3(func_);
}

extern Vector4 (*__cross_call_worker_CallFuncAliasVec4)(void*);

static Vector4 CallFuncAliasVec4(void* func_) {
	return __cross_call_worker_CallFuncAliasVec4(func_);
}

extern Matrix4x4 (*__cross_call_worker_CallFuncAliasMat4x4)(void*);

static Matrix4x4 CallFuncAliasMat4x4(void* func_) {
	return __cross_call_worker_CallFuncAliasMat4x4(func_);
}

extern String (*__cross_call_worker_CallFuncAliasAll)(void*);

static String CallFuncAliasAll(void* func_) {
	return __cross_call_worker_CallFuncAliasAll(func_);
}

extern int32_t (*__cross_call_worker_CallFunc1)(void*);

static int32_t CallFunc1(void* func_) {
	return __cross_call_worker_CallFunc1(func_);
}

extern int8_t (*__cross_call_worker_CallFunc2)(void*);

static int8_t CallFunc2(void* func_) {
	return __cross_call_worker_CallFunc2(func_);
}

extern void (*__cross_call_worker_CallFunc3)(void*);

static void CallFunc3(void* func_) {
	__cross_call_worker_CallFunc3(func_);
}

extern Vector4 (*__cross_call_worker_CallFunc4)(void*);

static Vector4 CallFunc4(void* func_) {
	return __cross_call_worker_CallFunc4(func_);
}

extern bool (*__cross_call_worker_CallFunc5)(void*);

static bool CallFunc5(void* func_) {
	return __cross_call_worker_CallFunc5(func_);
}

extern int64_t (*__cross_call_worker_CallFunc6)(void*);

static int64_t CallFunc6(void* func_) {
	return __cross_call_worker_CallFunc6(func_);
}

extern double (*__cross_call_worker_CallFunc7)(void*);

static double CallFunc7(void* func_) {
	return __cross_call_worker_CallFunc7(func_);
}

extern Matrix4x4 (*__cross_call_worker_CallFunc8)(void*);

static Matrix4x4 CallFunc8(void* func_) {
	return __cross_call_worker_CallFunc8(func_);
}

extern void (*__cross_call_worker_CallFunc9)(void*);

static void CallFunc9(void* func_) {
	__cross_call_worker_CallFunc9(func_);
}

extern uint32_t (*__cross_call_worker_CallFunc10)(void*);

static uint32_t CallFunc10(void* func_) {
	return __cross_call_worker_CallFunc10(func_);
}

extern uintptr_t (*__cross_call_worker_CallFunc11)(void*);

static uintptr_t CallFunc11(void* func_) {
	return __cross_call_worker_CallFunc11(func_);
}

extern bool (*__cross_call_worker_CallFunc12)(void*);

static bool CallFunc12(void* func_) {
	return __cross_call_worker_CallFunc12(func_);
}

extern String (*__cross_call_worker_CallFunc13)(void*);

static String CallFunc13(void* func_) {
	return __cross_call_worker_CallFunc13(func_);
}

extern Vector (*__cross_call_worker_CallFunc14)(void*);

static Vector CallFunc14(void* func_) {
	return __cross_call_worker_CallFunc14(func_);
}

extern int16_t (*__cross_call_worker_CallFunc15)(void*);

static int16_t CallFunc15(void* func_) {
	return __cross_call_worker_CallFunc15(func_);
}

extern uintptr_t (*__cross_call_worker_CallFunc16)(void*);

static uintptr_t CallFunc16(void* func_) {
	return __cross_call_worker_CallFunc16(func_);
}

extern String (*__cross_call_worker_CallFunc17)(void*);

static String CallFunc17(void* func_) {
	return __cross_call_worker_CallFunc17(func_);
}

extern String (*__cross_call_worker_CallFunc18)(void*);

static String CallFunc18(void* func_) {
	return __cross_call_worker_CallFunc18(func_);
}

extern String (*__cross_call_worker_CallFunc19)(void*);

static String CallFunc19(void* func_) {
	return __cross_call_worker_CallFunc19(func_);
}

extern String (*__cross_call_worker_CallFunc20)(void*);

static String CallFunc20(void* func_) {
	return __cross_call_worker_CallFunc20(func_);
}

extern String (*__cross_call_worker_CallFunc21)(void*);

static String CallFunc21(void* func_) {
	return __cross_call_worker_CallFunc21(func_);
}

extern String (*__cross_call_worker_CallFunc22)(void*);

static String CallFunc22(void* func_) {
	return __cross_call_worker_CallFunc22(func_);
}

extern String (*__cross_call_worker_CallFunc23)(void*);

static String CallFunc23(void* func_) {
	return __cross_call_worker_CallFunc23(func_);
}

extern String (*__cross_call_worker_CallFunc24)(void*);

static String CallFunc24(void* func_) {
	return __cross_call_worker_CallFunc24(func_);
}

extern String (*__cross_call_worker_CallFunc25)(void*);

static String CallFunc25(void* func_) {
	return __cross_call_worker_CallFunc25(func_);
}

extern String (*__cross_call_worker_CallFunc26)(void*);

static String CallFunc26(void* func_) {
	return __cross_call_worker_CallFunc26(func_);
}

extern String (*__cross_call_worker_CallFunc27)(void*);

static String CallFunc27(void* func_) {
	return __cross_call_worker_CallFunc27(func_);
}

extern String (*__cross_call_worker_CallFunc28)(void*);

static String CallFunc28(void* func_) {
	return __cross_call_worker_CallFunc28(func_);
}

extern String (*__cross_call_worker_CallFunc29)(void*);

static String CallFunc29(void* func_) {
	return __cross_call_worker_CallFunc29(func_);
}

extern String (*__cross_call_worker_CallFunc30)(void*);

static String CallFunc30(void* func_) {
	return __cross_call_worker_CallFunc30(func_);
}

extern String (*__cross_call_worker_CallFunc31)(void*);

static String CallFunc31(void* func_) {
	return __cross_call_worker_CallFunc31(func_);
}

extern String (*__cross_call_worker_CallFunc32)(void*);

static String CallFunc32(void* func_) {
	return __cross_call_worker_CallFunc32(func_);
}

extern String (*__cross_call_worker_CallFunc33)(void*);

static String CallFunc33(void* func_) {
	return __cross_call_worker_CallFunc33(func_);
}

extern String (*__cross_call_worker_CallFuncEnum)(void*);

static String CallFuncEnum(void* func_) {
	return __cross_call_worker_CallFuncEnum(func_);
}

extern void (*__cross_call_worker_ReverseCall)(String*);

static void ReverseCall(String* test) {
	__cross_call_worker_ReverseCall(test);
}

