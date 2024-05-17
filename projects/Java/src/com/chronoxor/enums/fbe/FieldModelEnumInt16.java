//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

package com.chronoxor.enums.fbe;

// Fast Binary Encoding EnumInt16 field model
public final class FieldModelEnumInt16 extends com.chronoxor.fbe.FieldModel
{
    public FieldModelEnumInt16(com.chronoxor.fbe.Buffer buffer, long offset) { super(buffer, offset); }

    // Get the field size
    @Override
    public long fbeSize() { return 2; }

    // Get the value
    public com.chronoxor.enums.EnumInt16 get() { return get(new com.chronoxor.enums.EnumInt16()); }
    public com.chronoxor.enums.EnumInt16 get(com.chronoxor.enums.EnumInt16 defaults)
    {
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return defaults;

        return new com.chronoxor.enums.EnumInt16(readInt16(fbeOffset()));
    }

    // Set the value
    public void set(com.chronoxor.enums.EnumInt16 value)
    {
        assert ((_buffer.getOffset() + fbeOffset() + fbeSize()) <= _buffer.getSize()) : "Model is broken!";
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return;

        write(fbeOffset(), value.getRaw());
    }
}
