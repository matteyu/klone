// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: epoch/epoch.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Epoch struct {
	GenesisTime           time.Time     `protobuf:"bytes,1,opt,name=genesis_time,json=genesisTime,proto3,stdtime" json:"genesis_time" yaml:"genesis_time"`
	EpochDuration         time.Duration `protobuf:"bytes,2,opt,name=epoch_duration,json=epochDuration,proto3,stdduration" json:"duration,omitempty" yaml:"epoch_duration"`
	CurrentEpoch          uint64        `protobuf:"varint,3,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch" yaml:"current_epoch"`
	CurrentEpochStartTime time.Time     `protobuf:"bytes,4,opt,name=current_epoch_start_time,json=currentEpochStartTime,proto3,stdtime" json:"current_epoch_start_time" yaml:"current_epoch_start_time"`
	CurrentEpochHeight    int64         `protobuf:"varint,5,opt,name=current_epoch_height,json=currentEpochHeight,proto3" json:"current_epoch_height" yaml:"current_epoch_height"`
}

func (m *Epoch) Reset()         { *m = Epoch{} }
func (m *Epoch) String() string { return proto.CompactTextString(m) }
func (*Epoch) ProtoMessage()    {}
func (*Epoch) Descriptor() ([]byte, []int) {
	return fileDescriptor_36a9d1673530db42, []int{0}
}
func (m *Epoch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Epoch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Epoch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Epoch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Epoch.Merge(m, src)
}
func (m *Epoch) XXX_Size() int {
	return m.Size()
}
func (m *Epoch) XXX_DiscardUnknown() {
	xxx_messageInfo_Epoch.DiscardUnknown(m)
}

var xxx_messageInfo_Epoch proto.InternalMessageInfo

func (m *Epoch) GetGenesisTime() time.Time {
	if m != nil {
		return m.GenesisTime
	}
	return time.Time{}
}

func (m *Epoch) GetEpochDuration() time.Duration {
	if m != nil {
		return m.EpochDuration
	}
	return 0
}

func (m *Epoch) GetCurrentEpoch() uint64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

func (m *Epoch) GetCurrentEpochStartTime() time.Time {
	if m != nil {
		return m.CurrentEpochStartTime
	}
	return time.Time{}
}

func (m *Epoch) GetCurrentEpochHeight() int64 {
	if m != nil {
		return m.CurrentEpochHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Epoch)(nil), "seiprotocol.seichain.epoch.Epoch")
}

func init() { proto.RegisterFile("epoch/epoch.proto", fileDescriptor_36a9d1673530db42) }

