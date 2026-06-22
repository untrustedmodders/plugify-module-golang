#pragma once

#include <stddef.h>

extern char error_buffer[1024];

typedef struct String { const char* p; ptrdiff_t n; } String;