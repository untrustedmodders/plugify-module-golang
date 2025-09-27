#include "module.hpp"

#include <plugify/logger.hpp>
#include <plugify/provider.hpp>
#include <plugify/call.hpp>

#include <plg/numerics.hpp>
#include <plg/any.hpp>
#include <plg/string.hpp>
#include <plg/numerics.hpp>

#include <plugify/assembly_loader.hpp>

#if __has_include(<stacktrace>)
#include <stacktrace>
#define HAS_STACKTRACE 1
#else
#define HAS_STACKTRACE 0
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

Result<InitData> GoLanguageModule::Initialize(const Provider& provider, [[maybe_unused]] const Extension& module) {
	_provider = std::make_unique<Provider>(provider);

	_provider->Log(LOG_PREFIX "Inited!", Severity::Debug);

	return InitData{{ .hasUpdate = false }};
}

void GoLanguageModule::Shutdown() {
	for (MemAddr* addr : _addresses) {
		*addr = nullptr;
	}
	_nativesMap.clear();
	_addresses.clear();
	_assemblies.clear();
	_provider.reset();
}

void GoLanguageModule::OnUpdate([[maybe_unused]] std::chrono::milliseconds dt) {
}

bool GoLanguageModule::IsDebugBuild() {
	return GOLM_IS_DEBUG;
}

void GoLanguageModule::OnMethodExport(const Extension& plugin) {
	const auto& methods = plugin.GetMethodsData();
	_nativesMap.reserve(_nativesMap.size() + methods.size());
	for (const auto& [method, addr] :methods) {
		_nativesMap.try_emplace(std::format("{}.{}", plugin.GetName(), method.GetName()), addr);
	}
}

Result<LoadData> GoLanguageModule::OnPluginLoad(const Extension& plugin) {
	fs::path assemblyPath(plugin.GetLocation());
	assemblyPath /= std::format("{}" GOLM_LIBRARY_SUFFIX, plugin.GetEntry());

	LoadFlag flags = LoadFlag::LazyBinding | LoadFlag::NoUnload;
	auto assemblyResult = _provider->Resolve<IAssemblyLoader>()->Load(assemblyPath, flags);
	if (!assemblyResult) {
		return MakeError(std::move(assemblyResult.error()));
	}

	auto& assembly = *assemblyResult;

	auto initResult = assembly->GetSymbol("Plugify_Init");
	if (!initResult) {
		return MakeError(std::move(initResult.error()));
	}
	auto callResult = assembly->GetSymbol("Plugify_InternalCall");
	if (!callResult) {
		return MakeError(std::move(callResult.error()));
	}
	auto startResult = assembly->GetSymbol("Plugify_PluginStart");
	if (!startResult) {
		return MakeError(std::move(startResult.error()));
	}
	auto updateResult = assembly->GetSymbol("Plugify_PluginUpdate");
	if (!updateResult) {
		return MakeError(std::move(updateResult.error()));
	}
	auto endResult = assembly->GetSymbol("Plugify_PluginEnd");
	if (!endResult) {
		return MakeError(std::move(endResult.error()));
	}
	auto contextResult = assembly->GetSymbol("Plugify_PluginContext");
	if (!contextResult) {
		return MakeError(std::move(contextResult.error()));
	}

	auto* initFunc = initResult->RCast<InitFunc>();
	auto* callFunc = callResult->RCast<CallFunc>();
	auto* startFunc = startResult->RCast<StartFunc>();
	auto* updateFunc = updateResult->RCast<UpdateFunc>();
	auto* endFunc = endResult->RCast<EndFunc>();
	auto* contextFunc = contextResult->RCast<ContextFunc>();

	std::vector<std::string> errors;

	const std::vector<Method>& exportedMethods = plugin.GetMethods();
	std::vector<MethodData> methods;
	methods.reserve(exportedMethods.size());

	for (size_t i = 0; i < exportedMethods.size(); ++i) {
		const auto& method = exportedMethods[i];
		if (auto funcResult = assembly->GetSymbol(method.GetFuncName())) {
			methods.emplace_back(method, *funcResult);
		} else {
			errors.emplace_back(std::format("{:>3}. {} {}", i + 1, method.GetName(), funcResult.error()));
			if (constexpr size_t kMaxDisplay = 100; errors.size() >= kMaxDisplay) {
				errors.emplace_back(std::format("... and {} more", exportedMethods.size() - kMaxDisplay));
				break;
			}
		}
	}
	if (!errors.empty()) {
		return MakeError("Invalid methods:\n{}", plg::join(errors, "\n"));
	}

	GoSlice api { const_cast<void**>(_pluginApi.data()), _pluginApi.size(), _pluginApi.size() };
	const int resultVersion = initFunc(api, kApiVersion, static_cast<const void *>(&plugin));
	if (resultVersion != 0) {
		return MakeError("Not supported plugin api {}, max supported {}", resultVersion, kApiVersion);
	}

	const auto& [hasUpdate, hasStart, hasEnd, _] = contextFunc ? *(contextFunc()) : PluginContext{};

	auto data = _assemblies.emplace_back(std::make_unique<AssemblyHolder>(std::move(assembly), updateFunc, startFunc, endFunc, contextFunc, callFunc)).get();
	return LoadData{ std::move(methods), data, { hasUpdate, hasStart, hasEnd, !exportedMethods.empty() } };
}

