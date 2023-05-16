// An ExchangeRate is a resource defining the exchange rate of currencies. This resource can set a custom exchange rate for a specific domain, separately from the exchange rate of the default domain set in `config`.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: spaceone/api/cost_analysis/v1/exchange_rate.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExchangeRateInfo_State int32

const (
	ExchangeRateInfo_NONE     ExchangeRateInfo_State = 0
	ExchangeRateInfo_ENABLED  ExchangeRateInfo_State = 1
	ExchangeRateInfo_DISABLED ExchangeRateInfo_State = 2
)

// Enum value maps for ExchangeRateInfo_State.
var (
	ExchangeRateInfo_State_name = map[int32]string{
		0: "NONE",
		1: "ENABLED",
		2: "DISABLED",
	}
	ExchangeRateInfo_State_value = map[string]int32{
		"NONE":     0,
		"ENABLED":  1,
		"DISABLED": 2,
	}
)

func (x ExchangeRateInfo_State) Enum() *ExchangeRateInfo_State {
	p := new(ExchangeRateInfo_State)
	*p = x
	return p
}

func (x ExchangeRateInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExchangeRateInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_cost_analysis_v1_exchange_rate_proto_enumTypes[0].Descriptor()
}

func (ExchangeRateInfo_State) Type() protoreflect.EnumType {
	return &file_spaceone_api_cost_analysis_v1_exchange_rate_proto_enumTypes[0]
}

func (x ExchangeRateInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExchangeRateInfo_State.Descriptor instead.
func (ExchangeRateInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescGZIP(), []int{3, 0}
}

//	{
//	   "currency": "KRW",
//	   "rate": 1300
//	}
type SetExchangeRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency string  `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Rate     float32 `protobuf:"fixed32,2,opt,name=rate,proto3" json:"rate,omitempty"`
	DomainId string  `protobuf:"bytes,12,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *SetExchangeRateRequest) Reset() {
	*x = SetExchangeRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetExchangeRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetExchangeRateRequest) ProtoMessage() {}

func (x *SetExchangeRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetExchangeRateRequest.ProtoReflect.Descriptor instead.
func (*SetExchangeRateRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescGZIP(), []int{0}
}

func (x *SetExchangeRateRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *SetExchangeRateRequest) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *SetExchangeRateRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

//	{
//	   "currency": "KRW"
//	}
type ExchangeRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	DomainId string `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *ExchangeRateRequest) Reset() {
	*x = ExchangeRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRateRequest) ProtoMessage() {}

func (x *ExchangeRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRateRequest.ProtoReflect.Descriptor instead.
func (*ExchangeRateRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeRateRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ExchangeRateRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

// {
//
// }
type ExchangeRateQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainId string `protobuf:"bytes,11,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *ExchangeRateQuery) Reset() {
	*x = ExchangeRateQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRateQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRateQuery) ProtoMessage() {}

func (x *ExchangeRateQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRateQuery.ProtoReflect.Descriptor instead.
func (*ExchangeRateQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescGZIP(), []int{2}
}

func (x *ExchangeRateQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

//	{
//	   "currency": "KRW",
//	   "rate": 1300.0,
//	   "state": "ENABLED",
//	   "is_default": true
//	}
type ExchangeRateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency  string                 `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Rate      float32                `protobuf:"fixed32,2,opt,name=rate,proto3" json:"rate,omitempty"`
	State     ExchangeRateInfo_State `protobuf:"varint,3,opt,name=state,proto3,enum=spaceone.api.cost_analysis.v1.ExchangeRateInfo_State" json:"state,omitempty"`
	IsDefault bool                   `protobuf:"varint,4,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	DomainId  string                 `protobuf:"bytes,11,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *ExchangeRateInfo) Reset() {
	*x = ExchangeRateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRateInfo) ProtoMessage() {}

func (x *ExchangeRateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRateInfo.ProtoReflect.Descriptor instead.
func (*ExchangeRateInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescGZIP(), []int{3}
}

func (x *ExchangeRateInfo) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ExchangeRateInfo) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *ExchangeRateInfo) GetState() ExchangeRateInfo_State {
	if x != nil {
		return x.State
	}
	return ExchangeRateInfo_NONE
}

func (x *ExchangeRateInfo) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

