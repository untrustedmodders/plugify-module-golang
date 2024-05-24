#include "assembly.h"
#include "module.h"
#include <plugify/plugin_descriptor.h>
#include <plugify/plugin.h>
#include <plugify/plugify_provider.h>
#include <module_export.h>

using namespace plugify;
namespace fs = std::filesystem;

namespace golm {
	GoLanguageModule::GoLanguageModule() = default;

	InitResult GoLanguageModule::Initialize(std::weak_ptr<IPlugifyProvider> provider, const IModule& module) {
		if (!(_provider = provider.lock())) {
			return ErrorData{ "Provider not exposed" };
		}
		return InitResultData{};
	}

	void GoLanguageModule::Shutdown() {
		_assemblyMap.clear();
		_provider.reset();
	}

	LoadResult GoLanguageModule::OnPluginLoad(const IPlugin& plugin) {
		const auto entryPoint = fs::path(plugin.GetDescriptor().entryPoint);
		fs::path assemblyPath = plugin.GetBaseDir() / entryPoint.parent_path() / std::format(BINARY_MODULE_PREFIX "{}" BINARY_MODULE_SUFFIX, entryPoint.filename().string());

		auto assembly = Assembly::LoadFromPath(assemblyPath);
		if (!assembly) {
			return ErrorData{ std::format("Failed to load assembly: {}", Assembly::GetError()) };
		}

		bool funcFail = false;
		std::vector<std::string_view> funcErrors;

		auto* const startFunc = assembly->GetFunction<StartFunc>("Plugify_PluginStart");
		if (!startFunc) {
			funcFail = true;
			funcErrors.emplace_back("Plugify_PluginStart");
		}

		auto* const endFunc = assembly->GetFunction<EndFunc>("Plugify_PluginEnd");
		if (!endFunc) {
			funcFail = true;
			funcErrors.emplace_back("Plugify_PluginEnd");
		}

		if (funcFail) {
			std::string funcs(funcErrors[0]);
			for (auto it = std::next(funcErrors.begin()); it != funcErrors.end(); ++it) {
				std::format_to(std::back_inserter(funcs), ", {}", *it);
			}
			return ErrorData{ std::format("Not found {} function(s)", funcs) };
		}

		funcFail = false;
		funcErrors.clear();

		const auto& exportedMethods = plugin.GetDescriptor().exportedMethods;
		std::vector<MethodData> methods;
		methods.reserve(exportedMethods.size());

		for (const auto& method : exportedMethods) {
			if (auto* const func = assembly->GetFunction(method.funcName.c_str())) {
				methods.emplace_back(method.name, func);
			}
			else {
				funcFail = true;
				funcErrors.emplace_back(method.name);
			}
		}
		if (funcFail) {
			std::string funcs(funcErrors[0]);
			for (auto it = std::next(funcErrors.begin()); it != funcErrors.end(); ++it) {
				std::format_to(std::back_inserter(funcs), ", {}", *it);
			}
			return ErrorData{ std::format("Not found {} method function(s)", funcs) };
		}

		const auto [_, result] = _assemblyMap.try_emplace(plugin.GetName(), std::move(assembly), startFunc, endFunc);
		if (!result) {
			return ErrorData{ std::format("Plugin name duplicate") };
		}

		return LoadResultData{ std::move(methods) };
	}

	void GoLanguageModule::OnPluginStart(const IPlugin& plugin) {
		if (const auto it = _assemblyMap.find(plugin.GetName()); it != _assemblyMap.end()) {
			const auto& assemblyHolder = std::get<AssemblyHolder>(*it);
			assemblyHolder.GetStartFunc()();
		}
	}

	void GoLanguageModule::OnPluginEnd(const IPlugin& plugin) {
		if (const auto it = _assemblyMap.find(plugin.GetName()); it != _assemblyMap.end()) {
			const auto& assemblyHolder = std::get<AssemblyHolder>(*it);
			assemblyHolder.GetEndFunc()();
		}
	}

	void GoLanguageModule::OnMethodExport(const IPlugin& plugin) {
		// TODO: implement
	}

	GoLanguageModule::AssemblyHolder::AssemblyHolder(std::unique_ptr<Assembly> assembly, StartFunc startFunc, EndFunc endFunc) : _assembly{ std::move(assembly) }, _startFunc{ startFunc }, _endFunc{ endFunc } {
	}

	StartFunc GoLanguageModule::AssemblyHolder::GetStartFunc() const {
		return _startFunc;
	}

	EndFunc GoLanguageModule::AssemblyHolder::GetEndFunc() const {
		return _endFunc;
	}

	GoLanguageModule g_golm;

	extern "C"
	GOLM_EXPORT ILanguageModule* GetLanguageModule() {
		return &g_golm;
	}
}
