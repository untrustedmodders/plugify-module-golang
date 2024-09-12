#include "module.h"

#include <plugify/module.h>
#include <plugify/plugin.h>
#include <plugify/log.h>
#include <plugify/math.h>
#include <plugify/string.h>
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

GoSlice* CreateGoSliceString(const std::vector<plg::string>& source, ArgumentList& args, StringHolder& holder) {
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

GoString* CreateGoString(const plg::string& source, ArgumentList& args) {
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
void CopyGoSliceToVector(const GoSlice& source, std::vector<plg::string>& dest) {
	dest.resize(static_cast<size_t>(source.len));
	for (size_t i = 0; i < dest.size(); ++i) {
		const auto& str = reinterpret_cast<GoString*>(source.data)[i];
		dest[i].assign(str.p, static_cast<size_t>(str.n));
	}
}

void CopyGoStringToString(const GoString& source, plg::string& dest) {
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
void CopyGoSliceToVectorReturn(const GoSlice& source, std::vector<plg::string>& dest) {
	if (source.data == nullptr || source.len == 0)
		std::construct_at(&dest, std::vector<plg::string>());
	else {
		std::construct_at(&dest, std::vector<plg::string>(static_cast<size_t>(source.len)));
		for (size_t i = 0; i < dest.size(); ++i) {
			const auto& str = reinterpret_cast<GoString*>(source.data)[i];
			dest[i].assign(str.p, static_cast<size_t>(str.n));
		}
	}
}

void CopyGoStringToStringReturn(const GoString& source, plg::string& dest) {
	if (source.p == nullptr || source.n == 0)
		std::construct_at(&dest, plg::string());
	else
		std::construct_at(&dest, plg::string(source.p, static_cast<size_t>(source.n)));
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

	DetectLeaks();

	_rt.reset();
	_provider.reset();
}

bool GoLanguageModule::IsDebugBuild() {
	return GOLM_IS_DEBUG;
}

void GoLanguageModule::OnMethodExport(PluginRef plugin) {
	for (const auto& [method, addr] : plugin.GetMethods()) {
		_nativesMap.try_emplace(std::format("{}.{}", plugin.GetName(), method.GetName()), addr);
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

	std::span<const MethodRef> exportedMethods = plugin.GetDescriptor().GetExportedMethods();
	std::vector<MethodData> methods;
	methods.reserve(exportedMethods.size());

	for (const auto& method : exportedMethods) {
		if (auto func = assembly->GetFunctionByName(method.GetFunctionName())) {
			if (IsMethodPrimitive(method)) {
				methods.emplace_back(method, func);
			} else {
				auto function = std::make_unique<Function>(_rt);
				func = function->GetJitFunc(method, &InternalCall, func);
				if (!func) {
					funcErrors.emplace_back(method.GetName());
					continue;
				}
				_functions.emplace_back(std::move(function));
				methods.emplace_back(method, func);
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
	std::span<const PropertyRef> paramProps = method.GetParamTypes();

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
					dcArgPointer(vm, CreateGoString(*p->GetArgument<plg::string*>(i), args));
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
					dcArgPointer(vm, CreateGoSliceString(*p->GetArgument<std::vector<plg::string>*>(i), args, stringHolder));
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
					dcArgAggr(vm, CreateDcAggr<GoString>(aggrs), CreateGoString(*p->GetArgument<plg::string*>(i), args));
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
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSliceString(*p->GetArgument<std::vector<plg::string>*>(i), args, stringHolder));
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
			auto* dest = p->GetArgument<plg::string*>(0);
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
			auto* dest = p->GetArgument<std::vector<plg::string>*>(0);
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
							CopyGoStringToString(*reinterpret_cast<GoString*>(&args[k++]), *p->GetArgument<plg::string*>(i));
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
							CopyGoSliceToVector(args[k++], *p->GetArgument<std::vector<plg::string>*>(i));
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

using type_index = uint32_t;
inline type_index type_id_seq = 0;
template<typename T> inline const type_index type_id = type_id_seq++;

std::map<type_index, int32_t> g_numberOfMalloc = { };
std::map<type_index, int32_t> g_numberOfAllocs = { };

std::string_view GetTypeName(type_index type) {
	static std::map<type_index, std::string_view> typeNameMap = {
			{type_id<plg::string>, "String"},
			{type_id<std::vector<bool>>, "VectorBool"},
			{type_id<std::vector<char>>, "VectorChar8"},
			{type_id<std::vector<char16_t>>, "VectorChar16"},
			{type_id<std::vector<int8_t>>, "VectorInt8"},
			{type_id<std::vector<int16_t>>, "VectorInt16"},
			{type_id<std::vector<int32_t>>, "VectorInt32"},
			{type_id<std::vector<int64_t>>, "VectorInt64"},
			{type_id<std::vector<uint8_t>>, "VectorUInt8"},
			{type_id<std::vector<uint16_t>>, "VectorUInt16"},
			{type_id<std::vector<uint32_t>>, "VectorUInt32"},
			{type_id<std::vector<uint64_t>>, "VectorUInt64"},
			{type_id<std::vector<uintptr_t>>, "VectorUIntPtr"},
			{type_id<std::vector<float>>, "VectorFloat"},
			{type_id<std::vector<double>>, "VectorDouble"},
			{type_id<std::vector<plg::string>>, "VectorString"},
			{type_id<bool*>, "BoolArray"},
			{type_id<char*>, "CString"},
			{type_id<char**>, "CStringArray"}
	};
	auto it = typeNameMap.find(type);
	if (it != typeNameMap.end()) {
		return std::get<std::string_view>(*it);
	}
	return "unknown";
}

void GoLanguageModule::DetectLeaks() {
	for (const auto& [type, count] : g_numberOfMalloc) {
		if (count > 0) {
			g_golm._provider->Log(std::format(LOG_PREFIX "Memory leaks detected: {} allocations. Related to {}!", count, GetTypeName(type)), Severity::Error);
		}
	}

	for (const auto& [type, count] : g_numberOfAllocs) {
		if (count > 0) {
			g_golm._provider->Log(std::format(LOG_PREFIX "Memory leaks detected: {} allocations. Related to {}!", count, GetTypeName(type)), Severity::Error);
		}
	}

	g_numberOfMalloc.clear();
	g_numberOfAllocs.clear();
}

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
	++g_numberOfAllocs[type_id<char*>];
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
	++g_numberOfAllocs[type_id<char*>];
	return dest;
}

const char** GetPluginDependencies(PluginRef plugin) {
	std::span<const PluginReferenceDescriptorRef> dependencies = plugin.GetDescriptor().GetDependencies();
	auto* deps = new const char*[dependencies.size()];
	for (size_t i = 0; i < dependencies.size(); ++i) {
		deps[i] = dependencies[i].GetName().data();
	}
	++g_numberOfAllocs[type_id<char**>];
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
		++g_numberOfAllocs[type_id<char*>];
		return dest;
	}
	return "";
}
void DeleteCStr(const char* path) {
	delete path;
	--g_numberOfAllocs[type_id<char*>];
	assert(g_numberOfAllocs[type_id<char*>] != -1);
}

void* AllocateString() {
	auto str = static_cast<plg::string*>(std::malloc(sizeof(plg::string)));
	++g_numberOfMalloc[type_id<plg::string>];
	return str;
}
void* CreateString(GoString source) {
	auto str = source.n == 0 || source.p == nullptr ? new plg::string() : new plg::string(source.p, static_cast<size_t>(source.n));
	++g_numberOfAllocs[type_id<plg::string>];
	return str;
}
const char* GetStringData(plg::string* string) {
	return string->c_str();
}
ptrdiff_t GetStringLength(plg::string* string) {
	return static_cast<ptrdiff_t>(string->length());
}
void ConstructString(plg::string* string, GoString source) {
	if (source.n == 0 || source.p == nullptr)
		std::construct_at(string, plg::string());
	else
		std::construct_at(string, plg::string(source.p, static_cast<size_t>(source.n)));
}
void AssignString(plg::string* string, GoString source) {
	if (source.p == nullptr || source.n == 0)
		string->clear();
	else
		string->assign(source.p, static_cast<size_t>(source.n));
}
void FreeString(plg::string* string) {
	string->~basic_string();
	std::free(string);
	--g_numberOfMalloc[type_id<plg::string>];
	assert(g_numberOfMalloc[type_id<plg::string>] != -1);
}
void DeleteString(plg::string* string) {
	delete string;
	--g_numberOfAllocs[type_id<plg::string>];
	assert(g_numberOfAllocs[type_id<plg::string>] != -1);
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

namespace {
	template<typename T>
	std::vector<T>* CreateVector(T* arr, ptrdiff_t len) requires(!std::is_same_v<T, GoString>) {
		auto vector = len == 0 ? new std::vector<T>() : new std::vector<T>(arr, arr + len);
		assert(vector);
		++g_numberOfAllocs[type_id<std::vector<T>>];
		return vector;
	}

	template<typename T>
	std::vector<plg::string>* CreateVector(T* arr, ptrdiff_t len) requires(std::is_same_v<T, GoString>) {
		auto vector = new std::vector<plg::string>();
		if (len != 0) {
			vector->reserve(static_cast<size_t>(len));
			for (ptrdiff_t i = 0; i < len; ++i) {
				const auto& str = arr[i];
				vector->emplace_back(str.p, str.n);
			}
		}
		assert(vector);
		++g_numberOfAllocs[type_id<std::vector<plg::string>>];
		return vector;
	}

	template<typename T>
	std::vector<T>* AllocateVector() requires(!std::is_same_v<T, GoString>) {
		auto vector = static_cast<std::vector<T>*>(std::malloc(sizeof(std::vector<T>)));
		assert(vector);
		++g_numberOfMalloc[type_id<std::vector<T>>];
		return vector;
	}

	template<typename T>
	std::vector<plg::string>* AllocateVector() requires(std::is_same_v<T, GoString>) {
		auto vector = static_cast<std::vector<plg::string>*>(std::malloc(sizeof(std::vector<plg::string>)));
		assert(vector);
		++g_numberOfMalloc[type_id<std::vector<plg::string>>];
		return vector;
	}


	template<typename T>
	void DeleteVector(std::vector<T>* vector) {
		delete vector;
		--g_numberOfAllocs[type_id<std::vector<T>>];
		assert(g_numberOfAllocs[type_id<std::vector<T>>] != -1);
	}

	template<typename T>
	void FreeVector(std::vector<T>* vector) {
		vector->~vector();
		std::free(vector);
		--g_numberOfMalloc[type_id<std::vector<T>>];
		assert(g_numberOfMalloc[type_id<std::vector<T>>] != -1);
	}

	template<typename T>
	ptrdiff_t GetVectorSize(std::vector<T>* vector) {
		return static_cast<ptrdiff_t>(vector->size());
	}

	template<typename T>
	void AssignVector(std::vector<T>* vector, T* arr, ptrdiff_t len) requires(!std::is_same_v<T, GoString>) {
		if (arr == nullptr || len == 0)
			vector->clear();
		else
			vector->assign(arr, arr + len);
	}

	template<typename T>
	void AssignVector(std::vector<plg::string>* vector, T* arr, ptrdiff_t len) requires(std::is_same_v<T, GoString>) {
		if (arr == nullptr || len == 0)
			vector->clear();
		else {
			size_t N = static_cast<size_t>(len);
			vector->resize(N);
			for (size_t i = 0; i < N; ++i) {
				const auto& str = arr[i];
				(*vector)[i].assign(str.p, static_cast<size_t>(str.n));
			}
		}
	}

	template<typename T>
	T* GetVectorData(std::vector<T>* vector) requires(!std::is_same_v<T, GoString> and !std::is_same_v<T, bool>) {
		return vector->data();
	}

	template<typename T>
	bool* GetVectorData(std::vector<T>* vector) requires(std::is_same_v<T, bool>) {
		bool* boolArray = new bool[vector->size()];

		// Manually copy values from std::vector<bool> to the bool array
		for (size_t i = 0; i < vector->size(); ++i) {
			boolArray[i] = (*vector)[i];
		}
		
		++g_numberOfAllocs[type_id<bool*>];

		return boolArray;
	}

	template<typename T>
	char** GetVectorData(std::vector<plg::string>* vector) requires(std::is_same_v<T, GoString>) {
		char** strArray = new char*[vector->size()];

		// Manually copy values from std::vector<std::string> to the char* array
		for (size_t i = 0; i < vector->size(); ++i) {
			strArray[i] = (*vector)[i].data();
		}
		
		++g_numberOfAllocs[type_id<char**>];

		return strArray;
	}

	template<typename T>
	void ConstructVector(std::vector<T>* vector, T* arr, ptrdiff_t len) requires(!std::is_same_v<T, GoString>) {
		std::construct_at(vector, len == 0 ? std::vector<T>() : std::vector<T>(arr, arr + len));
	}

	template<typename T>
	void ConstructVector(std::vector<plg::string>* vector, T* arr, ptrdiff_t len) requires(std::is_same_v<T, GoString>) {
		std::construct_at(vector, std::vector<plg::string>());
		size_t N = static_cast<size_t>(len);
		vector->reserve(N);
		for (size_t i = 0; i < N; ++i) {
			const auto& str = arr[i];
			vector->emplace_back(str.p, static_cast<size_t>(str.n));
		}
	}
}

void* CreateVector(void* arr, ptrdiff_t len, DataType type) {
	switch (type) {
		case Bool:
			return CreateVector(static_cast<bool*>(arr), len);
		case Char8:
			return CreateVector(static_cast<char*>(arr), len);
		case Char16:
			return CreateVector(static_cast<char16_t*>(arr), len);
		case Int8:
			return CreateVector(static_cast<int8_t*>(arr), len);
		case Int16:
			return CreateVector(static_cast<int16_t*>(arr), len);
		case Int32:
			return CreateVector(static_cast<int32_t*>(arr), len);
		case Int64:
			return CreateVector(static_cast<int64_t*>(arr), len);
		case UInt8:
			return CreateVector(static_cast<uint8_t*>(arr), len);
		case UInt16:
			return CreateVector(static_cast<uint16_t*>(arr), len);
		case UInt32:
			return CreateVector(static_cast<uint32_t*>(arr), len);
		case UInt64:
			return CreateVector(static_cast<uint64_t*>(arr), len);
		case Pointer:
			return CreateVector(static_cast<uintptr_t*>(arr), len);
		case Float:
			return CreateVector(static_cast<float*>(arr), len);
		case Double:
			return CreateVector(static_cast<double*>(arr), len);
		case String:
			return CreateVector(static_cast<GoString*>(arr), len);
		default:
			return nullptr;
	}
}

void* AllocateVector(DataType type) {
	switch (type) {
		case Bool:
			return AllocateVector<bool>();
		case Char8:
			return AllocateVector<char>();
		case Char16:
			return AllocateVector<char16_t>();
		case Int8:
			return AllocateVector<int8_t>();
		case Int16:
			return AllocateVector<int16_t>();
		case Int32:
			return AllocateVector<int32_t>();
		case Int64:
			return AllocateVector<int64_t>();
		case UInt8:
			return AllocateVector<uint8_t>();
		case UInt16:
			return AllocateVector<uint16_t>();
		case UInt32:
			return AllocateVector<uint32_t>();
		case UInt64:
			return AllocateVector<uint64_t>();
		case Pointer:
			return AllocateVector<uintptr_t>();
		case Float:
			return AllocateVector<float>();
		case Double:
			return AllocateVector<double>();
		case String:
			return AllocateVector<plg::string>();
		default:
			return nullptr;
	}
}

ptrdiff_t GetVectorSize(void* ptr, DataType type) {
	switch (type) {
		case Char8:
			return GetVectorSize(reinterpret_cast<std::vector<char>*>(ptr));
		case Char16:
			return GetVectorSize(reinterpret_cast<std::vector<char16_t>*>(ptr));
		case Int8:
			return GetVectorSize(reinterpret_cast<std::vector<int8_t>*>(ptr));
		case Int16:
			return GetVectorSize(reinterpret_cast<std::vector<int16_t>*>(ptr));
		case Int32:
			return GetVectorSize(reinterpret_cast<std::vector<int32_t>*>(ptr));
		case Int64:
			return GetVectorSize(reinterpret_cast<std::vector<int64_t>*>(ptr));
		case UInt8:
			return GetVectorSize(reinterpret_cast<std::vector<uint8_t>*>(ptr));
		case UInt16:
			return GetVectorSize(reinterpret_cast<std::vector<uint16_t>*>(ptr));
		case UInt32:
			return GetVectorSize(reinterpret_cast<std::vector<uint32_t>*>(ptr));
		case UInt64:
			return GetVectorSize(reinterpret_cast<std::vector<uint64_t>*>(ptr));
		case Pointer:
			return GetVectorSize(reinterpret_cast<std::vector<uintptr_t>*>(ptr));
		case Float:
			return GetVectorSize(reinterpret_cast<std::vector<float>*>(ptr));
		case Double:
			return GetVectorSize(reinterpret_cast<std::vector<double>*>(ptr));
		case Bool:
			return GetVectorSize(reinterpret_cast<std::vector<bool>*>(ptr));
		case String:
			return GetVectorSize(reinterpret_cast<std::vector<plg::string>*>(ptr));
		default:
			return -1; // Return -1 or some error code for invalid type
	}
}

void* GetVectorData(void* ptr, DataType type) {
	switch (type) {
		case Bool:
			return GetVectorData<bool>(reinterpret_cast<std::vector<bool>*>(ptr));
		case Char8:
			return GetVectorData<>(reinterpret_cast<std::vector<char>*>(ptr));
		case Char16:
			return GetVectorData<>(reinterpret_cast<std::vector<char16_t>*>(ptr));
		case Int8:
			return GetVectorData<>(reinterpret_cast<std::vector<int8_t>*>(ptr));
		case Int16:
			return GetVectorData<>(reinterpret_cast<std::vector<int16_t>*>(ptr));
		case Int32:
			return GetVectorData<>(reinterpret_cast<std::vector<int32_t>*>(ptr));
		case Int64:
			return GetVectorData<>(reinterpret_cast<std::vector<int64_t>*>(ptr));
		case UInt8:
			return GetVectorData<>(reinterpret_cast<std::vector<uint8_t>*>(ptr));
		case UInt16:
			return GetVectorData<>(reinterpret_cast<std::vector<uint16_t>*>(ptr));
		case UInt32:
			return GetVectorData<>(reinterpret_cast<std::vector<uint32_t>*>(ptr));
		case UInt64:
			return GetVectorData<>(reinterpret_cast<std::vector<uint64_t>*>(ptr));
		case Pointer:
			return GetVectorData<>(reinterpret_cast<std::vector<uintptr_t>*>(ptr));
		case Float:
			return GetVectorData<>(reinterpret_cast<std::vector<float>*>(ptr));
		case Double:
			return GetVectorData<>(reinterpret_cast<std::vector<double>*>(ptr));
		case String:
			return GetVectorData<GoString>(reinterpret_cast<std::vector<plg::string>*>(ptr));
		default:
			return nullptr; // Return nullptr for invalid type
	}
}

void ConstructVector(void* ptr, void* arr, ptrdiff_t len, DataType type) {
	switch (type) {
		case Bool:
			ConstructVector(reinterpret_cast<std::vector<bool>*>(ptr), static_cast<bool*>(arr), len);
			break;
		case Char8:
			ConstructVector(reinterpret_cast<std::vector<char>*>(ptr), static_cast<char*>(arr), len);
			break;
		case Char16:
			ConstructVector(reinterpret_cast<std::vector<char16_t>*>(ptr), static_cast<char16_t*>(arr), len);
			break;
		case Int8:
			ConstructVector(reinterpret_cast<std::vector<int8_t>*>(ptr), static_cast<int8_t*>(arr), len);
			break;
		case Int16:
			ConstructVector(reinterpret_cast<std::vector<int16_t>*>(ptr), static_cast<int16_t*>(arr), len);
			break;
		case Int32:
			ConstructVector(reinterpret_cast<std::vector<int32_t>*>(ptr), static_cast<int32_t*>(arr), len);
			break;
		case Int64:
			ConstructVector(reinterpret_cast<std::vector<int64_t>*>(ptr), static_cast<int64_t*>(arr), len);
			break;
		case UInt8:
			ConstructVector(reinterpret_cast<std::vector<uint8_t>*>(ptr), static_cast<uint8_t*>(arr), len);
			break;
		case UInt16:
			ConstructVector(reinterpret_cast<std::vector<uint16_t>*>(ptr), static_cast<uint16_t*>(arr), len);
			break;
		case UInt32:
			ConstructVector(reinterpret_cast<std::vector<uint32_t>*>(ptr), static_cast<uint32_t*>(arr), len);
			break;
		case UInt64:
			ConstructVector(reinterpret_cast<std::vector<uint64_t>*>(ptr), static_cast<uint64_t*>(arr), len);
			break;
		case Pointer:
			ConstructVector(reinterpret_cast<std::vector<uintptr_t>*>(ptr), static_cast<uintptr_t*>(arr), len);
			break;
		case Float:
			ConstructVector(reinterpret_cast<std::vector<float>*>(ptr), static_cast<float*>(arr), len);
			break;
		case Double:
			ConstructVector(reinterpret_cast<std::vector<double>*>(ptr), static_cast<double*>(arr), len);
			break;
		case String:
			ConstructVector(reinterpret_cast<std::vector<plg::string>*>(ptr), static_cast<GoString*>(arr), len);
			break;
		default:
			break;
	}
}

void AssignVector(void* ptr, void* arr, ptrdiff_t len, DataType type) {
	switch (type) {
		case Bool:
			AssignVector(reinterpret_cast<std::vector<bool>*>(ptr), static_cast<bool*>(arr), len);
			break;
		case Char8:
			AssignVector(reinterpret_cast<std::vector<char>*>(ptr), static_cast<char*>(arr), len);
			break;
		case Char16:
			AssignVector(reinterpret_cast<std::vector<char16_t>*>(ptr), static_cast<char16_t*>(arr), len);
			break;
		case Int8:
			AssignVector(reinterpret_cast<std::vector<int8_t>*>(ptr), static_cast<int8_t*>(arr), len);
			break;
		case Int16:
			AssignVector(reinterpret_cast<std::vector<int16_t>*>(ptr), static_cast<int16_t*>(arr), len);
			break;
		case Int32:
			AssignVector(reinterpret_cast<std::vector<int32_t>*>(ptr), static_cast<int32_t*>(arr), len);
			break;
		case Int64:
			AssignVector(reinterpret_cast<std::vector<int64_t>*>(ptr), static_cast<int64_t*>(arr), len);
			break;
		case UInt8:
			AssignVector(reinterpret_cast<std::vector<uint8_t>*>(ptr), static_cast<uint8_t*>(arr), len);
			break;
		case UInt16:
			AssignVector(reinterpret_cast<std::vector<uint16_t>*>(ptr), static_cast<uint16_t*>(arr), len);
			break;
		case UInt32:
			AssignVector(reinterpret_cast<std::vector<uint32_t>*>(ptr), static_cast<uint32_t*>(arr), len);
			break;
		case UInt64:
			AssignVector(reinterpret_cast<std::vector<uint64_t>*>(ptr), static_cast<uint64_t*>(arr), len);
			break;
		case Pointer:
			AssignVector(reinterpret_cast<std::vector<uintptr_t>*>(ptr), static_cast<uintptr_t*>(arr), len);
			break;
		case Float:
			AssignVector(reinterpret_cast<std::vector<float>*>(ptr), static_cast<float*>(arr), len);
			break;
		case Double:
			AssignVector(reinterpret_cast<std::vector<double>*>(ptr), static_cast<double*>(arr), len);
			break;
		case String:
			AssignVector(reinterpret_cast<std::vector<plg::string>*>(ptr), static_cast<GoString*>(arr), len);
			break;
		default:
			break;
	}
}

void DeleteVector(void* ptr, DataType type) {
	switch (type) {
		case Bool:
			DeleteVector(reinterpret_cast<std::vector<bool>*>(ptr));
			break;
		case Char8:
			DeleteVector(reinterpret_cast<std::vector<char>*>(ptr));
			break;
		case Char16:
			DeleteVector(reinterpret_cast<std::vector<char16_t>*>(ptr));
			break;
		case Int8:
			DeleteVector(reinterpret_cast<std::vector<int8_t>*>(ptr));
			break;
		case Int16:
			DeleteVector(reinterpret_cast<std::vector<int16_t>*>(ptr));
			break;
		case Int32:
			DeleteVector(reinterpret_cast<std::vector<int32_t>*>(ptr));
			break;
		case Int64:
			DeleteVector(reinterpret_cast<std::vector<int64_t>*>(ptr));
			break;
		case UInt8:
			DeleteVector(reinterpret_cast<std::vector<uint8_t>*>(ptr));
			break;
		case UInt16:
			DeleteVector(reinterpret_cast<std::vector<uint16_t>*>(ptr));
			break;
		case UInt32:
			DeleteVector(reinterpret_cast<std::vector<uint32_t>*>(ptr));
			break;
		case UInt64:
			DeleteVector(reinterpret_cast<std::vector<uint64_t>*>(ptr));
			break;
		case Pointer:
			DeleteVector(reinterpret_cast<std::vector<uintptr_t>*>(ptr));
			break;
		case Float:
			DeleteVector(reinterpret_cast<std::vector<float>*>(ptr));
			break;
		case Double:
			DeleteVector(reinterpret_cast<std::vector<double>*>(ptr));
			break;
		case String:
			DeleteVector(reinterpret_cast<std::vector<plg::string>*>(ptr));
			break;
		default:
			// Invalid type, do nothing or handle error if needed
			break;
	}
}

void FreeVector(void* ptr, DataType type) {
	switch (type) {
		case Bool:
			FreeVector(reinterpret_cast<std::vector<bool>*>(ptr));
			break;
		case Char8:
			FreeVector(reinterpret_cast<std::vector<char>*>(ptr));
			break;
		case Char16:
			FreeVector(reinterpret_cast<std::vector<char16_t>*>(ptr));
			break;
		case Int8:
			FreeVector(reinterpret_cast<std::vector<int8_t>*>(ptr));
			break;
		case Int16:
			FreeVector(reinterpret_cast<std::vector<int16_t>*>(ptr));
			break;
		case Int32:
			FreeVector(reinterpret_cast<std::vector<int32_t>*>(ptr));
			break;
		case Int64:
			FreeVector(reinterpret_cast<std::vector<int64_t>*>(ptr));
			break;
		case UInt8:
			FreeVector(reinterpret_cast<std::vector<uint8_t>*>(ptr));
			break;
		case UInt16:
			FreeVector(reinterpret_cast<std::vector<uint16_t>*>(ptr));
			break;
		case UInt32:
			FreeVector(reinterpret_cast<std::vector<uint32_t>*>(ptr));
			break;
		case UInt64:
			FreeVector(reinterpret_cast<std::vector<uint64_t>*>(ptr));
			break;
		case Pointer:
			FreeVector(reinterpret_cast<std::vector<uintptr_t>*>(ptr));
			break;
		case Float:
			FreeVector(reinterpret_cast<std::vector<float>*>(ptr));
			break;
		case Double:
			FreeVector(reinterpret_cast<std::vector<double>*>(ptr));
			break;
		case String:
			FreeVector(reinterpret_cast<std::vector<plg::string>*>(ptr));
			break;
		default:
			// Invalid type, do nothing or handle error if needed
			return;
	}
}

void DeleteVectorDataBool(void* ptr) {
	delete[] reinterpret_cast<bool*>(ptr);
	--g_numberOfAllocs[type_id<bool*>];
	assert(g_numberOfAllocs[type_id<bool*>] != -1);
}

void DeleteVectorDataCStr(void* ptr) {
	delete[] reinterpret_cast<char**>(ptr);
	--g_numberOfAllocs[type_id<char**>];
	assert(g_numberOfAllocs[type_id<char**>] != -1);
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
