# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/monitoring/v1/webhook.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(spaceone/api/monitoring/v1/webhook.proto\x12\x1aspaceone.api.monitoring.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"\x8c\x02\n\x11WebhookPluginInfo\x12\x11\n\tplugin_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08metadata\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12O\n\x0cupgrade_mode\x18\x05 \x01(\x0e\x32\x39.spaceone.api.monitoring.v1.WebhookPluginInfo.UpgradeMode\"-\n\x0bUpgradeMode\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x08\n\x04\x41UTO\x10\x02\"\xb6\x01\n\x14\x43reateWebhookRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x42\n\x0bplugin_info\x18\x02 \x01(\x0b\x32-.spaceone.api.monitoring.v1.WebhookPluginInfo\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nproject_id\x18\x0b \x01(\t\x12\x11\n\tdomain_id\x18\x0c \x01(\t\"r\n\x14UpdateWebhookRequest\x12\x12\n\nwebhook_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"\x87\x02\n\x1aUpdateWebhookPluginRequest\x12\x12\n\nwebhook_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12X\n\x0cupgrade_mode\x18\x04 \x01(\x0e\x32\x42.spaceone.api.monitoring.v1.UpdateWebhookPluginRequest.UpgradeMode\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"-\n\x0bUpgradeMode\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x08\n\x04\x41UTO\x10\x02\"7\n\x0eWebhookRequest\x12\x12\n\nwebhook_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"H\n\x11GetWebhookRequest\x12\x12\n\nwebhook_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\xa7\x02\n\x0cWebhookQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x12\n\nwebhook_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x44\n\x05state\x18\x04 \x01(\x0e\x32\x35.spaceone.api.monitoring.v1.WebhookQuery.WebhookState\x12\x12\n\naccess_key\x18\x05 \x01(\t\x12\x13\n\x0bwebhook_url\x18\x06 \x01(\t\x12\x12\n\nproject_id\x18\x07 \x01(\t\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"3\n\x0cWebhookState\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"\xa5\x03\n\x0bWebhookInfo\x12\x12\n\nwebhook_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x43\n\x05state\x18\x03 \x01(\x0e\x32\x34.spaceone.api.monitoring.v1.WebhookInfo.WebhookState\x12\x12\n\naccess_key\x18\x04 \x01(\t\x12\x13\n\x0bwebhook_url\x18\x05 \x01(\t\x12+\n\ncapability\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x42\n\x0bplugin_info\x18\x07 \x01(\x0b\x32-.spaceone.api.monitoring.v1.WebhookPluginInfo\x12%\n\x04tags\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nproject_id\x18\x0b \x01(\t\x12\x11\n\tdomain_id\x18\x0c \x01(\t\x12\x12\n\ncreated_at\x18\x15 \x01(\t\"3\n\x0cWebhookState\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"]\n\x0cWebhooksInfo\x12\x38\n\x07results\x18\x01 \x03(\x0b\x32\'.spaceone.api.monitoring.v1.WebhookInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"[\n\x10WebhookStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t2\xe5\n\n\x07Webhook\x12\x8d\x01\n\x06\x63reate\x12\x30.spaceone.api.monitoring.v1.CreateWebhookRequest\x1a\'.spaceone.api.monitoring.v1.WebhookInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/monitoring/v1/webhook/create:\x01*\x12\x8d\x01\n\x06update\x12\x30.spaceone.api.monitoring.v1.UpdateWebhookRequest\x1a\'.spaceone.api.monitoring.v1.WebhookInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/monitoring/v1/webhook/update:\x01*\x12\xa1\x01\n\rupdate_plugin\x12\x36.spaceone.api.monitoring.v1.UpdateWebhookPluginRequest\x1a\'.spaceone.api.monitoring.v1.WebhookInfo\"/\x82\xd3\xe4\x93\x02)\"$/monitoring/v1/webhook/update-plugin:\x01*\x12\x84\x01\n\rverify_plugin\x12*.spaceone.api.monitoring.v1.WebhookRequest\x1a\x16.google.protobuf.Empty\"/\x82\xd3\xe4\x93\x02)\"$/monitoring/v1/webhook/verify-plugin:\x01*\x12\x87\x01\n\x06\x65nable\x12*.spaceone.api.monitoring.v1.WebhookRequest\x1a\'.spaceone.api.monitoring.v1.WebhookInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/monitoring/v1/webhook/enable:\x01*\x12\x89\x01\n\x07\x64isable\x12*.spaceone.api.monitoring.v1.WebhookRequest\x1a\'.spaceone.api.monitoring.v1.WebhookInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/monitoring/v1/webhook/disable:\x01*\x12v\n\x06\x64\x65lete\x12*.spaceone.api.monitoring.v1.WebhookRequest\x1a\x16.google.protobuf.Empty\"(\x82\xd3\xe4\x93\x02\"\"\x1d/monitoring/v1/webhook/delete:\x01*\x12\x84\x01\n\x03get\x12-.spaceone.api.monitoring.v1.GetWebhookRequest\x1a\'.spaceone.api.monitoring.v1.WebhookInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/monitoring/v1/webhook/get:\x01*\x12\x82\x01\n\x04list\x12(.spaceone.api.monitoring.v1.WebhookQuery\x1a(.spaceone.api.monitoring.v1.WebhooksInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/monitoring/v1/webhook/list:\x01*\x12u\n\x04stat\x12,.spaceone.api.monitoring.v1.WebhookStatQuery\x1a\x17.google.protobuf.Struct\"&\x82\xd3\xe4\x93\x02 \"\x1b/monitoring/v1/webhook/stat:\x01*BAZ?github.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/v1b\x06proto3')



