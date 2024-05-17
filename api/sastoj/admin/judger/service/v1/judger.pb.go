// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: api/sastoj/admin/judger/service/v1/judger.proto

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

type UpdateJudgerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId      int64   `protobuf:"varint,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	JudgerGroupIds []int64 `protobuf:"varint,2,rep,packed,name=judger_group_ids,json=judgerGroupIds,proto3" json:"judger_group_ids,omitempty"`
}

func (x *UpdateJudgerRequest) Reset() {
	*x = UpdateJudgerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJudgerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJudgerRequest) ProtoMessage() {}

func (x *UpdateJudgerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJudgerRequest.ProtoReflect.Descriptor instead.
func (*UpdateJudgerRequest) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_judger_service_v1_judger_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateJudgerRequest) GetProblemId() int64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

func (x *UpdateJudgerRequest) GetJudgerGroupIds() []int64 {
	if x != nil {
		return x.JudgerGroupIds
	}
	return nil
}

type UpdateJudgerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateJudgerReply) Reset() {
	*x = UpdateJudgerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJudgerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJudgerReply) ProtoMessage() {}

func (x *UpdateJudgerReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJudgerReply.ProtoReflect.Descriptor instead.
func (*UpdateJudgerReply) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_judger_service_v1_judger_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateJudgerReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetJudgerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId int64 `protobuf:"varint,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
}

func (x *GetJudgerRequest) Reset() {
	*x = GetJudgerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJudgerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJudgerRequest) ProtoMessage() {}

func (x *GetJudgerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJudgerRequest.ProtoReflect.Descriptor instead.
func (*GetJudgerRequest) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_judger_service_v1_judger_proto_rawDescGZIP(), []int{2}
}

func (x *GetJudgerRequest) GetProblemId() int64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

type GetJudgerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*GetJudgerReply_Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *GetJudgerReply) Reset() {
	*x = GetJudgerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJudgerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJudgerReply) ProtoMessage() {}

func (x *GetJudgerReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJudgerReply.ProtoReflect.Descriptor instead.
func (*GetJudgerReply) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_judger_service_v1_judger_proto_rawDescGZIP(), []int{3}
}

func (x *GetJudgerReply) GetGroups() []*GetJudgerReply_Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type GetJudgerReply_Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetJudgerReply_Group) Reset() {
	*x = GetJudgerReply_Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJudgerReply_Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJudgerReply_Group) ProtoMessage() {}

func (x *GetJudgerReply_Group) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJudgerReply_Group.ProtoReflect.Descriptor instead.
func (*GetJudgerReply_Group) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_judger_service_v1_judger_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetJudgerReply_Group) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetJudgerReply_Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_sastoj_admin_judger_service_v1_judger_proto protoreflect.FileDescriptor

var file_api_sastoj_admin_judger_service_v1_judger_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x75, 0x64,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x0e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x73, 0x22, 0x2d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x75, 0x64,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x31, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x64,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x50, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x2b, 0x0a, 0x05, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xc0, 0x02, 0x0a, 0x06, 0x4a, 0x75, 0x64, 0x67,
	0x65, 0x72, 0x12, 0x9f, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x75, 0x64,
	0x67, 0x65, 0x72, 0x12, 0x37, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a,
	0x75, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14,
	0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0x93, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67,
	0x65, 0x72, 0x12, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x56, 0x0a, 0x23, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x2d, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_sastoj_admin_judger_service_v1_judger_proto_rawDescOnce sync.Once
	file_api_sastoj_admin_judger_service_v1_judger_proto_rawDescData = file_api_sastoj_admin_judger_service_v1_judger_proto_rawDesc
)

func file_api_sastoj_admin_judger_service_v1_judger_proto_rawDescGZIP() []byte {
	file_api_sastoj_admin_judger_service_v1_judger_proto_rawDescOnce.Do(func() {
		file_api_sastoj_admin_judger_service_v1_judger_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_sastoj_admin_judger_service_v1_judger_proto_rawDescData)
	})
	return file_api_sastoj_admin_judger_service_v1_judger_proto_rawDescData
}

var file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_sastoj_admin_judger_service_v1_judger_proto_goTypes = []interface{}{
	(*UpdateJudgerRequest)(nil),  // 0: api.sastoj.admin.judger.service.v1.UpdateJudgerRequest
	(*UpdateJudgerReply)(nil),    // 1: api.sastoj.admin.judger.service.v1.UpdateJudgerReply
	(*GetJudgerRequest)(nil),     // 2: api.sastoj.admin.judger.service.v1.GetJudgerRequest
	(*GetJudgerReply)(nil),       // 3: api.sastoj.admin.judger.service.v1.GetJudgerReply
	(*GetJudgerReply_Group)(nil), // 4: api.sastoj.admin.judger.service.v1.GetJudgerReply.Group
}
var file_api_sastoj_admin_judger_service_v1_judger_proto_depIdxs = []int32{
	4, // 0: api.sastoj.admin.judger.service.v1.GetJudgerReply.groups:type_name -> api.sastoj.admin.judger.service.v1.GetJudgerReply.Group
	0, // 1: api.sastoj.admin.judger.service.v1.Judger.UpdateJudger:input_type -> api.sastoj.admin.judger.service.v1.UpdateJudgerRequest
	2, // 2: api.sastoj.admin.judger.service.v1.Judger.GetJudger:input_type -> api.sastoj.admin.judger.service.v1.GetJudgerRequest
	1, // 3: api.sastoj.admin.judger.service.v1.Judger.UpdateJudger:output_type -> api.sastoj.admin.judger.service.v1.UpdateJudgerReply
	3, // 4: api.sastoj.admin.judger.service.v1.Judger.GetJudger:output_type -> api.sastoj.admin.judger.service.v1.GetJudgerReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_sastoj_admin_judger_service_v1_judger_proto_init() }
func file_api_sastoj_admin_judger_service_v1_judger_proto_init() {
	if File_api_sastoj_admin_judger_service_v1_judger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJudgerRequest); i {
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
		file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJudgerReply); i {
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
		file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJudgerRequest); i {
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
		file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJudgerReply); i {
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
		file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJudgerReply_Group); i {
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
			RawDescriptor: file_api_sastoj_admin_judger_service_v1_judger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_sastoj_admin_judger_service_v1_judger_proto_goTypes,
		DependencyIndexes: file_api_sastoj_admin_judger_service_v1_judger_proto_depIdxs,
		MessageInfos:      file_api_sastoj_admin_judger_service_v1_judger_proto_msgTypes,
	}.Build()
	File_api_sastoj_admin_judger_service_v1_judger_proto = out.File
	file_api_sastoj_admin_judger_service_v1_judger_proto_rawDesc = nil
	file_api_sastoj_admin_judger_service_v1_judger_proto_goTypes = nil
	file_api_sastoj_admin_judger_service_v1_judger_proto_depIdxs = nil
}
