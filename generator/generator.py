#!/usr/bin/python3
import sys
import argparse
import os
import json
from enum import IntEnum

TYPES_MAP = {
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
    'function': 'delegate',
    'string': 'string',
    'any': 'interface{}',
    'bool[]': '[]bool',
    'char8[]': '[]int8',
    'char16[]': '[]uint16',
    'int8[]': '[]int8',
    'int16[]': '[]int16',
    'int32[]': '[]int32',
    'int64[]': '[]int64',
    'uint8[]': '[]uint8',
    'uint16[]': '[]uint16',
    'uint32[]': '[]uint32',
    'uint64[]': '[]uint64',
    'ptr64[]': '[]uintptr',
    'float[]': '[]float32',
    'double[]': '[]float64',
    'string[]': '[]string',
    'any[]': '[]interface{}',
    'vec2[]': '[]plugify.Vector2',
    'vec3[]': '[]plugify.Vector3',
    'vec4[]': '[]plugify.Vector4',
    'mat4x4[]': '[]plugify.Matrix4x4',
    'vec2': 'plugify.Vector2',
    'vec3': 'plugify.Vector3',
    'vec4': 'plugify.Vector4',
    'mat4x4': 'plugify.Matrix4x4'
}

RTYPES_MAP = {
    'void': 'void',
    'bool': 'bool',
    'char8': 'int8_t',
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
    'function': 'void*',
    'string': 'String*',
    'any': 'Variant*',
    'bool[]': 'Vector*',
    'char8[]': 'Vector*',
    'char16[]': 'Vector*',
    'int8[]': 'Vector*',
    'int16[]': 'Vector*',
    'int32[]': 'Vector*',
    'int64[]': 'Vector*',
    'uint8[]': 'Vector*',
    'uint16[]': 'Vector*',
    'uint32[]': 'Vector*',
    'uint64[]': 'Vector*',
    'ptr64[]': 'Vector*',
    'float[]': 'Vector*',
    'double[]': 'Vector*',
    'string[]': 'Vector*',
    'any[]': 'Vector*',
    'vec2[]': 'Vector*',
    'vec3[]': 'Vector*',
    'vec4[]': 'Vector*',
    'mat4x4[]': 'Vector*',
    'vec2': 'Vector2*',
    'vec3': 'Vector3*',
    'vec4': 'Vector4*',
    'mat4x4': 'Matrix4x4*'
}

CTYPES_MAP = {
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
    'string': 'C.String',
    'any': 'C.Variant',
    'bool[]': 'C.Vector',
    'char8[]': 'C.Vector',
    'char16[]': 'C.Vector',
    'int8[]': 'C.Vector',
    'int16[]': 'C.Vector',
    'int32[]': 'C.Vector',
    'int64[]': 'C.Vector',
    'uint8[]': 'C.Vector',
    'uint16[]': 'C.Vector',
    'uint32[]': 'C.Vector',
    'uint64[]': 'C.Vector',
    'ptr64[]': 'C.Vector',
    'float[]': 'C.Vector',
    'double[]': 'C.Vector',
    'string[]': 'C.Vector',
    'any[]': 'C.Vector',
    'vec2[]': 'C.Vector',
    'vec3[]': 'C.Vector',
    'vec4[]': 'C.Vector',
    'mat4x4[]': 'C.Vector',
    'vec2': 'C.Vector2',
    'vec3': 'C.Vector3',
    'vec4': 'C.Vector4',
    'mat4x4': 'C.Matrix4x4'
}

FTYPES_MAP = {
    'void': '',
    'bool': 'Bool',
    'char8': 'Int8',
    'char16': 'UInt16',
    'int8': 'Int8',
    'int16': 'Int16',
    'int32': 'Int32',
    'int64': 'Int64',
    'uint8': 'UInt8',
    'uint16': 'UInt16',
    'uint32': 'UInt32',
    'uint64': 'UInt64',
    'ptr64': 'Pointer',
    'ptr32': 'Pointer',
    'float': 'Float',
    'double': 'Double',
    'string': 'String',
    'any': 'Variant',
    'vec2': 'Vector2',
    'vec3': 'Vector3',
    'vec4': 'Vector4',
    'mat4x4': 'Matrix4x4',
    'bool[]': 'Bool',
    'char8[]': 'Char8',
    'char16[]': 'Char16',
    'int8[]': 'Int8',
    'int16[]': 'Int16',
    'int32[]': 'Int32',
    'int64[]': 'Int64',
    'uint8[]': 'UInt8',
    'uint16[]': 'UInt16',
    'uint32[]': 'UInt32',
    'uint64[]': 'UInt64',
    'ptr32[]': 'Pointer',
    'ptr64[]': 'Pointer',
    'float[]': 'Float',
    'double[]': 'Double',
    'string[]': 'String',
    'any[]': 'Variant',
    'vec2[]': 'Vector2',
    'vec3[]': 'Vector3',
    'vec4[]': 'Vector4',
    'mat4x4[]': 'Matrix4x4'
}