void GoLanguageModule::OnPluginStart(const Extension& plugin) {
	plugin.GetUserData().RCast<AssemblyHolder*>()->startFunc();
}

void GoLanguageModule::OnPluginUpdate(const Extension& plugin, std::chrono::milliseconds dt) {
	plugin.GetUserData().RCast<AssemblyHolder*>()->updateFunc(std::chrono::duration<float>(dt).count());
}

void GoLanguageModule::OnPluginEnd(const Extension& plugin) {
	plugin.GetUserData().RCast<AssemblyHolder*>()->endFunc();
}

MemAddr GoLanguageModule::GetNativeMethod(std::string_view methodName) const {
	if (const auto it = _nativesMap.find(methodName); it != _nativesMap.end()) {
		return std::get<MemAddr>(*it);
	}
	_provider->Log(std::format(LOG_PREFIX "GetNativeMethod failed to find: '{}'", methodName), Severity::Fatal);
	return nullptr;
}

void GoLanguageModule::GetNativeMethod(std::string_view methodName, MemAddr* addressDest) {
	if (const auto it = _nativesMap.find(methodName); it != _nativesMap.end()) {
		*addressDest = std::get<MemAddr>(*it);
		_addresses.emplace_back(addressDest);
		return;
	}
	_provider->Log(std::format(LOG_PREFIX "GetNativeMethod failed to find: '{}'", methodName), Severity::Fatal);
}

const Method* GoLanguageModule::FindMethod(std::string_view name) {
	if (auto separated = Split(name, "."); separated.size() == 2) {
		if (auto plugin = _provider->FindExtension(separated[0])) {
			for (const auto& method : plugin->GetMethods()) {
				if (auto prototype = method.FindPrototype(separated[1])) {
					return prototype;
				}
			}
		}
	}
	_provider->Log(std::format(LOG_PREFIX "FindMethod failed to find: '{}'", name), Severity::Error);
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

bool IsExtensionLoaded(GoString name, GoString constraint) {
	if (constraint) {
		plg::range_set<> range;
		plg::parse(constraint, range);
		return g_golm.GetProvider()->IsExtensionLoaded(name, std::move(range));
	}
	else
		return g_golm.GetProvider()->IsExtensionLoaded(name);
}

const char* GetBaseDir() {
	const auto& source = plg::as_string(g_golm.GetProvider()->GetBaseDir());
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.data(), size);
	return dest;
}

const char* GetExtensionsDir() {
	const auto& source = plg::as_string(g_golm.GetProvider()->GetExtensionsDir());
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.data(), size);
	return dest;
}

const char* GetConfigsDir() {
	const auto& source = plg::as_string(g_golm.GetProvider()->GetConfigsDir());
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.data(), size);
	return dest;
}

const char* GetDataDir() {
	const auto& source = plg::as_string(g_golm.GetProvider()->GetDataDir());
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.data(), size);
	return dest;
}

const char* GetLogsDir() {
	const auto& source = plg::as_string(g_golm.GetProvider()->GetLogsDir());
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.data(), size);
	return dest;
}

const char* GetCacheDir() {
	const auto& source = plg::as_string(g_golm.GetProvider()->GetCacheDir());
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.data(), size);
	return dest;
}

void PrintException(GoString message) {
	if (const auto& provider = g_golm.GetProvider()) {
		provider->Log(std::format(LOG_PREFIX "[Exception] {}", std::string_view(message)), Severity::Error);
#if HAS_STACKTRACE
		auto trace = std::stacktrace::current();
		provider->Log(std::to_string(trace), Severity::Error);
#endif
	}
}

