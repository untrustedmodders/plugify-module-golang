#include <stddef.h>
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#include "plugify.h"

GetMethodFn getMethod = NULL;

void Plugify_SetMethodPtr(void* getMethodPtr) {
	getMethod = (GetMethodFn)getMethodPtr;
}

void* Plugify_GetMethod(_GoString_ str) {
	return getMethod(str);
}

void* Plugify_GetMethodCStr(const char* str) {
	_GoString_ gostr;
	gostr.p = str;
	gostr.n = strlen(str);
	return getMethod(gostr);
}