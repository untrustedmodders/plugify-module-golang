package main

/*
typedef struct { float x, y; } Vector2;
typedef struct { float x, y, z; } Vector3;
typedef struct { float x, y, z, w; } Vector4;
typedef struct { float m[4][4]; } Matrix4x4;
static int DaysInYear(void* funcPtr) {
	return ((int (*)(void)) funcPtr)();
}
*/
import "C"
import (
	"fmt"
	"github.com/untrustedmodders/go-plugify"
	"math"
	"plugify-plugin/cross_call_master"
	"strconv"
	"unsafe"
)

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

//export NoParamReturnPointer
func NoParamReturnPointer() uintptr {
	fmt.Println("Go: NoParamReturnPointer")
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
var globalPointer []uintptr = []uintptr{0, 1, 2, 3}
var globalFloat []float32 = []float32{-12.34, 0.0, 12.34}
var globalDouble []float64 = []float64{-12.345, 0.0, 12.345}
var globalString []string = []string{
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

//export NoParamReturnArrayPointer
func NoParamReturnArrayPointer() []uintptr {
	fmt.Println("Go: NoParamReturnArrayPointer")
	return globalPointer
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
	ret := plugify.Matrix4x4{
		M: [4][4]float32{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
	}
	return *(*C.Matrix4x4)(unsafe.Pointer(&ret))
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
func Param6(a int32, b float32, c float64, d C.Vector4, e []int64, f int8) {
	fmt.Printf("Param6: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c\n", f)
}

//export Param7
func Param7(a int32, b float32, c float64, d C.Vector4, e []int64, f int8, g string) {
	fmt.Printf("Param7: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s\n", f, g)
}

//export Param8
func Param8(a int32, b float32, c float64, d C.Vector4, e []int64, f int8, g string, h uint16) {
	fmt.Printf("Param8: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s, h = %f\n", f, g, h)
}

//export Param9
func Param9(a int32, b float32, c float64, d C.Vector4, e []int64, f int8, g string, h uint16, k int16) {
	fmt.Printf("Param9: a = %d, b = %f, c = %f, d = [%f,%f,%f,%f], e.size() = %d, e = [", a, b, c, d.x, d.y, d.z, d.w, len(e))
	for _, elem := range e {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("], f = %c, g = %s, h = %f, k = %d\n", f, g, h, k)
}

//export Param10
func Param10(a int32, b float32, c float64, d C.Vector4, e []int64, f int8, g string, h uint16, k int16, l uintptr) {
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
	*b = 3.14
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
func ParamRef6(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *int8) {
	*a = 750
	*b = 20.0
	*c = 1.23456
	*d = C.Vector4{10.0, 20.0, 30.0, 40.0}
	*e = []int64{-6, -5, -4}
	*f = 'Z'
}

//export ParamRef7
func ParamRef7(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *int8, g *string) {
	*a = -1000
	*b = 3.0
	*c = -1.0
	*d = C.Vector4{100.0, 200.0, 300.0, 400.0}
	*e = []int64{-6, -5, -4, -3}
	*f = 'Y'
	*g = "Hello, World!"
}

//export ParamRef8
func ParamRef8(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *int8, g *string, h *uint16) {
	*a = 999
	*b = -7.5
	*c = 0.123456
	*d = C.Vector4{-100.0, -200.0, -300.0, -400.0}
	*e = []int64{-6, -5, -4, -3, -2, -1}
	*f = 'X'
	*g = "Goodbye, World!"
	*h = 'A'
}

//export ParamRef9
func ParamRef9(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16) {
	*a = -1234
	*b = 123.45
	*c = -678.9
	*d = C.Vector4{987.65, 432.1, 123.456, 789.123}
	*e = []int64{-6, -5, -4, -3, -2, -1, 0, 1, 5, 9}
	*f = 'W'
	*g = "Testing, 1 2 3"
	*h = 'B'
	*k = 42
}

//export ParamRef10
func ParamRef10(a *int32, b *float32, c *float64, d *C.Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16, l *uintptr) {
	*a = 987
	*b = -0.123
	*c = 456.789
	*d = C.Vector4{-123.456, 0.987, 654.321, -789.123}
	*e = []int64{-6, -5, -4, -3, -2, -1, 0, 1, 5, 9, 4, -7}
	*f = 'V'
	*g = "Another string"
	*h = 'C'
	*k = -444
	*l = 0x12345678
}

//export ParamRefVectors
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

//export ParamAllPrimitives
func ParamAllPrimitives(p1 bool, p2 uint16, p3 int8, p4 int16, p5 int32, p6 int64, p7 uint8, p8 uint16, p9 uint32, p10 uint64, p11 uintptr, p12 float32, p13 float64) int64 {
	return 56
}

func main() {}

// Reverse staff

// Define the function implementations
func ReverseNoParamReturnVoid() string {
	cross_call_master.NoParamReturnVoidCallback()
	return ""
}

func ReverseNoParamReturnBool() string {
	result := cross_call_master.NoParamReturnBoolCallback()
	if result {
		return "true"
	}
	return "false"
}

func ReverseNoParamReturnChar8() string {
	result := cross_call_master.NoParamReturnChar8Callback()
	return strconv.Itoa(int(result))
}

func ReverseNoParamReturnChar16() string {
	result := cross_call_master.NoParamReturnChar16Callback()
	return strconv.Itoa(int(result))
}

func ReverseNoParamReturnInt8() string {
	result := cross_call_master.NoParamReturnInt8Callback()
	return strconv.Itoa(int(result))
}

func ReverseNoParamReturnInt16() string {
	result := cross_call_master.NoParamReturnInt16Callback()
	return strconv.Itoa(int(result))
}

func ReverseNoParamReturnInt32() string {
	result := cross_call_master.NoParamReturnInt32Callback()
	return strconv.Itoa(int(result))
}

func ReverseNoParamReturnInt64() string {
	result := cross_call_master.NoParamReturnInt64Callback()
	return strconv.FormatInt(result, 10)
}

func ReverseNoParamReturnUInt8() string {
	result := cross_call_master.NoParamReturnUInt8Callback()
	return strconv.Itoa(int(result))
}

func ReverseNoParamReturnUInt16() string {
	result := cross_call_master.NoParamReturnUInt16Callback()
	return strconv.Itoa(int(result))
}

func ReverseNoParamReturnUInt32() string {
	result := cross_call_master.NoParamReturnUInt32Callback()
	return strconv.Itoa(int(result))
}

func ReverseNoParamReturnUInt64() string {
	result := cross_call_master.NoParamReturnUInt64Callback()
	return strconv.FormatUint(result, 10)
}

func ReverseNoParamReturnPointer() string {
	result := cross_call_master.NoParamReturnPointerCallback()
	return fmt.Sprintf("0x%x", result)
}

func ReverseNoParamReturnFloat() string {
	result := cross_call_master.NoParamReturnFloatCallback()
	return strconv.FormatFloat(float64(result), 'f', -1, 32)
}

func ReverseNoParamReturnDouble() string {
	result := cross_call_master.NoParamReturnDoubleCallback()
	return strconv.FormatFloat(result, 'f', -1, 64)
}

func ReverseNoParamReturnFunction() string {
	result := cross_call_master.NoParamReturnFunctionCallback()
	return strconv.Itoa(int(C.DaysInYear(unsafe.Pointer(result))))
}

func ReverseNoParamReturnString() string {
	result := cross_call_master.NoParamReturnStringCallback()
	return result
}

func ReverseNoParamReturnArrayBool() string {
	result := cross_call_master.NoParamReturnArrayBoolCallback()
	str := "{"
	for _, val := range result {
		if val {
			str += "true, "
		} else {
			str += "false, "
		}
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayChar8() string {
	result := cross_call_master.NoParamReturnArrayChar8Callback()
	str := "{"
	for _, val := range result {
		str += strconv.Itoa(int(val)) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayChar16() string {
	result := cross_call_master.NoParamReturnArrayChar16Callback()
	str := "{"
	for _, val := range result {
		str += strconv.Itoa(int(val)) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayInt8() string {
	result := cross_call_master.NoParamReturnArrayInt8Callback()
	str := "{"
	for _, val := range result {
		str += strconv.Itoa(int(val)) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayInt16() string {
	result := cross_call_master.NoParamReturnArrayInt16Callback()
	str := "{"
	for _, val := range result {
		str += strconv.Itoa(int(val)) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayInt32() string {
	result := cross_call_master.NoParamReturnArrayInt32Callback()
	str := "{"
	for _, val := range result {
		str += strconv.Itoa(int(val)) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayInt64() string {
	result := cross_call_master.NoParamReturnArrayInt64Callback()
	str := "{"
	for _, val := range result {
		str += strconv.FormatInt(val, 10) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayUInt8() string {
	result := cross_call_master.NoParamReturnArrayUInt8Callback()
	str := "{"
	for _, val := range result {
		str += strconv.Itoa(int(val)) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayUInt16() string {
	result := cross_call_master.NoParamReturnArrayUInt16Callback()
	str := "{"
	for _, val := range result {
		str += strconv.Itoa(int(val)) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayUInt32() string {
	result := cross_call_master.NoParamReturnArrayUInt32Callback()
	str := "{"
	for _, val := range result {
		str += strconv.Itoa(int(val)) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayUInt64() string {
	result := cross_call_master.NoParamReturnArrayUInt64Callback()
	str := "{"
	for _, val := range result {
		str += strconv.FormatUint(val, 10) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayPointer() string {
	result := cross_call_master.NoParamReturnArrayPointerCallback()
	str := "{"
	for _, val := range result {
		str += fmt.Sprintf("0x%x", val) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayFloat() string {
	result := cross_call_master.NoParamReturnArrayFloatCallback()
	str := "{"
	for _, val := range result {
		str += strconv.FormatFloat(float64(val), 'f', -1, 32) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayDouble() string {
	result := cross_call_master.NoParamReturnArrayDoubleCallback()
	str := "{"
	for _, val := range result {
		str += strconv.FormatFloat(val, 'f', -1, 64) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnArrayString() string {
	result := cross_call_master.NoParamReturnArrayStringCallback()
	str := "{"
	for _, val := range result {
		str += "'" + val + "', "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
}

func ReverseNoParamReturnVector2() string {
	vector := cross_call_master.NoParamReturnVector2Callback()
	return fmt.Sprintf("{%.1f, %.1f}", vector.X, vector.Y)
}

func ReverseNoParamReturnVector3() string {
	vector := cross_call_master.NoParamReturnVector3Callback()
	return fmt.Sprintf("{%.1f, %.1f, %.1f}", vector.X, vector.Y, vector.Z)
}

func ReverseNoParamReturnVector4() string {
	vector := cross_call_master.NoParamReturnVector4Callback()
	return fmt.Sprintf("{%.1f, %.1f, %.1f, %.1f}", vector.X, vector.Y, vector.Z, vector.W)
}

func ReverseNoParamReturnMatrix4x4() string {
	result := cross_call_master.NoParamReturnMatrix4x4Callback()
	str := "{"
	for _, row := range result.M {
		str += "{"
		for _, val := range row {
			str += fmt.Sprintf("%.1f, ", val)
		}
		str = str[:len(str)-2] + "}, "
	}
	if len(str) > 0 {
		str = str[:len(str)-2] + "}"
	}
	return str
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
	cross_call_master.Param4Callback(666, 7.7, 8.7659, cross_call_master.Vector4{100.1, 200.2, 300.3, 400.4})
	return ""
}

func ReverseParam5() string {
	cross_call_master.Param5Callback(555, 6.6, 7.6598, cross_call_master.Vector4{-105.1, -205.2, -305.3, -405.4}, []int64{})
	return ""
}

func ReverseParam6() string {
	cross_call_master.Param6Callback(444, 5.5, 6.5987, cross_call_master.Vector4{110.1, 210.2, 310.3, 410.4}, []int64{90000, -100, 20000}, 'A')
	return ""
}

func ReverseParam7() string {
	cross_call_master.Param7Callback(333, 4.4, 5.9876, cross_call_master.Vector4{-115.1, -215.2, -315.3, -415.4}, []int64{800000, 30000, -4000000}, 'B', "red gold")
	return ""
}

func ReverseParam8() string {
	cross_call_master.Param8Callback(222, 3.3, 1.2345, cross_call_master.Vector4{120.1, 220.2, 320.3, 420.4}, []int64{7000000, 5000000, -600000000}, 'C', "blue ice", 'Z')
	return ""
}

func ReverseParam9() string {
	cross_call_master.Param9Callback(111, 2.2, 5.1234, cross_call_master.Vector4{-125.1, -225.2, -325.3, -425.4}, []int64{60000000, -700000000, 80000000000}, 'D', "pink metal", 'Y', -100)
	return ""
}

func ReverseParam10() string {
	cross_call_master.Param10Callback(1234, 1.1, 4.5123, cross_call_master.Vector4{130.1, 230.2, 330.3, 430.4}, []int64{500000000, 90000000000, 1000000000000}, 'E', "green wood", 'X', -200, 0xabeba)
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
	d := cross_call_master.Vector4{}
	cross_call_master.ParamRef4Callback(&a, &b, &c, &d)
	return fmt.Sprintf("%d|%.1f|%.5f|{%.1f, %.1f, %.1f, %.1f}", a, b, c, d.X, d.Y, d.Z, d.W)
}

func ReverseParamRef5() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := cross_call_master.Vector4{}
	e := []int64{}
	cross_call_master.ParamRef5Callback(&a, &b, &c, &d, &e)
	eStr := ""
	for _, v := range e {
		eStr += strconv.FormatInt(v, 10) + ", "
	}
	if len(e) > 0 {
		eStr = eStr[:len(eStr)-2]
	}
	return fmt.Sprintf("%d|%.1f|%.5f|{%.1f, %.1f, %.1f, %.1f}|{%s}", a, b, c, d.X, d.Y, d.Z, d.W, eStr)
}

func ReverseParamRef6() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := cross_call_master.Vector4{}
	e := []int64{}
	f := int8(0)
	cross_call_master.ParamRef6Callback(&a, &b, &c, &d, &e, &f)
	eStr := ""
	for _, v := range e {
		eStr += strconv.FormatInt(v, 10) + ", "
	}
	if len(e) > 0 {
		eStr = eStr[:len(eStr)-2]
	}
	return fmt.Sprintf("%d|%.1f|%.5f|{%.1f, %.1f, %.1f, %.1f}|{%s}|%d", a, b, c, d.X, d.Y, d.Z, d.W, eStr, f)
}

func ReverseParamRef7() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := cross_call_master.Vector4{}
	e := []int64{}
	f := int8(0)
	g := ""
	cross_call_master.ParamRef7Callback(&a, &b, &c, &d, &e, &f, &g)
	eStr := ""
	for _, v := range e {
		eStr += strconv.FormatInt(v, 10) + ", "
	}
	if len(e) > 0 {
		eStr = eStr[:len(eStr)-2]
	}
	return fmt.Sprintf("%d|%.1f|%.5f|{%.1f, %.1f, %.1f, %.1f}|{%s}|%d|%s", a, b, c, d.X, d.Y, d.Z, d.W, eStr, f, g)
}

func ReverseParamRef8() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := cross_call_master.Vector4{}
	e := []int64{}
	f := int8(0)
	g := ""
	h := uint16(0)
	cross_call_master.ParamRef8Callback(&a, &b, &c, &d, &e, &f, &g, &h)
	eStr := ""
	for _, v := range e {
		eStr += strconv.FormatInt(v, 10) + ", "
	}
	if len(e) > 0 {
		eStr = eStr[:len(eStr)-2]
	}
	return fmt.Sprintf("%d|%.1f|%.5f|{%.1f, %.1f, %.1f, %.1f}|{%s}|%d|%s|%d", a, b, c, d.X, d.Y, d.Z, d.W, eStr, f, g, h)
}

func ReverseParamRef9() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := cross_call_master.Vector4{}
	e := []int64{}
	f := int8(0)
	g := ""
	h := uint16(0)
	k := int16(0)
	cross_call_master.ParamRef9Callback(&a, &b, &c, &d, &e, &f, &g, &h, &k)
	eStr := ""
	for _, v := range e {
		eStr += strconv.FormatInt(v, 10) + ", "
	}
	if len(e) > 0 {
		eStr = eStr[:len(eStr)-2]
	}
	return fmt.Sprintf("%d|%.1f|%.5f|{%.1f, %.1f, %.1f, %.1f}|{%s}|%d|%s|%d|%d", a, b, c, d.X, d.Y, d.Z, d.W, eStr, f, g, h, k)
}

func ReverseParamRef10() string {
	a, b, c := int32(0), float32(0), float64(0)
	d := cross_call_master.Vector4{}
	e := []int64{}
	f := int8(0)
	g := ""
	h := uint16(0)
	k := int16(0)
	l := uintptr(0)
	cross_call_master.ParamRef10Callback(&a, &b, &c, &d, &e, &f, &g, &h, &k, &l)
	eStr := ""
	for _, v := range e {
		eStr += strconv.FormatInt(v, 10) + ", "
	}
	if len(e) > 0 {
		eStr = eStr[:len(eStr)-2]
	}
	return fmt.Sprintf("%d|%.1f|%.5f|{%.1f, %.1f, %.1f, %.1f}|{%s}|%d|%s|%d|%d|0x%x", a, b, c, d.X, d.Y, d.Z, d.W, eStr, f, g, h, k, l)
}

// formatArray formats a slice of any type into a string
func formatArray[T any](arr []T) string {
	if len(arr) == 0 {
		return ""
	}

	str := ""
	for i, v := range arr {
		str += fmt.Sprintf("%v", v)
		if i < len(arr)-1 {
			str += ", "
		}
	}
	return str
}

func formatArrayPtr[T any](arr []T) string {
	if len(arr) == 0 {
		return ""
	}

	str := ""
	for i, v := range arr {
		str += fmt.Sprintf("0x%x", v)
		if i < len(arr)-1 {
			str += ", "
		}
	}
	return str
}

func formatArrayStr[T any](arr []T) string {
	if len(arr) == 0 {
		return ""
	}

	str := ""
	for i, v := range arr {
		str += fmt.Sprintf("'%v'", v)
		if i < len(arr)-1 {
			str += ", "
		}
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
	p2Formatted := formatArray(p2)
	p3Formatted := formatArray(p3)
	p4Formatted := formatArray(p4)
	p5Formatted := formatArray(p5)
	p6Formatted := formatArray(p6)
	p7Formatted := formatArray(p7)
	p8Formatted := formatArray(p8)
	p9Formatted := formatArray(p9)
	p10Formatted := formatArray(p10)
	p11Formatted := formatArray(p11)
	p12Formatted := formatArrayPtr(p12)
	p13Formatted := formatArray(p13)
	p14Formatted := formatArray(p14)
	p15Formatted := formatArrayStr(p15)

	return fmt.Sprintf("{%s}|{%s}|{%s}|{%s}|{%s}|{%s}|{%s}|{%s}|{%s}|{%s}|{%s}|{%s}|{%s}|{%s}|{%s}",
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
	return strconv.Itoa(int(result))
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
}

//export ReverseCall
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
