# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc
from grpc.framework.common import cardinality
from grpc.framework.interfaces.face import utilities as face_utilities

import trinity_pb2 as trinity__pb2


class FIndexStub(object):

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetFMap = channel.unary_unary(
        '/trinity.FIndex/GetFMap',
        request_serializer=trinity__pb2.FMapRequest.SerializeToString,
        response_deserializer=trinity__pb2.ForwardMap.FromString,
        )
    self.SetFMap = channel.unary_unary(
        '/trinity.FIndex/SetFMap',
        request_serializer=trinity__pb2.ForwardMap.SerializeToString,
        response_deserializer=trinity__pb2.SetResult.FromString,
        )


class FIndexServicer(object):

  def GetFMap(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetFMap(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_FIndexServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetFMap': grpc.unary_unary_rpc_method_handler(
          servicer.GetFMap,
          request_deserializer=trinity__pb2.FMapRequest.FromString,
          response_serializer=trinity__pb2.ForwardMap.SerializeToString,
      ),
      'SetFMap': grpc.unary_unary_rpc_method_handler(
          servicer.SetFMap,
          request_deserializer=trinity__pb2.ForwardMap.FromString,
          response_serializer=trinity__pb2.SetResult.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'trinity.FIndex', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class IIndexStub(object):

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetIMap = channel.unary_unary(
        '/trinity.IIndex/GetIMap',
        request_serializer=trinity__pb2.IMapRequest.SerializeToString,
        response_deserializer=trinity__pb2.InvertedMap.FromString,
        )
    self.SetIMap = channel.unary_unary(
        '/trinity.IIndex/SetIMap',
        request_serializer=trinity__pb2.InvertedMap.SerializeToString,
        response_deserializer=trinity__pb2.SetResult.FromString,
        )


class IIndexServicer(object):

  def GetIMap(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetIMap(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_IIndexServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetIMap': grpc.unary_unary_rpc_method_handler(
          servicer.GetIMap,
          request_deserializer=trinity__pb2.IMapRequest.FromString,
          response_serializer=trinity__pb2.InvertedMap.SerializeToString,
      ),
      'SetIMap': grpc.unary_unary_rpc_method_handler(
          servicer.SetIMap,
          request_deserializer=trinity__pb2.InvertedMap.FromString,
          response_serializer=trinity__pb2.SetResult.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'trinity.IIndex', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class TrinityStub(object):

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetRootConfig = channel.unary_unary(
        '/trinity.Trinity/GetRootConfig',
        request_serializer=trinity__pb2.ConfigRequest.SerializeToString,
        response_deserializer=trinity__pb2.RootConfig.FromString,
        )
    self.SetRootConfig = channel.unary_unary(
        '/trinity.Trinity/SetRootConfig',
        request_serializer=trinity__pb2.RootConfig.SerializeToString,
        response_deserializer=trinity__pb2.SetResult.FromString,
        )


class TrinityServicer(object):

  def GetRootConfig(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetRootConfig(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_TrinityServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetRootConfig': grpc.unary_unary_rpc_method_handler(
          servicer.GetRootConfig,
          request_deserializer=trinity__pb2.ConfigRequest.FromString,
          response_serializer=trinity__pb2.RootConfig.SerializeToString,
      ),
      'SetRootConfig': grpc.unary_unary_rpc_method_handler(
          servicer.SetRootConfig,
          request_deserializer=trinity__pb2.RootConfig.FromString,
          response_serializer=trinity__pb2.SetResult.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'trinity.Trinity', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class SubsystemStub(object):

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.StartMainLoop = channel.unary_unary(
        '/trinity.Subsystem/StartMainLoop',
        request_serializer=trinity__pb2.StartSubsystemRequest.SerializeToString,
        response_deserializer=trinity__pb2.StartSubsystemResponse.FromString,
        )


class SubsystemServicer(object):

  def StartMainLoop(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_SubsystemServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'StartMainLoop': grpc.unary_unary_rpc_method_handler(
          servicer.StartMainLoop,
          request_deserializer=trinity__pb2.StartSubsystemRequest.FromString,
          response_serializer=trinity__pb2.StartSubsystemResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'trinity.Subsystem', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class GatewayStub(object):

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Search = channel.unary_unary(
        '/trinity.Gateway/Search',
        request_serializer=trinity__pb2.SearchRequest.SerializeToString,
        response_deserializer=trinity__pb2.SearchResponse.FromString,
        )


class GatewayServicer(object):

  def Search(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_GatewayServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Search': grpc.unary_unary_rpc_method_handler(
          servicer.Search,
          request_deserializer=trinity__pb2.SearchRequest.FromString,
          response_serializer=trinity__pb2.SearchResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'trinity.Gateway', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
