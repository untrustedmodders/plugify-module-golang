#pragma once

#include <plugify/assembly.hpp>
#include <plugify/callback.hpp>
#include <plugify/language_module.hpp>
#include <plugify/extension.hpp>

#include <module_export.h>

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
	operator bool() const { return n > 0;  }
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
	operator bool() const { return len > 0;  }
};

using namespace plugify;

namespace golm {
	constexpr int kApiVersion = 1;

	struct PluginContext {
		bool hasUpdate{};
		bool hasStart{};
		bool hasEnd{};
		bool hasPanic{};
	};

	using InitFunc = int (*)(GoSlice, int, const void*);
	using CallFunc = JitCallback::CallbackHandler;
	using StartFunc = void (*)();
	using UpdateFunc = void (*)(float);
	using EndFunc = void (*)();
	using ContextFunc = PluginContext* (*)();

	struct AssemblyHolder {
		std::shared_ptr<IAssembly> assembly;
		UpdateFunc updateFunc;
		StartFunc startFunc;
		EndFunc endFunc;
		ContextFunc contextFunc;
		CallFunc callFunc;
	};

	class GoLanguageModule final : public ILanguageModule {
	public:
		GoLanguageModule() = default;
		~GoLanguageModule() = default;

		// ILanguageModule
		Result<InitData> Initialize(const Provider& provider, const Extension& module) override;
		void Shutdown() override;
		void OnUpdate(std::chrono::milliseconds dt) override;
		Result<LoadData> OnPluginLoad(const Extension& plugin) override;
		void OnPluginStart(const Extension& plugin) override;
		void OnPluginUpdate(const Extension& plugin, std::chrono::milliseconds dt) override;
		void OnPluginEnd(const Extension& plugin) override;
		void OnMethodExport(const Extension& plugin) override;
		bool IsDebugBuild() override;

		const std::unique_ptr<Provider>& GetProvider() { return _provider; }

		MemAddr GetNativeMethod(std::string_view methodName) const;
		void GetNativeMethod(std::string_view methodName, MemAddr* addressDest);
		const Method* FindMethod(std::string_view name);

	private:
		std::unique_ptr<Provider> _provider;

		std::vector<std::unique_ptr<AssemblyHolder>> _assemblies;
		std::unordered_map<std::string, MemAddr, plg::string_hash, std::equal_to<>> _nativesMap;

		std::vector<MemAddr*> _addresses;

		static const std::array<void*, 140> _pluginApi;
	};

	extern GoLanguageModule g_golm;
}

extern "C" GOLM_EXPORT ILanguageModule* GetLanguageModule();
