// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.1.0.0

package protoex

import "strconv"
import "strings"
import "encoding/json"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Account key
type AccountKey struct {
    Uid int32
}

// Convert Account flags key to string
func (k AccountKey) String() string {
    var sb strings.Builder
    return sb.String()
}

// Account struct
type Account struct {
    Uid int32
    Name string
    State StateEx
    Wallet Balance
    Asset *Balance
    Orders []Order
}

// Create a new Account struct from JSON
func NewAccountFromJSON(buffer []byte) (*Account, error) {
    var result Account
    err := json.Unmarshal(buffer, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// Get the struct key
func (s Account) Key() AccountKey {
    return AccountKey{
        Uid: s.Uid,
    }
}

// Convert struct to string
func (s Account) String() string {
    var sb strings.Builder
    return sb.String()
}

// Convert struct to JSON
func (s Account) JSON() ([]byte, error) {
    return json.Marshal(s)
}