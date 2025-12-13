#include "shared.h"

PLUGIFY_EXPORT void (*__cross_call_master_ReverseReturn)(String*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_NoParamReturnVoidCallback)() = NULL;


PLUGIFY_EXPORT bool (*__cross_call_master_NoParamReturnBoolCallback)() = NULL;


PLUGIFY_EXPORT int8_t (*__cross_call_master_NoParamReturnChar8Callback)() = NULL;


PLUGIFY_EXPORT uint16_t (*__cross_call_master_NoParamReturnChar16Callback)() = NULL;


PLUGIFY_EXPORT int8_t (*__cross_call_master_NoParamReturnInt8Callback)() = NULL;


PLUGIFY_EXPORT int16_t (*__cross_call_master_NoParamReturnInt16Callback)() = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_master_NoParamReturnInt32Callback)() = NULL;


PLUGIFY_EXPORT int64_t (*__cross_call_master_NoParamReturnInt64Callback)() = NULL;


PLUGIFY_EXPORT uint8_t (*__cross_call_master_NoParamReturnUInt8Callback)() = NULL;


PLUGIFY_EXPORT uint16_t (*__cross_call_master_NoParamReturnUInt16Callback)() = NULL;


PLUGIFY_EXPORT uint32_t (*__cross_call_master_NoParamReturnUInt32Callback)() = NULL;


PLUGIFY_EXPORT uint64_t (*__cross_call_master_NoParamReturnUInt64Callback)() = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_master_NoParamReturnPointerCallback)() = NULL;


PLUGIFY_EXPORT float (*__cross_call_master_NoParamReturnFloatCallback)() = NULL;


PLUGIFY_EXPORT double (*__cross_call_master_NoParamReturnDoubleCallback)() = NULL;


PLUGIFY_EXPORT void* (*__cross_call_master_NoParamReturnFunctionCallback)() = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_NoParamReturnStringCallback)() = NULL;


PLUGIFY_EXPORT Variant (*__cross_call_master_NoParamReturnAnyCallback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayBoolCallback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayChar8Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayChar16Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayInt8Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayInt16Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayInt32Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayInt64Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayUInt8Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayUInt16Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayUInt32Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayUInt64Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayPointerCallback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayFloatCallback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayDoubleCallback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayStringCallback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayAnyCallback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayVector2Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayVector3Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayVector4Callback)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_NoParamReturnArrayMatrix4x4Callback)() = NULL;


PLUGIFY_EXPORT Vector2 (*__cross_call_master_NoParamReturnVector2Callback)() = NULL;


PLUGIFY_EXPORT Vector3 (*__cross_call_master_NoParamReturnVector3Callback)() = NULL;


PLUGIFY_EXPORT Vector4 (*__cross_call_master_NoParamReturnVector4Callback)() = NULL;


PLUGIFY_EXPORT Matrix4x4 (*__cross_call_master_NoParamReturnMatrix4x4Callback)() = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_Param1Callback)(int32_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_Param2Callback)(int32_t, float) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_Param3Callback)(int32_t, float, double) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_Param4Callback)(int32_t, float, double, Vector4*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_Param5Callback)(int32_t, float, double, Vector4*, Vector*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_Param6Callback)(int32_t, float, double, Vector4*, Vector*, int8_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_Param7Callback)(int32_t, float, double, Vector4*, Vector*, int8_t, String*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_Param8Callback)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_Param9Callback)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t, int16_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_Param10Callback)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t, int16_t, uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamRef1Callback)(int32_t*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamRef2Callback)(int32_t*, float*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamRef3Callback)(int32_t*, float*, double*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamRef4Callback)(int32_t*, float*, double*, Vector4*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamRef5Callback)(int32_t*, float*, double*, Vector4*, Vector*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamRef6Callback)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamRef7Callback)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamRef8Callback)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamRef9Callback)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*, int16_t*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamRef10Callback)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*, int16_t*, uintptr_t*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamRefVectorsCallback)(Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*) = NULL;


PLUGIFY_EXPORT int64_t (*__cross_call_master_ParamAllPrimitivesCallback)(bool, int8_t, uint16_t, int8_t, int16_t, int32_t, int64_t, uint8_t, uint16_t, uint32_t, uint64_t, uintptr_t, float, double) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_master_ParamEnumCallback)(int32_t, Vector*) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_master_ParamEnumRefCallback)(int32_t*, Vector*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamVariantCallback)(Variant*, Vector*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_ParamVariantRefCallback)(Variant*, Vector*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_CallFuncVoidCallback)(void*) = NULL;


