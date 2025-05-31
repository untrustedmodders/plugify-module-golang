#pragma once

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

typedef struct String { char* data; size_t size; size_t cap; } String;
typedef struct Vector { void* begin; void* end; void* capacity; } Vector;
typedef struct Vector2 { float x, y; } Vector2;
typedef struct Vector3 { float x, y, z; } Vector3;
typedef struct Vector4 { float x, y, z, w; } Vector4;
typedef struct Matrix4x4 { float m[4][4]; } Matrix4x4;
typedef struct Variant {
    union {
        bool boolean;
        char char8;
        wchar_t char16;
        int8_t int8;
        int16_t int16;
        int32_t int32;
        int64_t int64;
        uint8_t uint8;
        uint16_t uint16;
        uint32_t uint32;
        uint64_t uint64;
        void* ptr;
        float flt;
        double dbl;
        String str;
        Vector vec;
        Vector2 vec2;
        Vector3 vec3;
        Vector4 vec4;
    };
#if INTPTR_MAX == INT32_MAX
	volatile char pad[8];
#endif
    uint8_t current;
} Variant;

extern void* Plugify_GetMethodPtr(const char* methodName);
extern void Plugify_GetMethodPtr2(const char* methodName, void** addressDest);

static void ReverseReturn(String* returnString) {
	typedef void (*ReverseReturnFn)(String*);
	static ReverseReturnFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ReverseReturn", (void**)&__func);
	__func(returnString);
}

static void NoParamReturnVoidCallback() {
	typedef void (*NoParamReturnVoidCallbackFn)();
	static NoParamReturnVoidCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnVoidCallback", (void**)&__func);
	__func();
}

static bool NoParamReturnBoolCallback() {
	typedef bool (*NoParamReturnBoolCallbackFn)();
	static NoParamReturnBoolCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnBoolCallback", (void**)&__func);
	return __func();
}

static int8_t NoParamReturnChar8Callback() {
	typedef int8_t (*NoParamReturnChar8CallbackFn)();
	static NoParamReturnChar8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnChar8Callback", (void**)&__func);
	return __func();
}

static uint16_t NoParamReturnChar16Callback() {
	typedef uint16_t (*NoParamReturnChar16CallbackFn)();
	static NoParamReturnChar16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnChar16Callback", (void**)&__func);
	return __func();
}

static int8_t NoParamReturnInt8Callback() {
	typedef int8_t (*NoParamReturnInt8CallbackFn)();
	static NoParamReturnInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnInt8Callback", (void**)&__func);
	return __func();
}

static int16_t NoParamReturnInt16Callback() {
	typedef int16_t (*NoParamReturnInt16CallbackFn)();
	static NoParamReturnInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnInt16Callback", (void**)&__func);
	return __func();
}

static int32_t NoParamReturnInt32Callback() {
	typedef int32_t (*NoParamReturnInt32CallbackFn)();
	static NoParamReturnInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnInt32Callback", (void**)&__func);
	return __func();
}

static int64_t NoParamReturnInt64Callback() {
	typedef int64_t (*NoParamReturnInt64CallbackFn)();
	static NoParamReturnInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnInt64Callback", (void**)&__func);
	return __func();
}

static uint8_t NoParamReturnUInt8Callback() {
	typedef uint8_t (*NoParamReturnUInt8CallbackFn)();
	static NoParamReturnUInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnUInt8Callback", (void**)&__func);
	return __func();
}

static uint16_t NoParamReturnUInt16Callback() {
	typedef uint16_t (*NoParamReturnUInt16CallbackFn)();
	static NoParamReturnUInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnUInt16Callback", (void**)&__func);
	return __func();
}

static uint32_t NoParamReturnUInt32Callback() {
	typedef uint32_t (*NoParamReturnUInt32CallbackFn)();
	static NoParamReturnUInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnUInt32Callback", (void**)&__func);
	return __func();
}

