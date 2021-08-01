// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: bussiness/service/v1/restful.proto

package pbs_restful

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResourceFieldMask int32

const (
	ResourceFieldMask_RFM_NICK_NAME ResourceFieldMask = 0
	ResourceFieldMask_RFM_SEX       ResourceFieldMask = 1
	ResourceFieldMask_RFM_CONTENT   ResourceFieldMask = 2
	ResourceFieldMask_RFM_ADDRESS   ResourceFieldMask = 3
)

// Enum value maps for ResourceFieldMask.
var (
	ResourceFieldMask_name = map[int32]string{
		0: "RFM_NICK_NAME",
		1: "RFM_SEX",
		2: "RFM_CONTENT",
		3: "RFM_ADDRESS",
	}
	ResourceFieldMask_value = map[string]int32{
		"RFM_NICK_NAME": 0,
		"RFM_SEX":       1,
		"RFM_CONTENT":   2,
		"RFM_ADDRESS":   3,
	}
)

func (x ResourceFieldMask) Enum() *ResourceFieldMask {
	p := new(ResourceFieldMask)
	*p = x
	return p
}

func (x ResourceFieldMask) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceFieldMask) Descriptor() protoreflect.EnumDescriptor {
	return file_bussiness_service_v1_restful_proto_enumTypes[0].Descriptor()
}

func (ResourceFieldMask) Type() protoreflect.EnumType {
	return &file_bussiness_service_v1_restful_proto_enumTypes[0]
}

func (x ResourceFieldMask) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceFieldMask.Descriptor instead.
func (ResourceFieldMask) EnumDescriptor() ([]byte, []int) {
	return file_bussiness_service_v1_restful_proto_rawDescGZIP(), []int{0}
}

// Address 地址
type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	F1 string `protobuf:"bytes,1,opt,name=F1,proto3" json:"F1,omitempty" xorm:"varchar(255) unsigned not null default '' 'F1'"`
	F2 string `protobuf:"bytes,2,opt,name=F2,proto3" json:"F2,omitempty" xorm:"varchar(255) unsigned not null default '' 'F2'"`
	F3 string `protobuf:"bytes,3,opt,name=F3,proto3" json:"F3,omitempty" xorm:"varchar(255) unsigned not null default '' 'F3'"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bussiness_service_v1_restful_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_bussiness_service_v1_restful_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_bussiness_service_v1_restful_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetF1() string {
	if x != nil {
		return x.F1
	}
	return ""
}

func (x *Address) GetF2() string {
	if x != nil {
		return x.F2
	}
	return ""
}

func (x *Address) GetF3() string {
	if x != nil {
		return x.F3
	}
	return ""
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NickName 昵称
	NickName string `protobuf:"bytes,1,opt,name=NickName,proto3" json:"NickName,omitempty" xorm:"varchar(255) unsigned not null default '' 'NickName'"`
	// Sex 性别
	Sex int32 `protobuf:"varint,2,opt,name=Sex,proto3" json:"Sex,omitempty" xorm:"int not null default 0 'Sex'"`
	// Content 内容
	Content string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty" xorm:"varchar(255) unsigned not null default '' 'Content'"`
	// Addr 地址
	Addr *Address `protobuf:"bytes,4,opt,name=Addr,proto3" json:"Addr,omitempty" xorm:"text json 'Addr'"`
	// ID 资源唯一标识符
	ID int64 `protobuf:"varint,5,opt,name=ID,proto3" json:"ID,omitempty" xorm:"bigint not null default 0 'ID'"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bussiness_service_v1_restful_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_bussiness_service_v1_restful_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_bussiness_service_v1_restful_proto_rawDescGZIP(), []int{1}
}

func (x *Resource) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *Resource) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *Resource) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Resource) GetAddr() *Address {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *Resource) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID 资源ID
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty" xorm:"bigint not null default 0 'ID'"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bussiness_service_v1_restful_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bussiness_service_v1_restful_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_bussiness_service_v1_restful_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

//
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PageToken 页面Token
	PageToken string `protobuf:"bytes,1,opt,name=PageToken,proto3" json:"PageToken,omitempty" xorm:"varchar(255) unsigned not null default '' 'PageToken'"`
	PageSize  int32  `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty" xorm:"int not null default 0 'PageSize'"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bussiness_service_v1_restful_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bussiness_service_v1_restful_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_bussiness_service_v1_restful_proto_rawDescGZIP(), []int{3}
}

