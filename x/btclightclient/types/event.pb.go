// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/btclightclient/v1/event.proto

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

// The header included in the event is the block in the history
// of the current mainchain to which we are rolling back to.
// In other words, there is one rollback event emitted per re-org, to the
// greatest common ancestor of the old and the new fork.
type EventBTCRollBack struct {
	Header *BTCHeaderInfo `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (m *EventBTCRollBack) Reset()         { *m = EventBTCRollBack{} }
func (m *EventBTCRollBack) String() string { return proto.CompactTextString(m) }
func (*EventBTCRollBack) ProtoMessage()    {}
func (*EventBTCRollBack) Descriptor() ([]byte, []int) {
	return fileDescriptor_19c8f17b8989db54, []int{0}
}
func (m *EventBTCRollBack) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBTCRollBack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBTCRollBack.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBTCRollBack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBTCRollBack.Merge(m, src)
}
func (m *EventBTCRollBack) XXX_Size() int {
	return m.Size()
}
func (m *EventBTCRollBack) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBTCRollBack.DiscardUnknown(m)
}

var xxx_messageInfo_EventBTCRollBack proto.InternalMessageInfo

func (m *EventBTCRollBack) GetHeader() *BTCHeaderInfo {
	if m != nil {
		return m.Header
	}
	return nil
}

// EventBTCRollForward is emitted on Msg/InsertHeader
// The header included in the event is the one the main chain is extended with.
// In the event of a reorg, each block on the new fork that comes after
// the greatest common ancestor will have a corresponding roll forward event.
type EventBTCRollForward struct {
	Header *BTCHeaderInfo `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (m *EventBTCRollForward) Reset()         { *m = EventBTCRollForward{} }
func (m *EventBTCRollForward) String() string { return proto.CompactTextString(m) }
func (*EventBTCRollForward) ProtoMessage()    {}
func (*EventBTCRollForward) Descriptor() ([]byte, []int) {
	return fileDescriptor_19c8f17b8989db54, []int{1}
}
func (m *EventBTCRollForward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBTCRollForward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBTCRollForward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBTCRollForward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBTCRollForward.Merge(m, src)
}
func (m *EventBTCRollForward) XXX_Size() int {
	return m.Size()
}
func (m *EventBTCRollForward) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBTCRollForward.DiscardUnknown(m)
}

var xxx_messageInfo_EventBTCRollForward proto.InternalMessageInfo

func (m *EventBTCRollForward) GetHeader() *BTCHeaderInfo {
	if m != nil {
		return m.Header
	}
	return nil
}

