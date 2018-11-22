// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test
import _ "../proto"

import "encoding/json"

type EnumSimple int32

const (
    EnumSimple_ENUM_VALUE_0 = EnumSimple(0) + 0
    EnumSimple_ENUM_VALUE_1 = EnumSimple(1) + 0
    EnumSimple_ENUM_VALUE_2 = EnumSimple(1) + 1
    EnumSimple_ENUM_VALUE_3 = EnumSimple(3) + 0
    EnumSimple_ENUM_VALUE_4 = EnumSimple(0x4) + 0
    EnumSimple_ENUM_VALUE_5 = EnumSimple_ENUM_VALUE_3
)

func (e EnumSimple) String() string {
    switch e {
    case EnumSimple_ENUM_VALUE_0:
        return "ENUM_VALUE_0"
    case EnumSimple_ENUM_VALUE_1:
        return "ENUM_VALUE_1"
    case EnumSimple_ENUM_VALUE_2:
        return "ENUM_VALUE_2"
    case EnumSimple_ENUM_VALUE_3:
        return "ENUM_VALUE_3"
    case EnumSimple_ENUM_VALUE_4:
        return "ENUM_VALUE_4"
    case EnumSimple_ENUM_VALUE_5:
        return "ENUM_VALUE_5"
    }
    return "<unknown>"
}

func (e EnumSimple) MarshalJSON() ([]byte, error) {
    return json.Marshal(int32(e))
}

func (e *EnumSimple) UnmarshalJSON(b []byte) error {
    var value int32
    err := json.Unmarshal(b, &value)
    if err != nil {
        return err
    }
    *e = EnumSimple(value)
    return nil
}
