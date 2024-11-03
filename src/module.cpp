#include "module.hpp"

#include <plugify/module.hpp>
#include <plugify/plugin.hpp>
#include <plugify/log.hpp>
#include <plugify/math.hpp>
#include <plugify/string.hpp>
#include <plugify/plugin_descriptor.hpp>
#include <plugify/plugin_reference_descriptor.hpp>
#include <plugify/plugify_provider.hpp>

#include <dyncall/dyncall.h>

#if GOLM_PLATFORM_WINDOWS
#include <Windows.h>
#undef FindResource
#endif

#define LOG_PREFIX "[GOLM] "

using namespace golm;
using namespace plugify;

void std::default_delete<DCaggr>::operator()(DCaggr* p) const {
	dcFreeAggr(p);
}

void std::default_delete<DCCallVM>::operator()(DCCallVM* p) const {
	dcFree(p);
}

static thread_local VirtualMachine s_vm;

[[nodiscard]] DCCallVM& VirtualMachine::operator()() {
	if (_callVirtMachine == nullptr) {
		DCCallVM* vm = dcNewCallVM(4096);
		dcMode(vm, DC_CALL_C_DEFAULT);
		_callVirtMachine = std::unique_ptr<DCCallVM>(vm);
	}
	return *_callVirtMachine;
}

namespace {
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

