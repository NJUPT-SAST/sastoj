// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: sastoj/gojudge/judger/service/v1/judge.proto

package v1

import (
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

type JudgeType int32

const (
	JudgeType_rsjudge JudgeType = 0
	JudgeType_gojudge JudgeType = 1
)

// Enum value maps for JudgeType.
var (
	JudgeType_name = map[int32]string{
		0: "rsjudge",
		1: "gojudge",
	}
	JudgeType_value = map[string]int32{
		"rsjudge": 0,
		"gojudge": 1,
	}
)

func (x JudgeType) Enum() *JudgeType {
	p := new(JudgeType)
	*p = x
	return p
}

func (x JudgeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JudgeType) Descriptor() protoreflect.EnumDescriptor {
	return file_sastoj_gojudge_judger_service_v1_judge_proto_enumTypes[0].Descriptor()
}

func (JudgeType) Type() protoreflect.EnumType {
	return &file_sastoj_gojudge_judger_service_v1_judge_proto_enumTypes[0]
}

func (x JudgeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JudgeType.Descriptor instead.
func (JudgeType) EnumDescriptor() ([]byte, []int) {
	return file_sastoj_gojudge_judger_service_v1_judge_proto_rawDescGZIP(), []int{0}
}

type ListAllStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAllStatusRequest) Reset() {
	*x = ListAllStatusRequest{}
	mi := &file_sastoj_gojudge_judger_service_v1_judge_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllStatusRequest) ProtoMessage() {}

func (x *ListAllStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sastoj_gojudge_judger_service_v1_judge_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllStatusRequest.ProtoReflect.Descriptor instead.
func (*ListAllStatusRequest) Descriptor() ([]byte, []int) {
	return file_sastoj_gojudge_judger_service_v1_judge_proto_rawDescGZIP(), []int{0}
}

type ListAllStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status []*Status `protobuf:"bytes,2,rep,name=status,proto3" json:"status,omitempty"`
}

func (x *ListAllStatusReply) Reset() {
	*x = ListAllStatusReply{}
	mi := &file_sastoj_gojudge_judger_service_v1_judge_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllStatusReply) ProtoMessage() {}

func (x *ListAllStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_sastoj_gojudge_judger_service_v1_judge_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllStatusReply.ProtoReflect.Descriptor instead.
func (*ListAllStatusReply) Descriptor() ([]byte, []int) {
	return file_sastoj_gojudge_judger_service_v1_judge_proto_rawDescGZIP(), []int{1}
}

func (x *ListAllStatusReply) GetStatus() []*Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     JudgeType `protobuf:"varint,2,opt,name=type,proto3,enum=api.sastoj.gojudge.judger.service.v1.JudgeType" json:"type,omitempty"`
	Endpoint string    `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	mi := &file_sastoj_gojudge_judger_service_v1_judge_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_sastoj_gojudge_judger_service_v1_judge_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_sastoj_gojudge_judger_service_v1_judge_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetType() JudgeType {
	if x != nil {
		return x.Type
	}
	return JudgeType_rsjudge
}

func (x *Status) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

var File_sastoj_gojudge_judger_service_v1_judge_proto protoreflect.FileDescriptor

var file_sastoj_gojudge_judger_service_v1_judge_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f, 0x67, 0x6f, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x67, 0x6f, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5a, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e,
	0x67, 0x6f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x69, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x43, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x67, 0x6f,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2a, 0x25, 0x0a, 0x09, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x72, 0x73, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x67, 0x6f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x10, 0x01, 0x32, 0x8f, 0x01, 0x0a, 0x05, 0x4a,
	0x75, 0x64, 0x67, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73,
	0x74, 0x6f, 0x6a, 0x2e, 0x67, 0x6f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e,
	0x67, 0x6f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x9c, 0x02, 0x0a,
	0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e,
	0x67, 0x6f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4a, 0x75, 0x64, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2f, 0x67, 0x6f, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x41, 0x53, 0x47, 0x4a, 0x53, 0xaa, 0x02,
	0x24, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x2e, 0x47, 0x6f, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x24, 0x41, 0x70, 0x69, 0x5c, 0x53, 0x61, 0x73, 0x74,
	0x6f, 0x6a, 0x5c, 0x47, 0x6f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x5c, 0x4a, 0x75, 0x64, 0x67, 0x65,
	0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x30, 0x41,
	0x70, 0x69, 0x5c, 0x53, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x5c, 0x47, 0x6f, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x5c, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x29, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x53, 0x61, 0x73, 0x74, 0x6f, 0x6a, 0x3a, 0x3a, 0x47,
	0x6f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x3a, 0x3a, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x3a, 0x3a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sastoj_gojudge_judger_service_v1_judge_proto_rawDescOnce sync.Once
	file_sastoj_gojudge_judger_service_v1_judge_proto_rawDescData = file_sastoj_gojudge_judger_service_v1_judge_proto_rawDesc
)

func file_sastoj_gojudge_judger_service_v1_judge_proto_rawDescGZIP() []byte {
	file_sastoj_gojudge_judger_service_v1_judge_proto_rawDescOnce.Do(func() {
		file_sastoj_gojudge_judger_service_v1_judge_proto_rawDescData = protoimpl.X.CompressGZIP(file_sastoj_gojudge_judger_service_v1_judge_proto_rawDescData)
	})
	return file_sastoj_gojudge_judger_service_v1_judge_proto_rawDescData
}

var file_sastoj_gojudge_judger_service_v1_judge_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sastoj_gojudge_judger_service_v1_judge_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sastoj_gojudge_judger_service_v1_judge_proto_goTypes = []any{
	(JudgeType)(0),               // 0: api.sastoj.gojudge.judger.service.v1.JudgeType
	(*ListAllStatusRequest)(nil), // 1: api.sastoj.gojudge.judger.service.v1.ListAllStatusRequest
	(*ListAllStatusReply)(nil),   // 2: api.sastoj.gojudge.judger.service.v1.ListAllStatusReply
	(*Status)(nil),               // 3: api.sastoj.gojudge.judger.service.v1.Status
}
var file_sastoj_gojudge_judger_service_v1_judge_proto_depIdxs = []int32{
	3, // 0: api.sastoj.gojudge.judger.service.v1.ListAllStatusReply.status:type_name -> api.sastoj.gojudge.judger.service.v1.Status
	0, // 1: api.sastoj.gojudge.judger.service.v1.Status.type:type_name -> api.sastoj.gojudge.judger.service.v1.JudgeType
	1, // 2: api.sastoj.gojudge.judger.service.v1.Judge.ListAllStatus:input_type -> api.sastoj.gojudge.judger.service.v1.ListAllStatusRequest
	2, // 3: api.sastoj.gojudge.judger.service.v1.Judge.ListAllStatus:output_type -> api.sastoj.gojudge.judger.service.v1.ListAllStatusReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sastoj_gojudge_judger_service_v1_judge_proto_init() }
func file_sastoj_gojudge_judger_service_v1_judge_proto_init() {
	if File_sastoj_gojudge_judger_service_v1_judge_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sastoj_gojudge_judger_service_v1_judge_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sastoj_gojudge_judger_service_v1_judge_proto_goTypes,
		DependencyIndexes: file_sastoj_gojudge_judger_service_v1_judge_proto_depIdxs,
		EnumInfos:         file_sastoj_gojudge_judger_service_v1_judge_proto_enumTypes,
		MessageInfos:      file_sastoj_gojudge_judger_service_v1_judge_proto_msgTypes,
	}.Build()
	File_sastoj_gojudge_judger_service_v1_judge_proto = out.File
	file_sastoj_gojudge_judger_service_v1_judge_proto_rawDesc = nil
	file_sastoj_gojudge_judger_service_v1_judge_proto_goTypes = nil
	file_sastoj_gojudge_judger_service_v1_judge_proto_depIdxs = nil
}
