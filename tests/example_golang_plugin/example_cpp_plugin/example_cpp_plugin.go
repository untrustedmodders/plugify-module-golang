package example_cpp_plugin

//generated with https://github.com/untrustedmodders/cpp-lang-module/blob/main/generator/generator.py from example_cpp_plugin 

// #include "example_cpp_plugin.h"
import "C"
import "unsafe"

type Vector2 struct {
	X float32
	Y float32
}
type Vector3 struct {
	X float32
	Y float32
	Z float32
}
type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}
type Matrix4x4 struct {
	M00 float32
	M10 float32
	M20 float32
	M30 float32

	M01 float32
	M11 float32
	M21 float32
	M31 float32

	M02 float32
	M12 float32
	M22 float32
	M32 float32

	M03 float32
	M13 float32
	M23 float32
	M33 float32
}

func NoParamReturnVoid()  {
	C.NoParamReturnVoid()
}

func NoParamReturnBool() bool {
	result := bool(C.NoParamReturnBool())
	return result
}

func NoParamReturnChar8() int8 {
	result := int8(C.NoParamReturnChar8())
	return result
}

func NoParamReturnChar16() uint16 {
	result := uint16(C.NoParamReturnChar16())
	return result
}

func NoParamReturnInt8() int8 {
	result := int8(C.NoParamReturnInt8())
	return result
}

func NoParamReturnInt16() int16 {
	result := int16(C.NoParamReturnInt16())
	return result
}

func NoParamReturnInt32() int32 {
	result := int32(C.NoParamReturnInt32())
	return result
}

func NoParamReturnInt64() int64 {
	result := int64(C.NoParamReturnInt64())
	return result
}

func NoParamReturnUInt8() uint8 {
	result := uint8(C.NoParamReturnUInt8())
	return result
}

func NoParamReturnUInt16() uint16 {
	result := uint16(C.NoParamReturnUInt16())
	return result
}

func NoParamReturnUInt32() uint32 {
	result := uint32(C.NoParamReturnUInt32())
	return result
}

func NoParamReturnUInt64() uint64 {
	result := uint64(C.NoParamReturnUInt64())
	return result
}

func NoParamReturnPtr64() uintptr {
	result := uintptr(C.NoParamReturnPtr64())
	return result
}

func NoParamReturnFloat() float32 {
	result := float32(C.NoParamReturnFloat())
	return result
}

func NoParamReturnDouble() float64 {
	result := float64(C.NoParamReturnDouble())
	return result
}

func NoParamReturnFunction() uintptr {
	result := uintptr(C.NoParamReturnFunction())
	return result
}

func NoParamReturnString() string {
	C_output := C.Plugify_AllocateString()
	C.NoParamReturnString(C_output)

	P_output := C.Plugify_GetStringData(C_output)
	output := C.GoString(P_output)
	C.Plugify_FreeString(C_output)
	return output
}

func NoParamReturnArrayBool() []bool {
	C_output := C.Plugify_AllocateVector(C.BOOL)
	C.NoParamReturnArrayBool(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.BOOL)
	P_output := C.Plugify_GetVectorData(C_output, C.BOOL)
	output := make([]bool, L_output)
	for i := range output {
		output[i] = *(*bool)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_bool)))
	}

	C.Plugify_DeleteVectorDataBool(P_output)
	C.Plugify_FreeVector(C_output, C.BOOL)
	return output
}

func NoParamReturnArrayChar8() []int8 {
	C_output := C.Plugify_AllocateVector(C.CHAR8)
	C.NoParamReturnArrayChar8(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.CHAR8)
	P_output := C.Plugify_GetVectorData(C_output, C.CHAR8)
	output := make([]int8, L_output)
	for i := range output {
		output[i] = *(*int8)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_char)))
	}

	C.Plugify_FreeVector(C_output, C.CHAR8)
	return output
}

func NoParamReturnArrayChar16() []uint16 {
	C_output := C.Plugify_AllocateVector(C.CHAR16)
	C.NoParamReturnArrayChar16(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.CHAR16)
	P_output := C.Plugify_GetVectorData(C_output, C.CHAR16)
	output := make([]uint16, L_output)
	for i := range output {
		output[i] = *(*uint16)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uint16_t)))
	}

	C.Plugify_FreeVector(C_output, C.CHAR16)
	return output
}

