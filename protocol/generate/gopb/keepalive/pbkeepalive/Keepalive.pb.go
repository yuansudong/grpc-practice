// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: keepalive/Keepalive.proto

package pbkeepalive

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

// Code 用于描述一个返回码
type Code int32

const (
	// C_OK 代表成功
	Code_C_OK Code = 0
	// MainProtocol 代表协议错误
	Code_C_MAIN_PROTOCOL Code = 1
	// C_ARG 参数错误
	Code_C_ARG Code = 2
	// C_AUTH_TIMEOUT 超时
	Code_C_AUTH_TIMEOUT Code = 3
	// C_BODY_TOO_LARGE body太大
	Code_C_BODY_TOO_LARGE Code = 4
	// C_CMD cmd不对
	Code_C_CMD Code = 5
	// C_SUB_PROTOCOL 子协议错误
	Code_C_SUB_PROTOCOL Code = 6
	// C_NEED_AUTH 需要认证
	Code_C_NEED_AUTH Code = 7
	// C_NON_ACTIVE 不活跃,此Number下发之后,服务端会关闭连接
	Code_C_NON_ACTIVE Code = 8
	// C_AUTH 认证失败
	Code_C_AUTH Code = 9
	// C_GPOOL GRPC连接池报错
	Code_C_GPOOL Code = 10
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:  "C_OK",
		1:  "C_MAIN_PROTOCOL",
		2:  "C_ARG",
		3:  "C_AUTH_TIMEOUT",
		4:  "C_BODY_TOO_LARGE",
		5:  "C_CMD",
		6:  "C_SUB_PROTOCOL",
		7:  "C_NEED_AUTH",
		8:  "C_NON_ACTIVE",
		9:  "C_AUTH",
		10: "C_GPOOL",
	}
	Code_value = map[string]int32{
		"C_OK":             0,
		"C_MAIN_PROTOCOL":  1,
		"C_ARG":            2,
		"C_AUTH_TIMEOUT":   3,
		"C_BODY_TOO_LARGE": 4,
		"C_CMD":            5,
		"C_SUB_PROTOCOL":   6,
		"C_NEED_AUTH":      7,
		"C_NON_ACTIVE":     8,
		"C_AUTH":           9,
		"C_GPOOL":          10,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_keepalive_Keepalive_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_keepalive_Keepalive_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_keepalive_Keepalive_proto_rawDescGZIP(), []int{0}
}

var File_keepalive_Keepalive_proto protoreflect.FileDescriptor

var file_keepalive_Keepalive_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x4b, 0x65, 0x65, 0x70,
	0x61, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6b, 0x65, 0x65,
	0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x2a, 0xb5, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x43, 0x5f, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x5f, 0x4d,
	0x41, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x43, 0x5f, 0x41, 0x52, 0x47, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x5f, 0x41,
	0x55, 0x54, 0x48, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x14, 0x0a,
	0x10, 0x43, 0x5f, 0x42, 0x4f, 0x44, 0x59, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47,
	0x45, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x5f, 0x43, 0x4d, 0x44, 0x10, 0x05, 0x12, 0x12,
	0x0a, 0x0e, 0x43, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c,
	0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x5f, 0x4e, 0x45, 0x45, 0x44, 0x5f, 0x41, 0x55, 0x54,
	0x48, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x10,
	0x09, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x5f, 0x47, 0x50, 0x4f, 0x4f, 0x4c, 0x10, 0x0a, 0x42, 0x8d,
	0x01, 0x0a, 0x17, 0x7a, 0x77, 0x77, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x42, 0x0b, 0x50, 0x42, 0x4b, 0x65,
	0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x50, 0x01, 0x5a, 0x38, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x6f, 0x70,
	0x62, 0x2f, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x70, 0x62, 0x6b, 0x65,
	0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x3b, 0x70, 0x62, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x0b, 0x50, 0x42, 0x4b, 0x65, 0x65, 0x70, 0x61,
	0x6c, 0x69, 0x76, 0x65, 0xaa, 0x02, 0x17, 0x5a, 0x77, 0x77, 0x78, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keepalive_Keepalive_proto_rawDescOnce sync.Once
	file_keepalive_Keepalive_proto_rawDescData = file_keepalive_Keepalive_proto_rawDesc
)

func file_keepalive_Keepalive_proto_rawDescGZIP() []byte {
	file_keepalive_Keepalive_proto_rawDescOnce.Do(func() {
		file_keepalive_Keepalive_proto_rawDescData = protoimpl.X.CompressGZIP(file_keepalive_Keepalive_proto_rawDescData)
	})
	return file_keepalive_Keepalive_proto_rawDescData
}

var file_keepalive_Keepalive_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_keepalive_Keepalive_proto_goTypes = []interface{}{
	(Code)(0), // 0: keepalive.Code
}
var file_keepalive_Keepalive_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_keepalive_Keepalive_proto_init() }
func file_keepalive_Keepalive_proto_init() {
	if File_keepalive_Keepalive_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_keepalive_Keepalive_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_keepalive_Keepalive_proto_goTypes,
		DependencyIndexes: file_keepalive_Keepalive_proto_depIdxs,
		EnumInfos:         file_keepalive_Keepalive_proto_enumTypes,
	}.Build()
	File_keepalive_Keepalive_proto = out.File
	file_keepalive_Keepalive_proto_rawDesc = nil
	file_keepalive_Keepalive_proto_goTypes = nil
	file_keepalive_Keepalive_proto_depIdxs = nil
}
