#include <stddef.h>
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#include "plugify.h"

typedef void* (*GetMethodFn)(const char*);
typedef void* (*CreateStringEmptyFn)();
typedef void* (*CreateStringFn)(_GoString_);
typedef const char* (*GetStringFn)(void*);
typedef void (*DeleteStringFn)(void*, bool);
typedef void* (*CreateVectorFn)(bool*, ptrdiff_t);
typedef void* (*CreateVectorEmptyFn)();
typedef void* (*DeleteVectorFn)(void*, bool);
typedef ptrdiff_t (*GetVectorSizeFn)(void*);
typedef void* (*GetVectorDataFn)(void*);
typedef void* (*DeleteVectorDataBFn)(void*);
typedef void* (*DeleteVectorDataSFn)(void*);

GetMethodFn getMethod = NULL;
CreateStringEmptyFn createStringE = NULL;
CreateStringFn createString = NULL;
GetStringFn getString = NULL;
DeleteStringFn deleteString = NULL;
CreateVectorFn createVectorBool = NULL;
CreateVectorFn createVectorChar8 = NULL;
CreateVectorFn createVectorChar16 = NULL;
CreateVectorFn createVectorInt8 = NULL;
CreateVectorFn createVectorInt16 = NULL;
CreateVectorFn createVectorInt32 = NULL;
CreateVectorFn createVectorInt64 = NULL;
CreateVectorFn createVectorUInt8 = NULL;
CreateVectorFn createVectorUInt16 = NULL;
CreateVectorFn createVectorUInt32 = NULL;
CreateVectorFn createVectorUInt64 = NULL;
CreateVectorFn createVectorUIntPtr = NULL;
CreateVectorFn createVectorFloat = NULL;
CreateVectorFn createVectorDouble = NULL;
CreateVectorFn createVectorString = NULL;
CreateVectorEmptyFn createVectorBoolE = NULL;
CreateVectorEmptyFn createVectorChar8E = NULL;
CreateVectorEmptyFn createVectorChar16E = NULL;
CreateVectorEmptyFn createVectorInt8E = NULL;
CreateVectorEmptyFn createVectorInt16E = NULL;
CreateVectorEmptyFn createVectorInt32E = NULL;
CreateVectorEmptyFn createVectorInt64E = NULL;
CreateVectorEmptyFn createVectorUInt8E = NULL;
CreateVectorEmptyFn createVectorUInt16E = NULL;
CreateVectorEmptyFn createVectorUInt32E = NULL;
CreateVectorEmptyFn createVectorUInt64E = NULL;
CreateVectorEmptyFn createVectorUIntPtrE = NULL;
CreateVectorEmptyFn createVectorFloatE = NULL;
CreateVectorEmptyFn createVectorDoubleE = NULL;
CreateVectorEmptyFn createVectorStringE = NULL;
DeleteVectorFn deleteVectorBool = NULL;
DeleteVectorFn deleteVectorChar8 = NULL;
DeleteVectorFn deleteVectorChar16 = NULL;
DeleteVectorFn deleteVectorInt8 = NULL;
DeleteVectorFn deleteVectorInt16 = NULL;
DeleteVectorFn deleteVectorInt32 = NULL;
DeleteVectorFn deleteVectorInt64 = NULL;
DeleteVectorFn deleteVectorUInt8 = NULL;
DeleteVectorFn deleteVectorUInt16 = NULL;
DeleteVectorFn deleteVectorUInt32 = NULL;
DeleteVectorFn deleteVectorUInt64 = NULL;
DeleteVectorFn deleteVectorUIntPtr = NULL;
DeleteVectorFn deleteVectorFloat = NULL;
DeleteVectorFn deleteVectorDouble = NULL;
DeleteVectorFn deleteVectorString = NULL;
GetVectorSizeFn getVectorSizeBool = NULL;
GetVectorSizeFn getVectorSizeChar8 = NULL;
GetVectorSizeFn getVectorSizeChar16 = NULL;
GetVectorSizeFn getVectorSizeInt8 = NULL;
GetVectorSizeFn getVectorSizeInt16 = NULL;
GetVectorSizeFn getVectorSizeInt32 = NULL;
GetVectorSizeFn getVectorSizeInt64 = NULL;
GetVectorSizeFn getVectorSizeUInt8 = NULL;
GetVectorSizeFn getVectorSizeUInt16 = NULL;
GetVectorSizeFn getVectorSizeUInt32 = NULL;
GetVectorSizeFn getVectorSizeUInt64 = NULL;
GetVectorSizeFn getVectorSizeUIntPtr = NULL;
GetVectorSizeFn getVectorSizeFloat = NULL;
GetVectorSizeFn getVectorSizeDouble = NULL;
GetVectorSizeFn getVectorSizeString = NULL;
GetVectorDataFn getVectorData = NULL;
GetVectorDataFn getVectorDataB = NULL;
DeleteVectorDataBFn deleteVectorDataB = NULL;
GetVectorDataFn getVectorDataS = NULL;
DeleteVectorDataSFn deleteVectorDataS = NULL;

