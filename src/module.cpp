#include "module.h"

#include <plugify/module.h>
#include <plugify/plugin.h>
#include <plugify/log.h>
#include <plugify/math.h>
#include <plugify/plugin_descriptor.h>
#include <plugify/plugin_reference_descriptor.h>
#include <plugify/plugify_provider.h>

#if GOLM_PLATFORM_WINDOWS
#include <Windows.h>
#undef FindResource
#endif

#define LOG_PREFIX "[GOLM] "

using namespace golm;
using namespace plugify;

template<class T>
inline constexpr bool always_false_v = std::is_same_v<std::decay_t<T>, std::add_cv_t<std::decay_t<T>>>;

bool IsMethodPrimitive(MethodRef method) {
	if (ValueUtils::IsObject(method.GetReturnType().GetType()))
		return false;

	for (const auto& param : method.GetParamTypes()) {
		if (ValueUtils::IsObject(param.GetType()))
			return false;
	}

	return true;
}

GoSlice* CreateGoSliceBool(const std::vector<bool>& source, ArgumentList& args, BoolHolder& holder) {
	assert(args.size() != args.capacity() && "Resizing list will invalidate pointers!");
	size_t N = source.size();
	auto& boolArray = holder.emplace_back(std::make_unique<bool[]>(N));
	for (size_t i = 0; i < N; ++i) {
		boolArray[i] = source[i];
	}
	auto size = static_cast<GoInt>(N);
	auto& slice = args.emplace_back(boolArray.get(), size, size);
	return &slice;
}

GoSlice* CreateGoSliceString(const std::vector<std::string>& source, ArgumentList& args, StringHolder& holder) {
	assert(args.size() != args.capacity() && "Resizing list will invalidate pointers!");
	size_t N = source.size();
	auto& strArray = holder.emplace_back(std::make_unique<GoString[]>(N));
	for (size_t i = 0; i < N; ++i) {
		const auto& str = source[i];
		auto& dest = strArray[i];
		dest.p = str.c_str();
		dest.n = static_cast<GoInt>(str.length());
	}
	auto size = static_cast<GoInt>(N);
	auto& slice = args.emplace_back(strArray.get(), size, size);
	return &slice;
}

template<typename T>
GoSlice* CreateGoSlice(const std::vector<T>& source, ArgumentList& args) {
	assert(args.size() != args.capacity() && "Resizing list will invalidate pointers!");
	auto size = static_cast<GoInt>(source.size());
	auto& slice = args.emplace_back((void*)source.data(), size, size);
	return &slice;
}

GoSlice* CreateGoSlice(ArgumentList& args) {
	assert(args.size() != args.capacity() && "Resizing list will invalidate pointers!");
	auto& slice = args.emplace_back(nullptr, 0, 0);
	return &slice;
}

GoString* CreateGoString(const std::string& source, ArgumentList& args) {
	assert(args.size() != args.capacity() && "Resizing list will invalidate pointers!");
	auto size = static_cast<GoInt>(source.length());
	auto& slice = args.emplace_back((void*)source.c_str(), size, 0);
	return reinterpret_cast<GoString*>(&slice);
}

GoString* CreateGoString(ArgumentList& args) {
	assert(args.size() != args.capacity() && "Resizing list will invalidate pointers!");
	auto& slice = args.emplace_back((void*)"", 0, 0);
	return reinterpret_cast<GoString*>(&slice);
}

template<typename T>
void CopyGoSliceToVector(const GoSlice& source, std::vector<T>& dest) {
	if (source.data == nullptr || source.len == 0)
		dest.clear();
	else if (dest.data() != source.data)
		dest.assign(reinterpret_cast<T*>(source.data), reinterpret_cast<T*>(source.data) + static_cast<size_t>(source.len));
}

template<>
void CopyGoSliceToVector(const GoSlice& source, std::vector<bool>& dest) {
	dest.resize(static_cast<size_t>(source.len));
	for (size_t i = 0; i < dest.size(); ++i) {
		dest[i] = reinterpret_cast<bool*>(source.data)[i];
	}
}

template<>
void CopyGoSliceToVector(const GoSlice& source, std::vector<std::string>& dest) {
	dest.resize(static_cast<size_t>(source.len));
	for (size_t i = 0; i < dest.size(); ++i) {
		const auto& str = reinterpret_cast<GoString*>(source.data)[i];
		dest[i].assign(str.p, static_cast<size_t>(str.n));
	}
}

void CopyGoStringToString(const GoString& source, std::string& dest) {
	if (source.p == nullptr || source.n == 0)
		dest.clear();
	else if (dest.data() != source.p)
		dest.assign(source.p, static_cast<size_t>(source.n));
}

template<typename T>
void CopyGoSliceToVectorReturn(const GoSlice& source, std::vector<T>& dest) {
	if (source.data == nullptr || source.len == 0)
		std::construct_at(&dest, std::vector<T>());
	else
		std::construct_at(&dest, std::vector<T>(reinterpret_cast<T*>(source.data), reinterpret_cast<T*>(source.data) + static_cast<size_t>(source.len)));
}

template<>
void CopyGoSliceToVectorReturn(const GoSlice& source, std::vector<bool>& dest) {
	if (source.data == nullptr || source.len == 0)
		std::construct_at(&dest, std::vector<bool>());
	else {
		std::construct_at(&dest, std::vector<bool>(static_cast<size_t>(source.len)));
		for (size_t i = 0; i < dest.size(); ++i) {
			dest[i] = reinterpret_cast<bool*>(source.data)[i];
		}
	}
}

template<>
void CopyGoSliceToVectorReturn(const GoSlice& source, std::vector<std::string>& dest) {
	if (source.data == nullptr || source.len == 0)
		std::construct_at(&dest, std::vector<std::string>());
	else {
		std::construct_at(&dest, std::vector<std::string>(static_cast<size_t>(source.len)));
		for (size_t i = 0; i < dest.size(); ++i) {
			const auto& str = reinterpret_cast<GoString*>(source.data)[i];
			dest[i].assign(str.p, static_cast<size_t>(str.n));
		}
	}
}

void CopyGoStringToStringReturn(const GoString& source, std::string& dest) {
	if (source.p == nullptr || source.n == 0)
		std::construct_at(&dest, std::string());
	else
		std::construct_at(&dest, std::string(source.p, static_cast<size_t>(source.n)));
}

