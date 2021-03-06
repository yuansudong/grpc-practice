// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: common/channel.proto

package pbm_channel

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

// Channel 用于定义发布渠道
type Channel int32

const (
	// C_UNKNOWN 渠道未知
	Channel_C_UNKNWON Channel = 0
	// C_WX 渠道为微信
	Channel_C_WX Channel = 1
	// C_4399 渠道为C4399
	Channel_C_4399 Channel = 2
	// C_FB 渠道为脸书
	Channel_C_FB Channel = 3
	// C_THL 渠道为今日头条
	Channel_C_THL Channel = 4
	// C_ALIPAY 渠道为支付宝小程序
	Channel_C_ALIPAY Channel = 5
	// C_VIVO 渠道为vivo
	Channel_C_VIVO Channel = 6
	// C_HUAWEI 渠道为华为
	Channel_C_HUAWEI Channel = 7
	// C_MEIZU 渠道为魅族
	Channel_C_MEIZU Channel = 8
	// C_OPPO 渠道为OPPO
	Channel_C_OPPO Channel = 9
	// C_APPLE 渠道为苹果
	Channel_C_APPLE Channel = 10
	// C_STEAM 渠道为steam
	Channel_C_STEAM Channel = 11
	// C_GOOGLE 渠道为google
	Channel_C_GOOGLE Channel = 12
	// C_TWITTER 渠道为推特
	Channel_C_TWITTER Channel = 13
)

// Enum value maps for Channel.
var (
	Channel_name = map[int32]string{
		0:  "C_UNKNWON",
		1:  "C_WX",
		2:  "C_4399",
		3:  "C_FB",
		4:  "C_THL",
		5:  "C_ALIPAY",
		6:  "C_VIVO",
		7:  "C_HUAWEI",
		8:  "C_MEIZU",
		9:  "C_OPPO",
		10: "C_APPLE",
		11: "C_STEAM",
		12: "C_GOOGLE",
		13: "C_TWITTER",
	}
	Channel_value = map[string]int32{
		"C_UNKNWON": 0,
		"C_WX":      1,
		"C_4399":    2,
		"C_FB":      3,
		"C_THL":     4,
		"C_ALIPAY":  5,
		"C_VIVO":    6,
		"C_HUAWEI":  7,
		"C_MEIZU":   8,
		"C_OPPO":    9,
		"C_APPLE":   10,
		"C_STEAM":   11,
		"C_GOOGLE":  12,
		"C_TWITTER": 13,
	}
)

func (x Channel) Enum() *Channel {
	p := new(Channel)
	*p = x
	return p
}

func (x Channel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Channel) Descriptor() protoreflect.EnumDescriptor {
	return file_common_channel_proto_enumTypes[0].Descriptor()
}

func (Channel) Type() protoreflect.EnumType {
	return &file_common_channel_proto_enumTypes[0]
}

func (x Channel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Channel.Descriptor instead.
func (Channel) EnumDescriptor() ([]byte, []int) {
	return file_common_channel_proto_rawDescGZIP(), []int{0}
}

var File_common_channel_proto protoreflect.FileDescriptor

var file_common_channel_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0xbb,
	0x01, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x57, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x5f, 0x57,
	0x58, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x5f, 0x34, 0x33, 0x39, 0x39, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x43, 0x5f, 0x46, 0x42, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x5f, 0x54,
	0x48, 0x4c, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x5f, 0x41, 0x4c, 0x49, 0x50, 0x41, 0x59,
	0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x5f, 0x56, 0x49, 0x56, 0x4f, 0x10, 0x06, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x5f, 0x48, 0x55, 0x41, 0x57, 0x45, 0x49, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x5f, 0x4d, 0x45, 0x49, 0x5a, 0x55, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x5f, 0x4f,
	0x50, 0x50, 0x4f, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x45,
	0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x5f, 0x53, 0x54, 0x45, 0x41, 0x4d, 0x10, 0x0b, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x10, 0x0c, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x5f, 0x54, 0x57, 0x49, 0x54, 0x54, 0x45, 0x52, 0x10, 0x0d, 0x42, 0x57, 0x5a, 0x35,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x2f, 0x67, 0x6f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62,
	0x6d, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x3b, 0x70, 0x62, 0x6d, 0x5f, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x1a, 0x5a, 0x57, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_channel_proto_rawDescOnce sync.Once
	file_common_channel_proto_rawDescData = file_common_channel_proto_rawDesc
)

func file_common_channel_proto_rawDescGZIP() []byte {
	file_common_channel_proto_rawDescOnce.Do(func() {
		file_common_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_channel_proto_rawDescData)
	})
	return file_common_channel_proto_rawDescData
}

var file_common_channel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_channel_proto_goTypes = []interface{}{
	(Channel)(0), // 0: Common.Channel
}
var file_common_channel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_channel_proto_init() }
func file_common_channel_proto_init() {
	if File_common_channel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_channel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_channel_proto_goTypes,
		DependencyIndexes: file_common_channel_proto_depIdxs,
		EnumInfos:         file_common_channel_proto_enumTypes,
	}.Build()
	File_common_channel_proto = out.File
	file_common_channel_proto_rawDesc = nil
	file_common_channel_proto_goTypes = nil
	file_common_channel_proto_depIdxs = nil
}
