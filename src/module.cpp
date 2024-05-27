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

	template<typename T>
	void CopyGoSliceToVector(GoSlice* source, std::vector<T>& dest) {
		if constexpr (std::same_as<T, std::string>) {
			dest.resize(static_cast<size_t>(source->len));
			for (size_t k = 0; k < dest.size(); ++k) {
				auto str = reinterpret_cast<GoString**>(source->data)[k];
				dest[k].assign(str->p, static_cast<size_t>(str->n));
			}
		} else {
			if (source->data == nullptr || source->len == 0)
				dest.clear();
			else if (dest.data() != source->data)
				dest.assign(reinterpret_cast<T*>(source->data), reinterpret_cast<T*>(source->data) + static_cast<size_t>(source->len));
		}
	}

	void CopyGoStringToString(GoString* source, std::string& dest) {
		if (source->p == nullptr || source->n == 0)
			dest.clear();
		else if (dest.data() != source->p)
			dest.assign(source->p, static_cast<size_t>(source->n));
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
				funcAddr = function.GetJitFunc(method, &InternalCall, func);
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
				case ValueType::Function:
					dcArgPointer(vm, p->GetArgument<void*>(i));
					break;
				case ValueType::Float:
					dcArgPointer(vm, p->GetArgument<float*>(i));
					break;
				case ValueType::Double:
					dcArgPointer(vm, p->GetArgument<double*>(i));
					break;
				case ValueType::Vector2:
					dcArgPointer(vm, p->GetArgument<Vector2*>(i));
					break;
				case ValueType::Vector3:
					dcArgPointer(vm, p->GetArgument<Vector3*>(i));
					break;
				case ValueType::Vector4:
					dcArgPointer(vm, p->GetArgument<Vector4*>(i));
					break;
				case ValueType::Matrix4x4:
					dcArgPointer(vm, p->GetArgument<Matrix4x4*>(i));
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
				case ValueType::Function:
					dcArgPointer(vm, p->GetArgument<void*>(i));
					break;
				case ValueType::Float:
					dcArgFloat(vm, p->GetArgument<float>(i));
					break;
				case ValueType::Double:
					dcArgDouble(vm, p->GetArgument<double>(i));
					break;
				case ValueType::Vector2:
					dcArgPointer(vm, p->GetArgument<Vector2*>(i));
					break;
				case ValueType::Vector3:
					dcArgPointer(vm, p->GetArgument<Vector3*>(i));
					break;
				case ValueType::Vector4:
					dcArgPointer(vm, p->GetArgument<Vector4*>(i));
					break;
				case ValueType::Matrix4x4:
					dcArgPointer(vm, p->GetArgument<Matrix4x4*>(i));
					break;
				// GoString
				case ValueType::String: {
					GoString* string = utils::CreateGoString(*p->GetArgument<std::string*>(i), args);
					dcArgPointer(vm, const_cast<char*>(string->p));
					dcArgLongLong(vm, string->n);
					break;
				}
				// GoSlice
				/*case ValueType::ArrayBool: {
				 	GoSlice* array = utils::CreateGoSlice<bool>(*p->GetArgument<std::vector<bool>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}*/
				case ValueType::ArrayChar8: {
					GoSlice* array = utils::CreateGoSlice<char>(*p->GetArgument<std::vector<char>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayChar16: {
					GoSlice* array = utils::CreateGoSlice<char16_t>(*p->GetArgument<std::vector<char16_t>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayInt8: {
					GoSlice* array = utils::CreateGoSlice<int8_t>(*p->GetArgument<std::vector<int8_t>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayInt16: {
					GoSlice* array = utils::CreateGoSlice<int16_t>(*p->GetArgument<std::vector<int16_t>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayInt32: {
					GoSlice* array = utils::CreateGoSlice<int32_t>(*p->GetArgument<std::vector<int32_t>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayInt64: {
					GoSlice* array = utils::CreateGoSlice<int64_t>(*p->GetArgument<std::vector<int64_t>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayUInt8: {
					GoSlice* array = utils::CreateGoSlice<uint8_t>(*p->GetArgument<std::vector<uint8_t>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayUInt16: {
					GoSlice* array = utils::CreateGoSlice<uint16_t>(*p->GetArgument<std::vector<uint16_t>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayUInt32: {
					GoSlice* array = utils::CreateGoSlice<uint32_t>(*p->GetArgument<std::vector<uint32_t>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayUInt64: {
					GoSlice* array = utils::CreateGoSlice<uint64_t>(*p->GetArgument<std::vector<uint64_t>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayPointer: {
					GoSlice* array = utils::CreateGoSlice<uintptr_t>(*p->GetArgument<std::vector<uintptr_t>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayFloat: {
					GoSlice* array = utils::CreateGoSlice<float>(*p->GetArgument<std::vector<float>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayDouble: {
					GoSlice* array = utils::CreateGoSlice<double>(*p->GetArgument<std::vector<double>*>(i), args);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				case ValueType::ArrayString: {
					GoSlice* array = utils::CreateGoSliceString(*p->GetArgument<std::vector<std::string>*>(i), args, storage);
					dcArgPointer(vm, array->data);
					dcArgLongLong(vm, array->len);
					dcArgLongLong(vm, array->cap);
					break;
				}
				default:
					std::puts("Unsupported types!\n");
					std::terminate();
					break;
			}
		}
		hasRefs |= param.ref;
	}

	switch (method->retType.type) {
		case ValueType::Invalid:
			break;
			
		case ValueType::Void: {
			dcCallVoid(vm, addr);
			break;
		}
		case ValueType::Bool: {
			bool val = dcCallBool(vm, addr);
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::Char8: {
			char16_t val = static_cast<char16_t>(dcCallChar(vm, addr));
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::Char16: {
			char16_t val = static_cast<char16_t>(dcCallShort(vm, addr));
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::Int8: {
			int8_t val = dcCallChar(vm, addr);
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::Int16: {
			int16_t val = dcCallShort(vm, addr);
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::Int32: {
			int32_t val = dcCallInt(vm, addr);
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::Int64: {
			int64_t val = dcCallLongLong(vm, addr);
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::UInt8: {
			uint8_t val = static_cast<uint8_t>(dcCallChar(vm, addr));
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::UInt16: {
			uint16_t val = static_cast<uint16_t>(dcCallShort(vm, addr));
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::UInt32: {
			uint32_t val = static_cast<uint32_t>(dcCallInt(vm, addr));
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::UInt64: {
			uint64_t val = static_cast<uint64_t>(dcCallLongLong(vm, addr));
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::Pointer: {
			void* val = dcCallPointer(vm, addr);
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::Float: {
			float val = dcCallFloat(vm, addr);
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::Double: {
			double val = dcCallDouble(vm, addr);
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::Function: {
			void* val = dcCallPointer(vm, addr);
			ret->SetReturnPtr(val);
			break;
		}
		case ValueType::Vector2:
		case ValueType::Vector3:
		case ValueType::Vector4:
		case ValueType::Matrix4x4: {
			dcCallVoid(vm, addr);
			break;
		}
		case ValueType::String: {
			dcCallVoid(vm, addr);
			utils::CopyGoStringToString(reinterpret_cast<GoString*>(args[0]), *p->GetArgument<std::string*>(0));
			break;
		}
		case ValueType::ArrayBool: {
			dcCallVoid(vm, addr);
			//utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<bool>*>(0));
			break;
		}
		case ValueType::ArrayChar8: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<char>*>(0));
			break;
		}
		case ValueType::ArrayChar16: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<char16_t>*>(0));
			break;
		}
		case ValueType::ArrayInt8: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<int8_t>*>(0));
			break;
		}
		case ValueType::ArrayInt16: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<int16_t>*>(0));
			break;
		}
		case ValueType::ArrayInt32: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<int32_t>*>(0));
			break;
		}
		case ValueType::ArrayInt64: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<int64_t>*>(0));
			break;
		}
		case ValueType::ArrayUInt8: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<uint8_t>*>(0));
			break;
		}
		case ValueType::ArrayUInt16: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<uint16_t>*>(0));
			break;
		}
		case ValueType::ArrayUInt32: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<uint32_t>*>(0));
			break;
		}
		case ValueType::ArrayUInt64: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<uint64_t>*>(0));
			break;
		}
		case ValueType::ArrayPointer: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<uintptr_t>*>(0));
			break;
		}
		case ValueType::ArrayFloat: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<float>*>(0));
			break;
		}
		case ValueType::ArrayDouble: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<double>*>(0));
			break;
		}
		case ValueType::ArrayString: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[0]), *p->GetArgument<std::vector<std::string>*>(0));
			break;
		}
		default:
			std::puts("Unsupported types!\n");
			std::terminate();
			break;
	}

	if (hasRefs) {
		uint8_t j = hasRet; // skip first param if has return

		if (j < args.size()) {
			for (uint8_t i = 0; i < count; ++i) {
				const auto& param = method->paramTypes[i];
				if (param.ref) {
					switch (param.type) {
						case ValueType::String:
							utils::CopyGoStringToString(reinterpret_cast<GoString*>(args[j++]), *p->GetArgument<std::string*>(i));
							break;
                        /*case ValueType::ArrayBool:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<bool>*>(i));
							break;
                         }*/
						case ValueType::ArrayChar8:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<char>*>(i));
							break;
						case ValueType::ArrayChar16:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<char16_t>*>(i));
							break;
						case ValueType::ArrayInt8:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<int8_t>*>(i));
							break;
						case ValueType::ArrayInt16:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<int16_t>*>(i));
							break;
						case ValueType::ArrayInt32:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<int32_t>*>(i));
							break;
						case ValueType::ArrayInt64:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<int64_t>*>(i));
							break;
						case ValueType::ArrayUInt8:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<uint8_t>*>(i));
							break;
						case ValueType::ArrayUInt16:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<uint16_t>*>(i));
							break;
						case ValueType::ArrayUInt32:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<uint32_t>*>(i));
							break;
						case ValueType::ArrayUInt64:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<uint64_t>*>(i));
							break;
						case ValueType::ArrayPointer:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<uintptr_t>*>(i));
							break;
						case ValueType::ArrayFloat:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<float>*>(i));
							break;
						case ValueType::ArrayDouble:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<double>*>(i));
							break;
						case ValueType::ArrayString: {
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<std::string>*>(i));
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

