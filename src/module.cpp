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

	GoSlice* CreateGoSliceBool(const std::vector<bool>& source, ArgumentList& args, BoolStorage& storage) {
		size_t N = source.size();
		auto& boolArray = storage.emplace_back(std::make_unique<bool[]>(N));
		for (size_t i = 0; i < N; ++i) {
			boolArray[i] = source[i];
		}
		auto size = static_cast<GoInt>(N);
		auto* dest = new GoSlice(boolArray.get(), size, size);
		args.push_back(dest);
		return dest;
	}

	GoSlice* CreateGoSliceString(const std::vector<std::string>& source, ArgumentList& args, StringStorage& storage) {
		size_t N = source.size();
		auto& strArray = storage.emplace_back(std::make_unique<GoString[]>(N));
		for (size_t i = 0; i < N; ++i) {
			const auto& str = source[i];
			auto& dest = strArray[i];
			dest.p = str.c_str();
			dest.n = static_cast<GoInt>(str.size());
		}
		auto size = static_cast<GoInt>(N);
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

	GoSlice* CreateGoSlice(ArgumentList& args) {
		auto* dest = new GoSlice(nullptr, 0, 0);
		args.push_back(dest);
		return dest;
	}

	void DeleteGoSlice(void* ptr) {
		delete reinterpret_cast<GoSlice*>(ptr);
	}

	GoString* CreateGoString(ArgumentList& args) {
		auto* dest = new GoString("", 0);
		args.push_back(dest);
		return dest;
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
	void CopyGoSliceToVector(GoSlice* source, std::vector<T>& dest) {
		if constexpr (std::same_as<T, bool>) {
			dest.resize(static_cast<size_t>(source->len));
			for (size_t i = 0; i < dest.size(); ++i) {
				dest[i] = reinterpret_cast<bool*>(source->data)[i];
			}
		} else if constexpr (std::same_as<T, std::string>) {
			dest.resize(static_cast<size_t>(source->len));
			for (size_t i = 0; i < dest.size(); ++i) {
				const auto& str = reinterpret_cast<GoString*>(source->data)[i];
				dest[i].assign(str.p, static_cast<size_t>(str.n));
			}
		} else {
			if (source->data == nullptr || source->len == 0)
				dest.clear();
			else if (dest.data() != source->data)
				dest.assign(reinterpret_cast<T*>(source->data), reinterpret_cast<T*>(source->data) + static_cast<size_t>(source->len));
		}
	}

	template<typename T>
	void CopyGoSliceToVectorReturn(GoSlice* source, std::vector<T>* dest) {
		if constexpr (std::same_as<T, bool>) {
			if (source->data == nullptr || source->len == 0)
				std::construct_at(dest, std::vector<T>());
			else {
				std::construct_at(dest, std::vector<T>(static_cast<size_t>(source->len)));
				for (size_t i = 0; i < dest->size(); ++i) {
					(*dest)[i] = reinterpret_cast<bool*>(source->data)[i];
				}
			}
		} else if constexpr (std::same_as<T, std::string>) {
			if (source->data == nullptr || source->len == 0)
				std::construct_at(dest, std::vector<T>());
			else {
				std::construct_at(dest, std::vector<T>(static_cast<size_t>(source->len)));
				for (size_t i = 0; i < dest->size(); ++i) {
					const auto& str = reinterpret_cast<GoString*>(source->data)[i];
					(*dest)[i].assign(str.p, static_cast<size_t>(str.n));
				}
			}
		} else {
			if (source->data == nullptr || source->len == 0)
				std::construct_at(dest, std::vector<T>());
			else
				std::construct_at(dest, std::vector<T>(reinterpret_cast<T*>(source->data), reinterpret_cast<T*>(source->data) + static_cast<size_t>(source->len)));
		}
	}

	void CopyGoStringToString(GoString* source, std::string& dest) {
		if (source->p == nullptr || source->n == 0)
			dest.clear();
		else if (dest.data() != source->p)
			dest.assign(source->p, static_cast<size_t>(source->n));
	}

	void CopyGoStringToStringReturn(GoString* source, std::string* dest) {
		if (source->p == nullptr || source->n == 0)
			std::construct_at(dest, std::string());
		else
			std::construct_at(dest, std::string(source->p, static_cast<size_t>(source->n)));
	}
}

InitResult GoLanguageModule::Initialize(std::weak_ptr<IPlugifyProvider> provider, const IModule& /*module*/) {
	if (!(_provider = provider.lock())) {
		return ErrorData{ "Provider not exposed" };
	}

	_provider->Log(LOG_PREFIX "Inited!", Severity::Debug);

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
	_provider->Log(std::format(LOG_PREFIX "GetNativeMethod failed to find: '{}'", methodName), Severity::Fatal);
	return nullptr;
}

// C++ to Go
void GoLanguageModule::InternalCall(const plugify::Method* method, void* addr, const plugify::Parameters* p, uint8_t count, const plugify::ReturnValue* ret) {
	std::scoped_lock<std::mutex> lock(g_golm._mutex);
	ArgumentList args;

	StringStorage stringStorage;
	BoolStorage boolStorage;

	DCCallVM* vm = g_golm._callVirtMachine.get();
	dcReset(vm);

	bool hasRet = method->retType.type > ValueType::LastPrimitive;
	bool hasRefs = false;

	// Store parameters

	if (hasRet) {
		switch (method->retType.type) {
			// GoString*
			case ValueType::String:
				dcArgPointer(vm, utils::CreateGoString(args));
				--count;
				break;
			// GoSlice*
			case ValueType::ArrayBool:
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
				dcArgPointer(vm, utils::CreateGoSlice(args));
				--count;
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
				case ValueType::ArrayBool:
					dcArgPointer(vm, utils::CreateGoSliceBool(*p->GetArgument<std::vector<bool>*>(i), args, boolStorage));
					break;
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
					dcArgPointer(vm, utils::CreateGoSliceString(*p->GetArgument<std::vector<std::string>*>(i), args, stringStorage));
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
					dcArgShort(vm, p->GetArgument<short>(i));
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
				case ValueType::String:
#if GOLM_PLATFORM_WINDOWS
					dcArgPointer(vm, utils::CreateGoString(*p->GetArgument<std::string*>(i), args));
					break;
#else
				{
					GoString* string = utils::CreateGoString(*p->GetArgument<std::string*>(i), args);
					dcArgPointer(vm, const_cast<char*>(string->p));
					dcArgLongLong(vm, string->n);
					break;
				}
#endif
				// GoSlice*
				case ValueType::ArrayBool:
					dcArgPointer(vm, utils::CreateGoSliceBool(*p->GetArgument<std::vector<bool>*>(i), args, boolStorage));
					break;
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
					dcArgPointer(vm, utils::CreateGoSliceString(*p->GetArgument<std::vector<std::string>*>(i), args, stringStorage));
					break;
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
		case ValueType::Vector2: {
			dcCallVoid(vm, addr);
			std::construct_at(p->GetArgument<Vector2*>(0), *reinterpret_cast<Vector2*>(args[0]));
			break;
		}
		case ValueType::Vector3: {
			dcCallVoid(vm, addr);
			std::construct_at(p->GetArgument<Vector3*>(0), *reinterpret_cast<Vector3*>(args[0]));
			break;
		}
		case ValueType::Vector4: {
			dcCallVoid(vm, addr);
			std::construct_at(p->GetArgument<Vector4*>(0), *reinterpret_cast<Vector4*>(args[0]));
			break;
		}
		case ValueType::Matrix4x4: {
			dcCallVoid(vm, addr);
			std::construct_at(p->GetArgument<Matrix4x4*>(0), *reinterpret_cast<Matrix4x4*>(args[0]));
			break;
		}
		case ValueType::String: {
			dcCallVoid(vm, addr);
			utils::CopyGoStringToStringReturn(reinterpret_cast<GoString*>(args[0]), p->GetArgument<std::string*>(0));
			break;
		}
		case ValueType::ArrayBool: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<bool>*>(0));
			break;
		}
		case ValueType::ArrayChar8: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<char>*>(0));
			break;
		}
		case ValueType::ArrayChar16: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<char16_t>*>(0));
			break;
		}
		case ValueType::ArrayInt8: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<int8_t>*>(0));
			break;
		}
		case ValueType::ArrayInt16: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<int16_t>*>(0));
			break;
		}
		case ValueType::ArrayInt32: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<int32_t>*>(0));
			break;
		}
		case ValueType::ArrayInt64: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<int64_t>*>(0));
			break;
		}
		case ValueType::ArrayUInt8: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<uint8_t>*>(0));
			break;
		}
		case ValueType::ArrayUInt16: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<uint16_t>*>(0));
			break;
		}
		case ValueType::ArrayUInt32: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<uint32_t>*>(0));
			break;
		}
		case ValueType::ArrayUInt64: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<uint64_t>*>(0));
			break;
		}
		case ValueType::ArrayPointer: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<uintptr_t>*>(0));
			break;
		}
		case ValueType::ArrayFloat: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<float>*>(0));
			break;
		}
		case ValueType::ArrayDouble: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<double>*>(0));
			break;
		}
		case ValueType::ArrayString: {
			dcCallVoid(vm, addr);
			utils::CopyGoSliceToVectorReturn(reinterpret_cast<GoSlice*>(args[0]), p->GetArgument<std::vector<std::string>*>(0));
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
                        case ValueType::ArrayBool:
							utils::CopyGoSliceToVector(reinterpret_cast<GoSlice*>(args[j++]), *p->GetArgument<std::vector<bool>*>(i));
							break;
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
					utils::DeleteGoString(args[j++]);
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
					utils::DeleteGoSlice(args[j++]);
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
						utils::DeleteGoSlice(args[j++]);
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

void* AllocateString() {
	return malloc(sizeof(std::string));
}
void* CreateString(GoString source) {
	return source.n == 0 ? new std::string() : new std::string(source.p, source.n);
}
const char* GetStringData(void* ptr) {
	return reinterpret_cast<std::string*>(ptr)->c_str();
}
ptrdiff_t GetStringSize(void* ptr) {
	return static_cast<ptrdiff_t>(reinterpret_cast<std::string*>(ptr)->size());
}
void AssignString(void* ptr, GoString source) {
	auto str = reinterpret_cast<std::string*>(ptr);
	if (source.p == nullptr || source.n == 0)
		str->clear();
	else
		str->assign(source.p, static_cast<size_t>(source.n));
}
void FreeString(void* ptr) {
	reinterpret_cast<std::string*>(ptr)->~basic_string();
	free(ptr);
}

void DeleteString(void* ptr) {
	delete reinterpret_cast<std::string*>(ptr);
}

enum DataType {
	BOOL,
	CHAR8,
	CHAR16,
	INT8,
	INT16,
	INT32,
	INT64,
	UINT8,
	UINT16,
	UINT32,
	UINT64,
	UINTPTR,
	FLOAT,
	DOUBLE,
	STRING
};

void* CreateVector(void* arr, ptrdiff_t len, DataType type) {
	switch (type) {
		case BOOL:
			return len == 0 ? new std::vector<bool>() : new std::vector<bool>(static_cast<bool*>(arr), static_cast<bool*>(arr) + len);
		case CHAR8:
			return len == 0 ? new std::vector<char>() : new std::vector<char>(static_cast<char*>(arr), static_cast<char*>(arr) + len);
		case CHAR16:
			return len == 0 ? new std::vector<uint16_t>() : new std::vector<uint16_t>(static_cast<uint16_t*>(arr), static_cast<uint16_t*>(arr) + len);
		case INT8:
			return len == 0 ? new std::vector<int8_t>() : new std::vector<int8_t>(static_cast<int8_t*>(arr), static_cast<int8_t*>(arr) + len);
		case INT16:
			return len == 0 ? new std::vector<int16_t>() : new std::vector<int16_t>(static_cast<int16_t*>(arr), static_cast<int16_t*>(arr) + len);
		case INT32:
			return len == 0 ? new std::vector<int32_t>() : new std::vector<int32_t>(static_cast<int32_t*>(arr), static_cast<int32_t*>(arr) + len);
		case INT64:
			return len == 0 ? new std::vector<int64_t>() : new std::vector<int64_t>(static_cast<int64_t*>(arr), static_cast<int64_t*>(arr) + len);
		case UINT8:
			return len == 0 ? new std::vector<uint8_t>() : new std::vector<uint8_t>(static_cast<uint8_t*>(arr), static_cast<uint8_t*>(arr) + len);
		case UINT16:
			return len == 0 ? new std::vector<uint16_t>() : new std::vector<uint16_t>(static_cast<uint16_t*>(arr), static_cast<uint16_t*>(arr) + len);
		case UINT32:
			return len == 0 ? new std::vector<uint32_t>() : new std::vector<uint32_t>(static_cast<uint32_t*>(arr), static_cast<uint32_t*>(arr) + len);
		case UINT64:
			return len == 0 ? new std::vector<uint64_t>() : new std::vector<uint64_t>(static_cast<uint64_t*>(arr), static_cast<uint64_t*>(arr) + len);
		case UINTPTR:
			return len == 0 ? new std::vector<uintptr_t>() : new std::vector<uintptr_t>(static_cast<uintptr_t*>(arr), static_cast<uintptr_t*>(arr) + len);
		case FLOAT:
			return len == 0 ? new std::vector<float>() : new std::vector<float>(static_cast<float*>(arr), static_cast<float*>(arr) + len);
		case DOUBLE:
			return len == 0 ? new std::vector<double>() : new std::vector<double>(static_cast<double*>(arr), static_cast<double*>(arr) + len);
		case STRING:
		{
			auto* vector = new std::vector<std::string>();
			if (len != 0) {
				vector->reserve(len);
				for (ptrdiff_t i = 0; i < len; ++i) {
					const auto& str = static_cast<GoString*>(arr)[i];
					vector->emplace_back(str.p, str.n);
				}
			}
			return vector;
		}
		default:
			return nullptr;
	}
}

void* AllocateVector(DataType type) {
	switch (type) {
		case BOOL:
			return malloc(sizeof(std::vector<bool>));
		case CHAR8:
			return malloc(sizeof(std::vector<char>));
		case CHAR16:
			return malloc(sizeof(std::vector<uint16_t>));
		case INT8:
			return malloc(sizeof(std::vector<int8_t>));
		case INT16:
			return malloc(sizeof(std::vector<int16_t>));
		case INT32:
			return malloc(sizeof(std::vector<int32_t>));
		case INT64:
			return malloc(sizeof(std::vector<int64_t>));
		case UINT8:
			return malloc(sizeof(std::vector<uint8_t>));
		case UINT16:
			return malloc(sizeof(std::vector<uint16_t>));
		case UINT32:
			return malloc(sizeof(std::vector<uint32_t>));
		case UINT64:
			return malloc(sizeof(std::vector<uint64_t>));
		case UINTPTR:
			return malloc(sizeof(std::vector<uintptr_t>));
		case FLOAT:
			return malloc(sizeof(std::vector<float>));
		case DOUBLE:
			return malloc(sizeof(std::vector<double>));
		case STRING:
			return malloc(sizeof(std::vector<std::string>));
		default:
			return nullptr;
	}
}

ptrdiff_t GetVectorSize(void* ptr, DataType type) {
	switch (type) {
		case BOOL:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<bool>*>(ptr)->size());
		case CHAR8:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<char>*>(ptr)->size());
		case CHAR16:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint16_t>*>(ptr)->size());
		case INT8:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int8_t>*>(ptr)->size());
		case INT16:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int16_t>*>(ptr)->size());
		case INT32:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int32_t>*>(ptr)->size());
		case INT64:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int64_t>*>(ptr)->size());
		case UINT8:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint8_t>*>(ptr)->size());
		case UINT16:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint16_t>*>(ptr)->size());
		case UINT32:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint32_t>*>(ptr)->size());
		case UINT64:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint64_t>*>(ptr)->size());
		case UINTPTR:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uintptr_t>*>(ptr)->size());
		case FLOAT:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<float>*>(ptr)->size());
		case DOUBLE:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<double>*>(ptr)->size());
		case STRING:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<std::string>*>(ptr)->size());
		default:
			return -1; // Return -1 or some error code for invalid type
	}
}

