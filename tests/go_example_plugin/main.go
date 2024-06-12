package main

/*
typedef struct { float x, y; } Vector2;
typedef struct { float x, y, z; } Vector3;
typedef struct { float x, y, z, w; } Vector4;
typedef struct { float m00, m01, m02, m03, m10, m11, m12, m13, m20, m21, m22, m23, m30, m31, m32, m33; } Matrix4x4;
*/
import "C"
import (
	"fmt"
	"math"
	"github.com/untrustedmodders/go-plugify"
)

func assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func init() {
	plugify.OnPluginStart(func() {
		fmt.Println("Go: OnPluginStart")
	})

	plugify.OnPluginEnd(func() {
		fmt.Println("Go: OnPluginEnd")
	})
}

//export NoParamReturnVoid
func NoParamReturnVoid() {
	fmt.Println("Go: NoParamReturnVoid")
}

//export NoParamReturnBool
func NoParamReturnBool() bool {
	fmt.Println("Go: NoParamReturnBool")
	return true
}

//export NoParamReturnChar8
func NoParamReturnChar8() int8 {
	fmt.Println("Go: NoParamReturnChar16")
	return math.MaxInt8
}

//export NoParamReturnChar16
func NoParamReturnChar16() uint16 {
	fmt.Println("Go: NoParamReturnChar16")
	return math.MaxUint16
}

//export NoParamReturnInt8
func NoParamReturnInt8() int8 {
	fmt.Println("Go: NoParamReturnInt8")
	return math.MaxInt8
}

//export NoParamReturnInt16
func NoParamReturnInt16() int16 {
	fmt.Println("Go: NoParamReturnInt16")
	return math.MaxInt16
}

//export NoParamReturnInt32
func NoParamReturnInt32() int32 {
	fmt.Println("Go: NoParamReturnInt32")
	return math.MaxInt32
}

//export NoParamReturnInt64
func NoParamReturnInt64() int64 {
	fmt.Println("Go: NoParamReturnInt64")
	return math.MaxInt64
}

//export NoParamReturnUInt8
func NoParamReturnUInt8() uint8 {
	fmt.Println("Go: NoParamReturnUInt8")
	return math.MaxUint8
}

//export NoParamReturnUInt16
func NoParamReturnUInt16() uint16 {
	fmt.Println("Go: NoParamReturnUInt16")
	return math.MaxUint16
}

//export NoParamReturnUInt32
func NoParamReturnUInt32() uint32 {
	fmt.Println("Go: NoParamReturnUInt32")
	return math.MaxUint32
}

//export NoParamReturnUInt64
func NoParamReturnUInt64() uint64 {
	fmt.Println("Go: NoParamReturnUInt64")
	return math.MaxUint64
}

//export NoParamReturnPtr64
func NoParamReturnPtr64() uintptr {
	fmt.Println("Go: NoParamReturnPtr64")
	return uintptr(1)
}

//export NoParamReturnFloat
func NoParamReturnFloat() float32 {
	fmt.Println("Go: NoParamReturnFloat")
	return math.MaxFloat32
}

//export NoParamReturnDouble
func NoParamReturnDouble() float64 {
	fmt.Println("Go: NoParamReturnDouble")
	return math.MaxFloat64
}

//export NoParamReturnFunction
func NoParamReturnFunction() uintptr {
	fmt.Println("Go: NoParamReturnFunction")
	return uintptr(0)
}

//export NoParamReturnString
func NoParamReturnString() string {
	fmt.Println("Go: NoParamReturnString")
	return "Hello World"
}

// Declare global variables
var globalBool []bool = []bool{true, false}
var globalChar8 []int8 = []int8{'a', 'b', 'c', 'd'}
var globalChar16 []uint16 = []uint16{'a', 'b', 'c', 'd'}
var globalInt8 []int8 = []int8{-3, -2, -1, 0, 1}
var globalInt16 []int16 = []int16{-4, -3, -2, -1, 0, 1}
var globalInt32 []int32 = []int32{-5, -4, -3, -2, -1, 0, 1}
var globalInt64 []int64 = []int64{-6, -5, -4, -3, -2, -1, 0, 1}
var globalUInt8 []uint8 = []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8}
var globalUInt16 []uint16 = []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var globalUInt32 []uint32 = []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var globalUInt64 []uint64 = []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
var globalPtr64 []uintptr = []uintptr{0, 1, 2, 3}
var globalFloat []float32 = []float32{-12.34, 0.0, 12.34}
var globalDouble []float64 = []float64{-12.345, 0.0, 12.345}
var globalString []string= []string{
	"1st string", "2nd string",
	"3rd element string (Should be big enough to avoid small string optimization)",
}