template<typename T>
DCaggr* CreateDcAggr(AggrList& aggrs) {
	static_assert(always_false_v<T>, "CreateDcAggr specialization required");
	return nullptr;
}

template<>
DCaggr* CreateDcAggr<Vector2>(AggrList& aggrs) {
	DCaggr* ag = dcNewAggr(2, sizeof(Vector2));
	for (size_t i = 0; i < 2; ++i)
		dcAggrField(ag, DC_SIGCHAR_FLOAT, static_cast<int>(sizeof(float) * i), 1);
	dcCloseAggr(ag);
	aggrs.emplace_back(std::unique_ptr<DCaggr>(ag));
	return ag;
}

template<>
DCaggr* CreateDcAggr<Vector3>(AggrList& aggrs) {
	DCaggr* ag = dcNewAggr(3, sizeof(Vector3));
	for (size_t i = 0; i < 3; ++i)
		dcAggrField(ag, DC_SIGCHAR_FLOAT, static_cast<int>(sizeof(float) * i), 1);
	dcCloseAggr(ag);
	aggrs.emplace_back(std::unique_ptr<DCaggr>(ag));
	return ag;
}

template<>
DCaggr* CreateDcAggr<Vector4>(AggrList& aggrs) {
	DCaggr* ag = dcNewAggr(4, sizeof(Vector4));
	for (size_t i = 0; i < 4; ++i)
		dcAggrField(ag, DC_SIGCHAR_FLOAT, static_cast<int>(sizeof(float) * i), 1);
	dcCloseAggr(ag);
	aggrs.emplace_back(std::unique_ptr<DCaggr>(ag));
	return ag;
}

template<>
DCaggr* CreateDcAggr<Matrix4x4>(AggrList& aggrs) {
	DCaggr* ag = dcNewAggr(16, sizeof(Matrix4x4));
	for (size_t i = 0; i < 16; ++i)
		dcAggrField(ag, DC_SIGCHAR_FLOAT, static_cast<int>(sizeof(float) * i), 1);
	dcCloseAggr(ag);
	aggrs.emplace_back(std::unique_ptr<DCaggr>(ag));
	return ag;
}

template<>
DCaggr* CreateDcAggr<GoString>(AggrList& aggrs) {
	DCaggr* ag = dcNewAggr(2, sizeof(GoString));
	dcAggrField(ag, DC_SIGCHAR_STRING, 0, 1);
	dcAggrField(ag, DC_SIGCHAR_LONGLONG, sizeof(const char*), 1);
	dcCloseAggr(ag);
	aggrs.emplace_back(std::unique_ptr<DCaggr>(ag));
	return ag;
}

template<>
DCaggr* CreateDcAggr<GoSlice>(AggrList& aggrs) {
	DCaggr* ag = dcNewAggr(3, sizeof(GoSlice));
	dcAggrField(ag, DC_SIGCHAR_POINTER, 0, 1);
	dcAggrField(ag, DC_SIGCHAR_LONGLONG, sizeof(void*), 1);
	dcAggrField(ag, DC_SIGCHAR_LONGLONG, sizeof(void*) * 2, 1);
	dcCloseAggr(ag);
	aggrs.emplace_back(std::unique_ptr<DCaggr>(ag));
	return ag;
}

InitResult GoLanguageModule::Initialize(std::weak_ptr<IPlugifyProvider> provider, ModuleRef /*module*/) {
	if (!(_provider = provider.lock())) {
		return ErrorData{ "Provider not exposed" };
	}

	_provider->Log(LOG_PREFIX "Inited!", Severity::Debug);

	_rt = std::make_shared<asmjit::JitRuntime>();

	DCCallVM* vm = dcNewCallVM(4096);
	dcMode(vm, DC_CALL_C_DEFAULT);
	_callVirtMachine = std::unique_ptr<DCCallVM>(vm);

	return InitResultData{};
}

void GoLanguageModule::Shutdown() {
	for (MemAddr* addr : _addresses) {
		*addr = nullptr;
	}
	_nativesMap.clear();
	_functions.clear();
	_addresses.clear();
	_assemblyMap.clear();
	_callVirtMachine.reset();
	_rt.reset();
	_provider.reset();
}

void GoLanguageModule::OnMethodExport(PluginRef plugin) {
	auto pluginName = plugin.GetName();
	for (const auto& [name, addr] : plugin.GetMethods()) {
		_nativesMap.try_emplace(std::format("{}.{}", pluginName, name), addr);
	}
}

LoadResult GoLanguageModule::OnPluginLoad(PluginRef plugin) {
	fs::path assemblyPath(plugin.GetBaseDir());
	assemblyPath /= std::format("{}" GOLM_LIBRARY_SUFFIX, plugin.GetDescriptor().GetEntryPoint());

	auto assembly = std::make_unique<Assembly>(assemblyPath, LoadFlag::Lazy | LoadFlag::Nodelete | LoadFlag::PinInMemory);
	if (!assembly) {
		return ErrorData{ std::format("Failed to load assembly: {}", assembly->GetError()) };
	}

	std::vector<std::string_view> funcErrors;

	auto* const initFunc = assembly->GetFunctionByName("Plugify_Init").RCast<InitFunc>();
	if (!initFunc) {
		funcErrors.emplace_back("Plugify_Init");
	}

	auto* const startFunc = assembly->GetFunctionByName("Plugify_PluginStart").RCast<StartFunc>();
	if (!startFunc) {
		funcErrors.emplace_back("Plugify_PluginStart");
	}

	auto* const endFunc = assembly->GetFunctionByName("Plugify_PluginEnd").RCast<EndFunc>();
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

	std::vector<MethodRef> exportedMethods = plugin.GetDescriptor().GetExportedMethods();
	std::vector<MethodData> methods;
	methods.reserve(exportedMethods.size());

	for (const auto& method : exportedMethods) {
		if (auto func = assembly->GetFunctionByName(method.GetFunctionName())) {
			if (IsMethodPrimitive(method)) {
				methods.emplace_back(method.GetName(), func);
			} else {
				auto function = std::make_unique<Function>(_rt);
				func = function->GetJitFunc(method, &InternalCall, func);
				if (!func) {
					funcErrors.emplace_back(method.GetName());
					continue;
				}
				_functions.emplace_back(std::move(function));
				methods.emplace_back(method.GetName(), func);
			}
		} else {
			funcErrors.emplace_back(method.GetName());
		}
	}
	if (!funcErrors.empty()) {
		std::string funcs(funcErrors[0]);
		for (auto it = std::next(funcErrors.begin()); it != funcErrors.end(); ++it) {
			std::format_to(std::back_inserter(funcs), ", {}", *it);
		}
		return ErrorData{ std::format("Not found {} method function(s)", funcs) };
	}

	union {
		PluginRef plugin;
		void* ptr;
	} cast{plugin};

	GoSlice api { const_cast<void**>(_pluginApi.data()), _pluginApi.size(), _pluginApi.size() };
	const int resultVersion = initFunc(api, kApiVersion, cast.ptr);
	if (resultVersion != 0) {
		return ErrorData{ std::format("Not supported plugin api {}, max supported {}", resultVersion, kApiVersion) };
	}

	const auto [_, result] = _assemblyMap.try_emplace(plugin.GetId(), std::move(assembly), startFunc, endFunc);
	if (!result) {
		return ErrorData{ std::format("Plugin id duplicate") };
	}

	return LoadResultData{ std::move(methods) };
}

