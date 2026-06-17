#include "shared.h"

PLUGIFY_EXPORT void (*__cross_call_worker_NoParamReturnVoid)() = NULL;


PLUGIFY_EXPORT bool (*__cross_call_worker_NoParamReturnBool)() = NULL;


PLUGIFY_EXPORT int8_t (*__cross_call_worker_NoParamReturnChar8)() = NULL;


PLUGIFY_EXPORT uint16_t (*__cross_call_worker_NoParamReturnChar16)() = NULL;


PLUGIFY_EXPORT int8_t (*__cross_call_worker_NoParamReturnInt8)() = NULL;


PLUGIFY_EXPORT int16_t (*__cross_call_worker_NoParamReturnInt16)() = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_worker_NoParamReturnInt32)() = NULL;


PLUGIFY_EXPORT int64_t (*__cross_call_worker_NoParamReturnInt64)() = NULL;


PLUGIFY_EXPORT uint8_t (*__cross_call_worker_NoParamReturnUInt8)() = NULL;


PLUGIFY_EXPORT uint16_t (*__cross_call_worker_NoParamReturnUInt16)() = NULL;


PLUGIFY_EXPORT uint32_t (*__cross_call_worker_NoParamReturnUInt32)() = NULL;


PLUGIFY_EXPORT uint64_t (*__cross_call_worker_NoParamReturnUInt64)() = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_worker_NoParamReturnPointer)() = NULL;


PLUGIFY_EXPORT float (*__cross_call_worker_NoParamReturnFloat)() = NULL;


PLUGIFY_EXPORT double (*__cross_call_worker_NoParamReturnDouble)() = NULL;


PLUGIFY_EXPORT void* (*__cross_call_worker_NoParamReturnFunction)() = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_NoParamReturnString)() = NULL;


PLUGIFY_EXPORT Variant (*__cross_call_worker_NoParamReturnAny)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayBool)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayChar8)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayChar16)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayInt8)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayInt16)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayInt32)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayInt64)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayUInt8)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayUInt16)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayUInt32)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayUInt64)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayPointer)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayFloat)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayDouble)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayString)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayAny)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayVector2)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayVector3)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayVector4)() = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_NoParamReturnArrayMatrix4x4)() = NULL;


PLUGIFY_EXPORT Vector2 (*__cross_call_worker_NoParamReturnVector2)() = NULL;


PLUGIFY_EXPORT Vector3 (*__cross_call_worker_NoParamReturnVector3)() = NULL;


PLUGIFY_EXPORT Vector4 (*__cross_call_worker_NoParamReturnVector4)() = NULL;


PLUGIFY_EXPORT Matrix4x4 (*__cross_call_worker_NoParamReturnMatrix4x4)() = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_Param1)(int32_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_Param2)(int32_t, float) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_Param3)(int32_t, float, double) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_Param4)(int32_t, float, double, Vector4*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_Param5)(int32_t, float, double, Vector4*, Vector*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_Param6)(int32_t, float, double, Vector4*, Vector*, int8_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_Param7)(int32_t, float, double, Vector4*, Vector*, int8_t, String*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_Param8)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_Param9)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t, int16_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_Param10)(int32_t, float, double, Vector4*, Vector*, int8_t, String*, uint16_t, int16_t, uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamRef1)(int32_t*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamRef2)(int32_t*, float*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamRef3)(int32_t*, float*, double*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamRef4)(int32_t*, float*, double*, Vector4*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamRef5)(int32_t*, float*, double*, Vector4*, Vector*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamRef6)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamRef7)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamRef8)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamRef9)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*, int16_t*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamRef10)(int32_t*, float*, double*, Vector4*, Vector*, int8_t*, String*, uint16_t*, int16_t*, uintptr_t*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamRefVectors)(Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*) = NULL;