VAL_TYPESCAST_MAP = {
    'void': '',
    'bool': 'C.bool',
    'char8': 'C.int8_t',
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
    'function': 'plugify.GetFunctionPointerForDelegate',
    'string': 'plugify.ConstructString',
    'any': 'plugify.ConstructVariant',
    'bool[]': 'plugify.ConstructVectorBool',
    'char8[]': 'plugify.ConstructVectorChar8',
    'char16[]': 'plugify.ConstructVectorChar16',
    'int8[]': 'plugify.ConstructVectorInt8',
    'int16[]': 'plugify.ConstructVectorInt16',
    'int32[]': 'plugify.ConstructVectorInt32',
    'int64[]': 'plugify.ConstructVectorInt64',
    'uint8[]': 'plugify.ConstructVectorUInt8',
    'uint16[]': 'plugify.ConstructVectorUInt16',
    'uint32[]': 'plugify.ConstructVectorUInt32',
    'uint64[]': 'plugify.ConstructVectorUInt64',
    'ptr64[]': 'plugify.ConstructVectorPointer',
    'float[]': 'plugify.ConstructVectorFloat',
    'double[]': 'plugify.ConstructVectorDouble',
    'string[]': 'plugify.ConstructVectorString',
    'any[]': 'plugify.ConstructVectorVariant',
    'vec2[]': 'plugify.ConstructVectorVector2',
    'vec3[]': 'plugify.ConstructVectorVector3',
    'vec4[]': 'plugify.ConstructVectorVector4',
    'mat4x4[]': 'plugify.ConstructVectorMatrix4x4',
    'vec2': 'C.Vector2',
    'vec3': 'C.Vector3',
    'vec4': 'C.Vector4',
    'mat4x4': 'C.Matrix4x4'
}

RET_TYPESCAST_MAP = {
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
    'function': '',
    'string': 'plugify.PlgString',
    'any': 'plugify.PlgVariant',
    'bool[]': 'plugify.PlgVector',
    'char8[]': 'plugify.PlgVector',
    'char16[]': 'plugify.PlgVector',
    'int8[]': 'plugify.PlgVector',
    'int16[]': 'plugify.PlgVector',
    'int32[]': 'plugify.PlgVector',
    'int64[]': 'plugify.PlgVector',
    'uint8[]': 'plugify.PlgVector',
    'uint16[]': 'plugify.PlgVector',
    'uint32[]': 'plugify.PlgVector',
    'uint64[]': 'plugify.PlgVector',
    'ptr64[]': 'plugify.PlgVector',
    'float[]': 'plugify.PlgVector',
    'double[]': 'plugify.PlgVector',
    'string[]': 'plugify.PlgVector',
    'any[]': 'plugify.PlgVector',
    'vec2[]': 'plugify.PlgVector',
    'vec3[]': 'plugify.PlgVector',
    'vec4[]': 'plugify.PlgVector',
    'mat4x4[]': 'plugify.PlgVector',
    'vec2': 'plugify.Vector2',
    'vec3': 'plugify.Vector3',
    'vec4': 'plugify.Vector4',
    'mat4x4': 'plugify.Matrix4x4'
}

ASS_TYPESCAST_MAP = {
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
    'function': '',
    'string': 'plugify.GetStringData',
    'any': 'plugify.GetVariantData',
    'bool[]': 'plugify.GetVectorDataBool',
    'char8[]': 'plugify.GetVectorDataChar8',
    'char16[]': 'plugify.GetVectorDataChar16',
    'int8[]': 'plugify.GetVectorDataInt8',
    'int16[]': 'plugify.GetVectorDataInt16',
    'int32[]': 'plugify.GetVectorDataInt32',
    'int64[]': 'plugify.GetVectorDataInt64',
    'uint8[]': 'plugify.GetVectorDataUInt8',
    'uint16[]': 'plugify.GetVectorDataUInt16',
    'uint32[]': 'plugify.GetVectorDataUInt32',
    'uint64[]': 'plugify.GetVectorDataUInt64',
    'ptr64[]': 'plugify.GetVectorDataPointer',
    'float[]': 'plugify.GetVectorDataFloat',
    'double[]': 'plugify.GetVectorDataDouble',
    'string[]': 'plugify.GetVectorDataString',
    'any[]': 'plugify.GetVectorDataVariant',
    'vec2[]': 'plugify.GetVectorDataVector2',
    'vec3[]': 'plugify.GetVectorDataVector3',
    'vec4[]': 'plugify.GetVectorDataVector4',
    'mat4x4[]': 'plugify.GetVectorDataMatrix4x4',
    'vec2': 'plugify.Vector2',
    'vec3': 'plugify.Vector3',
    'vec4': 'plugify.Vector4',
    'mat4x4': 'plugify.Matrix4x4'
}

