# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/monitoring/v1/escalation_policy.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n2spaceone/api/monitoring/v1/escalation_policy.proto\x12\x1aspaceone.api.monitoring.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"\xe5\x01\n\x14\x45scalationPolicyRule\x12^\n\x12notification_level\x18\x01 \x01(\x0e\x32\x42.spaceone.api.monitoring.v1.EscalationPolicyRule.NotificationLevel\x12\x18\n\x10\x65scalate_minutes\x18\x02 \x01(\x05\"S\n\x11NotificationLevel\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03\x41LL\x10\x01\x12\x07\n\x03LV1\x10\x02\x12\x07\n\x03LV2\x10\x03\x12\x07\n\x03LV3\x10\x04\x12\x07\n\x03LV4\x10\x05\x12\x07\n\x03LV5\x10\x06\"\x88\x03\n\x1d\x43reateEscalationPolicyRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12?\n\x05rules\x18\x02 \x03(\x0b\x32\x30.spaceone.api.monitoring.v1.EscalationPolicyRule\x12\x14\n\x0crepeat_count\x18\x03 \x01(\x05\x12m\n\x10\x66inish_condition\x18\x04 \x01(\x0e\x32S.spaceone.api.monitoring.v1.CreateEscalationPolicyRequest.EscalationFinishCondition\x12\x12\n\nproject_id\x18\x05 \x01(\t\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"E\n\x19\x45scalationFinishCondition\x12\x08\n\x04NONE\x10\x00\x12\x10\n\x0c\x41\x43KNOWLEDGED\x10\x01\x12\x0c\n\x08RESOLVED\x10\x02\"\x92\x03\n\x1dUpdateEscalationPolicyRequest\x12\x1c\n\x14\x65scalation_policy_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12?\n\x05rules\x18\x03 \x03(\x0b\x32\x30.spaceone.api.monitoring.v1.EscalationPolicyRule\x12\x14\n\x0crepeat_count\x18\x04 \x01(\x05\x12m\n\x10\x66inish_condition\x18\x05 \x01(\x0e\x32S.spaceone.api.monitoring.v1.UpdateEscalationPolicyRequest.EscalationFinishCondition\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"E\n\x19\x45scalationFinishCondition\x12\x08\n\x04NONE\x10\x00\x12\x10\n\x0c\x41\x43KNOWLEDGED\x10\x01\x12\x0c\n\x08RESOLVED\x10\x02\"J\n\x17\x45scalationPolicyRequest\x12\x1c\n\x14\x65scalation_policy_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"[\n\x1aGetEscalationPolicyRequest\x12\x1c\n\x14\x65scalation_policy_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\xf2\x03\n\x15\x45scalationPolicyQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x1c\n\x14\x65scalation_policy_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x12\n\nis_default\x18\x04 \x01(\x08\x12\x65\n\x10\x66inish_condition\x18\x05 \x01(\x0e\x32K.spaceone.api.monitoring.v1.EscalationPolicyQuery.EscalationFinishCondition\x12V\n\x05scope\x18\x06 \x01(\x0e\x32G.spaceone.api.monitoring.v1.EscalationPolicyQuery.EscalationPolicyScope\x12\x12\n\nproject_id\x18\x07 \x01(\t\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"@\n\x15\x45scalationPolicyScope\x12\x0e\n\nSCOPE_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\x0b\n\x07PROJECT\x10\x02\"E\n\x19\x45scalationFinishCondition\x12\x08\n\x04NONE\x10\x00\x12\x10\n\x0c\x41\x43KNOWLEDGED\x10\x01\x12\x0c\n\x08RESOLVED\x10\x02\"\xd5\x04\n\x14\x45scalationPolicyInfo\x12\x1c\n\x14\x65scalation_policy_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x12\n\nis_default\x18\x03 \x01(\x08\x12?\n\x05rules\x18\x04 \x03(\x0b\x32\x30.spaceone.api.monitoring.v1.EscalationPolicyRule\x12\x14\n\x0crepeat_count\x18\x05 \x01(\x05\x12\x64\n\x10\x66inish_condition\x18\x06 \x01(\x0e\x32J.spaceone.api.monitoring.v1.EscalationPolicyInfo.EscalationFinishCondition\x12U\n\x05scope\x18\x07 \x01(\x0e\x32\x46.spaceone.api.monitoring.v1.EscalationPolicyInfo.EscalationPolicyScope\x12\x12\n\nproject_id\x18\x08 \x01(\t\x12%\n\x04tags\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x0b \x01(\t\x12\x12\n\ncreated_at\x18\x15 \x01(\t\"@\n\x15\x45scalationPolicyScope\x12\x0e\n\nSCOPE_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\x0b\n\x07PROJECT\x10\x02\"E\n\x19\x45scalationFinishCondition\x12\x08\n\x04NONE\x10\x00\x12\x10\n\x0c\x41\x43KNOWLEDGED\x10\x01\x12\x0c\n\x08RESOLVED\x10\x02\"p\n\x16\x45scalationPoliciesInfo\x12\x41\n\x07results\x18\x01 \x03(\x0b\x32\x30.spaceone.api.monitoring.v1.EscalationPolicyInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"d\n\x19\x45scalationPolicyStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t2\xf6\x08\n\x10\x45scalationPolicy\x12\xa9\x01\n\x06\x63reate\x12\x39.spaceone.api.monitoring.v1.CreateEscalationPolicyRequest\x1a\x30.spaceone.api.monitoring.v1.EscalationPolicyInfo\"2\x82\xd3\xe4\x93\x02,\"\'/monitoring/v1/escalation-policy/create:\x01*\x12\xa9\x01\n\x06update\x12\x39.spaceone.api.monitoring.v1.UpdateEscalationPolicyRequest\x1a\x30.spaceone.api.monitoring.v1.EscalationPolicyInfo\"2\x82\xd3\xe4\x93\x02,\"\'/monitoring/v1/escalation-policy/update:\x01*\x12\xad\x01\n\x0bset_default\x12\x33.spaceone.api.monitoring.v1.EscalationPolicyRequest\x1a\x30.spaceone.api.monitoring.v1.EscalationPolicyInfo\"7\x82\xd3\xe4\x93\x02\x31\",/monitoring/v1/escalation-policy/set-default:\x01*\x12\x89\x01\n\x06\x64\x65lete\x12\x33.spaceone.api.monitoring.v1.EscalationPolicyRequest\x1a\x16.google.protobuf.Empty\"2\x82\xd3\xe4\x93\x02,\"\'/monitoring/v1/escalation-policy/delete:\x01*\x12\xa0\x01\n\x03get\x12\x36.spaceone.api.monitoring.v1.GetEscalationPolicyRequest\x1a\x30.spaceone.api.monitoring.v1.EscalationPolicyInfo\"/\x82\xd3\xe4\x93\x02)\"$/monitoring/v1/escalation-policy/get:\x01*\x12\x9f\x01\n\x04list\x12\x31.spaceone.api.monitoring.v1.EscalationPolicyQuery\x1a\x32.spaceone.api.monitoring.v1.EscalationPoliciesInfo\"0\x82\xd3\xe4\x93\x02*\"%/monitoring/v1/escalation-policy/list:\x01*\x12\x88\x01\n\x04stat\x12\x35.spaceone.api.monitoring.v1.EscalationPolicyStatQuery\x1a\x17.google.protobuf.Struct\"0\x82\xd3\xe4\x93\x02*\"%/monitoring/v1/escalation-policy/stat:\x01*BAZ?github.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/v1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.monitoring.v1.escalation_policy_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z?github.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/v1'
  _ESCALATIONPOLICY.methods_by_name['create']._options = None
  _ESCALATIONPOLICY.methods_by_name['create']._serialized_options = b'\202\323\344\223\002,\"\'/monitoring/v1/escalation-policy/create:\001*'
  _ESCALATIONPOLICY.methods_by_name['update']._options = None
  _ESCALATIONPOLICY.methods_by_name['update']._serialized_options = b'\202\323\344\223\002,\"\'/monitoring/v1/escalation-policy/update:\001*'
  _ESCALATIONPOLICY.methods_by_name['set_default']._options = None
  _ESCALATIONPOLICY.methods_by_name['set_default']._serialized_options = b'\202\323\344\223\0021\",/monitoring/v1/escalation-policy/set-default:\001*'
  _ESCALATIONPOLICY.methods_by_name['delete']._options = None
  _ESCALATIONPOLICY.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002,\"\'/monitoring/v1/escalation-policy/delete:\001*'
  _ESCALATIONPOLICY.methods_by_name['get']._options = None
  _ESCALATIONPOLICY.methods_by_name['get']._serialized_options = b'\202\323\344\223\002)\"$/monitoring/v1/escalation-policy/get:\001*'
  _ESCALATIONPOLICY.methods_by_name['list']._options = None
  _ESCALATIONPOLICY.methods_by_name['list']._serialized_options = b'\202\323\344\223\002*\"%/monitoring/v1/escalation-policy/list:\001*'
  _ESCALATIONPOLICY.methods_by_name['stat']._options = None
  _ESCALATIONPOLICY.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002*\"%/monitoring/v1/escalation-policy/stat:\001*'
  _ESCALATIONPOLICYRULE._serialized_start=206
  _ESCALATIONPOLICYRULE._serialized_end=435
  _ESCALATIONPOLICYRULE_NOTIFICATIONLEVEL._serialized_start=352
  _ESCALATIONPOLICYRULE_NOTIFICATIONLEVEL._serialized_end=435
  _CREATEESCALATIONPOLICYREQUEST._serialized_start=438
  _CREATEESCALATIONPOLICYREQUEST._serialized_end=830
  _CREATEESCALATIONPOLICYREQUEST_ESCALATIONFINISHCONDITION._serialized_start=761
  _CREATEESCALATIONPOLICYREQUEST_ESCALATIONFINISHCONDITION._serialized_end=830
  _UPDATEESCALATIONPOLICYREQUEST._serialized_start=833
  _UPDATEESCALATIONPOLICYREQUEST._serialized_end=1235
  _UPDATEESCALATIONPOLICYREQUEST_ESCALATIONFINISHCONDITION._serialized_start=761
  _UPDATEESCALATIONPOLICYREQUEST_ESCALATIONFINISHCONDITION._serialized_end=830
  _ESCALATIONPOLICYREQUEST._serialized_start=1237
  _ESCALATIONPOLICYREQUEST._serialized_end=1311
  _GETESCALATIONPOLICYREQUEST._serialized_start=1313
  _GETESCALATIONPOLICYREQUEST._serialized_end=1404
  _ESCALATIONPOLICYQUERY._serialized_start=1407
  _ESCALATIONPOLICYQUERY._serialized_end=1905
  _ESCALATIONPOLICYQUERY_ESCALATIONPOLICYSCOPE._serialized_start=1770
  _ESCALATIONPOLICYQUERY_ESCALATIONPOLICYSCOPE._serialized_end=1834
  _ESCALATIONPOLICYQUERY_ESCALATIONFINISHCONDITION._serialized_start=761
  _ESCALATIONPOLICYQUERY_ESCALATIONFINISHCONDITION._serialized_end=830
  _ESCALATIONPOLICYINFO._serialized_start=1908
  _ESCALATIONPOLICYINFO._serialized_end=2505
  _ESCALATIONPOLICYINFO_ESCALATIONPOLICYSCOPE._serialized_start=1770
  _ESCALATIONPOLICYINFO_ESCALATIONPOLICYSCOPE._serialized_end=1834
  _ESCALATIONPOLICYINFO_ESCALATIONFINISHCONDITION._serialized_start=761
  _ESCALATIONPOLICYINFO_ESCALATIONFINISHCONDITION._serialized_end=830
  _ESCALATIONPOLICIESINFO._serialized_start=2507
  _ESCALATIONPOLICIESINFO._serialized_end=2619
  _ESCALATIONPOLICYSTATQUERY._serialized_start=2621
  _ESCALATIONPOLICYSTATQUERY._serialized_end=2721
  _ESCALATIONPOLICY._serialized_start=2724
  _ESCALATIONPOLICY._serialized_end=3866
# @@protoc_insertion_point(module_scope)