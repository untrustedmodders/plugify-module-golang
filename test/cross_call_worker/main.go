package main

import "C"
import (
	"fmt"
	"github.com/untrustedmodders/go-plugify"
	"math"
	"plugify-plugin/cross_call_master"
	"runtime/debug"
	"strconv"
	"strings"
	"unsafe"
)

type Example = int32

const (
	First  Example = 1
	Second Example = 2
	Third  Example = 3
	Forth  Example = 4
)

type NoParamReturnFunctionFunc func()
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
type FuncFunction func() uintptr
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
type Func1 func(*plugify.Vector3) int32
type Func2 func(float32, int64) int8
type Func3 func(uintptr, *plugify.Vector4, string)
type Func4 func(bool, int32, uint16, *plugify.Matrix4x4) plugify.Vector4
type Func5 func(int8, *plugify.Vector2, uintptr, float64, []uint64) bool
type Func6 func(string, float32, []float32, int16, []uint8, uintptr) int64
type Func7 func([]int8, uint16, uint16, []uint32, *plugify.Vector4, bool, uint64) float64
type Func8 func(*plugify.Vector3, []uint32, int16, bool, *plugify.Vector4, []uint16, uint16, int32) plugify.Matrix4x4
type Func9 func(float32, *plugify.Vector2, []int8, uint64, bool, string, *plugify.Vector4, int16, uintptr)
type Func10 func(*plugify.Vector4, *plugify.Matrix4x4, []uint32, uint64, []int8, int32, bool, *plugify.Vector2, int64, float64) uint32
type Func11 func([]bool, uint16, uint8, float64, *plugify.Vector3, []int8, int64, uint16, float32, *plugify.Vector2, uint32) uintptr
type Func12 func(uintptr, []float64, uint32, float64, bool, int32, int8, uint64, float32, []uintptr, int64, int8) bool
type Func13 func(int64, []int8, uint16, float32, []bool, *plugify.Vector4, string, int32, *plugify.Vector3, uintptr, *plugify.Vector2, []uint8, int16) string
type Func14 func([]int8, []uint32, *plugify.Matrix4x4, bool, uint16, int32, []float32, uint16, []uint8, int8, *plugify.Vector3, *plugify.Vector4, float64, uintptr) []string
type Func15 func([]int16, *plugify.Matrix4x4, *plugify.Vector4, uintptr, uint64, []uint32, bool, float32, []uint16, uint8, int32, *plugify.Vector2, uint16, float64, []uint8) int16
type Func16 func([]bool, int16, []int8, *plugify.Vector4, *plugify.Matrix4x4, *plugify.Vector2, []uint64, []int8, string, int64, []uint32, *plugify.Vector3, float32, float64, int8, uint16) uintptr
type Func17 func(*int32)
type Func18 func(*int8, *int16) plugify.Vector2
type Func19 func(*uint32, *plugify.Vector3, *[]uint32)
type Func20 func(*uint16, *plugify.Vector4, *[]uint64, *int8) int32
type Func21 func(*plugify.Matrix4x4, *[]int32, *plugify.Vector2, *bool, *float64) float32
type Func22 func(*uintptr, *uint32, *[]float64, *int16, *string, *plugify.Vector4) uint64
type Func23 func(*uint64, *plugify.Vector2, *[]int16, *uint16, *float32, *int8, *[]uint8)
type Func24 func(*[]int8, *int64, *[]uint8, *plugify.Vector4, *uint64, *[]uintptr, *float64, *[]uintptr) plugify.Matrix4x4
type Func25 func(*int32, *[]uintptr, *bool, *uint8, *string, *plugify.Vector3, *int64, *plugify.Vector4, *uint16) float64
type Func26 func(*uint16, *plugify.Vector2, *plugify.Matrix4x4, *[]float32, *int16, *uint64, *uint32, *[]uint16, *uintptr, *bool) int8
type Func27 func(*float32, *plugify.Vector3, *uintptr, *plugify.Vector2, *[]int16, *plugify.Matrix4x4, *bool, *plugify.Vector4, *int8, *int32, *[]uint8) uint8
type Func28 func(*uintptr, *uint16, *[]uint32, *plugify.Matrix4x4, *float32, *plugify.Vector4, *string, *[]uint64, *int64, *bool, *plugify.Vector3, *[]float32) string
type Func29 func(*plugify.Vector4, *int32, *[]int8, *float64, *bool, *int8, *[]uint16, *float32, *string, *plugify.Matrix4x4, *uint64, *plugify.Vector3, *[]int64) []string
type Func30 func(*uintptr, *plugify.Vector4, *int64, *[]uint32, *bool, *string, *plugify.Vector3, *[]uint8, *float32, *plugify.Vector2, *plugify.Matrix4x4, *int8, *[]float32, *float64) int32
type Func31 func(*int8, *uint32, *[]uint64, *plugify.Vector4, *string, *bool, *int64, *plugify.Vector2, *int8, *uint16, *[]int16, *plugify.Matrix4x4, *plugify.Vector3, *float32, *[]float64) plugify.Vector3
type Func32 func(*int32, *uint16, *[]int8, *plugify.Vector4, *uintptr, *[]uint32, *plugify.Matrix4x4, *uint64, *string, *int64, *plugify.Vector2, *[]int8, *bool, *plugify.Vector3, *uint8, *[]uint16)
type Func33 func(*any)
type FuncEnum func(Example, *[]Example) []Example

func init() {
	plugify.OnPluginStart(func() {
		fmt.Println("Go: OnPluginStart")
	})

	plugify.OnPluginEnd(func() {
		fmt.Println("Go: OnPluginEnd")
	})

	plugify.OnPluginPanic(func() []byte {
		return debug.Stack() // workaround for could not import runtime/debug inside plugify package
	})
}

func NoParamReturnVoid() {
	fmt.Println("Go: NoParamReturnVoid")
}

func NoParamReturnBool() bool {
	fmt.Println("Go: NoParamReturnBool")
	return true
}

func NoParamReturnChar8() int8 {
	fmt.Println("Go: NoParamReturnChar16")
	return math.MaxInt8
}

func NoParamReturnChar16() uint16 {
	fmt.Println("Go: NoParamReturnChar16")
	return math.MaxUint16
}

func NoParamReturnInt8() int8 {
	fmt.Println("Go: NoParamReturnInt8")
	return math.MaxInt8
}

func NoParamReturnInt16() int16 {
	fmt.Println("Go: NoParamReturnInt16")
	return math.MaxInt16
}

func NoParamReturnInt32() int32 {
	fmt.Println("Go: NoParamReturnInt32")
	return math.MaxInt32
}

func NoParamReturnInt64() int64 {
	fmt.Println("Go: NoParamReturnInt64")
	return math.MaxInt64
}

func NoParamReturnUInt8() uint8 {
	fmt.Println("Go: NoParamReturnUInt8")
	return math.MaxUint8
}

func NoParamReturnUInt16() uint16 {
	fmt.Println("Go: NoParamReturnUInt16")
	return math.MaxUint16
}

func NoParamReturnUInt32() uint32 {
	fmt.Println("Go: NoParamReturnUInt32")
	return math.MaxUint32
}

func NoParamReturnUInt64() uint64 {
	fmt.Println("Go: NoParamReturnUInt64")
	return math.MaxUint64
}

func NoParamReturnPointer() uintptr {
	fmt.Println("Go: NoParamReturnPointer")
	return uintptr(1)
}

func NoParamReturnFloat() float32 {
	fmt.Println("Go: NoParamReturnFloat")
	return math.MaxFloat32
}

func NoParamReturnDouble() float64 {
	fmt.Println("Go: NoParamReturnDouble")
	return math.MaxFloat64
}

func NoParamReturnFunction() NoParamReturnFunctionFunc {
	fmt.Println("Go: NoParamReturnFunction")
	return func() {
		fmt.Println("Go: NoParamReturnFunctionFunc")
	}
}

func NoParamReturnString() string {
	fmt.Println("Go: NoParamReturnString")
	return "Hello World"
}

func NoParamReturnAny() any {
	fmt.Println("NoParamReturnAny")
	return []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
}

func NoParamReturnArrayBool() []bool {
	fmt.Println("Go: NoParamReturnArrayBool")
	return []bool{true, false}
}

func NoParamReturnArrayChar8() []int8 {
	fmt.Println("Go: NoParamReturnArrayChar8")
	return []int8{'a', 'b', 'c', 'd'}
}

func NoParamReturnArrayChar16() []uint16 {
	fmt.Println("Go: NoParamReturnArrayChar16")
	return []uint16{'a', 'b', 'c', 'd'}
}

func NoParamReturnArrayInt8() []int8 {
	fmt.Println("Go: NoParamReturnArrayInt8")
	return []int8{-3, -2, -1, 0, 1}
}

func NoParamReturnArrayInt16() []int16 {
	fmt.Println("Go: NoParamReturnArrayInt16")
	return []int16{-4, -3, -2, -1, 0, 1}
}

func NoParamReturnArrayInt32() []int32 {
	fmt.Println("Go: NoParamReturnArrayInt32")
	return []int32{-5, -4, -3, -2, -1, 0, 1}
}

func NoParamReturnArrayInt64() []int64 {
	fmt.Println("Go: NoParamReturnArrayInt64")
	return []int64{-6, -5, -4, -3, -2, -1, 0, 1}
}

func NoParamReturnArrayUInt8() []uint8 {
	fmt.Println("Go: NoParamReturnArrayUInt8")
	return []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8}
}

