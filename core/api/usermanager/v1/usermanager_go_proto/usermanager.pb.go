// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usermanager/v1/usermanager.proto

package usermanager_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	type_go_proto "github.com/google/keytransparency/core/api/type/type_go_proto"
	tink_go_proto "github.com/google/tink/proto/tink_go_proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// GetKeySetsRequest requests the keyset of a domain_id/app_id
type GetKeySetRequest struct {
	// domain_id identifies the domain.
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// app_id identifies the application.
	AppId                string   `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeySetRequest) Reset()         { *m = GetKeySetRequest{} }
func (m *GetKeySetRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeySetRequest) ProtoMessage()    {}
func (*GetKeySetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e7438f46d8eef3c, []int{0}
}

func (m *GetKeySetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeySetRequest.Unmarshal(m, b)
}
func (m *GetKeySetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeySetRequest.Marshal(b, m, deterministic)
}
func (m *GetKeySetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeySetRequest.Merge(m, src)
}
func (m *GetKeySetRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeySetRequest.Size(m)
}
func (m *GetKeySetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeySetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeySetRequest proto.InternalMessageInfo

func (m *GetKeySetRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *GetKeySetRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

// CreateUserRequest specifies the information with which a new user should be initialized.
// New users will be signed with the current active key.
// It is the responsibility of authorized callers to verify that domain_id/app_id/user_id is correct.
type CreateUserRequest struct {
	// user is the user to create.
	User *type_go_proto.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// add_signing_keys specifies whether to add this service's signing keys to the set of authorized_keys.
	// This must be set to true if any further operations from this API are meant to succeed.
	// If set to false, there must be at least one key in authorized_keys.
	AddSigningKeys       bool     `protobuf:"varint,7,opt,name=add_signing_keys,json=addSigningKeys,proto3" json:"add_signing_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e7438f46d8eef3c, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *type_go_proto.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateUserRequest) GetAddSigningKeys() bool {
	if m != nil {
		return m.AddSigningKeys
	}
	return false
}

// UpdateUserRequest sets the data field for the user.
// The user must have the service's current signing key in its list of
// authorized_keys in order to succeed.
type UpdateUserRequest struct {
	// user contains data which will be applied to the user.
	User *type_go_proto.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// update_mask specifies which fields of user to update.
	// For example: "data" or "authorized_keys"
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e7438f46d8eef3c, []int{2}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *type_go_proto.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateUserRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// BatchCreateUserRequest creates multiple users all at once.
type BatchCreateUserRequest struct {
	// domain_id identifies the domain.
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// app_id identifies the application.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// users is the set of users to create.
	Users []*type_go_proto.User `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	// add_signing_keys specifies whether to add this service's signing_keys to the set of authorized_keys.
	// This must be set to true if any further operations from this API are meant to succeed.
	// If set to false, there must be at least one key in authorized_keys.
	AddSigningKeys       bool     `protobuf:"varint,4,opt,name=add_signing_keys,json=addSigningKeys,proto3" json:"add_signing_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchCreateUserRequest) Reset()         { *m = BatchCreateUserRequest{} }
func (m *BatchCreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*BatchCreateUserRequest) ProtoMessage()    {}
func (*BatchCreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e7438f46d8eef3c, []int{3}
}

func (m *BatchCreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchCreateUserRequest.Unmarshal(m, b)
}
func (m *BatchCreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchCreateUserRequest.Marshal(b, m, deterministic)
}
func (m *BatchCreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchCreateUserRequest.Merge(m, src)
}
func (m *BatchCreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_BatchCreateUserRequest.Size(m)
}
func (m *BatchCreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchCreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchCreateUserRequest proto.InternalMessageInfo

func (m *BatchCreateUserRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *BatchCreateUserRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *BatchCreateUserRequest) GetUsers() []*type_go_proto.User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *BatchCreateUserRequest) GetAddSigningKeys() bool {
	if m != nil {
		return m.AddSigningKeys
	}
	return false
}

// BatchCreateUserResponse creates multiple users at once.
type BatchCreateUserResponse struct {
	// users returns the list of created users.
	Users                []*type_go_proto.User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BatchCreateUserResponse) Reset()         { *m = BatchCreateUserResponse{} }
func (m *BatchCreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*BatchCreateUserResponse) ProtoMessage()    {}
func (*BatchCreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e7438f46d8eef3c, []int{4}
}

func (m *BatchCreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchCreateUserResponse.Unmarshal(m, b)
}
func (m *BatchCreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchCreateUserResponse.Marshal(b, m, deterministic)
}
func (m *BatchCreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchCreateUserResponse.Merge(m, src)
}
func (m *BatchCreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_BatchCreateUserResponse.Size(m)
}
func (m *BatchCreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchCreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchCreateUserResponse proto.InternalMessageInfo

func (m *BatchCreateUserResponse) GetUsers() []*type_go_proto.User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*GetKeySetRequest)(nil), "google.keytransparency.usermanager.v1.GetKeySetRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "google.keytransparency.usermanager.v1.CreateUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "google.keytransparency.usermanager.v1.UpdateUserRequest")
	proto.RegisterType((*BatchCreateUserRequest)(nil), "google.keytransparency.usermanager.v1.BatchCreateUserRequest")
	proto.RegisterType((*BatchCreateUserResponse)(nil), "google.keytransparency.usermanager.v1.BatchCreateUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserManagerClient is the client API for UserManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserManagerClient interface {
	// GetKeySet returns a list of public keys (a keyset) that corresponds to the signing keys
	// this service has for a given domain and app.
	GetKeySet(ctx context.Context, in *GetKeySetRequest, opts ...grpc.CallOption) (*tink_go_proto.Keyset, error)
	// CreateUser creates a new user and initializes it.
	// If the user already exists, this operation will fail.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*type_go_proto.User, error)
	// UpdateUserData sets the public key for an user.
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*type_go_proto.User, error)
	// BatchCreateUser creates a set of new users.
	BatchCreateUser(ctx context.Context, in *BatchCreateUserRequest, opts ...grpc.CallOption) (*type_go_proto.User, error)
}

