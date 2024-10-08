// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/agent/v1/agent.proto

package types

import (
	fmt "fmt"
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

// Agent defines the details of a project
type Agent struct {
	// id is the unique identifier of the agent
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// agent name,required
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// btc_receiving_address is agent’s fund escrow address,required
	BtcReceivingAddress string `protobuf:"bytes,3,opt,name=btc_receiving_address,json=btcReceivingAddress,proto3" json:"btc_receiving_address,omitempty"`
	// like 0xBAb28FF7659481F1c8516f616A576339936AFB06
	EthAddr string `protobuf:"bytes,4,opt,name=eth_addr,json=ethAddr,proto3" json:"eth_addr,omitempty"`
	// description is a brief description of the agent, optional
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// url is the agent's link, used for detailed introduction, optional
	Url string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *Agent) Reset()         { *m = Agent{} }
func (m *Agent) String() string { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()    {}
func (*Agent) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4eba5995b94e340, []int{0}
}
func (m *Agent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Agent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Agent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Agent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Agent.Merge(m, src)
}
func (m *Agent) XXX_Size() int {
	return m.Size()
}
func (m *Agent) XXX_DiscardUnknown() {
	xxx_messageInfo_Agent.DiscardUnknown(m)
}

var xxx_messageInfo_Agent proto.InternalMessageInfo

func (m *Agent) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Agent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Agent) GetBtcReceivingAddress() string {
	if m != nil {
		return m.BtcReceivingAddress
	}
	return ""
}

func (m *Agent) GetEthAddr() string {
	if m != nil {
		return m.EthAddr
	}
	return ""
}

func (m *Agent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Agent) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*Agent)(nil), "lorenzo.agent.v1.Agent")
}

func init() { proto.RegisterFile("lorenzo/agent/v1/agent.proto", fileDescriptor_c4eba5995b94e340) }

var fileDescriptor_c4eba5995b94e340 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x9b, 0xae, 0x9b, 0x1a, 0x41, 0x46, 0x44, 0x88, 0x20, 0xa1, 0x78, 0xda, 0xc5, 0x86,
	0x39, 0x5f, 0x60, 0x9e, 0x45, 0xa4, 0x47, 0x2f, 0xa3, 0x4d, 0x3e, 0xda, 0x40, 0x97, 0x94, 0x34,
	0x2b, 0xea, 0x53, 0xf8, 0x24, 0x3e, 0x87, 0xc7, 0x1d, 0x3d, 0x4a, 0xfb, 0x22, 0xd2, 0x74, 0x03,
	0x6f, 0xff, 0xef, 0xfb, 0xfd, 0x2f, 0xff, 0x1f, 0xbe, 0xa9, 0x8c, 0x05, 0xfd, 0x61, 0x78, 0x56,
	0x80, 0x76, 0xbc, 0x5d, 0x8e, 0x21, 0xa9, 0xad, 0x71, 0x86, 0xcc, 0x0f, 0x34, 0x19, 0x9f, 0xed,
	0xf2, 0xf6, 0x0b, 0xe1, 0xe9, 0x7a, 0x38, 0xc8, 0x05, 0x0e, 0x95, 0xa4, 0x28, 0x46, 0x8b, 0x28,
	0x0d, 0x95, 0x24, 0x04, 0x47, 0x3a, 0xdb, 0x02, 0x0d, 0x63, 0xb4, 0x38, 0x4b, 0x7d, 0x26, 0xf7,
	0xf8, 0x2a, 0x77, 0x62, 0x63, 0x41, 0x80, 0x6a, 0x95, 0x2e, 0x36, 0x99, 0x94, 0x16, 0x9a, 0x86,
	0x4e, 0x7c, 0xe9, 0x32, 0x77, 0x22, 0x3d, 0xb2, 0xf5, 0x88, 0xc8, 0x35, 0x3e, 0x05, 0x57, 0xfa,
	0x26, 0x8d, 0x7c, 0xed, 0x04, 0x5c, 0x39, 0x50, 0x12, 0xe3, 0x73, 0x09, 0x8d, 0xb0, 0xaa, 0x76,
	0xca, 0x68, 0x3a, 0xf5, 0xf4, 0xff, 0x8b, 0xcc, 0xf1, 0x64, 0x67, 0x2b, 0x3a, 0xf3, 0x64, 0x88,
	0x8f, 0xcf, 0xdf, 0x1d, 0x43, 0xfb, 0x8e, 0xa1, 0xdf, 0x8e, 0xa1, 0xcf, 0x9e, 0x05, 0xfb, 0x9e,
	0x05, 0x3f, 0x3d, 0x0b, 0x5e, 0x1f, 0x0a, 0xe5, 0xca, 0x5d, 0x9e, 0x08, 0xb3, 0xe5, 0x4f, 0xe3,
	0xce, 0xbb, 0x97, 0x61, 0xb6, 0x30, 0x15, 0x3f, 0x6a, 0x69, 0x57, 0xfc, 0xed, 0xe0, 0xc6, 0xbd,
	0xd7, 0xd0, 0xe4, 0x33, 0x6f, 0x66, 0xf5, 0x17, 0x00, 0x00, 0xff, 0xff, 0x93, 0x23, 0x3f, 0xb9,
	0x39, 0x01, 0x00, 0x00,
}

func (m *Agent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Agent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Agent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.EthAddr) > 0 {
		i -= len(m.EthAddr)
		copy(dAtA[i:], m.EthAddr)
		i = encodeVarintAgent(dAtA, i, uint64(len(m.EthAddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BtcReceivingAddress) > 0 {
		i -= len(m.BtcReceivingAddress)
		copy(dAtA[i:], m.BtcReceivingAddress)
		i = encodeVarintAgent(dAtA, i, uint64(len(m.BtcReceivingAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAgent(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAgent(dAtA []byte, offset int, v uint64) int {
	offset -= sovAgent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Agent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAgent(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.BtcReceivingAddress)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.EthAddr)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	return n
}

func sovAgent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAgent(x uint64) (n int) {
	return sovAgent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Agent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
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
			return fmt.Errorf("proto: Agent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Agent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
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
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcReceivingAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
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
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BtcReceivingAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
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
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
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
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
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
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAgent
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
func skipAgent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAgent
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
					return 0, ErrIntOverflowAgent
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
					return 0, ErrIntOverflowAgent
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
				return 0, ErrInvalidLengthAgent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAgent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAgent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAgent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAgent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAgent = fmt.Errorf("proto: unexpected end of group")
)