func NoParamReturnArrayUInt16() []uint16 {
	fmt.Println("Go: NoParamReturnArrayUInt16")
	return []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func NoParamReturnArrayUInt32() []uint32 {
	fmt.Println("Go: NoParamReturnArrayUInt32")
	return []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func NoParamReturnArrayUInt64() []uint64 {
	fmt.Println("Go: NoParamReturnArrayUInt64")
	return []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
}

func NoParamReturnArrayPointer() []uintptr {
	fmt.Println("Go: NoParamReturnArrayPointer")
	return []uintptr{0, 1, 2, 3}
}

func NoParamReturnArrayFloat() []float32 {
	fmt.Println("Go: NoParamReturnArrayFloat")
	return []float32{-12.34, 0.0, 12.34}
}

func NoParamReturnArrayDouble() []float64 {
	fmt.Println("Go: NoParamReturnArrayDouble")
	return []float64{-12.345, 0.0, 12.345}
}

func NoParamReturnArrayString() []string {
	fmt.Println("Go: NoParamReturnArrayString")
	return []string{
		"1st string", "2nd string",
		"3rd element string (Should be big enough to avoid small string optimization)",
	}
}

func NoParamReturnArrayAny() []any {
	fmt.Println("Go: NoParamReturnArrayAny")
	return []any{
		1.0,
		float32(2.0),
		"3rd element string (Should be big enough to avoid small string optimization)",
		[]string{"lolek", "and", "bolek"},
		int32(1),
	}
}

func NoParamReturnArrayVector2() []plugify.Vector2 {
	return []plugify.Vector2{
		{1.1, 2.2},
		{-3.3, 4.4},
		{5.5, -6.6},
		{7.7, 8.8},
		{0.0, 0.0},
	}
}

func NoParamReturnArrayVector3() []plugify.Vector3 {
	return []plugify.Vector3{
		{1.1, 2.2, 3.3},
		{-4.4, 5.5, -6.6},
		{7.7, 8.8, 9.9},
		{0.0, 0.0, 0.0},
		{10.1, -11.2, 12.3},
	}
}

func NoParamReturnArrayVector4() []plugify.Vector4 {
	return []plugify.Vector4{
		{1.1, 2.2, 3.3, 4.4},
		{-5.5, 6.6, -7.7, 8.8},
		{9.9, 0.0, -1.1, 2.2},
		{3.3, 4.4, 5.5, 6.6},
		{-7.7, -8.8, 9.9, -10.1},
	}
}

func NoParamReturnArrayMatrix4x4() []plugify.Matrix4x4 {
	return []plugify.Matrix4x4{
		{
			M: [4][4]float32{
				{1.0, 0.0, 0.0, 0.0},
				{0.0, 1.0, 0.0, 0.0},
				{0.0, 0.0, 1.0, 0.0},
				{0.0, 0.0, 0.0, 1.0},
			},
		},
		{
			M: [4][4]float32{
				{2.0, 3.0, 4.0, 5.0},
				{6.0, 7.0, 8.0, 9.0},
				{10.0, 11.0, 12.0, 13.0},
				{14.0, 15.0, 16.0, 17.0},
			},
		},
		{
			M: [4][4]float32{
				{-1.0, -2.0, -3.0, -4.0},
				{-5.0, -6.0, -7.0, -8.0},
				{-9.0, -10.0, -11.0, -12.0},
				{-13.0, -14.0, -15.0, -16.0},
			},
		},
	}
}

func NoParamReturnVector2() plugify.Vector2 {
	fmt.Println("Go: NoParamReturnVector2")
	return plugify.Vector2{X: 1, Y: 2}
}

func NoParamReturnVector3() plugify.Vector3 {
	fmt.Println("Go: NoParamReturnVector3")
	return plugify.Vector3{X: 1, Y: 2, Z: 3}
}

func NoParamReturnVector4() plugify.Vector4 {
	fmt.Println("Go: NoParamReturnVector4")
	return plugify.Vector4{X: 1, Y: 2, Z: 3, W: 4}
}

func NoParamReturnMatrix4x4() plugify.Matrix4x4 {
	fmt.Println("Go: NoParamReturnMatrix4x4")
	return plugify.Matrix4x4{
		M: [4][4]float32{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
	}
}

func Param1(a int32) {
	fmt.Printf("Param1: a = %d\n", a)
}

func Param2(a int32, b float32) {
	fmt.Printf("Param2: a = %d, b = %f\n", a, b)
}

func Param3(a int32, b float32, c float64) {
	fmt.Printf("Param3: a = %d, b = %f, c = %f\n", a, b, c)
}

func Param4(a int32, b float32, c float64, d plugify.Vector4) {
	fmt.Printf("Param4: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f]\n", a, b, c, formatVector4(d))
}

func Param5(a int32, b float32, c float64, d plugify.Vector4, e []int64) {
	fmt.Printf("Param5: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, formatVector4(d), len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Println("]")
}

func Param6(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8) {
	fmt.Printf("Param6: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, formatVector4(d), len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c\n", f)
}

func Param7(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string) {
	fmt.Printf("Param7: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, formatVector4(d), len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s\n", f, g)
}

func Param8(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16) {
	fmt.Printf("Param8: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, formatVector4(d), len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s, h = %f\n", f, g, h)
}

func Param9(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16, k int16) {
	fmt.Printf("Param9: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, formatVector4(d), len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s, h = %f, k = %d\n", f, g, h, k)
}

func Param10(a int32, b float32, c float64, d plugify.Vector4, e []int64, f int8, g string, h uint16, k int16, l uintptr) {
	fmt.Printf("Param10: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, formatVector4(d), len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s, h = %f, k = %d, l = %v\n", f, g, h, k, l)
}

func ParamRef1(a *int32) {
	*a = 42
}

func ParamRef2(a *int32, b *float32) {
	*a = 10
	*b = 3.14
}

func ParamRef3(a *int32, b *float32, c *float64) {
	*a = -20
	*b = 2.718
	*c = 3.14159
}

func ParamRef4(a *int32, b *float32, c *float64, d *plugify.Vector4) {
	*a = 100
	*b = -5.55
	*c = 1.618
	*d = plugify.Vector4{1.0, 2.0, 3.0, 4.0}
}

func ParamRef5(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64) {
	*a = 500
	*b = -10.5
	*c = 2.71828
	*d = plugify.Vector4{-1.0, -2.0, -3.0, -4.0}
	*e = []int64{-6, -5, -4, -3, -2, -1, 0, 1}
}

func ParamRef6(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8) {
	*a = 750
	*b = 20.0
	*c = 1.23456
	*d = plugify.Vector4{10.0, 20.0, 30.0, 40.0}
	*e = []int64{-6, -5, -4}
	*f = 'Z'
}

func ParamRef7(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string) {
	*a = -1000
	*b = 3.0
	*c = -1.0
	*d = plugify.Vector4{100.0, 200.0, 300.0, 400.0}
	*e = []int64{-6, -5, -4, -3}
	*f = 'Y'
	*g = "Hello, World!"
}

func ParamRef8(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16) {
	*a = 999
	*b = -7.5
	*c = 0.123456
	*d = plugify.Vector4{-100.0, -200.0, -300.0, -400.0}
	*e = []int64{-6, -5, -4, -3, -2, -1}
	*f = 'X'
	*g = "Goodbye, World!"
	*h = 'A'
}

func ParamRef9(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16) {
	*a = -1234
	*b = 123.45
	*c = -678.9
	*d = plugify.Vector4{987.65, 432.1, 123.456, 789.123}
	*e = []int64{-6, -5, -4, -3, -2, -1, 0, 1, 5, 9}
	*f = 'W'
	*g = "Testing, 1 2 3"
	*h = 'B'
	*k = 42
}

func ParamRef10(a *int32, b *float32, c *float64, d *plugify.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16, l *uintptr) {
	*a = 987
	*b = -0.123
	*c = 456.789
	*d = plugify.Vector4{-123.456, 0.987, 654.321, -789.123}
	*e = []int64{-6, -5, -4, -3, -2, -1, 0, 1, 5, 9, 4, -7}
	*f = 'V'
	*g = "Another string"
	*h = 'C'
	*k = -444
	*l = 0x12345678
}

func ParamRefVectors(p1 *[]bool, p2 *[]int8, p3 *[]uint16, p4 *[]int8, p5 *[]int16, p6 *[]int32, p7 *[]int64, p8 *[]uint8, p9 *[]uint16, p10 *[]uint32, p11 *[]uint64, p12 *[]uintptr, p13 *[]float32, p14 *[]float64, p15 *[]string) {
	*p1 = []bool{true}
	*p2 = []int8{'a', 'b', 'c'}
	*p3 = []uint16{'d', 'e', 'f'}
	*p4 = []int8{-3, -2, -1, 0, 1, 2, 3}
	*p5 = []int16{-4, -3, -2, -1, 0, 1, 2, 3, 4}
	*p6 = []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	*p7 = []int64{-6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6}
	*p8 = []uint8{0, 1, 2, 3, 4, 5, 6, 7}
	*p9 = []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8}
	*p10 = []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	*p11 = []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	*p12 = []uintptr{0, 1, 2}
	*p13 = []float32{-12.34, 0.0, 12.34}
	*p14 = []float64{-12.345, 0.0, 12.345}
	*p15 = []string{"1", "12", "123", "1234", "12345", "123456"}
}

func ParamAllPrimitives(p1 bool, p2 int8, p3 uint16, p4 int8, p5 int16, p6 int32, p7 int64, p8 uint8, p9 uint16, p10 uint32, p11 uint64, p12 uintptr, p13 float32, p14 float64) int64 {
	buffer := fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14)
	fmt.Println(buffer)
	return 56
}

func ParamEnum(p1 Example, p2 []Example) int32 {
	sum := p1
	for _, e := range p2 {
		sum += e
	}
	return int32(sum)
}

func ParamEnumRef(p1 *Example, p2 *[]Example) int32 {
	*p1 = Forth
	*p2 = []Example{First, Second, Third}
	sum := *p1
	for _, e := range *p2 {
		sum += e
	}
	return int32(sum)
}

func ParamVariant(p1 any, p2 []any) {
	buffer := fmt.Sprintf("%v%v", p1, p2)
	fmt.Println(buffer)
}

func ParamVariantRef(p1 *any, p2 *[]any) {
	*p1 = 'Z'
	*p2 = []any{false, 6.28, []float64{1, 2, 3}, uintptr(unsafe.Pointer(nil)), 123456789}
}

func CallFuncVoid(func_ FuncVoid) {
	func_()
}

func CallFuncBool(func_ FuncBool) bool {
	return func_()
}

func CallFuncChar8(func_ FuncChar8) int8 {
	return func_()
}

func CallFuncChar16(func_ FuncChar16) uint16 {
	return func_()
}

func CallFuncInt8(func_ FuncInt8) int8 {
	return func_()
}

func CallFuncInt16(func_ FuncInt16) int16 {
	return func_()
}

func CallFuncInt32(func_ FuncInt32) int32 {
	return func_()
}

func CallFuncInt64(func_ FuncInt64) int64 {
	return func_()
}

func CallFuncUInt8(func_ FuncUInt8) uint8 {
	return func_()
}

func CallFuncUInt16(func_ FuncUInt16) uint16 {
	return func_()
}

func CallFuncUInt32(func_ FuncUInt32) uint32 {
	return func_()
}

func CallFuncUInt64(func_ FuncUInt64) uint64 {
	return func_()
}

func CallFuncPtr(func_ FuncPtr) uintptr {
	return func_()
}

func CallFuncFloat(func_ FuncFloat) float32 {
	return func_()
}

func CallFuncDouble(func_ FuncDouble) float64 {
	return func_()
}

func CallFuncFunction(func_ FuncFunction) uintptr {
	return func_()
}

func CallFuncString(func_ FuncString) string {
	return func_()
}

func CallFuncAny(func_ FuncAny) any {
	return func_()
}

// Call functions for vector return types
func CallFuncBoolVector(func_ FuncBoolVector) []bool {
	result := func_()
	return result
}

func CallFuncChar8Vector(func_ FuncChar8Vector) []int8 {
	result := func_()
	return result
}

func CallFuncChar16Vector(func_ FuncChar16Vector) []uint16 {
	result := func_()
	return result
}

func CallFuncInt8Vector(func_ FuncInt8Vector) []int8 {
	result := func_()
	return result
}

func CallFuncInt16Vector(func_ FuncInt16Vector) []int16 {
	result := func_()
	return result
}

func CallFuncInt32Vector(func_ FuncInt32Vector) []int32 {
	result := func_()
	return result
}

func CallFuncInt64Vector(func_ FuncInt64Vector) []int64 {
	result := func_()
	return result
}

func CallFuncUInt8Vector(func_ FuncUInt8Vector) []uint8 {
	result := func_()
	return result
}

func CallFuncUInt16Vector(func_ FuncUInt16Vector) []uint16 {
	result := func_()
	return result
}

func CallFuncUInt32Vector(func_ FuncUInt32Vector) []uint32 {
	result := func_()
	return result
}

func CallFuncUInt64Vector(func_ FuncUInt64Vector) []uint64 {
	result := func_()
	return result
}

func CallFuncPtrVector(func_ FuncPtrVector) []uintptr {
	result := func_()
	return result
}

func CallFuncFloatVector(func_ FuncFloatVector) []float32 {
	result := func_()
	return result
}

func CallFuncDoubleVector(func_ FuncDoubleVector) []float64 {
	result := func_()
	return result
}

func CallFuncStringVector(func_ FuncStringVector) []string {
	result := func_()
	return result
}

func CallFuncAnyVector(func_ FuncAnyVector) []any {
	result := func_()
	return result
}

func CallFuncVec2Vector(func_ FuncVec2Vector) []plugify.Vector2 {
	result := func_()
	return result
}

func CallFuncVec3Vector(func_ FuncVec3Vector) []plugify.Vector3 {
	result := func_()
	return result
}

func CallFuncVec4Vector(func_ FuncVec4Vector) []plugify.Vector4 {
	result := func_()
	return result
}

func CallFuncMat4x4Vector(func_ FuncMat4x4Vector) []plugify.Matrix4x4 {
	result := func_()
	return result
}

func CallFuncVec2(func_ FuncVec2) plugify.Vector2 {
	result := func_()
	return result
}

func CallFuncVec3(func_ FuncVec3) plugify.Vector3 {
	result := func_()
	return result
}

func CallFuncVec4(func_ FuncVec4) plugify.Vector4 {
	result := func_()
	return result
}

func CallFuncMat4x4(func_ FuncMat4x4) plugify.Matrix4x4 {
	result := func_()
	return result
}

func CallFunc1(func_ Func1) int32 {
	vec := plugify.Vector3{4.5, 5.6, 6.7}
	return func_(&vec)
}

func CallFunc2(func_ Func2) int8 {
	f := float32(2.71)
	i64 := int64(200)
	return func_(f, i64)
}

func CallFunc3(func_ Func3) {
	ptr := uintptr(12345)
	vec4 := plugify.Vector4{7.8, 8.9, 9.1, 10.2}
	str := "RandomString"
	func_(ptr, &vec4, str)
}

func CallFunc4(func_ Func4) plugify.Vector4 {
	b := false
	u32 := int32(42)
	ch16 := uint16('B')
	mat := plugify.Matrix4x4Identity()
	return func_(b, u32, ch16, &mat)
}

func CallFunc5(func_ Func5) bool {
	i8 := int8(10)
	vec2 := plugify.Vector2{3.4, 5.6}
	ptr := uintptr(67890)
	d := 1.618
	vec64 := []uint64{4, 5, 6}
	return func_(i8, &vec2, ptr, d, vec64)
}

func CallFunc6(func_ Func6) int64 {
	str := "AnotherString"
	f := float32(4.56)
	vecF := []float32{4.0, 5.0, 6.0}
	i16 := int16(30)
	vecU8 := []uint8{3, 4, 5}
	ptr := uintptr(24680)
	return func_(str, f, vecF, i16, vecU8, ptr)
}

func CallFunc7(func_ Func7) float64 {
	vecC := []int8{'X', 'Y', 'Z'}
	u16 := uint16(20)
	ch16 := uint16('C')
	vecU32 := []uint32{4, 5, 6}
	vec4 := plugify.Vector4{4.5, 5.6, 6.7, 7.8}
	b := false
	u64 := uint64(200)
	return func_(vecC, u16, ch16, vecU32, &vec4, b, u64)
}

func CallFunc8(func_ Func8) plugify.Matrix4x4 {
	vec3 := plugify.Vector3{4.0, 5.0, 6.0}
	vecU32 := []uint32{4, 5, 6}
	i16 := int16(30)
	b := false
	vec4 := plugify.Vector4{4.5, 5.6, 6.7, 7.8}
	vecC16 := []uint16{'D', 'E'}
	ch16 := uint16('B')
	i32 := int32(50)
	return func_(&vec3, vecU32, i16, b, &vec4, vecC16, ch16, i32)
}

func CallFunc9(func_ Func9) {
	f := float32(2.71)
	vec2 := plugify.Vector2{3.4, 5.6}
	vecI8 := []int8{4, 5, 6}
	u64 := uint64(250)
	b := false
	str := "Random"
	vec4 := plugify.Vector4{4.5, 5.6, 6.7, 7.8}
	i16 := int16(30)
	ptr := uintptr(13579)
	func_(f, &vec2, vecI8, u64, b, str, &vec4, i16, ptr)
}

func CallFunc10(func_ Func10) uint32 {
	vec4 := plugify.Vector4{5.6, 7.8, 8.9, 9.0}
	mat := plugify.Matrix4x4Identity()
	vecU32 := []uint32{4, 5, 6}
	u64 := uint64(150)
	vecC := []int8{'X', 'Y', 'Z'}
	i32 := int32(60)
	b := false
	vec2 := plugify.Vector2{3.4, 5.6}
	i64 := int64(75)
	d := 2.71
	return func_(&vec4, &mat, vecU32, u64, vecC, i32, b, &vec2, i64, d)
}

func CallFunc11(func_ Func11) uintptr {
	vecB := []bool{false, true, false}
	ch16 := uint16('C')
	u8 := uint8(10)
	d := 2.71
	vec3 := plugify.Vector3{4.0, 5.0, 6.0}
	vecI8 := []int8{3, 4, 5}
	i64 := int64(150)
	u16 := uint16(20)
	f := float32(2.0)
	vec2 := plugify.Vector2{4.5, 6.7}
	u32 := uint32(30)
	return func_(vecB, ch16, u8, d, &vec3, vecI8, i64, u16, f, &vec2, u32)
}

func CallFunc12(func_ Func12) bool {
	ptr := uintptr(98765)
	vecD := []float64{4.0, 5.0, 6.0}
	u32 := uint32(30)
	d := 1.41
	b := false
	i32 := int32(25)
	i8 := int8(10)
	u64 := uint64(300)
	f := float32(2.72)
	vecPtr := []uintptr{2, 3, 4}
	i64 := int64(200)
	ch := int8('B')
	return func_(ptr, vecD, u32, d, b, i32, i8, u64, f, vecPtr, i64, ch)
}

func CallFunc13(func_ Func13) string {
	i64 := int64(75)
	vecC := []int8{'D', 'E', 'F'}
	u16 := uint16(20)
	f := float32(2.71)
	vecB := []bool{false, true, false}
	vec4 := plugify.Vector4{5.6, 7.8, 9.0, 10.1}
	str := "RandomString"
	i32 := int32(30)
	vec3 := plugify.Vector3{4.0, 5.0, 6.0}
	ptr := uintptr(13579)
	vec2 := plugify.Vector2{4.5, 6.7}
	vecU8 := []uint8{2, 3, 4}
	i16 := int16(20)
	return func_(i64, vecC, u16, f, vecB, &vec4, str, i32, &vec3, ptr, &vec2, vecU8, i16)
}

func CallFunc14(func_ Func14) []string {
	vecC := []int8{'D', 'E', 'F'}
	vecU32 := []uint32{4, 5, 6}
	mat := plugify.Matrix4x4Identity()
	b := false
	ch16 := uint16('B')
	i32 := int32(25)
	vecF := []float32{4.0, 5.0, 6.0}
	u16 := uint16(30)
	vecU8 := []uint8{3, 4, 5}
	i8 := int8(10)
	vec3 := plugify.Vector3{4.0, 5.0, 6.0}
	vec4 := plugify.Vector4{5.6, 7.8, 9.0, 10.1}
	d := 2.72
	ptr := uintptr(54321)
	return func_(vecC, vecU32, &mat, b, ch16, i32, vecF, u16, vecU8, i8, &vec3, &vec4, d, ptr)
}

func CallFunc15(func_ Func15) int16 {
	vecI16 := []int16{4, 5, 6}
	mat := plugify.Matrix4x4Identity()
	vec4 := plugify.Vector4{7.8, 8.9, 9.0, 10.1}
	ptr := uintptr(12345)
	u64 := uint64(200)
	vecU32 := []uint32{5, 6, 7}
	b := false
	f := float32(3.14)
	vecC16 := []uint16{'D', 'E'}
	u8 := uint8(6)
	i32 := int32(25)
	vec2 := plugify.Vector2{5.6, 7.8}
	u16 := uint16(40)
	d := 2.71
	vecU8 := []uint8{1, 3, 5}
	return func_(vecI16, &mat, &vec4, ptr, u64, vecU32, b, f, vecC16, u8, i32, &vec2, u16, d, vecU8)
}

func CallFunc16(func_ Func16) uintptr {
	vecB := []bool{true, true, false}
	i16 := int16(20)
	vecI8 := []int8{2, 3, 4}
	vec4 := plugify.Vector4{7.8, 8.9, 9.0, 10.1}
	mat := plugify.Matrix4x4Identity()
	vec2 := plugify.Vector2{5.6, 7.8}
	vecU64 := []uint64{5, 6, 7}
	vecC := []int8{'D', 'E', 'F'}
	str := "DifferentString"
	i64 := int64(300)
	vecU32 := []uint32{6, 7, 8}
	vec3 := plugify.Vector3{5.0, 6.0, 7.0}
	f := float32(3.14)
	d := 2.718
	i8 := int8(6)
	u16 := uint16(30)
	return func_(vecB, i16, vecI8, &vec4, &mat, &vec2, vecU64, vecC, str, i64, vecU32, &vec3, f, d, i8, u16)
}

// Implement the functions
func CallFunc17(func_ Func17) string {
	i32 := int32(42)
	func_(&i32)
	return fmt.Sprintf("%d", i32)
}

func CallFunc18(func_ Func18) string {
	i8 := int8(9)
	i16 := int16(25)
	ret := func_(&i8, &i16)
	return fmt.Sprintf("%s|%d|%d", formatVector2(ret), i8, i16)
}

func CallFunc19(func_ Func19) string {
	u32 := uint32(75)
	vec3 := plugify.Vector3{4.0, 5.0, 6.0}
	vecU32 := []uint32{4, 5, 6}
	func_(&u32, &vec3, &vecU32)
	return fmt.Sprintf("%d|%v|%v", u32, formatVector3(vec3), formatArray(vecU32))
}

func CallFunc20(func_ Func20) string {
	ch16 := uint16('Z')
	vec4 := plugify.Vector4{5.0, 6.0, 7.0, 8.0}
	vecU64 := []uint64{4, 5, 6}
	ch := int8('X')
	ret := func_(&ch16, &vec4, &vecU64, &ch)
	return fmt.Sprintf("%d|%d|%v|%v|%c", ret, ch16, formatVector4(vec4), formatArray(vecU64), ch)
}

func CallFunc21(func_ Func21) string {
	mat := plugify.Matrix4x4{}
	vecI32 := []int32{4, 5, 6}
	vec2 := plugify.Vector2{3.0, 4.0}
	b := false
	d := 6.28
	ret := func_(&mat, &vecI32, &vec2, &b, &d)
	return fmt.Sprintf("%v|%v|%v|%v|%t|%v", formatFlt32(ret), formatMatrix4x4(mat), formatArray(vecI32), formatVector2(vec2), b, formatFlt64(d))
}

func CallFunc22(func_ Func22) string {
	ptr := uintptr(1)
	u32 := uint32(20)
	vecD := []float64{4.0, 5.0, 6.0}
	i16 := int16(15)
	str := "Updated Test"
	vec4 := plugify.Vector4{5.0, 6.0, 7.0, 8.0}
	ret := func_(&ptr, &u32, &vecD, &i16, &str, &vec4)
	return fmt.Sprintf("%d|0x%x|%d|%v|%d|%v|%v", ret, ptr, u32, formatArrayFmt(vecD, formatFlt64), i16, str, formatVector4(vec4))
}

func CallFunc23(func_ Func23) string {
	u64 := uint64(200)
	vec2 := plugify.Vector2{3.0, 4.0}
	vecI16 := []int16{4, 5, 6}
	ch16 := uint16('Y')
	f := float32(2.34)
	i8 := int8(10)
	vecU8 := []uint8{3, 4, 5}
	func_(&u64, &vec2, &vecI16, &ch16, &f, &i8, &vecU8)
	return fmt.Sprintf("%d|%v|%v|%d|%v|%d|%v", u64, formatVector2(vec2), formatArray(vecI16), ch16, formatFlt32(f), i8, formatArray(vecU8))
}

func CallFunc24(func_ Func24) string {
	vecC := []int8{'D', 'E', 'F'}
	i64 := int64(100)
	vecU8 := []uint8{3, 4, 5}
	vec4 := plugify.Vector4{5.0, 6.0, 7.0, 8.0}
	u64 := uint64(200)
	vecPtr := []uintptr{3, 4, 5}
	d := 6.28
	vecV2 := []uintptr{4, 5, 6, 7}
	ret := func_(&vecC, &i64, &vecU8, &vec4, &u64, &vecPtr, &d, &vecV2)
	return fmt.Sprintf("%v|%v|%d|%v|%v|%d|%v|%v|%v", formatMatrix4x4(ret), formatArrayFmt(vecC, formatChr), i64, formatArray(vecU8), formatVector4(vec4), u64, formatArrayFmt(vecPtr, formatPtr), formatFlt64(d), formatArrayFmt(vecV2, formatPtr))
}

func CallFunc25(func_ Func25) string {
	i32 := int32(50)
	vecPtr := []uintptr{3, 4, 5}
	b := false
	u8 := uint8(10)
	str := "Updated Test String"
	vec3 := plugify.Vector3{4.0, 5.0, 6.0}
	i64 := int64(100)
	vec4 := plugify.Vector4{5.0, 6.0, 7.0, 8.0}
	u16 := uint16(20)
	ret := func_(&i32, &vecPtr, &b, &u8, &str, &vec3, &i64, &vec4, &u16)
	return fmt.Sprintf("%v|%d|%v|%t|%d|%v|%v|%d|%v|%d", formatFlt64(ret), i32, formatArrayFmt(vecPtr, formatPtr), b, u8, str, formatVector3(vec3), i64, formatVector4(vec4), u16)
}

func CallFunc26(func_ Func26) string {
	ch16 := uint16('B')
	vec2 := plugify.Vector2{3.0, 4.0}
	mat := plugify.Matrix4x4{}
	vecF := []float32{4.0, 5.0, 6.0}
	i16 := int16(20)
	u64 := uint64(200)
	u32 := uint32(20)
	vecU16 := []uint16{3, 4, 5}
	ptr := uintptr(0xDEADBEAFDEADBEAF)
	b := false
	ret := func_(&ch16, &vec2, &mat, &vecF, &i16, &u64, &u32, &vecU16, &ptr, &b)
	return fmt.Sprintf("%c|%d|%v|%v|%v|%d|%d|%v|0x%x|%t", ret, ch16, formatVector2(vec2), formatMatrix4x4(mat), formatArrayFmt(vecF, formatFlt32), u64, u32, formatArray(vecU16), ptr, b)
}

func CallFunc27(func_ Func27) string {
	f := float32(2.56)
	vec3 := plugify.Vector3{4.0, 5.0, 6.0}
	ptr := uintptr(0)
	vec2 := plugify.Vector2{3.0, 4.0}
	vecI16 := []int16{4, 5, 6}
	mat := plugify.Matrix4x4{}
	b := false
	vec4 := plugify.Vector4{5.0, 6.0, 7.0, 8.0}
	i8 := int8(10)
	i32 := int32(40)
	vecU8 := []uint8{3, 4, 5}
	ret := func_(&f, &vec3, &ptr, &vec2, &vecI16, &mat, &b, &vec4, &i8, &i32, &vecU8)
	return fmt.Sprintf("%d|%v|%v|0x%x|%v|%v|%v|%t|%v|%d|%d|%v", ret, formatFlt32(f), formatVector3(vec3), ptr, formatVector2(vec2), formatArray(vecI16), formatMatrix4x4(mat), b, formatVector4(vec4), i8, i32, formatArray(vecU8))
}

func CallFunc28(func_ Func28) string {
	ptr := uintptr(1)
	u16 := uint16(20)
	vecU32 := []uint32{4, 5, 6}
	mat := plugify.Matrix4x4{}
	f := float32(2.71)
	vec4 := plugify.Vector4{5.0, 6.0, 7.0, 8.0}
	str := "New example string"
	vecU64 := []uint64{400, 500, 600}
	i64 := int64(987654321)
	b := false
	vec3 := plugify.Vector3{4.0, 5.0, 6.0}
	vecF := []float32{4.0, 5.0, 6.0}
	ret := func_(&ptr, &u16, &vecU32, &mat, &f, &vec4, &str, &vecU64, &i64, &b, &vec3, &vecF)
	return fmt.Sprintf("%v|0x%x|%d|%v|%v|%v|%v|%v|%v|%d|%t|%v|%v", ret, ptr, u16, formatArray(vecU32), formatMatrix4x4(mat), formatFlt32(f), formatVector4(vec4), str, formatArray(vecU64), i64, b, formatVector3(vec3), formatArrayFmt(vecF, formatFlt32))
}

func CallFunc29(func_ Func29) string {
	vec4 := plugify.Vector4{2.0, 3.0, 4.0, 5.0}
	i32 := int32(99)
	vecI8 := []int8{4, 5, 6}
	d := float64(2.71)
	b := false
	i8 := int8(10)
	vecU16 := []uint16{4, 5, 6}
	f := float32(3.21)
	str := "Yet another example string"
	mat := plugify.Matrix4x4{}
	u64 := uint64(200)
	vec3 := plugify.Vector3{5.0, 6.0, 7.0}
	vecI64 := []int64{2000, 3000, 4000}
	ret := func_(&vec4, &i32, &vecI8, &d, &b, &i8, &vecU16, &f, &str, &mat, &u64, &vec3, &vecI64)
	return fmt.Sprintf("%v|%v|%d|%v|%v|%t|%d|%v|%v|%s|%v|%d|%v|%v", formatArrayFmt(ret, formatStr), formatVector4(vec4), i32, formatArray(vecI8), formatFlt64(d), b, i8, formatArray(vecU16), formatFlt32(f), str, formatMatrix4x4(mat), u64, formatVector3(vec3), formatArray(vecI64))
}

func CallFunc30(func_ Func30) string {
	ptr := uintptr(1)
	vec4 := plugify.Vector4{2.0, 3.0, 4.0, 5.0}
	i64 := int64(987654321)
	vecU32 := []uint32{4, 5, 6}
	b := false
	str := "Updated String for Func30"
	vec3 := plugify.Vector3{5.0, 6.0, 7.0}
	vecU8 := []uint8{1, 2, 3}
	f := float32(5.67)
	vec2 := plugify.Vector2{3.0, 4.0}
	mat := plugify.Matrix4x4{}
	i8 := int8(10)
	vecF := []float32{4.0, 5.0, 6.0}
	d := float64(8.90)
	ret := func_(&ptr, &vec4, &i64, &vecU32, &b, &str, &vec3, &vecU8, &f, &vec2, &mat, &i8, &vecF, &d)
	return fmt.Sprintf("%d|0x%x|%v|%d|%v|%t|%s|%v|%v|%v|%v|%v|%d|%v|%v", ret, ptr, formatVector4(vec4), i64, formatArray(vecU32), b, str, formatVector3(vec3), formatArray(vecU8), formatFlt32(f), formatVector2(vec2), formatMatrix4x4(mat), i8, formatArrayFmt(vecF, formatFlt32), formatFlt64(d))
}

func CallFunc31(func_ Func31) string {
	ch := int8('B')
	u32 := uint32(200)
	vecU64 := []uint64{4, 5, 6}
	vec4 := plugify.Vector4{2.0, 3.0, 4.0, 5.0}
	str := "Updated String for Func31"
	b := true
	i64 := int64(987654321)
	vec2 := plugify.Vector2{3.0, 4.0}
	i8 := int8(10)
	u16 := uint16(20)
	vecI16 := []int16{4, 5, 6}
	mat := plugify.Matrix4x4{}
	vec3 := plugify.Vector3{4.0, 5.0, 6.0}
	f := float32(5.67)
	vecD := []float64{4.0, 5.0, 6.0}
	ret := func_(&ch, &u32, &vecU64, &vec4, &str, &b, &i64, &vec2, &i8, &u16, &vecI16, &mat, &vec3, &f, &vecD)
	return fmt.Sprintf("%v|%c|%d|%v|%v|%s|%t|%d|%v|%d|%d|%v|%v|%v|%v|%v", formatVector3(ret), ch, u32, formatArray(vecU64), formatVector4(vec4), str, b, i64, formatVector2(vec2), i8, u16, formatArray(vecI16), formatMatrix4x4(mat), formatVector3(vec3), formatFlt32(f), formatArrayFmt(vecD, formatFlt64))
}

func CallFunc32(func_ Func32) string {
	i32 := int32(30)
	u16 := uint16(20)
	vecI8 := []int8{4, 5, 6}
	vec4 := plugify.Vector4{2.0, 3.0, 4.0, 5.0}
	ptr := uintptr(1)
	vecU32 := []uint32{4, 5, 6}
	mat := plugify.Matrix4x4{}
	u64 := uint64(200)
	str := "Updated String for Func32"
	i64 := int64(987654321)
	vec2 := plugify.Vector2{3.0, 4.0}
	vecI8_2 := []int8{7, 8, 9}
	b := false
	vec3 := plugify.Vector3{4.0, 5.0, 6.0}
	u8 := uint8(128)
	vecC16 := []uint16{'D', 'E', 'F'}

	func_(&i32, &u16, &vecI8, &vec4, &ptr, &vecU32, &mat, &u64, &str, &i64, &vec2, &vecI8_2, &b, &vec3, &u8, &vecC16)
	return fmt.Sprintf("%d|%d|%v|%v|0x%x|%v|%v|%d|%s|%d|%v|%v|%t|%v|%d|%v", i32, u16, formatArray(vecI8), formatVector4(vec4), ptr, formatArray(vecU32), formatMatrix4x4(mat), u64, str, i64, formatVector2(vec2), formatArray(vecI8_2), b, formatVector3(vec3), u8, formatArray(vecC16))
}

func CallFunc33(func_ Func33) string {
	var variant any
	variant = int32(30)
	func_(&variant)
	return fmt.Sprintf("%v", variant)
}

func CallFuncEnum(func_ FuncEnum) string {
	p1 := Forth
	p2 := []Example{}
	ret := func_(p1, &p2)
	retStr := fmt.Sprintf("%v|%v", formatArray(ret), formatArray(p2))
	return retStr
}

func main() {}

//
//
//
//
// Reverse staff

// Define the function implementations
func ReverseNoParamReturnVoid() string {
	cross_call_master.NoParamReturnVoidCallback()
	return ""
}

func ReverseNoParamReturnBool() string {
	result := cross_call_master.NoParamReturnBoolCallback()
	return formatBool(result)
}

func ReverseNoParamReturnChar8() string {
	result := cross_call_master.NoParamReturnChar8Callback()
	return fmt.Sprintf("%c", result)
}

func ReverseNoParamReturnChar16() string {
	result := cross_call_master.NoParamReturnChar16Callback()
	return fmt.Sprintf("%d", result)
}

func ReverseNoParamReturnInt8() string {
	result := cross_call_master.NoParamReturnInt8Callback()
	return fmt.Sprintf("%d", result)
}

func ReverseNoParamReturnInt16() string {
	result := cross_call_master.NoParamReturnInt16Callback()
	return fmt.Sprintf("%d", result)
}

func ReverseNoParamReturnInt32() string {
	result := cross_call_master.NoParamReturnInt32Callback()
	return fmt.Sprintf("%d", result)
}

func ReverseNoParamReturnInt64() string {
	result := cross_call_master.NoParamReturnInt64Callback()
	return fmt.Sprintf("%d", result)
}

func ReverseNoParamReturnUInt8() string {
	result := cross_call_master.NoParamReturnUInt8Callback()
	return fmt.Sprintf("%d", result)
}

func ReverseNoParamReturnUInt16() string {
	result := cross_call_master.NoParamReturnUInt16Callback()
	return fmt.Sprintf("%d", result)
}

func ReverseNoParamReturnUInt32() string {
	result := cross_call_master.NoParamReturnUInt32Callback()
	return fmt.Sprintf("%d", result)
}

func ReverseNoParamReturnUInt64() string {
	result := cross_call_master.NoParamReturnUInt64Callback()
	return fmt.Sprintf("%d", result)
}

func ReverseNoParamReturnPointer() string {
	result := cross_call_master.NoParamReturnPointerCallback()
	return fmt.Sprintf("0x%x", result)
}

func ReverseNoParamReturnFloat() string {
	result := cross_call_master.NoParamReturnFloatCallback()
	return fmt.Sprintf("%v", formatFlt32(result))
}

func ReverseNoParamReturnDouble() string {
	result := cross_call_master.NoParamReturnDoubleCallback()
	return fmt.Sprintf("%v", formatFlt64(result))
}

func ReverseNoParamReturnFunction() string {
	result := cross_call_master.NoParamReturnFunctionCallback()
	return fmt.Sprintf("%d", result())
}

func ReverseNoParamReturnString() string {
	result := cross_call_master.NoParamReturnStringCallback()
	return result
}

func ReverseNoParamReturnAny() string {
	result := cross_call_master.NoParamReturnAnyCallback()
	return fmt.Sprintf("%v", result)
}

func ReverseNoParamReturnArrayBool() string {
	result := cross_call_master.NoParamReturnArrayBoolCallback()
	return formatArrayFmt(result, formatBool)
}

func ReverseNoParamReturnArrayChar8() string {
	result := cross_call_master.NoParamReturnArrayChar8Callback()
	return formatArrayFmt(result, formatChr)
}

func ReverseNoParamReturnArrayChar16() string {
	result := cross_call_master.NoParamReturnArrayChar16Callback()
	return formatArray(result)
}

func ReverseNoParamReturnArrayInt8() string {
	result := cross_call_master.NoParamReturnArrayInt8Callback()
	return formatArray(result)
}

func ReverseNoParamReturnArrayInt16() string {
	result := cross_call_master.NoParamReturnArrayInt16Callback()
	return formatArray(result)
}

func ReverseNoParamReturnArrayInt32() string {
	result := cross_call_master.NoParamReturnArrayInt32Callback()
	return formatArray(result)
}

func ReverseNoParamReturnArrayInt64() string {
	result := cross_call_master.NoParamReturnArrayInt64Callback()
	return formatArray(result)
}

func ReverseNoParamReturnArrayUInt8() string {
	result := cross_call_master.NoParamReturnArrayUInt8Callback()
	return formatArray(result)
}

func ReverseNoParamReturnArrayUInt16() string {
	result := cross_call_master.NoParamReturnArrayUInt16Callback()
	return formatArray(result)
}

func ReverseNoParamReturnArrayUInt32() string {
	result := cross_call_master.NoParamReturnArrayUInt32Callback()
	return formatArray(result)
}

func ReverseNoParamReturnArrayUInt64() string {
	result := cross_call_master.NoParamReturnArrayUInt64Callback()
	return formatArray(result)
}

func ReverseNoParamReturnArrayPointer() string {
	result := cross_call_master.NoParamReturnArrayPointerCallback()
	return formatArrayFmt(result, formatPtr)
}

func ReverseNoParamReturnArrayFloat() string {
	result := cross_call_master.NoParamReturnArrayFloatCallback()
	return formatArrayFmt(result, formatFlt32)
}

func ReverseNoParamReturnArrayDouble() string {
	result := cross_call_master.NoParamReturnArrayDoubleCallback()
	return formatArrayFmt(result, formatFlt64)
}

func ReverseNoParamReturnArrayString() string {
	result := cross_call_master.NoParamReturnArrayStringCallback()
	return formatArrayFmt(result, formatStr)
}

func ReverseNoParamReturnArrayAny() string {
	result := cross_call_master.NoParamReturnArrayAnyCallback()
	return formatArrayFmt(result, formatAny)
}

func ReverseNoParamReturnVector2() string {
	result := cross_call_master.NoParamReturnVector2Callback()
	return formatVector2(result)
}

func ReverseNoParamReturnVector3() string {
	result := cross_call_master.NoParamReturnVector3Callback()
	return formatVector3(result)
}

func ReverseNoParamReturnVector4() string {
	result := cross_call_master.NoParamReturnVector4Callback()
	return formatVector4(result)
}

func ReverseNoParamReturnMatrix4x4() string {
	result := cross_call_master.NoParamReturnMatrix4x4Callback()
	return formatMatrix4x4(result)
}

func ReverseParam1() string {
	cross_call_master.Param1Callback(999)
	return ""
}

func ReverseParam2() string {
	cross_call_master.Param2Callback(888, 9.9)
	return ""
}

func ReverseParam3() string {
	cross_call_master.Param3Callback(777, 8.8, 9.8765)
	return ""
}

func ReverseParam4() string {
	cross_call_master.Param4Callback(666, 7.7, 8.7659, plugify.Vector4{100.1, 200.2, 300.3, 400.4})
	return ""
}

func ReverseParam5() string {
	cross_call_master.Param5Callback(555, 6.6, 7.6598, plugify.Vector4{-105.1, -205.2, -305.3, -405.4}, []int64{})
	return ""
}

func ReverseParam6() string {
	cross_call_master.Param6Callback(444, 5.5, 6.5987, plugify.Vector4{110.1, 210.2, 310.3, 410.4}, []int64{90000, -100, 20000}, 'A')
	return ""
}

func ReverseParam7() string {
	cross_call_master.Param7Callback(333, 4.4, 5.9876, plugify.Vector4{-115.1, -215.2, -315.3, -415.4}, []int64{800000, 30000, -4000000}, 'B', "red gold")
	return ""
}

func ReverseParam8() string {
	cross_call_master.Param8Callback(222, 3.3, 1.2345, plugify.Vector4{120.1, 220.2, 320.3, 420.4}, []int64{7000000, 5000000, -600000000}, 'C', "blue ice", 'Z')
	return ""
}

func ReverseParam9() string {
	cross_call_master.Param9Callback(111, 2.2, 5.1234, plugify.Vector4{-125.1, -225.2, -325.3, -425.4}, []int64{60000000, -700000000, 80000000000}, 'D', "pink metal", 'Y', -100)
	return ""
}

func ReverseParam10() string {
	cross_call_master.Param10Callback(1234, 1.1, 4.5123, plugify.Vector4{130.1, 230.2, 330.3, 430.4}, []int64{500000000, 90000000000, 1000000000000}, 'E', "green wood", 'X', -200, 0xabeba)
	return ""
}

func ReverseParamRef1() string {
	a := int32(0)
	cross_call_master.ParamRef1Callback(&a)
	return strconv.Itoa(int(a))
}

func ReverseParamRef2() string {
	a, b := int32(0), float32(0)
	cross_call_master.ParamRef2Callback(&a, &b)
	return fmt.Sprintf("%d|%.1f", a, b)
}

func ReverseParamRef3() string {
	a, b, c := int32(0), float32(0), float64(0)
	cross_call_master.ParamRef3Callback(&a, &b, &c)
	return fmt.Sprintf("%d|%.1f|%.5f", a, b, c)
}

func ReverseParamRef4() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := plugify.Vector4{}
	cross_call_master.ParamRef4Callback(&a, &b, &c, &d)
	return fmt.Sprintf("%d|%.1f|%.5f|%v", a, b, c, formatVector4(d))
}