_WEBHOOKPLUGININFO = DESCRIPTOR.message_types_by_name['WebhookPluginInfo']
_CREATEWEBHOOKREQUEST = DESCRIPTOR.message_types_by_name['CreateWebhookRequest']
_UPDATEWEBHOOKREQUEST = DESCRIPTOR.message_types_by_name['UpdateWebhookRequest']
_UPDATEWEBHOOKPLUGINREQUEST = DESCRIPTOR.message_types_by_name['UpdateWebhookPluginRequest']
_WEBHOOKREQUEST = DESCRIPTOR.message_types_by_name['WebhookRequest']
_GETWEBHOOKREQUEST = DESCRIPTOR.message_types_by_name['GetWebhookRequest']
_WEBHOOKQUERY = DESCRIPTOR.message_types_by_name['WebhookQuery']
_WEBHOOKINFO = DESCRIPTOR.message_types_by_name['WebhookInfo']
_WEBHOOKSINFO = DESCRIPTOR.message_types_by_name['WebhooksInfo']
_WEBHOOKSTATQUERY = DESCRIPTOR.message_types_by_name['WebhookStatQuery']
_WEBHOOKPLUGININFO_UPGRADEMODE = _WEBHOOKPLUGININFO.enum_types_by_name['UpgradeMode']
_UPDATEWEBHOOKPLUGINREQUEST_UPGRADEMODE = _UPDATEWEBHOOKPLUGINREQUEST.enum_types_by_name['UpgradeMode']
_WEBHOOKQUERY_WEBHOOKSTATE = _WEBHOOKQUERY.enum_types_by_name['WebhookState']
_WEBHOOKINFO_WEBHOOKSTATE = _WEBHOOKINFO.enum_types_by_name['WebhookState']
WebhookPluginInfo = _reflection.GeneratedProtocolMessageType('WebhookPluginInfo', (_message.Message,), {
  'DESCRIPTOR' : _WEBHOOKPLUGININFO,
  '__module__' : 'spaceone.api.monitoring.v1.webhook_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.monitoring.v1.WebhookPluginInfo)
  })