PLUGIFY_EXPORT int64_t (*__cross_call_worker_ParamAllPrimitives)(bool, int8_t, uint16_t, int8_t, int16_t, int32_t, int64_t, uint8_t, uint16_t, uint32_t, uint64_t, uintptr_t, float, double) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_worker_ParamAllAliases)(bool, int8_t, uint16_t, int8_t, int16_t, int32_t, int64_t, uintptr_t, float, double, String*, Variant*, Vector2*, Vector3*, Vector4*, Matrix4x4*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*) = NULL;


PLUGIFY_EXPORT int64_t (*__cross_call_worker_ParamAllRefAliases)(bool*, int8_t*, uint16_t*, int8_t*, int16_t*, int32_t*, int64_t*, uintptr_t*, float*, double*, String*, Variant*, Vector2*, Vector3*, Vector4*, Matrix4x4*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*, Vector*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamVariant)(Variant*, Vector*) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_worker_ParamEnum)(int32_t, Vector*) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_worker_ParamEnumRef)(int32_t*, Vector*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ParamVariantRef)(Variant*, Vector*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_CallFuncVoid)(void*) = NULL;


PLUGIFY_EXPORT bool (*__cross_call_worker_CallFuncBool)(void*) = NULL;


PLUGIFY_EXPORT int8_t (*__cross_call_worker_CallFuncChar8)(void*) = NULL;


PLUGIFY_EXPORT uint16_t (*__cross_call_worker_CallFuncChar16)(void*) = NULL;


PLUGIFY_EXPORT int8_t (*__cross_call_worker_CallFuncInt8)(void*) = NULL;


PLUGIFY_EXPORT int16_t (*__cross_call_worker_CallFuncInt16)(void*) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_worker_CallFuncInt32)(void*) = NULL;


PLUGIFY_EXPORT int64_t (*__cross_call_worker_CallFuncInt64)(void*) = NULL;


PLUGIFY_EXPORT uint8_t (*__cross_call_worker_CallFuncUInt8)(void*) = NULL;


PLUGIFY_EXPORT uint16_t (*__cross_call_worker_CallFuncUInt16)(void*) = NULL;


PLUGIFY_EXPORT uint32_t (*__cross_call_worker_CallFuncUInt32)(void*) = NULL;


PLUGIFY_EXPORT uint64_t (*__cross_call_worker_CallFuncUInt64)(void*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_worker_CallFuncPtr)(void*) = NULL;


PLUGIFY_EXPORT float (*__cross_call_worker_CallFuncFloat)(void*) = NULL;


PLUGIFY_EXPORT double (*__cross_call_worker_CallFuncDouble)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFuncString)(void*) = NULL;


PLUGIFY_EXPORT Variant (*__cross_call_worker_CallFuncAny)(void*) = NULL;


PLUGIFY_EXPORT void* (*__cross_call_worker_CallFuncFunction)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncBoolVector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncChar8Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncChar16Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncInt8Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncInt16Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncInt32Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncInt64Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncUInt8Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncUInt16Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncUInt32Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncUInt64Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncPtrVector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncFloatVector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncDoubleVector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncStringVector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAnyVector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncVec2Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncVec3Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncVec4Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncMat4x4Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector2 (*__cross_call_worker_CallFuncVec2)(void*) = NULL;


PLUGIFY_EXPORT Vector3 (*__cross_call_worker_CallFuncVec3)(void*) = NULL;


PLUGIFY_EXPORT Vector4 (*__cross_call_worker_CallFuncVec4)(void*) = NULL;


PLUGIFY_EXPORT Matrix4x4 (*__cross_call_worker_CallFuncMat4x4)(void*) = NULL;


PLUGIFY_EXPORT bool (*__cross_call_worker_CallFuncAliasBool)(void*) = NULL;


PLUGIFY_EXPORT int8_t (*__cross_call_worker_CallFuncAliasChar8)(void*) = NULL;


PLUGIFY_EXPORT uint16_t (*__cross_call_worker_CallFuncAliasChar16)(void*) = NULL;


PLUGIFY_EXPORT int8_t (*__cross_call_worker_CallFuncAliasInt8)(void*) = NULL;


