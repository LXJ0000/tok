// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: gateway/v1/common.proto

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

type SortOrder int32

const (
	SortOrder_ASC  SortOrder = 0
	SortOrder_DESC SortOrder = 1
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	SortOrder_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_gateway_v1_common_proto_enumTypes[0].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_gateway_v1_common_proto_enumTypes[0]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_gateway_v1_common_proto_rawDescGZIP(), []int{0}
}

type SearchOperator int32

const (
	SearchOperator_EQ       SearchOperator = 0  // 等于
	SearchOperator_NE       SearchOperator = 1  // 不等于
	SearchOperator_GT       SearchOperator = 2  // 大于
	SearchOperator_GE       SearchOperator = 3  // 大于等于
	SearchOperator_LT       SearchOperator = 4  // 小于
	SearchOperator_LE       SearchOperator = 5  // 小于等于
	SearchOperator_LIKE     SearchOperator = 6  // 使用like的模糊匹配
	SearchOperator_WILDCARD SearchOperator = 7  // 使用通配符的模糊匹配
	SearchOperator_IN       SearchOperator = 8  // 在指定的集合中
	SearchOperator_NOT_IN   SearchOperator = 9  // 不在指定的集合中
	SearchOperator_BETWEEN  SearchOperator = 10 // 在指定的范围内
	SearchOperator_RE       SearchOperator = 11 // 正则匹配
)

// Enum value maps for SearchOperator.
var (
	SearchOperator_name = map[int32]string{
		0:  "EQ",
		1:  "NE",
		2:  "GT",
		3:  "GE",
		4:  "LT",
		5:  "LE",
		6:  "LIKE",
		7:  "WILDCARD",
		8:  "IN",
		9:  "NOT_IN",
		10: "BETWEEN",
		11: "RE",
	}
	SearchOperator_value = map[string]int32{
		"EQ":       0,
		"NE":       1,
		"GT":       2,
		"GE":       3,
		"LT":       4,
		"LE":       5,
		"LIKE":     6,
		"WILDCARD": 7,
		"IN":       8,
		"NOT_IN":   9,
		"BETWEEN":  10,
		"RE":       11,
	}
)

func (x SearchOperator) Enum() *SearchOperator {
	p := new(SearchOperator)
	*p = x
	return p
}

func (x SearchOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_gateway_v1_common_proto_enumTypes[1].Descriptor()
}

func (SearchOperator) Type() protoreflect.EnumType {
	return &file_gateway_v1_common_proto_enumTypes[1]
}

func (x SearchOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchOperator.Descriptor instead.
func (SearchOperator) EnumDescriptor() ([]byte, []int) {
	return file_gateway_v1_common_proto_rawDescGZIP(), []int{1}
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizCode int32    `protobuf:"varint,1,opt,name=biz_code,json=bizCode,proto3" json:"biz_code,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Domain  string   `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Reason  []string `protobuf:"bytes,4,rep,name=reason,proto3" json:"reason,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	mi := &file_gateway_v1_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_gateway_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetBizCode() int32 {
	if x != nil {
		return x.BizCode
	}
	return 0
}

func (x *Metadata) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Metadata) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Metadata) GetReason() []string {
	if x != nil {
		return x.Reason
	}
	return nil
}

type SortField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string    `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`                            // 用于排序的字段名称
	Order SortOrder `protobuf:"varint,2,opt,name=order,proto3,enum=gateway.v1.SortOrder" json:"order,omitempty"` // 排序方式
}

func (x *SortField) Reset() {
	*x = SortField{}
	mi := &file_gateway_v1_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SortField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortField) ProtoMessage() {}

func (x *SortField) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortField.ProtoReflect.Descriptor instead.
func (*SortField) Descriptor() ([]byte, []int) {
	return file_gateway_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *SortField) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SortField) GetOrder() SortOrder {
	if x != nil {
		return x.Order
	}
	return SortOrder_ASC
}

type PaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32        `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"` // 页码 [1, +∞)
	Size int32        `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"` // 页面大小
	Sort []*SortField `protobuf:"bytes,3,rep,name=sort,proto3" json:"sort,omitempty"`  // 根据字段进行排序
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	mi := &file_gateway_v1_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_gateway_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *PaginationRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PaginationRequest) GetSort() []*SortField {
	if x != nil {
		return x.Sort
	}
	return nil
}

type PaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`   // 当前数据的所属页码
	Total int32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"` // 总页数
	Count int32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"` // 总条目数
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	mi := &file_gateway_v1_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_gateway_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *PaginationResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PaginationResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type SearchField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field     string         `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`                                       // 用于搜索的字段名称
	Value     string         `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`                                       // 搜索的值
	ValueList []string       `protobuf:"bytes,3,rep,name=value_list,json=valueList,proto3" json:"value_list,omitempty"`              // 搜索的值列表
	Operator  SearchOperator `protobuf:"varint,4,opt,name=operator,proto3,enum=gateway.v1.SearchOperator" json:"operator,omitempty"` // 操作符
}

