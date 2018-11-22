// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "../fbe"

import "errors"

// Fast Binary Encoding EnumInt32 final model class
type FinalModelEnumInt32 struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset
}

// Get the allocation size
func (fm FinalModelEnumInt32) FBEAllocationSize(value EnumInt32) int { return fm.FBESize() }

// Get the final size
func (fm FinalModelEnumInt32) FBESize() int { return 4 }

// Get the final offset
func (fm FinalModelEnumInt32) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelEnumInt32) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelEnumInt32) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelEnumInt32) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelEnumInt32(buffer *fbe.Buffer, offset int) *FinalModelEnumInt32 {
    return &FinalModelEnumInt32{ buffer: buffer, offset: offset }
}

// Check if the value is valid
func (fm FinalModelEnumInt32) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm FinalModelEnumInt32) Get() (EnumInt32, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumInt32(0), 0, errors.New("model is broken")
    }

    return EnumInt32(fbe.ReadInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelEnumInt32) Set(value EnumInt32) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), int32(value))
    return fm.FBESize(), nil
}
