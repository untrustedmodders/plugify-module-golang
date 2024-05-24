#pragma once

#include <plugify/language_module.h>
#include <unordered_map>

namespace golm {
	using StartFunc = void (*)();
	using EndFunc = void (*)();

	class Assembly;

	class GoLanguageModule final : public plugify::ILanguageModule {
	public:
		GoLanguageModule();

		// ILanguageModule
		plugify::InitResult Initialize(std::weak_ptr<plugify::IPlugifyProvider> provider, const plugify::IModule& module) override;
		void Shutdown() override;
		plugify::LoadResult OnPluginLoad(const plugify::IPlugin& plugin) override;
		void OnPluginStart(const plugify::IPlugin& plugin) override;
		void OnPluginEnd(const plugify::IPlugin& plugin) override;
		void OnMethodExport(const plugify::IPlugin& plugin) override;

	private:
		class AssemblyHolder {
		public:
			AssemblyHolder(std::unique_ptr<Assembly> assembly, StartFunc startFunc, EndFunc endFunc);

			StartFunc GetStartFunc() const;
			EndFunc GetEndFunc() const;

		private:
			std::unique_ptr<Assembly> _assembly;
			StartFunc _startFunc{ nullptr };
			EndFunc _endFunc{ nullptr };
		};

		std::shared_ptr<plugify::IPlugifyProvider> _provider;
		std::unordered_map<std::string, AssemblyHolder> _assemblyMap;
	};
}