func ReverseParamRef5() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := plugify.Vector4{}
	e := []int64{}
	cross_call_master.ParamRef5Callback(&a, &b, &c, &d, &e)
	return fmt.Sprintf("%d|%.1f|%.5f|%v|%s", a, b, c, formatVector4(d), formatArray(e))
}

func ReverseParamRef6() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := plugify.Vector4{}
	e := []int64{}
	f := int8(0)
	cross_call_master.ParamRef6Callback(&a, &b, &c, &d, &e, &f)
	return fmt.Sprintf("%d|%.1f|%.5f|%v|%s|%d", a, b, c, formatVector4(d), formatArray(e), f)
}

func ReverseParamRef7() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := plugify.Vector4{}
	e := []int64{}
	f := int8(0)
	g := ""
	cross_call_master.ParamRef7Callback(&a, &b, &c, &d, &e, &f, &g)
	return fmt.Sprintf("%d|%.1f|%.5f|%v|%s|%d|%s", a, b, c, formatVector4(d), formatArray(e), f, g)
}

func ReverseParamRef8() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := plugify.Vector4{}
	e := []int64{}
	f := int8(0)
	g := ""
	h := uint16(0)
	cross_call_master.ParamRef8Callback(&a, &b, &c, &d, &e, &f, &g, &h)
	return fmt.Sprintf("%d|%.1f|%.5f|%v|%s|%d|%s|%d", a, b, c, formatVector4(d), formatArray(e), f, g, h)
}

