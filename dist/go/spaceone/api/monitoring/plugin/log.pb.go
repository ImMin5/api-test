// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: spaceone/api/monitoring/plugin/log.proto

package plugin

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type Sort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Desc bool   `protobuf:"varint,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *Sort) Reset() {
	*x = Sort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_monitoring_plugin_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_plugin_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_log_proto_rawDescGZIP(), []int{0}
}

func (x *Sort) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Sort) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options    *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	SecretData *_struct.Struct `protobuf:"bytes,2,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	// +optional
	Schema string          `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	Query  *_struct.Struct `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	Keyword string `protobuf:"bytes,5,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Start   string `protobuf:"bytes,10,opt,name=start,proto3" json:"start,omitempty"`
	End     string `protobuf:"bytes,11,opt,name=end,proto3" json:"end,omitempty"`
	// +optional
	Sort *Sort `protobuf:"bytes,12,opt,name=sort,proto3" json:"sort,omitempty"`
	// +optional
	Limit int32 `protobuf:"varint,13,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_monitoring_plugin_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_plugin_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *LogRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *LogRequest) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *LogRequest) GetQuery() *_struct.Struct {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *LogRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *LogRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *LogRequest) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *LogRequest) GetSort() *Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *LogRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type LogsDataInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*_struct.Struct `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *LogsDataInfo) Reset() {
	*x = LogsDataInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_monitoring_plugin_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsDataInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsDataInfo) ProtoMessage() {}

func (x *LogsDataInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_plugin_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsDataInfo.ProtoReflect.Descriptor instead.
func (*LogsDataInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_log_proto_rawDescGZIP(), []int{2}
}

func (x *LogsDataInfo) GetResults() []*_struct.Struct {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_spaceone_api_monitoring_plugin_log_proto protoreflect.FileDescriptor

var file_spaceone_api_monitoring_plugin_log_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0xd2, 0x02, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x2d, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x38, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x41, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x73, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x6b,
	0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x64, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x42, 0x45, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66,
	0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_monitoring_plugin_log_proto_rawDescOnce sync.Once
	file_spaceone_api_monitoring_plugin_log_proto_rawDescData = file_spaceone_api_monitoring_plugin_log_proto_rawDesc
)

func file_spaceone_api_monitoring_plugin_log_proto_rawDescGZIP() []byte {
	file_spaceone_api_monitoring_plugin_log_proto_rawDescOnce.Do(func() {
		file_spaceone_api_monitoring_plugin_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_monitoring_plugin_log_proto_rawDescData)
	})
	return file_spaceone_api_monitoring_plugin_log_proto_rawDescData
}

var file_spaceone_api_monitoring_plugin_log_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spaceone_api_monitoring_plugin_log_proto_goTypes = []interface{}{
	(*Sort)(nil),           // 0: spaceone.api.monitoring.plugin.Sort
	(*LogRequest)(nil),     // 1: spaceone.api.monitoring.plugin.LogRequest
	(*LogsDataInfo)(nil),   // 2: spaceone.api.monitoring.plugin.LogsDataInfo
	(*_struct.Struct)(nil), // 3: google.protobuf.Struct
}
var file_spaceone_api_monitoring_plugin_log_proto_depIdxs = []int32{
	3, // 0: spaceone.api.monitoring.plugin.LogRequest.options:type_name -> google.protobuf.Struct
	3, // 1: spaceone.api.monitoring.plugin.LogRequest.secret_data:type_name -> google.protobuf.Struct
	3, // 2: spaceone.api.monitoring.plugin.LogRequest.query:type_name -> google.protobuf.Struct
	0, // 3: spaceone.api.monitoring.plugin.LogRequest.sort:type_name -> spaceone.api.monitoring.plugin.Sort
	3, // 4: spaceone.api.monitoring.plugin.LogsDataInfo.results:type_name -> google.protobuf.Struct
	1, // 5: spaceone.api.monitoring.plugin.Log.list:input_type -> spaceone.api.monitoring.plugin.LogRequest
	2, // 6: spaceone.api.monitoring.plugin.Log.list:output_type -> spaceone.api.monitoring.plugin.LogsDataInfo
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_spaceone_api_monitoring_plugin_log_proto_init() }
func file_spaceone_api_monitoring_plugin_log_proto_init() {
	if File_spaceone_api_monitoring_plugin_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_monitoring_plugin_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sort); i {
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
		file_spaceone_api_monitoring_plugin_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
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
		file_spaceone_api_monitoring_plugin_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsDataInfo); i {
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
			RawDescriptor: file_spaceone_api_monitoring_plugin_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_monitoring_plugin_log_proto_goTypes,
		DependencyIndexes: file_spaceone_api_monitoring_plugin_log_proto_depIdxs,
		MessageInfos:      file_spaceone_api_monitoring_plugin_log_proto_msgTypes,
	}.Build()
	File_spaceone_api_monitoring_plugin_log_proto = out.File
	file_spaceone_api_monitoring_plugin_log_proto_rawDesc = nil
	file_spaceone_api_monitoring_plugin_log_proto_goTypes = nil
	file_spaceone_api_monitoring_plugin_log_proto_depIdxs = nil
}
