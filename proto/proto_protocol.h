// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.3.0.0

#pragma once

#if defined(__clang__)
#pragma clang system_header
#elif defined(__GNUC__)
#pragma GCC system_header
#elif defined(_MSC_VER)
#pragma system_header
#endif

#include "proto_models.h"

namespace FBE {
namespace proto {

// Fast Binary Encoding proto sender
template <class TBuffer>
class Sender : public virtual FBE::Sender<TBuffer>
{
public:
    Sender()
        : OrderModel(this->_buffer)
        , BalanceModel(this->_buffer)
        , AccountModel(this->_buffer)
    {}
    Sender(const Sender&) = default;
    Sender(Sender&&) noexcept = default;
    virtual ~Sender() = default;

    Sender& operator=(const Sender&) = default;
    Sender& operator=(Sender&&) noexcept = default;

    size_t send(const ::proto::Order& value)
    {
        // Serialize the value into the FBE stream
        size_t serialized = OrderModel.serialize(value);
        assert((serialized > 0) && "proto::Order serialization failed!");
        assert(OrderModel.verify() && "proto::Order validation failed!");

        // Log the value
        if (this->_logging)
        {
            std::string message = value.string();
            this->onSendLog(message);
        }

        // Send the serialized value
        return this->send_serialized(serialized);
    }

    size_t send(const ::proto::Balance& value)
    {
        // Serialize the value into the FBE stream
        size_t serialized = BalanceModel.serialize(value);
        assert((serialized > 0) && "proto::Balance serialization failed!");
        assert(BalanceModel.verify() && "proto::Balance validation failed!");

        // Log the value
        if (this->_logging)
        {
            std::string message = value.string();
            this->onSendLog(message);
        }

        // Send the serialized value
        return this->send_serialized(serialized);
    }

    size_t send(const ::proto::Account& value)
    {
        // Serialize the value into the FBE stream
        size_t serialized = AccountModel.serialize(value);
        assert((serialized > 0) && "proto::Account serialization failed!");
        assert(AccountModel.verify() && "proto::Account validation failed!");

        // Log the value
        if (this->_logging)
        {
            std::string message = value.string();
            this->onSendLog(message);
        }

        // Send the serialized value
        return this->send_serialized(serialized);
    }

public:
    // Sender models accessors
    FBE::proto::OrderModel<TBuffer> OrderModel;
    FBE::proto::BalanceModel<TBuffer> BalanceModel;
    FBE::proto::AccountModel<TBuffer> AccountModel;
};

} // namespace proto
} // namespace FBE

namespace FBE {
namespace proto {

// Fast Binary Encoding proto receiver
template <class TBuffer>
class Receiver : public virtual FBE::Receiver<TBuffer>
{
public:
    Receiver() {}
    Receiver(const Receiver&) = default;
    Receiver(Receiver&&) = default;
    virtual ~Receiver() = default;

    Receiver& operator=(const Receiver&) = default;
    Receiver& operator=(Receiver&&) = default;

protected:
    // Receive handlers
    virtual void onReceive(const ::proto::Order& value) {}
    virtual void onReceive(const ::proto::Balance& value) {}
    virtual void onReceive(const ::proto::Account& value) {}

    // Receive message handler
    bool onReceive(size_t type, const void* data, size_t size) override
    {
        switch (type)
        {
            case FBE::proto::OrderModel<ReadBuffer>::fbe_type():
            {
                // Deserialize the value from the FBE stream
                OrderModel.attach(data, size);
                assert(OrderModel.verify() && "proto::Order validation failed!");
                [[maybe_unused]] size_t deserialized = OrderModel.deserialize(OrderValue);
                assert((deserialized > 0) && "proto::Order deserialization failed!");

                // Log the value
                if (this->_logging)
                {
                    std::string message = OrderValue.string();
                    this->onReceiveLog(message);
                }

                // Call receive handler with deserialized value
                onReceive(OrderValue);
                return true;
            }
            case FBE::proto::BalanceModel<ReadBuffer>::fbe_type():
            {
                // Deserialize the value from the FBE stream
                BalanceModel.attach(data, size);
                assert(BalanceModel.verify() && "proto::Balance validation failed!");
                [[maybe_unused]] size_t deserialized = BalanceModel.deserialize(BalanceValue);
                assert((deserialized > 0) && "proto::Balance deserialization failed!");

                // Log the value
                if (this->_logging)
                {
                    std::string message = BalanceValue.string();
                    this->onReceiveLog(message);
                }

                // Call receive handler with deserialized value
                onReceive(BalanceValue);
                return true;
            }
            case FBE::proto::AccountModel<ReadBuffer>::fbe_type():
            {
                // Deserialize the value from the FBE stream
                AccountModel.attach(data, size);
                assert(AccountModel.verify() && "proto::Account validation failed!");
                [[maybe_unused]] size_t deserialized = AccountModel.deserialize(AccountValue);
                assert((deserialized > 0) && "proto::Account deserialization failed!");

                // Log the value
                if (this->_logging)
                {
                    std::string message = AccountValue.string();
                    this->onReceiveLog(message);
                }

                // Call receive handler with deserialized value
                onReceive(AccountValue);
                return true;
            }
        }

        return false;
    }

private:
    // Receiver values accessors
    ::proto::Order OrderValue;
    ::proto::Balance BalanceValue;
    ::proto::Account AccountValue;

