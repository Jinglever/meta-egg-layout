// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.22.2
// source: meta_egg_layout.proto

package meta_egg_layout

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_meta_egg_layout_proto protoreflect.FileDescriptor

var file_meta_egg_layout_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x65, 0x67, 0x67, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x9b, 0x03, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x67, 0x67, 0x4c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x12, 0x4b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x20, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x00, 0x12, 0x51,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x23, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61,
	0x79, 0x6f, 0x75, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22,
	0x00, 0x12, 0x56, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x21, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67,
	0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x13, 0x5a,
	0x11, 0x2e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x65, 0x67, 0x67, 0x5f, 0x6c, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_meta_egg_layout_proto_goTypes = []any{
	(*CreateUserRequest)(nil),    // 0: metaegglayout.CreateUserRequest
	(*GetUserDetailRequest)(nil), // 1: metaegglayout.GetUserDetailRequest
	(*GetUserListRequest)(nil),   // 2: metaegglayout.GetUserListRequest
	(*UpdateUserRequest)(nil),    // 3: metaegglayout.UpdateUserRequest
	(*DeleteUserRequest)(nil),    // 4: metaegglayout.DeleteUserRequest
	(*UserDetail)(nil),           // 5: metaegglayout.UserDetail
	(*GetUserListResponse)(nil),  // 6: metaegglayout.GetUserListResponse
	(*emptypb.Empty)(nil),        // 7: google.protobuf.Empty
}
var file_meta_egg_layout_proto_depIdxs = []int32{
	0, // 0: metaegglayout.MetaEggLayout.CreateUser:input_type -> metaegglayout.CreateUserRequest
	1, // 1: metaegglayout.MetaEggLayout.GetUserDetail:input_type -> metaegglayout.GetUserDetailRequest
	2, // 2: metaegglayout.MetaEggLayout.GetUserList:input_type -> metaegglayout.GetUserListRequest
	3, // 3: metaegglayout.MetaEggLayout.UpdateUser:input_type -> metaegglayout.UpdateUserRequest
	4, // 4: metaegglayout.MetaEggLayout.DeleteUser:input_type -> metaegglayout.DeleteUserRequest
	5, // 5: metaegglayout.MetaEggLayout.CreateUser:output_type -> metaegglayout.UserDetail
	5, // 6: metaegglayout.MetaEggLayout.GetUserDetail:output_type -> metaegglayout.UserDetail
	6, // 7: metaegglayout.MetaEggLayout.GetUserList:output_type -> metaegglayout.GetUserListResponse
	7, // 8: metaegglayout.MetaEggLayout.UpdateUser:output_type -> google.protobuf.Empty
	7, // 9: metaegglayout.MetaEggLayout.DeleteUser:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_meta_egg_layout_proto_init() }
func file_meta_egg_layout_proto_init() {
	if File_meta_egg_layout_proto != nil {
		return
	}
	file_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meta_egg_layout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_meta_egg_layout_proto_goTypes,
		DependencyIndexes: file_meta_egg_layout_proto_depIdxs,
	}.Build()
	File_meta_egg_layout_proto = out.File
	file_meta_egg_layout_proto_rawDesc = nil
	file_meta_egg_layout_proto_goTypes = nil
	file_meta_egg_layout_proto_depIdxs = nil
}