void* Plugify_GetMethodPtr(const char* str) {
	return getMethod(str);
}

void* Plugify_CreateStringE() {
	return createStringE();
}
void* Plugify_CreateString(_GoString_ source) {
	return createString(source);
}
const char* Plugify_GetString(void* ptr) {
	return getString(ptr);
}
void Plugify_DeleteString(void* ptr, bool output) {
	deleteString(ptr, output);
}

void* Plugify_CreateVectorBool(bool* arr, ptrdiff_t len) {
	return createVectorBool((void*)arr, len);
}
void* Plugify_CreateVectorChar8(char* arr, ptrdiff_t len) {
	return createVectorChar8((void*)arr, len);
}
void* Plugify_CreateVectorChar16(uint16_t* arr, ptrdiff_t len) {
	return createVectorChar16((void*)arr, len);
}
void* Plugify_CreateVectorInt8(int8_t* arr, ptrdiff_t len) {
	return createVectorInt8((void*)arr, len);
}
void* Plugify_CreateVectorInt16(int16_t* arr, ptrdiff_t len) {
	return createVectorInt16((void*)arr, len);
}
void* Plugify_CreateVectorInt32(int32_t* arr, ptrdiff_t len) {
	return createVectorInt32((void*)arr, len);
}
void* Plugify_CreateVectorInt64(int64_t* arr, ptrdiff_t len) {
	return createVectorInt64((void*)arr, len);
}
void* Plugify_CreateVectorUInt8(uint8_t* arr, ptrdiff_t len) {
	return createVectorUInt8((void*)arr, len);
}
void* Plugify_CreateVectorUInt16(uint16_t* arr, ptrdiff_t len) {
	return createVectorUInt16((void*)arr, len);
}
void* Plugify_CreateVectorUInt32(uint32_t* arr, ptrdiff_t len) {
	return createVectorUInt32((void*)arr, len);
}
void* Plugify_CreateVectorUInt64(uint64_t* arr, ptrdiff_t len) {
	return createVectorUInt64((void*)arr, len);
}
void* Plugify_CreateVectorUIntPtr(uintptr_t* arr, ptrdiff_t len) {
	return createVectorUIntPtr((void*)arr, len);
}
void* Plugify_CreateVectorFloat(float* arr, ptrdiff_t len) {
	return createVectorFloat((void*)arr, len);
}
void* Plugify_CreateVectorDouble(double* arr, ptrdiff_t len) {
	return createVectorDouble((void*)arr, len);
}
void* Plugify_CreateVectorString(_GoString_* arr, ptrdiff_t len) {
	return createVectorString((void*)arr, len);
}

void* Plugify_CreateVectorBoolE() {
	return createVectorBoolE();
}
void* Plugify_CreateVectorChar8E() {
	return createVectorChar8E();
}
void* Plugify_CreateVectorChar16E() {
	return createVectorChar16E();
}
void* Plugify_CreateVectorInt8E() {
	return createVectorInt8E();
}
void* Plugify_CreateVectorInt16E() {
	return createVectorInt16E();
}
void* Plugify_CreateVectorInt32E() {
	return createVectorInt32E();
}
void* Plugify_CreateVectorInt64E() {
	return createVectorInt64E();
}
void* Plugify_CreateVectorUInt8E() {
	return createVectorUInt8E();
}
void* Plugify_CreateVectorUInt16E() {
	return createVectorUInt16E();
}
void* Plugify_CreateVectorUInt32E() {
	return createVectorUInt32E();
}
void* Plugify_CreateVectorUInt64E() {
	return createVectorUInt64E();
}
void* Plugify_CreateVectorUIntPtrE() {
	return createVectorUIntPtrE();
}
void* Plugify_CreateVectorFloatE() {
	return createVectorFloatE();
}
void* Plugify_CreateVectorDoubleE() {
	return createVectorDoubleE();
}
void* Plugify_CreateVectorStringE() {
	return createVectorStringE();
}

