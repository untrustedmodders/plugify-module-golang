package cross_call_master

//generated with https://github.com/untrustedmodders/plugify-module-cpp/blob/main/generator/generator.py from cross_call_master

// #include "cross_call_master.h"
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
	M [4][4]float32
}

func ReverseReturn(returnString string) {
	C_returnString := C.Plugify_ConstructString(returnString)

	C.ReverseReturn(&C_returnString)

	C.Plugify_DestroyString(&C_returnString)

}

func NoParamReturnVoidCallback() {
	C.NoParamReturnVoidCallback()
}

func NoParamReturnBoolCallback() bool {
	result := bool(C.NoParamReturnBoolCallback())
	return result
}

func NoParamReturnChar8Callback() int8 {
	result := int8(C.NoParamReturnChar8Callback())
	return result
}

func NoParamReturnChar16Callback() uint16 {
	result := uint16(C.NoParamReturnChar16Callback())
	return result
}

func NoParamReturnInt8Callback() int8 {
	result := int8(C.NoParamReturnInt8Callback())
	return result
}

func NoParamReturnInt16Callback() int16 {
	result := int16(C.NoParamReturnInt16Callback())
	return result
}

func NoParamReturnInt32Callback() int32 {
	result := int32(C.NoParamReturnInt32Callback())
	return result
}

func NoParamReturnInt64Callback() int64 {
	result := int64(C.NoParamReturnInt64Callback())
	return result
}

func NoParamReturnUInt8Callback() uint8 {
	result := uint8(C.NoParamReturnUInt8Callback())
	return result
}

func NoParamReturnUInt16Callback() uint16 {
	result := uint16(C.NoParamReturnUInt16Callback())
	return result
}

func NoParamReturnUInt32Callback() uint32 {
	result := uint32(C.NoParamReturnUInt32Callback())
	return result
}

func NoParamReturnUInt64Callback() uint64 {
	result := uint64(C.NoParamReturnUInt64Callback())
	return result
}

func NoParamReturnPointerCallback() uintptr {
	result := uintptr(C.NoParamReturnPointerCallback())
	return result
}

func NoParamReturnFloatCallback() float32 {
	result := float32(C.NoParamReturnFloatCallback())
	return result
}

func NoParamReturnDoubleCallback() float64 {
	result := float64(C.NoParamReturnDoubleCallback())
	return result
}

func NoParamReturnFunctionCallback() uintptr {
	result := uintptr(C.NoParamReturnFunctionCallback())
	return result
}

func NoParamReturnStringCallback() string {
	C_output := C.NoParamReturnStringCallback()

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func NoParamReturnArrayBoolCallback() []bool {
	C_output := C.NoParamReturnArrayBoolCallback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.BOOL)
	P_output := C.Plugify_GetVectorData(&C_output, C.BOOL)
	output := make([]bool, L_output)
	for i := range output {
		output[i] = *(*bool)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_bool)))
	}

	C.Plugify_DestroyVector(&C_output, C.BOOL)
	return output
}

func NoParamReturnArrayChar8Callback() []int8 {
	C_output := C.NoParamReturnArrayChar8Callback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.CHAR8)
	P_output := C.Plugify_GetVectorData(&C_output, C.CHAR8)
	output := make([]int8, L_output)
	for i := range output {
		output[i] = *(*int8)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_char)))
	}

	C.Plugify_DestroyVector(&C_output, C.CHAR8)
	return output
}

func NoParamReturnArrayChar16Callback() []uint16 {
	C_output := C.NoParamReturnArrayChar16Callback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.CHAR16)
	P_output := C.Plugify_GetVectorData(&C_output, C.CHAR16)
	output := make([]uint16, L_output)
	for i := range output {
		output[i] = *(*uint16)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_uint16_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.CHAR16)
	return output
}

func NoParamReturnArrayInt8Callback() []int8 {
	C_output := C.NoParamReturnArrayInt8Callback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.INT8)
	P_output := C.Plugify_GetVectorData(&C_output, C.INT8)
	output := make([]int8, L_output)
	for i := range output {
		output[i] = *(*int8)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_int8_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.INT8)
	return output
}

func NoParamReturnArrayInt16Callback() []int16 {
	C_output := C.NoParamReturnArrayInt16Callback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.INT16)
	P_output := C.Plugify_GetVectorData(&C_output, C.INT16)
	output := make([]int16, L_output)
	for i := range output {
		output[i] = *(*int16)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_int16_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.INT16)
	return output
}

func NoParamReturnArrayInt32Callback() []int32 {
	C_output := C.NoParamReturnArrayInt32Callback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.INT32)
	P_output := C.Plugify_GetVectorData(&C_output, C.INT32)
	output := make([]int32, L_output)
	for i := range output {
		output[i] = *(*int32)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_int32_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.INT32)
	return output
}

func NoParamReturnArrayInt64Callback() []int64 {
	C_output := C.NoParamReturnArrayInt64Callback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.INT64)
	P_output := C.Plugify_GetVectorData(&C_output, C.INT64)
	output := make([]int64, L_output)
	for i := range output {
		output[i] = *(*int64)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_int64_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.INT64)
	return output
}

func NoParamReturnArrayUInt8Callback() []uint8 {
	C_output := C.NoParamReturnArrayUInt8Callback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.UINT8)
	P_output := C.Plugify_GetVectorData(&C_output, C.UINT8)
	output := make([]uint8, L_output)
	for i := range output {
		output[i] = *(*uint8)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_uint8_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.UINT8)
	return output
}

func NoParamReturnArrayUInt16Callback() []uint16 {
	C_output := C.NoParamReturnArrayUInt16Callback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.UINT16)
	P_output := C.Plugify_GetVectorData(&C_output, C.UINT16)
	output := make([]uint16, L_output)
	for i := range output {
		output[i] = *(*uint16)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_uint16_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.UINT16)
	return output
}

func NoParamReturnArrayUInt32Callback() []uint32 {
	C_output := C.NoParamReturnArrayUInt32Callback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.UINT32)
	P_output := C.Plugify_GetVectorData(&C_output, C.UINT32)
	output := make([]uint32, L_output)
	for i := range output {
		output[i] = *(*uint32)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_uint32_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.UINT32)
	return output
}

