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

typedef struct { char* data; size_t size; size_t cap; } String;
typedef struct { size_t size; size_t cap; void* data; } Vector;

extern void* Plugify_GetMethodPtr(const char* methodName);
extern void Plugify_GetMethodPtr2(const char* methodName, void** addressDest);

extern String Plugify_ConstructString(_GoString_ source);
extern void Plugify_DestroyString(String* string);
extern const char* Plugify_GetStringData(String* string);
extern ptrdiff_t Plugify_GetStringLength(String* string);
extern void Plugify_AssignString(String* string, _GoString_ source);

extern Vector Plugify_ConstructVector(void* arr, ptrdiff_t len, enum DataType type);
extern void Plugify_DestroyVector(Vector* vector, enum DataType type);
extern void* Plugify_GetVectorData(Vector* vector, enum DataType type);
extern ptrdiff_t Plugify_GetVectorSize(Vector* vector, enum DataType type);
extern void Plugify_AssignVector(Vector* vector, void* arr, ptrdiff_t len, enum DataType type);

extern void Plugify_DeleteVectorDataCStr(void* arr);

extern void Plugify_SetGetMethodPtr(void* func);
extern void Plugify_SetGetMethodPtr2(void* func);

extern void Plugify_SetConstructString(void* func);
extern void Plugify_SetDestroyString(void* func);
extern void Plugify_SetGetStringData(void* func);
extern void Plugify_SetGetStringLength(void* func);
extern void Plugify_SetAssignString(void* func);
extern void Plugify_SetConstructVector(void* func);
extern void Plugify_SetDestroyVector(void* func);
extern void Plugify_SetGetVectorData(void* func);
extern void Plugify_SetGetVectorSize(void* func);
extern void Plugify_SetAssignVector(void* func);

extern void Plugify_SetDeleteVectorDataCStr(void* func);

//typedef struct { const char *p; ptrdiff_t n; } _GoString_;
typedef struct { void *data; ptrdiff_t len; ptrdiff_t cap; } _GoSlice_;

typedef struct { float x, y; } Vector2;
typedef struct { float x, y, z; } Vector3;
typedef struct { float x, y, z, w; } Vector4;
typedef struct { float m[4][4]; } Matrix4x4;

