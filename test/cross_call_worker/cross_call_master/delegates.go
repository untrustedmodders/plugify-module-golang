package cross_call_master

// Generated from cross_call_master

type NoParamReturnFunctionCallbackFunc func() int32

type FuncVoid func()

type FuncBool func() bool

type FuncChar8 func() int8

type FuncChar16 func() uint16

type FuncInt8 func() int8

type FuncInt16 func() int16

type FuncInt32 func() int32

type FuncInt64 func() int64

type FuncUInt8 func() uint8

type FuncUInt16 func() uint16

type FuncUInt32 func() uint32

type FuncUInt64 func() uint64

type FuncPtr func() uintptr

type FuncFloat func() float32

type FuncDouble func() float64

type FuncString func() string

type FuncAny func() any

type FuncFunction func() FuncFunctionInner

type FuncBoolVector func() []bool

type FuncChar8Vector func() []int8

type FuncChar16Vector func() []uint16

type FuncInt8Vector func() []int8

type FuncInt16Vector func() []int16

type FuncInt32Vector func() []int32

type FuncInt64Vector func() []int64

type FuncUInt8Vector func() []uint8

type FuncUInt16Vector func() []uint16

type FuncUInt32Vector func() []uint32

type FuncUInt64Vector func() []uint64

type FuncPtrVector func() []uintptr

type FuncFloatVector func() []float32

type FuncDoubleVector func() []float64

type FuncStringVector func() []string

type FuncAnyVector func() []any

type FuncVec2Vector func() []plugify.Vector2

type FuncVec3Vector func() []plugify.Vector3

type FuncVec4Vector func() []plugify.Vector4

type FuncMat4x4Vector func() []plugify.Matrix4x4

type FuncVec2 func() plugify.Vector2

type FuncVec3 func() plugify.Vector3

type FuncVec4 func() plugify.Vector4

type FuncMat4x4 func() plugify.Matrix4x4

type Func1 func(a plugify.Vector3) int32

type Func2 func(a float32, b int64) int8

type Func3 func(a uintptr, b plugify.Vector4, c string)

type Func4 func(a bool, b int32, c uint16, d plugify.Matrix4x4) plugify.Vector4

type Func5 func(a int8, b plugify.Vector2, c uintptr, d float64, e []uint64) bool

type Func6 func(a string, b float32, c []float32, d int16, e []uint8, f uintptr) int64

type Func7 func(vecC []int8, u16 uint16, ch16 uint16, vecU32 []uint32, vec4 plugify.Vector4, b bool, u64 uint64) float64

type Func8 func(vec3 plugify.Vector3, vecU32 []uint32, i16 int16, b bool, vec4 plugify.Vector4, vecC16 []uint16, ch16 uint16, i32 int32) plugify.Matrix4x4

type Func9 func(f float32, vec2 plugify.Vector2, vecI8 []int8, u64 uint64, b bool, str string, vec4 plugify.Vector4, i16 int16, ptr uintptr)

type Func10 func(vec4 plugify.Vector4, mat plugify.Matrix4x4, vecU32 []uint32, u64 uint64, vecC []int8, i32 int32, b bool, vec2 plugify.Vector2, i64 int64, d float64) uint32

type Func11 func(vecB []bool, ch16 uint16, u8 uint8, d float64, vec3 plugify.Vector3, vecI8 []int8, i64 int64, u16 uint16, f float32, vec2 plugify.Vector2, u32 uint32) uintptr

type Func12 func(ptr uintptr, vecD []float64, u32 uint32, d float64, b bool, i32 int32, i8 int8, u64 uint64, f float32, vecPtr []uintptr, i64 int64, ch int8) bool

type Func13 func(i64 int64, vecC []int8, d uint16, f float32, b []bool, vec4 plugify.Vector4, str string, int32_ int32, vec3 plugify.Vector3, ptr uintptr, vec2 plugify.Vector2, arr []uint8, i16 int16) string

type Func14 func(vecC []int8, vecU32 []uint32, mat plugify.Matrix4x4, b bool, ch16 uint16, i32 int32, vecF []float32, u16 uint16, vecU8 []uint8, i8 int8, vec3 plugify.Vector3, vec4 plugify.Vector4, d float64, ptr uintptr) []string

type Func15 func(vecI16 []int16, mat plugify.Matrix4x4, vec4 plugify.Vector4, ptr uintptr, u64 uint64, vecU32 []uint32, b bool, f float32, vecC16 []uint16, u8 uint8, i32 int32, vec2 plugify.Vector2, u16 uint16, d float64, vecU8 []uint8) int16

type Func16 func(vecB []bool, i16 int16, vecI8 []int8, vec4 plugify.Vector4, mat plugify.Matrix4x4, vec2 plugify.Vector2, vecU64 []uint64, vecC []int8, str string, i64 int64, vecU32 []uint32, vec3 plugify.Vector3, f float32, d float64, i8 int8, u16 uint16) uintptr

