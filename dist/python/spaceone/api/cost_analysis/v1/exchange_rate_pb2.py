# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/cost_analysis/v1/exchange_rate.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n1spaceone/api/cost_analysis/v1/exchange_rate.proto\x12\x1dspaceone.api.cost_analysis.v1\x1a\x1cgoogle/api/annotations.proto\"K\n\x16SetExchangeRateRequest\x12\x10\n\x08\x63urrency\x18\x01 \x01(\t\x12\x0c\n\x04rate\x18\x02 \x01(\x02\x12\x11\n\tdomain_id\x18\x0c \x01(\t\":\n\x13\x45xchangeRateRequest\x12\x10\n\x08\x63urrency\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"&\n\x11\x45xchangeRateQuery\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"\xcd\x01\n\x10\x45xchangeRateInfo\x12\x10\n\x08\x63urrency\x18\x01 \x01(\t\x12\x0c\n\x04rate\x18\x02 \x01(\x02\x12\x44\n\x05state\x18\x03 \x01(\x0e\x32\x35.spaceone.api.cost_analysis.v1.ExchangeRateInfo.State\x12\x12\n\nis_default\x18\x04 \x01(\x08\x12\x11\n\tdomain_id\x18\x0b \x01(\t\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"j\n\x11\x45xchangeRatesInfo\x12@\n\x07results\x18\x01 \x03(\x0b\x32/.spaceone.api.cost_analysis.v1.ExchangeRateInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\x32\xd2\x07\n\x0c\x45xchangeRate\x12\x9d\x01\n\x03set\x12\x35.spaceone.api.cost_analysis.v1.SetExchangeRateRequest\x1a/.spaceone.api.cost_analysis.v1.ExchangeRateInfo\".\x82\xd3\xe4\x93\x02(\"#/cost-analysis/v1/exchange-rate/set:\x01*\x12\x9e\x01\n\x05reset\x12\x32.spaceone.api.cost_analysis.v1.ExchangeRateRequest\x1a/.spaceone.api.cost_analysis.v1.ExchangeRateInfo\"0\x82\xd3\xe4\x93\x02*\"%/cost-analysis/v1/exchange-rate/reset:\x01*\x12\xa0\x01\n\x06\x65nable\x12\x32.spaceone.api.cost_analysis.v1.ExchangeRateRequest\x1a/.spaceone.api.cost_analysis.v1.ExchangeRateInfo\"1\x82\xd3\xe4\x93\x02+\"&/cost-analysis/v1/exchange-rate/enable:\x01*\x12\xa2\x01\n\x07\x64isable\x12\x32.spaceone.api.cost_analysis.v1.ExchangeRateRequest\x1a/.spaceone.api.cost_analysis.v1.ExchangeRateInfo\"2\x82\xd3\xe4\x93\x02,\"\'/cost-analysis/v1/exchange-rate/disable:\x01*\x12\x9a\x01\n\x03get\x12\x32.spaceone.api.cost_analysis.v1.ExchangeRateRequest\x1a/.spaceone.api.cost_analysis.v1.ExchangeRateInfo\".\x82\xd3\xe4\x93\x02(\"#/cost-analysis/v1/exchange-rate/get:\x01*\x12\x9b\x01\n\x04list\x12\x30.spaceone.api.cost_analysis.v1.ExchangeRateQuery\x1a\x30.spaceone.api.cost_analysis.v1.ExchangeRatesInfo\"/\x82\xd3\xe4\x93\x02)\"$/cost-analysis/v1/exchange-rate/list:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.cost_analysis.v1.exchange_rate_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1'
  _EXCHANGERATE.methods_by_name['set']._options = None
  _EXCHANGERATE.methods_by_name['set']._serialized_options = b'\202\323\344\223\002(\"#/cost-analysis/v1/exchange-rate/set:\001*'
  _EXCHANGERATE.methods_by_name['reset']._options = None
  _EXCHANGERATE.methods_by_name['reset']._serialized_options = b'\202\323\344\223\002*\"%/cost-analysis/v1/exchange-rate/reset:\001*'
  _EXCHANGERATE.methods_by_name['enable']._options = None
  _EXCHANGERATE.methods_by_name['enable']._serialized_options = b'\202\323\344\223\002+\"&/cost-analysis/v1/exchange-rate/enable:\001*'
  _EXCHANGERATE.methods_by_name['disable']._options = None
  _EXCHANGERATE.methods_by_name['disable']._serialized_options = b'\202\323\344\223\002,\"\'/cost-analysis/v1/exchange-rate/disable:\001*'
  _EXCHANGERATE.methods_by_name['get']._options = None
  _EXCHANGERATE.methods_by_name['get']._serialized_options = b'\202\323\344\223\002(\"#/cost-analysis/v1/exchange-rate/get:\001*'
  _EXCHANGERATE.methods_by_name['list']._options = None
  _EXCHANGERATE.methods_by_name['list']._serialized_options = b'\202\323\344\223\002)\"$/cost-analysis/v1/exchange-rate/list:\001*'
  _SETEXCHANGERATEREQUEST._serialized_start=114
  _SETEXCHANGERATEREQUEST._serialized_end=189
  _EXCHANGERATEREQUEST._serialized_start=191
  _EXCHANGERATEREQUEST._serialized_end=249
  _EXCHANGERATEQUERY._serialized_start=251
  _EXCHANGERATEQUERY._serialized_end=289
  _EXCHANGERATEINFO._serialized_start=292
  _EXCHANGERATEINFO._serialized_end=497
  _EXCHANGERATEINFO_STATE._serialized_start=453
  _EXCHANGERATEINFO_STATE._serialized_end=497
  _EXCHANGERATESINFO._serialized_start=499
  _EXCHANGERATESINFO._serialized_end=605
  _EXCHANGERATE._serialized_start=608
  _EXCHANGERATE._serialized_end=1586
# @@protoc_insertion_point(module_scope)
