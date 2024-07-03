// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/plan/v1/plan.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type PlanStatus int32

const (
	PlanStatus_Pause   PlanStatus = 0
	PlanStatus_Unpause PlanStatus = 1
)

var PlanStatus_name = map[int32]string{
	0: "Pause",
	1: "Unpause",
}

var PlanStatus_value = map[string]int32{
	"Pause":   0,
	"Unpause": 1,
}

func (x PlanStatus) String() string {
	return proto.EnumName(PlanStatus_name, int32(x))
}

func (PlanStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df1b3d6ed2d06d8a, []int{0}
}

// Plan defines the details of a project
type Plan struct {
	Id                 uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               string                                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PlanDescUri        string                                 `protobuf:"bytes,3,opt,name=plan_desc_uri,json=planDescUri,proto3" json:"plan_desc_uri,omitempty"`
	AgentId            uint64                                 `protobuf:"varint,4,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	PlanStartBlock     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=plan_start_block,json=planStartBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"plan_start_block"`
	PeriodBlocks       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=period_blocks,json=periodBlocks,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"period_blocks"`
	YatContractAddress string                                 `protobuf:"bytes,7,opt,name=yat_contract_address,json=yatContractAddress,proto3" json:"yat_contract_address,omitempty"`
	ContractAddress    string                                 `protobuf:"bytes,8,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Enabled            PlanStatus                             `protobuf:"varint,9,opt,name=enabled,proto3,enum=lorenzo.plan.v1.PlanStatus" json:"enabled,omitempty"`
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

func (m *Plan) GetYatContractAddress() string {
	if m != nil {
		return m.YatContractAddress
	}
	return ""
}

func (m *Plan) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *Plan) GetEnabled() PlanStatus {
	if m != nil {
		return m.Enabled
	}
	return PlanStatus_Pause
}

func init() {
	proto.RegisterEnum("lorenzo.plan.v1.PlanStatus", PlanStatus_name, PlanStatus_value)
	proto.RegisterType((*Plan)(nil), "lorenzo.plan.v1.Plan")
}

func init() { proto.RegisterFile("lorenzo/plan/v1/plan.proto", fileDescriptor_df1b3d6ed2d06d8a) }

