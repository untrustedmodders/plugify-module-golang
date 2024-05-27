// plugify.h
#pragma once

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

typedef void* (*GetMethodFn)(const char*);

#ifdef __cplusplus
extern "C" {
#endif

void Plugify_SetMethodPtr(void* ptr);
void* Plugify_GetMethodPtr(const char* str);

// std::string
void* Plugify_CreateStringE();
void* Plugify_CreateString(_GoString_ source);
const char* Plugify_GetString(void* ptr);
void Plugify_DeleteString(void* ptr);

// std::vector
void* Plugify_CreateVectorBool(bool* arr, ptrdiff_t len);
void* Plugify_CreateVectorChar8(char* arr, ptrdiff_t len);
void* Plugify_CreateVectorChar16(int16_t* arr, ptrdiff_t len);
void* Plugify_CreateVectorInt8(int8_t* arr, ptrdiff_t len);
void* Plugify_CreateVectorInt16(int16_t* arr, ptrdiff_t len);
void* Plugify_CreateVectorInt32(int32_t* arr, ptrdiff_t len);
void* Plugify_CreateVectorInt64(int64_t* arr, ptrdiff_t len);
void* Plugify_CreateVectorUInt8(uint8_t* arr, ptrdiff_t len);
void* Plugify_CreateVectorUInt16(uint16_t* arr, ptrdiff_t len);
void* Plugify_CreateVectorUInt32(uint32_t* arr, ptrdiff_t len);
void* Plugify_CreateVectorUInt64(uint64_t* arr, ptrdiff_t len);
void* Plugify_CreateVectorUIntPtr(uintptr_t* arr, ptrdiff_t len);
void* Plugify_CreateVectorFloat(float* arr, ptrdiff_t len);
void* Plugify_CreateVectorDouble(double* arr, ptrdiff_t len);
void* Plugify_CreateVectorString(_GoString_** arr, ptrdiff_t len);

void* Plugify_CreateVectorBoolE();
void* Plugify_CreateVectorChar8E();
void* Plugify_CreateVectorChar16E();
void* Plugify_CreateVectorInt8E();
void* Plugify_CreateVectorInt16E();
void* Plugify_CreateVectorInt32E();
void* Plugify_CreateVectorInt64E();
void* Plugify_CreateVectorUInt8E();
void* Plugify_CreateVectorUInt16E();
void* Plugify_CreateVectorUInt32E();
void* Plugify_CreateVectorUInt64E();
void* Plugify_CreateVectorUIntPtrE();
void* Plugify_CreateVectorFloatE();
void* Plugify_CreateVectorDoubleE();
void* Plugify_CreateVectorStringE();

void Plugify_DeleteVectorBool(void* ptr);
void Plugify_DeleteVectorChar8(void* ptr);
void Plugify_DeleteVectorChar16(void* ptr);
void Plugify_DeleteVectorInt8(void* ptr);
void Plugify_DeleteVectorInt16(void* ptr);
void Plugify_DeleteVectorInt32(void* ptr);
void Plugify_DeleteVectorInt64(void* ptr);
void Plugify_DeleteVectorUInt8(void* ptr);
void Plugify_DeleteVectorUInt16(void* ptr);
void Plugify_DeleteVectorUInt32(void* ptr);
void Plugify_DeleteVectorUInt64(void* ptr);
void Plugify_DeleteVectorUIntPtr(void* ptr);
void Plugify_DeleteVectorFloat(void* ptr);
void Plugify_DeleteVectorDouble(void* ptr);
void Plugify_DeleteVectorString(void* ptr);

ptrdiff_t Plugify_GetVectorSize(void* ptr);
void* Plugify_GetVectorData(void* ptr);

// std::vector<bool>
ptrdiff_t Plugify_GetVectorSizeB(void* ptr);
void* Plugify_GetVectorDataB(void* ptr);
void Plugify_DeleteVectorDataB(void* ptr);

// std::vector<std::string>
ptrdiff_t Plugify_GetVectorSizeS(void* ptr);
void* Plugify_GetVectorDataS(void* ptr);
void Plugify_DeleteVectorDataS(void* ptr, ptrdiff_t len);
#ifdef __cplusplus
}
#endif