    // Receiver models accessors
    FBE::proto::OrderModel<ReadBuffer> OrderModel;
    FBE::proto::BalanceModel<ReadBuffer> BalanceModel;
    FBE::proto::AccountModel<ReadBuffer> AccountModel;
};

} // namespace proto
} // namespace FBE

namespace FBE {
namespace proto {

// Fast Binary Encoding proto client
template <class TBuffer>
class Client : public virtual Sender<TBuffer>, protected virtual Receiver<TBuffer>
{
public:
    Client() = default;
    Client(const Client&) = default;
    Client(Client&&) = default;
    virtual ~Client() = default;

    Client& operator=(const Client&) = default;
    Client& operator=(Client&&) = default;

    // Reset client buffers
    void reset()
    {
        std::scoped_lock locker(this->_lock);
        reset_requests();
    }

    // Watchdog for timeouts
    void watchdog(uint64_t utc)
    {
        std::scoped_lock locker(this->_lock);
        watchdog_requests(utc);
    }

protected:
    std::mutex _lock;
    uint64_t _timestamp{0};

    virtual bool onReceiveResponse(const ::proto::Order& response) { return false; }
    virtual bool onReceiveResponse(const ::proto::Balance& response) { return false; }
    virtual bool onReceiveResponse(const ::proto::Account& response) { return false; }

    virtual bool onReceiveReject(const ::proto::Order& reject) { return false; }
    virtual bool onReceiveReject(const ::proto::Balance& reject) { return false; }
    virtual bool onReceiveReject(const ::proto::Account& reject) { return false; }

    virtual void onReceiveNotify(const ::proto::Order& notify) {}
    virtual void onReceiveNotify(const ::proto::Balance& notify) {}
    virtual void onReceiveNotify(const ::proto::Account& notify) {}

    virtual void onReceive(const ::proto::Order& value) override { if (!onReceiveResponse(value) && !onReceiveReject(value)) onReceiveNotify(value); }
    virtual void onReceive(const ::proto::Balance& value) override { if (!onReceiveResponse(value) && !onReceiveReject(value)) onReceiveNotify(value); }
    virtual void onReceive(const ::proto::Account& value) override { if (!onReceiveResponse(value) && !onReceiveReject(value)) onReceiveNotify(value); }

    // Reset client requests
    virtual void reset_requests()
    {
        Sender<TBuffer>::reset();
        Receiver<TBuffer>::reset();
    }

    // Watchdog client requests for timeouts
    virtual void watchdog_requests(uint64_t utc)
    {
    }
};

} // namespace proto
} // namespace FBE

namespace FBE {
namespace proto {

// Fast Binary Encoding proto proxy
template <class TBuffer>
class Proxy : public virtual FBE::Receiver<TBuffer>
{
public:
    Proxy() {}
    Proxy(const Proxy&) = default;
    Proxy(Proxy&&) = default;
    virtual ~Proxy() = default;

    Proxy& operator=(const Proxy&) = default;
    Proxy& operator=(Proxy&&) = default;

protected:
    // Proxy handlers
    virtual void onProxy(FBE::proto::OrderModel<ReadBuffer>& model, size_t type, const void* data, size_t size) {}
    virtual void onProxy(FBE::proto::BalanceModel<ReadBuffer>& model, size_t type, const void* data, size_t size) {}
    virtual void onProxy(FBE::proto::AccountModel<ReadBuffer>& model, size_t type, const void* data, size_t size) {}