void* GetMethodPtr(const char* methodName) {
	return g_golm.GetNativeMethod(methodName);
}

void* CreateStringE() {
	return malloc(sizeof(std::string));
}
void* CreateString(GoString source) {
	return new std::string(source.p, source.n);
}
const char* GetString(void* ptr) {
	return reinterpret_cast<std::string*>(ptr)->c_str();
}
void DeleteString(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::string*>(ptr)->~basic_string();
		free(ptr);
	} else
		delete reinterpret_cast<std::string*>(ptr);
}

void* CreateVectorBool(bool* arr, ptrdiff_t len) {
	return new std::vector<bool>(arr, arr + len);
}
void* CreateVectorChar8(char* arr, ptrdiff_t len) {
	return new std::vector<char>(arr, arr + len);
}
void* CreateVectorChar16(uint16_t* arr, ptrdiff_t len) {
	return new std::vector<uint16_t>(arr, arr + len);
}
void* CreateVectorInt8(int8_t* arr, ptrdiff_t len) {
	return new std::vector<int8_t>(arr, arr + len);
}
void* CreateVectorInt16(int16_t* arr, ptrdiff_t len) {
	return new std::vector<int16_t>(arr, arr + len);
}
void* CreateVectorInt32(int32_t* arr, ptrdiff_t len) {
	return new std::vector<int32_t>(arr, arr + len);
}
void* CreateVectorInt64(int64_t* arr, ptrdiff_t len) {
	return new std::vector<int64_t>(arr, arr + len);
}
void* CreateVectorUInt8(uint8_t* arr, ptrdiff_t len) {
	return new std::vector<uint8_t>(arr, arr + len);
}
void* CreateVectorUInt16(uint16_t* arr, ptrdiff_t len) {
	return new std::vector<uint16_t>(arr, arr + len);
}
void* CreateVectorUInt32(uint32_t* arr, ptrdiff_t len) {
	return new std::vector<uint32_t>(arr, arr + len);
}
void* CreateVectorUInt64(uint64_t* arr, ptrdiff_t len) {
	return new std::vector<uint64_t>(arr, arr + len);
}
void* CreateVectorUIntPtr(uintptr_t* arr, ptrdiff_t len) {
	return new std::vector<uintptr_t>(arr, arr + len);
}
void* CreateVectorFloat(float* arr, ptrdiff_t len) {
	return new std::vector<float>(arr, arr + len);
}
void* CreateVectorDouble(double* arr, ptrdiff_t len) {
	return new std::vector<double>(arr, arr + len);
}
void* CreateVectorString(GoString* arr, ptrdiff_t len) {
	auto* vector = new std::vector<std::string>();
	vector->reserve(len);
	for (ptrdiff_t i = 0; i < len; ++i) {
		const auto& element = arr[i];
		vector->emplace_back(element.p, element.n);
	}
	return vector;
}