DEL_TYPESCAST_MAP = {
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
    'string': 'plugify.DestroyString',
    'any': 'plugify.DestroyVariant',
    'bool[]': 'plugify.DestroyVectorBool',
    'char8[]': 'plugify.DestroyVectorChar8',
    'char16[]': 'plugify.DestroyVectorChar16',
    'int8[]': 'plugify.DestroyVectorInt8',
    'int16[]': 'plugify.DestroyVectorInt16',
    'int32[]': 'plugify.DestroyVectorInt32',
    'int64[]': 'plugify.DestroyVectorInt64',
    'uint8[]': 'plugify.DestroyVectorUInt8',
    'uint16[]': 'plugify.DestroyVectorUInt16',
    'uint32[]': 'plugify.DestroyVectorUInt32',
    'uint64[]': 'plugify.DestroyVectorUInt64',
    'ptr64[]': 'plugify.DestroyVectorPointer',
    'float[]': 'plugify.DestroyVectorFloat',
    'double[]': 'plugify.DestroyVectorDouble',
    'string[]': 'plugify.DestroyVectorString',
    'any[]': 'plugify.DestroyVectorVariant',
    'vec2[]': 'plugify.DestroyVectorVector2',
    'vec3[]': 'plugify.DestroyVectorVector3',
    'vec4[]': 'plugify.DestroyVectorVector4',
    'mat4x4[]': 'plugify.DestroyVectorMatrix4x4',
    'vec2': '',
    'vec3': '',
    'vec4': '',
    'mat4x4': ''
}

INVALID_NAMES = {
    'break',
    'case',
    'chan',
    'const',
    'continue',
    'default',
    'defer',
    'else',
    'fallthrough',
    'for',
    'func',
    'go',
    'goto',
    'if',
    'import',
    'interface',
    'map',
    'package',
    'range',
    'return',
    'select',
    'struct',
    'switch',
    'type',
    'var',
    'append',
    'bool',
    'byte',
    'cap',
    'close',
    'complex',
    'complex64',
    'complex128',
    'copy',
    'delete',
    'error',
    'false',
    'float32',
    'float64',
    'imag',
    'int',
    'int8',
    'int16',
    'int32',
    'int64',
    'iota',
    'len',
    'make',
    'new',
    'nil',
    'panic',
    'print',
    'println',
    'real',
    'recover',
    'rune',
    'string',
    'true',
    'uint',
    'uint8',
    'uint16',
    'uint32',
    'uint64',
    'uintptr',
}


def is_obj_type(type_name: str) -> bool:
    """
    Checks if the type is a reference type or a collection type (array or object).
    """
    return '[]' in type_name or type_name in {'string', 'any'}


def is_pod_type(type_name: str) -> bool:
    """
    Checks if the type is a Plain Old Data (POD) type, such as vectors or matrices.
    """
    return type_name in {'vec2', 'vec3', 'vec4', 'mat4x4'}


def convert_type(type_name: str, is_ref: bool = False) -> str:
    """
    Converts a type name to its corresponding Go type.
    If the type is a reference, it appends 'ref' to the type.
    """
    base_type = TYPES_MAP.get(type_name, None)
    return f'*{base_type}' if is_ref else base_type


def convert_ctype(type_name: str, is_ref: bool = False, is_ret: bool = False) -> str:
    """
    Converts a type name to its corresponding C type.
    If the type is a reference, it appends '*' (pointer) to the type.
    """
    base_type = RTYPES_MAP.get(type_name, None)
    if '*' in base_type:
        if is_ret and base_type != 'void*':
            return base_type[:-1]  # Remove pointer if it's a return type
        else:
            return base_type
    elif is_ref:
        return f'{base_type}*' if '*' not in base_type else base_type
    return base_type


def adjust_type_name(param: dict, type_name: str, is_ref: bool) -> str:
    if 'delegate' in type_name and 'prototype' in param:
        return generate_name(param['prototype'].get('name', 'UnnamedDelegate'))
    elif 'enum' in param:
        name = generate_name(param['enum'].get('name', 'UnnamedEnum'))
        if '[]' in type_name:
            return f'*[]{name}' if is_ref else f'[]{name}'
        else:
            return f'*{name}' if is_ref else name
    return type_name


def adjust_ctype_name(param: dict, type_name: str, is_ref: bool) -> str:
    return type_name


def get_type_name(param: dict) -> str:
    is_ref = 'ref' in param and param['ref']
    type_name = convert_type(param['type'], is_ref)
    return adjust_type_name(param, type_name, is_ref)


def get_ctype_name(param: dict, is_ret: bool) -> str:
    is_ref = 'ref' in param and param['ref']
    type_name = convert_ctype(param['type'], is_ref, is_ret)
    return adjust_ctype_name(param, type_name, is_ref)


def generate_name(name: str) -> str:
    """
    Generates a valid Go name by appending an underscore if the name is invalid.
    """
    return f'{name}_' if name in INVALID_NAMES else name


class ParamGen(IntEnum):
    Types = 1
    Names = 2
    TypesNames = 3
    CastNames = 5


