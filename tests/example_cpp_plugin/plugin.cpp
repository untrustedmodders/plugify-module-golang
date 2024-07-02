#include <plugify/cpp_plugin.h>
#include <plugin_export.h>
#include <test/test.h>
#include <pps/example_golang_plugin.h>
#include <cassert>

class ExampleCppPlugin : public plugify::IPluginEntry {
public:
    void OnPluginStart() override {
        // No Params, Only Return
        {
            std::cout << "=================================C++=======================================" << std::endl;

            example_golang_plugin::NoParamReturnVoid();
            assert(example_golang_plugin::NoParamReturnBool() == true);
            assert(example_golang_plugin::NoParamReturnChar8() == std::numeric_limits<char>::max());
            assert(example_golang_plugin::NoParamReturnChar16() == std::numeric_limits<char16_t>::max());
            assert(example_golang_plugin::NoParamReturnInt8() == std::numeric_limits<int8_t>::max());
            assert(example_golang_plugin::NoParamReturnInt16() == std::numeric_limits<int16_t>::max());
            assert(example_golang_plugin::NoParamReturnInt32() == std::numeric_limits<int32_t>::max());
            assert(example_golang_plugin::NoParamReturnInt64() == std::numeric_limits<int64_t>::max());
            assert(example_golang_plugin::NoParamReturnUInt8() == std::numeric_limits<uint8_t>::max());
            assert(example_golang_plugin::NoParamReturnUInt16() == std::numeric_limits<uint16_t>::max());
            assert(example_golang_plugin::NoParamReturnUInt32() == std::numeric_limits<uint32_t>::max());
            assert(example_golang_plugin::NoParamReturnUInt64() == std::numeric_limits<uint64_t>::max());
            assert(example_golang_plugin::NoParamReturnPtr64() == reinterpret_cast<void*>(0x1));
            assert(example_golang_plugin::NoParamReturnFloat() == std::numeric_limits<float>::max());
            assert(example_golang_plugin::NoParamReturnDouble() == std::numeric_limits<double>::max());
            assert(example_golang_plugin::NoParamReturnFunction() == nullptr);
            assert(example_golang_plugin::NoParamReturnString() == "Hello World");
            assert((example_golang_plugin::NoParamReturnArrayBool() == std::vector<bool>{true, false}));
            assert((example_golang_plugin::NoParamReturnArrayChar16() == std::vector<char16_t>{u'a', u'b', u'c', u'd'}));
            assert((example_golang_plugin::NoParamReturnArrayInt8() == std::vector<int8_t>{-3, -2, -1, 0, 1}));
            assert((example_golang_plugin::NoParamReturnArrayInt16() == std::vector<int16_t>{-4, -3, -2, -1, 0, 1}));
            assert((example_golang_plugin::NoParamReturnArrayInt32() == std::vector<int32_t>{-5, -4, -3, -2, -1, 0, 1}));
            assert((example_golang_plugin::NoParamReturnArrayInt64() == std::vector<int64_t>{-6, -5, -4, -3, -2, -1, 0, 1}));
            assert((example_golang_plugin::NoParamReturnArrayUInt8() == std::vector<uint8_t>{0, 1, 2, 3, 4, 5, 6, 7, 8}));
            assert((example_golang_plugin::NoParamReturnArrayUInt16() == std::vector<uint16_t>{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}));
            assert((example_golang_plugin::NoParamReturnArrayUInt32() == std::vector<uint32_t>{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}));
            assert((example_golang_plugin::NoParamReturnArrayUInt64() == std::vector<uint64_t>{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}));
            assert((example_golang_plugin::NoParamReturnArrayPtr64() == std::vector<void*>{reinterpret_cast<void*>(0), reinterpret_cast<void*>(1), reinterpret_cast<void*>(2), reinterpret_cast<void*>(3)}));
            assert((example_golang_plugin::NoParamReturnArrayFloat() == std::vector<float>{-12.34f, 0.0f, 12.34f}));
            assert((example_golang_plugin::NoParamReturnArrayDouble() == std::vector<double>{-12.345, 0.0, 12.345}));
            assert((example_golang_plugin::NoParamReturnArrayString() == std::vector<std::string>{"1st string", "2nd string", "3rd element string (Should be big enough to avoid small string optimization)"}));
            assert((example_golang_plugin::NoParamReturnVector2() == plugify::Vector2(1, 2)));
            assert((example_golang_plugin::NoParamReturnVector3() == plugify::Vector3(1, 2, 3)));
            assert((example_golang_plugin::NoParamReturnVector4() == plugify::Vector4(1, 2, 3, 4)));
            assert((example_golang_plugin::NoParamReturnMatrix4x4() == plugify::Matrix4x4({1, 5, 9, 13, 2, 6, 10, 14, 3, 7, 11, 15, 4, 8, 12, 16})));
        }

        int32_t intValue = 42;
        float floatValue = 3.14f;
        double doubleValue = 6.28;
        plugify::Vector4 vector4Value(1.0f, 2.0f, 3.0f, 4.0f);
        std::vector<int64_t> longListValue = { 100, 200, 300 };
        char16_t charValue = 'A';
        std::string stringValue = "Hello";
        float floatValue2 = 2.71f;
        int16_t shortValue = 10;
        void* ptrValue = nullptr; // Provide appropriate value for pointer

        // Params (no refs)
        {
            example_golang_plugin::Param1(intValue);
            example_golang_plugin::Param2(intValue, floatValue);
            example_golang_plugin::Param3(intValue, floatValue, doubleValue);
            example_golang_plugin::Param4(intValue, floatValue, doubleValue, vector4Value);
            example_golang_plugin::Param5(intValue, floatValue, doubleValue, vector4Value, longListValue);
            example_golang_plugin::Param6(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue);
            example_golang_plugin::Param7(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue);
            example_golang_plugin::Param8(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue, floatValue2);
            example_golang_plugin::Param9(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue, floatValue2, shortValue);
            example_golang_plugin::Param10(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue, floatValue2, shortValue, ptrValue);
        }

        // Params (with refs)
        {
            example_golang_plugin::ParamRef1(intValue);
            assert((intValue == 42));

            example_golang_plugin::ParamRef2(intValue, floatValue);
            assert((intValue == 10));
            assert((floatValue == 3.14f));

            example_golang_plugin::ParamRef3(intValue, floatValue, doubleValue);
            assert((intValue == -20));
            assert((floatValue == 2.718f));
            assert((doubleValue == 3.14159));

            example_golang_plugin::ParamRef4(intValue, floatValue, doubleValue, vector4Value);
            assert((intValue == 100));
            assert((floatValue == -5.55f));
            assert((doubleValue == 1.618));
            assert((vector4Value.x == 1.0f && vector4Value.y == 2.0f && vector4Value.z == 3.0f && vector4Value.w == 4.0f));

            example_golang_plugin::ParamRef5(intValue, floatValue, doubleValue, vector4Value, longListValue);
            assert((intValue == 500));
            assert((floatValue == -10.5f));
            assert((doubleValue == 2.71828));
            assert((vector4Value.x == -1.0f && vector4Value.y == -2.0f && vector4Value.z == -3.0f && vector4Value.w == -4.0f));
            assert((longListValue == std::vector<int64_t>{ -6, -5, -4, -3, -2, -1, 0, 1 }));

            example_golang_plugin::ParamRef6(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue);
            assert((intValue == 750));
            assert((floatValue == 20.0f));
            assert((doubleValue == 1.23456));
            assert((vector4Value.x == 10.0f && vector4Value.y == 20.0f && vector4Value.z == 30.0f && vector4Value.w == 40.0f));
            assert((longListValue == std::vector<int64_t>{ -6, -5, -4 }));
            assert((charValue == 'Z'));

            example_golang_plugin::ParamRef7(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue);
            assert((intValue == -1000));
            assert((floatValue == 3.0f));
            assert((doubleValue == -1));
            assert((vector4Value.x == 100.0f && vector4Value.y == 200.0f && vector4Value.z == 300.0f && vector4Value.w == 400.0f));
            assert((longListValue == std::vector<int64_t>{ -6, -5, -4, -3 }));
            assert((charValue == 'X'));
            assert((stringValue == "Hello, World!"));

            example_golang_plugin::ParamRef8(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue, floatValue2);
            assert((intValue == 999));
            assert((floatValue == -7.5f));
            assert((doubleValue == 0.123456));
            assert((vector4Value.x == -100.0f && vector4Value.y == -200.0f && vector4Value.z == -300.0f && vector4Value.w == -400.0f));
            assert((longListValue == std::vector<int64_t>{ -6, -5, -4, -3, -2, -1 }));
            assert((charValue == 'Y'));
            assert((stringValue == "Goodbye, World!"));
            assert((floatValue2 == 99.99f));

            example_golang_plugin::ParamRef9(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue, floatValue2, shortValue);
            assert((intValue == -1234));
            assert((floatValue == 123.45f));
            assert((doubleValue == -678.9));
            assert((vector4Value.x == 987.65f && vector4Value.y == 432.1f && vector4Value.z == 123.456f && vector4Value.w == 789.123f));
            assert((longListValue == std::vector<int64_t>{ -6, -5, -4, -3, -2, -1, 0, 1, 5, 9 }));
            assert((charValue == 'A'));
            assert((stringValue == "Testing, 1 2 3"));
            assert((floatValue2 == -987.654f));
            assert((shortValue == 42));

            example_golang_plugin::ParamRef10(intValue, floatValue, doubleValue, vector4Value, longListValue, charValue, stringValue, floatValue2, shortValue, ptrValue);
            assert((intValue == 987));
            assert((floatValue == -0.123f));
            assert((doubleValue == 456.789));
            assert((vector4Value.x == -123.456f && vector4Value.y == 0.987f && vector4Value.z == 654.321f && vector4Value.w == -789.123f));
            assert((longListValue == std::vector<int64_t>{ -6, -5, -4, -3, -2, -1, 0, 1, 5, 9, 4, -7 }));
            assert((charValue == 'B'));
            assert((stringValue == "Another string"));
            assert((floatValue2 == 3.141592f));
            assert((ptrValue == nullptr));
        }

        // Initialize arrays
        std::vector<bool> boolArray = { true, false, true };
        std::vector<char> char8Array = { 'A', 'B', 'C' };
        std::vector<char16_t> char16Array = { 'A', 'B', 'C' };
        std::vector<int8_t> sbyteArray = { -1, -2, -3 };
        std::vector<int16_t> shortArray = { 10, 20, 30 };
        std::vector<int32_t> intArray = { 100, 200, 300 };
        std::vector<int64_t> longArray = { 1000, 2000, 3000 };
        std::vector<uint8_t> byteArray = { 1, 2, 3 };
        std::vector<uint16_t> ushortArray = { 1000, 2000, 3000 };
        std::vector<uint32_t> uintArray = { 10000, 20000, 30000 };
        std::vector<uint64_t> ulongArray = { 100000, 200000, 300000 };
        std::vector<void*> intPtrArray = { nullptr, nullptr, nullptr };
        std::vector<float> floatArray = { 1.1f, 2.2f, 3.3f };
        std::vector<double> doubleArray = { 1.1, 2.2, 3.3 };
        std::vector<std::string> stringArray = { "Hello", "World", "!" };

        {
            // Call function ParamRefVectors and print how values change
            example_golang_plugin::ParamRefVectors(boolArray, char8Array, char16Array, sbyteArray, shortArray, intArray, longArray,
                            byteArray, ushortArray, uintArray, ulongArray, intPtrArray, floatArray, doubleArray, stringArray);

            assert((boolArray == std::vector<bool>{ true }));
            assert((char8Array == std::vector<char>{ 'a', 'b', 'c' }));
            assert((char16Array == std::vector<char16_t>{ 'a', 'b', 'c' }));
            assert((sbyteArray == std::vector<int8_t>{ -3, -2, -1, 0, 1, 2, 3 }));
            assert((shortArray == std::vector<int16_t>{ -4, -3, -2, -1, 0, 1, 2, 3, 4 }));
            assert((intArray == std::vector<int32_t>{ -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5 }));
            assert((longArray == std::vector<int64_t>{ -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6 }));
            assert((byteArray == std::vector<uint8_t>{ 0, 1, 2, 3, 4, 5, 6, 7 }));
            assert((ushortArray == std::vector<uint16_t>{ 0, 1, 2, 3, 4, 5, 6, 7, 8 }));
            assert((uintArray == std::vector<uint32_t>{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }));
            assert((ulongArray == std::vector<uint64_t>{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }));
            assert((intPtrArray == std::vector<void*>{ nullptr, (void*)1, (void*)2 }));
            assert((floatArray == std::vector<float>{ -12.34f, 0.0f, 12.34f }));
            assert((doubleArray == std::vector<double>{ -12.345, 0.0, 12.345 }));
            assert((stringArray == std::vector<std::string>{ "1", "12", "123", "1234", "12345", "123456" }));

            // Call function ParamAllPrimitives and print the returned value
            int64_t returnValue = example_golang_plugin::ParamAllPrimitives(boolArray[0], char16Array[0], sbyteArray[0], shortArray[0], intArray[0], longArray[0],
                                                     byteArray[0], ushortArray[0], uintArray[0], ulongArray[0], intPtrArray[0], floatArray[0], doubleArray[0]);

            assert((returnValue == 56));
        }
    }
};

ExampleCppPlugin g_examplePlugin;
EXPOSE_PLUGIN(PLUGIN_API, &g_examplePlugin)