static uint64_t NoParamReturnUInt64Callback() {
	typedef uint64_t (*NoParamReturnUInt64CallbackFn)();
	static NoParamReturnUInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnUInt64Callback", (void**)&__func);
	return __func();
}

static uintptr_t NoParamReturnPointerCallback() {
	typedef uintptr_t (*NoParamReturnPointerCallbackFn)();
	static NoParamReturnPointerCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnPointerCallback", (void**)&__func);
	return __func();
}

static float NoParamReturnFloatCallback() {
	typedef float (*NoParamReturnFloatCallbackFn)();
	static NoParamReturnFloatCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnFloatCallback", (void**)&__func);
	return __func();
}

static double NoParamReturnDoubleCallback() {
	typedef double (*NoParamReturnDoubleCallbackFn)();
	static NoParamReturnDoubleCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnDoubleCallback", (void**)&__func);
	return __func();
}

static void* NoParamReturnFunctionCallback() {
	typedef void* (*NoParamReturnFunctionCallbackFn)();
	static NoParamReturnFunctionCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnFunctionCallback", (void**)&__func);
	return __func();
}

static String NoParamReturnStringCallback() {
	typedef String (*NoParamReturnStringCallbackFn)();
	static NoParamReturnStringCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnStringCallback", (void**)&__func);
	return __func();
}

