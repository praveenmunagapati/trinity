// Generated by the gRPC protobuf plugin.
// If you make any local change, they will be lost.
// source: trinity.proto
#ifndef GRPC_trinity_2eproto__INCLUDED
#define GRPC_trinity_2eproto__INCLUDED

#include "trinity.pb.h"

#include <grpc++/impl/codegen/async_stream.h>
#include <grpc++/impl/codegen/async_unary_call.h>
#include <grpc++/impl/codegen/method_handler_impl.h>
#include <grpc++/impl/codegen/proto_utils.h>
#include <grpc++/impl/codegen/rpc_method.h>
#include <grpc++/impl/codegen/service_type.h>
#include <grpc++/impl/codegen/status.h>
#include <grpc++/impl/codegen/stub_options.h>
#include <grpc++/impl/codegen/sync_stream.h>

namespace grpc {
class CompletionQueue;
class Channel;
class RpcService;
class ServerCompletionQueue;
class ServerContext;
}  // namespace grpc

namespace trinity {

class FMapper GRPC_FINAL {
 public:
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status GetDocMap(::grpc::ClientContext* context, const ::trinity::DocMapRequest& request, ::trinity::ForwardMap* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::trinity::ForwardMap>> AsyncGetDocMap(::grpc::ClientContext* context, const ::trinity::DocMapRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::trinity::ForwardMap>>(AsyncGetDocMapRaw(context, request, cq));
    }
    virtual ::grpc::Status SetDocMap(::grpc::ClientContext* context, const ::trinity::ForwardMap& request, ::trinity::SetResult* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::trinity::SetResult>> AsyncSetDocMap(::grpc::ClientContext* context, const ::trinity::ForwardMap& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::trinity::SetResult>>(AsyncSetDocMapRaw(context, request, cq));
    }
  private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::trinity::ForwardMap>* AsyncGetDocMapRaw(::grpc::ClientContext* context, const ::trinity::DocMapRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::trinity::SetResult>* AsyncSetDocMapRaw(::grpc::ClientContext* context, const ::trinity::ForwardMap& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub GRPC_FINAL : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel);
    ::grpc::Status GetDocMap(::grpc::ClientContext* context, const ::trinity::DocMapRequest& request, ::trinity::ForwardMap* response) GRPC_OVERRIDE;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::trinity::ForwardMap>> AsyncGetDocMap(::grpc::ClientContext* context, const ::trinity::DocMapRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::trinity::ForwardMap>>(AsyncGetDocMapRaw(context, request, cq));
    }
    ::grpc::Status SetDocMap(::grpc::ClientContext* context, const ::trinity::ForwardMap& request, ::trinity::SetResult* response) GRPC_OVERRIDE;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::trinity::SetResult>> AsyncSetDocMap(::grpc::ClientContext* context, const ::trinity::ForwardMap& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::trinity::SetResult>>(AsyncSetDocMapRaw(context, request, cq));
    }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    ::grpc::ClientAsyncResponseReader< ::trinity::ForwardMap>* AsyncGetDocMapRaw(::grpc::ClientContext* context, const ::trinity::DocMapRequest& request, ::grpc::CompletionQueue* cq) GRPC_OVERRIDE;
    ::grpc::ClientAsyncResponseReader< ::trinity::SetResult>* AsyncSetDocMapRaw(::grpc::ClientContext* context, const ::trinity::ForwardMap& request, ::grpc::CompletionQueue* cq) GRPC_OVERRIDE;
    const ::grpc::RpcMethod rpcmethod_GetDocMap_;
    const ::grpc::RpcMethod rpcmethod_SetDocMap_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status GetDocMap(::grpc::ServerContext* context, const ::trinity::DocMapRequest* request, ::trinity::ForwardMap* response);
    virtual ::grpc::Status SetDocMap(::grpc::ServerContext* context, const ::trinity::ForwardMap* request, ::trinity::SetResult* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_GetDocMap : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_GetDocMap() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_GetDocMap() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetDocMap(::grpc::ServerContext* context, const ::trinity::DocMapRequest* request, ::trinity::ForwardMap* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetDocMap(::grpc::ServerContext* context, ::trinity::DocMapRequest* request, ::grpc::ServerAsyncResponseWriter< ::trinity::ForwardMap>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_SetDocMap : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_SetDocMap() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_SetDocMap() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SetDocMap(::grpc::ServerContext* context, const ::trinity::ForwardMap* request, ::trinity::SetResult* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestSetDocMap(::grpc::ServerContext* context, ::trinity::ForwardMap* request, ::grpc::ServerAsyncResponseWriter< ::trinity::SetResult>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_GetDocMap<WithAsyncMethod_SetDocMap<Service > > AsyncService;
  template <class BaseClass>
  class WithGenericMethod_GetDocMap : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_GetDocMap() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_GetDocMap() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetDocMap(::grpc::ServerContext* context, const ::trinity::DocMapRequest* request, ::trinity::ForwardMap* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_SetDocMap : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_SetDocMap() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_SetDocMap() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SetDocMap(::grpc::ServerContext* context, const ::trinity::ForwardMap* request, ::trinity::SetResult* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_GetDocMap : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_GetDocMap() {
      ::grpc::Service::MarkMethodStreamedUnary(0,
        new ::grpc::StreamedUnaryHandler< ::trinity::DocMapRequest, ::trinity::ForwardMap>(std::bind(&WithStreamedUnaryMethod_GetDocMap<BaseClass>::StreamedGetDocMap, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_GetDocMap() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status GetDocMap(::grpc::ServerContext* context, const ::trinity::DocMapRequest* request, ::trinity::ForwardMap* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedGetDocMap(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::trinity::DocMapRequest,::trinity::ForwardMap>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_SetDocMap : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_SetDocMap() {
      ::grpc::Service::MarkMethodStreamedUnary(1,
        new ::grpc::StreamedUnaryHandler< ::trinity::ForwardMap, ::trinity::SetResult>(std::bind(&WithStreamedUnaryMethod_SetDocMap<BaseClass>::StreamedSetDocMap, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_SetDocMap() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status SetDocMap(::grpc::ServerContext* context, const ::trinity::ForwardMap* request, ::trinity::SetResult* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedSetDocMap(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::trinity::ForwardMap,::trinity::SetResult>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_GetDocMap<WithStreamedUnaryMethod_SetDocMap<Service > > StreamedUnaryService;
};

class Trinity GRPC_FINAL {
 public:
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status GetRootConfig(::grpc::ClientContext* context, const ::trinity::ConfigRequest& request, ::trinity::RootConfig* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::trinity::RootConfig>> AsyncGetRootConfig(::grpc::ClientContext* context, const ::trinity::ConfigRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::trinity::RootConfig>>(AsyncGetRootConfigRaw(context, request, cq));
    }
    virtual ::grpc::Status SetRootConfig(::grpc::ClientContext* context, const ::trinity::RootConfig& request, ::trinity::SetResult* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::trinity::SetResult>> AsyncSetRootConfig(::grpc::ClientContext* context, const ::trinity::RootConfig& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::trinity::SetResult>>(AsyncSetRootConfigRaw(context, request, cq));
    }
  private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::trinity::RootConfig>* AsyncGetRootConfigRaw(::grpc::ClientContext* context, const ::trinity::ConfigRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::trinity::SetResult>* AsyncSetRootConfigRaw(::grpc::ClientContext* context, const ::trinity::RootConfig& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub GRPC_FINAL : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel);
    ::grpc::Status GetRootConfig(::grpc::ClientContext* context, const ::trinity::ConfigRequest& request, ::trinity::RootConfig* response) GRPC_OVERRIDE;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::trinity::RootConfig>> AsyncGetRootConfig(::grpc::ClientContext* context, const ::trinity::ConfigRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::trinity::RootConfig>>(AsyncGetRootConfigRaw(context, request, cq));
    }
    ::grpc::Status SetRootConfig(::grpc::ClientContext* context, const ::trinity::RootConfig& request, ::trinity::SetResult* response) GRPC_OVERRIDE;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::trinity::SetResult>> AsyncSetRootConfig(::grpc::ClientContext* context, const ::trinity::RootConfig& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::trinity::SetResult>>(AsyncSetRootConfigRaw(context, request, cq));
    }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    ::grpc::ClientAsyncResponseReader< ::trinity::RootConfig>* AsyncGetRootConfigRaw(::grpc::ClientContext* context, const ::trinity::ConfigRequest& request, ::grpc::CompletionQueue* cq) GRPC_OVERRIDE;
    ::grpc::ClientAsyncResponseReader< ::trinity::SetResult>* AsyncSetRootConfigRaw(::grpc::ClientContext* context, const ::trinity::RootConfig& request, ::grpc::CompletionQueue* cq) GRPC_OVERRIDE;
    const ::grpc::RpcMethod rpcmethod_GetRootConfig_;
    const ::grpc::RpcMethod rpcmethod_SetRootConfig_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status GetRootConfig(::grpc::ServerContext* context, const ::trinity::ConfigRequest* request, ::trinity::RootConfig* response);
    virtual ::grpc::Status SetRootConfig(::grpc::ServerContext* context, const ::trinity::RootConfig* request, ::trinity::SetResult* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_GetRootConfig : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_GetRootConfig() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_GetRootConfig() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetRootConfig(::grpc::ServerContext* context, const ::trinity::ConfigRequest* request, ::trinity::RootConfig* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetRootConfig(::grpc::ServerContext* context, ::trinity::ConfigRequest* request, ::grpc::ServerAsyncResponseWriter< ::trinity::RootConfig>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_SetRootConfig : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_SetRootConfig() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_SetRootConfig() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SetRootConfig(::grpc::ServerContext* context, const ::trinity::RootConfig* request, ::trinity::SetResult* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestSetRootConfig(::grpc::ServerContext* context, ::trinity::RootConfig* request, ::grpc::ServerAsyncResponseWriter< ::trinity::SetResult>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_GetRootConfig<WithAsyncMethod_SetRootConfig<Service > > AsyncService;
  template <class BaseClass>
  class WithGenericMethod_GetRootConfig : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_GetRootConfig() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_GetRootConfig() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetRootConfig(::grpc::ServerContext* context, const ::trinity::ConfigRequest* request, ::trinity::RootConfig* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_SetRootConfig : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_SetRootConfig() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_SetRootConfig() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SetRootConfig(::grpc::ServerContext* context, const ::trinity::RootConfig* request, ::trinity::SetResult* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_GetRootConfig : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_GetRootConfig() {
      ::grpc::Service::MarkMethodStreamedUnary(0,
        new ::grpc::StreamedUnaryHandler< ::trinity::ConfigRequest, ::trinity::RootConfig>(std::bind(&WithStreamedUnaryMethod_GetRootConfig<BaseClass>::StreamedGetRootConfig, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_GetRootConfig() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status GetRootConfig(::grpc::ServerContext* context, const ::trinity::ConfigRequest* request, ::trinity::RootConfig* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedGetRootConfig(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::trinity::ConfigRequest,::trinity::RootConfig>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_SetRootConfig : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_SetRootConfig() {
      ::grpc::Service::MarkMethodStreamedUnary(1,
        new ::grpc::StreamedUnaryHandler< ::trinity::RootConfig, ::trinity::SetResult>(std::bind(&WithStreamedUnaryMethod_SetRootConfig<BaseClass>::StreamedSetRootConfig, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_SetRootConfig() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status SetRootConfig(::grpc::ServerContext* context, const ::trinity::RootConfig* request, ::trinity::SetResult* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedSetRootConfig(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::trinity::RootConfig,::trinity::SetResult>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_GetRootConfig<WithStreamedUnaryMethod_SetRootConfig<Service > > StreamedUnaryService;
};

class Subsystem GRPC_FINAL {
 public:
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status StartMainLoop(::grpc::ClientContext* context, const ::trinity::StartSubsystemRequest& request, ::trinity::StartSubsystemResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::trinity::StartSubsystemResponse>> AsyncStartMainLoop(::grpc::ClientContext* context, const ::trinity::StartSubsystemRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::trinity::StartSubsystemResponse>>(AsyncStartMainLoopRaw(context, request, cq));
    }
  private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::trinity::StartSubsystemResponse>* AsyncStartMainLoopRaw(::grpc::ClientContext* context, const ::trinity::StartSubsystemRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub GRPC_FINAL : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel);
    ::grpc::Status StartMainLoop(::grpc::ClientContext* context, const ::trinity::StartSubsystemRequest& request, ::trinity::StartSubsystemResponse* response) GRPC_OVERRIDE;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::trinity::StartSubsystemResponse>> AsyncStartMainLoop(::grpc::ClientContext* context, const ::trinity::StartSubsystemRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::trinity::StartSubsystemResponse>>(AsyncStartMainLoopRaw(context, request, cq));
    }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    ::grpc::ClientAsyncResponseReader< ::trinity::StartSubsystemResponse>* AsyncStartMainLoopRaw(::grpc::ClientContext* context, const ::trinity::StartSubsystemRequest& request, ::grpc::CompletionQueue* cq) GRPC_OVERRIDE;
    const ::grpc::RpcMethod rpcmethod_StartMainLoop_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status StartMainLoop(::grpc::ServerContext* context, const ::trinity::StartSubsystemRequest* request, ::trinity::StartSubsystemResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_StartMainLoop : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_StartMainLoop() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_StartMainLoop() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status StartMainLoop(::grpc::ServerContext* context, const ::trinity::StartSubsystemRequest* request, ::trinity::StartSubsystemResponse* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestStartMainLoop(::grpc::ServerContext* context, ::trinity::StartSubsystemRequest* request, ::grpc::ServerAsyncResponseWriter< ::trinity::StartSubsystemResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_StartMainLoop<Service > AsyncService;
  template <class BaseClass>
  class WithGenericMethod_StartMainLoop : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_StartMainLoop() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_StartMainLoop() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status StartMainLoop(::grpc::ServerContext* context, const ::trinity::StartSubsystemRequest* request, ::trinity::StartSubsystemResponse* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_StartMainLoop : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_StartMainLoop() {
      ::grpc::Service::MarkMethodStreamedUnary(0,
        new ::grpc::StreamedUnaryHandler< ::trinity::StartSubsystemRequest, ::trinity::StartSubsystemResponse>(std::bind(&WithStreamedUnaryMethod_StartMainLoop<BaseClass>::StreamedStartMainLoop, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_StartMainLoop() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status StartMainLoop(::grpc::ServerContext* context, const ::trinity::StartSubsystemRequest* request, ::trinity::StartSubsystemResponse* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedStartMainLoop(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::trinity::StartSubsystemRequest,::trinity::StartSubsystemResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_StartMainLoop<Service > StreamedUnaryService;
};

class Gateway GRPC_FINAL {
 public:
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status GetSearchQuery(::grpc::ClientContext* context, const ::trinity::SearchQuery& request, ::trinity::SearchResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::trinity::SearchResponse>> AsyncGetSearchQuery(::grpc::ClientContext* context, const ::trinity::SearchQuery& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::trinity::SearchResponse>>(AsyncGetSearchQueryRaw(context, request, cq));
    }
  private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::trinity::SearchResponse>* AsyncGetSearchQueryRaw(::grpc::ClientContext* context, const ::trinity::SearchQuery& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub GRPC_FINAL : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel);
    ::grpc::Status GetSearchQuery(::grpc::ClientContext* context, const ::trinity::SearchQuery& request, ::trinity::SearchResponse* response) GRPC_OVERRIDE;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::trinity::SearchResponse>> AsyncGetSearchQuery(::grpc::ClientContext* context, const ::trinity::SearchQuery& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::trinity::SearchResponse>>(AsyncGetSearchQueryRaw(context, request, cq));
    }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    ::grpc::ClientAsyncResponseReader< ::trinity::SearchResponse>* AsyncGetSearchQueryRaw(::grpc::ClientContext* context, const ::trinity::SearchQuery& request, ::grpc::CompletionQueue* cq) GRPC_OVERRIDE;
    const ::grpc::RpcMethod rpcmethod_GetSearchQuery_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status GetSearchQuery(::grpc::ServerContext* context, const ::trinity::SearchQuery* request, ::trinity::SearchResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_GetSearchQuery : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_GetSearchQuery() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_GetSearchQuery() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetSearchQuery(::grpc::ServerContext* context, const ::trinity::SearchQuery* request, ::trinity::SearchResponse* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetSearchQuery(::grpc::ServerContext* context, ::trinity::SearchQuery* request, ::grpc::ServerAsyncResponseWriter< ::trinity::SearchResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_GetSearchQuery<Service > AsyncService;
  template <class BaseClass>
  class WithGenericMethod_GetSearchQuery : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_GetSearchQuery() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_GetSearchQuery() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetSearchQuery(::grpc::ServerContext* context, const ::trinity::SearchQuery* request, ::trinity::SearchResponse* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_GetSearchQuery : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_GetSearchQuery() {
      ::grpc::Service::MarkMethodStreamedUnary(0,
        new ::grpc::StreamedUnaryHandler< ::trinity::SearchQuery, ::trinity::SearchResponse>(std::bind(&WithStreamedUnaryMethod_GetSearchQuery<BaseClass>::StreamedGetSearchQuery, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_GetSearchQuery() GRPC_OVERRIDE {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status GetSearchQuery(::grpc::ServerContext* context, const ::trinity::SearchQuery* request, ::trinity::SearchResponse* response) GRPC_FINAL GRPC_OVERRIDE {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedGetSearchQuery(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::trinity::SearchQuery,::trinity::SearchResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_GetSearchQuery<Service > StreamedUnaryService;
};

}  // namespace trinity


#endif  // GRPC_trinity_2eproto__INCLUDED