void GoLanguageModule::OnPluginStart(PluginRef plugin) {
	if (const auto it = _assemblyMap.find(plugin.GetId()); it != _assemblyMap.end()) {
		const auto& assemblyHolder = std::get<AssemblyHolder>(*it);
		assemblyHolder.GetStartFunc()();
	}
}

void GoLanguageModule::OnPluginEnd(PluginRef plugin) {
	if (const auto it = _assemblyMap.find(plugin.GetId()); it != _assemblyMap.end()) {
		const auto& assemblyHolder = std::get<AssemblyHolder>(*it);
		assemblyHolder.GetEndFunc()();
	}
}

MemAddr GoLanguageModule::GetNativeMethod(std::string_view methodName) const {
	if (const auto it = _nativesMap.find(methodName); it != _nativesMap.end()) {
		return std::get<MemAddr>(*it);
	}
	_provider->Log(std::format(LOG_PREFIX "GetNativeMethod failed to find: '{}'", methodName), Severity::Fatal);
	return nullptr;
}

void GoLanguageModule::GetNativeMethod(std::string_view methodName, plugify::MemAddr* addressDest) {
	if (const auto it = _nativesMap.find(methodName); it != _nativesMap.end()) {
		*addressDest = std::get<MemAddr>(*it);
		_addresses.emplace_back(addressDest);
		return;
	}
	_provider->Log(std::format(LOG_PREFIX "GetNativeMethod failed to find: '{}'", methodName), Severity::Fatal);
}