var fileDescriptor_36a9d1673530db42 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x31, 0x8b, 0xdb, 0x30,
	0x1c, 0xc5, 0xad, 0x26, 0xe9, 0xe0, 0x24, 0x85, 0xba, 0x09, 0xb8, 0x2e, 0x58, 0xc1, 0x53, 0x4a,
	0x5b, 0x0b, 0xda, 0x21, 0xd0, 0xd1, 0xb4, 0xd0, 0x2e, 0x1d, 0xd2, 0x4e, 0x1d, 0x6a, 0x1c, 0x57,
	0xb5, 0x05, 0xb1, 0x65, 0x2c, 0x19, 0xea, 0xad, 0x1f, 0x21, 0x63, 0xc7, 0xfb, 0x38, 0x19, 0x33,
	0xde, 0xa4, 0x3b, 0x92, 0x2d, 0xa3, 0x3f, 0xc1, 0x61, 0xc9, 0x06, 0xfb, 0x2e, 0x70, 0x4b, 0x90,
	0xfe, 0xef, 0xfd, 0xdf, 0x2f, 0x7e, 0x48, 0x7f, 0x8e, 0x33, 0x1a, 0xc6, 0x48, 0xfe, 0xba, 0x59,
	0x4e, 0x39, 0x35, 0x2c, 0x86, 0x89, 0x3c, 0x85, 0x74, 0xeb, 0x32, 0x4c, 0xc2, 0x38, 0x20, 0xa9,
	0x2b, 0x1d, 0xd6, 0x2c, 0xa2, 0x11, 0x95, 0x22, 0xaa, 0x4f, 0x6a, 0xc3, 0x82, 0x11, 0xa5, 0xd1,
	0x16, 0x23, 0x79, 0xdb, 0x14, 0x7f, 0x10, 0x27, 0x09, 0x66, 0x3c, 0x48, 0xb2, 0xc6, 0x60, 0xdf,
	0x37, 0xfc, 0x2e, 0xf2, 0x80, 0x13, 0x9a, 0x2a, 0xdd, 0xb9, 0x1a, 0xea, 0xa3, 0xcf, 0x35, 0xc0,
	0xf8, 0xa5, 0x4f, 0x22, 0x9c, 0x62, 0x46, 0x98, 0x5f, 0x87, 0x98, 0x60, 0x01, 0x96, 0xe3, 0xf7,
	0x96, 0xab, 0x02, 0xdc, 0x36, 0xc0, 0xfd, 0xd1, 0x12, 0x3c, 0xb8, 0x17, 0x50, 0xab, 0x04, 0x7c,
	0x51, 0x06, 0xc9, 0xf6, 0xa3, 0xd3, 0xdd, 0x76, 0x76, 0x37, 0x10, 0xac, 0xc7, 0xcd, 0xa8, 0x5e,
	0x31, 0x4a, 0xfd, 0x99, 0xfc, 0x12, 0xbf, 0xfd, 0x07, 0xe6, 0x13, 0x49, 0x78, 0xf9, 0x80, 0xf0,
	0xa9, 0x31, 0x78, 0xab, 0x1a, 0x70, 0x16, 0xd0, 0x68, 0x57, 0xde, 0xd2, 0x84, 0x70, 0x9c, 0x64,
	0xbc, 0xac, 0x04, 0x9c, 0x2b, 0x6c, 0x3f, 0xd4, 0xf9, 0x5f, 0x83, 0xa7, 0x72, 0xd8, 0xe6, 0x18,
	0xdf, 0xf4, 0x69, 0x58, 0xe4, 0x39, 0x4e, 0xb9, 0x2f, 0x05, 0x73, 0xb0, 0x00, 0xcb, 0xa1, 0xf7,
	0xfa, 0x2c, 0x60, 0x5f, 0xa8, 0x04, 0x9c, 0xa9, 0xd4, 0xde, 0xd8, 0x59, 0x4f, 0x9a, 0xbb, 0xaa,
	0xea, 0x1f, 0xd0, 0xcd, 0x9e, 0xc1, 0x67, 0x3c, 0xc8, 0xb9, 0xea, 0x6d, 0xf8, 0x68, 0x6f, 0x6f,
	0x9a, 0xde, 0xe0, 0x05, 0x54, 0x27, 0x49, 0x75, 0x38, 0xef, 0x92, 0xbf, 0xd7, 0xa2, 0x6c, 0x93,
	0xe8, 0xb3, 0xfe, 0x5e, 0x8c, 0x49, 0x14, 0x73, 0x73, 0xb4, 0x00, 0xcb, 0x81, 0xb7, 0x3a, 0x0b,
	0x78, 0x51, 0xaf, 0x04, 0x7c, 0x75, 0x89, 0xaa, 0x54, 0x67, 0x6d, 0x74, 0x69, 0x5f, 0xe4, 0xd0,
	0xfb, 0xba, 0x3f, 0xda, 0xe0, 0x70, 0xb4, 0xc1, 0xed, 0xd1, 0x06, 0xbb, 0x93, 0xad, 0x1d, 0x4e,
	0xb6, 0x76, 0x7d, 0xb2, 0xb5, 0x9f, 0x28, 0x22, 0x3c, 0x2e, 0x36, 0x6e, 0x48, 0x13, 0xc4, 0x30,
	0x79, 0xd7, 0xbe, 0x5d, 0x79, 0x91, 0x8f, 0x17, 0xfd, 0x55, 0x0f, 0x1c, 0xf1, 0x32, 0xc3, 0x6c,
	0xf3, 0x54, 0x3a, 0x3e, 0xdc, 0x05, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xcf, 0x24, 0x22, 0xfc, 0x02,
	0x00, 0x00,
}

func (m *Epoch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Epoch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Epoch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentEpochHeight != 0 {
		i = encodeVarintEpoch(dAtA, i, uint64(m.CurrentEpochHeight))
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CurrentEpochStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CurrentEpochStartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEpoch(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.CurrentEpoch != 0 {
		i = encodeVarintEpoch(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x18
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.EpochDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.EpochDuration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintEpoch(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.GenesisTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.GenesisTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintEpoch(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEpoch(dAtA []byte, offset int, v uint64) int {
	offset -= sovEpoch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Epoch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.GenesisTime)
	n += 1 + l + sovEpoch(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.EpochDuration)
	n += 1 + l + sovEpoch(uint64(l))
	if m.CurrentEpoch != 0 {
		n += 1 + sovEpoch(uint64(m.CurrentEpoch))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CurrentEpochStartTime)
	n += 1 + l + sovEpoch(uint64(l))
	if m.CurrentEpochHeight != 0 {
		n += 1 + sovEpoch(uint64(m.CurrentEpochHeight))
	}
	return n
}

func sovEpoch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEpoch(x uint64) (n int) {
	return sovEpoch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Epoch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpoch
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
			return fmt.Errorf("proto: Epoch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Epoch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.GenesisTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.EpochDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CurrentEpochStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochHeight", wireType)
			}
			m.CurrentEpochHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpochHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpoch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpoch
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
func skipEpoch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpoch
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
					return 0, ErrIntOverflowEpoch
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
					return 0, ErrIntOverflowEpoch
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
				return 0, ErrInvalidLengthEpoch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEpoch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEpoch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEpoch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpoch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEpoch = fmt.Errorf("proto: unexpected end of group")
)