func (x *ListRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NextPageToken 下一页的页面Token
	NextPageToken string `protobuf:"bytes,1,opt,name=NextPageToken,proto3" json:"NextPageToken,omitempty" xorm:"varchar(255) unsigned not null default '' 'NextPageToken'"`
	// LastPageToken 上一页的页面Token
	LastPageToken string `protobuf:"bytes,2,opt,name=LastPageToken,proto3" json:"LastPageToken,omitempty" xorm:"varchar(255) unsigned not null default '' 'LastPageToken'"`
	// List 列表
	List []*Resource `protobuf:"bytes,3,rep,name=List,proto3" json:"List,omitempty" xorm:"text json 'List'"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bussiness_service_v1_restful_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bussiness_service_v1_restful_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_bussiness_service_v1_restful_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListResponse) GetLastPageToken() string {
	if x != nil {
		return x.LastPageToken
	}
	return ""
}

func (x *ListResponse) GetList() []*Resource {
	if x != nil {
		return x.List
	}
	return nil
}

type UpdateItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mask 掩码
	Mask ResourceFieldMask `protobuf:"varint,1,opt,name=Mask,proto3,enum=Bussiness.Service.Restful.ResourceFieldMask" json:"Mask,omitempty" xorm:"int unsigned not null default 0 'Mask'"`
	// Data 与mask相对应的值
	Data *Resource `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty" xorm:"text json 'Data'"`
}

func (x *UpdateItem) Reset() {
	*x = UpdateItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bussiness_service_v1_restful_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItem) ProtoMessage() {}

