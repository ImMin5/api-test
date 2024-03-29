# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/cost_analysis/v1/schedule.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n,spaceone/api/cost_analysis/v1/schedule.proto\x12\x1dspaceone.api.cost_analysis.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"K\n\tScheduled\x12\x0c\n\x04\x63ron\x18\x01 \x01(\t\x12\x10\n\x08interval\x18\x02 \x01(\x05\x12\x0f\n\x07minutes\x18\x03 \x03(\x05\x12\r\n\x05hours\x18\x04 \x03(\x05\"\xdd\x01\n\x15\x43reateScheduleRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12:\n\x08schedule\x18\x03 \x01(\x0b\x32(.spaceone.api.cost_analysis.v1.Scheduled\x12(\n\x07options\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x16\n\x0e\x64\x61ta_source_id\x18\x0b \x01(\t\x12\x11\n\tdomain_id\x18\x0c \x01(\t\"\xda\x01\n\x15UpdateScheduleRequest\x12\x13\n\x0bschedule_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12:\n\x08schedule\x18\x03 \x01(\x0b\x32(.spaceone.api.cost_analysis.v1.Scheduled\x12(\n\x07options\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"9\n\x0fScheduleRequest\x12\x13\n\x0bschedule_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"J\n\x12GetScheduleRequest\x12\x13\n\x0bschedule_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\x98\x01\n\rScheduleQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x13\n\x0bschedule_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\r\n\x05state\x18\x04 \x01(\t\x12\x16\n\x0e\x64\x61ta_source_id\x18\x05 \x01(\t\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"\x88\x03\n\x0cScheduleInfo\x12\x13\n\x0bschedule_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12@\n\x05state\x18\x03 \x01(\x0e\x32\x31.spaceone.api.cost_analysis.v1.ScheduleInfo.State\x12:\n\x08schedule\x18\x04 \x01(\x0b\x32(.spaceone.api.cost_analysis.v1.Scheduled\x12(\n\x07options\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x16\n\x0e\x64\x61ta_source_id\x18\x0b \x01(\t\x12\x11\n\tdomain_id\x18\x0c \x01(\t\x12\x12\n\ncreated_at\x18\x15 \x01(\t\x12\x19\n\x11last_scheduled_at\x18\x16 \x01(\t\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"b\n\rSchedulesInfo\x12<\n\x07results\x18\x01 \x03(\x0b\x32+.spaceone.api.cost_analysis.v1.ScheduleInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"\\\n\x11ScheduleStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t2\x93\t\n\x08Schedule\x12\x99\x01\n\x06\x63reate\x12\x34.spaceone.api.cost_analysis.v1.CreateScheduleRequest\x1a+.spaceone.api.cost_analysis.v1.ScheduleInfo\",\x82\xd3\xe4\x93\x02&\"!/cost-analysis/v1/schedule/create:\x01*\x12\x99\x01\n\x06update\x12\x34.spaceone.api.cost_analysis.v1.UpdateScheduleRequest\x1a+.spaceone.api.cost_analysis.v1.ScheduleInfo\",\x82\xd3\xe4\x93\x02&\"!/cost-analysis/v1/schedule/update:\x01*\x12\x93\x01\n\x06\x65nable\x12..spaceone.api.cost_analysis.v1.ScheduleRequest\x1a+.spaceone.api.cost_analysis.v1.ScheduleInfo\",\x82\xd3\xe4\x93\x02&\"!/cost-analysis/v1/schedule/enable:\x01*\x12\x95\x01\n\x07\x64isable\x12..spaceone.api.cost_analysis.v1.ScheduleRequest\x1a+.spaceone.api.cost_analysis.v1.ScheduleInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/cost-analysis/v1/schedule/disable:\x01*\x12~\n\x06\x64\x65lete\x12..spaceone.api.cost_analysis.v1.ScheduleRequest\x1a\x16.google.protobuf.Empty\",\x82\xd3\xe4\x93\x02&\"!/cost-analysis/v1/schedule/delete:\x01*\x12\x90\x01\n\x03get\x12\x31.spaceone.api.cost_analysis.v1.GetScheduleRequest\x1a+.spaceone.api.cost_analysis.v1.ScheduleInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/cost-analysis/v1/schedule/get:\x01*\x12\x8e\x01\n\x04list\x12,.spaceone.api.cost_analysis.v1.ScheduleQuery\x1a,.spaceone.api.cost_analysis.v1.SchedulesInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/cost-analysis/v1/schedule/list:\x01*\x12}\n\x04stat\x12\x30.spaceone.api.cost_analysis.v1.ScheduleStatQuery\x1a\x17.google.protobuf.Struct\"*\x82\xd3\xe4\x93\x02$\"\x1f/cost-analysis/v1/schedule/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.cost_analysis.v1.schedule_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1'
  _SCHEDULE.methods_by_name['create']._options = None
  _SCHEDULE.methods_by_name['create']._serialized_options = b'\202\323\344\223\002&\"!/cost-analysis/v1/schedule/create:\001*'
  _SCHEDULE.methods_by_name['update']._options = None
  _SCHEDULE.methods_by_name['update']._serialized_options = b'\202\323\344\223\002&\"!/cost-analysis/v1/schedule/update:\001*'
  _SCHEDULE.methods_by_name['enable']._options = None
  _SCHEDULE.methods_by_name['enable']._serialized_options = b'\202\323\344\223\002&\"!/cost-analysis/v1/schedule/enable:\001*'
  _SCHEDULE.methods_by_name['disable']._options = None
  _SCHEDULE.methods_by_name['disable']._serialized_options = b'\202\323\344\223\002\'\"\"/cost-analysis/v1/schedule/disable:\001*'
  _SCHEDULE.methods_by_name['delete']._options = None
  _SCHEDULE.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002&\"!/cost-analysis/v1/schedule/delete:\001*'
  _SCHEDULE.methods_by_name['get']._options = None
  _SCHEDULE.methods_by_name['get']._serialized_options = b'\202\323\344\223\002#\"\036/cost-analysis/v1/schedule/get:\001*'
  _SCHEDULE.methods_by_name['list']._options = None
  _SCHEDULE.methods_by_name['list']._serialized_options = b'\202\323\344\223\002$\"\037/cost-analysis/v1/schedule/list:\001*'
  _SCHEDULE.methods_by_name['stat']._options = None
  _SCHEDULE.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002$\"\037/cost-analysis/v1/schedule/stat:\001*'
  _SCHEDULED._serialized_start=202
  _SCHEDULED._serialized_end=277
  _CREATESCHEDULEREQUEST._serialized_start=280
  _CREATESCHEDULEREQUEST._serialized_end=501
  _UPDATESCHEDULEREQUEST._serialized_start=504
  _UPDATESCHEDULEREQUEST._serialized_end=722
  _SCHEDULEREQUEST._serialized_start=724
  _SCHEDULEREQUEST._serialized_end=781
  _GETSCHEDULEREQUEST._serialized_start=783
  _GETSCHEDULEREQUEST._serialized_end=857
  _SCHEDULEQUERY._serialized_start=860
  _SCHEDULEQUERY._serialized_end=1012
  _SCHEDULEINFO._serialized_start=1015
  _SCHEDULEINFO._serialized_end=1407
  _SCHEDULEINFO_STATE._serialized_start=1363
  _SCHEDULEINFO_STATE._serialized_end=1407
  _SCHEDULESINFO._serialized_start=1409
  _SCHEDULESINFO._serialized_end=1507
  _SCHEDULESTATQUERY._serialized_start=1509
  _SCHEDULESTATQUERY._serialized_end=1601
  _SCHEDULE._serialized_start=1604
  _SCHEDULE._serialized_end=2775
# @@protoc_insertion_point(module_scope)
