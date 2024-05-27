#include "module.h"

#include <plugify/module.h>
#include <plugify/plugin.h>
#include <plugify/log.h>
#include <plugify/math.h>
#include <plugify/plugin_descriptor.h>
#include <plugify/plugify_provider.h>

#define LOG_PREFIX "[GOLM] "

using namespace golm;
using namespace plugify;

namespace golm::utils {
	bool IsMethodPrimitive(const plugify::Method& method) {
		if (method.retType.type >= ValueType::LastPrimitive)
			return false;

		for (const auto& param : method.paramTypes) {
			if (param.type >= ValueType::LastPrimitive)
				return false;
		}

		return true;
	}

	template<typename T>
	void* GoSliceToVector(GoSlice* source, ArgumentList& args) {
		std::vector<T>* dest;
		if (source != nullptr) {
			dest = new std::vector<T>(reinterpret_cast<T*>(source->data), reinterpret_cast<T*>(source->data) + static_cast<size_t>(source->len));
		} else {
			dest = new std::vector<T>();
		}
		args.push_back(dest);
		return dest;
	}

	void* GoStringToString(GoString* source, ArgumentList& args) {
		std::string* dest;
		if (source != nullptr) {
			dest = new std::string(source->p,  static_cast<size_t>(source->n));
		} else {
			dest = new std::string();
		}
		args.push_back(dest);
		return dest;
	}

	GoSlice* CreateGoSliceString(const std::vector<std::string>& source, ArgumentList& args, StringStorage& storage) {
		auto& [arrays, strings] = storage;
		auto& strArray = arrays.emplace_back(std::make_unique<GoString*[]>(source.size()));
		for (size_t i = 0; i < source.size(); ++i) {
			const auto& str = source[i];
			auto& it = strings.emplace_back(std::make_unique<GoString>(str.c_str(), static_cast<GoInt>(str.size())));
			strArray[i] = it.get();
		}
		auto size = static_cast<GoInt>(source.size());
		auto* dest = new GoSlice(strArray.get(), size, size);
		args.push_back(dest);
		return dest;
	}

	template<typename T, GoInt N>
	void* CreateGoSliceVector(const T& source, ArgumentList& args) {
		auto* dest = new GoSlice((void*)source.data.data(), N, N);
		args.push_back(dest);
		return dest;
	}

	template<typename T>
	GoSlice* CreateGoSlice(const std::vector<T>& source, ArgumentList& args) {
		auto size = static_cast<GoInt>(source.size());
		auto* dest = new GoSlice((void*)source.data(), size, size);
		args.push_back(dest);
		return dest;
	}

	template<typename T>
	void DeleteGoSlice(void* ptr) {
		delete reinterpret_cast<GoSlice*>(ptr);
	}

	GoString* CreateGoString(const std::string& source, ArgumentList& args) {
		auto size = static_cast<GoInt>(source.size());
		auto* dest = new GoString(source.c_str(), size);
		args.push_back(dest);
		return dest;
	}

	void DeleteGoString(void* ptr) {
		delete reinterpret_cast<GoString*>(ptr);
	}

	template<typename T>
	void* AllocateMemory(ArgumentList& args) {
		void* ptr = malloc(sizeof(T));
		args.push_back(ptr);
		return ptr;
	}

	template<typename T>
	void FreeMemory(void* ptr) {
		reinterpret_cast<T*>(ptr)->~T();
		free(ptr);
	}
}

InitResult GoLanguageModule::Initialize(std::weak_ptr<IPlugifyProvider> provider, const IModule& /*module*/) {
	if (!(_provider = provider.lock())) {
		return ErrorData{ "Provider not exposed" };
	}

	_provider->Log("[GOLM] Inited!", Severity::Debug);

	_rt = std::make_shared<asmjit::JitRuntime>();

	DCCallVM* vm = dcNewCallVM(4096);
	dcMode(vm, DC_CALL_C_DEFAULT);
	_callVirtMachine = std::unique_ptr<DCCallVM, VMDeleter>(vm);

	return InitResultData{};
}

