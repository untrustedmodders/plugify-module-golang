#!/usr/bin/python3
import sys
import argparse
import os
import json
from enum import Enum


VAL_TYPES_MAP = {
    'void': 'void',
    'bool': 'bool',
    'char8': 'char',
    'char16': 'uint16_t',
    'int8': 'int8_t',
    'int16': 'int16_t',
    'int32': 'int32_t',
    'int64': 'int64_t',
    'uint8': 'uint8_t',
    'uint16': 'uint16_t',
    'uint32': 'uint32_t',
    'uint64': 'uint64_t',
    'ptr64': 'uintptr_t',
    'float': 'float',
    'double': 'double',
    'function': 'delegate',
    'string': 'String*',
    'bool*': 'Vector*',
    'char8*': 'Vector*',
    'char16*': 'Vector*',
    'int8*': 'Vector*',
    'int16*': 'Vector*',
    'int32*': 'Vector*',
    'int64*': 'Vector*',
    'uint8*': 'Vector*',
    'uint16*': 'Vector*',
    'uint32*': 'Vector*',
    'uint64*': 'Vector*',
    'ptr64*': 'Vector*',
    'float*': 'Vector*',
    'double*': 'Vector*',
    'string*': 'Vector*',
    'vec2': 'Vector2*',
    'vec3': 'Vector3*',
    'vec4': 'Vector4*',
    'mat4x4': 'Matrix4x4*'
}

REF_TYPES_MAP = {
    'void': 'void',
    'bool': 'bool*',
    'char8': 'char*',
    'char16': 'uint16_t*',
    'int8': 'int8_t*',
    'int16': 'int16_t*',
    'int32': 'int32_t*',
    'int64': 'int64_t*',
    'uint8': 'uint8_t*',
    'uint16': 'uint16_t*',
    'uint32': 'uint32_t*',
    'uint64': 'uint64_t*',
    'ptr64': 'uintptr_t*',
    'float': 'float*',
    'double': 'double*',
    'function': 'delegate',
    'string': 'String*',
    'bool*': 'Vector*',
    'char8*': 'Vector*',
    'char16*': 'Vector*',
    'int8*': 'Vector*',
    'int16*': 'Vector*',
    'int32*': 'Vector*',
    'int64*': 'Vector*',
    'uint8*': 'Vector*',
    'uint16*': 'Vector*',
    'uint32*': 'Vector*',
    'uint64*': 'Vector*',
    'ptr64*': 'Vector*',
    'float*': 'Vector*',
    'double*': 'Vector*',
    'string*': 'Vector*',
    'vec2': 'Vector2*',
    'vec3': 'Vector3*',
    'vec4': 'Vector4*',
    'mat4x4': 'Matrix4x4*'
}

RET_TYPES_MAP = {
    'void': 'void',
    'bool': 'bool',
    'char8': 'char',
    'char16': 'uint16_t',
    'int8': 'int8_t',
    'int16': 'int16_t',
    'int32': 'int32_t',
    'int64': 'int64_t',
    'uint8': 'uint8_t',
    'uint16': 'uint16_t',
    'uint32': 'uint32_t',
    'uint64': 'uint64_t',
    'ptr64': 'uintptr_t',
    'float': 'float',
    'double': 'double',
    'function': 'delegate',
    'string': 'String',
    'bool*': 'Vector',
    'char8*': 'Vector',
    'char16*': 'Vector',
    'int8*': 'Vector',
    'int16*': 'Vector',
    'int32*': 'Vector',
    'int64*': 'Vector',
    'uint8*': 'Vector',
    'uint16*': 'Vector',
    'uint32*': 'Vector',
    'uint64*': 'Vector',
    'ptr64*': 'Vector',
    'float*': 'Vector',
    'double*': 'Vector',
    'string*': 'Vector',
    'vec2': 'Vector2',
    'vec3': 'Vector3',
    'vec4': 'Vector4',
    'mat4x4': 'Matrix4x4'
}