type Func17 func(i32 *int32)

type Func18 func(i8 *int8, i16 *int16) plugify.Vector2

type Func19 func(u32 *uint32, vec3 *plugify.Vector3, vecU32 *[]uint32)

type Func20 func(ch16 *uint16, vec4 *plugify.Vector4, vecU64 *[]uint64, ch *int8) int32

type Func21 func(mat *plugify.Matrix4x4, vecI32 *[]int32, vec2 *plugify.Vector2, b *bool, extraParam *float64) float32

type Func22 func(ptr64Ref *uintptr, uint32Ref *uint32, vectorDoubleRef *[]float64, int16Ref *int16, plgStringRef *string, plgVector4Ref *plugify.Vector4) uint64

type Func23 func(uint64Ref *uint64, plgVector2Ref *plugify.Vector2, vectorInt16Ref *[]int16, char16Ref *uint16, floatRef *float32, int8Ref *int8, vectorUInt8Ref *[]uint8)

type Func24 func(vectorCharRef *[]int8, int64Ref *int64, vectorUInt8Ref *[]uint8, plgVector4Ref *plugify.Vector4, uint64Ref *uint64, vectorptr64Ref *[]uintptr, doubleRef *float64, vectorptr64Ref2 *[]uintptr) plugify.Matrix4x4

type Func25 func(int32Ref *int32, vectorptr64Ref *[]uintptr, boolRef *bool, uint8Ref *uint8, plgStringRef *string, plgVector3Ref *plugify.Vector3, int64Ref *int64, plgVector4Ref *plugify.Vector4, uint16Ref *uint16) float64

type Func26 func(char16Ref *uint16, plgVector2Ref *plugify.Vector2, plgMatrix4x4Ref *plugify.Matrix4x4, vectorFloatRef *[]float32, int16Ref *int16, uint64Ref *uint64, uint32Ref *uint32, vectorUInt16Ref *[]uint16, ptr64Ref *uintptr, boolRef *bool) int8

type Func27 func(floatRef *float32, plgVector3Ref *plugify.Vector3, ptr64Ref *uintptr, plgVector2Ref *plugify.Vector2, vectorInt16Ref *[]int16, plgMatrix4x4Ref *plugify.Matrix4x4, boolRef *bool, plgVector4Ref *plugify.Vector4, int8Ref *int8, int32Ref *int32, vectorUInt8Ref *[]uint8) uint8

type Func28 func(ptr64Ref *uintptr, uint16Ref *uint16, vectorUInt32Ref *[]uint32, plgMatrix4x4Ref *plugify.Matrix4x4, floatRef *float32, plgVector4Ref *plugify.Vector4, plgStringRef *string, vectorUInt64Ref *[]uint64, int64Ref *int64, boolRef *bool, plgVector3Ref *plugify.Vector3, vectorFloatRef *[]float32) string

type Func29 func(plgVector4Ref *plugify.Vector4, int32Ref *int32, vectorInt8Ref *[]int8, doubleRef *float64, boolRef *bool, int8Ref *int8, vectorUInt16Ref *[]uint16, floatRef *float32, plgStringRef *string, plgMatrix4x4Ref *plugify.Matrix4x4, uint64Ref *uint64, plgVector3Ref *plugify.Vector3, vectorInt64Ref *[]int64) []string

type Func30 func(ptr64Ref *uintptr, plgVector4Ref *plugify.Vector4, int64Ref *int64, vectorUInt32Ref *[]uint32, boolRef *bool, plgStringRef *string, plgVector3Ref *plugify.Vector3, vectorUInt8Ref *[]uint8, floatRef *float32, plgVector2Ref *plugify.Vector2, plgMatrix4x4Ref *plugify.Matrix4x4, int8Ref *int8, vectorFloatRef *[]float32, doubleRef *float64) int32

type Func31 func(charRef *int8, uint32Ref *uint32, vectorUInt64Ref *[]uint64, plgVector4Ref *plugify.Vector4, plgStringRef *string, boolRef *bool, int64Ref *int64, vec2Ref *plugify.Vector2, int8Ref *int8, uint16Ref *uint16, vectorInt16Ref *[]int16, mat4x4Ref *plugify.Matrix4x4, vec3Ref *plugify.Vector3, floatRef *float32, vectorDoubleRef *[]float64) plugify.Vector3

type Func32 func(p1 *int32, p2 *uint16, p3 *[]int8, p4 *plugify.Vector4, p5 *uintptr, p6 *[]uint32, p7 *plugify.Matrix4x4, p8 *uint64, p9 *string, p10 *int64, p11 *plugify.Vector2, p12 *[]int8, p13 *bool, p14 *plugify.Vector3, p15 *uint8, p16 *[]uint16) float64

type Func33 func(variant *any)

type FuncEnum func(p1 Example, p2 *[]Example) []Example

