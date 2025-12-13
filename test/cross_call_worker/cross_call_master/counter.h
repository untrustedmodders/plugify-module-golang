#pragma once

#include "shared.h"

extern uintptr_t (*__cross_call_master_CounterCreate)(int64_t);

static uintptr_t CounterCreate(int64_t initialValue) {
	return __cross_call_master_CounterCreate(initialValue);
}

extern uintptr_t (*__cross_call_master_CounterCreateZero)();

static uintptr_t CounterCreateZero() {
	return __cross_call_master_CounterCreateZero();
}

extern int64_t (*__cross_call_master_CounterGetValue)(uintptr_t);

static int64_t CounterGetValue(uintptr_t counter) {
	return __cross_call_master_CounterGetValue(counter);
}

extern void (*__cross_call_master_CounterSetValue)(uintptr_t, int64_t);

static void CounterSetValue(uintptr_t counter, int64_t value) {
	__cross_call_master_CounterSetValue(counter, value);
}

extern void (*__cross_call_master_CounterIncrement)(uintptr_t);

static void CounterIncrement(uintptr_t counter) {
	__cross_call_master_CounterIncrement(counter);
}

extern void (*__cross_call_master_CounterDecrement)(uintptr_t);

static void CounterDecrement(uintptr_t counter) {
	__cross_call_master_CounterDecrement(counter);
}

extern void (*__cross_call_master_CounterAdd)(uintptr_t, int64_t);

static void CounterAdd(uintptr_t counter, int64_t amount) {
	__cross_call_master_CounterAdd(counter, amount);
}

extern void (*__cross_call_master_CounterReset)(uintptr_t);

static void CounterReset(uintptr_t counter) {
	__cross_call_master_CounterReset(counter);
}

extern bool (*__cross_call_master_CounterIsPositive)(uintptr_t);

static bool CounterIsPositive(uintptr_t counter) {
	return __cross_call_master_CounterIsPositive(counter);
}

extern int32_t (*__cross_call_master_CounterCompare)(int64_t, int64_t);

static int32_t CounterCompare(int64_t value1, int64_t value2) {
	return __cross_call_master_CounterCompare(value1, value2);
}

extern int64_t (*__cross_call_master_CounterSum)(Vector*);

static int64_t CounterSum(Vector* values) {
	return __cross_call_master_CounterSum(values);
}

