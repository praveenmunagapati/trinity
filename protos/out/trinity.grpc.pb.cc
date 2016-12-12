// Generated by the gRPC protobuf plugin.
// If you make any local change, they will be lost.
// source: trinity.proto

#include "trinity.pb.h"
#include "trinity.grpc.pb.h"

#include <grpc++/impl/codegen/async_stream.h>
#include <grpc++/impl/codegen/async_unary_call.h>
#include <grpc++/impl/codegen/channel_interface.h>
#include <grpc++/impl/codegen/client_unary_call.h>
#include <grpc++/impl/codegen/method_handler_impl.h>
#include <grpc++/impl/codegen/rpc_service_method.h>
#include <grpc++/impl/codegen/service_type.h>
#include <grpc++/impl/codegen/sync_stream.h>
namespace trinity {

static const char* FMapper_method_names[] = {
  "/trinity.FMapper/GetDocMap",
  "/trinity.FMapper/SetDocMap",
};

std::unique_ptr< FMapper::Stub> FMapper::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  std::unique_ptr< FMapper::Stub> stub(new FMapper::Stub(channel));
  return stub;
}

FMapper::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_GetDocMap_(FMapper_method_names[0], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SetDocMap_(FMapper_method_names[1], ::grpc::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status FMapper::Stub::GetDocMap(::grpc::ClientContext* context, const ::trinity::DocMapRequest& request, ::trinity::ForwardMap* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_GetDocMap_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::trinity::ForwardMap>* FMapper::Stub::AsyncGetDocMapRaw(::grpc::ClientContext* context, const ::trinity::DocMapRequest& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::trinity::ForwardMap>(channel_.get(), cq, rpcmethod_GetDocMap_, context, request);
}

::grpc::Status FMapper::Stub::SetDocMap(::grpc::ClientContext* context, const ::trinity::ForwardMap& request, ::trinity::SetResult* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_SetDocMap_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::trinity::SetResult>* FMapper::Stub::AsyncSetDocMapRaw(::grpc::ClientContext* context, const ::trinity::ForwardMap& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::trinity::SetResult>(channel_.get(), cq, rpcmethod_SetDocMap_, context, request);
}

FMapper::Service::Service() {
  (void)FMapper_method_names;
  AddMethod(new ::grpc::RpcServiceMethod(
      FMapper_method_names[0],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< FMapper::Service, ::trinity::DocMapRequest, ::trinity::ForwardMap>(
          std::mem_fn(&FMapper::Service::GetDocMap), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      FMapper_method_names[1],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< FMapper::Service, ::trinity::ForwardMap, ::trinity::SetResult>(
          std::mem_fn(&FMapper::Service::SetDocMap), this)));
}

FMapper::Service::~Service() {
}

::grpc::Status FMapper::Service::GetDocMap(::grpc::ServerContext* context, const ::trinity::DocMapRequest* request, ::trinity::ForwardMap* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status FMapper::Service::SetDocMap(::grpc::ServerContext* context, const ::trinity::ForwardMap* request, ::trinity::SetResult* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


static const char* Trinity_method_names[] = {
  "/trinity.Trinity/GetRootConfig",
  "/trinity.Trinity/SetRootConfig",
};

std::unique_ptr< Trinity::Stub> Trinity::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  std::unique_ptr< Trinity::Stub> stub(new Trinity::Stub(channel));
  return stub;
}

Trinity::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_GetRootConfig_(Trinity_method_names[0], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SetRootConfig_(Trinity_method_names[1], ::grpc::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status Trinity::Stub::GetRootConfig(::grpc::ClientContext* context, const ::trinity::ConfigRequest& request, ::trinity::RootConfig* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_GetRootConfig_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::trinity::RootConfig>* Trinity::Stub::AsyncGetRootConfigRaw(::grpc::ClientContext* context, const ::trinity::ConfigRequest& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::trinity::RootConfig>(channel_.get(), cq, rpcmethod_GetRootConfig_, context, request);
}

::grpc::Status Trinity::Stub::SetRootConfig(::grpc::ClientContext* context, const ::trinity::RootConfig& request, ::trinity::SetResult* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_SetRootConfig_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::trinity::SetResult>* Trinity::Stub::AsyncSetRootConfigRaw(::grpc::ClientContext* context, const ::trinity::RootConfig& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::trinity::SetResult>(channel_.get(), cq, rpcmethod_SetRootConfig_, context, request);
}

Trinity::Service::Service() {
  (void)Trinity_method_names;
  AddMethod(new ::grpc::RpcServiceMethod(
      Trinity_method_names[0],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< Trinity::Service, ::trinity::ConfigRequest, ::trinity::RootConfig>(
          std::mem_fn(&Trinity::Service::GetRootConfig), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      Trinity_method_names[1],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< Trinity::Service, ::trinity::RootConfig, ::trinity::SetResult>(
          std::mem_fn(&Trinity::Service::SetRootConfig), this)));
}

Trinity::Service::~Service() {
}

::grpc::Status Trinity::Service::GetRootConfig(::grpc::ServerContext* context, const ::trinity::ConfigRequest* request, ::trinity::RootConfig* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Trinity::Service::SetRootConfig(::grpc::ServerContext* context, const ::trinity::RootConfig* request, ::trinity::SetResult* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


static const char* Subsystem_method_names[] = {
  "/trinity.Subsystem/StartMainLoop",
};

std::unique_ptr< Subsystem::Stub> Subsystem::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  std::unique_ptr< Subsystem::Stub> stub(new Subsystem::Stub(channel));
  return stub;
}

Subsystem::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_StartMainLoop_(Subsystem_method_names[0], ::grpc::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status Subsystem::Stub::StartMainLoop(::grpc::ClientContext* context, const ::trinity::StartSubsystemRequest& request, ::trinity::StartSubsystemResponse* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_StartMainLoop_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::trinity::StartSubsystemResponse>* Subsystem::Stub::AsyncStartMainLoopRaw(::grpc::ClientContext* context, const ::trinity::StartSubsystemRequest& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::trinity::StartSubsystemResponse>(channel_.get(), cq, rpcmethod_StartMainLoop_, context, request);
}

Subsystem::Service::Service() {
  (void)Subsystem_method_names;
  AddMethod(new ::grpc::RpcServiceMethod(
      Subsystem_method_names[0],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< Subsystem::Service, ::trinity::StartSubsystemRequest, ::trinity::StartSubsystemResponse>(
          std::mem_fn(&Subsystem::Service::StartMainLoop), this)));
}

Subsystem::Service::~Service() {
}

::grpc::Status Subsystem::Service::StartMainLoop(::grpc::ServerContext* context, const ::trinity::StartSubsystemRequest* request, ::trinity::StartSubsystemResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


static const char* Gateway_method_names[] = {
  "/trinity.Gateway/GetSearchQuery",
};

std::unique_ptr< Gateway::Stub> Gateway::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  std::unique_ptr< Gateway::Stub> stub(new Gateway::Stub(channel));
  return stub;
}

Gateway::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_GetSearchQuery_(Gateway_method_names[0], ::grpc::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status Gateway::Stub::GetSearchQuery(::grpc::ClientContext* context, const ::trinity::SearchQuery& request, ::trinity::SearchResponse* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_GetSearchQuery_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::trinity::SearchResponse>* Gateway::Stub::AsyncGetSearchQueryRaw(::grpc::ClientContext* context, const ::trinity::SearchQuery& request, ::grpc::CompletionQueue* cq) {
  return new ::grpc::ClientAsyncResponseReader< ::trinity::SearchResponse>(channel_.get(), cq, rpcmethod_GetSearchQuery_, context, request);
}

Gateway::Service::Service() {
  (void)Gateway_method_names;
  AddMethod(new ::grpc::RpcServiceMethod(
      Gateway_method_names[0],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< Gateway::Service, ::trinity::SearchQuery, ::trinity::SearchResponse>(
          std::mem_fn(&Gateway::Service::GetSearchQuery), this)));
}

Gateway::Service::~Service() {
}

::grpc::Status Gateway::Service::GetSearchQuery(::grpc::ServerContext* context, const ::trinity::SearchQuery* request, ::trinity::SearchResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace trinity