_sym_db.RegisterMessage(WebhookPluginInfo)

CreateWebhookRequest = _reflection.GeneratedProtocolMessageType('CreateWebhookRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEWEBHOOKREQUEST,
  '__module__' : 'spaceone.api.monitoring.v1.webhook_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.monitoring.v1.CreateWebhookRequest)
  })
_sym_db.RegisterMessage(CreateWebhookRequest)

UpdateWebhookRequest = _reflection.GeneratedProtocolMessageType('UpdateWebhookRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEWEBHOOKREQUEST,
  '__module__' : 'spaceone.api.monitoring.v1.webhook_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.monitoring.v1.UpdateWebhookRequest)
  })
_sym_db.RegisterMessage(UpdateWebhookRequest)

UpdateWebhookPluginRequest = _reflection.GeneratedProtocolMessageType('UpdateWebhookPluginRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEWEBHOOKPLUGINREQUEST,
  '__module__' : 'spaceone.api.monitoring.v1.webhook_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.monitoring.v1.UpdateWebhookPluginRequest)
  })
_sym_db.RegisterMessage(UpdateWebhookPluginRequest)

WebhookRequest = _reflection.GeneratedProtocolMessageType('WebhookRequest', (_message.Message,), {
  'DESCRIPTOR' : _WEBHOOKREQUEST,
  '__module__' : 'spaceone.api.monitoring.v1.webhook_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.monitoring.v1.WebhookRequest)
  })
_sym_db.RegisterMessage(WebhookRequest)

GetWebhookRequest = _reflection.GeneratedProtocolMessageType('GetWebhookRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETWEBHOOKREQUEST,
  '__module__' : 'spaceone.api.monitoring.v1.webhook_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.monitoring.v1.GetWebhookRequest)
  })
_sym_db.RegisterMessage(GetWebhookRequest)

WebhookQuery = _reflection.GeneratedProtocolMessageType('WebhookQuery', (_message.Message,), {
  'DESCRIPTOR' : _WEBHOOKQUERY,
  '__module__' : 'spaceone.api.monitoring.v1.webhook_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.monitoring.v1.WebhookQuery)
  })
_sym_db.RegisterMessage(WebhookQuery)

WebhookInfo = _reflection.GeneratedProtocolMessageType('WebhookInfo', (_message.Message,), {
  'DESCRIPTOR' : _WEBHOOKINFO,
  '__module__' : 'spaceone.api.monitoring.v1.webhook_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.monitoring.v1.WebhookInfo)
  })
_sym_db.RegisterMessage(WebhookInfo)

WebhooksInfo = _reflection.GeneratedProtocolMessageType('WebhooksInfo', (_message.Message,), {
  'DESCRIPTOR' : _WEBHOOKSINFO,
  '__module__' : 'spaceone.api.monitoring.v1.webhook_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.monitoring.v1.WebhooksInfo)
  })
_sym_db.RegisterMessage(WebhooksInfo)

WebhookStatQuery = _reflection.GeneratedProtocolMessageType('WebhookStatQuery', (_message.Message,), {
  'DESCRIPTOR' : _WEBHOOKSTATQUERY,
  '__module__' : 'spaceone.api.monitoring.v1.webhook_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.monitoring.v1.WebhookStatQuery)
  })
_sym_db.RegisterMessage(WebhookStatQuery)