typedef void (*ReverseReturnFn)(String*);
static void ReverseReturn(String* returnString) {
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
typedef String (*NoParamReturnStringCallbackFn)();
static String NoParamReturnStringCallback() {
	static NoParamReturnStringCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnStringCallback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayBoolCallbackFn)();
static Vector NoParamReturnArrayBoolCallback() {
	static NoParamReturnArrayBoolCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayBoolCallback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayChar8CallbackFn)();
static Vector NoParamReturnArrayChar8Callback() {
	static NoParamReturnArrayChar8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayChar8Callback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayChar16CallbackFn)();
static Vector NoParamReturnArrayChar16Callback() {
	static NoParamReturnArrayChar16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayChar16Callback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayInt8CallbackFn)();
static Vector NoParamReturnArrayInt8Callback() {
	static NoParamReturnArrayInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayInt8Callback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayInt16CallbackFn)();
static Vector NoParamReturnArrayInt16Callback() {
	static NoParamReturnArrayInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayInt16Callback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayInt32CallbackFn)();
static Vector NoParamReturnArrayInt32Callback() {
	static NoParamReturnArrayInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayInt32Callback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayInt64CallbackFn)();
static Vector NoParamReturnArrayInt64Callback() {
	static NoParamReturnArrayInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayInt64Callback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayUInt8CallbackFn)();
static Vector NoParamReturnArrayUInt8Callback() {
	static NoParamReturnArrayUInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayUInt8Callback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayUInt16CallbackFn)();
static Vector NoParamReturnArrayUInt16Callback() {
	static NoParamReturnArrayUInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayUInt16Callback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayUInt32CallbackFn)();
static Vector NoParamReturnArrayUInt32Callback() {
	static NoParamReturnArrayUInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayUInt32Callback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayUInt64CallbackFn)();
static Vector NoParamReturnArrayUInt64Callback() {
	static NoParamReturnArrayUInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayUInt64Callback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayPointerCallbackFn)();
static Vector NoParamReturnArrayPointerCallback() {
	static NoParamReturnArrayPointerCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayPointerCallback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayFloatCallbackFn)();
static Vector NoParamReturnArrayFloatCallback() {
	static NoParamReturnArrayFloatCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayFloatCallback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayDoubleCallbackFn)();
static Vector NoParamReturnArrayDoubleCallback() {
	static NoParamReturnArrayDoubleCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayDoubleCallback", (void**)&__func);
	return __func();
}
typedef Vector (*NoParamReturnArrayStringCallbackFn)();
static Vector NoParamReturnArrayStringCallback() {
	static NoParamReturnArrayStringCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayStringCallback", (void**)&__func);
	return __func();
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
typedef void (*Param5CallbackFn)(int32_t, float, double, Vector4*, Vector*);
static void Param5Callback(int32_t a, float b, double c, Vector4* d, Vector* e) {
	static Param5CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param5Callback", (void**)&__func);
	__func(a, b, c, d, e);
}
typedef void (*Param6CallbackFn)(int32_t, float, double, Vector4*, Vector*, char);
static void Param6Callback(int32_t a, float b, double c, Vector4* d, Vector* e, char f) {
	static Param6CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param6Callback", (void**)&__func);
	__func(a, b, c, d, e, f);
}
typedef void (*Param7CallbackFn)(int32_t, float, double, Vector4*, Vector*, char, String*);
static void Param7Callback(int32_t a, float b, double c, Vector4* d, Vector* e, char f, String* g) {
	static Param7CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param7Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g);
}
typedef void (*Param8CallbackFn)(int32_t, float, double, Vector4*, Vector*, char, String*, uint16_t);
static void Param8Callback(int32_t a, float b, double c, Vector4* d, Vector* e, char f, String* g, uint16_t h) {
	static Param8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param8Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h);
}
typedef void (*Param9CallbackFn)(int32_t, float, double, Vector4*, Vector*, char, String*, uint16_t, int16_t);
static void Param9Callback(int32_t a, float b, double c, Vector4* d, Vector* e, char f, String* g, uint16_t h, int16_t k) {
	static Param9CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param9Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k);
}
typedef void (*Param10CallbackFn)(int32_t, float, double, Vector4*, Vector*, char, String*, uint16_t, int16_t, uintptr_t);
static void Param10Callback(int32_t a, float b, double c, Vector4* d, Vector* e, char f, String* g, uint16_t h, int16_t k, uintptr_t l) {
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
typedef void (*ParamRef5CallbackFn)(int32_t*, float*, double*, Vector4*, Vector*);
static void ParamRef5Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e) {
	static ParamRef5CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef5Callback", (void**)&__func);
	__func(a, b, c, d, e);
}
typedef void (*ParamRef6CallbackFn)(int32_t*, float*, double*, Vector4*, Vector*, char*);
static void ParamRef6Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, char* f) {
	static ParamRef6CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef6Callback", (void**)&__func);
	__func(a, b, c, d, e, f);
}
typedef void (*ParamRef7CallbackFn)(int32_t*, float*, double*, Vector4*, Vector*, char*, String*);
static void ParamRef7Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, char* f, String* g) {
	static ParamRef7CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef7Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g);
}
typedef void (*ParamRef8CallbackFn)(int32_t*, float*, double*, Vector4*, Vector*, char*, String*, uint16_t*);
static void ParamRef8Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, char* f, String* g, uint16_t* h) {
	static ParamRef8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef8Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h);
}
typedef void (*ParamRef9CallbackFn)(int32_t*, float*, double*, Vector4*, Vector*, char*, String*, uint16_t*, int16_t*);
static void ParamRef9Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, char* f, String* g, uint16_t* h, int16_t* k) {
	static ParamRef9CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef9Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k);
}
typedef void (*ParamRef10CallbackFn)(int32_t*, float*, double*, Vector4*, Vector*, char*, String*, uint16_t*, int16_t*, uintptr_t*);
static void ParamRef10Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, char* f, String* g, uint16_t* h, int16_t* k, uintptr_t* l) {
	static ParamRef10CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef10Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k, l);
}
typedef void (*ParamRefVectorsCallbackFn)(Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*);
static void ParamRefVectorsCallback(Vector* p1, Vector* p2, Vector* p3, Vector* p4, Vector* p5, Vector* p6, Vector* p7, Vector* p8, Vector* p9, Vector* p10, Vector* p11, Vector* p12, Vector* p13, Vector* p14, Vector* p15) {
	static ParamRefVectorsCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRefVectorsCallback", (void**)&__func);
	__func(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15);
}
typedef int64_t (*ParamAllPrimitivesCallbackFn)(bool, char, uint16_t, int8_t, int16_t, int32_t, int64_t, uint8_t, uint16_t, uint32_t, uint64_t, uintptr_t, float, double);
static int64_t ParamAllPrimitivesCallback(bool p1, char p2, uint16_t p3, int8_t p4, int16_t p5, int32_t p6, int64_t p7, uint8_t p8, uint16_t p9, uint32_t p10, uint64_t p11, uintptr_t p12, float p13, double p14) {
	static ParamAllPrimitivesCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamAllPrimitivesCallback", (void**)&__func);
	return __func(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14);
}/*
typedef void (*CallFuncVoidCallbackFn)(delegate);
static void CallFuncVoidCallback(delegate func) {
	static CallFuncVoidCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncVoidCallback", (void**)&__func);
	__func(func);
}
typedef bool (*CallFuncBoolCallbackFn)(delegate);
static bool CallFuncBoolCallback(delegate func) {
	static CallFuncBoolCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncBoolCallback", (void**)&__func);
	return __func(func);
}
typedef char (*CallFuncChar8CallbackFn)(delegate);
static char CallFuncChar8Callback(delegate func) {
	static CallFuncChar8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncChar8Callback", (void**)&__func);
	return __func(func);
}
typedef uint16_t (*CallFuncChar16CallbackFn)(delegate);
static uint16_t CallFuncChar16Callback(delegate func) {
	static CallFuncChar16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncChar16Callback", (void**)&__func);
	return __func(func);
}
typedef int8_t (*CallFuncInt8CallbackFn)(delegate);
static int8_t CallFuncInt8Callback(delegate func) {
	static CallFuncInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt8Callback", (void**)&__func);
	return __func(func);
}
typedef int16_t (*CallFuncInt16CallbackFn)(delegate);
static int16_t CallFuncInt16Callback(delegate func) {
	static CallFuncInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt16Callback", (void**)&__func);
	return __func(func);
}
typedef int32_t (*CallFuncInt32CallbackFn)(delegate);
static int32_t CallFuncInt32Callback(delegate func) {
	static CallFuncInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt32Callback", (void**)&__func);
	return __func(func);
}
typedef int64_t (*CallFuncInt64CallbackFn)(delegate);
static int64_t CallFuncInt64Callback(delegate func) {
	static CallFuncInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt64Callback", (void**)&__func);
	return __func(func);
}
typedef uint8_t (*CallFuncUInt8CallbackFn)(delegate);
static uint8_t CallFuncUInt8Callback(delegate func) {
	static CallFuncUInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt8Callback", (void**)&__func);
	return __func(func);
}
typedef uint16_t (*CallFuncUInt16CallbackFn)(delegate);
static uint16_t CallFuncUInt16Callback(delegate func) {
	static CallFuncUInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt16Callback", (void**)&__func);
	return __func(func);
}
typedef uint32_t (*CallFuncUInt32CallbackFn)(delegate);
static uint32_t CallFuncUInt32Callback(delegate func) {
	static CallFuncUInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt32Callback", (void**)&__func);
	return __func(func);
}
typedef uint64_t (*CallFuncUInt64CallbackFn)(delegate);
static uint64_t CallFuncUInt64Callback(delegate func) {
	static CallFuncUInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt64Callback", (void**)&__func);
	return __func(func);
}
typedef uintptr_t (*CallFuncPtrCallbackFn)(delegate);
static uintptr_t CallFuncPtrCallback(delegate func) {
	static CallFuncPtrCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncPtrCallback", (void**)&__func);
	return __func(func);
}
typedef float (*CallFuncFloatCallbackFn)(delegate);
static float CallFuncFloatCallback(delegate func) {
	static CallFuncFloatCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncFloatCallback", (void**)&__func);
	return __func(func);
}
typedef double (*CallFuncDoubleCallbackFn)(delegate);
static double CallFuncDoubleCallback(delegate func) {
	static CallFuncDoubleCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncDoubleCallback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFuncStringCallbackFn)(delegate);
static String CallFuncStringCallback(delegate func) {
	static CallFuncStringCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncStringCallback", (void**)&__func);
	return __func(func);
}
typedef uintptr_t (*CallFuncFunctionCallbackFn)(delegate);
static uintptr_t CallFuncFunctionCallback(delegate func) {
	static CallFuncFunctionCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncFunctionCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncBoolVectorCallbackFn)(delegate);
static Vector CallFuncBoolVectorCallback(delegate func) {
	static CallFuncBoolVectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncBoolVectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncChar8VectorCallbackFn)(delegate);
static Vector CallFuncChar8VectorCallback(delegate func) {
	static CallFuncChar8VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncChar8VectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncChar16VectorCallbackFn)(delegate);
static Vector CallFuncChar16VectorCallback(delegate func) {
	static CallFuncChar16VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncChar16VectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncInt8VectorCallbackFn)(delegate);
static Vector CallFuncInt8VectorCallback(delegate func) {
	static CallFuncInt8VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt8VectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncInt16VectorCallbackFn)(delegate);
static Vector CallFuncInt16VectorCallback(delegate func) {
	static CallFuncInt16VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt16VectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncInt32VectorCallbackFn)(delegate);
static Vector CallFuncInt32VectorCallback(delegate func) {
	static CallFuncInt32VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt32VectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncInt64VectorCallbackFn)(delegate);
static Vector CallFuncInt64VectorCallback(delegate func) {
	static CallFuncInt64VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt64VectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncUInt8VectorCallbackFn)(delegate);
static Vector CallFuncUInt8VectorCallback(delegate func) {
	static CallFuncUInt8VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt8VectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncUInt16VectorCallbackFn)(delegate);
static Vector CallFuncUInt16VectorCallback(delegate func) {
	static CallFuncUInt16VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt16VectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncUInt32VectorCallbackFn)(delegate);
static Vector CallFuncUInt32VectorCallback(delegate func) {
	static CallFuncUInt32VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt32VectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncUInt64VectorCallbackFn)(delegate);
static Vector CallFuncUInt64VectorCallback(delegate func) {
	static CallFuncUInt64VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt64VectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncPtrVectorCallbackFn)(delegate);
static Vector CallFuncPtrVectorCallback(delegate func) {
	static CallFuncPtrVectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncPtrVectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncFloatVectorCallbackFn)(delegate);
static Vector CallFuncFloatVectorCallback(delegate func) {
	static CallFuncFloatVectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncFloatVectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncStringVectorCallbackFn)(delegate);
static Vector CallFuncStringVectorCallback(delegate func) {
	static CallFuncStringVectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncStringVectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFuncDoubleVectorCallbackFn)(delegate);
static Vector CallFuncDoubleVectorCallback(delegate func) {
	static CallFuncDoubleVectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncDoubleVectorCallback", (void**)&__func);
	return __func(func);
}
typedef Vector2 (*CallFuncVec2CallbackFn)(delegate);
static Vector2 CallFuncVec2Callback(delegate func) {
	static CallFuncVec2CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncVec2Callback", (void**)&__func);
	return __func(func);
}
typedef Vector3 (*CallFuncVec3CallbackFn)(delegate);
static Vector3 CallFuncVec3Callback(delegate func) {
	static CallFuncVec3CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncVec3Callback", (void**)&__func);
	return __func(func);
}
typedef Vector4 (*CallFuncVec4CallbackFn)(delegate);
static Vector4 CallFuncVec4Callback(delegate func) {
	static CallFuncVec4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncVec4Callback", (void**)&__func);
	return __func(func);
}
typedef Matrix4x4 (*CallFuncMat4x4CallbackFn)(delegate);
static Matrix4x4 CallFuncMat4x4Callback(delegate func) {
	static CallFuncMat4x4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncMat4x4Callback", (void**)&__func);
	return __func(func);
}
typedef int32_t (*CallFunc1CallbackFn)(delegate);
static int32_t CallFunc1Callback(delegate func) {
	static CallFunc1CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc1Callback", (void**)&__func);
	return __func(func);
}
typedef char (*CallFunc2CallbackFn)(delegate);
static char CallFunc2Callback(delegate func) {
	static CallFunc2CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc2Callback", (void**)&__func);
	return __func(func);
}
typedef void (*CallFunc3CallbackFn)(delegate);
static void CallFunc3Callback(delegate func) {
	static CallFunc3CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc3Callback", (void**)&__func);
	__func(func);
}
typedef Vector4 (*CallFunc4CallbackFn)(delegate);
static Vector4 CallFunc4Callback(delegate func) {
	static CallFunc4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc4Callback", (void**)&__func);
	return __func(func);
}
typedef bool (*CallFunc5CallbackFn)(delegate);
static bool CallFunc5Callback(delegate func) {
	static CallFunc5CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc5Callback", (void**)&__func);
	return __func(func);
}
typedef int64_t (*CallFunc6CallbackFn)(delegate);
static int64_t CallFunc6Callback(delegate func) {
	static CallFunc6CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc6Callback", (void**)&__func);
	return __func(func);
}
typedef double (*CallFunc7CallbackFn)(delegate);
static double CallFunc7Callback(delegate func) {
	static CallFunc7CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc7Callback", (void**)&__func);
	return __func(func);
}
typedef Matrix4x4 (*CallFunc8CallbackFn)(delegate);
static Matrix4x4 CallFunc8Callback(delegate func) {
	static CallFunc8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc8Callback", (void**)&__func);
	return __func(func);
}
typedef void (*CallFunc9CallbackFn)(delegate);
static void CallFunc9Callback(delegate func) {
	static CallFunc9CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc9Callback", (void**)&__func);
	__func(func);
}
typedef uint32_t (*CallFunc10CallbackFn)(delegate);
static uint32_t CallFunc10Callback(delegate func) {
	static CallFunc10CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc10Callback", (void**)&__func);
	return __func(func);
}
typedef uintptr_t (*CallFunc11CallbackFn)(delegate);
static uintptr_t CallFunc11Callback(delegate func) {
	static CallFunc11CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc11Callback", (void**)&__func);
	return __func(func);
}
typedef bool (*CallFunc12CallbackFn)(delegate);
static bool CallFunc12Callback(delegate func) {
	static CallFunc12CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc12Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc13CallbackFn)(delegate);
static String CallFunc13Callback(delegate func) {
	static CallFunc13CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc13Callback", (void**)&__func);
	return __func(func);
}
typedef Vector (*CallFunc14CallbackFn)(delegate);
static Vector CallFunc14Callback(delegate func) {
	static CallFunc14CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc14Callback", (void**)&__func);
	return __func(func);
}
typedef int16_t (*CallFunc15CallbackFn)(delegate);
static int16_t CallFunc15Callback(delegate func) {
	static CallFunc15CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc15Callback", (void**)&__func);
	return __func(func);
}
typedef uintptr_t (*CallFunc16CallbackFn)(delegate);
static uintptr_t CallFunc16Callback(delegate func) {
	static CallFunc16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc16Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc17CallbackFn)(delegate);
static String CallFunc17Callback(delegate func) {
	static CallFunc17CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc17Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc18CallbackFn)(delegate);
static String CallFunc18Callback(delegate func) {
	static CallFunc18CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc18Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc19CallbackFn)(delegate);
static String CallFunc19Callback(delegate func) {
	static CallFunc19CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc19Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc20CallbackFn)(delegate);
static String CallFunc20Callback(delegate func) {
	static CallFunc20CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc20Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc21CallbackFn)(delegate);
static String CallFunc21Callback(delegate func) {
	static CallFunc21CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc21Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc22CallbackFn)(delegate);
static String CallFunc22Callback(delegate func) {
	static CallFunc22CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc22Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc23CallbackFn)(delegate);
static String CallFunc23Callback(delegate func) {
	static CallFunc23CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc23Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc24CallbackFn)(delegate);
static String CallFunc24Callback(delegate func) {
	static CallFunc24CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc24Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc25CallbackFn)(delegate);
static String CallFunc25Callback(delegate func) {
	static CallFunc25CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc25Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc26CallbackFn)(delegate);
static String CallFunc26Callback(delegate func) {
	static CallFunc26CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc26Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc27CallbackFn)(delegate);
static String CallFunc27Callback(delegate func) {
	static CallFunc27CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc27Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc28CallbackFn)(delegate);
static String CallFunc28Callback(delegate func) {
	static CallFunc28CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc28Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc29CallbackFn)(delegate);
static String CallFunc29Callback(delegate func) {
	static CallFunc29CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc29Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc30CallbackFn)(delegate);
static String CallFunc30Callback(delegate func) {
	static CallFunc30CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc30Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc31CallbackFn)(delegate);
static String CallFunc31Callback(delegate func) {
	static CallFunc31CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc31Callback", (void**)&__func);
	return __func(func);
}
typedef String (*CallFunc32CallbackFn)(delegate);
static String CallFunc32Callback(delegate func) {
	static CallFunc32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc32Callback", (void**)&__func);
	return __func(func);
}
*/