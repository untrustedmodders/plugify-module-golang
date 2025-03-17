#pragma once

#include <module_export.h>
#include <plugify/assembly.hpp>
#include <plugify/jit/callback.hpp>
#include <plugify/language_module.hpp>
#include <plugify/module.hpp>

#include <asmjit/asmjit.h>
#include <cpptrace/cpptrace.hpp>

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef void* GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;

typedef struct { const char* p; GoInt n; } GoString;
typedef void* GoMap;
typedef void* GoChan;
typedef struct { void* t; void* v; } GoInterface;
typedef struct { void* data; GoInt len; GoInt cap; } GoSlice;

namespace golm {
	struct string_hash {
		using is_transparent = void;
		[[nodiscard]] size_t operator()(const char* txt) const {
			return std::hash<std::string_view>{}(txt);
		}
		[[nodiscard]] size_t operator()(std::string_view txt) const {
			return std::hash<std::string_view>{}(txt);
		}
		[[nodiscard]] size_t operator()(const std::string& txt) const {
			return std::hash<std::string>{}(txt);
		}
	};

	constexpr int kApiVersion = 1;

	struct PluginContext {
		bool hasUpdate{};
		bool hasStart{};
		bool hasEnd{};
		bool hasPanic{};
	};

	using InitFunc = int (*)(GoSlice, int, const void*);
	using StartFunc = void (*)();
	using UpdateFunc = void (*)(float);
	using EndFunc = void (*)();
	using ContextFunc = PluginContext* (*)();

	struct AssemblyHolder {
		std::unique_ptr<plugify::Assembly> assembly;
		UpdateFunc updateFunc;
		StartFunc startFunc;
		EndFunc endFunc;
		ContextFunc contextFunc;
		plugify::JitCallback::CallbackHandler callFunc;
	};

	class GoLanguageModule final : public plugify::ILanguageModule {
	public:
		GoLanguageModule() = default;
		~GoLanguageModule() = default;

		// ILanguageModule
		plugify::InitResult Initialize(std::weak_ptr<plugify::IPlugifyProvider> provider, plugify::ModuleHandle module) override;
		void Shutdown() override;
		void OnUpdate(plugify::DateTime dt) override;
		plugify::LoadResult OnPluginLoad(plugify::PluginHandle plugin) override;
		void OnPluginStart(plugify::PluginHandle plugin) override;
		void OnPluginUpdate(plugify::PluginHandle plugin, plugify::DateTime dt) override;
		void OnPluginEnd(plugify::PluginHandle plugin) override;
		void OnMethodExport(plugify::PluginHandle plugin) override;
		bool IsDebugBuild() override;

		const std::shared_ptr<plugify::IPlugifyProvider>& GetProvider() { return _provider; }
		const std::shared_ptr<asmjit::JitRuntime>& GetRuntime() { return _rt; }

		plugify::MemAddr GetNativeMethod(std::string_view methodName) const;
		void GetNativeMethod(std::string_view methodName, plugify::MemAddr* addressDest);
		plugify::MethodHandle FindMethod(std::string_view name);

	private:
		std::shared_ptr<plugify::IPlugifyProvider> _provider;
		std::shared_ptr<asmjit::JitRuntime> _rt;

		std::vector<std::unique_ptr<AssemblyHolder>> _assemblies;
		std::unordered_map<std::string, plugify::MemAddr, string_hash, std::equal_to<>> _nativesMap;

		std::vector<plugify::MemAddr*> _addresses;

		static const std::array<void*, 139> _pluginApi;
	};

	extern GoLanguageModule g_golm;
}

extern "C" GOLM_EXPORT plugify::ILanguageModule* GetLanguageModule();