// EventBTCHeaderInserted is emitted on Msg/InsertHeader
// The header included in the event is the one that was added to the
// on chain BTC storage.
type EventBTCHeaderInserted struct {
	Header *BTCHeaderInfo `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (m *EventBTCHeaderInserted) Reset()         { *m = EventBTCHeaderInserted{} }
func (m *EventBTCHeaderInserted) String() string { return proto.CompactTextString(m) }
func (*EventBTCHeaderInserted) ProtoMessage()    {}
func (*EventBTCHeaderInserted) Descriptor() ([]byte, []int) {
	return fileDescriptor_19c8f17b8989db54, []int{2}
}
func (m *EventBTCHeaderInserted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBTCHeaderInserted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBTCHeaderInserted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBTCHeaderInserted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBTCHeaderInserted.Merge(m, src)
}
func (m *EventBTCHeaderInserted) XXX_Size() int {
	return m.Size()
}
func (m *EventBTCHeaderInserted) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBTCHeaderInserted.DiscardUnknown(m)
}

var xxx_messageInfo_EventBTCHeaderInserted proto.InternalMessageInfo

func (m *EventBTCHeaderInserted) GetHeader() *BTCHeaderInfo {
	if m != nil {
		return m.Header
	}
	return nil
}

type EventBTCFeeRateUpdated struct {
	FeeRate uint64 `protobuf:"varint,1,opt,name=fee_rate,json=feeRate,proto3" json:"fee_rate,omitempty"`
}

func (m *EventBTCFeeRateUpdated) Reset()         { *m = EventBTCFeeRateUpdated{} }
func (m *EventBTCFeeRateUpdated) String() string { return proto.CompactTextString(m) }
func (*EventBTCFeeRateUpdated) ProtoMessage()    {}
func (*EventBTCFeeRateUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_19c8f17b8989db54, []int{3}
}
func (m *EventBTCFeeRateUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBTCFeeRateUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBTCFeeRateUpdated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBTCFeeRateUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBTCFeeRateUpdated.Merge(m, src)
}
func (m *EventBTCFeeRateUpdated) XXX_Size() int {
	return m.Size()
}
func (m *EventBTCFeeRateUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBTCFeeRateUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_EventBTCFeeRateUpdated proto.InternalMessageInfo

func (m *EventBTCFeeRateUpdated) GetFeeRate() uint64 {
	if m != nil {
		return m.FeeRate
	}
	return 0
}

func init() {
	proto.RegisterType((*EventBTCRollBack)(nil), "lorenzo.btclightclient.v1.EventBTCRollBack")
	proto.RegisterType((*EventBTCRollForward)(nil), "lorenzo.btclightclient.v1.EventBTCRollForward")
	proto.RegisterType((*EventBTCHeaderInserted)(nil), "lorenzo.btclightclient.v1.EventBTCHeaderInserted")
	proto.RegisterType((*EventBTCFeeRateUpdated)(nil), "lorenzo.btclightclient.v1.EventBTCFeeRateUpdated")
}

func init() {
	proto.RegisterFile("lorenzo/btclightclient/v1/event.proto", fileDescriptor_19c8f17b8989db54)
}

var fileDescriptor_19c8f17b8989db54 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0xc9, 0x2f, 0x4a,
	0xcd, 0xab, 0xca, 0xd7, 0x4f, 0x2a, 0x49, 0xce, 0xc9, 0x4c, 0xcf, 0x00, 0x91, 0xa9, 0x79, 0x25,
	0xfa, 0x65, 0x86, 0xfa, 0xa9, 0x65, 0xa9, 0x79, 0x25, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x92, 0x50, 0x65, 0x7a, 0xa8, 0xca, 0xf4, 0xca, 0x0c, 0xa5, 0xf4, 0x70, 0x9b, 0x80, 0xa6, 0x18,
	0x6c, 0x94, 0x52, 0x08, 0x97, 0x80, 0x2b, 0xc8, 0x64, 0xa7, 0x10, 0xe7, 0xa0, 0xfc, 0x9c, 0x1c,
	0xa7, 0xc4, 0xe4, 0x6c, 0x21, 0x07, 0x2e, 0xb6, 0x8c, 0xd4, 0xc4, 0x94, 0xd4, 0x22, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x6e, 0x23, 0x0d, 0x3d, 0x9c, 0xf6, 0xe9, 0x39, 0x85, 0x38, 0x7b, 0x80, 0xd5,
	0x7a, 0xe6, 0xa5, 0xe5, 0x07, 0x41, 0xf5, 0x29, 0x85, 0x73, 0x09, 0x23, 0x9b, 0xea, 0x96, 0x5f,
	0x54, 0x9e, 0x58, 0x94, 0x42, 0x05, 0x83, 0xa3, 0xb8, 0xc4, 0x60, 0x06, 0xc3, 0x64, 0x8b, 0x53,
	0x8b, 0x4a, 0x52, 0xa9, 0x61, 0xb6, 0x31, 0xc2, 0x6c, 0xb7, 0xd4, 0xd4, 0xa0, 0xc4, 0x92, 0xd4,
	0xd0, 0x82, 0x94, 0x44, 0x90, 0xd9, 0x92, 0x5c, 0x1c, 0x69, 0xa9, 0xa9, 0xf1, 0x45, 0x89, 0x25,
	0xa9, 0x60, 0xd3, 0x59, 0x82, 0xd8, 0xd3, 0x20, 0x2a, 0x9c, 0xc2, 0x4f, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63,
	0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x36, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f,
	0x57, 0xdf, 0x07, 0xe2, 0x14, 0xdd, 0x00, 0x50, 0x98, 0x27, 0xe7, 0xe7, 0xe8, 0xc3, 0x62, 0xa9,
	0xcc, 0x48, 0xbf, 0x02, 0x3d, 0xaa, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xf1, 0x63,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x46, 0x44, 0x83, 0x25, 0x13, 0x02, 0x00, 0x00,
}

func (m *EventBTCRollBack) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBTCRollBack) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBTCRollBack) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBTCRollForward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBTCRollForward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBTCRollForward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBTCHeaderInserted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBTCHeaderInserted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBTCHeaderInserted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBTCFeeRateUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBTCFeeRateUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBTCFeeRateUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeeRate != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.FeeRate))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventBTCRollBack) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventBTCRollForward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventBTCHeaderInserted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventBTCFeeRateUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FeeRate != 0 {
		n += 1 + sovEvent(uint64(m.FeeRate))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventBTCRollBack) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventBTCRollBack: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBTCRollBack: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &BTCHeaderInfo{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventBTCRollForward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventBTCRollForward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBTCRollForward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &BTCHeaderInfo{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventBTCHeaderInserted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventBTCHeaderInserted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBTCHeaderInserted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &BTCHeaderInfo{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventBTCFeeRateUpdated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventBTCFeeRateUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBTCFeeRateUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			m.FeeRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
