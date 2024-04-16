// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: api/sastoj/admin/case/service/v1/case.proto

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

type Case struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Point  int64 `protobuf:"varint,2,opt,name=point,proto3" json:"point,omitempty"`
	Index  int64 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	IsAuto bool  `protobuf:"varint,4,opt,name=is_auto,json=isAuto,proto3" json:"is_auto,omitempty"`
}

func (x *Case) Reset() {
	*x = Case{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Case) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Case) ProtoMessage() {}

func (x *Case) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Case.ProtoReflect.Descriptor instead.
func (*Case) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_case_service_v1_case_proto_rawDescGZIP(), []int{0}
}

func (x *Case) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Case) GetPoint() int64 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *Case) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Case) GetIsAuto() bool {
	if x != nil {
		return x.IsAuto
	}
	return false
}

type CreateCasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId int64   `protobuf:"varint,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Requests  []*Case `protobuf:"bytes,2,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *CreateCasesRequest) Reset() {
	*x = CreateCasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCasesRequest) ProtoMessage() {}

func (x *CreateCasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCasesRequest.ProtoReflect.Descriptor instead.
func (*CreateCasesRequest) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_case_service_v1_case_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCasesRequest) GetProblemId() int64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

func (x *CreateCasesRequest) GetRequests() []*Case {
	if x != nil {
		return x.Requests
	}
	return nil
}

type CreateCasesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *CreateCasesReply) Reset() {
	*x = CreateCasesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCasesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCasesReply) ProtoMessage() {}

func (x *CreateCasesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCasesReply.ProtoReflect.Descriptor instead.
func (*CreateCasesReply) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_case_service_v1_case_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCasesReply) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UpdateCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Point     int64 `protobuf:"varint,2,opt,name=point,proto3" json:"point,omitempty"`
	Index     int64 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	IsAuto    bool  `protobuf:"varint,4,opt,name=is_auto,json=isAuto,proto3" json:"is_auto,omitempty"`
	ProblemId int64 `protobuf:"varint,5,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
}

func (x *UpdateCaseRequest) Reset() {
	*x = UpdateCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCaseRequest) ProtoMessage() {}

func (x *UpdateCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCaseRequest.ProtoReflect.Descriptor instead.
func (*UpdateCaseRequest) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_case_service_v1_case_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCaseRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCaseRequest) GetPoint() int64 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *UpdateCaseRequest) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *UpdateCaseRequest) GetIsAuto() bool {
	if x != nil {
		return x.IsAuto
	}
	return false
}

func (x *UpdateCaseRequest) GetProblemId() int64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

type UpdateCaseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCaseReply) Reset() {
	*x = UpdateCaseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCaseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCaseReply) ProtoMessage() {}

func (x *UpdateCaseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCaseReply.ProtoReflect.Descriptor instead.
func (*UpdateCaseReply) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_case_service_v1_case_proto_rawDescGZIP(), []int{4}
}

type DeleteCaseByCaseIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteCaseByCaseIdsRequest) Reset() {
	*x = DeleteCaseByCaseIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCaseByCaseIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCaseByCaseIdsRequest) ProtoMessage() {}

func (x *DeleteCaseByCaseIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCaseByCaseIdsRequest.ProtoReflect.Descriptor instead.
func (*DeleteCaseByCaseIdsRequest) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_case_service_v1_case_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCaseByCaseIdsRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteCaseByCaseIdsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCaseByCaseIdsReply) Reset() {
	*x = DeleteCaseByCaseIdsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCaseByCaseIdsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCaseByCaseIdsReply) ProtoMessage() {}

func (x *DeleteCaseByCaseIdsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCaseByCaseIdsReply.ProtoReflect.Descriptor instead.
func (*DeleteCaseByCaseIdsReply) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_case_service_v1_case_proto_rawDescGZIP(), []int{6}
}

type DeleteCasesByProblemIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId int64 `protobuf:"varint,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
}

func (x *DeleteCasesByProblemIdRequest) Reset() {
	*x = DeleteCasesByProblemIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCasesByProblemIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCasesByProblemIdRequest) ProtoMessage() {}

func (x *DeleteCasesByProblemIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCasesByProblemIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteCasesByProblemIdRequest) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_case_service_v1_case_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCasesByProblemIdRequest) GetProblemId() int64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

type DeleteCasesByProblemIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCasesByProblemIdReply) Reset() {
	*x = DeleteCasesByProblemIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCasesByProblemIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCasesByProblemIdReply) ProtoMessage() {}

func (x *DeleteCasesByProblemIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCasesByProblemIdReply.ProtoReflect.Descriptor instead.
func (*DeleteCasesByProblemIdReply) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_case_service_v1_case_proto_rawDescGZIP(), []int{8}
}

type GetCasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId int64 `protobuf:"varint,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
}

func (x *GetCasesRequest) Reset() {
	*x = GetCasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCasesRequest) ProtoMessage() {}

func (x *GetCasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCasesRequest.ProtoReflect.Descriptor instead.
func (*GetCasesRequest) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_case_service_v1_case_proto_rawDescGZIP(), []int{9}
}

func (x *GetCasesRequest) GetProblemId() int64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

type GetCasesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Case `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GetCasesReply) Reset() {
	*x = GetCasesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCasesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCasesReply) ProtoMessage() {}

func (x *GetCasesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCasesReply.ProtoReflect.Descriptor instead.
func (*GetCasesReply) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_case_service_v1_case_proto_rawDescGZIP(), []int{10}
}

func (x *GetCasesReply) GetResults() []*Case {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_api_sastoj_admin_case_service_v1_case_proto protoreflect.FileDescriptor

var file_api_sastoj_admin_case_service_v1_case_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a,
	0x04, 0x43, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12,
	0x42, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x73, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x41, 0x75, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x49, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2e, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x73, 0x65, 0x42, 0x79, 0x43, 0x61, 0x73, 0x65, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x73, 0x65, 0x42, 0x79, 0x43, 0x61, 0x73, 0x65, 0x49, 0x64, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x3e, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65,
	0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x49, 0x64, 0x22, 0x1d, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65,
	0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74,
	0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x9f, 0x06, 0x0a, 0x0b, 0x43, 0x61, 0x73, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x73, 0x65, 0x73, 0x12, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73,
	0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x86, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x12, 0x33,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a, 0x01,
	0x2a, 0x1a, 0x05, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x12, 0xa9, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x73, 0x42, 0x79, 0x43, 0x61, 0x73, 0x65, 0x49, 0x64,
	0x73, 0x12, 0x3c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x42,
	0x79, 0x43, 0x61, 0x73, 0x65, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x42, 0x79, 0x43,
	0x61, 0x73, 0x65, 0x49, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x1a, 0x0c, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0xb4, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x73, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12,
	0x3f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x73, 0x42, 0x79,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x73, 0x42,
	0x79, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x8a, 0x01, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63,
	0x61, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x51, 0x0a, 0x20, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x2b,
	0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x61, 0x73, 0x74,
	0x6f, 0x6a, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_sastoj_admin_case_service_v1_case_proto_rawDescOnce sync.Once
	file_api_sastoj_admin_case_service_v1_case_proto_rawDescData = file_api_sastoj_admin_case_service_v1_case_proto_rawDesc
)

func file_api_sastoj_admin_case_service_v1_case_proto_rawDescGZIP() []byte {
	file_api_sastoj_admin_case_service_v1_case_proto_rawDescOnce.Do(func() {
		file_api_sastoj_admin_case_service_v1_case_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_sastoj_admin_case_service_v1_case_proto_rawDescData)
	})
	return file_api_sastoj_admin_case_service_v1_case_proto_rawDescData
}