void Plugify_DeleteVectorBool(void* ptr, bool output) {
	deleteVectorBool(ptr, output);
}
void Plugify_DeleteVectorChar8(void* ptr, bool output) {
	deleteVectorChar8(ptr, output);
}
void Plugify_DeleteVectorChar16(void* ptr, bool output) {
	deleteVectorChar16(ptr, output);
}
void Plugify_DeleteVectorInt8(void* ptr, bool output) {
	deleteVectorInt8(ptr, output);
}
void Plugify_DeleteVectorInt16(void* ptr, bool output) {
	deleteVectorInt16(ptr, output);
}
void Plugify_DeleteVectorInt32(void* ptr, bool output) {
	deleteVectorInt32(ptr, output);
}
void Plugify_DeleteVectorInt64(void* ptr, bool output) {
	deleteVectorInt64(ptr, output);
}
void Plugify_DeleteVectorUInt8(void* ptr, bool output) {
	deleteVectorUInt8(ptr, output);
}
void Plugify_DeleteVectorUInt16(void* ptr, bool output) {
	deleteVectorUInt16(ptr, output);
}
void Plugify_DeleteVectorUInt32(void* ptr, bool output) {
	deleteVectorUInt32(ptr, output);
}
void Plugify_DeleteVectorUInt64(void* ptr, bool output) {
	deleteVectorUInt64(ptr, output);
}
void Plugify_DeleteVectorUIntPtr(void* ptr, bool output) {
	deleteVectorUIntPtr(ptr, output);
}
void Plugify_DeleteVectorFloat(void* ptr, bool output) {
	deleteVectorFloat(ptr, output);
}
void Plugify_DeleteVectorDouble(void* ptr, bool output) {
	deleteVectorDouble(ptr, output);
}
void Plugify_DeleteVectorString(void* ptr, bool output) {
	deleteVectorString(ptr, output);
}

ptrdiff_t Plugify_GetVectorSizeBool(void* ptr) {
	return getVectorSizeBool(ptr);
}
ptrdiff_t Plugify_GetVectorSizeChar8(void* ptr) {
	return getVectorSizeChar8(ptr);
}
ptrdiff_t Plugify_GetVectorSizeChar16(void* ptr) {
	return getVectorSizeChar16(ptr);
}
ptrdiff_t Plugify_GetVectorSizeInt8(void* ptr) {
	return getVectorSizeInt8(ptr);
}
ptrdiff_t Plugify_GetVectorSizeInt16(void* ptr) {
	return getVectorSizeInt16(ptr);
}
ptrdiff_t Plugify_GetVectorSizeInt32(void* ptr) {
	return getVectorSizeInt32(ptr);
}
ptrdiff_t Plugify_GetVectorSizeInt64(void* ptr) {
	return getVectorSizeInt64(ptr);
}
ptrdiff_t Plugify_GetVectorSizeUInt8(void* ptr) {
	return getVectorSizeUInt8(ptr);
}
ptrdiff_t Plugify_GetVectorSizeUInt16(void* ptr) {
	return getVectorSizeUInt16(ptr);
}
ptrdiff_t Plugify_GetVectorSizeUInt32(void* ptr) {
	return getVectorSizeUInt32(ptr);
}
ptrdiff_t Plugify_GetVectorSizeUInt64(void* ptr) {
	return getVectorSizeUInt64(ptr);
}
ptrdiff_t Plugify_GetVectorSizeUIntPtr(void* ptr) {
	return getVectorSizeUIntPtr(ptr);
}
ptrdiff_t Plugify_GetVectorSizeFloat(void* ptr) {
	return getVectorSizeFloat(ptr);
}
ptrdiff_t Plugify_GetVectorSizeDouble(void* ptr) {
	return getVectorSizeDouble(ptr);
}
ptrdiff_t Plugify_GetVectorSizeString(void* ptr) {
	return getVectorSizeString(ptr);
}
void* Plugify_GetVectorData(void* ptr) {
	return getVectorData(ptr);
}

void* Plugify_GetVectorDataB(void* ptr) {
	return getVectorDataB(ptr);
}
void Plugify_DeleteVectorDataB(void* ptr) {
	deleteVectorDataB(ptr);
}

void* Plugify_GetVectorDataS(void* ptr) {
	return getVectorDataS(ptr);
}
void Plugify_DeleteVectorDataS(void* ptr) {
	deleteVectorDataS(ptr);
}

void Plugify_SetGetMethodPtr(void* fn) {
	getMethod = (GetMethodFn)fn;
}

void Plugify_SetCreateStringEmpty(void* fn) {
	createStringE = (CreateStringEmptyFn)fn;
}

void Plugify_SetCreateString(void* fn) {
	createString = (CreateStringFn)fn;
}

void Plugify_SetGetString(void* fn) {
	getString = (GetStringFn)fn;
}

void Plugify_SetDeleteString(void* fn) {
	deleteString = (DeleteStringFn)fn;
}

