// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: sastoj/admin/judge/service/v1/judge.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content     string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Point       int32  `protobuf:"varint,4,opt,name=point,proto3" json:"point,omitempty"`
	ContestId   int64  `protobuf:"varint,5,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	CaseVersion int32  `protobuf:"varint,6,opt,name=case_version,json=caseVersion,proto3" json:"case_version,omitempty"`
	Index       int32  `protobuf:"varint,7,opt,name=index,proto3" json:"index,omitempty"`
	Config      string `protobuf:"bytes,8,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Problem) Reset() {
	*x = Problem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem.ProtoReflect.Descriptor instead.
func (*Problem) Descriptor() ([]byte, []int) {
	return file_sastoj_admin_judge_service_v1_judge_proto_rawDescGZIP(), []int{0}
}

func (x *Problem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Problem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Problem) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Problem) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *Problem) GetContestId() int64 {
	if x != nil {
		return x.ContestId
	}
	return 0
}

func (x *Problem) GetCaseVersion() int32 {
	if x != nil {
		return x.CaseVersion
	}
	return 0
}

func (x *Problem) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Problem) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

type Submission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code       string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Status     int32                  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Point      int32                  `protobuf:"varint,4,opt,name=point,proto3" json:"point,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Submission) Reset() {
	*x = Submission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission) ProtoMessage() {}

func (x *Submission) ProtoReflect() protoreflect.Message {
	mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submission.ProtoReflect.Descriptor instead.
func (*Submission) Descriptor() ([]byte, []int) {
	return file_sastoj_admin_judge_service_v1_judge_proto_rawDescGZIP(), []int{1}
}

func (x *Submission) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Submission) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Submission) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Submission) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *Submission) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type SubmitJudgeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmissionId int64 `protobuf:"varint,1,opt,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
	Point        int32 `protobuf:"varint,2,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *SubmitJudgeRequest) Reset() {
	*x = SubmitJudgeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitJudgeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitJudgeRequest) ProtoMessage() {}

func (x *SubmitJudgeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitJudgeRequest.ProtoReflect.Descriptor instead.
func (*SubmitJudgeRequest) Descriptor() ([]byte, []int) {
	return file_sastoj_admin_judge_service_v1_judge_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitJudgeRequest) GetSubmissionId() int64 {
	if x != nil {
		return x.SubmissionId
	}
	return 0
}

func (x *SubmitJudgeRequest) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

type SubmitJudgeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubmitJudgeReply) Reset() {
	*x = SubmitJudgeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitJudgeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitJudgeReply) ProtoMessage() {}

func (x *SubmitJudgeReply) ProtoReflect() protoreflect.Message {
	mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitJudgeReply.ProtoReflect.Descriptor instead.
func (*SubmitJudgeReply) Descriptor() ([]byte, []int) {
	return file_sastoj_admin_judge_service_v1_judge_proto_rawDescGZIP(), []int{3}
}

type GetJudgableProblemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetJudgableProblemsRequest) Reset() {
	*x = GetJudgableProblemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJudgableProblemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJudgableProblemsRequest) ProtoMessage() {}

func (x *GetJudgableProblemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJudgableProblemsRequest.ProtoReflect.Descriptor instead.
func (*GetJudgableProblemsRequest) Descriptor() ([]byte, []int) {
	return file_sastoj_admin_judge_service_v1_judge_proto_rawDescGZIP(), []int{4}
}

type GetJudgableProblemsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Problem `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GetJudgableProblemsReply) Reset() {
	*x = GetJudgableProblemsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJudgableProblemsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJudgableProblemsReply) ProtoMessage() {}