func NoParamReturnArrayInt8() []int8 {
	C_output := C.Plugify_AllocateVector(C.INT8)
	C.NoParamReturnArrayInt8(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.INT8)
	P_output := C.Plugify_GetVectorData(C_output, C.INT8)
	output := make([]int8, L_output)
	for i := range output {
		output[i] = *(*int8)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_int8_t)))
	}

	C.Plugify_FreeVector(C_output, C.INT8)
	return output
}

func NoParamReturnArrayInt16() []int16 {
	C_output := C.Plugify_AllocateVector(C.INT16)
	C.NoParamReturnArrayInt16(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.INT16)
	P_output := C.Plugify_GetVectorData(C_output, C.INT16)
	output := make([]int16, L_output)
	for i := range output {
		output[i] = *(*int16)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_int16_t)))
	}

	C.Plugify_FreeVector(C_output, C.INT16)
	return output
}

func NoParamReturnArrayInt32() []int32 {
	C_output := C.Plugify_AllocateVector(C.INT32)
	C.NoParamReturnArrayInt32(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.INT32)
	P_output := C.Plugify_GetVectorData(C_output, C.INT32)
	output := make([]int32, L_output)
	for i := range output {
		output[i] = *(*int32)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_int32_t)))
	}

	C.Plugify_FreeVector(C_output, C.INT32)
	return output
}

func NoParamReturnArrayInt64() []int64 {
	C_output := C.Plugify_AllocateVector(C.INT64)
	C.NoParamReturnArrayInt64(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.INT64)
	P_output := C.Plugify_GetVectorData(C_output, C.INT64)
	output := make([]int64, L_output)
	for i := range output {
		output[i] = *(*int64)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_int64_t)))
	}

	C.Plugify_FreeVector(C_output, C.INT64)
	return output
}

func NoParamReturnArrayUInt8() []uint8 {
	C_output := C.Plugify_AllocateVector(C.UINT8)
	C.NoParamReturnArrayUInt8(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.UINT8)
	P_output := C.Plugify_GetVectorData(C_output, C.UINT8)
	output := make([]uint8, L_output)
	for i := range output {
		output[i] = *(*uint8)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uint8_t)))
	}

	C.Plugify_FreeVector(C_output, C.UINT8)
	return output
}

func NoParamReturnArrayUInt16() []uint16 {
	C_output := C.Plugify_AllocateVector(C.UINT16)
	C.NoParamReturnArrayUInt16(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.UINT16)
	P_output := C.Plugify_GetVectorData(C_output, C.UINT16)
	output := make([]uint16, L_output)
	for i := range output {
		output[i] = *(*uint16)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uint16_t)))
	}

	C.Plugify_FreeVector(C_output, C.UINT16)
	return output
}

func NoParamReturnArrayUInt32() []uint32 {
	C_output := C.Plugify_AllocateVector(C.UINT32)
	C.NoParamReturnArrayUInt32(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.UINT32)
	P_output := C.Plugify_GetVectorData(C_output, C.UINT32)
	output := make([]uint32, L_output)
	for i := range output {
		output[i] = *(*uint32)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uint32_t)))
	}

	C.Plugify_FreeVector(C_output, C.UINT32)
	return output
}

func NoParamReturnArrayUInt64() []uint64 {
	C_output := C.Plugify_AllocateVector(C.UINT64)
	C.NoParamReturnArrayUInt64(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.UINT64)
	P_output := C.Plugify_GetVectorData(C_output, C.UINT64)
	output := make([]uint64, L_output)
	for i := range output {
		output[i] = *(*uint64)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uint64_t)))
	}

	C.Plugify_FreeVector(C_output, C.UINT64)
	return output
}

