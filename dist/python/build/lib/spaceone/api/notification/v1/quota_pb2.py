# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/notification/v1/quota.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(spaceone/api/notification/v1/quota.proto\x12\x1cspaceone.api.notification.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"d\n\x12\x43reateQuotaRequest\x12\x13\n\x0bprotocol_id\x18\x01 \x01(\t\x12&\n\x05limit\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"a\n\x12UpdateQuotaRequest\x12\x10\n\x08quota_id\x18\x01 \x01(\t\x12&\n\x05limit\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"3\n\x0cQuotaRequest\x12\x10\n\x08quota_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"r\n\nQuotaQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x10\n\x08quota_id\x18\x02 \x01(\t\x12\x13\n\x0bprotocol_id\x18\x03 \x01(\t\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"m\n\tQuotaInfo\x12\x10\n\x08quota_id\x18\x01 \x01(\t\x12\x13\n\x0bprotocol_id\x18\x02 \x01(\t\x12&\n\x05limit\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"[\n\nQuotasInfo\x12\x38\n\x07results\x18\x01 \x03(\x0b\x32\'.spaceone.api.notification.v1.QuotaInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"Y\n\x0eQuotaStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t2\x9f\x06\n\x05Quota\x12\x8d\x01\n\x06\x63reate\x12\x30.spaceone.api.notification.v1.CreateQuotaRequest\x1a\'.spaceone.api.notification.v1.QuotaInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/notification/v1/quota/create:\x01*\x12\x8d\x01\n\x06update\x12\x30.spaceone.api.notification.v1.UpdateQuotaRequest\x1a\'.spaceone.api.notification.v1.QuotaInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/notification/v1/quota/update:\x01*\x12v\n\x06\x64\x65lete\x12*.spaceone.api.notification.v1.QuotaRequest\x1a\x16.google.protobuf.Empty\"(\x82\xd3\xe4\x93\x02\"\"\x1d/notification/v1/quota/delete:\x01*\x12\x81\x01\n\x03get\x12*.spaceone.api.notification.v1.QuotaRequest\x1a\'.spaceone.api.notification.v1.QuotaInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/notification/v1/quota/get:\x01*\x12\x82\x01\n\x04list\x12(.spaceone.api.notification.v1.QuotaQuery\x1a(.spaceone.api.notification.v1.QuotasInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/notification/v1/quota/list:\x01*\x12u\n\x04stat\x12,.spaceone.api.notification.v1.QuotaStatQuery\x1a\x17.google.protobuf.Struct\"&\x82\xd3\xe4\x93\x02 \"\x1b/notification/v1/quota/stat:\x01*BCZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/notification/v1b\x06proto3')



_CREATEQUOTAREQUEST = DESCRIPTOR.message_types_by_name['CreateQuotaRequest']
_UPDATEQUOTAREQUEST = DESCRIPTOR.message_types_by_name['UpdateQuotaRequest']
_QUOTAREQUEST = DESCRIPTOR.message_types_by_name['QuotaRequest']
_QUOTAQUERY = DESCRIPTOR.message_types_by_name['QuotaQuery']
_QUOTAINFO = DESCRIPTOR.message_types_by_name['QuotaInfo']
_QUOTASINFO = DESCRIPTOR.message_types_by_name['QuotasInfo']
_QUOTASTATQUERY = DESCRIPTOR.message_types_by_name['QuotaStatQuery']
CreateQuotaRequest = _reflection.GeneratedProtocolMessageType('CreateQuotaRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEQUOTAREQUEST,
  '__module__' : 'spaceone.api.notification.v1.quota_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.notification.v1.CreateQuotaRequest)
  })
_sym_db.RegisterMessage(CreateQuotaRequest)

UpdateQuotaRequest = _reflection.GeneratedProtocolMessageType('UpdateQuotaRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEQUOTAREQUEST,
  '__module__' : 'spaceone.api.notification.v1.quota_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.notification.v1.UpdateQuotaRequest)
  })
