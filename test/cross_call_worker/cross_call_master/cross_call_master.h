#pragma once
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

enum DataType {
	BOOL,
	CHAR8,
	CHAR16,
	INT8,
	INT16,
	INT32,
	INT64,
	UINT8,
	UINT16,
	UINT32,
	UINT64,
	POINTER,
	FLOAT,
	DOUBLE,
	STRING
};

extern void* Plugify_GetMethodPtr(const char* methodName);
extern void Plugify_GetMethodPtr2(const char* methodName, void** addressDest);

extern void* Plugify_AllocateString();
extern void* Plugify_CreateString(_GoString_ source);
extern const char* Plugify_GetStringData(void* ptr);
extern ptrdiff_t Plugify_GetStringLength(void* ptr);
extern void Plugify_AssignString(void* ptr, _GoString_ source);
extern void Plugify_FreeString(void* ptr);
extern void Plugify_DeleteString(void* ptr);

extern void* Plugify_CreateVector(void* arr, ptrdiff_t len, enum DataType type);
extern void* Plugify_AllocateVector(enum DataType type);
extern ptrdiff_t Plugify_GetVectorSize(void* ptr, enum DataType type);
extern void* Plugify_GetVectorData(void* ptr, enum DataType type);
extern void Plugify_AssignVector(void* ptr, void* arr, ptrdiff_t len, enum DataType type);
extern void Plugify_DeleteVector(void* ptr, enum DataType type);
extern void Plugify_FreeVector(void* ptr, enum DataType type);

extern void Plugify_DeleteVectorDataBool(void* ptr);
extern void Plugify_DeleteVectorDataCStr(void* ptr);

extern void Plugify_SetGetMethodPtr(void* func);
extern void Plugify_SetAllocateString(void* func);
extern void Plugify_SetCreateString(void* func);
extern void Plugify_SetGetStringData(void* func);
extern void Plugify_SetGetStringLength(void* func);
extern void Plugify_SetAssignString(void* func);
extern void Plugify_SetFreeString(void* func);
extern void Plugify_SetDeleteString(void* func);
extern void Plugify_SetCreateVector(void* func);
extern void Plugify_SetAllocateVector(void* func);
extern void Plugify_SetGetVectorSize(void* func);
extern void Plugify_SetGetVectorData(void* func);
extern void Plugify_SetAssignVector(void* func);
extern void Plugify_SetDeleteVector(void* func);
extern void Plugify_SetFreeVector(void* func);

extern void Plugify_SetDeleteVectorDataBool(void* func);
extern void Plugify_SetDeleteVectorDataCStr(void* func);

//typedef struct { const char *p; ptrdiff_t n; } _GoString_;
typedef struct { void *data; ptrdiff_t len; ptrdiff_t cap; } _GoSlice_;

typedef struct { float x, y; } Vector2;
typedef struct { float x, y, z; } Vector3;
typedef struct { float x, y, z, w; } Vector4;
typedef struct { float m[4][4]; } Matrix4x4;

