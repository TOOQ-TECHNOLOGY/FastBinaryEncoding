//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.test.fbe

// Fast Binary Encoding StructMap final model
@Suppress("MemberVisibilityCanBePrivate", "RemoveRedundantCallsOfConversionMethods", "ReplaceGetOrSet")
class FinalModelStructMap(buffer: com.chronoxor.fbe.Buffer, offset: Long) : com.chronoxor.fbe.FinalModel(buffer, offset)
{
    val f1: FinalModelMapInt32Byte = FinalModelMapInt32Byte(buffer, 0)
    val f2: FinalModelMapInt32OptionalByte = FinalModelMapInt32OptionalByte(buffer, 0)
    val f3: FinalModelMapInt32Bytes = FinalModelMapInt32Bytes(buffer, 0)
    val f4: FinalModelMapInt32OptionalBytes = FinalModelMapInt32OptionalBytes(buffer, 0)
    val f5: FinalModelMapInt32EnumSimple = FinalModelMapInt32EnumSimple(buffer, 0)
    val f6: FinalModelMapInt32OptionalEnumSimple = FinalModelMapInt32OptionalEnumSimple(buffer, 0)
    val f7: FinalModelMapInt32FlagsSimple = FinalModelMapInt32FlagsSimple(buffer, 0)
    val f8: FinalModelMapInt32OptionalFlagsSimple = FinalModelMapInt32OptionalFlagsSimple(buffer, 0)
    val f9: FinalModelMapInt32StructSimple = FinalModelMapInt32StructSimple(buffer, 0)
    val f10: FinalModelMapInt32OptionalStructSimple = FinalModelMapInt32OptionalStructSimple(buffer, 0)

    // Get the allocation size
    @Suppress("UNUSED_PARAMETER")
    fun fbeAllocationSize(fbeValue: com.chronoxor.test.StructMap): Long = (0
        + f1.fbeAllocationSize(fbeValue.f1)
        + f2.fbeAllocationSize(fbeValue.f2)
        + f3.fbeAllocationSize(fbeValue.f3)
        + f4.fbeAllocationSize(fbeValue.f4)
        + f5.fbeAllocationSize(fbeValue.f5)
        + f6.fbeAllocationSize(fbeValue.f6)
        + f7.fbeAllocationSize(fbeValue.f7)
        + f8.fbeAllocationSize(fbeValue.f8)
        + f9.fbeAllocationSize(fbeValue.f9)
        + f10.fbeAllocationSize(fbeValue.f10)
        )

    // Field type
    var fbeType: Long = fbeTypeConst

    companion object
    {
        const val fbeTypeConst: Long = 140
    }

    // Check if the struct value is valid
    override fun verify(): Long
    {
        _buffer.shift(fbeOffset)
        val fbeResult = verifyFields()
        _buffer.unshift(fbeOffset)
        return fbeResult
    }

    // Check if the struct fields are valid
    fun verifyFields(): Long
    {
        var fbeCurrentOffset = 0L
        @Suppress("VARIABLE_WITH_REDUNDANT_INITIALIZER")
        var fbeFieldSize = 0L

        f1.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f1.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f2.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f2.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f3.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f3.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f4.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f4.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f5.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f5.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f6.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f6.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f7.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f7.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f8.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f8.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f9.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f9.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f10.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f10.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        return fbeCurrentOffset
    }

    // Get the struct value
    fun get(fbeSize: com.chronoxor.fbe.Size, fbeValue: com.chronoxor.test.StructMap = com.chronoxor.test.StructMap()): com.chronoxor.test.StructMap
    {
        _buffer.shift(fbeOffset)
        fbeSize.value = getFields(fbeValue)
        _buffer.unshift(fbeOffset)
        return fbeValue
    }

    // Get the struct fields values
    @Suppress("UNUSED_PARAMETER")
    fun getFields(fbeValue: com.chronoxor.test.StructMap): Long
    {
        var fbeCurrentOffset = 0L
        var fbeCurrentSize = 0L
        val fbeFieldSize = com.chronoxor.fbe.Size()

        f1.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f1.get(fbeValue.f1)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f2.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f2.get(fbeValue.f2)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f3.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f3.get(fbeValue.f3)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f4.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f4.get(fbeValue.f4)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f5.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f5.get(fbeValue.f5)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f6.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f6.get(fbeValue.f6)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f7.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f7.get(fbeValue.f7)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f8.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f8.get(fbeValue.f8)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f9.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f9.get(fbeValue.f9)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f10.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f10.get(fbeValue.f10)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        return fbeCurrentSize
    }

    // Set the struct value
    fun set(fbeValue: com.chronoxor.test.StructMap): Long
    {
        _buffer.shift(fbeOffset)
        val fbeSize = setFields(fbeValue)
        _buffer.unshift(fbeOffset)
        return fbeSize
    }

    // Set the struct fields values
    @Suppress("UNUSED_PARAMETER")
    fun setFields(fbeValue: com.chronoxor.test.StructMap): Long
    {
        var fbeCurrentOffset = 0L
        var fbeCurrentSize = 0L
        val fbeFieldSize = com.chronoxor.fbe.Size()

        f1.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f1.set(fbeValue.f1)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f2.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f2.set(fbeValue.f2)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f3.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f3.set(fbeValue.f3)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f4.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f4.set(fbeValue.f4)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f5.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f5.set(fbeValue.f5)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f6.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f6.set(fbeValue.f6)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f7.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f7.set(fbeValue.f7)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f8.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f8.set(fbeValue.f8)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f9.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f9.set(fbeValue.f9)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f10.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f10.set(fbeValue.f10)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        return fbeCurrentSize
    }
}
