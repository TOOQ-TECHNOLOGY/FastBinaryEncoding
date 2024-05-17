//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

package protoex

import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding BalanceMessage model
type BalanceMessageModel struct {
    // Model buffer
    buffer *fbe.Buffer

    // Field model
    model *FieldModelBalanceMessage
}

// Create a new BalanceMessage model
func NewBalanceMessageModel(buffer *fbe.Buffer) *BalanceMessageModel {
    return &BalanceMessageModel{buffer: buffer, model: NewFieldModelBalanceMessage(buffer, 4)}
}

// Get the model buffer
func (m *BalanceMessageModel) Buffer() *fbe.Buffer { return m.buffer }
// Get the field model
func (m *BalanceMessageModel) Model() *FieldModelBalanceMessage { return m.model }

// Get the model size
func (m *BalanceMessageModel) FBESize() int { return m.model.FBESize() + m.model.FBEExtra() }
// // Get the model type
func (m *BalanceMessageModel) FBEType() int { return m.model.FBEType() }

// Check if the struct value is valid
func (m *BalanceMessageModel) Verify() bool {
    if (m.buffer.Offset() + m.model.FBEOffset() - 4) > m.buffer.Size() {
        return false
    }

    fbeFullSize := int(fbe.ReadUInt32(m.buffer.Data(), m.buffer.Offset() + m.model.FBEOffset() - 4))
    if fbeFullSize < m.model.FBESize() {
        return false
    }

    return m.model.Verify()
}

// Create a new model (begin phase)
func (m *BalanceMessageModel) CreateBegin() int {
    fbeBegin := m.buffer.Allocate(4 + m.model.FBESize())
    return fbeBegin
}

// Create a new model (end phase)
func (m *BalanceMessageModel) CreateEnd(fbeBegin int) int {
    fbeEnd := m.buffer.Size()
    fbeFullSize := fbeEnd - fbeBegin
    fbe.WriteUInt32(m.buffer.Data(), m.buffer.Offset() + m.model.FBEOffset() - 4, uint32(fbeFullSize))
    return fbeFullSize
}

// Serialize the struct value
func (m *BalanceMessageModel) Serialize(value *BalanceMessage) (int, error) {
    fbeBegin := m.CreateBegin()
    err := m.model.Set(value)
    fbeFullSize := m.CreateEnd(fbeBegin)
    return fbeFullSize, err
}

// Deserialize the struct value
func (m *BalanceMessageModel) Deserialize() (*BalanceMessage, int, error) {
    value := NewBalanceMessage()
    fbeFullSize, err := m.DeserializeValue(value)
    return value, fbeFullSize, err
}

// Deserialize the struct value by the given pointer
func (m *BalanceMessageModel) DeserializeValue(value *BalanceMessage) (int, error) {
    if (m.buffer.Offset() + m.model.FBEOffset() - 4) > m.buffer.Size() {
        value = NewBalanceMessage()
        return 0, nil
    }

    fbeFullSize := int(fbe.ReadUInt32(m.buffer.Data(), m.buffer.Offset() + m.model.FBEOffset() - 4))
    if fbeFullSize < m.model.FBESize() {
        value = NewBalanceMessage()
        return 0, errors.New("model is broken")
    }

    err := m.model.GetValue(value)
    return fbeFullSize, err
}

// Move to the next struct value
func (m *BalanceMessageModel) Next(prev int) {
    m.model.FBEShift(prev)
}