typedef void (*ReverseReturnFn)(void*);
static void ReverseReturn(void* returnString) {
	static ReverseReturnFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ReverseReturn", (void**)&__func);
	__func(returnString);
}
typedef void (*NoParamReturnVoidCallbackFn)();
static void NoParamReturnVoidCallback() {
	static NoParamReturnVoidCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnVoidCallback", (void**)&__func);
	__func();
}
typedef bool (*NoParamReturnBoolCallbackFn)();
static bool NoParamReturnBoolCallback() {
	static NoParamReturnBoolCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnBoolCallback", (void**)&__func);
	return __func();
}
typedef char (*NoParamReturnChar8CallbackFn)();
static char NoParamReturnChar8Callback() {
	static NoParamReturnChar8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnChar8Callback", (void**)&__func);
	return __func();
}
typedef uint16_t (*NoParamReturnChar16CallbackFn)();
static uint16_t NoParamReturnChar16Callback() {
	static NoParamReturnChar16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnChar16Callback", (void**)&__func);
	return __func();
}
typedef int8_t (*NoParamReturnInt8CallbackFn)();
static int8_t NoParamReturnInt8Callback() {
	static NoParamReturnInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnInt8Callback", (void**)&__func);
	return __func();
}
typedef int16_t (*NoParamReturnInt16CallbackFn)();
static int16_t NoParamReturnInt16Callback() {
	static NoParamReturnInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnInt16Callback", (void**)&__func);
	return __func();
}
typedef int32_t (*NoParamReturnInt32CallbackFn)();
static int32_t NoParamReturnInt32Callback() {
	static NoParamReturnInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnInt32Callback", (void**)&__func);
	return __func();
}
typedef int64_t (*NoParamReturnInt64CallbackFn)();
static int64_t NoParamReturnInt64Callback() {
	static NoParamReturnInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnInt64Callback", (void**)&__func);
	return __func();
}
typedef uint8_t (*NoParamReturnUInt8CallbackFn)();
static uint8_t NoParamReturnUInt8Callback() {
	static NoParamReturnUInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnUInt8Callback", (void**)&__func);
	return __func();
}
typedef uint16_t (*NoParamReturnUInt16CallbackFn)();
static uint16_t NoParamReturnUInt16Callback() {
	static NoParamReturnUInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnUInt16Callback", (void**)&__func);
	return __func();
}
typedef uint32_t (*NoParamReturnUInt32CallbackFn)();
static uint32_t NoParamReturnUInt32Callback() {
	static NoParamReturnUInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnUInt32Callback", (void**)&__func);
	return __func();
}
typedef uint64_t (*NoParamReturnUInt64CallbackFn)();
static uint64_t NoParamReturnUInt64Callback() {
	static NoParamReturnUInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnUInt64Callback", (void**)&__func);
	return __func();
}
typedef uintptr_t (*NoParamReturnPointerCallbackFn)();
static uintptr_t NoParamReturnPointerCallback() {
	static NoParamReturnPointerCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnPointerCallback", (void**)&__func);
	return __func();
}
typedef float (*NoParamReturnFloatCallbackFn)();
static float NoParamReturnFloatCallback() {
	static NoParamReturnFloatCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnFloatCallback", (void**)&__func);
	return __func();
}
typedef double (*NoParamReturnDoubleCallbackFn)();
static double NoParamReturnDoubleCallback() {
	static NoParamReturnDoubleCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnDoubleCallback", (void**)&__func);
	return __func();
}
typedef void* (*NoParamReturnFunctionCallbackFn)();
static void* NoParamReturnFunctionCallback() {
	static NoParamReturnFunctionCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnFunctionCallback", (void**)&__func);
	return __func();
}
typedef void (*NoParamReturnStringCallbackFn)(void*);
static void NoParamReturnStringCallback(void* output) {
	static NoParamReturnStringCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnStringCallback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayBoolCallbackFn)(void*);
static void NoParamReturnArrayBoolCallback(void* output) {
	static NoParamReturnArrayBoolCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayBoolCallback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayChar8CallbackFn)(void*);
static void NoParamReturnArrayChar8Callback(void* output) {
	static NoParamReturnArrayChar8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayChar8Callback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayChar16CallbackFn)(void*);
static void NoParamReturnArrayChar16Callback(void* output) {
	static NoParamReturnArrayChar16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayChar16Callback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayInt8CallbackFn)(void*);
static void NoParamReturnArrayInt8Callback(void* output) {
	static NoParamReturnArrayInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayInt8Callback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayInt16CallbackFn)(void*);
static void NoParamReturnArrayInt16Callback(void* output) {
	static NoParamReturnArrayInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayInt16Callback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayInt32CallbackFn)(void*);
static void NoParamReturnArrayInt32Callback(void* output) {
	static NoParamReturnArrayInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayInt32Callback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayInt64CallbackFn)(void*);
static void NoParamReturnArrayInt64Callback(void* output) {
	static NoParamReturnArrayInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayInt64Callback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayUInt8CallbackFn)(void*);
static void NoParamReturnArrayUInt8Callback(void* output) {
	static NoParamReturnArrayUInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayUInt8Callback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayUInt16CallbackFn)(void*);
