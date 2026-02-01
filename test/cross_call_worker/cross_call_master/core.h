#pragma once

#include "shared.h"

extern void (*__cross_call_master_ReverseReturn)(String*);

static void ReverseReturn(String* returnString) {
	__cross_call_master_ReverseReturn(returnString);
}

extern void (*__cross_call_master_NoParamReturnVoidCallback)();

static void NoParamReturnVoidCallback() {
	__cross_call_master_NoParamReturnVoidCallback();
}

extern bool (*__cross_call_master_NoParamReturnBoolCallback)();

static bool NoParamReturnBoolCallback() {
	return __cross_call_master_NoParamReturnBoolCallback();
}

extern int8_t (*__cross_call_master_NoParamReturnChar8Callback)();

static int8_t NoParamReturnChar8Callback() {
	return __cross_call_master_NoParamReturnChar8Callback();
}

extern uint16_t (*__cross_call_master_NoParamReturnChar16Callback)();

static uint16_t NoParamReturnChar16Callback() {
	return __cross_call_master_NoParamReturnChar16Callback();
}

extern int8_t (*__cross_call_master_NoParamReturnInt8Callback)();

static int8_t NoParamReturnInt8Callback() {
	return __cross_call_master_NoParamReturnInt8Callback();
}

extern int16_t (*__cross_call_master_NoParamReturnInt16Callback)();

static int16_t NoParamReturnInt16Callback() {
	return __cross_call_master_NoParamReturnInt16Callback();
}

extern int32_t (*__cross_call_master_NoParamReturnInt32Callback)();

static int32_t NoParamReturnInt32Callback() {
	return __cross_call_master_NoParamReturnInt32Callback();
}

extern int64_t (*__cross_call_master_NoParamReturnInt64Callback)();

static int64_t NoParamReturnInt64Callback() {
	return __cross_call_master_NoParamReturnInt64Callback();
}

extern uint8_t (*__cross_call_master_NoParamReturnUInt8Callback)();

static uint8_t NoParamReturnUInt8Callback() {
	return __cross_call_master_NoParamReturnUInt8Callback();
}

extern uint16_t (*__cross_call_master_NoParamReturnUInt16Callback)();

static uint16_t NoParamReturnUInt16Callback() {
	return __cross_call_master_NoParamReturnUInt16Callback();
}

extern uint32_t (*__cross_call_master_NoParamReturnUInt32Callback)();

static uint32_t NoParamReturnUInt32Callback() {
	return __cross_call_master_NoParamReturnUInt32Callback();
}

extern uint64_t (*__cross_call_master_NoParamReturnUInt64Callback)();

static uint64_t NoParamReturnUInt64Callback() {
	return __cross_call_master_NoParamReturnUInt64Callback();
}

extern uintptr_t (*__cross_call_master_NoParamReturnPointerCallback)();

static uintptr_t NoParamReturnPointerCallback() {
	return __cross_call_master_NoParamReturnPointerCallback();
}

extern float (*__cross_call_master_NoParamReturnFloatCallback)();

static float NoParamReturnFloatCallback() {
	return __cross_call_master_NoParamReturnFloatCallback();
}

extern double (*__cross_call_master_NoParamReturnDoubleCallback)();

static double NoParamReturnDoubleCallback() {
	return __cross_call_master_NoParamReturnDoubleCallback();
}

extern void* (*__cross_call_master_NoParamReturnFunctionCallback)();

static void* NoParamReturnFunctionCallback() {
	return __cross_call_master_NoParamReturnFunctionCallback();
}

extern String (*__cross_call_master_NoParamReturnStringCallback)();

static String NoParamReturnStringCallback() {
	return __cross_call_master_NoParamReturnStringCallback();
}

extern Variant (*__cross_call_master_NoParamReturnAnyCallback)();

static Variant NoParamReturnAnyCallback() {
	return __cross_call_master_NoParamReturnAnyCallback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayBoolCallback)();

static Vector NoParamReturnArrayBoolCallback() {
	return __cross_call_master_NoParamReturnArrayBoolCallback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayChar8Callback)();

static Vector NoParamReturnArrayChar8Callback() {
	return __cross_call_master_NoParamReturnArrayChar8Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayChar16Callback)();

static Vector NoParamReturnArrayChar16Callback() {
	return __cross_call_master_NoParamReturnArrayChar16Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayInt8Callback)();

static Vector NoParamReturnArrayInt8Callback() {
	return __cross_call_master_NoParamReturnArrayInt8Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayInt16Callback)();

static Vector NoParamReturnArrayInt16Callback() {
	return __cross_call_master_NoParamReturnArrayInt16Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayInt32Callback)();

static Vector NoParamReturnArrayInt32Callback() {
	return __cross_call_master_NoParamReturnArrayInt32Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayInt64Callback)();

static Vector NoParamReturnArrayInt64Callback() {
	return __cross_call_master_NoParamReturnArrayInt64Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayUInt8Callback)();

