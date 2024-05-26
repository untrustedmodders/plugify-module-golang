// plugify.h
#pragma once

typedef void* (*GetMethodFn)(_GoString_);

static GetMethodFn* GetMethodPtr() {
	static GetMethodFn getMethod = NULL;
	return &getMethod;
}

static void SetMethodPtr(void* getMethodPtr) {
	GetMethodFn* getMethod = GetMethodPtr();
	*getMethod = (GetMethodFn)getMethodPtr;
}

static void* GetNativeMethod(_GoString_ str) {
	GetMethodFn& getMethod = *GetMethodPtr();
	return getMethod(str);
}