// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: hardware/system.proto

package hardware

import (
	proto "github.com/golang/protobuf/proto"
	protocol_go "github.com/triclass/protocol-go"
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

// 系统控制
type System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//	*System_Status
	Action isSystem_Action `protobuf_oneof:"action"`
}

func (x *System) Reset() {
	*x = System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hardware_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*System) ProtoMessage() {}

func (x *System) ProtoReflect() protoreflect.Message {
	mi := &file_hardware_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use System.ProtoReflect.Descriptor instead.
func (*System) Descriptor() ([]byte, []int) {
	return file_hardware_system_proto_rawDescGZIP(), []int{0}
}

func (m *System) GetAction() isSystem_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *System) GetStatus() protocol_go.SystemStatus {
	if x, ok := x.GetAction().(*System_Status); ok {
		return x.Status
	}
	return protocol_go.SystemStatus_UNKNOWN
}

type isSystem_Action interface {
	isSystem_Action()
}

type System_Status struct {
	// 设置系统状态
	Status protocol_go.SystemStatus `protobuf:"varint,1,opt,name=status,proto3,enum=triclass.system.status.SystemStatus,oneof"`
}

func (*System_Status) isSystem_Action() {}

var File_hardware_system_proto protoreflect.FileDescriptor

var file_hardware_system_proto_rawDesc = []byte{
	0x0a, 0x15, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x72, 0x69, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x2e, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x1a, 0x13, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x52, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x74, 0x72, 0x69, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x2e, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x50, 0x01, 0x5a, 0x13,
	0x2e, 0x2f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x3b, 0x68, 0x61, 0x72, 0x64, 0x77,
	0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hardware_system_proto_rawDescOnce sync.Once
	file_hardware_system_proto_rawDescData = file_hardware_system_proto_rawDesc
)

func file_hardware_system_proto_rawDescGZIP() []byte {
	file_hardware_system_proto_rawDescOnce.Do(func() {
		file_hardware_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_hardware_system_proto_rawDescData)
	})
	return file_hardware_system_proto_rawDescData
}

var file_hardware_system_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hardware_system_proto_goTypes = []interface{}{
	(*System)(nil),                // 0: triclass.hardware.System
	(protocol_go.SystemStatus)(0), // 1: triclass.system.status.SystemStatus
}
var file_hardware_system_proto_depIdxs = []int32{
	1, // 0: triclass.hardware.System.status:type_name -> triclass.system.status.SystemStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hardware_system_proto_init() }
func file_hardware_system_proto_init() {
	if File_hardware_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hardware_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*System); i {
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
	file_hardware_system_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*System_Status)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hardware_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hardware_system_proto_goTypes,
		DependencyIndexes: file_hardware_system_proto_depIdxs,
		MessageInfos:      file_hardware_system_proto_msgTypes,
	}.Build()
	File_hardware_system_proto = out.File
	file_hardware_system_proto_rawDesc = nil
	file_hardware_system_proto_goTypes = nil
	file_hardware_system_proto_depIdxs = nil
}
