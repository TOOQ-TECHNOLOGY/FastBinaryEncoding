// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.2.0.0

package proto

import "errors"
import "../fbe"

// Workaround for Go unused imports issue
var _ = fbe.Version

// Proxy Order interface
type OnProxyOrder interface {
    OnProxyOrder(model *OrderModel, fbeType int, buffer []byte)
}

// Proxy Order function
type OnProxyOrderFunc func(model *OrderModel, fbeType int, buffer []byte)
func (f OnProxyOrderFunc) OnProxyOrder(model *OrderModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Proxy Balance interface
type OnProxyBalance interface {
    OnProxyBalance(model *BalanceModel, fbeType int, buffer []byte)
}

// Proxy Balance function
type OnProxyBalanceFunc func(model *BalanceModel, fbeType int, buffer []byte)
func (f OnProxyBalanceFunc) OnProxyBalance(model *BalanceModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Proxy Account interface
type OnProxyAccount interface {
    OnProxyAccount(model *AccountModel, fbeType int, buffer []byte)
}

// Proxy Account function
type OnProxyAccountFunc func(model *AccountModel, fbeType int, buffer []byte)
func (f OnProxyAccountFunc) OnProxyAccount(model *AccountModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Fast Binary Encoding proto proxy
type Proxy struct {
    *fbe.Receiver
    orderModel *OrderModel
    balanceModel *BalanceModel
    accountModel *AccountModel

    // Proxy Order handler
    HandlerOnProxyOrder OnProxyOrder
    // Proxy Balance handler
    HandlerOnProxyBalance OnProxyBalance
    // Proxy Account handler
    HandlerOnProxyAccount OnProxyAccount
}

// Create a new proto proxy with an empty buffer
func NewProxy() *Proxy {
    return NewProxyWithBuffer(fbe.NewEmptyBuffer())
}

// Create a new proto proxy with the given buffer
func NewProxyWithBuffer(buffer *fbe.Buffer) *Proxy {
    proxy := &Proxy{
        fbe.NewReceiver(buffer, false),
        NewOrderModel(buffer),
        NewBalanceModel(buffer),
        NewAccountModel(buffer),
        nil,
        nil,
        nil,
    }
    proxy.SetupHandlerOnReceive(proxy)
    proxy.SetupHandlerOnProxyOrderFunc(func(model *OrderModel, fbeType int, buffer []byte) {})
    proxy.SetupHandlerOnProxyBalanceFunc(func(model *BalanceModel, fbeType int, buffer []byte) {})
    proxy.SetupHandlerOnProxyAccountFunc(func(model *AccountModel, fbeType int, buffer []byte) {})
    return proxy
}

// Setup handlers
func (p *Proxy) SetupHandlers(handlers interface{}) {
    p.Receiver.SetupHandlers(handlers)
    if handler, ok := handlers.(OnProxyOrder); ok {
        p.SetupHandlerOnProxyOrder(handler)
    }
    if handler, ok := handlers.(OnProxyBalance); ok {
        p.SetupHandlerOnProxyBalance(handler)
    }
    if handler, ok := handlers.(OnProxyAccount); ok {
        p.SetupHandlerOnProxyAccount(handler)
    }
}

// Setup proxy Order handler
func (p *Proxy) SetupHandlerOnProxyOrder(handler OnProxyOrder) { p.HandlerOnProxyOrder = handler }
// Setup proxy Order handler function
func (p *Proxy) SetupHandlerOnProxyOrderFunc(function func(model *OrderModel, fbeType int, buffer []byte)) { p.HandlerOnProxyOrder = OnProxyOrderFunc(function) }
// Setup proxy Balance handler
func (p *Proxy) SetupHandlerOnProxyBalance(handler OnProxyBalance) { p.HandlerOnProxyBalance = handler }
// Setup proxy Balance handler function
func (p *Proxy) SetupHandlerOnProxyBalanceFunc(function func(model *BalanceModel, fbeType int, buffer []byte)) { p.HandlerOnProxyBalance = OnProxyBalanceFunc(function) }
// Setup proxy Account handler
func (p *Proxy) SetupHandlerOnProxyAccount(handler OnProxyAccount) { p.HandlerOnProxyAccount = handler }
// Setup proxy Account handler function
func (p *Proxy) SetupHandlerOnProxyAccountFunc(function func(model *AccountModel, fbeType int, buffer []byte)) { p.HandlerOnProxyAccount = OnProxyAccountFunc(function) }

// Receive message handler
func (p *Proxy) OnReceive(fbeType int, buffer []byte) (bool, error) {
    switch fbeType {
    case p.orderModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.orderModel.Buffer().Attach(buffer)
        if !p.orderModel.Verify() {
            return false, errors.New("proto.Order validation failed")
        }

        // Call proxy handler
        fbeBegin, err := p.orderModel.model.GetBegin()
        if fbeBegin == 0 {
            return false, err
        }
        p.HandlerOnProxyOrder.OnProxyOrder(p.orderModel, fbeType, buffer)
        p.orderModel.model.GetEnd(fbeBegin)
        return true, nil
    case p.balanceModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.balanceModel.Buffer().Attach(buffer)
        if !p.balanceModel.Verify() {
            return false, errors.New("proto.Balance validation failed")
        }

        // Call proxy handler
        fbeBegin, err := p.balanceModel.model.GetBegin()
        if fbeBegin == 0 {
            return false, err
        }
        p.HandlerOnProxyBalance.OnProxyBalance(p.balanceModel, fbeType, buffer)
        p.balanceModel.model.GetEnd(fbeBegin)
        return true, nil
    case p.accountModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.accountModel.Buffer().Attach(buffer)
        if !p.accountModel.Verify() {
            return false, errors.New("proto.Account validation failed")
        }

        // Call proxy handler
        fbeBegin, err := p.accountModel.model.GetBegin()
        if fbeBegin == 0 {
            return false, err
        }
        p.HandlerOnProxyAccount.OnProxyAccount(p.accountModel, fbeType, buffer)
        p.accountModel.model.GetEnd(fbeBegin)
        return true, nil
    }

    return false, nil
}