void* CreateVectorBoolE() {
	return malloc(sizeof(std::vector<bool>));
}
void* CreateVectorChar8E() {
	return malloc(sizeof(std::vector<char>));
}
void* CreateVectorChar16E() {
	return malloc(sizeof(std::vector<uint16_t>));
}
void* CreateVectorInt8E() {
	return malloc(sizeof(std::vector<int8_t>));
}
void* CreateVectorInt16E() {
	return malloc(sizeof(std::vector<int16_t>));
}
void* CreateVectorInt32E() {
	return malloc(sizeof(std::vector<int32_t>));
}
void* CreateVectorInt64E() {
	return malloc(sizeof(std::vector<int64_t>));
}
void* CreateVectorUInt8E() {
	return malloc(sizeof(std::vector<uint8_t>));
}
void* CreateVectorUInt16E() {
	return malloc(sizeof(std::vector<uint16_t>));
}
void* CreateVectorUInt32E() {
	return malloc(sizeof(std::vector<uint32_t>));
}
void* CreateVectorUInt64E() {
	return malloc(sizeof(std::vector<uint64_t>));
}
void* CreateVectorUIntPtrE() {
	return malloc(sizeof(std::vector<uintptr_t>));
}
void* CreateVectorFloatE() {
	return malloc(sizeof(std::vector<float>));
}
void* CreateVectorDoubleE() {
	return malloc(sizeof(std::vector<double>));
}
void* CreateVectorStringE() {
	return malloc(sizeof(std::vector<std::string>));
}