_sym_db.RegisterMessage(UpdateQuotaRequest)

QuotaRequest = _reflection.GeneratedProtocolMessageType('QuotaRequest', (_message.Message,), {
  'DESCRIPTOR' : _QUOTAREQUEST,
  '__module__' : 'spaceone.api.notification.v1.quota_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.notification.v1.QuotaRequest)
  })
_sym_db.RegisterMessage(QuotaRequest)

QuotaQuery = _reflection.GeneratedProtocolMessageType('QuotaQuery', (_message.Message,), {
  'DESCRIPTOR' : _QUOTAQUERY,
  '__module__' : 'spaceone.api.notification.v1.quota_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.notification.v1.QuotaQuery)
  })
_sym_db.RegisterMessage(QuotaQuery)

QuotaInfo = _reflection.GeneratedProtocolMessageType('QuotaInfo', (_message.Message,), {
  'DESCRIPTOR' : _QUOTAINFO,
  '__module__' : 'spaceone.api.notification.v1.quota_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.notification.v1.QuotaInfo)
  })
_sym_db.RegisterMessage(QuotaInfo)

QuotasInfo = _reflection.GeneratedProtocolMessageType('QuotasInfo', (_message.Message,), {
  'DESCRIPTOR' : _QUOTASINFO,
  '__module__' : 'spaceone.api.notification.v1.quota_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.notification.v1.QuotasInfo)
  })
_sym_db.RegisterMessage(QuotasInfo)

QuotaStatQuery = _reflection.GeneratedProtocolMessageType('QuotaStatQuery', (_message.Message,), {
  'DESCRIPTOR' : _QUOTASTATQUERY,
  '__module__' : 'spaceone.api.notification.v1.quota_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.notification.v1.QuotaStatQuery)
  })
_sym_db.RegisterMessage(QuotaStatQuery)

_QUOTA = DESCRIPTOR.services_by_name['Quota']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/notification/v1'
  _QUOTA.methods_by_name['create']._options = None
  _QUOTA.methods_by_name['create']._serialized_options = b'\202\323\344\223\002\"\"\035/notification/v1/quota/create:\001*'
  _QUOTA.methods_by_name['update']._options = None
  _QUOTA.methods_by_name['update']._serialized_options = b'\202\323\344\223\002\"\"\035/notification/v1/quota/update:\001*'
  _QUOTA.methods_by_name['delete']._options = None
  _QUOTA.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\"\"\035/notification/v1/quota/delete:\001*'
  _QUOTA.methods_by_name['get']._options = None
  _QUOTA.methods_by_name['get']._serialized_options = b'\202\323\344\223\002\037\"\032/notification/v1/quota/get:\001*'
  _QUOTA.methods_by_name['list']._options = None
  _QUOTA.methods_by_name['list']._serialized_options = b'\202\323\344\223\002 \"\033/notification/v1/quota/list:\001*'
  _QUOTA.methods_by_name['stat']._options = None
  _QUOTA.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002 \"\033/notification/v1/quota/stat:\001*'
  _CREATEQUOTAREQUEST._serialized_start=197
  _CREATEQUOTAREQUEST._serialized_end=297
  _UPDATEQUOTAREQUEST._serialized_start=299
  _UPDATEQUOTAREQUEST._serialized_end=396
  _QUOTAREQUEST._serialized_start=398
  _QUOTAREQUEST._serialized_end=449
  _QUOTAQUERY._serialized_start=451
  _QUOTAQUERY._serialized_end=565
  _QUOTAINFO._serialized_start=567
  _QUOTAINFO._serialized_end=676
  _QUOTASINFO._serialized_start=678
  _QUOTASINFO._serialized_end=769
  _QUOTASTATQUERY._serialized_start=771
  _QUOTASTATQUERY._serialized_end=860
  _QUOTA._serialized_start=863
  _QUOTA._serialized_end=1662
# @@protoc_insertion_point(module_scope)
