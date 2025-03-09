package main

import (
	"fmt"
	"github.com/untrustedmodders/go-plugify"
	"plugify-plugin/cross_call_master"
	"unsafe"
)

func MockVoid() {
	// fmt.Println("Void function called")
}

// Functions returning primitive types
func MockBool() bool {
	return true
}

func MockChar8() int8 {
	return 'A'
}

func MockChar16() uint16 {
	return 'Z'
}

func MockInt8() int8 {
	return 10
}

func MockInt16() int16 {
	return 100
}

func MockInt32() int32 {
	return 1000
}

func MockInt64() int64 {
	return 10000
}

func MockUInt8() uint8 {
	return 20
}

func MockUInt16() uint16 {
	return 200
}

func MockUInt32() uint32 {
	return 2000
}

func MockUInt64() uint64 {
	return 20000
}

func MockPtr() uintptr {
	return uintptr(0)
}

func MockFloat() float32 {
	return 3.14
}

func MockDouble() float64 {
	return 6.28
}

func MockFunction() uintptr {
	return uintptr(2)
}

func MockString() string {
	return "Test string"
}

func MockAny() any {
	return uint16('A')
}

// Functions returning arrays
func MockBoolArray() []bool {
	return []bool{true, false}
}

func MockChar8Array() []int8 {
	return []int8{'A', 'B'}
}

func MockChar16Array() []uint16 {
	return []uint16{'A', 'B'}
}

func MockInt8Array() []int8 {
	return []int8{10, 20}
}

func MockInt16Array() []int16 {
	return []int16{100, 200}
}

func MockInt32Array() []int32 {
	return []int32{1000, 2000}
}

func MockInt64Array() []int64 {
	return []int64{10000, 20000}
}

func MockUInt8Array() []uint8 {
	return []uint8{20, 30}
}

func MockUInt16Array() []uint16 {
	return []uint16{200, 300}
}

func MockUInt32Array() []uint32 {
	return []uint32{2000, 3000}
}

func MockUInt64Array() []uint64 {
	return []uint64{20000, 30000}
}

func MockPtrArray() []uintptr {
	return []uintptr{uintptr(0), uintptr(1)}
}

func MockFloatArray() []float32 {
	return []float32{1.1, 2.2}
}

func MockDoubleArray() []float64 {
	return []float64{3.3, 4.4}
}

func MockStringArray() []string {
	return []string{"Hello", "World"}
}

func MockAnyArray() []any {
	return []any{"Hello", 3.14, 6.28, 1, 0xDEADBEAF}
}

func MockVec2Array() []plugify.Vector2 {
	return []plugify.Vector2{
		{0.5, -1.2},
		{3.4, 7.8},
		{-6.7, 2.3},
		{8.9, -4.5},
		{0.0, 0.0},
	}
}

func MockVec3Array() []plugify.Vector3 {
	return []plugify.Vector3{
		{2.1, 3.2, 4.3},
		{-5.4, 6.5, -7.6},
		{8.7, 9.8, 0.1},
		{1.2, -3.3, 4.4},
		{-5.5, 6.6, -7.7},
	}
}

func MockVec4Array() []plugify.Vector4 {
	return []plugify.Vector4{
		{0.1, 1.2, 2.3, 3.4},
		{-4.5, 5.6, 6.7, -7.8},
		{8.9, -9.0, 10.1, -11.2},
		{12.3, 13.4, 14.5, 15.6},
		{-16.7, 17.8, 18.9, -19.0},
	}
}