void DeleteVectorBool(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<bool>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<bool>*>(ptr);
}
void DeleteVectorChar8(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<char>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<char>*>(ptr);
}
void DeleteVectorChar16(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<uint16_t>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<uint16_t>*>(ptr);
}
void DeleteVectorInt8(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<int8_t>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<int8_t>*>(ptr);
}
void DeleteVectorInt16(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<int16_t>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<int16_t>*>(ptr);
}
void DeleteVectorInt32(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<int32_t>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<int32_t>*>(ptr);
}
void DeleteVectorInt64(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<int64_t>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<int64_t>*>(ptr);
}
void DeleteVectorUInt8(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<uint8_t>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<uint8_t>*>(ptr);
}
void DeleteVectorUInt16(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<uint16_t>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<uint16_t>*>(ptr);
}
void DeleteVectorUInt32(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<uint32_t>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<uint32_t>*>(ptr);
}
void DeleteVectorUInt64(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<uint64_t>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<uint64_t>*>(ptr);
}
void DeleteVectorUIntPtr(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<uintptr_t>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<uintptr_t>*>(ptr);
}
void DeleteVectorFloat(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<float>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<float>*>(ptr);
}
void DeleteVectorDouble(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<double>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<double>*>(ptr);
}
void DeleteVectorString(void* ptr, bool output) {
	if (output) {
		reinterpret_cast<std::vector<std::string>*>(ptr)->~vector();
		free(ptr);
	} else
		delete reinterpret_cast<std::vector<std::string>*>(ptr);
}

