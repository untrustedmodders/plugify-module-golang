#include "module.hpp"

#include <plugify/module.hpp>
#include <plugify/plugin.hpp>
#include <plugify/log.hpp>
#include <plugify/string.hpp>
#include <plugify/numerics.hpp>
#include <plugify/plugin_descriptor.hpp>
#include <plugify/plugin_reference_descriptor.hpp>
#include <plugify/plugify_provider.hpp>
#include <plugify/jit/call.hpp>
#include <plugify/jit/helpers.hpp>
#include <plugify/numerics.hpp>
#include <plugify/any.hpp>

#include <cpptrace/cpptrace.hpp>

#if GOLM_PLATFORM_WINDOWS
#include <windows.h>
#undef FindResource
#endif

#define LOG_PREFIX "[GOLM] "

using namespace golm;
using namespace plugify;

namespace {
	std::vector<std::string_view> Split(std::string_view strv, std::string_view delims) {
		std::vector<std::string_view> output;
		size_t first = 0;

		while (first < strv.size()) {
			const size_t second = strv.find_first_of(delims, first);

			if (first != second)
				output.emplace_back(strv.substr(first, second-first));

			if (second == std::string_view::npos)
				break;

			first = second + 1;
		}

		return output;
	}
}

InitResult GoLanguageModule::Initialize(std::weak_ptr<IPlugifyProvider> provider, ModuleHandle /*module*/) {
	if (!((_provider = provider.lock()))) {
		return ErrorData{ "Provider not exposed" };
	}

	_rt = std::make_shared<asmjit::JitRuntime>();

	_provider->Log(LOG_PREFIX "Inited!", Severity::Debug);

	return InitResultData{{ .hasUpdate = false }};
}

void GoLanguageModule::Shutdown() {
	for (MemAddr* addr : _addresses) {
		*addr = nullptr;
	}
	_nativesMap.clear();
	_addresses.clear();
	_assemblies.clear();
	_rt.reset();
	_provider.reset();
}

void GoLanguageModule::OnUpdate(plugify::DateTime) {
	
}

bool GoLanguageModule::IsDebugBuild() {
	return GOLM_IS_DEBUG;
}

void GoLanguageModule::OnMethodExport(PluginHandle plugin) {
	for (const auto& [method, addr] : plugin.GetMethods()) {
		_nativesMap.try_emplace(std::format("{}.{}", plugin.GetName(), method.GetName()), addr);
	}
}

