#pragma once

#include <asmjit/asmjit.h>
#include <module_export.h>
#include <plugify/assembly.hpp>
#include <plugify/jit/callback.hpp>
#include <plugify/language_module.hpp>
#include <plugify/module.hpp>

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

extern "C" {
	typedef struct DCCallVM_ DCCallVM;
	typedef struct DCaggr_ DCaggr;
}

template <>
struct std::default_delete<DCaggr> {
	void operator()(DCaggr* p) const;
};

template <>
struct std::default_delete<DCCallVM> {
	void operator()(DCCallVM* p) const;
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

	using InitFunc = int (*)(GoSlice, int, const void*);
	using StartFunc = void (*)();
	using EndFunc = void (*)();

	class AssemblyHolder {
	public:
		AssemblyHolder(std::unique_ptr<plugify::Assembly> assembly, StartFunc startFunc, EndFunc endFunc) : _assembly{std::move(assembly)}, _startFunc{startFunc}, _endFunc{endFunc} {}

		[[nodiscard]] StartFunc GetStartFunc() const { return _startFunc; }
		[[nodiscard]] EndFunc GetEndFunc() const { return _endFunc; }

	private:
		std::unique_ptr<plugify::Assembly> _assembly;
		StartFunc _startFunc;
		EndFunc _endFunc;
	};

	struct VirtualMachine {
		[[nodiscard]] DCCallVM& operator()();
	private:
		std::unique_ptr<DCCallVM> _callVirtMachine;
	};

	//using MethodRef = std::reference_wrapper<const plugify::Method>;
	using ArgumentList = std::vector<GoSlice>;
	using AggrList = std::vector<std::unique_ptr<DCaggr>>;
	using StringHolder = std::vector<std::unique_ptr<GoString[]>>;
	using BoolHolder = std::vector<std::unique_ptr<bool[]>>;

	class GoLanguageModule final : public plugify::ILanguageModule {
	public:
		GoLanguageModule() = default;
		~GoLanguageModule() = default;

		// ILanguageModule
		plugify::InitResult Initialize(std::weak_ptr<plugify::IPlugifyProvider> provider, plugify::ModuleRef module) override;
		void Shutdown() override;
		plugify::LoadResult OnPluginLoad(plugify::PluginRef plugin) override;
		void OnPluginStart(plugify::PluginRef plugin) override;
		void OnPluginEnd(plugify::PluginRef plugin) override;
		void OnMethodExport(plugify::PluginRef plugin) override;
		bool IsDebugBuild() override;

		const std::shared_ptr<plugify::IPlugifyProvider>& GetProvider() { return _provider; }
		plugify::MemAddr GetNativeMethod(std::string_view methodName) const;
		void GetNativeMethod(std::string_view methodName, plugify::MemAddr* addressDest);

	private:
		static void InternalCall(plugify::MethodRef method, plugify::MemAddr data, const plugify::JitCallback::Parameters* params, uint8_t count, const plugify::JitCallback::Return* ret);

	private:
		std::shared_ptr<asmjit::JitRuntime> _rt;
		std::shared_ptr<plugify::IPlugifyProvider> _provider;

		std::map<plugify::UniqueId, AssemblyHolder> _assemblyMap;
		std::unordered_map<std::string, plugify::MemAddr, string_hash, std::equal_to<>> _nativesMap;

		std::vector<plugify::JitCallback> _functions;
		std::vector<plugify::MemAddr*> _addresses;

		static const std::array<void*, 28> _pluginApi;
	};

	extern GoLanguageModule g_golm;
}

extern "C" GOLM_EXPORT plugify::ILanguageModule* GetLanguageModule();