static Vector NoParamReturnArrayUInt8Callback() {
	return __cross_call_master_NoParamReturnArrayUInt8Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayUInt16Callback)();

static Vector NoParamReturnArrayUInt16Callback() {
	return __cross_call_master_NoParamReturnArrayUInt16Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayUInt32Callback)();

static Vector NoParamReturnArrayUInt32Callback() {
	return __cross_call_master_NoParamReturnArrayUInt32Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayUInt64Callback)();

static Vector NoParamReturnArrayUInt64Callback() {
	return __cross_call_master_NoParamReturnArrayUInt64Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayPointerCallback)();

static Vector NoParamReturnArrayPointerCallback() {
	return __cross_call_master_NoParamReturnArrayPointerCallback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayFloatCallback)();

static Vector NoParamReturnArrayFloatCallback() {
	return __cross_call_master_NoParamReturnArrayFloatCallback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayDoubleCallback)();

static Vector NoParamReturnArrayDoubleCallback() {
	return __cross_call_master_NoParamReturnArrayDoubleCallback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayStringCallback)();

static Vector NoParamReturnArrayStringCallback() {
	return __cross_call_master_NoParamReturnArrayStringCallback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayAnyCallback)();

static Vector NoParamReturnArrayAnyCallback() {
	return __cross_call_master_NoParamReturnArrayAnyCallback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayVector2Callback)();

static Vector NoParamReturnArrayVector2Callback() {
	return __cross_call_master_NoParamReturnArrayVector2Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayVector3Callback)();

static Vector NoParamReturnArrayVector3Callback() {
	return __cross_call_master_NoParamReturnArrayVector3Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayVector4Callback)();

static Vector NoParamReturnArrayVector4Callback() {
	return __cross_call_master_NoParamReturnArrayVector4Callback();
}

extern Vector (*__cross_call_master_NoParamReturnArrayMatrix4x4Callback)();

static Vector NoParamReturnArrayMatrix4x4Callback() {
	return __cross_call_master_NoParamReturnArrayMatrix4x4Callback();
}

extern Vector2 (*__cross_call_master_NoParamReturnVector2Callback)();

static Vector2 NoParamReturnVector2Callback() {
	return __cross_call_master_NoParamReturnVector2Callback();
}

extern Vector3 (*__cross_call_master_NoParamReturnVector3Callback)();

static Vector3 NoParamReturnVector3Callback() {
	return __cross_call_master_NoParamReturnVector3Callback();
}

extern Vector4 (*__cross_call_master_NoParamReturnVector4Callback)();

static Vector4 NoParamReturnVector4Callback() {
	return __cross_call_master_NoParamReturnVector4Callback();
}

extern Matrix4x4 (*__cross_call_master_NoParamReturnMatrix4x4Callback)();

static Matrix4x4 NoParamReturnMatrix4x4Callback() {
	return __cross_call_master_NoParamReturnMatrix4x4Callback();
}

extern void (*__cross_call_master_Param1Callback)(int32_t);

static void Param1Callback(int32_t a) {
	__cross_call_master_Param1Callback(a);
}

extern void (*__cross_call_master_Param2Callback)(int32_t, float);

static void Param2Callback(int32_t a, float b) {
	__cross_call_master_Param2Callback(a, b);
}

extern void (*__cross_call_master_Param3Callback)(int32_t, float, double);

static void Param3Callback(int32_t a, float b, double c) {
	__cross_call_master_Param3Callback(a, b, c);
}

extern void (*__cross_call_master_Param4Callback)(int32_t, float, double, Vector4*);

static void Param4Callback(int32_t a, float b, double c, Vector4* d) {
	__cross_call_master_Param4Callback(a, b, c, d);
}

extern void (*__cross_call_master_Param5Callback)(int32_t, float, double, Vector4*, Vector*);

static void Param5Callback(int32_t a, float b, double c, Vector4* d, Vector* e) {
	__cross_call_master_Param5Callback(a, b, c, d, e);
}

extern void (*__cross_call_master_Param6Callback)(int32_t, float, double, Vector4*, Vector*, int8_t);

static void Param6Callback(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f) {
	__cross_call_master_Param6Callback(a, b, c, d, e, f);
}

extern void (*__cross_call_master_Param7Callback)(int32_t, float, double, Vector4*, Vector*, int8_t, String*);

static void Param7Callback(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f, String* g) {
	__cross_call_master_Param7Callback(a, b, c, d, e, f, g);
}

extern void (*__cross_call_master_Param8Callback)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t);

static void Param8Callback(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f, String* g, uint16_t h) {
	__cross_call_master_Param8Callback(a, b, c, d, e, f, g, h);
}

extern void (*__cross_call_master_Param9Callback)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t, int16_t);

static void Param9Callback(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f, String* g, uint16_t h, int16_t k) {
	__cross_call_master_Param9Callback(a, b, c, d, e, f, g, h, k);
}

extern void (*__cross_call_master_Param10Callback)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t, int16_t, uintptr_t);

