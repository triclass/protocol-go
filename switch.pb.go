// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: switch.proto

package triclass

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 开关状态
type Switch int32

const (
	// 开
	Switch_ON Switch = 0
	// 关
	Switch_OFF Switch = 1
)

// Enum value maps for Switch.
var (
	Switch_name = map[int32]string{
		0: "ON",
		1: "OFF",
	}
	Switch_value = map[string]int32{
		"ON":  0,
		"OFF": 1,
	}
)

func (x Switch) Enum() *Switch {
	p := new(Switch)
	*p = x
	return p
}

func (x Switch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Switch) Descriptor() protoreflect.EnumDescriptor {
	return file_switch_proto_enumTypes[0].Descriptor()
}

func (Switch) Type() protoreflect.EnumType {
	return &file_switch_proto_enumTypes[0]
}

func (x Switch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Switch.Descriptor instead.
func (Switch) EnumDescriptor() ([]byte, []int) {
	return file_switch_proto_rawDescGZIP(), []int{0}
}

var File_switch_proto protoreflect.FileDescriptor

var file_switch_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x74, 0x72, 0x69, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2a,
	0x19, 0x0a, 0x06, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4e, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x46, 0x46, 0x10, 0x01, 0x42, 0x3a, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x67, 0x6f, 0x3b, 0x74, 0x72,
	0x69, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_switch_proto_rawDescOnce sync.Once
	file_switch_proto_rawDescData = file_switch_proto_rawDesc
)

func file_switch_proto_rawDescGZIP() []byte {
	file_switch_proto_rawDescOnce.Do(func() {
		file_switch_proto_rawDescData = protoimpl.X.CompressGZIP(file_switch_proto_rawDescData)
	})
	return file_switch_proto_rawDescData
}

var file_switch_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_switch_proto_goTypes = []interface{}{
	(Switch)(0), // 0: triclass.switch.Switch
}
var file_switch_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_switch_proto_init() }
func file_switch_proto_init() {
	if File_switch_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_switch_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_switch_proto_goTypes,
		DependencyIndexes: file_switch_proto_depIdxs,
		EnumInfos:         file_switch_proto_enumTypes,
	}.Build()
	File_switch_proto = out.File
	file_switch_proto_rawDesc = nil
	file_switch_proto_goTypes = nil
	file_switch_proto_depIdxs = nil
}
