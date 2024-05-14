// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/btcstaking/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_Lorenzo_Protocol_lorenzo_types "github.com/Lorenzo-Protocol/lorenzo/types"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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

type TransactionKey struct {
	Index uint32                                                        `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Hash  *github_com_Lorenzo_Protocol_lorenzo_types.BTCHeaderHashBytes `protobuf:"bytes,2,opt,name=hash,proto3,customtype=github.com/Lorenzo-Protocol/lorenzo/types.BTCHeaderHashBytes" json:"hash,omitempty"`
}

func (m *TransactionKey) Reset()         { *m = TransactionKey{} }
func (m *TransactionKey) String() string { return proto.CompactTextString(m) }
func (*TransactionKey) ProtoMessage()    {}
func (*TransactionKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be51bab5db52b8e, []int{0}
}
func (m *TransactionKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionKey.Merge(m, src)
}
func (m *TransactionKey) XXX_Size() int {
	return m.Size()
}
func (m *TransactionKey) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionKey.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionKey proto.InternalMessageInfo

func (m *TransactionKey) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

// TransactionInfo is the info of a tx on Bitcoin,
// including
// - the position of the tx on BTC blockchain
// - the full tx content
// - the Merkle proof that this tx is on the above position
type TransactionInfo struct {
	// key is the position (txIdx, blockHash) of this tx on BTC blockchain
	// Although it is already a part of SubmissionKey, we store it here again
	// to make TransactionInfo self-contained.
	// For example, storing the key allows TransactionInfo to not relay on
	// the fact that TransactionInfo will be ordered in the same order as
	// TransactionKeys in SubmissionKey.
	Key *TransactionKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// transaction is the full transaction in bytes
	Transaction []byte `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// proof is the Merkle proof that this tx is included in the position in `key`
	// TODO: maybe it could use here better format as we already processed and
	// validated the proof?
	Proof []byte `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *TransactionInfo) Reset()         { *m = TransactionInfo{} }
func (m *TransactionInfo) String() string { return proto.CompactTextString(m) }
func (*TransactionInfo) ProtoMessage()    {}
func (*TransactionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be51bab5db52b8e, []int{1}
}
func (m *TransactionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInfo.Merge(m, src)
}
func (m *TransactionInfo) XXX_Size() int {
	return m.Size()
}
func (m *TransactionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInfo proto.InternalMessageInfo

func (m *TransactionInfo) GetKey() *TransactionKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *TransactionInfo) GetTransaction() []byte {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *TransactionInfo) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

type MsgCreateBTCStaking struct {
	Signer    string           `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	StakingTx *TransactionInfo `protobuf:"bytes,2,opt,name=staking_tx,json=stakingTx,proto3" json:"staking_tx,omitempty"`
}

func (m *MsgCreateBTCStaking) Reset()         { *m = MsgCreateBTCStaking{} }
func (m *MsgCreateBTCStaking) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBTCStaking) ProtoMessage()    {}
func (*MsgCreateBTCStaking) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be51bab5db52b8e, []int{2}
}
func (m *MsgCreateBTCStaking) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBTCStaking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBTCStaking.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBTCStaking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBTCStaking.Merge(m, src)
}
func (m *MsgCreateBTCStaking) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBTCStaking) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBTCStaking.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBTCStaking proto.InternalMessageInfo

func (m *MsgCreateBTCStaking) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgCreateBTCStaking) GetStakingTx() *TransactionInfo {
	if m != nil {
		return m.StakingTx
	}
	return nil
}

type MsgCreateBTCStakingResponse struct {
}

