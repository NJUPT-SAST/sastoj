// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: api/sastoj/admin/group/service/v1/group.proto

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

type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_group_service_v1_group_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateGroupReply) Reset() {
	*x = CreateGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupReply) ProtoMessage() {}

func (x *CreateGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupReply.ProtoReflect.Descriptor instead.
func (*CreateGroupReply) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_group_service_v1_group_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGroupReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateGroupRequest) Reset() {
	*x = UpdateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupRequest) ProtoMessage() {}

func (x *UpdateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_group_service_v1_group_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateGroupRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateGroupReply) Reset() {
	*x = UpdateGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupReply) ProtoMessage() {}

func (x *UpdateGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupReply.ProtoReflect.Descriptor instead.
func (*UpdateGroupReply) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_group_service_v1_group_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateGroupReply) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

type DeleteGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGroupRequest) Reset() {
	*x = DeleteGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupRequest) ProtoMessage() {}

func (x *DeleteGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_group_service_v1_group_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteGroupRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteGroupReply) Reset() {
	*x = DeleteGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupReply) ProtoMessage() {}

func (x *DeleteGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupReply.ProtoReflect.Descriptor instead.
func (*DeleteGroupReply) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_group_service_v1_group_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteGroupReply) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

type GetGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGroupRequest) Reset() {
	*x = GetGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupRequest) ProtoMessage() {}

func (x *GetGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupRequest.ProtoReflect.Descriptor instead.
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_group_service_v1_group_proto_rawDescGZIP(), []int{6}
}

func (x *GetGroupRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetGroupReply) Reset() {
	*x = GetGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupReply) ProtoMessage() {}

func (x *GetGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupReply.ProtoReflect.Descriptor instead.
func (*GetGroupReply) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_group_service_v1_group_proto_rawDescGZIP(), []int{7}
}

func (x *GetGroupReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetGroupReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListGroupRequest) Reset() {
	*x = ListGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupRequest) ProtoMessage() {}

func (x *ListGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupRequest.ProtoReflect.Descriptor instead.
func (*ListGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_group_service_v1_group_proto_rawDescGZIP(), []int{8}
}

type ListGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*ListGroupReply_Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *ListGroupReply) Reset() {
	*x = ListGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupReply) ProtoMessage() {}

func (x *ListGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupReply.ProtoReflect.Descriptor instead.
func (*ListGroupReply) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_group_service_v1_group_proto_rawDescGZIP(), []int{9}
}

func (x *ListGroupReply) GetGroups() []*ListGroupReply_Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type ListGroupReply_Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListGroupReply_Group) Reset() {
	*x = ListGroupReply_Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupReply_Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupReply_Group) ProtoMessage() {}

func (x *ListGroupReply_Group) ProtoReflect() protoreflect.Message {
	mi := &file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupReply_Group.ProtoReflect.Descriptor instead.
func (*ListGroupReply_Group) Descriptor() ([]byte, []int) {
	return file_api_sastoj_admin_group_service_v1_group_proto_rawDescGZIP(), []int{9, 0}
}

func (x *ListGroupReply_Group) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListGroupReply_Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_sastoj_admin_group_service_v1_group_proto protoreflect.FileDescriptor

var file_api_sastoj_admin_group_service_v1_group_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x28, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4f, 0x0a, 0x06, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x2b, 0x0a, 0x05, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x89, 0x05, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x8c, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x35, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x3a, 0x01, 0x2a, 0x1a, 0x06, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x79, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x35, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73,
	0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x79, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x35, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x85, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61,
	0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0d, 0x12, 0x0b, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x73,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x33, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x52, 0x0a, 0x21, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f,
	0x6a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x2b, 0x73, 0x61, 0x73, 0x74,
	0x6f, 0x6a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_sastoj_admin_group_service_v1_group_proto_rawDescOnce sync.Once
	file_api_sastoj_admin_group_service_v1_group_proto_rawDescData = file_api_sastoj_admin_group_service_v1_group_proto_rawDesc
)

func file_api_sastoj_admin_group_service_v1_group_proto_rawDescGZIP() []byte {
	file_api_sastoj_admin_group_service_v1_group_proto_rawDescOnce.Do(func() {
		file_api_sastoj_admin_group_service_v1_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_sastoj_admin_group_service_v1_group_proto_rawDescData)
	})
	return file_api_sastoj_admin_group_service_v1_group_proto_rawDescData
}

