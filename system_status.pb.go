// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: system_status.proto

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

// 系统状态
type SystemStatus int32

const (
	// 开机
	SystemStatus_ON SystemStatus = 0
	// 关机
	SystemStatus_OFF SystemStatus = 1
	// 待机
	SystemStatus_STANDBY SystemStatus = 2
)

// Enum value maps for SystemStatus.
var (
	SystemStatus_name = map[int32]string{
		0: "ON",
		1: "OFF",
		2: "STANDBY",
	}
	SystemStatus_value = map[string]int32{
		"ON":      0,
		"OFF":     1,
		"STANDBY": 2,
	}
)

func (x SystemStatus) Enum() *SystemStatus {
	p := new(SystemStatus)
	*p = x
	return p
}

func (x SystemStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SystemStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_system_status_proto_enumTypes[0].Descriptor()
}

func (SystemStatus) Type() protoreflect.EnumType {
	return &file_system_status_proto_enumTypes[0]
}

func (x SystemStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SystemStatus.Descriptor instead.
func (SystemStatus) EnumDescriptor() ([]byte, []int) {
	return file_system_status_proto_rawDescGZIP(), []int{0}
}

var File_system_status_proto protoreflect.FileDescriptor

var file_system_status_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x74, 0x72, 0x69, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x2c, 0x0a,
	0x0c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a,
	0x02, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x46, 0x46, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x42, 0x59, 0x10, 0x02, 0x42, 0x1c, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x50, 0x01, 0x5a, 0x0a, 0x2e,
	0x3b, 0x74, 0x72, 0x69, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_system_status_proto_rawDescOnce sync.Once
	file_system_status_proto_rawDescData = file_system_status_proto_rawDesc
)

func file_system_status_proto_rawDescGZIP() []byte {
	file_system_status_proto_rawDescOnce.Do(func() {
		file_system_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_status_proto_rawDescData)
	})
	return file_system_status_proto_rawDescData
}

var file_system_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_system_status_proto_goTypes = []interface{}{
	(SystemStatus)(0), // 0: triclass.system.status.SystemStatus
}
var file_system_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_system_status_proto_init() }
func file_system_status_proto_init() {
	if File_system_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_system_status_proto_goTypes,
		DependencyIndexes: file_system_status_proto_depIdxs,
		EnumInfos:         file_system_status_proto_enumTypes,
	}.Build()
	File_system_status_proto = out.File
	file_system_status_proto_rawDesc = nil
	file_system_status_proto_goTypes = nil
	file_system_status_proto_depIdxs = nil
}