void* GetVectorData(void* ptr, DataType type) {
	switch (type) {
		case CHAR8:
			return reinterpret_cast<std::vector<char>*>(ptr)->data();
		case CHAR16:
			return reinterpret_cast<std::vector<uint16_t>*>(ptr)->data();
		case INT8:
			return reinterpret_cast<std::vector<int8_t>*>(ptr)->data();
		case INT16:
			return reinterpret_cast<std::vector<int16_t>*>(ptr)->data();
		case INT32:
			return reinterpret_cast<std::vector<int32_t>*>(ptr)->data();
		case INT64:
			return reinterpret_cast<std::vector<int64_t>*>(ptr)->data();
		case UINT8:
			return reinterpret_cast<std::vector<uint8_t>*>(ptr)->data();
		case UINT16:
			return reinterpret_cast<std::vector<uint16_t>*>(ptr)->data();
		case UINT32:
			return reinterpret_cast<std::vector<uint32_t>*>(ptr)->data();
		case UINT64:
			return reinterpret_cast<std::vector<uint64_t>*>(ptr)->data();
		case UINTPTR:
			return reinterpret_cast<std::vector<uintptr_t>*>(ptr)->data();
		case FLOAT:
			return reinterpret_cast<std::vector<float>*>(ptr)->data();
		case DOUBLE:
			return reinterpret_cast<std::vector<double>*>(ptr)->data();
		case BOOL: {
			auto& vector = *reinterpret_cast<std::vector<bool>*>(ptr);

			bool* boolArray = new bool[vector.size()];

			// Manually copy values from std::vector<bool> to the bool array
			for (size_t i = 0; i < vector.size(); ++i) {
				boolArray[i] = vector[i];
			}

			return boolArray;
		}
		case STRING: {
			auto& vector = *reinterpret_cast<std::vector<std::string>*>(ptr);

			char** strArray = new char*[vector.size()];

			// Manually copy values from std::vector<std::string> to the char* array
			for (size_t i = 0; i < vector.size(); ++i) {
				strArray[i] = vector[i].data();
			}

			return strArray;
		}
		default:
			return nullptr; // Return nullptr for invalid type
	}
}

