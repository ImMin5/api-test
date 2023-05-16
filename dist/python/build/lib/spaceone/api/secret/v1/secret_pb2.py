# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/secret/v1/secret.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#spaceone/api/secret/v1/secret.proto\x12\x16spaceone.api.secret.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"\x98\x02\n\x13\x43reateSecretRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x37\n\x0bsecret_type\x18\x03 \x01(\x0e\x32\".spaceone.api.secret.v1.SecretType\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0e\n\x06schema\x18\x05 \x01(\t\x12\x1a\n\x12service_account_id\x18\x06 \x01(\t\x12\x12\n\nproject_id\x18\x07 \x01(\t\x12\x19\n\x11trusted_secret_id\x18\x0b \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"\x9d\x01\n\x13UpdateSecretRequest\x12\x11\n\tsecret_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nproject_id\x18\x04 \x01(\t\x12\x11\n\tdomain_id\x18\x05 \x01(\t\x12\x17\n\x0frelease_project\x18\x06 \x01(\x08\"5\n\rSecretRequest\x12\x11\n\tsecret_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"v\n\x17UpdateSecretDataRequest\x12\x11\n\tsecret_id\x18\x01 \x01(\t\x12%\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x03 \x01(\t\x12\x0e\n\x06schema\x18\x04 \x01(\t\"F\n\x10GetSecretRequest\x12\x11\n\tsecret_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\xff\x01\n\x0bSecretQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x11\n\tsecret_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x37\n\x0bsecret_type\x18\x04 \x01(\x0e\x32\".spaceone.api.secret.v1.SecretType\x12\x0e\n\x06schema\x18\x06 \x01(\t\x12\x10\n\x08provider\x18\x07 \x01(\t\x12\x1a\n\x12service_account_id\x18\x08 \x01(\t\x12\x19\n\x11trusted_secret_id\x18\t \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"|\n\x0eSecretDataInfo\x12%\n\x04\x64\x61ta\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tencrypted\x18\x02 \x01(\x08\x12\x30\n\x0f\x65ncrypt_options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"\xa1\x02\n\nSecretInfo\x12\x11\n\tsecret_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x37\n\x0bsecret_type\x18\x03 \x01(\x0e\x32\".spaceone.api.secret.v1.SecretType\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0e\n\x06schema\x18\x06 \x01(\t\x12\x10\n\x08provider\x18\x07 \x01(\t\x12\x1a\n\x12service_account_id\x18\x08 \x01(\t\x12\x12\n\nproject_id\x18\t \x01(\t\x12\x19\n\x11trusted_secret_id\x18\n \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x16 \x01(\t\"W\n\x0bSecretsInfo\x12\x33\n\x07results\x18\x01 \x03(\x0b\x32\".spaceone.api.secret.v1.SecretInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"Z\n\x0fSecretStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t*\'\n\nSecretType\x12\x08\n\x04NONE\x10\x00\x12\x0f\n\x0b\x43REDENTIALS\x10\x01\x32\xd5\x07\n\x06Secret\x12~\n\x06\x63reate\x12+.spaceone.api.secret.v1.CreateSecretRequest\x1a\".spaceone.api.secret.v1.SecretInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/secret/v1/secret/create:\x01*\x12~\n\x06update\x12+.spaceone.api.secret.v1.UpdateSecretRequest\x1a\".spaceone.api.secret.v1.SecretInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/secret/v1/secret/update:\x01*\x12l\n\x06\x64\x65lete\x12%.spaceone.api.secret.v1.SecretRequest\x1a\x16.google.protobuf.Empty\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/secret/v1/secret/delete:\x01*\x12\x80\x01\n\x0bupdate_data\x12/.spaceone.api.secret.v1.UpdateSecretDataRequest\x1a\x16.google.protobuf.Empty\"(\x82\xd3\xe4\x93\x02\"\"\x1d/secret/v1/secret/update-data:\x01*\x12\x80\x01\n\x08get_data\x12%.spaceone.api.secret.v1.SecretRequest\x1a&.spaceone.api.secret.v1.SecretDataInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/secret/v1/secret/get-data:\x01*\x12u\n\x03get\x12(.spaceone.api.secret.v1.GetSecretRequest\x1a\".spaceone.api.secret.v1.SecretInfo\" \x82\xd3\xe4\x93\x02\x1a\"\x15/secret/v1/secret/get:\x01*\x12s\n\x04list\x12#.spaceone.api.secret.v1.SecretQuery\x1a#.spaceone.api.secret.v1.SecretsInfo\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/secret/v1/secret/list:\x01*\x12k\n\x04stat\x12\'.spaceone.api.secret.v1.SecretStatQuery\x1a\x17.google.protobuf.Struct\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/secret/v1/secret/stat:\x01*B=Z;github.com/cloudforet-io/api/dist/go/spaceone/api/secret/v1b\x06proto3')