static Variant NoParamReturnAnyCallback() {
	typedef Variant (*NoParamReturnAnyCallbackFn)();
	static NoParamReturnAnyCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnAnyCallback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayBoolCallback() {
	typedef Vector (*NoParamReturnArrayBoolCallbackFn)();
	static NoParamReturnArrayBoolCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayBoolCallback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayChar8Callback() {
	typedef Vector (*NoParamReturnArrayChar8CallbackFn)();
	static NoParamReturnArrayChar8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayChar8Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayChar16Callback() {
	typedef Vector (*NoParamReturnArrayChar16CallbackFn)();
	static NoParamReturnArrayChar16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayChar16Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayInt8Callback() {
	typedef Vector (*NoParamReturnArrayInt8CallbackFn)();
	static NoParamReturnArrayInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayInt8Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayInt16Callback() {
	typedef Vector (*NoParamReturnArrayInt16CallbackFn)();
	static NoParamReturnArrayInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayInt16Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayInt32Callback() {
	typedef Vector (*NoParamReturnArrayInt32CallbackFn)();
	static NoParamReturnArrayInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayInt32Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayInt64Callback() {
	typedef Vector (*NoParamReturnArrayInt64CallbackFn)();
	static NoParamReturnArrayInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayInt64Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayUInt8Callback() {
	typedef Vector (*NoParamReturnArrayUInt8CallbackFn)();
	static NoParamReturnArrayUInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayUInt8Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayUInt16Callback() {
	typedef Vector (*NoParamReturnArrayUInt16CallbackFn)();
	static NoParamReturnArrayUInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayUInt16Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayUInt32Callback() {
	typedef Vector (*NoParamReturnArrayUInt32CallbackFn)();
	static NoParamReturnArrayUInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayUInt32Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayUInt64Callback() {
	typedef Vector (*NoParamReturnArrayUInt64CallbackFn)();
	static NoParamReturnArrayUInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayUInt64Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayPointerCallback() {
	typedef Vector (*NoParamReturnArrayPointerCallbackFn)();
	static NoParamReturnArrayPointerCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayPointerCallback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayFloatCallback() {
	typedef Vector (*NoParamReturnArrayFloatCallbackFn)();
	static NoParamReturnArrayFloatCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayFloatCallback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayDoubleCallback() {
	typedef Vector (*NoParamReturnArrayDoubleCallbackFn)();
	static NoParamReturnArrayDoubleCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayDoubleCallback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayStringCallback() {
	typedef Vector (*NoParamReturnArrayStringCallbackFn)();
	static NoParamReturnArrayStringCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayStringCallback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayAnyCallback() {
	typedef Vector (*NoParamReturnArrayAnyCallbackFn)();
	static NoParamReturnArrayAnyCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayAnyCallback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayVector2Callback() {
	typedef Vector (*NoParamReturnArrayVector2CallbackFn)();
	static NoParamReturnArrayVector2CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayVector2Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayVector3Callback() {
	typedef Vector (*NoParamReturnArrayVector3CallbackFn)();
	static NoParamReturnArrayVector3CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayVector3Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayVector4Callback() {
	typedef Vector (*NoParamReturnArrayVector4CallbackFn)();
	static NoParamReturnArrayVector4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayVector4Callback", (void**)&__func);
	return __func();
}

static Vector NoParamReturnArrayMatrix4x4Callback() {
	typedef Vector (*NoParamReturnArrayMatrix4x4CallbackFn)();
	static NoParamReturnArrayMatrix4x4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnArrayMatrix4x4Callback", (void**)&__func);
	return __func();
}

static Vector2 NoParamReturnVector2Callback() {
	typedef Vector2 (*NoParamReturnVector2CallbackFn)();
	static NoParamReturnVector2CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnVector2Callback", (void**)&__func);
	return __func();
}

static Vector3 NoParamReturnVector3Callback() {
	typedef Vector3 (*NoParamReturnVector3CallbackFn)();
	static NoParamReturnVector3CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnVector3Callback", (void**)&__func);
	return __func();
}

static Vector4 NoParamReturnVector4Callback() {
	typedef Vector4 (*NoParamReturnVector4CallbackFn)();
	static NoParamReturnVector4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnVector4Callback", (void**)&__func);
	return __func();
}

static Matrix4x4 NoParamReturnMatrix4x4Callback() {
	typedef Matrix4x4 (*NoParamReturnMatrix4x4CallbackFn)();
	static NoParamReturnMatrix4x4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.NoParamReturnMatrix4x4Callback", (void**)&__func);
	return __func();
}

static void Param1Callback(int32_t a) {
	typedef void (*Param1CallbackFn)(int32_t);
	static Param1CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param1Callback", (void**)&__func);
	__func(a);
}

static void Param2Callback(int32_t a, float b) {
	typedef void (*Param2CallbackFn)(int32_t, float);
	static Param2CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param2Callback", (void**)&__func);
	__func(a, b);
}

static void Param3Callback(int32_t a, float b, double c) {
	typedef void (*Param3CallbackFn)(int32_t, float, double);
	static Param3CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param3Callback", (void**)&__func);
	__func(a, b, c);
}

static void Param4Callback(int32_t a, float b, double c, Vector4* d) {
	typedef void (*Param4CallbackFn)(int32_t, float, double, Vector4*);
	static Param4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param4Callback", (void**)&__func);
	__func(a, b, c, d);
}

static void Param5Callback(int32_t a, float b, double c, Vector4* d, Vector* e) {
	typedef void (*Param5CallbackFn)(int32_t, float, double, Vector4*, Vector*);
	static Param5CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param5Callback", (void**)&__func);
	__func(a, b, c, d, e);
}

static void Param6Callback(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f) {
	typedef void (*Param6CallbackFn)(int32_t, float, double, Vector4*, Vector*, int8_t);
	static Param6CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param6Callback", (void**)&__func);
	__func(a, b, c, d, e, f);
}

static void Param7Callback(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f, String* g) {
	typedef void (*Param7CallbackFn)(int32_t, float, double, Vector4*, Vector*, int8_t, String*);
	static Param7CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param7Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g);
}

static void Param8Callback(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f, String* g, uint16_t h) {
	typedef void (*Param8CallbackFn)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t);
	static Param8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param8Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h);
}

static void Param9Callback(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f, String* g, uint16_t h, int16_t k) {
	typedef void (*Param9CallbackFn)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t, int16_t);
	static Param9CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param9Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k);
}

static void Param10Callback(int32_t a, float b, double c, Vector4* d, Vector* e, int8_t f, String* g, uint16_t h, int16_t k, uintptr_t l) {
	typedef void (*Param10CallbackFn)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t, int16_t, uintptr_t);
	static Param10CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.Param10Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k, l);
}

static void ParamRef1Callback(int32_t* a) {
	typedef void (*ParamRef1CallbackFn)(int32_t*);
	static ParamRef1CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef1Callback", (void**)&__func);
	__func(a);
}

static void ParamRef2Callback(int32_t* a, float* b) {
	typedef void (*ParamRef2CallbackFn)(int32_t*, float*);
	static ParamRef2CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef2Callback", (void**)&__func);
	__func(a, b);
}

static void ParamRef3Callback(int32_t* a, float* b, double* c) {
	typedef void (*ParamRef3CallbackFn)(int32_t*, float*, double*);
	static ParamRef3CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef3Callback", (void**)&__func);
	__func(a, b, c);
}

static void ParamRef4Callback(int32_t* a, float* b, double* c, Vector4* d) {
	typedef void (*ParamRef4CallbackFn)(int32_t*, float*, double*, Vector4*);
	static ParamRef4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef4Callback", (void**)&__func);
	__func(a, b, c, d);
}

static void ParamRef5Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e) {
	typedef void (*ParamRef5CallbackFn)(int32_t*, float*, double*, Vector4*, Vector*);
	static ParamRef5CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef5Callback", (void**)&__func);
	__func(a, b, c, d, e);
}

static void ParamRef6Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f) {
	typedef void (*ParamRef6CallbackFn)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*);
	static ParamRef6CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef6Callback", (void**)&__func);
	__func(a, b, c, d, e, f);
}

static void ParamRef7Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f, String* g) {
	typedef void (*ParamRef7CallbackFn)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*);
	static ParamRef7CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef7Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g);
}

static void ParamRef8Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f, String* g, uint16_t* h) {
	typedef void (*ParamRef8CallbackFn)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*);
	static ParamRef8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef8Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h);
}

static void ParamRef9Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f, String* g, uint16_t* h, int16_t* k) {
	typedef void (*ParamRef9CallbackFn)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*, int16_t*);
	static ParamRef9CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef9Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k);
}