static void Param10Callback(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f, String* g, uint16_t h, int16_t k, uintptr_t l) {
	__cross_call_master_Param10Callback(a, b, c, d, e, f, g, h, k, l);
}

extern void (*__cross_call_master_ParamRef1Callback)(int32_t*);

static void ParamRef1Callback(int32_t* a) {
	__cross_call_master_ParamRef1Callback(a);
}

extern void (*__cross_call_master_ParamRef2Callback)(int32_t*, float*);

static void ParamRef2Callback(int32_t* a, float* b) {
	__cross_call_master_ParamRef2Callback(a, b);
}

extern void (*__cross_call_master_ParamRef3Callback)(int32_t*, float*, double*);

static void ParamRef3Callback(int32_t* a, float* b, double* c) {
	__cross_call_master_ParamRef3Callback(a, b, c);
}

extern void (*__cross_call_master_ParamRef4Callback)(int32_t*, float*, double*, Vector4*);

static void ParamRef4Callback(int32_t* a, float* b, double* c, Vector4* d) {
	__cross_call_master_ParamRef4Callback(a, b, c, d);
}

extern void (*__cross_call_master_ParamRef5Callback)(int32_t*, float*, double*, Vector4*, Vector*);

static void ParamRef5Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e) {
	__cross_call_master_ParamRef5Callback(a, b, c, d, e);
}

extern void (*__cross_call_master_ParamRef6Callback)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*);

static void ParamRef6Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f) {
	__cross_call_master_ParamRef6Callback(a, b, c, d, e, f);
}

extern void (*__cross_call_master_ParamRef7Callback)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*);

static void ParamRef7Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f, String* g) {
	__cross_call_master_ParamRef7Callback(a, b, c, d, e, f, g);
}

extern void (*__cross_call_master_ParamRef8Callback)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*);

static void ParamRef8Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f, String* g, uint16_t* h) {
	__cross_call_master_ParamRef8Callback(a, b, c, d, e, f, g, h);
}

extern void (*__cross_call_master_ParamRef9Callback)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*, int16_t*);

static void ParamRef9Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f, String* g, uint16_t* h, int16_t* k) {
	__cross_call_master_ParamRef9Callback(a, b, c, d, e, f, g, h, k);
}

extern void (*__cross_call_master_ParamRef10Callback)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*, int16_t*, uintptr_t*);

static void ParamRef10Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f, String* g, uint16_t* h, int16_t* k, uintptr_t* l) {
	__cross_call_master_ParamRef10Callback(a, b, c, d, e, f, g, h, k, l);
}

extern void (*__cross_call_master_ParamRefVectorsCallback)(Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*);

static void ParamRefVectorsCallback(Vector* p1, Vector* p2, Vector* p3, Vector* p4, Vector* p5, Vector* p6, Vector* p7, Vector* p8, Vector* p9, Vector* p10, Vector* p11, Vector* p12, Vector* p13, Vector* p14, Vector* p15) {
	__cross_call_master_ParamRefVectorsCallback(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15);
}

extern int64_t (*__cross_call_master_ParamAllPrimitivesCallback)(bool, int8_t, uint16_t, int8_t, int16_t, int32_t, int64_t, uint8_t, uint16_t, uint32_t, uint64_t, uintptr_t, float, double);

static int64_t ParamAllPrimitivesCallback(bool p1, int8_t p2, uint16_t p3, int8_t p4, int16_t p5, int32_t p6, int64_t p7, uint8_t p8, uint16_t p9, uint32_t p10, uint64_t p11, uintptr_t p12, float p13, double p14) {
	return __cross_call_master_ParamAllPrimitivesCallback(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14);
}

extern int32_t (*__cross_call_master_ParamAllAliasesCallback)(bool, int8_t, uint16_t, int8_t, int16_t, int32_t, int64_t, uintptr_t, float, double, String*, Variant*, Vector2*, Vector3*, Vector4*, Matrix4x4*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*);

static int32_t ParamAllAliasesCallback(bool aBool, int8_t aChar8, uint16_t aChar16, int8_t aInt8, int16_t aInt16, int32_t aInt32, int64_t aInt64, uintptr_t aPtr, float aFloat, double aDouble, String* aString, Variant* aAny, Vector2* aVec2, Vector3* aVec3, Vector4* aVec4, Matrix4x4* aMat4x4, Vector* aBoolVec, Vector* aChar8Vec, Vector* aChar16Vec, Vector* aInt8Vec, Vector* aInt16Vec, Vector* aInt32Vec, Vector* aInt64Vec, Vector* aPtrVec, Vector* aFloatVec, Vector* aDoubleVec, Vector* aStringVec, Vector* aAnyVec, Vector* aVec2Vec, Vector* aVec3Vec, Vector* aVec4Vec) {
	return __cross_call_master_ParamAllAliasesCallback(aBool, aChar8, aChar16, aInt8, aInt16, aInt32, aInt64, aPtr, aFloat, aDouble, aString, aAny, aVec2, aVec3, aVec4, aMat4x4, aBoolVec, aChar8Vec, aChar16Vec, aInt8Vec, aInt16Vec, aInt32Vec, aInt64Vec, aPtrVec, aFloatVec, aDoubleVec, aStringVec, aAnyVec, aVec2Vec, aVec3Vec, aVec4Vec);
}

