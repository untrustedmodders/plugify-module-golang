#include <stddef.h>
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#include "plugify.h"
#include <string>
#include <vector>

GetMethodFn getMethod = NULL;

void Plugify_SetMethodPtr(void* ptr) {
	getMethod = (GetMethodFn)ptr;
}
void* Plugify_GetMethodPtr(const char* str) {
	return getMethod(str);
}

void* Plugify_CreateStringE() {
	return new std::string();
}
void* Plugify_CreateString(_GoString_ source) {
	return new std::string(source.p, source.n);
}
const char* Plugify_GetString(void* ptr) {
	return reinterpret_cast<std::string*>(ptr)->c_str();
}
void Plugify_DeleteString(void* ptr) {
	 delete reinterpret_cast<std::string*>(ptr);
}

void* Plugify_CreateVectorBool(bool* arr, ptrdiff_t len) {
	return new std::vector<bool>(arr, arr + len);
}
void* Plugify_CreateVectorChar8(char* arr, ptrdiff_t len) {
	return new std::vector<char>(arr, arr + len);
}
void* Plugify_CreateVectorChar16(int16_t* arr, ptrdiff_t len) {
	return new std::vector<int16_t>(arr, arr + len);
}
void* Plugify_CreateVectorInt8(int8_t* arr, ptrdiff_t len) {
	return new std::vector<int8_t>(arr, arr + len);
}
void* Plugify_CreateVectorInt16(int16_t* arr, ptrdiff_t len) {
	return new std::vector<int16_t>(arr, arr + len);
}
void* Plugify_CreateVectorInt32(int32_t* arr, ptrdiff_t len) {
	return new std::vector<int32_t>(arr, arr + len);
}
void* Plugify_CreateVectorInt64(int64_t* arr, ptrdiff_t len) {
	return new std::vector<int64_t>(arr, arr + len);
}
void* Plugify_CreateVectorUInt8(uint8_t* arr, ptrdiff_t len) {
	return new std::vector<uint8_t>(arr, arr + len);
}
void* Plugify_CreateVectorUInt16(uint16_t* arr, ptrdiff_t len) {
	return new std::vector<uint16_t>(arr, arr + len);
}
void* Plugify_CreateVectorUInt32(uint32_t* arr, ptrdiff_t len) {
	return new std::vector<uint32_t>(arr, arr + len);
}
void* Plugify_CreateVectorUInt64(uint64_t* arr, ptrdiff_t len) {
	return new std::vector<uint64_t>(arr, arr + len);
}
void* Plugify_CreateVectorUIntPtr(uintptr_t* arr, ptrdiff_t len) {
	return new std::vector<uintptr_t>(arr, arr + len);
}
void* Plugify_CreateVectorFloat(float* arr, ptrdiff_t len) {
	return new std::vector<float>(arr, arr + len);
}
void* Plugify_CreateVectorDouble(double* arr, ptrdiff_t len) {
	return new std::vector<double>(arr, arr + len);
}
void* Plugify_CreateVectorString(_GoString_** arr, ptrdiff_t len) {
	auto* vector = new std::vector<std::string>();
	vector->reserve(len);
	for (ptrdiff_t i = 0; i < len; ++i) {
		auto* element = arr[i];
		vector->emplace_back(element->p, element->n);
	}
	return vector;
}

void* Plugify_CreateVectorBoolE() {
	return new std::vector<bool>();
}
void* Plugify_CreateVectorChar8E() {
	return new std::vector<char>();
}
void* Plugify_CreateVectorChar16E() {
	return new std::vector<int16_t>();
}
void* Plugify_CreateVectorInt8E() {
	return new std::vector<int8_t>();
}
void* Plugify_CreateVectorInt16E() {
	return new std::vector<int16_t>();
}
void* Plugify_CreateVectorInt32E() {
	return new std::vector<int32_t>();
}
void* Plugify_CreateVectorInt64E() {
	return new std::vector<int64_t>();
}
void* Plugify_CreateVectorUInt8E() {
	return new std::vector<uint8_t>();
}
void* Plugify_CreateVectorUInt16E() {
	return new std::vector<uint16_t>();
}
void* Plugify_CreateVectorUInt32E() {
	return new std::vector<uint32_t>();
}
void* Plugify_CreateVectorUInt64E() {
	return new std::vector<uint64_t>();
}
void* Plugify_CreateVectorUIntPtrE() {
	return new std::vector<uintptr_t>();
}
void* Plugify_CreateVectorFloatE() {
	return new std::vector<float>();
}
void* Plugify_CreateVectorDoubleE() {
	return new std::vector<double>();
}
void* Plugify_CreateVectorStringE() {
	return new std::vector<std::string>();
}

