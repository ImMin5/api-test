# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/statistics/v1/schedule.proto
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
from spaceone.api.statistics.v1 import resource_pb2 as spaceone_dot_api_dot_statistics_dot_v1_dot_resource__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)spaceone/api/statistics/v1/schedule.proto\x12\x1aspaceone.api.statistics.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\x1a)spaceone/api/statistics/v1/resource.proto\"K\n\tScheduled\x12\x0c\n\x04\x63ron\x18\x01 \x01(\t\x12\x10\n\x08interval\x18\x02 \x01(\x05\x12\x0f\n\x07minutes\x18\x03 \x03(\x05\x12\r\n\x05hours\x18\x04 \x03(\x05\"\x7f\n\x0bQueryOption\x12<\n\taggregate\x18\x01 \x01(\x0b\x32).spaceone.api.statistics.v1.StatAggregate\x12\x32\n\x04page\x18\x02 \x01(\x0b\x32$.spaceone.api.statistics.v1.StatPage\"\xd1\x01\n\x12\x41\x64\x64ScheduleRequest\x12\r\n\x05topic\x18\x01 \x01(\t\x12(\n\x07options\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x37\n\x08schedule\x18\x03 \x01(\x0b\x32%.spaceone.api.statistics.v1.Scheduled\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0f\n\x07user_id\x18\x0b \x01(\t\x12\x11\n\tdomain_id\x18\x0c \x01(\t\"\xb3\x01\n\x15UpdateScheduleRequest\x12\x13\n\x0bschedule_id\x18\x01 \x01(\t\x12\x37\n\x08schedule\x18\x02 \x01(\x0b\x32%.spaceone.api.statistics.v1.Scheduled\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nstorage_id\x18\x04 \x01(\t\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"9\n\x0fScheduleRequest\x12\x13\n\x0bschedule_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"J\n\x12GetScheduleRequest\x12\x13\n\x0bschedule_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\x98\x01\n\rScheduleQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x13\n\x0bschedule_id\x18\x02 \x01(\t\x12\r\n\x05topic\x18\x03 \x01(\t\x12\r\n\x05state\x18\x04 \x01(\t\x12\x15\n\rresource_type\x18\x05 \x01(\t\x12\x11\n\tdomain_id\x18\x06 \x01(\t\"\xeb\x02\n\x0cScheduleInfo\x12\x13\n\x0bschedule_id\x18\x01 \x01(\t\x12\r\n\x05topic\x18\x02 \x01(\t\x12=\n\x05state\x18\x03 \x01(\x0e\x32..spaceone.api.statistics.v1.ScheduleInfo.State\x12(\n\x07options\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x37\n\x08schedule\x18\x05 \x01(\x0b\x32%.spaceone.api.statistics.v1.Scheduled\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x07 \x01(\t\x12\x12\n\ncreated_at\x18\x08 \x01(\t\x12\x19\n\x11last_scheduled_at\x18\t \x01(\t\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"_\n\rSchedulesInfo\x12\x39\n\x07results\x18\x01 \x03(\x0b\x32(.spaceone.api.statistics.v1.ScheduleInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"\\\n\x11ScheduleStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t2\xc5\x08\n\x08Schedule\x12\x87\x01\n\x03\x61\x64\x64\x12..spaceone.api.statistics.v1.AddScheduleRequest\x1a(.spaceone.api.statistics.v1.ScheduleInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/statistics/v1/schedule/add:\x01*\x12\x8d\x01\n\x06update\x12\x31.spaceone.api.statistics.v1.UpdateScheduleRequest\x1a(.spaceone.api.statistics.v1.ScheduleInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/statistics/schedule/update:\x01*\x12\x8a\x01\n\x06\x65nable\x12+.spaceone.api.statistics.v1.ScheduleRequest\x1a(.spaceone.api.statistics.v1.ScheduleInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/statistics/v1/schedule/enable:\x01*\x12\x8c\x01\n\x07\x64isable\x12+.spaceone.api.statistics.v1.ScheduleRequest\x1a(.spaceone.api.statistics.v1.ScheduleInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/statistics/v1/schedule/disable:\x01*\x12x\n\x06\x64\x65lete\x12+.spaceone.api.statistics.v1.ScheduleRequest\x1a\x16.google.protobuf.Empty\")\x82\xd3\xe4\x93\x02#\"\x1e/statistics/v1/schedule/delete:\x01*\x12\x87\x01\n\x03get\x12..spaceone.api.statistics.v1.GetScheduleRequest\x1a(.spaceone.api.statistics.v1.ScheduleInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/statistics/v1/schedule/get:\x01*\x12\x85\x01\n\x04list\x12).spaceone.api.statistics.v1.ScheduleQuery\x1a).spaceone.api.statistics.v1.SchedulesInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/statistics/v1/schedule/list:\x01*\x12w\n\x04stat\x12-.spaceone.api.statistics.v1.ScheduleStatQuery\x1a\x17.google.protobuf.Struct\"\'\x82\xd3\xe4\x93\x02!\"\x1c/statistics/v1/schedule/stat:\x01*BAZ?github.com/cloudforet-io/api/dist/go/spaceone/api/statistics/v1b\x06proto3')



_SCHEDULED = DESCRIPTOR.message_types_by_name['Scheduled']
_QUERYOPTION = DESCRIPTOR.message_types_by_name['QueryOption']
_ADDSCHEDULEREQUEST = DESCRIPTOR.message_types_by_name['AddScheduleRequest']
_UPDATESCHEDULEREQUEST = DESCRIPTOR.message_types_by_name['UpdateScheduleRequest']
_SCHEDULEREQUEST = DESCRIPTOR.message_types_by_name['ScheduleRequest']
_GETSCHEDULEREQUEST = DESCRIPTOR.message_types_by_name['GetScheduleRequest']
_SCHEDULEQUERY = DESCRIPTOR.message_types_by_name['ScheduleQuery']
_SCHEDULEINFO = DESCRIPTOR.message_types_by_name['ScheduleInfo']
_SCHEDULESINFO = DESCRIPTOR.message_types_by_name['SchedulesInfo']
_SCHEDULESTATQUERY = DESCRIPTOR.message_types_by_name['ScheduleStatQuery']
_SCHEDULEINFO_STATE = _SCHEDULEINFO.enum_types_by_name['State']
Scheduled = _reflection.GeneratedProtocolMessageType('Scheduled', (_message.Message,), {
  'DESCRIPTOR' : _SCHEDULED,
  '__module__' : 'spaceone.api.statistics.v1.schedule_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.statistics.v1.Scheduled)
  })
_sym_db.RegisterMessage(Scheduled)

QueryOption = _reflection.GeneratedProtocolMessageType('QueryOption', (_message.Message,), {
  'DESCRIPTOR' : _QUERYOPTION,
  '__module__' : 'spaceone.api.statistics.v1.schedule_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.statistics.v1.QueryOption)
  })
_sym_db.RegisterMessage(QueryOption)

AddScheduleRequest = _reflection.GeneratedProtocolMessageType('AddScheduleRequest', (_message.Message,), {
  'DESCRIPTOR' : _ADDSCHEDULEREQUEST,
  '__module__' : 'spaceone.api.statistics.v1.schedule_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.statistics.v1.AddScheduleRequest)
  })
_sym_db.RegisterMessage(AddScheduleRequest)

UpdateScheduleRequest = _reflection.GeneratedProtocolMessageType('UpdateScheduleRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATESCHEDULEREQUEST,
  '__module__' : 'spaceone.api.statistics.v1.schedule_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.statistics.v1.UpdateScheduleRequest)
  })
_sym_db.RegisterMessage(UpdateScheduleRequest)

ScheduleRequest = _reflection.GeneratedProtocolMessageType('ScheduleRequest', (_message.Message,), {
  'DESCRIPTOR' : _SCHEDULEREQUEST,
  '__module__' : 'spaceone.api.statistics.v1.schedule_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.statistics.v1.ScheduleRequest)
  })
_sym_db.RegisterMessage(ScheduleRequest)

GetScheduleRequest = _reflection.GeneratedProtocolMessageType('GetScheduleRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETSCHEDULEREQUEST,
  '__module__' : 'spaceone.api.statistics.v1.schedule_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.statistics.v1.GetScheduleRequest)
  })
