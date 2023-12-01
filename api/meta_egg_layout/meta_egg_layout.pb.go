// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.2
// source: meta_egg_layout.proto

package meta_egg_layout

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 分页请求
type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_egg_layout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_meta_egg_layout_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// UserDetail 用户详情
type UserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                      //
	Name      *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`                             // 用户名
	Gender    uint64  `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`                              // 性别
	Age       uint32  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`                                    // 年龄
	IsOnJob   bool    `protobuf:"varint,5,opt,name=is_on_job,json=isOnJob,proto3" json:"is_on_job,omitempty"`           // 是否在职
	Birthday  *string `protobuf:"bytes,6,opt,name=birthday,proto3,oneof" json:"birthday,omitempty"`                     // 生日
	CreatedBy *uint64 `protobuf:"varint,7,opt,name=created_by,json=createdBy,proto3,oneof" json:"created_by,omitempty"` // 创建者
	CreatedAt string  `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`        // 创建时间
	UpdatedBy *uint64 `protobuf:"varint,9,opt,name=updated_by,json=updatedBy,proto3,oneof" json:"updated_by,omitempty"` // 更新者
	UpdatedAt string  `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`       // 更新时间
}

func (x *UserDetail) Reset() {
	*x = UserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_egg_layout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetail) ProtoMessage() {}

func (x *UserDetail) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserDetail.ProtoReflect.Descriptor instead.
func (*UserDetail) Descriptor() ([]byte, []int) {
	return file_meta_egg_layout_proto_rawDescGZIP(), []int{1}
}

func (x *UserDetail) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserDetail) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UserDetail) GetGender() uint64 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserDetail) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserDetail) GetIsOnJob() bool {
	if x != nil {
		return x.IsOnJob
	}
	return false
}

func (x *UserDetail) GetBirthday() string {
	if x != nil && x.Birthday != nil {
		return *x.Birthday
	}
	return ""
}

func (x *UserDetail) GetCreatedBy() uint64 {
	if x != nil && x.CreatedBy != nil {
		return *x.CreatedBy
	}
	return 0
}

func (x *UserDetail) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserDetail) GetUpdatedBy() uint64 {
	if x != nil && x.UpdatedBy != nil {
		return *x.UpdatedBy
	}
	return 0
}

func (x *UserDetail) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名
	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// 性别
	Gender uint64 `protobuf:"varint,2,opt,name=gender,proto3" json:"gender,omitempty"`
	// 年龄
	Age uint32 `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	// 是否在职
	IsOnJob bool `protobuf:"varint,4,opt,name=is_on_job,json=isOnJob,proto3" json:"is_on_job,omitempty"`
	// 生日 格式: 2006-01-02
	Birthday *string `protobuf:"bytes,5,opt,name=birthday,proto3,oneof" json:"birthday,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_egg_layout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meta_egg_layout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_meta_egg_layout_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetGender() uint64 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *CreateUserRequest) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreateUserRequest) GetIsOnJob() bool {
	if x != nil {
		return x.IsOnJob
	}
	return false
}

func (x *CreateUserRequest) GetBirthday() string {
	if x != nil && x.Birthday != nil {
		return *x.Birthday
	}
	return ""
}

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
		mi := &file_meta_egg_layout_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailRequest) ProtoMessage() {}

func (x *GetUserDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meta_egg_layout_proto_msgTypes[3]
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
	return file_meta_egg_layout_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserDetailRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// UserListInfo 用户列表信息
type UserListInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`          //
	Name   *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"` // 用户名
	Gender uint64  `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`  // 性别
}

func (x *UserListInfo) Reset() {
	*x = UserListInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_egg_layout_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListInfo) ProtoMessage() {}

func (x *UserListInfo) ProtoReflect() protoreflect.Message {
	mi := &file_meta_egg_layout_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListInfo.ProtoReflect.Descriptor instead.
func (*UserListInfo) Descriptor() ([]byte, []int) {
	return file_meta_egg_layout_proto_rawDescGZIP(), []int{4}
}

func (x *UserListInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserListInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UserListInfo) GetGender() uint64 {
	if x != nil {
		return x.Gender
	}
	return 0
}

// GetUserListRequest 获取用户列表请求
type GetUserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页请求（可选, 不传则不分页）
	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3,oneof" json:"pagination,omitempty"`
	// 筛选项: 性别 (可选)
	Gender *uint64 `protobuf:"varint,2,opt,name=gender,proto3,oneof" json:"gender,omitempty"`
	// 筛选项: 是否在职 (可选)
	IsOnJob *bool `protobuf:"varint,3,opt,name=is_on_job,json=isOnJob,proto3,oneof" json:"is_on_job,omitempty"`
	// 排序字段, 可选: id
	OrderBy *string `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3,oneof" json:"order_by,omitempty"`
	// 排序方式, 默认 desc, 可选: asc, desc
	OrderType *string `protobuf:"bytes,5,opt,name=order_type,json=orderType,proto3,oneof" json:"order_type,omitempty"`
}

