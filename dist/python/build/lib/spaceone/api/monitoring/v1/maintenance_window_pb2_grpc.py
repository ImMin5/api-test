# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.monitoring.v1 import maintenance_window_pb2 as spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2


class MaintenanceWindowStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.monitoring.v1.MaintenanceWindow/create',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.CreateMaintenanceWindowRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowInfo.FromString,
                )
        self.update = channel.unary_unary(
                '/spaceone.api.monitoring.v1.MaintenanceWindow/update',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.UpdateMaintenanceWindowRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowInfo.FromString,
                )
        self.close = channel.unary_unary(
                '/spaceone.api.monitoring.v1.MaintenanceWindow/close',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowInfo.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.monitoring.v1.MaintenanceWindow/get',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.GetMaintenanceWindowRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.monitoring.v1.MaintenanceWindow/list',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowsInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.monitoring.v1.MaintenanceWindow/stat',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class MaintenanceWindowServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Creates a new MaintenanceWindow. When creating a MaintenanceWindow, you can set the title and maintenance schedule of the MaintenanceWindow. From the `start_time` to the `end_time` specified by the schedule set in this method, alerts in the Projects linked with the MaintenanceWindow are ceased.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Updates a specific MaintenanceWindow. You can make changes in MaintenanceWindow settings including, the `title` and the schedule.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def close(self, request, context):
        """Closes a MaintenanceWindow by changing the state of the MaintenanceWindow to `CLOSED` when the maintenance is completed. As the MaintenanceWindow is not deleted but closed, the maintenance history remains undeleted.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Gets a specific MaintenanceWindow. Prints detailed information about the MaintenanceWindow, including the title and the schedule.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Gets a list of all MaintenanceWindows. You can use a query to get a filtered list of MaintenanceWindows.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MaintenanceWindowServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.CreateMaintenanceWindowRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.UpdateMaintenanceWindowRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowInfo.SerializeToString,
            ),
            'close': grpc.unary_unary_rpc_method_handler(
                    servicer.close,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowInfo.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.GetMaintenanceWindowRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowsInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.monitoring.v1.MaintenanceWindow', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class MaintenanceWindow(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.monitoring.v1.MaintenanceWindow/create',
            spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.CreateMaintenanceWindowRequest.SerializeToString,
            spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.monitoring.v1.MaintenanceWindow/update',
            spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.UpdateMaintenanceWindowRequest.SerializeToString,
            spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def close(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.monitoring.v1.MaintenanceWindow/close',
            spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowRequest.SerializeToString,
            spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.monitoring.v1.MaintenanceWindow/get',
            spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.GetMaintenanceWindowRequest.SerializeToString,
            spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def list(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.monitoring.v1.MaintenanceWindow/list',
            spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowQuery.SerializeToString,
            spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowsInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def stat(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.monitoring.v1.MaintenanceWindow/stat',
            spaceone_dot_api_dot_monitoring_dot_v1_dot_maintenance__window__pb2.MaintenanceWindowStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