    // Receive message handler
    bool onReceive(size_t type, const void* data, size_t size) override
    {
        switch (type)
        {
            case FBE::proto::OrderModel<ReadBuffer>::fbe_type():
            {
                // Attach the FBE stream to the proxy model
                OrderModel.attach(data, size);
                assert(OrderModel.verify() && "proto::Order validation failed!");

                size_t fbe_begin = OrderModel.model.get_begin();
                if (fbe_begin == 0)
                    return false;
                // Call proxy handler
                onProxy(OrderModel, type, data, size);
                OrderModel.model.get_end(fbe_begin);
                return true;
            }
            case FBE::proto::BalanceModel<ReadBuffer>::fbe_type():
            {
                // Attach the FBE stream to the proxy model
                BalanceModel.attach(data, size);
                assert(BalanceModel.verify() && "proto::Balance validation failed!");

                size_t fbe_begin = BalanceModel.model.get_begin();
                if (fbe_begin == 0)
                    return false;
                // Call proxy handler
                onProxy(BalanceModel, type, data, size);
                BalanceModel.model.get_end(fbe_begin);
                return true;
            }
            case FBE::proto::AccountModel<ReadBuffer>::fbe_type():
            {
                // Attach the FBE stream to the proxy model
                AccountModel.attach(data, size);
                assert(AccountModel.verify() && "proto::Account validation failed!");

                size_t fbe_begin = AccountModel.model.get_begin();
                if (fbe_begin == 0)
                    return false;
                // Call proxy handler
                onProxy(AccountModel, type, data, size);
                AccountModel.model.get_end(fbe_begin);
                return true;
            }
        }

        return false;
    }

private:
    // Proxy models accessors
    FBE::proto::OrderModel<ReadBuffer> OrderModel;
    FBE::proto::BalanceModel<ReadBuffer> BalanceModel;
    FBE::proto::AccountModel<ReadBuffer> AccountModel;
};

} // namespace proto
} // namespace FBE

namespace FBE {
namespace proto {

// Fast Binary Encoding proto final sender
template <class TBuffer>
class FinalSender : public virtual FBE::Sender<TBuffer>
{
public:
    FinalSender()
        : OrderModel(this->_buffer)
        , BalanceModel(this->_buffer)
        , AccountModel(this->_buffer)
    { this->final(true); }
    FinalSender(const FinalSender&) = default;
    FinalSender(FinalSender&&) noexcept = default;
    virtual ~FinalSender() = default;

    FinalSender& operator=(const FinalSender&) = default;
    FinalSender& operator=(FinalSender&&) noexcept = default;

    size_t send(const ::proto::Order& value)
    {
        // Serialize the value into the FBE stream
        size_t serialized = OrderModel.serialize(value);
        assert((serialized > 0) && "proto::Order serialization failed!");
        assert(OrderModel.verify() && "proto::Order validation failed!");

        // Log the value
        if (this->_logging)
        {
            std::string message = value.string();
            this->onSendLog(message);
        }

        // Send the serialized value
        return this->send_serialized(serialized);
    }

    size_t send(const ::proto::Balance& value)
    {
        // Serialize the value into the FBE stream
        size_t serialized = BalanceModel.serialize(value);
        assert((serialized > 0) && "proto::Balance serialization failed!");
        assert(BalanceModel.verify() && "proto::Balance validation failed!");

        // Log the value
        if (this->_logging)
        {
            std::string message = value.string();
            this->onSendLog(message);
        }

        // Send the serialized value
        return this->send_serialized(serialized);
    }

    size_t send(const ::proto::Account& value)
    {
        // Serialize the value into the FBE stream
        size_t serialized = AccountModel.serialize(value);
        assert((serialized > 0) && "proto::Account serialization failed!");
        assert(AccountModel.verify() && "proto::Account validation failed!");

        // Log the value
        if (this->_logging)
        {
            std::string message = value.string();
            this->onSendLog(message);
        }

        // Send the serialized value
        return this->send_serialized(serialized);
    }

public:
    // Sender models accessors
    FBE::proto::OrderFinalModel<TBuffer> OrderModel;
    FBE::proto::BalanceFinalModel<TBuffer> BalanceModel;
    FBE::proto::AccountFinalModel<TBuffer> AccountModel;
};

} // namespace proto
} // namespace FBE

namespace FBE {
namespace proto {

// Fast Binary Encoding proto final receiver
template <class TBuffer>
class FinalReceiver : public virtual FBE::Receiver<TBuffer>
{
public:
    FinalReceiver() { this->final(true); }
    FinalReceiver(const FinalReceiver&) = default;
    FinalReceiver(FinalReceiver&&) = default;
    virtual ~FinalReceiver() = default;

    FinalReceiver& operator=(const FinalReceiver&) = default;
    FinalReceiver& operator=(FinalReceiver&&) = default;

protected:
    // Receive handlers
    virtual void onReceive(const ::proto::Order& value) {}
    virtual void onReceive(const ::proto::Balance& value) {}
    virtual void onReceive(const ::proto::Account& value) {}

