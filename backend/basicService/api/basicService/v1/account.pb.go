// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: basicService/v1/account.proto

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

type VoucherType int32

const (
	VoucherType_VOUCHER_PHONE VoucherType = 0
	VoucherType_VOUCHER_EMAIL VoucherType = 1
)

// Enum value maps for VoucherType.
var (
	VoucherType_name = map[int32]string{
		0: "VOUCHER_PHONE",
		1: "VOUCHER_EMAIL",
	}
	VoucherType_value = map[string]int32{
		"VOUCHER_PHONE": 0,
		"VOUCHER_EMAIL": 1,
	}
)

func (x VoucherType) Enum() *VoucherType {
	p := new(VoucherType)
	*p = x
	return p
}

func (x VoucherType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VoucherType) Descriptor() protoreflect.EnumDescriptor {
	return file_basicService_v1_account_proto_enumTypes[0].Descriptor()
}

func (VoucherType) Type() protoreflect.EnumType {
	return &file_basicService_v1_account_proto_enumTypes[0]
}

func (x VoucherType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VoucherType.Descriptor instead.
func (VoucherType) EnumDescriptor() ([]byte, []int) {
	return file_basicService_v1_account_proto_rawDescGZIP(), []int{0}
}

// 手机号和邮箱至少要存在一个
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile   string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_basicService_v1_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basicService_v1_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_basicService_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	AccountId int64     `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"` // 账户id，通过此id绑定跨平台账号
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_basicService_v1_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basicService_v1_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_basicService_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RegisterResponse) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

// 手机号、邮箱、账户id三者至少要存在一个
type CheckAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile    string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	AccountId int64  `protobuf:"varint,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CheckAccountRequest) Reset() {
	*x = CheckAccountRequest{}
	mi := &file_basicService_v1_account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccountRequest) ProtoMessage() {}

func (x *CheckAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basicService_v1_account_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccountRequest.ProtoReflect.Descriptor instead.
func (*CheckAccountRequest) Descriptor() ([]byte, []int) {
	return file_basicService_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *CheckAccountRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *CheckAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CheckAccountRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *CheckAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CheckAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	AccountId int64     `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"` // 账户id，通过此id绑定跨平台账号
}

func (x *CheckAccountResponse) Reset() {
	*x = CheckAccountResponse{}
	mi := &file_basicService_v1_account_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccountResponse) ProtoMessage() {}

func (x *CheckAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basicService_v1_account_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccountResponse.ProtoReflect.Descriptor instead.
func (*CheckAccountResponse) Descriptor() ([]byte, []int) {
	return file_basicService_v1_account_proto_rawDescGZIP(), []int{3}
}

func (x *CheckAccountResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *CheckAccountResponse) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

type BindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId   int64       `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	VoucherType VoucherType `protobuf:"varint,2,opt,name=voucher_type,json=voucherType,proto3,enum=basicService.v1.VoucherType" json:"voucher_type,omitempty"`
	Voucher     string      `protobuf:"bytes,3,opt,name=voucher,proto3" json:"voucher,omitempty"`
}

func (x *BindRequest) Reset() {
	*x = BindRequest{}
	mi := &file_basicService_v1_account_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindRequest) ProtoMessage() {}

func (x *BindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basicService_v1_account_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindRequest.ProtoReflect.Descriptor instead.
func (*BindRequest) Descriptor() ([]byte, []int) {
	return file_basicService_v1_account_proto_rawDescGZIP(), []int{4}
}

func (x *BindRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *BindRequest) GetVoucherType() VoucherType {
	if x != nil {
		return x.VoucherType
	}
	return VoucherType_VOUCHER_PHONE
}

func (x *BindRequest) GetVoucher() string {
	if x != nil {
		return x.Voucher
	}
	return ""
}

type BindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *BindResponse) Reset() {
	*x = BindResponse{}
	mi := &file_basicService_v1_account_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindResponse) ProtoMessage() {}

func (x *BindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basicService_v1_account_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindResponse.ProtoReflect.Descriptor instead.
func (*BindResponse) Descriptor() ([]byte, []int) {
	return file_basicService_v1_account_proto_rawDescGZIP(), []int{5}
}

func (x *BindResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

type UnbindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId   int64       `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	VoucherType VoucherType `protobuf:"varint,2,opt,name=voucher_type,json=voucherType,proto3,enum=basicService.v1.VoucherType" json:"voucher_type,omitempty"`
	Voucher     string      `protobuf:"bytes,3,opt,name=voucher,proto3" json:"voucher,omitempty"`
}

func (x *UnbindRequest) Reset() {
	*x = UnbindRequest{}
	mi := &file_basicService_v1_account_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnbindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbindRequest) ProtoMessage() {}

func (x *UnbindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basicService_v1_account_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnbindRequest.ProtoReflect.Descriptor instead.
func (*UnbindRequest) Descriptor() ([]byte, []int) {
	return file_basicService_v1_account_proto_rawDescGZIP(), []int{6}
}

func (x *UnbindRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *UnbindRequest) GetVoucherType() VoucherType {
	if x != nil {
		return x.VoucherType
	}
	return VoucherType_VOUCHER_PHONE
}

func (x *UnbindRequest) GetVoucher() string {
	if x != nil {
		return x.Voucher
	}
	return ""
}

type UnbindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *UnbindResponse) Reset() {
	*x = UnbindResponse{}
	mi := &file_basicService_v1_account_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnbindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbindResponse) ProtoMessage() {}

