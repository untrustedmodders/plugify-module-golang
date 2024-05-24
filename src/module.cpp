#include "module.h"
#include <module_export.h>

using namespace plugify;

namespace golm {
	InitResult GoLanguageModule::Initialize(std::weak_ptr<IPlugifyProvider> provider, const IModule& module) {
		// TODO: implement
		return InitResultData{};
	}

	void GoLanguageModule::Shutdown() {
		// TODO: implement
	}

	LoadResult GoLanguageModule::OnPluginLoad(const IPlugin& plugin) {
		// TODO: implement
		return ErrorData{ "Loading not implemented" };
	}

	void GoLanguageModule::OnPluginStart(const IPlugin& plugin) {
		// TODO: implement
	}

	void GoLanguageModule::OnPluginEnd(const IPlugin& plugin) {
		// TODO: implement
	}

	void GoLanguageModule::OnMethodExport(const IPlugin& plugin) {
		// TODO: implement
	}

	GoLanguageModule g_golm;

	extern "C"
	GOLM_EXPORT ILanguageModule* GetLanguageModule() {
		return &g_golm;
	}
}