VAL_GOTYPES_MAP = {
    'void': '',
    'bool': 'bool',
    'char8': 'int8',
    'char16': 'uint16',
    'int8': 'int8',
    'int16': 'int16',
    'int32': 'int32',
    'int64': 'int64',
    'uint8': 'uint8',
    'uint16': 'uint16',
    'uint32': 'uint32',
    'uint64': 'uint64',
    'ptr64': 'uintptr',
    'float': 'float32',
    'double': 'float64',
    'function': 'func',
    'string': 'string',
    'bool*': '[]bool',
    'char8*': '[]int8',
    'char16*': '[]uint16',
    'int8*': '[]int8',
    'int16*': '[]int16',
    'int32*': '[]int32',
    'int64*': '[]int64',
    'uint8*': '[]uint8',
    'uint16*': '[]uint16',
    'uint32*': '[]uint32',
    'uint64*': '[]uint64',
    'ptr64*': '[]uintptr',
    'float*': '[]float32',
    'double*': '[]float64',
    'string*': '[]string',
    'vec2': 'Vector2',
    'vec3': 'Vector3',
    'vec4': 'Vector4',
    'mat4x4': 'Matrix4x4'
}

REF_GOTYPES_MAP = {
    'void': '',
    'bool': '*bool',
    'char8': '*int8',
    'char16': '*uint16',
    'int8': '*int8',
    'int16': '*int16',
    'int32': '*int32',
    'int64': '*int64',
    'uint8': '*uint8',
    'uint16': '*uint16',
    'uint32': '*uint32',
    'uint64': '*uint64',
    'ptr64': '*uintptr',
    'float': '*float32',
    'double': '*float64',
    'function': '*func',
    'string': '*string',
    'bool*': '*[]bool',
    'char8*': '*[]int8',
    'char16*': '*[]uint16',
    'int8*': '*[]int8',
    'int16*': '*[]int16',
    'int32*': '*[]int32',
    'int64*': '*[]int64',
    'uint8*': '*[]uint8',
    'uint16*': '*[]uint16',
    'uint32*': '*[]uint32',
    'uint64*': '*[]uint64',
    'ptr64*': '*[]uintptr',
    'float*': '*[]float32',
    'double*': '*[]float64',
    'string*': '*[]string',
    'vec2': '*Vector2',
    'vec3': '*Vector3',
    'vec4': '*Vector4',
    'mat4x4': '*Matrix4x4'
}

RET_GOTYPES_MAP = {
    'void': '',
    'bool': 'bool',
    'char8': 'int8',
    'char16': 'uint16',
    'int8': 'int8',
    'int16': 'int16',
    'int32': 'int32',
    'int64': 'int64',
    'uint8': 'uint8',
    'uint16': 'uint16',
    'uint32': 'uint32',
    'uint64': 'uint64',
    'ptr64': 'uintptr',
    'float': 'float32',
    'double': 'float64',
    'function': 'func',
    'string': 'string',
    'bool*': '[]bool',
    'char8*': '[]int8',
    'char16*': '[]uint16',
    'int8*': '[]int8',
    'int16*': '[]int16',
    'int32*': '[]int32',
    'int64*': '[]int64',
    'uint8*': '[]uint8',
    'uint16*': '[]uint16',
    'uint32*': '[]uint32',
    'uint64*': '[]uint64',
    'ptr64*': '[]uintptr',
    'float*': '[]float32',
    'double*': '[]float64',
    'string*': '[]string',
    'vec2': 'Vector2',
    'vec3': 'Vector3',
    'vec4': 'Vector4',
    'mat4x4': 'Matrix4x4'
}

