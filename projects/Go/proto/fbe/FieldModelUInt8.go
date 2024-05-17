//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

package fbe

import "errors"

// Fast Binary Encoding UInt8 field model
type FieldModelUInt8 struct {
    // Field model buffer
    buffer *Buffer
    // Field model buffer offset
    offset int
}

// Create a new UInt8 field model
func NewFieldModelUInt8(buffer *Buffer, offset int) *FieldModelUInt8 {
    return &FieldModelUInt8{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelUInt8) FBESize() int { return 1 }
// Get the field extra size
func (fm *FieldModelUInt8) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelUInt8) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelUInt8) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelUInt8) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelUInt8) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FieldModelUInt8) Verify() bool { return true }

// Get the value
func (fm *FieldModelUInt8) Get() (uint8, error) {
    return fm.GetDefault(0)
}

// Get the value with provided default value
func (fm *FieldModelUInt8) GetDefault(defaults uint8) (uint8, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return defaults, nil
    }

    return ReadUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), nil
}

// Set the value
func (fm *FieldModelUInt8) Set(value uint8) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    WriteUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return nil
}