extern int64_t (*__cross_call_master_ParamAllRefAliasesCallback)(bool*, int8_t*, uint16_t*, int8_t*, int16_t*, int32_t*, int64_t*, uintptr_t*, float*, double*, String*, Variant*, Vector2*, Vector3*, Vector4*, Matrix4x4*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*);

static int64_t ParamAllRefAliasesCallback(bool* aBool, int8_t* aChar8, uint16_t* aChar16, int8_t* aInt8, int16_t* aInt16, int32_t* aInt32, int64_t* aInt64, uintptr_t* aPtr, float* aFloat, double* aDouble, String* aString, Variant* aAny, Vector2* aVec2, Vector3* aVec3, Vector4* aVec4, Matrix4x4* aMat4x4, Vector* aBoolVec, Vector* aChar8Vec, Vector* aChar16Vec, Vector* aInt8Vec, Vector* aInt16Vec, Vector* aInt32Vec, Vector* aInt64Vec, Vector* aPtrVec, Vector* aFloatVec, Vector* aDoubleVec, Vector* aStringVec, Vector* aAnyVec, Vector* aVec2Vec, Vector* aVec3Vec, Vector* aVec4Vec) {
	return __cross_call_master_ParamAllRefAliasesCallback(aBool, aChar8, aChar16, aInt8, aInt16, aInt32, aInt64, aPtr, aFloat, aDouble, aString, aAny, aVec2, aVec3, aVec4, aMat4x4, aBoolVec, aChar8Vec, aChar16Vec, aInt8Vec, aInt16Vec, aInt32Vec, aInt64Vec, aPtrVec, aFloatVec, aDoubleVec, aStringVec, aAnyVec, aVec2Vec, aVec3Vec, aVec4Vec);
}

extern int32_t (*__cross_call_master_ParamEnumCallback)(int32_t, Vector*);

static int32_t ParamEnumCallback(int32_t p1, Vector* p2) {
	return __cross_call_master_ParamEnumCallback(p1, p2);
}

extern int32_t (*__cross_call_master_ParamEnumRefCallback)(int32_t*, Vector*);

static int32_t ParamEnumRefCallback(int32_t* p1, Vector* p2) {
	return __cross_call_master_ParamEnumRefCallback(p1, p2);
}

extern void (*__cross_call_master_ParamVariantCallback)(Variant*, Vector*);

static void ParamVariantCallback(Variant* p1, Vector* p2) {
	__cross_call_master_ParamVariantCallback(p1, p2);
}

extern void (*__cross_call_master_ParamVariantRefCallback)(Variant*, Vector*);

static void ParamVariantRefCallback(Variant* p1, Vector* p2) {
	__cross_call_master_ParamVariantRefCallback(p1, p2);
}

extern void (*__cross_call_master_CallFuncVoidCallback)(void*);

static void CallFuncVoidCallback(void* func_) {
	__cross_call_master_CallFuncVoidCallback(func_);
}

extern bool (*__cross_call_master_CallFuncBoolCallback)(void*);

static bool CallFuncBoolCallback(void* func_) {
	return __cross_call_master_CallFuncBoolCallback(func_);
}

extern int8_t (*__cross_call_master_CallFuncChar8Callback)(void*);

static int8_t CallFuncChar8Callback(void* func_) {
	return __cross_call_master_CallFuncChar8Callback(func_);
}

extern uint16_t (*__cross_call_master_CallFuncChar16Callback)(void*);

static uint16_t CallFuncChar16Callback(void* func_) {
	return __cross_call_master_CallFuncChar16Callback(func_);
}

extern int8_t (*__cross_call_master_CallFuncInt8Callback)(void*);

static int8_t CallFuncInt8Callback(void* func_) {
	return __cross_call_master_CallFuncInt8Callback(func_);
}

extern int16_t (*__cross_call_master_CallFuncInt16Callback)(void*);

static int16_t CallFuncInt16Callback(void* func_) {
	return __cross_call_master_CallFuncInt16Callback(func_);
}

extern int32_t (*__cross_call_master_CallFuncInt32Callback)(void*);

static int32_t CallFuncInt32Callback(void* func_) {
	return __cross_call_master_CallFuncInt32Callback(func_);
}

extern int64_t (*__cross_call_master_CallFuncInt64Callback)(void*);

static int64_t CallFuncInt64Callback(void* func_) {
	return __cross_call_master_CallFuncInt64Callback(func_);
}

extern uint8_t (*__cross_call_master_CallFuncUInt8Callback)(void*);

static uint8_t CallFuncUInt8Callback(void* func_) {
	return __cross_call_master_CallFuncUInt8Callback(func_);
}

extern uint16_t (*__cross_call_master_CallFuncUInt16Callback)(void*);