static void ParamRef10Callback(int32_t* a, float* b, double* c, Vector4* d, Vector* e, int8_t* f, String* g, uint16_t* h, int16_t* k, uintptr_t* l) {
	typedef void (*ParamRef10CallbackFn)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*, int16_t*, uintptr_t*);
	static ParamRef10CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRef10Callback", (void**)&__func);
	__func(a, b, c, d, e, f, g, h, k, l);
}

static void ParamRefVectorsCallback(Vector* p1, Vector* p2, Vector* p3, Vector* p4, Vector* p5, Vector* p6, Vector* p7, Vector* p8, Vector* p9, Vector* p10, Vector* p11, Vector* p12, Vector* p13, Vector* p14, Vector* p15) {
	typedef void (*ParamRefVectorsCallbackFn)(Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*);
	static ParamRefVectorsCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamRefVectorsCallback", (void**)&__func);
	__func(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15);
}

static int64_t ParamAllPrimitivesCallback(bool p1, int8_t p2, uint16_t p3, int8_t p4, int16_t p5, int32_t p6, int64_t p7, uint8_t p8, uint16_t p9, uint32_t p10, uint64_t p11, uintptr_t p12, float p13, double p14) {
	typedef int64_t (*ParamAllPrimitivesCallbackFn)(bool, int8_t, uint16_t, int8_t, int16_t, int32_t, int64_t, uint8_t, uint16_t, uint32_t, uint64_t, uintptr_t, float, double);
	static ParamAllPrimitivesCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamAllPrimitivesCallback", (void**)&__func);
	return __func(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14);
}