//export NoParamReturnArrayBool
func NoParamReturnArrayBool() []bool {
	fmt.Println("Go: NoParamReturnArrayBool")
	return globalBool
}

//export NoParamReturnArrayChar8
func NoParamReturnArrayChar8() []int8 {
	fmt.Println("Go: NoParamReturnArrayChar8")
	return globalChar8
}

//export NoParamReturnArrayChar16
func NoParamReturnArrayChar16() []uint16 {
	fmt.Println("Go: NoParamReturnArrayChar16")
	return globalChar16
}

//export NoParamReturnArrayInt8
func NoParamReturnArrayInt8() []int8 {
	fmt.Println("Go: NoParamReturnArrayInt8")
	return globalInt8
}

//export NoParamReturnArrayInt16
func NoParamReturnArrayInt16() []int16 {
	fmt.Println("Go: NoParamReturnArrayInt16")
	return globalInt16
}

//export NoParamReturnArrayInt32
func NoParamReturnArrayInt32() []int32 {
	fmt.Println("Go: NoParamReturnArrayInt32")
	return globalInt32
}

//export NoParamReturnArrayInt64
func NoParamReturnArrayInt64() []int64 {
	fmt.Println("Go: NoParamReturnArrayInt64")
	return globalInt64
}

//export NoParamReturnArrayUInt8
func NoParamReturnArrayUInt8() []uint8 {
	fmt.Println("Go: NoParamReturnArrayUInt8")
	return globalUInt8
}

//export NoParamReturnArrayUInt16
func NoParamReturnArrayUInt16() []uint16 {
	fmt.Println("Go: NoParamReturnArrayUInt16")
	return globalUInt16
}

//export NoParamReturnArrayUInt32
func NoParamReturnArrayUInt32() []uint32 {
	fmt.Println("Go: NoParamReturnArrayUInt32")
	return globalUInt32
}

//export NoParamReturnArrayUInt64
func NoParamReturnArrayUInt64() []uint64 {
	fmt.Println("Go: NoParamReturnArrayUInt64")
	return globalUInt64
}

//export NoParamReturnArrayPtr64
func NoParamReturnArrayPtr64() []uintptr {
	fmt.Println("Go: NoParamReturnArrayPtr64")
	return globalPtr64
}

//export NoParamReturnArrayFloat
func NoParamReturnArrayFloat() []float32 {
	fmt.Println("Go: NoParamReturnArrayFloat")
	return globalFloat
}

//export NoParamReturnArrayDouble
func NoParamReturnArrayDouble() []float64 {
	fmt.Println("Go: NoParamReturnArrayDouble")
	return globalDouble
}

//export NoParamReturnArrayString
func NoParamReturnArrayString() []string {
	fmt.Println("Go: NoParamReturnArrayString")
	return globalString
}

//export NoParamReturnVector2
func NoParamReturnVector2() C.Vector2 {
	fmt.Println("Go: NoParamReturnVector2")
	return C.Vector2{1, 2}
}
//export NoParamReturnVector3
func NoParamReturnVector3() C.Vector3 {
	fmt.Println("Go: NoParamReturnVector3")
	return C.Vector3{1, 2, 3}
}
//export NoParamReturnVector4
func NoParamReturnVector4() C.Vector4 {
	fmt.Println("Go: NoParamReturnVector4")
	return C.Vector4{1, 2, 3, 4}
}
//export NoParamReturnMatrix4x4
func NoParamReturnMatrix4x4() C.Matrix4x4 {
	fmt.Println("Go: NoParamReturnMatrix4x4")
	return C.Matrix4x4{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16};
}

//export Param1
func Param1(a int32) {
	fmt.Printf("Param1: a = %d\n", a)
}

//export Param2
func Param2(a int32, b float32) {
	fmt.Printf("Param2: a = %d, b = %f\n", a, b)
}

//export Param3
func Param3(a int32, b float32, c float64) {
	fmt.Printf("Param3: a = %d, b = %f, c = %f\n", a, b, c)
}

//export Param4
func Param4(a int32, b float32, c float64, d C.Vector4) {
	fmt.Printf("Param4: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f]\n", a, b, c, d.x, d.y, d.z, d.w)
}

//export Param5
func Param5(a int32, b float32, c float64, d C.Vector4, e []int64) {
	fmt.Printf("Param5: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Println("]")
}

//export Param6
func Param6(a int32, b float32, c float64, d C.Vector4, e []int64, f uint16) {
	fmt.Printf("Param6: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c\n", f)
}

//export Param7
func Param7(a int32, b float32, c float64, d C.Vector4, e []int64, f uint16, g string) {
	fmt.Printf("Param7: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s\n", f, g)
}

//export Param8
func Param8(a int32, b float32, c float64, d C.Vector4, e []int64, f uint16, g string, h float32) {
	fmt.Printf("Param8: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s, h = %f\n", f, g, h)
}

