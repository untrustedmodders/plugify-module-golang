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
	UINTPTR,
	FLOAT,
	DOUBLE,
	STRING
};

extern void* Plugify_GetMethodPtr(const char* methodName);

extern void* Plugify_AllocateString();
extern void* Plugify_CreateString(_GoString_ source);
extern const char* Plugify_GetStringData(void* ptr);
extern ptrdiff_t Plugify_GetStringSize(void* ptr);
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
extern void Plugify_SetGetStringSize(void* func);
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

typedef void (*NoParamReturnVoidFn)();
static void NoParamReturnVoid() {
	static NoParamReturnVoidFn func = NULL;
	if (func == NULL) func = (NoParamReturnVoidFn)Plugify_GetMethodPtr("cpp_test.NoParamReturnVoid");
	func();
}
typedef bool (*NoParamReturnBoolFn)();
static bool NoParamReturnBool() {
	static NoParamReturnBoolFn func = NULL;
	if (func == NULL) func = (NoParamReturnBoolFn)Plugify_GetMethodPtr("cpp_test.NoParamReturnBool");
	return func();
}
typedef char (*NoParamReturnChar8Fn)();
static char NoParamReturnChar8() {
	static NoParamReturnChar8Fn func = NULL;
	if (func == NULL) func = (NoParamReturnChar8Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnChar8");
	return func();
}
typedef uint16_t (*NoParamReturnChar16Fn)();
static uint16_t NoParamReturnChar16() {
	static NoParamReturnChar16Fn func = NULL;
	if (func == NULL) func = (NoParamReturnChar16Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnChar16");
	return func();
}
typedef int8_t (*NoParamReturnInt8Fn)();
static int8_t NoParamReturnInt8() {
	static NoParamReturnInt8Fn func = NULL;
	if (func == NULL) func = (NoParamReturnInt8Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnInt8");
	return func();
}
typedef int16_t (*NoParamReturnInt16Fn)();
static int16_t NoParamReturnInt16() {
	static NoParamReturnInt16Fn func = NULL;
	if (func == NULL) func = (NoParamReturnInt16Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnInt16");
	return func();
}
typedef int32_t (*NoParamReturnInt32Fn)();
static int32_t NoParamReturnInt32() {
	static NoParamReturnInt32Fn func = NULL;
	if (func == NULL) func = (NoParamReturnInt32Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnInt32");
	return func();
}
typedef int64_t (*NoParamReturnInt64Fn)();
static int64_t NoParamReturnInt64() {
	static NoParamReturnInt64Fn func = NULL;
	if (func == NULL) func = (NoParamReturnInt64Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnInt64");
	return func();
}
typedef uint8_t (*NoParamReturnUInt8Fn)();
static uint8_t NoParamReturnUInt8() {
	static NoParamReturnUInt8Fn func = NULL;
	if (func == NULL) func = (NoParamReturnUInt8Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnUInt8");
	return func();
}
typedef uint16_t (*NoParamReturnUInt16Fn)();
static uint16_t NoParamReturnUInt16() {
	static NoParamReturnUInt16Fn func = NULL;
	if (func == NULL) func = (NoParamReturnUInt16Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnUInt16");
	return func();
}
typedef uint32_t (*NoParamReturnUInt32Fn)();
static uint32_t NoParamReturnUInt32() {
	static NoParamReturnUInt32Fn func = NULL;
	if (func == NULL) func = (NoParamReturnUInt32Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnUInt32");
	return func();
}
typedef uint64_t (*NoParamReturnUInt64Fn)();
static uint64_t NoParamReturnUInt64() {
	static NoParamReturnUInt64Fn func = NULL;
	if (func == NULL) func = (NoParamReturnUInt64Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnUInt64");
	return func();
}
typedef uintptr_t (*NoParamReturnPtr64Fn)();
static uintptr_t NoParamReturnPtr64() {
	static NoParamReturnPtr64Fn func = NULL;
	if (func == NULL) func = (NoParamReturnPtr64Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnPtr64");
	return func();
}
typedef float (*NoParamReturnFloatFn)();
static float NoParamReturnFloat() {
	static NoParamReturnFloatFn func = NULL;
	if (func == NULL) func = (NoParamReturnFloatFn)Plugify_GetMethodPtr("cpp_test.NoParamReturnFloat");
	return func();
}
typedef double (*NoParamReturnDoubleFn)();
static double NoParamReturnDouble() {
	static NoParamReturnDoubleFn func = NULL;
	if (func == NULL) func = (NoParamReturnDoubleFn)Plugify_GetMethodPtr("cpp_test.NoParamReturnDouble");
	return func();
}
typedef uintptr_t (*NoParamReturnFunctionFn)();
static uintptr_t NoParamReturnFunction() {
	static NoParamReturnFunctionFn func = NULL;
	if (func == NULL) func = (NoParamReturnFunctionFn)Plugify_GetMethodPtr("cpp_test.NoParamReturnFunction");
	return func();
}
typedef void (*NoParamReturnStringFn)(void*);
static void NoParamReturnString(void* output) {
	static NoParamReturnStringFn func = NULL;
	if (func == NULL) func = (NoParamReturnStringFn)Plugify_GetMethodPtr("cpp_test.NoParamReturnString");
	return func(output);
}
typedef void (*NoParamReturnArrayBoolFn)(void*);
static void NoParamReturnArrayBool(void* output) {
	static NoParamReturnArrayBoolFn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayBoolFn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayBool");
	return func(output);
}
typedef void (*NoParamReturnArrayChar8Fn)(void*);
static void NoParamReturnArrayChar8(void* output) {
	static NoParamReturnArrayChar8Fn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayChar8Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayChar8");
	return func(output);
}
typedef void (*NoParamReturnArrayChar16Fn)(void*);
static void NoParamReturnArrayChar16(void* output) {
	static NoParamReturnArrayChar16Fn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayChar16Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayChar16");
	return func(output);
}
typedef void (*NoParamReturnArrayInt8Fn)(void*);
static void NoParamReturnArrayInt8(void* output) {
	static NoParamReturnArrayInt8Fn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayInt8Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayInt8");
	return func(output);
}
typedef void (*NoParamReturnArrayInt16Fn)(void*);
static void NoParamReturnArrayInt16(void* output) {
	static NoParamReturnArrayInt16Fn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayInt16Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayInt16");
	return func(output);
}
typedef void (*NoParamReturnArrayInt32Fn)(void*);
static void NoParamReturnArrayInt32(void* output) {
	static NoParamReturnArrayInt32Fn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayInt32Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayInt32");
	return func(output);
}
typedef void (*NoParamReturnArrayInt64Fn)(void*);
static void NoParamReturnArrayInt64(void* output) {
	static NoParamReturnArrayInt64Fn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayInt64Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayInt64");
	return func(output);
}
typedef void (*NoParamReturnArrayUInt8Fn)(void*);
static void NoParamReturnArrayUInt8(void* output) {
	static NoParamReturnArrayUInt8Fn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayUInt8Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayUInt8");
	return func(output);
}
typedef void (*NoParamReturnArrayUInt16Fn)(void*);
static void NoParamReturnArrayUInt16(void* output) {
	static NoParamReturnArrayUInt16Fn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayUInt16Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayUInt16");
	return func(output);
}
typedef void (*NoParamReturnArrayUInt32Fn)(void*);
static void NoParamReturnArrayUInt32(void* output) {
	static NoParamReturnArrayUInt32Fn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayUInt32Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayUInt32");
	return func(output);
}
typedef void (*NoParamReturnArrayUInt64Fn)(void*);
static void NoParamReturnArrayUInt64(void* output) {
	static NoParamReturnArrayUInt64Fn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayUInt64Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayUInt64");
	return func(output);
}
typedef void (*NoParamReturnArrayPtr64Fn)(void*);
static void NoParamReturnArrayPtr64(void* output) {
	static NoParamReturnArrayPtr64Fn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayPtr64Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayPtr64");
	return func(output);
}
typedef void (*NoParamReturnArrayFloatFn)(void*);
static void NoParamReturnArrayFloat(void* output) {
	static NoParamReturnArrayFloatFn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayFloatFn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayFloat");
	return func(output);
}
typedef void (*NoParamReturnArrayDoubleFn)(void*);
static void NoParamReturnArrayDouble(void* output) {
	static NoParamReturnArrayDoubleFn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayDoubleFn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayDouble");
	return func(output);
}
typedef void (*NoParamReturnArrayStringFn)(void*);
static void NoParamReturnArrayString(void* output) {
	static NoParamReturnArrayStringFn func = NULL;
	if (func == NULL) func = (NoParamReturnArrayStringFn)Plugify_GetMethodPtr("cpp_test.NoParamReturnArrayString");
	return func(output);
}
typedef void (*NoParamReturnVector2Fn)(Vector2*);
static void NoParamReturnVector2(Vector2* output) {
	static NoParamReturnVector2Fn func = NULL;
	if (func == NULL) func = (NoParamReturnVector2Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnVector2");
	return func(output);
}
typedef void (*NoParamReturnVector3Fn)(Vector3*);
static void NoParamReturnVector3(Vector3* output) {
	static NoParamReturnVector3Fn func = NULL;
	if (func == NULL) func = (NoParamReturnVector3Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnVector3");
	return func(output);
}
typedef void (*NoParamReturnVector4Fn)(Vector4*);
static void NoParamReturnVector4(Vector4* output) {
	static NoParamReturnVector4Fn func = NULL;
	if (func == NULL) func = (NoParamReturnVector4Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnVector4");
	return func(output);
}
typedef void (*NoParamReturnMatrix4x4Fn)(Matrix4x4*);
static void NoParamReturnMatrix4x4(Matrix4x4* output) {
	static NoParamReturnMatrix4x4Fn func = NULL;
	if (func == NULL) func = (NoParamReturnMatrix4x4Fn)Plugify_GetMethodPtr("cpp_test.NoParamReturnMatrix4x4");
	return func(output);
}
typedef void (*Param1Fn)(int32_t);
static void Param1(int32_t a) {
	static Param1Fn func = NULL;
	if (func == NULL) func = (Param1Fn)Plugify_GetMethodPtr("cpp_test.Param1");
	func(a);
}
typedef void (*Param2Fn)(int32_t, float);
static void Param2(int32_t a, float b) {
	static Param2Fn func = NULL;
	if (func == NULL) func = (Param2Fn)Plugify_GetMethodPtr("cpp_test.Param2");
	func(a, b);
}
typedef void (*Param3Fn)(int32_t, float, double);
static void Param3(int32_t a, float b, double c) {
	static Param3Fn func = NULL;
	if (func == NULL) func = (Param3Fn)Plugify_GetMethodPtr("cpp_test.Param3");
	func(a, b, c);
}
typedef void (*Param4Fn)(int32_t, float, double, Vector4*);
static void Param4(int32_t a, float b, double c, Vector4* d) {
	static Param4Fn func = NULL;
	if (func == NULL) func = (Param4Fn)Plugify_GetMethodPtr("cpp_test.Param4");
	func(a, b, c, d);
}
typedef void (*Param5Fn)(int32_t, float, double, Vector4*, void*);
static void Param5(int32_t a, float b, double c, Vector4* d, void* e) {
	static Param5Fn func = NULL;
	if (func == NULL) func = (Param5Fn)Plugify_GetMethodPtr("cpp_test.Param5");
	func(a, b, c, d, e);
}
typedef void (*Param6Fn)(int32_t, float, double, Vector4*, void*, char);
static void Param6(int32_t a, float b, double c, Vector4* d, void* e, char f) {
	static Param6Fn func = NULL;
	if (func == NULL) func = (Param6Fn)Plugify_GetMethodPtr("cpp_test.Param6");
	func(a, b, c, d, e, f);
}
typedef void (*Param7Fn)(int32_t, float, double, Vector4*, void*, char, void*);
static void Param7(int32_t a, float b, double c, Vector4* d, void* e, char f, void* g) {
	static Param7Fn func = NULL;
	if (func == NULL) func = (Param7Fn)Plugify_GetMethodPtr("cpp_test.Param7");
	func(a, b, c, d, e, f, g);
}
typedef void (*Param8Fn)(int32_t, float, double, Vector4*, void*, char, void*, float);
static void Param8(int32_t a, float b, double c, Vector4* d, void* e, char f, void* g, float h) {
	static Param8Fn func = NULL;
	if (func == NULL) func = (Param8Fn)Plugify_GetMethodPtr("cpp_test.Param8");
	func(a, b, c, d, e, f, g, h);
}
typedef void (*Param9Fn)(int32_t, float, double, Vector4*, void*, char, void*, float, int16_t);
static void Param9(int32_t a, float b, double c, Vector4* d, void* e, char f, void* g, float h, int16_t k) {
	static Param9Fn func = NULL;
	if (func == NULL) func = (Param9Fn)Plugify_GetMethodPtr("cpp_test.Param9");
	func(a, b, c, d, e, f, g, h, k);
}
typedef void (*Param10Fn)(int32_t, float, double, Vector4*, void*, char, void*, float, int16_t, uintptr_t);
static void Param10(int32_t a, float b, double c, Vector4* d, void* e, char f, void* g, float h, int16_t k, uintptr_t l) {
	static Param10Fn func = NULL;
	if (func == NULL) func = (Param10Fn)Plugify_GetMethodPtr("cpp_test.Param10");
	func(a, b, c, d, e, f, g, h, k, l);
}
typedef void (*ParamRef1Fn)(int32_t*);
static void ParamRef1(int32_t* a) {
	static ParamRef1Fn func = NULL;
	if (func == NULL) func = (ParamRef1Fn)Plugify_GetMethodPtr("cpp_test.ParamRef1");
	func(a);
}
typedef void (*ParamRef2Fn)(int32_t*, float*);
static void ParamRef2(int32_t* a, float* b) {
	static ParamRef2Fn func = NULL;
	if (func == NULL) func = (ParamRef2Fn)Plugify_GetMethodPtr("cpp_test.ParamRef2");
	func(a, b);
}
typedef void (*ParamRef3Fn)(int32_t*, float*, double*);
static void ParamRef3(int32_t* a, float* b, double* c) {
	static ParamRef3Fn func = NULL;
	if (func == NULL) func = (ParamRef3Fn)Plugify_GetMethodPtr("cpp_test.ParamRef3");
	func(a, b, c);
}
typedef void (*ParamRef4Fn)(int32_t*, float*, double*, Vector4*);
static void ParamRef4(int32_t* a, float* b, double* c, Vector4* d) {
	static ParamRef4Fn func = NULL;
	if (func == NULL) func = (ParamRef4Fn)Plugify_GetMethodPtr("cpp_test.ParamRef4");
	func(a, b, c, d);
}
typedef void (*ParamRef5Fn)(int32_t*, float*, double*, Vector4*, void*);
static void ParamRef5(int32_t* a, float* b, double* c, Vector4* d, void* e) {
	static ParamRef5Fn func = NULL;
	if (func == NULL) func = (ParamRef5Fn)Plugify_GetMethodPtr("cpp_test.ParamRef5");
	func(a, b, c, d, e);
}
typedef void (*ParamRef6Fn)(int32_t*, float*, double*, Vector4*, void*, char*);
static void ParamRef6(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f) {
	static ParamRef6Fn func = NULL;
	if (func == NULL) func = (ParamRef6Fn)Plugify_GetMethodPtr("cpp_test.ParamRef6");
	func(a, b, c, d, e, f);
}
typedef void (*ParamRef7Fn)(int32_t*, float*, double*, Vector4*, void*, char*, void*);
static void ParamRef7(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f, void* g) {
	static ParamRef7Fn func = NULL;
	if (func == NULL) func = (ParamRef7Fn)Plugify_GetMethodPtr("cpp_test.ParamRef7");
	func(a, b, c, d, e, f, g);
}
typedef void (*ParamRef8Fn)(int32_t*, float*, double*, Vector4*, void*, char*, void*, float*);
static void ParamRef8(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f, void* g, float* h) {
	static ParamRef8Fn func = NULL;
	if (func == NULL) func = (ParamRef8Fn)Plugify_GetMethodPtr("cpp_test.ParamRef8");
	func(a, b, c, d, e, f, g, h);
}
typedef void (*ParamRef9Fn)(int32_t*, float*, double*, Vector4*, void*, char*, void*, float*, int16_t*);
static void ParamRef9(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f, void* g, float* h, int16_t* k) {
	static ParamRef9Fn func = NULL;
	if (func == NULL) func = (ParamRef9Fn)Plugify_GetMethodPtr("cpp_test.ParamRef9");
	func(a, b, c, d, e, f, g, h, k);
}
typedef void (*ParamRef10Fn)(int32_t*, float*, double*, Vector4*, void*, char*, void*, float*, int16_t*, uintptr_t*);
static void ParamRef10(int32_t* a, float* b, double* c, Vector4* d, void* e, char* f, void* g, float* h, int16_t* k, uintptr_t* l) {
	static ParamRef10Fn func = NULL;
	if (func == NULL) func = (ParamRef10Fn)Plugify_GetMethodPtr("cpp_test.ParamRef10");
	func(a, b, c, d, e, f, g, h, k, l);
}
typedef void (*ParamRefVectorsFn)(void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*);
static void ParamRefVectors(void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void* p7, void* p8, void* p9, void* p10, void* p11, void* p12, void* p13, void* p14, void* p15) {
	static ParamRefVectorsFn func = NULL;
	if (func == NULL) func = (ParamRefVectorsFn)Plugify_GetMethodPtr("cpp_test.ParamRefVectors");
	func(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15);
}
typedef int64_t (*ParamAllPrimitivesFn)(bool, uint16_t, int8_t, int16_t, int32_t, int64_t, uint8_t, uint16_t, uint32_t, uint64_t, uintptr_t, float, double);
static int64_t ParamAllPrimitives(bool p1, uint16_t p2, int8_t p3, int16_t p4, int32_t p5, int64_t p6, uint8_t p7, uint16_t p8, uint32_t p9, uint64_t p10, uintptr_t p11, float p12, double p13) {
	static ParamAllPrimitivesFn func = NULL;
	if (func == NULL) func = (ParamAllPrimitivesFn)Plugify_GetMethodPtr("cpp_test.ParamAllPrimitives");
	return func(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13);
}