func (x *SearchField) Reset() {
	*x = SearchField{}
	mi := &file_gateway_v1_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchField) ProtoMessage() {}

func (x *SearchField) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchField.ProtoReflect.Descriptor instead.
func (*SearchField) Descriptor() ([]byte, []int) {
	return file_gateway_v1_common_proto_rawDescGZIP(), []int{4}
}

func (x *SearchField) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SearchField) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SearchField) GetValueList() []string {
	if x != nil {
		return x.ValueList
	}
	return nil
}

func (x *SearchField) GetOperator() SearchOperator {
	if x != nil {
		return x.Operator
	}
	return SearchOperator_EQ
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search []*SearchField `protobuf:"bytes,1,rep,name=search,proto3" json:"search,omitempty"` // 搜索条件
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	mi := &file_gateway_v1_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_common_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_gateway_v1_common_proto_rawDescGZIP(), []int{5}
}

func (x *SearchRequest) GetSearch() []*SearchField {
	if x != nil {
		return x.Search
	}
	return nil
}

type VideoAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id,omitempty,string"
	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar      string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	IsFollowing bool   `protobuf:"varint,4,opt,name=isFollowing,proto3" json:"isFollowing,omitempty"`
}

func (x *VideoAuthor) Reset() {
	*x = VideoAuthor{}
	mi := &file_gateway_v1_common_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoAuthor) ProtoMessage() {}

func (x *VideoAuthor) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_common_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoAuthor.ProtoReflect.Descriptor instead.
func (*VideoAuthor) Descriptor() ([]byte, []int) {
	return file_gateway_v1_common_proto_rawDescGZIP(), []int{6}
}

func (x *VideoAuthor) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VideoAuthor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VideoAuthor) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *VideoAuthor) GetIsFollowing() bool {
	if x != nil {
		return x.IsFollowing
	}
	return false
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id,omitempty,string"
	Id       int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                            // 视频唯一标识
	Author   *VideoAuthor `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`                     // 视频作者信息
	PlayUrl  string       `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`    // 视频播放地址
	CoverUrl string       `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"` // 视频封面地址
	// @gotags: json:"favoriteCount,omitempty,string"
	FavoriteCount int64 `protobuf:"varint,5,opt,name=favoriteCount,proto3" json:"favoriteCount,omitempty"` // 视频的点赞总数
	// @gotags: json:"commentCount,omitempty,string"
	CommentCount int64  `protobuf:"varint,6,opt,name=commentCount,proto3" json:"commentCount,omitempty"` // 视频的评论总数
	IsFavorite   bool   `protobuf:"varint,7,opt,name=isFavorite,proto3" json:"isFavorite,omitempty"`     // true-已点赞，false-未点赞
	Title        string `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`                // 视频标题
	IsCollected  bool   `protobuf:"varint,9,opt,name=isCollected,proto3" json:"isCollected,omitempty"`   // 是否已收藏
	// @gotags: json:"collectedCount,omitempty,string"
	CollectedCount int64 `protobuf:"varint,10,opt,name=collectedCount,proto3" json:"collectedCount,omitempty"` // 收藏数
}

