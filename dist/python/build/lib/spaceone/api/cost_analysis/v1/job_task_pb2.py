# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/cost_analysis/v1/job_task.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n,spaceone/api/cost_analysis/v1/job_task.proto\x12\x1dspaceone.api.cost_analysis.v1\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"I\n\x11GetJobTaskRequest\x12\x13\n\x0bjob_task_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\xbb\x02\n\x0cJobTaskQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x13\n\x0bjob_task_id\x18\x02 \x01(\t\x12\x42\n\x06status\x18\x03 \x01(\x0e\x32\x32.spaceone.api.cost_analysis.v1.JobTaskQuery.Status\x12\x0e\n\x06job_id\x18\x0b \x01(\t\x12\x16\n\x0e\x64\x61ta_source_id\x18\x0c \x01(\t\x12\x11\n\tdomain_id\x18\r \x01(\t\"k\n\x06Status\x12\x0e\n\nSCOPE_NONE\x10\x00\x12\x0b\n\x07PENDING\x10\x01\x12\x0f\n\x0bIN_PROGRESS\x10\x02\x12\x0b\n\x07SUCCESS\x10\x03\x12\x0b\n\x07\x46\x41ILURE\x10\x04\x12\x0b\n\x07TIMEOUT\x10\x05\x12\x0c\n\x08\x43\x41NCELED\x10\x06\"\xca\x03\n\x0bJobTaskInfo\x12\x13\n\x0bjob_task_id\x18\x01 \x01(\t\x12\x41\n\x06status\x18\x02 \x01(\x0e\x32\x31.spaceone.api.cost_analysis.v1.JobTaskInfo.Status\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x15\n\rcreated_count\x18\x04 \x01(\x05\x12\x12\n\nerror_code\x18\x05 \x01(\t\x12\x15\n\rerror_message\x18\x06 \x01(\t\x12\x0e\n\x06job_id\x18\x0b \x01(\t\x12\x16\n\x0e\x64\x61ta_source_id\x18\x0c \x01(\t\x12\x11\n\tdomain_id\x18\r \x01(\t\x12\x12\n\ncreated_at\x18\x15 \x01(\t\x12\x12\n\nstarted_at\x18\x16 \x01(\t\x12\x12\n\nupdated_at\x18\x17 \x01(\t\x12\x13\n\x0b\x66inished_at\x18\x18 \x01(\t\"k\n\x06Status\x12\x0e\n\nSCOPE_NONE\x10\x00\x12\x0b\n\x07PENDING\x10\x01\x12\x0f\n\x0bIN_PROGRESS\x10\x02\x12\x0b\n\x07SUCCESS\x10\x03\x12\x0b\n\x07\x46\x41ILURE\x10\x04\x12\x0b\n\x07TIMEOUT\x10\x05\x12\x0c\n\x08\x43\x41NCELED\x10\x06\"`\n\x0cJobTasksInfo\x12;\n\x07results\x18\x01 \x03(\x0b\x32*.spaceone.api.cost_analysis.v1.JobTaskInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"[\n\x10JobTaskStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t2\xa7\x03\n\x07JobTask\x12\x8e\x01\n\x03get\x12\x30.spaceone.api.cost_analysis.v1.GetJobTaskRequest\x1a*.spaceone.api.cost_analysis.v1.JobTaskInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/cost-analysis/v1/job-task/get:\x01*\x12\x8c\x01\n\x04list\x12+.spaceone.api.cost_analysis.v1.JobTaskQuery\x1a+.spaceone.api.cost_analysis.v1.JobTasksInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/cost-analysis/v1/job-task/list:\x01*\x12|\n\x04stat\x12/.spaceone.api.cost_analysis.v1.JobTaskStatQuery\x1a\x17.google.protobuf.Struct\"*\x82\xd3\xe4\x93\x02$\"\x1f/cost-analysis/v1/job-task/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1b\x06proto3')