func NoParamReturnArrayPtr64() []uintptr {
	C_output := C.Plugify_AllocateVector(C.UINTPTR)
	C.NoParamReturnArrayPtr64(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.UINTPTR)
	P_output := C.Plugify_GetVectorData(C_output, C.UINTPTR)
	output := make([]uintptr, L_output)
	for i := range output {
		output[i] = *(*uintptr)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uintptr_t)))
	}

	C.Plugify_FreeVector(C_output, C.UINTPTR)
	return output
}

func NoParamReturnArrayFloat() []float32 {
	C_output := C.Plugify_AllocateVector(C.FLOAT)
	C.NoParamReturnArrayFloat(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.FLOAT)
	P_output := C.Plugify_GetVectorData(C_output, C.FLOAT)
	output := make([]float32, L_output)
	for i := range output {
		output[i] = *(*float32)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_float)))
	}

	C.Plugify_FreeVector(C_output, C.FLOAT)
	return output
}

func NoParamReturnArrayDouble() []float64 {
	C_output := C.Plugify_AllocateVector(C.DOUBLE)
	C.NoParamReturnArrayDouble(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.DOUBLE)
	P_output := C.Plugify_GetVectorData(C_output, C.DOUBLE)
	output := make([]float64, L_output)
	for i := range output {
		output[i] = *(*float64)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_double)))
	}

	C.Plugify_FreeVector(C_output, C.DOUBLE)
	return output
}

func NoParamReturnArrayString() []string {
	C_output := C.Plugify_AllocateVector(C.STRING)
	C.NoParamReturnArrayString(C_output)

	L_output := C.Plugify_GetVectorSize(C_output, C.STRING)
	P_output := C.Plugify_GetVectorData(C_output, C.STRING)
	output := make([]string, L_output)
	for i := range output {
		output[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uintptr_t))))
	}

	C.Plugify_DeleteVectorDataCStr(P_output)
	C.Plugify_FreeVector(C_output, C.STRING)
	return output
}

func NoParamReturnVector2() Vector2 {
	C_result := C.NoParamReturnVector2()
	return *(*Vector2)(unsafe.Pointer(&C_result))
}

func NoParamReturnVector3() Vector3 {
	C_result := C.NoParamReturnVector3()
	return *(*Vector3)(unsafe.Pointer(&C_result))
}

func NoParamReturnVector4() Vector4 {
	C_result := C.NoParamReturnVector4()
	return *(*Vector4)(unsafe.Pointer(&C_result))
}

func NoParamReturnMatrix4x4() Matrix4x4 {
	C_result := C.NoParamReturnMatrix4x4()
	return *(*Matrix4x4)(unsafe.Pointer(&C_result))
}

func Param1(a int32)  {
	C_a := C.int32_t(a)

	C.Param1(C_a)
}

func Param2(a int32, b float32)  {
	C_a := C.int32_t(a)
	C_b := C.float(b)

	C.Param2(C_a, C_b)
}

func Param3(a int32, b float32, c float64)  {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)

	C.Param3(C_a, C_b, C_c)
}

func Param4(a int32, b float32, c float64, d Vector4)  {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))

	C.Param4(C_a, C_b, C_c, &C_d)
}

func Param5(a int32, b float32, c float64, d Vector4, e []int64)  {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))
	C_e := C.Plugify_CreateVector(unsafe.Pointer(&e[0]), C.ptrdiff_t(len(e)), C.INT64)

	C.Param5(C_a, C_b, C_c, &C_d, C_e)

	C.Plugify_DeleteVector(C_e, C.INT64)

}

func Param6(a int32, b float32, c float64, d Vector4, e []int64, f int8)  {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))
	C_e := C.Plugify_CreateVector(unsafe.Pointer(&e[0]), C.ptrdiff_t(len(e)), C.INT64)
	C_f := C.char(f)

	C.Param6(C_a, C_b, C_c, &C_d, C_e, C_f)

	C.Plugify_DeleteVector(C_e, C.INT64)

}

func Param7(a int32, b float32, c float64, d Vector4, e []int64, f int8, g string)  {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))
	C_e := C.Plugify_CreateVector(unsafe.Pointer(&e[0]), C.ptrdiff_t(len(e)), C.INT64)
	C_f := C.char(f)
	C_g := C.Plugify_CreateString(g)

	C.Param7(C_a, C_b, C_c, &C_d, C_e, C_f, C_g)

	C.Plugify_DeleteVector(C_e, C.INT64)
	C.Plugify_DeleteString(C_g)

}