func NoParamReturnArrayUInt64Callback() []uint64 {
	C_output := C.NoParamReturnArrayUInt64Callback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.UINT64)
	P_output := C.Plugify_GetVectorData(&C_output, C.UINT64)
	output := make([]uint64, L_output)
	for i := range output {
		output[i] = *(*uint64)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_uint64_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.UINT64)
	return output
}

func NoParamReturnArrayPointerCallback() []uintptr {
	C_output := C.NoParamReturnArrayPointerCallback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.POINTER)
	P_output := C.Plugify_GetVectorData(&C_output, C.POINTER)
	output := make([]uintptr, L_output)
	for i := range output {
		output[i] = *(*uintptr)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_uintptr_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.POINTER)
	return output
}

func NoParamReturnArrayFloatCallback() []float32 {
	C_output := C.NoParamReturnArrayFloatCallback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.FLOAT)
	P_output := C.Plugify_GetVectorData(&C_output, C.FLOAT)
	output := make([]float32, L_output)
	for i := range output {
		output[i] = *(*float32)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_float)))
	}

	C.Plugify_DestroyVector(&C_output, C.FLOAT)
	return output
}

func NoParamReturnArrayDoubleCallback() []float64 {
	C_output := C.NoParamReturnArrayDoubleCallback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.DOUBLE)
	P_output := C.Plugify_GetVectorData(&C_output, C.DOUBLE)
	output := make([]float64, L_output)
	for i := range output {
		output[i] = *(*float64)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_double)))
	}

	C.Plugify_DestroyVector(&C_output, C.DOUBLE)
	return output
}

func NoParamReturnArrayStringCallback() []string {
	C_output := C.NoParamReturnArrayStringCallback()

	L_output := C.Plugify_GetVectorSize(&C_output, C.STRING)
	P_output := C.Plugify_GetVectorData(&C_output, C.STRING)
	output := make([]string, L_output)
	for i := range output {
		output[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(P_output) + uintptr(i*C.sizeof_uintptr_t))))
	}

	C.Plugify_DeleteVectorDataCStr(P_output)
	C.Plugify_DestroyVector(&C_output, C.STRING)
	return output
}

func NoParamReturnVector2Callback() Vector2 {
	C_result := C.NoParamReturnVector2Callback()
	return *(*Vector2)(unsafe.Pointer(&C_result))
}

func NoParamReturnVector3Callback() Vector3 {
	C_result := C.NoParamReturnVector3Callback()
	return *(*Vector3)(unsafe.Pointer(&C_result))
}

func NoParamReturnVector4Callback() Vector4 {
	C_result := C.NoParamReturnVector4Callback()
	return *(*Vector4)(unsafe.Pointer(&C_result))
}

func NoParamReturnMatrix4x4Callback() Matrix4x4 {
	C_result := C.NoParamReturnMatrix4x4Callback()
	return *(*Matrix4x4)(unsafe.Pointer(&C_result))
}

func Param1Callback(a int32) {
	C_a := C.int32_t(a)

	C.Param1Callback(C_a)
}

func Param2Callback(a int32, b float32) {
	C_a := C.int32_t(a)
	C_b := C.float(b)

	C.Param2Callback(C_a, C_b)
}

func Param3Callback(a int32, b float32, c float64) {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)

	C.Param3Callback(C_a, C_b, C_c)
}

func Param4Callback(a int32, b float32, c float64, d Vector4) {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))

	C.Param4Callback(C_a, C_b, C_c, &C_d)
}

func Param5Callback(a int32, b float32, c float64, d Vector4, e []int64) {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))
	var A_e unsafe.Pointer
	S_e := len(e)
	if S_e > 0 {
		A_e = unsafe.Pointer(&e[0])
	} else {
		A_e = nil
	}
	C_e := C.Plugify_ConstructVector(A_e, C.ptrdiff_t(S_e), C.INT64)

	C.Param5Callback(C_a, C_b, C_c, &C_d, &C_e)

	C.Plugify_DestroyVector(&C_e, C.INT64)

}

func Param6Callback(a int32, b float32, c float64, d Vector4, e []int64, f int8) {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))
	var A_e unsafe.Pointer
	S_e := len(e)
	if S_e > 0 {
		A_e = unsafe.Pointer(&e[0])
	} else {
		A_e = nil
	}
	C_e := C.Plugify_ConstructVector(A_e, C.ptrdiff_t(S_e), C.INT64)
	C_f := C.char(f)

	C.Param6Callback(C_a, C_b, C_c, &C_d, &C_e, C_f)

	C.Plugify_DestroyVector(&C_e, C.INT64)

}

func Param7Callback(a int32, b float32, c float64, d Vector4, e []int64, f int8, g string) {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))
	var A_e unsafe.Pointer
	S_e := len(e)
	if S_e > 0 {
		A_e = unsafe.Pointer(&e[0])
	} else {
		A_e = nil
	}
	C_e := C.Plugify_ConstructVector(A_e, C.ptrdiff_t(S_e), C.INT64)
	C_f := C.char(f)
	C_g := C.Plugify_ConstructString(g)

	C.Param7Callback(C_a, C_b, C_c, &C_d, &C_e, C_f, &C_g)

	C.Plugify_DestroyVector(&C_e, C.INT64)
	C.Plugify_DestroyString(&C_g)

}

func Param8Callback(a int32, b float32, c float64, d Vector4, e []int64, f int8, g string, h uint16) {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))
	var A_e unsafe.Pointer
	S_e := len(e)
	if S_e > 0 {
		A_e = unsafe.Pointer(&e[0])
	} else {
		A_e = nil
	}
	C_e := C.Plugify_ConstructVector(A_e, C.ptrdiff_t(S_e), C.INT64)
	C_f := C.char(f)
	C_g := C.Plugify_ConstructString(g)
	C_h := C.uint16_t(h)

	C.Param8Callback(C_a, C_b, C_c, &C_d, &C_e, C_f, &C_g, C_h)

	C.Plugify_DestroyVector(&C_e, C.INT64)
	C.Plugify_DestroyString(&C_g)

}