ptrdiff_t GetVectorSizeBool(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<bool>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeChar8(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<char>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeChar16(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<char16_t>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeInt8(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int8_t>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeInt16(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int16_t>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeInt32(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int32_t>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeInt64(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int64_t>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeUInt8(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint8_t>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeUInt16(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint16_t>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeUInt32(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint32_t>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeUInt64(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint64_t>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeUIntPtr(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uintptr_t>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeFloat(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<float>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeDouble(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<double>*>(ptr)->size());
}
ptrdiff_t GetVectorSizeString(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<std::string>*>(ptr)->size());
}

void* GetVectorData(void* ptr) {
	return reinterpret_cast<std::vector<int>*>(ptr)->data();
}

void* GetVectorDataB(void* ptr) {
	auto& vec = *reinterpret_cast<std::vector<bool>*>(ptr);

	uint8_t* boolArray = new uint8_t[vec.size()];

	// Manually copy values from std::vector<bool> to the bool array
	for (size_t i = 0; i < vec.size(); ++i) {
		boolArray[i] = vec[i];
	}

	return boolArray;
}
void DeleteVectorDataB(void* ptr) {
	delete[] reinterpret_cast<bool*>(ptr);
}

void* GetVectorDataS(void* ptr) {
	auto& vec = *reinterpret_cast<std::vector<std::string>*>(ptr);

	char** strArray = new char*[vec.size()];

	// Manually copy values from std::vector<bool> to the bool array
	for (size_t i = 0; i < vec.size(); ++i) {
		auto& element = vec[i];
		strArray[i] = element.data();
	}

	return strArray;
}
void DeleteVectorDataS(void* ptr) {
	delete[] reinterpret_cast<char**>(ptr);
}

const std::array<void*, 70> GoLanguageModule::_pluginApi = {
		reinterpret_cast<void*>(&GetMethodPtr),
		reinterpret_cast<void*>(&CreateStringE),
		reinterpret_cast<void*>(&CreateString),
		reinterpret_cast<void*>(&GetString),
		reinterpret_cast<void*>(&DeleteString),
		reinterpret_cast<void*>(&CreateVectorBool),
		reinterpret_cast<void*>(&CreateVectorChar8),
		reinterpret_cast<void*>(&CreateVectorChar16),
		reinterpret_cast<void*>(&CreateVectorInt8),
		reinterpret_cast<void*>(&CreateVectorInt16),
		reinterpret_cast<void*>(&CreateVectorInt32),
		reinterpret_cast<void*>(&CreateVectorInt64),
		reinterpret_cast<void*>(&CreateVectorUInt8),
		reinterpret_cast<void*>(&CreateVectorUInt16),
		reinterpret_cast<void*>(&CreateVectorUInt32),
		reinterpret_cast<void*>(&CreateVectorUInt64),
		reinterpret_cast<void*>(&CreateVectorUIntPtr),
		reinterpret_cast<void*>(&CreateVectorFloat),
		reinterpret_cast<void*>(&CreateVectorDouble),
		reinterpret_cast<void*>(&CreateVectorString),
		reinterpret_cast<void*>(&CreateVectorBoolE),
		reinterpret_cast<void*>(&CreateVectorChar8E),
		reinterpret_cast<void*>(&CreateVectorChar16E),
		reinterpret_cast<void*>(&CreateVectorInt8E),
		reinterpret_cast<void*>(&CreateVectorInt16E),
		reinterpret_cast<void*>(&CreateVectorInt32E),
		reinterpret_cast<void*>(&CreateVectorInt64E),
		reinterpret_cast<void*>(&CreateVectorUInt8E),
		reinterpret_cast<void*>(&CreateVectorUInt16E),
		reinterpret_cast<void*>(&CreateVectorUInt32E),
		reinterpret_cast<void*>(&CreateVectorUInt64E),
		reinterpret_cast<void*>(&CreateVectorUIntPtrE),
		reinterpret_cast<void*>(&CreateVectorFloatE),
		reinterpret_cast<void*>(&CreateVectorDoubleE),
		reinterpret_cast<void*>(&CreateVectorStringE),
		reinterpret_cast<void*>(&DeleteVectorBool),
		reinterpret_cast<void*>(&DeleteVectorChar8),
		reinterpret_cast<void*>(&DeleteVectorChar16),
		reinterpret_cast<void*>(&DeleteVectorInt8),
		reinterpret_cast<void*>(&DeleteVectorInt16),
		reinterpret_cast<void*>(&DeleteVectorInt32),
		reinterpret_cast<void*>(&DeleteVectorInt64),
		reinterpret_cast<void*>(&DeleteVectorUInt8),
		reinterpret_cast<void*>(&DeleteVectorUInt16),
		reinterpret_cast<void*>(&DeleteVectorUInt32),
		reinterpret_cast<void*>(&DeleteVectorUInt64),
		reinterpret_cast<void*>(&DeleteVectorUIntPtr),
		reinterpret_cast<void*>(&DeleteVectorFloat),
		reinterpret_cast<void*>(&DeleteVectorDouble),
		reinterpret_cast<void*>(&DeleteVectorString),
		reinterpret_cast<void*>(&GetVectorSizeBool),
		reinterpret_cast<void*>(&GetVectorSizeChar8),
		reinterpret_cast<void*>(&GetVectorSizeChar16),
		reinterpret_cast<void*>(&GetVectorSizeInt8),
		reinterpret_cast<void*>(&GetVectorSizeInt16),
		reinterpret_cast<void*>(&GetVectorSizeInt32),
		reinterpret_cast<void*>(&GetVectorSizeInt64),
		reinterpret_cast<void*>(&GetVectorSizeUInt8),
		reinterpret_cast<void*>(&GetVectorSizeUInt16),
		reinterpret_cast<void*>(&GetVectorSizeUInt32),
		reinterpret_cast<void*>(&GetVectorSizeUInt64),
		reinterpret_cast<void*>(&GetVectorSizeUIntPtr),
		reinterpret_cast<void*>(&GetVectorSizeFloat),
		reinterpret_cast<void*>(&GetVectorSizeDouble),
		reinterpret_cast<void*>(&GetVectorSizeString),
		reinterpret_cast<void*>(&GetVectorData),
		reinterpret_cast<void*>(&GetVectorDataB),
		reinterpret_cast<void*>(&DeleteVectorDataB),
		reinterpret_cast<void*>(&GetVectorDataS),
		reinterpret_cast<void*>(&DeleteVectorDataS)
};

namespace golm {
	GoLanguageModule g_golm;
}

plugify::ILanguageModule* GetLanguageModule() {
	return &golm::g_golm;
}