# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/monitoring/plugin/event.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n*spaceone/api/monitoring/plugin/event.proto\x12\x1espaceone.api.monitoring.plugin\x1a\x1cgoogle/protobuf/struct.proto\"_\n\x0cParseRequest\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\"\xab\x03\n\tEventInfo\x12\x11\n\tevent_key\x18\x01 \x01(\t\x12\x12\n\nevent_type\x18\x02 \x01(\t\x12\r\n\x05title\x18\x03 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\x12\x44\n\x08severity\x18\x05 \x01(\x0e\x32\x32.spaceone.api.monitoring.plugin.EventInfo.Severity\x12)\n\x08resource\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0c\n\x04rule\x18\x07 \x01(\t\x12\x13\n\x0boccurred_at\x18\x08 \x01(\t\x12\x30\n\x0f\x61\x64\x64itional_info\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\timage_url\x18\n \x01(\t\x12\x10\n\x08provider\x18\x0b \x01(\t\x12\x0f\n\x07\x61\x63\x63ount\x18\x0c \x01(\t\"W\n\x08Severity\x12\x08\n\x04NONE\x10\x00\x12\x0c\n\x08\x43RITICAL\x10\x01\x12\t\n\x05\x45RROR\x10\x02\x12\x0b\n\x07WARNING\x10\x03\x12\x08\n\x04INFO\x10\x04\x12\x11\n\rNOT_AVAILABLE\x10\x05\"H\n\nEventsInfo\x12:\n\x07results\x18\x01 \x03(\x0b\x32).spaceone.api.monitoring.plugin.EventInfo2l\n\x05\x45vent\x12\x63\n\x05parse\x12,.spaceone.api.monitoring.plugin.ParseRequest\x1a*.spaceone.api.monitoring.plugin.EventsInfo\"\x00\x42\x45ZCgithub.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/pluginb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.monitoring.plugin.event_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZCgithub.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/plugin'
  _PARSEREQUEST._serialized_start=108
  _PARSEREQUEST._serialized_end=203
  _EVENTINFO._serialized_start=206
  _EVENTINFO._serialized_end=633
  _EVENTINFO_SEVERITY._serialized_start=546
  _EVENTINFO_SEVERITY._serialized_end=633
  _EVENTSINFO._serialized_start=635
  _EVENTSINFO._serialized_end=707
  _EVENT._serialized_start=709
  _EVENT._serialized_end=817
# @@protoc_insertion_point(module_scope)
