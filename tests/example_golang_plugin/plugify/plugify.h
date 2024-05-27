// plugify.h
#pragma once

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

void* Plugify_GetMethodPtr(const char* str);

// std::string
void* Plugify_CreateStringE();
void* Plugify_CreateString(_GoString_ source);
const char* Plugify_GetString(void* ptr);
void Plugify_DeleteString(void* ptr, bool output);

// std::vector
void* Plugify_CreateVectorBool(bool* arr, ptrdiff_t len);
void* Plugify_CreateVectorChar8(char* arr, ptrdiff_t len);
void* Plugify_CreateVectorChar16(uint16_t* arr, ptrdiff_t len);
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
void* Plugify_CreateVectorString(_GoString_* arr, ptrdiff_t len);

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

void Plugify_DeleteVectorBool(void* ptr, bool output);
void Plugify_DeleteVectorChar8(void* ptr, bool output);
void Plugify_DeleteVectorChar16(void* ptr, bool output);
void Plugify_DeleteVectorInt8(void* ptr, bool output);
void Plugify_DeleteVectorInt16(void* ptr, bool output);
void Plugify_DeleteVectorInt32(void* ptr, bool output);
void Plugify_DeleteVectorInt64(void* ptr, bool output);
void Plugify_DeleteVectorUInt8(void* ptr, bool output);
void Plugify_DeleteVectorUInt16(void* ptr, bool output);
void Plugify_DeleteVectorUInt32(void* ptr, bool output);
void Plugify_DeleteVectorUInt64(void* ptr, bool output);
void Plugify_DeleteVectorUIntPtr(void* ptr, bool output);
void Plugify_DeleteVectorFloat(void* ptr, bool output);
void Plugify_DeleteVectorDouble(void* ptr, bool output);
void Plugify_DeleteVectorString(void* ptr, bool output);

ptrdiff_t Plugify_GetVectorSizeBool(void* ptr);
ptrdiff_t Plugify_GetVectorSizeChar8(void* ptr);
ptrdiff_t Plugify_GetVectorSizeChar16(void* ptr);
ptrdiff_t Plugify_GetVectorSizeInt8(void* ptr);
ptrdiff_t Plugify_GetVectorSizeInt16(void* ptr);
ptrdiff_t Plugify_GetVectorSizeInt32(void* ptr);
ptrdiff_t Plugify_GetVectorSizeInt64(void* ptr);
ptrdiff_t Plugify_GetVectorSizeUInt8(void* ptr);
ptrdiff_t Plugify_GetVectorSizeUInt16(void* ptr);
ptrdiff_t Plugify_GetVectorSizeUInt32(void* ptr);
ptrdiff_t Plugify_GetVectorSizeUInt64(void* ptr);
ptrdiff_t Plugify_GetVectorSizeUIntPtr(void* ptr);
ptrdiff_t Plugify_GetVectorSizeFloat(void* ptr);
ptrdiff_t Plugify_GetVectorSizeDouble(void* ptr);
ptrdiff_t Plugify_GetVectorSizeString(void* ptr);

void* Plugify_GetVectorData(void* ptr);

// std::vector<bool>
void* Plugify_GetVectorDataB(void* ptr);
void Plugify_DeleteVectorDataB(void* ptr);

// std::vector<std::string>
void* Plugify_GetVectorDataS(void* ptr);
void Plugify_DeleteVectorDataS(void* ptr);