VAL_GOTYPESCAST_MAP = {
    'void': '',
    'bool': 'C.bool',
    'char8': 'C.char',
    'char16': 'C.uint16_t',
    'int8': 'C.int8_t',
    'int16': 'C.int16_t',
    'int32': 'C.int32_t',
    'int64': 'C.int64_t',
    'uint8': 'C.uint8_t',
    'uint16': 'C.uint16_t',
    'uint32': 'C.uint32_t',
    'uint64': 'C.uint64_t',
    'ptr64': 'C.uintptr_t',
    'float': 'C.float',
    'double': 'C.double',
    'function': '',
    'string': 'C.Plugify_ConstructString',
    'bool*': 'C.Plugify_ConstructVectorC.BOOL',
    'char8*': 'C.Plugify_ConstructVectorC.CHAR8',
    'char16*': 'C.Plugify_ConstructVectorC.CHAR16',
    'int8*': 'C.Plugify_ConstructVectorC.INT8',
    'int16*': 'C.Plugify_ConstructVectorC.INT16',
    'int32*': 'C.Plugify_ConstructVectorC.INT32',
    'int64*': 'C.Plugify_ConstructVectorC.INT64',
    'uint8*': 'C.Plugify_ConstructVectorC.UINT8',
    'uint16*': 'C.Plugify_ConstructVectorC.UINT16',
    'uint32*': 'C.Plugify_ConstructVectorC.UINT32',
    'uint64*': 'C.Plugify_ConstructVectorC.UINT64',
    'ptr64*': 'C.Plugify_ConstructVectorC.POINTER',
    'float*': 'C.Plugify_ConstructVectorC.FLOAT',
    'double*': 'C.Plugify_ConstructVectorC.DOUBLE',
    'string*': 'C.Plugify_ConstructVectorC.STRING',
    'vec2': 'C.Vector2',
    'vec3': 'C.Vector3',
    'vec4': 'C.Vector4',
    'mat4x4': 'C.Matrix4x4'
}


DEL_GOTYPESCAST_MAP = {
    'void': '',
    'bool': '',
    'char8': '',
    'char16': '',
    'int8': '',
    'int16': '',
    'int32': '',
    'int64': '',
    'uint8': '',
    'uint16': '',
    'uint32': '',
    'uint64': '',
    'ptr64': '',
    'float': '',
    'double': '',
    'function': '',
    'string': 'C.Plugify_DestroyString',
    'bool*': 'C.Plugify_DestroyVectorC.BOOL',
    'char8*': 'C.Plugify_DestroyVectorC.CHAR8',
    'char16*': 'C.Plugify_DestroyVectorC.CHAR16',
    'int8*': 'C.Plugify_DestroyVectorC.INT8',
    'int16*': 'C.Plugify_DestroyVectorC.INT16',
    'int32*': 'C.Plugify_DestroyVectorC.INT32',
    'int64*': 'C.Plugify_DestroyVectorC.INT64',
    'uint8*': 'C.Plugify_DestroyVectorC.UINT8',
    'uint16*': 'C.Plugify_DestroyVectorC.UINT16',
    'uint32*': 'C.Plugify_DestroyVectorC.UINT32',
    'uint64*': 'C.Plugify_DestroyVectorC.UINT64',
    'ptr64*': 'C.Plugify_DestroyVectorC.POINTER',
    'float*': 'C.Plugify_DestroyVectorC.FLOAT',
    'double*': 'C.Plugify_DestroyVectorC.DOUBLE',
    'string*': 'C.Plugify_DestroyVectorC.STRING',
    'vec2': '',
    'vec3': '',
    'vec4': '',
    'mat4x4': ''
}

ASS_GOTYPESCAST_MAP = {
    'void': '',
    'bool': 'bool',
    'char8': 'byte',
    'char16': 'uint16',
    'int8': 'int8',
    'int16': 'int16',
    'int32': 'int32',
    'int64': 'int64',
    'uint8': 'uint8',
    'uint16': 'uint16',
    'uint32': 'uint32',
    'uint64': 'uint64',
    'ptr64': 'uintptr',
    'float': 'float32',
    'double': 'float64',
    'function': '',
    'string': 'C.GoString',
    'bool*': 'C.bool',
    'char8*': 'C.char',
    'char16*': 'C.uint16_t',
    'int8*': 'C.int8_t',
    'int16*': 'C.int16_t',
    'int32*': 'C.int32_t',
    'int64*': 'C.int64_t',
    'uint8*': 'C.uint8_t',
    'uint16*': 'C.uint16_t',
    'uint32*': 'C.uint32_t',
    'uint64*': 'C.uint64_t',
    'ptr64*': 'C.uintptr_t',
    'float*': 'C.float',
    'double*': 'C.double',
    'string*': 'string',
    'vec2': 'C.Vector2',
    'vec3': 'C.Vector3',
    'vec4': 'C.Vector4',
    'mat4x4': 'C.Matrix4x4'
}

