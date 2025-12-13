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

static void CallFuncVoidCallback(void* func) {
	__cross_call_master_CallFuncVoidCallback(func);
}

extern bool (*__cross_call_master_CallFuncBoolCallback)(void*);

static bool CallFuncBoolCallback(void* func) {
	return __cross_call_master_CallFuncBoolCallback(func);
}

extern int8_t (*__cross_call_master_CallFuncChar8Callback)(void*);

static int8_t CallFuncChar8Callback(void* func) {
	return __cross_call_master_CallFuncChar8Callback(func);
}

extern uint16_t (*__cross_call_master_CallFuncChar16Callback)(void*);

static uint16_t CallFuncChar16Callback(void* func) {
	return __cross_call_master_CallFuncChar16Callback(func);
}

extern int8_t (*__cross_call_master_CallFuncInt8Callback)(void*);

static int8_t CallFuncInt8Callback(void* func) {
	return __cross_call_master_CallFuncInt8Callback(func);
}

extern int16_t (*__cross_call_master_CallFuncInt16Callback)(void*);

static int16_t CallFuncInt16Callback(void* func) {
	return __cross_call_master_CallFuncInt16Callback(func);
}

extern int32_t (*__cross_call_master_CallFuncInt32Callback)(void*);

static int32_t CallFuncInt32Callback(void* func) {
	return __cross_call_master_CallFuncInt32Callback(func);
}

extern int64_t (*__cross_call_master_CallFuncInt64Callback)(void*);

static int64_t CallFuncInt64Callback(void* func) {
	return __cross_call_master_CallFuncInt64Callback(func);
}

extern uint8_t (*__cross_call_master_CallFuncUInt8Callback)(void*);

static uint8_t CallFuncUInt8Callback(void* func) {
	return __cross_call_master_CallFuncUInt8Callback(func);
}

extern uint16_t (*__cross_call_master_CallFuncUInt16Callback)(void*);

static uint16_t CallFuncUInt16Callback(void* func) {
	return __cross_call_master_CallFuncUInt16Callback(func);
}

extern uint32_t (*__cross_call_master_CallFuncUInt32Callback)(void*);

static uint32_t CallFuncUInt32Callback(void* func) {
	return __cross_call_master_CallFuncUInt32Callback(func);
}

extern uint64_t (*__cross_call_master_CallFuncUInt64Callback)(void*);

static uint64_t CallFuncUInt64Callback(void* func) {
	return __cross_call_master_CallFuncUInt64Callback(func);
}

extern uintptr_t (*__cross_call_master_CallFuncPtrCallback)(void*);

static uintptr_t CallFuncPtrCallback(void* func) {
	return __cross_call_master_CallFuncPtrCallback(func);
}

extern float (*__cross_call_master_CallFuncFloatCallback)(void*);

static float CallFuncFloatCallback(void* func) {
	return __cross_call_master_CallFuncFloatCallback(func);
}

extern double (*__cross_call_master_CallFuncDoubleCallback)(void*);

static double CallFuncDoubleCallback(void* func) {
	return __cross_call_master_CallFuncDoubleCallback(func);
}

extern String (*__cross_call_master_CallFuncStringCallback)(void*);

static String CallFuncStringCallback(void* func) {
	return __cross_call_master_CallFuncStringCallback(func);
}

extern Variant (*__cross_call_master_CallFuncAnyCallback)(void*);

static Variant CallFuncAnyCallback(void* func) {
	return __cross_call_master_CallFuncAnyCallback(func);
}

extern uintptr_t (*__cross_call_master_CallFuncFunctionCallback)(void*);

static uintptr_t CallFuncFunctionCallback(void* func) {
	return __cross_call_master_CallFuncFunctionCallback(func);
}

extern Vector (*__cross_call_master_CallFuncBoolVectorCallback)(void*);

static Vector CallFuncBoolVectorCallback(void* func) {
	return __cross_call_master_CallFuncBoolVectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncChar8VectorCallback)(void*);

static Vector CallFuncChar8VectorCallback(void* func) {
	return __cross_call_master_CallFuncChar8VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncChar16VectorCallback)(void*);

static Vector CallFuncChar16VectorCallback(void* func) {
	return __cross_call_master_CallFuncChar16VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncInt8VectorCallback)(void*);

static Vector CallFuncInt8VectorCallback(void* func) {
	return __cross_call_master_CallFuncInt8VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncInt16VectorCallback)(void*);

static Vector CallFuncInt16VectorCallback(void* func) {
	return __cross_call_master_CallFuncInt16VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncInt32VectorCallback)(void*);

static Vector CallFuncInt32VectorCallback(void* func) {
	return __cross_call_master_CallFuncInt32VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncInt64VectorCallback)(void*);