void AssignVector(void* ptr, void* arr, ptrdiff_t len, DataType type) {
	switch (type) {
		case BOOL: {
			auto* vector = reinterpret_cast<std::vector<bool>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<bool*>(arr), static_cast<bool*>(arr) + len);
			break;
		}
		case CHAR8: {
			auto* vector = reinterpret_cast<std::vector<char>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<char*>(arr), static_cast<char*>(arr) + len);
			break;
		}
		case CHAR16: {
			auto* vector = reinterpret_cast<std::vector<uint16_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<uint16_t*>(arr), static_cast<uint16_t*>(arr) + len);
			break;
		}
		case INT8: {
			auto* vector = reinterpret_cast<std::vector<int8_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<int8_t*>(arr), static_cast<int8_t*>(arr) + len);
			break;
		}
		case INT16: {
			auto* vector = reinterpret_cast<std::vector<int16_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<int16_t*>(arr), static_cast<int16_t*>(arr) + len);
			break;
		}
		case INT32: {
			auto* vector = reinterpret_cast<std::vector<int32_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<int32_t*>(arr), static_cast<int32_t*>(arr) + len);
			break;
		}
		case INT64: {
			auto* vector = reinterpret_cast<std::vector<int64_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<int64_t*>(arr), static_cast<int64_t*>(arr) + len);
			break;
		}
		case UINT8: {
			auto* vector = reinterpret_cast<std::vector<uint8_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<uint8_t*>(arr), static_cast<uint8_t*>(arr) + len);
			break;
		}
		case UINT16: {
			auto* vector = reinterpret_cast<std::vector<uint16_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<uint16_t*>(arr), static_cast<uint16_t*>(arr) + len);
			break;
		}
		case UINT32: {
			auto* vector = reinterpret_cast<std::vector<uint32_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<uint32_t*>(arr), static_cast<uint32_t*>(arr) + len);
			break;
		}
		case UINT64: {
			auto* vector = reinterpret_cast<std::vector<uint64_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<uint64_t*>(arr), static_cast<uint64_t*>(arr) + len);
			break;
		}
		case UINTPTR: {
			auto* vector = reinterpret_cast<std::vector<uintptr_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<uintptr_t*>(arr), static_cast<uintptr_t*>(arr) + len);
			break;
		}
		case FLOAT: {
			auto* vector = reinterpret_cast<std::vector<float>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<float*>(arr), static_cast<float*>(arr) + len);
			break;
		}
		case DOUBLE: {
			auto* vector = reinterpret_cast<std::vector<double>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<double*>(arr), static_cast<double*>(arr) + len);
			break;
		}
		case STRING: {
			auto* vector = reinterpret_cast<std::vector<std::string>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else {
				vector->resize(static_cast<size_t>(len));
				for (ptrdiff_t i = 0; i < len; ++i) {
					const auto& str = static_cast<GoString*>(arr)[i];
					vector->at(static_cast<size_t>(i)).assign(str.p, static_cast<size_t>(str.n));
				}
			}
			break;
		}
		default:
			break;
	}
}

