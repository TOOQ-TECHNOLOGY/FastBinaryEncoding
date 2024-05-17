//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

#pragma once

#if defined(__clang__)
#pragma clang system_header
#elif defined(__GNUC__)
#pragma GCC system_header
#elif defined(_MSC_VER)
#pragma system_header
#endif

#include "fbe.h"

namespace FBE {

// Fast Binary Encoding base final model
template <typename T, typename TBase = T>
class FinalModelBase
{
public:
    FinalModelBase(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset) {}

    // Get the allocation size
    size_t fbe_allocation_size(T value) const noexcept { return fbe_size(); }

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Set the field offset
    size_t fbe_offset(size_t offset) const noexcept { return _offset = offset; }

    // Get the final size
    size_t fbe_size() const noexcept { return sizeof(TBase); }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the value is valid
    size_t verify() const noexcept;

    // Get the field value
    size_t get(T& value) const noexcept;
    // Set the field value
    size_t set(T value) noexcept;

private:
    FBEBuffer& _buffer;
    mutable size_t _offset;
};

// Fast Binary Encoding final model
template <typename T>
class FinalModel : public FinalModelBase<T>
{
public:
    using FinalModelBase<T>::FinalModelBase;
};

// Fast Binary Encoding final model bool specialization
template <>
class FinalModel<bool> : public FinalModelBase<bool, uint8_t>
{
public:
    using FinalModelBase<bool, uint8_t>::FinalModelBase;
};

// Fast Binary Encoding final model char specialization
template <>
class FinalModel<char> : public FinalModelBase<char, uint8_t>
{
public:
    using FinalModelBase<char, uint8_t>::FinalModelBase;
};

// Fast Binary Encoding final model wchar specialization
template <>
class FinalModel<wchar_t> : public FinalModelBase<wchar_t, uint32_t>
{
public:
    using FinalModelBase<wchar_t, uint32_t>::FinalModelBase;
};

// Fast Binary Encoding final model decimal specialization
template <>
class FinalModel<decimal_t>
{
public:
    FinalModel(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset) {}

    // Get the allocation size
    size_t fbe_allocation_size(decimal_t value) const noexcept { return fbe_size(); }

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Set the field offset
    size_t fbe_offset(size_t offset) const noexcept { return _offset = offset; }

    // Get the final size
    size_t fbe_size() const noexcept { return 16; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the decimal value is valid
    size_t verify() const noexcept;

    // Get the decimal value
    size_t get(decimal_t& value) const noexcept;
    // Set the decimal value
    size_t set(decimal_t value) noexcept;

private:
    FBEBuffer& _buffer;
    mutable size_t _offset;

    static uint64_t extract(double a) noexcept;
    static uint64_t uint32x32(uint32_t a, uint32_t b) noexcept;
    static void uint64x64(uint64_t a, uint64_t b, uint64_t& low64, uint32_t& high32) noexcept;
};

// Fast Binary Encoding final model UUID specialization
template <>
class FinalModel<uuid_t>
{
public:
    FinalModel(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset) {}

    // Get the allocation size
    size_t fbe_allocation_size(uuid_t value) const noexcept { return fbe_size(); }

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Set the field offset
    size_t fbe_offset(size_t offset) const noexcept { return _offset = offset; }

    // Get the final size
    size_t fbe_size() const noexcept { return 16; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the UUID value is valid
    size_t verify() const noexcept;

    // Get the UUID value
    size_t get(uuid_t& value) const noexcept;
    // Set the UUID value
    size_t set(uuid_t value) noexcept;

private:
    FBEBuffer& _buffer;
    mutable size_t _offset;
};

// Fast Binary Encoding final model bytes specialization
template <>
class FinalModel<buffer_t>
{
public:
    FinalModel(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset) {}