PLUGIFY_EXPORT int16_t (*__cross_call_worker_CallFuncAliasInt16)(void*) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_worker_CallFuncAliasInt32)(void*) = NULL;


PLUGIFY_EXPORT int64_t (*__cross_call_worker_CallFuncAliasInt64)(void*) = NULL;


PLUGIFY_EXPORT uint8_t (*__cross_call_worker_CallFuncAliasUInt8)(void*) = NULL;


PLUGIFY_EXPORT uint16_t (*__cross_call_worker_CallFuncAliasUInt16)(void*) = NULL;


PLUGIFY_EXPORT uint32_t (*__cross_call_worker_CallFuncAliasUInt32)(void*) = NULL;


PLUGIFY_EXPORT uint64_t (*__cross_call_worker_CallFuncAliasUInt64)(void*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_worker_CallFuncAliasPtr)(void*) = NULL;


PLUGIFY_EXPORT float (*__cross_call_worker_CallFuncAliasFloat)(void*) = NULL;


PLUGIFY_EXPORT double (*__cross_call_worker_CallFuncAliasDouble)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFuncAliasString)(void*) = NULL;


PLUGIFY_EXPORT Variant (*__cross_call_worker_CallFuncAliasAny)(void*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_worker_CallFuncAliasFunction)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasBoolVector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasChar8Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasChar16Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasInt8Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasInt16Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasInt32Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasInt64Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasUInt8Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasUInt16Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasUInt32Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasUInt64Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasPtrVector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasFloatVector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasDoubleVector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasStringVector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasAnyVector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasVec2Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasVec3Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasVec4Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFuncAliasMat4x4Vector)(void*) = NULL;


PLUGIFY_EXPORT Vector2 (*__cross_call_worker_CallFuncAliasVec2)(void*) = NULL;


PLUGIFY_EXPORT Vector3 (*__cross_call_worker_CallFuncAliasVec3)(void*) = NULL;


PLUGIFY_EXPORT Vector4 (*__cross_call_worker_CallFuncAliasVec4)(void*) = NULL;


PLUGIFY_EXPORT Matrix4x4 (*__cross_call_worker_CallFuncAliasMat4x4)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFuncAliasAll)(void*) = NULL;


PLUGIFY_EXPORT int32_t (*__cross_call_worker_CallFunc1)(void*) = NULL;


PLUGIFY_EXPORT int8_t (*__cross_call_worker_CallFunc2)(void*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_CallFunc3)(void*) = NULL;


PLUGIFY_EXPORT Vector4 (*__cross_call_worker_CallFunc4)(void*) = NULL;


PLUGIFY_EXPORT bool (*__cross_call_worker_CallFunc5)(void*) = NULL;


PLUGIFY_EXPORT int64_t (*__cross_call_worker_CallFunc6)(void*) = NULL;


PLUGIFY_EXPORT double (*__cross_call_worker_CallFunc7)(void*) = NULL;


PLUGIFY_EXPORT Matrix4x4 (*__cross_call_worker_CallFunc8)(void*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_CallFunc9)(void*) = NULL;


PLUGIFY_EXPORT uint32_t (*__cross_call_worker_CallFunc10)(void*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_worker_CallFunc11)(void*) = NULL;


PLUGIFY_EXPORT bool (*__cross_call_worker_CallFunc12)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc13)(void*) = NULL;


PLUGIFY_EXPORT Vector (*__cross_call_worker_CallFunc14)(void*) = NULL;


PLUGIFY_EXPORT int16_t (*__cross_call_worker_CallFunc15)(void*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__cross_call_worker_CallFunc16)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc17)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc18)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc19)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc20)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc21)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc22)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc23)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc24)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc25)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc26)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc27)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc28)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc29)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc30)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc31)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc32)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFunc33)(void*) = NULL;


PLUGIFY_EXPORT String (*__cross_call_worker_CallFuncEnum)(void*) = NULL;


PLUGIFY_EXPORT void (*__cross_call_worker_ReverseCall)(String*) = NULL;