_SECRETTYPE = DESCRIPTOR.enum_types_by_name['SecretType']
SecretType = enum_type_wrapper.EnumTypeWrapper(_SECRETTYPE)
NONE = 0
CREDENTIALS = 1


_CREATESECRETREQUEST = DESCRIPTOR.message_types_by_name['CreateSecretRequest']
_UPDATESECRETREQUEST = DESCRIPTOR.message_types_by_name['UpdateSecretRequest']
_SECRETREQUEST = DESCRIPTOR.message_types_by_name['SecretRequest']
_UPDATESECRETDATAREQUEST = DESCRIPTOR.message_types_by_name['UpdateSecretDataRequest']
_GETSECRETREQUEST = DESCRIPTOR.message_types_by_name['GetSecretRequest']
_SECRETQUERY = DESCRIPTOR.message_types_by_name['SecretQuery']
_SECRETDATAINFO = DESCRIPTOR.message_types_by_name['SecretDataInfo']
_SECRETINFO = DESCRIPTOR.message_types_by_name['SecretInfo']
_SECRETSINFO = DESCRIPTOR.message_types_by_name['SecretsInfo']
_SECRETSTATQUERY = DESCRIPTOR.message_types_by_name['SecretStatQuery']
CreateSecretRequest = _reflection.GeneratedProtocolMessageType('CreateSecretRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATESECRETREQUEST,
  '__module__' : 'spaceone.api.secret.v1.secret_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.secret.v1.CreateSecretRequest)
  })
_sym_db.RegisterMessage(CreateSecretRequest)

UpdateSecretRequest = _reflection.GeneratedProtocolMessageType('UpdateSecretRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATESECRETREQUEST,
  '__module__' : 'spaceone.api.secret.v1.secret_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.secret.v1.UpdateSecretRequest)
  })
_sym_db.RegisterMessage(UpdateSecretRequest)

SecretRequest = _reflection.GeneratedProtocolMessageType('SecretRequest', (_message.Message,), {
  'DESCRIPTOR' : _SECRETREQUEST,
  '__module__' : 'spaceone.api.secret.v1.secret_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.secret.v1.SecretRequest)
  })
_sym_db.RegisterMessage(SecretRequest)

UpdateSecretDataRequest = _reflection.GeneratedProtocolMessageType('UpdateSecretDataRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATESECRETDATAREQUEST,
  '__module__' : 'spaceone.api.secret.v1.secret_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.secret.v1.UpdateSecretDataRequest)
  })
_sym_db.RegisterMessage(UpdateSecretDataRequest)

GetSecretRequest = _reflection.GeneratedProtocolMessageType('GetSecretRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETSECRETREQUEST,
  '__module__' : 'spaceone.api.secret.v1.secret_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.secret.v1.GetSecretRequest)
  })
_sym_db.RegisterMessage(GetSecretRequest)

SecretQuery = _reflection.GeneratedProtocolMessageType('SecretQuery', (_message.Message,), {
  'DESCRIPTOR' : _SECRETQUERY,
  '__module__' : 'spaceone.api.secret.v1.secret_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.secret.v1.SecretQuery)
  })
_sym_db.RegisterMessage(SecretQuery)

SecretDataInfo = _reflection.GeneratedProtocolMessageType('SecretDataInfo', (_message.Message,), {
  'DESCRIPTOR' : _SECRETDATAINFO,
  '__module__' : 'spaceone.api.secret.v1.secret_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.secret.v1.SecretDataInfo)
  })
