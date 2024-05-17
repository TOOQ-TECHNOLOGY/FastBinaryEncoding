//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

package test

import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding StructList field model
type FieldModelStructList struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int

    F1 *FieldModelVectorByte
    F2 *FieldModelVectorOptionalByte
    F3 *FieldModelVectorBytes
    F4 *FieldModelVectorOptionalBytes
    F5 *FieldModelVectorEnumSimple
    F6 *FieldModelVectorOptionalEnumSimple
    F7 *FieldModelVectorFlagsSimple
    F8 *FieldModelVectorOptionalFlagsSimple
    F9 *FieldModelVectorStructSimple
    F10 *FieldModelVectorOptionalStructSimple
}

// Create a new StructList field model
func NewFieldModelStructList(buffer *fbe.Buffer, offset int) *FieldModelStructList {
    fbeResult := FieldModelStructList{buffer: buffer, offset: offset}
    fbeResult.F1 = NewFieldModelVectorByte(buffer, 4 + 4)
    fbeResult.F2 = NewFieldModelVectorOptionalByte(buffer, fbeResult.F1.FBEOffset() + fbeResult.F1.FBESize())
    fbeResult.F3 = NewFieldModelVectorBytes(buffer, fbeResult.F2.FBEOffset() + fbeResult.F2.FBESize())
    fbeResult.F4 = NewFieldModelVectorOptionalBytes(buffer, fbeResult.F3.FBEOffset() + fbeResult.F3.FBESize())
    fbeResult.F5 = NewFieldModelVectorEnumSimple(buffer, fbeResult.F4.FBEOffset() + fbeResult.F4.FBESize())
    fbeResult.F6 = NewFieldModelVectorOptionalEnumSimple(buffer, fbeResult.F5.FBEOffset() + fbeResult.F5.FBESize())
    fbeResult.F7 = NewFieldModelVectorFlagsSimple(buffer, fbeResult.F6.FBEOffset() + fbeResult.F6.FBESize())
    fbeResult.F8 = NewFieldModelVectorOptionalFlagsSimple(buffer, fbeResult.F7.FBEOffset() + fbeResult.F7.FBESize())
    fbeResult.F9 = NewFieldModelVectorStructSimple(buffer, fbeResult.F8.FBEOffset() + fbeResult.F8.FBESize())
    fbeResult.F10 = NewFieldModelVectorOptionalStructSimple(buffer, fbeResult.F9.FBEOffset() + fbeResult.F9.FBESize())
    return &fbeResult
}

// Get the field size
func (fm *FieldModelStructList) FBESize() int { return 4 }

// Get the field body size
func (fm *FieldModelStructList) FBEBody() int {
    fbeResult := 4 + 4 +
        fm.F1.FBESize() +
        fm.F2.FBESize() +
        fm.F3.FBESize() +
        fm.F4.FBESize() +
        fm.F5.FBESize() +
        fm.F6.FBESize() +
        fm.F7.FBESize() +
        fm.F8.FBESize() +
        fm.F9.FBESize() +
        fm.F10.FBESize() +
        0
    return fbeResult
}

// Get the field extra size
func (fm *FieldModelStructList) FBEExtra() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4) > fm.buffer.Size()) {
        return 0
    }

    fm.buffer.Shift(fbeStructOffset)

    fbeResult := fm.FBEBody() +
        fm.F1.FBEExtra() +
        fm.F2.FBEExtra() +
        fm.F3.FBEExtra() +
        fm.F4.FBEExtra() +
        fm.F5.FBEExtra() +
        fm.F6.FBEExtra() +
        fm.F7.FBEExtra() +
        fm.F8.FBEExtra() +
        fm.F9.FBEExtra() +
        fm.F10.FBEExtra() +
        0

    fm.buffer.Unshift(fbeStructOffset)

    return fbeResult
}

// Get the field type
func (fm *FieldModelStructList) FBEType() int { return 131 }

// Get the field offset
func (fm *FieldModelStructList) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelStructList) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelStructList) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelStructList) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FieldModelStructList) Verify() bool { return fm.VerifyType(true) }

