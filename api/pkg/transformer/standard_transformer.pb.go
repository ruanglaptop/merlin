// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/transformer/standard_transformer.proto

package transformer

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

type StandardTransformerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransformerConfig *TransformerConfig `protobuf:"bytes,1,opt,name=transformerConfig,proto3" json:"transformerConfig,omitempty"`
}

func (x *StandardTransformerConfig) Reset() {
	*x = StandardTransformerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transformer_standard_transformer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandardTransformerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandardTransformerConfig) ProtoMessage() {}

func (x *StandardTransformerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transformer_standard_transformer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandardTransformerConfig.ProtoReflect.Descriptor instead.
func (*StandardTransformerConfig) Descriptor() ([]byte, []int) {
	return file_pkg_transformer_standard_transformer_proto_rawDescGZIP(), []int{0}
}

func (x *StandardTransformerConfig) GetTransformerConfig() *TransformerConfig {
	if x != nil {
		return x.TransformerConfig
	}
	return nil
}

type TransformerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feast []*FeatureTable `protobuf:"bytes,1,rep,name=feast,proto3" json:"feast,omitempty"`
}

func (x *TransformerConfig) Reset() {
	*x = TransformerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transformer_standard_transformer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformerConfig) ProtoMessage() {}

func (x *TransformerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transformer_standard_transformer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformerConfig.ProtoReflect.Descriptor instead.
func (*TransformerConfig) Descriptor() ([]byte, []int) {
	return file_pkg_transformer_standard_transformer_proto_rawDescGZIP(), []int{1}
}

func (x *TransformerConfig) GetFeast() []*FeatureTable {
	if x != nil {
		return x.Feast
	}
	return nil
}

type FeatureTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project  string     `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Entities []*Entity  `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
	Features []*Feature `protobuf:"bytes,3,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *FeatureTable) Reset() {
	*x = FeatureTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transformer_standard_transformer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureTable) ProtoMessage() {}

func (x *FeatureTable) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transformer_standard_transformer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureTable.ProtoReflect.Descriptor instead.
func (*FeatureTable) Descriptor() ([]byte, []int) {
	return file_pkg_transformer_standard_transformer_proto_rawDescGZIP(), []int{2}
}

func (x *FeatureTable) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *FeatureTable) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *FeatureTable) GetFeatures() []*Feature {
	if x != nil {
		return x.Features
	}
	return nil
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ValueType string `protobuf:"bytes,2,opt,name=valueType,proto3" json:"valueType,omitempty"`
	JsonPath  string `protobuf:"bytes,3,opt,name=jsonPath,proto3" json:"jsonPath,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transformer_standard_transformer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transformer_standard_transformer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_pkg_transformer_standard_transformer_proto_rawDescGZIP(), []int{3}
}

func (x *Entity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Entity) GetValueType() string {
	if x != nil {
		return x.ValueType
	}
	return ""
}

func (x *Entity) GetJsonPath() string {
	if x != nil {
		return x.JsonPath
	}
	return ""
}

type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ValueType    string `protobuf:"bytes,2,opt,name=valueType,proto3" json:"valueType,omitempty"`
	DefaultValue string `protobuf:"bytes,3,opt,name=defaultValue,proto3" json:"defaultValue,omitempty"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_transformer_standard_transformer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_transformer_standard_transformer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_pkg_transformer_standard_transformer_proto_rawDescGZIP(), []int{4}
}

func (x *Feature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Feature) GetValueType() string {
	if x != nil {
		return x.ValueType
	}
	return ""
}

func (x *Feature) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

var File_pkg_transformer_standard_transformer_proto protoreflect.FileDescriptor

var file_pkg_transformer_standard_transformer_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6d, 0x65,
	0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72,
	0x22, 0x70, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x53, 0x0a,
	0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69,
	0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x4b, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x36, 0x0a, 0x05, 0x66, 0x65, 0x61, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x65, 0x61, 0x73, 0x74, 0x22,
	0x99, 0x01, 0x0a, 0x0c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d,
	0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x56, 0x0a, 0x06, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50,
	0x61, 0x74, 0x68, 0x22, 0x5f, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6a, 0x65, 0x6b, 0x2f, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_transformer_standard_transformer_proto_rawDescOnce sync.Once
	file_pkg_transformer_standard_transformer_proto_rawDescData = file_pkg_transformer_standard_transformer_proto_rawDesc
)

func file_pkg_transformer_standard_transformer_proto_rawDescGZIP() []byte {
	file_pkg_transformer_standard_transformer_proto_rawDescOnce.Do(func() {
		file_pkg_transformer_standard_transformer_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_transformer_standard_transformer_proto_rawDescData)
	})
	return file_pkg_transformer_standard_transformer_proto_rawDescData
}

var file_pkg_transformer_standard_transformer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_transformer_standard_transformer_proto_goTypes = []interface{}{
	(*StandardTransformerConfig)(nil), // 0: merlin.transformer.StandardTransformerConfig
	(*TransformerConfig)(nil),         // 1: merlin.transformer.TransformerConfig
	(*FeatureTable)(nil),              // 2: merlin.transformer.FeatureTable
	(*Entity)(nil),                    // 3: merlin.transformer.Entity
	(*Feature)(nil),                   // 4: merlin.transformer.Feature
}
var file_pkg_transformer_standard_transformer_proto_depIdxs = []int32{
	1, // 0: merlin.transformer.StandardTransformerConfig.transformerConfig:type_name -> merlin.transformer.TransformerConfig
	2, // 1: merlin.transformer.TransformerConfig.feast:type_name -> merlin.transformer.FeatureTable
	3, // 2: merlin.transformer.FeatureTable.entities:type_name -> merlin.transformer.Entity
	4, // 3: merlin.transformer.FeatureTable.features:type_name -> merlin.transformer.Feature
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_transformer_standard_transformer_proto_init() }
func file_pkg_transformer_standard_transformer_proto_init() {
	if File_pkg_transformer_standard_transformer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_transformer_standard_transformer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandardTransformerConfig); i {
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
		file_pkg_transformer_standard_transformer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformerConfig); i {
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
		file_pkg_transformer_standard_transformer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureTable); i {
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
		file_pkg_transformer_standard_transformer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_pkg_transformer_standard_transformer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
			RawDescriptor: file_pkg_transformer_standard_transformer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_transformer_standard_transformer_proto_goTypes,
		DependencyIndexes: file_pkg_transformer_standard_transformer_proto_depIdxs,
		MessageInfos:      file_pkg_transformer_standard_transformer_proto_msgTypes,
	}.Build()
	File_pkg_transformer_standard_transformer_proto = out.File
	file_pkg_transformer_standard_transformer_proto_rawDesc = nil
	file_pkg_transformer_standard_transformer_proto_goTypes = nil
	file_pkg_transformer_standard_transformer_proto_depIdxs = nil
}