_sym_db.RegisterMessage(SecretDataInfo)

SecretInfo = _reflection.GeneratedProtocolMessageType('SecretInfo', (_message.Message,), {
  'DESCRIPTOR' : _SECRETINFO,
  '__module__' : 'spaceone.api.secret.v1.secret_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.secret.v1.SecretInfo)
  })
_sym_db.RegisterMessage(SecretInfo)

SecretsInfo = _reflection.GeneratedProtocolMessageType('SecretsInfo', (_message.Message,), {
  'DESCRIPTOR' : _SECRETSINFO,
  '__module__' : 'spaceone.api.secret.v1.secret_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.secret.v1.SecretsInfo)
  })
_sym_db.RegisterMessage(SecretsInfo)

SecretStatQuery = _reflection.GeneratedProtocolMessageType('SecretStatQuery', (_message.Message,), {
  'DESCRIPTOR' : _SECRETSTATQUERY,
  '__module__' : 'spaceone.api.secret.v1.secret_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.secret.v1.SecretStatQuery)
  })
_sym_db.RegisterMessage(SecretStatQuery)

_SECRET = DESCRIPTOR.services_by_name['Secret']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z;github.com/cloudforet-io/api/dist/go/spaceone/api/secret/v1'
  _SECRET.methods_by_name['create']._options = None
  _SECRET.methods_by_name['create']._serialized_options = b'\202\323\344\223\002\035\"\030/secret/v1/secret/create:\001*'
  _SECRET.methods_by_name['update']._options = None
  _SECRET.methods_by_name['update']._serialized_options = b'\202\323\344\223\002\035\"\030/secret/v1/secret/update:\001*'
  _SECRET.methods_by_name['delete']._options = None
  _SECRET.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\035\"\030/secret/v1/secret/delete:\001*'
  _SECRET.methods_by_name['update_data']._options = None
  _SECRET.methods_by_name['update_data']._serialized_options = b'\202\323\344\223\002\"\"\035/secret/v1/secret/update-data:\001*'
  _SECRET.methods_by_name['get_data']._options = None
  _SECRET.methods_by_name['get_data']._serialized_options = b'\202\323\344\223\002\037\"\032/secret/v1/secret/get-data:\001*'
  _SECRET.methods_by_name['get']._options = None
  _SECRET.methods_by_name['get']._serialized_options = b'\202\323\344\223\002\032\"\025/secret/v1/secret/get:\001*'
  _SECRET.methods_by_name['list']._options = None
  _SECRET.methods_by_name['list']._serialized_options = b'\202\323\344\223\002\033\"\026/secret/v1/secret/list:\001*'
  _SECRET.methods_by_name['stat']._options = None
  _SECRET.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\033\"\026/secret/v1/secret/stat:\001*'
  _SECRETTYPE._serialized_start=1733
  _SECRETTYPE._serialized_end=1772
  _CREATESECRETREQUEST._serialized_start=187
  _CREATESECRETREQUEST._serialized_end=467
  _UPDATESECRETREQUEST._serialized_start=470
  _UPDATESECRETREQUEST._serialized_end=627
  _SECRETREQUEST._serialized_start=629
  _SECRETREQUEST._serialized_end=682
  _UPDATESECRETDATAREQUEST._serialized_start=684
  _UPDATESECRETDATAREQUEST._serialized_end=802
  _GETSECRETREQUEST._serialized_start=804
  _GETSECRETREQUEST._serialized_end=874
  _SECRETQUERY._serialized_start=877
  _SECRETQUERY._serialized_end=1132
  _SECRETDATAINFO._serialized_start=1134
  _SECRETDATAINFO._serialized_end=1258
  _SECRETINFO._serialized_start=1261
  _SECRETINFO._serialized_end=1550
  _SECRETSINFO._serialized_start=1552
  _SECRETSINFO._serialized_end=1639
  _SECRETSTATQUERY._serialized_start=1641
  _SECRETSTATQUERY._serialized_end=1731
  _SECRET._serialized_start=1775
  _SECRET._serialized_end=2756
# @@protoc_insertion_point(module_scope)
