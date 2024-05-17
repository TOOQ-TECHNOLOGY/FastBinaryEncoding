//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.enums.fbe

// Fast Binary Encoding EnumInt16 field model
class FieldModelEnumInt16(buffer: com.chronoxor.fbe.Buffer, offset: Long) : com.chronoxor.fbe.FieldModel(buffer, offset)
{
    // Field size
    override val fbeSize: Long = 2

    // Get the value
    fun get(defaults: com.chronoxor.enums.EnumInt16 = com.chronoxor.enums.EnumInt16()): com.chronoxor.enums.EnumInt16
    {
        if ((_buffer.offset + fbeOffset + fbeSize) > _buffer.size)
            return defaults

        return com.chronoxor.enums.EnumInt16(readInt16(fbeOffset))
    }

    // Set the value
    fun set(value: com.chronoxor.enums.EnumInt16)
    {
        assert((_buffer.offset + fbeOffset + fbeSize) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeOffset + fbeSize) > _buffer.size)
            return

        write(fbeOffset, value.raw)
    }
}
