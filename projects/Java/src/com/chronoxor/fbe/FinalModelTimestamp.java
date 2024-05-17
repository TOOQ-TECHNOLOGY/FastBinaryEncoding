//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

package com.chronoxor.fbe;

// Fast Binary Encoding timestamp final model
public final class FinalModelTimestamp extends FinalModel
{
    public FinalModelTimestamp(Buffer buffer, long offset) { super(buffer, offset); }

    // Get the allocation size
    public long fbeAllocationSize(java.time.Instant value) { return fbeSize(); }

    // Get the final size
    @Override
    public long fbeSize() { return 8; }

    // Check if the timestamp value is valid
    @Override
    public long verify()
    {
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return Long.MAX_VALUE;

        return fbeSize();
    }

    // Get the timestamp value
    public java.time.Instant get(Size size)
    {
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return java.time.Instant.EPOCH;

        size.value = fbeSize();
        long nanoseconds = readInt64(fbeOffset());
        return java.time.Instant.ofEpochSecond(nanoseconds / 1000000000, nanoseconds % 1000000000);
    }

    // Set the timestamp value
    public long set(java.time.Instant value)
    {
        assert ((_buffer.getOffset() + fbeOffset() + fbeSize()) <= _buffer.getSize()) : "Model is broken!";
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return 0;

        long nanoseconds = value.getEpochSecond() * 1000000000 + value.getNano();
        write(fbeOffset(), nanoseconds);
        return fbeSize();
    }
}
