# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/cost_analysis/v1/job.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\'spaceone/api/cost_analysis/v1/job.proto\x12\x1dspaceone.api.cost_analysis.v1\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"/\n\nJobRequest\x12\x0e\n\x06job_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"@\n\rGetJobRequest\x12\x0e\n\x06job_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\x91\x02\n\x08JobQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x0e\n\x06job_id\x18\x02 \x01(\t\x12>\n\x06status\x18\x03 \x01(\x0e\x32..spaceone.api.cost_analysis.v1.JobQuery.Status\x12\x16\n\x0e\x64\x61ta_source_id\x18\x0b \x01(\t\x12\x11\n\tdomain_id\x18\x0c \x01(\t\"^\n\x06Status\x12\x0e\n\nSCOPE_NONE\x10\x00\x12\x0f\n\x0bIN_PROGRESS\x10\x01\x12\x0b\n\x07SUCCESS\x10\x02\x12\x0b\n\x07\x46\x41ILURE\x10\x03\x12\x0b\n\x07TIMEOUT\x10\x04\x12\x0c\n\x08\x43\x41NCELED\x10\x05\"R\n\x0b\x43hangedInfo\x12\r\n\x05start\x18\x01 \x01(\t\x12\x0b\n\x03\x65nd\x18\x02 \x01(\t\x12\'\n\x06\x66ilter\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"\xdf\x03\n\x07JobInfo\x12\x0e\n\x06job_id\x18\x01 \x01(\t\x12=\n\x06status\x18\x02 \x01(\x0e\x32-.spaceone.api.cost_analysis.v1.JobInfo.Status\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nerror_code\x18\x04 \x01(\t\x12\x15\n\rerror_message\x18\x05 \x01(\t\x12\x13\n\x0btotal_tasks\x18\x06 \x01(\x05\x12\x16\n\x0eremained_tasks\x18\x07 \x01(\x05\x12\x16\n\x0e\x64\x61ta_source_id\x18\x0b \x01(\t\x12\x11\n\tdomain_id\x18\x0c \x01(\t\x12\x12\n\ncreated_at\x18\x15 \x01(\t\x12\x12\n\nupdated_at\x18\x16 \x01(\t\x12\x13\n\x0b\x66inished_at\x18\x17 \x01(\t\x12;\n\x07\x63hanged\x18\x18 \x03(\x0b\x32*.spaceone.api.cost_analysis.v1.ChangedInfo\"^\n\x06Status\x12\x0e\n\nSCOPE_NONE\x10\x00\x12\x0f\n\x0bIN_PROGRESS\x10\x01\x12\x0b\n\x07SUCCESS\x10\x02\x12\x0b\n\x07\x46\x41ILURE\x10\x03\x12\x0b\n\x07TIMEOUT\x10\x04\x12\x0c\n\x08\x43\x41NCELED\x10\x05\"X\n\x08JobsInfo\x12\x37\n\x07results\x18\x01 \x03(\x0b\x32&.spaceone.api.cost_analysis.v1.JobInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"W\n\x0cJobStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t2\x86\x04\n\x03Job\x12\x84\x01\n\x06\x63\x61ncel\x12).spaceone.api.cost_analysis.v1.JobRequest\x1a&.spaceone.api.cost_analysis.v1.JobInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/cost-analysis/v1/job/cancel:\x01*\x12\x81\x01\n\x03get\x12,.spaceone.api.cost_analysis.v1.GetJobRequest\x1a&.spaceone.api.cost_analysis.v1.JobInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/cost-analysis/v1/job/get:\x01*\x12\x7f\n\x04list\x12\'.spaceone.api.cost_analysis.v1.JobQuery\x1a\'.spaceone.api.cost_analysis.v1.JobsInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/cost-analysis/v1/job/list:\x01*\x12s\n\x04stat\x12+.spaceone.api.cost_analysis.v1.JobStatQuery\x1a\x17.google.protobuf.Struct\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/cost-analysis/v1/job/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.cost_analysis.v1.job_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1'
  _JOB.methods_by_name['cancel']._options = None
  _JOB.methods_by_name['cancel']._serialized_options = b'\202\323\344\223\002!\"\034/cost-analysis/v1/job/cancel:\001*'
  _JOB.methods_by_name['get']._options = None
  _JOB.methods_by_name['get']._serialized_options = b'\202\323\344\223\002\036\"\031/cost-analysis/v1/job/get:\001*'
  _JOB.methods_by_name['list']._options = None
  _JOB.methods_by_name['list']._serialized_options = b'\202\323\344\223\002\037\"\032/cost-analysis/v1/job/list:\001*'
  _JOB.methods_by_name['stat']._options = None
  _JOB.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\037\"\032/cost-analysis/v1/job/stat:\001*'
  _JOBREQUEST._serialized_start=168
  _JOBREQUEST._serialized_end=215
  _GETJOBREQUEST._serialized_start=217
  _GETJOBREQUEST._serialized_end=281
  _JOBQUERY._serialized_start=284
  _JOBQUERY._serialized_end=557
  _JOBQUERY_STATUS._serialized_start=463
  _JOBQUERY_STATUS._serialized_end=557
  _CHANGEDINFO._serialized_start=559
  _CHANGEDINFO._serialized_end=641
  _JOBINFO._serialized_start=644
  _JOBINFO._serialized_end=1123
  _JOBINFO_STATUS._serialized_start=463
  _JOBINFO_STATUS._serialized_end=557
  _JOBSINFO._serialized_start=1125
  _JOBSINFO._serialized_end=1213
  _JOBSTATQUERY._serialized_start=1215
  _JOBSTATQUERY._serialized_end=1302
  _JOB._serialized_start=1305
  _JOB._serialized_end=1823
# @@protoc_insertion_point(module_scope)
