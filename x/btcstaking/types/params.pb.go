// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/btcstaking/v1/params.proto

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

type Receiver struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	// like 0xBAb28FF7659481F1c8516f616A576339936AFB06
	EthAddr string `protobuf:"bytes,3,opt,name=eth_addr,json=ethAddr,proto3" json:"eth_addr,omitempty"`
}

func (m *Receiver) Reset()         { *m = Receiver{} }
func (m *Receiver) String() string { return proto.CompactTextString(m) }
func (*Receiver) ProtoMessage()    {}
func (*Receiver) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc0c3789eee73a9d, []int{0}
}
func (m *Receiver) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Receiver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Receiver.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Receiver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receiver.Merge(m, src)
}
func (m *Receiver) XXX_Size() int {
	return m.Size()
}
func (m *Receiver) XXX_DiscardUnknown() {
	xxx_messageInfo_Receiver.DiscardUnknown(m)
}

var xxx_messageInfo_Receiver proto.InternalMessageInfo

func (m *Receiver) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Receiver) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Receiver) GetEthAddr() string {
	if m != nil {
		return m.EthAddr
	}
	return ""
}

// GenesisState defines the btcstaking module's genesis state.
type Params struct {
	Receivers             []*Receiver `protobuf:"bytes,1,rep,name=receivers,proto3" json:"receivers,omitempty"`
	BtcConfirmationsDepth uint32      `protobuf:"varint,2,opt,name=btc_confirmations_depth,json=btcConfirmationsDepth,proto3" json:"btc_confirmations_depth,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc0c3789eee73a9d, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetReceivers() []*Receiver {
	if m != nil {
		return m.Receivers
	}
	return nil
}

func (m *Params) GetBtcConfirmationsDepth() uint32 {
	if m != nil {
		return m.BtcConfirmationsDepth
	}
	return 0
}

func init() {
	proto.RegisterType((*Receiver)(nil), "lorenzo.btcstaking.v1.Receiver")
	proto.RegisterType((*Params)(nil), "lorenzo.btcstaking.v1.Params")
}

func init() {
	proto.RegisterFile("lorenzo/btcstaking/v1/params.proto", fileDescriptor_cc0c3789eee73a9d)
}

var fileDescriptor_cc0c3789eee73a9d = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x63, 0x8a, 0x4a, 0x6b, 0xc4, 0x62, 0x51, 0x11, 0x18, 0x4c, 0x95, 0xa9, 0x0b, 0xb6,
	0x0a, 0x52, 0x37, 0x06, 0x7e, 0x46, 0x90, 0xa2, 0x8c, 0x2c, 0x91, 0xe3, 0x98, 0x24, 0xa2, 0x89,
	0x23, 0xe7, 0x12, 0x01, 0x0b, 0xaf, 0xc0, 0x63, 0x31, 0x76, 0x64, 0x44, 0xc9, 0x8b, 0xa0, 0x38,
	0xad, 0xda, 0x81, 0xed, 0xe8, 0xdc, 0x63, 0xeb, 0xd3, 0x87, 0xbd, 0xa5, 0x36, 0xaa, 0xf8, 0xd0,
	0x3c, 0x02, 0x59, 0x81, 0x78, 0xc9, 0x8a, 0x84, 0xd7, 0x73, 0x5e, 0x0a, 0x23, 0xf2, 0x8a, 0x95,
	0x46, 0x83, 0x26, 0x93, 0xf5, 0x86, 0x6d, 0x37, 0xac, 0x9e, 0x9f, 0x1d, 0x27, 0x3a, 0xd1, 0x76,
	0xc1, 0xbb, 0xd4, 0x8f, 0xbd, 0x47, 0x3c, 0x0a, 0x94, 0x54, 0x59, 0xad, 0x0c, 0x21, 0x78, 0xbf,
	0x10, 0xb9, 0x72, 0xd1, 0x14, 0xcd, 0xc6, 0x81, 0xcd, 0x5d, 0x27, 0xe2, 0xd8, 0xb8, 0x7b, 0x7d,
	0xd7, 0x65, 0x72, 0x8a, 0x47, 0x0a, 0xd2, 0xd0, 0xf6, 0x03, 0xdb, 0x1f, 0x28, 0x48, 0x6f, 0xe2,
	0xd8, 0x78, 0x9f, 0x78, 0xe8, 0x5b, 0x16, 0x72, 0x8d, 0xc7, 0x66, 0xfd, 0x71, 0xe5, 0xa2, 0xe9,
	0x60, 0x76, 0x78, 0x79, 0xce, 0xfe, 0x25, 0x63, 0x1b, 0x80, 0x60, 0xfb, 0x82, 0x2c, 0xf0, 0x49,
	0x04, 0x32, 0x94, 0xba, 0x78, 0xce, 0x4c, 0x2e, 0x20, 0xd3, 0x45, 0x15, 0xc6, 0xaa, 0x84, 0xd4,
	0xa2, 0x1c, 0x05, 0x93, 0x08, 0xe4, 0xdd, 0xee, 0xf5, 0xbe, 0x3b, 0xde, 0xfa, 0xdf, 0x0d, 0x45,
	0xab, 0x86, 0xa2, 0xdf, 0x86, 0xa2, 0xaf, 0x96, 0x3a, 0xab, 0x96, 0x3a, 0x3f, 0x2d, 0x75, 0x9e,
	0x16, 0x49, 0x06, 0xe9, 0x6b, 0xc4, 0xa4, 0xce, 0xf9, 0x43, 0xcf, 0x71, 0xe1, 0x77, 0x0e, 0xa4,
	0x5e, 0xf2, 0x8d, 0xd6, 0xb7, 0x5d, 0xb1, 0xf0, 0x5e, 0xaa, 0x2a, 0x1a, 0x5a, 0x51, 0x57, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x56, 0x7c, 0x93, 0x3c, 0x7b, 0x01, 0x00, 0x00,
}

func (m *Receiver) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Receiver) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Receiver) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EthAddr) > 0 {
		i -= len(m.EthAddr)
		copy(dAtA[i:], m.EthAddr)
		i = encodeVarintParams(dAtA, i, uint64(len(m.EthAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BtcConfirmationsDepth != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BtcConfirmationsDepth))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Receivers) > 0 {
		for iNdEx := len(m.Receivers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Receivers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Receiver) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.EthAddr)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Receivers) > 0 {
		for _, e := range m.Receivers {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.BtcConfirmationsDepth != 0 {
		n += 1 + sovParams(uint64(m.BtcConfirmationsDepth))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Receiver) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Receiver: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Receiver: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receivers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receivers = append(m.Receivers, &Receiver{})
			if err := m.Receivers[len(m.Receivers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcConfirmationsDepth", wireType)
			}
			m.BtcConfirmationsDepth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BtcConfirmationsDepth |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
