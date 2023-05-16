# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.repository.v1 import schema_pb2 as spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2


class SchemaStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.repository.v1.Schema/create',
                request_serializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.CreateSchemaRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaInfo.FromString,
                )
        self.update = channel.unary_unary(
                '/spaceone.api.repository.v1.Schema/update',
                request_serializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.UpdateSchemaRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaInfo.FromString,
                )
        self.delete = channel.unary_unary(
                '/spaceone.api.repository.v1.Schema/delete',
                request_serializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.repository.v1.Schema/get',
                request_serializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.GetRepositorySchemaRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.repository.v1.Schema/list',
                request_serializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemasInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.repository.v1.Schema/stat',
                request_serializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class SchemaServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Creates a new Schema. You must specify the parameters: `service_type`, `name`, and `schema`(data structure). With the parameter `domain_id`, you can choose whether you will create a Schema in `Local` or externally. The Schema created includes `repository_info`, information about where the resource is managed.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Updates a specific Schema. You can make changes in Schema settings, including `name`, `schema`, `labels`, and `tags`.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
        """Deletes a specific Schema. You must specify the `name` of the Schema to delete, as the `name` is an identifier of Schema resources.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Gets a specific Schema. You must specify the `name` of the Schema to get, as the `name` is an identifier of Schema resources. You can use the parameter `repository_id` to limit the scope of the method to a specific Repository.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Gets a list of all Schemas in a specific Repository. The parameter `repository_id` is used as an identifier of a Repository to get its list of Schemas. You can use a query to get a filtered list of Schemas.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SchemaServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.CreateSchemaRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.UpdateSchemaRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.GetRepositorySchemaRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemasInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.repository.v1.Schema', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Schema(object):
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.repository.v1.Schema/create',
            spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.CreateSchemaRequest.SerializeToString,
            spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.repository.v1.Schema/update',
            spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.UpdateSchemaRequest.SerializeToString,
            spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.repository.v1.Schema/delete',
            spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.repository.v1.Schema/get',
            spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.GetRepositorySchemaRequest.SerializeToString,
            spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.repository.v1.Schema/list',
            spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaQuery.SerializeToString,
            spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemasInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.repository.v1.Schema/stat',
            spaceone_dot_api_dot_repository_dot_v1_dot_schema__pb2.SchemaStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)