static uint16_t CallFuncUInt16Callback(void* func_) {
	return __cross_call_master_CallFuncUInt16Callback(func_);
}

extern uint32_t (*__cross_call_master_CallFuncUInt32Callback)(void*);

static uint32_t CallFuncUInt32Callback(void* func_) {
	return __cross_call_master_CallFuncUInt32Callback(func_);
}

extern uint64_t (*__cross_call_master_CallFuncUInt64Callback)(void*);

static uint64_t CallFuncUInt64Callback(void* func_) {
	return __cross_call_master_CallFuncUInt64Callback(func_);
}

extern uintptr_t (*__cross_call_master_CallFuncPtrCallback)(void*);

static uintptr_t CallFuncPtrCallback(void* func_) {
	return __cross_call_master_CallFuncPtrCallback(func_);
}

extern float (*__cross_call_master_CallFuncFloatCallback)(void*);

static float CallFuncFloatCallback(void* func_) {
	return __cross_call_master_CallFuncFloatCallback(func_);
}

extern double (*__cross_call_master_CallFuncDoubleCallback)(void*);

static double CallFuncDoubleCallback(void* func_) {
	return __cross_call_master_CallFuncDoubleCallback(func_);
}

extern String (*__cross_call_master_CallFuncStringCallback)(void*);

static String CallFuncStringCallback(void* func_) {
	return __cross_call_master_CallFuncStringCallback(func_);
}

extern Variant (*__cross_call_master_CallFuncAnyCallback)(void*);

static Variant CallFuncAnyCallback(void* func_) {
	return __cross_call_master_CallFuncAnyCallback(func_);
}

extern uintptr_t (*__cross_call_master_CallFuncFunctionCallback)(void*);

static uintptr_t CallFuncFunctionCallback(void* func_) {
	return __cross_call_master_CallFuncFunctionCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncBoolVectorCallback)(void*);

static Vector CallFuncBoolVectorCallback(void* func_) {
	return __cross_call_master_CallFuncBoolVectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncChar8VectorCallback)(void*);

static Vector CallFuncChar8VectorCallback(void* func_) {
	return __cross_call_master_CallFuncChar8VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncChar16VectorCallback)(void*);

static Vector CallFuncChar16VectorCallback(void* func_) {
	return __cross_call_master_CallFuncChar16VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncInt8VectorCallback)(void*);

static Vector CallFuncInt8VectorCallback(void* func_) {
	return __cross_call_master_CallFuncInt8VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncInt16VectorCallback)(void*);

static Vector CallFuncInt16VectorCallback(void* func_) {
	return __cross_call_master_CallFuncInt16VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncInt32VectorCallback)(void*);

static Vector CallFuncInt32VectorCallback(void* func_) {
	return __cross_call_master_CallFuncInt32VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncInt64VectorCallback)(void*);

static Vector CallFuncInt64VectorCallback(void* func_) {
	return __cross_call_master_CallFuncInt64VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncUInt8VectorCallback)(void*);

static Vector CallFuncUInt8VectorCallback(void* func_) {
	return __cross_call_master_CallFuncUInt8VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncUInt16VectorCallback)(void*);

static Vector CallFuncUInt16VectorCallback(void* func_) {
	return __cross_call_master_CallFuncUInt16VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncUInt32VectorCallback)(void*);

static Vector CallFuncUInt32VectorCallback(void* func_) {
	return __cross_call_master_CallFuncUInt32VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncUInt64VectorCallback)(void*);

static Vector CallFuncUInt64VectorCallback(void* func_) {
	return __cross_call_master_CallFuncUInt64VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncPtrVectorCallback)(void*);

static Vector CallFuncPtrVectorCallback(void* func_) {
	return __cross_call_master_CallFuncPtrVectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncFloatVectorCallback)(void*);

static Vector CallFuncFloatVectorCallback(void* func_) {
	return __cross_call_master_CallFuncFloatVectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncDoubleVectorCallback)(void*);

static Vector CallFuncDoubleVectorCallback(void* func_) {
	return __cross_call_master_CallFuncDoubleVectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncStringVectorCallback)(void*);

static Vector CallFuncStringVectorCallback(void* func_) {
	return __cross_call_master_CallFuncStringVectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAnyVectorCallback)(void*);

static Vector CallFuncAnyVectorCallback(void* func_) {
	return __cross_call_master_CallFuncAnyVectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncVec2VectorCallback)(void*);

static Vector CallFuncVec2VectorCallback(void* func_) {
	return __cross_call_master_CallFuncVec2VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncVec3VectorCallback)(void*);

static Vector CallFuncVec3VectorCallback(void* func_) {
	return __cross_call_master_CallFuncVec3VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncVec4VectorCallback)(void*);

static Vector CallFuncVec4VectorCallback(void* func_) {
	return __cross_call_master_CallFuncVec4VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncMat4x4VectorCallback)(void*);