_GETJOBTASKREQUEST = DESCRIPTOR.message_types_by_name['GetJobTaskRequest']
_JOBTASKQUERY = DESCRIPTOR.message_types_by_name['JobTaskQuery']
_JOBTASKINFO = DESCRIPTOR.message_types_by_name['JobTaskInfo']
_JOBTASKSINFO = DESCRIPTOR.message_types_by_name['JobTasksInfo']
_JOBTASKSTATQUERY = DESCRIPTOR.message_types_by_name['JobTaskStatQuery']
_JOBTASKQUERY_STATUS = _JOBTASKQUERY.enum_types_by_name['Status']
_JOBTASKINFO_STATUS = _JOBTASKINFO.enum_types_by_name['Status']
GetJobTaskRequest = _reflection.GeneratedProtocolMessageType('GetJobTaskRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETJOBTASKREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.job_task_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.GetJobTaskRequest)
  })
_sym_db.RegisterMessage(GetJobTaskRequest)

JobTaskQuery = _reflection.GeneratedProtocolMessageType('JobTaskQuery', (_message.Message,), {
  'DESCRIPTOR' : _JOBTASKQUERY,
  '__module__' : 'spaceone.api.cost_analysis.v1.job_task_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.JobTaskQuery)
  })
_sym_db.RegisterMessage(JobTaskQuery)

JobTaskInfo = _reflection.GeneratedProtocolMessageType('JobTaskInfo', (_message.Message,), {
  'DESCRIPTOR' : _JOBTASKINFO,
  '__module__' : 'spaceone.api.cost_analysis.v1.job_task_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.JobTaskInfo)
  })
_sym_db.RegisterMessage(JobTaskInfo)

JobTasksInfo = _reflection.GeneratedProtocolMessageType('JobTasksInfo', (_message.Message,), {
  'DESCRIPTOR' : _JOBTASKSINFO,
  '__module__' : 'spaceone.api.cost_analysis.v1.job_task_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.JobTasksInfo)
  })
_sym_db.RegisterMessage(JobTasksInfo)

JobTaskStatQuery = _reflection.GeneratedProtocolMessageType('JobTaskStatQuery', (_message.Message,), {
  'DESCRIPTOR' : _JOBTASKSTATQUERY,
  '__module__' : 'spaceone.api.cost_analysis.v1.job_task_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.JobTaskStatQuery)
  })
_sym_db.RegisterMessage(JobTaskStatQuery)

_JOBTASK = DESCRIPTOR.services_by_name['JobTask']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1'
  _JOBTASK.methods_by_name['get']._options = None
  _JOBTASK.methods_by_name['get']._serialized_options = b'\202\323\344\223\002#\"\036/cost-analysis/v1/job-task/get:\001*'
  _JOBTASK.methods_by_name['list']._options = None
  _JOBTASK.methods_by_name['list']._serialized_options = b'\202\323\344\223\002$\"\037/cost-analysis/v1/job-task/list:\001*'
  _JOBTASK.methods_by_name['stat']._options = None
  _JOBTASK.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002$\"\037/cost-analysis/v1/job-task/stat:\001*'
  _GETJOBTASKREQUEST._serialized_start=173
  _GETJOBTASKREQUEST._serialized_end=246
  _JOBTASKQUERY._serialized_start=249
  _JOBTASKQUERY._serialized_end=564
  _JOBTASKQUERY_STATUS._serialized_start=457
  _JOBTASKQUERY_STATUS._serialized_end=564
  _JOBTASKINFO._serialized_start=567
  _JOBTASKINFO._serialized_end=1025
  _JOBTASKINFO_STATUS._serialized_start=457
  _JOBTASKINFO_STATUS._serialized_end=564
  _JOBTASKSINFO._serialized_start=1027
  _JOBTASKSINFO._serialized_end=1123
  _JOBTASKSTATQUERY._serialized_start=1125
  _JOBTASKSTATQUERY._serialized_end=1216
  _JOBTASK._serialized_start=1219
  _JOBTASK._serialized_end=1642
# @@protoc_insertion_point(module_scope)
