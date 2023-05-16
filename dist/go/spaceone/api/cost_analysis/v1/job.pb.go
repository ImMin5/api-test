// A Job is an act of collecting external cost data through plugins. The data to collect is defined in a plugin.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: spaceone/api/cost_analysis/v1/job.proto

package v1

import (
	v1 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v1"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type JobQuery_Status int32

const (
	JobQuery_SCOPE_NONE  JobQuery_Status = 0
	JobQuery_IN_PROGRESS JobQuery_Status = 1
	JobQuery_SUCCESS     JobQuery_Status = 2
	JobQuery_FAILURE     JobQuery_Status = 3
	JobQuery_TIMEOUT     JobQuery_Status = 4
	JobQuery_CANCELED    JobQuery_Status = 5
)

// Enum value maps for JobQuery_Status.
var (
	JobQuery_Status_name = map[int32]string{
		0: "SCOPE_NONE",
		1: "IN_PROGRESS",
		2: "SUCCESS",
		3: "FAILURE",
		4: "TIMEOUT",
		5: "CANCELED",
	}
	JobQuery_Status_value = map[string]int32{
		"SCOPE_NONE":  0,
		"IN_PROGRESS": 1,
		"SUCCESS":     2,
		"FAILURE":     3,
		"TIMEOUT":     4,
		"CANCELED":    5,
	}
)

func (x JobQuery_Status) Enum() *JobQuery_Status {
	p := new(JobQuery_Status)
	*p = x
	return p
}

func (x JobQuery_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobQuery_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_cost_analysis_v1_job_proto_enumTypes[0].Descriptor()
}

func (JobQuery_Status) Type() protoreflect.EnumType {
	return &file_spaceone_api_cost_analysis_v1_job_proto_enumTypes[0]
}

func (x JobQuery_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobQuery_Status.Descriptor instead.
func (JobQuery_Status) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_proto_rawDescGZIP(), []int{2, 0}
}

type JobInfo_Status int32

const (
	JobInfo_SCOPE_NONE  JobInfo_Status = 0
	JobInfo_IN_PROGRESS JobInfo_Status = 1
	JobInfo_SUCCESS     JobInfo_Status = 2
	JobInfo_FAILURE     JobInfo_Status = 3
	JobInfo_TIMEOUT     JobInfo_Status = 4
	JobInfo_CANCELED    JobInfo_Status = 5
)

// Enum value maps for JobInfo_Status.
var (
	JobInfo_Status_name = map[int32]string{
		0: "SCOPE_NONE",
		1: "IN_PROGRESS",
		2: "SUCCESS",
		3: "FAILURE",
		4: "TIMEOUT",
		5: "CANCELED",
	}
	JobInfo_Status_value = map[string]int32{
		"SCOPE_NONE":  0,
		"IN_PROGRESS": 1,
		"SUCCESS":     2,
		"FAILURE":     3,
		"TIMEOUT":     4,
		"CANCELED":    5,
	}
)

func (x JobInfo_Status) Enum() *JobInfo_Status {
	p := new(JobInfo_Status)
	*p = x
	return p
}

func (x JobInfo_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobInfo_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_cost_analysis_v1_job_proto_enumTypes[1].Descriptor()
}

func (JobInfo_Status) Type() protoreflect.EnumType {
	return &file_spaceone_api_cost_analysis_v1_job_proto_enumTypes[1]
}

func (x JobInfo_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobInfo_Status.Descriptor instead.
func (JobInfo_Status) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_proto_rawDescGZIP(), []int{4, 0}
}

//	{
//	   "job_id": "job-07994c7c9021"
//	}
type JobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId    string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	DomainId string `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *JobRequest) Reset() {
	*x = JobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobRequest) ProtoMessage() {}

func (x *JobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobRequest.ProtoReflect.Descriptor instead.
func (*JobRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_proto_rawDescGZIP(), []int{0}
}

func (x *JobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

//	{
//	   "job_id": "job-85cf2c385252"
//	}
type GetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId    string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	DomainId string `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// +optional
	Only []string `protobuf:"bytes,3,rep,name=only,proto3" json:"only,omitempty"`
}

func (x *GetJobRequest) Reset() {
	*x = GetJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobRequest) ProtoMessage() {}

func (x *GetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobRequest.ProtoReflect.Descriptor instead.
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_proto_rawDescGZIP(), []int{1}
}

