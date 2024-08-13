// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/token/v1/token.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
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

// Source defines the source type of token asset, if source is:
// - module: token origin is sdk module;
// - contract: token origin is erc20 contract;
type Source int32

const (
	// undefined source
	OWNER_UNDEFINED Source = 0
	// token source is module
	OWNER_MODULE Source = 1
	// token source is erc20 contract
	OWNER_CONTRACT Source = 2
)

var Source_name = map[int32]string{
	0: "OWNER_UNDEFINED",
	1: "OWNER_MODULE",
	2: "OWNER_CONTRACT",
}

var Source_value = map[string]int32{
	"OWNER_UNDEFINED": 0,
	"OWNER_MODULE":    1,
	"OWNER_CONTRACT":  2,
}

func (x Source) String() string {
	return proto.EnumName(Source_name, int32(x))
}

func (Source) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1095540cf3b10f35, []int{0}
}

// TokenPair defines a pairing of a cosmos coin and an erc20 token
type TokenPair struct {
	// erc20 contract hex format address
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// sdk coin base denomination
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	// allows for token conversion
	Enabled bool `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// source of token asset
	Source Source `protobuf:"varint,4,opt,name=source,proto3,enum=lorenzo.token.v1.Source" json:"source,omitempty"`
}

func (m *TokenPair) Reset()         { *m = TokenPair{} }
func (m *TokenPair) String() string { return proto.CompactTextString(m) }
func (*TokenPair) ProtoMessage()    {}
func (*TokenPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_1095540cf3b10f35, []int{0}
}
func (m *TokenPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPair.Merge(m, src)
}
func (m *TokenPair) XXX_Size() int {
	return m.Size()
}
func (m *TokenPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPair.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPair proto.InternalMessageInfo

func (m *TokenPair) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *TokenPair) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *TokenPair) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *TokenPair) GetSource() Source {
	if m != nil {
		return m.Source
	}
	return OWNER_UNDEFINED
}

func init() {
	proto.RegisterEnum("lorenzo.token.v1.Source", Source_name, Source_value)
	proto.RegisterType((*TokenPair)(nil), "lorenzo.token.v1.TokenPair")
}

func init() { proto.RegisterFile("lorenzo/token/v1/token.proto", fileDescriptor_1095540cf3b10f35) }

var fileDescriptor_1095540cf3b10f35 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0xc9, 0x2f, 0x4a,
	0xcd, 0xab, 0xca, 0xd7, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x2f, 0x33, 0x84, 0x30, 0xf4, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0xa0, 0xb2, 0x7a, 0x10, 0xc1, 0x32, 0x43, 0x29, 0xb9, 0xe4,
	0xfc, 0xe2, 0xdc, 0xfc, 0x62, 0xfd, 0xa4, 0xc4, 0xbc, 0x6c, 0xfd, 0x32, 0xc3, 0xa4, 0xd4, 0x92,
	0x44, 0x43, 0x30, 0x07, 0xa2, 0x43, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xcc, 0xd4, 0x07, 0xb1,
	0x20, 0xa2, 0x4a, 0xf3, 0x18, 0xb9, 0x38, 0x43, 0x40, 0x46, 0x04, 0x24, 0x66, 0x16, 0x09, 0x69,
	0x72, 0x09, 0x24, 0xe7, 0xe7, 0x95, 0x14, 0x25, 0x26, 0x97, 0xc4, 0x27, 0xa6, 0xa4, 0x14, 0xa5,
	0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xf1, 0xc3, 0xc4, 0x1d, 0x21, 0xc2, 0x42,
	0x22, 0x5c, 0xac, 0x29, 0xa9, 0x79, 0xf9, 0xb9, 0x12, 0x4c, 0x60, 0x79, 0x08, 0x47, 0x48, 0x82,
	0x8b, 0x3d, 0x35, 0x2f, 0x31, 0x29, 0x27, 0x35, 0x45, 0x82, 0x59, 0x81, 0x51, 0x83, 0x23, 0x08,
	0xc6, 0x15, 0x32, 0xe0, 0x62, 0x2b, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x95, 0x60, 0x51, 0x60, 0xd4,
	0xe0, 0x33, 0x92, 0xd0, 0x43, 0xf7, 0x81, 0x5e, 0x30, 0x58, 0x3e, 0x08, 0xaa, 0xce, 0x8a, 0xe5,
	0xc5, 0x02, 0x79, 0x46, 0x2d, 0x4f, 0x2e, 0x36, 0x88, 0xb8, 0x90, 0x30, 0x17, 0xbf, 0x7f, 0xb8,
	0x9f, 0x6b, 0x50, 0x7c, 0xa8, 0x9f, 0x8b, 0xab, 0x9b, 0xa7, 0x9f, 0xab, 0x8b, 0x00, 0x83, 0x90,
	0x00, 0x17, 0x0f, 0x44, 0xd0, 0xd7, 0xdf, 0x25, 0xd4, 0xc7, 0x55, 0x80, 0x51, 0x48, 0x88, 0x8b,
	0x0f, 0x22, 0xe2, 0xec, 0xef, 0x17, 0x12, 0xe4, 0xe8, 0x1c, 0x22, 0xc0, 0x24, 0xc5, 0xd2, 0xb1,
	0x58, 0x8e, 0xc1, 0xc9, 0xef, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x4c,
	0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x7d, 0x20, 0xce, 0xd2, 0x0d,
	0x00, 0x85, 0x4f, 0x72, 0x7e, 0x8e, 0x3e, 0x2c, 0x1e, 0xca, 0x8c, 0xf5, 0x2b, 0xa0, 0x91, 0x51,
	0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x42, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x3a, 0x26, 0xf6, 0xaf, 0xaa, 0x01, 0x00, 0x00,
}

func (this *TokenPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenPair)
	if !ok {
		that2, ok := that.(TokenPair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ContractAddress != that1.ContractAddress {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if this.Source != that1.Source {
		return false
	}
	return true
}
func (m *TokenPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Source != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.Source))
		i--
		dAtA[i] = 0x20
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintToken(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if m.Source != 0 {
		n += 1 + sovToken(uint64(m.Source))
	}
	return n
}

func sovToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: TokenPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			m.Source = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Source |= Source(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToken = fmt.Errorf("proto: unexpected end of group")
)