var fileDescriptor_df1b3d6ed2d06d8a = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xe3, 0x2e, 0x5b, 0x56, 0x8f, 0x75, 0x95, 0xb5, 0x83, 0x19, 0x52, 0x56, 0x4d, 0x08,
	0x15, 0xa4, 0x25, 0x1b, 0x88, 0x07, 0xa0, 0x70, 0x19, 0xe2, 0x10, 0xa5, 0xaa, 0x84, 0xb8, 0x44,
	0x8e, 0x6d, 0x85, 0xa8, 0xa9, 0x1d, 0xd9, 0x4e, 0x45, 0x79, 0x0a, 0x1e, 0x84, 0x07, 0xe9, 0xb1,
	0x47, 0xc4, 0xa1, 0x42, 0xed, 0x8b, 0x20, 0x3b, 0xad, 0x40, 0xe5, 0xc6, 0xc9, 0xfe, 0xbe, 0xff,
	0x3f, 0xbf, 0xe8, 0xfb, 0xfc, 0x87, 0x57, 0x95, 0x54, 0x5c, 0x7c, 0x95, 0x71, 0x5d, 0x11, 0x11,
	0xcf, 0xef, 0xdd, 0x19, 0xd5, 0x4a, 0x1a, 0x89, 0x2e, 0x76, 0x5a, 0xe4, 0x7a, 0xf3, 0xfb, 0xab,
	0xcb, 0x42, 0x16, 0xd2, 0x69, 0xb1, 0xbd, 0xb5, 0xb6, 0x9b, 0xef, 0x47, 0xd0, 0x4f, 0x2a, 0x22,
	0x50, 0x0f, 0x76, 0x4a, 0x86, 0xc1, 0x00, 0x0c, 0xfd, 0xb4, 0x53, 0x32, 0x84, 0xa0, 0x2f, 0xc8,
	0x8c, 0xe3, 0xce, 0x00, 0x0c, 0xbb, 0xa9, 0xbb, 0xa3, 0x1b, 0x78, 0x6e, 0x69, 0x19, 0xe3, 0x9a,
	0x66, 0x8d, 0x2a, 0xf1, 0x91, 0x13, 0xcf, 0x6c, 0xf3, 0x1d, 0xd7, 0x74, 0xa2, 0x4a, 0xf4, 0x18,
	0x9e, 0x92, 0x82, 0x0b, 0x93, 0x95, 0x0c, 0xfb, 0x8e, 0x16, 0xb8, 0xfa, 0x81, 0xa1, 0x8f, 0xb0,
	0xef, 0x3e, 0xd7, 0x86, 0x28, 0x93, 0xe5, 0x95, 0xa4, 0x53, 0x7c, 0x6c, 0x09, 0xa3, 0x68, 0xb9,
	0xbe, 0xf6, 0x7e, 0xae, 0xaf, 0x9f, 0x15, 0xa5, 0xf9, 0xdc, 0xe4, 0x11, 0x95, 0xb3, 0x98, 0x4a,
	0x3d, 0x93, 0x7a, 0x77, 0xdc, 0x6a, 0x36, 0x8d, 0xcd, 0xa2, 0xe6, 0x3a, 0x7a, 0x10, 0x26, 0xed,
	0x59, 0xce, 0xd8, 0x62, 0x46, 0x96, 0x82, 0xc6, 0xf0, 0xbc, 0xe6, 0xaa, 0x94, 0xac, 0xa5, 0x6a,
	0x7c, 0xf2, 0x5f, 0xd8, 0x47, 0x2d, 0xc4, 0x31, 0x35, 0xba, 0x83, 0x97, 0x0b, 0x62, 0x32, 0x2a,
	0x85, 0x51, 0x84, 0x9a, 0x8c, 0x30, 0xa6, 0xb8, 0xd6, 0x38, 0x70, 0x43, 0xa3, 0x05, 0x31, 0x6f,
	0x77, 0xd2, 0x9b, 0x56, 0x41, 0xcf, 0x61, 0xff, 0x1f, 0xf7, 0xa9, 0x73, 0x5f, 0xd0, 0x03, 0xeb,
	0x6b, 0x18, 0x70, 0x41, 0xf2, 0x8a, 0x33, 0xdc, 0x1d, 0x80, 0x61, 0xef, 0xe5, 0x93, 0xe8, 0xe0,
	0xc1, 0xa2, 0xa4, 0x9d, 0xd1, 0x34, 0x3a, 0xdd, 0x7b, 0x5f, 0x3c, 0x85, 0xf0, 0x4f, 0x1b, 0x75,
	0xe1, 0x71, 0x42, 0x1a, 0xcd, 0xfb, 0x1e, 0x3a, 0x83, 0xc1, 0x44, 0xd4, 0xae, 0x00, 0xa3, 0xf7,
	0xcb, 0x4d, 0x08, 0x56, 0x9b, 0x10, 0xfc, 0xda, 0x84, 0xe0, 0xdb, 0x36, 0xf4, 0x56, 0xdb, 0xd0,
	0xfb, 0xb1, 0x0d, 0xbd, 0x4f, 0x77, 0x7f, 0x6d, 0xe2, 0x43, 0xfb, 0xbf, 0xdb, 0xc4, 0x06, 0x81,
	0xca, 0x2a, 0xde, 0xa7, 0xe9, 0x4b, 0x9b, 0x27, 0xb7, 0x97, 0xfc, 0xc4, 0xe5, 0xe4, 0xd5, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x1a, 0xbf, 0x57, 0x6c, 0x02, 0x00, 0x00,
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
	if m.Enabled != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.Enabled))
		i--
		dAtA[i] = 0x48
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.YatContractAddress) > 0 {
		i -= len(m.YatContractAddress)
		copy(dAtA[i:], m.YatContractAddress)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.YatContractAddress)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size := m.PeriodBlocks.Size()
		i -= size
		if _, err := m.PeriodBlocks.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPlan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.PlanStartBlock.Size()
		i -= size
		if _, err := m.PlanStartBlock.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPlan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.AgentId != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.AgentId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PlanDescUri) > 0 {
		i -= len(m.PlanDescUri)
		copy(dAtA[i:], m.PlanDescUri)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.PlanDescUri)))
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
	l = len(m.PlanDescUri)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.AgentId != 0 {
		n += 1 + sovPlan(uint64(m.AgentId))
	}
	l = m.PlanStartBlock.Size()
	n += 1 + l + sovPlan(uint64(l))
	l = m.PeriodBlocks.Size()
	n += 1 + l + sovPlan(uint64(l))
	l = len(m.YatContractAddress)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.Enabled != 0 {
		n += 1 + sovPlan(uint64(m.Enabled))
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
		case 4:
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanStartBlock", wireType)
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
			if err := m.PlanStartBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodBlocks", wireType)
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
			if err := m.PeriodBlocks.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field YatContractAddress", wireType)
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
			m.YatContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
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
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			m.Enabled = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Enabled |= PlanStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