func ReverseParamRef9() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := plugify.Vector4{}
	e := []int64{}
	f := int8(0)
	g := ""
	h := uint16(0)
	k := int16(0)
	cross_call_master.ParamRef9Callback(&a, &b, &c, &d, &e, &f, &g, &h, &k)
	return fmt.Sprintf("%d|%.1f|%.5f|%v|%s|%d|%s|%d|%d", a, b, c, formatVector4(d), formatArray(e), f, g, h, k)
}

func ReverseParamRef10() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := plugify.Vector4{}
	e := []int64{}
	f := int8(0)
	g := ""
	h := uint16(0)
	k := int16(0)
	l := uintptr(0)
	cross_call_master.ParamRef10Callback(&a, &b, &c, &d, &e, &f, &g, &h, &k, &l)
	return fmt.Sprintf("%d|%.1f|%.5f|%v|%s|%d|%s|%d|%d|0x%x", a, b, c, formatVector4(d), formatArray(e), f, g, h, k, l)
}

func formatBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func formatPtr(p uintptr) string {
	return fmt.Sprintf("0x%x", p)
}

func formatChr(c int8) string {
	return fmt.Sprintf("%c", c)
}

func formatStr(s string) string {
	return fmt.Sprintf("'%s'", s)
}

func formatFlt32(d float32) string {
	s := fmt.Sprintf("%.3f", d)
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	return s
}