func (x *Video) Reset() {
	*x = Video{}
	mi := &file_gateway_v1_common_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_common_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_gateway_v1_common_proto_rawDescGZIP(), []int{7}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *VideoAuthor {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Video) GetIsCollected() bool {
	if x != nil {
		return x.IsCollected
	}
	return false
}

func (x *Video) GetCollectedCount() int64 {
	if x != nil {
		return x.CollectedCount
	}
	return 0
}

var File_gateway_v1_common_proto protoreflect.FileDescriptor

var file_gateway_v1_common_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x6f, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x7a, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x69, 0x7a, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2b, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x66, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x6f, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x54,
	0x0a, 0x12, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x36, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x40, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x6b, 0x0a, 0x0b, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x22, 0xca, 0x02, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2f, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x2a, 0x1e, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53,
	0x43, 0x10, 0x01, 0x2a, 0x81, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x51, 0x10, 0x00, 0x12, 0x06,
	0x0a, 0x02, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x54, 0x10, 0x02, 0x12, 0x06,
	0x0a, 0x02, 0x47, 0x45, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x54, 0x10, 0x04, 0x12, 0x06,
	0x0a, 0x02, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x06,
	0x12, 0x0c, 0x0a, 0x08, 0x57, 0x49, 0x4c, 0x44, 0x43, 0x41, 0x52, 0x44, 0x10, 0x07, 0x12, 0x06,
	0x0a, 0x02, 0x49, 0x4e, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e,
	0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x45, 0x54, 0x57, 0x45, 0x45, 0x4e, 0x10, 0x0a, 0x12,
	0x06, 0x0a, 0x02, 0x52, 0x45, 0x10, 0x0b, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x58, 0x4a, 0x30, 0x30, 0x30, 0x30, 0x2f, 0x74, 0x6f,
	0x6b, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_v1_common_proto_rawDescOnce sync.Once
	file_gateway_v1_common_proto_rawDescData = file_gateway_v1_common_proto_rawDesc
)

func file_gateway_v1_common_proto_rawDescGZIP() []byte {
	file_gateway_v1_common_proto_rawDescOnce.Do(func() {
		file_gateway_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_v1_common_proto_rawDescData)
	})
	return file_gateway_v1_common_proto_rawDescData
}

var file_gateway_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_gateway_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_gateway_v1_common_proto_goTypes = []any{
	(SortOrder)(0),             // 0: gateway.v1.SortOrder
	(SearchOperator)(0),        // 1: gateway.v1.SearchOperator
	(*Metadata)(nil),           // 2: gateway.v1.Metadata
	(*SortField)(nil),          // 3: gateway.v1.SortField
	(*PaginationRequest)(nil),  // 4: gateway.v1.PaginationRequest
	(*PaginationResponse)(nil), // 5: gateway.v1.PaginationResponse
	(*SearchField)(nil),        // 6: gateway.v1.SearchField
	(*SearchRequest)(nil),      // 7: gateway.v1.SearchRequest
	(*VideoAuthor)(nil),        // 8: gateway.v1.VideoAuthor
	(*Video)(nil),              // 9: gateway.v1.Video
}
var file_gateway_v1_common_proto_depIdxs = []int32{
	0, // 0: gateway.v1.SortField.order:type_name -> gateway.v1.SortOrder
	3, // 1: gateway.v1.PaginationRequest.sort:type_name -> gateway.v1.SortField
	1, // 2: gateway.v1.SearchField.operator:type_name -> gateway.v1.SearchOperator
	6, // 3: gateway.v1.SearchRequest.search:type_name -> gateway.v1.SearchField
	8, // 4: gateway.v1.Video.author:type_name -> gateway.v1.VideoAuthor
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_gateway_v1_common_proto_init() }
func file_gateway_v1_common_proto_init() {
	if File_gateway_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_v1_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_v1_common_proto_goTypes,
		DependencyIndexes: file_gateway_v1_common_proto_depIdxs,
		EnumInfos:         file_gateway_v1_common_proto_enumTypes,
		MessageInfos:      file_gateway_v1_common_proto_msgTypes,
	}.Build()
	File_gateway_v1_common_proto = out.File
	file_gateway_v1_common_proto_rawDesc = nil
	file_gateway_v1_common_proto_goTypes = nil
	file_gateway_v1_common_proto_depIdxs = nil
}