func (x *GetJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *GetJobRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *GetJobRequest) GetOnly() []string {
	if x != nil {
		return x.Only
	}
	return nil
}

//	{
//	   "query": {}
//	}
type JobQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v1.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	JobId string `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// +optional
	Status JobQuery_Status `protobuf:"varint,3,opt,name=status,proto3,enum=spaceone.api.cost_analysis.v1.JobQuery_Status" json:"status,omitempty"`
	// +optional
	DataSourceId string `protobuf:"bytes,11,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	DomainId     string `protobuf:"bytes,12,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *JobQuery) Reset() {
	*x = JobQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobQuery) ProtoMessage() {}

func (x *JobQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobQuery.ProtoReflect.Descriptor instead.
func (*JobQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_proto_rawDescGZIP(), []int{2}
}

func (x *JobQuery) GetQuery() *v1.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *JobQuery) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobQuery) GetStatus() JobQuery_Status {
	if x != nil {
		return x.Status
	}
	return JobQuery_SCOPE_NONE
}

func (x *JobQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *JobQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type ChangedInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start  string          `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End    string          `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	Filter *_struct.Struct `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ChangedInfo) Reset() {
	*x = ChangedInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangedInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangedInfo) ProtoMessage() {}

func (x *ChangedInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangedInfo.ProtoReflect.Descriptor instead.
func (*ChangedInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_proto_rawDescGZIP(), []int{3}
}

func (x *ChangedInfo) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *ChangedInfo) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *ChangedInfo) GetFilter() *_struct.Struct {
	if x != nil {
		return x.Filter
	}
	return nil
}

//	{
//	   "job_id": "job-07994c7c9021",
//	   "status": "CANCELED",
//	   "options": {
//	       "no_preload_cache": false,
//	       "start": "2021-01-01T00:00:00Z"
//	   },
//	   "total_tasks": 2,
//	   "remained_tasks": 2,
//	   "data_source_id": "ds-fcba92ca73b1",
//	   "domain_id": "domain-58010aa2e451",
//	   "created_at": "2022-04-02T09:17:44.031Z",
//	   "updated_at": "2022-04-02T09:19:47.715Z",
//	   "finished_at": "2022-04-02T09:19:47.715Z",
//	   "changed": [
//	       {
//	           "start": "2021-01-01T00:00:00.000Z"
//	       }
//	   ]
//	}
type JobInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId         string          `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Status        JobInfo_Status  `protobuf:"varint,2,opt,name=status,proto3,enum=spaceone.api.cost_analysis.v1.JobInfo_Status" json:"status,omitempty"`
	Options       *_struct.Struct `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	ErrorCode     string          `protobuf:"bytes,4,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage  string          `protobuf:"bytes,5,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	TotalTasks    int32           `protobuf:"varint,6,opt,name=total_tasks,json=totalTasks,proto3" json:"total_tasks,omitempty"`
	RemainedTasks int32           `protobuf:"varint,7,opt,name=remained_tasks,json=remainedTasks,proto3" json:"remained_tasks,omitempty"`
	DataSourceId  string          `protobuf:"bytes,11,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	DomainId      string          `protobuf:"bytes,12,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CreatedAt     string          `protobuf:"bytes,21,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string          `protobuf:"bytes,22,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	FinishedAt    string          `protobuf:"bytes,23,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	Changed       []*ChangedInfo  `protobuf:"bytes,24,rep,name=changed,proto3" json:"changed,omitempty"`
}

func (x *JobInfo) Reset() {
	*x = JobInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobInfo) ProtoMessage() {}

func (x *JobInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobInfo.ProtoReflect.Descriptor instead.
func (*JobInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_proto_rawDescGZIP(), []int{4}
}

func (x *JobInfo) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobInfo) GetStatus() JobInfo_Status {
	if x != nil {
		return x.Status
	}
	return JobInfo_SCOPE_NONE
}

func (x *JobInfo) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *JobInfo) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *JobInfo) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *JobInfo) GetTotalTasks() int32 {
	if x != nil {
		return x.TotalTasks
	}
	return 0
}

func (x *JobInfo) GetRemainedTasks() int32 {
	if x != nil {
		return x.RemainedTasks
	}
	return 0
}

func (x *JobInfo) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *JobInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *JobInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *JobInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *JobInfo) GetFinishedAt() string {
	if x != nil {
		return x.FinishedAt
	}
	return ""
}