// internal
void Plugify_SetGetMethodPtr(void* fn);
void Plugify_SetCreateStringEmpty(void* fn);
void Plugify_SetCreateString(void* fn);
void Plugify_SetGetString(void* fn);
void Plugify_SetDeleteString(void* fn);
void Plugify_SetCreateVectorBool(void* fn);
void Plugify_SetCreateVectorChar8(void* fn);
void Plugify_SetCreateVectorChar16(void* fn);
void Plugify_SetCreateVectorInt8(void* fn);
void Plugify_SetCreateVectorInt16(void* fn);
void Plugify_SetCreateVectorInt32(void* fn);
void Plugify_SetCreateVectorInt64(void* fn);
void Plugify_SetCreateVectorUInt8(void* fn);
void Plugify_SetCreateVectorUInt16(void* fn);
void Plugify_SetCreateVectorUInt32(void* fn);
void Plugify_SetCreateVectorUInt64(void* fn);
void Plugify_SetCreateVectorUIntPtr(void* fn);
void Plugify_SetCreateVectorFloat(void* fn);
void Plugify_SetCreateVectorDouble(void* fn);
void Plugify_SetCreateVectorString(void* fn);
void Plugify_SetCreateVectorBoolE(void* fn);
void Plugify_SetCreateVectorChar8E(void* fn);
void Plugify_SetCreateVectorChar16E(void* fn);
void Plugify_SetCreateVectorInt8E(void* fn);
void Plugify_SetCreateVectorInt16E(void* fn);
void Plugify_SetCreateVectorInt32E(void* fn);
void Plugify_SetCreateVectorInt64E(void* fn);
void Plugify_SetCreateVectorUInt8E(void* fn);
void Plugify_SetCreateVectorUInt16E(void* fn);
void Plugify_SetCreateVectorUInt32E(void* fn);
void Plugify_SetCreateVectorUInt64E(void* fn);
void Plugify_SetCreateVectorUIntPtrE(void* fn);
void Plugify_SetCreateVectorFloatE(void* fn);
void Plugify_SetCreateVectorDoubleE(void* fn);
void Plugify_SetCreateVectorStringE(void* fn);
void Plugify_SetDeleteVectorBool(void* fn);
void Plugify_SetDeleteVectorChar8(void* fn);
void Plugify_SetDeleteVectorChar16(void* fn);
void Plugify_SetDeleteVectorInt8(void* fn);
void Plugify_SetDeleteVectorInt16(void* fn);
void Plugify_SetDeleteVectorInt32(void* fn);
void Plugify_SetDeleteVectorInt64(void* fn);
void Plugify_SetDeleteVectorUInt8(void* fn);
void Plugify_SetDeleteVectorUInt16(void* fn);
void Plugify_SetDeleteVectorUInt32(void* fn);
void Plugify_SetDeleteVectorUInt64(void* fn);
void Plugify_SetDeleteVectorUIntPtr(void* fn);
void Plugify_SetDeleteVectorFloat(void* fn);
void Plugify_SetDeleteVectorDouble(void* fn);
void Plugify_SetDeleteVectorString(void* fn);
void Plugify_SetGetVectorSizeBool(void* fn);
void Plugify_SetGetVectorSizeChar8(void* fn);
void Plugify_SetGetVectorSizeChar16(void* fn);
void Plugify_SetGetVectorSizeInt8(void* fn);
void Plugify_SetGetVectorSizeInt16(void* fn);
void Plugify_SetGetVectorSizeInt32(void* fn);
void Plugify_SetGetVectorSizeInt64(void* fn);
void Plugify_SetGetVectorSizeUInt8(void* fn);
void Plugify_SetGetVectorSizeUInt16(void* fn);
void Plugify_SetGetVectorSizeUInt32(void* fn);
void Plugify_SetGetVectorSizeUInt64(void* fn);
void Plugify_SetGetVectorSizeUIntPtr(void* fn);
void Plugify_SetGetVectorSizeFloat(void* fn);
void Plugify_SetGetVectorSizeDouble(void* fn);
void Plugify_SetGetVectorSizeString(void* fn);
void Plugify_SetGetVectorData(void* fn);
void Plugify_SetGetVectorDataB(void* fn);
void Plugify_SetDeleteVectorDataB(void* fn);
void Plugify_SetGetVectorDataS(void* fn);
void Plugify_SetDeleteVectorDataS(void* fn);
#ifdef __cplusplus
}
#endif