static int32_t ParamEnumCallback(int32_t p1, Vector* p2) {
	typedef int32_t (*ParamEnumCallbackFn)(int32_t, Vector*);
	static ParamEnumCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamEnumCallback", (void**)&__func);
	return __func(p1, p2);
}

static int32_t ParamEnumRefCallback(int32_t* p1, Vector* p2) {
	typedef int32_t (*ParamEnumRefCallbackFn)(int32_t*, Vector*);
	static ParamEnumRefCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamEnumRefCallback", (void**)&__func);
	return __func(p1, p2);
}

static void ParamVariantCallback(Variant* p1, Vector* p2) {
	typedef void (*ParamVariantCallbackFn)(Variant*, Vector*);
	static ParamVariantCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamVariantCallback", (void**)&__func);
	__func(p1, p2);
}

static void ParamVariantRefCallback(Variant* p1, Vector* p2) {
	typedef void (*ParamVariantRefCallbackFn)(Variant*, Vector*);
	static ParamVariantRefCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.ParamVariantRefCallback", (void**)&__func);
	__func(p1, p2);
}

static void CallFuncVoidCallback(void* func) {
	typedef void (*CallFuncVoidCallbackFn)(void*);
	static CallFuncVoidCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncVoidCallback", (void**)&__func);
	__func(func);
}

static bool CallFuncBoolCallback(void* func) {
	typedef bool (*CallFuncBoolCallbackFn)(void*);
	static CallFuncBoolCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncBoolCallback", (void**)&__func);
	return __func(func);
}

static int8_t CallFuncChar8Callback(void* func) {
	typedef int8_t (*CallFuncChar8CallbackFn)(void*);
	static CallFuncChar8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncChar8Callback", (void**)&__func);
	return __func(func);
}

static uint16_t CallFuncChar16Callback(void* func) {
	typedef uint16_t (*CallFuncChar16CallbackFn)(void*);
	static CallFuncChar16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncChar16Callback", (void**)&__func);
	return __func(func);
}

static int8_t CallFuncInt8Callback(void* func) {
	typedef int8_t (*CallFuncInt8CallbackFn)(void*);
	static CallFuncInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt8Callback", (void**)&__func);
	return __func(func);
}

static int16_t CallFuncInt16Callback(void* func) {
	typedef int16_t (*CallFuncInt16CallbackFn)(void*);
	static CallFuncInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt16Callback", (void**)&__func);
	return __func(func);
}

static int32_t CallFuncInt32Callback(void* func) {
	typedef int32_t (*CallFuncInt32CallbackFn)(void*);
	static CallFuncInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt32Callback", (void**)&__func);
	return __func(func);
}

static int64_t CallFuncInt64Callback(void* func) {
	typedef int64_t (*CallFuncInt64CallbackFn)(void*);
	static CallFuncInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt64Callback", (void**)&__func);
	return __func(func);
}

static uint8_t CallFuncUInt8Callback(void* func) {
	typedef uint8_t (*CallFuncUInt8CallbackFn)(void*);
	static CallFuncUInt8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt8Callback", (void**)&__func);
	return __func(func);
}

static uint16_t CallFuncUInt16Callback(void* func) {
	typedef uint16_t (*CallFuncUInt16CallbackFn)(void*);
	static CallFuncUInt16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt16Callback", (void**)&__func);
	return __func(func);
}

static uint32_t CallFuncUInt32Callback(void* func) {
	typedef uint32_t (*CallFuncUInt32CallbackFn)(void*);
	static CallFuncUInt32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt32Callback", (void**)&__func);
	return __func(func);
}

static uint64_t CallFuncUInt64Callback(void* func) {
	typedef uint64_t (*CallFuncUInt64CallbackFn)(void*);
	static CallFuncUInt64CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt64Callback", (void**)&__func);
	return __func(func);
}

