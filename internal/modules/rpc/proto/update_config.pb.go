// Code generated by protoc-gen-go. DO NOT EDIT.
// source: update_config.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ConfigRequest struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Command              string   `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c46896a2f144953b, []int{0}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ConfigRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

type ConfigResponse struct {
	Res                  string   `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigResponse) Reset()         { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()    {}
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c46896a2f144953b, []int{1}
}

func (m *ConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigResponse.Unmarshal(m, b)
}
func (m *ConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigResponse.Marshal(b, m, deterministic)
}
func (m *ConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigResponse.Merge(m, src)
}
func (m *ConfigResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigResponse.Size(m)
}
func (m *ConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigResponse proto.InternalMessageInfo

func (m *ConfigResponse) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "rpc.ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "rpc.ConfigResponse")
}

func init() { proto.RegisterFile("update_config.proto", fileDescriptor_c46896a2f144953b) }

var fileDescriptor_c46896a2f144953b = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2d, 0x48, 0x49,
	0x2c, 0x49, 0x8d, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2e, 0x2a, 0x48, 0x56, 0xb2, 0xe4, 0xe2, 0x75, 0x06, 0x0b, 0x06, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x08, 0xf1, 0x71, 0x31, 0x65, 0x16, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31,
	0x65, 0x16, 0x08, 0x49, 0x70, 0xb1, 0x27, 0xe7, 0xe7, 0xe6, 0x26, 0xe6, 0xa5, 0x48, 0x30, 0x81,
	0x05, 0x61, 0x5c, 0x25, 0x25, 0x2e, 0x3e, 0x98, 0xd6, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21,
	0x01, 0x2e, 0xe6, 0xa2, 0xd4, 0x62, 0xa8, 0x66, 0x10, 0xd3, 0xc8, 0x81, 0x8b, 0x27, 0x14, 0x6c,
	0x35, 0x44, 0xa5, 0x90, 0x01, 0x17, 0x73, 0x50, 0x69, 0x9e, 0x90, 0x90, 0x5e, 0x51, 0x41, 0xb2,
	0x1e, 0x8a, 0xc5, 0x52, 0xc2, 0x28, 0x62, 0x10, 0x13, 0x95, 0x18, 0x9c, 0x38, 0xa2, 0xd8, 0xf4,
	0xf4, 0xad, 0x8b, 0x0a, 0x92, 0x93, 0xd8, 0xc0, 0xce, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xc9, 0xfb, 0x98, 0xee, 0xcd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UpdateConfigClient is the client API for UpdateConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpdateConfigClient interface {
	Run(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
}

type updateConfigClient struct {
	cc *grpc.ClientConn
}

func NewUpdateConfigClient(cc *grpc.ClientConn) UpdateConfigClient {
	return &updateConfigClient{cc}
}

func (c *updateConfigClient) Run(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/rpc.UpdateConfig/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateConfigServer is the server API for UpdateConfig service.
type UpdateConfigServer interface {
	Run(context.Context, *ConfigRequest) (*ConfigResponse, error)
}

// UnimplementedUpdateConfigServer can be embedded to have forward compatible implementations.
type UnimplementedUpdateConfigServer struct {
}

func (*UnimplementedUpdateConfigServer) Run(ctx context.Context, req *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}

func RegisterUpdateConfigServer(s *grpc.Server, srv UpdateConfigServer) {
	s.RegisterService(&_UpdateConfig_serviceDesc, srv)
}

func _UpdateConfig_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateConfigServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.UpdateConfig/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateConfigServer).Run(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UpdateConfig_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.UpdateConfig",
	HandlerType: (*UpdateConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _UpdateConfig_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "update_config.proto",
}