LoadResult GoLanguageModule::OnPluginLoad(PluginHandle plugin) {
	fs::path assemblyPath(plugin.GetBaseDir());
	assemblyPath /= std::format("{}" GOLM_LIBRARY_SUFFIX, plugin.GetDescriptor().GetEntryPoint());

	auto assembly = std::make_unique<Assembly>(assemblyPath, LoadFlag::Lazy | LoadFlag::Nodelete | LoadFlag::PinInMemory);
	if (!assembly->IsValid()) {
		return ErrorData{ std::format("Failed to load assembly: {}", assembly->GetError()) };
	}

	auto* const initFunc = assembly->GetFunctionByName("Plugify_Init").RCast<InitFunc>();
	if (!initFunc) {
		return ErrorData{ "Not found 'Plugify_Init' function" };
	}

	auto* const callFunc = assembly->GetFunctionByName("Plugify_InternalCall").RCast<JitCallback::CallbackHandler>();
	if (!callFunc) {
		return ErrorData{ "Not found 'Plugify_InternalCall' function" };
	}

	auto* const startFunc = assembly->GetFunctionByName("Plugify_PluginStart").RCast<StartFunc>();
	auto* const updateFunc = assembly->GetFunctionByName("Plugify_PluginUpdate").RCast<UpdateFunc>();
	auto* const endFunc = assembly->GetFunctionByName("Plugify_PluginEnd").RCast<EndFunc>();
	auto* const contextFunc = assembly->GetFunctionByName("Plugify_PluginContext").RCast<ContextFunc>();

	std::vector<std::string_view> funcErrors;

	std::span<const MethodHandle> exportedMethods = plugin.GetDescriptor().GetExportedMethods();
	std::vector<MethodData> methods;
	methods.reserve(exportedMethods.size());

	for (const auto& method : exportedMethods) {
		if (auto func = assembly->GetFunctionByName(method.GetFunctionName())) {
			methods.emplace_back(method, func);
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

	GoSlice api { const_cast<void**>(_pluginApi.data()), _pluginApi.size(), _pluginApi.size() };
	const int resultVersion = initFunc(api, kApiVersion, plugin);
	if (resultVersion != 0) {
		return ErrorData{ std::format("Not supported plugin api {}, max supported {}", resultVersion, kApiVersion) };
	}

	const auto& [hasUpdate, hasStart, hasEnd, _] = contextFunc ? *(contextFunc()) : PluginContext{};

	auto data = _assemblies.emplace_back(std::make_unique<AssemblyHolder>(std::move(assembly), updateFunc, startFunc, endFunc, contextFunc, callFunc)).get();
	return LoadResultData{ std::move(methods), data, { hasUpdate, hasStart, hasEnd, !exportedMethods.empty() } };
}

void GoLanguageModule::OnPluginStart(PluginHandle plugin) {
	plugin.GetData().RCast<AssemblyHolder*>()->startFunc();
}

void GoLanguageModule::OnPluginUpdate(plugify::PluginHandle plugin, plugify::DateTime dt) {
	plugin.GetData().RCast<AssemblyHolder*>()->updateFunc(dt.AsSeconds());
}

void GoLanguageModule::OnPluginEnd(PluginHandle plugin) {
	plugin.GetData().RCast<AssemblyHolder*>()->endFunc();
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

MethodHandle GoLanguageModule::FindMethod(std::string_view name) {
	auto separated = Split(name, ".");
	if (separated.size() != 2)
		return {};

	auto plugin = _provider->FindPlugin(separated[0]);
	if (plugin) {
		for (const auto& method : plugin.GetDescriptor().GetExportedMethods()) {
			auto prototype = method.FindPrototype(separated[1]);
			if (prototype) {
				return prototype;
			}
		}
	}
	return {};
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

bool IsModuleLoaded(GoString moduleName, GoString versionName, bool minimum) {
	if (std::string_view version = versionName; !version.empty())
		return g_golm.GetProvider()->IsModuleLoaded(moduleName, plg::version(version), minimum);
	else
		return g_golm.GetProvider()->IsModuleLoaded(moduleName, std::nullopt, minimum);
}

bool IsPluginLoaded(GoString pluginName, GoString versionName, bool minimum) {
	if (std::string_view version = versionName; !version.empty())
		return g_golm.GetProvider()->IsPluginLoaded(pluginName, plg::version(version), minimum);
	else
		return g_golm.GetProvider()->IsPluginLoaded(pluginName, std::nullopt, minimum);
}

void PrintException(GoString message) {
	if (const auto& provider = g_golm.GetProvider()) {
		provider->Log(std::format(LOG_PREFIX "[Exception] {}", std::string_view(message)), Severity::Error);

		std::stringstream stream;
		cpptrace::generate_trace().print(stream);
		provider->Log(stream.str(), Severity::Error);
	}
}

ptrdiff_t GetPluginId(PluginHandle plugin) {
	return static_cast<ptrdiff_t>(plugin.GetId());
}

const char* GetPluginName(PluginHandle plugin) {
	return plugin.GetName().data();
}

const char* GetPluginFullName(PluginHandle plugin) {
	return plugin.GetFriendlyName().data();
}

const char* GetPluginDescription(PluginHandle plugin) {
	return plugin.GetDescriptor().GetDescription().data();
}

const char* GetPluginVersion(PluginHandle plugin) {
	return plugin.GetDescriptor().GetVersionName().data();
}

const char* GetPluginAuthor(PluginHandle plugin) {
	return plugin.GetDescriptor().GetCreatedBy().data();
}

const char* GetPluginWebsite(PluginHandle plugin) {
	return plugin.GetDescriptor().GetCreatedByURL().data();
}

#if GOLM_PLATFORM_WINDOWS
namespace {
	bool ConvertUtf8ToWide(std::wstring& dest, std::string_view str) {
		int wlen = MultiByteToWideChar(CP_UTF8, 0, str.data(), static_cast<int>(str.length()), nullptr, 0);
		if (wlen < 0)
			return false;

		dest.resize(static_cast<size_t>(wlen));
		if (wlen > 0 && MultiByteToWideChar(CP_UTF8, 0, str.data(), static_cast<int>(str.length()), dest.data(), wlen) < 0)
			return false;

		return true;
	}

	std::wstring ConvertUtf8ToWide(std::string_view str){
		std::wstring ret;
		if (!ConvertUtf8ToWide(ret, str))
			return {};
		return ret;
	}

	bool ConvertWideToUtf8(std::string& dest, std::wstring_view str) {
		int mblen = WideCharToMultiByte(CP_UTF8, 0, str.data(), static_cast<int>(str.length()), nullptr, 0, nullptr, nullptr);
		if (mblen < 0)
			return false;

		dest.resize(static_cast<size_t>(mblen));
		if (mblen > 0 && WideCharToMultiByte(CP_UTF8, 0, str.data(), static_cast<int>(str.length()), dest.data(), mblen, nullptr, nullptr) < 0)
			return false;

		return true;
	}

	std::string ConvertWideToUtf8(std::wstring_view str) {
		std::string ret;
		if (!ConvertWideToUtf8(ret, str))
			return {};
		return ret;
	}
}
#define GOLM_UTF8(str) ConvertWideToUtf8(str)
#define GOLM_PSTR(str) ConvertUtf8ToWide(str)
#else
#define GOLM_UTF8(str) str
#define GOLM_PSTR(str) str
#endif

const char* GetPluginBaseDir(PluginHandle plugin) {
	auto source = GOLM_UTF8(plugin.GetBaseDir());
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.data(), size);
	return dest;
}

const char* GetPluginConfigsDir(PluginHandle plugin) {
	auto source = GOLM_UTF8(plugin.GetConfigsDir());
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.data(), size);
	return dest;
}

const char* GetPluginDataDir(PluginHandle plugin) {
	auto source = GOLM_UTF8(plugin.GetDataDir());
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.data(), size);
	return dest;
}

const char* GetPluginLogsDir(PluginHandle plugin) {
	auto source = GOLM_UTF8(plugin.GetLogsDir());
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.data(), size);
	return dest;
}

