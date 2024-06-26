// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/btcstaking/v1/staking_record.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type BTCStakingRecord struct {
	TxHash          []byte `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Amount          uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	MintToAddr      []byte `protobuf:"bytes,3,opt,name=mint_to_addr,json=mintToAddr,proto3" json:"mint_to_addr,omitempty"`
	BtcReceiverName string `protobuf:"bytes,4,opt,name=btc_receiver_name,json=btcReceiverName,proto3" json:"btc_receiver_name,omitempty"`
	BtcReceiverAddr string `protobuf:"bytes,5,opt,name=btc_receiver_addr,json=btcReceiverAddr,proto3" json:"btc_receiver_addr,omitempty"`
}

func (m *BTCStakingRecord) Reset()         { *m = BTCStakingRecord{} }
func (m *BTCStakingRecord) String() string { return proto.CompactTextString(m) }
func (*BTCStakingRecord) ProtoMessage()    {}
func (*BTCStakingRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fe1ffedd1828bb5, []int{0}
}
func (m *BTCStakingRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BTCStakingRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BTCStakingRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BTCStakingRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BTCStakingRecord.Merge(m, src)
}
func (m *BTCStakingRecord) XXX_Size() int {
	return m.Size()
}
func (m *BTCStakingRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_BTCStakingRecord.DiscardUnknown(m)
}

var xxx_messageInfo_BTCStakingRecord proto.InternalMessageInfo

func (m *BTCStakingRecord) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *BTCStakingRecord) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *BTCStakingRecord) GetMintToAddr() []byte {
	if m != nil {
		return m.MintToAddr
	}
	return nil
}

func (m *BTCStakingRecord) GetBtcReceiverName() string {
	if m != nil {
		return m.BtcReceiverName
	}
	return ""
}

func (m *BTCStakingRecord) GetBtcReceiverAddr() string {
	if m != nil {
		return m.BtcReceiverAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*BTCStakingRecord)(nil), "lorenzo.btcstaking.v1.BTCStakingRecord")
}

func init() {
	proto.RegisterFile("lorenzo/btcstaking/v1/staking_record.proto", fileDescriptor_8fe1ffedd1828bb5)
}

var fileDescriptor_8fe1ffedd1828bb5 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xbb, 0xff, 0x7f, 0xad, 0xb8, 0x14, 0xd4, 0xa0, 0x36, 0xf6, 0x10, 0x82, 0xa7, 0x52,
	0x30, 0x4b, 0x11, 0xbc, 0x5b, 0x2f, 0x1e, 0x44, 0x4a, 0xec, 0xc9, 0x4b, 0xd8, 0x6c, 0x96, 0x24,
	0xd8, 0xcd, 0x94, 0xdd, 0x69, 0xa9, 0x3e, 0x85, 0x2f, 0xe4, 0xdd, 0x63, 0x8f, 0x1e, 0xa5, 0x7d,
	0x11, 0xc9, 0x6e, 0x44, 0xa1, 0xb7, 0xfd, 0xbe, 0xf9, 0xcd, 0xb7, 0xcc, 0x0c, 0x1d, 0xce, 0x40,
	0xcb, 0xea, 0x15, 0x58, 0x8a, 0xc2, 0x20, 0x7f, 0x2e, 0xab, 0x9c, 0x2d, 0x47, 0xac, 0x79, 0x26,
	0x5a, 0x0a, 0xd0, 0x59, 0x34, 0xd7, 0x80, 0xe0, 0x9d, 0x36, 0x6c, 0xf4, 0xcb, 0x46, 0xcb, 0x51,
	0xbf, 0x27, 0xc0, 0x28, 0x30, 0x4c, 0x19, 0xdb, 0xaa, 0x4c, 0xee, 0xf8, 0xfe, 0xb9, 0x2b, 0x24,
	0x56, 0x31, 0x27, 0x9a, 0xd2, 0x49, 0x0e, 0x39, 0x38, 0xbf, 0x7e, 0x39, 0xf7, 0xe2, 0x9d, 0xd0,
	0xa3, 0xf1, 0xf4, 0xf6, 0xd1, 0x65, 0xc7, 0xf6, 0x6f, 0xaf, 0x47, 0xf7, 0x71, 0x95, 0x14, 0xdc,
	0x14, 0x3e, 0x09, 0xc9, 0xa0, 0x1b, 0x77, 0x70, 0x75, 0xc7, 0x4d, 0xe1, 0x9d, 0xd1, 0x0e, 0x57,
	0xb0, 0xa8, 0xd0, 0xff, 0x17, 0x92, 0x41, 0x3b, 0x6e, 0x94, 0x17, 0xd2, 0xae, 0x2a, 0x2b, 0x4c,
	0x10, 0x12, 0x9e, 0x65, 0xda, 0xff, 0x6f, 0xbb, 0x68, 0xed, 0x4d, 0xe1, 0x26, 0xcb, 0xb4, 0x37,
	0xa4, 0xc7, 0x29, 0x8a, 0x7a, 0x38, 0x59, 0x2e, 0xa5, 0x4e, 0x2a, 0xae, 0xa4, 0xdf, 0x0e, 0xc9,
	0xe0, 0x20, 0x3e, 0x4c, 0x51, 0xc4, 0x8d, 0xff, 0xc0, 0x95, 0xdc, 0x61, 0x6d, 0xe4, 0xde, 0x0e,
	0x5b, 0xe7, 0x8e, 0x27, 0x1f, 0x9b, 0x80, 0xac, 0x37, 0x01, 0xf9, 0xda, 0x04, 0xe4, 0x6d, 0x1b,
	0xb4, 0xd6, 0xdb, 0xa0, 0xf5, 0xb9, 0x0d, 0x5a, 0x4f, 0xd7, 0x79, 0x89, 0xc5, 0x22, 0x8d, 0x04,
	0x28, 0x76, 0xef, 0xb6, 0x78, 0x39, 0xa9, 0x67, 0x16, 0x30, 0x63, 0x3f, 0x27, 0x58, 0xfd, 0x3d,
	0x02, 0xbe, 0xcc, 0xa5, 0x49, 0x3b, 0x76, 0x31, 0x57, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x14,
	0x91, 0x25, 0x4b, 0xa7, 0x01, 0x00, 0x00,
}

func (m *BTCStakingRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BTCStakingRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BTCStakingRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BtcReceiverAddr) > 0 {
		i -= len(m.BtcReceiverAddr)
		copy(dAtA[i:], m.BtcReceiverAddr)
		i = encodeVarintStakingRecord(dAtA, i, uint64(len(m.BtcReceiverAddr)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BtcReceiverName) > 0 {
		i -= len(m.BtcReceiverName)
		copy(dAtA[i:], m.BtcReceiverName)
		i = encodeVarintStakingRecord(dAtA, i, uint64(len(m.BtcReceiverName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MintToAddr) > 0 {
		i -= len(m.MintToAddr)
		copy(dAtA[i:], m.MintToAddr)
		i = encodeVarintStakingRecord(dAtA, i, uint64(len(m.MintToAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Amount != 0 {
		i = encodeVarintStakingRecord(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintStakingRecord(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStakingRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovStakingRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BTCStakingRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovStakingRecord(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovStakingRecord(uint64(m.Amount))
	}
	l = len(m.MintToAddr)
	if l > 0 {
		n += 1 + l + sovStakingRecord(uint64(l))
	}
	l = len(m.BtcReceiverName)
	if l > 0 {
		n += 1 + l + sovStakingRecord(uint64(l))
	}
	l = len(m.BtcReceiverAddr)
	if l > 0 {
		n += 1 + l + sovStakingRecord(uint64(l))
	}
	return n
}

func sovStakingRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStakingRecord(x uint64) (n int) {
	return sovStakingRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BTCStakingRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStakingRecord
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
			return fmt.Errorf("proto: BTCStakingRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BTCStakingRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakingRecord
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
				return ErrInvalidLengthStakingRecord
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStakingRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakingRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintToAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakingRecord
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
				return ErrInvalidLengthStakingRecord
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStakingRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintToAddr = append(m.MintToAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.MintToAddr == nil {
				m.MintToAddr = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcReceiverName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakingRecord
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
				return ErrInvalidLengthStakingRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakingRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BtcReceiverName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcReceiverAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakingRecord
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
				return ErrInvalidLengthStakingRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakingRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BtcReceiverAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStakingRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStakingRecord
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
func skipStakingRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStakingRecord
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
					return 0, ErrIntOverflowStakingRecord
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
					return 0, ErrIntOverflowStakingRecord
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
				return 0, ErrInvalidLengthStakingRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStakingRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStakingRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStakingRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStakingRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStakingRecord = fmt.Errorf("proto: unexpected end of group")
)