func formatFlt64(d float64) string {
	s := fmt.Sprintf("%.6f", d)
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	return s
}

func formatAny(a any) string {
	return fmt.Sprintf("%v", a)
}

func formatVector2(t plugify.Vector2) string {
	return fmt.Sprintf("{%s, %s}", formatFlt32(t.X), formatFlt32(t.Y))
}

func formatVector3(t plugify.Vector3) string {
	return fmt.Sprintf("{%s, %s, %s}", formatFlt32(t.X), formatFlt32(t.Y), formatFlt32(t.Z))
}

func formatVector4(t plugify.Vector4) string {
	return fmt.Sprintf("{%s, %s, %s, %s}", formatFlt32(t.X), formatFlt32(t.Y), formatFlt32(t.Z), formatFlt32(t.W))
}

func formatMatrix4x4(t plugify.Matrix4x4) string {
	formattedRow1 := fmt.Sprintf("{%s, %s, %s, %s}", formatFlt32(t.M[0][0]), formatFlt32(t.M[0][1]), formatFlt32(t.M[0][2]), formatFlt32(t.M[0][3]))
	formattedRow2 := fmt.Sprintf("{%s, %s, %s, %s}", formatFlt32(t.M[1][0]), formatFlt32(t.M[1][1]), formatFlt32(t.M[1][2]), formatFlt32(t.M[1][3]))
	formattedRow3 := fmt.Sprintf("{%s, %s, %s, %s}", formatFlt32(t.M[2][0]), formatFlt32(t.M[2][1]), formatFlt32(t.M[2][2]), formatFlt32(t.M[2][3]))
	formattedRow4 := fmt.Sprintf("{%s, %s, %s, %s}", formatFlt32(t.M[3][0]), formatFlt32(t.M[3][1]), formatFlt32(t.M[3][2]), formatFlt32(t.M[3][3]))
	return fmt.Sprintf("{%s, %s, %s, %s}", formattedRow1, formattedRow2, formattedRow3, formattedRow4)
}