REF_GOTYPESCAST_MAP = {
    'void': '',
    'bool': '',
    'char8': '',
    'char16': '',
    'int8': '',
    'int16': '',
    'int32': '',
    'int64': '',
    'uint8': '',
    'uint16': '',
    'uint32': '',
    'uint64': '',
    'ptr64': '',
    'float': '',
    'double': '',
    'function': '',
    'string': '&',
    'bool*': '&',
    'char8*': '&',
    'char16*': '&',
    'int8*': '&',
    'int16*': '&',
    'int32*': '&',
    'int64*': '&',
    'uint8*': '&',
    'uint16*': '&',
    'uint32*': '&',
    'uint64*': '&',
    'ptr64*': '&',
    'float*': '&',
    'double*': '&',
    'string*': '&',
    'vec2': '&',
    'vec3': '&',
    'vec4': '&',
    'mat4x4': '&'
}

def validate_manifest(pplugin):
    parse_errors = []
    methods = pplugin.get('exportedMethods')
    if type(methods) is list:
        for i, method in enumerate(methods):
            if type(method) is dict:
                if type(method.get('type')) is str:
                    parse_errors += [f'root.exportedMethods[{i}].type not string']
            else:
                parse_errors += [f'root.exportedMethods[{i}] not object']
    else:
        parse_errors += ['root.exportedMethods not array']
    return parse_errors


def convert_type(type_name, is_ref: False, is_ret: False):
    if is_ret:
        return RET_TYPES_MAP.get(type_name, 'int')
    elif is_ref:
        return REF_TYPES_MAP.get(type_name, 'int')
    else:
        return VAL_TYPES_MAP.get(type_name, 'int')


def convert_gotype(type_name, is_ref: False, is_ret: False):
    if is_ret:
        return RET_GOTYPES_MAP.get(type_name, 'int')
    elif is_ref:
        return REF_GOTYPES_MAP.get(type_name, 'int')
    else:
        return VAL_GOTYPES_MAP.get(type_name, 'int')


def is_obj_return(type_name):
    return '*' in type_name or type_name == 'string'      


def is_pod_return(type_name):
    return 'vec' in type_name or 'mat' in type_name       


class ParamGen(Enum):
    Types = 1
    Names = 2
    TypesNames = 3


def gen_params(method, param_gen: ParamGen):
    def gen_param(param):
        if param_gen == ParamGen.Types:
            type = convert_type(param['type'], 'ref' in param and param['ref'] is True, False)
            return type
        if param_gen == ParamGen.Names:
            return param['name']
        type = convert_type(param['type'], 'ref' in param and param['ref'] is True, False)
        return f'{type} {param['name']}'

    def gen_return(param):
        if param_gen == ParamGen.Types:
            type = convert_type(param['type'], 'ref' in param and param['ref'] is True, False)
            return type
        if param_gen == ParamGen.Names:
            return 'output'
        type = convert_type(param['type'], 'ref' in param and param['ref'] is True, False)
        return f'{type} output'

    output_string = ''
    ret_type = method['retType']
    if method["paramTypes"]:
        it = iter(method["paramTypes"])
        output_string += gen_param(next(it))
        for p in it:
            output_string += f', {gen_param(p)}'
    return output_string


def gen_goparams(method, param_gen: ParamGen):
    def gen_param(param):
        type = convert_gotype(param['type'], 'ref' in param and param['ref'] is True, False)
        if param_gen == ParamGen.Types:
            return type
        if param_gen == ParamGen.Names:
            ref_type = REF_GOTYPESCAST_MAP.get(param['type'], 'int')
            if ref_type == '&':
                return f'&C_{param['name']}'
            elif ref_type != '':
                return f'{ref_type}(C_{param['name']})'
            else:
                return f'C_{param['name']}'
        return f'{param['name']} {type}'

    def gen_return(param):
        type = convert_gotype(param['type'], 'ref' in param and param['ref'] is True, False)
        if param_gen == ParamGen.Types:
            return type
        if param_gen == ParamGen.Names:
            ref_type = REF_GOTYPESCAST_MAP.get(param['type'], 'int')
            if ref_type == '&':
                return '&C_output'
            elif ref_type != '':
                return f'{ref_type}(C_output)'
            else:
                return 'C_output'
        return ''

    output_string = ''
    ret_type = method['retType']
    if method["paramTypes"]:
        it = iter(method["paramTypes"])
        output_string += gen_param(next(it))
        for p in it:
            output_string += f', {gen_param(p)}'
    return output_string


