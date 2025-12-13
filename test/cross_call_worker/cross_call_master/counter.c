#include "shared.h"

PLUGIFY_EXPORT uintptr_t (*__cross_call_master_CounterCreate)(int64_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_master_CounterCreateZero)() = NULL;


PLUGIFY_EXPORT int64_t (*__cross_call_master_CounterGetValue)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_CounterSetValue)(uintptr_t, int64_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_CounterIncrement)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_CounterDecrement)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_CounterAdd)(uintptr_t, int64_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_CounterReset)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__cross_call_master_CounterIsPositive)(uintptr_t) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_master_CounterCompare)(int64_t, int64_t) = NULL;


PLUGIFY_EXPORT int64_t (*__cross_call_master_CounterSum)(Vector*) = NULL;


