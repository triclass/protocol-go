// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: dm/heartbeat.proto

package dm

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

// 心跳
type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 学校编号
	SchoolId int64 `protobuf:"varint,1,opt,name=schoolId,proto3" json:"schoolId,omitempty"`
	// 厂商
	Manufacturer protocol_go.Manufacturer `protobuf:"varint,2,opt,name=manufacturer,proto3,enum=triclass.manufacturer.Manufacturer" json:"manufacturer,omitempty"`
	// 设备序列号
	Sn string `protobuf:"bytes,3,opt,name=sn,proto3" json:"sn,omitempty"`
	// 设备状态
	Status protocol_go.SystemStatus `protobuf:"varint,4,opt,name=status,proto3,enum=triclass.system.status.SystemStatus" json:"status,omitempty"`
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dm_heartbeat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_dm_heartbeat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_dm_heartbeat_proto_rawDescGZIP(), []int{0}
}

func (x *Heartbeat) GetSchoolId() int64 {
	if x != nil {
		return x.SchoolId
	}
	return 0
}

func (x *Heartbeat) GetManufacturer() protocol_go.Manufacturer {
	if x != nil {
		return x.Manufacturer
	}
	return protocol_go.Manufacturer_UNKNOWN
}

func (x *Heartbeat) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

func (x *Heartbeat) GetStatus() protocol_go.SystemStatus {
	if x != nil {
		return x.Status
	}
	return protocol_go.SystemStatus_UNKNOWN
}

var File_dm_heartbeat_proto protoreflect.FileDescriptor

var file_dm_heartbeat_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x6d, 0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x69, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x64,
	0x6d, 0x1a, 0x12, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x09, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x72, 0x69,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x72, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x52,
	0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x73, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x3c, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e,
	0x74, 0x72, 0x69, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x1c, 0x0a, 0x0f, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x64, 0x6d, 0x50, 0x01,
	0x5a, 0x07, 0x2e, 0x2f, 0x64, 0x6d, 0x3b, 0x64, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_dm_heartbeat_proto_rawDescOnce sync.Once
	file_dm_heartbeat_proto_rawDescData = file_dm_heartbeat_proto_rawDesc
)

func file_dm_heartbeat_proto_rawDescGZIP() []byte {
	file_dm_heartbeat_proto_rawDescOnce.Do(func() {
		file_dm_heartbeat_proto_rawDescData = protoimpl.X.CompressGZIP(file_dm_heartbeat_proto_rawDescData)
	})
	return file_dm_heartbeat_proto_rawDescData
}

var file_dm_heartbeat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dm_heartbeat_proto_goTypes = []interface{}{
	(*Heartbeat)(nil),             // 0: triclass.dm.Heartbeat
	(protocol_go.Manufacturer)(0), // 1: triclass.manufacturer.Manufacturer
	(protocol_go.SystemStatus)(0), // 2: triclass.system.status.SystemStatus
}
var file_dm_heartbeat_proto_depIdxs = []int32{
	1, // 0: triclass.dm.Heartbeat.manufacturer:type_name -> triclass.manufacturer.Manufacturer
	2, // 1: triclass.dm.Heartbeat.status:type_name -> triclass.system.status.SystemStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dm_heartbeat_proto_init() }
func file_dm_heartbeat_proto_init() {
	if File_dm_heartbeat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dm_heartbeat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
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
			RawDescriptor: file_dm_heartbeat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dm_heartbeat_proto_goTypes,
		DependencyIndexes: file_dm_heartbeat_proto_depIdxs,
		MessageInfos:      file_dm_heartbeat_proto_msgTypes,
	}.Build()
	File_dm_heartbeat_proto = out.File
	file_dm_heartbeat_proto_rawDesc = nil
	file_dm_heartbeat_proto_goTypes = nil
	file_dm_heartbeat_proto_depIdxs = nil
}