func (x *UnbindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basicService_v1_account_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnbindResponse.ProtoReflect.Descriptor instead.
func (*UnbindResponse) Descriptor() ([]byte, []int) {
	return file_basicService_v1_account_proto_rawDescGZIP(), []int{7}
}

func (x *UnbindResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_basicService_v1_account_proto protoreflect.FileDescriptor

var file_basicService_v1_account_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b,
	0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x60, 0x0a, 0x10, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x7e, 0x0a,
	0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x64, 0x0a,
	0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0c, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x75, 0x63, 0x68,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x22, 0x3d, 0x0a,
	0x0c, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x89, 0x01, 0x0a,
	0x0d, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3f, 0x0a,
	0x0c, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x0e, 0x55, 0x6e, 0x62, 0x69,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x2a, 0x33, 0x0a, 0x0b, 0x56, 0x6f, 0x75,
	0x63, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x4f, 0x55, 0x43,
	0x48, 0x45, 0x52, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x56,
	0x4f, 0x55, 0x43, 0x48, 0x45, 0x52, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x32, 0xce,
	0x02, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4f, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x20, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x24, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x43, 0x0a, 0x04, 0x42, 0x69, 0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x12, 0x1e,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x58,
	0x4a, 0x30, 0x30, 0x30, 0x30, 0x2f, 0x74, 0x6f, 0x6b, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basicService_v1_account_proto_rawDescOnce sync.Once
	file_basicService_v1_account_proto_rawDescData = file_basicService_v1_account_proto_rawDesc
)

func file_basicService_v1_account_proto_rawDescGZIP() []byte {
	file_basicService_v1_account_proto_rawDescOnce.Do(func() {
		file_basicService_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_basicService_v1_account_proto_rawDescData)
	})
	return file_basicService_v1_account_proto_rawDescData
}

var file_basicService_v1_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_basicService_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_basicService_v1_account_proto_goTypes = []any{
	(VoucherType)(0),             // 0: basicService.v1.VoucherType
	(*RegisterRequest)(nil),      // 1: basicService.v1.RegisterRequest
	(*RegisterResponse)(nil),     // 2: basicService.v1.RegisterResponse
	(*CheckAccountRequest)(nil),  // 3: basicService.v1.CheckAccountRequest
	(*CheckAccountResponse)(nil), // 4: basicService.v1.CheckAccountResponse
	(*BindRequest)(nil),          // 5: basicService.v1.BindRequest
	(*BindResponse)(nil),         // 6: basicService.v1.BindResponse
	(*UnbindRequest)(nil),        // 7: basicService.v1.UnbindRequest
	(*UnbindResponse)(nil),       // 8: basicService.v1.UnbindResponse
	(*Metadata)(nil),             // 9: basicService.v1.Metadata
}
var file_basicService_v1_account_proto_depIdxs = []int32{
	9,  // 0: basicService.v1.RegisterResponse.meta:type_name -> basicService.v1.Metadata
	9,  // 1: basicService.v1.CheckAccountResponse.meta:type_name -> basicService.v1.Metadata
	0,  // 2: basicService.v1.BindRequest.voucher_type:type_name -> basicService.v1.VoucherType
	9,  // 3: basicService.v1.BindResponse.meta:type_name -> basicService.v1.Metadata
	0,  // 4: basicService.v1.UnbindRequest.voucher_type:type_name -> basicService.v1.VoucherType
	9,  // 5: basicService.v1.UnbindResponse.meta:type_name -> basicService.v1.Metadata
	1,  // 6: basicService.v1.AccountService.Register:input_type -> basicService.v1.RegisterRequest
	3,  // 7: basicService.v1.AccountService.CheckAccount:input_type -> basicService.v1.CheckAccountRequest
	5,  // 8: basicService.v1.AccountService.Bind:input_type -> basicService.v1.BindRequest
	7,  // 9: basicService.v1.AccountService.Unbind:input_type -> basicService.v1.UnbindRequest
	2,  // 10: basicService.v1.AccountService.Register:output_type -> basicService.v1.RegisterResponse
	4,  // 11: basicService.v1.AccountService.CheckAccount:output_type -> basicService.v1.CheckAccountResponse
	6,  // 12: basicService.v1.AccountService.Bind:output_type -> basicService.v1.BindResponse
	8,  // 13: basicService.v1.AccountService.Unbind:output_type -> basicService.v1.UnbindResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_basicService_v1_account_proto_init() }
func file_basicService_v1_account_proto_init() {
	if File_basicService_v1_account_proto != nil {
		return
	}
	file_basicService_v1_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_basicService_v1_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basicService_v1_account_proto_goTypes,
		DependencyIndexes: file_basicService_v1_account_proto_depIdxs,
		EnumInfos:         file_basicService_v1_account_proto_enumTypes,
		MessageInfos:      file_basicService_v1_account_proto_msgTypes,
	}.Build()
	File_basicService_v1_account_proto = out.File
	file_basicService_v1_account_proto_rawDesc = nil
	file_basicService_v1_account_proto_goTypes = nil
	file_basicService_v1_account_proto_depIdxs = nil
}