    // Get the allocation size
    size_t fbe_allocation_size(const void* data, size_t size) const noexcept { return 4 + size; }
    template <size_t N>
    size_t fbe_allocation_size(const uint8_t (&data)[N]) const noexcept { return 4 + N; }
    template <size_t N>
    size_t fbe_allocation_size(const std::array<uint8_t, N>& data) const noexcept { return 4 + N; }
    size_t fbe_allocation_size(const std::vector<uint8_t>& value) const noexcept { return 4 + value.size(); }
    size_t fbe_allocation_size(const buffer_t& value) const noexcept { return 4 + value.size(); }

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Set the field offset
    size_t fbe_offset(size_t offset) const noexcept { return _offset = offset; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the bytes value is valid
    size_t verify() const noexcept;

    // Get the bytes value
    size_t get(void* data, size_t size) const noexcept;
    // Get the bytes value
    template <size_t N>
    size_t get(uint8_t (&data)[N]) const noexcept { return get(data, N); }
    // Get the bytes value
    template <size_t N>
    size_t get(std::array<uint8_t, N>& data) const noexcept { return get(data.data(), data.size()); }
    // Get the bytes value
    size_t get(std::vector<uint8_t>& value) const noexcept;
    // Get the bytes value
    size_t get(buffer_t& value) const noexcept { return get(value.buffer()); }

    // Set the bytes value
    size_t set(const void* data, size_t size);
    // Set the bytes value
    template <size_t N>
    size_t set(const uint8_t (&data)[N]) { return set(data, N); }
    // Set the bytes value
    template <size_t N>
    size_t set(const std::array<uint8_t, N>& data) { return set(data.data(), data.size()); }
    // Set the bytes value
    size_t set(const std::vector<uint8_t>& value) { return set(value.data(), value.size()); }
    // Set the bytes value
    size_t set(const buffer_t& value) { return set(value.buffer()); }

private:
    FBEBuffer& _buffer;
    mutable size_t _offset;
};

// Fast Binary Encoding final model string specialization
template <>
class FinalModel<std::string>
{
public:
    FinalModel(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset) {}

    // Get the allocation size
    size_t fbe_allocation_size(const char* data, size_t size) const noexcept { return 4 + size; }
    template <size_t N>
    size_t fbe_allocation_size(const char (&data)[N]) const noexcept { return 4 + N; }
    template <size_t N>
    size_t fbe_allocation_size(const std::array<char, N>& data) const noexcept { return 4 + N; }
    size_t fbe_allocation_size(const std::string& value) const noexcept { return 4 + value.size(); }

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Set the field offset
    size_t fbe_offset(size_t offset) const noexcept { return _offset = offset; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the string value is valid
    size_t verify() const noexcept;

    // Get the string value
    size_t get(char* data, size_t size) const noexcept;
    // Get the string value
    template <size_t N>
    size_t get(char (&data)[N]) const noexcept { return get(data, N); }
    // Get the string value
    template <size_t N>
    size_t get(std::array<char, N>& data) const noexcept { return get(data.data(), data.size()); }
    // Get the string value
    size_t get(std::string& value) const noexcept;

    // Set the string value
    size_t set(const char* data, size_t size);
    // Set the string value
    template <size_t N>
    size_t set(const char (&data)[N]) { return set(data, N); }
    // Set the string value
    template <size_t N>
    size_t set(const std::array<char, N>& data) { return set(data.data(), data.size()); }
    // Set the string value
    size_t set(const std::string& value);

private:
    FBEBuffer& _buffer;
    mutable size_t _offset;
};

// Fast Binary Encoding final model optional specialization
template <typename T>
class FinalModel<std::optional<T>>
{
public:
    FinalModel(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset), value(buffer, 0) {}

