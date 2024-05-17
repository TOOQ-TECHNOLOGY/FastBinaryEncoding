//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

package com.chronoxor.protoex.fbe;

// Fast Binary Encoding Balance final model
public final class FinalModelBalance extends com.chronoxor.fbe.FinalModel
{
    public final com.chronoxor.proto.fbe.FinalModelBalance parent;
    public final com.chronoxor.fbe.FinalModelDouble locked;

    public FinalModelBalance(com.chronoxor.fbe.Buffer buffer, long offset)
    {
        super(buffer, offset);
        parent = new com.chronoxor.proto.fbe.FinalModelBalance(buffer, 0);
        locked = new com.chronoxor.fbe.FinalModelDouble(buffer, 0);
    }

    // Get the allocation size
    public long fbeAllocationSize(com.chronoxor.protoex.Balance fbeValue)
    {
        long fbeResult = 0
            + parent.fbeAllocationSize(fbeValue)
            + locked.fbeAllocationSize(fbeValue.locked)
            ;
        return fbeResult;
    }

    // Get the final type
    public static final long fbeTypeConst = com.chronoxor.proto.fbe.FinalModelBalance.fbeTypeConst;
    public long fbeType() { return fbeTypeConst; }

    // Check if the struct value is valid
    @Override
    public long verify()
    {
        _buffer.shift(fbeOffset());
        long fbeResult = verifyFields();
        _buffer.unshift(fbeOffset());
        return fbeResult;
    }

    // Check if the struct fields are valid
    public long verifyFields()
    {
        long fbeCurrentOffset = 0;
        long fbeFieldSize = 0;

        parent.fbeOffset(fbeCurrentOffset);
        fbeFieldSize = parent.verifyFields();
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE;
        fbeCurrentOffset += fbeFieldSize;

        locked.fbeOffset(fbeCurrentOffset);
        fbeFieldSize = locked.verify();
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE;
        fbeCurrentOffset += fbeFieldSize;

        return fbeCurrentOffset;
    }

    // Get the struct value
    public com.chronoxor.protoex.Balance get(com.chronoxor.fbe.Size fbeSize) { return get(fbeSize, new com.chronoxor.protoex.Balance()); }
    public com.chronoxor.protoex.Balance get(com.chronoxor.fbe.Size fbeSize, com.chronoxor.protoex.Balance fbeValue)
    {
        _buffer.shift(fbeOffset());
        fbeSize.value = getFields(fbeValue);
        _buffer.unshift(fbeOffset());
        return fbeValue;
    }

    // Get the struct fields values
    public long getFields(com.chronoxor.protoex.Balance fbeValue)
    {
        long fbeCurrentOffset = 0;
        long fbeCurrentSize = 0;
        var fbeFieldSize = new com.chronoxor.fbe.Size(0);

        parent.fbeOffset(fbeCurrentOffset);
        fbeFieldSize.value = parent.getFields(fbeValue);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        locked.fbeOffset(fbeCurrentOffset);
        fbeValue.locked = locked.get(fbeFieldSize);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        return fbeCurrentSize;
    }

    // Set the struct value
    public long set(com.chronoxor.protoex.Balance fbeValue)
    {
        _buffer.shift(fbeOffset());
        long fbeSize = setFields(fbeValue);
        _buffer.unshift(fbeOffset());
        return fbeSize;
    }

    // Set the struct fields values
    public long setFields(com.chronoxor.protoex.Balance fbeValue)
    {
        long fbeCurrentOffset = 0;
        long fbeCurrentSize = 0;
        var fbeFieldSize = new com.chronoxor.fbe.Size();

        parent.fbeOffset(fbeCurrentOffset);
        fbeFieldSize.value = parent.setFields(fbeValue);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        locked.fbeOffset(fbeCurrentOffset);
        fbeFieldSize.value = locked.set(fbeValue.locked);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        return fbeCurrentSize;
    }
}