static Vector CallFuncInt64VectorCallback(void* func) {
	return __cross_call_master_CallFuncInt64VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncUInt8VectorCallback)(void*);

static Vector CallFuncUInt8VectorCallback(void* func) {
	return __cross_call_master_CallFuncUInt8VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncUInt16VectorCallback)(void*);

static Vector CallFuncUInt16VectorCallback(void* func) {
	return __cross_call_master_CallFuncUInt16VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncUInt32VectorCallback)(void*);

static Vector CallFuncUInt32VectorCallback(void* func) {
	return __cross_call_master_CallFuncUInt32VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncUInt64VectorCallback)(void*);

static Vector CallFuncUInt64VectorCallback(void* func) {
	return __cross_call_master_CallFuncUInt64VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncPtrVectorCallback)(void*);

static Vector CallFuncPtrVectorCallback(void* func) {
	return __cross_call_master_CallFuncPtrVectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncFloatVectorCallback)(void*);

static Vector CallFuncFloatVectorCallback(void* func) {
	return __cross_call_master_CallFuncFloatVectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncDoubleVectorCallback)(void*);

static Vector CallFuncDoubleVectorCallback(void* func) {
	return __cross_call_master_CallFuncDoubleVectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncStringVectorCallback)(void*);

static Vector CallFuncStringVectorCallback(void* func) {
	return __cross_call_master_CallFuncStringVectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncAnyVectorCallback)(void*);

static Vector CallFuncAnyVectorCallback(void* func) {
	return __cross_call_master_CallFuncAnyVectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncVec2VectorCallback)(void*);

static Vector CallFuncVec2VectorCallback(void* func) {
	return __cross_call_master_CallFuncVec2VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncVec3VectorCallback)(void*);

static Vector CallFuncVec3VectorCallback(void* func) {
	return __cross_call_master_CallFuncVec3VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncVec4VectorCallback)(void*);

static Vector CallFuncVec4VectorCallback(void* func) {
	return __cross_call_master_CallFuncVec4VectorCallback(func);
}

extern Vector (*__cross_call_master_CallFuncMat4x4VectorCallback)(void*);

static Vector CallFuncMat4x4VectorCallback(void* func) {
	return __cross_call_master_CallFuncMat4x4VectorCallback(func);
}

extern Vector2 (*__cross_call_master_CallFuncVec2Callback)(void*);

static Vector2 CallFuncVec2Callback(void* func) {
	return __cross_call_master_CallFuncVec2Callback(func);
}

extern Vector3 (*__cross_call_master_CallFuncVec3Callback)(void*);

static Vector3 CallFuncVec3Callback(void* func) {
	return __cross_call_master_CallFuncVec3Callback(func);
}

extern Vector4 (*__cross_call_master_CallFuncVec4Callback)(void*);

static Vector4 CallFuncVec4Callback(void* func) {
	return __cross_call_master_CallFuncVec4Callback(func);
}

extern Matrix4x4 (*__cross_call_master_CallFuncMat4x4Callback)(void*);

static Matrix4x4 CallFuncMat4x4Callback(void* func) {
	return __cross_call_master_CallFuncMat4x4Callback(func);
}

extern int32_t (*__cross_call_master_CallFunc1Callback)(void*);

static int32_t CallFunc1Callback(void* func) {
	return __cross_call_master_CallFunc1Callback(func);
}

extern int8_t (*__cross_call_master_CallFunc2Callback)(void*);

static int8_t CallFunc2Callback(void* func) {
	return __cross_call_master_CallFunc2Callback(func);
}

extern void (*__cross_call_master_CallFunc3Callback)(void*);

static void CallFunc3Callback(void* func) {
	__cross_call_master_CallFunc3Callback(func);
}

extern Vector4 (*__cross_call_master_CallFunc4Callback)(void*);

static Vector4 CallFunc4Callback(void* func) {
	return __cross_call_master_CallFunc4Callback(func);
}

extern bool (*__cross_call_master_CallFunc5Callback)(void*);

static bool CallFunc5Callback(void* func) {
	return __cross_call_master_CallFunc5Callback(func);
}

extern int64_t (*__cross_call_master_CallFunc6Callback)(void*);

static int64_t CallFunc6Callback(void* func) {
	return __cross_call_master_CallFunc6Callback(func);
}

extern double (*__cross_call_master_CallFunc7Callback)(void*);

static double CallFunc7Callback(void* func) {
	return __cross_call_master_CallFunc7Callback(func);
}

extern Matrix4x4 (*__cross_call_master_CallFunc8Callback)(void*);

static Matrix4x4 CallFunc8Callback(void* func) {
	return __cross_call_master_CallFunc8Callback(func);
}