func (m *MsgCreateBTCStakingResponse) Reset()         { *m = MsgCreateBTCStakingResponse{} }
func (m *MsgCreateBTCStakingResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBTCStakingResponse) ProtoMessage()    {}
func (*MsgCreateBTCStakingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be51bab5db52b8e, []int{3}
}
func (m *MsgCreateBTCStakingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBTCStakingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBTCStakingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBTCStakingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBTCStakingResponse.Merge(m, src)
}
func (m *MsgCreateBTCStakingResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBTCStakingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBTCStakingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBTCStakingResponse proto.InternalMessageInfo

type MsgBurnRequest struct {
	Signer           string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	BtcTargetAddress string `protobuf:"bytes,2,opt,name=btc_target_address,json=btcTargetAddress,proto3" json:"btc_target_address,omitempty"`
	Amount           string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *MsgBurnRequest) Reset()         { *m = MsgBurnRequest{} }
func (m *MsgBurnRequest) String() string { return proto.CompactTextString(m) }
func (*MsgBurnRequest) ProtoMessage()    {}
func (*MsgBurnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be51bab5db52b8e, []int{4}
}
func (m *MsgBurnRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnRequest.Merge(m, src)
}
func (m *MsgBurnRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnRequest proto.InternalMessageInfo

func (m *MsgBurnRequest) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgBurnRequest) GetBtcTargetAddress() string {
	if m != nil {
		return m.BtcTargetAddress
	}
	return ""
}

func (m *MsgBurnRequest) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type MsgBurnResponse struct {
}

func (m *MsgBurnResponse) Reset()         { *m = MsgBurnResponse{} }
func (m *MsgBurnResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBurnResponse) ProtoMessage()    {}
func (*MsgBurnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be51bab5db52b8e, []int{5}
}
func (m *MsgBurnResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnResponse.Merge(m, src)
}
func (m *MsgBurnResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TransactionKey)(nil), "lorenzo.btcstaking.v1.TransactionKey")
	proto.RegisterType((*TransactionInfo)(nil), "lorenzo.btcstaking.v1.TransactionInfo")
	proto.RegisterType((*MsgCreateBTCStaking)(nil), "lorenzo.btcstaking.v1.MsgCreateBTCStaking")
	proto.RegisterType((*MsgCreateBTCStakingResponse)(nil), "lorenzo.btcstaking.v1.MsgCreateBTCStakingResponse")
	proto.RegisterType((*MsgBurnRequest)(nil), "lorenzo.btcstaking.v1.MsgBurnRequest")
	proto.RegisterType((*MsgBurnResponse)(nil), "lorenzo.btcstaking.v1.MsgBurnResponse")
}

func init() { proto.RegisterFile("lorenzo/btcstaking/v1/tx.proto", fileDescriptor_6be51bab5db52b8e) }

