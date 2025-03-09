#pragma once
// autoexports.h
#pragma once

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include <uchar.h>

#ifdef __cplusplus
extern "C" {
#endif

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
        char16_t char16;
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
    uint8_t current;
} Variant;

#ifdef __cplusplus
}
#endif