func Param9Callback(a int32, b float32, c float64, d Vector4, e []int64, f int8, g string, h uint16, k int16) {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))
	var A_e unsafe.Pointer
	S_e := len(e)
	if S_e > 0 {
		A_e = unsafe.Pointer(&e[0])
	} else {
		A_e = nil
	}
	C_e := C.Plugify_ConstructVector(A_e, C.ptrdiff_t(S_e), C.INT64)
	C_f := C.char(f)
	C_g := C.Plugify_ConstructString(g)
	C_h := C.uint16_t(h)
	C_k := C.int16_t(k)

	C.Param9Callback(C_a, C_b, C_c, &C_d, &C_e, C_f, &C_g, C_h, C_k)

	C.Plugify_DestroyVector(&C_e, C.INT64)
	C.Plugify_DestroyString(&C_g)

}

func Param10Callback(a int32, b float32, c float64, d Vector4, e []int64, f int8, g string, h uint16, k int16, l uintptr) {
	C_a := C.int32_t(a)
	C_b := C.float(b)
	C_c := C.double(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(&d))
	var A_e unsafe.Pointer
	S_e := len(e)
	if S_e > 0 {
		A_e = unsafe.Pointer(&e[0])
	} else {
		A_e = nil
	}
	C_e := C.Plugify_ConstructVector(A_e, C.ptrdiff_t(S_e), C.INT64)
	C_f := C.char(f)
	C_g := C.Plugify_ConstructString(g)
	C_h := C.uint16_t(h)
	C_k := C.int16_t(k)
	C_l := C.uintptr_t(l)

	C.Param10Callback(C_a, C_b, C_c, &C_d, &C_e, C_f, &C_g, C_h, C_k, C_l)

	C.Plugify_DestroyVector(&C_e, C.INT64)
	C.Plugify_DestroyString(&C_g)

}

func ParamRef1Callback(a *int32) {
	C_a := (*C.int32_t)(a)

	C.ParamRef1Callback(C_a)
}

func ParamRef2Callback(a *int32, b *float32) {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)

	C.ParamRef2Callback(C_a, C_b)
}

func ParamRef3Callback(a *int32, b *float32, c *float64) {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)

	C.ParamRef3Callback(C_a, C_b, C_c)
}

func ParamRef4Callback(a *int32, b *float32, c *float64, d *Vector4) {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))

	C.ParamRef4Callback(C_a, C_b, C_c, &C_d)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
}

func ParamRef5Callback(a *int32, b *float32, c *float64, d *Vector4, e *[]int64) {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))
	var A_e unsafe.Pointer
	S_e := len(*e)
	if S_e > 0 {
		A_e = unsafe.Pointer(&(*e)[0])
	} else {
		A_e = nil
	}
	C_e := C.Plugify_ConstructVector(A_e, C.ptrdiff_t(S_e), C.INT64)

	C.ParamRef5Callback(C_a, C_b, C_c, &C_d, &C_e)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
	L_e := C.Plugify_GetVectorSize(&C_e, C.INT64)
	P_e := C.Plugify_GetVectorData(&C_e, C.INT64)
	*e = make([]int64, L_e)
	for i := range *e {
		(*e)[i] = *(*int64)(unsafe.Pointer(uintptr(P_e) + uintptr(i*C.sizeof_int64_t)))
	}

	C.Plugify_DestroyVector(&C_e, C.INT64)

}

func ParamRef6Callback(a *int32, b *float32, c *float64, d *Vector4, e *[]int64, f *int8) {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))
	var A_e unsafe.Pointer
	S_e := len(*e)
	if S_e > 0 {
		A_e = unsafe.Pointer(&(*e)[0])
	} else {
		A_e = nil
	}
	C_e := C.Plugify_ConstructVector(A_e, C.ptrdiff_t(S_e), C.INT64)
	C_f := (*C.char)(f)

	C.ParamRef6Callback(C_a, C_b, C_c, &C_d, &C_e, C_f)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
	L_e := C.Plugify_GetVectorSize(&C_e, C.INT64)
	P_e := C.Plugify_GetVectorData(&C_e, C.INT64)
	*e = make([]int64, L_e)
	for i := range *e {
		(*e)[i] = *(*int64)(unsafe.Pointer(uintptr(P_e) + uintptr(i*C.sizeof_int64_t)))
	}

	C.Plugify_DestroyVector(&C_e, C.INT64)

}

func ParamRef7Callback(a *int32, b *float32, c *float64, d *Vector4, e *[]int64, f *int8, g *string) {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))
	var A_e unsafe.Pointer
	S_e := len(*e)
	if S_e > 0 {
		A_e = unsafe.Pointer(&(*e)[0])
	} else {
		A_e = nil
	}
	C_e := C.Plugify_ConstructVector(A_e, C.ptrdiff_t(S_e), C.INT64)
	C_f := (*C.char)(f)
	C_g := C.Plugify_ConstructString(*g)

	C.ParamRef7Callback(C_a, C_b, C_c, &C_d, &C_e, C_f, &C_g)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
	L_e := C.Plugify_GetVectorSize(&C_e, C.INT64)
	P_e := C.Plugify_GetVectorData(&C_e, C.INT64)
	*e = make([]int64, L_e)
	for i := range *e {
		(*e)[i] = *(*int64)(unsafe.Pointer(uintptr(P_e) + uintptr(i*C.sizeof_int64_t)))
	}
	P_g := C.Plugify_GetStringData(&C_g)
	*g = C.GoString(P_g)

	C.Plugify_DestroyVector(&C_e, C.INT64)
	C.Plugify_DestroyString(&C_g)

}

