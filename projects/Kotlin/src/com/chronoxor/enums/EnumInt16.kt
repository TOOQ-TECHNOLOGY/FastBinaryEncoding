//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.enums

@Suppress("MemberVisibilityCanBePrivate", "RemoveRedundantCallsOfConversionMethods")
class EnumInt16 : Comparable<EnumInt16>
{
    companion object
    {
        val ENUM_VALUE_0 = EnumInt16(EnumInt16Enum.ENUM_VALUE_0)
        val ENUM_VALUE_1 = EnumInt16(EnumInt16Enum.ENUM_VALUE_1)
        val ENUM_VALUE_2 = EnumInt16(EnumInt16Enum.ENUM_VALUE_2)
        val ENUM_VALUE_3 = EnumInt16(EnumInt16Enum.ENUM_VALUE_3)
        val ENUM_VALUE_4 = EnumInt16(EnumInt16Enum.ENUM_VALUE_4)
        val ENUM_VALUE_5 = EnumInt16(EnumInt16Enum.ENUM_VALUE_5)
    }

    var value: EnumInt16Enum? = EnumInt16Enum.values()[0]
        private set

    val raw: Short
        get() = value!!.raw

    constructor()
    constructor(value: Short) { setEnum(value) }
    constructor(value: EnumInt16Enum) { setEnum(value) }
    constructor(value: EnumInt16) { setEnum(value) }

    fun setDefault() { setEnum(0.toShort()) }

    fun setEnum(value: Short) { this.value = EnumInt16Enum.mapValue(value) }
    fun setEnum(value: EnumInt16Enum) { this.value = value }
    fun setEnum(value: EnumInt16) { this.value = value.value }

    override fun compareTo(other: EnumInt16): Int
    {
        if (value == null)
            return -1
        if (other.value == null)
            return 1
        return (value!!.raw - other.value!!.raw).toInt()
    }

    override fun equals(other: Any?): Boolean
    {
        if (other == null)
            return false

        if (!EnumInt16::class.java.isAssignableFrom(other.javaClass))
            return false

        val enm = other as EnumInt16? ?: return false

        if (enm.value == null)
            return false
        if (value!!.raw != enm.value!!.raw)
            return false
        return true
    }

    override fun hashCode(): Int
    {
        var hash = 17
        hash = hash * 31 + if (value != null) value!!.hashCode() else 0
        return hash
    }

    override fun toString(): String
    {
        return if (value != null) value!!.toString() else "<unknown>"
    }
}
