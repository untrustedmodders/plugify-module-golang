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
    'string': 'void*',
    'bool*': 'void*',
    'char8*': 'void*',
    'char16*': 'void*',
    'int8*': 'void*',
    'int16*': 'void*',
    'int32*': 'void*',
    'int64*': 'void*',
    'uint8*': 'void*',
    'uint16*': 'void*',
    'uint32*': 'void*',
    'uint64*': 'void*',
    'ptr64*': 'void*',
    'float*': 'void*',
    'double*': 'void*',
    'string*': 'void*',
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
    'string': 'void*',
    'bool*': 'void*',
    'char8*': 'void*',
    'char16*': 'void*',
    'int8*': 'void*',
    'int16*': 'void*',
    'int32*': 'void*',
    'int64*': 'void*',
    'uint8*': 'void*',
    'uint16*': 'void*',
    'uint32*': 'void*',
    'uint64*': 'void*',
    'ptr64*': 'void*',
    'float*': 'void*',
    'double*': 'void*',
    'string*': 'void*',
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
    'string': 'void',
    'bool*': 'void',
    'char8*': 'void',
    'char16*': 'void',
    'int8*': 'void',
    'int16*': 'void',
    'int32*': 'void',
    'int64*': 'void',
    'uint8*': 'void',
    'uint16*': 'void',
    'uint32*': 'void',
    'uint64*': 'void',
    'ptr64*': 'void',
    'float*': 'void',
    'double*': 'void',
    'string*': 'void',
    'vec2': 'void',
    'vec3': 'void',
    'vec4': 'void',
    'mat4x4': 'void'
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
    'string': 'C.Plugify_CreateString',
    'bool*': 'C.Plugify_CreateVectorC.BOOL',
    'char8*': 'C.Plugify_CreateVectorC.CHAR8',
    'char16*': 'C.Plugify_CreateVectorC.CHAR16',
    'int8*': 'C.Plugify_CreateVectorC.INT8',
    'int16*': 'C.Plugify_CreateVectorC.INT16',
    'int32*': 'C.Plugify_CreateVectorC.INT32',
    'int64*': 'C.Plugify_CreateVectorC.INT64',
    'uint8*': 'C.Plugify_CreateVectorC.UINT8',
    'uint16*': 'C.Plugify_CreateVectorC.UINT16',
    'uint32*': 'C.Plugify_CreateVectorC.UINT32',
    'uint64*': 'C.Plugify_CreateVectorC.UINT64',
    'ptr64*': 'C.Plugify_CreateVectorC.UINTPTR',
    'float*': 'C.Plugify_CreateVectorC.FLOAT',
    'double*': 'C.Plugify_CreateVectorC.DOUBLE',
    'string*': 'C.Plugify_CreateVectorC.STRING',
    'vec2': 'C.Vector2',
    'vec3': 'C.Vector3',
    'vec4': 'C.Vector4',
    'mat4x4': 'C.Matrix4x4'
}

RET_GOTYPESCAST_MAP = {
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
    'string': 'C.Plugify_AllocateString',
    'bool*': 'C.Plugify_AllocateVectorC.BOOL',
    'char8*': 'C.Plugify_AllocateVectorC.CHAR8',
    'char16*': 'C.Plugify_AllocateVectorC.CHAR16',
    'int8*': 'C.Plugify_AllocateVectorC.INT8',
    'int16*': 'C.Plugify_AllocateVectorC.INT16',
    'int32*': 'C.Plugify_AllocateVectorC.INT32',
    'int64*': 'C.Plugify_AllocateVectorC.INT64',
    'uint8*': 'C.Plugify_AllocateVectorC.UINT8',
    'uint16*': 'C.Plugify_AllocateVectorC.UINT16',
    'uint32*': 'C.Plugify_AllocateVectorC.UINT32',
    'uint64*': 'C.Plugify_AllocateVectorC.UINT64',
    'ptr64*': 'C.Plugify_AllocateVectorC.UINTPTR',
    'float*': 'C.Plugify_AllocateVectorC.FLOAT',
    'double*': 'C.Plugify_AllocateVectorC.DOUBLE',
    'string*': 'C.Plugify_AllocateVectorC.STRING',
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
    'string': 'C.Plugify_DeleteString',
    'bool*': 'C.Plugify_DeleteVectorC.BOOL',
    'char8*': 'C.Plugify_DeleteVectorC.CHAR8',
    'char16*': 'C.Plugify_DeleteVectorC.CHAR16',
    'int8*': 'C.Plugify_DeleteVectorC.INT8',
    'int16*': 'C.Plugify_DeleteVectorC.INT16',
    'int32*': 'C.Plugify_DeleteVectorC.INT32',
    'int64*': 'C.Plugify_DeleteVectorC.INT64',
    'uint8*': 'C.Plugify_DeleteVectorC.UINT8',
    'uint16*': 'C.Plugify_DeleteVectorC.UINT16',
    'uint32*': 'C.Plugify_DeleteVectorC.UINT32',
    'uint64*': 'C.Plugify_DeleteVectorC.UINT64',
    'ptr64*': 'C.Plugify_DeleteVectorC.UINTPTR',
    'float*': 'C.Plugify_DeleteVectorC.FLOAT',
    'double*': 'C.Plugify_DeleteVectorC.DOUBLE',
    'string*': 'C.Plugify_DeleteVectorC.STRING',
    'vec2': '',
    'vec3': '',
    'vec4': '',
    'mat4x4': ''
}

