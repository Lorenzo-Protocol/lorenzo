// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/plan/v1/plan.proto

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

// Plan defines the details of a project
type Plan struct {
	Id                    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol                string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	PlanDescUri           string `protobuf:"bytes,4,opt,name=plan_desc_uri,json=planDescUri,proto3" json:"plan_desc_uri,omitempty"`
	AgentId               uint64 `protobuf:"varint,5,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	SubscriptionStartTime uint64 `protobuf:"varint,6,opt,name=subscription_start_time,json=subscriptionStartTime,proto3" json:"subscription_start_time,omitempty"`
	SubscriptionEndTime   uint64 `protobuf:"varint,7,opt,name=subscription_end_time,json=subscriptionEndTime,proto3" json:"subscription_end_time,omitempty"`
	EndTime               uint64 `protobuf:"varint,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	MerkleRoot            string `protobuf:"bytes,9,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	ContractAddress       string `protobuf:"bytes,10,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_df1b3d6ed2d06d8a, []int{0}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Plan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Plan) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Plan) GetPlanDescUri() string {
	if m != nil {
		return m.PlanDescUri
	}
	return ""
}

func (m *Plan) GetAgentId() uint64 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *Plan) GetSubscriptionStartTime() uint64 {
	if m != nil {
		return m.SubscriptionStartTime
	}
	return 0
}

func (m *Plan) GetSubscriptionEndTime() uint64 {
	if m != nil {
		return m.SubscriptionEndTime
	}
	return 0
}

func (m *Plan) GetEndTime() uint64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Plan) GetMerkleRoot() string {
	if m != nil {
		return m.MerkleRoot
	}
	return ""
}

func (m *Plan) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Plan)(nil), "lorenzo.plan.v1.Plan")
}

func init() { proto.RegisterFile("lorenzo/plan/v1/plan.proto", fileDescriptor_df1b3d6ed2d06d8a) }

var fileDescriptor_df1b3d6ed2d06d8a = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbd, 0x4e, 0xeb, 0x30,
	0x14, 0xc7, 0x9b, 0xdc, 0xdc, 0x7e, 0x9c, 0xea, 0xde, 0x22, 0x23, 0xc0, 0x65, 0x08, 0x55, 0xa7,
	0x32, 0xd0, 0x50, 0x90, 0xd8, 0x41, 0x30, 0x80, 0x18, 0xaa, 0x02, 0x0b, 0x8b, 0x95, 0xd8, 0x56,
	0xb1, 0x48, 0xec, 0xc8, 0x76, 0x2b, 0xca, 0x53, 0xf0, 0x58, 0x8c, 0x1d, 0x18, 0x18, 0x51, 0xfb,
	0x22, 0x28, 0x4e, 0x8b, 0xca, 0x64, 0x9f, 0xdf, 0xef, 0x1c, 0x7f, 0xe8, 0x0f, 0xfb, 0xa9, 0xd2,
	0x5c, 0xbe, 0xaa, 0x28, 0x4f, 0x63, 0x19, 0x4d, 0x07, 0x6e, 0xed, 0xe7, 0x5a, 0x59, 0x85, 0x5a,
	0x2b, 0xd7, 0x77, 0x6c, 0x3a, 0xe8, 0x7e, 0xf8, 0x10, 0x0c, 0xd3, 0x58, 0xa2, 0xff, 0xe0, 0x0b,
	0x86, 0xbd, 0x8e, 0xd7, 0x0b, 0x46, 0xbe, 0x60, 0x08, 0x41, 0x20, 0xe3, 0x8c, 0x63, 0xbf, 0xe3,
	0xf5, 0x1a, 0x23, 0xb7, 0x47, 0xbb, 0x50, 0x35, 0xb3, 0x2c, 0x51, 0x29, 0xfe, 0xe3, 0xe8, 0xaa,
	0x42, 0x5d, 0xf8, 0x57, 0x9c, 0x47, 0x18, 0x37, 0x94, 0x4c, 0xb4, 0xc0, 0x81, 0xd3, 0xcd, 0x02,
	0x5e, 0x72, 0x43, 0x1f, 0xb4, 0x40, 0x6d, 0xa8, 0xc7, 0x63, 0x2e, 0x2d, 0x11, 0x0c, 0xff, 0x75,
	0xb7, 0xd4, 0x5c, 0x7d, 0xcd, 0xd0, 0x19, 0xec, 0x99, 0x49, 0x62, 0xa8, 0x16, 0xb9, 0x15, 0x4a,
	0x12, 0x63, 0x63, 0x6d, 0x89, 0x15, 0x19, 0xc7, 0x55, 0xd7, 0xb9, 0xb3, 0xa9, 0xef, 0x0a, 0x7b,
	0x2f, 0x32, 0x8e, 0x4e, 0xe0, 0x97, 0x20, 0x5c, 0xb2, 0x72, 0xaa, 0xe6, 0xa6, 0xb6, 0x37, 0xe5,
	0x95, 0x64, 0x6e, 0xa6, 0x0d, 0xf5, 0x9f, 0xb6, 0x7a, 0xf9, 0x0c, 0xbe, 0x52, 0x07, 0xd0, 0xcc,
	0xb8, 0x7e, 0x4e, 0x39, 0xd1, 0x4a, 0x59, 0xdc, 0x70, 0x7f, 0x80, 0x12, 0x8d, 0x94, 0xb2, 0xe8,
	0x10, 0xb6, 0xa8, 0x92, 0x56, 0xc7, 0xd4, 0x92, 0x98, 0x31, 0xcd, 0x8d, 0xc1, 0xe0, 0xba, 0x5a,
	0x6b, 0x7e, 0x5e, 0xe2, 0x8b, 0x9b, 0xf7, 0x45, 0xe8, 0xcd, 0x17, 0xa1, 0xf7, 0xb5, 0x08, 0xbd,
	0xb7, 0x65, 0x58, 0x99, 0x2f, 0xc3, 0xca, 0xe7, 0x32, 0xac, 0x3c, 0x1e, 0x8f, 0x85, 0x7d, 0x9a,
	0x24, 0x7d, 0xaa, 0xb2, 0xe8, 0xb6, 0x0c, 0xe3, 0x68, 0x58, 0x64, 0x43, 0x55, 0x1a, 0xad, 0x93,
	0x7b, 0x29, 0xb3, 0xb3, 0xb3, 0x9c, 0x9b, 0xa4, 0xea, 0xa2, 0x3b, 0xfd, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xe8, 0x1d, 0x05, 0x75, 0xd8, 0x01, 0x00, 0x00,
}

func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.MerkleRoot) > 0 {
		i -= len(m.MerkleRoot)
		copy(dAtA[i:], m.MerkleRoot)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.MerkleRoot)))
		i--
		dAtA[i] = 0x4a
	}
	if m.EndTime != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x40
	}
	if m.SubscriptionEndTime != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.SubscriptionEndTime))
		i--
		dAtA[i] = 0x38
	}
	if m.SubscriptionStartTime != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.SubscriptionStartTime))
		i--
		dAtA[i] = 0x30
	}
	if m.AgentId != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.AgentId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PlanDescUri) > 0 {
		i -= len(m.PlanDescUri)
		copy(dAtA[i:], m.PlanDescUri)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.PlanDescUri)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlan(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPlan(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	l = len(m.PlanDescUri)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.AgentId != 0 {
		n += 1 + sovPlan(uint64(m.AgentId))
	}
	if m.SubscriptionStartTime != 0 {
		n += 1 + sovPlan(uint64(m.SubscriptionStartTime))
	}
	if m.SubscriptionEndTime != 0 {
		n += 1 + sovPlan(uint64(m.SubscriptionEndTime))
	}
	if m.EndTime != 0 {
		n += 1 + sovPlan(uint64(m.EndTime))
	}
	l = len(m.MerkleRoot)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	return n
}

func sovPlan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlan(x uint64) (n int) {
	return sovPlan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
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
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanDescUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlanDescUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentId", wireType)
			}
			m.AgentId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AgentId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionStartTime", wireType)
			}
			m.SubscriptionStartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubscriptionStartTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionEndTime", wireType)
			}
			m.SubscriptionEndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubscriptionEndTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkleRoot", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MerkleRoot = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlan
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
func skipPlan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
				return 0, ErrInvalidLengthPlan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlan = fmt.Errorf("proto: unexpected end of group")
)