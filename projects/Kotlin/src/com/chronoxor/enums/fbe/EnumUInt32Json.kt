//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.enums.fbe

class EnumUInt32Json : com.google.gson.JsonSerializer<com.chronoxor.enums.EnumUInt32>, com.google.gson.JsonDeserializer<com.chronoxor.enums.EnumUInt32>
{
    override fun serialize(src: com.chronoxor.enums.EnumUInt32, typeOfSrc: java.lang.reflect.Type, context: com.google.gson.JsonSerializationContext): com.google.gson.JsonElement
    {
        return com.google.gson.JsonPrimitive(src.raw.toInt())
    }

    @Throws(com.google.gson.JsonParseException::class)
    override fun deserialize(json: com.google.gson.JsonElement, type: java.lang.reflect.Type, context: com.google.gson.JsonDeserializationContext): com.chronoxor.enums.EnumUInt32
    {
        return com.chronoxor.enums.EnumUInt32(json.asJsonPrimitive.asInt.toUInt())
    }
}
