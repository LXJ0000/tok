// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: gateway/v1/comment.proto

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

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id,omitempty,string"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 评论id
	// @gotags: json:"videoId,omitempty,string"
	VideoId int64 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` // 视频id
	// @gotags: json:"parentId,omitempty,string"
	ParentId   int64        `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`      // 父评论id
	User       *CommentUser `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`                               // 评论用户
	ReplyUser  *CommentUser `protobuf:"bytes,5,opt,name=reply_user,json=replyUser,proto3" json:"reply_user,omitempty"`    // 回复用户
	Content    string       `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`                         // 评论内容
	Date       string       `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`                               // 评论日期
	LikeCount  string       `protobuf:"bytes,8,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`    // 点赞数
	ReplyCount string       `protobuf:"bytes,9,opt,name=reply_count,json=replyCount,proto3" json:"reply_count,omitempty"` // 回复数
	Comments   []*Comment   `protobuf:"bytes,10,rep,name=comments,proto3" json:"comments,omitempty"`                      // 子评论
}

func (x *Comment) Reset() {
	*x = Comment{}
	mi := &file_gateway_v1_comment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_comment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_gateway_v1_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *Comment) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Comment) GetUser() *CommentUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Comment) GetReplyUser() *CommentUser {
	if x != nil {
		return x.ReplyUser
	}
	return nil
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Comment) GetLikeCount() string {
	if x != nil {
		return x.LikeCount
	}
	return ""
}

func (x *Comment) GetReplyCount() string {
	if x != nil {
		return x.ReplyCount
	}
	return ""
}

func (x *Comment) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type CommentUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id,omitempty,string"
	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                      // 用户id
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                   // 用户名称
	Avatar      string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`                               // 用户头像
	IsFollowing bool   `protobuf:"varint,4,opt,name=is_following,json=isFollowing,proto3" json:"is_following,omitempty"` // 是否关注
}

func (x *CommentUser) Reset() {
	*x = CommentUser{}
	mi := &file_gateway_v1_comment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommentUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentUser) ProtoMessage() {}

func (x *CommentUser) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_comment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentUser.ProtoReflect.Descriptor instead.
func (*CommentUser) Descriptor() ([]byte, []int) {
	return file_gateway_v1_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentUser) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommentUser) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *CommentUser) GetIsFollowing() bool {
	if x != nil {
		return x.IsFollowing
	}
	return false
}

type CreateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"videoId,omitempty,string"
	VideoId int64  `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` // 视频id
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`                 // 评论内容
	// @gotags: json:"parentId,omitempty,string"
	ParentId int64 `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// @gotags: json:"replyUserId,omitempty,string"
	ReplyUserId int64 `protobuf:"varint,4,opt,name=reply_user_id,json=replyUserId,proto3" json:"reply_user_id,omitempty"`
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	mi := &file_gateway_v1_comment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_comment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_gateway_v1_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCommentRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *CreateCommentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateCommentRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CreateCommentRequest) GetReplyUserId() int64 {
	if x != nil {
		return x.ReplyUserId
	}
	return 0
}

type CreateCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CreateCommentResponse) Reset() {
	*x = CreateCommentResponse{}
	mi := &file_gateway_v1_comment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentResponse) ProtoMessage() {}

func (x *CreateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_comment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentResponse.ProtoReflect.Descriptor instead.
func (*CreateCommentResponse) Descriptor() ([]byte, []int) {
	return file_gateway_v1_comment_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCommentResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type RemoveCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id,omitempty,string"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 评论id
}

func (x *RemoveCommentRequest) Reset() {
	*x = RemoveCommentRequest{}
	mi := &file_gateway_v1_comment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCommentRequest) ProtoMessage() {}

func (x *RemoveCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_comment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCommentRequest.ProtoReflect.Descriptor instead.
func (*RemoveCommentRequest) Descriptor() ([]byte, []int) {
	return file_gateway_v1_comment_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveCommentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveCommentResponse) Reset() {
	*x = RemoveCommentResponse{}
	mi := &file_gateway_v1_comment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCommentResponse) ProtoMessage() {}

func (x *RemoveCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_comment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCommentResponse.ProtoReflect.Descriptor instead.
func (*RemoveCommentResponse) Descriptor() ([]byte, []int) {
	return file_gateway_v1_comment_proto_rawDescGZIP(), []int{5}
}

type ListComment4VideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"videoId,omitempty,string"
	VideoId    int64              `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` // 视频id
	Pagination *PaginationRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListComment4VideoRequest) Reset() {
	*x = ListComment4VideoRequest{}
	mi := &file_gateway_v1_comment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListComment4VideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComment4VideoRequest) ProtoMessage() {}

func (x *ListComment4VideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_comment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComment4VideoRequest.ProtoReflect.Descriptor instead.
func (*ListComment4VideoRequest) Descriptor() ([]byte, []int) {
	return file_gateway_v1_comment_proto_rawDescGZIP(), []int{6}
}

func (x *ListComment4VideoRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *ListComment4VideoRequest) GetPagination() *PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListComment4VideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments   []*Comment          `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Pagination *PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListComment4VideoResponse) Reset() {
	*x = ListComment4VideoResponse{}
	mi := &file_gateway_v1_comment_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListComment4VideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComment4VideoResponse) ProtoMessage() {}