void DeleteVector(void* ptr, DataType type) {
	switch (type) {
		case BOOL:
			delete reinterpret_cast<std::vector<bool>*>(ptr);
			break;
		case CHAR8:
			delete reinterpret_cast<std::vector<char>*>(ptr);
			break;
		case CHAR16:
			delete reinterpret_cast<std::vector<uint16_t>*>(ptr);
			break;
		case INT8:
			delete reinterpret_cast<std::vector<int8_t>*>(ptr);
			break;
		case INT16:
			delete reinterpret_cast<std::vector<int16_t>*>(ptr);
			break;
		case INT32:
			delete reinterpret_cast<std::vector<int32_t>*>(ptr);
			break;
		case INT64:
			delete reinterpret_cast<std::vector<int64_t>*>(ptr);
			break;
		case UINT8:
			delete reinterpret_cast<std::vector<uint8_t>*>(ptr);
			break;
		case UINT16:
			delete reinterpret_cast<std::vector<uint16_t>*>(ptr);
			break;
		case UINT32:
			delete reinterpret_cast<std::vector<uint32_t>*>(ptr);
			break;
		case UINT64:
			delete reinterpret_cast<std::vector<uint64_t>*>(ptr);
			break;
		case UINTPTR:
			delete reinterpret_cast<std::vector<uintptr_t>*>(ptr);
			break;
		case FLOAT:
			delete reinterpret_cast<std::vector<float>*>(ptr);
			break;
		case DOUBLE:
			delete reinterpret_cast<std::vector<double>*>(ptr);
			break;
		case STRING:
			delete reinterpret_cast<std::vector<std::string>*>(ptr);
			break;
		default:
			// Invalid type, do nothing or handle error if needed
			break;
	}
}

