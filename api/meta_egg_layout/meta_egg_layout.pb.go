// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: meta_egg_layout.proto

package meta_egg_layout

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// GetUserDetailRequest 获取用户详情请求
type GetUserDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserDetailRequest) Reset() {
	*x = GetUserDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_egg_layout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailRequest) ProtoMessage() {}

func (x *GetUserDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meta_egg_layout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailRequest.ProtoReflect.Descriptor instead.
func (*GetUserDetailRequest) Descriptor() ([]byte, []int) {
	return file_meta_egg_layout_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserDetailRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// GetUserDetailResponse 获取用户详情响应
type GetUserDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                      //
	Name      *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`                             // 用户名
	Gender    uint64  `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`                              // 性别
	CreatedBy *uint64 `protobuf:"varint,4,opt,name=created_by,json=createdBy,proto3,oneof" json:"created_by,omitempty"` // 创建者
	CreatedAt string  `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`        // 创建时间
	UpdatedBy *uint64 `protobuf:"varint,6,opt,name=updated_by,json=updatedBy,proto3,oneof" json:"updated_by,omitempty"` // 更新者
	UpdatedAt string  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`        // 更新时间
}

func (x *GetUserDetailResponse) Reset() {
	*x = GetUserDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_egg_layout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailResponse) ProtoMessage() {}

func (x *GetUserDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meta_egg_layout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailResponse.ProtoReflect.Descriptor instead.
func (*GetUserDetailResponse) Descriptor() ([]byte, []int) {
	return file_meta_egg_layout_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserDetailResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserDetailResponse) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GetUserDetailResponse) GetGender() uint64 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *GetUserDetailResponse) GetCreatedBy() uint64 {
	if x != nil && x.CreatedBy != nil {
		return *x.CreatedBy
	}
	return 0
}

func (x *GetUserDetailResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetUserDetailResponse) GetUpdatedBy() uint64 {
	if x != nil && x.UpdatedBy != nil {
		return *x.UpdatedBy
	}
	return 0
}

func (x *GetUserDetailResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_meta_egg_layout_proto protoreflect.FileDescriptor

var file_meta_egg_layout_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x65, 0x67, 0x67, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x32, 0x02, 0x20, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x85, 0x02, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x32, 0x6d, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x67, 0x67, 0x4c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x12, 0x5c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61,
	0x79, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x65, 0x67, 0x67, 0x5f,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meta_egg_layout_proto_rawDescOnce sync.Once
	file_meta_egg_layout_proto_rawDescData = file_meta_egg_layout_proto_rawDesc
)

func file_meta_egg_layout_proto_rawDescGZIP() []byte {
	file_meta_egg_layout_proto_rawDescOnce.Do(func() {
		file_meta_egg_layout_proto_rawDescData = protoimpl.X.CompressGZIP(file_meta_egg_layout_proto_rawDescData)
	})
	return file_meta_egg_layout_proto_rawDescData
}

var file_meta_egg_layout_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_meta_egg_layout_proto_goTypes = []interface{}{
	(*GetUserDetailRequest)(nil),  // 0: metaegglayout.GetUserDetailRequest
	(*GetUserDetailResponse)(nil), // 1: metaegglayout.GetUserDetailResponse
}
var file_meta_egg_layout_proto_depIdxs = []int32{
	0, // 0: metaegglayout.MetaEggLayout.GetUserDetail:input_type -> metaegglayout.GetUserDetailRequest
	1, // 1: metaegglayout.MetaEggLayout.GetUserDetail:output_type -> metaegglayout.GetUserDetailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_meta_egg_layout_proto_init() }
func file_meta_egg_layout_proto_init() {
	if File_meta_egg_layout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meta_egg_layout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDetailRequest); i {
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
		file_meta_egg_layout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDetailResponse); i {
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
	file_meta_egg_layout_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meta_egg_layout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_meta_egg_layout_proto_goTypes,
		DependencyIndexes: file_meta_egg_layout_proto_depIdxs,
		MessageInfos:      file_meta_egg_layout_proto_msgTypes,
	}.Build()
	File_meta_egg_layout_proto = out.File
	file_meta_egg_layout_proto_rawDesc = nil
	file_meta_egg_layout_proto_goTypes = nil
	file_meta_egg_layout_proto_depIdxs = nil
}
