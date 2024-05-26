// plugify.h
#pragma once

#include <stdlib.h>
#include <stdint.h>
#include <string.h>

typedef struct { const char *p; ptrdiff_t n; } _GoString_;

typedef void* (*GetMethodFn)(_GoString_);

extern void Plugify_SetMethodPtr(void* getMethodPtr);
extern void* Plugify_GetMethod(_GoString_ str);
extern void* Plugify_GetMethodCStr(const char* str);