void GoLanguageModule::Shutdown() {
	_callVirtMachine.reset();
	_functions.clear();
	_nativesMap.clear();
	_assemblies.clear();
	_provider.reset();
	_rt.reset();
}

void GoLanguageModule::OnMethodExport(const IPlugin& plugin) {
	const auto& pluginName = plugin.GetName();
	for (const auto& [name, addr] : plugin.GetMethods()) {
		_nativesMap.try_emplace(std::format("{}.{}", pluginName, name), addr);
	}
}

LoadResult GoLanguageModule::OnPluginLoad(const IPlugin& plugin) {
	fs::path assemblyPath = plugin.GetBaseDir() / std::format("{}" BINARY_MODULE_SUFFIX, plugin.GetDescriptor().entryPoint);

	auto assembly = Assembly::LoadFromPath(assemblyPath);
	if (!assembly) {
		return ErrorData{ std::format("Failed to load assembly: {}", Assembly::GetError()) };
	}

	std::vector<std::string_view> funcErrors;

	auto* const initFunc = assembly->GetFunction<InitFunc>("Plugify_Init");
	if (!initFunc) {
		funcErrors.emplace_back("Plugify_Init");
	}

	auto* const startFunc = assembly->GetFunction<StartFunc>("Plugify_PluginStart");
	if (!startFunc) {
		funcErrors.emplace_back("Plugify_PluginStart");
	}

	auto* const endFunc = assembly->GetFunction<EndFunc>("Plugify_PluginEnd");
	if (!endFunc) {
		funcErrors.emplace_back("Plugify_PluginEnd");
	}

	if (!funcErrors.empty()) {
		std::string funcs(funcErrors[0]);
		for (auto it = std::next(funcErrors.begin()); it != funcErrors.end(); ++it) {
			std::format_to(std::back_inserter(funcs), ", {}", *it);
		}
		return ErrorData{ std::format("Not found {} function(s)", funcs) };
	}

	const auto& exportedMethods = plugin.GetDescriptor().exportedMethods;
	std::vector<MethodData> methods;
	methods.reserve(exportedMethods.size());

	for (const auto& method : exportedMethods) {
		if (auto* const func = assembly->GetFunction(method.funcName.c_str())) {
			void* funcAddr;
			if (utils::IsMethodPrimitive(method)) {
				funcAddr = func;
			} else {
				Function function(_rt);
				funcAddr = function.GetJitFunc(method, &InternalCall, func); //// TODO: Implement
				if (!funcAddr) {
					funcErrors.emplace_back(method.name);
					continue;
				}
				_functions.emplace(funcAddr, std::move(function));
			}

			methods.emplace_back(method.name, funcAddr);
		}
		else {
			funcErrors.emplace_back(method.name);
		}
	}
	if (!funcErrors.empty()) {
		std::string funcs(funcErrors[0]);
		for (auto it = std::next(funcErrors.begin()); it != funcErrors.end(); ++it) {
			std::format_to(std::back_inserter(funcs), ", {}", *it);
		}
		return ErrorData{ std::format("Not found {} method function(s)", funcs) };
	}

	GoSlice api { const_cast<void**>(_pluginApi.data()), _pluginApi.size(), _pluginApi.size() };
	const int resultVersion = initFunc(api, kApiVersion);
	if (resultVersion != 0) {
		return ErrorData{ std::format("Not supported plugin api {}, max supported {}", resultVersion, kApiVersion) };
	}

	const auto [_, result] = _assemblies.try_emplace(plugin.GetName(), std::move(assembly), startFunc, endFunc);
	if (!result) {
		return ErrorData{ std::format("Plugin name duplicate") };
	}

	return LoadResultData{ std::move(methods) };
}

void GoLanguageModule::OnPluginStart(const IPlugin& plugin) {
	if (const auto it = _assemblies.find(plugin.GetName()); it != _assemblies.end()) {
		const auto& assemblyHolder = std::get<AssemblyHolder>(*it);
		assemblyHolder.GetStartFunc()();
	}
}

void GoLanguageModule::OnPluginEnd(const IPlugin& plugin) {
	if (const auto it = _assemblies.find(plugin.GetName()); it != _assemblies.end()) {
		const auto& assemblyHolder = std::get<AssemblyHolder>(*it);
		assemblyHolder.GetEndFunc()();
	}
}

