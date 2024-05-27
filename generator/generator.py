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
    'char16': 'int16_t',
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
    'char16': 'int16_t*',
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
    'char16': 'int16_t',
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
    'char16': 'int16',
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
    'char16*': '[]int16',
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
    'char16': '*int16',
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
    'char16*': '*[]int16',
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
    'char16': 'int16',
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
    'char16*': '[]int16',
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
    'char16': 'C.int16_t',
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
    'bool*': 'C.Plugify_CreateVectorBool',
    'char8*': 'C.Plugify_CreateVectorChar8',
    'char16*': 'C.Plugify_CreateVectorChar16',
    'int8*': 'C.Plugify_CreateVectorInt8',
    'int16*': 'C.Plugify_CreateVectorInt16',
    'int32*': 'C.Plugify_CreateVectorInt32',
    'int64*': 'C.Plugify_CreateVectorInt64',
    'uint8*': 'C.Plugify_CreateVectorUInt8',
    'uint16*': 'C.Plugify_CreateVectorUInt16',
    'uint32*': 'C.Plugify_CreateVectorUInt32',
    'uint64*': 'C.Plugify_CreateVectorUInt64',
    'ptr64*': 'C.Plugify_CreateVectorUIntPtr',
    'float*': 'C.Plugify_CreateVectorFloat',
    'double*': 'C.Plugify_CreateVectorDouble',
    'string*': 'C.Plugify_CreateVectorString',
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
    'bool*': 'C.Plugify_DeleteVectorBool',
    'char8*': 'C.Plugify_DeleteVectorChar8',
    'char16*': 'C.Plugify_DeleteVectorChar16',
    'int8*': 'C.Plugify_DeleteVectorInt8',
    'int16*': 'C.Plugify_DeleteVectorInt16',
    'int32*': 'C.Plugify_DeleteVectorInt32',
    'int64*': 'C.Plugify_DeleteVectorInt64',
    'uint8*': 'C.Plugify_DeleteVectorUInt8',
    'uint16*': 'C.Plugify_DeleteVectorUInt16',
    'uint32*': 'C.Plugify_DeleteVectorUInt32',
    'uint64*': 'C.Plugify_DeleteVectorUInt64',
    'ptr64*': 'C.Plugify_DeleteVectorUIntPtr',
    'float*': 'C.Plugify_DeleteVectorFloat',
    'double*': 'C.Plugify_DeleteVectorDouble',
    'string*': 'C.Plugify_DeleteVectorString',
    'vec2': '',
    'vec3': '',
    'vec4': '',
    'mat4x4': ''
}