_sym_db.RegisterMessage(GetScheduleRequest)

ScheduleQuery = _reflection.GeneratedProtocolMessageType('ScheduleQuery', (_message.Message,), {
  'DESCRIPTOR' : _SCHEDULEQUERY,
  '__module__' : 'spaceone.api.statistics.v1.schedule_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.statistics.v1.ScheduleQuery)
  })
_sym_db.RegisterMessage(ScheduleQuery)

ScheduleInfo = _reflection.GeneratedProtocolMessageType('ScheduleInfo', (_message.Message,), {
  'DESCRIPTOR' : _SCHEDULEINFO,
  '__module__' : 'spaceone.api.statistics.v1.schedule_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.statistics.v1.ScheduleInfo)
  })
_sym_db.RegisterMessage(ScheduleInfo)

SchedulesInfo = _reflection.GeneratedProtocolMessageType('SchedulesInfo', (_message.Message,), {
  'DESCRIPTOR' : _SCHEDULESINFO,
  '__module__' : 'spaceone.api.statistics.v1.schedule_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.statistics.v1.SchedulesInfo)
  })
_sym_db.RegisterMessage(SchedulesInfo)

ScheduleStatQuery = _reflection.GeneratedProtocolMessageType('ScheduleStatQuery', (_message.Message,), {
  'DESCRIPTOR' : _SCHEDULESTATQUERY,
  '__module__' : 'spaceone.api.statistics.v1.schedule_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.statistics.v1.ScheduleStatQuery)
  })
