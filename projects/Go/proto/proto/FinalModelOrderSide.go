// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.1.0.0

package proto

import "../fbe"

import "errors"

// Fast Binary Encoding OrderSide final model class
type FinalModelOrderSide struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset
}

// Get the allocation size
func (fm FinalModelOrderSide) FBEAllocationSize(value OrderSide) int { return fm.FBESize() }

// Get the final size
func (fm FinalModelOrderSide) FBESize() int { return 1 }

// Get the final offset
func (fm FinalModelOrderSide) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelOrderSide) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelOrderSide) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelOrderSide) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelOrderSide(buffer *fbe.Buffer, offset int) *FinalModelOrderSide {
    return &FinalModelOrderSide{ buffer: buffer, offset: offset }
}

// Check if the value is valid
func (fm FinalModelOrderSide) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm FinalModelOrderSide) Get() (OrderSide, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return OrderSide(0), 0, errors.New("model is broken")
    }

    return OrderSide(fbe.ReadByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelOrderSide) Set(value OrderSide) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), byte(value))
    return fm.FBESize(), nil
}