func (x *UpdateItem) ProtoReflect() protoreflect.Message {
	mi := &file_bussiness_service_v1_restful_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItem.ProtoReflect.Descriptor instead.
func (*UpdateItem) Descriptor() ([]byte, []int) {
	return file_bussiness_service_v1_restful_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateItem) GetMask() ResourceFieldMask {
	if x != nil {
		return x.Mask
	}
	return ResourceFieldMask_RFM_NICK_NAME
}

func (x *UpdateItem) GetData() *Resource {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Items 更新条目的一个集合
	Items []*UpdateItem `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty" xorm:"text json 'Items'"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bussiness_service_v1_restful_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bussiness_service_v1_restful_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_bussiness_service_v1_restful_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRequest) GetItems() []*UpdateItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_bussiness_service_v1_restful_proto protoreflect.FileDescriptor

var file_bussiness_service_v1_restful_proto_rawDesc = []byte{
	0x0a, 0x22, 0x62, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x46, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x46, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x46, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x46, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x46, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x46, 0x33, 0x22, 0x9a, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x53, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x53, 0x65, 0x78,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x41, 0x64,
	0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x42, 0x75, 0x73, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x66, 0x75, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x04, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x47, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x4e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x24, 0x0a, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x87, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x40,
	0x0a, 0x04, 0x4d, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x42,
	0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x37, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x05, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x42, 0x75, 0x73, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0x55, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x11, 0x0a, 0x0d,
	0x52, 0x46, 0x4d, 0x5f, 0x4e, 0x49, 0x43, 0x4b, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x52, 0x46, 0x4d, 0x5f, 0x53, 0x45, 0x58, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x52, 0x46, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0f, 0x0a,
	0x0b, 0x52, 0x46, 0x4d, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x03, 0x32, 0xfa,
	0x05, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x12, 0xc5, 0x01, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x66, 0x75,
	0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x23, 0x2e, 0x42, 0x75, 0x73,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0x71, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x6b, 0x22, 0x35, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x2f, 0x7b, 0x53, 0x65, 0x78, 0x7d,
	0x2f, 0x7b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x7d, 0x2f, 0x7b, 0x41, 0x64, 0x64, 0x72,
	0x2e, 0x46, 0x31, 0x7d, 0x2f, 0x7b, 0x41, 0x64, 0x64, 0x72, 0x2e, 0x46, 0x32, 0x7d, 0x3a, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x2e, 0x46, 0x33, 0x5a, 0x29, 0x32, 0x21, 0x2f, 0x76, 0x32, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x2f, 0x7b, 0x53, 0x65,
	0x78, 0x7d, 0x2f, 0x7b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x7d, 0x3a, 0x04, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x6c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x42, 0x75, 0x73, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x7b, 0x49, 0x44, 0x7d,
	0x12, 0x72, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x42, 0x75, 0x73, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x66, 0x75, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x70, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28,
	0x2e, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x42, 0x75, 0x73, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x66, 0x75, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x32, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x12, 0x28, 0x2e, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x42, 0x75,
	0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x1a, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x60, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x42, 0x4a, 0x5a, 0x48, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x2f, 0x67, 0x6f, 0x70, 0x62, 0x2f, 0x62, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x2f,
	0x70, 0x62, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x3b, 0x70, 0x62, 0x73, 0x5f,
	0x72, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bussiness_service_v1_restful_proto_rawDescOnce sync.Once
	file_bussiness_service_v1_restful_proto_rawDescData = file_bussiness_service_v1_restful_proto_rawDesc
)

func file_bussiness_service_v1_restful_proto_rawDescGZIP() []byte {
	file_bussiness_service_v1_restful_proto_rawDescOnce.Do(func() {
		file_bussiness_service_v1_restful_proto_rawDescData = protoimpl.X.CompressGZIP(file_bussiness_service_v1_restful_proto_rawDescData)
	})
	return file_bussiness_service_v1_restful_proto_rawDescData
}

var file_bussiness_service_v1_restful_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bussiness_service_v1_restful_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bussiness_service_v1_restful_proto_goTypes = []interface{}{
	(ResourceFieldMask)(0), // 0: Bussiness.Service.Restful.ResourceFieldMask
	(*Address)(nil),        // 1: Bussiness.Service.Restful.Address
	(*Resource)(nil),       // 2: Bussiness.Service.Restful.Resource
	(*GetRequest)(nil),     // 3: Bussiness.Service.Restful.GetRequest
	(*ListRequest)(nil),    // 4: Bussiness.Service.Restful.ListRequest
	(*ListResponse)(nil),   // 5: Bussiness.Service.Restful.ListResponse
	(*UpdateItem)(nil),     // 6: Bussiness.Service.Restful.UpdateItem
	(*UpdateRequest)(nil),  // 7: Bussiness.Service.Restful.UpdateRequest
	(*emptypb.Empty)(nil),  // 8: google.protobuf.Empty
}
var file_bussiness_service_v1_restful_proto_depIdxs = []int32{
	1,  // 0: Bussiness.Service.Restful.Resource.Addr:type_name -> Bussiness.Service.Restful.Address
	2,  // 1: Bussiness.Service.Restful.ListResponse.List:type_name -> Bussiness.Service.Restful.Resource
	0,  // 2: Bussiness.Service.Restful.UpdateItem.Mask:type_name -> Bussiness.Service.Restful.ResourceFieldMask
	2,  // 3: Bussiness.Service.Restful.UpdateItem.Data:type_name -> Bussiness.Service.Restful.Resource
	6,  // 4: Bussiness.Service.Restful.UpdateRequest.Items:type_name -> Bussiness.Service.Restful.UpdateItem
	2,  // 5: Bussiness.Service.Restful.Restful.Create:input_type -> Bussiness.Service.Restful.Resource
	3,  // 6: Bussiness.Service.Restful.Restful.Get:input_type -> Bussiness.Service.Restful.GetRequest
	4,  // 7: Bussiness.Service.Restful.Restful.List:input_type -> Bussiness.Service.Restful.ListRequest
	7,  // 8: Bussiness.Service.Restful.Restful.Update:input_type -> Bussiness.Service.Restful.UpdateRequest
	7,  // 9: Bussiness.Service.Restful.Restful.Replace:input_type -> Bussiness.Service.Restful.UpdateRequest
	2,  // 10: Bussiness.Service.Restful.Restful.Delete:input_type -> Bussiness.Service.Restful.Resource
	2,  // 11: Bussiness.Service.Restful.Restful.Create:output_type -> Bussiness.Service.Restful.Resource
	2,  // 12: Bussiness.Service.Restful.Restful.Get:output_type -> Bussiness.Service.Restful.Resource
	5,  // 13: Bussiness.Service.Restful.Restful.List:output_type -> Bussiness.Service.Restful.ListResponse
	2,  // 14: Bussiness.Service.Restful.Restful.Update:output_type -> Bussiness.Service.Restful.Resource
	2,  // 15: Bussiness.Service.Restful.Restful.Replace:output_type -> Bussiness.Service.Restful.Resource
	8,  // 16: Bussiness.Service.Restful.Restful.Delete:output_type -> google.protobuf.Empty
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_bussiness_service_v1_restful_proto_init() }
func file_bussiness_service_v1_restful_proto_init() {
	if File_bussiness_service_v1_restful_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bussiness_service_v1_restful_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_bussiness_service_v1_restful_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_bussiness_service_v1_restful_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_bussiness_service_v1_restful_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_bussiness_service_v1_restful_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_bussiness_service_v1_restful_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItem); i {
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
		file_bussiness_service_v1_restful_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
			RawDescriptor: file_bussiness_service_v1_restful_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bussiness_service_v1_restful_proto_goTypes,
		DependencyIndexes: file_bussiness_service_v1_restful_proto_depIdxs,
		EnumInfos:         file_bussiness_service_v1_restful_proto_enumTypes,
		MessageInfos:      file_bussiness_service_v1_restful_proto_msgTypes,
	}.Build()
	File_bussiness_service_v1_restful_proto = out.File
	file_bussiness_service_v1_restful_proto_rawDesc = nil
	file_bussiness_service_v1_restful_proto_goTypes = nil
	file_bussiness_service_v1_restful_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RestfulClient is the client API for Restful service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RestfulClient interface {
	// Create 创建
	Create(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*Resource, error)
	// Get 获取单个
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Resource, error)
	// List 拉取结果集
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Update 更新
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Resource, error)
	// Update 更新
	Replace(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Resource, error)
	// Delete 删除
	Delete(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type restfulClient struct {
	cc grpc.ClientConnInterface
}

func NewRestfulClient(cc grpc.ClientConnInterface) RestfulClient {
	return &restfulClient{cc}
}

func (c *restfulClient) Create(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, "/Bussiness.Service.Restful.Restful/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restfulClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, "/Bussiness.Service.Restful.Restful/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restfulClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/Bussiness.Service.Restful.Restful/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restfulClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, "/Bussiness.Service.Restful.Restful/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restfulClient) Replace(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, "/Bussiness.Service.Restful.Restful/Replace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restfulClient) Delete(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Bussiness.Service.Restful.Restful/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestfulServer is the server API for Restful service.
type RestfulServer interface {
	// Create 创建
	Create(context.Context, *Resource) (*Resource, error)
	// Get 获取单个
	Get(context.Context, *GetRequest) (*Resource, error)
	// List 拉取结果集
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Update 更新
	Update(context.Context, *UpdateRequest) (*Resource, error)
	// Update 更新
	Replace(context.Context, *UpdateRequest) (*Resource, error)
	// Delete 删除
	Delete(context.Context, *Resource) (*emptypb.Empty, error)
}

// UnimplementedRestfulServer can be embedded to have forward compatible implementations.
type UnimplementedRestfulServer struct {
}

func (*UnimplementedRestfulServer) Create(context.Context, *Resource) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRestfulServer) Get(context.Context, *GetRequest) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRestfulServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedRestfulServer) Update(context.Context, *UpdateRequest) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedRestfulServer) Replace(context.Context, *UpdateRequest) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replace not implemented")
}
func (*UnimplementedRestfulServer) Delete(context.Context, *Resource) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterRestfulServer(s *grpc.Server, srv RestfulServer) {
	s.RegisterService(&_Restful_serviceDesc, srv)
}

func _Restful_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestfulServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bussiness.Service.Restful.Restful/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestfulServer).Create(ctx, req.(*Resource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restful_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestfulServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bussiness.Service.Restful.Restful/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestfulServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restful_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestfulServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bussiness.Service.Restful.Restful/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestfulServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restful_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestfulServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bussiness.Service.Restful.Restful/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestfulServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restful_Replace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestfulServer).Replace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bussiness.Service.Restful.Restful/Replace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestfulServer).Replace(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restful_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestfulServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bussiness.Service.Restful.Restful/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestfulServer).Delete(ctx, req.(*Resource))
	}
	return interceptor(ctx, in, info, handler)
}

var _Restful_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Bussiness.Service.Restful.Restful",
	HandlerType: (*RestfulServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Restful_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Restful_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Restful_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Restful_Update_Handler,
		},
		{
			MethodName: "Replace",
			Handler:    _Restful_Replace_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Restful_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bussiness/service/v1/restful.proto",
}