def gen_params(method: dict, param_gen: ParamGen) -> str:
    """
    Generates the parameters string for the method based on the param_gen type.
    Handles different modes such as Types, Names, CastNames, and TypesNames.
    """

    # Helper function to generate the type of a parameter, considering references and function types
    def gen_param_type(param: dict) -> str:
        if param['type'] == 'function':
            return generate_name(param['prototype']['name'])
        return convert_type(param['type'], 'ref' in param and param['ref'])

    # Helper function to generate the appropriate cast name for parameters, handling function pointers and references
    def gen_param_cast_name(param: dict) -> str:
        if is_obj_type(param['type']):
            return f'(*{CTYPES_MAP[param['type']]})(unsafe.Pointer(&__{generate_name(param["name"])}))'
        if 'vec' in param['type'] or 'mat' in param['type']:
            return gen_vector_matrix_cast(param)
        return gen_ref_cast(param)

    # Helper function to generate the cast name for vector or matrix types
    def gen_vector_matrix_cast(param: dict) -> str:
        return f'&__{generate_name(param["name"])}'

    # Helper function to handle ref type parameters
    def gen_ref_cast(param: dict) -> str:
        if 'ref' in param and param['ref']:
            return f'&__{generate_name(param["name"])}'
        return f'__{generate_name(param["name"])}'

    # Helper function to generate the parameter type and cast name for TypesNames
    def gen_param_type_name(param: dict) -> str:
        return f'{generate_name(param["name"])} {get_type_name(param)}'

    # Function to generate the appropriate string for a single parameter based on the param_gen mode
    def gen_param(param: dict) -> str:
        match param_gen:
            # Handle parameter type generation (e.g., ref, function, delegate)
            case ParamGen.Types:
                return gen_param_type(param)
            # Handle parameter name generation
            case ParamGen.Names:
                return generate_name(param['name'])
            # Handle cast names for marshaling or references
            case ParamGen.CastNames:
                return gen_param_cast_name(param)
            # Handle types with type names
            case ParamGen.TypesNames:
                return gen_param_type_name(param)
            case _:
                return ''

    # Generate the full parameters list as a string
    parts = []

    # Add each parameter to the list using gen_param
    if method.get('paramTypes'):
        parts.extend([gen_param(param) for param in method['paramTypes']])

    # Join the parts with a comma and space to form the final string
    return ', '.join(parts)


def gen_cparams(method: dict, param_gen: ParamGen) -> str:
    """Generate function parameters based on the specified type."""
    def gen_param(index: int, param: dict) -> str:
        type_name = get_ctype_name(param, False)
        match param_gen:
            case ParamGen.Types:
                return type_name
            case ParamGen.Names:
                return param.get('name', f'p{index}')
            case ParamGen.TypesNames:
                return f'{type_name} {param.get("name", f"p{index}")}'
            case _:
                return ''

    # Generate the full parameters list as a string
    parts = []

    # Add each parameter to the list using gen_param
    if method.get('paramTypes'):
        parts.extend([gen_param(i, p) for i, p in enumerate(method['paramTypes'])])

    # Join the parts with a comma and space to form the final string
    return ', '.join(parts)


def gen_types(method: dict) -> str:
    """
    Generates a string representation of the types for the method, including its parameters and return type.
    """
    # Initialize list to hold type representations
    parts = []

    # Process parameters if they exist
    if method.get('paramTypes'):
        parts.extend([get_type_name(param) for param in method['paramTypes']])

    # Add the return type
    ret_type = method.get('retType', {})
    if ret_type:
        parts.append(get_type_name(ret_type))

    # Join all parts together and return the final string
    return ', '.join(parts)


def gen_paramscast(method: dict, tabs: str) -> list[str]:
    """
    Generates parameter casting code for a method, considering various cases like
    function pointers, references, and vector constructions.
    """

    def gen_param(param: dict) -> str:
        """
        Generates the appropriate casting code for a parameter, considering its type,
        references, and other special cases like function pointers or vector constructions.
        """
        # Check if the parameter type requires a casting transformation (from VAL_TYPESCAST_MAP)
        param_type = VAL_TYPESCAST_MAP.get(param['type'], None)
        name = generate_name(param['name'])

        # Handle vector type-base casting
        if 'C.Vector' in param_type or 'C.Matrix' in param_type:
            if 'ref' in param and param['ref'] is True:
                return f'__{name} := *(*{param_type})(unsafe.Pointer({name}))'
            else:
                return f'__{name} := *(*{param_type})(unsafe.Pointer(&{name}))'

        # Handle general type-based casting (other than function pointers)
        elif param_type:
            if 'ref' in param and param['ref'] is True:
                return f'__{name} := {param_type}(*{name})'
            else:
                return f'__{name} := {param_type}({name})'

        # Return an empty string for unhandled cases
        return ''

    def gen_return(param) -> str:
        """
        Generates the return value casting code for the method, if applicable.
        """
        return f'var __retVal_native {RET_TYPESCAST_MAP.get(param['type'], None)}'

    output_parts = []

    # Handle return type casting if the return type is an object (requires marshaling)
    ret_type = method.get('retType', {})
    if ret_type:
        if is_obj_type(ret_type.get('type')):
            ret_val = gen_return(ret_type)
            if ret_val:
                output_parts.append(f'{tabs}{ret_val}')

    # Handle casting for each parameter in the method
    if method.get('paramTypes'):
        for param in method['paramTypes']:
            param_cast = gen_param(param)
            if param_cast:
                output_parts.append(f'{tabs}{param_cast}')

    # Return the full generated string of parameter and return type casting code
    return output_parts


