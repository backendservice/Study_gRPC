// Code generated by protoc-gen-go. DO NOT EDIT.
// source: main_service.proto

package grpcService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type GetUserRequest struct {
	Serial               int32    `protobuf:"varint,1,opt,name=Serial,proto3" json:"Serial,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_main_service_e21156c859426789, []int{0}
}
func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(dst, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetSerial() int32 {
	if m != nil {
		return m.Serial
	}
	return 0
}

type GetUserReplay struct {
	UserId               string   `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserReplay) Reset()         { *m = GetUserReplay{} }
func (m *GetUserReplay) String() string { return proto.CompactTextString(m) }
func (*GetUserReplay) ProtoMessage()    {}
func (*GetUserReplay) Descriptor() ([]byte, []int) {
	return fileDescriptor_main_service_e21156c859426789, []int{1}
}
func (m *GetUserReplay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserReplay.Unmarshal(m, b)
}
func (m *GetUserReplay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserReplay.Marshal(b, m, deterministic)
}
func (dst *GetUserReplay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserReplay.Merge(dst, src)
}
func (m *GetUserReplay) XXX_Size() int {
	return xxx_messageInfo_GetUserReplay.Size(m)
}
func (m *GetUserReplay) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserReplay.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserReplay proto.InternalMessageInfo

func (m *GetUserReplay) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "grpcService.GetUserRequest")
	proto.RegisterType((*GetUserReplay)(nil), "grpcService.GetUserReplay")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataEngineClient is the client API for DataEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataEngineClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReplay, error)
}

type dataEngineClient struct {
	cc *grpc.ClientConn
}

func NewDataEngineClient(cc *grpc.ClientConn) DataEngineClient {
	return &dataEngineClient{cc}
}

func (c *dataEngineClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReplay, error) {
	out := new(GetUserReplay)
	err := c.cc.Invoke(ctx, "/grpcService.DataEngine/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataEngineServer is the server API for DataEngine service.
type DataEngineServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserReplay, error)
}

func RegisterDataEngineServer(s *grpc.Server, srv DataEngineServer) {
	s.RegisterService(&_DataEngine_serviceDesc, srv)
}

func _DataEngine_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataEngineServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcService.DataEngine/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataEngineServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataEngine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcService.DataEngine",
	HandlerType: (*DataEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _DataEngine_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main_service.proto",
}

func init() { proto.RegisterFile("main_service.proto", fileDescriptor_main_service_e21156c859426789) }

var fileDescriptor_main_service_e21156c859426789 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0x4d, 0xcc, 0xcc,
	0x8b, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x4e, 0x2f, 0x2a, 0x48, 0x0e, 0x86, 0x08, 0x29, 0x69, 0x70, 0xf1, 0xb9, 0xa7, 0x96, 0x84, 0x16,
	0xa7, 0x16, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x71, 0xb1, 0x05, 0xa7, 0x16,
	0x65, 0x26, 0xe6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x79, 0x4a, 0xea, 0x5c, 0xbc,
	0x70, 0x95, 0x05, 0x39, 0x89, 0x95, 0x20, 0x85, 0x20, 0x9e, 0x67, 0x0a, 0x58, 0x21, 0x67, 0x10,
	0x94, 0x67, 0x14, 0xc4, 0xc5, 0xe5, 0x92, 0x58, 0x92, 0xe8, 0x9a, 0x97, 0x9e, 0x99, 0x97, 0x2a,
	0xe4, 0xc2, 0xc5, 0x0e, 0xd5, 0x26, 0x24, 0xad, 0x87, 0x64, 0xb3, 0x1e, 0xaa, 0xb5, 0x52, 0x52,
	0xd8, 0x25, 0x41, 0x36, 0x29, 0x31, 0x24, 0xb1, 0x81, 0x9d, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x99, 0x81, 0xf8, 0x28, 0xd0, 0x00, 0x00, 0x00,
}