static Vector CallFuncMat4x4VectorCallback(void* func_) {
	return __cross_call_master_CallFuncMat4x4VectorCallback(func_);
}

extern Vector2 (*__cross_call_master_CallFuncVec2Callback)(void*);

static Vector2 CallFuncVec2Callback(void* func_) {
	return __cross_call_master_CallFuncVec2Callback(func_);
}

extern Vector3 (*__cross_call_master_CallFuncVec3Callback)(void*);

static Vector3 CallFuncVec3Callback(void* func_) {
	return __cross_call_master_CallFuncVec3Callback(func_);
}

extern Vector4 (*__cross_call_master_CallFuncVec4Callback)(void*);

static Vector4 CallFuncVec4Callback(void* func_) {
	return __cross_call_master_CallFuncVec4Callback(func_);
}

extern Matrix4x4 (*__cross_call_master_CallFuncMat4x4Callback)(void*);

static Matrix4x4 CallFuncMat4x4Callback(void* func_) {
	return __cross_call_master_CallFuncMat4x4Callback(func_);
}

extern bool (*__cross_call_master_CallFuncAliasBoolCallback)(void*);

static bool CallFuncAliasBoolCallback(void* func_) {
	return __cross_call_master_CallFuncAliasBoolCallback(func_);
}

extern int8_t (*__cross_call_master_CallFuncAliasChar8Callback)(void*);

static int8_t CallFuncAliasChar8Callback(void* func_) {
	return __cross_call_master_CallFuncAliasChar8Callback(func_);
}

extern uint16_t (*__cross_call_master_CallFuncAliasChar16Callback)(void*);

static uint16_t CallFuncAliasChar16Callback(void* func_) {
	return __cross_call_master_CallFuncAliasChar16Callback(func_);
}

extern int8_t (*__cross_call_master_CallFuncAliasInt8Callback)(void*);

static int8_t CallFuncAliasInt8Callback(void* func_) {
	return __cross_call_master_CallFuncAliasInt8Callback(func_);
}

extern int16_t (*__cross_call_master_CallFuncAliasInt16Callback)(void*);

static int16_t CallFuncAliasInt16Callback(void* func_) {
	return __cross_call_master_CallFuncAliasInt16Callback(func_);
}

extern int32_t (*__cross_call_master_CallFuncAliasInt32Callback)(void*);

static int32_t CallFuncAliasInt32Callback(void* func_) {
	return __cross_call_master_CallFuncAliasInt32Callback(func_);
}

extern int64_t (*__cross_call_master_CallFuncAliasInt64Callback)(void*);

static int64_t CallFuncAliasInt64Callback(void* func_) {
	return __cross_call_master_CallFuncAliasInt64Callback(func_);
}

extern uint8_t (*__cross_call_master_CallFuncAliasUInt8Callback)(void*);

static uint8_t CallFuncAliasUInt8Callback(void* func_) {
	return __cross_call_master_CallFuncAliasUInt8Callback(func_);
}

extern uint16_t (*__cross_call_master_CallFuncAliasUInt16Callback)(void*);

static uint16_t CallFuncAliasUInt16Callback(void* func_) {
	return __cross_call_master_CallFuncAliasUInt16Callback(func_);
}

extern uint32_t (*__cross_call_master_CallFuncAliasUInt32Callback)(void*);

static uint32_t CallFuncAliasUInt32Callback(void* func_) {
	return __cross_call_master_CallFuncAliasUInt32Callback(func_);
}

extern uint64_t (*__cross_call_master_CallFuncAliasUInt64Callback)(void*);

static uint64_t CallFuncAliasUInt64Callback(void* func_) {
	return __cross_call_master_CallFuncAliasUInt64Callback(func_);
}

extern uintptr_t (*__cross_call_master_CallFuncAliasPtrCallback)(void*);

static uintptr_t CallFuncAliasPtrCallback(void* func_) {
	return __cross_call_master_CallFuncAliasPtrCallback(func_);
}

extern float (*__cross_call_master_CallFuncAliasFloatCallback)(void*);

static float CallFuncAliasFloatCallback(void* func_) {
	return __cross_call_master_CallFuncAliasFloatCallback(func_);
}

extern double (*__cross_call_master_CallFuncAliasDoubleCallback)(void*);

static double CallFuncAliasDoubleCallback(void* func_) {
	return __cross_call_master_CallFuncAliasDoubleCallback(func_);
}

extern String (*__cross_call_master_CallFuncAliasStringCallback)(void*);

static String CallFuncAliasStringCallback(void* func_) {
	return __cross_call_master_CallFuncAliasStringCallback(func_);
}

extern Variant (*__cross_call_master_CallFuncAliasAnyCallback)(void*);

static Variant CallFuncAliasAnyCallback(void* func_) {
	return __cross_call_master_CallFuncAliasAnyCallback(func_);
}

extern uintptr_t (*__cross_call_master_CallFuncAliasFunctionCallback)(void*);