var file_api_sastoj_admin_case_service_v1_case_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_sastoj_admin_case_service_v1_case_proto_goTypes = []interface{}{
	(*Case)(nil),                          // 0: api.sastoj.admin.case.service.v1.Case
	(*CreateCasesRequest)(nil),            // 1: api.sastoj.admin.case.service.v1.CreateCasesRequest
	(*CreateCasesReply)(nil),              // 2: api.sastoj.admin.case.service.v1.CreateCasesReply
	(*UpdateCaseRequest)(nil),             // 3: api.sastoj.admin.case.service.v1.UpdateCaseRequest
	(*UpdateCaseReply)(nil),               // 4: api.sastoj.admin.case.service.v1.UpdateCaseReply
	(*DeleteCaseByCaseIdsRequest)(nil),    // 5: api.sastoj.admin.case.service.v1.DeleteCaseByCaseIdsRequest
	(*DeleteCaseByCaseIdsReply)(nil),      // 6: api.sastoj.admin.case.service.v1.DeleteCaseByCaseIdsReply
	(*DeleteCasesByProblemIdRequest)(nil), // 7: api.sastoj.admin.case.service.v1.DeleteCasesByProblemIdRequest
	(*DeleteCasesByProblemIdReply)(nil),   // 8: api.sastoj.admin.case.service.v1.DeleteCasesByProblemIdReply
	(*GetCasesRequest)(nil),               // 9: api.sastoj.admin.case.service.v1.GetCasesRequest
	(*GetCasesReply)(nil),                 // 10: api.sastoj.admin.case.service.v1.GetCasesReply
}
var file_api_sastoj_admin_case_service_v1_case_proto_depIdxs = []int32{
	0,  // 0: api.sastoj.admin.case.service.v1.CreateCasesRequest.requests:type_name -> api.sastoj.admin.case.service.v1.Case
	0,  // 1: api.sastoj.admin.case.service.v1.GetCasesReply.results:type_name -> api.sastoj.admin.case.service.v1.Case
	1,  // 2: api.sastoj.admin.case.service.v1.CaseService.CreateCases:input_type -> api.sastoj.admin.case.service.v1.CreateCasesRequest
	3,  // 3: api.sastoj.admin.case.service.v1.CaseService.UpdateCase:input_type -> api.sastoj.admin.case.service.v1.UpdateCaseRequest
	5,  // 4: api.sastoj.admin.case.service.v1.CaseService.DeleteCasesByCaseIds:input_type -> api.sastoj.admin.case.service.v1.DeleteCaseByCaseIdsRequest
	7,  // 5: api.sastoj.admin.case.service.v1.CaseService.DeleteCasesByProblemId:input_type -> api.sastoj.admin.case.service.v1.DeleteCasesByProblemIdRequest
	9,  // 6: api.sastoj.admin.case.service.v1.CaseService.GetCases:input_type -> api.sastoj.admin.case.service.v1.GetCasesRequest
	2,  // 7: api.sastoj.admin.case.service.v1.CaseService.CreateCases:output_type -> api.sastoj.admin.case.service.v1.CreateCasesReply
	4,  // 8: api.sastoj.admin.case.service.v1.CaseService.UpdateCase:output_type -> api.sastoj.admin.case.service.v1.UpdateCaseReply
	6,  // 9: api.sastoj.admin.case.service.v1.CaseService.DeleteCasesByCaseIds:output_type -> api.sastoj.admin.case.service.v1.DeleteCaseByCaseIdsReply
	8,  // 10: api.sastoj.admin.case.service.v1.CaseService.DeleteCasesByProblemId:output_type -> api.sastoj.admin.case.service.v1.DeleteCasesByProblemIdReply
	10, // 11: api.sastoj.admin.case.service.v1.CaseService.GetCases:output_type -> api.sastoj.admin.case.service.v1.GetCasesReply
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_sastoj_admin_case_service_v1_case_proto_init() }
func file_api_sastoj_admin_case_service_v1_case_proto_init() {
	if File_api_sastoj_admin_case_service_v1_case_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Case); i {
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
		file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCasesRequest); i {
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
		file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCasesReply); i {
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
		file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCaseRequest); i {
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
		file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCaseReply); i {
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
		file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCaseByCaseIdsRequest); i {
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
		file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCaseByCaseIdsReply); i {
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
		file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCasesByProblemIdRequest); i {
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
		file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCasesByProblemIdReply); i {
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
		file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCasesRequest); i {
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
		file_api_sastoj_admin_case_service_v1_case_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCasesReply); i {
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
			RawDescriptor: file_api_sastoj_admin_case_service_v1_case_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_sastoj_admin_case_service_v1_case_proto_goTypes,
		DependencyIndexes: file_api_sastoj_admin_case_service_v1_case_proto_depIdxs,
		MessageInfos:      file_api_sastoj_admin_case_service_v1_case_proto_msgTypes,
	}.Build()
	File_api_sastoj_admin_case_service_v1_case_proto = out.File
	file_api_sastoj_admin_case_service_v1_case_proto_rawDesc = nil
	file_api_sastoj_admin_case_service_v1_case_proto_goTypes = nil
	file_api_sastoj_admin_case_service_v1_case_proto_depIdxs = nil
}