static uintptr_t CallFuncPtrCallback(void* func) {
	typedef uintptr_t (*CallFuncPtrCallbackFn)(void*);
	static CallFuncPtrCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncPtrCallback", (void**)&__func);
	return __func(func);
}

static float CallFuncFloatCallback(void* func) {
	typedef float (*CallFuncFloatCallbackFn)(void*);
	static CallFuncFloatCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncFloatCallback", (void**)&__func);
	return __func(func);
}

static double CallFuncDoubleCallback(void* func) {
	typedef double (*CallFuncDoubleCallbackFn)(void*);
	static CallFuncDoubleCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncDoubleCallback", (void**)&__func);
	return __func(func);
}

static String CallFuncStringCallback(void* func) {
	typedef String (*CallFuncStringCallbackFn)(void*);
	static CallFuncStringCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncStringCallback", (void**)&__func);
	return __func(func);
}

static Variant CallFuncAnyCallback(void* func) {
	typedef Variant (*CallFuncAnyCallbackFn)(void*);
	static CallFuncAnyCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncAnyCallback", (void**)&__func);
	return __func(func);
}

static uintptr_t CallFuncFunctionCallback(void* func) {
	typedef uintptr_t (*CallFuncFunctionCallbackFn)(void*);
	static CallFuncFunctionCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncFunctionCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncBoolVectorCallback(void* func) {
	typedef Vector (*CallFuncBoolVectorCallbackFn)(void*);
	static CallFuncBoolVectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncBoolVectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncChar8VectorCallback(void* func) {
	typedef Vector (*CallFuncChar8VectorCallbackFn)(void*);
	static CallFuncChar8VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncChar8VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncChar16VectorCallback(void* func) {
	typedef Vector (*CallFuncChar16VectorCallbackFn)(void*);
	static CallFuncChar16VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncChar16VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncInt8VectorCallback(void* func) {
	typedef Vector (*CallFuncInt8VectorCallbackFn)(void*);
	static CallFuncInt8VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt8VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncInt16VectorCallback(void* func) {
	typedef Vector (*CallFuncInt16VectorCallbackFn)(void*);
	static CallFuncInt16VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt16VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncInt32VectorCallback(void* func) {
	typedef Vector (*CallFuncInt32VectorCallbackFn)(void*);
	static CallFuncInt32VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt32VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncInt64VectorCallback(void* func) {
	typedef Vector (*CallFuncInt64VectorCallbackFn)(void*);
	static CallFuncInt64VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncInt64VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncUInt8VectorCallback(void* func) {
	typedef Vector (*CallFuncUInt8VectorCallbackFn)(void*);
	static CallFuncUInt8VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt8VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncUInt16VectorCallback(void* func) {
	typedef Vector (*CallFuncUInt16VectorCallbackFn)(void*);
	static CallFuncUInt16VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt16VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncUInt32VectorCallback(void* func) {
	typedef Vector (*CallFuncUInt32VectorCallbackFn)(void*);
	static CallFuncUInt32VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt32VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncUInt64VectorCallback(void* func) {
	typedef Vector (*CallFuncUInt64VectorCallbackFn)(void*);
	static CallFuncUInt64VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncUInt64VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncPtrVectorCallback(void* func) {
	typedef Vector (*CallFuncPtrVectorCallbackFn)(void*);
	static CallFuncPtrVectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncPtrVectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncFloatVectorCallback(void* func) {
	typedef Vector (*CallFuncFloatVectorCallbackFn)(void*);
	static CallFuncFloatVectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncFloatVectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncDoubleVectorCallback(void* func) {
	typedef Vector (*CallFuncDoubleVectorCallbackFn)(void*);
	static CallFuncDoubleVectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncDoubleVectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncStringVectorCallback(void* func) {
	typedef Vector (*CallFuncStringVectorCallbackFn)(void*);
	static CallFuncStringVectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncStringVectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncAnyVectorCallback(void* func) {
	typedef Vector (*CallFuncAnyVectorCallbackFn)(void*);
	static CallFuncAnyVectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncAnyVectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncVec2VectorCallback(void* func) {
	typedef Vector (*CallFuncVec2VectorCallbackFn)(void*);
	static CallFuncVec2VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncVec2VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncVec3VectorCallback(void* func) {
	typedef Vector (*CallFuncVec3VectorCallbackFn)(void*);
	static CallFuncVec3VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncVec3VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncVec4VectorCallback(void* func) {
	typedef Vector (*CallFuncVec4VectorCallbackFn)(void*);
	static CallFuncVec4VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncVec4VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector CallFuncMat4x4VectorCallback(void* func) {
	typedef Vector (*CallFuncMat4x4VectorCallbackFn)(void*);
	static CallFuncMat4x4VectorCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncMat4x4VectorCallback", (void**)&__func);
	return __func(func);
}

static Vector2 CallFuncVec2Callback(void* func) {
	typedef Vector2 (*CallFuncVec2CallbackFn)(void*);
	static CallFuncVec2CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncVec2Callback", (void**)&__func);
	return __func(func);
}

static Vector3 CallFuncVec3Callback(void* func) {
	typedef Vector3 (*CallFuncVec3CallbackFn)(void*);
	static CallFuncVec3CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncVec3Callback", (void**)&__func);
	return __func(func);
}

static Vector4 CallFuncVec4Callback(void* func) {
	typedef Vector4 (*CallFuncVec4CallbackFn)(void*);
	static CallFuncVec4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncVec4Callback", (void**)&__func);
	return __func(func);
}

static Matrix4x4 CallFuncMat4x4Callback(void* func) {
	typedef Matrix4x4 (*CallFuncMat4x4CallbackFn)(void*);
	static CallFuncMat4x4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncMat4x4Callback", (void**)&__func);
	return __func(func);
}

static int32_t CallFunc1Callback(void* func) {
	typedef int32_t (*CallFunc1CallbackFn)(void*);
	static CallFunc1CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc1Callback", (void**)&__func);
	return __func(func);
}

static int8_t CallFunc2Callback(void* func) {
	typedef int8_t (*CallFunc2CallbackFn)(void*);
	static CallFunc2CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc2Callback", (void**)&__func);
	return __func(func);
}

static void CallFunc3Callback(void* func) {
	typedef void (*CallFunc3CallbackFn)(void*);
	static CallFunc3CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc3Callback", (void**)&__func);
	__func(func);
}

static Vector4 CallFunc4Callback(void* func) {
	typedef Vector4 (*CallFunc4CallbackFn)(void*);
	static CallFunc4CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc4Callback", (void**)&__func);
	return __func(func);
}

static bool CallFunc5Callback(void* func) {
	typedef bool (*CallFunc5CallbackFn)(void*);
	static CallFunc5CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc5Callback", (void**)&__func);
	return __func(func);
}

static int64_t CallFunc6Callback(void* func) {
	typedef int64_t (*CallFunc6CallbackFn)(void*);
	static CallFunc6CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc6Callback", (void**)&__func);
	return __func(func);
}

static double CallFunc7Callback(void* func) {
	typedef double (*CallFunc7CallbackFn)(void*);
	static CallFunc7CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc7Callback", (void**)&__func);
	return __func(func);
}

static Matrix4x4 CallFunc8Callback(void* func) {
	typedef Matrix4x4 (*CallFunc8CallbackFn)(void*);
	static CallFunc8CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc8Callback", (void**)&__func);
	return __func(func);
}

static void CallFunc9Callback(void* func) {
	typedef void (*CallFunc9CallbackFn)(void*);
	static CallFunc9CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc9Callback", (void**)&__func);
	__func(func);
}

static uint32_t CallFunc10Callback(void* func) {
	typedef uint32_t (*CallFunc10CallbackFn)(void*);
	static CallFunc10CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc10Callback", (void**)&__func);
	return __func(func);
}

static uintptr_t CallFunc11Callback(void* func) {
	typedef uintptr_t (*CallFunc11CallbackFn)(void*);
	static CallFunc11CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc11Callback", (void**)&__func);
	return __func(func);
}

static bool CallFunc12Callback(void* func) {
	typedef bool (*CallFunc12CallbackFn)(void*);
	static CallFunc12CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc12Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc13Callback(void* func) {
	typedef String (*CallFunc13CallbackFn)(void*);
	static CallFunc13CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc13Callback", (void**)&__func);
	return __func(func);
}

static Vector CallFunc14Callback(void* func) {
	typedef Vector (*CallFunc14CallbackFn)(void*);
	static CallFunc14CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc14Callback", (void**)&__func);
	return __func(func);
}

static int16_t CallFunc15Callback(void* func) {
	typedef int16_t (*CallFunc15CallbackFn)(void*);
	static CallFunc15CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc15Callback", (void**)&__func);
	return __func(func);
}

static uintptr_t CallFunc16Callback(void* func) {
	typedef uintptr_t (*CallFunc16CallbackFn)(void*);
	static CallFunc16CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc16Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc17Callback(void* func) {
	typedef String (*CallFunc17CallbackFn)(void*);
	static CallFunc17CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc17Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc18Callback(void* func) {
	typedef String (*CallFunc18CallbackFn)(void*);
	static CallFunc18CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc18Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc19Callback(void* func) {
	typedef String (*CallFunc19CallbackFn)(void*);
	static CallFunc19CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc19Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc20Callback(void* func) {
	typedef String (*CallFunc20CallbackFn)(void*);
	static CallFunc20CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc20Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc21Callback(void* func) {
	typedef String (*CallFunc21CallbackFn)(void*);
	static CallFunc21CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc21Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc22Callback(void* func) {
	typedef String (*CallFunc22CallbackFn)(void*);
	static CallFunc22CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc22Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc23Callback(void* func) {
	typedef String (*CallFunc23CallbackFn)(void*);
	static CallFunc23CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc23Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc24Callback(void* func) {
	typedef String (*CallFunc24CallbackFn)(void*);
	static CallFunc24CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc24Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc25Callback(void* func) {
	typedef String (*CallFunc25CallbackFn)(void*);
	static CallFunc25CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc25Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc26Callback(void* func) {
	typedef String (*CallFunc26CallbackFn)(void*);
	static CallFunc26CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc26Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc27Callback(void* func) {
	typedef String (*CallFunc27CallbackFn)(void*);
	static CallFunc27CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc27Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc28Callback(void* func) {
	typedef String (*CallFunc28CallbackFn)(void*);
	static CallFunc28CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc28Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc29Callback(void* func) {
	typedef String (*CallFunc29CallbackFn)(void*);
	static CallFunc29CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc29Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc30Callback(void* func) {
	typedef String (*CallFunc30CallbackFn)(void*);
	static CallFunc30CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc30Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc31Callback(void* func) {
	typedef String (*CallFunc31CallbackFn)(void*);
	static CallFunc31CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc31Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc32Callback(void* func) {
	typedef String (*CallFunc32CallbackFn)(void*);
	static CallFunc32CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc32Callback", (void**)&__func);
	return __func(func);
}

static String CallFunc33Callback(void* func) {
	typedef String (*CallFunc33CallbackFn)(void*);
	static CallFunc33CallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFunc33Callback", (void**)&__func);
	return __func(func);
}

static String CallFuncEnumCallback(void* func) {
	typedef String (*CallFuncEnumCallbackFn)(void*);
	static CallFuncEnumCallbackFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("cross_call_master.CallFuncEnumCallback", (void**)&__func);
	return __func(func);
}