    // Receive message handler
    bool onReceive(size_t type, const void* data, size_t size) override
    {
        switch (type)
        {
            case FBE::proto::OrderFinalModel<ReadBuffer>::fbe_type():
            {
                // Deserialize the value from the FBE stream
                OrderModel.attach(data, size);
                assert(OrderModel.verify() && "proto::Order validation failed!");
                [[maybe_unused]] size_t deserialized = OrderModel.deserialize(OrderValue);
                assert((deserialized > 0) && "proto::Order deserialization failed!");

                // Log the value
                if (this->_logging)
                {
                    std::string message = OrderValue.string();
                    this->onReceiveLog(message);
                }

                // Call receive handler with deserialized value
                onReceive(OrderValue);
                return true;
            }
            case FBE::proto::BalanceFinalModel<ReadBuffer>::fbe_type():
            {
                // Deserialize the value from the FBE stream
                BalanceModel.attach(data, size);
                assert(BalanceModel.verify() && "proto::Balance validation failed!");
                [[maybe_unused]] size_t deserialized = BalanceModel.deserialize(BalanceValue);
                assert((deserialized > 0) && "proto::Balance deserialization failed!");

                // Log the value
                if (this->_logging)
                {
                    std::string message = BalanceValue.string();
                    this->onReceiveLog(message);
                }

                // Call receive handler with deserialized value
                onReceive(BalanceValue);
                return true;
            }
            case FBE::proto::AccountFinalModel<ReadBuffer>::fbe_type():
            {
                // Deserialize the value from the FBE stream
                AccountModel.attach(data, size);
                assert(AccountModel.verify() && "proto::Account validation failed!");
                [[maybe_unused]] size_t deserialized = AccountModel.deserialize(AccountValue);
                assert((deserialized > 0) && "proto::Account deserialization failed!");

                // Log the value
                if (this->_logging)
                {
                    std::string message = AccountValue.string();
                    this->onReceiveLog(message);
                }

                // Call receive handler with deserialized value
                onReceive(AccountValue);
                return true;
            }
        }

        return false;
    }

private:
    // Receiver values accessors
    ::proto::Order OrderValue;
    ::proto::Balance BalanceValue;
    ::proto::Account AccountValue;

    // Receiver models accessors
    FBE::proto::OrderFinalModel<ReadBuffer> OrderModel;
    FBE::proto::BalanceFinalModel<ReadBuffer> BalanceModel;
    FBE::proto::AccountFinalModel<ReadBuffer> AccountModel;
};

} // namespace proto
} // namespace FBE

namespace FBE {
namespace proto {

// Fast Binary Encoding proto final client
template <class TBuffer>
class FinalClient : public virtual FinalSender<TBuffer>, protected virtual FinalReceiver<TBuffer>
{
public:
    FinalClient() = default;
    FinalClient(const FinalClient&) = default;
    FinalClient(FinalClient&&) = default;
    virtual ~FinalClient() = default;

    FinalClient& operator=(const FinalClient&) = default;
    FinalClient& operator=(FinalClient&&) = default;

    // Reset client buffers
    void reset()
    {
        std::scoped_lock locker(this->_lock);
        reset_requests();
    }

    // Watchdog for timeouts
    void watchdog(uint64_t utc)
    {
        std::scoped_lock locker(this->_lock);
        watchdog_requests(utc);
    }

protected:
    std::mutex _lock;
    uint64_t _timestamp{0};

    virtual bool onReceiveResponse(const ::proto::Order& response) { return false; }
    virtual bool onReceiveResponse(const ::proto::Balance& response) { return false; }
    virtual bool onReceiveResponse(const ::proto::Account& response) { return false; }

    virtual bool onReceiveReject(const ::proto::Order& reject) { return false; }
    virtual bool onReceiveReject(const ::proto::Balance& reject) { return false; }
    virtual bool onReceiveReject(const ::proto::Account& reject) { return false; }

    virtual void onReceiveNotify(const ::proto::Order& notify) {}
    virtual void onReceiveNotify(const ::proto::Balance& notify) {}
    virtual void onReceiveNotify(const ::proto::Account& notify) {}

    virtual void onReceive(const ::proto::Order& value) override { if (!onReceiveResponse(value) && !onReceiveReject(value)) onReceiveNotify(value); }
    virtual void onReceive(const ::proto::Balance& value) override { if (!onReceiveResponse(value) && !onReceiveReject(value)) onReceiveNotify(value); }
    virtual void onReceive(const ::proto::Account& value) override { if (!onReceiveResponse(value) && !onReceiveReject(value)) onReceiveNotify(value); }

    // Reset client requests
    virtual void reset_requests()
    {
        Sender<TBuffer>::reset();
        Receiver<TBuffer>::reset();
    }

    // Watchdog client requests for timeouts
    virtual void watchdog_requests(uint64_t utc)
    {
    }
};

} // namespace proto
} // namespace FBE