func formatArray[T any](arr []T) string {
	if len(arr) == 0 {
		return "{}"
	}

	str := "{"
	for _, v := range arr {
		str += fmt.Sprintf("%v, ", v)
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func formatArrayFmt[T any](arr []T, format func(T) string) string {
	if len(arr) == 0 {
		return "{}"
	}

	str := "{"
	for _, v := range arr {
		str += fmt.Sprintf("%s, ", format(v))
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseParamRefVectors() string {
	p1 := []bool{}
	p2 := []int8{}
	p3 := []uint16{}
	p4 := []int8{}
	p5 := []int16{}
	p6 := []int32{}
	p7 := []int64{}
	p8 := []uint8{}
	p9 := []uint16{}
	p10 := []uint32{}
	p11 := []uint64{}
	p12 := []uintptr{}
	p13 := []float32{}
	p14 := []float64{}
	p15 := []string{}

	cross_call_master.ParamRefVectorsCallback(&p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &p9, &p10, &p11, &p12, &p13, &p14, &p15)

	p1Formatted := formatArray(p1)
	p2Formatted := formatArrayFmt(p2, formatChr)
	p3Formatted := formatArray(p3)
	p4Formatted := formatArray(p4)
	p5Formatted := formatArray(p5)
	p6Formatted := formatArray(p6)
	p7Formatted := formatArray(p7)
	p8Formatted := formatArray(p8)
	p9Formatted := formatArray(p9)
	p10Formatted := formatArray(p10)
	p11Formatted := formatArray(p11)
	p12Formatted := formatArrayFmt(p12, formatPtr)
	p13Formatted := formatArray(p13)
	p14Formatted := formatArray(p14)
	p15Formatted := formatArrayFmt(p15, formatStr)

	return fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s",
		p1Formatted, p2Formatted, p3Formatted, p4Formatted, p5Formatted, p6Formatted, p7Formatted,
		p8Formatted, p9Formatted, p10Formatted, p11Formatted, p12Formatted, p13Formatted, p14Formatted,
		p15Formatted)
}

func ReverseParamAllPrimitives() string {
	result := cross_call_master.ParamAllPrimitivesCallback(
		true,           // bool
		'%',            // char8
		'☢',            // char16
		-1,             // int8
		-1000,          // int16
		-1000000,       // int32
		-1000000000000, // int64
		200,            // uint8
		50000,          // uint16
		3000000000,     // uint32
		9999999999,     // uint64
		0xfedcbaabcdef, // uintptr (used for pointer simulation)
		0.001,          // float32
		987654.456789,  // float64
	)
	return fmt.Sprintf("%d", result)
}

func ReverseParamEnum() string {
	p1 := cross_call_master.Forth
	p2 := []cross_call_master.Example{cross_call_master.First, cross_call_master.Second, cross_call_master.Third}
	result := cross_call_master.ParamEnumCallback(p1, p2)
	return fmt.Sprintf("%d", result)
}

func ReverseParamEnumRef() string {
	p1 := cross_call_master.First
	p2 := []cross_call_master.Example{cross_call_master.First, cross_call_master.First, cross_call_master.Second}
	result := cross_call_master.ParamEnumRefCallback(&p1, &p2)
	return fmt.Sprintf("%d|%d|%s", result, p1, formatArray(p2))
}

func ReverseParamVariant() string {
	p1 := "my custom string with enough chars"
	p2 := []any{'X', '☢', -1, -1000, -1000000, -1000000000000, 200, 50000, 3000000000, 9999999999, 0xfedcbaabcdef, 0.001, 987654.456789}
	cross_call_master.ParamVariantCallback(p1, p2)
	return fmt.Sprintf("%s|%v", p1, p2)
}

func ReverseParamVariantRef() string {
	var p1 any = "my custom string with enough chars"
	p2 := []any{int8('X'), uint16('☢'), int32(-1), int32(-1000), int32(-1000000), int64(-1000000000000), int32(200), int32(50000), int64(3000000000), int64(9999999999), uintptr(0xfedcbaabcdef), float32(0.001), float32(987654.456789)}
	cross_call_master.ParamVariantRefCallback(&p1, &p2)
	return fmt.Sprintf("%s|{%v, %v, %s}", formatArray(p1.([]int32)), p2[0].(bool), formatFlt32(p2[1].(float32)), p2[2].(string))
}

func ReverseCallFuncVoid() string {
	cross_call_master.CallFuncVoidCallback(MockVoid)
	return ""
}

func ReverseCallFuncBool() string {
	result := cross_call_master.CallFuncBoolCallback(MockBool)
	return formatBool(result)
}

func ReverseCallFuncChar8() string {
	result := cross_call_master.CallFuncChar8Callback(MockChar8)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFuncChar16() string {
	result := cross_call_master.CallFuncChar16Callback(MockChar16)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFuncInt8() string {
	result := cross_call_master.CallFuncInt8Callback(MockInt8)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFuncInt16() string {
	result := cross_call_master.CallFuncInt16Callback(MockInt16)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFuncInt32() string {
	result := cross_call_master.CallFuncInt32Callback(MockInt32)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFuncInt64() string {
	result := cross_call_master.CallFuncInt64Callback(MockInt64)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFuncUInt8() string {
	result := cross_call_master.CallFuncUInt8Callback(MockUInt8)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFuncUInt16() string {
	result := cross_call_master.CallFuncUInt16Callback(MockUInt16)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFuncUInt32() string {
	result := cross_call_master.CallFuncUInt32Callback(MockUInt32)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFuncUInt64() string {
	result := cross_call_master.CallFuncUInt64Callback(MockUInt64)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFuncPtr() string {
	result := cross_call_master.CallFuncPtrCallback(MockPtr)
	return fmt.Sprintf("0x%x", result)
}

func ReverseCallFuncFloat() string {
	result := cross_call_master.CallFuncFloatCallback(MockFloat)
	return fmt.Sprintf("%v", formatFlt32(result))
}

func ReverseCallFuncDouble() string {
	result := cross_call_master.CallFuncDoubleCallback(MockDouble)
	return fmt.Sprintf("%v", formatFlt64(result))
}

func ReverseCallFuncString() string {
	result := cross_call_master.CallFuncStringCallback(MockString)
	return result
}

func ReverseCallFuncAny() string {
	result := cross_call_master.CallFuncAnyCallback(MockAny)
	return fmt.Sprintf("%c", result.(uint16))
}

func ReverseCallFuncBoolVector() string {
	result := cross_call_master.CallFuncBoolVectorCallback(MockBoolArray)
	return formatArrayFmt(result, formatBool)
}

func ReverseCallFuncChar8Vector() string {
	result := cross_call_master.CallFuncChar8VectorCallback(MockChar8Array)
	return formatArrayFmt(result, formatChr)
}

func ReverseCallFuncChar16Vector() string {
	result := cross_call_master.CallFuncChar16VectorCallback(MockChar16Array)
	return formatArray(result)
}

func ReverseCallFuncInt8Vector() string {
	result := cross_call_master.CallFuncInt8VectorCallback(MockInt8Array)
	return formatArray(result)
}

func ReverseCallFuncInt16Vector() string {
	result := cross_call_master.CallFuncInt16VectorCallback(MockInt16Array)
	return formatArray(result)
}

func ReverseCallFuncInt32Vector() string {
	result := cross_call_master.CallFuncInt32VectorCallback(MockInt32Array)
	return formatArray(result)
}

func ReverseCallFuncInt64Vector() string {
	result := cross_call_master.CallFuncInt64VectorCallback(MockInt64Array)
	return formatArray(result)
}

func ReverseCallFuncUInt8Vector() string {
	result := cross_call_master.CallFuncUInt8VectorCallback(MockUInt8Array)
	return formatArray(result)
}

func ReverseCallFuncUInt16Vector() string {
	result := cross_call_master.CallFuncUInt16VectorCallback(MockUInt16Array)
	return formatArray(result)
}

func ReverseCallFuncUInt32Vector() string {
	result := cross_call_master.CallFuncUInt32VectorCallback(MockUInt32Array)
	return formatArray(result)
}

func ReverseCallFuncUInt64Vector() string {
	result := cross_call_master.CallFuncUInt64VectorCallback(MockUInt64Array)
	return formatArray(result)
}

func ReverseCallFuncPtrVector() string {
	result := cross_call_master.CallFuncPtrVectorCallback(MockPtrArray)
	return formatArrayFmt(result, formatPtr)
}

func ReverseCallFuncFloatVector() string {
	result := cross_call_master.CallFuncFloatVectorCallback(MockFloatArray)
	return formatArrayFmt(result, formatFlt32)
}

func ReverseCallFuncDoubleVector() string {
	result := cross_call_master.CallFuncDoubleVectorCallback(MockDoubleArray)
	return formatArrayFmt(result, formatFlt64)
}

func ReverseCallFuncStringVector() string {
	result := cross_call_master.CallFuncStringVectorCallback(MockStringArray)
	return formatArrayFmt(result, formatStr)
}

func ReverseCallFuncAnyVector() string {
	result := cross_call_master.CallFuncAnyVectorCallback(MockAnyArray)
	return formatArrayFmt(result, formatAny)
}

func ReverseCallFuncVec2Vector() string {
	result := cross_call_master.CallFuncVec2VectorCallback(MockVec2Array)
	return formatArrayFmt(result, formatVector2)
}

func ReverseCallFuncVec3Vector() string {
	result := cross_call_master.CallFuncVec3VectorCallback(MockVec3Array)
	return formatArrayFmt(result, formatVector3)
}

func ReverseCallFuncVec4Vector() string {
	result := cross_call_master.CallFuncVec4VectorCallback(MockVec4Array)
	return formatArrayFmt(result, formatVector4)
}

func ReverseCallFuncMat4x4Vector() string {
	result := cross_call_master.CallFuncMat4x4VectorCallback(MockMat4x4Array)
	return formatArrayFmt(result, formatMatrix4x4)
}

func ReverseCallFuncVec2() string {
	result := cross_call_master.CallFuncVec2Callback(MockVec2)
	return formatVector2(result)
}

func ReverseCallFuncVec3() string {
	result := cross_call_master.CallFuncVec3Callback(MockVec3)
	return formatVector3(result)
}

func ReverseCallFuncVec4() string {
	result := cross_call_master.CallFuncVec4Callback(MockVec4)
	return formatVector4(result)
}

func ReverseCallFuncMat4x4() string {
	result := cross_call_master.CallFuncMat4x4Callback(MockMat4x4)
	return formatMatrix4x4(result)
}

func ReverseCallFunc1() string {
	result := cross_call_master.CallFunc1Callback(MockFunc1)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFunc2() string {
	result := cross_call_master.CallFunc2Callback(MockFunc2)
	return fmt.Sprintf("%c", result)
}

func ReverseCallFunc3() string {
	cross_call_master.CallFunc3Callback(MockFunc3)
	return ""
}

func ReverseCallFunc4() string {
	result := cross_call_master.CallFunc4Callback(MockFunc4)
	return formatVector4(result)
}

func ReverseCallFunc5() string {
	result := cross_call_master.CallFunc5Callback(MockFunc5)
	return fmt.Sprintf("%t", result)
}

func ReverseCallFunc6() string {
	result := cross_call_master.CallFunc6Callback(MockFunc6)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFunc7() string {
	result := cross_call_master.CallFunc7Callback(MockFunc7)
	return fmt.Sprintf("%v", formatFlt64(result))
}

func ReverseCallFunc8() string {
	result := cross_call_master.CallFunc8Callback(MockFunc8)
	return formatMatrix4x4(result)
}

func ReverseCallFunc9() string {
	cross_call_master.CallFunc9Callback(MockFunc9)
	return ""
}

func ReverseCallFunc10() string {
	result := cross_call_master.CallFunc10Callback(MockFunc10)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFunc11() string {
	result := cross_call_master.CallFunc11Callback(MockFunc11)
	return fmt.Sprintf("0x%x", result)
}

func ReverseCallFunc12() string {
	result := cross_call_master.CallFunc12Callback(MockFunc12)
	return fmt.Sprintf("%t", result)
}

func ReverseCallFunc13() string {
	result := cross_call_master.CallFunc13Callback(MockFunc13)
	return result
}

func ReverseCallFunc14() string {
	result := cross_call_master.CallFunc14Callback(MockFunc14)
	return formatArrayFmt(result, formatStr)
}

func ReverseCallFunc15() string {
	result := cross_call_master.CallFunc15Callback(MockFunc15)
	return fmt.Sprintf("%d", result)
}

func ReverseCallFunc16() string {
	result := cross_call_master.CallFunc16Callback(MockFunc16)
	return fmt.Sprintf("0x%x", result)
}

func ReverseCallFunc17() string {
	result := cross_call_master.CallFunc17Callback(MockFunc17)
	return result
}

func ReverseCallFunc18() string {
	result := cross_call_master.CallFunc18Callback(MockFunc18)
	return result
}

func ReverseCallFunc19() string {
	result := cross_call_master.CallFunc19Callback(MockFunc19)
	return result
}

func ReverseCallFunc20() string {
	result := cross_call_master.CallFunc20Callback(MockFunc20)
	return result
}

func ReverseCallFunc21() string {
	result := cross_call_master.CallFunc21Callback(MockFunc21)
	return result
}

func ReverseCallFunc22() string {
	result := cross_call_master.CallFunc22Callback(MockFunc22)
	return result
}

func ReverseCallFunc23() string {
	result := cross_call_master.CallFunc23Callback(MockFunc23)
	return result
}

func ReverseCallFunc24() string {
	result := cross_call_master.CallFunc24Callback(MockFunc24)
	return result
}

func ReverseCallFunc25() string {
	result := cross_call_master.CallFunc25Callback(MockFunc25)
	return result
}

func ReverseCallFunc26() string {
	result := cross_call_master.CallFunc26Callback(MockFunc26)
	return result
}

func ReverseCallFunc27() string {
	result := cross_call_master.CallFunc27Callback(MockFunc27)
	return result
}

func ReverseCallFunc28() string {
	result := cross_call_master.CallFunc28Callback(MockFunc28)
	return result
}

func ReverseCallFunc29() string {
	result := cross_call_master.CallFunc29Callback(MockFunc29)
	return result
}

func ReverseCallFunc30() string {
	result := cross_call_master.CallFunc30Callback(MockFunc30)
	return result
}

func ReverseCallFunc31() string {
	result := cross_call_master.CallFunc31Callback(MockFunc31)
	return result
}

func ReverseCallFunc32() string {
	result := cross_call_master.CallFunc32Callback(MockFunc32)
	return result
}

func ReverseCallFunc33() string {
	result := cross_call_master.CallFunc33Callback(MockFunc33)
	return result
}

func ReverseCallFuncEnum() string {
	result := cross_call_master.CallFuncEnumCallback(MockFuncEnum)
	return result
}

var ReverseTest = map[string]func() string{
	"NoParamReturnVoid":         ReverseNoParamReturnVoid,
	"NoParamReturnBool":         ReverseNoParamReturnBool,
	"NoParamReturnChar8":        ReverseNoParamReturnChar8,
	"NoParamReturnChar16":       ReverseNoParamReturnChar16,
	"NoParamReturnInt8":         ReverseNoParamReturnInt8,
	"NoParamReturnInt16":        ReverseNoParamReturnInt16,
	"NoParamReturnInt32":        ReverseNoParamReturnInt32,
	"NoParamReturnInt64":        ReverseNoParamReturnInt64,
	"NoParamReturnUInt8":        ReverseNoParamReturnUInt8,
	"NoParamReturnUInt16":       ReverseNoParamReturnUInt16,
	"NoParamReturnUInt32":       ReverseNoParamReturnUInt32,
	"NoParamReturnUInt64":       ReverseNoParamReturnUInt64,
	"NoParamReturnPointer":      ReverseNoParamReturnPointer,
	"NoParamReturnFloat":        ReverseNoParamReturnFloat,
	"NoParamReturnDouble":       ReverseNoParamReturnDouble,
	"NoParamReturnFunction":     ReverseNoParamReturnFunction,
	"NoParamReturnString":       ReverseNoParamReturnString,
	"NoParamReturnAny":          ReverseNoParamReturnAny,
	"NoParamReturnArrayBool":    ReverseNoParamReturnArrayBool,
	"NoParamReturnArrayChar8":   ReverseNoParamReturnArrayChar8,
	"NoParamReturnArrayChar16":  ReverseNoParamReturnArrayChar16,
	"NoParamReturnArrayInt8":    ReverseNoParamReturnArrayInt8,
	"NoParamReturnArrayInt16":   ReverseNoParamReturnArrayInt16,
	"NoParamReturnArrayInt32":   ReverseNoParamReturnArrayInt32,
	"NoParamReturnArrayInt64":   ReverseNoParamReturnArrayInt64,
	"NoParamReturnArrayUInt8":   ReverseNoParamReturnArrayUInt8,
	"NoParamReturnArrayUInt16":  ReverseNoParamReturnArrayUInt16,
	"NoParamReturnArrayUInt32":  ReverseNoParamReturnArrayUInt32,
	"NoParamReturnArrayUInt64":  ReverseNoParamReturnArrayUInt64,
	"NoParamReturnArrayPointer": ReverseNoParamReturnArrayPointer,
	"NoParamReturnArrayFloat":   ReverseNoParamReturnArrayFloat,
	"NoParamReturnArrayDouble":  ReverseNoParamReturnArrayDouble,
	"NoParamReturnArrayString":  ReverseNoParamReturnArrayString,
	"NoParamReturnArrayAny":     ReverseNoParamReturnArrayAny,
	"NoParamReturnVector2":      ReverseNoParamReturnVector2,
	"NoParamReturnVector3":      ReverseNoParamReturnVector3,
	"NoParamReturnVector4":      ReverseNoParamReturnVector4,
	"NoParamReturnMatrix4x4":    ReverseNoParamReturnMatrix4x4,
	"Param1":                    ReverseParam1,
	"Param2":                    ReverseParam2,
	"Param3":                    ReverseParam3,
	"Param4":                    ReverseParam4,
	"Param5":                    ReverseParam5,
	"Param6":                    ReverseParam6,
	"Param7":                    ReverseParam7,
	"Param8":                    ReverseParam8,
	"Param9":                    ReverseParam9,
	"Param10":                   ReverseParam10,
	"ParamRef1":                 ReverseParamRef1,
	"ParamRef2":                 ReverseParamRef2,
	"ParamRef3":                 ReverseParamRef3,
	"ParamRef4":                 ReverseParamRef4,
	"ParamRef5":                 ReverseParamRef5,
	"ParamRef6":                 ReverseParamRef6,
	"ParamRef7":                 ReverseParamRef7,
	"ParamRef8":                 ReverseParamRef8,
	"ParamRef9":                 ReverseParamRef9,
	"ParamRef10":                ReverseParamRef10,
	"ParamRefArrays":            ReverseParamRefVectors,
	"ParamAllPrimitives":        ReverseParamAllPrimitives,
	"ParamEnum":                 ReverseParamEnum,
	"ParamEnumRef":              ReverseParamEnumRef,
	"ParamVariant":              ReverseParamVariant,
	"ParamVariantRef":           ReverseParamVariantRef,
	"CallFuncVoid":              ReverseCallFuncVoid,
	"CallFuncBool":              ReverseCallFuncBool,
	"CallFuncChar8":             ReverseCallFuncChar8,
	"CallFuncChar16":            ReverseCallFuncChar16,
	"CallFuncInt8":              ReverseCallFuncInt8,
	"CallFuncInt16":             ReverseCallFuncInt16,
	"CallFuncInt32":             ReverseCallFuncInt32,
	"CallFuncInt64":             ReverseCallFuncInt64,
	"CallFuncUInt8":             ReverseCallFuncUInt8,
	"CallFuncUInt16":            ReverseCallFuncUInt16,
	"CallFuncUInt32":            ReverseCallFuncUInt32,
	"CallFuncUInt64":            ReverseCallFuncUInt64,
	"CallFuncPtr":               ReverseCallFuncPtr,
	"CallFuncFloat":             ReverseCallFuncFloat,
	"CallFuncDouble":            ReverseCallFuncDouble,
	"CallFuncString":            ReverseCallFuncString,
	"CallFuncAny":               ReverseCallFuncAny,
	"CallFuncBoolVector":        ReverseCallFuncBoolVector,
	"CallFuncChar8Vector":       ReverseCallFuncChar8Vector,
	"CallFuncChar16Vector":      ReverseCallFuncChar16Vector,
	"CallFuncInt8Vector":        ReverseCallFuncInt8Vector,
	"CallFuncInt16Vector":       ReverseCallFuncInt16Vector,
	"CallFuncInt32Vector":       ReverseCallFuncInt32Vector,
	"CallFuncInt64Vector":       ReverseCallFuncInt64Vector,
	"CallFuncUInt8Vector":       ReverseCallFuncUInt8Vector,
	"CallFuncUInt16Vector":      ReverseCallFuncUInt16Vector,
	"CallFuncUInt32Vector":      ReverseCallFuncUInt32Vector,
	"CallFuncUInt64Vector":      ReverseCallFuncUInt64Vector,
	"CallFuncPtrVector":         ReverseCallFuncPtrVector,
	"CallFuncFloatVector":       ReverseCallFuncFloatVector,
	"CallFuncDoubleVector":      ReverseCallFuncDoubleVector,
	"CallFuncStringVector":      ReverseCallFuncStringVector,
	"CallFuncAnyVector":         ReverseCallFuncAnyVector,
	"CallFuncVec2Vector":        ReverseCallFuncVec2Vector,
	"CallFuncVec3Vector":        ReverseCallFuncVec3Vector,
	"CallFuncVec4Vector":        ReverseCallFuncVec4Vector,
	"CallFuncMat4x4Vector":      ReverseCallFuncMat4x4Vector,
	"CallFuncVec2":              ReverseCallFuncVec2,
	"CallFuncVec3":              ReverseCallFuncVec3,
	"CallFuncVec4":              ReverseCallFuncVec4,
	"CallFuncMat4x4":            ReverseCallFuncMat4x4,
	"CallFunc1":                 ReverseCallFunc1,
	"CallFunc2":                 ReverseCallFunc2,
	"CallFunc3":                 ReverseCallFunc3,
	"CallFunc4":                 ReverseCallFunc4,
	"CallFunc5":                 ReverseCallFunc5,
	"CallFunc6":                 ReverseCallFunc6,
	"CallFunc7":                 ReverseCallFunc7,
	"CallFunc8":                 ReverseCallFunc8,
	"CallFunc9":                 ReverseCallFunc9,
	"CallFunc10":                ReverseCallFunc10,
	"CallFunc11":                ReverseCallFunc11,
	"CallFunc12":                ReverseCallFunc12,
	"CallFunc13":                ReverseCallFunc13,
	"CallFunc14":                ReverseCallFunc14,
	"CallFunc15":                ReverseCallFunc15,
	"CallFunc16":                ReverseCallFunc16,
	"CallFunc17":                ReverseCallFunc17,
	"CallFunc18":                ReverseCallFunc18,
	"CallFunc19":                ReverseCallFunc19,
	"CallFunc20":                ReverseCallFunc20,
	"CallFunc21":                ReverseCallFunc21,
	"CallFunc22":                ReverseCallFunc22,
	"CallFunc23":                ReverseCallFunc23,
	"CallFunc24":                ReverseCallFunc24,
	"CallFunc25":                ReverseCallFunc25,
	"CallFunc26":                ReverseCallFunc26,
	"CallFunc27":                ReverseCallFunc27,
	"CallFunc28":                ReverseCallFunc28,
	"CallFunc29":                ReverseCallFunc29,
	"CallFunc30":                ReverseCallFunc30,
	"CallFunc31":                ReverseCallFunc31,
	"CallFunc32":                ReverseCallFunc32,
	"CallFunc33":                ReverseCallFunc33,
	"CallFuncEnum":              ReverseCallFuncEnum,
}

func ReverseCall(test string) {
	// Retrieve the function from the map
	if method, exists := ReverseTest[test]; exists {
		// Call the function and get the result
		result := method()
		if result != "" {
			// Call the ReverseReturn function with the result
			cross_call_master.ReverseReturn(result)
		}
	} else {
		// Print an error message if the function is not found
		fmt.Printf("Method '%s' not found.\n", test)
	}
}