func (x *ExchangeRateInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

//	{
//	   "results": [
//	       {
//	           "currency": "JPY",
//	           "rate": 129.8,
//	           "state": "ENABLED",
//	           "is_default": true,
//	           "domain_id": "domain-58010aa2e451"
//	       },
//	       {
//	           "currency": "KRW",
//	           "rate": 1242.7,
//	           "state": "ENABLED",
//	           "is_default": true,
//	           "domain_id": "domain-58010aa2e451"
//	       }
//	   ],
//	   "total_count": 2
//	}
type ExchangeRatesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*ExchangeRateInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32               `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ExchangeRatesInfo) Reset() {
	*x = ExchangeRatesInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRatesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRatesInfo) ProtoMessage() {}

func (x *ExchangeRatesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRatesInfo.ProtoReflect.Descriptor instead.
func (*ExchangeRatesInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescGZIP(), []int{4}
}

func (x *ExchangeRatesInfo) GetResults() []*ExchangeRateInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ExchangeRatesInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_spaceone_api_cost_analysis_v1_exchange_rate_proto protoreflect.FileDescriptor

var file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDesc = []byte{
	0x0a, 0x31, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x65, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x13, 0x45, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x11, 0x45, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xf9, 0x01, 0x0a, 0x10, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x4b,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x22, 0x7f, 0x0a, 0x11, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xd2, 0x07, 0x0a, 0x0c, 0x45, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x12,
	0x35, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a,
	0x01, 0x2a, 0x22, 0x23, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x72,
	0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x74, 0x12, 0x9e, 0x01, 0x0a, 0x05, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x12, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01,
	0x2a, 0x22, 0x25, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x72, 0x61,
	0x74, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0xa0, 0x01, 0x0a, 0x06, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b,
	0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2d,
	0x72, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0xa2, 0x01, 0x0a, 0x07,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x32, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2d, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x9a, 0x01, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2d, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x9b, 0x01,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2d, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x44, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66,
	0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescData = file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDesc
)

func file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescData)
	})
	return file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDescData
}

var file_spaceone_api_cost_analysis_v1_exchange_rate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_cost_analysis_v1_exchange_rate_proto_goTypes = []interface{}{
	(ExchangeRateInfo_State)(0),    // 0: spaceone.api.cost_analysis.v1.ExchangeRateInfo.State
	(*SetExchangeRateRequest)(nil), // 1: spaceone.api.cost_analysis.v1.SetExchangeRateRequest
	(*ExchangeRateRequest)(nil),    // 2: spaceone.api.cost_analysis.v1.ExchangeRateRequest
	(*ExchangeRateQuery)(nil),      // 3: spaceone.api.cost_analysis.v1.ExchangeRateQuery
	(*ExchangeRateInfo)(nil),       // 4: spaceone.api.cost_analysis.v1.ExchangeRateInfo
	(*ExchangeRatesInfo)(nil),      // 5: spaceone.api.cost_analysis.v1.ExchangeRatesInfo
}
var file_spaceone_api_cost_analysis_v1_exchange_rate_proto_depIdxs = []int32{
	0, // 0: spaceone.api.cost_analysis.v1.ExchangeRateInfo.state:type_name -> spaceone.api.cost_analysis.v1.ExchangeRateInfo.State
	4, // 1: spaceone.api.cost_analysis.v1.ExchangeRatesInfo.results:type_name -> spaceone.api.cost_analysis.v1.ExchangeRateInfo
	1, // 2: spaceone.api.cost_analysis.v1.ExchangeRate.set:input_type -> spaceone.api.cost_analysis.v1.SetExchangeRateRequest
	2, // 3: spaceone.api.cost_analysis.v1.ExchangeRate.reset:input_type -> spaceone.api.cost_analysis.v1.ExchangeRateRequest
	2, // 4: spaceone.api.cost_analysis.v1.ExchangeRate.enable:input_type -> spaceone.api.cost_analysis.v1.ExchangeRateRequest
	2, // 5: spaceone.api.cost_analysis.v1.ExchangeRate.disable:input_type -> spaceone.api.cost_analysis.v1.ExchangeRateRequest
	2, // 6: spaceone.api.cost_analysis.v1.ExchangeRate.get:input_type -> spaceone.api.cost_analysis.v1.ExchangeRateRequest
	3, // 7: spaceone.api.cost_analysis.v1.ExchangeRate.list:input_type -> spaceone.api.cost_analysis.v1.ExchangeRateQuery
	4, // 8: spaceone.api.cost_analysis.v1.ExchangeRate.set:output_type -> spaceone.api.cost_analysis.v1.ExchangeRateInfo
	4, // 9: spaceone.api.cost_analysis.v1.ExchangeRate.reset:output_type -> spaceone.api.cost_analysis.v1.ExchangeRateInfo
	4, // 10: spaceone.api.cost_analysis.v1.ExchangeRate.enable:output_type -> spaceone.api.cost_analysis.v1.ExchangeRateInfo
	4, // 11: spaceone.api.cost_analysis.v1.ExchangeRate.disable:output_type -> spaceone.api.cost_analysis.v1.ExchangeRateInfo
	4, // 12: spaceone.api.cost_analysis.v1.ExchangeRate.get:output_type -> spaceone.api.cost_analysis.v1.ExchangeRateInfo
	5, // 13: spaceone.api.cost_analysis.v1.ExchangeRate.list:output_type -> spaceone.api.cost_analysis.v1.ExchangeRatesInfo
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_v1_exchange_rate_proto_init() }
func file_spaceone_api_cost_analysis_v1_exchange_rate_proto_init() {
	if File_spaceone_api_cost_analysis_v1_exchange_rate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetExchangeRateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRateQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRateInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRatesInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_v1_exchange_rate_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_v1_exchange_rate_proto_depIdxs,
		EnumInfos:         file_spaceone_api_cost_analysis_v1_exchange_rate_proto_enumTypes,
		MessageInfos:      file_spaceone_api_cost_analysis_v1_exchange_rate_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_v1_exchange_rate_proto = out.File
	file_spaceone_api_cost_analysis_v1_exchange_rate_proto_rawDesc = nil
	file_spaceone_api_cost_analysis_v1_exchange_rate_proto_goTypes = nil
	file_spaceone_api_cost_analysis_v1_exchange_rate_proto_depIdxs = nil
}