func (x *GetJudgableProblemsReply) ProtoReflect() protoreflect.Message {
	mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJudgableProblemsReply.ProtoReflect.Descriptor instead.
func (*GetJudgableProblemsReply) Descriptor() ([]byte, []int) {
	return file_sastoj_admin_judge_service_v1_judge_proto_rawDescGZIP(), []int{5}
}

func (x *GetJudgableProblemsReply) GetResults() []*Problem {
	if x != nil {
		return x.Results
	}
	return nil
}

type GetSubmissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId int64 `protobuf:"varint,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Status    int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetSubmissionsRequest) Reset() {
	*x = GetSubmissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubmissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubmissionsRequest) ProtoMessage() {}

func (x *GetSubmissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubmissionsRequest.ProtoReflect.Descriptor instead.
func (*GetSubmissionsRequest) Descriptor() ([]byte, []int) {
	return file_sastoj_admin_judge_service_v1_judge_proto_rawDescGZIP(), []int{6}
}

func (x *GetSubmissionsRequest) GetProblemId() int64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

func (x *GetSubmissionsRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetSubmissionsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submissions []*Submission `protobuf:"bytes,1,rep,name=submissions,proto3" json:"submissions,omitempty"`
}

func (x *GetSubmissionsReply) Reset() {
	*x = GetSubmissionsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubmissionsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubmissionsReply) ProtoMessage() {}

func (x *GetSubmissionsReply) ProtoReflect() protoreflect.Message {
	mi := &file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubmissionsReply.ProtoReflect.Descriptor instead.
func (*GetSubmissionsReply) Descriptor() ([]byte, []int) {
	return file_sastoj_admin_judge_service_v1_judge_proto_rawDescGZIP(), []int{7}
}

func (x *GetSubmissionsReply) GetSubmissions() []*Submission {
	if x != nil {
		return x.Submissions
	}
	return nil
}

var File_sastoj_admin_judge_service_v1_judge_proto protoreflect.FileDescriptor

var file_sastoj_admin_judge_service_v1_judge_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x73, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x9b, 0x01, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4f, 0x0a,
	0x12, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x12,
	0x0a, 0x10, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x61, 0x62, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x60, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x50,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x22, 0x4e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x66, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4f, 0x0a, 0x0b, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xec, 0x03, 0x0a, 0x05, 0x4a,
	0x75, 0x64, 0x67, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a,
	0x75, 0x64, 0x67, 0x65, 0x12, 0x35, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f,
	0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a,
	0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0xa1, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x61,
	0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x12, 0x3d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12,
	0x06, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0x12, 0x9f, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f,
	0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x8c, 0x02, 0x0a, 0x25, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2c, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xa2,
	0x02, 0x05, 0x41, 0x53, 0x41, 0x4a, 0x53, 0xaa, 0x02, 0x21, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x61,
	0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x21, 0x41, 0x70,
	0x69, 0x5c, 0x53, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x4a,
	0x75, 0x64, 0x67, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x2d, 0x41, 0x70, 0x69, 0x5c, 0x53, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x5c, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x5c, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x26, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x53, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x3a, 0x3a, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x3a, 0x3a, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sastoj_admin_judge_service_v1_judge_proto_rawDescOnce sync.Once
	file_sastoj_admin_judge_service_v1_judge_proto_rawDescData = file_sastoj_admin_judge_service_v1_judge_proto_rawDesc
)

func file_sastoj_admin_judge_service_v1_judge_proto_rawDescGZIP() []byte {
	file_sastoj_admin_judge_service_v1_judge_proto_rawDescOnce.Do(func() {
		file_sastoj_admin_judge_service_v1_judge_proto_rawDescData = protoimpl.X.CompressGZIP(file_sastoj_admin_judge_service_v1_judge_proto_rawDescData)
	})
	return file_sastoj_admin_judge_service_v1_judge_proto_rawDescData
}

var file_sastoj_admin_judge_service_v1_judge_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sastoj_admin_judge_service_v1_judge_proto_goTypes = []interface{}{
	(*Problem)(nil),                    // 0: api.sastoj.admin.judge.service.v1.Problem
	(*Submission)(nil),                 // 1: api.sastoj.admin.judge.service.v1.Submission
	(*SubmitJudgeRequest)(nil),         // 2: api.sastoj.admin.judge.service.v1.SubmitJudgeRequest
	(*SubmitJudgeReply)(nil),           // 3: api.sastoj.admin.judge.service.v1.SubmitJudgeReply
	(*GetJudgableProblemsRequest)(nil), // 4: api.sastoj.admin.judge.service.v1.GetJudgableProblemsRequest
	(*GetJudgableProblemsReply)(nil),   // 5: api.sastoj.admin.judge.service.v1.GetJudgableProblemsReply
	(*GetSubmissionsRequest)(nil),      // 6: api.sastoj.admin.judge.service.v1.GetSubmissionsRequest
	(*GetSubmissionsReply)(nil),        // 7: api.sastoj.admin.judge.service.v1.GetSubmissionsReply
	(*timestamppb.Timestamp)(nil),      // 8: google.protobuf.Timestamp
}
var file_sastoj_admin_judge_service_v1_judge_proto_depIdxs = []int32{
	8, // 0: api.sastoj.admin.judge.service.v1.Submission.create_time:type_name -> google.protobuf.Timestamp
	0, // 1: api.sastoj.admin.judge.service.v1.GetJudgableProblemsReply.results:type_name -> api.sastoj.admin.judge.service.v1.Problem
	1, // 2: api.sastoj.admin.judge.service.v1.GetSubmissionsReply.submissions:type_name -> api.sastoj.admin.judge.service.v1.Submission
	2, // 3: api.sastoj.admin.judge.service.v1.Judge.SubmitJudge:input_type -> api.sastoj.admin.judge.service.v1.SubmitJudgeRequest
	4, // 4: api.sastoj.admin.judge.service.v1.Judge.GetJudgableProblems:input_type -> api.sastoj.admin.judge.service.v1.GetJudgableProblemsRequest
	6, // 5: api.sastoj.admin.judge.service.v1.Judge.GetSubmissions:input_type -> api.sastoj.admin.judge.service.v1.GetSubmissionsRequest
	3, // 6: api.sastoj.admin.judge.service.v1.Judge.SubmitJudge:output_type -> api.sastoj.admin.judge.service.v1.SubmitJudgeReply
	5, // 7: api.sastoj.admin.judge.service.v1.Judge.GetJudgableProblems:output_type -> api.sastoj.admin.judge.service.v1.GetJudgableProblemsReply
	7, // 8: api.sastoj.admin.judge.service.v1.Judge.GetSubmissions:output_type -> api.sastoj.admin.judge.service.v1.GetSubmissionsReply
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sastoj_admin_judge_service_v1_judge_proto_init() }
func file_sastoj_admin_judge_service_v1_judge_proto_init() {
	if File_sastoj_admin_judge_service_v1_judge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Problem); i {
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
		file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submission); i {
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
		file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitJudgeRequest); i {
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
		file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitJudgeReply); i {
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
		file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJudgableProblemsRequest); i {
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
		file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJudgableProblemsReply); i {
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
		file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubmissionsRequest); i {
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
		file_sastoj_admin_judge_service_v1_judge_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubmissionsReply); i {
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
			RawDescriptor: file_sastoj_admin_judge_service_v1_judge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sastoj_admin_judge_service_v1_judge_proto_goTypes,
		DependencyIndexes: file_sastoj_admin_judge_service_v1_judge_proto_depIdxs,
		MessageInfos:      file_sastoj_admin_judge_service_v1_judge_proto_msgTypes,
	}.Build()
	File_sastoj_admin_judge_service_v1_judge_proto = out.File
	file_sastoj_admin_judge_service_v1_judge_proto_rawDesc = nil
	file_sastoj_admin_judge_service_v1_judge_proto_goTypes = nil
	file_sastoj_admin_judge_service_v1_judge_proto_depIdxs = nil
}