def gen_goparamscast(method):
    def gen_param(param):
        type = VAL_GOTYPESCAST_MAP.get(param['type'], 'int')
        if 'ConstructVector' in type:
            if 'ref' in param and param['ref'] is True:
                output = f'var A_{param['name']} unsafe.Pointer\n'
                output += f'\tS_{param['name']} := len(*{param['name']})\n'
                output += f'\tif S_{param['name']} > 0 {{\n'
                output += f'\t\tA_{param['name']} = unsafe.Pointer(&(*{param['name']})[0])\n'
                output += '\t} else {\n'
                output += f'\t\tA_{param['name']} = nil\n'
                output += '\t}\n'
                output += f'\tC_{param['name']} := {type[:25]}(A_{param['name']}, C.ptrdiff_t(S_{param['name']}), {type[25:]})'
                return output;
            else:
                output = f'var A_{param['name']} unsafe.Pointer\n'
                output += f'\tS_{param['name']} := len({param['name']})\n'
                output += f'\tif S_{param['name']} > 0 {{\n'
                output += f'\t\tA_{param['name']} = unsafe.Pointer(&{param['name']}[0])\n'
                output += '\t} else {\n'
                output += f'\t\tA_{param['name']} = nil\n'
                output += '\t}\n'
                output += f'\tC_{param['name']} := {type[:25]}(A_{param['name']}, C.ptrdiff_t(S_{param['name']}), {type[25:]})'
                return output;
        elif 'ConstructString' in type:
            if 'ref' in param and param['ref'] is True:
                return f'C_{param['name']} := {type}(*{param['name']})'
            else:
                return f'C_{param['name']} := {type}({param['name']})'
        elif 'C.Vector' in type or 'C.Matrix' in type:
            if 'ref' in param and param['ref'] is True:
                return f'C_{param['name']} := *(*{type})(unsafe.Pointer({param['name']}))'
            else:
                return f'C_{param['name']} := *(*{type})(unsafe.Pointer(&{param['name']}))'
        elif type == 'C.uintptr_t':
            if 'ref' in param and param['ref'] is True:
                return f'C_{param['name']} := (*{type})(unsafe.Pointer({param['name']}))'
            else:
                return f'C_{param['name']} := {type}({param['name']})'
        else:
            if 'ref' in param and param['ref'] is True:
                return f'C_{param['name']} := (*{type})({param['name']})'
            else:
                return f'C_{param['name']} := {type}({param['name']})'

    def gen_return(param):
        return ''

    output_string = ''
    ret_type = method['retType']
    if method["paramTypes"]:
        it = iter(method["paramTypes"])
        output_string += f'{gen_param(next(it))}\n'
        for p in it:
            output_string += f'\t{gen_param(p)}\n'
    return output_string


