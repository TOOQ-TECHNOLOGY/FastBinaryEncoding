// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test

import "strconv"
import "strings"
import "encoding/json"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// StructBytes key
type StructBytesKey struct {
}

// Convert StructBytes flags key to string
func (k StructBytesKey) String() string {
    var sb strings.Builder
    return sb.String()
}

// StructBytes struct
type StructBytes struct {
    F1 []byte
    F2 *[]byte
    F3 *[]byte
}

// Create a new StructBytes struct from JSON
func NewStructBytesFromJSON(buffer []byte) (*StructBytes, error) {
    var result StructBytes
    err := json.Unmarshal(buffer, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// Get the struct key
func (s StructBytes) Key() StructBytesKey {
    return StructBytesKey{
    }
}

// Convert struct to string
func (s StructBytes) String() string {
    var sb strings.Builder
    return sb.String()
}

// Convert struct to JSON
func (s StructBytes) JSON() ([]byte, error) {
    return json.Marshal(s)
}