func (x *GetUserListRequest) Reset() {
	*x = GetUserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_egg_layout_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListRequest) ProtoMessage() {}

func (x *GetUserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meta_egg_layout_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListRequest.ProtoReflect.Descriptor instead.
func (*GetUserListRequest) Descriptor() ([]byte, []int) {
	return file_meta_egg_layout_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserListRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *GetUserListRequest) GetGender() uint64 {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return 0
}

func (x *GetUserListRequest) GetIsOnJob() bool {
	if x != nil && x.IsOnJob != nil {
		return *x.IsOnJob
	}
	return false
}

func (x *GetUserListRequest) GetOrderBy() string {
	if x != nil && x.OrderBy != nil {
		return *x.OrderBy
	}
	return ""
}

func (x *GetUserListRequest) GetOrderType() string {
	if x != nil && x.OrderType != nil {
		return *x.OrderType
	}
	return ""
}

// GetUserListResponse 获取用户列表响应
type GetUserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*UserListInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`    // 列表数据
	Total int64           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"` // 结果集总数
}

func (x *GetUserListResponse) Reset() {
	*x = GetUserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_egg_layout_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListResponse) ProtoMessage() {}

func (x *GetUserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meta_egg_layout_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListResponse.ProtoReflect.Descriptor instead.
func (*GetUserListResponse) Descriptor() ([]byte, []int) {
	return file_meta_egg_layout_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserListResponse) GetList() []*UserListInfo {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *GetUserListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 更新项: 用户名 (可选)
	Name *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// 更新项: 性别 (可选)
	Gender *uint64 `protobuf:"varint,3,opt,name=gender,proto3,oneof" json:"gender,omitempty"`
	// 更新项: 年龄 (可选)
	Age *uint32 `protobuf:"varint,4,opt,name=age,proto3,oneof" json:"age,omitempty"`
	// 更新项: 是否在职 (可选)
	IsOnJob *bool `protobuf:"varint,5,opt,name=is_on_job,json=isOnJob,proto3,oneof" json:"is_on_job,omitempty"`
	// 更新项: 生日 格式: 2006-01-02 (可选)
	Birthday *string `protobuf:"bytes,6,opt,name=birthday,proto3,oneof" json:"birthday,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_egg_layout_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meta_egg_layout_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_meta_egg_layout_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateUserRequest) GetGender() uint64 {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return 0
}

func (x *UpdateUserRequest) GetAge() uint32 {
	if x != nil && x.Age != nil {
		return *x.Age
	}
	return 0
}

func (x *UpdateUserRequest) GetIsOnJob() bool {
	if x != nil && x.IsOnJob != nil {
		return *x.IsOnJob
	}
	return false
}

func (x *UpdateUserRequest) GetBirthday() string {
	if x != nil && x.Birthday != nil {
		return *x.Birthday
	}
	return ""
}

// DeleteUserRequest 删除用户请求
type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_egg_layout_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meta_egg_layout_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_meta_egg_layout_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteUserRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_meta_egg_layout_proto protoreflect.FileDescriptor

var file_meta_egg_layout_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x65, 0x67, 0x67, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x26, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x10, 0x64, 0x20, 0x00, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xd6, 0x02, 0x0a, 0x0a, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x6f, 0x6e, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x4f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x1f, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x48,
	0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x40, 0x48,
	0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x28, 0x01, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x4f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x3b, 0x0a, 0x08, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xfa, 0x42, 0x17,
	0x72, 0x15, 0x32, 0x13, 0x5e, 0x5c, 0x64, 0x7b, 0x34, 0x7d, 0x2d, 0x5c, 0x64, 0x7b, 0x32, 0x7d,
	0x2d, 0x5c, 0x64, 0x7b, 0x32, 0x7d, 0x24, 0x48, 0x01, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x22, 0x2f, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x58, 0x0a,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc0, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x24,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x48, 0x01, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x5f, 0x6a, 0x6f,
	0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x07, 0x69, 0x73, 0x4f, 0x6e, 0x4a,
	0x6f, 0x62, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x48, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x34, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0x72, 0x0b, 0x52, 0x03, 0x61, 0x73, 0x63,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x48, 0x04, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x5f, 0x6a, 0x6f, 0x62, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5c, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xa0, 0x02, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x40, 0x48, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x28, 0x01, 0x48, 0x01, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12,
	0x15, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x5f,
	0x6a, 0x6f, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x07, 0x69, 0x73, 0x4f,
	0x6e, 0x4a, 0x6f, 0x62, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xfa, 0x42, 0x17, 0x72, 0x15,
	0x32, 0x13, 0x5e, 0x5c, 0x64, 0x7b, 0x34, 0x7d, 0x2d, 0x5c, 0x64, 0x7b, 0x32, 0x7d, 0x2d, 0x5c,
	0x64, 0x7b, 0x32, 0x7d, 0x24, 0x48, 0x04, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61,
	0x79, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x61, 0x67, 0x65,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x5f, 0x6a, 0x6f, 0x62, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x22, 0x2c, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x32, 0x9b, 0x03, 0x0a, 0x0d, 0x4d, 0x65,
	0x74, 0x61, 0x45, 0x67, 0x67, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x4b, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x65, 0x67, 0x67, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x5f, 0x65, 0x67, 0x67, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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

var file_meta_egg_layout_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_meta_egg_layout_proto_goTypes = []interface{}{
	(*Pagination)(nil),           // 0: metaegglayout.Pagination
	(*UserDetail)(nil),           // 1: metaegglayout.UserDetail
	(*CreateUserRequest)(nil),    // 2: metaegglayout.CreateUserRequest
	(*GetUserDetailRequest)(nil), // 3: metaegglayout.GetUserDetailRequest
	(*UserListInfo)(nil),         // 4: metaegglayout.UserListInfo
	(*GetUserListRequest)(nil),   // 5: metaegglayout.GetUserListRequest
	(*GetUserListResponse)(nil),  // 6: metaegglayout.GetUserListResponse
	(*UpdateUserRequest)(nil),    // 7: metaegglayout.UpdateUserRequest
	(*DeleteUserRequest)(nil),    // 8: metaegglayout.DeleteUserRequest
	(*emptypb.Empty)(nil),        // 9: google.protobuf.Empty
}
var file_meta_egg_layout_proto_depIdxs = []int32{
	0, // 0: metaegglayout.GetUserListRequest.pagination:type_name -> metaegglayout.Pagination
	4, // 1: metaegglayout.GetUserListResponse.list:type_name -> metaegglayout.UserListInfo
	2, // 2: metaegglayout.MetaEggLayout.CreateUser:input_type -> metaegglayout.CreateUserRequest
	3, // 3: metaegglayout.MetaEggLayout.GetUserDetail:input_type -> metaegglayout.GetUserDetailRequest
	5, // 4: metaegglayout.MetaEggLayout.GetUserList:input_type -> metaegglayout.GetUserListRequest
	7, // 5: metaegglayout.MetaEggLayout.UpdateUser:input_type -> metaegglayout.UpdateUserRequest
	8, // 6: metaegglayout.MetaEggLayout.DeleteUser:input_type -> metaegglayout.DeleteUserRequest
	1, // 7: metaegglayout.MetaEggLayout.CreateUser:output_type -> metaegglayout.UserDetail
	1, // 8: metaegglayout.MetaEggLayout.GetUserDetail:output_type -> metaegglayout.UserDetail
	6, // 9: metaegglayout.MetaEggLayout.GetUserList:output_type -> metaegglayout.GetUserListResponse
	9, // 10: metaegglayout.MetaEggLayout.UpdateUser:output_type -> google.protobuf.Empty
	9, // 11: metaegglayout.MetaEggLayout.DeleteUser:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_meta_egg_layout_proto_init() }
func file_meta_egg_layout_proto_init() {
	if File_meta_egg_layout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meta_egg_layout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
			switch v := v.(*UserDetail); i {
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
		file_meta_egg_layout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_meta_egg_layout_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_meta_egg_layout_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListInfo); i {
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
		file_meta_egg_layout_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListRequest); i {
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
		file_meta_egg_layout_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListResponse); i {
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
		file_meta_egg_layout_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_meta_egg_layout_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
	file_meta_egg_layout_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_meta_egg_layout_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_meta_egg_layout_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_meta_egg_layout_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meta_egg_layout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
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