func Param8(a int32, b float32, c float64, d Vector4, e []int64, f int8, g string, h float32)  {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))
	C_e := C.Plugify_CreateVector(unsafe.Pointer(&e[0]), C.ptrdiff_t(len(e)), C.INT64)
	C_f := C.char(f)
	C_g := C.Plugify_CreateString(g)
	C_h := C.float(h)

	C.Param8(C_a, C_b, C_c, &C_d, C_e, C_f, C_g, C_h)

	C.Plugify_DeleteVector(C_e, C.INT64)
	C.Plugify_DeleteString(C_g)

}

func Param9(a int32, b float32, c float64, d Vector4, e []int64, f int8, g string, h float32, k int16)  {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))
	C_e := C.Plugify_CreateVector(unsafe.Pointer(&e[0]), C.ptrdiff_t(len(e)), C.INT64)
	C_f := C.char(f)
	C_g := C.Plugify_CreateString(g)
	C_h := C.float(h)
	C_k := C.int16_t(k)

	C.Param9(C_a, C_b, C_c, &C_d, C_e, C_f, C_g, C_h, C_k)

	C.Plugify_DeleteVector(C_e, C.INT64)
	C.Plugify_DeleteString(C_g)

}

func Param10(a int32, b float32, c float64, d Vector4, e []int64, f int8, g string, h float32, k int16, l uintptr)  {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))
	C_e := C.Plugify_CreateVector(unsafe.Pointer(&e[0]), C.ptrdiff_t(len(e)), C.INT64)
	C_f := C.char(f)
	C_g := C.Plugify_CreateString(g)
	C_h := C.float(h)
	C_k := C.int16_t(k)
	C_l := C.uintptr_t(l)

	C.Param10(C_a, C_b, C_c, &C_d, C_e, C_f, C_g, C_h, C_k, C_l)

	C.Plugify_DeleteVector(C_e, C.INT64)
	C.Plugify_DeleteString(C_g)

}

func ParamRef1(a *int32)  {
	C_a := (*C.int32_t)(a)

	C.ParamRef1(C_a)
}

func ParamRef2(a *int32, b *float32)  {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)

	C.ParamRef2(C_a, C_b)
}

func ParamRef3(a *int32, b *float32, c *float64)  {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)

	C.ParamRef3(C_a, C_b, C_c)
}

func ParamRef4(a *int32, b *float32, c *float64, d *Vector4)  {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))

	C.ParamRef4(C_a, C_b, C_c, &C_d)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
}

func ParamRef5(a *int32, b *float32, c *float64, d *Vector4, e *[]int64)  {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))
	C_e := C.Plugify_CreateVector(unsafe.Pointer(&(*e)[0]), C.ptrdiff_t(len(*e)), C.INT64)

	C.ParamRef5(C_a, C_b, C_c, &C_d, C_e)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
	L_e := C.Plugify_GetVectorSize(C_e, C.INT64)
	P_e := C.Plugify_GetVectorData(C_e, C.INT64)
	*e = make([]int64, L_e)
	for i := range (*e) {
		(*e)[i] = *(*int64)(unsafe.Pointer(uintptr(P_e) + uintptr(i * C.sizeof_int64_t)))
	}

	C.Plugify_DeleteVector(C_e, C.INT64)

}

func ParamRef6(a *int32, b *float32, c *float64, d *Vector4, e *[]int64, f *int8)  {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))
	C_e := C.Plugify_CreateVector(unsafe.Pointer(&(*e)[0]), C.ptrdiff_t(len(*e)), C.INT64)
	C_f := (*C.char)(f)

	C.ParamRef6(C_a, C_b, C_c, &C_d, C_e, C_f)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
	L_e := C.Plugify_GetVectorSize(C_e, C.INT64)
	P_e := C.Plugify_GetVectorData(C_e, C.INT64)
	*e = make([]int64, L_e)
	for i := range (*e) {
		(*e)[i] = *(*int64)(unsafe.Pointer(uintptr(P_e) + uintptr(i * C.sizeof_int64_t)))
	}

	C.Plugify_DeleteVector(C_e, C.INT64)

}