// C++ to Go
void GoLanguageModule::InternalCall(MethodRef method, MemAddr addr, const Parameters* p, uint8_t count, const ReturnValue* ret) {
	std::scoped_lock<std::mutex> lock(g_golm._mutex);

	PropertyRef retProp = method.GetReturnType();
	ValueType retType = retProp.GetType();
	std::vector<PropertyRef> paramProps = method.GetParamTypes();

	size_t argsCount = static_cast<size_t>(std::count_if(paramProps.begin(), paramProps.end(), [](const PropertyRef& param) {
		return ValueUtils::IsObject(param.GetType());
	}));

	AggrList aggrs;
	ArgumentList args;
	args.reserve(argsCount);

	StringHolder stringHolder;
	BoolHolder boolHolder;

	DCCallVM* vm = g_golm._callVirtMachine.get();
	dcReset(vm);

	bool hasRet = ValueUtils::IsHiddenParam(retType);
	bool hasRefs = false;

	DCaggr* ag = nullptr;

	switch (retType) {
		case ValueType::String:
			ag = CreateDcAggr<GoString>(aggrs);
			dcBeginCallAggr(vm, ag);
			break;
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
			ag = CreateDcAggr<GoSlice>(aggrs);
			dcBeginCallAggr(vm, ag);
			break;
		case ValueType::Vector2:
			ag = CreateDcAggr<Vector2>(aggrs);
			dcBeginCallAggr(vm, ag);
			break;
		case ValueType::Vector3:
			ag = CreateDcAggr<Vector3>(aggrs);
			dcBeginCallAggr(vm, ag);
			break;
		case ValueType::Vector4:
			ag = CreateDcAggr<Vector4>(aggrs);
			dcBeginCallAggr(vm, ag);
			break;
		case ValueType::Matrix4x4:
			ag = CreateDcAggr<Matrix4x4>(aggrs);
			dcBeginCallAggr(vm, ag);
			break;
		default:
			// Should not require storage
			break;
	}

	for (uint8_t i = hasRet, j = 0; i < count; ++i, ++j) {
		const auto& param = paramProps[j];
		if (param.IsReference()) {
			switch (param.GetType()) {
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
					dcArgPointer(vm, CreateGoString(*p->GetArgument<std::string*>(i), args));
					break;
				// GoSlice*
				case ValueType::ArrayBool:
					dcArgPointer(vm, CreateGoSliceBool(*p->GetArgument<std::vector<bool>*>(i), args, boolHolder));
					break;
				case ValueType::ArrayChar8:
					dcArgPointer(vm, CreateGoSlice<char>(*p->GetArgument<std::vector<char>*>(i), args));
					break;
				case ValueType::ArrayChar16:
					dcArgPointer(vm, CreateGoSlice<char16_t>(*p->GetArgument<std::vector<char16_t>*>(i), args));
					break;
				case ValueType::ArrayInt8:
					dcArgPointer(vm, CreateGoSlice<int8_t>(*p->GetArgument<std::vector<int8_t>*>(i), args));
					break;
				case ValueType::ArrayInt16:
					dcArgPointer(vm, CreateGoSlice<int16_t>(*p->GetArgument<std::vector<int16_t>*>(i), args));
					break;
				case ValueType::ArrayInt32:
					dcArgPointer(vm, CreateGoSlice<int32_t>(*p->GetArgument<std::vector<int32_t>*>(i), args));
					break;
				case ValueType::ArrayInt64:
					dcArgPointer(vm, CreateGoSlice<int64_t>(*p->GetArgument<std::vector<int64_t>*>(i), args));
					break;
				case ValueType::ArrayUInt8:
					dcArgPointer(vm, CreateGoSlice<uint8_t>(*p->GetArgument<std::vector<uint8_t>*>(i), args));
					break;
				case ValueType::ArrayUInt16:
					dcArgPointer(vm, CreateGoSlice<uint16_t>(*p->GetArgument<std::vector<uint16_t>*>(i), args));
					break;
				case ValueType::ArrayUInt32:
					dcArgPointer(vm, CreateGoSlice<uint32_t>(*p->GetArgument<std::vector<uint32_t>*>(i), args));
					break;
				case ValueType::ArrayUInt64:
					dcArgPointer(vm, CreateGoSlice<uint64_t>(*p->GetArgument<std::vector<uint64_t>*>(i), args));
					break;
				case ValueType::ArrayPointer:
					dcArgPointer(vm, CreateGoSlice<uintptr_t>(*p->GetArgument<std::vector<uintptr_t>*>(i), args));
					break;
				case ValueType::ArrayFloat:
					dcArgPointer(vm, CreateGoSlice<float>(*p->GetArgument<std::vector<float>*>(i), args));
					break;
				case ValueType::ArrayDouble:
					dcArgPointer(vm, CreateGoSlice<double>(*p->GetArgument<std::vector<double>*>(i), args));
					break;
				case ValueType::ArrayString:
					dcArgPointer(vm, CreateGoSliceString(*p->GetArgument<std::vector<std::string>*>(i), args, stringHolder));
					break;
				default:
					std::puts(LOG_PREFIX "Unsupported types!\n");
					std::terminate();
					break;
			}
		} else {
			switch (param.GetType()) {
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
					dcArgAggr(vm, CreateDcAggr<Vector2>(aggrs), p->GetArgument<Vector2*>(i));
					break;
				case ValueType::Vector3:
					dcArgAggr(vm, CreateDcAggr<Vector3>(aggrs), p->GetArgument<Vector3*>(i));
					break;
				case ValueType::Vector4:
					dcArgAggr(vm, CreateDcAggr<Vector4>(aggrs), p->GetArgument<Vector4*>(i));
					break;
				case ValueType::Matrix4x4:
					dcArgAggr(vm, CreateDcAggr<Matrix4x4>(aggrs), p->GetArgument<Matrix4x4*>(i));
					break;
				// GoString
				case ValueType::String:
					dcArgAggr(vm, CreateDcAggr<GoString>(aggrs), CreateGoString(*p->GetArgument<std::string*>(i), args));
					break;
				// GoSlice
				case ValueType::ArrayBool:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSliceBool(*p->GetArgument<std::vector<bool>*>(i), args, boolHolder));
					break;
				case ValueType::ArrayChar8:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<char>(*p->GetArgument<std::vector<char>*>(i), args));
					break;
				case ValueType::ArrayChar16:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<char16_t>(*p->GetArgument<std::vector<char16_t>*>(i), args));
					break;
				case ValueType::ArrayInt8:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<int8_t>(*p->GetArgument<std::vector<int8_t>*>(i), args));
					break;
				case ValueType::ArrayInt16:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<int16_t>(*p->GetArgument<std::vector<int16_t>*>(i), args));
					break;
				case ValueType::ArrayInt32:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<int32_t>(*p->GetArgument<std::vector<int32_t>*>(i), args));
					break;
				case ValueType::ArrayInt64:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<int64_t>(*p->GetArgument<std::vector<int64_t>*>(i), args));
					break;
				case ValueType::ArrayUInt8:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<uint8_t>(*p->GetArgument<std::vector<uint8_t>*>(i), args));
					break;
				case ValueType::ArrayUInt16:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<uint16_t>(*p->GetArgument<std::vector<uint16_t>*>(i), args));
					break;
				case ValueType::ArrayUInt32:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<uint32_t>(*p->GetArgument<std::vector<uint32_t>*>(i), args));
					break;
				case ValueType::ArrayUInt64:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<uint64_t>(*p->GetArgument<std::vector<uint64_t>*>(i), args));
					break;
				case ValueType::ArrayPointer:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<uintptr_t>(*p->GetArgument<std::vector<uintptr_t>*>(i), args));
					break;
				case ValueType::ArrayFloat:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<float>(*p->GetArgument<std::vector<float>*>(i), args));
					break;
				case ValueType::ArrayDouble:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<double>(*p->GetArgument<std::vector<double>*>(i), args));
					break;
				case ValueType::ArrayString:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSliceString(*p->GetArgument<std::vector<std::string>*>(i), args, stringHolder));
					break;
				default:
					std::puts(LOG_PREFIX "Unsupported types!\n");
					std::terminate();
					break;
			}
		}
		hasRefs |= param.IsReference();
	}

	switch (retType) {
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
			Vector2 source;
			dcCallAggr(vm, addr, ag, &source);
			ret->SetReturnPtr(source);
			break;
		}
