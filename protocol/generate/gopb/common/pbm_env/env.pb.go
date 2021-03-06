// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: common/env.proto

package pbm_env

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

// 代表环境
type Environment int32

const (
	// E_UNKNOWN 环境未知
	Environment_E_UNKNOWN Environment = 0
	// E_TEST 测试环境
	Environment_E_TEST Environment = 1
	// E_DEBUG 开发环境
	Environment_E_DEV Environment = 2
	// E_PROD 正式环境
	Environment_E_PROD Environment = 4 //
)

// Enum value maps for Environment.
var (
	Environment_name = map[int32]string{
		0: "E_UNKNOWN",
		1: "E_TEST",
		2: "E_DEV",
		4: "E_PROD",
	}
	Environment_value = map[string]int32{
		"E_UNKNOWN": 0,
		"E_TEST":    1,
		"E_DEV":     2,
		"E_PROD":    4,
	}
)

func (x Environment) Enum() *Environment {
	p := new(Environment)
	*p = x
	return p
}

func (x Environment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Environment) Descriptor() protoreflect.EnumDescriptor {
	return file_common_env_proto_enumTypes[0].Descriptor()
}

func (Environment) Type() protoreflect.EnumType {
	return &file_common_env_proto_enumTypes[0]
}

func (x Environment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Environment.Descriptor instead.
func (Environment) EnumDescriptor() ([]byte, []int) {
	return file_common_env_proto_rawDescGZIP(), []int{0}
}

var File_common_env_proto protoreflect.FileDescriptor

var file_common_env_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x76, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0x3f, 0x0a, 0x0b, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x5f, 0x54, 0x45,
	0x53, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x10, 0x04, 0x42, 0x49, 0x5a, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x2f, 0x67, 0x6f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x6d,
	0x5f, 0x65, 0x6e, 0x76, 0x3b, 0x70, 0x62, 0x6d, 0x5f, 0x65, 0x6e, 0x76, 0xf8, 0x01, 0x01, 0xaa,
	0x02, 0x14, 0x5a, 0x77, 0x77, 0x78, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_env_proto_rawDescOnce sync.Once
	file_common_env_proto_rawDescData = file_common_env_proto_rawDesc
)

func file_common_env_proto_rawDescGZIP() []byte {
	file_common_env_proto_rawDescOnce.Do(func() {
		file_common_env_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_env_proto_rawDescData)
	})
	return file_common_env_proto_rawDescData
}

var file_common_env_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_env_proto_goTypes = []interface{}{
	(Environment)(0), // 0: Common.Environment
}
var file_common_env_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_env_proto_init() }
func file_common_env_proto_init() {
	if File_common_env_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_env_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_env_proto_goTypes,
		DependencyIndexes: file_common_env_proto_depIdxs,
		EnumInfos:         file_common_env_proto_enumTypes,
	}.Build()
	File_common_env_proto = out.File
	file_common_env_proto_rawDesc = nil
	file_common_env_proto_goTypes = nil
	file_common_env_proto_depIdxs = nil
}