func ParamRef8Callback(a *int32, b *float32, c *float64, d *Vector4, e *[]int64, f *int8, g *string, h *uint16) {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))
	var A_e unsafe.Pointer
	S_e := len(*e)
	if S_e > 0 {
		A_e = unsafe.Pointer(&(*e)[0])
	} else {
		A_e = nil
	}
	C_e := C.Plugify_ConstructVector(A_e, C.ptrdiff_t(S_e), C.INT64)
	C_f := (*C.char)(f)
	C_g := C.Plugify_ConstructString(*g)
	C_h := (*C.uint16_t)(h)

	C.ParamRef8Callback(C_a, C_b, C_c, &C_d, &C_e, C_f, &C_g, C_h)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
	L_e := C.Plugify_GetVectorSize(&C_e, C.INT64)
	P_e := C.Plugify_GetVectorData(&C_e, C.INT64)
	*e = make([]int64, L_e)
	for i := range *e {
		(*e)[i] = *(*int64)(unsafe.Pointer(uintptr(P_e) + uintptr(i*C.sizeof_int64_t)))
	}
	P_g := C.Plugify_GetStringData(&C_g)
	*g = C.GoString(P_g)

	C.Plugify_DestroyVector(&C_e, C.INT64)
	C.Plugify_DestroyString(&C_g)

}

func ParamRef9Callback(a *int32, b *float32, c *float64, d *Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16) {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))
	var A_e unsafe.Pointer
	S_e := len(*e)
	if S_e > 0 {
		A_e = unsafe.Pointer(&(*e)[0])
	} else {
		A_e = nil
	}
	C_e := C.Plugify_ConstructVector(A_e, C.ptrdiff_t(S_e), C.INT64)
	C_f := (*C.char)(f)
	C_g := C.Plugify_ConstructString(*g)
	C_h := (*C.uint16_t)(h)
	C_k := (*C.int16_t)(k)

	C.ParamRef9Callback(C_a, C_b, C_c, &C_d, &C_e, C_f, &C_g, C_h, C_k)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
	L_e := C.Plugify_GetVectorSize(&C_e, C.INT64)
	P_e := C.Plugify_GetVectorData(&C_e, C.INT64)
	*e = make([]int64, L_e)
	for i := range *e {
		(*e)[i] = *(*int64)(unsafe.Pointer(uintptr(P_e) + uintptr(i*C.sizeof_int64_t)))
	}
	P_g := C.Plugify_GetStringData(&C_g)
	*g = C.GoString(P_g)

	C.Plugify_DestroyVector(&C_e, C.INT64)
	C.Plugify_DestroyString(&C_g)

}

func ParamRef10Callback(a *int32, b *float32, c *float64, d *Vector4, e *[]int64, f *int8, g *string, h *uint16, k *int16, l *uintptr) {
	C_a := (*C.int32_t)(a)
	C_b := (*C.float)(b)
	C_c := (*C.double)(c)
	C_d := *(*C.Vector4)(unsafe.Pointer(d))
	var A_e unsafe.Pointer
	S_e := len(*e)
	if S_e > 0 {
		A_e = unsafe.Pointer(&(*e)[0])
	} else {
		A_e = nil
	}
	C_e := C.Plugify_ConstructVector(A_e, C.ptrdiff_t(S_e), C.INT64)
	C_f := (*C.char)(f)
	C_g := C.Plugify_ConstructString(*g)
	C_h := (*C.uint16_t)(h)
	C_k := (*C.int16_t)(k)
	C_l := (*C.uintptr_t)(unsafe.Pointer(l))

	C.ParamRef10Callback(C_a, C_b, C_c, &C_d, &C_e, C_f, &C_g, C_h, C_k, C_l)

	*d = *(*Vector4)(unsafe.Pointer(&C_d))
	L_e := C.Plugify_GetVectorSize(&C_e, C.INT64)
	P_e := C.Plugify_GetVectorData(&C_e, C.INT64)
	*e = make([]int64, L_e)
	for i := range *e {
		(*e)[i] = *(*int64)(unsafe.Pointer(uintptr(P_e) + uintptr(i*C.sizeof_int64_t)))
	}
	P_g := C.Plugify_GetStringData(&C_g)
	*g = C.GoString(P_g)

	C.Plugify_DestroyVector(&C_e, C.INT64)
	C.Plugify_DestroyString(&C_g)

}