void Plugify_DeleteVectorBool(void* ptr) {
	delete reinterpret_cast<std::vector<bool>*>(ptr);
}
void Plugify_DeleteVectorChar8(void* ptr) {
	delete reinterpret_cast<std::vector<char>*>(ptr);
}
void Plugify_DeleteVectorChar16(void* ptr) {
	delete reinterpret_cast<std::vector<int16_t>*>(ptr);
}
void Plugify_DeleteVectorInt8(void* ptr) {
	delete reinterpret_cast<std::vector<int8_t>*>(ptr);
}
void Plugify_DeleteVectorInt16(void* ptr) {
	delete reinterpret_cast<std::vector<int16_t>*>(ptr);
}
void Plugify_DeleteVectorInt32(void* ptr) {
	delete reinterpret_cast<std::vector<int32_t>*>(ptr);
}
void Plugify_DeleteVectorInt64(void* ptr) {
	delete reinterpret_cast<std::vector<int64_t>*>(ptr);
}
void Plugify_DeleteVectorUInt8(void* ptr) {
	delete reinterpret_cast<std::vector<uint8_t>*>(ptr);
}
void Plugify_DeleteVectorUInt16(void* ptr) {
	delete reinterpret_cast<std::vector<uint16_t>*>(ptr);
}
void Plugify_DeleteVectorUInt32(void* ptr) {
	delete reinterpret_cast<std::vector<uint32_t>*>(ptr);
}
void Plugify_DeleteVectorUInt64(void* ptr) {
	delete reinterpret_cast<std::vector<uint64_t>*>(ptr);
}
void Plugify_DeleteVectorUIntPtr(void* ptr) {
	delete reinterpret_cast<std::vector<uintptr_t>*>(ptr);
}
void Plugify_DeleteVectorFloat(void* ptr) {
	delete reinterpret_cast<std::vector<float>*>(ptr);
}
void Plugify_DeleteVectorDouble(void* ptr) {
	delete reinterpret_cast<std::vector<double>*>(ptr);
}
void Plugify_DeleteVectorString(void* ptr) {
	delete reinterpret_cast<std::vector<std::string>*>(ptr);
}

ptrdiff_t Plugify_GetVectorSize(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int>*>(ptr)->size());
}
void* Plugify_GetVectorData(void* ptr) {
	return reinterpret_cast<std::vector<int>*>(ptr)->data();
}
ptrdiff_t Plugify_GetVectorSizeB(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<bool>*>(ptr)->size());
}
void* Plugify_GetVectorDataB(void* ptr) {
	auto& vec = *reinterpret_cast<std::vector<bool>*>(ptr);

	bool* boolArray = new bool[vec.size()];

	// Manually copy values from std::vector<bool> to the bool array
	for (size_t i = 0; i < vec.size(); ++i) {
		boolArray[i] = vec[i];
	}

	return boolArray;
}

void Plugify_DeleteVectorDataB(void* ptr) {
	delete[] reinterpret_cast<bool*>(ptr);
}

ptrdiff_t Plugify_GetVectorSizeS(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<std::string>*>(ptr)->size());
}
void* Plugify_GetVectorDataS(void* ptr) {
	auto& vec = *reinterpret_cast<std::vector<std::string>*>(ptr);

	_GoString_** strArray = new _GoString_*[vec.size()];

	// Manually copy values from std::vector<bool> to the bool array
	for (size_t i = 0; i < vec.size(); ++i) {
		auto& element = vec[i];
		strArray[i] = new _GoString_{element.c_str(), static_cast<ptrdiff_t>(element.size())};
	}

	return strArray;
}

void Plugify_DeleteVectorDataS(void* ptr, ptrdiff_t len) {
	auto* arr = reinterpret_cast<_GoString_**>(ptr);
	for (ptrdiff_t i = 0; i < len; ++i) {
		delete arr[i];
	}
	delete[] arr;
}