func ParamRef7(a *int32, b *float32, c *float64, d *Vector4, e *[]int64, f *int8, g *string)  {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))
	C_e := C.Plugify_CreateVector(unsafe.Pointer(&(*e)[0]), C.ptrdiff_t(len(*e)), C.INT64)
	C_f := (*C.char)(f)
	C_g := C.Plugify_CreateString(*g)

	C.ParamRef7(C_a, C_b, C_c, &C_d, C_e, C_f, C_g)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
	L_e := C.Plugify_GetVectorSize(C_e, C.INT64)
	P_e := C.Plugify_GetVectorData(C_e, C.INT64)
	*e = make([]int64, L_e)
	for i := range (*e) {
		(*e)[i] = *(*int64)(unsafe.Pointer(uintptr(P_e) + uintptr(i * C.sizeof_int64_t)))
	}
	P_g := C.Plugify_GetStringData(C_g)
	*g = C.GoString(P_g)

	C.Plugify_DeleteVector(C_e, C.INT64)
	C.Plugify_DeleteString(C_g)

}

func ParamRef8(a *int32, b *float32, c *float64, d *Vector4, e *[]int64, f *int8, g *string, h *float32)  {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))
	C_e := C.Plugify_CreateVector(unsafe.Pointer(&(*e)[0]), C.ptrdiff_t(len(*e)), C.INT64)
	C_f := (*C.char)(f)
	C_g := C.Plugify_CreateString(*g)
	C_h := (*C.float)(h)

	C.ParamRef8(C_a, C_b, C_c, &C_d, C_e, C_f, C_g, C_h)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
	L_e := C.Plugify_GetVectorSize(C_e, C.INT64)
	P_e := C.Plugify_GetVectorData(C_e, C.INT64)
	*e = make([]int64, L_e)
	for i := range (*e) {
		(*e)[i] = *(*int64)(unsafe.Pointer(uintptr(P_e) + uintptr(i * C.sizeof_int64_t)))
	}
	P_g := C.Plugify_GetStringData(C_g)
	*g = C.GoString(P_g)

	C.Plugify_DeleteVector(C_e, C.INT64)
	C.Plugify_DeleteString(C_g)

}

func ParamRef9(a *int32, b *float32, c *float64, d *Vector4, e *[]int64, f *int8, g *string, h *float32, k *int16)  {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))
	C_e := C.Plugify_CreateVector(unsafe.Pointer(&(*e)[0]), C.ptrdiff_t(len(*e)), C.INT64)
	C_f := (*C.char)(f)
	C_g := C.Plugify_CreateString(*g)
	C_h := (*C.float)(h)
	C_k := (*C.int16_t)(k)

	C.ParamRef9(C_a, C_b, C_c, &C_d, C_e, C_f, C_g, C_h, C_k)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
	L_e := C.Plugify_GetVectorSize(C_e, C.INT64)
	P_e := C.Plugify_GetVectorData(C_e, C.INT64)
	*e = make([]int64, L_e)
	for i := range (*e) {
		(*e)[i] = *(*int64)(unsafe.Pointer(uintptr(P_e) + uintptr(i * C.sizeof_int64_t)))
	}
	P_g := C.Plugify_GetStringData(C_g)
	*g = C.GoString(P_g)

	C.Plugify_DeleteVector(C_e, C.INT64)
	C.Plugify_DeleteString(C_g)

}