def gen_paramscast_assign(method: dict, tabs: str) -> list[str]:
    """
    Generates parameter assignment code for a method, including casting, marshaling for 'ref' types,
    and handling special types like 'VectorData'.
    """

    def gen_param(param: dict) -> str:
        """
        Generates assignment code for a single parameter, considering 'ref' types and other special cases like 'VectorData'.
        """
        if 'ref' in param and param['ref'] is True:
            param_type = ASS_TYPESCAST_MAP.get(param['type'], None)
            name = generate_name(param['name'])

            # Handle vector type-base casting
            if 'plugify.Vector' in param_type or 'plugify.Matrix' in param_type:
                return f'*{name} = *(*{param_type})(unsafe.Pointer(&__{name}))'

            # Handle 'VectorData' type
            elif 'VectorData' in param_type:
                return f'{param_type}To(&__{name}, {name})'
            elif 'plugify.' in param_type:  # Handle other types
                return f'*{name} = {param_type}(&__{name})'
            elif param_type != '':  # Handle other types
                if 'enum' in param:
                    return f'*{name} = {param['enum'].get('name', 'UnnamedEnum')}(__{name})'
                else:
                    return f'*{name} = {param_type}(__{name})'
            else:
                return ''
        return ''  # Return empty if no 'ref' is in param

    def gen_return(param) -> str:
        """
        Generates assignment code for the return type, considering 'VectorData' and other special cases.
        """
        param_type = ASS_TYPESCAST_MAP.get(param['type'], None)

        if param_type != '':  # Handle other types
            if 'enum' in param:
                return f'__retVal = {param_type}T[{param['enum'].get('name', 'UnnamedEnum')}](&__retVal_native)'
            else:
                return f'__retVal = {param_type}(&__retVal_native)'
        else:
            return ''  # Return empty if no valid type

    # List to hold the code parts for parameters and return
    output_parts = []

    # Handle return type if needed
    ret_type = method.get('retType', {})
    if ret_type:
        if is_obj_type(ret_type.get('type')):
            ret_val = gen_return(ret_type)
            if ret_val:
                output_parts.append(f'{tabs}{ret_val}')

    # Handle parameter types
    if method.get('paramTypes'):
        for param in method['paramTypes']:
            param_cast = gen_param(param)
            if param_cast:
                output_parts.append(f'{tabs}{param_cast}')

    return output_parts


def gen_paramscast_cleanup(method: dict, tabs: str) -> list[str]:
    """
    Generate code for parameter and return value cleanup using type casts.
    """

    def gen_param(param: dict) -> str:
        """
        Generate type casting cleanup code for a parameter.
        """
        param_type = DEL_TYPESCAST_MAP.get(param.get('type'))
        name = generate_name(param['name'])

        if param_type:
            return f'{param_type}(&__{name})'

        return ''  # Return empty string if type is not in DEL_TYPESCAST_MAP

    def gen_return(ret_type: dict):
        """
        Generate type casting cleanup code for a return value.
        """
        return_type = DEL_TYPESCAST_MAP.get(ret_type.get('type'))
        if return_type:
            return f'{return_type}(&__retVal_native)'
        return ''  # Return empty string if type is not in DEL_TYPESCAST_MAP

    # Initialize the output parts
    output_parts = []

    # Handle return type cleanup
    ret_type = method.get('retType', {})
    if is_obj_type(ret_type.get('type', '')):
        ret_cast = gen_return(ret_type)
        if ret_cast:
            output_parts.append(f'{tabs}{ret_cast}')

    # Handle parameter cleanup
    for param in method.get('paramTypes', []):
        param_cast = gen_param(param)
        if param_cast:
            output_parts.append(f'{tabs}{param_cast}')

    return output_parts


def gen_delegate_body(prototype: dict, delegates: set[str]) -> str:
    """
    Generates a Go-style function type definition.
    """
    # Extract delegate details
    delegate_name = prototype.get('name', 'UnnamedDelegate')
    delegate_description = prototype.get('description', '')

    # Check for duplicate delegates
    if delegate_name in delegates:
        return ''  # Skip if already generated

    # Add the delegate name to the set
    delegates.add(delegate_name)

    # Get the return type and convert it
    ret_type = prototype.get('retType', {})
    return_type = get_type_name(ret_type)

    # Get the parameter list
    param_list = gen_params(prototype, ParamGen.TypesNames)

    # Start building the Go function type definition
    delegate_code = []
    if delegate_description:
        delegate_code.append(f'// {delegate_name} - {delegate_description}')

    delegate_code.append(f'type {delegate_name} func({param_list}) {return_type}')

    # Join the list into a single formatted string
    return '\n'.join(delegate_code)


