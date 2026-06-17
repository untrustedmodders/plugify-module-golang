#pragma once

#include <map>

#include <plugify/assembly.hpp>
#include <plugify/callback.hpp>
#include <plugify/language_module.hpp>
#include <plugify/logger.hpp>
#include <plugify/profiler.hpp>
#include <plugify/assembly_loader.hpp>
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
	const char* p{};
	GoInt n{};

	GoString() = default;

	explicit GoString(std::string_view sv)
		: p(sv.data())
		, n(static_cast<GoInt>(sv.size()))
	{}

	operator std::string_view() const { return {p, static_cast<size_t>(n)};  }
	operator bool() const { return n > 0;  }
};

using GoMap = void*;
using GoChan = void*;

struct GoInterface {
	void* t{};
	void* v{};
};

struct GoSlice {
	void* data{};
	GoInt len{};
	GoInt cap{};

	GoSlice() = default;

	template<typename T, size_t N>
	explicit GoSlice(std::span<const T, N> sp)
		: data(const_cast<T*>(sp.data()))
		, len(static_cast<GoInt>(sp.size()))
		, cap(static_cast<GoInt>(sp.size()))
	{}

	template<typename T>
	operator std::span<const T>() const { return { static_cast<T*>(data), static_cast<size_t>(len)}; }
	operator bool() const { return len > 0;  }
};

using namespace plugify;

namespace golm {
	constexpr GoInt kApiVersion = 3;

	enum class PluginCode { Ok, Failed };

	struct PluginResult {
		PluginCode code{};
		plg::string	message{};

		explicit operator bool() const noexcept { return code == PluginCode::Ok; }
		operator std::string_view() const noexcept { return message; }
	};

	struct PluginContext {
		bool hasUpdate{};
		bool hasStart{};
		bool hasEnd{};
	};

	using InitFunc = GoInt (*)(GoSlice, GoInt, GoString);
	using ShutdownFunc = void (*)();
	using StartFunc = PluginResult (*)(GoString);
	using UpdateFunc = PluginResult (*)(GoString, GoFloat32);
	using EndFunc = PluginResult (*)(GoString);
	using ContextFunc = PluginContext (*)(GoString);
	using CallbackFunc = void (*)(const Method* method, Address data, uint64_t* params, size_t count, /*uint128_t*/ void* ret);

	using OpenFunc = GoInt (*)(GoString);
	using BindFunc = bool (*)(GoInt, GoInt, GoString, GoString);
	using CallFunc = bool (*)(GoInt, GoString);
	using FindFunc = bool (*)(GoInt, GoString);
	using ErrorFunc = GoString (*)();

	struct Symbols {
		StartFunc startFunc{};
		UpdateFunc updateFunc{};
		EndFunc endFunc{};
		ContextFunc contextFunc{};
		CallbackFunc callbackFunc{};

		InitFunc initFunc{};
		ShutdownFunc shutdownFunc{};
	};

	struct RuntimeSymbols : Symbols {
		OpenFunc openFunc{};
		BindFunc bindFunc{};
		CallFunc callFunc{};
		FindFunc findFunc{};
		ErrorFunc errorFunc{};
	};

	struct AssemblyHolder {
		std::shared_ptr<IAssembly> assembly;
		Symbols symbols;
		GoInt id; // only valid for -buildmode=plugin
	};

	using AssemblyMap = std::map<UniqueId, AssemblyHolder>;

	class GoLanguageModule final : public ILanguageModule {
	public:
		GoLanguageModule() = default;
		~GoLanguageModule() = default;

		// ILanguageModule
		Result<InitData> Initialize(const Provider& provider, const Extension& module) override;
		Result<void> Shutdown() override;
		Result<void> OnUpdate(std::chrono::milliseconds dt) override;

		Result<LoadData> OnPluginLoad(const Extension& plugin) override;
		Result<void> OnPluginStart(const Extension& plugin) override;
		Result<void> OnPluginUpdate(const Extension& plugin, std::chrono::milliseconds dt) override;
		Result<void> OnPluginEnd(const Extension& plugin) override;
		Result<void> OnMethodExport(const Extension& plugin) override;

		bool IsDebugBuild() const noexcept override;

		const std::unique_ptr<Provider>& GetProvider() { return _provider; }
		const std::shared_ptr<ILogger>& GetLogger() { return _logger; }
		const std::shared_ptr<IProfiler>& GetProfiler() const { return _profiler; }

		const AssemblyMap& GetAssemblies() const { return _assemblies; }
		const AssemblyHolder* FindAssembly(UniqueId pluginId) const;
		const Extension* FindExtension(std::string_view name) const;
		std::shared_ptr<Method> FindMethod(std::string_view name) const;

	private:
		std::unique_ptr<Provider> _provider;
		std::shared_ptr<ILogger> _logger;
		std::shared_ptr<IAssemblyLoader> _loader;
		std::shared_ptr<IProfiler> _profiler;
		std::shared_ptr<IAssembly> _runtime;
		RuntimeSymbols _symbols;
		AssemblyMap _assemblies;

		static const std::array<void*, 140> _pluginApi;
	};

	extern GoLanguageModule g_golm;
}

extern "C" GOLM_EXPORT ILanguageModule* GetLanguageModule();
