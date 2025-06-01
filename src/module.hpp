#pragma once

#include <module_export.h>
#include <plugify/assembly.hpp>
#include <plugify/jit/callback.hpp>
#include <plugify/language_module.hpp>
#include <plugify/module.hpp>

#include <asmjit/asmjit.h>
#include <cpptrace/cpptrace.hpp>

using GoInt8 = signed char;
using GoUint8 = unsigned char;
using GoInt16 = short;
using GoUint16 = unsigned short;
using GoInt32 = int;
using GoUint32 = unsigned int;
using GoInt64 = long long;
using GoUint64 = unsigned long long;
#if INTPTR_MAX == INT64_MAX
using GoInt = GoInt64;
using GoUint = GoUint64;
#elif INTPTR_MAX == INT32_MAX
using GoInt = GoInt32 ;
using GoUint = GoUint32;
#endif // INTPTR_MAX
using GoUintptr = void*;
using GoFloat32 = float;
using GoFloat64 = double;

struct GoString {
	const char* p;
	GoInt n;

	operator std::string_view() const { return {p, static_cast<size_t>(n)};  }
};

using GoMap = void*;
using GoChan = void*;

struct GoInterface {
	void* t;
	void* v;
};

struct GoSlice {
	void* data;
	GoInt len;
	GoInt cap;

	template<typename T>
	operator std::span<T>() const { return { static_cast<T*>(data), static_cast<size_t>(len)}; }
};

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
