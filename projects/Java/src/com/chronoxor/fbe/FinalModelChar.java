//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

package com.chronoxor.fbe;

// Fast Binary Encoding char final model
public final class FinalModelChar extends FinalModel
{
    public FinalModelChar(Buffer buffer, long offset) { super(buffer, offset); }

    // Get the allocation size
    public long fbeAllocationSize(char value) { return fbeSize(); }

    // Get the final size
    @Override
    public long fbeSize() { return 1; }

    // Check if the value is valid
    @Override
    public long verify()
    {
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return Long.MAX_VALUE;

        return fbeSize();
    }

    // Get the value
    public char get(Size size)
    {
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return '\u0000';

        size.value = fbeSize();
        return readChar(fbeOffset());
    }

    // Set the value
    public long set(char value)
    {
        assert ((_buffer.getOffset() + fbeOffset() + fbeSize()) <= _buffer.getSize()) : "Model is broken!";
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return 0;

        write(fbeOffset(), (byte)value);
        return fbeSize();
    }
}
