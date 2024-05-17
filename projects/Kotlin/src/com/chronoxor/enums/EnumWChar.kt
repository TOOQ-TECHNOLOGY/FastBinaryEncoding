//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.enums

@Suppress("MemberVisibilityCanBePrivate", "RemoveRedundantCallsOfConversionMethods")
class EnumWChar : Comparable<EnumWChar>
{
    companion object
    {
        val ENUM_VALUE_0 = EnumWChar(EnumWCharEnum.ENUM_VALUE_0)
        val ENUM_VALUE_1 = EnumWChar(EnumWCharEnum.ENUM_VALUE_1)
        val ENUM_VALUE_2 = EnumWChar(EnumWCharEnum.ENUM_VALUE_2)
        val ENUM_VALUE_3 = EnumWChar(EnumWCharEnum.ENUM_VALUE_3)
        val ENUM_VALUE_4 = EnumWChar(EnumWCharEnum.ENUM_VALUE_4)
        val ENUM_VALUE_5 = EnumWChar(EnumWCharEnum.ENUM_VALUE_5)
    }

    var value: EnumWCharEnum? = EnumWCharEnum.values()[0]
        private set

    val raw: Int
        get() = value!!.raw

    constructor()
    constructor(value: Int) { setEnum(value) }
    constructor(value: EnumWCharEnum) { setEnum(value) }
    constructor(value: EnumWChar) { setEnum(value) }

    fun setDefault() { setEnum(0.toInt()) }

    fun setEnum(value: Int) { this.value = EnumWCharEnum.mapValue(value) }
    fun setEnum(value: EnumWCharEnum) { this.value = value }
    fun setEnum(value: EnumWChar) { this.value = value.value }

    override fun compareTo(other: EnumWChar): Int
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

        if (!EnumWChar::class.java.isAssignableFrom(other.javaClass))
            return false

        val enm = other as EnumWChar? ?: return false

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