func MockMat4x4Array() []plugify.Matrix4x4 {
	return []plugify.Matrix4x4{
		// Identity matrix
		{
			M: [4][4]float32{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
		},
		// Random matrix #1
		{
			M: [4][4]float32{
				{0.5, 1.0, 1.5, 2.0},
				{2.5, 3.0, 3.5, 4.0},
				{4.5, 5.0, 5.5, 6.0},
				{6.5, 7.0, 7.5, 8.0},
			}},
		{
			M: [4][4]float32{
				{-1.0, -2.0, -3.0, -4.0},
				{-5.0, -6.0, -7.0, -8.0},
				{-9.0, -10.0, -11.0, -12.0},
				{-13.0, -14.0, -15.0, -16.0},
			}},
		// Random matrix #3
		{
			M: [4][4]float32{
				{1.1, 2.2, 3.3, 4.4},
				{5.5, 6.6, 7.7, 8.8},
				{9.9, 10.0, 11.1, 12.2},
				{13.3, 14.4, 15.5, 16.6},
			}},
	}
}

// Functions returning plugify.Vectors and matrices
func MockVec2() plugify.Vector2 {
	return plugify.Vector2{1.0, 2.0}
}

func MockVec3() plugify.Vector3 {
	return plugify.Vector3{1.0, 2.0, 3.0}
}

func MockVec4() plugify.Vector4 {
	return plugify.Vector4{1.0, 2.0, 3.0, 4.0}
}

func MockMat4x4() plugify.Matrix4x4 {
	return plugify.Matrix4x4{
		M: [4][4]float32{
			{1.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0},
		},
	}
}

func MockFunc1(v plugify.Vector3) int32 {
	_ = fmt.Sprintf("%f%f%f", v.X, v.Y, v.Z)
	return int32(v.X + v.Y + v.Z)
}

// Mock implementation for 2 parameter function
func MockFunc2(a float32, b int64) int8 {
	_ = fmt.Sprintf("%f%d", a, b)

	return '&'
}

// Mock implementation for 3 parameter function
func MockFunc3(p uintptr, v plugify.Vector4, s string) {
	_ = fmt.Sprintf("%d%f%f%f%f%s", p, v.X, v.Y, v.Z, v.W, s)

}

func MockFunc4(flag bool, u int32, c uint16, m plugify.Matrix4x4) plugify.Vector4 {
	_ = fmt.Sprintf("%t%d%d%f", flag, u, c, m.M[0][0])
	return plugify.Vector4{1.0, 2.0, 3.0, 4.0}
}

// Mock implementation for 5 parameter function
func MockFunc5(i int8, v plugify.Vector2, p uintptr, d float64, vec []uint64) bool {
	_ = fmt.Sprintf("%d%f%f%d%f%d", i, v.X, v.Y, p, d, len(vec))
	return true
}

// Mock implementation for 6 parameter function
func MockFunc6(s string, f float32, vec []float32, i int16, uVec []uint8, p uintptr) int64 {
	_ = fmt.Sprintf("%s%f%d%d%d%d", s, f, len(vec), i, len(uVec), p)
	return int64(f + float32(i))
}

// Mock implementation for 7 parameter function
func MockFunc7(vec []int8, u uint16, c uint16, uVec []uint32, v plugify.Vector4, flag bool, l uint64) float64 {
	_ = fmt.Sprintf("%d%d%d%d%f%f%f%f", len(vec), u, c, len(uVec), v.X, v.Y, v.Z, v.W)
	return 3.14
}

// Mock implementation for 8 parameter function
func MockFunc8(v plugify.Vector3, uVec []uint32, i int16, flag bool, v4 plugify.Vector4, cVec []uint16, c uint16, a int32) plugify.Matrix4x4 {
	_ = fmt.Sprintf("%f%f%f%d%d%t%f%d%d", v.X, v.Y, v.Z, len(uVec), i, flag, v4.W, len(cVec), c)
	return plugify.Matrix4x4{}
}

// Mock implementation for 9 parameter function
func MockFunc9(f float32, v plugify.Vector2, iVec []int8, l uint64, flag bool, s string, v4 plugify.Vector4, i int16, p uintptr) {
	_ = fmt.Sprintf("%f%f%f%d%d%t%s%f%d%d", f, v.X, v.Y, len(iVec), l, flag, s, v4.W, i, p)
}

// Mock implementation for 10 parameter function
func MockFunc10(v4 plugify.Vector4, m plugify.Matrix4x4, uVec []uint32, l uint64, cVec []int8, a int32, flag bool, v plugify.Vector2, i int64, d float64) uint32 {
	_ = fmt.Sprintf("%f%f%f%f%f%d%d%d%d%d%t", v4.X, v4.Y, v4.Z, v4.W, m.M[1][1], len(uVec), l, len(cVec), a, flag)
	return 42
}

// Mock implementation for 11 parameter function
func MockFunc11(bVec []bool, c uint16, u uint8, d float64, v3 plugify.Vector3, iVec []int8, i int64, u16 uint16, f float32, v plugify.Vector2, u32 uint32) uintptr {
	_ = fmt.Sprintf("%d%d%d%f%f%d%d%d%f%f%d%d", len(bVec), c, u, d, v3.X, len(iVec), i, u16, f, v.X, u32)
	return uintptr(0)
}

// Mock implementation for 12 parameter function
func MockFunc12(p uintptr, dVec []float64, u uint32, d float64, flag bool, a int32, i int8, l uint64, f float32, pVec []uintptr, i64 int64, c int8) bool {
	_ = fmt.Sprintf("%d%d%d%f%t%d%d%d%f%d%d%d", p, len(dVec), u, d, flag, a, i, l, f, len(pVec), i64, c)
	return false
}

// Mock implementation for 13 parameter function
func MockFunc13(i64 int64, cVec []int8, u16 uint16, f float32, bVec []bool, v4 plugify.Vector4, s string, a int32, v3 plugify.Vector3, p uintptr, v2 plugify.Vector2, u8Vec []uint8, i16 int16) string {
	_ = fmt.Sprintf("%d%d%d%f%d%f%s%d%f%d%f%d%d", i64, len(cVec), u16, f, len(bVec), v4.Z, s, a, v3.X, p, v2.X, len(u8Vec), i16)
	return "Dummy String"
}

// Mock implementation for 14 parameter function
func MockFunc14(cVec []int8, uVec []uint32, m plugify.Matrix4x4, flag bool, c uint16, a int32, fVec []float32, u16 uint16, u8Vec []uint8, i8 int8, v3 plugify.Vector3, v4 plugify.Vector4, d float64, p uintptr) []string {
	_ = fmt.Sprintf("%d%d%f%t%d%d%d%d%d%d%f%f%f%d", len(cVec), len(uVec), m.M[2][2], flag, c, a, len(fVec), u16, len(u8Vec), i8, v3.X, v4.X, d, p)
	return []string{"String1", "String2"}
}

// Mock implementation for 15 parameter function
func MockFunc15(iVec []int16, m plugify.Matrix4x4, v4 plugify.Vector4, p uintptr, l uint64, uVec []uint32, flag bool, f float32, cVec []uint16, u uint8, a int32, v2 plugify.Vector2, u16 uint16, d float64, u8Vec []uint8) int16 {
	_ = fmt.Sprintf("%d%f%f%d%d%d%t%f%d%d%d%f%d%f%d", len(iVec), m.M[1][0], v4.X, p, l, len(uVec), flag, f, len(cVec), u, a, v2.X, u16, d, len(u8Vec))
	return 257
}

// Mock implementation for 16 parameter function
func MockFunc16(bVec []bool, i16 int16, iVec []int8, v4 plugify.Vector4, m plugify.Matrix4x4, v2 plugify.Vector2, uVec []uint64, cVec []int8, s string, i64 int64, u32Vec []uint32, v3 plugify.Vector3, f float32, d float64, i8 int8, u16 uint16) uintptr {
	_ = fmt.Sprintf("%d%d%d%f%f%f%f%f%d%d%d%s%d%d%f%f%d%d", len(bVec), i16, len(iVec), v4.X, v4.Y, v4.Z, v4.W, m.M[2][3], v2.X, len(uVec), len(cVec), s, i64, len(u32Vec), v3.X, f, d, i8, u16)
	return uintptr(0)
}

func MockFunc17(refVal *int32) {
	*refVal += 10
}

func MockFunc18(i8 *int8, i16 *int16) plugify.Vector2 {
	*i8 = 5
	*i16 = 10
	return plugify.Vector2{float32(*i8), float32(*i16)}
}

func MockFunc19(u32 *uint32, v3 *plugify.Vector3, uVec *[]uint32) {
	*u32 = 42
	*v3 = plugify.Vector3{1.0, 2.0, 3.0}
	*uVec = []uint32{1, 2, 3}
}

func MockFunc20(c *uint16, v4 *plugify.Vector4, uVec *[]uint64, ch *int8) int32 {
	*c = 't'
	*v4 = plugify.Vector4{1.0, 2.0, 3.0, 4.0}
	*uVec = []uint64{100, 200}
	*ch = 'F'
	return 0
}

func MockFunc21(m *plugify.Matrix4x4, iVec *[]int32, v2 *plugify.Vector2, flag *bool, d *float64) float32 {
	*flag = true
	*d = 3.14
	*v2 = plugify.Vector2{1.0, 2.0}
	*m = plugify.Matrix4x4{
		M: [4][4]float32{
			{1.3, 0.6, 0.8, 0.5},
			{0.7, 1.1, 0.2, 0.4},
			{0.9, 0.3, 1.2, 0.7},
			{0.2, 0.8, 0.5, 1.0},
		},
	}
	*iVec = []int32{1, 2, 3}
	return 0.0
}

func MockFunc22(p *uintptr, u32 *uint32, dVec *[]float64, i16 *int16, s *string, v4 *plugify.Vector4) uint64 {
	*p = uintptr(0)
	*u32 = 99
	*i16 = 123
	*s = "Hello"
	*v4 = plugify.Vector4{1.0, 2.0, 3.0, 4.0}
	*dVec = []float64{1.1, 2.2, 3.3}
	return 0
}

func MockFunc23(u64 *uint64, v2 *plugify.Vector2, iVec *[]int16, c *uint16, f *float32, i8 *int8, u8Vec *[]uint8) {
	*u64 = 50
	*f = 1.5
	*i8 = -1
	*v2 = plugify.Vector2{3.0, 4.0}
	*u8Vec = []uint8{1, 2, 3}
	*c = 'Ⅴ'
	*iVec = []int16{1, 2, 3, 4}
}

func MockFunc24(cVec *[]int8, i64 *int64, u8Vec *[]uint8, v4 *plugify.Vector4, u64 *uint64, pVec *[]uintptr, d *float64, vVec *[]uintptr) plugify.Matrix4x4 {
	*i64 = 64
	*d = 2.71
	*v4 = plugify.Vector4{1.0, 2.0, 3.0, 4.0}
	*cVec = []int8{'a', 'b', 'c'}
	*u8Vec = []uint8{5, 6, 7}
	*pVec = []uintptr{uintptr(0)}
	*vVec = []uintptr{uintptr(1), uintptr(1), uintptr(2), uintptr(2)}
	*u64 = 0xffffffff
	return plugify.Matrix4x4{}
}

func MockFunc25(i32 *int32, pVec *[]uintptr, flag *bool, u8 *uint8, s *string, v3 *plugify.Vector3, i64 *int64, v4 *plugify.Vector4, u16 *uint16) float64 {
	*flag = false
	*i32 = 100
	*u8 = 250
	*v3 = plugify.Vector3{1.0, 2.0, 3.0}
	*v4 = plugify.Vector4{4.0, 5.0, 6.0, 7.0}
	*s = "MockFunc25"
	*pVec = []uintptr{uintptr(0)}
	*i64 = 1337
	*u16 = 64222
	return 0.0
}

func MockFunc26(c *uint16, v2 *plugify.Vector2, m *plugify.Matrix4x4, fVec *[]float32, i16 *int16, u64 *uint64, u32 *uint32, u16Vec *[]uint16, p *uintptr, flag *bool) int8 {
	*c = 'Z'
	*flag = true
	*v2 = plugify.Vector2{2.0, 3.0}
	*m = plugify.Matrix4x4{
		M: [4][4]float32{
			{0.9, 0.2, 0.4, 0.8},
			{0.1, 1.0, 0.6, 0.3},
			{0.7, 0.5, 0.2, 0.9},
			{0.3, 0.4, 1.5, 0.1},
		},
	}
	*fVec = []float32{1.1, 2.2}
	*u64 = 64
	*u32 = 32
	*u16Vec = []uint16{100, 200}
	*i16 = 332
	*p = uintptr(unsafe.Pointer(uintptr(0xDEADBEAFDEADBEAF)))
	return 'A'
}

func MockFunc27(f *float32, v3 *plugify.Vector3, p *uintptr, v2 *plugify.Vector2, i16Vec *[]int16, m *plugify.Matrix4x4, flag *bool, v4 *plugify.Vector4, i8 *int8, i32 *int32, u8Vec *[]uint8) uint8 {
	*f = 1.0
	*v3 = plugify.Vector3{-1.0, -2.0, -3.0}
	*p = uintptr(unsafe.Pointer(uintptr(0xDEADBEAFDEADBEAF)))
	*v2 = plugify.Vector2{-111.0, 111.0}
	*i16Vec = []int16{1, 2, 3, 4}
	*m = plugify.Matrix4x4{
		M: [4][4]float32{
			{1.0, 0.5, 0.3, 0.7},
			{0.8, 1.2, 0.6, 0.9},
			{1.5, 1.1, 0.4, 0.2},
			{0.3, 0.9, 0.7, 1.0},
		},
	}
	*flag = true
	*v4 = plugify.Vector4{1.0, 2.0, 3.0, 4.0}
	*i8 = 111
	*i32 = 30
	*u8Vec = []uint8{0, 0, 0, 0, 0, 0, 1, 0}
	return 0
}

func MockFunc28(ptr *uintptr, u16 *uint16, u32Vec *[]uint32, m *plugify.Matrix4x4, f *float32, v4 *plugify.Vector4, str *string, u64Vec *[]uint64, i64 *int64, b *bool, vec3 *plugify.Vector3, fVec *[]float32) string {
	*ptr = uintptr(0)
	*u16 = 65500
	*u32Vec = []uint32{1, 2, 3, 4, 5, 7}
	*m = plugify.Matrix4x4{
		M: [4][4]float32{
			{1.4, 0.7, 0.2, 0.5},
			{0.3, 1.1, 0.6, 0.8},
			{0.9, 0.4, 1.3, 0.1},
			{0.6, 0.2, 0.7, 1.0},
		},
	}
	*f = 5.5
	*v4 = plugify.Vector4{1.0, 2.0, 3.0, 4.0}
	*u64Vec = []uint64{1, 2}
	*i64 = 834748377834
	*b = true
	*vec3 = plugify.Vector3{10, 20, 30}
	*str = "MockFunc28"
	*fVec = []float32{1.0, -1000.0, 2000.0}
	return *str
}

func MockFunc29(v4 *plugify.Vector4, i32 *int32, iVec *[]int8, d *float64, flag *bool, i8 *int8, u16Vec *[]uint16, f *float32, s *string, m *plugify.Matrix4x4, u64 *uint64, v3 *plugify.Vector3, i64Vec *[]int64) []string {
	*i32 = 30
	*flag = true
	*v4 = plugify.Vector4{1.0, 2.0, 3.0, 4.0}
	*d = 3.14
	*i8 = 8
	*u16Vec = []uint16{100, 200}
	*f = 1.5
	*s = "MockFunc29"
	*m = plugify.Matrix4x4{
		M: [4][4]float32{
			{0.4, 1.0, 0.6, 0.3},
			{1.2, 0.8, 0.5, 0.9},
			{0.7, 0.3, 1.4, 0.6},
			{0.1, 0.9, 0.8, 1.3},
		},
	}
	*u64 = 64
	*v3 = plugify.Vector3{1.0, 2.0, 3.0}
	*i64Vec = []int64{1, 2, 3}
	*iVec = []int8{127, 126, 125}
	return []string{"Example", "MockFunc29"}
}

func MockFunc30(p *uintptr, v4 *plugify.Vector4, i64 *int64, uVec *[]uint32, flag *bool, s *string, v3 *plugify.Vector3, u8Vec *[]uint8, f *float32, v2 *plugify.Vector2, m *plugify.Matrix4x4, i8 *int8, vVec *[]float32, d *float64) int32 {
	*flag = false
	*f = 1.1
	*i64 = 1000
	*v2 = plugify.Vector2{3.0, 4.0}
	*v4 = plugify.Vector4{1.0, 2.0, 3.0, 4.0}
	*s = "MockFunc30"
	*p = uintptr(0)
	*uVec = []uint32{100, 200}
	*m = plugify.Matrix4x4{
		M: [4][4]float32{
			{0.5, 0.3, 1.0, 0.7},
			{1.1, 0.9, 0.6, 0.4},
			{0.2, 0.8, 1.5, 0.3},
			{0.7, 0.4, 0.9, 1.0},
		},
	}
	*i8 = 8
	*vVec = []float32{1.0, 1.0, 2.0, 2.0}
	*d = 2.718
	*v3 = plugify.Vector3{1, 2, 3}
	*u8Vec = []uint8{255, 0, 255, 200, 100, 200}
	return 42
}

func MockFunc31(c *int8, u32 *uint32, uVec *[]uint64, v4 *plugify.Vector4, s *string, flag *bool, i64 *int64, v2 *plugify.Vector2, i8 *int8, u16 *uint16, iVec *[]int16, m *plugify.Matrix4x4, v3 *plugify.Vector3, f *float32, v4Vec *[]float64) plugify.Vector3 {
	*u32 = 12345
	*flag = true
	*v3 = plugify.Vector3{1.0, 2.0, 3.0}
	*s = "MockFunc31"
	*v2 = plugify.Vector2{5.0, 6.0}
	*i8 = 7
	*u16 = 255
	*m = plugify.Matrix4x4{
		M: [4][4]float32{
			{0.8, 0.5, 1.2, 0.3},
			{1.0, 0.7, 0.4, 0.6},
			{0.9, 0.2, 0.5, 1.4},
			{0.6, 0.8, 1.1, 0.7},
		},
	}
	*iVec = []int16{1, 2}
	*v4 = plugify.Vector4{1.0, 2.0, 3.0, 4.0}
	*i64 = 123456789
	*c = 'C'
	*v4Vec = []float64{1.0, 1.0, 1.0, 1.0, 2.0, 2.0, 2.0, 2.0}
	*uVec = []uint64{1, 2, 3, 4, 5}
	*f = -1.0
	return plugify.Vector3{1.0, 2.0, 3.0}
}

func MockFunc32(i32 *int32, u16 *uint16, iVec *[]int8, v4 *plugify.Vector4, p *uintptr, uVec *[]uint32, m *plugify.Matrix4x4, u64 *uint64, s *string, i64 *int64, v2 *plugify.Vector2, u8Vec *[]int8, flag *bool, v3 *plugify.Vector3, u8 *uint8, cVec *[]uint16) float64 {
	*i32 = 42
	*u16 = 255
	*flag = false
	*v2 = plugify.Vector2{2.5, 3.5}
	*u8Vec = []int8{1, 2, 3, 4, 5, 9}
	*v4 = plugify.Vector4{4.0, 5.0, 6.0, 7.0}
	*s = "MockFunc32"
	*p = uintptr(0)
	*m = plugify.Matrix4x4{
		M: [4][4]float32{
			{1.0, 0.4, 0.3, 0.9},
			{0.7, 1.2, 0.5, 0.8},
			{0.2, 0.6, 1.1, 0.4},
			{0.9, 0.3, 0.8, 1.5},
		},
	}
	*u64 = 123456789
	*uVec = []uint32{100, 200}
	*i64 = 1000
	*v3 = plugify.Vector3{0.0, 0.0, 0.0}
	*u8 = 8
	*cVec = []uint16{'a', 'b', 'c'}
	*iVec = []int8{0, 1}
	return 1.0
}

func MockFunc33(variant *any) {
	*variant = "MockFunc33"
}

func MockFuncEnum(p1 cross_call_master.Example, p2 *[]cross_call_master.Example) []cross_call_master.Example {
	*p2 = []cross_call_master.Example{cross_call_master.First, cross_call_master.Second, cross_call_master.Third}
	return []cross_call_master.Example{p1, cross_call_master.Forth}
}