	GoSlice* CreateGoSliceBool(const plg::vector<bool>& source, ArgumentList& args, BoolHolder& holder) {
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

	GoSlice* CreateGoSliceString(const plg::vector<plg::string>& source, ArgumentList& args, StringHolder& holder) {
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
	GoSlice* CreateGoSlice(const plg::vector<T>& source, ArgumentList& args) {
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
	void CopyGoSliceToVector(const GoSlice& source, plg::vector<T>& dest) {
		if (source.data == nullptr || source.len == 0) [[unlikely]]
			dest.clear();
		else if (dest.data() != source.data)
			dest.assign(reinterpret_cast<T*>(source.data), reinterpret_cast<T*>(source.data) + static_cast<size_t>(source.len));
	}

	template<>
	void CopyGoSliceToVector(const GoSlice& source, plg::vector<plg::string>& dest) {
		dest.resize(static_cast<size_t>(source.len));
		for (size_t i = 0; i < dest.size(); ++i) {
			const auto& str = reinterpret_cast<GoString*>(source.data)[i];
			dest[i].assign(str.p, static_cast<size_t>(str.n));
		}
	}

	void CopyGoStringToString(const GoString& source, plg::string& dest) {
		if (source.p == nullptr || source.n == 0) [[unlikely]]
			dest.clear();
		else if (dest.data() != source.p)
			dest.assign(source.p, static_cast<size_t>(source.n));
	}

	template<typename T>
	void CopyGoSliceToVectorReturn(const GoSlice& source, plg::vector<T>& dest) {
		if (source.data == nullptr || source.len == 0) [[unlikely]]
			std::construct_at(&dest);
		else
			std::construct_at(&dest, reinterpret_cast<T*>(source.data), reinterpret_cast<T*>(source.data) + static_cast<size_t>(source.len));
	}

	template<>
	void CopyGoSliceToVectorReturn(const GoSlice& source, plg::vector<plg::string>& dest) {
		if (source.data == nullptr || source.len == 0) [[unlikely]]
			std::construct_at(&dest);
		else {
			std::construct_at(&dest, static_cast<size_t>(source.len));
			for (size_t i = 0; i < dest.size(); ++i) {
				const auto& str = reinterpret_cast<GoString*>(source.data)[i];
				dest[i].assign(str.p, static_cast<size_t>(str.n));
			}
		}
	}

	void CopyGoStringToStringReturn(const GoString& source, plg::string& dest) {
		if (source.p == nullptr || source.n == 0) [[unlikely]]
			std::construct_at(&dest);
		else
			std::construct_at(&dest, source.p, static_cast<size_t>(source.n));
	}

	template<typename T>
	DCaggr* CreateDcAggr(AggrList&) {
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
}

InitResult GoLanguageModule::Initialize(std::weak_ptr<IPlugifyProvider> provider, ModuleRef /*module*/) {
	if (!(_provider = provider.lock())) {
		return ErrorData{ "Provider not exposed" };
	}

	_provider->Log(LOG_PREFIX "Inited!", Severity::Debug);

	_rt = std::make_shared<asmjit::JitRuntime>();

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
				JitCallback callback(_rt);
				func = callback.GetJitFunc(method, &InternalCall, func);
				if (!func) {
					funcErrors.emplace_back(method.GetName());
					continue;
				}
				_functions.emplace_back(std::move(callback));
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
void GoLanguageModule::InternalCall(MethodRef method, MemAddr addr, const JitCallback::Parameters* p, uint8_t count, const JitCallback::Return* ret) {
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

	DCCallVM* vm = &s_vm();
	dcReset(vm);

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

	for (uint8_t i = 0; i < count; ++i) {
		const auto& param = paramProps[i];
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
					dcArgPointer(vm, CreateGoSliceBool(*p->GetArgument<plg::vector<bool>*>(i), args, boolHolder));
					break;
				case ValueType::ArrayChar8:
					dcArgPointer(vm, CreateGoSlice<char>(*p->GetArgument<plg::vector<char>*>(i), args));
					break;
				case ValueType::ArrayChar16:
					dcArgPointer(vm, CreateGoSlice<char16_t>(*p->GetArgument<plg::vector<char16_t>*>(i), args));
					break;
				case ValueType::ArrayInt8:
					dcArgPointer(vm, CreateGoSlice<int8_t>(*p->GetArgument<plg::vector<int8_t>*>(i), args));
					break;
				case ValueType::ArrayInt16:
					dcArgPointer(vm, CreateGoSlice<int16_t>(*p->GetArgument<plg::vector<int16_t>*>(i), args));
					break;
				case ValueType::ArrayInt32:
					dcArgPointer(vm, CreateGoSlice<int32_t>(*p->GetArgument<plg::vector<int32_t>*>(i), args));
					break;
				case ValueType::ArrayInt64:
					dcArgPointer(vm, CreateGoSlice<int64_t>(*p->GetArgument<plg::vector<int64_t>*>(i), args));
					break;
				case ValueType::ArrayUInt8:
					dcArgPointer(vm, CreateGoSlice<uint8_t>(*p->GetArgument<plg::vector<uint8_t>*>(i), args));
					break;
				case ValueType::ArrayUInt16:
					dcArgPointer(vm, CreateGoSlice<uint16_t>(*p->GetArgument<plg::vector<uint16_t>*>(i), args));
					break;
				case ValueType::ArrayUInt32:
					dcArgPointer(vm, CreateGoSlice<uint32_t>(*p->GetArgument<plg::vector<uint32_t>*>(i), args));
					break;
				case ValueType::ArrayUInt64:
					dcArgPointer(vm, CreateGoSlice<uint64_t>(*p->GetArgument<plg::vector<uint64_t>*>(i), args));
					break;
				case ValueType::ArrayPointer:
					dcArgPointer(vm, CreateGoSlice<uintptr_t>(*p->GetArgument<plg::vector<uintptr_t>*>(i), args));
					break;
				case ValueType::ArrayFloat:
					dcArgPointer(vm, CreateGoSlice<float>(*p->GetArgument<plg::vector<float>*>(i), args));
					break;
				case ValueType::ArrayDouble:
					dcArgPointer(vm, CreateGoSlice<double>(*p->GetArgument<plg::vector<double>*>(i), args));
					break;
				case ValueType::ArrayString:
					dcArgPointer(vm, CreateGoSliceString(*p->GetArgument<plg::vector<plg::string>*>(i), args, stringHolder));
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
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSliceBool(*p->GetArgument<plg::vector<bool>*>(i), args, boolHolder));
					break;
				case ValueType::ArrayChar8:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<char>(*p->GetArgument<plg::vector<char>*>(i), args));
					break;
				case ValueType::ArrayChar16:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<char16_t>(*p->GetArgument<plg::vector<char16_t>*>(i), args));
					break;
				case ValueType::ArrayInt8:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<int8_t>(*p->GetArgument<plg::vector<int8_t>*>(i), args));
					break;
				case ValueType::ArrayInt16:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<int16_t>(*p->GetArgument<plg::vector<int16_t>*>(i), args));
					break;
				case ValueType::ArrayInt32:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<int32_t>(*p->GetArgument<plg::vector<int32_t>*>(i), args));
					break;
				case ValueType::ArrayInt64:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<int64_t>(*p->GetArgument<plg::vector<int64_t>*>(i), args));
					break;
				case ValueType::ArrayUInt8:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<uint8_t>(*p->GetArgument<plg::vector<uint8_t>*>(i), args));
					break;
				case ValueType::ArrayUInt16:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<uint16_t>(*p->GetArgument<plg::vector<uint16_t>*>(i), args));
					break;
				case ValueType::ArrayUInt32:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<uint32_t>(*p->GetArgument<plg::vector<uint32_t>*>(i), args));
					break;
				case ValueType::ArrayUInt64:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<uint64_t>(*p->GetArgument<plg::vector<uint64_t>*>(i), args));
					break;
				case ValueType::ArrayPointer:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<uintptr_t>(*p->GetArgument<plg::vector<uintptr_t>*>(i), args));
					break;
				case ValueType::ArrayFloat:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<float>(*p->GetArgument<plg::vector<float>*>(i), args));
					break;
				case ValueType::ArrayDouble:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSlice<double>(*p->GetArgument<plg::vector<double>*>(i), args));
					break;
				case ValueType::ArrayString:
					dcArgAggr(vm, CreateDcAggr<GoSlice>(aggrs), CreateGoSliceString(*p->GetArgument<plg::vector<plg::string>*>(i), args, stringHolder));
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
			ret->SetReturn(val);
			break;
		}
		case ValueType::Char8: {
			char16_t val = static_cast<char16_t>(dcCallChar(vm, addr));
			ret->SetReturn(val);
			break;
		}
		case ValueType::Char16: {
			char16_t val = static_cast<char16_t>(dcCallShort(vm, addr));
			ret->SetReturn(val);
			break;
		}
		case ValueType::Int8: {
			int8_t val = dcCallChar(vm, addr);
			ret->SetReturn(val);
			break;
		}
		case ValueType::Int16: {
			int16_t val = dcCallShort(vm, addr);
			ret->SetReturn(val);
			break;
		}
		case ValueType::Int32: {
			int32_t val = dcCallInt(vm, addr);
			ret->SetReturn(val);
			break;
		}
		case ValueType::Int64: {
			int64_t val = dcCallLongLong(vm, addr);
			ret->SetReturn(val);
			break;
		}
		case ValueType::UInt8: {
			uint8_t val = static_cast<uint8_t>(dcCallChar(vm, addr));
			ret->SetReturn(val);
			break;
		}
		case ValueType::UInt16: {
			uint16_t val = static_cast<uint16_t>(dcCallShort(vm, addr));
			ret->SetReturn(val);
			break;
		}
		case ValueType::UInt32: {
			uint32_t val = static_cast<uint32_t>(dcCallInt(vm, addr));
			ret->SetReturn(val);
			break;
		}
		case ValueType::UInt64: {
			uint64_t val = static_cast<uint64_t>(dcCallLongLong(vm, addr));
			ret->SetReturn(val);
			break;
		}
		case ValueType::Pointer: {
			void* val = dcCallPointer(vm, addr);
			ret->SetReturn(val);
			break;
		}
		case ValueType::Float: {
			float val = dcCallFloat(vm, addr);
			ret->SetReturn(val);
			break;
		}
		case ValueType::Double: {
			double val = dcCallDouble(vm, addr);
			ret->SetReturn(val);
			break;
		}
		case ValueType::Function: {
			void* val = dcCallPointer(vm, addr);
			ret->SetReturn(val);
			break;
		}
		case ValueType::Vector2:
		case ValueType::Vector3:
		case ValueType::Vector4:
		case ValueType::Matrix4x4: {
			dcCallAggr(vm, addr, ag, ret->GetReturnPtr());
			break;
		}
		case ValueType::String: {
			GoString source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoStringToStringReturn(source, *reinterpret_cast<plg::string*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayBool: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<bool>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayChar8: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<char>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayChar16: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<char16_t>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayInt8: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<int8_t>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayInt16: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<int16_t>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayInt32: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<int32_t>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayInt64: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<int64_t>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayUInt8: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<uint8_t>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayUInt16: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<uint16_t>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayUInt32: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<uint32_t>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayUInt64: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<uint64_t>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayPointer: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<uintptr_t>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayFloat: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<float>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayDouble: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<double>*>(ret->GetReturnPtr()));
			break;
		}
		case ValueType::ArrayString: {
			GoSlice source;
			dcCallAggr(vm, addr, ag, &source);
			CopyGoSliceToVectorReturn(source, *reinterpret_cast<plg::vector<plg::string>*>(ret->GetReturnPtr()));
			break;
		}
		default:
			std::puts(LOG_PREFIX "Unsupported types!\n");
			std::terminate();
			break;
	}

	if (argsCount != 0) {
		if (hasRefs) {
			size_t j = 0;
			for (uint8_t i = 0; i < count; ++i) {
				const auto& param = paramProps[i];
				if (param.IsReference()) {
					switch (param.GetType()) {
						case ValueType::String:
							CopyGoStringToString(*reinterpret_cast<GoString*>(&args[j++]), *p->GetArgument<plg::string*>(i));
							break;
						case ValueType::ArrayBool:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<bool>*>(i));
							break;
						case ValueType::ArrayChar8:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<char>*>(i));
							break;
						case ValueType::ArrayChar16:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<char16_t>*>(i));
							break;
						case ValueType::ArrayInt8:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<int8_t>*>(i));
							break;
						case ValueType::ArrayInt16:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<int16_t>*>(i));
							break;
						case ValueType::ArrayInt32:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<int32_t>*>(i));
							break;
						case ValueType::ArrayInt64:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<int64_t>*>(i));
							break;
						case ValueType::ArrayUInt8:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<uint8_t>*>(i));
							break;
						case ValueType::ArrayUInt16:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<uint16_t>*>(i));
							break;
						case ValueType::ArrayUInt32:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<uint32_t>*>(i));
							break;
						case ValueType::ArrayUInt64:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<uint64_t>*>(i));
							break;
						case ValueType::ArrayPointer:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<uintptr_t>*>(i));
							break;
						case ValueType::ArrayFloat:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<float>*>(i));
							break;
						case ValueType::ArrayDouble:
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<double>*>(i));
							break;
						case ValueType::ArrayString: {
							CopyGoSliceToVector(args[j++], *p->GetArgument<plg::vector<plg::string>*>(i));
							break;
						}
						default:
							break;
					}
				}
				if (j == argsCount)
					break;
			}
		}
	}
}