void Plugify_SetCreateVectorBool(void* fn) {
	createVectorBool = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorChar8(void* fn) {
	createVectorChar8 = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorChar16(void* fn) {
	createVectorChar16 = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorInt8(void* fn) {
	createVectorInt8 = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorInt16(void* fn) {
	createVectorInt16 = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorInt32(void* fn) {
	createVectorInt32 = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorInt64(void* fn) {
	createVectorInt64 = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorUInt8(void* fn) {
	createVectorUInt8 = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorUInt16(void* fn) {
	createVectorUInt16 = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorUInt32(void* fn) {
	createVectorUInt32 = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorUInt64(void* fn) {
	createVectorUInt64 = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorUIntPtr(void* fn) {
	createVectorUIntPtr = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorFloat(void* fn) {
	createVectorFloat = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorDouble(void* fn) {
	createVectorDouble = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorString(void* fn) {
	createVectorString = (CreateVectorFn)fn;
}

void Plugify_SetCreateVectorBoolE(void* fn) {
	createVectorBoolE = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorChar8E(void* fn) {
	createVectorChar8E = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorChar16E(void* fn) {
	createVectorChar16E = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorInt8E(void* fn) {
	createVectorInt8E = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorInt16E(void* fn) {
	createVectorInt16E = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorInt32E(void* fn) {
	createVectorInt32E = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorInt64E(void* fn) {
	createVectorInt64E = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorUInt8E(void* fn) {
	createVectorUInt8E = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorUInt16E(void* fn) {
	createVectorUInt16E = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorUInt32E(void* fn) {
	createVectorUInt32E = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorUInt64E(void* fn) {
	createVectorUInt64E = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorUIntPtrE(void* fn) {
	createVectorUIntPtrE = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorFloatE(void* fn) {
	createVectorFloatE = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorDoubleE(void* fn) {
	createVectorDoubleE = (CreateVectorEmptyFn)fn;
}

void Plugify_SetCreateVectorStringE(void* fn) {
	createVectorStringE = (CreateVectorEmptyFn)fn;
}

void Plugify_SetDeleteVectorBool(void* fn) {
	deleteVectorBool = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorChar8(void* fn) {
	deleteVectorChar8 = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorChar16(void* fn) {
	deleteVectorChar16 = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorInt8(void* fn) {
	deleteVectorInt8 = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorInt16(void* fn) {
	deleteVectorInt16 = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorInt32(void* fn) {
	deleteVectorInt32 = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorInt64(void* fn) {
	deleteVectorInt64 = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorUInt8(void* fn) {
	deleteVectorUInt8 = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorUInt16(void* fn) {
	deleteVectorUInt16 = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorUInt32(void* fn) {
	deleteVectorUInt32 = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorUInt64(void* fn) {
	deleteVectorUInt64 = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorUIntPtr(void* fn) {
	deleteVectorUIntPtr = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorFloat(void* fn) {
	deleteVectorFloat = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorDouble(void* fn) {
	deleteVectorDouble = (DeleteVectorFn)fn;
}

void Plugify_SetDeleteVectorString(void* fn) {
	deleteVectorString = (DeleteVectorFn)fn;
}

void Plugify_SetGetVectorSizeBool(void* fn) {
	getVectorSizeBool = (GetVectorSizeFn) fn;
}

void Plugify_SetGetVectorSizeChar8(void* fn) {
	getVectorSizeChar8 = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeChar16(void* fn) {
	getVectorSizeChar16 = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeInt8(void* fn) {
	getVectorSizeInt8 = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeInt16(void* fn) {
	getVectorSizeInt16 = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeInt32(void* fn) {
	getVectorSizeInt32 = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeInt64(void* fn) {
	getVectorSizeInt64 = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeUInt8(void* fn) {
	getVectorSizeUInt8 = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeUInt16(void* fn) {
	getVectorSizeUInt16 = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeUInt32(void* fn) {
	getVectorSizeUInt32 = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeUInt64(void* fn) {
	getVectorSizeUInt64 = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeUIntPtr(void* fn) {
	getVectorSizeUIntPtr = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeFloat(void* fn) {
	getVectorSizeFloat = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeDouble(void* fn) {
	getVectorSizeDouble = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorSizeString(void* fn) {
	getVectorSizeString = (GetVectorSizeFn)fn;
}

void Plugify_SetGetVectorData(void* fn) {
	getVectorData = (GetVectorDataFn)fn;
}

void Plugify_SetGetVectorDataB(void* fn) {
	getVectorDataB = (GetVectorDataFn)fn;
}

void Plugify_SetDeleteVectorDataB(void* fn) {
	deleteVectorDataB = (DeleteVectorDataBFn)fn;
}

void Plugify_SetGetVectorDataS(void* fn) {
	getVectorDataS = (GetVectorDataFn)fn;
}

void Plugify_SetDeleteVectorDataS(void* fn) {
	deleteVectorDataS = (DeleteVectorDataSFn)fn;
}