const char** GetPluginDependencies(PluginHandle plugin) {
	std::span<const PluginReferenceDescriptorHandle> dependencies = plugin.GetDescriptor().GetDependencies();
	auto* deps = new const char*[dependencies.size()];
	for (size_t i = 0; i < dependencies.size(); ++i) {
		deps[i] = dependencies[i].GetName().data();
	}
	return deps;
}

ptrdiff_t GetPluginDependenciesSize(PluginHandle plugin) {
	return static_cast<ptrdiff_t>(plugin.GetDescriptor().GetDependencies().size());
}

const char* FindPluginResource(PluginHandle plugin, GoString path) {
	auto resource = plugin.FindResource(GOLM_PSTR(path));
	if (resource.has_value()) {
		auto source= GOLM_UTF8(*resource);
		size_t size = source.length() + 1;
		char* dest = new char[size];
		std::memcpy(dest, source.data(), size);
		return dest;
	}
	return "";
}

void DeleteCStr(const char* str) {
	delete str;
}

void DeleteCStrArr(const char** arr) {
	delete[] arr;
}

// String Functions
plg::string ConstructString(GoString source) {
	if (source.p == nullptr || source.n == 0) [[unlikely]]
		return {};
	else
		return { source.p, static_cast<size_t>(source.n) };
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

// Variant Functions
void DestroyVariant(plg::any* any) {
	any->~variant();
}

namespace {
	template<typename T>
	PLUGIFY_FORCE_INLINE plg::vector<T> ConstructVector(T* arr, ptrdiff_t len) requires(!std::is_same_v<T, GoString>) {
		if (arr == nullptr || len == 0) [[unlikely]]
			if (len > 0)
				return plg::vector<T>(static_cast<size_t>(len));
			else
				return {};
		else
			return plg::vector<T>(arr, arr + len);
	}

	template<typename T>
	/*PLUGIFY_FORCE_INLINE*/plg::vector<plg::string> ConstructVector(T* arr, ptrdiff_t len) requires(std::is_same_v<T, GoString>) {
		if (arr == nullptr || len == 0) [[unlikely]]
			if (len > 0)
				return plg::vector<plg::string>(static_cast<size_t>(len));
			else
				return {};
		else {
			plg::vector<plg::string> vector;
			size_t N = static_cast<size_t>(len);
			vector.reserve(N);
			for (size_t i = 0; i < N; ++i) {
				const auto& str = arr[i];
				vector.emplace_back(str.p, static_cast<size_t>(str.n));
			}
			return vector;
		}
	}

	template<typename T>
	PLUGIFY_FORCE_INLINE void DestroyVector(plg::vector<T>* vector) {
		vector->~vector();
	}

	template<typename T>
	PLUGIFY_FORCE_INLINE ptrdiff_t GetVectorSize(plg::vector<T>* vector) {
		return static_cast<ptrdiff_t>(vector->size());
	}

	template<typename T>
	PLUGIFY_FORCE_INLINE void AssignVector(plg::vector<T>* vector, T* arr, ptrdiff_t len) requires(!std::is_same_v<T, GoString>) {
		if (arr == nullptr || len == 0) [[unlikely]]
			vector->clear();
		else
			vector->assign(arr, arr + len);
	}

	template<typename T>
	PLUGIFY_FORCE_INLINE void AssignVector(plg::vector<plg::string>* vector, T* arr, ptrdiff_t len) requires(std::is_same_v<T, GoString>) {
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
	PLUGIFY_FORCE_INLINE T* GetVectorData(plg::vector<T>* vector) {
		return vector->data();
	}
}

plg::vector<bool> ConstructVectorBool(bool* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<char> ConstructVectorChar8(char* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<char16_t> ConstructVectorChar16(char16_t* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<int8_t> ConstructVectorInt8(int8_t* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<int16_t> ConstructVectorInt16(int16_t* arr, ptrdiff_t len) { return ConstructVector(arr, len); }

plg::vector<int32_t> ConstructVectorInt32(int32_t* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<int64_t> ConstructVectorInt64(int64_t* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<uint8_t> ConstructVectorUInt8(uint8_t* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<uint16_t> ConstructVectorUInt16(uint16_t* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<uint32_t> ConstructVectorUInt32(uint32_t* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<uint64_t> ConstructVectorUInt64(uint64_t* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<uintptr_t> ConstructVectorPointer(uintptr_t* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<float> ConstructVectorFloat(float* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<double> ConstructVectorDouble(double* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<plg::string> ConstructVectorString(GoString* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<plg::any> ConstructVectorVariant(ptrdiff_t len) { return plg::vector<plg::any>(static_cast<size_t>(len)); }
plg::vector<plg::vec2> ConstructVectorVector2(plg::vec2* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<plg::vec3> ConstructVectorVector3(plg::vec3* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<plg::vec4> ConstructVectorVector4(plg::vec4* arr, ptrdiff_t len) { return ConstructVector(arr, len); }
plg::vector<plg::mat4x4> ConstructVectorMatrix4x4(plg::mat4x4* arr, ptrdiff_t len) { return ConstructVector(arr, len); }

void DestroyVectorBool(plg::vector<bool>* vec) { DestroyVector(vec); }
void DestroyVectorChar8(plg::vector<char>* vec) { DestroyVector(vec); }
void DestroyVectorChar16(plg::vector<char16_t>* vec) { DestroyVector(vec); }
void DestroyVectorInt8(plg::vector<int8_t>* vec) { DestroyVector(vec); }
void DestroyVectorInt16(plg::vector<int16_t>* vec) { DestroyVector(vec); }
void DestroyVectorInt32(plg::vector<int32_t>* vec) { DestroyVector(vec); }
void DestroyVectorInt64(plg::vector<int64_t>* vec) { DestroyVector(vec); }
void DestroyVectorUInt8(plg::vector<uint8_t>* vec) { DestroyVector(vec); }
void DestroyVectorUInt16(plg::vector<uint16_t>* vec) { DestroyVector(vec); }
void DestroyVectorUInt32(plg::vector<uint32_t>* vec) { DestroyVector(vec); }
void DestroyVectorUInt64(plg::vector<uint64_t>* vec) { DestroyVector(vec); }
void DestroyVectorPointer(plg::vector<uintptr_t>* vec) { DestroyVector(vec); }
void DestroyVectorFloat(plg::vector<float>* vec) { DestroyVector(vec); }
void DestroyVectorDouble(plg::vector<double>* vec) { DestroyVector(vec); }
void DestroyVectorString(plg::vector<plg::string>* vec) { DestroyVector(vec); }
void DestroyVectorVariant(plg::vector<plg::any>* vector) { DestroyVector(vector); }
void DestroyVectorVector2(plg::vector<plg::vec2>* vec) { DestroyVector(vec); }
void DestroyVectorVector3(plg::vector<plg::vec3>* vec) { DestroyVector(vec); }
void DestroyVectorVector4(plg::vector<plg::vec4>* vec) { DestroyVector(vec); }
void DestroyVectorMatrix4x4(plg::vector<plg::mat4x4>* vec) { DestroyVector(vec); }

ptrdiff_t GetVectorSizeBool(plg::vector<bool>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeChar8(plg::vector<char>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeChar16(plg::vector<char16_t>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeInt8(plg::vector<int8_t>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeInt16(plg::vector<int16_t>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeInt32(plg::vector<int32_t>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeInt64(plg::vector<int64_t>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeUInt8(plg::vector<uint8_t>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeUInt16(plg::vector<uint16_t>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeUInt32(plg::vector<uint32_t>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeUInt64(plg::vector<uint64_t>* vec) { return GetVectorSize(vec); } 
ptrdiff_t GetVectorSizePointer(plg::vector<uintptr_t>* vec) { return GetVectorSize(vec); } 
ptrdiff_t GetVectorSizeFloat(plg::vector<float>* vec) { return GetVectorSize(vec); } 
ptrdiff_t GetVectorSizeDouble(plg::vector<double>* vec) { return GetVectorSize(vec); } 
ptrdiff_t GetVectorSizeString(plg::vector<plg::string>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeVariant(plg::vector<plg::any>* vector) { return GetVectorSize(vector); }
ptrdiff_t GetVectorSizeVector2(plg::vector<plg::vec2>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeVector3(plg::vector<plg::vec3>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeVector4(plg::vector<plg::vec4>* vec) { return GetVectorSize(vec); }
ptrdiff_t GetVectorSizeMatrix4x4(plg::vector<plg::mat4x4>* vec) { return GetVectorSize(vec); }

bool* GetVectorDataBool(plg::vector<bool>* vec) { return GetVectorData(vec); }
char* GetVectorDataChar8(plg::vector<char>* vec) { return GetVectorData(vec); }
char16_t* GetVectorDataChar16(plg::vector<char16_t>* vec) { return GetVectorData(vec); }
int8_t* GetVectorDataInt8(plg::vector<int8_t>* vec) { return GetVectorData(vec); }
int16_t* GetVectorDataInt16(plg::vector<int16_t>* vec) { return GetVectorData(vec); }
int32_t* GetVectorDataInt32(plg::vector<int32_t>* vec) { return GetVectorData(vec); }
int64_t* GetVectorDataInt64(plg::vector<int64_t>* vec) { return GetVectorData(vec); }
uint8_t* GetVectorDataUInt8(plg::vector<uint8_t>* vec) { return GetVectorData(vec); }
uint16_t* GetVectorDataUInt16(plg::vector<uint16_t>* vec) { return GetVectorData(vec); }
uint32_t* GetVectorDataUInt32(plg::vector<uint32_t>* vec) { return GetVectorData(vec); }
uint64_t* GetVectorDataUInt64(plg::vector<uint64_t>* vec) { return GetVectorData(vec); }
uintptr_t* GetVectorDataPointer(plg::vector<uintptr_t>* vec) { return GetVectorData(vec); }
float* GetVectorDataFloat(plg::vector<float>* vec) { return GetVectorData(vec); }
double* GetVectorDataDouble(plg::vector<double>* vec) { return GetVectorData(vec); }
plg::string* GetVectorDataString(plg::vector<plg::string>* vec) { return GetVectorData(vec); }
plg::any* GetVectorDataVariant(plg::vector<plg::any>* vec, ptrdiff_t at) { return &vec->at(static_cast<size_t>(at)); }
plg::vec2* GetVectorDataVector2(plg::vector<plg::vec2>* vec) { return GetVectorData(vec); }
plg::vec3* GetVectorDataVector3(plg::vector<plg::vec3>* vec) { return GetVectorData(vec); }
plg::vec4* GetVectorDataVector4(plg::vector<plg::vec4>* vec) { return GetVectorData(vec); }
plg::mat4x4* GetVectorDataMatrix4x4(plg::vector<plg::mat4x4>* vec) { return GetVectorData(vec); }

void AssignVectorBool(plg::vector<bool>* ptr, bool* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorChar8(plg::vector<char>* ptr, char* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorChar16(plg::vector<char16_t>* ptr, char16_t* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorInt8(plg::vector<int8_t>* ptr, int8_t* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorInt16(plg::vector<int16_t>* ptr, int16_t* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorInt32(plg::vector<int32_t>* ptr, int32_t* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorInt64(plg::vector<int64_t>* ptr, int64_t* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorUInt8(plg::vector<uint8_t>* ptr, uint8_t* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorUInt16(plg::vector<uint16_t>* ptr, uint16_t* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorUInt32(plg::vector<uint32_t>* ptr, uint32_t* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorUInt64(plg::vector<uint64_t>* ptr, uint64_t * arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorPointer(plg::vector<uintptr_t>* ptr, uintptr_t* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorFloat(plg::vector<float>* ptr, float* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorDouble(plg::vector<double>* ptr, double* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorString(plg::vector<plg::string>* ptr, GoString* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorVariant(plg::vector<plg::any>* vector, ptrdiff_t len) { vector->resize(static_cast<size_t>(len)); }
void AssignVectorVector2(plg::vector<plg::vec2>* ptr, plg::vec2* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorVector3(plg::vector<plg::vec3>* ptr, plg::vec3* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorVector4(plg::vector<plg::vec4>* ptr, plg::vec4* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }
void AssignVectorMatrix4x4(plg::vector<plg::mat4x4>* ptr, plg::mat4x4* arr, ptrdiff_t len) { AssignVector(ptr, arr, len); }

struct ManagedType {
	plugify::ValueType type{};
	bool ref{};
};

static_assert(sizeof(ManagedType) == 2, "ManagedType size mismatch with C#");

JitCall* NewCall(void* target, ManagedType* params, ptrdiff_t count, ManagedType ret) {
	if (target == nullptr)
		return nullptr;

#if GOLM_ARCH_ARM
	ValueType typeHidden = ValueType::Void;
#else
	ValueType typeHidden = ValueType::Pointer;
#endif

	bool retHidden = ValueUtils::IsHiddenParam(ret.type);
	asmjit::FuncSignature sig(asmjit::CallConvId::kCDecl, asmjit::FuncSignature::kNoVarArgs, JitUtils::GetRetTypeId(retHidden ? typeHidden : ret.type));

#if !GOLM_ARCH_ARM
	if (retHidden) {
		sig.addArg(JitUtils::GetValueTypeId(ret.type));
	}
#endif

	for (ptrdiff_t i = 0; i < count; ++i) {
		const auto& [type, ref] = params[i];
		sig.addArg(JitUtils::GetValueTypeId(ref ? ValueType::Pointer : type));
	}

	JitCall* call = new JitCall(g_golm.GetRuntime());
	call->GetJitFunc(sig, target, JitCall::WaitType::None, retHidden);
	return call;
}

void DeleteCall(JitCall* call) {
	delete call;
}

void* GetCallFunction(JitCall* call) {
	return call ? call->GetFunction() : nullptr;
}

const char* GetCallError(JitCall* call) {
	return call ? call->GetError().data() : "Target invalid";
}

JitCallback* NewCallback(PluginHandle plugin, GoString name, void* delegate) {
	MethodHandle method = g_golm.FindMethod(name);
	if (method == nullptr || delegate == nullptr)
		return nullptr;

	JitCallback* callback = new JitCallback(g_golm.GetRuntime());
	callback->GetJitFunc(method, plugin.GetData().RCast<AssemblyHolder*>()->callFunc, delegate);
	return callback;
}

void DeleteCallback(JitCallback* callback) {
	delete callback;
}

void* GetCallbackFunction(JitCallback* callback) {
	return callback ? callback->GetFunction() : nullptr;
}

const char* GetCallbackError(JitCallback* callback) {
	return callback ? callback->GetError().data() : "Method invalid";
}

ptrdiff_t GetMethodParamCount(MethodHandle handle) {
	return static_cast<ptrdiff_t>(handle.GetParamTypes().size());
}

ManagedType GetMethodParamType(MethodHandle handle, ptrdiff_t index) {
	PropertyHandle param;
	if (index < 0) {
		param = handle.GetReturnType();
	} else {
		param = handle.GetParamTypes()[static_cast<size_t>(index)];
	}
	return { param.GetType(), param.IsReference() };
}

MethodHandle GetMethodPrototype(MethodHandle handle, ptrdiff_t index) {
	if (index < 0) {
		return handle.GetReturnType().GetPrototype();
	} else {
		return handle.GetParamTypes()[static_cast<size_t>(index)].GetPrototype();
	}
}

EnumHandle GetMethodEnum(MethodHandle handle, ptrdiff_t index) {
	if (index < 0) {
		return handle.GetReturnType().GetEnum();
	} else {
		return handle.GetParamTypes()[static_cast<size_t>(index)].GetEnum();
	}
}

const std::array<void*, 139> GoLanguageModule::_pluginApi = {
		static_cast<void*>(&::GetMethodPtr),
		static_cast<void*>(&::GetMethodPtr2),
		static_cast<void*>(&::IsModuleLoaded),
		static_cast<void*>(&::IsPluginLoaded),
		static_cast<void*>(&::PrintException),

		static_cast<void*>(&::GetPluginId),
		static_cast<void*>(&::GetPluginName),
		static_cast<void*>(&::GetPluginFullName),
		static_cast<void*>(&::GetPluginDescription),
		static_cast<void*>(&::GetPluginVersion),
		static_cast<void*>(&::GetPluginAuthor),
		static_cast<void*>(&::GetPluginWebsite),
		static_cast<void*>(&::GetPluginBaseDir),
		static_cast<void*>(&::GetPluginConfigsDir),
		static_cast<void*>(&::GetPluginDataDir),
		static_cast<void*>(&::GetPluginLogsDir),
		static_cast<void*>(&::GetPluginDependencies),
		static_cast<void*>(&::GetPluginDependenciesSize),
		static_cast<void*>(&::FindPluginResource),

		static_cast<void*>(&::DeleteCStr),
		static_cast<void*>(&::DeleteCStrArr),

		static_cast<void*>(&::ConstructString),
		static_cast<void*>(&::DestroyString),
		static_cast<void*>(&::GetStringData),
		static_cast<void*>(&::GetStringLength),
		static_cast<void*>(&::AssignString),

		static_cast<void*>(&::DestroyVariant),

		static_cast<void*>(&::ConstructVectorBool),
		static_cast<void*>(&::ConstructVectorChar8),
		static_cast<void*>(&::ConstructVectorChar16),
		static_cast<void*>(&::ConstructVectorInt8),
		static_cast<void*>(&::ConstructVectorInt16),
		static_cast<void*>(&::ConstructVectorInt32),
		static_cast<void*>(&::ConstructVectorInt64),
		static_cast<void*>(&::ConstructVectorUInt8),
		static_cast<void*>(&::ConstructVectorUInt16),
		static_cast<void*>(&::ConstructVectorUInt32),
		static_cast<void*>(&::ConstructVectorUInt64),
		static_cast<void*>(&::ConstructVectorPointer),
		static_cast<void*>(&::ConstructVectorFloat),
		static_cast<void*>(&::ConstructVectorDouble),
		static_cast<void*>(&::ConstructVectorString),
		static_cast<void*>(&::ConstructVectorVariant),
		static_cast<void*>(&::ConstructVectorVector2),
		static_cast<void*>(&::ConstructVectorVector3),
		static_cast<void*>(&::ConstructVectorVector4),
		static_cast<void*>(&::ConstructVectorMatrix4x4),

		static_cast<void*>(&::DestroyVectorBool),
		static_cast<void*>(&::DestroyVectorChar8),
		static_cast<void*>(&::DestroyVectorChar16),
		static_cast<void*>(&::DestroyVectorInt8),
		static_cast<void*>(&::DestroyVectorInt16),
		static_cast<void*>(&::DestroyVectorInt32),
		static_cast<void*>(&::DestroyVectorInt64),
		static_cast<void*>(&::DestroyVectorUInt8),
		static_cast<void*>(&::DestroyVectorUInt16),
		static_cast<void*>(&::DestroyVectorUInt32),
		static_cast<void*>(&::DestroyVectorUInt64),
		static_cast<void*>(&::DestroyVectorPointer),
		static_cast<void*>(&::DestroyVectorFloat),
		static_cast<void*>(&::DestroyVectorDouble),
		static_cast<void*>(&::DestroyVectorString),
		static_cast<void*>(&::DestroyVectorVariant),
		static_cast<void*>(&::DestroyVectorVector2),
		static_cast<void*>(&::DestroyVectorVector3),
		static_cast<void*>(&::DestroyVectorVector4),
		static_cast<void*>(&::DestroyVectorMatrix4x4),

		static_cast<void*>(&::GetVectorSizeBool),
		static_cast<void*>(&::GetVectorSizeChar8),
		static_cast<void*>(&::GetVectorSizeChar16),
		static_cast<void*>(&::GetVectorSizeInt8),
		static_cast<void*>(&::GetVectorSizeInt16),
		static_cast<void*>(&::GetVectorSizeInt32),
		static_cast<void*>(&::GetVectorSizeInt64),
		static_cast<void*>(&::GetVectorSizeUInt8),
		static_cast<void*>(&::GetVectorSizeUInt16),
		static_cast<void*>(&::GetVectorSizeUInt32),
		static_cast<void*>(&::GetVectorSizeUInt64),
		static_cast<void*>(&::GetVectorSizePointer),
		static_cast<void*>(&::GetVectorSizeFloat),
		static_cast<void*>(&::GetVectorSizeDouble),
		static_cast<void*>(&::GetVectorSizeString),
		static_cast<void*>(&::GetVectorSizeVariant),
		static_cast<void*>(&::GetVectorSizeVector2),
		static_cast<void*>(&::GetVectorSizeVector3),
		static_cast<void*>(&::GetVectorSizeVector4),
		static_cast<void*>(&::GetVectorSizeMatrix4x4),

		static_cast<void*>(&::GetVectorDataBool),
		static_cast<void*>(&::GetVectorDataChar8),
		static_cast<void*>(&::GetVectorDataChar16),
		static_cast<void*>(&::GetVectorDataInt8),
		static_cast<void*>(&::GetVectorDataInt16),
		static_cast<void*>(&::GetVectorDataInt32),
		static_cast<void*>(&::GetVectorDataInt64),
		static_cast<void*>(&::GetVectorDataUInt8),
		static_cast<void*>(&::GetVectorDataUInt16),
		static_cast<void*>(&::GetVectorDataUInt32),
		static_cast<void*>(&::GetVectorDataUInt64),
		static_cast<void*>(&::GetVectorDataPointer),
		static_cast<void*>(&::GetVectorDataFloat),
		static_cast<void*>(&::GetVectorDataDouble),
		static_cast<void*>(&::GetVectorDataString),
		static_cast<void*>(&::GetVectorDataVariant),
		static_cast<void*>(&::GetVectorDataVector2),
		static_cast<void*>(&::GetVectorDataVector3),
		static_cast<void*>(&::GetVectorDataVector4),
		static_cast<void*>(&::GetVectorDataMatrix4x4),

		static_cast<void*>(&::AssignVectorBool),
		static_cast<void*>(&::AssignVectorChar8),
		static_cast<void*>(&::AssignVectorChar16),
		static_cast<void*>(&::AssignVectorInt8),
		static_cast<void*>(&::AssignVectorInt16),
		static_cast<void*>(&::AssignVectorInt32),
		static_cast<void*>(&::AssignVectorInt64),
		static_cast<void*>(&::AssignVectorUInt8),
		static_cast<void*>(&::AssignVectorUInt16),
		static_cast<void*>(&::AssignVectorUInt32),
		static_cast<void*>(&::AssignVectorUInt64),
		static_cast<void*>(&::AssignVectorPointer),
		static_cast<void*>(&::AssignVectorFloat),
		static_cast<void*>(&::AssignVectorDouble),
		static_cast<void*>(&::AssignVectorString),
		static_cast<void*>(&::AssignVectorVariant),
		static_cast<void*>(&::AssignVectorVector2),
		static_cast<void*>(&::AssignVectorVector3),
		static_cast<void*>(&::AssignVectorVector4),
		static_cast<void*>(&::AssignVectorMatrix4x4),

		static_cast<void*>(&::NewCall),
		static_cast<void*>(&::DeleteCall),
		static_cast<void*>(&::GetCallFunction),
		static_cast<void*>(&::GetCallError),

		static_cast<void*>(&::NewCallback),
		static_cast<void*>(&::DeleteCallback),
		static_cast<void*>(&::GetCallbackFunction),
		static_cast<void*>(&::GetCallbackError),

		static_cast<void*>(&::GetMethodParamCount),
		static_cast<void*>(&::GetMethodParamType),
		static_cast<void*>(&::GetMethodPrototype),
		static_cast<void*>(&::GetMethodEnum),
};

ILanguageModule* GetLanguageModule() {
	return &golm::g_golm;
}
