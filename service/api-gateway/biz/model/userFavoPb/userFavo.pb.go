// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: userFavo.proto

package userFavoPb

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

// ------------------------------------
// Messages
// ------------------------------------
type FavoActionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty" form:"userId" query:"userId"`     // 用户id
	Token   string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`          //用户鉴权token
	VideoId int64  `protobuf:"varint,3,opt,name=videoId,proto3" json:"videoId,omitempty" form:"videoId" query:"videoId"` // 视频id
	Type    int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty" form:"type" query:"type"`             // 1-点赞，2-取消点赞
}

func (x *FavoActionReq) Reset() {
	*x = FavoActionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userFavo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoActionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoActionReq) ProtoMessage() {}

func (x *FavoActionReq) ProtoReflect() protoreflect.Message {
	mi := &file_userFavo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoActionReq.ProtoReflect.Descriptor instead.
func (*FavoActionReq) Descriptor() ([]byte, []int) {
	return file_userFavo_proto_rawDescGZIP(), []int{0}
}

func (x *FavoActionReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FavoActionReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *FavoActionReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *FavoActionReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type FavoActionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty" form:"statusCode" query:"statusCode"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=statusMsg,proto3" json:"statusMsg,omitempty" form:"statusMsg" query:"statusMsg"`      // 返回状态描述
}

func (x *FavoActionResp) Reset() {
	*x = FavoActionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userFavo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoActionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoActionResp) ProtoMessage() {}

func (x *FavoActionResp) ProtoReflect() protoreflect.Message {
	mi := &file_userFavo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoActionResp.ProtoReflect.Descriptor instead.
func (*FavoActionResp) Descriptor() ([]byte, []int) {
	return file_userFavo_proto_rawDescGZIP(), []int{1}
}

func (x *FavoActionResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FavoActionResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type FavoListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty" form:"userId" query:"userId"` // 用户id
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`      //用户鉴权token
}

func (x *FavoListReq) Reset() {
	*x = FavoListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userFavo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoListReq) ProtoMessage() {}

func (x *FavoListReq) ProtoReflect() protoreflect.Message {
	mi := &file_userFavo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoListReq.ProtoReflect.Descriptor instead.
func (*FavoListReq) Descriptor() ([]byte, []int) {
	return file_userFavo_proto_rawDescGZIP(), []int{2}
}

func (x *FavoListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FavoListReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type FavoListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty" form:"statusCode" query:"statusCode"` // 状态码，0-成功，其他值-失败
	StatusMsg  string   `protobuf:"bytes,2,opt,name=statusMsg,proto3" json:"statusMsg,omitempty" form:"statusMsg" query:"statusMsg"`      // 返回状态描述
	VideoList  []*Video `protobuf:"bytes,3,rep,name=videoList,proto3" json:"videoList" form:"videoList" query:"videoList"`                // 用户点赞视频列表
}

func (x *FavoListResp) Reset() {
	*x = FavoListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userFavo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoListResp) ProtoMessage() {}

func (x *FavoListResp) ProtoReflect() protoreflect.Message {
	mi := &file_userFavo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoListResp.ProtoReflect.Descriptor instead.
func (*FavoListResp) Descriptor() ([]byte, []int) {
	return file_userFavo_proto_rawDescGZIP(), []int{3}
}

func (x *FavoListResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FavoListResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *FavoListResp) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type FavoStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty" form:"userId" query:"userId"`     // 用户id
	VideoId int64 `protobuf:"varint,2,opt,name=videoId,proto3" json:"videoId,omitempty" form:"videoId" query:"videoId"` // 视频id
}

func (x *FavoStatusReq) Reset() {
	*x = FavoStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userFavo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoStatusReq) ProtoMessage() {}

func (x *FavoStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_userFavo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoStatusReq.ProtoReflect.Descriptor instead.
func (*FavoStatusReq) Descriptor() ([]byte, []int) {
	return file_userFavo_proto_rawDescGZIP(), []int{4}
}

func (x *FavoStatusReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FavoStatusReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type FavoStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty" form:"statusCode" query:"statusCode"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=statusMsg,proto3" json:"statusMsg,omitempty" form:"statusMsg" query:"statusMsg"`      // 返回状态描述
	IsFavorite bool   `protobuf:"varint,3,opt,name=isFavorite,proto3" json:"isFavorite,omitempty" form:"isFavorite" query:"isFavorite"` // 用户是否点赞视频
}

func (x *FavoStatusResp) Reset() {
	*x = FavoStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userFavo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoStatusResp) ProtoMessage() {}

func (x *FavoStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_userFavo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoStatusResp.ProtoReflect.Descriptor instead.
func (*FavoStatusResp) Descriptor() ([]byte, []int) {
	return file_userFavo_proto_rawDescGZIP(), []int{5}
}

func (x *FavoStatusResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FavoStatusResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *FavoStatusResp) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`                                                                    // 视频唯一标识
	Author        *User  `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty" form:"author" query:"author"`                                                     // 视频作者信息
	PlayUrl       string `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty" form:"play_url" query:"play_url"`                                // 视频播放地址
	CoverUrl      string `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty" form:"cover_url" query:"cover_url"`                           // 视频封面地址
	FavoriteCount int64  `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty" form:"favorite_count" query:"favorite_count"` // 视频的点赞总数
	CommentCount  int64  `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty" form:"comment_count" query:"comment_count"`      // 视频的评论总数
	IsFavorite    bool   `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty" form:"is_favorite" query:"is_favorite"`                // true-已点赞，false-未点赞
	Title         string `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty" form:"title" query:"title"`                                                         // 视频标题
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userFavo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_userFavo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_userFavo_proto_rawDescGZIP(), []int{6}
}

func (x *Video) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *User {
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`                                                                    // 用户id
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" form:"name" query:"name"`                                                             // 用户名称
	FollowCount   int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty" form:"follow_count" query:"follow_count"`           // 关注总数
	FollowerCount int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty" form:"follower_count" query:"follower_count"` // 粉丝总数
	IsFollow      bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty" form:"is_follow" query:"is_follow"`                          // true-已关注，false-未关注
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userFavo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_userFavo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_userFavo_proto_rawDescGZIP(), []int{7}
}

func (x *User) GetId() uint64 {
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

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

var File_userFavo_proto protoreflect.FileDescriptor

var file_userFavo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x50, 0x62, 0x22, 0x6b, 0x0a, 0x0d,
	0x46, 0x61, 0x76, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4e, 0x0a, 0x0e, 0x46, 0x61, 0x76,
	0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x3b, 0x0a, 0x0b, 0x46, 0x61, 0x76,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7d, 0x0a, 0x0c, 0x46, 0x61, 0x76, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x12, 0x2f, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x46, 0x61,
	0x76, 0x6f, 0x50, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x6e, 0x0a, 0x0e, 0x46, 0x61, 0x76, 0x6f,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x22, 0xfc, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x50, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x32, 0xd6, 0x01, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x52, 0x70, 0x63, 0x12, 0x43, 0x0a, 0x0a, 0x46,
	0x61, 0x76, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x46, 0x61, 0x76, 0x6f, 0x50, 0x62, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x50,
	0x62, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x3d, 0x0a, 0x08, 0x46, 0x61, 0x76, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x50, 0x62, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f,
	0x50, 0x62, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x43, 0x0a, 0x0a, 0x46, 0x61, 0x76, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x50, 0x62, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x46,
	0x61, 0x76, 0x6f, 0x50, 0x62, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x27, 0x5a, 0x25, 0x70, 0x61, 0x69, 0x67, 0x75, 0x31, 0x39, 0x30,
	0x32, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x50, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userFavo_proto_rawDescOnce sync.Once
	file_userFavo_proto_rawDescData = file_userFavo_proto_rawDesc
)

func file_userFavo_proto_rawDescGZIP() []byte {
	file_userFavo_proto_rawDescOnce.Do(func() {
		file_userFavo_proto_rawDescData = protoimpl.X.CompressGZIP(file_userFavo_proto_rawDescData)
	})
	return file_userFavo_proto_rawDescData
}

var file_userFavo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_userFavo_proto_goTypes = []interface{}{
	(*FavoActionReq)(nil),  // 0: userFavoPb.FavoActionReq
	(*FavoActionResp)(nil), // 1: userFavoPb.FavoActionResp
	(*FavoListReq)(nil),    // 2: userFavoPb.FavoListReq
	(*FavoListResp)(nil),   // 3: userFavoPb.FavoListResp
	(*FavoStatusReq)(nil),  // 4: userFavoPb.FavoStatusReq
	(*FavoStatusResp)(nil), // 5: userFavoPb.FavoStatusResp
	(*Video)(nil),          // 6: userFavoPb.Video
	(*User)(nil),           // 7: userFavoPb.User
}
var file_userFavo_proto_depIdxs = []int32{
	6, // 0: userFavoPb.FavoListResp.videoList:type_name -> userFavoPb.Video
	7, // 1: userFavoPb.Video.author:type_name -> userFavoPb.User
	0, // 2: userFavoPb.UserFavoRpc.FavoAction:input_type -> userFavoPb.FavoActionReq
	2, // 3: userFavoPb.UserFavoRpc.FavoList:input_type -> userFavoPb.FavoListReq
	4, // 4: userFavoPb.UserFavoRpc.FavoStatus:input_type -> userFavoPb.FavoStatusReq
	1, // 5: userFavoPb.UserFavoRpc.FavoAction:output_type -> userFavoPb.FavoActionResp
	3, // 6: userFavoPb.UserFavoRpc.FavoList:output_type -> userFavoPb.FavoListResp
	5, // 7: userFavoPb.UserFavoRpc.FavoStatus:output_type -> userFavoPb.FavoStatusResp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_userFavo_proto_init() }
func file_userFavo_proto_init() {
	if File_userFavo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userFavo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoActionReq); i {
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
		file_userFavo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoActionResp); i {
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
		file_userFavo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoListReq); i {
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
		file_userFavo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoListResp); i {
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
		file_userFavo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoStatusReq); i {
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
		file_userFavo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoStatusResp); i {
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
		file_userFavo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_userFavo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_userFavo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userFavo_proto_goTypes,
		DependencyIndexes: file_userFavo_proto_depIdxs,
		MessageInfos:      file_userFavo_proto_msgTypes,
	}.Build()
	File_userFavo_proto = out.File
	file_userFavo_proto_rawDesc = nil
	file_userFavo_proto_goTypes = nil
	file_userFavo_proto_depIdxs = nil
}