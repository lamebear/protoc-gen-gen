// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package example // import "github.com/Nais777/gserve/example"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type ID struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_c6a4e847c6ebc6f5, []int{0}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (dst *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(dst, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type User struct {
	ID                   *ID      `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=LastName" json:"LastName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_c6a4e847c6ebc6f5, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() *ID {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func init() {
	proto.RegisterType((*ID)(nil), "example.ID")
	proto.RegisterType((*User)(nil), "example.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Users service

type UsersClient interface {
	User(ctx context.Context, in *ID, opts ...grpc.CallOption) (*User, error)
	Users(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Users_UsersClient, error)
	BatchUpdate(ctx context.Context, opts ...grpc.CallOption) (Users_BatchUpdateClient, error)
	BatchDelete(ctx context.Context, opts ...grpc.CallOption) (Users_BatchDeleteClient, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) User(ctx context.Context, in *ID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/example.Users/User", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Users(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Users_UsersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Users_serviceDesc.Streams[0], c.cc, "/example.Users/Users", opts...)
	if err != nil {
		return nil, err
	}
	x := &usersUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Users_UsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type usersUsersClient struct {
	grpc.ClientStream
}

func (x *usersUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *usersClient) BatchUpdate(ctx context.Context, opts ...grpc.CallOption) (Users_BatchUpdateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Users_serviceDesc.Streams[1], c.cc, "/example.Users/BatchUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &usersBatchUpdateClient{stream}
	return x, nil
}

type Users_BatchUpdateClient interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ClientStream
}

type usersBatchUpdateClient struct {
	grpc.ClientStream
}

func (x *usersBatchUpdateClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *usersBatchUpdateClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *usersClient) BatchDelete(ctx context.Context, opts ...grpc.CallOption) (Users_BatchDeleteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Users_serviceDesc.Streams[2], c.cc, "/example.Users/BatchDelete", opts...)
	if err != nil {
		return nil, err
	}
	x := &usersBatchDeleteClient{stream}
	return x, nil
}

type Users_BatchDeleteClient interface {
	Send(*ID) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type usersBatchDeleteClient struct {
	grpc.ClientStream
}

func (x *usersBatchDeleteClient) Send(m *ID) error {
	return x.ClientStream.SendMsg(m)
}

func (x *usersBatchDeleteClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Users service

type UsersServer interface {
	User(context.Context, *ID) (*User, error)
	Users(*empty.Empty, Users_UsersServer) error
	BatchUpdate(Users_BatchUpdateServer) error
	BatchDelete(Users_BatchDeleteServer) error
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Users/User",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).User(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Users_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UsersServer).Users(m, &usersUsersServer{stream})
}

type Users_UsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type usersUsersServer struct {
	grpc.ServerStream
}

func (x *usersUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _Users_BatchUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UsersServer).BatchUpdate(&usersBatchUpdateServer{stream})
}

type Users_BatchUpdateServer interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ServerStream
}

type usersBatchUpdateServer struct {
	grpc.ServerStream
}

func (x *usersBatchUpdateServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *usersBatchUpdateServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Users_BatchDelete_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UsersServer).BatchDelete(&usersBatchDeleteServer{stream})
}

type Users_BatchDeleteServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*ID, error)
	grpc.ServerStream
}

type usersBatchDeleteServer struct {
	grpc.ServerStream
}

func (x *usersBatchDeleteServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *usersBatchDeleteServer) Recv() (*ID, error) {
	m := new(ID)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "User",
			Handler:    _Users_User_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Users",
			Handler:       _Users_Users_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BatchUpdate",
			Handler:       _Users_BatchUpdate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "BatchDelete",
			Handler:       _Users_BatchDelete_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "example.proto",
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_example_c6a4e847c6ebc6f5) }

var fileDescriptor_example_c6a4e847c6ebc6f5 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x4f, 0x4f, 0x4b, 0xc3, 0x30,
	0x14, 0x37, 0xd5, 0xa9, 0x4b, 0x99, 0x87, 0x87, 0xc8, 0xe8, 0x3c, 0xcc, 0xea, 0xa1, 0xa7, 0x64,
	0x56, 0xb0, 0xf7, 0x51, 0x85, 0x82, 0xec, 0x30, 0xd8, 0x45, 0xf0, 0x90, 0xce, 0x67, 0x57, 0x68,
	0x49, 0x69, 0x32, 0xd1, 0xef, 0xe8, 0x87, 0x92, 0x74, 0x69, 0x75, 0x94, 0x5d, 0x02, 0xbf, 0xbf,
	0x2f, 0x3f, 0x3a, 0xc2, 0x2f, 0x51, 0x56, 0x05, 0xb2, 0xaa, 0x96, 0x5a, 0xc2, 0x99, 0x85, 0xde,
	0x24, 0x93, 0x32, 0x2b, 0x90, 0x37, 0x74, 0xba, 0xfd, 0xe0, 0x58, 0x56, 0xfa, 0x7b, 0xe7, 0xf2,
	0x2f, 0xa9, 0x93, 0xc4, 0x70, 0x61, 0xde, 0x31, 0x99, 0x92, 0x60, 0xb0, 0x74, 0x92, 0xd8, 0x7f,
	0xa3, 0x27, 0x2b, 0x85, 0x35, 0x4c, 0x3a, 0xde, 0x0d, 0x5d, 0xd6, 0xf6, 0x27, 0xb1, 0x31, 0xc1,
	0x35, 0x1d, 0x3e, 0xe7, 0xb5, 0xd2, 0x0b, 0x51, 0xe2, 0xd8, 0x99, 0x92, 0x60, 0xb8, 0xfc, 0x23,
	0xc0, 0xa3, 0xe7, 0x2f, 0xc2, 0x8a, 0xc7, 0x8d, 0xd8, 0xe1, 0xf0, 0x87, 0xd0, 0x81, 0xe9, 0x57,
	0x70, 0x67, 0x0f, 0xfd, 0x2f, 0xf7, 0x46, 0x1d, 0x30, 0x9a, 0x7f, 0x04, 0x61, 0x6b, 0xbf, 0x62,
	0xbb, 0x2d, 0xac, 0xdd, 0xc2, 0x9e, 0xcc, 0x96, 0x5e, 0x62, 0x46, 0xe0, 0x9e, 0xba, 0x73, 0xa1,
	0xd7, 0x9b, 0x55, 0xf5, 0x2e, 0x34, 0xc2, 0xbe, 0xa3, 0x17, 0x08, 0xc8, 0x8c, 0xc0, 0xa3, 0x8d,
	0xc4, 0x58, 0xa0, 0xc6, 0xfd, 0x3f, 0x1d, 0xb8, 0x6c, 0x92, 0xf3, 0xdb, 0xd7, 0x9b, 0x2c, 0xd7,
	0x9b, 0x6d, 0xca, 0xd6, 0xb2, 0xe4, 0x0b, 0x91, 0xab, 0x28, 0x8a, 0x78, 0xa6, 0xb0, 0xfe, 0x44,
	0x6e, 0x3b, 0xd2, 0xd3, 0x26, 0xf8, 0xf0, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xba, 0x81, 0x3a,
	0xa6, 0x01, 0x00, 0x00,
}