static uintptr_t CallFuncAliasFunctionCallback(void* func_) {
	return __cross_call_master_CallFuncAliasFunctionCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasBoolVectorCallback)(void*);

static Vector CallFuncAliasBoolVectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasBoolVectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasChar8VectorCallback)(void*);

static Vector CallFuncAliasChar8VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasChar8VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasChar16VectorCallback)(void*);

static Vector CallFuncAliasChar16VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasChar16VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasInt8VectorCallback)(void*);

static Vector CallFuncAliasInt8VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasInt8VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasInt16VectorCallback)(void*);

static Vector CallFuncAliasInt16VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasInt16VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasInt32VectorCallback)(void*);

static Vector CallFuncAliasInt32VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasInt32VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasInt64VectorCallback)(void*);

static Vector CallFuncAliasInt64VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasInt64VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasUInt8VectorCallback)(void*);

static Vector CallFuncAliasUInt8VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasUInt8VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasUInt16VectorCallback)(void*);

static Vector CallFuncAliasUInt16VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasUInt16VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasUInt32VectorCallback)(void*);

static Vector CallFuncAliasUInt32VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasUInt32VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasUInt64VectorCallback)(void*);

static Vector CallFuncAliasUInt64VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasUInt64VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasPtrVectorCallback)(void*);

static Vector CallFuncAliasPtrVectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasPtrVectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasFloatVectorCallback)(void*);

static Vector CallFuncAliasFloatVectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasFloatVectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasDoubleVectorCallback)(void*);

static Vector CallFuncAliasDoubleVectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasDoubleVectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasStringVectorCallback)(void*);

static Vector CallFuncAliasStringVectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasStringVectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasAnyVectorCallback)(void*);

static Vector CallFuncAliasAnyVectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasAnyVectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasVec2VectorCallback)(void*);

static Vector CallFuncAliasVec2VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasVec2VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasVec3VectorCallback)(void*);

static Vector CallFuncAliasVec3VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasVec3VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasVec4VectorCallback)(void*);

static Vector CallFuncAliasVec4VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasVec4VectorCallback(func_);
}

extern Vector (*__cross_call_master_CallFuncAliasMat4x4VectorCallback)(void*);

static Vector CallFuncAliasMat4x4VectorCallback(void* func_) {
	return __cross_call_master_CallFuncAliasMat4x4VectorCallback(func_);
}

extern Vector2 (*__cross_call_master_CallFuncAliasVec2Callback)(void*);

static Vector2 CallFuncAliasVec2Callback(void* func_) {
	return __cross_call_master_CallFuncAliasVec2Callback(func_);
}

extern Vector3 (*__cross_call_master_CallFuncAliasVec3Callback)(void*);

static Vector3 CallFuncAliasVec3Callback(void* func_) {
	return __cross_call_master_CallFuncAliasVec3Callback(func_);
}

extern Vector4 (*__cross_call_master_CallFuncAliasVec4Callback)(void*);

static Vector4 CallFuncAliasVec4Callback(void* func_) {
	return __cross_call_master_CallFuncAliasVec4Callback(func_);
}

extern Matrix4x4 (*__cross_call_master_CallFuncAliasMat4x4Callback)(void*);

static Matrix4x4 CallFuncAliasMat4x4Callback(void* func_) {
	return __cross_call_master_CallFuncAliasMat4x4Callback(func_);
}

extern String (*__cross_call_master_CallFuncAliasAllCallback)(void*);

static String CallFuncAliasAllCallback(void* func_) {
	return __cross_call_master_CallFuncAliasAllCallback(func_);
}

extern int32_t (*__cross_call_master_CallFunc1Callback)(void*);

static int32_t CallFunc1Callback(void* func_) {
	return __cross_call_master_CallFunc1Callback(func_);
}

extern int8_t (*__cross_call_master_CallFunc2Callback)(void*);

static int8_t CallFunc2Callback(void* func_) {
	return __cross_call_master_CallFunc2Callback(func_);
}

extern void (*__cross_call_master_CallFunc3Callback)(void*);

static void CallFunc3Callback(void* func_) {
	__cross_call_master_CallFunc3Callback(func_);
}

extern Vector4 (*__cross_call_master_CallFunc4Callback)(void*);

static Vector4 CallFunc4Callback(void* func_) {
	return __cross_call_master_CallFunc4Callback(func_);
}

extern bool (*__cross_call_master_CallFunc5Callback)(void*);

static bool CallFunc5Callback(void* func_) {
	return __cross_call_master_CallFunc5Callback(func_);
}

extern int64_t (*__cross_call_master_CallFunc6Callback)(void*);

static int64_t CallFunc6Callback(void* func_) {
	return __cross_call_master_CallFunc6Callback(func_);
}

extern double (*__cross_call_master_CallFunc7Callback)(void*);

static double CallFunc7Callback(void* func_) {
	return __cross_call_master_CallFunc7Callback(func_);
}

extern Matrix4x4 (*__cross_call_master_CallFunc8Callback)(void*);

