// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: basicService/v1/verify.proto

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

type CreateVerificationCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 验证码的位数
	Bits int64 `protobuf:"varint,1,opt,name=bits,proto3" json:"bits,omitempty"`
	// 过期时间戳，毫秒
	ExpireTime int64 `protobuf:"varint,2,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *CreateVerificationCodeRequest) Reset() {
	*x = CreateVerificationCodeRequest{}
	mi := &file_basicService_v1_verify_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVerificationCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVerificationCodeRequest) ProtoMessage() {}

func (x *CreateVerificationCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basicService_v1_verify_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVerificationCodeRequest.ProtoReflect.Descriptor instead.
func (*CreateVerificationCodeRequest) Descriptor() ([]byte, []int) {
	return file_basicService_v1_verify_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVerificationCodeRequest) GetBits() int64 {
	if x != nil {
		return x.Bits
	}
	return 0
}

func (x *CreateVerificationCodeRequest) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

type CreateVerificationCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta               *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	VerificationCodeId int64     `protobuf:"varint,2,opt,name=verification_code_id,json=verificationCodeId,proto3" json:"verification_code_id,omitempty"`
}

func (x *CreateVerificationCodeResponse) Reset() {
	*x = CreateVerificationCodeResponse{}
	mi := &file_basicService_v1_verify_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVerificationCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVerificationCodeResponse) ProtoMessage() {}

func (x *CreateVerificationCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basicService_v1_verify_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVerificationCodeResponse.ProtoReflect.Descriptor instead.
func (*CreateVerificationCodeResponse) Descriptor() ([]byte, []int) {
	return file_basicService_v1_verify_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVerificationCodeResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *CreateVerificationCodeResponse) GetVerificationCodeId() int64 {
	if x != nil {
		return x.VerificationCodeId
	}
	return 0
}

type ValidateVerificationCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerificationCodeId int64  `protobuf:"varint,1,opt,name=verification_code_id,json=verificationCodeId,proto3" json:"verification_code_id,omitempty"`
	Code               string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ValidateVerificationCodeRequest) Reset() {
	*x = ValidateVerificationCodeRequest{}
	mi := &file_basicService_v1_verify_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateVerificationCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateVerificationCodeRequest) ProtoMessage() {}

func (x *ValidateVerificationCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basicService_v1_verify_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateVerificationCodeRequest.ProtoReflect.Descriptor instead.
func (*ValidateVerificationCodeRequest) Descriptor() ([]byte, []int) {
	return file_basicService_v1_verify_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateVerificationCodeRequest) GetVerificationCodeId() int64 {
	if x != nil {
		return x.VerificationCodeId
	}
	return 0
}

func (x *ValidateVerificationCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ValidateVerificationCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ValidateVerificationCodeResponse) Reset() {
	*x = ValidateVerificationCodeResponse{}
	mi := &file_basicService_v1_verify_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateVerificationCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateVerificationCodeResponse) ProtoMessage() {}

func (x *ValidateVerificationCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basicService_v1_verify_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateVerificationCodeResponse.ProtoReflect.Descriptor instead.
func (*ValidateVerificationCodeResponse) Descriptor() ([]byte, []int) {
	return file_basicService_v1_verify_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateVerificationCodeResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_basicService_v1_verify_proto protoreflect.FileDescriptor

var file_basicService_v1_verify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a,
	0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62, 0x69,
	0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x14, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x1f, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x51, 0x0a, 0x20, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x32, 0x89, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x79, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7f,
	0x0a, 0x18, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x58,
	0x4a, 0x30, 0x30, 0x30, 0x30, 0x2f, 0x74, 0x6f, 0x6b, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basicService_v1_verify_proto_rawDescOnce sync.Once
	file_basicService_v1_verify_proto_rawDescData = file_basicService_v1_verify_proto_rawDesc
)

func file_basicService_v1_verify_proto_rawDescGZIP() []byte {
	file_basicService_v1_verify_proto_rawDescOnce.Do(func() {
		file_basicService_v1_verify_proto_rawDescData = protoimpl.X.CompressGZIP(file_basicService_v1_verify_proto_rawDescData)
	})
	return file_basicService_v1_verify_proto_rawDescData
}

var file_basicService_v1_verify_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_basicService_v1_verify_proto_goTypes = []any{
	(*CreateVerificationCodeRequest)(nil),    // 0: basicService.v1.CreateVerificationCodeRequest
	(*CreateVerificationCodeResponse)(nil),   // 1: basicService.v1.CreateVerificationCodeResponse
	(*ValidateVerificationCodeRequest)(nil),  // 2: basicService.v1.ValidateVerificationCodeRequest
	(*ValidateVerificationCodeResponse)(nil), // 3: basicService.v1.ValidateVerificationCodeResponse
	(*Metadata)(nil),                         // 4: basicService.v1.Metadata
}
var file_basicService_v1_verify_proto_depIdxs = []int32{
	4, // 0: basicService.v1.CreateVerificationCodeResponse.meta:type_name -> basicService.v1.Metadata
	4, // 1: basicService.v1.ValidateVerificationCodeResponse.meta:type_name -> basicService.v1.Metadata
	0, // 2: basicService.v1.AuthService.CreateVerificationCode:input_type -> basicService.v1.CreateVerificationCodeRequest
	2, // 3: basicService.v1.AuthService.ValidateVerificationCode:input_type -> basicService.v1.ValidateVerificationCodeRequest
	1, // 4: basicService.v1.AuthService.CreateVerificationCode:output_type -> basicService.v1.CreateVerificationCodeResponse
	3, // 5: basicService.v1.AuthService.ValidateVerificationCode:output_type -> basicService.v1.ValidateVerificationCodeResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_basicService_v1_verify_proto_init() }
func file_basicService_v1_verify_proto_init() {
	if File_basicService_v1_verify_proto != nil {
		return
	}
	file_basicService_v1_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_basicService_v1_verify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basicService_v1_verify_proto_goTypes,
		DependencyIndexes: file_basicService_v1_verify_proto_depIdxs,
		MessageInfos:      file_basicService_v1_verify_proto_msgTypes,
	}.Build()
	File_basicService_v1_verify_proto = out.File
	file_basicService_v1_verify_proto_rawDesc = nil
	file_basicService_v1_verify_proto_goTypes = nil
	file_basicService_v1_verify_proto_depIdxs = nil
}