extern void (*__cross_call_master_CallFunc9Callback)(void*);

static void CallFunc9Callback(void* func) {
	__cross_call_master_CallFunc9Callback(func);
}

extern uint32_t (*__cross_call_master_CallFunc10Callback)(void*);

static uint32_t CallFunc10Callback(void* func) {
	return __cross_call_master_CallFunc10Callback(func);
}

extern uintptr_t (*__cross_call_master_CallFunc11Callback)(void*);

static uintptr_t CallFunc11Callback(void* func) {
	return __cross_call_master_CallFunc11Callback(func);
}

extern bool (*__cross_call_master_CallFunc12Callback)(void*);

static bool CallFunc12Callback(void* func) {
	return __cross_call_master_CallFunc12Callback(func);
}

extern String (*__cross_call_master_CallFunc13Callback)(void*);

static String CallFunc13Callback(void* func) {
	return __cross_call_master_CallFunc13Callback(func);
}

extern Vector (*__cross_call_master_CallFunc14Callback)(void*);

static Vector CallFunc14Callback(void* func) {
	return __cross_call_master_CallFunc14Callback(func);
}

extern int16_t (*__cross_call_master_CallFunc15Callback)(void*);

static int16_t CallFunc15Callback(void* func) {
	return __cross_call_master_CallFunc15Callback(func);
}

extern uintptr_t (*__cross_call_master_CallFunc16Callback)(void*);

static uintptr_t CallFunc16Callback(void* func) {
	return __cross_call_master_CallFunc16Callback(func);
}

extern String (*__cross_call_master_CallFunc17Callback)(void*);

static String CallFunc17Callback(void* func) {
	return __cross_call_master_CallFunc17Callback(func);
}

extern String (*__cross_call_master_CallFunc18Callback)(void*);

static String CallFunc18Callback(void* func) {
	return __cross_call_master_CallFunc18Callback(func);
}

extern String (*__cross_call_master_CallFunc19Callback)(void*);

static String CallFunc19Callback(void* func) {
	return __cross_call_master_CallFunc19Callback(func);
}

extern String (*__cross_call_master_CallFunc20Callback)(void*);

static String CallFunc20Callback(void* func) {
	return __cross_call_master_CallFunc20Callback(func);
}

extern String (*__cross_call_master_CallFunc21Callback)(void*);

static String CallFunc21Callback(void* func) {
	return __cross_call_master_CallFunc21Callback(func);
}

extern String (*__cross_call_master_CallFunc22Callback)(void*);

static String CallFunc22Callback(void* func) {
	return __cross_call_master_CallFunc22Callback(func);
}

extern String (*__cross_call_master_CallFunc23Callback)(void*);

static String CallFunc23Callback(void* func) {
	return __cross_call_master_CallFunc23Callback(func);
}

extern String (*__cross_call_master_CallFunc24Callback)(void*);

static String CallFunc24Callback(void* func) {
	return __cross_call_master_CallFunc24Callback(func);
}

extern String (*__cross_call_master_CallFunc25Callback)(void*);

static String CallFunc25Callback(void* func) {
	return __cross_call_master_CallFunc25Callback(func);
}

extern String (*__cross_call_master_CallFunc26Callback)(void*);

static String CallFunc26Callback(void* func) {
	return __cross_call_master_CallFunc26Callback(func);
}

extern String (*__cross_call_master_CallFunc27Callback)(void*);

static String CallFunc27Callback(void* func) {
	return __cross_call_master_CallFunc27Callback(func);
}

extern String (*__cross_call_master_CallFunc28Callback)(void*);

static String CallFunc28Callback(void* func) {
	return __cross_call_master_CallFunc28Callback(func);
}

extern String (*__cross_call_master_CallFunc29Callback)(void*);

static String CallFunc29Callback(void* func) {
	return __cross_call_master_CallFunc29Callback(func);
}

extern String (*__cross_call_master_CallFunc30Callback)(void*);

static String CallFunc30Callback(void* func) {
	return __cross_call_master_CallFunc30Callback(func);
}

extern String (*__cross_call_master_CallFunc31Callback)(void*);

static String CallFunc31Callback(void* func) {
	return __cross_call_master_CallFunc31Callback(func);
}

extern String (*__cross_call_master_CallFunc32Callback)(void*);

static String CallFunc32Callback(void* func) {
	return __cross_call_master_CallFunc32Callback(func);
}

extern String (*__cross_call_master_CallFunc33Callback)(void*);

static String CallFunc33Callback(void* func) {
	return __cross_call_master_CallFunc33Callback(func);
}

extern String (*__cross_call_master_CallFuncEnumCallback)(void*);

static String CallFuncEnumCallback(void* func) {
	return __cross_call_master_CallFuncEnumCallback(func);
}