type userManagerClient struct {
	cc *grpc.ClientConn
}

func NewUserManagerClient(cc *grpc.ClientConn) UserManagerClient {
	return &userManagerClient{cc}
}

func (c *userManagerClient) GetKeySet(ctx context.Context, in *GetKeySetRequest, opts ...grpc.CallOption) (*tink_go_proto.Keyset, error) {
	out := new(tink_go_proto.Keyset)
	err := c.cc.Invoke(ctx, "/google.keytransparency.usermanager.v1.UserManager/GetKeySet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*type_go_proto.User, error) {
	out := new(type_go_proto.User)
	err := c.cc.Invoke(ctx, "/google.keytransparency.usermanager.v1.UserManager/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*type_go_proto.User, error) {
	out := new(type_go_proto.User)
	err := c.cc.Invoke(ctx, "/google.keytransparency.usermanager.v1.UserManager/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) BatchCreateUser(ctx context.Context, in *BatchCreateUserRequest, opts ...grpc.CallOption) (*type_go_proto.User, error) {
	out := new(type_go_proto.User)
	err := c.cc.Invoke(ctx, "/google.keytransparency.usermanager.v1.UserManager/BatchCreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagerServer is the server API for UserManager service.
type UserManagerServer interface {
	// GetKeySet returns a list of public keys (a keyset) that corresponds to the signing keys
	// this service has for a given domain and app.
	GetKeySet(context.Context, *GetKeySetRequest) (*tink_go_proto.Keyset, error)
	// CreateUser creates a new user and initializes it.
	// If the user already exists, this operation will fail.
	CreateUser(context.Context, *CreateUserRequest) (*type_go_proto.User, error)
	// UpdateUserData sets the public key for an user.
	UpdateUser(context.Context, *UpdateUserRequest) (*type_go_proto.User, error)
	// BatchCreateUser creates a set of new users.
	BatchCreateUser(context.Context, *BatchCreateUserRequest) (*type_go_proto.User, error)
}

func RegisterUserManagerServer(s *grpc.Server, srv UserManagerServer) {
	s.RegisterService(&_UserManager_serviceDesc, srv)
}

func _UserManager_GetKeySet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeySetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).GetKeySet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.usermanager.v1.UserManager/GetKeySet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).GetKeySet(ctx, req.(*GetKeySetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.usermanager.v1.UserManager/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.usermanager.v1.UserManager/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_BatchCreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).BatchCreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.usermanager.v1.UserManager/BatchCreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).BatchCreateUser(ctx, req.(*BatchCreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.keytransparency.usermanager.v1.UserManager",
	HandlerType: (*UserManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeySet",
			Handler:    _UserManager_GetKeySet_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserManager_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserManager_UpdateUser_Handler,
		},
		{
			MethodName: "BatchCreateUser",
			Handler:    _UserManager_BatchCreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermanager/v1/usermanager.proto",
}

func init() { proto.RegisterFile("usermanager/v1/usermanager.proto", fileDescriptor_8e7438f46d8eef3c) }

var fileDescriptor_8e7438f46d8eef3c = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdf, 0x6a, 0x13, 0x4f,
	0x14, 0x66, 0x93, 0x36, 0xbf, 0x66, 0x02, 0xbf, 0xb6, 0x03, 0x6a, 0x88, 0x5e, 0xc4, 0x05, 0x21,
	0x78, 0x31, 0x43, 0x23, 0xd2, 0x12, 0x29, 0x48, 0x84, 0x6a, 0x29, 0x95, 0x92, 0x52, 0x10, 0x2f,
	0x5c, 0xa6, 0x99, 0xe9, 0x76, 0xd9, 0x66, 0x66, 0xdc, 0x99, 0x2d, 0x2c, 0xa5, 0x37, 0xde, 0x08,
	0xde, 0xfa, 0x06, 0xbe, 0x80, 0x4f, 0xe0, 0x03, 0x78, 0xed, 0x2b, 0xf8, 0x20, 0x32, 0x33, 0x9b,
	0x3f, 0x64, 0x13, 0xdc, 0x88, 0xde, 0x6c, 0x32, 0x67, 0xce, 0x9c, 0xf3, 0x7d, 0xdf, 0x7e, 0x67,
	0x16, 0xb4, 0x53, 0xc5, 0x92, 0x11, 0xe1, 0x24, 0x64, 0x09, 0xbe, 0xde, 0xc1, 0x33, 0x4b, 0x24,
	0x13, 0xa1, 0x05, 0x7c, 0x14, 0x0a, 0x11, 0x5e, 0x31, 0x14, 0xb3, 0x4c, 0x27, 0x84, 0x2b, 0x49,
	0x12, 0xc6, 0x87, 0x19, 0x9a, 0xcd, 0xbc, 0xde, 0x69, 0x3d, 0x70, 0x69, 0x98, 0xc8, 0x08, 0x13,
	0xce, 0x85, 0x26, 0x3a, 0x12, 0x5c, 0xb9, 0x22, 0xad, 0x76, 0xbe, 0x6b, 0x57, 0xe7, 0xe9, 0x05,
	0xbe, 0x88, 0xd8, 0x15, 0x0d, 0x46, 0x44, 0xc5, 0x79, 0xc6, 0xa6, 0xce, 0x24, 0xc3, 0xe6, 0x91,
	0x07, 0x80, 0x8e, 0x78, 0xbe, 0xe9, 0x1f, 0x80, 0xad, 0x97, 0x4c, 0x1f, 0xb1, 0xec, 0x94, 0xe9,
	0x01, 0x7b, 0x9f, 0x32, 0xa5, 0xe1, 0x7d, 0x50, 0xa7, 0x62, 0x44, 0x22, 0x1e, 0x44, 0xb4, 0xe9,
	0xb5, 0xbd, 0x4e, 0x7d, 0xb0, 0xe1, 0x02, 0x87, 0x14, 0xde, 0x01, 0x35, 0x22, 0xa5, 0xd9, 0xa9,
	0xd8, 0x9d, 0x75, 0x22, 0xe5, 0x21, 0xf5, 0x35, 0xd8, 0x7e, 0x91, 0x30, 0xa2, 0xd9, 0x99, 0x62,
	0xc9, 0xb8, 0xd0, 0x53, 0xb0, 0x66, 0xb8, 0xd8, 0xcc, 0x46, 0xf7, 0x21, 0x5a, 0xc2, 0xd7, 0x42,
	0xb3, 0xe7, 0x6c, 0x3a, 0xec, 0x80, 0x2d, 0x42, 0x69, 0xa0, 0xa2, 0x90, 0x47, 0x3c, 0x0c, 0x62,
	0x96, 0xa9, 0xe6, 0x7f, 0x6d, 0xaf, 0xb3, 0x31, 0xf8, 0x9f, 0x50, 0x7a, 0xea, 0xc2, 0x47, 0x2c,
	0x53, 0xfe, 0x47, 0x0f, 0x6c, 0x9f, 0x49, 0xfa, 0x77, 0xda, 0x3e, 0x03, 0x8d, 0xd4, 0xd6, 0xb2,
	0xe2, 0x35, 0xab, 0xf6, 0x74, 0x6b, 0x7c, 0x7a, 0xac, 0x2f, 0x3a, 0x30, 0xfa, 0x1e, 0x13, 0x15,
	0x0f, 0x80, 0x4b, 0x37, 0xff, 0xfd, 0xaf, 0x1e, 0xb8, 0xdb, 0x27, 0x7a, 0x78, 0x59, 0x54, 0xe1,
	0x0f, 0xe4, 0x84, 0xbb, 0x60, 0xdd, 0x60, 0x52, 0xcd, 0x6a, 0xbb, 0x5a, 0x8e, 0x83, 0xcb, 0x5f,
	0xa8, 0xdd, 0xda, 0x42, 0xed, 0x06, 0xe0, 0x5e, 0x01, 0xb0, 0x92, 0x82, 0x2b, 0x36, 0xed, 0x5e,
	0x59, 0xad, 0x7b, 0xf7, 0x53, 0x0d, 0x34, 0xcc, 0xfa, 0xd8, 0xb9, 0x17, 0x7e, 0xf1, 0x40, 0x7d,
	0x62, 0x2f, 0xb8, 0x8b, 0x4a, 0x19, 0x1e, 0xcd, 0x1b, 0xb2, 0x35, 0x79, 0x09, 0xc3, 0x24, 0x93,
	0x5a, 0x20, 0xeb, 0x5f, 0x43, 0x82, 0x69, 0xff, 0xf9, 0x87, 0x1f, 0x3f, 0x3f, 0x57, 0x7a, 0x70,
	0x0f, 0xcf, 0xcd, 0x9b, 0x93, 0x58, 0xe1, 0x9b, 0x89, 0xf8, 0xb7, 0x98, 0x48, 0xa9, 0xf0, 0x8d,
	0x13, 0xfc, 0x16, 0xc7, 0xb6, 0x02, 0xfc, 0xee, 0x01, 0x30, 0x15, 0x01, 0xee, 0x95, 0x44, 0x59,
	0x78, 0xd1, 0xad, 0xdf, 0xeb, 0xe4, 0xbf, 0xb3, 0x68, 0xdf, 0xf8, 0x27, 0x4b, 0xd1, 0x9a, 0x38,
	0x2a, 0x40, 0xb6, 0xd1, 0x31, 0x6e, 0xab, 0x75, 0x1e, 0x33, 0x0f, 0x13, 0xec, 0x39, 0x0f, 0x1b,
	0x2e, 0xd3, 0x81, 0x28, 0xcd, 0xa5, 0x30, 0x43, 0x2b, 0x70, 0xe9, 0xfe, 0x2b, 0x2e, 0xdf, 0x3c,
	0xb0, 0x39, 0xe7, 0x50, 0xb8, 0x5f, 0x92, 0xd0, 0xe2, 0x51, 0x2c, 0xc3, 0xea, 0x95, 0x65, 0xd5,
	0xf7, 0xf7, 0x57, 0xf6, 0x53, 0x6f, 0xa6, 0x69, 0xcf, 0x7b, 0xdc, 0x3f, 0x79, 0xfb, 0x3a, 0x8c,
	0xf4, 0x65, 0x7a, 0x8e, 0x86, 0x62, 0x84, 0xf3, 0x6b, 0x7a, 0xae, 0x31, 0x1e, 0x8a, 0xc4, 0xdd,
	0xec, 0xcb, 0xbf, 0x16, 0x41, 0x28, 0x02, 0x77, 0xfb, 0xd4, 0xec, 0xcf, 0x93, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xb0, 0x65, 0x5d, 0x96, 0x5b, 0x06, 0x00, 0x00,
}