#if GOLM_PLATFORM_WINDOWS
		case ValueType::Vector3: {
			auto* dest = p->GetArgument<Vector3*>(0);
			dcCallAggr(vm, addr, ag, dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::Vector4: {
			auto* dest = p->GetArgument<Vector4*>(0);
			dcCallAggr(vm, addr, ag, dest);
			ret->SetReturnPtr(dest);
			break;
		}
#else
		case ValueType::Vector3: {
			Vector3 source;
			dcCallAggr(vm, addr, ag, &source);
			ret->SetReturnPtr(source);
			break;
		}
		case ValueType::Vector4: {
			Vector4 source;
			dcCallAggr(vm, addr, ag, &source);
			ret->SetReturnPtr(source);
			break;
		}
#endif
		case ValueType::Matrix4x4: {
			auto* dest = p->GetArgument<Matrix4x4*>(0);
			dcCallAggr(vm, addr, ag, dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::String: {
			auto* dest = p->GetArgument<std::string*>(0);
			GoString source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoStringToStringReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayBool: {
			auto* dest = p->GetArgument<std::vector<bool>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayChar8: {
			auto* dest = p->GetArgument<std::vector<char>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayChar16: {
			auto* dest = p->GetArgument<std::vector<char16_t>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayInt8: {
			auto* dest = p->GetArgument<std::vector<int8_t>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayInt16: {
			auto* dest = p->GetArgument<std::vector<int16_t>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayInt32: {
			auto* dest = p->GetArgument<std::vector<int32_t>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayInt64: {
			auto* dest = p->GetArgument<std::vector<int64_t>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayUInt8: {
			auto* dest = p->GetArgument<std::vector<uint8_t>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayUInt16: {
			auto* dest = p->GetArgument<std::vector<uint16_t>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayUInt32: {
			auto* dest = p->GetArgument<std::vector<uint32_t>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayUInt64: {
			auto* dest = p->GetArgument<std::vector<uint64_t>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayPointer: {
			auto* dest = p->GetArgument<std::vector<uintptr_t>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayFloat: {
			auto* dest = p->GetArgument<std::vector<float>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayDouble: {
			auto* dest = p->GetArgument<std::vector<double>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		case ValueType::ArrayString: {
			auto* dest = p->GetArgument<std::vector<std::string>*>(0);
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *dest);
			ret->SetReturnPtr(dest);
			break;
		}
		default:
			std::puts(LOG_PREFIX "Unsupported types!\n");
			std::terminate();
			break;
	}

	if (argsCount != 0) {
		if (hasRefs) {
			size_t k = 0;
			for (uint8_t i = hasRet, j = 0; i < count; ++i, ++j) {
				const auto& param = paramProps[j];
				if (param.IsReference()) {
					switch (param.GetType()) {
						case ValueType::String:
							CopyGoStringToString(*reinterpret_cast<GoString*>(&args[k++]), *p->GetArgument<std::string*>(i));
							break;
						case ValueType::ArrayBool:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<bool>*>(i));
							break;
						case ValueType::ArrayChar8:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<char>*>(i));
							break;
						case ValueType::ArrayChar16:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<char16_t>*>(i));
							break;
						case ValueType::ArrayInt8:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<int8_t>*>(i));
							break;
						case ValueType::ArrayInt16:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<int16_t>*>(i));
							break;
						case ValueType::ArrayInt32:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<int32_t>*>(i));
							break;
						case ValueType::ArrayInt64:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<int64_t>*>(i));
							break;
						case ValueType::ArrayUInt8:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<uint8_t>*>(i));
							break;
						case ValueType::ArrayUInt16:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<uint16_t>*>(i));
							break;
						case ValueType::ArrayUInt32:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<uint32_t>*>(i));
							break;
						case ValueType::ArrayUInt64:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<uint64_t>*>(i));
							break;
						case ValueType::ArrayPointer:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<uintptr_t>*>(i));
							break;
						case ValueType::ArrayFloat:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<float>*>(i));
							break;
						case ValueType::ArrayDouble:
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<double>*>(i));
							break;
						case ValueType::ArrayString: {
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<std::string>*>(i));
							break;
						}
						default:
							break;
					}
				}
				if (k == argsCount)
					break;
			}
		}
	}
}

namespace golm {
	GoLanguageModule g_golm;
}

// TODO: Detect leaks

MemAddr GetMethodPtr(const char* methodName) {
	return g_golm.GetNativeMethod(methodName);
}

void GetMethodPtr2(const char* methodName, MemAddr* addressDest) {
	g_golm.GetNativeMethod(methodName, addressDest);
}

const char* GetBaseDir() {
	auto source = g_golm.GetProvider()->GetBaseDir().string();
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.c_str(), size);
	return dest;
}

bool IsModuleLoaded(const char* moduleName, int version, bool minimum) {
	auto requiredVersion = (version >= 0 && version != INT_MAX) ? std::make_optional(version) : std::nullopt;
	return g_golm.GetProvider()->IsModuleLoaded(moduleName, requiredVersion, minimum);
}

bool IsPluginLoaded(const char* pluginName, int version, bool minimum) {
	auto requiredVersion = (version >= 0 && version != INT_MAX) ? std::make_optional(version) : std::nullopt;
	return g_golm.GetProvider()->IsPluginLoaded(pluginName, requiredVersion, minimum);
}

UniqueId GetPluginId(PluginRef plugin) {
	return plugin.GetId();
}

const char* GetPluginName(PluginRef plugin) {
	return plugin.GetName().data();
}

const char* GetPluginFullName(PluginRef plugin) {
	return plugin.GetFriendlyName().data();
}

const char* GetPluginDescription(PluginRef plugin) {
	return plugin.GetDescriptor().GetDescription().data();
}

const char* GetPluginVersion(PluginRef plugin) {
	return plugin.GetDescriptor().GetVersionName().data();
}

const char* GetPluginAuthor(PluginRef plugin) {
	return plugin.GetDescriptor().GetCreatedBy().data();
}

const char* GetPluginWebsite(PluginRef plugin) {
	return plugin.GetDescriptor().GetCreatedByURL().data();
}

const char* GetPluginBaseDir(PluginRef plugin) {
	auto source = plugin.GetBaseDir().string();
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.c_str(), size);
	return dest;
}

const char** GetPluginDependencies(PluginRef plugin) {
	std::vector<PluginReferenceDescriptorRef> dependencies = plugin.GetDescriptor().GetDependencies();
	auto* deps = new const char*[dependencies.size()];
	for (size_t i = 0; i < dependencies.size(); ++i) {
		deps[i] = dependencies[i].GetName().data();
	}
	return deps;
}

ptrdiff_t GetPluginDependenciesSize(PluginRef plugin) {
	return static_cast<ptrdiff_t>(plugin.GetDescriptor().GetDependencies().size());
}

const char* FindPluginResource(PluginRef plugin, const char* path) {
	auto resource = plugin.FindResource(path);
	if (resource.has_value()) {
		auto source= resource->string();
		size_t size = source.length() + 1;
		char* dest = new char[size];
		std::memcpy(dest, source.c_str(), size);
		return dest;
	}
	return "";
}
void DeleteCStr(const char* path) {
	delete path;
}

void* AllocateString() {
	return malloc(sizeof(std::string));
}
void* CreateString(GoString source) {
	return source.n == 0 ? new std::string() : new std::string(source.p, static_cast<size_t>(source.n));
}
const char* GetStringData(std::string* ptr) {
	return ptr->c_str();
}
ptrdiff_t GetStringLength(std::string* ptr) {
	return static_cast<ptrdiff_t>(ptr->length());
}
void ConstructString(std::string* ptr, GoString source) {
	std::construct_at(ptr, std::string(source.p, static_cast<size_t>(source.n)));
}
void AssignString(std::string* ptr, GoString source) {
	if (source.p == nullptr || source.n == 0)
		ptr->clear();
	else
		ptr->assign(source.p, static_cast<size_t>(source.n));
}
void FreeString(std::string* ptr) {
	ptr->~basic_string();
	free(ptr);
}

void DeleteString(std::string* ptr) {
	delete ptr;
}

enum DataType {
	Bool,
	Char8,
	Char16,
	Int8,
	Int16,
	Int32,
	Int64,
	UInt8,
	UInt16,
	UInt32,
	UInt64,
	Pointer,
	Float,
	Double,
	String
};

