// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/bnblightclient/v1/client.proto

package types

import (
	fmt "fmt"
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

type Header struct {
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c005374b948758, []int{0}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

type Header_Header struct {
	// header defines the bnb header bytes
	RawHeader []byte `protobuf:"bytes,1,opt,name=raw_header,json=rawHeader,proto3" json:"raw_header,omitempty"`
	// hash defines the bnb header hash
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// number defines the block number
	Number uint64 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	// receipt_root defines the receipts merkle root hash
	ReceiptRoot []byte `protobuf:"bytes,4,opt,name=receipt_root,json=receiptRoot,proto3" json:"receipt_root,omitempty"`
}

func (m *Header_Header) Reset()         { *m = Header_Header{} }
func (m *Header_Header) String() string { return proto.CompactTextString(m) }
func (*Header_Header) ProtoMessage()    {}
func (*Header_Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c005374b948758, []int{0, 0}
}
func (m *Header_Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header_Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header_Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header_Header.Merge(m, src)
}
func (m *Header_Header) XXX_Size() int {
	return m.Size()
}
func (m *Header_Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header_Header proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Header)(nil), "lorenzo.bnblightclient.v1.Header")
	proto.RegisterType((*Header_Header)(nil), "lorenzo.bnblightclient.v1.Header.Header")
}

func init() {
	proto.RegisterFile("lorenzo/bnblightclient/v1/client.proto", fileDescriptor_30c005374b948758)
}

var fileDescriptor_30c005374b948758 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0xc9, 0x2f, 0x4a,
	0xcd, 0xab, 0xca, 0xd7, 0x4f, 0xca, 0x4b, 0xca, 0xc9, 0x4c, 0xcf, 0x28, 0x49, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2f, 0x33, 0xd4, 0x87, 0xb0, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x24,
	0xa1, 0xea, 0xf4, 0x50, 0xd5, 0xe9, 0x95, 0x19, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x55,
	0xe9, 0x83, 0x58, 0x10, 0x0d, 0x4a, 0x6d, 0x8c, 0x5c, 0x6c, 0x1e, 0xa9, 0x89, 0x29, 0xa9, 0x45,
	0x52, 0x35, 0x30, 0x96, 0x90, 0x2c, 0x17, 0x57, 0x51, 0x62, 0x79, 0x7c, 0x06, 0x98, 0x27, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x13, 0xc4, 0x59, 0x94, 0x58, 0x0e, 0x95, 0x16, 0xe2, 0x62, 0xc9, 0x48,
	0x2c, 0xce, 0x90, 0x60, 0x02, 0x4b, 0x80, 0xd9, 0x42, 0x62, 0x5c, 0x6c, 0x79, 0xa5, 0xb9, 0x49,
	0xa9, 0x45, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x50, 0x9e, 0x90, 0x22, 0x17, 0x4f, 0x51,
	0x6a, 0x72, 0x6a, 0x66, 0x41, 0x49, 0x7c, 0x51, 0x7e, 0x7e, 0x89, 0x04, 0x0b, 0x58, 0x0f, 0x37,
	0x54, 0x2c, 0x28, 0x3f, 0xbf, 0xc4, 0x8a, 0xa5, 0x63, 0x81, 0x3c, 0x83, 0x53, 0xc8, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85,
	0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x59, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0xea, 0xfb, 0x40, 0xbc, 0xa7, 0x1b, 0x00, 0x72, 0x7c, 0x72, 0x7e, 0x8e, 0x3e,
	0x2c, 0x5c, 0x2a, 0xd0, 0x43, 0xa6, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x4b, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0x85, 0x42, 0x8d, 0x40, 0x01, 0x00, 0x00,
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Header_Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header_Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header_Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReceiptRoot) > 0 {
		i -= len(m.ReceiptRoot)
		copy(dAtA[i:], m.ReceiptRoot)
		i = encodeVarintClient(dAtA, i, uint64(len(m.ReceiptRoot)))
		i--
		dAtA[i] = 0x22
	}
	if m.Number != 0 {
		i = encodeVarintClient(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintClient(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RawHeader) > 0 {
		i -= len(m.RawHeader)
		copy(dAtA[i:], m.RawHeader)
		i = encodeVarintClient(dAtA, i, uint64(len(m.RawHeader)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClient(dAtA []byte, offset int, v uint64) int {
	offset -= sovClient(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Header_Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RawHeader)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	if m.Number != 0 {
		n += 1 + sovClient(uint64(m.Number))
	}
	l = len(m.ReceiptRoot)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	return n
}

func sovClient(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClient(x uint64) (n int) {
	return sovClient(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClient
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
func (m *Header_Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawHeader", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawHeader = append(m.RawHeader[:0], dAtA[iNdEx:postIndex]...)
			if m.RawHeader == nil {
				m.RawHeader = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiptRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiptRoot = append(m.ReceiptRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.ReceiptRoot == nil {
				m.ReceiptRoot = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClient
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
func skipClient(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClient
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
					return 0, ErrIntOverflowClient
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
					return 0, ErrIntOverflowClient
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
				return 0, ErrInvalidLengthClient
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClient
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClient
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClient        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClient          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClient = fmt.Errorf("proto: unexpected end of group")
)