func ParamRef10(a *int32, b *float32, c *float64, d *Vector4, e *[]int64, f *int8, g *string, h *float32, k *int16, l *uintptr)  {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))
	C_e := C.Plugify_CreateVector(unsafe.Pointer(&(*e)[0]), C.ptrdiff_t(len(*e)), C.INT64)
	C_f := (*C.char)(f)
	C_g := C.Plugify_CreateString(*g)
	C_h := (*C.float)(h)
	C_k := (*C.int16_t)(k)
	C_l := (*C.uintptr_t)(unsafe.Pointer(l))

	C.ParamRef10(C_a, C_b, C_c, &C_d, C_e, C_f, C_g, C_h, C_k, C_l)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
	L_e := C.Plugify_GetVectorSize(C_e, C.INT64)
	P_e := C.Plugify_GetVectorData(C_e, C.INT64)
	*e = make([]int64, L_e)
	for i := range (*e) {
		(*e)[i] = *(*int64)(unsafe.Pointer(uintptr(P_e) + uintptr(i * C.sizeof_int64_t)))
	}
	P_g := C.Plugify_GetStringData(C_g)
	*g = C.GoString(P_g)

	C.Plugify_DeleteVector(C_e, C.INT64)
	C.Plugify_DeleteString(C_g)

}

