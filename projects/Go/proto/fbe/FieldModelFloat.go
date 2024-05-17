//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

package fbe

import "errors"

// Fast Binary Encoding Float field model
type FieldModelFloat struct {
    // Field model buffer
    buffer *Buffer
    // Field model buffer offset
    offset int
}

// Create a new Float field model
func NewFieldModelFloat(buffer *Buffer, offset int) *FieldModelFloat {
    return &FieldModelFloat{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelFloat) FBESize() int { return 4 }
// Get the field extra size
func (fm *FieldModelFloat) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelFloat) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelFloat) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelFloat) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelFloat) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FieldModelFloat) Verify() bool { return true }

// Get the value
func (fm *FieldModelFloat) Get() (float32, error) {
    return fm.GetDefault(0.0)
}

// Get the value with provided default value
func (fm *FieldModelFloat) GetDefault(defaults float32) (float32, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return defaults, nil
    }

    return ReadFloat(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), nil
}

// Set the value
func (fm *FieldModelFloat) Set(value float32) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    WriteFloat(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return nil
}