//export Param9
func Param9(a int32, b float32, c float64, d C.Vector4, e []int64, f uint16, g string, h float32, k int16) {
	fmt.Printf("Param9: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s, h = %f, k = %d\n", f, g, h, k)
}

//export Param10
func Param10(a int32, b float32, c float64, d C.Vector4, e []int64, f uint16, g string, h float32, k int16, l uintptr) {
	fmt.Printf("Param10: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s, h = %f, k = %d, l = %v\n", f, g, h, k, l)
}

//export ParamRef1
func ParamRef1(a *int32) {
	*a = 42
}

//export ParamRef2
func ParamRef2(a *int32, b *float32) {
	*a = 10
	*b = 3.14;
}

//export ParamRef3
func ParamRef3(a *int32, b *float32, c *float64) {
	*a = -20
	*b = 2.718
	*c = 3.14159
}

//export ParamRef4
func ParamRef4(a *int32, b *float32, c *float64, d *C.Vector4) {
	*a = 100
	*b = -5.55
	*c = 1.618
	*d = C.Vector4{1.0, 2.0, 3.0, 4.0}
}

//export ParamRef5
func ParamRef5(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64) {
	*a = 500
	*b = -10.5
	*c = 2.71828
	*d = C.Vector4{-1.0, -2.0, -3.0, -4.0}
	*e = []int64{-6, -5, -4, -3, -2, -1, 0, 1}
}

//export ParamRef6
func ParamRef6(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *uint16) {
	*a = 750
	*b = 20.0
	*c = 1.23456
	*d = C.Vector4{10.0, 20.0, 30.0, 40.0}
	*e = []int64{-6, -5, -4}
	*f = 'Z'
}

//export ParamRef7
func ParamRef7(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *uint16, g *string) {
	*a = -1000
	*b = 3.0
	*c = -1.0
	*d = C.Vector4{100.0, 200.0, 300.0, 400.0}
	*e = []int64{-6, -5, -4, -3}
	*f = uint16('X')
	*g = "Hello, World!"
}

//export ParamRef8
func ParamRef8(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *uint16, g *string, h *float32) {
	*a = 999
	*b = -7.5
	*c = 0.123456
	*d = C.Vector4{-100.0, -200.0, -300.0, -400.0}
	*e = []int64{-6, -5, -4, -3, -2, -1}
	*f = uint16('Y')
	*g = "Goodbye, World!"
	*h = 99.99
}

//export ParamRef9
func ParamRef9(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *uint16, g *string, h *float32, k *int16) {
	*a = -1234
	*b = 123.45
	*c = -678.9
	*d = C.Vector4{987.65, 432.1, 123.456, 789.123}
	*e = []int64{-6, -5, -4, -3, -2, -1, 0, 1, 5, 9}
	*f = uint16('A')
	*g = "Testing, 1 2 3"
	*h = -987.654
	*k = 42
}

//export ParamRef10
func ParamRef10(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *uint16, g *string, h *float32, k *int16, l *uintptr) {
	*a = 987
	*b = -0.123
	*c = 456.789
	*d = C.Vector4{-123.456, 0.987, 654.321, -789.123}
	*e = []int64{-6, -5, -4, -3, -2, -1, 0, 1, 5, 9, 4, -7}
	*f = uint16('B')
	*g = "Another string"
	*h = 3.141592
	*k = -32768
	*l = 0 // Null pointer
}

//export ParamRefVectors
func ParamRefVectors(p1 *[]bool, p2 *[]int8, p3 *[]uint16, p4 *[]int8, p5 *[]int16, p6 *[]int32, p7 *[]int64, p8 *[]uint8, p9 *[]uint16, p10 *[]uint32, p11 *[]uint64, p12 *[]uintptr, p13 *[]float32, p14 *[]float64, p15 *[]string) {
	*p1 = []bool{true}
	*p2 = []int8{'a', 'b', 'c'}
	*p3 = []uint16{'a', 'b', 'c'}
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

//export ParamAllPrimitives
func ParamAllPrimitives(p1 bool, p2 uint16, p3 int8, p4 int16, p5 int32, p6 int64, p7 uint8, p8 uint16, p9 uint32, p10 uint64, p11 uintptr, p12 float32, p13 float64) int64 {
	sum := int64(0)
	if p1 {
		sum += int64(1)
	}
	sum += int64(p2)
	sum += int64(p3)
	sum += int64(p4)
	sum += int64(p5)
	sum += p6
	sum += int64(p7)
	sum += int64(p8)
	sum += int64(p9)
	sum += int64(p10)
	sum += int64(p11)
	sum += int64(p12)
	sum += int64(p13)
	return sum
}

func main() {}