var fileDescriptor_6be51bab5db52b8e = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0xa6, 0x2d, 0x64, 0xa2, 0x6d, 0x1d, 0xab, 0xc6, 0x88, 0x6b, 0x58, 0x68, 0x29,
	0x41, 0x77, 0x69, 0x04, 0x05, 0xf1, 0xa0, 0x1b, 0x84, 0x8a, 0x06, 0xca, 0xb8, 0x5e, 0xbc, 0x84,
	0xd9, 0xcd, 0x74, 0xb2, 0xb4, 0x3b, 0x13, 0x67, 0x26, 0x65, 0xa3, 0x1e, 0xb4, 0x9f, 0xc0, 0x8f,
	0xd2, 0x8f, 0xe1, 0xb1, 0x37, 0xc5, 0x83, 0x48, 0x72, 0xe8, 0xd7, 0x90, 0x9d, 0xd9, 0xda, 0x14,
	0x13, 0xda, 0xdb, 0xbc, 0xf7, 0x7f, 0xf3, 0xfe, 0xbf, 0xf7, 0x98, 0x81, 0xce, 0xbe, 0x90, 0x94,
	0x7f, 0x14, 0x7e, 0xa4, 0x63, 0xa5, 0xc9, 0x5e, 0xc2, 0x99, 0x7f, 0xb0, 0xe5, 0xeb, 0xcc, 0x1b,
	0x48, 0xa1, 0x05, 0xba, 0x59, 0xe8, 0xde, 0x99, 0xee, 0x1d, 0x6c, 0xd5, 0x6f, 0xc7, 0x42, 0xa5,
	0x42, 0xf9, 0xa9, 0x32, 0xe5, 0xa9, 0x62, 0xb6, 0xbe, 0x7e, 0xc7, 0x0a, 0x5d, 0x13, 0xf9, 0x36,
	0x28, 0xa4, 0x35, 0x26, 0x98, 0xb0, 0xf9, 0xfc, 0x64, 0xb3, 0xee, 0x67, 0xb8, 0x1c, 0x4a, 0xc2,
	0x15, 0x89, 0x75, 0x22, 0xf8, 0x6b, 0x3a, 0x42, 0x6b, 0x70, 0x31, 0xe1, 0x3d, 0x9a, 0xd5, 0x40,
	0x03, 0x6c, 0x5e, 0xc3, 0x36, 0x40, 0x21, 0x5c, 0xe8, 0x13, 0xd5, 0xaf, 0x5d, 0x69, 0x80, 0xcd,
	0xab, 0xc1, 0xf3, 0x5f, 0xbf, 0xef, 0x3f, 0x63, 0x89, 0xee, 0x0f, 0x23, 0x2f, 0x16, 0xa9, 0xff,
	0xc6, 0x52, 0x3e, 0xdc, 0xc9, 0x7b, 0xc6, 0x62, 0xdf, 0x3f, 0x1d, 0x4b, 0x8f, 0x06, 0x54, 0x79,
	0x41, 0xd8, 0xde, 0xa6, 0xa4, 0x47, 0xe5, 0x36, 0x51, 0xfd, 0x60, 0xa4, 0xa9, 0xc2, 0xa6, 0x9b,
	0x7b, 0x08, 0xe0, 0xca, 0x94, 0xfd, 0x2b, 0xbe, 0x2b, 0xd0, 0x13, 0x58, 0xde, 0xa3, 0x23, 0xe3,
	0x5e, 0x6d, 0xad, 0x7b, 0x33, 0x17, 0xe0, 0x9d, 0x67, 0xc6, 0xf9, 0x0d, 0xd4, 0x80, 0x55, 0x7d,
	0x96, 0xb6, 0xa4, 0x78, 0x3a, 0x95, 0x8f, 0x36, 0x90, 0x42, 0xec, 0xd6, 0xca, 0x46, 0xb3, 0x81,
	0xfb, 0x15, 0xc0, 0x1b, 0x1d, 0xc5, 0xda, 0x92, 0x12, 0x4d, 0x83, 0xb0, 0xfd, 0xd6, 0xda, 0xa0,
	0x5b, 0x70, 0x49, 0x25, 0x8c, 0x53, 0x69, 0x58, 0x2a, 0xb8, 0x88, 0xd0, 0x4b, 0x08, 0x0b, 0x92,
	0xae, 0xce, 0x8c, 0x4d, 0xb5, 0xb5, 0x71, 0x31, 0x67, 0x3e, 0x1c, 0xae, 0x14, 0x5a, 0x98, 0x3d,
	0xad, 0x1e, 0x9e, 0x1c, 0x35, 0x8b, 0x9e, 0xee, 0x3d, 0x78, 0x77, 0x06, 0x02, 0xa6, 0x6a, 0x20,
	0xb8, 0xa2, 0xee, 0x27, 0xb8, 0xdc, 0x51, 0x2c, 0x18, 0x4a, 0x8e, 0xe9, 0x87, 0x21, 0x55, 0x7a,
	0x2e, 0xdc, 0x03, 0x88, 0x22, 0x1d, 0x77, 0x35, 0x91, 0x8c, 0xea, 0x2e, 0xe9, 0xf5, 0x24, 0x55,
	0xca, 0x40, 0x56, 0xf0, 0x6a, 0xa4, 0xe3, 0xd0, 0x08, 0x2f, 0x6c, 0x3e, 0xef, 0x42, 0x52, 0x31,
	0xe4, 0xda, 0x6c, 0xa4, 0x82, 0x8b, 0xe8, 0x3c, 0xdb, 0x75, 0xb8, 0xf2, 0xcf, 0xdc, 0xf2, 0xb4,
	0x7e, 0x00, 0x58, 0xee, 0x28, 0x86, 0x24, 0x5c, 0xfd, 0x6f, 0x6d, 0xcd, 0x39, 0xab, 0x98, 0x31,
	0x5f, 0xbd, 0x75, 0xf9, 0xda, 0x53, 0x6f, 0xf4, 0x0e, 0x2e, 0xe4, 0x2c, 0x68, 0x7d, 0xfe, 0xdd,
	0xa9, 0x45, 0xd5, 0x37, 0x2e, 0x2a, 0xb3, 0x6d, 0xeb, 0x8b, 0x5f, 0x4e, 0x8e, 0x9a, 0x20, 0xd8,
	0xf9, 0x3e, 0x76, 0xc0, 0xf1, 0xd8, 0x01, 0x7f, 0xc6, 0x0e, 0xf8, 0x36, 0x71, 0x4a, 0xc7, 0x13,
	0xa7, 0xf4, 0x73, 0xe2, 0x94, 0xde, 0x3f, 0xbe, 0xcc, 0x7b, 0xcf, 0xa6, 0x3f, 0xb2, 0x79, 0xfc,
	0xd1, 0x92, 0xf9, 0x68, 0x8f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x81, 0x85, 0x1a, 0xef, 0xeb,
	0x03, 0x00, 0x00,
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
	// CreateBTCDelegation creates a new BTC delegation
	CreateBTCStaking(ctx context.Context, in *MsgCreateBTCStaking, opts ...grpc.CallOption) (*MsgCreateBTCStakingResponse, error)
	Burn(ctx context.Context, in *MsgBurnRequest, opts ...grpc.CallOption) (*MsgBurnResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateBTCStaking(ctx context.Context, in *MsgCreateBTCStaking, opts ...grpc.CallOption) (*MsgCreateBTCStakingResponse, error) {
	out := new(MsgCreateBTCStakingResponse)
	err := c.cc.Invoke(ctx, "/lorenzo.btcstaking.v1.Msg/CreateBTCStaking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Burn(ctx context.Context, in *MsgBurnRequest, opts ...grpc.CallOption) (*MsgBurnResponse, error) {
	out := new(MsgBurnResponse)
	err := c.cc.Invoke(ctx, "/lorenzo.btcstaking.v1.Msg/Burn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateBTCDelegation creates a new BTC delegation
	CreateBTCStaking(context.Context, *MsgCreateBTCStaking) (*MsgCreateBTCStakingResponse, error)
	Burn(context.Context, *MsgBurnRequest) (*MsgBurnResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateBTCStaking(ctx context.Context, req *MsgCreateBTCStaking) (*MsgCreateBTCStakingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBTCStaking not implemented")
}
func (*UnimplementedMsgServer) Burn(ctx context.Context, req *MsgBurnRequest) (*MsgBurnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Burn not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateBTCStaking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBTCStaking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBTCStaking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorenzo.btcstaking.v1.Msg/CreateBTCStaking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBTCStaking(ctx, req.(*MsgCreateBTCStaking))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Burn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Burn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorenzo.btcstaking.v1.Msg/Burn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Burn(ctx, req.(*MsgBurnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lorenzo.btcstaking.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBTCStaking",
			Handler:    _Msg_CreateBTCStaking_Handler,
		},
		{
			MethodName: "Burn",
			Handler:    _Msg_Burn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorenzo/btcstaking/v1/tx.proto",
}

func (m *TransactionKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Hash != nil {
		{
			size := m.Hash.Size()
			i -= size
			if _, err := m.Hash.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TransactionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Transaction) > 0 {
		i -= len(m.Transaction)
		copy(dAtA[i:], m.Transaction)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Transaction)))
		i--
		dAtA[i] = 0x12
	}
	if m.Key != nil {
		{
			size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateBTCStaking) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBTCStaking) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBTCStaking) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StakingTx != nil {
		{
			size, err := m.StakingTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
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

func (m *MsgCreateBTCStakingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBTCStakingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBTCStakingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgBurnRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BtcTargetAddress) > 0 {
		i -= len(m.BtcTargetAddress)
		copy(dAtA[i:], m.BtcTargetAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BtcTargetAddress)))
		i--
		dAtA[i] = 0x12
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

func (m *MsgBurnResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *TransactionKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovTx(uint64(m.Index))
	}
	if m.Hash != nil {
		l = m.Hash.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *TransactionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Transaction)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateBTCStaking) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.StakingTx != nil {
		l = m.StakingTx.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateBTCStakingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgBurnRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BtcTargetAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgBurnResponse) Size() (n int) {
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
func (m *TransactionKey) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TransactionKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
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
			var v github_com_Lorenzo_Protocol_lorenzo_types.BTCHeaderHashBytes
			m.Hash = &v
			if err := m.Hash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *TransactionInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TransactionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
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
			if m.Key == nil {
				m.Key = &TransactionKey{}
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transaction", wireType)
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
			m.Transaction = append(m.Transaction[:0], dAtA[iNdEx:postIndex]...)
			if m.Transaction == nil {
				m.Transaction = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
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
			m.Proof = append(m.Proof[:0], dAtA[iNdEx:postIndex]...)
			if m.Proof == nil {
				m.Proof = []byte{}
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
func (m *MsgCreateBTCStaking) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateBTCStaking: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBTCStaking: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field StakingTx", wireType)
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
			if m.StakingTx == nil {
				m.StakingTx = &TransactionInfo{}
			}
			if err := m.StakingTx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgCreateBTCStakingResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateBTCStakingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBTCStakingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgBurnRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgBurnRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field BtcTargetAddress", wireType)
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
			m.BtcTargetAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = string(dAtA[iNdEx:postIndex])
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
func (m *MsgBurnResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgBurnResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