func (x *ListComment4VideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_comment_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComment4VideoResponse.ProtoReflect.Descriptor instead.
func (*ListComment4VideoResponse) Descriptor() ([]byte, []int) {
	return file_gateway_v1_comment_proto_rawDescGZIP(), []int{7}
}

func (x *ListComment4VideoResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *ListComment4VideoResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListChildCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"commentId,omitempty,string"
	CommentId  int64              `protobuf:"varint,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"` // 评论id
	Pagination *PaginationRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListChildCommentRequest) Reset() {
	*x = ListChildCommentRequest{}
	mi := &file_gateway_v1_comment_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListChildCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChildCommentRequest) ProtoMessage() {}

func (x *ListChildCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_comment_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChildCommentRequest.ProtoReflect.Descriptor instead.
func (*ListChildCommentRequest) Descriptor() ([]byte, []int) {
	return file_gateway_v1_comment_proto_rawDescGZIP(), []int{8}
}

func (x *ListChildCommentRequest) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *ListChildCommentRequest) GetPagination() *PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListChildCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments   []*Comment          `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Pagination *PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListChildCommentResponse) Reset() {
	*x = ListChildCommentResponse{}
	mi := &file_gateway_v1_comment_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListChildCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChildCommentResponse) ProtoMessage() {}

func (x *ListChildCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_comment_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChildCommentResponse.ProtoReflect.Descriptor instead.
func (*ListChildCommentResponse) Descriptor() ([]byte, []int) {
	return file_gateway_v1_comment_proto_rawDescGZIP(), []int{9}
}

func (x *ListChildCommentResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *ListChildCommentResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_gateway_v1_comment_proto protoreflect.FileDescriptor

var file_gateway_v1_comment_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x02,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x36,
	0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x6c, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x22, 0x8c, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x46, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x74, 0x0a, 0x18, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x34, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x34, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x3e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x77, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8b, 0x01, 0x0a, 0x18, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xda, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x66, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x2a, 0x08, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x7b, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x34, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x12, 0x24, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x34, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x34, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x78, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4c, 0x58, 0x4a, 0x30, 0x30, 0x30, 0x30, 0x2f, 0x74, 0x6f, 0x6b, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_v1_comment_proto_rawDescOnce sync.Once
	file_gateway_v1_comment_proto_rawDescData = file_gateway_v1_comment_proto_rawDesc
)

func file_gateway_v1_comment_proto_rawDescGZIP() []byte {
	file_gateway_v1_comment_proto_rawDescOnce.Do(func() {
		file_gateway_v1_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_v1_comment_proto_rawDescData)
	})
	return file_gateway_v1_comment_proto_rawDescData
}

var file_gateway_v1_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_gateway_v1_comment_proto_goTypes = []any{
	(*Comment)(nil),                   // 0: gateway.v1.Comment
	(*CommentUser)(nil),               // 1: gateway.v1.CommentUser
	(*CreateCommentRequest)(nil),      // 2: gateway.v1.CreateCommentRequest
	(*CreateCommentResponse)(nil),     // 3: gateway.v1.CreateCommentResponse
	(*RemoveCommentRequest)(nil),      // 4: gateway.v1.RemoveCommentRequest
	(*RemoveCommentResponse)(nil),     // 5: gateway.v1.RemoveCommentResponse
	(*ListComment4VideoRequest)(nil),  // 6: gateway.v1.ListComment4VideoRequest
	(*ListComment4VideoResponse)(nil), // 7: gateway.v1.ListComment4VideoResponse
	(*ListChildCommentRequest)(nil),   // 8: gateway.v1.ListChildCommentRequest
	(*ListChildCommentResponse)(nil),  // 9: gateway.v1.ListChildCommentResponse
	(*PaginationRequest)(nil),         // 10: gateway.v1.PaginationRequest
	(*PaginationResponse)(nil),        // 11: gateway.v1.PaginationResponse
}
var file_gateway_v1_comment_proto_depIdxs = []int32{
	1,  // 0: gateway.v1.Comment.user:type_name -> gateway.v1.CommentUser
	1,  // 1: gateway.v1.Comment.reply_user:type_name -> gateway.v1.CommentUser
	0,  // 2: gateway.v1.Comment.comments:type_name -> gateway.v1.Comment
	0,  // 3: gateway.v1.CreateCommentResponse.comment:type_name -> gateway.v1.Comment
	10, // 4: gateway.v1.ListComment4VideoRequest.pagination:type_name -> gateway.v1.PaginationRequest
	0,  // 5: gateway.v1.ListComment4VideoResponse.comments:type_name -> gateway.v1.Comment
	11, // 6: gateway.v1.ListComment4VideoResponse.pagination:type_name -> gateway.v1.PaginationResponse
	10, // 7: gateway.v1.ListChildCommentRequest.pagination:type_name -> gateway.v1.PaginationRequest
	0,  // 8: gateway.v1.ListChildCommentResponse.comments:type_name -> gateway.v1.Comment
	11, // 9: gateway.v1.ListChildCommentResponse.pagination:type_name -> gateway.v1.PaginationResponse
	2,  // 10: gateway.v1.CommentService.CreateComment:input_type -> gateway.v1.CreateCommentRequest
	4,  // 11: gateway.v1.CommentService.RemoveComment:input_type -> gateway.v1.RemoveCommentRequest
	6,  // 12: gateway.v1.CommentService.ListComment4Video:input_type -> gateway.v1.ListComment4VideoRequest
	8,  // 13: gateway.v1.CommentService.ListChildComment:input_type -> gateway.v1.ListChildCommentRequest
	3,  // 14: gateway.v1.CommentService.CreateComment:output_type -> gateway.v1.CreateCommentResponse
	5,  // 15: gateway.v1.CommentService.RemoveComment:output_type -> gateway.v1.RemoveCommentResponse
	7,  // 16: gateway.v1.CommentService.ListComment4Video:output_type -> gateway.v1.ListComment4VideoResponse
	9,  // 17: gateway.v1.CommentService.ListChildComment:output_type -> gateway.v1.ListChildCommentResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_gateway_v1_comment_proto_init() }
func file_gateway_v1_comment_proto_init() {
	if File_gateway_v1_comment_proto != nil {
		return
	}
	file_gateway_v1_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_v1_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_v1_comment_proto_goTypes,
		DependencyIndexes: file_gateway_v1_comment_proto_depIdxs,
		MessageInfos:      file_gateway_v1_comment_proto_msgTypes,
	}.Build()
	File_gateway_v1_comment_proto = out.File
	file_gateway_v1_comment_proto_rawDesc = nil
	file_gateway_v1_comment_proto_goTypes = nil
	file_gateway_v1_comment_proto_depIdxs = nil
}