def gen_goparamscast_assign(method):
    def gen_param(param):
        if 'ref' in param and param['ref'] is True:
            type = ASS_GOTYPESCAST_MAP.get(param['type'], 'int')
            if type == 'C.GoString':
                output = f'P_{param['name']} := C.Plugify_GetStringData(&C_{param['name']})\n'
                output += f'\t*{param['name']} = {type}(P_{param['name']})'
                return output
            elif 'C.Vector' in type or 'C.Matrix' in type:
                return f'*{param['name']} = *(*{type[2:]})(unsafe.Pointer(&C_{param['name']}))'
            elif 'C.' in type or type[0] == '*' or type == 'string':
                name = f'&C_{param['name']}'
                mode = ''
                if param['type'] == 'string*':
                    mode = 'CStr'
                go_type = VAL_GOTYPES_MAP.get(param['type'], 'int')[2:]

                val_type = VAL_GOTYPESCAST_MAP.get(param['type'], 'int')[25:]
                output = f'L_{param['name']} := C.Plugify_GetVectorSize({name}, {val_type})\n'
                output += f'\tP_{param['name']} := C.Plugify_GetVectorData({name}, {val_type})\n'
                #output += f'\t*{param['name']} = unsafe.Slice((*{go_type})(P_{param['name']}), L_{param['name']})'
                output += f'\t*{param['name']} = make([]{go_type}, L_{param['name']})\n'
                output += f'\tfor i := range (*{param['name']})' + " {\n"
                if mode == 'CStr':
                    output += f'\t\t(*{param['name']})[i] = C.GoString(C.Plugify_GetStringData((*C.String)(unsafe.Pointer(uintptr(P_{param['name']}) + uintptr(i * C.sizeof_uintptr_t)))))\n'
                else:
                    output += f'\t\t(*{param['name']})[i] = *(*{go_type})(unsafe.Pointer(uintptr(P_{param['name']}) + uintptr(i * C.sizeof_{type[2:]})))\n'
                output += "\t}"
                return output
            else:
                return ''#f'*{param['name']} = {type}(C_{param['name']})'
        else:
            return ''

    def gen_return(param):
        type = ASS_GOTYPESCAST_MAP.get(param['type'], 'int')
        if type == 'C.GoString':
            output = f'P_output := C.Plugify_GetStringData(&C_output)\n'
            output += f'\toutput := {type}(P_output)'
            return output
        elif 'C.Vector' in type or 'C.Matrix' in type:
            return f'output := *(*{type[2:]})(unsafe.Pointer(&C_output))'
        elif 'C.' in type or type[0] == '*' or type == 'string':
            name = f'&C_output'
            mode = ''
            if param['type'] == 'string*':
                mode = 'CStr'
            go_type = VAL_GOTYPES_MAP.get(param['type'], 'int')[2:]
            val_type = VAL_GOTYPESCAST_MAP.get(param['type'], 'int')[25:]
            output = f'L_output := C.Plugify_GetVectorSize({name}, {val_type})\n'
            output += f'\tP_output := C.Plugify_GetVectorData({name}, {val_type})\n'
            #output += f'\toutput := unsafe.Slice((*{go_type})(P_output), L_output)'
            output += f'\toutput := make([]{go_type}, L_output)\n'
            output += f'\tfor i := range output' + " {\n"
            if mode == 'CStr':
                output += f'\t\toutput[i] = C.GoString(C.Plugify_GetStringData((*C.String)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uintptr_t)))))\n'
            else:
                output += f'\t\toutput[i] = *(*{go_type})(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_{type[2:]})))\n'
            output += "\t}\n"
            return output
        else:
            return f'output := {type}(C_output)'

    output_string = ''
    ret_type = method['retType']
    is_obj_ret = is_obj_return(ret_type["type"])
    if is_obj_ret:
        output_string += f'\t{gen_return(ret_type)}'
    if method["paramTypes"]:
        it = iter(method["paramTypes"])
        param = gen_param(next(it))
        if param != '':
            output_string += f'\t{param}\n'
        for p in it:
            param = gen_param(p)
            if param != '':
                output_string += f'\t{param}\n'
    return output_string


def gen_goparamscast_cleanup(method):
    def gen_param(param):
        type = DEL_GOTYPESCAST_MAP.get(param['type'], 'int')
        if type == '':
            return type
        elif 'DestroyVector' in type:
            return f'{type[:23]}(&C_{param['name']}, {type[23:]})'
        else:
            return f'{type}(&C_{param['name']})'

    def gen_return(param):
        type = DEL_GOTYPESCAST_MAP.get(param['type'], 'int')
        if type == '':
            return type
        elif 'DestroyVector' in type:
            return f'{type[:23]}(&C_output, {type[23:]})'
        else:
            return f'{type}(&C_output)'

    output_string = ''
    ret_type = method['retType']
    is_obj_ret = is_obj_return(ret_type["type"])
    if is_obj_ret:
        output_string += f'\t{gen_return(ret_type)}'
    if method["paramTypes"]:
        it = iter(method["paramTypes"])
        param = gen_param(next(it))
        if param != '':
            output_string += f'\t{param}\n'
        for p in it:
            param = gen_param(p)
            if param != '':
                output_string += f'\t{param}\n'
    return output_string


