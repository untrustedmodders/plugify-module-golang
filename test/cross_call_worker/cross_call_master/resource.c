#include "shared.h"

PLUGIFY_EXPORT uintptr_t (*__cross_call_master_ResourceHandleCreate)(int32_t, String*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_master_ResourceHandleCreateDefault)() = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ResourceHandleDestroy)(uintptr_t) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_master_ResourceHandleGetId)(uintptr_t) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_ResourceHandleGetName)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ResourceHandleSetName)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ResourceHandleIncrementCounter)(uintptr_t) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_master_ResourceHandleGetCounter)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ResourceHandleAddData)(uintptr_t, float) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_ResourceHandleGetData)(uintptr_t) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_master_ResourceHandleGetAliveCount)() = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_master_ResourceHandleGetTotalCreated)() = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_master_ResourceHandleGetTotalDestroyed)() = NULL;