def gen_enum_body(enum: dict, enum_type: str, enums: set[str], used_names: set[str]) -> str:
    """
    Generates Go-style constant definitions for an enum.
    """
    # Get the enum details
    enum_name = enum.get('name', 'InvalidEnum')
    enum_description = enum.get('description', '')
    enum_values = enum.get('values', [])

    # Check for duplicate enums
    if enum_name in enums:
        return ''  # Skip if already generated

    # Add the enum name to the set
    enums.add(enum_name)

    underlying = convert_type(enum_type)

    # Start building the Go enum definition
    enum_code: list[str] = []

    # Add enum description
    if enum_description:
        enum_code.append(f'// {enum_name} - {enum_description}')

    enum_code.append(f'type {enum_name} = {underlying}\n')

    # Define constants
    enum_code.append('const (')

    local_seen: dict[str, int] = {}

    for i, value in enumerate(enum_values):
        raw_name = value.get('name', f'Value{i}')
        base_name = generate_name(raw_name)
        enum_val = value.get('value', str(i))
        descr = value.get('description', '')

        ## Counter for local repeats
        local_seen[base_name] = local_seen.get(base_name, 0) + 1
        is_local_duplicate = local_seen[base_name] > 1

        candidate = base_name

        def resolve_conflict(name: str) -> str:
            if name not in used_names:
                return name
            # trying with prefix
            pref = f'{enum_name}_{name}'
            if pref not in used_names:
                return pref
            # if the prefix is also occupied, add suffixes
            suffix = 2
            while True:
                trial = f'{pref}_{suffix}'
                if trial not in used_names:
                    return trial
                suffix += 1

        # If it is already globally occupied or a local duplicate, we run the logic of the conflict.
        if candidate in used_names or is_local_duplicate:
            candidate = resolve_conflict(candidate)

        # Registration
        used_names.add(candidate)

        if descr:
            enum_code.append(f'\t// {raw_name} - {descr}')
        enum_code.append(f'\t{candidate} {enum_name} = {enum_val}')

    enum_code.append(')')
    return '\n'.join(enum_code)


def gen_documentation(method: dict, tab_level: int = 0) -> str:
    """
    Generate a Go-style documentation comment from a JSON block with customizable tabulation.
    """
    # Extract general details
    name = method.get('name', 'UnnamedFunction')
    description = method.get('description', 'No description provided.')
    param_types = method.get('paramTypes', [])
    ret_type = method.get('retType', {}).get('type', 'void')

    # Determine tabulation
    tab = '\t' * tab_level

    # Start building the Go documentation comment
    docstring = [f'{tab}// {name} - {description}']

    # Add parameters
    for param in param_types:
        param_name = param.get('name', 'UnnamedParam')
        param_desc = param.get('description', 'No description available.')
        docstring.append(f'\n{tab}// @param {param_name}: {param_desc}')

    # Add return type description
    if ret_type.lower() != 'void':
        ret_desc = method.get('retType', {}).get('description', 'No description available.')
        docstring.append(f'\n{tab}// @return {ret_desc}')

    return ''.join(docstring)


def generate_method_body(method: dict, ret_type: dict, return_type: str) -> list[str]:
    """Generate the method body for a given method."""
    body = []  # List to hold the method body code
    indent = '\t'  # Default indentation
    inner_indent = indent  # Inner indentation for try/catch block

    # Check if the return type is an object and if it has a return value
    is_obj_ret = is_obj_type(ret_type['type'])
    is_pod_ret = is_pod_type(ret_type['type'])
    has_ret = not is_obj_ret and ret_type['type'] != 'void'

    # Generate and add parameter casting code
    index = 0
    params_cast = gen_paramscast(method, indent)

    # For object return types, declare a return value variable
    has_cast = len(params_cast) > 0 and ret_type['type'] != 'void'
    if is_obj_ret or has_cast:
        body.append(f'{indent}var __retVal {return_type}')

    has_try = False
    if len(params_cast) > 0:
        for params in params_cast:
            body.append(params)
        index = len(body)  # Mark position to insert try block
        body.append(f'{indent}plugify.Block {{')
        body.append(f'{indent}\tTry: func() {{')  # Start try block
        inner_indent = '\t\t\t'  # Adjust inner indentation
        has_try = has_cast

    # Generate the function call with marshaled parameters
    function_call = f'C.{method["name"]}({gen_params(method, ParamGen.CastNames)})'
    if is_obj_ret:
        body.append(f'{inner_indent}__native := {function_call}')
        body.append(f'{inner_indent}__retVal_native = *(*{RET_TYPESCAST_MAP[ret_type['type']]})(unsafe.Pointer(&__native))')
    elif has_try:
        if is_pod_ret:
            body.append(f'{inner_indent}__native := {function_call}')
            body.append(f'{inner_indent}__retVal = *(*{RET_TYPESCAST_MAP[ret_type['type']]})(unsafe.Pointer(&__native))')
        else:
            body.append(f'{inner_indent}__retVal = {RET_TYPESCAST_MAP[ret_type['type']]}({function_call})')
    elif has_ret:
        if is_pod_ret:
            body.append(f'{inner_indent}__native := {function_call}')
            body.append(f'{inner_indent}__retVal := *(*{RET_TYPESCAST_MAP[ret_type['type']]})(unsafe.Pointer(&__native))')
        elif ret_type['type'] == 'function':
            body.append(f'{inner_indent}__retVal := plugify.GetDelegateForFunctionPointer({function_call}, reflect.TypeOf({return_type}(nil))).({return_type})')
        else:
            body.append(f'{inner_indent}__retVal := {return_type}({function_call})')
    else:
        body.append(f'{inner_indent}{function_call}')

    # Unmarshal native data back into managed data
    assign_cast = gen_paramscast_assign(method, inner_indent)
    if len(assign_cast) > 0:
        body.append(f'{inner_indent}// Unmarshal - Convert native data to managed data.')
        for assign in assign_cast:
            body.append(assign)

    # Cleanup after function call
    cleanup_cast = gen_paramscast_cleanup(method, inner_indent)
    if len(cleanup_cast) > 0:
        body.append(f'{indent}\t}},\n{indent}\tFinally: func() {{\n{inner_indent}// Perform cleanup.')
        for cleanup in cleanup_cast:
            body.append(cleanup)
        body.append(f'{indent}\t}},\n{indent}}}.Do()')
    elif len(params_cast) > 0:
        # If no cleanup, remove try/catch block and adjust indentation
        body.pop(index)
        body.pop(index)
        for i in range(index, len(body)):
            body[i] = body[i][2:]  # Remove one level of indentation

    # Handle the return value for the method
    if ret_type['type'] != 'void':
        body.append(f'{indent}return __retVal')

    return body  # Return the generated method body