FRE_GOTYPESCAST_MAP = {
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
    'string': 'C.Plugify_FreeString',
    'bool*': 'C.Plugify_FreeVectorC.BOOL',
    'char8*': 'C.Plugify_FreeVectorC.CHAR8',
    'char16*': 'C.Plugify_FreeVectorC.CHAR16',
    'int8*': 'C.Plugify_FreeVectorC.INT8',
    'int16*': 'C.Plugify_FreeVectorC.INT16',
    'int32*': 'C.Plugify_FreeVectorC.INT32',
    'int64*': 'C.Plugify_FreeVectorC.INT64',
    'uint8*': 'C.Plugify_FreeVectorC.UINT8',
    'uint16*': 'C.Plugify_FreeVectorC.UINT16',
    'uint32*': 'C.Plugify_FreeVectorC.UINT32',
    'uint64*': 'C.Plugify_FreeVectorC.UINT64',
    'ptr64*': 'C.Plugify_FreeVectorC.UINTPTR',
    'float*': 'C.Plugify_FreeVectorC.FLOAT',
    'double*': 'C.Plugify_FreeVectorC.DOUBLE',
    'string*': 'C.Plugify_FreeVectorC.STRING',
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
    'string': '',
    'bool*': '',
    'char8*': '',
    'char16*': '',
    'int8*': '',
    'int16*': '',
    'int32*': '',
    'int64*': '',
    'uint8*': '',
    'uint16*': '',
    'uint32*': '',
    'uint64*': '',
    'ptr64*': '',
    'float*': '',
    'double*': '',
    'string*': '',
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
    return '*' in type_name or type_name == 'string' or 'vec' in type_name or 'mat' in type_name       


class ParamGen(Enum):
    Types = 1
    Names = 2
    TypesNames = 3


def gen_params_string(method, param_gen: ParamGen):
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
    is_obj_ret = is_obj_return(ret_type["type"])
    if is_obj_ret:
        output_string += f'{gen_return(ret_type)}'
    if method["paramTypes"]:
        if is_obj_ret:
            output_string += ', '
        it = iter(method["paramTypes"])
        output_string += gen_param(next(it))
        for p in it:
            output_string += f', {gen_param(p)}'
    return output_string


def gen_goparams_string(method, param_gen: ParamGen):
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
    is_obj_ret = is_obj_return(ret_type["type"])
    if is_obj_ret:
        output_string += f'{gen_return(ret_type)}'
    if method["paramTypes"]:
        it = iter(method["paramTypes"])
        output_string += gen_param(next(it))
        for p in it:
            output_string += f', {gen_param(p)}'
    return output_string


def gen_goparamscast_string(method):
    def gen_param(param):
        type = VAL_GOTYPESCAST_MAP.get(param['type'], 'int')
        if 'CreateVector' in type:
            if 'ref' in param and param['ref'] is True:
                return f'C_{param['name']} := {type[:22]}(unsafe.Pointer(&(*{param['name']})[0]), C.ptrdiff_t(len(*{param['name']})), {type[22:]})'
            else:
                return f'C_{param['name']} := {type[:22]}(unsafe.Pointer(&{param['name']}[0]), C.ptrdiff_t(len({param['name']})), {type[22:]})'
        elif 'CreateString' in type:
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
        type = RET_GOTYPESCAST_MAP.get(param['type'], 'int')
        if 'C.Vector' in type or 'C.Matrix' in type:
            return f'C_output := {type}' + "{}"
        else:
            return f'C_output := {type[:24]}({type[24:]})'

    output_string = ''
    ret_type = method['retType']
    is_obj_ret = is_obj_return(ret_type["type"])
    if is_obj_ret:
        output_string += f'{gen_return(ret_type)}'
    if method["paramTypes"]:
        it = iter(method["paramTypes"])
        output_string += f'{gen_param(next(it))}\n'
        for p in it:
            output_string += f'\t{gen_param(p)}\n'
    return output_string


def gen_goparamscast_assign_string(method):
    def gen_param(param):
        if 'ref' in param and param['ref'] is True:
            type = ASS_GOTYPESCAST_MAP.get(param['type'], 'int')
            if type == 'C.GoString':
                output = f'P_{param['name']} := C.Plugify_GetStringData(C_{param['name']})\n'
                output += f'\t*{param['name']} = {type}(P_{param['name']})'
                return output
            elif 'C.Vector' in type or 'C.Matrix' in type:
                return f'*{param['name']} = *(*{type[2:]})(unsafe.Pointer(&C_{param['name']}))'
            elif 'C.' in type or type[0] == '*' or type == 'string':
                name = f'C_{param['name']}'
                mode = ''
                if param['type'] == 'bool*':
                    mode = 'Bool'
                elif param['type'] == 'string*':
                    mode = 'CStr'
                go_type = VAL_GOTYPES_MAP.get(param['type'], 'int')[2:]

                val_type = VAL_GOTYPESCAST_MAP.get(param['type'], 'int')[22:]
                output = f'L_{param['name']} := C.Plugify_GetVectorSize({name}, {val_type})\n'
                output += f'\tP_{param['name']} := C.Plugify_GetVectorData({name}, {val_type})\n'
                #output += f'\t*{param['name']} = unsafe.Slice((*{go_type})(P_{param['name']}), L_{param['name']})'
                output += f'\t*{param['name']} = make([]{go_type}, L_{param['name']})\n'
                output += f'\tfor i := range (*{param['name']})' + " {\n"
                if mode == 'CStr':
                    output += f'\t\t(*{param['name']})[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(P_{param['name']}) + uintptr(i * C.sizeof_uintptr_t))))\n'
                else:
                    output += f'\t\t(*{param['name']})[i] = *(*{go_type})(unsafe.Pointer(uintptr(P_{param['name']}) + uintptr(i * C.sizeof_{type[2:]})))\n'
                output += "\t}"
                if mode == 'Bool' or mode == 'CStr':
                    output += f'\n\tC.Plugify_DeleteVectorData{mode}(P_{param['name']})'
                return output
            else:
                return ''#f'*{param['name']} = {type}(C_{param['name']})'
        else:
            return ''

    def gen_return(param):
        type = ASS_GOTYPESCAST_MAP.get(param['type'], 'int')
        if type == 'C.GoString':
            output = f'P_output := C.Plugify_GetStringData(C_output)\n'
            output += f'\toutput := {type}(P_output)'
            return output
        elif 'C.Vector' in type or 'C.Matrix' in type:
            return f'output := *(*{type[2:]})(unsafe.Pointer(&C_output))'
        elif 'C.' in type or type[0] == '*' or type == 'string':
            name = f'C_output'
            mode = ''
            if param['type'] == 'bool*':
                mode = 'Bool'
            elif param['type'] == 'string*':
                mode = 'CStr'
            go_type = VAL_GOTYPES_MAP.get(param['type'], 'int')[2:]
            val_type = VAL_GOTYPESCAST_MAP.get(param['type'], 'int')[22:]
            output = f'L_output := C.Plugify_GetVectorSize({name}, {val_type})\n'
            output += f'\tP_output := C.Plugify_GetVectorData({name}, {val_type})\n'
            #output += f'\toutput := unsafe.Slice((*{go_type})(P_output), L_output)'
            output += f'\toutput := make([]{go_type}, L_output)\n'
            output += f'\tfor i := range output' + " {\n"
            if mode == 'CStr':
                output += f'\t\toutput[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_uintptr_t))))\n'
            else:
                output += f'\t\toutput[i] = *(*{go_type})(unsafe.Pointer(uintptr(P_output) + uintptr(i * C.sizeof_{type[2:]})))\n'
            output += "\t}\n"
            if mode == 'Bool' or mode == 'CStr':
                output += f'\n\tC.Plugify_DeleteVectorData{mode}(P_output)'
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


def gen_goparamscast_cleanup_string(method):
    def gen_param(param):
        type = DEL_GOTYPESCAST_MAP.get(param['type'], 'int')
        if type == '':
            return type
        elif 'DeleteVector' in type:
            return f'{type[:22]}(C_{param['name']}, {type[22:]})'
        else:
            return f'{type}(C_{param['name']})'

    def gen_return(param):
        type = FRE_GOTYPESCAST_MAP.get(param['type'], 'int')
        if type == '':
            return type
        elif 'FreeVector' in type:
            return f'{type[:20]}(C_output, {type[20:]})'
        else:
            return f'{type}(C_output)'

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

    content = '#pragma once\n'
    content += '#include <stdlib.h>\n'
    content += '#include <stdint.h>\n'
    content += '#include <stdbool.h>\n\n'
    content += 'enum DataType {\n'
    content += '\tBOOL,\n'
    content += '\tCHAR8,\n'
    content += '\tCHAR16,\n'
    content += '\tINT8,\n'
    content += '\tINT16,\n'
    content += '\tINT32,\n'
    content += '\tINT64,\n'
    content += '\tUINT8,\n'
    content += '\tUINT16,\n'
    content += '\tUINT32,\n'
    content += '\tUINT64,\n'
    content += '\tUINTPTR,\n'
    content += '\tFLOAT,\n'
    content += '\tDOUBLE,\n'
    content += '\tSTRING\n'
    content += '};\n\n'
    content += 'extern void* Plugify_GetMethodPtr(const char* methodName);\n'
    content += '\n'
    content += 'extern void* Plugify_AllocateString();\n'
    content += 'extern void* Plugify_CreateString(_GoString_ source);\n'
    content += 'extern const char* Plugify_GetStringData(void* ptr);\n'
    content += 'extern ptrdiff_t Plugify_GetStringLength(void* ptr);\n'
    content += 'extern void Plugify_AssignString(void* ptr, _GoString_ source);\n'
    content += 'extern void Plugify_FreeString(void* ptr);\n'
    content += 'extern void Plugify_DeleteString(void* ptr);\n'
    content += '\n'
    content += 'extern void* Plugify_CreateVector(void* arr, ptrdiff_t len, enum DataType type);\n'
    content += 'extern void* Plugify_AllocateVector(enum DataType type);\n'
    content += 'extern ptrdiff_t Plugify_GetVectorSize(void* ptr, enum DataType type);\n'
    content += 'extern void* Plugify_GetVectorData(void* ptr, enum DataType type);\n'
    content += 'extern void Plugify_AssignVector(void* ptr, void* arr, ptrdiff_t len, enum DataType type);\n'
    content += 'extern void Plugify_DeleteVector(void* ptr, enum DataType type);\n'
    content += 'extern void Plugify_FreeVector(void* ptr, enum DataType type);\n'
    content += '\n'
    content += 'extern void Plugify_DeleteVectorDataBool(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorDataCStr(void* ptr);\n'
    content += '\n'
    content += 'extern void Plugify_SetGetMethodPtr(void* func);\n'
    content += 'extern void Plugify_SetAllocateString(void* func);\n'
    content += 'extern void Plugify_SetCreateString(void* func);\n'
    content += 'extern void Plugify_SetGetStringData(void* func);\n'
    content += 'extern void Plugify_SetGetStringLength(void* func);\n'
    content += 'extern void Plugify_SetAssignString(void* func);\n'
    content += 'extern void Plugify_SetFreeString(void* func);\n'
    content += 'extern void Plugify_SetDeleteString(void* func);\n'
    content += 'extern void Plugify_SetCreateVector(void* func);\n'
    content += 'extern void Plugify_SetAllocateVector(void* func);\n'
    content += 'extern void Plugify_SetGetVectorSize(void* func);\n'
    content += 'extern void Plugify_SetGetVectorData(void* func);\n'
    content += 'extern void Plugify_SetAssignVector(void* func);\n'
    content += 'extern void Plugify_SetDeleteVector(void* func);\n'
    content += 'extern void Plugify_SetFreeVector(void* func);\n'
    content += '\n'
    content += 'extern void Plugify_SetDeleteVectorDataBool(void* func);\n'
    content += 'extern void Plugify_SetDeleteVectorDataCStr(void* func);\n'
    content += '\n'
    content += '//typedef struct { const char *p; ptrdiff_t n; } _GoString_;\n'
    content += 'typedef struct { void *data; ptrdiff_t len; ptrdiff_t cap; } _GoSlice_;\n'
    content += '\n'
    content += 'typedef struct { float x, y; } Vector2;\n'
    content += 'typedef struct { float x, y, z; } Vector3;\n'
    content += 'typedef struct { float x, y, z, w; } Vector4;\n'
    content += 'typedef struct { float m00, m01, m02, m03, m10, m11, m12, m13, m20, m21, m22, m23, m30, m31, m32, m33; } Matrix4x4;\n'

    content += '\n'

    for method in pplugin['exportedMethods']:
        ret_type = method['retType']
        return_type = convert_type(ret_type["type"], "ref" in ret_type, True)

        content += (f'typedef {return_type} '
                    f'(*{method['name']}Fn)({gen_params_string(method, ParamGen.Types)});\n')
        content += (f'static {return_type} '
                    f'{method['name']}({gen_params_string(method, ParamGen.TypesNames)}) {{\n')
        content += f'\tstatic {method['name']}Fn func = NULL;\n'
        content += f'\tif (func == NULL) func = ({method['name']}Fn)Plugify_GetMethodPtr("{plugin_name}.{method['name']}");\n'
        content += (f'\t{"return " if ret_type['type'] != 'void' else ""}'
                    f'func({gen_params_string(method, ParamGen.Names)});\n')
        content += '}\n'

    header_cfile = os.path.join(header_dir, f'{plugin_name}.h')
    with open(header_cfile, 'w', encoding='utf-8') as fd:
        fd.write(content)

    content = ''

    # TODO: Make POD structures pass as fist argument for return types

    link = 'https://github.com/untrustedmodders/cpp-lang-module/blob/main/generator/generator.py'

    content += f'package {plugin_name}\n'
    content += '\n'
    content += f'//generated with {link} from {plugin_name} \n'
    content += '\n'
    content += f'// #include "{plugin_name}.h"\n'
    content += 'import "C"\n'
    content += 'import "unsafe"\n'
    content += '\n'
    content += 'type Vector2 struct {\n\tX float32\n\tY float32\n}\n'
    content += 'type Vector3 struct {\n\tX float32\n\tY float32\n\tZ float32\n}\n'
    content += 'type Vector4 struct {\n\tX float32\n\tY float32\n\tZ float32\n\tW float32\n}\n'
    content += 'type Matrix4x4 struct {\n\tM [4][4]float32\n}\n'
    content += '\n'

    for method in pplugin['exportedMethods']:
        ret_type = method['retType']
        return_type = convert_gotype(ret_type["type"], "ref" in ret_type, True)
        if return_type != '':
            return_type += ' '

        content += (f'func '
                    f'{method['name']}({gen_goparams_string(method, ParamGen.TypesNames)}) {return_type}{{\n')

        params = gen_goparamscast_string(method)
        if params != '':
            content += f'\t{params}\n'

        is_obj_ret = is_obj_return(ret_type["type"])
        if not is_obj_ret and ret_type['type'] != 'void':
            content += f'\tresult := {return_type}(C.{method['name']}({gen_goparams_string(method, ParamGen.Names)}))\n'
        else:
            content += f'\tC.{method['name']}({gen_goparams_string(method, ParamGen.Names)})\n'

        params = gen_goparamscast_assign_string(method)
        if params != '':
            content += f'\n{params}'

        params = gen_goparamscast_cleanup_string(method)
        if params != '':
            content += f'\n{params}\n'

        if is_obj_ret:
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
    parser.add_argument('--override')
    return parser.parse_args()


if __name__ == '__main__':
    args = get_args()
    sys.exit(main(args.manifest, args.output, args.override))