var file_api_sastoj_admin_group_service_v1_group_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_sastoj_admin_group_service_v1_group_proto_goTypes = []interface{}{
	(*CreateGroupRequest)(nil),   // 0: api.sastoj.admin.group.service.v1.CreateGroupRequest
	(*CreateGroupReply)(nil),     // 1: api.sastoj.admin.group.service.v1.CreateGroupReply
	(*UpdateGroupRequest)(nil),   // 2: api.sastoj.admin.group.service.v1.UpdateGroupRequest
	(*UpdateGroupReply)(nil),     // 3: api.sastoj.admin.group.service.v1.UpdateGroupReply
	(*DeleteGroupRequest)(nil),   // 4: api.sastoj.admin.group.service.v1.DeleteGroupRequest
	(*DeleteGroupReply)(nil),     // 5: api.sastoj.admin.group.service.v1.DeleteGroupReply
	(*GetGroupRequest)(nil),      // 6: api.sastoj.admin.group.service.v1.GetGroupRequest
	(*GetGroupReply)(nil),        // 7: api.sastoj.admin.group.service.v1.GetGroupReply
	(*ListGroupRequest)(nil),     // 8: api.sastoj.admin.group.service.v1.ListGroupRequest
	(*ListGroupReply)(nil),       // 9: api.sastoj.admin.group.service.v1.ListGroupReply
	(*ListGroupReply_Group)(nil), // 10: api.sastoj.admin.group.service.v1.ListGroupReply.Group
}
var file_api_sastoj_admin_group_service_v1_group_proto_depIdxs = []int32{
	10, // 0: api.sastoj.admin.group.service.v1.ListGroupReply.groups:type_name -> api.sastoj.admin.group.service.v1.ListGroupReply.Group
	0,  // 1: api.sastoj.admin.group.service.v1.Group.CreateGroup:input_type -> api.sastoj.admin.group.service.v1.CreateGroupRequest
	2,  // 2: api.sastoj.admin.group.service.v1.Group.UpdateGroup:input_type -> api.sastoj.admin.group.service.v1.UpdateGroupRequest
	4,  // 3: api.sastoj.admin.group.service.v1.Group.DeleteGroup:input_type -> api.sastoj.admin.group.service.v1.DeleteGroupRequest
	6,  // 4: api.sastoj.admin.group.service.v1.Group.GetGroup:input_type -> api.sastoj.admin.group.service.v1.GetGroupRequest
	8,  // 5: api.sastoj.admin.group.service.v1.Group.ListGroup:input_type -> api.sastoj.admin.group.service.v1.ListGroupRequest
	1,  // 6: api.sastoj.admin.group.service.v1.Group.CreateGroup:output_type -> api.sastoj.admin.group.service.v1.CreateGroupReply
	3,  // 7: api.sastoj.admin.group.service.v1.Group.UpdateGroup:output_type -> api.sastoj.admin.group.service.v1.UpdateGroupReply
	5,  // 8: api.sastoj.admin.group.service.v1.Group.DeleteGroup:output_type -> api.sastoj.admin.group.service.v1.DeleteGroupReply
	7,  // 9: api.sastoj.admin.group.service.v1.Group.GetGroup:output_type -> api.sastoj.admin.group.service.v1.GetGroupReply
	9,  // 10: api.sastoj.admin.group.service.v1.Group.ListGroup:output_type -> api.sastoj.admin.group.service.v1.ListGroupReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_sastoj_admin_group_service_v1_group_proto_init() }
func file_api_sastoj_admin_group_service_v1_group_proto_init() {
	if File_api_sastoj_admin_group_service_v1_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupReply); i {
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
		file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupRequest); i {
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
		file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupReply); i {
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
		file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupRequest); i {
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
		file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupReply); i {
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
		file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupRequest); i {
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
		file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupReply); i {
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
		file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupRequest); i {
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
		file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupReply); i {
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
		file_api_sastoj_admin_group_service_v1_group_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupReply_Group); i {
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
			RawDescriptor: file_api_sastoj_admin_group_service_v1_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_sastoj_admin_group_service_v1_group_proto_goTypes,
		DependencyIndexes: file_api_sastoj_admin_group_service_v1_group_proto_depIdxs,
		MessageInfos:      file_api_sastoj_admin_group_service_v1_group_proto_msgTypes,
	}.Build()
	File_api_sastoj_admin_group_service_v1_group_proto = out.File
	file_api_sastoj_admin_group_service_v1_group_proto_rawDesc = nil
	file_api_sastoj_admin_group_service_v1_group_proto_goTypes = nil
	file_api_sastoj_admin_group_service_v1_group_proto_depIdxs = nil
}
