// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/vesting/v1beta1/vesting.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// LazyGradedVestingAccount implements the LazyGradedVestingAccount interface. It vests all
// coins according to a predefined schedule.
type LazyGradedVestingAccount struct {
	*types.BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
	VestingSchedules          VestingSchedules `protobuf:"bytes,2,rep,name=vesting_schedules,json=vestingSchedules,proto3,castrepeated=VestingSchedules" json:"vesting_schedules" yaml:"vesting_schedules"`
}

func (m *LazyGradedVestingAccount) Reset()      { *m = LazyGradedVestingAccount{} }
func (*LazyGradedVestingAccount) ProtoMessage() {}
func (*LazyGradedVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a9bc06e563192a, []int{0}
}
func (m *LazyGradedVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LazyGradedVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LazyGradedVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LazyGradedVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LazyGradedVestingAccount.Merge(m, src)
}
func (m *LazyGradedVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *LazyGradedVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_LazyGradedVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_LazyGradedVestingAccount proto.InternalMessageInfo

// Schedule - represent single schedule data for a vesting schedule
type Schedule struct {
	StartTime int64                                  `protobuf:"varint,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" yaml:"start_time"`
	EndTime   int64                                  `protobuf:"varint,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" yaml:"end_time"`
	Ratio     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=ratio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"ratio" yaml:"ratio"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a9bc06e563192a, []int{1}
}
func (m *Schedule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Schedule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule.Merge(m, src)
}
func (m *Schedule) XXX_Size() int {
	return m.Size()
}
func (m *Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule proto.InternalMessageInfo

// VestingSchedule defines vesting schedule for a denom
type VestingSchedule struct {
	Denom     string    `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty" yaml:"start_time"`
	Schedules Schedules `protobuf:"bytes,2,rep,name=schedules,proto3,castrepeated=Schedules" json:"schedules" yaml:"schedules"`
}

func (m *VestingSchedule) Reset()         { *m = VestingSchedule{} }
func (m *VestingSchedule) String() string { return proto.CompactTextString(m) }
func (*VestingSchedule) ProtoMessage()    {}
func (*VestingSchedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4a9bc06e563192a, []int{2}
}
func (m *VestingSchedule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingSchedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingSchedule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingSchedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingSchedule.Merge(m, src)
}
func (m *VestingSchedule) XXX_Size() int {
	return m.Size()
}
func (m *VestingSchedule) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingSchedule.DiscardUnknown(m)
}

var xxx_messageInfo_VestingSchedule proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LazyGradedVestingAccount)(nil), "terra.vesting.v1beta1.LazyGradedVestingAccount")
	proto.RegisterType((*Schedule)(nil), "terra.vesting.v1beta1.Schedule")
	proto.RegisterType((*VestingSchedule)(nil), "terra.vesting.v1beta1.VestingSchedule")
}

func init() {
	proto.RegisterFile("terra/vesting/v1beta1/vesting.proto", fileDescriptor_c4a9bc06e563192a)
}