// Check if the struct value and its type are valid
func (fm *FieldModelStructList) VerifyType(fbeVerifyType bool) bool {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return true
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4 + 4) > fm.buffer.Size()) {
        return false
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset))
    if fbeStructSize < (4 + 4) {
        return false
    }

    fbeStructType := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset + 4))
    if fbeVerifyType && (fbeStructType != fm.FBEType()) {
        return false
    }

    fm.buffer.Shift(fbeStructOffset)
    fbeResult := fm.VerifyFields(fbeStructSize)
    fm.buffer.Unshift(fbeStructOffset)
    return fbeResult
}

// // Check if the struct value fields are valid
func (fm *FieldModelStructList) VerifyFields(fbeStructSize int) bool {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.F1.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.F1.Verify() {
        return false
    }
    fbeCurrentSize += fm.F1.FBESize()

    if (fbeCurrentSize + fm.F2.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.F2.Verify() {
        return false
    }
    fbeCurrentSize += fm.F2.FBESize()

    if (fbeCurrentSize + fm.F3.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.F3.Verify() {
        return false
    }
    fbeCurrentSize += fm.F3.FBESize()

    if (fbeCurrentSize + fm.F4.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.F4.Verify() {
        return false
    }
    fbeCurrentSize += fm.F4.FBESize()

    if (fbeCurrentSize + fm.F5.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.F5.Verify() {
        return false
    }
    fbeCurrentSize += fm.F5.FBESize()

    if (fbeCurrentSize + fm.F6.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.F6.Verify() {
        return false
    }
    fbeCurrentSize += fm.F6.FBESize()

    if (fbeCurrentSize + fm.F7.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.F7.Verify() {
        return false
    }
    fbeCurrentSize += fm.F7.FBESize()

    if (fbeCurrentSize + fm.F8.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.F8.Verify() {
        return false
    }
    fbeCurrentSize += fm.F8.FBESize()

    if (fbeCurrentSize + fm.F9.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.F9.Verify() {
        return false
    }
    fbeCurrentSize += fm.F9.FBESize()

    if (fbeCurrentSize + fm.F10.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.F10.Verify() {
        return false
    }
    fbeCurrentSize += fm.F10.FBESize()

    return true
}

// Get the struct value (begin phase)
func (fm *FieldModelStructList) GetBegin() (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, nil
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4 + 4) > fm.buffer.Size()) {
        return 0, errors.New("model is broken")
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset))
    if fbeStructSize < (4 + 4) {
        return 0, errors.New("model is broken")
    }

    fm.buffer.Shift(fbeStructOffset)
    return fbeStructOffset, nil
}

// Get the struct value (end phase)
func (fm *FieldModelStructList) GetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Get the struct value
func (fm *FieldModelStructList) Get() (*StructList, error) {
    fbeResult := NewStructList()
    return fbeResult, fm.GetValue(fbeResult)
}

// Get the struct value by the given pointer
func (fm *FieldModelStructList) GetValue(fbeValue *StructList) error {
    fbeBegin, err := fm.GetBegin()
    if fbeBegin == 0 {
        return err
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset()))
    fm.GetFields(fbeValue, fbeStructSize)
    fm.GetEnd(fbeBegin)
    return nil
}

// Get the struct fields values
func (fm *FieldModelStructList) GetFields(fbeValue *StructList, fbeStructSize int) {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.F1.FBESize()) <= fbeStructSize {
        fbeValue.F1, _ = fm.F1.Get()
    } else {
        fbeValue.F1 = make([]byte, 0)
    }
    fbeCurrentSize += fm.F1.FBESize()

    if (fbeCurrentSize + fm.F2.FBESize()) <= fbeStructSize {
        fbeValue.F2, _ = fm.F2.Get()
    } else {
        fbeValue.F2 = make([]*byte, 0)
    }
    fbeCurrentSize += fm.F2.FBESize()

    if (fbeCurrentSize + fm.F3.FBESize()) <= fbeStructSize {
        fbeValue.F3, _ = fm.F3.Get()
    } else {
        fbeValue.F3 = make([][]byte, 0)
    }
    fbeCurrentSize += fm.F3.FBESize()

    if (fbeCurrentSize + fm.F4.FBESize()) <= fbeStructSize {
        fbeValue.F4, _ = fm.F4.Get()
    } else {
        fbeValue.F4 = make([]*[]byte, 0)
    }
    fbeCurrentSize += fm.F4.FBESize()

    if (fbeCurrentSize + fm.F5.FBESize()) <= fbeStructSize {
        fbeValue.F5, _ = fm.F5.Get()
    } else {
        fbeValue.F5 = make([]EnumSimple, 0)
    }
    fbeCurrentSize += fm.F5.FBESize()

    if (fbeCurrentSize + fm.F6.FBESize()) <= fbeStructSize {
        fbeValue.F6, _ = fm.F6.Get()
    } else {
        fbeValue.F6 = make([]*EnumSimple, 0)
    }
    fbeCurrentSize += fm.F6.FBESize()

    if (fbeCurrentSize + fm.F7.FBESize()) <= fbeStructSize {
        fbeValue.F7, _ = fm.F7.Get()
    } else {
        fbeValue.F7 = make([]FlagsSimple, 0)
    }
    fbeCurrentSize += fm.F7.FBESize()

    if (fbeCurrentSize + fm.F8.FBESize()) <= fbeStructSize {
        fbeValue.F8, _ = fm.F8.Get()
    } else {
        fbeValue.F8 = make([]*FlagsSimple, 0)
    }
    fbeCurrentSize += fm.F8.FBESize()

    if (fbeCurrentSize + fm.F9.FBESize()) <= fbeStructSize {
        fbeValue.F9, _ = fm.F9.Get()
    } else {
        fbeValue.F9 = make([]StructSimple, 0)
    }
    fbeCurrentSize += fm.F9.FBESize()

    if (fbeCurrentSize + fm.F10.FBESize()) <= fbeStructSize {
        fbeValue.F10, _ = fm.F10.Get()
    } else {
        fbeValue.F10 = make([]*StructSimple, 0)
    }
    fbeCurrentSize += fm.F10.FBESize()
}

// Set the struct value (begin phase)
func (fm *FieldModelStructList) SetBegin() (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbeStructSize := fm.FBEBody()
    fbeStructOffset := fm.buffer.Allocate(fbeStructSize) - fm.buffer.Offset()
    if (fbeStructOffset <= 0) || ((fm.buffer.Offset() + fbeStructOffset + fbeStructSize) > fm.buffer.Size()) {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint32(fbeStructOffset))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset, uint32(fbeStructSize))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset + 4, uint32(fm.FBEType()))

    fm.buffer.Shift(fbeStructOffset)
    return fbeStructOffset, nil
}

// Set the struct value (end phase)
func (fm *FieldModelStructList) SetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Set the struct value
func (fm *FieldModelStructList) Set(fbeValue *StructList) error {
    fbeBegin, err := fm.SetBegin()
    if fbeBegin == 0 {
        return err
    }

    err = fm.SetFields(fbeValue)
    fm.SetEnd(fbeBegin)
    return err
}

// Set the struct fields values
func (fm *FieldModelStructList) SetFields(fbeValue *StructList) error {
    var err error = nil

    if err = fm.F1.Set(fbeValue.F1); err != nil {
        return err
    }
    if err = fm.F2.Set(fbeValue.F2); err != nil {
        return err
    }
    if err = fm.F3.Set(fbeValue.F3); err != nil {
        return err
    }
    if err = fm.F4.Set(fbeValue.F4); err != nil {
        return err
    }
    if err = fm.F5.Set(fbeValue.F5); err != nil {
        return err
    }
    if err = fm.F6.Set(fbeValue.F6); err != nil {
        return err
    }
    if err = fm.F7.Set(fbeValue.F7); err != nil {
        return err
    }
    if err = fm.F8.Set(fbeValue.F8); err != nil {
        return err
    }
    if err = fm.F9.Set(fbeValue.F9); err != nil {
        return err
    }
    if err = fm.F10.Set(fbeValue.F10); err != nil {
        return err
    }
    return err
}
