// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: coreService/v1/user.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type UpdateUserImageType int32

const (
	UpdateUserImageType_AVATAR           UpdateUserImageType = 0
	UpdateUserImageType_BACKGROUND_IMAGE UpdateUserImageType = 1
)

// Enum value maps for UpdateUserImageType.
var (
	UpdateUserImageType_name = map[int32]string{
		0: "AVATAR",
		1: "BACKGROUND_IMAGE",
	}
	UpdateUserImageType_value = map[string]int32{
		"AVATAR":           0,
		"BACKGROUND_IMAGE": 1,
	}
)

func (x UpdateUserImageType) Enum() *UpdateUserImageType {
	p := new(UpdateUserImageType)
	*p = x
	return p
}

func (x UpdateUserImageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateUserImageType) Descriptor() protoreflect.EnumDescriptor {
	return file_coreService_v1_user_proto_enumTypes[0].Descriptor()
}

func (UpdateUserImageType) Type() protoreflect.EnumType {
	return &file_coreService_v1_user_proto_enumTypes[0]
}

func (x UpdateUserImageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateUserImageType.Descriptor instead.
func (UpdateUserImageType) EnumDescriptor() ([]byte, []int) {
	return file_coreService_v1_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                 // 用户id
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                              // 用户名称
	Avatar          string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`                                          // 用户头像Url
	BackgroundImage string `protobuf:"bytes,4,opt,name=background_image,json=backgroundImage,proto3" json:"background_image,omitempty"` // 用户个人页顶部大图
	Signature       string `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`                                    // 个人简介
	Mobile          string `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile,omitempty"`                                          // 手机号
	Email           string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`                                            // 邮箱
	FollowCount     int64  `protobuf:"varint,8,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`            // 关注总数
	FollowerCount   int64  `protobuf:"varint,9,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`      // 粉丝总数
	TotalFavorited  int64  `protobuf:"varint,10,opt,name=total_favorited,json=totalFavorited,proto3" json:"total_favorited,omitempty"`  // 获赞数量
	WorkCount       int64  `protobuf:"varint,11,opt,name=work_count,json=workCount,proto3" json:"work_count,omitempty"`                 // 作品数量
	FavoriteCount   int64  `protobuf:"varint,12,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`     // 点赞数量
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_coreService_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_coreService_v1_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_coreService_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetBackgroundImage() string {
	if x != nil {
		return x.BackgroundImage
	}
	return ""
}

func (x *User) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *User) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetTotalFavorited() int64 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

func (x *User) GetWorkCount() int64 {
	if x != nil {
		return x.WorkCount
	}
	return 0
}

func (x *User) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile    string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	AccountId int64  `protobuf:"varint,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_coreService_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coreService_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_coreService_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	UserId int64     `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_coreService_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coreService_v1_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_coreService_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *CreateUserResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AccountId int64 `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *GetUserInfoRequest) Reset() {
	*x = GetUserInfoRequest{}
	mi := &file_coreService_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRequest) ProtoMessage() {}

func (x *GetUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coreService_v1_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRequest.ProtoReflect.Descriptor instead.
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_coreService_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserInfoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserInfoRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

type GetUserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	User *User     `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserInfoResponse) Reset() {
	*x = GetUserInfoResponse{}
	mi := &file_coreService_v1_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResponse) ProtoMessage() {}