func (x *JobInfo) GetChanged() []*ChangedInfo {
	if x != nil {
		return x.Changed
	}
	return nil
}

//	{
//	       "results": [
//	           {
//	               "job_id": "job-85cf2c385252",
//	               "status": "SUCCESS",
//	               "options": {
//	                   "start": null,
//	                   "no_preload_cache": false
//	               },
//	               "total_tasks": 1,
//	               "data_source_id": "ds-c96609f5afeb",
//	               "domain_id": "domain-58010aa2e451",
//	               "created_at": "2022-07-17T16:00:08.254Z",
//	               "updated_at": "2022-07-17T16:01:30.637Z",
//	               "finished_at": "2022-07-17T16:01:30.637Z",
//	               "changed": [
//	                   {
//	                       "start": "2022-07-01T00:00:00.000Z"
//	                   }
//	               ]
//	           },
//	           {
//	               "job_id": "job-6b6765f757a9",
//	               "status": "SUCCESS",
//	               "options": {
//	                   "start": null,
//	                   "no_preload_cache": false
//	               },
//	               "total_tasks": 2,
//	               "data_source_id": "ds-fcba92ca73b1",
//	               "domain_id": "domain-58010aa2e451",
//	               "created_at": "2022-07-17T16:00:05.077Z",
//	               "updated_at": "2022-07-17T16:01:28.206Z",
//	               "finished_at": "2022-07-17T16:01:28.206Z",
//	               "changed": [
//	                   {
//	                       "start": "2022-07-01T00:00:00.000Z"
//	                   }
//	               ]
//	           }
//	       ],
//	       "total_count": 372
//	}
type JobsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*JobInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32      `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *JobsInfo) Reset() {
	*x = JobsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobsInfo) ProtoMessage() {}

func (x *JobsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobsInfo.ProtoReflect.Descriptor instead.
func (*JobsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_proto_rawDescGZIP(), []int{5}
}

func (x *JobsInfo) GetResults() []*JobInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *JobsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type JobStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    *v1.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DomainId string              `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *JobStatQuery) Reset() {
	*x = JobStatQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatQuery) ProtoMessage() {}

func (x *JobStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatQuery.ProtoReflect.Descriptor instead.
func (*JobStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_proto_rawDescGZIP(), []int{6}
}

func (x *JobStatQuery) GetQuery() *v1.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *JobStatQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_spaceone_api_cost_analysis_v1_job_proto protoreflect.FileDescriptor

var file_spaceone_api_cost_analysis_v1_job_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x6e, 0x6c,
	0x79, 0x22, 0xbf, 0x02, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x46,
	0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45,
	0x4f, 0x55, 0x54, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45,
	0x44, 0x10, 0x05, 0x22, 0x66, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xee, 0x04, 0x0a, 0x07,
	0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x45,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a,
	0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x22, 0x5e, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f,
	0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10,
	0x03, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x04, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x22, 0x6d, 0x0a, 0x08,
	0x4a, 0x6f, 0x62, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x68, 0x0a, 0x0c, 0x4a,
	0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x32, 0x86, 0x04, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x84, 0x01,
	0x0a, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x27, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x12, 0x81, 0x01, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x2c, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f,
	0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x7f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4a, 0x6f, 0x62, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f,
	0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x73, 0x0a, 0x04, 0x73, 0x74, 0x61,
	0x74, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a,
	0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x44,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_cost_analysis_v1_job_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_v1_job_proto_rawDescData = file_spaceone_api_cost_analysis_v1_job_proto_rawDesc
)

func file_spaceone_api_cost_analysis_v1_job_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_v1_job_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_v1_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_cost_analysis_v1_job_proto_rawDescData)
	})
	return file_spaceone_api_cost_analysis_v1_job_proto_rawDescData
}