void* GoLanguageModule::GetNativeMethod(const std::string& methodName) const {
	if (const auto it = _nativesMap.find(methodName); it != _nativesMap.end()) {
		return std::get<void*>(*it);
	}
	_provider->Log(std::format("[GOLM] GetNativeMethod failed to find: '{}'", methodName), Severity::Fatal);
	return nullptr;
}

// C++ to Go
void GoLanguageModule::InternalCall(const plugify::Method* method, void* addr, const plugify::Parameters* p, uint8_t count, const plugify::ReturnValue* ret) {
	std::scoped_lock<std::mutex> lock(g_golm._mutex);
	ArgumentList args;
	StringStorage storage;

	DCCallVM* vm = g_golm._callVirtMachine.get();
	dcReset(vm);

	bool hasRet = method->retType.type > ValueType::LastPrimitive;
	bool hasRefs = false;

	// Store parameters

	if (hasRet) {
		switch (method->retType.type) {
			// GoString*
			case ValueType::String:
				dcArgPointer(vm, utils::AllocateMemory<GoString>(args));
				break;
			// GoSlice*
			case ValueType::Vector2:
			case ValueType::Vector3:
			case ValueType::Vector4:
			case ValueType::Matrix4x4:
			//case ValueType::ArrayBool:
			case ValueType::ArrayChar8:
			case ValueType::ArrayChar16:
			case ValueType::ArrayInt8:
			case ValueType::ArrayInt16:
			case ValueType::ArrayInt32:
			case ValueType::ArrayInt64:
			case ValueType::ArrayUInt8:
			case ValueType::ArrayUInt16:
			case ValueType::ArrayUInt32:
			case ValueType::ArrayUInt64:
			case ValueType::ArrayPointer:
			case ValueType::ArrayFloat:
			case ValueType::ArrayDouble:
			case ValueType::ArrayString:
				dcArgPointer(vm, utils::AllocateMemory<GoSlice>(args));
				break;
			default:
				// Should not require storage
				break;
		}
	}

	for (uint8_t i = 0; i < count; ++i) {
		const auto& param = method->paramTypes[i];
		if (param.ref) {
			switch (param.type) {
				case ValueType::Invalid:
				case ValueType::Void:
					break;
				case ValueType::Bool:
					dcArgPointer(vm, p->GetArgument<bool*>(i));
					break;
				case ValueType::Char8:
					dcArgPointer(vm, p->GetArgument<char*>(i));
					break;
				case ValueType::Char16:
					dcArgPointer(vm, p->GetArgument<char16_t*>(i));
					break;
				case ValueType::Int8:
					dcArgPointer(vm, p->GetArgument<int8_t*>(i));
					break;
				case ValueType::Int16:
					dcArgPointer(vm, p->GetArgument<int16_t*>(i));
					break;
				case ValueType::Int32:
					dcArgPointer(vm, p->GetArgument<int32_t*>(i));
					break;
				case ValueType::Int64:
					dcArgPointer(vm, p->GetArgument<int64_t*>(i));
					break;
				case ValueType::UInt8:
					dcArgPointer(vm, p->GetArgument<uint8_t*>(i));
					break;
				case ValueType::UInt16:
					dcArgPointer(vm, p->GetArgument<uint16_t*>(i));
					break;
				case ValueType::UInt32:
					dcArgPointer(vm, p->GetArgument<uint32_t*>(i));
					break;
				case ValueType::UInt64:
					dcArgPointer(vm, p->GetArgument<uint64_t*>(i));
					break;
				case ValueType::Pointer:
					dcArgPointer(vm, p->GetArgument<void*>(i));
					break;
				case ValueType::Float:
					dcArgPointer(vm, p->GetArgument<float*>(i));
					break;
				case ValueType::Double:
					dcArgPointer(vm, p->GetArgument<double*>(i));
					break;
				case ValueType::Vector2:
					dcArgPointer(vm, utils::CreateGoSliceVector<Vector2, 2>(*p->GetArgument<Vector2*>(i), args));
					break;
				case ValueType::Vector3:
					dcArgPointer(vm, utils::CreateGoSliceVector<Vector3, 3>(*p->GetArgument<Vector3*>(i), args));
					break;
				case ValueType::Vector4:
					dcArgPointer(vm, utils::CreateGoSliceVector<Vector4, 4>(*p->GetArgument<Vector4*>(i), args));
					break;
				case ValueType::Matrix4x4:
					dcArgPointer(vm, utils::CreateGoSliceVector<Matrix4x4, 16>(*p->GetArgument<Matrix4x4*>(i), args));
					break;
				// ??
				case ValueType::Function:
					//dcArgPointer(vm, g_csharplm.MonoDelegateToArg(*p->GetArgument<MonoDelegate**>(i), *param.prototype));
					break;
				// GoString*
				case ValueType::String:
					dcArgPointer(vm, utils::CreateGoString(*p->GetArgument<std::string*>(i), args));
					break;
				// GoSlice*
				/*case ValueType::ArrayBool:
					dcArgPointer(vm, utils::CreateGoSlice<bool>(*p->GetArgument<std::vector<bool>*>(i), args));
					break;*/
				case ValueType::ArrayChar8:
					dcArgPointer(vm, { utils::CreateGoSlice<char>(*p->GetArgument<std::vector<char>*>(i), args) });
					break;
				case ValueType::ArrayChar16:
					dcArgPointer(vm, { utils::CreateGoSlice<char16_t>(*p->GetArgument<std::vector<char16_t>*>(i), args) });
					break;
				case ValueType::ArrayInt8:
					dcArgPointer(vm, { utils::CreateGoSlice<int8_t>(*p->GetArgument<std::vector<int8_t>*>(i), args) });
					break;
				case ValueType::ArrayInt16:
					dcArgPointer(vm, { utils::CreateGoSlice<int16_t>(*p->GetArgument<std::vector<int16_t>*>(i), args) });
					break;
				case ValueType::ArrayInt32:
					dcArgPointer(vm, { utils::CreateGoSlice<int32_t>(*p->GetArgument<std::vector<int32_t>*>(i), args) });
					break;
				case ValueType::ArrayInt64:
					dcArgPointer(vm, { utils::CreateGoSlice<int64_t>(*p->GetArgument<std::vector<int64_t>*>(i), args) });
					break;
				case ValueType::ArrayUInt8:
					dcArgPointer(vm, { utils::CreateGoSlice<uint8_t>(*p->GetArgument<std::vector<uint8_t>*>(i), args) });
					break;
				case ValueType::ArrayUInt16:
					dcArgPointer(vm, { utils::CreateGoSlice<uint16_t>(*p->GetArgument<std::vector<uint16_t>*>(i), args) });
					break;
				case ValueType::ArrayUInt32:
					dcArgPointer(vm, { utils::CreateGoSlice<uint32_t>(*p->GetArgument<std::vector<uint32_t>*>(i), args) });
					break;
				case ValueType::ArrayUInt64:
					dcArgPointer(vm, { utils::CreateGoSlice<uint64_t>(*p->GetArgument<std::vector<uint64_t>*>(i), args) });
					break;
				case ValueType::ArrayPointer:
					dcArgPointer(vm, { utils::CreateGoSlice<uintptr_t>(*p->GetArgument<std::vector<uintptr_t>*>(i), args) });
					break;
				case ValueType::ArrayFloat:
					dcArgPointer(vm, { utils::CreateGoSlice<float>(*p->GetArgument<std::vector<float>*>(i), args) });
					break;
				case ValueType::ArrayDouble:
					dcArgPointer(vm, { utils::CreateGoSlice<double>(*p->GetArgument<std::vector<double>*>(i), args) });
					break;
				case ValueType::ArrayString:
					dcArgPointer(vm, { utils::CreateGoSliceString(*p->GetArgument<std::vector<std::string>*>(i), args, storage) });
					break;
				default:
					std::puts("Unsupported types!\n");
					std::terminate();
					break;
			}
		} else {
			switch (param.type) {
				case ValueType::Invalid:
				case ValueType::Void:
					break;
				case ValueType::Bool:
					dcArgBool(vm, p->GetArgument<bool>(i));
					break;
				case ValueType::Char8:
					dcArgChar(vm, p->GetArgument<char>(i));
					break;
				case ValueType::Char16:
					dcArgShort(vm, static_cast<short>(p->GetArgument<char16_t>(i)));
					break;
				case ValueType::Int8:
				case ValueType::UInt8:
					dcArgChar(vm, p->GetArgument<int8_t>(i));
					break;
				case ValueType::Int16:
				case ValueType::UInt16:
					dcArgShort(vm, p->GetArgument<int16_t>(i));
					break;
				case ValueType::Int32:
				case ValueType::UInt32:
					dcArgInt(vm, p->GetArgument<int32_t>(i));
					break;
				case ValueType::Int64:
				case ValueType::UInt64:
					dcArgLongLong(vm, p->GetArgument<int64_t>(i));
					break;
				case ValueType::Pointer:
					dcArgPointer(vm, p->GetArgument<void*>(i));
					break;
				case ValueType::Float:
					dcArgFloat(vm, p->GetArgument<float>(i));
					break;
				case ValueType::Double:
					dcArgDouble(vm, p->GetArgument<double>(i));
					break;
				case ValueType::Vector2:
					dcArgPointer(vm, utils::CreateGoSliceVector<Vector2, 2>(*p->GetArgument<Vector2*>(i), args));
					break;
				case ValueType::Vector3:
					dcArgPointer(vm, utils::CreateGoSliceVector<Vector3, 3>(*p->GetArgument<Vector3*>(i), args));
					break;
				case ValueType::Vector4:
					dcArgPointer(vm, utils::CreateGoSliceVector<Vector4, 4>(*p->GetArgument<Vector4*>(i), args));
					break;
				case ValueType::Matrix4x4:
					dcArgPointer(vm, utils::CreateGoSliceVector<Matrix4x4, 16>(*p->GetArgument<Matrix4x4*>(i), args));
					break;
					// ??
				case ValueType::Function:
					//dcArgPointer(vm, g_csharplm.MonoDelegateToArg(*p->GetArgument<MonoDelegate**>(i), *param.prototype));
					break;
					// GoString*
				case ValueType::String:
					dcArgPointer(vm, utils::CreateGoString(*p->GetArgument<std::string*>(i), args));
					break;
					// GoSlice*
					/*case ValueType::ArrayBool:
						dcArgPointer(vm, utils::CreateGoSlice<bool>(*p->GetArgument<std::vector<bool>*>(i), args));
						break;*/
				case ValueType::ArrayChar8:
					dcArgPointer(vm, utils::CreateGoSlice<char>(*p->GetArgument<std::vector<char>*>(i), args));
					break;
				case ValueType::ArrayChar16:
					dcArgPointer(vm, utils::CreateGoSlice<char16_t>(*p->GetArgument<std::vector<char16_t>*>(i), args));
					break;
				case ValueType::ArrayInt8:
					dcArgPointer(vm, utils::CreateGoSlice<int8_t>(*p->GetArgument<std::vector<int8_t>*>(i), args));
					break;
				case ValueType::ArrayInt16:
					dcArgPointer(vm, utils::CreateGoSlice<int16_t>(*p->GetArgument<std::vector<int16_t>*>(i), args));
					break;
				case ValueType::ArrayInt32:
					dcArgPointer(vm, utils::CreateGoSlice<int32_t>(*p->GetArgument<std::vector<int32_t>*>(i), args));
					break;
				case ValueType::ArrayInt64:
					dcArgPointer(vm, utils::CreateGoSlice<int64_t>(*p->GetArgument<std::vector<int64_t>*>(i), args));
					break;
				case ValueType::ArrayUInt8:
					dcArgPointer(vm, utils::CreateGoSlice<uint8_t>(*p->GetArgument<std::vector<uint8_t>*>(i), args));
					break;
				case ValueType::ArrayUInt16:
					dcArgPointer(vm, utils::CreateGoSlice<uint16_t>(*p->GetArgument<std::vector<uint16_t>*>(i), args));
					break;
				case ValueType::ArrayUInt32:
					dcArgPointer(vm, utils::CreateGoSlice<uint32_t>(*p->GetArgument<std::vector<uint32_t>*>(i), args));
					break;
				case ValueType::ArrayUInt64:
					dcArgPointer(vm, utils::CreateGoSlice<uint64_t>(*p->GetArgument<std::vector<uint64_t>*>(i), args));
					break;
				case ValueType::ArrayPointer:
					dcArgPointer(vm, utils::CreateGoSlice<uintptr_t>(*p->GetArgument<std::vector<uintptr_t>*>(i), args));
					break;
				case ValueType::ArrayFloat:
					dcArgPointer(vm, utils::CreateGoSlice<float>(*p->GetArgument<std::vector<float>*>(i), args));
					break;
				case ValueType::ArrayDouble:
					dcArgPointer(vm, utils::CreateGoSlice<double>(*p->GetArgument<std::vector<double>*>(i), args));
					break;
				case ValueType::ArrayString:
					dcArgPointer(vm, utils::CreateGoSliceString(*p->GetArgument<std::vector<std::string>*>(i), args, storage));
					break;
				default:
					std::puts("Unsupported types!\n");
					std::terminate();
					break;
			}
		}
		hasRefs |= param.ref;
	}

	// TODO: Add dmCall here
	(void*)addr;
	(void*)ret;

	if (hasRefs) {
		uint8_t j = hasRet; // skip first param if has return

		if (j < args.size()) {
			for (uint8_t i = 0; i < count; ++i) {
				const auto& param = method->paramTypes[i];
				if (param.ref) {
					switch (param.type) {
						case ValueType::String: {
							auto source = reinterpret_cast<GoString*>(args[j++]);
							auto dest = p->GetArgument<std::string*>(i);
							if (source->p == nullptr || source->n == 0)
								dest->clear();
							else if (dest->data() != source->p)
								dest->assign(source->p, static_cast<size_t>(source->n));
							break;
						}
                        /*case ValueType::ArrayBool: {
                            auto source = reinterpret_cast<GoSlice*>(args[j++]);
                            p->GetArgument<std::vector<bool>*>(i);
                            if (source->data == nullptr || source->len == 0)
                                dest->clear();
                            else if (dest->data() != source->data)
                                dest->assign(reinterpret_cast<bool*>(source->data), reinterpret_cast<bool*>(source->data) + static_cast<size_t>(source->len));
                            break;
                        }*/
						case ValueType::ArrayChar8: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<char>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<char*>(source->data), reinterpret_cast<char*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayChar16: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<char16_t>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<char16_t*>(source->data), reinterpret_cast<char16_t*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayInt8: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<int8_t>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<int8_t*>(source->data), reinterpret_cast<int8_t*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayInt16: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<int16_t>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<int16_t*>(source->data), reinterpret_cast<int16_t*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayInt32: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<int32_t>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<int32_t*>(source->data), reinterpret_cast<int32_t*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayInt64: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<int64_t>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<int64_t*>(source->data), reinterpret_cast<int64_t*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayUInt8: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<uint8_t>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<uint8_t*>(source->data), reinterpret_cast<uint8_t*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayUInt16: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<uint16_t>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<uint16_t*>(source->data), reinterpret_cast<uint16_t*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayUInt32: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<uint32_t>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<uint32_t*>(source->data), reinterpret_cast<uint32_t*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayUInt64: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<uint64_t>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<uint64_t*>(source->data), reinterpret_cast<uint64_t*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayPointer: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<uintptr_t>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<uintptr_t*>(source->data), reinterpret_cast<uintptr_t*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayFloat: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<float>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<float*>(source->data), reinterpret_cast<float*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayDouble: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto dest = p->GetArgument<std::vector<double>*>(i);
							if (source->data == nullptr || source->len == 0)
								dest->clear();
							else if (dest->data() != source->data)
								dest->assign(reinterpret_cast<double*>(source->data), reinterpret_cast<double*>(source->data) + static_cast<size_t>(source->len));
							break;
						}
						case ValueType::ArrayString: {
							auto source = reinterpret_cast<GoSlice*>(args[j++]);
							auto& dest=  *p->GetArgument<std::vector<std::string>*>(i);
							dest.resize(static_cast<size_t>(source->len));
							for (size_t k = 0; k < dest.size(); ++k) {
								auto str = reinterpret_cast<GoString**>(source->data)[k];
								dest[k].assign(str->p, static_cast<size_t>(str->n));
							}
							break;
						}
						default:
							break;
					}
				}
				if (j == args.size())
					break;
			}
		}
	}

	if (!args.empty()) {
		uint8_t j = 0;

		if (hasRet) {
			switch (method->retType.type) {
				case ValueType::String:
					utils::FreeMemory<GoString>(args[j++]);
					break;
				case ValueType::Vector2:
				case ValueType::Vector3:
				case ValueType::Vector4:
				case ValueType::Matrix4x4:
				//case ValueType::ArrayBool:
				case ValueType::ArrayChar8:
				case ValueType::ArrayChar16:
				case ValueType::ArrayInt8:
				case ValueType::ArrayInt16:
				case ValueType::ArrayInt32:
				case ValueType::ArrayInt64:
				case ValueType::ArrayUInt8:
				case ValueType::ArrayUInt16:
				case ValueType::ArrayUInt32:
				case ValueType::ArrayUInt64:
				case ValueType::ArrayPointer:
				case ValueType::ArrayFloat:
				case ValueType::ArrayDouble:
				case ValueType::ArrayString:
					utils::FreeMemory<GoSlice>(args[j++]);
					break;
				default:
					break;
			}
		}

		if (j < args.size()) {
			for (uint8_t i = 0; i < count; ++i) {
				switch (method->paramTypes[i].type) {
					case ValueType::String:
						utils::DeleteGoString(args[j++]);
						break;
                    /*case ValueType::ArrayBool:
                        utils::DeleteGoSlice<bool>(args[j++]);
                        break;*/
					case ValueType::ArrayChar8:
						utils::DeleteGoSlice<char>(args[j++]);
						break;
					case ValueType::ArrayChar16:
						utils::DeleteGoSlice<char16_t>(args[j++]);
						break;
					case ValueType::ArrayInt8:
						utils::DeleteGoSlice<int16_t>(args[j++]);
						break;
					case ValueType::ArrayInt16:
						utils::DeleteGoSlice<int16_t>(args[j++]);
						break;
					case ValueType::ArrayInt32:
						utils::DeleteGoSlice<int32_t>(args[j++]);
						break;
					case ValueType::ArrayInt64:
						utils::DeleteGoSlice<int64_t>(args[j++]);
						break;
					case ValueType::ArrayUInt8:
						utils::DeleteGoSlice<uint8_t>(args[j++]);
						break;
					case ValueType::ArrayUInt16:
						utils::DeleteGoSlice<uint16_t>(args[j++]);
						break;
					case ValueType::ArrayUInt32:
						utils::DeleteGoSlice<uint32_t>(args[j++]);
						break;
					case ValueType::ArrayUInt64:
						utils::DeleteGoSlice<uint64_t>(args[j++]);
						break;
					case ValueType::ArrayPointer:
						utils::DeleteGoSlice<uintptr_t>(args[j++]);
						break;
					case ValueType::ArrayFloat:
						utils::DeleteGoSlice<float>(args[j++]);
						break;
					case ValueType::ArrayDouble:
						utils::DeleteGoSlice<double>(args[j++]);
						break;
					case ValueType::ArrayString:
						utils::DeleteGoSlice<GoString>(args[j++]);
						break;
					default:
						break;
				}
				if (j == args.size())
					break;
			}
		}
	}
}

void* GetNativeMethodImpl(GoString methodName) {
	return g_golm.GetNativeMethod(std::string(methodName.p, methodName.n));
}

const std::array<void*, 1> GoLanguageModule::_pluginApi = {
		reinterpret_cast<void*>(&GetNativeMethodImpl),
};

namespace golm {
	GoLanguageModule g_golm;
}

plugify::ILanguageModule* GetLanguageModule() {
	return &golm::g_golm;
}