void FreeVector(void* ptr, DataType type) {
	switch (type) {
		case BOOL:
			reinterpret_cast<std::vector<bool>*>(ptr)->~vector();
			break;
		case CHAR8:
			reinterpret_cast<std::vector<char>*>(ptr)->~vector();
			break;
		case CHAR16:
			reinterpret_cast<std::vector<uint16_t>*>(ptr)->~vector();
			break;
		case INT8:
			reinterpret_cast<std::vector<int8_t>*>(ptr)->~vector();
			break;
		case INT16:
			reinterpret_cast<std::vector<int16_t>*>(ptr)->~vector();
			break;
		case INT32:
			reinterpret_cast<std::vector<int32_t>*>(ptr)->~vector();
			break;
		case INT64:
			reinterpret_cast<std::vector<int64_t>*>(ptr)->~vector();
			break;
		case UINT8:
			reinterpret_cast<std::vector<uint8_t>*>(ptr)->~vector();
			break;
		case UINT16:
			reinterpret_cast<std::vector<uint16_t>*>(ptr)->~vector();
			break;
		case UINT32:
			reinterpret_cast<std::vector<uint32_t>*>(ptr)->~vector();
			break;
		case UINT64:
			reinterpret_cast<std::vector<uint64_t>*>(ptr)->~vector();
			break;
		case UINTPTR:
			reinterpret_cast<std::vector<uintptr_t>*>(ptr)->~vector();
			break;
		case FLOAT:
			reinterpret_cast<std::vector<float>*>(ptr)->~vector();
			break;
		case DOUBLE:
			reinterpret_cast<std::vector<double>*>(ptr)->~vector();
			break;
		case STRING:
			reinterpret_cast<std::vector<std::string>*>(ptr)->~vector();
			break;
		default:
			// Invalid type, do nothing or handle error if needed
			return;
	}
	free(ptr);
}