    // Get the allocation size
    size_t fbe_allocation_size(const std::optional<T>& opt) const noexcept { return 1 + (opt ? value.fbe_allocation_size(opt.value()) : 0); }

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Set the field offset
    size_t fbe_offset(size_t offset) const noexcept { return _offset = offset; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    //! Is the value present?
    explicit operator bool() const noexcept { return has_value(); }

    // Checks if the object contains a value
    bool has_value() const noexcept;

    // Check if the optional value is valid
    size_t verify() const noexcept;

    // Get the optional value
    size_t get(std::optional<T>& opt) const noexcept;
    // Set the optional value
    size_t set(const std::optional<T>& opt);

private:
    FBEBuffer& _buffer;
    mutable size_t _offset;

public:
    // Base final model value
    FinalModel<T> value;
};

// Fast Binary Encoding final model array
template <typename T, size_t N>
class FinalModelArray
{
public:
    FinalModelArray(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset) {}

    // Get the allocation size
    template <size_t S>
    size_t fbe_allocation_size(const T (&values)[S]) const noexcept;
    template <size_t S>
    size_t fbe_allocation_size(const std::array<T, S>& values) const noexcept;
    size_t fbe_allocation_size(const std::vector<T>& values) const noexcept;

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Set the field offset
    size_t fbe_offset(size_t offset) const noexcept { return _offset = offset; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the array is valid
    size_t verify() const noexcept;

    // Get the array as C-array
    template <size_t S>
    size_t get(T (&values)[S]) const noexcept;
    // Get the array as std::array
    template <size_t S>
    size_t get(std::array<T, S>& values) const noexcept;
    // Get the array as std::vector
    size_t get(std::vector<T>& values) const noexcept;

    // Set the array as C-array
    template <size_t S>
    size_t set(const T (&values)[S]) noexcept;
    // Set the array as std::array
    template <size_t S>
    size_t set(const std::array<T, S>& values) noexcept;
    // Set the array as std::vector
    size_t set(const std::vector<T>& values) noexcept;

private:
    FBEBuffer& _buffer;
    mutable size_t _offset;
};

// Fast Binary Encoding final model vector
template <typename T>
class FinalModelVector
{
public:
    FinalModelVector(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset) {}

    // Get the allocation size
    size_t fbe_allocation_size(const std::vector<T>& values) const noexcept;
    size_t fbe_allocation_size(const std::list<T>& values) const noexcept;
    size_t fbe_allocation_size(const std::set<T>& values) const noexcept;

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Set the field offset
    size_t fbe_offset(size_t offset) const noexcept { return _offset = offset; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the vector is valid
    size_t verify() const noexcept;

    // Get the vector as std::vector
    size_t get(std::vector<T>& values) const noexcept;
    // Get the vector as std::list
    size_t get(std::list<T>& values) const noexcept;
    // Get the vector as std::set
    size_t get(std::set<T>& values) const noexcept;

    // Set the vector as std::vector
    size_t set(const std::vector<T>& values) noexcept;
    // Set the vector as std::list
    size_t set(const std::list<T>& values) noexcept;
    // Set the vector as std::set
    size_t set(const std::set<T>& values) noexcept;

private:
    FBEBuffer& _buffer;
    mutable size_t _offset;
};

// Fast Binary Encoding final model map
template <typename TKey, typename TValue>
class FinalModelMap
{
public:
    FinalModelMap(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset) {}

    // Get the allocation size
    size_t fbe_allocation_size(const std::map<TKey, TValue>& values) const noexcept;
    size_t fbe_allocation_size(const std::unordered_map<TKey, TValue>& values) const noexcept;

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Set the field offset
    size_t fbe_offset(size_t offset) const noexcept { return _offset = offset; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the map is valid
    size_t verify() const noexcept;

    // Get the map as std::map
    size_t get(std::map<TKey, TValue>& values) const noexcept;
    // Get the map as std::unordered_map
    size_t get(std::unordered_map<TKey, TValue>& values) const noexcept;

    // Set the map as std::map
    size_t set(const std::map<TKey, TValue>& values) noexcept;
    // Set the map as std::unordered_map
    size_t set(const std::unordered_map<TKey, TValue>& values) noexcept;

private:
    FBEBuffer& _buffer;
    mutable size_t _offset;
};

} // namespace FBE

#include "fbe_final_models.inl"