static void NoParamReturnArrayUInt16Callback(void* output) {
	static NoParamReturnArrayUInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayUInt16Callback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayUInt32CallbackFn)(void*);
static void NoParamReturnArrayUInt32Callback(void* output) {
	static NoParamReturnArrayUInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayUInt32Callback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayUInt64CallbackFn)(void*);
static void NoParamReturnArrayUInt64Callback(void* output) {
	static NoParamReturnArrayUInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayUInt64Callback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayPointerCallbackFn)(void*);
static void NoParamReturnArrayPointerCallback(void* output) {
	static NoParamReturnArrayPointerCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayPointerCallback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayFloatCallbackFn)(void*);
static void NoParamReturnArrayFloatCallback(void* output) {
	static NoParamReturnArrayFloatCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayFloatCallback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayDoubleCallbackFn)(void*);
static void NoParamReturnArrayDoubleCallback(void* output) {
	static NoParamReturnArrayDoubleCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayDoubleCallback", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayStringCallbackFn)(void*);
static void NoParamReturnArrayStringCallback(void* output) {
	static NoParamReturnArrayStringCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayStringCallback", (void**)&__func);
	return __func(output);
}
typedef Vector2 (*NoParamReturnVector2CallbackFn)();
static Vector2 NoParamReturnVector2Callback() {
	static NoParamReturnVector2CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnVector2Callback", (void**)&__func);
	return __func();
}
typedef Vector3 (*NoParamReturnVector3CallbackFn)();
static Vector3 NoParamReturnVector3Callback() {
	static NoParamReturnVector3CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnVector3Callback", (void**)&__func);
	return __func();
}
typedef Vector4 (*NoParamReturnVector4CallbackFn)();
static Vector4 NoParamReturnVector4Callback() {
	static NoParamReturnVector4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnVector4Callback", (void**)&__func);
	return __func();
}
typedef Matrix4x4 (*NoParamReturnMatrix4x4CallbackFn)();
static Matrix4x4 NoParamReturnMatrix4x4Callback() {
	static NoParamReturnMatrix4x4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnMatrix4x4Callback", (void**)&__func);
	return __func();
}
typedef void (*Param1CallbackFn)(int32_t);
static void Param1Callback(int32_t a) {
	static Param1CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param1Callback", (void**)&__func);
	__func(a);
}
typedef void (*Param2CallbackFn)(int32_t, float);
static void Param2Callback(int32_t a, float b) {
	static Param2CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param2Callback", (void**)&__func);
	__func(a, b);
}
typedef void (*Param3CallbackFn)(int32_t, float, double);
static void Param3Callback(int32_t a, float b, double c) {
	static Param3CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param3Callback", (void**)&__func);
	__func(a, b, c);
}
typedef void (*Param4CallbackFn)(int32_t, float, double, Vector4*);
static void Param4Callback(int32_t a, float b, double c, Vector4* d) {
	static Param4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param4Callback", (void**)&__func);
	__func(a, b, c, d);
}
typedef void (*Param5CallbackFn)(int32_t, float, double, Vector4*, void*);
static void Param5Callback(int32_t a, float b, double c, Vector4* d, void* e) {
	static Param5CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param5Callback", (void**)&__func);
	__func(a, b, c, d, e);
}
typedef void (*Param6CallbackFn)(int32_t, float, double, Vector4*, void*, char);
static void Param6Callback(int32_t a, float b, double c, Vector4* d, void* e, char f) {
	static Param6CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param6Callback", (void**)&__func);
	__func(a, b, c, d, e, f);
}
typedef void (*Param7CallbackFn)(int32_t, float, double, Vector4*, void*, char, void*);
static void Param7Callback(int32_t a, float b, double c, Vector4* d, void* e, char f, void* g) {
	static Param7CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param7Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g);
}
typedef void (*Param8CallbackFn)(int32_t, float, double, Vector4*, void*, char, void*, uint16_t);
static void Param8Callback(int32_t a, float b, double c, Vector4* d, void* e, char f, void* g, uint16_t h) {
	static Param8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param8Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h);
}
typedef void (*Param9CallbackFn)(int32_t, float, double, Vector4*, void*, char, void*, uint16_t, int16_t);
static void Param9Callback(int32_t a, float b, double c, Vector4* d, void* e, char f, void* g, uint16_t h, int16_t k) {
	static Param9CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param9Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k);
}
typedef void (*Param10CallbackFn)(int32_t, float, double, Vector4*, void*, char, void*, uint16_t, int16_t, uintptr_t);
static void Param10Callback(int32_t a, float b, double c, Vector4* d, void* e, char f, void* g, uint16_t h, int16_t k, uintptr_t l) {
	static Param10CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param10Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k, l);
}
typedef void (*ParamRef1CallbackFn)(int32_t*);
static void ParamRef1Callback(int32_t* a) {
	static ParamRef1CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef1Callback", (void**)&__func);
	__func(a);
}
typedef void (*ParamRef2CallbackFn)(int32_t*, float*);
static void ParamRef2Callback(int32_t* a, float* b) {
	static ParamRef2CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef2Callback", (void**)&__func);
	__func(a, b);
}
typedef void (*ParamRef3CallbackFn)(int32_t*, float*, double*);
static void ParamRef3Callback(int32_t* a, float* b, double* c) {
	static ParamRef3CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef3Callback", (void**)&__func);
	__func(a, b, c);
}
typedef void (*ParamRef4CallbackFn)(int32_t*, float*, double*, Vector4*);
static void ParamRef4Callback(int32_t* a, float* b, double* c, Vector4* d) {
	static ParamRef4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef4Callback", (void**)&__func);
	__func(a, b, c, d);
}
typedef void (*ParamRef5CallbackFn)(int32_t*, float*, double*, Vector4*, void*);
static void ParamRef5Callback(int32_t* a, float* b, double* c, Vector4* d, void* e) {
	static ParamRef5CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef5Callback", (void**)&__func);
	__func(a, b, c, d, e);
}
typedef void (*ParamRef6CallbackFn)(int32_t*, float*, double*, Vector4*, void*, char*);
static void ParamRef6Callback(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f) {
	static ParamRef6CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef6Callback", (void**)&__func);
	__func(a, b, c, d, e, f);
}
typedef void (*ParamRef7CallbackFn)(int32_t*, float*, double*, Vector4*, void*, char*, void*);
static void ParamRef7Callback(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f, void* g) {
	static ParamRef7CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef7Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g);
}
typedef void (*ParamRef8CallbackFn)(int32_t*, float*, double*, Vector4*, void*, char*, void*, uint16_t*);
static void ParamRef8Callback(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f, void* g, uint16_t* h) {
	static ParamRef8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef8Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h);
}
typedef void (*ParamRef9CallbackFn)(int32_t*, float*, double*, Vector4*, void*, char*, void*, uint16_t*, int16_t*);
static void ParamRef9Callback(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f, void* g, uint16_t* h, int16_t* k) {
	static ParamRef9CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef9Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k);
}
typedef void (*ParamRef10CallbackFn)(int32_t*, float*, double*, Vector4*, void*, char*, void*, uint16_t*, int16_t*, uintptr_t*);
static void ParamRef10Callback(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f, void* g, uint16_t* h, int16_t* k, uintptr_t* l) {
	static ParamRef10CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef10Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k, l);
}
typedef void (*ParamRefVectorsCallbackFn)(void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*);
static void ParamRefVectorsCallback(void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void* p7, void* p8, void* p9, void* p10, void* p11, void* p12, void* p13, void* p14, void* p15) {
	static ParamRefVectorsCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRefVectorsCallback", (void**)&__func);
	__func(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15);
}
typedef int64_t (*ParamAllPrimitivesCallbackFn)(bool, char, uint16_t, int8_t, int16_t, int32_t, int64_t, uint8_t, uint16_t, uint32_t, uint64_t, uintptr_t, float, double);
static int64_t ParamAllPrimitivesCallback(bool p1, char p2, uint16_t p3, int8_t p4, int16_t p5, int32_t p6, int64_t p7, uint8_t p8, uint16_t p9, uint32_t p10, uint64_t p11, uintptr_t p12, float p13, double p14) {
	static ParamAllPrimitivesCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamAllPrimitivesCallback", (void**)&__func);
	return __func(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14);
}