PLUGIFY_EXPORT bool (*__cross_call_master_CallFuncBoolCallback)(void*) = NULL;


PLUGIFY_EXPORT int8_t (*__cross_call_master_CallFuncChar8Callback)(void*) = NULL;


PLUGIFY_EXPORT uint16_t (*__cross_call_master_CallFuncChar16Callback)(void*) = NULL;


PLUGIFY_EXPORT int8_t (*__cross_call_master_CallFuncInt8Callback)(void*) = NULL;


PLUGIFY_EXPORT int16_t (*__cross_call_master_CallFuncInt16Callback)(void*) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_master_CallFuncInt32Callback)(void*) = NULL;


PLUGIFY_EXPORT int64_t (*__cross_call_master_CallFuncInt64Callback)(void*) = NULL;


PLUGIFY_EXPORT uint8_t (*__cross_call_master_CallFuncUInt8Callback)(void*) = NULL;


PLUGIFY_EXPORT uint16_t (*__cross_call_master_CallFuncUInt16Callback)(void*) = NULL;


PLUGIFY_EXPORT uint32_t (*__cross_call_master_CallFuncUInt32Callback)(void*) = NULL;


PLUGIFY_EXPORT uint64_t (*__cross_call_master_CallFuncUInt64Callback)(void*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_master_CallFuncPtrCallback)(void*) = NULL;


PLUGIFY_EXPORT float (*__cross_call_master_CallFuncFloatCallback)(void*) = NULL;


PLUGIFY_EXPORT double (*__cross_call_master_CallFuncDoubleCallback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFuncStringCallback)(void*) = NULL;


PLUGIFY_EXPORT Variant (*__cross_call_master_CallFuncAnyCallback)(void*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_master_CallFuncFunctionCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncBoolVectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncChar8VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncChar16VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncInt8VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncInt16VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncInt32VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncInt64VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncUInt8VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncUInt16VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncUInt32VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncUInt64VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncPtrVectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncFloatVectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncDoubleVectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncStringVectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncAnyVectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncVec2VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncVec3VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncVec4VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFuncMat4x4VectorCallback)(void*) = NULL;


PLUGIFY_EXPORT Vector2 (*__cross_call_master_CallFuncVec2Callback)(void*) = NULL;


PLUGIFY_EXPORT Vector3 (*__cross_call_master_CallFuncVec3Callback)(void*) = NULL;


PLUGIFY_EXPORT Vector4 (*__cross_call_master_CallFuncVec4Callback)(void*) = NULL;


PLUGIFY_EXPORT Matrix4x4 (*__cross_call_master_CallFuncMat4x4Callback)(void*) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_master_CallFunc1Callback)(void*) = NULL;


PLUGIFY_EXPORT int8_t (*__cross_call_master_CallFunc2Callback)(void*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_CallFunc3Callback)(void*) = NULL;


PLUGIFY_EXPORT Vector4 (*__cross_call_master_CallFunc4Callback)(void*) = NULL;


PLUGIFY_EXPORT bool (*__cross_call_master_CallFunc5Callback)(void*) = NULL;


PLUGIFY_EXPORT int64_t (*__cross_call_master_CallFunc6Callback)(void*) = NULL;


PLUGIFY_EXPORT double (*__cross_call_master_CallFunc7Callback)(void*) = NULL;


PLUGIFY_EXPORT Matrix4x4 (*__cross_call_master_CallFunc8Callback)(void*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_master_CallFunc9Callback)(void*) = NULL;


PLUGIFY_EXPORT uint32_t (*__cross_call_master_CallFunc10Callback)(void*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_master_CallFunc11Callback)(void*) = NULL;


PLUGIFY_EXPORT bool (*__cross_call_master_CallFunc12Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc13Callback)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_master_CallFunc14Callback)(void*) = NULL;


PLUGIFY_EXPORT int16_t (*__cross_call_master_CallFunc15Callback)(void*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_master_CallFunc16Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc17Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc18Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc19Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc20Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc21Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc22Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc23Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc24Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc25Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc26Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc27Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc28Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc29Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc30Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc31Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc32Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFunc33Callback)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_master_CallFuncEnumCallback)(void*) = NULL;