func ParamRefVectors(p1 *[]bool, p2 *[]int8, p3 *[]uint16, p4 *[]int8, p5 *[]int16, p6 *[]int32, p7 *[]int64, p8 *[]uint8, p9 *[]uint16, p10 *[]uint32, p11 *[]uint64, p12 *[]uintptr, p13 *[]float32, p14 *[]float64, p15 *[]string)  {
	C_p1 := C.Plugify_CreateVector(unsafe.Pointer(&(*p1)[0]), C.ptrdiff_t(len(*p1)), C.BOOL)
	C_p2 := C.Plugify_CreateVector(unsafe.Pointer(&(*p2)[0]), C.ptrdiff_t(len(*p2)), C.CHAR8)
	C_p3 := C.Plugify_CreateVector(unsafe.Pointer(&(*p3)[0]), C.ptrdiff_t(len(*p3)), C.CHAR16)
	C_p4 := C.Plugify_CreateVector(unsafe.Pointer(&(*p4)[0]), C.ptrdiff_t(len(*p4)), C.INT8)
	C_p5 := C.Plugify_CreateVector(unsafe.Pointer(&(*p5)[0]), C.ptrdiff_t(len(*p5)), C.INT16)
	C_p6 := C.Plugify_CreateVector(unsafe.Pointer(&(*p6)[0]), C.ptrdiff_t(len(*p6)), C.INT32)
	C_p7 := C.Plugify_CreateVector(unsafe.Pointer(&(*p7)[0]), C.ptrdiff_t(len(*p7)), C.INT64)
	C_p8 := C.Plugify_CreateVector(unsafe.Pointer(&(*p8)[0]), C.ptrdiff_t(len(*p8)), C.UINT8)
	C_p9 := C.Plugify_CreateVector(unsafe.Pointer(&(*p9)[0]), C.ptrdiff_t(len(*p9)), C.UINT16)
	C_p10 := C.Plugify_CreateVector(unsafe.Pointer(&(*p10)[0]), C.ptrdiff_t(len(*p10)), C.UINT32)
	C_p11 := C.Plugify_CreateVector(unsafe.Pointer(&(*p11)[0]), C.ptrdiff_t(len(*p11)), C.UINT64)
	C_p12 := C.Plugify_CreateVector(unsafe.Pointer(&(*p12)[0]), C.ptrdiff_t(len(*p12)), C.UINTPTR)
	C_p13 := C.Plugify_CreateVector(unsafe.Pointer(&(*p13)[0]), C.ptrdiff_t(len(*p13)), C.FLOAT)
	C_p14 := C.Plugify_CreateVector(unsafe.Pointer(&(*p14)[0]), C.ptrdiff_t(len(*p14)), C.DOUBLE)
	C_p15 := C.Plugify_CreateVector(unsafe.Pointer(&(*p15)[0]), C.ptrdiff_t(len(*p15)), C.STRING)

	C.ParamRefVectors(C_p1, C_p2, C_p3, C_p4, C_p5, C_p6, C_p7, C_p8, C_p9, C_p10, C_p11, C_p12, C_p13, C_p14, C_p15)

	L_p1 := C.Plugify_GetVectorSize(C_p1, C.BOOL)
	P_p1 := C.Plugify_GetVectorData(C_p1, C.BOOL)
	*p1 = make([]bool, L_p1)
	for i := range (*p1) {
		(*p1)[i] = *(*bool)(unsafe.Pointer(uintptr(P_p1) + uintptr(i * C.sizeof_bool)))
	}
	C.Plugify_DeleteVectorDataBool(P_p1)
	L_p2 := C.Plugify_GetVectorSize(C_p2, C.CHAR8)
	P_p2 := C.Plugify_GetVectorData(C_p2, C.CHAR8)
	*p2 = make([]int8, L_p2)
	for i := range (*p2) {
		(*p2)[i] = *(*int8)(unsafe.Pointer(uintptr(P_p2) + uintptr(i * C.sizeof_char)))
	}
	L_p3 := C.Plugify_GetVectorSize(C_p3, C.CHAR16)
	P_p3 := C.Plugify_GetVectorData(C_p3, C.CHAR16)
	*p3 = make([]uint16, L_p3)
	for i := range (*p3) {
		(*p3)[i] = *(*uint16)(unsafe.Pointer(uintptr(P_p3) + uintptr(i * C.sizeof_uint16_t)))
	}
	L_p4 := C.Plugify_GetVectorSize(C_p4, C.INT8)
	P_p4 := C.Plugify_GetVectorData(C_p4, C.INT8)
	*p4 = make([]int8, L_p4)
	for i := range (*p4) {
		(*p4)[i] = *(*int8)(unsafe.Pointer(uintptr(P_p4) + uintptr(i * C.sizeof_int8_t)))
	}
	L_p5 := C.Plugify_GetVectorSize(C_p5, C.INT16)
	P_p5 := C.Plugify_GetVectorData(C_p5, C.INT16)
	*p5 = make([]int16, L_p5)
	for i := range (*p5) {
		(*p5)[i] = *(*int16)(unsafe.Pointer(uintptr(P_p5) + uintptr(i * C.sizeof_int16_t)))
	}
	L_p6 := C.Plugify_GetVectorSize(C_p6, C.INT32)
	P_p6 := C.Plugify_GetVectorData(C_p6, C.INT32)
	*p6 = make([]int32, L_p6)
	for i := range (*p6) {
		(*p6)[i] = *(*int32)(unsafe.Pointer(uintptr(P_p6) + uintptr(i * C.sizeof_int32_t)))
	}
	L_p7 := C.Plugify_GetVectorSize(C_p7, C.INT64)
	P_p7 := C.Plugify_GetVectorData(C_p7, C.INT64)
	*p7 = make([]int64, L_p7)
	for i := range (*p7) {
		(*p7)[i] = *(*int64)(unsafe.Pointer(uintptr(P_p7) + uintptr(i * C.sizeof_int64_t)))
	}
	L_p8 := C.Plugify_GetVectorSize(C_p8, C.UINT8)
	P_p8 := C.Plugify_GetVectorData(C_p8, C.UINT8)
	*p8 = make([]uint8, L_p8)
	for i := range (*p8) {
		(*p8)[i] = *(*uint8)(unsafe.Pointer(uintptr(P_p8) + uintptr(i * C.sizeof_uint8_t)))
	}
	L_p9 := C.Plugify_GetVectorSize(C_p9, C.UINT16)
	P_p9 := C.Plugify_GetVectorData(C_p9, C.UINT16)
	*p9 = make([]uint16, L_p9)
	for i := range (*p9) {
		(*p9)[i] = *(*uint16)(unsafe.Pointer(uintptr(P_p9) + uintptr(i * C.sizeof_uint16_t)))
	}
	L_p10 := C.Plugify_GetVectorSize(C_p10, C.UINT32)
	P_p10 := C.Plugify_GetVectorData(C_p10, C.UINT32)
	*p10 = make([]uint32, L_p10)
	for i := range (*p10) {
		(*p10)[i] = *(*uint32)(unsafe.Pointer(uintptr(P_p10) + uintptr(i * C.sizeof_uint32_t)))
	}
	L_p11 := C.Plugify_GetVectorSize(C_p11, C.UINT64)
	P_p11 := C.Plugify_GetVectorData(C_p11, C.UINT64)
	*p11 = make([]uint64, L_p11)
	for i := range (*p11) {
		(*p11)[i] = *(*uint64)(unsafe.Pointer(uintptr(P_p11) + uintptr(i * C.sizeof_uint64_t)))
	}
	L_p12 := C.Plugify_GetVectorSize(C_p12, C.UINTPTR)
	P_p12 := C.Plugify_GetVectorData(C_p12, C.UINTPTR)
	*p12 = make([]uintptr, L_p12)
	for i := range (*p12) {
		(*p12)[i] = *(*uintptr)(unsafe.Pointer(uintptr(P_p12) + uintptr(i * C.sizeof_uintptr_t)))
	}
	L_p13 := C.Plugify_GetVectorSize(C_p13, C.FLOAT)
	P_p13 := C.Plugify_GetVectorData(C_p13, C.FLOAT)
	*p13 = make([]float32, L_p13)
	for i := range (*p13) {
		(*p13)[i] = *(*float32)(unsafe.Pointer(uintptr(P_p13) + uintptr(i * C.sizeof_float)))
	}
	L_p14 := C.Plugify_GetVectorSize(C_p14, C.DOUBLE)
	P_p14 := C.Plugify_GetVectorData(C_p14, C.DOUBLE)
	*p14 = make([]float64, L_p14)
	for i := range (*p14) {
		(*p14)[i] = *(*float64)(unsafe.Pointer(uintptr(P_p14) + uintptr(i * C.sizeof_double)))
	}
	L_p15 := C.Plugify_GetVectorSize(C_p15, C.STRING)
	P_p15 := C.Plugify_GetVectorData(C_p15, C.STRING)
	*p15 = make([]string, L_p15)
	for i := range (*p15) {
		(*p15)[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(P_p15) + uintptr(i * C.sizeof_uintptr_t))))
	}
	C.Plugify_DeleteVectorDataCStr(P_p15)

	C.Plugify_DeleteVector(C_p1, C.BOOL)
	C.Plugify_DeleteVector(C_p2, C.CHAR8)
	C.Plugify_DeleteVector(C_p3, C.CHAR16)
	C.Plugify_DeleteVector(C_p4, C.INT8)
	C.Plugify_DeleteVector(C_p5, C.INT16)
	C.Plugify_DeleteVector(C_p6, C.INT32)
	C.Plugify_DeleteVector(C_p7, C.INT64)
	C.Plugify_DeleteVector(C_p8, C.UINT8)
	C.Plugify_DeleteVector(C_p9, C.UINT16)
	C.Plugify_DeleteVector(C_p10, C.UINT32)
	C.Plugify_DeleteVector(C_p11, C.UINT64)
	C.Plugify_DeleteVector(C_p12, C.UINTPTR)
	C.Plugify_DeleteVector(C_p13, C.FLOAT)
	C.Plugify_DeleteVector(C_p14, C.DOUBLE)
	C.Plugify_DeleteVector(C_p15, C.STRING)

}

func ParamAllPrimitives(p1 bool, p2 uint16, p3 int8, p4 int16, p5 int32, p6 int64, p7 uint8, p8 uint16, p9 uint32, p10 uint64, p11 uintptr, p12 float32, p13 float64) int64 {
	C_p1 := C.bool(p1)
	C_p2 := C.uint16_t(p2)
	C_p3 := C.int8_t(p3)
	C_p4 := C.int16_t(p4)
	C_p5 := C.int32_t(p5)
	C_p6 := C.int64_t(p6)
	C_p7 := C.uint8_t(p7)
	C_p8 := C.uint16_t(p8)
	C_p9 := C.uint32_t(p9)
	C_p10 := C.uint64_t(p10)
	C_p11 := C.uintptr_t(p11)
	C_p12 := C.float(p12)
	C_p13 := C.double(p13)

	result := int64(C.ParamAllPrimitives(C_p1, C_p2, C_p3, C_p4, C_p5, C_p6, C_p7, C_p8, C_p9, C_p10, C_p11, C_p12, C_p13))
	return result
}