void* CreateVector(void* arr, ptrdiff_t len, DataType type) {
	switch (type) {
		case Bool:
			return len == 0 ? new std::vector<bool>() : new std::vector<bool>(static_cast<bool*>(arr), static_cast<bool*>(arr) + len);
		case Char8:
			return len == 0 ? new std::vector<char>() : new std::vector<char>(static_cast<char*>(arr), static_cast<char*>(arr) + len);
		case Char16:
			return len == 0 ? new std::vector<uint16_t>() : new std::vector<uint16_t>(static_cast<uint16_t*>(arr), static_cast<uint16_t*>(arr) + len);
		case Int8:
			return len == 0 ? new std::vector<int8_t>() : new std::vector<int8_t>(static_cast<int8_t*>(arr), static_cast<int8_t*>(arr) + len);
		case Int16:
			return len == 0 ? new std::vector<int16_t>() : new std::vector<int16_t>(static_cast<int16_t*>(arr), static_cast<int16_t*>(arr) + len);
		case Int32:
			return len == 0 ? new std::vector<int32_t>() : new std::vector<int32_t>(static_cast<int32_t*>(arr), static_cast<int32_t*>(arr) + len);
		case Int64:
			return len == 0 ? new std::vector<int64_t>() : new std::vector<int64_t>(static_cast<int64_t*>(arr), static_cast<int64_t*>(arr) + len);
		case UInt8:
			return len == 0 ? new std::vector<uint8_t>() : new std::vector<uint8_t>(static_cast<uint8_t*>(arr), static_cast<uint8_t*>(arr) + len);
		case UInt16:
			return len == 0 ? new std::vector<uint16_t>() : new std::vector<uint16_t>(static_cast<uint16_t*>(arr), static_cast<uint16_t*>(arr) + len);
		case UInt32:
			return len == 0 ? new std::vector<uint32_t>() : new std::vector<uint32_t>(static_cast<uint32_t*>(arr), static_cast<uint32_t*>(arr) + len);
		case UInt64:
			return len == 0 ? new std::vector<uint64_t>() : new std::vector<uint64_t>(static_cast<uint64_t*>(arr), static_cast<uint64_t*>(arr) + len);
		case Pointer:
			return len == 0 ? new std::vector<uintptr_t>() : new std::vector<uintptr_t>(static_cast<uintptr_t*>(arr), static_cast<uintptr_t*>(arr) + len);
		case Float:
			return len == 0 ? new std::vector<float>() : new std::vector<float>(static_cast<float*>(arr), static_cast<float*>(arr) + len);
		case Double:
			return len == 0 ? new std::vector<double>() : new std::vector<double>(static_cast<double*>(arr), static_cast<double*>(arr) + len);
		case String:
		{
			auto* vector = new std::vector<std::string>();
			if (len != 0) {
				vector->reserve(static_cast<size_t>(len));
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
		case Bool:
			return malloc(sizeof(std::vector<bool>));
		case Char8:
			return malloc(sizeof(std::vector<char>));
		case Char16:
			return malloc(sizeof(std::vector<uint16_t>));
		case Int8:
			return malloc(sizeof(std::vector<int8_t>));
		case Int16:
			return malloc(sizeof(std::vector<int16_t>));
		case Int32:
			return malloc(sizeof(std::vector<int32_t>));
		case Int64:
			return malloc(sizeof(std::vector<int64_t>));
		case UInt8:
			return malloc(sizeof(std::vector<uint8_t>));
		case UInt16:
			return malloc(sizeof(std::vector<uint16_t>));
		case UInt32:
			return malloc(sizeof(std::vector<uint32_t>));
		case UInt64:
			return malloc(sizeof(std::vector<uint64_t>));
		case Pointer:
			return malloc(sizeof(std::vector<uintptr_t>));
		case Float:
			return malloc(sizeof(std::vector<float>));
		case Double:
			return malloc(sizeof(std::vector<double>));
		case String:
			return malloc(sizeof(std::vector<std::string>));
		default:
			return nullptr;
	}
}

ptrdiff_t GetVectorSize(void* ptr, DataType type) {
	switch (type) {
		case Bool:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<bool>*>(ptr)->size());
		case Char8:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<char>*>(ptr)->size());
		case Char16:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint16_t>*>(ptr)->size());
		case Int8:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int8_t>*>(ptr)->size());
		case Int16:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int16_t>*>(ptr)->size());
		case Int32:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int32_t>*>(ptr)->size());
		case Int64:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<int64_t>*>(ptr)->size());
		case UInt8:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint8_t>*>(ptr)->size());
		case UInt16:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint16_t>*>(ptr)->size());
		case UInt32:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint32_t>*>(ptr)->size());
		case UInt64:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uint64_t>*>(ptr)->size());
		case Pointer:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<uintptr_t>*>(ptr)->size());
		case Float:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<float>*>(ptr)->size());
		case Double:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<double>*>(ptr)->size());
		case String:
			return static_cast<ptrdiff_t>(reinterpret_cast<std::vector<std::string>*>(ptr)->size());
		default:
			return -1; // Return -1 or some error code for invalid type
	}
}