func ParamRefVectorsCallback(p1 *[]bool, p2 *[]int8, p3 *[]uint16, p4 *[]int8, p5 *[]int16, p6 *[]int32, p7 *[]int64, p8 *[]uint8, p9 *[]uint16, p10 *[]uint32, p11 *[]uint64, p12 *[]uintptr, p13 *[]float32, p14 *[]float64, p15 *[]string) {
	var A_p1 unsafe.Pointer
	S_p1 := len(*p1)
	if S_p1 > 0 {
		A_p1 = unsafe.Pointer(&(*p1)[0])
	} else {
		A_p1 = nil
	}
	C_p1 := C.Plugify_ConstructVector(A_p1, C.ptrdiff_t(S_p1), C.BOOL)
	var A_p2 unsafe.Pointer
	S_p2 := len(*p2)
	if S_p2 > 0 {
		A_p2 = unsafe.Pointer(&(*p2)[0])
	} else {
		A_p2 = nil
	}
	C_p2 := C.Plugify_ConstructVector(A_p2, C.ptrdiff_t(S_p2), C.CHAR8)
	var A_p3 unsafe.Pointer
	S_p3 := len(*p3)
	if S_p3 > 0 {
		A_p3 = unsafe.Pointer(&(*p3)[0])
	} else {
		A_p3 = nil
	}
	C_p3 := C.Plugify_ConstructVector(A_p3, C.ptrdiff_t(S_p3), C.CHAR16)
	var A_p4 unsafe.Pointer
	S_p4 := len(*p4)
	if S_p4 > 0 {
		A_p4 = unsafe.Pointer(&(*p4)[0])
	} else {
		A_p4 = nil
	}
	C_p4 := C.Plugify_ConstructVector(A_p4, C.ptrdiff_t(S_p4), C.INT8)
	var A_p5 unsafe.Pointer
	S_p5 := len(*p5)
	if S_p5 > 0 {
		A_p5 = unsafe.Pointer(&(*p5)[0])
	} else {
		A_p5 = nil
	}
	C_p5 := C.Plugify_ConstructVector(A_p5, C.ptrdiff_t(S_p5), C.INT16)
	var A_p6 unsafe.Pointer
	S_p6 := len(*p6)
	if S_p6 > 0 {
		A_p6 = unsafe.Pointer(&(*p6)[0])
	} else {
		A_p6 = nil
	}
	C_p6 := C.Plugify_ConstructVector(A_p6, C.ptrdiff_t(S_p6), C.INT32)
	var A_p7 unsafe.Pointer
	S_p7 := len(*p7)
	if S_p7 > 0 {
		A_p7 = unsafe.Pointer(&(*p7)[0])
	} else {
		A_p7 = nil
	}
	C_p7 := C.Plugify_ConstructVector(A_p7, C.ptrdiff_t(S_p7), C.INT64)
	var A_p8 unsafe.Pointer
	S_p8 := len(*p8)
	if S_p8 > 0 {
		A_p8 = unsafe.Pointer(&(*p8)[0])
	} else {
		A_p8 = nil
	}
	C_p8 := C.Plugify_ConstructVector(A_p8, C.ptrdiff_t(S_p8), C.UINT8)
	var A_p9 unsafe.Pointer
	S_p9 := len(*p9)
	if S_p9 > 0 {
		A_p9 = unsafe.Pointer(&(*p9)[0])
	} else {
		A_p9 = nil
	}
	C_p9 := C.Plugify_ConstructVector(A_p9, C.ptrdiff_t(S_p9), C.UINT16)
	var A_p10 unsafe.Pointer
	S_p10 := len(*p10)
	if S_p10 > 0 {
		A_p10 = unsafe.Pointer(&(*p10)[0])
	} else {
		A_p10 = nil
	}
	C_p10 := C.Plugify_ConstructVector(A_p10, C.ptrdiff_t(S_p10), C.UINT32)
	var A_p11 unsafe.Pointer
	S_p11 := len(*p11)
	if S_p11 > 0 {
		A_p11 = unsafe.Pointer(&(*p11)[0])
	} else {
		A_p11 = nil
	}
	C_p11 := C.Plugify_ConstructVector(A_p11, C.ptrdiff_t(S_p11), C.UINT64)
	var A_p12 unsafe.Pointer
	S_p12 := len(*p12)
	if S_p12 > 0 {
		A_p12 = unsafe.Pointer(&(*p12)[0])
	} else {
		A_p12 = nil
	}
	C_p12 := C.Plugify_ConstructVector(A_p12, C.ptrdiff_t(S_p12), C.POINTER)
	var A_p13 unsafe.Pointer
	S_p13 := len(*p13)
	if S_p13 > 0 {
		A_p13 = unsafe.Pointer(&(*p13)[0])
	} else {
		A_p13 = nil
	}
	C_p13 := C.Plugify_ConstructVector(A_p13, C.ptrdiff_t(S_p13), C.FLOAT)
	var A_p14 unsafe.Pointer
	S_p14 := len(*p14)
	if S_p14 > 0 {
		A_p14 = unsafe.Pointer(&(*p14)[0])
	} else {
		A_p14 = nil
	}
	C_p14 := C.Plugify_ConstructVector(A_p14, C.ptrdiff_t(S_p14), C.DOUBLE)
	var A_p15 unsafe.Pointer
	S_p15 := len(*p15)
	if S_p15 > 0 {
		A_p15 = unsafe.Pointer(&(*p15)[0])
	} else {
		A_p15 = nil
	}
	C_p15 := C.Plugify_ConstructVector(A_p15, C.ptrdiff_t(S_p15), C.STRING)

	C.ParamRefVectorsCallback(&C_p1, &C_p2, &C_p3, &C_p4, &C_p5, &C_p6, &C_p7, &C_p8, &C_p9, &C_p10, &C_p11, &C_p12, &C_p13, &C_p14, &C_p15)

	L_p1 := C.Plugify_GetVectorSize(&C_p1, C.BOOL)
	P_p1 := C.Plugify_GetVectorData(&C_p1, C.BOOL)
	*p1 = make([]bool, L_p1)
	for i := range *p1 {
		(*p1)[i] = *(*bool)(unsafe.Pointer(uintptr(P_p1) + uintptr(i*C.sizeof_bool)))
	}
	L_p2 := C.Plugify_GetVectorSize(&C_p2, C.CHAR8)
	P_p2 := C.Plugify_GetVectorData(&C_p2, C.CHAR8)
	*p2 = make([]int8, L_p2)
	for i := range *p2 {
		(*p2)[i] = *(*int8)(unsafe.Pointer(uintptr(P_p2) + uintptr(i*C.sizeof_char)))
	}
	L_p3 := C.Plugify_GetVectorSize(&C_p3, C.CHAR16)
	P_p3 := C.Plugify_GetVectorData(&C_p3, C.CHAR16)
	*p3 = make([]uint16, L_p3)
	for i := range *p3 {
		(*p3)[i] = *(*uint16)(unsafe.Pointer(uintptr(P_p3) + uintptr(i*C.sizeof_uint16_t)))
	}
	L_p4 := C.Plugify_GetVectorSize(&C_p4, C.INT8)
	P_p4 := C.Plugify_GetVectorData(&C_p4, C.INT8)
	*p4 = make([]int8, L_p4)
	for i := range *p4 {
		(*p4)[i] = *(*int8)(unsafe.Pointer(uintptr(P_p4) + uintptr(i*C.sizeof_int8_t)))
	}
	L_p5 := C.Plugify_GetVectorSize(&C_p5, C.INT16)
	P_p5 := C.Plugify_GetVectorData(&C_p5, C.INT16)
	*p5 = make([]int16, L_p5)
	for i := range *p5 {
		(*p5)[i] = *(*int16)(unsafe.Pointer(uintptr(P_p5) + uintptr(i*C.sizeof_int16_t)))
	}
	L_p6 := C.Plugify_GetVectorSize(&C_p6, C.INT32)
	P_p6 := C.Plugify_GetVectorData(&C_p6, C.INT32)
	*p6 = make([]int32, L_p6)
	for i := range *p6 {
		(*p6)[i] = *(*int32)(unsafe.Pointer(uintptr(P_p6) + uintptr(i*C.sizeof_int32_t)))
	}
	L_p7 := C.Plugify_GetVectorSize(&C_p7, C.INT64)
	P_p7 := C.Plugify_GetVectorData(&C_p7, C.INT64)
	*p7 = make([]int64, L_p7)
	for i := range *p7 {
		(*p7)[i] = *(*int64)(unsafe.Pointer(uintptr(P_p7) + uintptr(i*C.sizeof_int64_t)))
	}
	L_p8 := C.Plugify_GetVectorSize(&C_p8, C.UINT8)
	P_p8 := C.Plugify_GetVectorData(&C_p8, C.UINT8)
	*p8 = make([]uint8, L_p8)
	for i := range *p8 {
		(*p8)[i] = *(*uint8)(unsafe.Pointer(uintptr(P_p8) + uintptr(i*C.sizeof_uint8_t)))
	}
	L_p9 := C.Plugify_GetVectorSize(&C_p9, C.UINT16)
	P_p9 := C.Plugify_GetVectorData(&C_p9, C.UINT16)
	*p9 = make([]uint16, L_p9)
	for i := range *p9 {
		(*p9)[i] = *(*uint16)(unsafe.Pointer(uintptr(P_p9) + uintptr(i*C.sizeof_uint16_t)))
	}
	L_p10 := C.Plugify_GetVectorSize(&C_p10, C.UINT32)
	P_p10 := C.Plugify_GetVectorData(&C_p10, C.UINT32)
	*p10 = make([]uint32, L_p10)
	for i := range *p10 {
		(*p10)[i] = *(*uint32)(unsafe.Pointer(uintptr(P_p10) + uintptr(i*C.sizeof_uint32_t)))
	}
	L_p11 := C.Plugify_GetVectorSize(&C_p11, C.UINT64)
	P_p11 := C.Plugify_GetVectorData(&C_p11, C.UINT64)
	*p11 = make([]uint64, L_p11)
	for i := range *p11 {
		(*p11)[i] = *(*uint64)(unsafe.Pointer(uintptr(P_p11) + uintptr(i*C.sizeof_uint64_t)))
	}
	L_p12 := C.Plugify_GetVectorSize(&C_p12, C.POINTER)
	P_p12 := C.Plugify_GetVectorData(&C_p12, C.POINTER)
	*p12 = make([]uintptr, L_p12)
	for i := range *p12 {
		(*p12)[i] = *(*uintptr)(unsafe.Pointer(uintptr(P_p12) + uintptr(i*C.sizeof_uintptr_t)))
	}
	L_p13 := C.Plugify_GetVectorSize(&C_p13, C.FLOAT)
	P_p13 := C.Plugify_GetVectorData(&C_p13, C.FLOAT)
	*p13 = make([]float32, L_p13)
	for i := range *p13 {
		(*p13)[i] = *(*float32)(unsafe.Pointer(uintptr(P_p13) + uintptr(i*C.sizeof_float)))
	}
	L_p14 := C.Plugify_GetVectorSize(&C_p14, C.DOUBLE)
	P_p14 := C.Plugify_GetVectorData(&C_p14, C.DOUBLE)
	*p14 = make([]float64, L_p14)
	for i := range *p14 {
		(*p14)[i] = *(*float64)(unsafe.Pointer(uintptr(P_p14) + uintptr(i*C.sizeof_double)))
	}
	L_p15 := C.Plugify_GetVectorSize(&C_p15, C.STRING)
	P_p15 := C.Plugify_GetVectorData(&C_p15, C.STRING)
	*p15 = make([]string, L_p15)
	for i := range *p15 {
		(*p15)[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(P_p15) + uintptr(i*C.sizeof_uintptr_t))))
	}
	C.Plugify_DeleteVectorDataCStr(P_p15)

	C.Plugify_DestroyVector(&C_p1, C.BOOL)
	C.Plugify_DestroyVector(&C_p2, C.CHAR8)
	C.Plugify_DestroyVector(&C_p3, C.CHAR16)
	C.Plugify_DestroyVector(&C_p4, C.INT8)
	C.Plugify_DestroyVector(&C_p5, C.INT16)
	C.Plugify_DestroyVector(&C_p6, C.INT32)
	C.Plugify_DestroyVector(&C_p7, C.INT64)
	C.Plugify_DestroyVector(&C_p8, C.UINT8)
	C.Plugify_DestroyVector(&C_p9, C.UINT16)
	C.Plugify_DestroyVector(&C_p10, C.UINT32)
	C.Plugify_DestroyVector(&C_p11, C.UINT64)
	C.Plugify_DestroyVector(&C_p12, C.POINTER)
	C.Plugify_DestroyVector(&C_p13, C.FLOAT)
	C.Plugify_DestroyVector(&C_p14, C.DOUBLE)
	C.Plugify_DestroyVector(&C_p15, C.STRING)

}