ptrdiff_t GetPluginId(const Extension& plugin) {
	return static_cast<ptrdiff_t>(plugin.GetId());
}

const char* GetPluginName(const Extension& plugin) {
	return plugin.GetName().c_str();
}

const char* GetPluginDescription(const Extension& plugin) {
	return plugin.GetDescription().c_str();
}

const char* GetPluginVersion(const Extension& plugin) {
	return plugin.GetVersionString().c_str();
}

const char* GetPluginAuthor(const Extension& plugin) {
	return plugin.GetAuthor().c_str();
}

const char* GetPluginWebsite(const Extension& plugin) {
	return plugin.GetWebsite().c_str();
}

const char* GetPluginLicense(const Extension& plugin) {
	return plugin.GetLicense().c_str();
}

const char* GetPluginLocation(const Extension& plugin) {
	const auto& source = plg::as_string(plugin.GetLocation());
	size_t size = source.length() + 1;
	char* dest = new char[size];
	std::memcpy(dest, source.data(), size);
	return dest;
}

const char** GetPluginDependencies(const Extension& plugin) {
	const std::vector<Dependency>& dependencies = plugin.GetDependencies();
	auto* deps = new const char*[dependencies.size()];
	for (size_t i = 0; i < dependencies.size(); ++i) {
		deps[i] = dependencies[i].GetName().c_str();
	}
	return deps;
}

