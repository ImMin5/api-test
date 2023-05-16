# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/dashboard/v1/custom_widget.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n-spaceone/api/dashboard/v1/custom_widget.proto\x12\x19spaceone.api.dashboard.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"\xac\x02\n\x19\x43reateCustomWidgetRequest\x12\x13\n\x0bwidget_name\x18\x01 \x01(\t\x12\r\n\x05title\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x30\n\x0finherit_options\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\x06labels\x18\n \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"\xb1\x02\n\x19UpdateCustomWidgetRequest\x12\x18\n\x10\x63ustom_widget_id\x18\x01 \x01(\t\x12\r\n\x05title\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x30\n\x0finherit_options\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\x06labels\x18\n \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"B\n\x13\x43ustomWidgetRequest\x12\x18\n\x10\x63ustom_widget_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"S\n\x16GetCustomWidgetRequest\x12\x18\n\x10\x63ustom_widget_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\xa1\x01\n\x11\x43ustomWidgetQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x18\n\x10\x63ustom_widget_id\x18\x02 \x01(\t\x12\x13\n\x0bwidget_name\x18\x03 \x01(\t\x12\r\n\x05title\x18\x04 \x01(\t\x12\x0f\n\x07user_id\x18\n \x01(\t\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"\x87\x03\n\x10\x43ustomWidgetInfo\x12\x18\n\x10\x63ustom_widget_id\x18\x01 \x01(\t\x12\x13\n\x0bwidget_name\x18\x02 \x01(\t\x12\r\n\x05title\x18\x03 \x01(\t\x12\x0f\n\x07version\x18\x04 \x01(\t\x12(\n\x07options\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x30\n\x0finherit_options\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\x06labels\x18\n \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0f\n\x07user_id\x18\x14 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x16 \x01(\t\x12\x12\n\nupdated_at\x18\x17 \x01(\t\"f\n\x11\x43ustomWidgetsInfo\x12<\n\x07results\x18\x01 \x03(\x0b\x32+.spaceone.api.dashboard.v1.CustomWidgetInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"`\n\x15\x43ustomWidgetStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t2\xef\x06\n\x0c\x43ustomWidget\x12\x9a\x01\n\x06\x63reate\x12\x34.spaceone.api.dashboard.v1.CreateCustomWidgetRequest\x1a+.spaceone.api.dashboard.v1.CustomWidgetInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/dashboard/v1/custom-widget/create:\x01*\x12\x9a\x01\n\x06update\x12\x34.spaceone.api.dashboard.v1.UpdateCustomWidgetRequest\x1a+.spaceone.api.dashboard.v1.CustomWidgetInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/dashboard/v1/custom-widget/update:\x01*\x12\x7f\n\x06\x64\x65lete\x12..spaceone.api.dashboard.v1.CustomWidgetRequest\x1a\x16.google.protobuf.Empty\"-\x82\xd3\xe4\x93\x02\'\"\"/dashboard/v1/custom-widget/delete:\x01*\x12\x91\x01\n\x03get\x12\x31.spaceone.api.dashboard.v1.GetCustomWidgetRequest\x1a+.spaceone.api.dashboard.v1.CustomWidgetInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/dashboard/v1/custom-widget/get:\x01*\x12\x8f\x01\n\x04list\x12,.spaceone.api.dashboard.v1.CustomWidgetQuery\x1a,.spaceone.api.dashboard.v1.CustomWidgetsInfo\"+\x82\xd3\xe4\x93\x02%\" /dashboard/v1/custom-widget/list:\x01*\x12~\n\x04stat\x12\x30.spaceone.api.dashboard.v1.CustomWidgetStatQuery\x1a\x17.google.protobuf.Struct\"+\x82\xd3\xe4\x93\x02%\" /dashboard/v1/custom-widget/stat:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.dashboard.v1.custom_widget_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1'
  _CUSTOMWIDGET.methods_by_name['create']._options = None
  _CUSTOMWIDGET.methods_by_name['create']._serialized_options = b'\202\323\344\223\002\'\"\"/dashboard/v1/custom-widget/create:\001*'
  _CUSTOMWIDGET.methods_by_name['update']._options = None
  _CUSTOMWIDGET.methods_by_name['update']._serialized_options = b'\202\323\344\223\002\'\"\"/dashboard/v1/custom-widget/update:\001*'
  _CUSTOMWIDGET.methods_by_name['delete']._options = None
  _CUSTOMWIDGET.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\'\"\"/dashboard/v1/custom-widget/delete:\001*'
  _CUSTOMWIDGET.methods_by_name['get']._options = None
  _CUSTOMWIDGET.methods_by_name['get']._serialized_options = b'\202\323\344\223\002$\"\037/dashboard/v1/custom-widget/get:\001*'
  _CUSTOMWIDGET.methods_by_name['list']._options = None
  _CUSTOMWIDGET.methods_by_name['list']._serialized_options = b'\202\323\344\223\002%\" /dashboard/v1/custom-widget/list:\001*'
  _CUSTOMWIDGET.methods_by_name['stat']._options = None
  _CUSTOMWIDGET.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002%\" /dashboard/v1/custom-widget/stat:\001*'
  _CREATECUSTOMWIDGETREQUEST._serialized_start=200
  _CREATECUSTOMWIDGETREQUEST._serialized_end=500
  _UPDATECUSTOMWIDGETREQUEST._serialized_start=503
  _UPDATECUSTOMWIDGETREQUEST._serialized_end=808
  _CUSTOMWIDGETREQUEST._serialized_start=810
  _CUSTOMWIDGETREQUEST._serialized_end=876
  _GETCUSTOMWIDGETREQUEST._serialized_start=878
  _GETCUSTOMWIDGETREQUEST._serialized_end=961
  _CUSTOMWIDGETQUERY._serialized_start=964
  _CUSTOMWIDGETQUERY._serialized_end=1125
  _CUSTOMWIDGETINFO._serialized_start=1128
  _CUSTOMWIDGETINFO._serialized_end=1519
  _CUSTOMWIDGETSINFO._serialized_start=1521
  _CUSTOMWIDGETSINFO._serialized_end=1623
  _CUSTOMWIDGETSTATQUERY._serialized_start=1625
  _CUSTOMWIDGETSTATQUERY._serialized_end=1721
  _CUSTOMWIDGET._serialized_start=1724
  _CUSTOMWIDGET._serialized_end=2603
# @@protoc_insertion_point(module_scope)