var file_spaceone_api_cost_analysis_v1_job_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spaceone_api_cost_analysis_v1_job_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spaceone_api_cost_analysis_v1_job_proto_goTypes = []interface{}{
	(JobQuery_Status)(0),       // 0: spaceone.api.cost_analysis.v1.JobQuery.Status
	(JobInfo_Status)(0),        // 1: spaceone.api.cost_analysis.v1.JobInfo.Status
	(*JobRequest)(nil),         // 2: spaceone.api.cost_analysis.v1.JobRequest
	(*GetJobRequest)(nil),      // 3: spaceone.api.cost_analysis.v1.GetJobRequest
	(*JobQuery)(nil),           // 4: spaceone.api.cost_analysis.v1.JobQuery
	(*ChangedInfo)(nil),        // 5: spaceone.api.cost_analysis.v1.ChangedInfo
	(*JobInfo)(nil),            // 6: spaceone.api.cost_analysis.v1.JobInfo
	(*JobsInfo)(nil),           // 7: spaceone.api.cost_analysis.v1.JobsInfo
	(*JobStatQuery)(nil),       // 8: spaceone.api.cost_analysis.v1.JobStatQuery
	(*v1.Query)(nil),           // 9: spaceone.api.core.v1.Query
	(*_struct.Struct)(nil),     // 10: google.protobuf.Struct
	(*v1.StatisticsQuery)(nil), // 11: spaceone.api.core.v1.StatisticsQuery
}
var file_spaceone_api_cost_analysis_v1_job_proto_depIdxs = []int32{
	9,  // 0: spaceone.api.cost_analysis.v1.JobQuery.query:type_name -> spaceone.api.core.v1.Query
	0,  // 1: spaceone.api.cost_analysis.v1.JobQuery.status:type_name -> spaceone.api.cost_analysis.v1.JobQuery.Status
	10, // 2: spaceone.api.cost_analysis.v1.ChangedInfo.filter:type_name -> google.protobuf.Struct
	1,  // 3: spaceone.api.cost_analysis.v1.JobInfo.status:type_name -> spaceone.api.cost_analysis.v1.JobInfo.Status
	10, // 4: spaceone.api.cost_analysis.v1.JobInfo.options:type_name -> google.protobuf.Struct
	5,  // 5: spaceone.api.cost_analysis.v1.JobInfo.changed:type_name -> spaceone.api.cost_analysis.v1.ChangedInfo
	6,  // 6: spaceone.api.cost_analysis.v1.JobsInfo.results:type_name -> spaceone.api.cost_analysis.v1.JobInfo
	11, // 7: spaceone.api.cost_analysis.v1.JobStatQuery.query:type_name -> spaceone.api.core.v1.StatisticsQuery
	2,  // 8: spaceone.api.cost_analysis.v1.Job.cancel:input_type -> spaceone.api.cost_analysis.v1.JobRequest
	3,  // 9: spaceone.api.cost_analysis.v1.Job.get:input_type -> spaceone.api.cost_analysis.v1.GetJobRequest
	4,  // 10: spaceone.api.cost_analysis.v1.Job.list:input_type -> spaceone.api.cost_analysis.v1.JobQuery
	8,  // 11: spaceone.api.cost_analysis.v1.Job.stat:input_type -> spaceone.api.cost_analysis.v1.JobStatQuery
	6,  // 12: spaceone.api.cost_analysis.v1.Job.cancel:output_type -> spaceone.api.cost_analysis.v1.JobInfo
	6,  // 13: spaceone.api.cost_analysis.v1.Job.get:output_type -> spaceone.api.cost_analysis.v1.JobInfo
	7,  // 14: spaceone.api.cost_analysis.v1.Job.list:output_type -> spaceone.api.cost_analysis.v1.JobsInfo
	10, // 15: spaceone.api.cost_analysis.v1.Job.stat:output_type -> google.protobuf.Struct
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_v1_job_proto_init() }
func file_spaceone_api_cost_analysis_v1_job_proto_init() {
	if File_spaceone_api_cost_analysis_v1_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobRequest); i {
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
		file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobRequest); i {
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
		file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobQuery); i {
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
		file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangedInfo); i {
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
		file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobInfo); i {
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
		file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobsInfo); i {
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
		file_spaceone_api_cost_analysis_v1_job_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStatQuery); i {
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
			RawDescriptor: file_spaceone_api_cost_analysis_v1_job_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_v1_job_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_v1_job_proto_depIdxs,
		EnumInfos:         file_spaceone_api_cost_analysis_v1_job_proto_enumTypes,
		MessageInfos:      file_spaceone_api_cost_analysis_v1_job_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_v1_job_proto = out.File
	file_spaceone_api_cost_analysis_v1_job_proto_rawDesc = nil
	file_spaceone_api_cost_analysis_v1_job_proto_goTypes = nil
	file_spaceone_api_cost_analysis_v1_job_proto_depIdxs = nil
}
