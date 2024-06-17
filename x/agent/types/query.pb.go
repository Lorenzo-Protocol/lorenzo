// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/agent/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryAgentsRequest is the request type for the Query/Agents RPC method.
type QueryAgentsRequest struct {
}

func (m *QueryAgentsRequest) Reset()         { *m = QueryAgentsRequest{} }
func (m *QueryAgentsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAgentsRequest) ProtoMessage()    {}
func (*QueryAgentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8962260ae9a0a6e1, []int{0}
}
func (m *QueryAgentsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAgentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAgentsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAgentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAgentsRequest.Merge(m, src)
}
func (m *QueryAgentsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAgentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAgentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAgentsRequest proto.InternalMessageInfo

// QueryAgentsResponse is the response type for the Query/Agents RPC method.
type QueryAgentsResponse struct {
	// Agent Contains the details of the agent.
	Agents []Agent `protobuf:"bytes,1,rep,name=agents,proto3" json:"agents"`
}

func (m *QueryAgentsResponse) Reset()         { *m = QueryAgentsResponse{} }
func (m *QueryAgentsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAgentsResponse) ProtoMessage()    {}
func (*QueryAgentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8962260ae9a0a6e1, []int{1}
}
func (m *QueryAgentsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAgentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAgentsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAgentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAgentsResponse.Merge(m, src)
}
func (m *QueryAgentsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAgentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAgentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAgentsResponse proto.InternalMessageInfo

func (m *QueryAgentsResponse) GetAgents() []Agent {
	if m != nil {
		return m.Agents
	}
	return nil
}

// QueryAgentRequest is the request type for the Query/Agent RPC method.
type QueryAgentRequest struct {
	// id is the unique identifier of the agent
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryAgentRequest) Reset()         { *m = QueryAgentRequest{} }
func (m *QueryAgentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAgentRequest) ProtoMessage()    {}
func (*QueryAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8962260ae9a0a6e1, []int{2}
}
func (m *QueryAgentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAgentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAgentRequest.Merge(m, src)
}
func (m *QueryAgentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAgentRequest proto.InternalMessageInfo

func (m *QueryAgentRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// QueryAgentResponse is the response type for the Query/Agent RPC method.
type QueryAgentResponse struct {
	// Agent Contains the details of the agent.
	Agent Agent `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent"`
}

func (m *QueryAgentResponse) Reset()         { *m = QueryAgentResponse{} }
func (m *QueryAgentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAgentResponse) ProtoMessage()    {}
func (*QueryAgentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8962260ae9a0a6e1, []int{3}
}
func (m *QueryAgentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAgentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAgentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAgentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAgentResponse.Merge(m, src)
}
func (m *QueryAgentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAgentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAgentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAgentResponse proto.InternalMessageInfo

func (m *QueryAgentResponse) GetAgent() Agent {
	if m != nil {
		return m.Agent
	}
	return Agent{}
}

type QueryAdminRequest struct {
}

func (m *QueryAdminRequest) Reset()         { *m = QueryAdminRequest{} }
func (m *QueryAdminRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAdminRequest) ProtoMessage()    {}
func (*QueryAdminRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8962260ae9a0a6e1, []int{4}
}
func (m *QueryAdminRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAdminRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAdminRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAdminRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAdminRequest.Merge(m, src)
}
func (m *QueryAdminRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAdminRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAdminRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAdminRequest proto.InternalMessageInfo

// QueryAgentResponse is the response type for the Query/Agent RPC method.
type QueryAdminResponse struct {
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (m *QueryAdminResponse) Reset()         { *m = QueryAdminResponse{} }
func (m *QueryAdminResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAdminResponse) ProtoMessage()    {}
func (*QueryAdminResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8962260ae9a0a6e1, []int{5}
}
func (m *QueryAdminResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAdminResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAdminResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAdminResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAdminResponse.Merge(m, src)
}
func (m *QueryAdminResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAdminResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAdminResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAdminResponse proto.InternalMessageInfo

func (m *QueryAdminResponse) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryAgentsRequest)(nil), "lorenzo.agent.v1.QueryAgentsRequest")
	proto.RegisterType((*QueryAgentsResponse)(nil), "lorenzo.agent.v1.QueryAgentsResponse")
	proto.RegisterType((*QueryAgentRequest)(nil), "lorenzo.agent.v1.QueryAgentRequest")
	proto.RegisterType((*QueryAgentResponse)(nil), "lorenzo.agent.v1.QueryAgentResponse")
	proto.RegisterType((*QueryAdminRequest)(nil), "lorenzo.agent.v1.QueryAdminRequest")
	proto.RegisterType((*QueryAdminResponse)(nil), "lorenzo.agent.v1.QueryAdminResponse")
}

func init() { proto.RegisterFile("lorenzo/agent/v1/query.proto", fileDescriptor_8962260ae9a0a6e1) }

var fileDescriptor_8962260ae9a0a6e1 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0xdb, 0x7a, 0x4b, 0xe2, 0x98, 0x18, 0x9d, 0x4b, 0x72, 0x7b, 0x1b, 0x52, 0x48, 0xc1,
	0x84, 0x0d, 0x9d, 0x00, 0xf1, 0x01, 0x20, 0x6e, 0x8c, 0x2c, 0xb4, 0xee, 0xdc, 0x98, 0x42, 0x27,
	0x75, 0x12, 0x98, 0x29, 0x9d, 0x01, 0x41, 0xe3, 0xc6, 0x27, 0x30, 0xf1, 0x55, 0x7c, 0x08, 0x96,
	0x44, 0x37, 0xae, 0x8c, 0x01, 0xf7, 0xbe, 0x82, 0xe9, 0xcc, 0x14, 0x51, 0xfe, 0xed, 0xda, 0xf3,
	0x7d, 0xe7, 0xfb, 0x9d, 0x39, 0x33, 0xa0, 0x32, 0x66, 0x19, 0xa6, 0xef, 0x18, 0x8a, 0x12, 0x4c,
	0x05, 0x9a, 0xb7, 0xd1, 0x74, 0x86, 0xb3, 0x65, 0x90, 0x66, 0x4c, 0x30, 0xf8, 0x40, 0xab, 0x81,
	0x54, 0x83, 0x79, 0xdb, 0x2d, 0x27, 0x2c, 0x61, 0x52, 0x44, 0xf9, 0x97, 0xf2, 0xb9, 0xb7, 0x23,
	0xc6, 0x27, 0x8c, 0xbf, 0x56, 0x82, 0xfa, 0xd1, 0x52, 0x25, 0x61, 0x2c, 0x19, 0x63, 0x14, 0xa5,
	0x04, 0x45, 0x94, 0x32, 0x11, 0x09, 0xc2, 0xe8, 0x4e, 0x3d, 0xc0, 0x2b, 0x92, 0x54, 0xfd, 0x32,
	0x80, 0x2f, 0xf2, 0x69, 0x7a, 0x79, 0x8d, 0x87, 0x78, 0x3a, 0xc3, 0x5c, 0xf8, 0x03, 0x70, 0xfd,
	0x4f, 0x95, 0xa7, 0x8c, 0x72, 0x0c, 0x1f, 0x83, 0x92, 0xec, 0xe5, 0x8e, 0x59, 0xbb, 0xd3, 0xbc,
	0xd7, 0xb9, 0x09, 0xfe, 0x1f, 0x3e, 0x90, 0x1d, 0xfd, 0xab, 0xd5, 0x8f, 0xaa, 0x11, 0x6a, 0xb3,
	0x5f, 0x07, 0x0f, 0xff, 0xa6, 0x69, 0x04, 0xbc, 0x0f, 0x2c, 0x12, 0x3b, 0x66, 0xcd, 0x6c, 0x5e,
	0x85, 0x16, 0x89, 0xfd, 0xa7, 0xfb, 0x83, 0xec, 0x88, 0x5d, 0x60, 0xcb, 0x10, 0x69, 0xbc, 0x08,
	0x54, 0x5e, 0xff, 0xba, 0xe0, 0xc5, 0x13, 0x42, 0x8b, 0x23, 0x3d, 0x29, 0xf2, 0x55, 0x51, 0xe7,
	0x07, 0xc0, 0x8e, 0xf2, 0x82, 0xcc, 0xbf, 0xdb, 0x77, 0xbe, 0x7e, 0x69, 0x95, 0xf5, 0x6e, 0x7b,
	0x71, 0x9c, 0x61, 0xce, 0x5f, 0x8a, 0x8c, 0xd0, 0x24, 0x54, 0xb6, 0xce, 0x6f, 0x0b, 0xd8, 0x32,
	0x06, 0xbe, 0x05, 0x25, 0xb5, 0x1d, 0xd8, 0x38, 0x1c, 0xea, 0x70, 0xa5, 0xee, 0xa3, 0x0b, 0x2e,
	0x35, 0x90, 0x5f, 0xfb, 0xf8, 0xed, 0xd7, 0x67, 0xcb, 0x85, 0x0e, 0x3a, 0x7e, 0x6d, 0x1c, 0x2e,
	0x80, 0x2d, 0x7b, 0x60, 0xfd, 0x5c, 0x62, 0x81, 0x6d, 0x9c, 0x37, 0x69, 0x6a, 0x43, 0x52, 0x3d,
	0x58, 0x39, 0x41, 0x45, 0xef, 0x49, 0xfc, 0x01, 0x72, 0x60, 0xcb, 0xed, 0x9d, 0x26, 0xef, 0x2d,
	0xfc, 0x34, 0x79, 0xff, 0x02, 0xfc, 0xaa, 0x24, 0xdf, 0xc2, 0x9b, 0x23, 0xe4, 0xdc, 0xd8, 0x7f,
	0xb6, 0xda, 0x78, 0xe6, 0x7a, 0xe3, 0x99, 0x3f, 0x37, 0x9e, 0xf9, 0x69, 0xeb, 0x19, 0xeb, 0xad,
	0x67, 0x7c, 0xdf, 0x7a, 0xc6, 0xab, 0x76, 0x42, 0xc4, 0x9b, 0xd9, 0x30, 0x18, 0xb1, 0x09, 0x1a,
	0xa8, 0xe6, 0xd6, 0xf3, 0xfc, 0x51, 0x8f, 0xd8, 0x78, 0x97, 0xb6, 0xd0, 0x79, 0x62, 0x99, 0x62,
	0x3e, 0x2c, 0xc9, 0x47, 0xdf, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x26, 0xb8, 0xfa, 0x93,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Agent queries all agents
	Agents(ctx context.Context, in *QueryAgentsRequest, opts ...grpc.CallOption) (*QueryAgentsResponse, error)
	// Agent queries the agent of the specified escrow_address
	Agent(ctx context.Context, in *QueryAgentRequest, opts ...grpc.CallOption) (*QueryAgentResponse, error)
	// Admin queries the admin of the agent module
	Admin(ctx context.Context, in *QueryAdminRequest, opts ...grpc.CallOption) (*QueryAdminResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Agents(ctx context.Context, in *QueryAgentsRequest, opts ...grpc.CallOption) (*QueryAgentsResponse, error) {
	out := new(QueryAgentsResponse)
	err := c.cc.Invoke(ctx, "/lorenzo.agent.v1.Query/Agents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Agent(ctx context.Context, in *QueryAgentRequest, opts ...grpc.CallOption) (*QueryAgentResponse, error) {
	out := new(QueryAgentResponse)
	err := c.cc.Invoke(ctx, "/lorenzo.agent.v1.Query/Agent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Admin(ctx context.Context, in *QueryAdminRequest, opts ...grpc.CallOption) (*QueryAdminResponse, error) {
	out := new(QueryAdminResponse)
	err := c.cc.Invoke(ctx, "/lorenzo.agent.v1.Query/Admin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Agent queries all agents
	Agents(context.Context, *QueryAgentsRequest) (*QueryAgentsResponse, error)
	// Agent queries the agent of the specified escrow_address
	Agent(context.Context, *QueryAgentRequest) (*QueryAgentResponse, error)
	// Admin queries the admin of the agent module
	Admin(context.Context, *QueryAdminRequest) (*QueryAdminResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Agents(ctx context.Context, req *QueryAgentsRequest) (*QueryAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Agents not implemented")
}
func (*UnimplementedQueryServer) Agent(ctx context.Context, req *QueryAgentRequest) (*QueryAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Agent not implemented")
}
func (*UnimplementedQueryServer) Admin(ctx context.Context, req *QueryAdminRequest) (*QueryAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Admin not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Agents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Agents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorenzo.agent.v1.Query/Agents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Agents(ctx, req.(*QueryAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Agent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Agent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorenzo.agent.v1.Query/Agent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Agent(ctx, req.(*QueryAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Admin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Admin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorenzo.agent.v1.Query/Admin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Admin(ctx, req.(*QueryAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lorenzo.agent.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Agents",
			Handler:    _Query_Agents_Handler,
		},
		{
			MethodName: "Agent",
			Handler:    _Query_Agent_Handler,
		},
		{
			MethodName: "Admin",
			Handler:    _Query_Admin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorenzo/agent/v1/query.proto",
}

func (m *QueryAgentsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAgentsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAgentsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAgentsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAgentsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAgentsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Agents) > 0 {
		for iNdEx := len(m.Agents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Agents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryAgentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAgentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAgentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryAgentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAgentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAgentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Agent.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAdminRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAdminRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAdminRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAdminResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAdminResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAdminResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAgentsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAgentsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Agents) > 0 {
		for _, e := range m.Agents {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryAgentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryAgentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Agent.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAdminRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAdminResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAgentsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAgentsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAgentsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAgentsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAgentsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAgentsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Agents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Agents = append(m.Agents, Agent{})
			if err := m.Agents[len(m.Agents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAgentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAgentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAgentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAgentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAgentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAgentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Agent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Agent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAdminRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAdminRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAdminRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAdminResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAdminResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAdminResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
