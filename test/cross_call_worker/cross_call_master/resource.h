#pragma once

#include "shared.h"

extern uintptr_t (*__cross_call_master_ResourceHandleCreate)(int32_t, String*);

static uintptr_t ResourceHandleCreate(int32_t id, String* name) {
	return __cross_call_master_ResourceHandleCreate(id, name);
}

extern uintptr_t (*__cross_call_master_ResourceHandleCreateDefault)();

static uintptr_t ResourceHandleCreateDefault() {
	return __cross_call_master_ResourceHandleCreateDefault();
}

extern void (*__cross_call_master_ResourceHandleDestroy)(uintptr_t);

static void ResourceHandleDestroy(uintptr_t handle) {
	__cross_call_master_ResourceHandleDestroy(handle);
}

extern int32_t (*__cross_call_master_ResourceHandleGetId)(uintptr_t);

static int32_t ResourceHandleGetId(uintptr_t handle) {
	return __cross_call_master_ResourceHandleGetId(handle);
}

extern String (*__cross_call_master_ResourceHandleGetName)(uintptr_t);

static String ResourceHandleGetName(uintptr_t handle) {
	return __cross_call_master_ResourceHandleGetName(handle);
}

extern void (*__cross_call_master_ResourceHandleSetName)(uintptr_t, String*);

static void ResourceHandleSetName(uintptr_t handle, String* name) {
	__cross_call_master_ResourceHandleSetName(handle, name);
}

extern void (*__cross_call_master_ResourceHandleIncrementCounter)(uintptr_t);

static void ResourceHandleIncrementCounter(uintptr_t handle) {
	__cross_call_master_ResourceHandleIncrementCounter(handle);
}

extern int32_t (*__cross_call_master_ResourceHandleGetCounter)(uintptr_t);

static int32_t ResourceHandleGetCounter(uintptr_t handle) {
	return __cross_call_master_ResourceHandleGetCounter(handle);
}

extern void (*__cross_call_master_ResourceHandleAddData)(uintptr_t, float);

static void ResourceHandleAddData(uintptr_t handle, float value) {
	__cross_call_master_ResourceHandleAddData(handle, value);
}

extern Vector (*__cross_call_master_ResourceHandleGetData)(uintptr_t);

static Vector ResourceHandleGetData(uintptr_t handle) {
	return __cross_call_master_ResourceHandleGetData(handle);
}

extern int32_t (*__cross_call_master_ResourceHandleGetAliveCount)();

static int32_t ResourceHandleGetAliveCount() {
	return __cross_call_master_ResourceHandleGetAliveCount();
}

extern int32_t (*__cross_call_master_ResourceHandleGetTotalCreated)();

static int32_t ResourceHandleGetTotalCreated() {
	return __cross_call_master_ResourceHandleGetTotalCreated();
}

extern int32_t (*__cross_call_master_ResourceHandleGetTotalDestroyed)();

static int32_t ResourceHandleGetTotalDestroyed() {
	return __cross_call_master_ResourceHandleGetTotalDestroyed();
}