namespace golm {
	GoLanguageModule g_golm;
}

MemAddr GetMethodPtr(const char* methodName) {
	return g_golm.GetNativeMethod(methodName);
}

void GetMethodPtr2(const char* methodName, MemAddr* addressDest) {
	g_golm.GetNativeMethod(methodName, addressDest);
}

const char* GetBaseDir() {
	auto source = fs::path(g_golm.GetProvider()->GetBaseDir()).string();
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
	auto source = fs::path(plugin.GetBaseDir()).string();
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.c_str(), size);
	return dest;
}

const char** GetPluginDependencies(PluginRef plugin) {
	std::span<const PluginReferenceDescriptorRef> dependencies = plugin.GetDescriptor().GetDependencies();
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
	auto resource = plugin.FindResource(fs::path(path).c_str());
	if (resource.has_value()) {
		auto source= fs::path(*resource).string();
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

plugify::String ConstructString(GoString source) {
	plugify::String ret;
	if (source.n == 0 || source.p == nullptr) [[unlikely]]
		std::construct_at(reinterpret_cast<plg::string*>(&ret));
	else
		std::construct_at(reinterpret_cast<plg::string*>(&ret), source.p, static_cast<size_t>(source.n));
	return ret;
}
void DestroyString(plg::string* string) {
	string->~basic_string();
}
const char* GetStringData(plg::string* string) {
	return string->c_str();
}
ptrdiff_t GetStringLength(plg::string* string) {
	return static_cast<ptrdiff_t>(string->length());
}
void AssignString(plg::string* string, GoString source) {
	if (source.p == nullptr || source.n == 0) [[unlikely]]
		string->clear();
	else
		string->assign(source.p, static_cast<size_t>(source.n));
}

enum class DataType : uint8_t {
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

#if defined(__clang__)
#define GOLM_FORCE_INLINE [[gnu::always_inline]] [[gnu::gnu_inline]] extern inline
#elif defined(__GNUC__)
#define GOLM_FORCE_INLINE [[gnu::always_inline]] inline
#elif defined(_MSC_VER)
#pragma warning(error: 4714)
#define GOLM_FORCE_INLINE __forceinline
#else
#define GOTLM_FORCE_INLINE inline
#endif

namespace {
	template<typename T>
	GOLM_FORCE_INLINE plugify::Vector ConstructVector(T* arr, ptrdiff_t len) requires(!std::is_same_v<T, GoString>) {
		plugify::Vector ret;
		if (arr == nullptr || len == 0) [[unlikely]]
			std::construct_at(reinterpret_cast<plg::vector<T>*>(&ret));
		else
			std::construct_at(reinterpret_cast<plg::vector<T>*>(&ret), arr, arr + len);
		return ret;
	}

	template<typename T>
	GOLM_FORCE_INLINE plugify::Vector ConstructVector(T* arr, ptrdiff_t len) requires(std::is_same_v<T, GoString>) {
		plugify::Vector ret;
		auto* vector = reinterpret_cast<plg::vector<plg::string>*>(&ret);
		std::construct_at(vector);
		size_t N = static_cast<size_t>(len);
		if (arr == nullptr || N == 0) [[unlikely]]
			return ret;
		vector->reserve(N);
		for (size_t i = 0; i < N; ++i) {
			const auto& str = arr[i];
			vector->emplace_back(str.p, static_cast<size_t>(str.n));
		}
		return ret;
	}

	template<typename T>
	GOLM_FORCE_INLINE void DestroyVector(plg::vector<T>* vector) {
		vector->~vector_base();
	}

	template<typename T>
	GOLM_FORCE_INLINE ptrdiff_t GetVectorSize(plg::vector<T>* vector) {
		return static_cast<ptrdiff_t>(vector->size());
	}

	template<typename T>
	GOLM_FORCE_INLINE void AssignVector(plg::vector<T>* vector, T* arr, ptrdiff_t len) requires(!std::is_same_v<T, GoString>) {
		if (arr == nullptr || len == 0) [[unlikely]]
			vector->clear();
		else
			vector->assign(arr, arr + len);
	}

	template<typename T>
	GOLM_FORCE_INLINE void AssignVector(plg::vector<plg::string>* vector, T* arr, ptrdiff_t len) requires(std::is_same_v<T, GoString>) {
		if (arr == nullptr || len == 0) [[unlikely]]
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
	GOLM_FORCE_INLINE T* GetVectorData(plg::vector<T>* vector) {
		return vector->data();
	}
}


plugify::Vector ConstructVector(void* arr, ptrdiff_t len, DataType type) {
	switch (type) {
		case DataType::Bool:
			return ConstructVector(static_cast<bool*>(arr), len);
		case DataType::Char8:
			return ConstructVector(static_cast<char*>(arr), len);
		case DataType::Char16:
			return ConstructVector(static_cast<char16_t*>(arr), len);
		case DataType::Int8:
			return ConstructVector(static_cast<int8_t*>(arr), len);
		case DataType::Int16:
			return ConstructVector(static_cast<int16_t*>(arr), len);
		case DataType::Int32:
			return ConstructVector(static_cast<int32_t*>(arr), len);
		case DataType::Int64:
			return ConstructVector(static_cast<int64_t*>(arr), len);
		case DataType::UInt8:
			return ConstructVector(static_cast<uint8_t*>(arr), len);
		case DataType::UInt16:
			return ConstructVector(static_cast<uint16_t*>(arr), len);
		case DataType::UInt32:
			return ConstructVector(static_cast<uint32_t*>(arr), len);
		case DataType::UInt64:
			return ConstructVector(static_cast<uint64_t*>(arr), len);
		case DataType::Pointer:
			return ConstructVector(static_cast<uintptr_t*>(arr), len);
		case DataType::Float:
			return ConstructVector(static_cast<float*>(arr), len);
		case DataType::Double:
			return ConstructVector(static_cast<double*>(arr), len);
		case DataType::String:
			return ConstructVector(static_cast<GoString*>(arr), len);
		default:
			break;
	}
	return {};
}

void DestroyVector(void* ptr, DataType type) {
	switch (type) {
		case DataType::Bool:
			DestroyVector(reinterpret_cast<plg::vector<bool>*>(ptr));
			break;
		case DataType::Char8:
			DestroyVector(reinterpret_cast<plg::vector<char>*>(ptr));
			break;
		case DataType::Char16:
			DestroyVector(reinterpret_cast<plg::vector<char16_t>*>(ptr));
			break;
		case DataType::Int8:
			DestroyVector(reinterpret_cast<plg::vector<int8_t>*>(ptr));
			break;
		case DataType::Int16:
			DestroyVector(reinterpret_cast<plg::vector<int16_t>*>(ptr));
			break;
		case DataType::Int32:
			DestroyVector(reinterpret_cast<plg::vector<int32_t>*>(ptr));
			break;
		case DataType::Int64:
			DestroyVector(reinterpret_cast<plg::vector<int64_t>*>(ptr));
			break;
		case DataType::UInt8:
			DestroyVector(reinterpret_cast<plg::vector<uint8_t>*>(ptr));
			break;
		case DataType::UInt16:
			DestroyVector(reinterpret_cast<plg::vector<uint16_t>*>(ptr));
			break;
		case DataType::UInt32:
			DestroyVector(reinterpret_cast<plg::vector<uint32_t>*>(ptr));
			break;
		case DataType::UInt64:
			DestroyVector(reinterpret_cast<plg::vector<uint64_t>*>(ptr));
			break;
		case DataType::Pointer:
			DestroyVector(reinterpret_cast<plg::vector<uintptr_t>*>(ptr));
			break;
		case DataType::Float:
			DestroyVector(reinterpret_cast<plg::vector<float>*>(ptr));
			break;
		case DataType::Double:
			DestroyVector(reinterpret_cast<plg::vector<double>*>(ptr));
			break;
		case DataType::String:
			DestroyVector(reinterpret_cast<plg::vector<plg::string>*>(ptr));
			break;
		default:
			// Invalid type, do nothing or handle error if needed
			break;
	}
}

ptrdiff_t GetVectorSize(void* ptr, DataType type) {
	switch (type) {
		case DataType::Char8:
			return GetVectorSize(reinterpret_cast<plg::vector<char>*>(ptr));
		case DataType::Char16:
			return GetVectorSize(reinterpret_cast<plg::vector<char16_t>*>(ptr));
		case DataType::Int8:
			return GetVectorSize(reinterpret_cast<plg::vector<int8_t>*>(ptr));
		case DataType::Int16:
			return GetVectorSize(reinterpret_cast<plg::vector<int16_t>*>(ptr));
		case DataType::Int32:
			return GetVectorSize(reinterpret_cast<plg::vector<int32_t>*>(ptr));
		case DataType::Int64:
			return GetVectorSize(reinterpret_cast<plg::vector<int64_t>*>(ptr));
		case DataType::UInt8:
			return GetVectorSize(reinterpret_cast<plg::vector<uint8_t>*>(ptr));
		case DataType::UInt16:
			return GetVectorSize(reinterpret_cast<plg::vector<uint16_t>*>(ptr));
		case DataType::UInt32:
			return GetVectorSize(reinterpret_cast<plg::vector<uint32_t>*>(ptr));
		case DataType::UInt64:
			return GetVectorSize(reinterpret_cast<plg::vector<uint64_t>*>(ptr));
		case DataType::Pointer:
			return GetVectorSize(reinterpret_cast<plg::vector<uintptr_t>*>(ptr));
		case DataType::Float:
			return GetVectorSize(reinterpret_cast<plg::vector<float>*>(ptr));
		case DataType::Double:
			return GetVectorSize(reinterpret_cast<plg::vector<double>*>(ptr));
		case DataType::Bool:
			return GetVectorSize(reinterpret_cast<plg::vector<bool>*>(ptr));
		case DataType::String:
			return GetVectorSize(reinterpret_cast<plg::vector<plg::string>*>(ptr));
		default:
			return -1; // Return -1 or some error code for invalid type
	}
}

void* GetVectorData(void* ptr, DataType type) {
	switch (type) {
		case DataType::Bool:
			return GetVectorData<>(reinterpret_cast<plg::vector<bool>*>(ptr));
		case DataType::Char8:
			return GetVectorData<>(reinterpret_cast<plg::vector<char>*>(ptr));
		case DataType::Char16:
			return GetVectorData<>(reinterpret_cast<plg::vector<char16_t>*>(ptr));
		case DataType::Int8:
			return GetVectorData<>(reinterpret_cast<plg::vector<int8_t>*>(ptr));
		case DataType::Int16:
			return GetVectorData<>(reinterpret_cast<plg::vector<int16_t>*>(ptr));
		case DataType::Int32:
			return GetVectorData<>(reinterpret_cast<plg::vector<int32_t>*>(ptr));
		case DataType::Int64:
			return GetVectorData<>(reinterpret_cast<plg::vector<int64_t>*>(ptr));
		case DataType::UInt8:
			return GetVectorData<>(reinterpret_cast<plg::vector<uint8_t>*>(ptr));
		case DataType::UInt16:
			return GetVectorData<>(reinterpret_cast<plg::vector<uint16_t>*>(ptr));
		case DataType::UInt32:
			return GetVectorData<>(reinterpret_cast<plg::vector<uint32_t>*>(ptr));
		case DataType::UInt64:
			return GetVectorData<>(reinterpret_cast<plg::vector<uint64_t>*>(ptr));
		case DataType::Pointer:
			return GetVectorData<>(reinterpret_cast<plg::vector<uintptr_t>*>(ptr));
		case DataType::Float:
			return GetVectorData<>(reinterpret_cast<plg::vector<float>*>(ptr));
		case DataType::Double:
			return GetVectorData<>(reinterpret_cast<plg::vector<double>*>(ptr));
		case DataType::String:
			return GetVectorData<>(reinterpret_cast<plg::vector<plg::string>*>(ptr));
		default:
			return nullptr; // Return nullptr for invalid type
	}
}

void AssignVector(void* ptr, void* arr, ptrdiff_t len, DataType type) {
	switch (type) {
		case DataType::Bool:
			AssignVector(reinterpret_cast<plg::vector<bool>*>(ptr), static_cast<bool*>(arr), len);
			break;
		case DataType::Char8:
			AssignVector(reinterpret_cast<plg::vector<char>*>(ptr), static_cast<char*>(arr), len);
			break;
		case DataType::Char16:
			AssignVector(reinterpret_cast<plg::vector<char16_t>*>(ptr), static_cast<char16_t*>(arr), len);
			break;
		case DataType::Int8:
			AssignVector(reinterpret_cast<plg::vector<int8_t>*>(ptr), static_cast<int8_t*>(arr), len);
			break;
		case DataType::Int16:
			AssignVector(reinterpret_cast<plg::vector<int16_t>*>(ptr), static_cast<int16_t*>(arr), len);
			break;
		case DataType::Int32:
			AssignVector(reinterpret_cast<plg::vector<int32_t>*>(ptr), static_cast<int32_t*>(arr), len);
			break;
		case DataType::Int64:
			AssignVector(reinterpret_cast<plg::vector<int64_t>*>(ptr), static_cast<int64_t*>(arr), len);
			break;
		case DataType::UInt8:
			AssignVector(reinterpret_cast<plg::vector<uint8_t>*>(ptr), static_cast<uint8_t*>(arr), len);
			break;
		case DataType::UInt16:
			AssignVector(reinterpret_cast<plg::vector<uint16_t>*>(ptr), static_cast<uint16_t*>(arr), len);
			break;
		case DataType::UInt32:
			AssignVector(reinterpret_cast<plg::vector<uint32_t>*>(ptr), static_cast<uint32_t*>(arr), len);
			break;
		case DataType::UInt64:
			AssignVector(reinterpret_cast<plg::vector<uint64_t>*>(ptr), static_cast<uint64_t*>(arr), len);
			break;
		case DataType::Pointer:
			AssignVector(reinterpret_cast<plg::vector<uintptr_t>*>(ptr), static_cast<uintptr_t*>(arr), len);
			break;
		case DataType::Float:
			AssignVector(reinterpret_cast<plg::vector<float>*>(ptr), static_cast<float*>(arr), len);
			break;
		case DataType::Double:
			AssignVector(reinterpret_cast<plg::vector<double>*>(ptr), static_cast<double*>(arr), len);
			break;
		case DataType::String:
			AssignVector(reinterpret_cast<plg::vector<plg::string>*>(ptr), static_cast<GoString*>(arr), len);
			break;
		default:
			break;
	}
}

void DeleteVectorDataCStr(void* ptr) {
	delete[] reinterpret_cast<char**>(ptr);
}

const std::array<void*, 28> GoLanguageModule::_pluginApi = {
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
		reinterpret_cast<void*>(&::ConstructString),
		reinterpret_cast<void*>(&::DestroyString),
		reinterpret_cast<void*>(&::GetStringData),
		reinterpret_cast<void*>(&::GetStringLength),
		reinterpret_cast<void*>(&::AssignString),
		reinterpret_cast<void*>(&::ConstructVector),
		reinterpret_cast<void*>(&::DestroyVector),
		reinterpret_cast<void*>(&::GetVectorData),
		reinterpret_cast<void*>(&::GetVectorSize),
		reinterpret_cast<void*>(&::AssignVector),
		reinterpret_cast<void*>(&::DeleteVectorDataCStr)
};

ILanguageModule* GetLanguageModule() {
	return &golm::g_golm;
}