static Matrix4x4 CallFunc8Callback(void* func_) {
	return __cross_call_master_CallFunc8Callback(func_);
}

extern void (*__cross_call_master_CallFunc9Callback)(void*);

static void CallFunc9Callback(void* func_) {
	__cross_call_master_CallFunc9Callback(func_);
}

extern uint32_t (*__cross_call_master_CallFunc10Callback)(void*);

static uint32_t CallFunc10Callback(void* func_) {
	return __cross_call_master_CallFunc10Callback(func_);
}

extern uintptr_t (*__cross_call_master_CallFunc11Callback)(void*);

static uintptr_t CallFunc11Callback(void* func_) {
	return __cross_call_master_CallFunc11Callback(func_);
}

extern bool (*__cross_call_master_CallFunc12Callback)(void*);

static bool CallFunc12Callback(void* func_) {
	return __cross_call_master_CallFunc12Callback(func_);
}

extern String (*__cross_call_master_CallFunc13Callback)(void*);

static String CallFunc13Callback(void* func_) {
	return __cross_call_master_CallFunc13Callback(func_);
}

extern Vector (*__cross_call_master_CallFunc14Callback)(void*);

static Vector CallFunc14Callback(void* func_) {
	return __cross_call_master_CallFunc14Callback(func_);
}

extern int16_t (*__cross_call_master_CallFunc15Callback)(void*);

static int16_t CallFunc15Callback(void* func_) {
	return __cross_call_master_CallFunc15Callback(func_);
}

extern uintptr_t (*__cross_call_master_CallFunc16Callback)(void*);

static uintptr_t CallFunc16Callback(void* func_) {
	return __cross_call_master_CallFunc16Callback(func_);
}

extern String (*__cross_call_master_CallFunc17Callback)(void*);

static String CallFunc17Callback(void* func_) {
	return __cross_call_master_CallFunc17Callback(func_);
}

extern String (*__cross_call_master_CallFunc18Callback)(void*);

static String CallFunc18Callback(void* func_) {
	return __cross_call_master_CallFunc18Callback(func_);
}

extern String (*__cross_call_master_CallFunc19Callback)(void*);

static String CallFunc19Callback(void* func_) {
	return __cross_call_master_CallFunc19Callback(func_);
}

extern String (*__cross_call_master_CallFunc20Callback)(void*);

static String CallFunc20Callback(void* func_) {
	return __cross_call_master_CallFunc20Callback(func_);
}

extern String (*__cross_call_master_CallFunc21Callback)(void*);

static String CallFunc21Callback(void* func_) {
	return __cross_call_master_CallFunc21Callback(func_);
}

extern String (*__cross_call_master_CallFunc22Callback)(void*);

static String CallFunc22Callback(void* func_) {
	return __cross_call_master_CallFunc22Callback(func_);
}

extern String (*__cross_call_master_CallFunc23Callback)(void*);

static String CallFunc23Callback(void* func_) {
	return __cross_call_master_CallFunc23Callback(func_);
}

extern String (*__cross_call_master_CallFunc24Callback)(void*);

static String CallFunc24Callback(void* func_) {
	return __cross_call_master_CallFunc24Callback(func_);
}

extern String (*__cross_call_master_CallFunc25Callback)(void*);

static String CallFunc25Callback(void* func_) {
	return __cross_call_master_CallFunc25Callback(func_);
}

extern String (*__cross_call_master_CallFunc26Callback)(void*);

static String CallFunc26Callback(void* func_) {
	return __cross_call_master_CallFunc26Callback(func_);
}

extern String (*__cross_call_master_CallFunc27Callback)(void*);

static String CallFunc27Callback(void* func_) {
	return __cross_call_master_CallFunc27Callback(func_);
}

extern String (*__cross_call_master_CallFunc28Callback)(void*);

static String CallFunc28Callback(void* func_) {
	return __cross_call_master_CallFunc28Callback(func_);
}

extern String (*__cross_call_master_CallFunc29Callback)(void*);

static String CallFunc29Callback(void* func_) {
	return __cross_call_master_CallFunc29Callback(func_);
}

extern String (*__cross_call_master_CallFunc30Callback)(void*);

static String CallFunc30Callback(void* func_) {
	return __cross_call_master_CallFunc30Callback(func_);
}

extern String (*__cross_call_master_CallFunc31Callback)(void*);

static String CallFunc31Callback(void* func_) {
	return __cross_call_master_CallFunc31Callback(func_);
}

extern String (*__cross_call_master_CallFunc32Callback)(void*);

static String CallFunc32Callback(void* func_) {
	return __cross_call_master_CallFunc32Callback(func_);
}

extern String (*__cross_call_master_CallFunc33Callback)(void*);

static String CallFunc33Callback(void* func_) {
	return __cross_call_master_CallFunc33Callback(func_);
}

extern String (*__cross_call_master_CallFuncEnumCallback)(void*);

static String CallFuncEnumCallback(void* func_) {
	return __cross_call_master_CallFuncEnumCallback(func_);
}

