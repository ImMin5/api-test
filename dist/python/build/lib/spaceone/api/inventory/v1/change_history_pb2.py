# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/inventory/v1/change_history.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.spaceone/api/inventory/v1/change_history.proto\x12\x19spaceone.api.inventory.v1\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"\xc2\x02\n\x12\x43hangeHistoryQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x18\n\x10\x63loud_service_id\x18\x02 \x01(\t\x12J\n\x06\x61\x63tion\x18\x03 \x01(\x0e\x32:.spaceone.api.inventory.v1.ChangeHistoryQuery.RecordAction\x12\x0f\n\x07user_id\x18\x04 \x01(\t\x12\x14\n\x0c\x63ollector_id\x18\x05 \x01(\t\x12\x0e\n\x06job_id\x18\x06 \x01(\t\x12\x12\n\nupdated_by\x18\x07 \x01(\t\x12\x11\n\tdomain_id\x18\x08 \x01(\t\"<\n\x0cRecordAction\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06\x43REATE\x10\x01\x12\n\n\x06UPDATE\x10\x02\x12\n\n\x06\x44\x45LETE\x10\x03\"\xeb\x02\n\nRecordInfo\x12\x11\n\trecord_id\x18\x01 \x01(\t\x12\x18\n\x10\x63loud_service_id\x18\x02 \x01(\t\x12\x42\n\x06\x61\x63tion\x18\x03 \x01(\x0e\x32\x32.spaceone.api.inventory.v1.RecordInfo.RecordAction\x12(\n\x04\x64iff\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12\x12\n\ndiff_count\x18\x05 \x01(\x05\x12\x0f\n\x07user_id\x18\x06 \x01(\t\x12\x14\n\x0c\x63ollector_id\x18\x07 \x01(\t\x12\x0e\n\x06job_id\x18\x08 \x01(\t\x12\x12\n\nupdated_by\x18\t \x01(\t\x12\x11\n\tdomain_id\x18\n \x01(\t\x12\x12\n\ncreated_at\x18\x0b \x01(\t\"<\n\x0cRecordAction\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06\x43REATE\x10\x01\x12\n\n\x06UPDATE\x10\x02\x12\n\n\x06\x44\x45LETE\x10\x03\"`\n\x11\x43hangeHistoryInfo\x12\x36\n\x07results\x18\x01 \x03(\x0b\x32%.spaceone.api.inventory.v1.RecordInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"{\n\x16\x43hangeHistoryStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x18\n\x10\x63loud_service_id\x18\x03 \x01(\t2\xa6\x02\n\rChangeHistory\x12\x91\x01\n\x04list\x12-.spaceone.api.inventory.v1.ChangeHistoryQuery\x1a,.spaceone.api.inventory.v1.ChangeHistoryInfo\",\x82\xd3\xe4\x93\x02&\"!/inventory/v1/change-history/list:\x01*\x12\x80\x01\n\x04stat\x12\x31.spaceone.api.inventory.v1.ChangeHistoryStatQuery\x1a\x17.google.protobuf.Struct\",\x82\xd3\xe4\x93\x02&\"!/inventory/v1/change-history/stat:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1b\x06proto3')



_CHANGEHISTORYQUERY = DESCRIPTOR.message_types_by_name['ChangeHistoryQuery']
_RECORDINFO = DESCRIPTOR.message_types_by_name['RecordInfo']
_CHANGEHISTORYINFO = DESCRIPTOR.message_types_by_name['ChangeHistoryInfo']
_CHANGEHISTORYSTATQUERY = DESCRIPTOR.message_types_by_name['ChangeHistoryStatQuery']
_CHANGEHISTORYQUERY_RECORDACTION = _CHANGEHISTORYQUERY.enum_types_by_name['RecordAction']
_RECORDINFO_RECORDACTION = _RECORDINFO.enum_types_by_name['RecordAction']
ChangeHistoryQuery = _reflection.GeneratedProtocolMessageType('ChangeHistoryQuery', (_message.Message,), {
  'DESCRIPTOR' : _CHANGEHISTORYQUERY,
  '__module__' : 'spaceone.api.inventory.v1.change_history_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.ChangeHistoryQuery)
  })
_sym_db.RegisterMessage(ChangeHistoryQuery)

RecordInfo = _reflection.GeneratedProtocolMessageType('RecordInfo', (_message.Message,), {
  'DESCRIPTOR' : _RECORDINFO,
  '__module__' : 'spaceone.api.inventory.v1.change_history_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.RecordInfo)
  })
_sym_db.RegisterMessage(RecordInfo)

ChangeHistoryInfo = _reflection.GeneratedProtocolMessageType('ChangeHistoryInfo', (_message.Message,), {
  'DESCRIPTOR' : _CHANGEHISTORYINFO,
  '__module__' : 'spaceone.api.inventory.v1.change_history_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.ChangeHistoryInfo)
  })
_sym_db.RegisterMessage(ChangeHistoryInfo)

ChangeHistoryStatQuery = _reflection.GeneratedProtocolMessageType('ChangeHistoryStatQuery', (_message.Message,), {
  'DESCRIPTOR' : _CHANGEHISTORYSTATQUERY,
  '__module__' : 'spaceone.api.inventory.v1.change_history_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.inventory.v1.ChangeHistoryStatQuery)
  })
_sym_db.RegisterMessage(ChangeHistoryStatQuery)

_CHANGEHISTORY = DESCRIPTOR.services_by_name['ChangeHistory']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1'
  _CHANGEHISTORY.methods_by_name['list']._options = None
  _CHANGEHISTORY.methods_by_name['list']._serialized_options = b'\202\323\344\223\002&\"!/inventory/v1/change-history/list:\001*'
  _CHANGEHISTORY.methods_by_name['stat']._options = None
  _CHANGEHISTORY.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002&\"!/inventory/v1/change-history/stat:\001*'
  _CHANGEHISTORYQUERY._serialized_start=172
  _CHANGEHISTORYQUERY._serialized_end=494
  _CHANGEHISTORYQUERY_RECORDACTION._serialized_start=434
  _CHANGEHISTORYQUERY_RECORDACTION._serialized_end=494
  _RECORDINFO._serialized_start=497
  _RECORDINFO._serialized_end=860
  _RECORDINFO_RECORDACTION._serialized_start=434
  _RECORDINFO_RECORDACTION._serialized_end=494
  _CHANGEHISTORYINFO._serialized_start=862
  _CHANGEHISTORYINFO._serialized_end=958
  _CHANGEHISTORYSTATQUERY._serialized_start=960
  _CHANGEHISTORYSTATQUERY._serialized_end=1083
  _CHANGEHISTORY._serialized_start=1086
  _CHANGEHISTORY._serialized_end=1380
# @@protoc_insertion_point(module_scope)