ptrdiff_t GetPluginDependenciesSize(const Extension& plugin) {
	return static_cast<ptrdiff_t>(plugin.GetDependencies().size());
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
	ValueType type{};
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
	Signature sig(CallConv::CDecl, retHidden ? typeHidden : ret.type);

#if !GOLM_ARCH_ARM
	if (retHidden) {
		sig.AddArg(ret.type);
	}
#endif

	for (ptrdiff_t i = 0; i < count; ++i) {
		const auto& [type, ref] = params[i];
		sig.AddArg(ref ? ValueType::Pointer : type);
	}

	JitCall* call = new JitCall{};
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

JitCallback* NewCallback(const Extension& plugin, GoString name, void* delegate) {
	const Method* method = g_golm.FindMethod(name);
	if (method == nullptr || delegate == nullptr)
		return nullptr;

	JitCallback* callback = new JitCallback{};
	callback->GetJitFunc(*method, plugin.GetUserData().RCast<AssemblyHolder*>()->callFunc, delegate);
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

ptrdiff_t GetMethodParamCount(const Method& handle) {
	return static_cast<ptrdiff_t>(handle.GetParamTypes().size());
}

ManagedType GetMethodParamType(const Method& handle, ptrdiff_t index) {
	const Property& param = index < 0 ? handle.GetRetType() : handle.GetParamTypes()[static_cast<size_t>(index)];
	return { param.GetType(), param.IsRef() };
}

const Method& GetMethodPrototype(const Method& handle, ptrdiff_t index) {
	if (index < 0) {
		return *handle.GetRetType().GetPrototype();
	} else {
		return *handle.GetParamTypes()[static_cast<size_t>(index)].GetPrototype();
	}
}

const EnumObject& GetMethodEnum(const Method& handle, ptrdiff_t index) {
	if (index < 0) {
		return *handle.GetRetType().GetEnumerate();
	} else {
		return *handle.GetParamTypes()[static_cast<size_t>(index)].GetEnumerate();
	}
}

const std::array<void*, 140> GoLanguageModule::_pluginApi = {
		reinterpret_cast<void*>(&::GetMethodPtr),
		reinterpret_cast<void*>(&::GetMethodPtr2),
		reinterpret_cast<void*>(&::GetBaseDir),
		reinterpret_cast<void*>(&::GetExtensionsDir),
		reinterpret_cast<void*>(&::GetConfigsDir),
		reinterpret_cast<void*>(&::GetDataDir),
		reinterpret_cast<void*>(&::GetLogsDir),
		reinterpret_cast<void*>(&::GetCacheDir),
		reinterpret_cast<void*>(&::IsExtensionLoaded),
		reinterpret_cast<void*>(&::PrintException),

		reinterpret_cast<void*>(&::GetPluginId),
		reinterpret_cast<void*>(&::GetPluginName),
		reinterpret_cast<void*>(&::GetPluginDescription),
		reinterpret_cast<void*>(&::GetPluginVersion),
		reinterpret_cast<void*>(&::GetPluginAuthor),
		reinterpret_cast<void*>(&::GetPluginWebsite),
		reinterpret_cast<void*>(&::GetPluginLicense),
		reinterpret_cast<void*>(&::GetPluginLocation),
		reinterpret_cast<void*>(&::GetPluginDependencies),
		reinterpret_cast<void*>(&::GetPluginDependenciesSize),

		reinterpret_cast<void*>(&::DeleteCStr),
		reinterpret_cast<void*>(&::DeleteCStrArr),

		reinterpret_cast<void*>(&::ConstructString),
		reinterpret_cast<void*>(&::DestroyString),
		reinterpret_cast<void*>(&::GetStringData),
		reinterpret_cast<void*>(&::GetStringLength),
		reinterpret_cast<void*>(&::AssignString),

		reinterpret_cast<void*>(&::DestroyVariant),

		reinterpret_cast<void*>(&::ConstructVectorBool),
		reinterpret_cast<void*>(&::ConstructVectorChar8),
		reinterpret_cast<void*>(&::ConstructVectorChar16),
		reinterpret_cast<void*>(&::ConstructVectorInt8),
		reinterpret_cast<void*>(&::ConstructVectorInt16),
		reinterpret_cast<void*>(&::ConstructVectorInt32),
		reinterpret_cast<void*>(&::ConstructVectorInt64),
		reinterpret_cast<void*>(&::ConstructVectorUInt8),
		reinterpret_cast<void*>(&::ConstructVectorUInt16),
		reinterpret_cast<void*>(&::ConstructVectorUInt32),
		reinterpret_cast<void*>(&::ConstructVectorUInt64),
		reinterpret_cast<void*>(&::ConstructVectorPointer),
		reinterpret_cast<void*>(&::ConstructVectorFloat),
		reinterpret_cast<void*>(&::ConstructVectorDouble),
		reinterpret_cast<void*>(&::ConstructVectorString),
		reinterpret_cast<void*>(&::ConstructVectorVariant),
		reinterpret_cast<void*>(&::ConstructVectorVector2),
		reinterpret_cast<void*>(&::ConstructVectorVector3),
		reinterpret_cast<void*>(&::ConstructVectorVector4),
		reinterpret_cast<void*>(&::ConstructVectorMatrix4x4),

		reinterpret_cast<void*>(&::DestroyVectorBool),
		reinterpret_cast<void*>(&::DestroyVectorChar8),
		reinterpret_cast<void*>(&::DestroyVectorChar16),
		reinterpret_cast<void*>(&::DestroyVectorInt8),
		reinterpret_cast<void*>(&::DestroyVectorInt16),
		reinterpret_cast<void*>(&::DestroyVectorInt32),
		reinterpret_cast<void*>(&::DestroyVectorInt64),
		reinterpret_cast<void*>(&::DestroyVectorUInt8),
		reinterpret_cast<void*>(&::DestroyVectorUInt16),
		reinterpret_cast<void*>(&::DestroyVectorUInt32),
		reinterpret_cast<void*>(&::DestroyVectorUInt64),
		reinterpret_cast<void*>(&::DestroyVectorPointer),
		reinterpret_cast<void*>(&::DestroyVectorFloat),
		reinterpret_cast<void*>(&::DestroyVectorDouble),
		reinterpret_cast<void*>(&::DestroyVectorString),
		reinterpret_cast<void*>(&::DestroyVectorVariant),
		reinterpret_cast<void*>(&::DestroyVectorVector2),
		reinterpret_cast<void*>(&::DestroyVectorVector3),
		reinterpret_cast<void*>(&::DestroyVectorVector4),
		reinterpret_cast<void*>(&::DestroyVectorMatrix4x4),

		reinterpret_cast<void*>(&::GetVectorSizeBool),
		reinterpret_cast<void*>(&::GetVectorSizeChar8),
		reinterpret_cast<void*>(&::GetVectorSizeChar16),
		reinterpret_cast<void*>(&::GetVectorSizeInt8),
		reinterpret_cast<void*>(&::GetVectorSizeInt16),
		reinterpret_cast<void*>(&::GetVectorSizeInt32),
		reinterpret_cast<void*>(&::GetVectorSizeInt64),
		reinterpret_cast<void*>(&::GetVectorSizeUInt8),
		reinterpret_cast<void*>(&::GetVectorSizeUInt16),
		reinterpret_cast<void*>(&::GetVectorSizeUInt32),
		reinterpret_cast<void*>(&::GetVectorSizeUInt64),
		reinterpret_cast<void*>(&::GetVectorSizePointer),
		reinterpret_cast<void*>(&::GetVectorSizeFloat),
		reinterpret_cast<void*>(&::GetVectorSizeDouble),
		reinterpret_cast<void*>(&::GetVectorSizeString),
		reinterpret_cast<void*>(&::GetVectorSizeVariant),
		reinterpret_cast<void*>(&::GetVectorSizeVector2),
		reinterpret_cast<void*>(&::GetVectorSizeVector3),
		reinterpret_cast<void*>(&::GetVectorSizeVector4),
		reinterpret_cast<void*>(&::GetVectorSizeMatrix4x4),

		reinterpret_cast<void*>(&::GetVectorDataBool),
		reinterpret_cast<void*>(&::GetVectorDataChar8),
		reinterpret_cast<void*>(&::GetVectorDataChar16),
		reinterpret_cast<void*>(&::GetVectorDataInt8),
		reinterpret_cast<void*>(&::GetVectorDataInt16),
		reinterpret_cast<void*>(&::GetVectorDataInt32),
		reinterpret_cast<void*>(&::GetVectorDataInt64),
		reinterpret_cast<void*>(&::GetVectorDataUInt8),
		reinterpret_cast<void*>(&::GetVectorDataUInt16),
		reinterpret_cast<void*>(&::GetVectorDataUInt32),
		reinterpret_cast<void*>(&::GetVectorDataUInt64),
		reinterpret_cast<void*>(&::GetVectorDataPointer),
		reinterpret_cast<void*>(&::GetVectorDataFloat),
		reinterpret_cast<void*>(&::GetVectorDataDouble),
		reinterpret_cast<void*>(&::GetVectorDataString),
		reinterpret_cast<void*>(&::GetVectorDataVariant),
		reinterpret_cast<void*>(&::GetVectorDataVector2),
		reinterpret_cast<void*>(&::GetVectorDataVector3),
		reinterpret_cast<void*>(&::GetVectorDataVector4),
		reinterpret_cast<void*>(&::GetVectorDataMatrix4x4),

		reinterpret_cast<void*>(&::AssignVectorBool),
		reinterpret_cast<void*>(&::AssignVectorChar8),
		reinterpret_cast<void*>(&::AssignVectorChar16),
		reinterpret_cast<void*>(&::AssignVectorInt8),
		reinterpret_cast<void*>(&::AssignVectorInt16),
		reinterpret_cast<void*>(&::AssignVectorInt32),
		reinterpret_cast<void*>(&::AssignVectorInt64),
		reinterpret_cast<void*>(&::AssignVectorUInt8),
		reinterpret_cast<void*>(&::AssignVectorUInt16),
		reinterpret_cast<void*>(&::AssignVectorUInt32),
		reinterpret_cast<void*>(&::AssignVectorUInt64),
		reinterpret_cast<void*>(&::AssignVectorPointer),
		reinterpret_cast<void*>(&::AssignVectorFloat),
		reinterpret_cast<void*>(&::AssignVectorDouble),
		reinterpret_cast<void*>(&::AssignVectorString),
		reinterpret_cast<void*>(&::AssignVectorVariant),
		reinterpret_cast<void*>(&::AssignVectorVector2),
		reinterpret_cast<void*>(&::AssignVectorVector3),
		reinterpret_cast<void*>(&::AssignVectorVector4),
		reinterpret_cast<void*>(&::AssignVectorMatrix4x4),

		reinterpret_cast<void*>(&::NewCall),
		reinterpret_cast<void*>(&::DeleteCall),
		reinterpret_cast<void*>(&::GetCallFunction),
		reinterpret_cast<void*>(&::GetCallError),

		reinterpret_cast<void*>(&::NewCallback),
		reinterpret_cast<void*>(&::DeleteCallback),
		reinterpret_cast<void*>(&::GetCallbackFunction),
		reinterpret_cast<void*>(&::GetCallbackError),

		reinterpret_cast<void*>(&::GetMethodParamCount),
		reinterpret_cast<void*>(&::GetMethodParamType),
		reinterpret_cast<void*>(&::GetMethodPrototype),
		reinterpret_cast<void*>(&::GetMethodEnum),
};

ILanguageModule* GetLanguageModule() {
	return &golm::g_golm;
}
