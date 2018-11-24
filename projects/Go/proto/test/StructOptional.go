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

// StructOptional key
type StructOptionalKey struct {
    StructSimpleKey
}

// Convert StructOptional flags key to string
func (k StructOptionalKey) String() string {
    var sb strings.Builder
    return sb.String()
}

// StructOptional struct
type StructOptional struct {
    StructSimple
    F100 *bool
    F101 *bool
    F102 *bool
    F103 *byte
    F104 *byte
    F105 *byte
    F106 *rune
    F107 *rune
    F108 *rune
    F109 *rune
    F110 *rune
    F111 *rune
    F112 *int8
    F113 *int8
    F114 *int8
    F115 *uint8
    F116 *uint8
    F117 *uint8
    F118 *int16
    F119 *int16
    F120 *int16
    F121 *uint16
    F122 *uint16
    F123 *uint16
    F124 *int32
    F125 *int32
    F126 *int32
    F127 *uint32
    F128 *uint32
    F129 *uint32
    F130 *int64
    F131 *int64
    F132 *int64
    F133 *uint64
    F134 *uint64
    F135 *uint64
    F136 *float32
    F137 *float32
    F138 *float32
    F139 *float64
    F140 *float64
    F141 *float64
    F142 *fbe.Decimal
    F143 *fbe.Decimal
    F144 *fbe.Decimal
    F145 *string
    F146 *string
    F147 *string
    F148 *fbe.Timestamp
    F149 *fbe.Timestamp
    F150 *fbe.Timestamp
    F151 *fbe.UUID
    F152 *fbe.UUID
    F153 *fbe.UUID
    F154 *proto.OrderSide
    F155 *proto.OrderSide
    F156 *proto.OrderType
    F157 *proto.OrderType
    F158 *proto.Order
    F159 *proto.Order
    F160 *proto.Balance
    F161 *proto.Balance
    F162 *proto.State
    F163 *proto.State
    F164 *proto.Account
    F165 *proto.Account
}

// Create a new StructOptional struct from JSON
func NewStructOptionalFromJSON(buffer []byte) (*StructOptional, error) {
    var result StructOptional
    err := json.Unmarshal(buffer, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// Get the struct key
func (s StructOptional) Key() StructOptionalKey {
    return StructOptionalKey{
        s.StructSimple.Key(),
    }
}

// Convert struct to string
func (s StructOptional) String() string {
    var sb strings.Builder
    return sb.String()
}

// Convert struct to JSON
func (s StructOptional) JSON() ([]byte, error) {
    return json.Marshal(s)
}