def generate_method_code(method: dict) -> str:
    """
    Generate the implementation code for a single method.
    """
    # Safely extract return type and determine if it uses a delegate prototype
    ret_type = method.get('retType', {})
    return_type = get_type_name(ret_type)

    method_name = method.get('name', 'UnnamedFunction')

    # Generate the method signature and body
    func_code = [
        gen_documentation(method),
        f'func {method_name}({gen_params(method, ParamGen.TypesNames)}) {return_type} {{',
        *generate_method_body(method, ret_type, return_type),
        '}\n'
    ]

    return '\n'.join(func_code)


def generate_delegate_code(pplugin: dict, delegates: set[str]) -> str:
    """
    Generate Go delegate code from a plugin definition.
    """
    # Container for all generated delegate code
    content = []

    def process_prototype(prototype: dict):
        """
        Generate delegate code from the given prototype if it hasn't been processed.
        """
        delegate_code = gen_delegate_body(prototype, delegates)
        if delegate_code:
            content.append(delegate_code)

    # Main loop: Process all exported methods in the plugin
    for method in pplugin.get('exportedMethods', []):
        # Check the return type for a delegate
        ret_type = method.get('retType', {})
        if 'prototype' in ret_type:
            process_prototype(ret_type['prototype'])

        # Check parameters for delegates
        for param in method.get('paramTypes', []):
            if 'prototype' in param:
                process_prototype(param['prototype'])

    content.append('')

    # Join all generated delegates into a single string
    return '\n'.join(content)


def generate_enum_code(pplugin: dict, enums: set[str]) -> str:
    """
    Iterates over all methods and collects enums.
    used_names — global set of already occupied identifiers (constants).
    """
    content: list[str] = []
    used_names: set[str] = set()

    def process_enum(e: dict, etype: str):
        code = gen_enum_body(e, etype, enums, used_names)
        if code:
            content.append(code)

    def process_prototype(proto: dict):
        # Return type
        if 'enum' in proto.get('retType', {}):
            rt = proto['retType']
            process_enum(rt['enum'], rt.get('type', ''))
        # Parameters
        for p in proto.get('paramTypes', []):
            if 'enum' in p:
                process_enum(p['enum'], p.get('type', ''))
            if 'prototype' in p:
                process_prototype(p['prototype'])

    for method in pplugin.get('exportedMethods', []):
        # Return type
        if 'retType' in method and 'enum' in method['retType']:
            rt = method['retType']
            process_enum(rt['enum'], rt.get('type', ''))
        # Parameters
        for param in method.get('paramTypes', []):
            if 'enum' in param:
                process_enum(param['enum'], param.get('type', ''))
            if 'prototype' in param:
                process_prototype(param['prototype'])

    content.append('')
    return '\n'.join(content)