_WEBHOOK = DESCRIPTOR.services_by_name['Webhook']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z?github.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/v1'
  _WEBHOOK.methods_by_name['create']._options = None
  _WEBHOOK.methods_by_name['create']._serialized_options = b'\202\323\344\223\002\"\"\035/monitoring/v1/webhook/create:\001*'
  _WEBHOOK.methods_by_name['update']._options = None
  _WEBHOOK.methods_by_name['update']._serialized_options = b'\202\323\344\223\002\"\"\035/monitoring/v1/webhook/update:\001*'
  _WEBHOOK.methods_by_name['update_plugin']._options = None
  _WEBHOOK.methods_by_name['update_plugin']._serialized_options = b'\202\323\344\223\002)\"$/monitoring/v1/webhook/update-plugin:\001*'
  _WEBHOOK.methods_by_name['verify_plugin']._options = None
  _WEBHOOK.methods_by_name['verify_plugin']._serialized_options = b'\202\323\344\223\002)\"$/monitoring/v1/webhook/verify-plugin:\001*'
  _WEBHOOK.methods_by_name['enable']._options = None
  _WEBHOOK.methods_by_name['enable']._serialized_options = b'\202\323\344\223\002\"\"\035/monitoring/v1/webhook/enable:\001*'
  _WEBHOOK.methods_by_name['disable']._options = None
  _WEBHOOK.methods_by_name['disable']._serialized_options = b'\202\323\344\223\002#\"\036/monitoring/v1/webhook/disable:\001*'
  _WEBHOOK.methods_by_name['delete']._options = None
  _WEBHOOK.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\"\"\035/monitoring/v1/webhook/delete:\001*'
  _WEBHOOK.methods_by_name['get']._options = None
  _WEBHOOK.methods_by_name['get']._serialized_options = b'\202\323\344\223\002\037\"\032/monitoring/v1/webhook/get:\001*'
  _WEBHOOK.methods_by_name['list']._options = None
  _WEBHOOK.methods_by_name['list']._serialized_options = b'\202\323\344\223\002 \"\033/monitoring/v1/webhook/list:\001*'
  _WEBHOOK.methods_by_name['stat']._options = None
  _WEBHOOK.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002 \"\033/monitoring/v1/webhook/stat:\001*'
  _WEBHOOKPLUGININFO._serialized_start=196
  _WEBHOOKPLUGININFO._serialized_end=464
  _WEBHOOKPLUGININFO_UPGRADEMODE._serialized_start=419
  _WEBHOOKPLUGININFO_UPGRADEMODE._serialized_end=464
  _CREATEWEBHOOKREQUEST._serialized_start=467
  _CREATEWEBHOOKREQUEST._serialized_end=649
  _UPDATEWEBHOOKREQUEST._serialized_start=651
  _UPDATEWEBHOOKREQUEST._serialized_end=765
  _UPDATEWEBHOOKPLUGINREQUEST._serialized_start=768
  _UPDATEWEBHOOKPLUGINREQUEST._serialized_end=1031
  _UPDATEWEBHOOKPLUGINREQUEST_UPGRADEMODE._serialized_start=419
  _UPDATEWEBHOOKPLUGINREQUEST_UPGRADEMODE._serialized_end=464
  _WEBHOOKREQUEST._serialized_start=1033
  _WEBHOOKREQUEST._serialized_end=1088
  _GETWEBHOOKREQUEST._serialized_start=1090
  _GETWEBHOOKREQUEST._serialized_end=1162
  _WEBHOOKQUERY._serialized_start=1165
  _WEBHOOKQUERY._serialized_end=1460
  _WEBHOOKQUERY_WEBHOOKSTATE._serialized_start=1409
  _WEBHOOKQUERY_WEBHOOKSTATE._serialized_end=1460
  _WEBHOOKINFO._serialized_start=1463
  _WEBHOOKINFO._serialized_end=1884
  _WEBHOOKINFO_WEBHOOKSTATE._serialized_start=1409
  _WEBHOOKINFO_WEBHOOKSTATE._serialized_end=1460
  _WEBHOOKSINFO._serialized_start=1886
  _WEBHOOKSINFO._serialized_end=1979
  _WEBHOOKSTATQUERY._serialized_start=1981
  _WEBHOOKSTATQUERY._serialized_end=2072
  _WEBHOOK._serialized_start=2075
  _WEBHOOK._serialized_end=3456
# @@protoc_insertion_point(module_scope)