void* GetVectorData(void* ptr, DataType type) {
	switch (type) {
		case Char8:
			return reinterpret_cast<std::vector<char>*>(ptr)->data();
		case Char16:
			return reinterpret_cast<std::vector<uint16_t>*>(ptr)->data();
		case Int8:
			return reinterpret_cast<std::vector<int8_t>*>(ptr)->data();
		case Int16:
			return reinterpret_cast<std::vector<int16_t>*>(ptr)->data();
		case Int32:
			return reinterpret_cast<std::vector<int32_t>*>(ptr)->data();
		case Int64:
			return reinterpret_cast<std::vector<int64_t>*>(ptr)->data();
		case UInt8:
			return reinterpret_cast<std::vector<uint8_t>*>(ptr)->data();
		case UInt16:
			return reinterpret_cast<std::vector<uint16_t>*>(ptr)->data();
		case UInt32:
			return reinterpret_cast<std::vector<uint32_t>*>(ptr)->data();
		case UInt64:
			return reinterpret_cast<std::vector<uint64_t>*>(ptr)->data();
		case Pointer:
			return reinterpret_cast<std::vector<uintptr_t>*>(ptr)->data();
		case Float:
			return reinterpret_cast<std::vector<float>*>(ptr)->data();
		case Double:
			return reinterpret_cast<std::vector<double>*>(ptr)->data();
		case Bool: {
			auto& vector = *reinterpret_cast<std::vector<bool>*>(ptr);

			bool* boolArray = new bool[vector.size()];

			// Manually copy values from std::vector<bool> to the bool array
			for (size_t i = 0; i < vector.size(); ++i) {
				boolArray[i] = vector[i];
			}

			return boolArray;
		}
		case String: {
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

void ConstructVector(void* ptr, void* arr, ptrdiff_t len, DataType type) {
	switch (type) {
		case Bool: {
			auto* vector = reinterpret_cast<std::vector<bool>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<bool>() : std::vector<bool>(static_cast<bool*>(arr), static_cast<bool*>(arr) + len));
			break;
		}
		case Char8: {
			auto* vector = reinterpret_cast<std::vector<char>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<char>() : std::vector<char>(static_cast<char*>(arr), static_cast<char*>(arr) + len));
			break;
		}
		case Char16: {
			auto* vector = reinterpret_cast<std::vector<uint16_t>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<uint16_t>() : std::vector<uint16_t>(static_cast<uint16_t*>(arr), static_cast<uint16_t*>(arr) + len));
			break;
		}
		case Int8: {
			auto* vector = reinterpret_cast<std::vector<int8_t>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<int8_t>() : std::vector<int8_t>(static_cast<int8_t*>(arr), static_cast<int8_t*>(arr) + len));
			break;
		}
		case Int16: {
			auto* vector = reinterpret_cast<std::vector<int16_t>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<int16_t>() : std::vector<int16_t>(static_cast<int16_t*>(arr), static_cast<int16_t*>(arr) + len));
			break;
		}
		case Int32: {
			auto* vector = reinterpret_cast<std::vector<int32_t>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<int32_t>() : std::vector<int32_t>(static_cast<int32_t*>(arr), static_cast<int32_t*>(arr) + len));
			break;
		}
		case Int64: {
			auto* vector = reinterpret_cast<std::vector<int64_t>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<int64_t>() : std::vector<int64_t>(static_cast<int64_t*>(arr), static_cast<int64_t*>(arr) + len));
			break;
		}
		case UInt8: {
			auto* vector = reinterpret_cast<std::vector<uint8_t>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<uint8_t>() : std::vector<uint8_t>(static_cast<uint8_t*>(arr), static_cast<uint8_t*>(arr) + len));
			break;
		}
		case UInt16: {
			auto* vector = reinterpret_cast<std::vector<uint16_t>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<uint16_t>() : std::vector<uint16_t>(static_cast<uint16_t*>(arr), static_cast<uint16_t*>(arr) + len));
			break;
		}
		case UInt32: {
			auto* vector = reinterpret_cast<std::vector<uint32_t>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<uint32_t>() : std::vector<uint32_t>(static_cast<uint32_t*>(arr), static_cast<uint32_t*>(arr) + len));
			break;
		}
		case UInt64: {
			auto* vector = reinterpret_cast<std::vector<uint64_t>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<uint64_t>() : std::vector<uint64_t>(static_cast<uint64_t*>(arr), static_cast<uint64_t*>(arr) + len));
			break;
		}
		case Pointer: {
			auto* vector = reinterpret_cast<std::vector<uintptr_t>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<uintptr_t>() : std::vector<uintptr_t>(static_cast<uintptr_t*>(arr), static_cast<uintptr_t*>(arr) + len));
			break;
		}
		case Float: {
			auto* vector = reinterpret_cast<std::vector<float>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<float>() : std::vector<float>(static_cast<float*>(arr), static_cast<float*>(arr) + len));
			break;
		}
		case Double: {
			auto* vector = reinterpret_cast<std::vector<double>*>(ptr);
			std::construct_at(vector, len == 0 ? std::vector<double>() : std::vector<double>(static_cast<double*>(arr), static_cast<double*>(arr) + len));
			break;
		}
		case String: {
			auto* vector = reinterpret_cast<std::vector<std::string>*>(ptr);
			std::construct_at(vector, std::vector<std::string>());
			size_t N = static_cast<size_t>(len);
			vector->reserve(N);
			for (size_t i = 0; i < N; ++i) {
				const auto& str = static_cast<GoString*>(arr)[i];
				vector->emplace_back(str.p, static_cast<size_t>(str.n));
			}
			break;
		}
		default:
			break;
	}
}

