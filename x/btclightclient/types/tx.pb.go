// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/btclightclient/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_Lorenzo_Protocol_lorenzo_v3_types "github.com/Lorenzo-Protocol/lorenzo/v3/types"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgInsertHeaders defines the message for multiple incoming header bytes
type MsgInsertHeaders struct {
	Signer  string                                                        `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Headers []github_com_Lorenzo_Protocol_lorenzo_v3_types.BTCHeaderBytes `protobuf:"bytes,2,rep,name=headers,proto3,customtype=github.com/Lorenzo-Protocol/lorenzo/v3/types.BTCHeaderBytes" json:"headers,omitempty"`
}

func (m *MsgInsertHeaders) Reset()         { *m = MsgInsertHeaders{} }
func (m *MsgInsertHeaders) String() string { return proto.CompactTextString(m) }
func (*MsgInsertHeaders) ProtoMessage()    {}
func (*MsgInsertHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_16c2a8a24fcd05e8, []int{0}
}
func (m *MsgInsertHeaders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInsertHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInsertHeaders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInsertHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInsertHeaders.Merge(m, src)
}
func (m *MsgInsertHeaders) XXX_Size() int {
	return m.Size()
}
func (m *MsgInsertHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInsertHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInsertHeaders proto.InternalMessageInfo

func (m *MsgInsertHeaders) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

// MsgInsertHeadersResponse defines the response for the InsertHeaders
// transaction
type MsgInsertHeadersResponse struct {
}

func (m *MsgInsertHeadersResponse) Reset()         { *m = MsgInsertHeadersResponse{} }
func (m *MsgInsertHeadersResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInsertHeadersResponse) ProtoMessage()    {}
func (*MsgInsertHeadersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16c2a8a24fcd05e8, []int{1}
}
func (m *MsgInsertHeadersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInsertHeadersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInsertHeadersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInsertHeadersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInsertHeadersResponse.Merge(m, src)
}
func (m *MsgInsertHeadersResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInsertHeadersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInsertHeadersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInsertHeadersResponse proto.InternalMessageInfo

// MsgUpdateParams defines a message for updating btc light client module
// parameters.
type MsgUpdateParams struct {
	// authority is the address of the governance account.
	// just FYI: cosmos.AddressString marks that this field should use type alias
	// for AddressString instead of string, but the functionality is not yet
	// implemented in cosmos-proto
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the btc light client parameters to update.
	//
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_16c2a8a24fcd05e8, []int{2}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse is the response to the MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16c2a8a24fcd05e8, []int{3}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

type MsgUpdateFeeRate struct {
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// sat/vbyte * 1000
	FeeRate uint64 `protobuf:"varint,2,opt,name=fee_rate,json=feeRate,proto3" json:"fee_rate,omitempty"`
}

func (m *MsgUpdateFeeRate) Reset()         { *m = MsgUpdateFeeRate{} }
func (m *MsgUpdateFeeRate) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateFeeRate) ProtoMessage()    {}
func (*MsgUpdateFeeRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_16c2a8a24fcd05e8, []int{4}
}
func (m *MsgUpdateFeeRate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateFeeRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateFeeRate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateFeeRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateFeeRate.Merge(m, src)
}
func (m *MsgUpdateFeeRate) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateFeeRate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateFeeRate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateFeeRate proto.InternalMessageInfo

func (m *MsgUpdateFeeRate) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgUpdateFeeRate) GetFeeRate() uint64 {
	if m != nil {
		return m.FeeRate
	}
	return 0
}

type MsgUpdateFeeRateResponse struct {
}

func (m *MsgUpdateFeeRateResponse) Reset()         { *m = MsgUpdateFeeRateResponse{} }
func (m *MsgUpdateFeeRateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateFeeRateResponse) ProtoMessage()    {}
func (*MsgUpdateFeeRateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16c2a8a24fcd05e8, []int{5}
}
func (m *MsgUpdateFeeRateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateFeeRateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateFeeRateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateFeeRateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateFeeRateResponse.Merge(m, src)
}
func (m *MsgUpdateFeeRateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateFeeRateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateFeeRateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateFeeRateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgInsertHeaders)(nil), "lorenzo.btclightclient.v1.MsgInsertHeaders")
	proto.RegisterType((*MsgInsertHeadersResponse)(nil), "lorenzo.btclightclient.v1.MsgInsertHeadersResponse")
	proto.RegisterType((*MsgUpdateParams)(nil), "lorenzo.btclightclient.v1.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "lorenzo.btclightclient.v1.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgUpdateFeeRate)(nil), "lorenzo.btclightclient.v1.MsgUpdateFeeRate")
	proto.RegisterType((*MsgUpdateFeeRateResponse)(nil), "lorenzo.btclightclient.v1.MsgUpdateFeeRateResponse")
}

func init() {
	proto.RegisterFile("lorenzo/btclightclient/v1/tx.proto", fileDescriptor_16c2a8a24fcd05e8)
}

var fileDescriptor_16c2a8a24fcd05e8 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6b, 0x13, 0x41,
	0x18, 0xcd, 0xa6, 0x35, 0xb5, 0x63, 0x8b, 0xba, 0x14, 0xbb, 0xd9, 0xc3, 0x36, 0xee, 0xa1, 0x86,
	0x94, 0xee, 0xd2, 0x04, 0x14, 0x2b, 0x52, 0x5c, 0x41, 0x14, 0x0c, 0x94, 0x55, 0x11, 0xbd, 0x94,
	0x4d, 0x32, 0x9d, 0x2c, 0x64, 0x77, 0xd6, 0xf9, 0xa6, 0xa1, 0xf1, 0x20, 0xe2, 0xd1, 0x93, 0x3f,
	0x25, 0x07, 0xcf, 0xde, 0x84, 0x1e, 0x8b, 0x27, 0xf1, 0x50, 0x24, 0x11, 0xf2, 0x37, 0x64, 0x77,
	0x26, 0x6d, 0x77, 0x71, 0x6b, 0x7a, 0x19, 0x32, 0xdf, 0xbc, 0xf9, 0xde, 0xf7, 0xde, 0xcb, 0x2c,
	0x32, 0x7b, 0x94, 0xe1, 0xf0, 0x3d, 0xb5, 0x5b, 0xbc, 0xdd, 0xf3, 0x49, 0x37, 0x5e, 0x71, 0xc8,
	0xed, 0xfe, 0x96, 0xcd, 0x0f, 0xad, 0x88, 0x51, 0x4e, 0xd5, 0xb2, 0xc4, 0x58, 0x69, 0x8c, 0xd5,
	0xdf, 0xd2, 0x57, 0x08, 0x25, 0x34, 0x41, 0xd9, 0xf1, 0x2f, 0x71, 0x41, 0x5f, 0x6d, 0x53, 0x08,
	0x28, 0xd8, 0x01, 0x90, 0xb8, 0x51, 0x00, 0x44, 0x1e, 0xac, 0xe7, 0xb3, 0x45, 0x1e, 0xf3, 0x02,
	0x90, 0xb8, 0x9b, 0x5e, 0xe0, 0x87, 0xd4, 0x4e, 0x56, 0x59, 0x2a, 0x8b, 0x9e, 0x7b, 0x82, 0x4c,
	0x6c, 0xc4, 0x91, 0xf9, 0x4d, 0x41, 0x37, 0x9a, 0x40, 0x9e, 0x85, 0x80, 0x19, 0x7f, 0x8a, 0xbd,
	0x0e, 0x66, 0xa0, 0xde, 0x42, 0x25, 0xf0, 0x49, 0x88, 0x99, 0xa6, 0x54, 0x94, 0xea, 0xa2, 0x2b,
	0x77, 0xea, 0x1b, 0xb4, 0xd0, 0x15, 0x10, 0xad, 0x58, 0x99, 0xab, 0x2e, 0x39, 0x3b, 0xbf, 0x4e,
	0xd6, 0x1e, 0x10, 0x9f, 0x77, 0x0f, 0x5a, 0x56, 0x9b, 0x06, 0xf6, 0x73, 0x31, 0xe2, 0xe6, 0x6e,
	0xdc, 0xbb, 0x4d, 0x7b, 0xf6, 0x74, 0xe6, 0x7e, 0xc3, 0xe6, 0x83, 0x08, 0x83, 0xe5, 0xbc, 0x7c,
	0x2c, 0x58, 0x9c, 0x01, 0xc7, 0xe0, 0x4e, 0xfb, 0x6d, 0xdf, 0xfb, 0x34, 0x19, 0xd6, 0x24, 0xcf,
	0xe7, 0xc9, 0xb0, 0x76, 0x27, 0x47, 0x6d, 0x76, 0x56, 0x53, 0x47, 0x5a, 0xb6, 0xe6, 0x62, 0x88,
	0x68, 0x08, 0xd8, 0xfc, 0xae, 0xa0, 0xeb, 0x4d, 0x20, 0xaf, 0xa2, 0x8e, 0xc7, 0xf1, 0x6e, 0x62,
	0x92, 0x7a, 0x17, 0x2d, 0x7a, 0x07, 0xbc, 0x4b, 0x99, 0xcf, 0x07, 0x42, 0x9e, 0xa3, 0xfd, 0xf8,
	0xba, 0xb9, 0x22, 0x5d, 0x79, 0xd4, 0xe9, 0x30, 0x0c, 0xf0, 0x82, 0x33, 0x3f, 0x24, 0xee, 0x19,
	0x54, 0xdd, 0x41, 0x25, 0x61, 0xb3, 0x56, 0xac, 0x28, 0xd5, 0x6b, 0xf5, 0xdb, 0x56, 0x6e, 0xb2,
	0x96, 0xa0, 0x72, 0xe6, 0x8f, 0x4e, 0xd6, 0x0a, 0xae, 0xbc, 0xb6, 0x7d, 0x3f, 0x56, 0x78, 0xd6,
	0x30, 0x16, 0xb9, 0x9e, 0x2f, 0xf2, 0xfc, 0xcc, 0x66, 0x19, 0xad, 0x66, 0x4a, 0xa7, 0x12, 0x3f,
	0x24, 0xf1, 0x89, 0xa3, 0x27, 0x18, 0xbb, 0x1e, 0xc7, 0xb9, 0xf1, 0x95, 0xd1, 0xd5, 0x7d, 0x8c,
	0xf7, 0x98, 0xc7, 0x71, 0x22, 0x62, 0xde, 0x5d, 0xd8, 0x17, 0x57, 0x2e, 0x63, 0x7f, 0x8a, 0x4b,
	0xda, 0x9f, 0xaa, 0x4d, 0x67, 0xab, 0xff, 0x29, 0xa2, 0xb9, 0x26, 0x10, 0x15, 0xd0, 0x72, 0xfa,
	0xff, 0xb5, 0x71, 0x81, 0x77, 0xd9, 0x30, 0xf5, 0xc6, 0x25, 0xc0, 0xa7, 0xb6, 0x14, 0xd4, 0x10,
	0x2d, 0xa5, 0x72, 0xaf, 0x5d, 0xdc, 0xe6, 0x3c, 0x56, 0xaf, 0xcf, 0x8e, 0x9d, 0x32, 0xaa, 0xef,
	0xd0, 0x72, 0x3a, 0x85, 0x8d, 0x59, 0x9a, 0x48, 0xf0, 0xff, 0x44, 0xfe, 0xd3, 0x5f, 0xfd, 0xca,
	0xc7, 0xc9, 0xb0, 0xa6, 0x38, 0xaf, 0x8f, 0x46, 0x86, 0x72, 0x3c, 0x32, 0x94, 0xdf, 0x23, 0x43,
	0xf9, 0x32, 0x36, 0x0a, 0xc7, 0x63, 0xa3, 0xf0, 0x73, 0x6c, 0x14, 0xde, 0x3e, 0x9c, 0xf1, 0x69,
	0x1e, 0x66, 0x63, 0x4e, 0xde, 0x6a, 0xab, 0x94, 0x7c, 0x22, 0x1a, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x60, 0x2f, 0x92, 0x40, 0xe8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// InsertHeaders adds a batch of headers to the BTC light client chain
	InsertHeaders(ctx context.Context, in *MsgInsertHeaders, opts ...grpc.CallOption) (*MsgInsertHeadersResponse, error)
	// UpdateParams defines a method for updating btc light client module
	// parameters.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	UpdateFeeRate(ctx context.Context, in *MsgUpdateFeeRate, opts ...grpc.CallOption) (*MsgUpdateFeeRateResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) InsertHeaders(ctx context.Context, in *MsgInsertHeaders, opts ...grpc.CallOption) (*MsgInsertHeadersResponse, error) {
	out := new(MsgInsertHeadersResponse)
	err := c.cc.Invoke(ctx, "/lorenzo.btclightclient.v1.Msg/InsertHeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/lorenzo.btclightclient.v1.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateFeeRate(ctx context.Context, in *MsgUpdateFeeRate, opts ...grpc.CallOption) (*MsgUpdateFeeRateResponse, error) {
	out := new(MsgUpdateFeeRateResponse)
	err := c.cc.Invoke(ctx, "/lorenzo.btclightclient.v1.Msg/UpdateFeeRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// InsertHeaders adds a batch of headers to the BTC light client chain
	InsertHeaders(context.Context, *MsgInsertHeaders) (*MsgInsertHeadersResponse, error)
	// UpdateParams defines a method for updating btc light client module
	// parameters.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	UpdateFeeRate(context.Context, *MsgUpdateFeeRate) (*MsgUpdateFeeRateResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) InsertHeaders(ctx context.Context, req *MsgInsertHeaders) (*MsgInsertHeadersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertHeaders not implemented")
}
func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) UpdateFeeRate(ctx context.Context, req *MsgUpdateFeeRate) (*MsgUpdateFeeRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFeeRate not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_InsertHeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInsertHeaders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InsertHeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorenzo.btclightclient.v1.Msg/InsertHeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InsertHeaders(ctx, req.(*MsgInsertHeaders))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorenzo.btclightclient.v1.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateFeeRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateFeeRate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateFeeRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorenzo.btclightclient.v1.Msg/UpdateFeeRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateFeeRate(ctx, req.(*MsgUpdateFeeRate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lorenzo.btclightclient.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertHeaders",
			Handler:    _Msg_InsertHeaders_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "UpdateFeeRate",
			Handler:    _Msg_UpdateFeeRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorenzo/btclightclient/v1/tx.proto",
}

func (m *MsgInsertHeaders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInsertHeaders) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInsertHeaders) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Headers[iNdEx].Size()
				i -= size
				if _, err := m.Headers[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInsertHeadersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInsertHeadersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInsertHeadersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateFeeRate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateFeeRate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateFeeRate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeeRate != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.FeeRate))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateFeeRateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateFeeRateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateFeeRateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInsertHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgInsertHeadersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateFeeRate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.FeeRate != 0 {
		n += 1 + sovTx(uint64(m.FeeRate))
	}
	return n
}

func (m *MsgUpdateFeeRateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInsertHeaders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInsertHeaders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInsertHeaders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_Lorenzo_Protocol_lorenzo_v3_types.BTCHeaderBytes
			m.Headers = append(m.Headers, v)
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInsertHeadersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInsertHeadersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInsertHeadersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateFeeRate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateFeeRate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateFeeRate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			m.FeeRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FeeRate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateFeeRateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateFeeRateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateFeeRateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