def generate_cheader(plugin_name: str, pplugin: dict) -> str:
    content = ['#pragma once\n',
               '#include <stdlib.h>',
               '#include <stdint.h>',
               '#include <stdbool.h>\n',
               'typedef struct String { char* data; size_t size; size_t cap; } String;',
               'typedef struct Vector { void* begin; void* end; void* capacity; } Vector;',
               'typedef struct Vector2 { float x, y; } Vector2;',
               'typedef struct Vector3 { float x, y, z; } Vector3;',
               'typedef struct Vector4 { float x, y, z, w; } Vector4;',
               'typedef struct Matrix4x4 { float m[4][4]; } Matrix4x4;',
               'typedef struct Variant {',
               '\tunion {',
               '\tbool boolean;',
               '\tchar char8;',
               '\twchar_t char16;',
               '\tint8_t int8;',
               '\tint16_t int16;',
               '\tint32_t int32;',
               '\tint64_t int64;',
               '\tuint8_t uint8;',
               '\tuint16_t uint16;',
               '\tuint32_t uint32;',
               '\tuint64_t uint64;',
               '\tvoid* ptr;',
               '\tfloat flt;',
               '\tdouble dbl;',
               '\tString str;',
               '\tVector vec;',
               '\tVector2 vec2;',
               '\tVector3 vec3;',
               '\tVector4 vec4;',
               '\t};',
               '#if INTPTR_MAX == INT32_MAX',
               '\tvolatile char pad[8];',
               '#endif',
               '\tuint8_t current;',
               '} Variant;'
               '\n',
               'extern void* Plugify_GetMethodPtr(const char* methodName);',
               'extern void Plugify_GetMethodPtr2(const char* methodName, void** addressDest);\n'
    ]

    # Append method implementations
    for method in pplugin.get('exportedMethods', []):
        method_name = method.get('name', 'UnnamedMethod')
        #param_types_data = method.get('paramTypes', [])
        ret_type_data = method.get('retType', {})

        ret_type = get_ctype_name(ret_type_data, True)
        param_list = gen_cparams(method, ParamGen.TypesNames)
        param_types = gen_cparams(method, ParamGen.Types)
        param_names = gen_cparams(method, ParamGen.Names)

        content.append(
            f'static {ret_type} {method_name}({param_list}) {{\n'
            f'\ttypedef {ret_type} (*{method_name}Fn)({param_types});\n'
            f'\tstatic {method_name}Fn __func = NULL;\n'
            f'\tif (__func == NULL) '
            f'Plugify_GetMethodPtr2("{plugin_name}.{method_name}", (void**)&__func);\n'
            f'\t{"return " if ret_type_data.get("type", "void") != "void" else ""}__func({param_names});\n'
            '}\n'
        )

    # Join and return the complete content as a single string
    return '\n'.join(content)


def generate_header(plugin_name: str, pplugin: dict) -> str:
    """Generate the full header content for the plugin."""
    # GitHub link for reference
    link = 'https://github.com/untrustedmodders/plugify-module-golang/blob/main/generator/generator.py'

    directives = []

    # Append directives implementations
    for method in pplugin.get('exportedMethods', []):
        directives.append(f'#cgo noescape {method.get("name", "UnnamedMethod")}')
        #languageModule = pplugin.get('languageModule', {})
        #if languageModule.get('name', '') == 'golang':
        #   directives.append(f'#cgo noescape {method.get("name", "UnnamedMethod")}')

    # Initialize content with common imports and namespace declaration
    content = [
        f'package {plugin_name}',
        '',
        '/*',
        f'#include "{plugin_name}.h"',
        '\n'.join(directives),
        '*/',
        'import "C"',
        'import (',
        '\t"unsafe"',
        '\t_ "reflect"',
        '\t"github.com/untrustedmodders/go-plugify"',
        ')',
        '',
        f'// Generated with {link} from {plugin_name}',
        '',
    ]

    # Append enum definitions
    enums = set()
    content.append(generate_enum_code(pplugin, enums))

    # Append delegate definitions
    delegates = set()
    content.append(generate_delegate_code(pplugin, delegates))

    # Append method implementations
    for method in pplugin.get('exportedMethods', []):
        content.append(generate_method_code(method))

    # Join and return the complete content as a single string
    return '\n'.join(content)


def main(manifest_path: str, output_dir: str, override: bool):
    """Main function to process the plugin and generate the Go header file."""
    # Validate inputs
    if not os.path.isfile(manifest_path):
        print(f'Manifest file does not exist: {manifest_path}')
        return 1
    if not os.path.isdir(output_dir):
        print(f'Output folder does not exist: {output_dir}')
        return 1

    # Determine plugin name and output file path
    plugin_name = os.path.basename(manifest_path).rsplit('.', 3)[0]
    output_path = os.path.join(output_dir, plugin_name, f'{plugin_name}.go')
    output_path2 = os.path.join(output_dir, plugin_name, f'{plugin_name}.h')
    os.makedirs(os.path.dirname(output_path), exist_ok=True)

    # Handle existing file
    if os.path.isfile(output_path) and not override:
        print(f'Output file already exists: {output_path}. Use --override to overwrite existing file.')
        return 1

    try:
        # Read and parse manifest
        with open(manifest_path, 'r', encoding='utf-8') as file:
            pplugin = json.load(file)

        # Generate header content
        content = generate_header(plugin_name, pplugin)

        # Write content to file
        with open(output_path, 'w', encoding='utf-8') as fd:
            fd.write(''.join(content))

        # Generate c header content
        content = generate_cheader(plugin_name, pplugin)

        # Write content to file
        with open(output_path2, 'w', encoding='utf-8') as fd:
            fd.write(content)

    except Exception as e:
        print(f'An error occurred: {e}')
        return 1

    print(f'Header generated at: {output_path}')
    return 0


def get_args():
    """Parse command-line arguments."""
    parser = argparse.ArgumentParser(description='Generate Go .go include files for plugin manifests.')
    parser.add_argument('manifest', help='Path to the plugin manifest file')
    parser.add_argument('output', help='Path to the output directory')
    parser.add_argument('--override', action='store_true', help='Override existing files')
    return parser.parse_args()


if __name__ == '__main__':
    args = get_args()
    sys.exit(main(args.manifest, args.output, args.override))
