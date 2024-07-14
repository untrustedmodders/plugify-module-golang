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
typedef struct { float m00, m01, m02, m03, m10, m11, m12, m13, m20, m21, m22, m23, m30, m31, m32, m33; } Matrix4x4;

typedef void (*NoParamReturnVoidFn)();
static void NoParamReturnVoid() {
	static NoParamReturnVoidFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnVoid", (void**)&__func);
	__func();
}
typedef bool (*NoParamReturnBoolFn)();
static bool NoParamReturnBool() {
	static NoParamReturnBoolFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnBool", (void**)&__func);
	return __func();
}
typedef char (*NoParamReturnChar8Fn)();
static char NoParamReturnChar8() {
	static NoParamReturnChar8Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnChar8", (void**)&__func);
	return __func();
}
typedef uint16_t (*NoParamReturnChar16Fn)();
static uint16_t NoParamReturnChar16() {
	static NoParamReturnChar16Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnChar16", (void**)&__func);
	return __func();
}
typedef int8_t (*NoParamReturnInt8Fn)();
static int8_t NoParamReturnInt8() {
	static NoParamReturnInt8Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnInt8", (void**)&__func);
	return __func();
}
typedef int16_t (*NoParamReturnInt16Fn)();
static int16_t NoParamReturnInt16() {
	static NoParamReturnInt16Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnInt16", (void**)&__func);
	return __func();
}
typedef int32_t (*NoParamReturnInt32Fn)();
static int32_t NoParamReturnInt32() {
	static NoParamReturnInt32Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnInt32", (void**)&__func);
	return __func();
}
typedef int64_t (*NoParamReturnInt64Fn)();
static int64_t NoParamReturnInt64() {
	static NoParamReturnInt64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnInt64", (void**)&__func);
	return __func();
}
typedef uint8_t (*NoParamReturnUInt8Fn)();
static uint8_t NoParamReturnUInt8() {
	static NoParamReturnUInt8Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnUInt8", (void**)&__func);
	return __func();
}
typedef uint16_t (*NoParamReturnUInt16Fn)();
static uint16_t NoParamReturnUInt16() {
	static NoParamReturnUInt16Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnUInt16", (void**)&__func);
	return __func();
}
typedef uint32_t (*NoParamReturnUInt32Fn)();
static uint32_t NoParamReturnUInt32() {
	static NoParamReturnUInt32Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnUInt32", (void**)&__func);
	return __func();
}
typedef uint64_t (*NoParamReturnUInt64Fn)();
static uint64_t NoParamReturnUInt64() {
	static NoParamReturnUInt64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnUInt64", (void**)&__func);
	return __func();
}
typedef uintptr_t (*NoParamReturnPtr64Fn)();
static uintptr_t NoParamReturnPtr64() {
	static NoParamReturnPtr64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnPtr64", (void**)&__func);
	return __func();
}
typedef float (*NoParamReturnFloatFn)();
static float NoParamReturnFloat() {
	static NoParamReturnFloatFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnFloat", (void**)&__func);
	return __func();
}
typedef double (*NoParamReturnDoubleFn)();
static double NoParamReturnDouble() {
	static NoParamReturnDoubleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnDouble", (void**)&__func);
	return __func();
}
typedef uintptr_t (*NoParamReturnFunctionFn)();
static uintptr_t NoParamReturnFunction() {
	static NoParamReturnFunctionFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnFunction", (void**)&__func);
	return __func();
}
typedef void (*NoParamReturnStringFn)(void*);
static void NoParamReturnString(void* output) {
	static NoParamReturnStringFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnString", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayBoolFn)(void*);
static void NoParamReturnArrayBool(void* output) {
	static NoParamReturnArrayBoolFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayBool", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayChar8Fn)(void*);
static void NoParamReturnArrayChar8(void* output) {
	static NoParamReturnArrayChar8Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayChar8", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayChar16Fn)(void*);
static void NoParamReturnArrayChar16(void* output) {
	static NoParamReturnArrayChar16Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayChar16", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayInt8Fn)(void*);
static void NoParamReturnArrayInt8(void* output) {
	static NoParamReturnArrayInt8Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayInt8", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayInt16Fn)(void*);
static void NoParamReturnArrayInt16(void* output) {
	static NoParamReturnArrayInt16Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayInt16", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayInt32Fn)(void*);
static void NoParamReturnArrayInt32(void* output) {
	static NoParamReturnArrayInt32Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayInt32", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayInt64Fn)(void*);
static void NoParamReturnArrayInt64(void* output) {
	static NoParamReturnArrayInt64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayInt64", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayUInt8Fn)(void*);
static void NoParamReturnArrayUInt8(void* output) {
	static NoParamReturnArrayUInt8Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayUInt8", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayUInt16Fn)(void*);
static void NoParamReturnArrayUInt16(void* output) {
	static NoParamReturnArrayUInt16Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayUInt16", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayUInt32Fn)(void*);
static void NoParamReturnArrayUInt32(void* output) {
	static NoParamReturnArrayUInt32Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayUInt32", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayUInt64Fn)(void*);
static void NoParamReturnArrayUInt64(void* output) {
	static NoParamReturnArrayUInt64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayUInt64", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayPtr64Fn)(void*);
static void NoParamReturnArrayPtr64(void* output) {
	static NoParamReturnArrayPtr64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayPtr64", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayFloatFn)(void*);
static void NoParamReturnArrayFloat(void* output) {
	static NoParamReturnArrayFloatFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayFloat", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayDoubleFn)(void*);
static void NoParamReturnArrayDouble(void* output) {
	static NoParamReturnArrayDoubleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayDouble", (void**)&__func);
	return __func(output);
}
typedef void (*NoParamReturnArrayStringFn)(void*);
static void NoParamReturnArrayString(void* output) {
	static NoParamReturnArrayStringFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnArrayString", (void**)&__func);
	return __func(output);
}
typedef Vector2 (*NoParamReturnVector2Fn)();
static Vector2 NoParamReturnVector2() {
	static NoParamReturnVector2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnVector2", (void**)&__func);
	return __func();
}
typedef Vector3 (*NoParamReturnVector3Fn)();
static Vector3 NoParamReturnVector3() {
	static NoParamReturnVector3Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnVector3", (void**)&__func);
	return __func();
}
typedef Vector4 (*NoParamReturnVector4Fn)();
static Vector4 NoParamReturnVector4() {
	static NoParamReturnVector4Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnVector4", (void**)&__func);
	return __func();
}
typedef Matrix4x4 (*NoParamReturnMatrix4x4Fn)();
static Matrix4x4 NoParamReturnMatrix4x4() {
	static NoParamReturnMatrix4x4Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.NoParamReturnMatrix4x4", (void**)&__func);
	return __func();
}
typedef void (*Param1Fn)(int32_t);
static void Param1(int32_t a) {
	static Param1Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.Param1", (void**)&__func);
	__func(a);
}
typedef void (*Param2Fn)(int32_t, float);
static void Param2(int32_t a, float b) {
	static Param2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.Param2", (void**)&__func);
	__func(a, b);
}
typedef void (*Param3Fn)(int32_t, float, double);
static void Param3(int32_t a, float b, double c) {
	static Param3Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.Param3", (void**)&__func);
	__func(a, b, c);
}
typedef void (*Param4Fn)(int32_t, float, double, Vector4*);
static void Param4(int32_t a, float b, double c, Vector4* d) {
	static Param4Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.Param4", (void**)&__func);
	__func(a, b, c, d);
}
typedef void (*Param5Fn)(int32_t, float, double, Vector4*, void*);
static void Param5(int32_t a, float b, double c, Vector4* d, void* e) {
	static Param5Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.Param5", (void**)&__func);
	__func(a, b, c, d, e);
}
typedef void (*Param6Fn)(int32_t, float, double, Vector4*, void*, char);
static void Param6(int32_t a, float b, double c, Vector4* d, void* e, char f) {
	static Param6Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.Param6", (void**)&__func);
	__func(a, b, c, d, e, f);
}
typedef void (*Param7Fn)(int32_t, float, double, Vector4*, void*, char, void*);
static void Param7(int32_t a, float b, double c, Vector4* d, void* e, char f, void* g) {
	static Param7Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.Param7", (void**)&__func);
	__func(a, b, c, d, e, f, g);
}
typedef void (*Param8Fn)(int32_t, float, double, Vector4*, void*, char, void*, float);
static void Param8(int32_t a, float b, double c, Vector4* d, void* e, char f, void* g, float h) {
	static Param8Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.Param8", (void**)&__func);
	__func(a, b, c, d, e, f, g, h);
}
typedef void (*Param9Fn)(int32_t, float, double, Vector4*, void*, char, void*, float, int16_t);
static void Param9(int32_t a, float b, double c, Vector4* d, void* e, char f, void* g, float h, int16_t k) {
	static Param9Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.Param9", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k);
}
typedef void (*Param10Fn)(int32_t, float, double, Vector4*, void*, char, void*, float, int16_t, uintptr_t);
static void Param10(int32_t a, float b, double c, Vector4* d, void* e, char f, void* g, float h, int16_t k, uintptr_t l) {
	static Param10Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.Param10", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k, l);
}
typedef void (*ParamRef1Fn)(int32_t*);
static void ParamRef1(int32_t* a) {
	static ParamRef1Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.ParamRef1", (void**)&__func);
	__func(a);
}
typedef void (*ParamRef2Fn)(int32_t*, float*);
static void ParamRef2(int32_t* a, float* b) {
	static ParamRef2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.ParamRef2", (void**)&__func);
	__func(a, b);
}
typedef void (*ParamRef3Fn)(int32_t*, float*, double*);
static void ParamRef3(int32_t* a, float* b, double* c) {
	static ParamRef3Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.ParamRef3", (void**)&__func);
	__func(a, b, c);
}
typedef void (*ParamRef4Fn)(int32_t*, float*, double*, Vector4*);
static void ParamRef4(int32_t* a, float* b, double* c, Vector4* d) {
	static ParamRef4Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.ParamRef4", (void**)&__func);
	__func(a, b, c, d);
}
typedef void (*ParamRef5Fn)(int32_t*, float*, double*, Vector4*, void*);
static void ParamRef5(int32_t* a, float* b, double* c, Vector4* d, void* e) {
	static ParamRef5Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.ParamRef5", (void**)&__func);
	__func(a, b, c, d, e);
}
typedef void (*ParamRef6Fn)(int32_t*, float*, double*, Vector4*, void*, char*);
static void ParamRef6(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f) {
	static ParamRef6Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.ParamRef6", (void**)&__func);
	__func(a, b, c, d, e, f);
}
typedef void (*ParamRef7Fn)(int32_t*, float*, double*, Vector4*, void*, char*, void*);
static void ParamRef7(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f, void* g) {
	static ParamRef7Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.ParamRef7", (void**)&__func);
	__func(a, b, c, d, e, f, g);
}
typedef void (*ParamRef8Fn)(int32_t*, float*, double*, Vector4*, void*, char*, void*, float*);
static void ParamRef8(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f, void* g, float* h) {
	static ParamRef8Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.ParamRef8", (void**)&__func);
	__func(a, b, c, d, e, f, g, h);
}
typedef void (*ParamRef9Fn)(int32_t*, float*, double*, Vector4*, void*, char*, void*, float*, int16_t*);
static void ParamRef9(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f, void* g, float* h, int16_t* k) {
	static ParamRef9Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.ParamRef9", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k);
}
typedef void (*ParamRef10Fn)(int32_t*, float*, double*, Vector4*, void*, char*, void*, float*, int16_t*, uintptr_t*);
static void ParamRef10(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f, void* g, float* h, int16_t* k, uintptr_t* l) {
	static ParamRef10Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.ParamRef10", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k, l);
}
typedef void (*ParamRefVectorsFn)(void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*);
static void ParamRefVectors(void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void* p7, void* p8, void* p9, void* p10, void* p11, void* p12, void* p13, void* p14, void* p15) {
	static ParamRefVectorsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.ParamRefVectors", (void**)&__func);
	__func(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15);
}
typedef int64_t (*ParamAllPrimitivesFn)(bool, uint16_t, int8_t, int16_t, int32_t, int64_t, uint8_t, uint16_t, uint32_t, uint64_t, uintptr_t, float, double);
static int64_t ParamAllPrimitives(bool p1, uint16_t p2, int8_t p3, int16_t p4, int32_t p5, int64_t p6, uint8_t p7, uint16_t p8, uint32_t p9, uint64_t p10, uintptr_t p11, float p12, double p13) {
	static ParamAllPrimitivesFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("example_cpp_plugin.ParamAllPrimitives", (void**)&__func);
	return __func(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13);
}
