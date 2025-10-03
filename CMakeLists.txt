cmake_minimum_required(VERSION 3.14 FATAL_ERROR)

if(POLICY CMP0092)
    cmake_policy(SET CMP0092 NEW) # Don't add -W3 warning level by default.
endif()


file(READ "${CMAKE_CURRENT_SOURCE_DIR}/version.txt" VERSION_FILE_CONTENTS)
string(STRIP "${VERSION_FILE_CONTENTS}" VERSION_FILE_CONTENTS)
set(GOLM_VERSION "${VERSION_FILE_CONTENTS}" CACHE STRING "Set version name")
set(GOLM_PACKAGE "golang_module" CACHE STRING "Set package name")
string(REPLACE "v" "" GOLM_VERSION "${GOLM_VERSION}")
string(REGEX REPLACE "[.+-]" ";" GOLM_VERSION_LIST ${GOLM_VERSION})
list(GET GOLM_VERSION_LIST 0 GOLM_VERSION_MAJOR)
list(GET GOLM_VERSION_LIST 1 GOLM_VERSION_MINOR)
list(GET GOLM_VERSION_LIST 2 GOLM_VERSION_PATCH)

project(plugify-module-golang 
        VERSION "${GOLM_VERSION_MAJOR}.${GOLM_VERSION_MINOR}.${GOLM_VERSION_PATCH}"
        DESCRIPTION "Go Language Module" 
        HOMEPAGE_URL "https://github.com/untrustedmodders/plugify-module-golang" 
        LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 23)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)
#set(CMAKE_POSITION_INDEPENDENT_CODE ON)

if(NOT CMAKE_BUILD_TYPE MATCHES "Debug|Devel|MinSizeRel|RelWithDebInfo|Release")
    message(STATUS "CMAKE_BUILD_TYPE not set, defaulting to Release.")
    set(CMAKE_BUILD_TYPE Release)
endif()

if(UNIX AND NOT APPLE)
    set(LINUX TRUE)
endif()

set(BUILD_SHARED_LIBS ON CACHE INTERNAL "")

#
# Plugify
#
set(PLUGIFY_BUILD_SHARED_LIB ON CACHE INTERNAL "")
set(PLUGIFY_BUILD_TESTS OFF CACHE INTERNAL "")
if(LINUX)
    set(PLUGIFY_USE_STATIC_STDLIB ON CACHE INTERNAL "")
    set(PLUGIFY_USE_ABI0 OFF CACHE INTERNAL "")
endif()
set(PLUGIFY_USE_SANITIZER ON CACHE INTERNAL "")
add_subdirectory(external/plugify)

#
# Go Language Module for Plugify
#
file(GLOB_RECURSE GOLM_SOURCES RELATIVE ${CMAKE_CURRENT_SOURCE_DIR} "src/*.cpp")
set(GOLM_PCH_FILE "src/pch.hpp")

add_library(${PROJECT_NAME} SHARED ${GOLM_SOURCES})
target_include_directories(${PROJECT_NAME} PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/include)

set(GOLM_LINK_LIBRARIES plugify::plugify)

if(NOT COMPILER_SUPPORTS_FORMAT)
    set(GOLM_LINK_LIBRARIES ${GOLM_LINK_LIBRARIES} fmt::fmt-header-only)
endif()

target_link_libraries(${PROJECT_NAME} PRIVATE ${GOLM_LINK_LIBRARIES})
target_precompile_headers(${PROJECT_NAME} PRIVATE ${GOLM_PCH_FILE})

if(MSVC)
    target_compile_options(${PROJECT_NAME} PRIVATE /W4 /WX /Zc:preprocessor $<$<CONFIG:Release>:/Zi>)
    target_link_options(${PROJECT_NAME} PRIVATE $<$<CONFIG:Release>:/DEBUG /OPT:REF /OPT:ICF>)
else()
    target_compile_options(${PROJECT_NAME} PRIVATE -Wextra -Wconversion -Werror -Wshadow -Wpedantic $<$<CONFIG:Release>:-g1>)
endif()

if(APPLE)
    target_link_options(${PROJECT_NAME} PRIVATE "-Wl,-exported_symbols_list,${CMAKE_CURRENT_SOURCE_DIR}/sym/exported_symbols.lds")
elseif(UNIX)
    target_link_options(${PROJECT_NAME} PRIVATE "-Wl,--version-script,${CMAKE_CURRENT_SOURCE_DIR}/sym/version_script.lds")
endif()

if(PLUGIFY_USE_SANITIZER)
    if(MSVC)
        target_compile_options(${PROJECT_NAME} PRIVATE $<$<CONFIG:Debug>: /fsanitize=address>)
        target_compile_definitions(${PROJECT_NAME} PRIVATE $<$<CONFIG:Debug>: _DISABLE_VECTOR_ANNOTATION>)
    else()
        target_compile_options(${PROJECT_NAME} PRIVATE $<$<CONFIG:Debug>: -fsanitize=address -fsanitize=leak -fsanitize=undefined>)
        target_link_libraries(${PROJECT_NAME} PRIVATE $<$<CONFIG:Debug>: -fsanitize=address -fsanitize=leak -fsanitize=undefined>)
    endif()
endif()

if(LINUX)
    target_link_libraries(${PROJECT_NAME} PRIVATE -lstdc++exp)
endif()

include(GenerateExportHeader)
generate_export_header(${PROJECT_NAME} EXPORT_MACRO_NAME GOLM_EXPORT EXPORT_FILE_NAME ${CMAKE_BINARY_DIR}/exports/module_export.h)
target_include_directories(${PROJECT_NAME} PRIVATE ${CMAKE_BINARY_DIR}/exports)

target_compile_definitions(${PROJECT_NAME} PRIVATE
        GOLM_PLATFORM_WINDOWS=$<BOOL:${WIN32}>
        GOLM_PLATFORM_APPLE=$<BOOL:${APPLE}>
        GOLM_PLATFORM_LINUX=$<BOOL:${LINUX}>
        GOLM_ARCH_ARM=$<BOOL:${PLUGIFY_USE_ARM}>
        GOLM_IS_DEBUG=$<STREQUAL:${CMAKE_BUILD_TYPE},Debug>
        GOLM_LIBRARY_SUFFIX="${CMAKE_SHARED_LIBRARY_SUFFIX}"
        GOLM_LIBRARY_PREFIX="${CMAKE_SHARED_LIBRARY_PREFIX}"
)

configure_file(
    ${CMAKE_CURRENT_SOURCE_DIR}/${PROJECT_NAME}.pmodule.in
    ${CMAKE_CURRENT_BINARY_DIR}/${PROJECT_NAME}.pmodule
)