def main(manifest_path, output_dir, override):
    if not os.path.isfile(manifest_path):
        print(f'Manifest file not exists {manifest_path}')
        return 1
    if not os.path.isdir(output_dir):
        print(f'Output folder not exists {output_dir}')
        return 1

    plugin_name = os.path.splitext(os.path.basename(manifest_path))[0]

    header_dir = os.path.join(output_dir, plugin_name)
    if not os.path.exists(header_dir):
        os.makedirs(header_dir, exist_ok=True)
    header_file = os.path.join(header_dir, f'{plugin_name}.go')
    if os.path.isfile(header_file) and not override:
        print(f'Already exists {header_file}')
        return 1

    with open(manifest_path, 'r', encoding='utf-8') as fd:
        pplugin = json.load(fd)

    parse_errors = validate_manifest(pplugin)
    if parse_errors:
        print('Parse fail:')
        for error in parse_errors:
            print(f'  {error}')
        return 1

    content = ('#pragma once\n'
               '#include <stdlib.h>\n'
               '#include <stdint.h>\n'
               '#include <stdbool.h>\n\n'
               'enum DataType {\n'
               '\tBOOL,\n'
               '\tCHAR8,\n'
               '\tCHAR16,\n'
               '\tINT8,\n'
               '\tINT16,\n'
               '\tINT32,\n'
               '\tINT64,\n'
               '\tUINT8,\n'
               '\tUINT16,\n'
               '\tUINT32,\n'
               '\tUINT64,\n'
               '\tPOINTER,\n'
               '\tFLOAT,\n'
               '\tDOUBLE,\n'
               '\tSTRING\n'
               '};\n\n'
               'typedef struct { char* data; size_t size; size_t cap; } String;\n'
               'typedef struct { void* begin; void* end; void* capacity; } Vector;\n'
               '\n'
               'extern void* Plugify_GetMethodPtr(const char* methodName);\n'
               'extern void Plugify_GetMethodPtr2(const char* methodName, void** addressDest);\n'
               '\n'
               'extern String Plugify_ConstructString(_GoString_ source);\n'
               'extern void Plugify_DestroyString(String* string);\n'
               'extern const char* Plugify_GetStringData(String* string);\n'
               'extern ptrdiff_t Plugify_GetStringLength(String* string);\n'
               'extern void Plugify_AssignString(String* string, _GoString_ source);\n'
               '\n'
               'extern Vector Plugify_ConstructVector(void* arr, ptrdiff_t len, enum DataType type);\n'
               'extern void Plugify_DestroyVector(Vector* vector, enum DataType type);\n'
               'extern void* Plugify_GetVectorData(Vector* vector, enum DataType type);\n'
               'extern ptrdiff_t Plugify_GetVectorSize(Vector* vector, enum DataType type);\n'
               'extern void Plugify_AssignVector(Vector* vector, void* arr, ptrdiff_t len, enum DataType type);\n'
               '\n'
               'extern void Plugify_DeleteVectorDataCStr(void* arr);\n'
               '\n'
               'extern void Plugify_SetGetMethodPtr(void* func);\n'
               'extern void Plugify_SetGetMethodPtr2(void* func);\n'
               '\n'
               'extern void Plugify_SetConstructString(void* func);\n'
               'extern void Plugify_SetDestroyString(void* func);\n'
               'extern void Plugify_SetGetStringData(void* func);\n'
               'extern void Plugify_SetGetStringLength(void* func);\n'
               'extern void Plugify_SetAssignString(void* func);\n'
               'extern void Plugify_SetConstructVector(void* func);\n'
               'extern void Plugify_SetDestroyVector(void* func);\n'
               'extern void Plugify_SetGetVectorData(void* func);\n'
               'extern void Plugify_SetGetVectorSize(void* func);\n'
               'extern void Plugify_SetAssignVector(void* func);\n'
               '\n'
               'extern void Plugify_SetDeleteVectorDataCStr(void* func);\n'
               '\n'
               '//typedef struct { const char *p; ptrdiff_t n; } _GoString_;\n'
               'typedef struct { void *data; ptrdiff_t len; ptrdiff_t cap; } _GoSlice_;\n'
               '\n'
               'typedef struct { float x, y; } Vector2;\n'
               'typedef struct { float x, y, z; } Vector3;\n'
               'typedef struct { float x, y, z, w; } Vector4;\n'
               'typedef struct { float m[4][4]; } Matrix4x4;\n')

    content += '\n'

    for method in pplugin['exportedMethods']:
        ret_type = method['retType']
        return_type = convert_type(ret_type["type"], "ref" in ret_type, True)

        content += (f'typedef {return_type} '
                    f'(*{method['name']}Fn)({gen_params(method, ParamGen.Types)});\n'
                    f'static {return_type} '
                    f'{method['name']}({gen_params(method, ParamGen.TypesNames)}) {{\n'
                    f'\tstatic {method['name']}Fn __func = NULL;\n'
                    f'\tif (__func == NULL) Plugify_GetMethodPtr2("{plugin_name}.{method['name']}", (void**)&__func);\n'
                    f'\t{"return " if ret_type['type'] != 'void' else ""}'
                    f'__func({gen_params(method, ParamGen.Names)});\n'
                    '}\n')

    header_cfile = os.path.join(header_dir, f'{plugin_name}.h')
    with open(header_cfile, 'w', encoding='utf-8') as fd:
        fd.write(content)

    link = 'https://github.com/untrustedmodders/plugify-module-cpp/blob/main/generator/generator.py'

    content = (f'package {plugin_name}\n'
               '\n'
               f'//generated with {link} from {plugin_name} \n'
               '\n'
               f'// #include "{plugin_name}.h"\n'
               'import "C"\n'
               'import "unsafe"\n'
               '\n'
               'type Vector2 struct {\n\tX float32\n\tY float32\n}\n'
               'type Vector3 struct {\n\tX float32\n\tY float32\n\tZ float32\n}\n'
               'type Vector4 struct {\n\tX float32\n\tY float32\n\tZ float32\n\tW float32\n}\n'
               'type Matrix4x4 struct {\n\tM[4][4] float32\n}\n'
               '\n')

    for method in pplugin['exportedMethods']:
        ret_type = method['retType']
        return_type = convert_gotype(ret_type["type"], "ref" in ret_type, True)

        content += (f'func '
                    f'{method['name']}({gen_goparams(method, ParamGen.TypesNames)}) {return_type} {{\n')

        params = gen_goparamscast(method)
        if params != '':
            content += f'\t{params}\n'

        is_obj_ret = is_obj_return(ret_type["type"])
        is_pod_ret = is_pod_return(ret_type["type"])
        if is_pod_ret:
            content += f'\tC_result := C.{method['name']}({gen_goparams(method, ParamGen.Names)})\n'
        elif is_obj_ret:
            content += f'\tC_output := C.{method['name']}({gen_goparams(method, ParamGen.Names)})\n'
        elif ret_type['type'] != 'void':
            content += f'\tresult := {return_type}(C.{method['name']}({gen_goparams(method, ParamGen.Names)}))\n'
        else:
            content += f'\tC.{method['name']}({gen_goparams(method, ParamGen.Names)})\n'

        params = gen_goparamscast_assign(method)
        if params != '':
            content += f'\n{params}'

        params = gen_goparamscast_cleanup(method)
        if params != '':
            content += f'\n{params}\n'

        if is_pod_ret:
            content += f'\treturn *(*{return_type})(unsafe.Pointer(&C_result))\n'
        elif is_obj_ret:
            content += '\treturn output\n'
        elif ret_type['type'] != 'void':
            content += '\treturn result\n'
        content += '}\n\n'

    if not 'unsafe.' in content:
        content = content.replace('import "unsafe"\n', '', 1)

    with open(header_file, 'w', encoding='utf-8') as fd:
        fd.write(content)

    return 0


def get_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('manifest')
    parser.add_argument('output')
    parser.add_argument('--override', action='store_true')
    return parser.parse_args()


if __name__ == '__main__':
    args = get_args()
    sys.exit(main(args.manifest, args.output, args.override))