_sym_db.RegisterMessage(ScheduleStatQuery)

_SCHEDULE = DESCRIPTOR.services_by_name['Schedule']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z?github.com/cloudforet-io/api/dist/go/spaceone/api/statistics/v1'
  _SCHEDULE.methods_by_name['add']._options = None
  _SCHEDULE.methods_by_name['add']._serialized_options = b'\202\323\344\223\002 \"\033/statistics/v1/schedule/add:\001*'
  _SCHEDULE.methods_by_name['update']._options = None
  _SCHEDULE.methods_by_name['update']._serialized_options = b'\202\323\344\223\002 \"\033/statistics/schedule/update:\001*'
  _SCHEDULE.methods_by_name['enable']._options = None
  _SCHEDULE.methods_by_name['enable']._serialized_options = b'\202\323\344\223\002#\"\036/statistics/v1/schedule/enable:\001*'
  _SCHEDULE.methods_by_name['disable']._options = None
  _SCHEDULE.methods_by_name['disable']._serialized_options = b'\202\323\344\223\002$\"\037/statistics/v1/schedule/disable:\001*'
  _SCHEDULE.methods_by_name['delete']._options = None
  _SCHEDULE.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002#\"\036/statistics/v1/schedule/delete:\001*'
  _SCHEDULE.methods_by_name['get']._options = None
  _SCHEDULE.methods_by_name['get']._serialized_options = b'\202\323\344\223\002 \"\033/statistics/v1/schedule/get:\001*'
  _SCHEDULE.methods_by_name['list']._options = None
  _SCHEDULE.methods_by_name['list']._serialized_options = b'\202\323\344\223\002!\"\034/statistics/v1/schedule/list:\001*'
  _SCHEDULE.methods_by_name['stat']._options = None
  _SCHEDULE.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002!\"\034/statistics/v1/schedule/stat:\001*'
  _SCHEDULED._serialized_start=239
  _SCHEDULED._serialized_end=314
  _QUERYOPTION._serialized_start=316
  _QUERYOPTION._serialized_end=443
  _ADDSCHEDULEREQUEST._serialized_start=446
  _ADDSCHEDULEREQUEST._serialized_end=655
  _UPDATESCHEDULEREQUEST._serialized_start=658
  _UPDATESCHEDULEREQUEST._serialized_end=837
  _SCHEDULEREQUEST._serialized_start=839
  _SCHEDULEREQUEST._serialized_end=896
  _GETSCHEDULEREQUEST._serialized_start=898
  _GETSCHEDULEREQUEST._serialized_end=972
  _SCHEDULEQUERY._serialized_start=975
  _SCHEDULEQUERY._serialized_end=1127
  _SCHEDULEINFO._serialized_start=1130
  _SCHEDULEINFO._serialized_end=1493
  _SCHEDULEINFO_STATE._serialized_start=1449
  _SCHEDULEINFO_STATE._serialized_end=1493
  _SCHEDULESINFO._serialized_start=1495
  _SCHEDULESINFO._serialized_end=1590
  _SCHEDULESTATQUERY._serialized_start=1592
  _SCHEDULESTATQUERY._serialized_end=1684
  _SCHEDULE._serialized_start=1687
  _SCHEDULE._serialized_end=2780
# @@protoc_insertion_point(module_scope)