void AssignVector(void* ptr, void* arr, ptrdiff_t len, DataType type) {
	switch (type) {
		case Bool: {
			auto* vector = reinterpret_cast<std::vector<bool>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<bool*>(arr), static_cast<bool*>(arr) + len);
			break;
		}
		case Char8: {
			auto* vector = reinterpret_cast<std::vector<char>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<char*>(arr), static_cast<char*>(arr) + len);
			break;
		}
		case Char16: {
			auto* vector = reinterpret_cast<std::vector<uint16_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<uint16_t*>(arr), static_cast<uint16_t*>(arr) + len);
			break;
		}
		case Int8: {
			auto* vector = reinterpret_cast<std::vector<int8_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<int8_t*>(arr), static_cast<int8_t*>(arr) + len);
			break;
		}
		case Int16: {
			auto* vector = reinterpret_cast<std::vector<int16_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<int16_t*>(arr), static_cast<int16_t*>(arr) + len);
			break;
		}
		case Int32: {
			auto* vector = reinterpret_cast<std::vector<int32_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<int32_t*>(arr), static_cast<int32_t*>(arr) + len);
			break;
		}
		case Int64: {
			auto* vector = reinterpret_cast<std::vector<int64_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<int64_t*>(arr), static_cast<int64_t*>(arr) + len);
			break;
		}
		case UInt8: {
			auto* vector = reinterpret_cast<std::vector<uint8_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<uint8_t*>(arr), static_cast<uint8_t*>(arr) + len);
			break;
		}
		case UInt16: {
			auto* vector = reinterpret_cast<std::vector<uint16_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<uint16_t*>(arr), static_cast<uint16_t*>(arr) + len);
			break;
		}
		case UInt32: {
			auto* vector = reinterpret_cast<std::vector<uint32_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<uint32_t*>(arr), static_cast<uint32_t*>(arr) + len);
			break;
		}
		case UInt64: {
			auto* vector = reinterpret_cast<std::vector<uint64_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<uint64_t*>(arr), static_cast<uint64_t*>(arr) + len);
			break;
		}
		case Pointer: {
			auto* vector = reinterpret_cast<std::vector<uintptr_t>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<uintptr_t*>(arr), static_cast<uintptr_t*>(arr) + len);
			break;
		}
		case Float: {
			auto* vector = reinterpret_cast<std::vector<float>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<float*>(arr), static_cast<float*>(arr) + len);
			break;
		}
		case Double: {
			auto* vector = reinterpret_cast<std::vector<double>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else
				vector->assign(static_cast<double*>(arr), static_cast<double*>(arr) + len);
			break;
		}
		case String: {
			auto* vector = reinterpret_cast<std::vector<std::string>*>(ptr);
			if (arr == nullptr || len == 0)
				vector->clear();
			else {
				size_t N = static_cast<size_t>(len);
				vector->resize(N);
				for (size_t i = 0; i < N; ++i) {
					const auto& str = static_cast<GoString*>(arr)[i];
					(*vector)[i].assign(str.p, static_cast<size_t>(str.n));
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
		case Bool:
			delete reinterpret_cast<std::vector<bool>*>(ptr);
			break;
		case Char8:
			delete reinterpret_cast<std::vector<char>*>(ptr);
			break;
		case Char16:
			delete reinterpret_cast<std::vector<uint16_t>*>(ptr);
			break;
		case Int8:
			delete reinterpret_cast<std::vector<int8_t>*>(ptr);
			break;
		case Int16:
			delete reinterpret_cast<std::vector<int16_t>*>(ptr);
			break;
		case Int32:
			delete reinterpret_cast<std::vector<int32_t>*>(ptr);
			break;
		case Int64:
			delete reinterpret_cast<std::vector<int64_t>*>(ptr);
			break;
		case UInt8:
			delete reinterpret_cast<std::vector<uint8_t>*>(ptr);
			break;
		case UInt16:
			delete reinterpret_cast<std::vector<uint16_t>*>(ptr);
			break;
		case UInt32:
			delete reinterpret_cast<std::vector<uint32_t>*>(ptr);
			break;
		case UInt64:
			delete reinterpret_cast<std::vector<uint64_t>*>(ptr);
			break;
		case Pointer:
			delete reinterpret_cast<std::vector<uintptr_t>*>(ptr);
			break;
		case Float:
			delete reinterpret_cast<std::vector<float>*>(ptr);
			break;
		case Double:
			delete reinterpret_cast<std::vector<double>*>(ptr);
			break;
		case String:
			delete reinterpret_cast<std::vector<std::string>*>(ptr);
			break;
		default:
			// Invalid type, do nothing or handle error if needed
			break;
	}
}

void FreeVector(void* ptr, DataType type) {
	switch (type) {
		case Bool:
			reinterpret_cast<std::vector<bool>*>(ptr)->~vector();
			break;
		case Char8:
			reinterpret_cast<std::vector<char>*>(ptr)->~vector();
			break;
		case Char16:
			reinterpret_cast<std::vector<uint16_t>*>(ptr)->~vector();
			break;
		case Int8:
			reinterpret_cast<std::vector<int8_t>*>(ptr)->~vector();
			break;
		case Int16:
			reinterpret_cast<std::vector<int16_t>*>(ptr)->~vector();
			break;
		case Int32:
			reinterpret_cast<std::vector<int32_t>*>(ptr)->~vector();
			break;
		case Int64:
			reinterpret_cast<std::vector<int64_t>*>(ptr)->~vector();
			break;
		case UInt8:
			reinterpret_cast<std::vector<uint8_t>*>(ptr)->~vector();
			break;
		case UInt16:
			reinterpret_cast<std::vector<uint16_t>*>(ptr)->~vector();
			break;
		case UInt32:
			reinterpret_cast<std::vector<uint32_t>*>(ptr)->~vector();
			break;
		case UInt64:
			reinterpret_cast<std::vector<uint64_t>*>(ptr)->~vector();
			break;
		case Pointer:
			reinterpret_cast<std::vector<uintptr_t>*>(ptr)->~vector();
			break;
		case Float:
			reinterpret_cast<std::vector<float>*>(ptr)->~vector();
			break;
		case Double:
			reinterpret_cast<std::vector<double>*>(ptr)->~vector();
			break;
		case String:
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

const std::array<void*, 35> GoLanguageModule::_pluginApi = {
		reinterpret_cast<void*>(&::GetMethodPtr),
		reinterpret_cast<void*>(&::GetMethodPtr2),
		reinterpret_cast<void*>(&::GetBaseDir),
		reinterpret_cast<void*>(&::IsModuleLoaded),
		reinterpret_cast<void*>(&::IsPluginLoaded),
		reinterpret_cast<void*>(&::GetPluginId),
		reinterpret_cast<void*>(&::GetPluginName),
		reinterpret_cast<void*>(&::GetPluginFullName),
		reinterpret_cast<void*>(&::GetPluginDescription),
		reinterpret_cast<void*>(&::GetPluginVersion),
		reinterpret_cast<void*>(&::GetPluginAuthor),
		reinterpret_cast<void*>(&::GetPluginWebsite),
		reinterpret_cast<void*>(&::GetPluginBaseDir),
		reinterpret_cast<void*>(&::GetPluginDependencies),
		reinterpret_cast<void*>(&::GetPluginDependenciesSize),
		reinterpret_cast<void*>(&::FindPluginResource),
		reinterpret_cast<void*>(&::DeleteCStr),
		reinterpret_cast<void*>(&::AllocateString),
		reinterpret_cast<void*>(&::CreateString),
		reinterpret_cast<void*>(&::GetStringData),
		reinterpret_cast<void*>(&::GetStringLength),
		reinterpret_cast<void*>(&::ConstructString),
		reinterpret_cast<void*>(&::AssignString),
		reinterpret_cast<void*>(&::FreeString),
		reinterpret_cast<void*>(&::DeleteString),
		reinterpret_cast<void*>(&::CreateVector),
		reinterpret_cast<void*>(&::AllocateVector),
		reinterpret_cast<void*>(&::GetVectorSize),
		reinterpret_cast<void*>(&::GetVectorData),
		reinterpret_cast<void*>(&::ConstructVector),
		reinterpret_cast<void*>(&::AssignVector),
		reinterpret_cast<void*>(&::DeleteVector),
		reinterpret_cast<void*>(&::FreeVector),
		reinterpret_cast<void*>(&::DeleteVectorDataBool),
		reinterpret_cast<void*>(&::DeleteVectorDataCStr)
};

ILanguageModule* GetLanguageModule() {
	return &golm::g_golm;
}