var fileDescriptor_c4a9bc06e563192a = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xe3, 0x76, 0x83, 0xc6, 0x43, 0x5a, 0x17, 0x36, 0x29, 0xda, 0x21, 0xae, 0x02, 0x4c,
	0x15, 0x68, 0x0e, 0x2b, 0x3b, 0xf5, 0x80, 0x44, 0x84, 0x84, 0x84, 0x38, 0x85, 0x89, 0x03, 0x97,
	0xca, 0x89, 0xad, 0x2e, 0xa2, 0x89, 0xa7, 0xd8, 0xad, 0x28, 0x4f, 0x00, 0x37, 0x0e, 0x1c, 0x38,
	0xee, 0xbc, 0x27, 0xd9, 0x81, 0x43, 0x8f, 0x88, 0x43, 0x40, 0xed, 0x1b, 0xe4, 0x09, 0x50, 0xed,
	0x64, 0x1d, 0x19, 0xdd, 0xa9, 0xf5, 0xf7, 0xfd, 0xbf, 0x9f, 0xfd, 0xf7, 0xdf, 0x81, 0x0f, 0x24,
	0xcb, 0x32, 0xe2, 0x4d, 0x98, 0x90, 0x71, 0x3a, 0xf4, 0x26, 0x47, 0x21, 0x93, 0xe4, 0xa8, 0x5a,
	0xe3, 0xb3, 0x8c, 0x4b, 0x6e, 0xed, 0x29, 0x11, 0xae, 0x8a, 0xa5, 0x68, 0x7f, 0x77, 0xc8, 0x87,
	0x5c, 0x29, 0xbc, 0xe5, 0x3f, 0x2d, 0xde, 0x7f, 0x18, 0x71, 0x91, 0x70, 0x71, 0x3b, 0xd2, 0xfd,
	0xd6, 0x80, 0xf6, 0x1b, 0xf2, 0x69, 0xfa, 0x2a, 0x23, 0x94, 0xd1, 0x77, 0xba, 0xf7, 0x22, 0x8a,
	0xf8, 0x38, 0x95, 0x56, 0x08, 0x77, 0x43, 0x22, 0xd8, 0xa0, 0x1c, 0x19, 0x10, 0x5d, 0xb7, 0x41,
	0x07, 0x74, 0xb7, 0x7a, 0x8f, 0xb1, 0xde, 0xa1, 0x7e, 0x1e, 0xec, 0x13, 0xc1, 0xfe, 0x25, 0xf9,
	0x1b, 0xb3, 0x1c, 0x81, 0xc0, 0x0a, 0x6f, 0x74, 0xac, 0x2f, 0x00, 0xee, 0x54, 0x7c, 0x11, 0x9d,
	0x32, 0x3a, 0x1e, 0x31, 0x61, 0x37, 0x3a, 0xcd, 0xee, 0x56, 0xef, 0x00, 0xff, 0xd7, 0x30, 0x2e,
	0x11, 0x6f, 0x4b, 0xb9, 0x7f, 0x7c, 0x99, 0x23, 0xa3, 0xc8, 0x91, 0x3d, 0x25, 0xc9, 0xa8, 0xef,
	0xde, 0xc0, 0xb9, 0x17, 0xbf, 0x51, 0xbb, 0x36, 0x24, 0x82, 0xf6, 0xa4, 0x56, 0xe9, 0xb7, 0x3e,
	0x9f, 0x23, 0xe3, 0xfb, 0x39, 0x32, 0xdc, 0x1f, 0x00, 0xb6, 0xaa, 0xba, 0x75, 0x0c, 0xa1, 0x90,
	0x24, 0x93, 0x03, 0x19, 0x27, 0x4c, 0x99, 0x6f, 0xfa, 0x7b, 0x45, 0x8e, 0x76, 0xf4, 0x76, 0xab,
	0x9e, 0x1b, 0x98, 0x6a, 0x71, 0x12, 0x27, 0xcc, 0xc2, 0xb0, 0xc5, 0x52, 0xaa, 0x67, 0x1a, 0x6a,
	0xe6, 0x7e, 0x91, 0xa3, 0x6d, 0x3d, 0x53, 0x75, 0xdc, 0xe0, 0x2e, 0x4b, 0xa9, 0xd2, 0x9f, 0xc0,
	0xcd, 0x8c, 0xc8, 0x98, 0xdb, 0xcd, 0x0e, 0xe8, 0x9a, 0xfe, 0xf3, 0xa5, 0xa7, 0x5f, 0x39, 0x3a,
	0x18, 0xc6, 0xf2, 0x74, 0x1c, 0xe2, 0x88, 0x27, 0x5e, 0x99, 0xa8, 0xfe, 0x39, 0x14, 0xf4, 0x83,
	0x27, 0xa7, 0x67, 0x4c, 0xe0, 0x97, 0x2c, 0x2a, 0x72, 0x74, 0x4f, 0xa3, 0x15, 0xc4, 0x0d, 0x34,
	0xac, 0xbf, 0xb1, 0xb4, 0xe4, 0x5e, 0x00, 0xb8, 0x5d, 0xf3, 0x6f, 0x3d, 0x81, 0x9b, 0x94, 0xa5,
	0x3c, 0x51, 0x86, 0xcc, 0x75, 0x86, 0xb4, 0xc6, 0xa2, 0xd0, 0xac, 0x87, 0x83, 0xd6, 0x84, 0x73,
	0x95, 0xca, 0xa3, 0x32, 0x95, 0x76, 0x49, 0xbd, 0x9e, 0x86, 0xb9, 0x8a, 0x61, 0x05, 0xd6, 0x87,
	0xf5, 0x5f, 0x5f, 0xce, 0x1d, 0x30, 0x9b, 0x3b, 0xe0, 0xcf, 0xdc, 0x01, 0x5f, 0x17, 0x8e, 0x31,
	0x5b, 0x38, 0xc6, 0xcf, 0x85, 0x63, 0xbc, 0x7f, 0x7a, 0xfd, 0x2e, 0x46, 0x44, 0x88, 0x38, 0x3a,
	0xd4, 0xdf, 0x4d, 0xc4, 0x33, 0xe6, 0x4d, 0x7a, 0xde, 0xc7, 0xab, 0xe7, 0xae, 0x6e, 0x26, 0xbc,
	0xa3, 0x5e, 0xf9, 0xb3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x5c, 0x52, 0x15, 0x5f, 0x03,
	0x00, 0x00,
}

func (m *LazyGradedVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LazyGradedVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LazyGradedVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingSchedules) > 0 {
		for iNdEx := len(m.VestingSchedules) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingSchedules[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Schedule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schedule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schedule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Ratio.Size()
		i -= size
		if _, err := m.Ratio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVesting(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.EndTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x10
	}
	if m.StartTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *VestingSchedule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingSchedule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingSchedule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Schedules) > 0 {
		for iNdEx := len(m.Schedules) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Schedules[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVesting(dAtA []byte, offset int, v uint64) int {
	offset -= sovVesting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LazyGradedVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	if len(m.VestingSchedules) > 0 {
		for _, e := range m.VestingSchedules {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	return n
}

func (m *Schedule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartTime != 0 {
		n += 1 + sovVesting(uint64(m.StartTime))
	}
	if m.EndTime != 0 {
		n += 1 + sovVesting(uint64(m.EndTime))
	}
	l = m.Ratio.Size()
	n += 1 + l + sovVesting(uint64(l))
	return n
}

func (m *VestingSchedule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	if len(m.Schedules) > 0 {
		for _, e := range m.Schedules {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	return n
}

func sovVesting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVesting(x uint64) (n int) {
	return sovVesting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LazyGradedVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: LazyGradedVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LazyGradedVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &types.BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingSchedules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingSchedules = append(m.VestingSchedules, VestingSchedule{})
			if err := m.VestingSchedules[len(m.VestingSchedules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *Schedule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: Schedule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Schedule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ratio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ratio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *VestingSchedule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: VestingSchedule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingSchedule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schedules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schedules = append(m.Schedules, Schedule{})
			if err := m.Schedules[len(m.Schedules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func skipVesting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
				return 0, ErrInvalidLengthVesting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVesting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVesting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVesting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVesting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVesting = fmt.Errorf("proto: unexpected end of group")
)