func (x *GetUserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coreService_v1_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResponse.ProtoReflect.Descriptor instead.
func (*GetUserInfoResponse) Descriptor() ([]byte, []int) {
	return file_coreService_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserInfoResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *GetUserInfoResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar          string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	BackgroundImage string `protobuf:"bytes,4,opt,name=background_image,json=backgroundImage,proto3" json:"background_image,omitempty"`
	Signature       string `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *UpdateUserInfoRequest) Reset() {
	*x = UpdateUserInfoRequest{}
	mi := &file_coreService_v1_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoRequest) ProtoMessage() {}

func (x *UpdateUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coreService_v1_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_coreService_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserInfoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateUserInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserInfoRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UpdateUserInfoRequest) GetBackgroundImage() string {
	if x != nil {
		return x.BackgroundImage
	}
	return ""
}

func (x *UpdateUserInfoRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type UpdateUserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *UpdateUserInfoResponse) Reset() {
	*x = UpdateUserInfoResponse{}
	mi := &file_coreService_v1_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoResponse) ProtoMessage() {}

func (x *UpdateUserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coreService_v1_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoResponse) Descriptor() ([]byte, []int) {
	return file_coreService_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserInfoResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

type GetUserByIdListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIdList []int64 `protobuf:"varint,1,rep,packed,name=user_id_list,json=userIdList,proto3" json:"user_id_list,omitempty"`
}

func (x *GetUserByIdListRequest) Reset() {
	*x = GetUserByIdListRequest{}
	mi := &file_coreService_v1_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIdListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdListRequest) ProtoMessage() {}

func (x *GetUserByIdListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coreService_v1_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdListRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIdListRequest) Descriptor() ([]byte, []int) {
	return file_coreService_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserByIdListRequest) GetUserIdList() []int64 {
	if x != nil {
		return x.UserIdList
	}
	return nil
}

type GetUserByIdListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserList []*User   `protobuf:"bytes,1,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
	Meta     *Metadata `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *GetUserByIdListResponse) Reset() {
	*x = GetUserByIdListResponse{}
	mi := &file_coreService_v1_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIdListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdListResponse) ProtoMessage() {}

func (x *GetUserByIdListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coreService_v1_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdListResponse.ProtoReflect.Descriptor instead.
func (*GetUserByIdListResponse) Descriptor() ([]byte, []int) {
	return file_coreService_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserByIdListResponse) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

func (x *GetUserByIdListResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_coreService_v1_user_proto protoreflect.FileDescriptor

var file_coreService_v1_user_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6f, 0x72,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf2, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x69, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x5b, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x6d, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xae, 0x01, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x46, 0x0a, 0x16, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x22, 0x3a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x7a,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x2a, 0x37, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x42, 0x41, 0x43, 0x4b, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x49, 0x4d, 0x41, 0x47,
	0x45, 0x10, 0x01, 0x32, 0xc2, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x6d, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x74, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x1a, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x64, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x58, 0x4a, 0x30, 0x30, 0x30, 0x30, 0x2f, 0x74,
	0x6f, 0x6b, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coreService_v1_user_proto_rawDescOnce sync.Once
	file_coreService_v1_user_proto_rawDescData = file_coreService_v1_user_proto_rawDesc
)

func file_coreService_v1_user_proto_rawDescGZIP() []byte {
	file_coreService_v1_user_proto_rawDescOnce.Do(func() {
		file_coreService_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_coreService_v1_user_proto_rawDescData)
	})
	return file_coreService_v1_user_proto_rawDescData
}

var file_coreService_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_coreService_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_coreService_v1_user_proto_goTypes = []any{
	(UpdateUserImageType)(0),        // 0: coreService.v1.UpdateUserImageType
	(*User)(nil),                    // 1: coreService.v1.User
	(*CreateUserRequest)(nil),       // 2: coreService.v1.CreateUserRequest
	(*CreateUserResponse)(nil),      // 3: coreService.v1.CreateUserResponse
	(*GetUserInfoRequest)(nil),      // 4: coreService.v1.GetUserInfoRequest
	(*GetUserInfoResponse)(nil),     // 5: coreService.v1.GetUserInfoResponse
	(*UpdateUserInfoRequest)(nil),   // 6: coreService.v1.UpdateUserInfoRequest
	(*UpdateUserInfoResponse)(nil),  // 7: coreService.v1.UpdateUserInfoResponse
	(*GetUserByIdListRequest)(nil),  // 8: coreService.v1.GetUserByIdListRequest
	(*GetUserByIdListResponse)(nil), // 9: coreService.v1.GetUserByIdListResponse
	(*Metadata)(nil),                // 10: coreService.v1.Metadata
}
var file_coreService_v1_user_proto_depIdxs = []int32{
	10, // 0: coreService.v1.CreateUserResponse.meta:type_name -> coreService.v1.Metadata
	10, // 1: coreService.v1.GetUserInfoResponse.meta:type_name -> coreService.v1.Metadata
	1,  // 2: coreService.v1.GetUserInfoResponse.user:type_name -> coreService.v1.User
	10, // 3: coreService.v1.UpdateUserInfoResponse.meta:type_name -> coreService.v1.Metadata
	1,  // 4: coreService.v1.GetUserByIdListResponse.user_list:type_name -> coreService.v1.User
	10, // 5: coreService.v1.GetUserByIdListResponse.meta:type_name -> coreService.v1.Metadata
	2,  // 6: coreService.v1.UserService.CreateUser:input_type -> coreService.v1.CreateUserRequest
	4,  // 7: coreService.v1.UserService.GetUserInfo:input_type -> coreService.v1.GetUserInfoRequest
	6,  // 8: coreService.v1.UserService.UpdateUserInfo:input_type -> coreService.v1.UpdateUserInfoRequest
	8,  // 9: coreService.v1.UserService.GetUserByIdList:input_type -> coreService.v1.GetUserByIdListRequest
	3,  // 10: coreService.v1.UserService.CreateUser:output_type -> coreService.v1.CreateUserResponse
	5,  // 11: coreService.v1.UserService.GetUserInfo:output_type -> coreService.v1.GetUserInfoResponse
	7,  // 12: coreService.v1.UserService.UpdateUserInfo:output_type -> coreService.v1.UpdateUserInfoResponse
	9,  // 13: coreService.v1.UserService.GetUserByIdList:output_type -> coreService.v1.GetUserByIdListResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_coreService_v1_user_proto_init() }
func file_coreService_v1_user_proto_init() {
	if File_coreService_v1_user_proto != nil {
		return
	}
	file_coreService_v1_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coreService_v1_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coreService_v1_user_proto_goTypes,
		DependencyIndexes: file_coreService_v1_user_proto_depIdxs,
		EnumInfos:         file_coreService_v1_user_proto_enumTypes,
		MessageInfos:      file_coreService_v1_user_proto_msgTypes,
	}.Build()
	File_coreService_v1_user_proto = out.File
	file_coreService_v1_user_proto_rawDesc = nil
	file_coreService_v1_user_proto_goTypes = nil
	file_coreService_v1_user_proto_depIdxs = nil
}