void DeleteVectorDataBool(void* ptr) {
	delete[] reinterpret_cast<bool*>(ptr);
}

void DeleteVectorDataCStr(void* ptr) {
	delete[] reinterpret_cast<char**>(ptr);
}

const std::array<void*, 17> GoLanguageModule::_pluginApi = {
		reinterpret_cast<void*>(&GetMethodPtr),
		reinterpret_cast<void*>(&AllocateString),
		reinterpret_cast<void*>(&CreateString),
		reinterpret_cast<void*>(&GetStringData),
		reinterpret_cast<void*>(&GetStringSize),
		reinterpret_cast<void*>(&AssignString),
		reinterpret_cast<void*>(&FreeString),
		reinterpret_cast<void*>(&DeleteString),
		reinterpret_cast<void*>(&CreateVector),
		reinterpret_cast<void*>(&AllocateVector),
		reinterpret_cast<void*>(&GetVectorSize),
		reinterpret_cast<void*>(&GetVectorData),
		reinterpret_cast<void*>(&AssignVector),
		reinterpret_cast<void*>(&DeleteVector),
		reinterpret_cast<void*>(&FreeVector),
		reinterpret_cast<void*>(&DeleteVectorDataBool),
		reinterpret_cast<void*>(&DeleteVectorDataCStr)
};

namespace golm {
	GoLanguageModule g_golm;
}

plugify::ILanguageModule* GetLanguageModule() {
	return &golm::g_golm;
}