ASS_GOTYPESCAST_MAP = {
    'void': '',
    'bool': 'bool',
    'char8': 'byte',
    'char16': 'int16',
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
    'char16*': 'C.int16_t',
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
    'string*': '*string',
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
    'string': 'unsafe.Pointer',
    'bool*': 'unsafe.Pointer',
    'char8*': 'unsafe.Pointer',
    'char16*': 'unsafe.Pointer',
    'int8*': 'unsafe.Pointer',
    'int16*': 'unsafe.Pointer',
    'int32*': 'unsafe.Pointer',
    'int64*': 'unsafe.Pointer',
    'uint8*': 'unsafe.Pointer',
    'uint16*': 'unsafe.Pointer',
    'uint32*': 'unsafe.Pointer',
    'uint64*': 'unsafe.Pointer',
    'ptr64*': 'unsafe.Pointer',
    'float*': 'unsafe.Pointer',
    'double*': 'unsafe.Pointer',
    'string*': 'unsafe.Pointer',
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
            rtype = REF_GOTYPESCAST_MAP.get(param['type'], 'int')
            if rtype == '&':
                return f'&C_{param['name']}'
            elif rtype != '':
                return f'{rtype}(C_{param['name']})'
            else:
                return f'C_{param['name']}'
        return f'{param['name']} {type}'

    def gen_return(param):
        type = convert_gotype(param['type'], 'ref' in param and param['ref'] is True, False)
        if param_gen == ParamGen.Types:
            return type
        if param_gen == ParamGen.Names:
            rtype = REF_GOTYPESCAST_MAP.get(param['type'], 'int')
            if rtype == '&':
                return '&C_output'
            elif rtype != '':
                return f'{rtype}(C_output)'
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
            rtype = ASS_GOTYPESCAST_MAP.get(param['type'], 'int')
            if rtype == 'C.uintptr_t' or rtype == '*string':
                if 'ref' in param and param['ref'] is True:
                    return f'C_{param['name']} := {type}((*{rtype})(unsafe.Pointer(&(*{param['name']})[0])), C.ptrdiff_t(len(*{param['name']})))'
                else:
                    return f'C_{param['name']} := {type}((*{rtype})(unsafe.Pointer(&{param['name']}[0])), C.ptrdiff_t(len({param['name']})))'
            else:
                if 'ref' in param and param['ref'] is True:
                    return f'C_{param['name']} := {type}((*{rtype})(&(*{param['name']})[0]), C.ptrdiff_t(len(*{param['name']})))'
                else:
                    return f'C_{param['name']} := {type}((*{rtype})(&{param['name']}[0]), C.ptrdiff_t(len({param['name']})))'
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
        type = VAL_GOTYPESCAST_MAP.get(param['type'], 'int')
        if 'C.Vector' in type or 'C.Matrix' in type:
            return f'C_output := {type}' + "{}"
        else:
            return f'C_output := {type}E()'

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
                output = f'P_{param['name']} := C.Plugify_GetString(unsafe.Pointer(C_{param['name']}))\n'
                output += f'\t*{param['name']} = {type}(P_{param['name']})'
                return output
            elif 'C.Vector' in type or 'C.Matrix' in type:
                stype = type[2:]
                return f'*{param['name']} = *(*{stype})(unsafe.Pointer(&C_{param['name']}))'
            elif 'C.' in type or type[0] == '*':
                name = f'unsafe.Pointer(C_{param['name']})'
                mode = ''
                if param['type'] == 'bool*':
                    mode = 'B'
                elif param['type'] == 'string*':
                    mode = 'S'
                gotype = VAL_GOTYPES_MAP.get(param['type'], 'int')[2:]
                output = f'L_{param['name']} := C.Plugify_GetVectorSize{mode}({name})\n'
                output += f'\tP_{param['name']} := C.Plugify_GetVectorData{mode}({name})\n'
                output += f'\t*{param['name']} = unsafe.Slice((*{gotype})(unsafe.Pointer(P_{param['name']})), L_{param['name']})'
                if mode == 'B':
                    output += f'\n\tC.Plugify_DeleteVectorDataB(unsafe.Pointer(P_{param['name']}))'
                elif mode == 'S':
                    output += f'\n\tC.Plugify_DeleteVectorDataS(unsafe.Pointer(P_{param['name']}), L_{param['name']})'
                return output
            else:
                return ''#f'*{param['name']} = {type}(C_{param['name']})'
        else:
            return ''

    def gen_return(param):
        type = ASS_GOTYPESCAST_MAP.get(param['type'], 'int')
        if type == 'C.GoString':
            output = f'P_output := C.Plugify_GetString(unsafe.Pointer(C_output))\n'
            output += f'\toutput := {type}(P_output)'
            return output
        elif 'C.Vector' in type or 'C.Matrix' in type:
            stype = type[2:]
            return f'output := *(*{stype})(unsafe.Pointer(&C_output))'
        elif 'C.' in type or type[0] == '*':
            name = f'unsafe.Pointer(C_output)'
            mode = ''
            if param['type'] == 'bool*':
                mode = 'B'
            elif param['type'] == 'string*':
                mode = 'S'
            gotype = VAL_GOTYPES_MAP.get(param['type'], 'int')[2:]
            output = f'L_output := C.Plugify_GetVectorSize{mode}({name})\n'
            output += f'\tP_output := C.Plugify_GetVectorData{mode}({name})\n'
            output += f'\toutput := unsafe.Slice((*{gotype})(unsafe.Pointer(P_output)), L_output)'
            if mode == 'B':
                output += '\n\tC.Plugify_DeleteVectorDataB(unsafe.Pointer(P_output))'
            elif mode == 'S':
                output += '\n\tC.Plugify_DeleteVectorDataS(unsafe.Pointer(P_output), L_output)'
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
        else:
            return f'{type}(unsafe.Pointer(C_{param['name']}))'

    def gen_return(param):
        type = DEL_GOTYPESCAST_MAP.get(param['type'], 'int')
        if type == '':
            return type
        else:
            return f'{type}(unsafe.Pointer(C_output))'

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
    content += 'extern void* Plugify_GetMethodPtr(const char* str);\n'
    content += 'extern void* Plugify_CreateStringE();\n'
    content += 'extern void* Plugify_CreateString(_GoString_ source);\n'
    content += 'extern const char* Plugify_GetString(void* ptr);\n'
    content += 'extern void Plugify_DeleteString(void* ptr);\n'
    content += 'extern void* Plugify_CreateVectorBool(bool* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorChar8(char* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorChar16(int16_t* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorInt8(int8_t* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorInt16(int16_t* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorInt32(int32_t* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorInt64(int64_t* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorUInt8(uint8_t* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorUInt16(uint16_t* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorUInt32(uint32_t* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorUInt64(uint64_t* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorUIntPtr(uintptr_t* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorFloat(float* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorDouble(double* arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorString(_GoString_** arr, ptrdiff_t len);\n'
    content += 'extern void* Plugify_CreateVectorBoolE();\n'
    content += 'extern void* Plugify_CreateVectorChar8E();\n'
    content += 'extern void* Plugify_CreateVectorChar16E();\n'
    content += 'extern void* Plugify_CreateVectorInt8E();\n'
    content += 'extern void* Plugify_CreateVectorInt16E();\n'
    content += 'extern void* Plugify_CreateVectorInt32E();\n'
    content += 'extern void* Plugify_CreateVectorInt64E();\n'
    content += 'extern void* Plugify_CreateVectorUInt8E();\n'
    content += 'extern void* Plugify_CreateVectorUInt16E();\n'
    content += 'extern void* Plugify_CreateVectorUInt32E();\n'
    content += 'extern void* Plugify_CreateVectorUInt64E();\n'
    content += 'extern void* Plugify_CreateVectorUIntPtrE();\n'
    content += 'extern void* Plugify_CreateVectorFloatE();\n'
    content += 'extern void* Plugify_CreateVectorDoubleE();\n'
    content += 'extern void* Plugify_CreateVectorStringE();\n'
    content += 'extern void Plugify_DeleteVectorBool(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorChar8(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorChar16(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorInt8(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorInt16(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorInt32(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorInt64(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorUInt8(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorUInt16(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorUInt32(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorUInt64(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorUIntPtr(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorFloat(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorDouble(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorString(void* ptr);\n'
    content += 'extern ptrdiff_t Plugify_GetVectorSize(void* ptr);\n'
    content += 'extern void* Plugify_GetVectorData(void* ptr);\n'
    content += 'extern ptrdiff_t Plugify_GetVectorSizeB(void* ptr);\n'
    content += 'extern void* Plugify_GetVectorDataB(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorDataB(void* ptr);\n'
    content += 'extern ptrdiff_t Plugify_GetVectorSizeS(void* ptr);\n'
    content += 'extern void* Plugify_GetVectorDataS(void* ptr);\n'
    content += 'extern void Plugify_DeleteVectorDataS(void* ptr, ptrdiff_t len);\n'
    content += '//typedef struct { const char *p; ptrdiff_t n; } _GoString_;\n'
    content += 'typedef struct { void *data; ptrdiff_t len; ptrdiff_t cap; } _GoSlice_;\n'
    content += 'typedef struct { float x, y; } Vector2;\n'
    content += 'typedef struct { float x, y, z; } Vector3;\n'
    content += 'typedef struct { float x, y, z, w; } Vector4;\n'
    content += 'typedef struct { float m[4][4]; } Matrix4x4;\n'

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
    content += 'type Vector2 struct {\n\tVector2 C.Vector2\n}\n'
    content += 'type Vector3 struct {\n\tVector3 C.Vector3\n}\n'
    content += 'type Vector4 struct {\n\tVector4 C.Vector4\n}\n'
    content += 'type Matrix4x4 struct {\n\tMatrix4x4 C.Matrix4x4\n}\n'
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
