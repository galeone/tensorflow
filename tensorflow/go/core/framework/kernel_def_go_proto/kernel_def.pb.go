// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: tensorflow/core/framework/kernel_def.proto

package kernel_def_go_proto

import (
	attr_value_go_proto "github.com/galeone/tensorflow/tensorflow/go/core/framework/attr_value_go_proto"
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

type KernelDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must match the name of an Op.
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// Type of device this kernel runs on.
	DeviceType string                      `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	Constraint []*KernelDef_AttrConstraint `protobuf:"bytes,3,rep,name=constraint,proto3" json:"constraint,omitempty"`
	// Names of the Op's input_/output_args that reside in host memory
	// instead of device memory.
	HostMemoryArg []string `protobuf:"bytes,4,rep,name=host_memory_arg,json=hostMemoryArg,proto3" json:"host_memory_arg,omitempty"`
	// This allows experimental kernels to be registered for an op that
	// won't be used unless the user specifies a "_kernel" attr with
	// value matching this.
	Label string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	// Prioritization of kernel amongst different devices. By default we assume
	// priority is 0. The higher the priority the better. By default (i.e. if
	// this is not set), we prefer GPU kernels over CPU.
	Priority int32 `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *KernelDef) Reset() {
	*x = KernelDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_kernel_def_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KernelDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KernelDef) ProtoMessage() {}

func (x *KernelDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_kernel_def_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KernelDef.ProtoReflect.Descriptor instead.
func (*KernelDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_kernel_def_proto_rawDescGZIP(), []int{0}
}

func (x *KernelDef) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *KernelDef) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *KernelDef) GetConstraint() []*KernelDef_AttrConstraint {
	if x != nil {
		return x.Constraint
	}
	return nil
}

func (x *KernelDef) GetHostMemoryArg() []string {
	if x != nil {
		return x.HostMemoryArg
	}
	return nil
}

func (x *KernelDef) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *KernelDef) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

// A collection of KernelDefs
type KernelList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kernel []*KernelDef `protobuf:"bytes,1,rep,name=kernel,proto3" json:"kernel,omitempty"`
}

func (x *KernelList) Reset() {
	*x = KernelList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_kernel_def_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KernelList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KernelList) ProtoMessage() {}

func (x *KernelList) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_kernel_def_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KernelList.ProtoReflect.Descriptor instead.
func (*KernelList) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_kernel_def_proto_rawDescGZIP(), []int{1}
}

func (x *KernelList) GetKernel() []*KernelDef {
	if x != nil {
		return x.Kernel
	}
	return nil
}

type KernelDef_AttrConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an attr from the Op.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of values that this kernel supports for this attr.
	// Like OpDef.AttrDef.allowed_values, except for kernels instead of Ops.
	AllowedValues *attr_value_go_proto.AttrValue `protobuf:"bytes,2,opt,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
}

func (x *KernelDef_AttrConstraint) Reset() {
	*x = KernelDef_AttrConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_kernel_def_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KernelDef_AttrConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KernelDef_AttrConstraint) ProtoMessage() {}

func (x *KernelDef_AttrConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_kernel_def_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KernelDef_AttrConstraint.ProtoReflect.Descriptor instead.
func (*KernelDef_AttrConstraint) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_kernel_def_proto_rawDescGZIP(), []int{0, 0}
}

func (x *KernelDef_AttrConstraint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KernelDef_AttrConstraint) GetAllowedValues() *attr_value_go_proto.AttrValue {
	if x != nil {
		return x.AllowedValues
	}
	return nil
}

var File_tensorflow_core_framework_kernel_def_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_kernel_def_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6b, 0x65, 0x72, 0x6e,
	0x65, 0x6c, 0x5f, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x09, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x44,
	0x65, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x6f, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x61, 0x72, 0x67, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x41, 0x72,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x1a, 0x62, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x0a, 0x4b, 0x65, 0x72, 0x6e, 0x65,
	0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x52, 0x06, 0x6b, 0x65,
	0x72, 0x6e, 0x65, 0x6c, 0x42, 0x83, 0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x42, 0x0f, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_kernel_def_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_kernel_def_proto_rawDescData = file_tensorflow_core_framework_kernel_def_proto_rawDesc
)

func file_tensorflow_core_framework_kernel_def_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_kernel_def_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_kernel_def_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_kernel_def_proto_rawDescData)
	})
	return file_tensorflow_core_framework_kernel_def_proto_rawDescData
}

var file_tensorflow_core_framework_kernel_def_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_core_framework_kernel_def_proto_goTypes = []interface{}{
	(*KernelDef)(nil),                     // 0: tensorflow.KernelDef
	(*KernelList)(nil),                    // 1: tensorflow.KernelList
	(*KernelDef_AttrConstraint)(nil),      // 2: tensorflow.KernelDef.AttrConstraint
	(*attr_value_go_proto.AttrValue)(nil), // 3: tensorflow.AttrValue
}
var file_tensorflow_core_framework_kernel_def_proto_depIdxs = []int32{
	2, // 0: tensorflow.KernelDef.constraint:type_name -> tensorflow.KernelDef.AttrConstraint
	0, // 1: tensorflow.KernelList.kernel:type_name -> tensorflow.KernelDef
	3, // 2: tensorflow.KernelDef.AttrConstraint.allowed_values:type_name -> tensorflow.AttrValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_kernel_def_proto_init() }
func file_tensorflow_core_framework_kernel_def_proto_init() {
	if File_tensorflow_core_framework_kernel_def_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_kernel_def_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KernelDef); i {
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
		file_tensorflow_core_framework_kernel_def_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KernelList); i {
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
		file_tensorflow_core_framework_kernel_def_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KernelDef_AttrConstraint); i {
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
			RawDescriptor: file_tensorflow_core_framework_kernel_def_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_kernel_def_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_kernel_def_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_kernel_def_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_kernel_def_proto = out.File
	file_tensorflow_core_framework_kernel_def_proto_rawDesc = nil
	file_tensorflow_core_framework_kernel_def_proto_goTypes = nil
	file_tensorflow_core_framework_kernel_def_proto_depIdxs = nil
}
