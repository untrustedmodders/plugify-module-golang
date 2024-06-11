#include "assembly.h"

#if GOLM_PLATFORM_WINDOWS
#include <windows.h>
#elif GOLM_PLATFORM_LINUX
#include <dlfcn.h>
#elif GOLM_PLATFORM_APPLE
#include <dlfcn.h>
#else
#error "Platform is not supported!"
#endif

namespace golm {
	thread_local static std::string lastError;

	std::unique_ptr<Assembly> Assembly::LoadFromPath(const std::filesystem::path& assemblyPath) {
#if GOLM_PLATFORM_WINDOWS
		void* handle = static_cast<void*>(LoadLibraryW(assemblyPath.c_str()));
#elif GOLM_PLATFORM_LINUX || GOLM_PLATFORM_APPLE
		void* handle = dlopen(assemblyPath.string().c_str(), RTLD_LAZY);
#else
		void* handle = nullptr;
#endif
		if (handle) {
			return std::unique_ptr<Assembly>(new Assembly(handle));
		}
#if GOLM_PLATFORM_WINDOWS
		uint32_t errorCode = GetLastError();
		if (errorCode != 0) {
			LPSTR messageBuffer = nullptr;
			size_t size = FormatMessageA(FORMAT_MESSAGE_ALLOCATE_BUFFER | FORMAT_MESSAGE_FROM_SYSTEM | FORMAT_MESSAGE_IGNORE_INSERTS,
				NULL, errorCode, MAKELANGID(LANG_NEUTRAL, SUBLANG_DEFAULT), reinterpret_cast<LPSTR>(&messageBuffer), 0, NULL);
			lastError = std::string(messageBuffer, size);
			LocalFree(messageBuffer);
		}
#elif GOLM_PLATFORM_LINUX || GOLM_PLATFORM_APPLE
		lastError = dlerror();
#endif
		return nullptr;
	}

	std::string Assembly::GetError() {
		return lastError;
	}

	Assembly::Assembly(void* handle) : _handle{handle} {
	}

	Assembly::~Assembly() {
#if GOLM_PLATFORM_WINDOWS
		FreeLibrary(static_cast<HMODULE>(_handle));
#elif GOLM_PLATFORM_LINUX || GOLM_PLATFORM_APPLE
		dlclose(_handle);
#endif
	}

	void* Assembly::GetFunction(const char* functionName) const {
#if GOLM_PLATFORM_WINDOWS
		return reinterpret_cast<void*>(GetProcAddress(static_cast<HMODULE>(_handle), functionName));
#elif GOLM_PLATFORM_LINUX || GOLM_PLATFORM_APPLE
		return dlsym(_handle, functionName);
#else
		return nullptr;
#endif
	}
}