func ParamAllPrimitivesCallback(p1 bool, p2 int8, p3 uint16, p4 int8, p5 int16, p6 int32, p7 int64, p8 uint8, p9 uint16, p10 uint32, p11 uint64, p12 uintptr, p13 float32, p14 float64) int64 {
	C_p1 := C.bool(p1)
	C_p2 := C.char(p2)
	C_p3 := C.uint16_t(p3)
	C_p4 := C.int8_t(p4)
	C_p5 := C.int16_t(p5)
	C_p6 := C.int32_t(p6)
	C_p7 := C.int64_t(p7)
	C_p8 := C.uint8_t(p8)
	C_p9 := C.uint16_t(p9)
	C_p10 := C.uint32_t(p10)
	C_p11 := C.uint64_t(p11)
	C_p12 := C.uintptr_t(p12)
	C_p13 := C.float(p13)
	C_p14 := C.double(p14)

	result := int64(C.ParamAllPrimitivesCallback(C_p1, C_p2, C_p3, C_p4, C_p5, C_p6, C_p7, C_p8, C_p9, C_p10, C_p11, C_p12, C_p13, C_p14))
	return result
}

/*
func CallFuncVoidCallback(func func)  {
	C_func := (func)

	C.CallFuncVoidCallback(C_func)
}

func CallFuncBoolCallback(func func) bool {
	C_func := (func)

	result := bool(C.CallFuncBoolCallback(C_func))
	return result
}

func CallFuncChar8Callback(func func) int8 {
	C_func := (func)

	result := int8(C.CallFuncChar8Callback(C_func))
	return result
}

func CallFuncChar16Callback(func func) uint16 {
	C_func := (func)

	result := uint16(C.CallFuncChar16Callback(C_func))
	return result
}

func CallFuncInt8Callback(func func) int8 {
	C_func := (func)

	result := int8(C.CallFuncInt8Callback(C_func))
	return result
}

func CallFuncInt16Callback(func func) int16 {
	C_func := (func)

	result := int16(C.CallFuncInt16Callback(C_func))
	return result
}

func CallFuncInt32Callback(func func) int32 {
	C_func := (func)

	result := int32(C.CallFuncInt32Callback(C_func))
	return result
}

func CallFuncInt64Callback(func func) int64 {
	C_func := (func)

	result := int64(C.CallFuncInt64Callback(C_func))
	return result
}

func CallFuncUInt8Callback(func func) uint8 {
	C_func := (func)

	result := uint8(C.CallFuncUInt8Callback(C_func))
	return result
}

func CallFuncUInt16Callback(func func) uint16 {
	C_func := (func)

	result := uint16(C.CallFuncUInt16Callback(C_func))
	return result
}

func CallFuncUInt32Callback(func func) uint32 {
	C_func := (func)

	result := uint32(C.CallFuncUInt32Callback(C_func))
	return result
}

func CallFuncUInt64Callback(func func) uint64 {
	C_func := (func)

	result := uint64(C.CallFuncUInt64Callback(C_func))
	return result
}

func CallFuncPtrCallback(func func) uintptr {
	C_func := (func)

	result := uintptr(C.CallFuncPtrCallback(C_func))
	return result
}

func CallFuncFloatCallback(func func) float32 {
	C_func := (func)

	result := float32(C.CallFuncFloatCallback(C_func))
	return result
}

func CallFuncDoubleCallback(func func) float64 {
	C_func := (func)

	result := float64(C.CallFuncDoubleCallback(C_func))
	return result
}

func CallFuncStringCallback(func func) string {
	C_func := (func)

	C_output := C.CallFuncStringCallback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFuncFunctionCallback(func func) uintptr {
	C_func := (func)

	result := uintptr(C.CallFuncFunctionCallback(C_func))
	return result
}

func CallFuncBoolVectorCallback(func func) []bool {
	C_func := (func)

	C_output := C.CallFuncBoolVectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.BOOL)
	P_output := C.Plugify_GetVectorData(&C_output, C.BOOL)
	output := make([]bool, L_output)
	for i := range output {
		output[i] = *(*bool)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_bool)))
	}

	C.Plugify_DestroyVector(&C_output, C.BOOL)
	return output
}

func CallFuncChar8VectorCallback(func func) []int8 {
	C_func := (func)

	C_output := C.CallFuncChar8VectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.CHAR8)
	P_output := C.Plugify_GetVectorData(&C_output, C.CHAR8)
	output := make([]int8, L_output)
	for i := range output {
		output[i] = *(*int8)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_char)))
	}

	C.Plugify_DestroyVector(&C_output, C.CHAR8)
	return output
}

func CallFuncChar16VectorCallback(func func) []uint16 {
	C_func := (func)

	C_output := C.CallFuncChar16VectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.CHAR16)
	P_output := C.Plugify_GetVectorData(&C_output, C.CHAR16)
	output := make([]uint16, L_output)
	for i := range output {
		output[i] = *(*uint16)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uint16_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.CHAR16)
	return output
}

func CallFuncInt8VectorCallback(func func) []int8 {
	C_func := (func)

	C_output := C.CallFuncInt8VectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.INT8)
	P_output := C.Plugify_GetVectorData(&C_output, C.INT8)
	output := make([]int8, L_output)
	for i := range output {
		output[i] = *(*int8)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_int8_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.INT8)
	return output
}

func CallFuncInt16VectorCallback(func func) []int16 {
	C_func := (func)

	C_output := C.CallFuncInt16VectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.INT16)
	P_output := C.Plugify_GetVectorData(&C_output, C.INT16)
	output := make([]int16, L_output)
	for i := range output {
		output[i] = *(*int16)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_int16_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.INT16)
	return output
}

func CallFuncInt32VectorCallback(func func) []int32 {
	C_func := (func)

	C_output := C.CallFuncInt32VectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.INT32)
	P_output := C.Plugify_GetVectorData(&C_output, C.INT32)
	output := make([]int32, L_output)
	for i := range output {
		output[i] = *(*int32)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_int32_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.INT32)
	return output
}

func CallFuncInt64VectorCallback(func func) []int64 {
	C_func := (func)

	C_output := C.CallFuncInt64VectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.INT64)
	P_output := C.Plugify_GetVectorData(&C_output, C.INT64)
	output := make([]int64, L_output)
	for i := range output {
		output[i] = *(*int64)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_int64_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.INT64)
	return output
}

func CallFuncUInt8VectorCallback(func func) []uint8 {
	C_func := (func)

	C_output := C.CallFuncUInt8VectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.UINT8)
	P_output := C.Plugify_GetVectorData(&C_output, C.UINT8)
	output := make([]uint8, L_output)
	for i := range output {
		output[i] = *(*uint8)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uint8_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.UINT8)
	return output
}

func CallFuncUInt16VectorCallback(func func) []uint16 {
	C_func := (func)

	C_output := C.CallFuncUInt16VectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.UINT16)
	P_output := C.Plugify_GetVectorData(&C_output, C.UINT16)
	output := make([]uint16, L_output)
	for i := range output {
		output[i] = *(*uint16)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uint16_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.UINT16)
	return output
}

func CallFuncUInt32VectorCallback(func func) []uint32 {
	C_func := (func)

	C_output := C.CallFuncUInt32VectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.UINT32)
	P_output := C.Plugify_GetVectorData(&C_output, C.UINT32)
	output := make([]uint32, L_output)
	for i := range output {
		output[i] = *(*uint32)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uint32_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.UINT32)
	return output
}

func CallFuncUInt64VectorCallback(func func) []uint64 {
	C_func := (func)

	C_output := C.CallFuncUInt64VectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.UINT64)
	P_output := C.Plugify_GetVectorData(&C_output, C.UINT64)
	output := make([]uint64, L_output)
	for i := range output {
		output[i] = *(*uint64)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uint64_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.UINT64)
	return output
}

func CallFuncPtrVectorCallback(func func) []uintptr {
	C_func := (func)

	C_output := C.CallFuncPtrVectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.POINTER)
	P_output := C.Plugify_GetVectorData(&C_output, C.POINTER)
	output := make([]uintptr, L_output)
	for i := range output {
		output[i] = *(*uintptr)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uintptr_t)))
	}

	C.Plugify_DestroyVector(&C_output, C.POINTER)
	return output
}

func CallFuncFloatVectorCallback(func func) []float32 {
	C_func := (func)

	C_output := C.CallFuncFloatVectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.FLOAT)
	P_output := C.Plugify_GetVectorData(&C_output, C.FLOAT)
	output := make([]float32, L_output)
	for i := range output {
		output[i] = *(*float32)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_float)))
	}

	C.Plugify_DestroyVector(&C_output, C.FLOAT)
	return output
}

func CallFuncStringVectorCallback(func func) []string {
	C_func := (func)

	C_output := C.CallFuncStringVectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.STRING)
	P_output := C.Plugify_GetVectorData(&C_output, C.STRING)
	output := make([]string, L_output)
	for i := range output {
		output[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uintptr_t))))
	}

	C.Plugify_DeleteVectorDataCStr(P_output)
	C.Plugify_DestroyVector(&C_output, C.STRING)
	return output
}

func CallFuncDoubleVectorCallback(func func) []float64 {
	C_func := (func)

	C_output := C.CallFuncDoubleVectorCallback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.DOUBLE)
	P_output := C.Plugify_GetVectorData(&C_output, C.DOUBLE)
	output := make([]float64, L_output)
	for i := range output {
		output[i] = *(*float64)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_double)))
	}

	C.Plugify_DestroyVector(&C_output, C.DOUBLE)
	return output
}

func CallFuncVec2Callback(func func) Vector2 {
	C_func := (func)

	C_result := C.CallFuncVec2Callback(C_func)
	return *(*Vector2)(unsafe.Pointer(&C_result))
}

func CallFuncVec3Callback(func func) Vector3 {
	C_func := (func)

	C_result := C.CallFuncVec3Callback(C_func)
	return *(*Vector3)(unsafe.Pointer(&C_result))
}

func CallFuncVec4Callback(func func) Vector4 {
	C_func := (func)

	C_result := C.CallFuncVec4Callback(C_func)
	return *(*Vector4)(unsafe.Pointer(&C_result))
}

func CallFuncMat4x4Callback(func func) Matrix4x4 {
	C_func := (func)

	C_result := C.CallFuncMat4x4Callback(C_func)
	return *(*Matrix4x4)(unsafe.Pointer(&C_result))
}

func CallFunc1Callback(func func) int32 {
	C_func := (func)

	result := int32(C.CallFunc1Callback(C_func))
	return result
}

func CallFunc2Callback(func func) int8 {
	C_func := (func)

	result := int8(C.CallFunc2Callback(C_func))
	return result
}

func CallFunc3Callback(func func)  {
	C_func := (func)

	C.CallFunc3Callback(C_func)
}

func CallFunc4Callback(func func) Vector4 {
	C_func := (func)

	C_result := C.CallFunc4Callback(C_func)
	return *(*Vector4)(unsafe.Pointer(&C_result))
}

func CallFunc5Callback(func func) bool {
	C_func := (func)

	result := bool(C.CallFunc5Callback(C_func))
	return result
}

func CallFunc6Callback(func func) int64 {
	C_func := (func)

	result := int64(C.CallFunc6Callback(C_func))
	return result
}

func CallFunc7Callback(func func) float64 {
	C_func := (func)

	result := float64(C.CallFunc7Callback(C_func))
	return result
}

func CallFunc8Callback(func func) Matrix4x4 {
	C_func := (func)

	C_result := C.CallFunc8Callback(C_func)
	return *(*Matrix4x4)(unsafe.Pointer(&C_result))
}

func CallFunc9Callback(func func)  {
	C_func := (func)

	C.CallFunc9Callback(C_func)
}

func CallFunc10Callback(func func) uint32 {
	C_func := (func)

	result := uint32(C.CallFunc10Callback(C_func))
	return result
}

func CallFunc11Callback(func func) uintptr {
	C_func := (func)

	result := uintptr(C.CallFunc11Callback(C_func))
	return result
}

func CallFunc12Callback(func func) bool {
	C_func := (func)

	result := bool(C.CallFunc12Callback(C_func))
	return result
}

func CallFunc13Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc13Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc14Callback(func func) []string {
	C_func := (func)

	C_output := C.CallFunc14Callback(C_func)

	L_output := C.Plugify_GetVectorSize(&C_output, C.STRING)
	P_output := C.Plugify_GetVectorData(&C_output, C.STRING)
	output := make([]string, L_output)
	for i := range output {
		output[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uintptr_t))))
	}

	C.Plugify_DeleteVectorDataCStr(P_output)
	C.Plugify_DestroyVector(&C_output, C.STRING)
	return output
}

func CallFunc15Callback(func func) int16 {
	C_func := (func)

	result := int16(C.CallFunc15Callback(C_func))
	return result
}

func CallFunc16Callback(func func) uintptr {
	C_func := (func)

	result := uintptr(C.CallFunc16Callback(C_func))
	return result
}

func CallFunc17Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc17Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc18Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc18Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc19Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc19Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc20Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc20Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc21Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc21Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc22Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc22Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc23Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc23Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc24Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc24Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc25Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc25Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc26Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc26Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc27Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc27Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc28Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc28Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc29Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc29Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc30Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc30Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc31Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc31Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}

func CallFunc32Callback(func func) string {
	C_func := (func)

	C_output := C.CallFunc32Callback(C_func)

	P_output := C.Plugify_GetStringData(&C_output)
	output := C.GoString(P_output)
	C.Plugify_DestroyString(&C_output)
	return output
}
*/
