// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Params defines the parameters for the module
type Params struct {
	// string base_denom = 1 [
	//   (gogoproto.moretags)   = "yaml:\"base_denom\"",
	//   (gogoproto.jsontag) = "base_denom"
	// ];
	PriorityNormalizer github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=priority_normalizer,json=priorityNormalizer,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"priority_normalizer" yaml:"priority_normalizer"`
	BaseFeePerGas      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=base_fee_per_gas,json=baseFeePerGas,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_fee_per_gas" yaml:"base_fee_per_gas"`
	MinimumFeePerGas   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=minimum_fee_per_gas,json=minimumFeePerGas,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"minimum_fee_per_gas" yaml:"minimum_fee_per_gas"`
	// ChainConfig chain_config = 5 [(gogoproto.moretags) = "yaml:\"chain_config\"", (gogoproto.nullable) = false];
	//   string chain_id = 6 [
	//   (gogoproto.moretags)   = "yaml:\"chain_id\"",
	//   (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
	//   (gogoproto.nullable)   = false,
	//   (gogoproto.jsontag) = "chain_id"
	// ];
	// repeated string whitelisted_codehashes_bank_send = 7 [
	//   (gogoproto.moretags)   = "yaml:\"whitelisted_codehashes_bank_send\"",
	//   (gogoproto.jsontag) = "whitelisted_codehashes_bank_send"
	// ];
	WhitelistedCwCodeHashesForDelegateCall [][]byte                               `protobuf:"bytes,8,rep,name=whitelisted_cw_code_hashes_for_delegate_call,json=whitelistedCwCodeHashesForDelegateCall,proto3" json:"whitelisted_cw_code_hashes_for_delegate_call" yaml:"whitelisted_cw_code_hashes_for_delegate_call"`
	DeliverTxHookWasmGasLimit              uint64                                 `protobuf:"varint,9,opt,name=deliver_tx_hook_wasm_gas_limit,json=deliverTxHookWasmGasLimit,proto3" json:"deliver_tx_hook_wasm_gas_limit,omitempty"`
	MaxDynamicBaseFeeUpwardAdjustment      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=max_dynamic_base_fee_upward_adjustment,json=maxDynamicBaseFeeUpwardAdjustment,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_dynamic_base_fee_upward_adjustment" yaml:"max_dynamic_base_fee_upward_adjustment"`
	MaxDynamicBaseFeeDownwardAdjustment    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=max_dynamic_base_fee_downward_adjustment,json=maxDynamicBaseFeeDownwardAdjustment,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_dynamic_base_fee_downward_adjustment" yaml:"max_dynamic_base_fee_downward_adjustment"`
	TargetGasUsedPerBlock                  uint64                                 `protobuf:"varint,12,opt,name=target_gas_used_per_block,json=targetGasUsedPerBlock,proto3" json:"target_gas_used_per_block,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9272f3679901ea94, []int{0}
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

func (m *Params) GetWhitelistedCwCodeHashesForDelegateCall() [][]byte {
	if m != nil {
		return m.WhitelistedCwCodeHashesForDelegateCall
	}
	return nil
}

func (m *Params) GetDeliverTxHookWasmGasLimit() uint64 {
	if m != nil {
		return m.DeliverTxHookWasmGasLimit
	}
	return 0
}

func (m *Params) GetTargetGasUsedPerBlock() uint64 {
	if m != nil {
		return m.TargetGasUsedPerBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "seiprotocol.seichain.evm.Params")
}

func init() { proto.RegisterFile("evm/params.proto", fileDescriptor_9272f3679901ea94) }

var fileDescriptor_9272f3679901ea94 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x6b, 0x13, 0x4f,
	0x18, 0xc7, 0x33, 0xbf, 0x96, 0xf2, 0xeb, 0x5a, 0xa1, 0x6c, 0x15, 0xb7, 0x3d, 0xec, 0xd6, 0x15,
	0x4a, 0x0e, 0x26, 0x7b, 0xe8, 0x45, 0x7a, 0x6b, 0x1a, 0x9a, 0x1e, 0x44, 0xca, 0x62, 0x11, 0x04,
	0x19, 0x26, 0x3b, 0x4f, 0x37, 0x63, 0x66, 0x32, 0x61, 0x66, 0xd2, 0x24, 0xbe, 0x00, 0xc1, 0x83,
	0x20, 0xe2, 0xc1, 0xa3, 0x6f, 0x46, 0xe8, 0xb1, 0x47, 0x11, 0x5c, 0xa4, 0xc5, 0x4b, 0x8f, 0x79,
	0x05, 0x92, 0xc9, 0xf6, 0x8f, 0xed, 0x22, 0x89, 0xa7, 0x4c, 0x9e, 0xe7, 0x33, 0xcf, 0x7c, 0xbf,
	0xcf, 0x33, 0x3b, 0xce, 0x32, 0x1c, 0x89, 0xa8, 0x4b, 0x14, 0x11, 0xba, 0xda, 0x55, 0xd2, 0x48,
	0xd7, 0xd3, 0xc0, 0xec, 0x2a, 0x91, 0xbc, 0xaa, 0x81, 0x25, 0x2d, 0xc2, 0x3a, 0x55, 0x38, 0x12,
	0x6b, 0xf7, 0x52, 0x99, 0x4a, 0x9b, 0x8a, 0xc6, 0xab, 0x09, 0x1f, 0x9e, 0x2c, 0x3a, 0x0b, 0xfb,
	0xb6, 0x80, 0xfb, 0x09, 0x39, 0x2b, 0x5d, 0xc5, 0xa4, 0x62, 0x66, 0x88, 0x3b, 0x52, 0x09, 0xc2,
	0xd9, 0x1b, 0x50, 0xde, 0x7f, 0xeb, 0xa8, 0xbc, 0x58, 0x4b, 0x8e, 0xb3, 0xa0, 0xf4, 0x3d, 0x0b,
	0x36, 0x52, 0x66, 0x5a, 0xbd, 0x66, 0x35, 0x91, 0x22, 0x4a, 0xa4, 0x16, 0x52, 0xe7, 0x3f, 0x15,
	0x4d, 0xdb, 0x91, 0x19, 0x76, 0x41, 0x57, 0xeb, 0x90, 0x9c, 0x67, 0x41, 0x51, 0xb1, 0x51, 0x16,
	0xac, 0x0d, 0x89, 0xe0, 0x5b, 0x61, 0x41, 0x32, 0x8c, 0xdd, 0x8b, 0xe8, 0xb3, 0xcb, 0xa0, 0xfb,
	0x16, 0x39, 0xcb, 0x4d, 0xa2, 0x01, 0x1f, 0x02, 0xe0, 0x2e, 0x28, 0x9c, 0x12, 0xed, 0xcd, 0x59,
	0x4d, 0xaf, 0x66, 0xd6, 0x74, 0xab, 0xd2, 0x28, 0x0b, 0x1e, 0x4c, 0x04, 0xdd, 0xcc, 0x84, 0xf1,
	0xdd, 0x71, 0x68, 0x17, 0x60, 0x1f, 0x54, 0x83, 0x68, 0xf7, 0x23, 0x72, 0x56, 0x04, 0xeb, 0x30,
	0xd1, 0x13, 0x7f, 0x68, 0x99, 0xff, 0xd7, 0xfe, 0x14, 0x14, 0xbb, 0xea, 0x4f, 0x41, 0x32, 0x8c,
	0x97, 0xf3, 0xe8, 0x95, 0xa8, 0xaf, 0xc8, 0x79, 0xdc, 0x6f, 0x31, 0x03, 0x9c, 0x69, 0x03, 0x14,
	0x27, 0x7d, 0x9c, 0x48, 0x0a, 0xb8, 0x45, 0x74, 0x0b, 0x34, 0x3e, 0x94, 0x0a, 0x53, 0xe0, 0x90,
	0x12, 0x03, 0x38, 0x21, 0x9c, 0x7b, 0xff, 0xaf, 0xcf, 0x95, 0x97, 0x6a, 0xe9, 0x79, 0x16, 0xcc,
	0xb4, 0x6f, 0x94, 0x05, 0x9b, 0x13, 0x61, 0xb3, 0xec, 0x0a, 0xe3, 0x8d, 0x6b, 0xf8, 0x4e, 0x7f,
	0x47, 0x52, 0xd8, 0xb3, 0xec, 0xae, 0x54, 0xf5, 0x9c, 0xdc, 0x21, 0x9c, 0xbb, 0xdb, 0x8e, 0x4f,
	0x81, 0xb3, 0x23, 0x50, 0xd8, 0x0c, 0x70, 0x4b, 0xca, 0x36, 0xee, 0x13, 0x2d, 0xc6, 0xb6, 0x31,
	0x67, 0x82, 0x19, 0x6f, 0x71, 0x1d, 0x95, 0xe7, 0xe3, 0xd5, 0x9c, 0x7a, 0x3e, 0xd8, 0x93, 0xb2,
	0xfd, 0x82, 0x68, 0xd1, 0x20, 0xfa, 0xe9, 0x18, 0x70, 0x7f, 0x20, 0x67, 0x43, 0x90, 0x01, 0xa6,
	0xc3, 0x0e, 0x11, 0x2c, 0xc1, 0x97, 0x03, 0xed, 0x75, 0xfb, 0x44, 0x51, 0x4c, 0xe8, 0xeb, 0x9e,
	0x36, 0x02, 0x3a, 0xc6, 0x73, 0xec, 0xc8, 0xde, 0xa1, 0x99, 0x67, 0x36, 0xe5, 0x01, 0xa3, 0x2c,
	0xa8, 0xe4, 0x63, 0x9c, 0x8a, 0x0f, 0xe3, 0x87, 0x82, 0x0c, 0xea, 0x13, 0xae, 0x36, 0xb9, 0x75,
	0x07, 0x16, 0xda, 0xbe, 0x64, 0xdc, 0x5f, 0xc8, 0x29, 0x17, 0x96, 0xa3, 0xb2, 0xdf, 0xb9, 0xe9,
	0xf0, 0x8e, 0x75, 0xf8, 0x7e, 0x76, 0x87, 0x53, 0x1f, 0x31, 0xca, 0x82, 0xe8, 0x2f, 0x1e, 0x0b,
	0x76, 0x84, 0xf1, 0xa3, 0x5b, 0x2e, 0xeb, 0x39, 0x76, 0xcd, 0xe7, 0x13, 0x67, 0xd5, 0x10, 0x95,
	0x82, 0xb1, 0xc3, 0xef, 0x69, 0xa0, 0xf6, 0x03, 0x68, 0x72, 0x99, 0xb4, 0xbd, 0x25, 0x7b, 0x0b,
	0xee, 0x4f, 0x80, 0x06, 0xd1, 0x07, 0x1a, 0xe8, 0x3e, 0xa8, 0xda, 0x38, 0xb9, 0x35, 0xff, 0xf9,
	0x4b, 0x50, 0xaa, 0x35, 0x8e, 0x4f, 0x7d, 0x74, 0x72, 0xea, 0xa3, 0x9f, 0xa7, 0x3e, 0xfa, 0x70,
	0xe6, 0x97, 0x4e, 0xce, 0xfc, 0xd2, 0xb7, 0x33, 0xbf, 0xf4, 0xb2, 0x72, 0xad, 0x0b, 0x1a, 0x58,
	0xe5, 0xe2, 0xa1, 0xb4, 0x7f, 0xec, 0x4b, 0x19, 0x0d, 0xa2, 0xf1, 0x93, 0x6a, 0x1b, 0xd2, 0x5c,
	0xb0, 0xf9, 0xcd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x48, 0xda, 0x43, 0x66, 0x05, 0x00,
	0x00,
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
	if m.TargetGasUsedPerBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TargetGasUsedPerBlock))
		i--
		dAtA[i] = 0x60
	}
	{
		size := m.MaxDynamicBaseFeeDownwardAdjustment.Size()
		i -= size
		if _, err := m.MaxDynamicBaseFeeDownwardAdjustment.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.MaxDynamicBaseFeeUpwardAdjustment.Size()
		i -= size
		if _, err := m.MaxDynamicBaseFeeUpwardAdjustment.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.DeliverTxHookWasmGasLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DeliverTxHookWasmGasLimit))
		i--
		dAtA[i] = 0x48
	}
	if len(m.WhitelistedCwCodeHashesForDelegateCall) > 0 {
		for iNdEx := len(m.WhitelistedCwCodeHashesForDelegateCall) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WhitelistedCwCodeHashesForDelegateCall[iNdEx])
			copy(dAtA[i:], m.WhitelistedCwCodeHashesForDelegateCall[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.WhitelistedCwCodeHashesForDelegateCall[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	{
		size := m.MinimumFeePerGas.Size()
		i -= size
		if _, err := m.MinimumFeePerGas.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.BaseFeePerGas.Size()
		i -= size
		if _, err := m.BaseFeePerGas.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.PriorityNormalizer.Size()
		i -= size
		if _, err := m.PriorityNormalizer.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
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
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PriorityNormalizer.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.BaseFeePerGas.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinimumFeePerGas.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.WhitelistedCwCodeHashesForDelegateCall) > 0 {
		for _, b := range m.WhitelistedCwCodeHashesForDelegateCall {
			l = len(b)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.DeliverTxHookWasmGasLimit != 0 {
		n += 1 + sovParams(uint64(m.DeliverTxHookWasmGasLimit))
	}
	l = m.MaxDynamicBaseFeeUpwardAdjustment.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxDynamicBaseFeeDownwardAdjustment.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.TargetGasUsedPerBlock != 0 {
		n += 1 + sovParams(uint64(m.TargetGasUsedPerBlock))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriorityNormalizer", wireType)
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
			if err := m.PriorityNormalizer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFeePerGas", wireType)
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
			if err := m.BaseFeePerGas.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumFeePerGas", wireType)
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
			if err := m.MinimumFeePerGas.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistedCwCodeHashesForDelegateCall", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhitelistedCwCodeHashesForDelegateCall = append(m.WhitelistedCwCodeHashesForDelegateCall, make([]byte, postIndex-iNdEx))
			copy(m.WhitelistedCwCodeHashesForDelegateCall[len(m.WhitelistedCwCodeHashesForDelegateCall)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeliverTxHookWasmGasLimit", wireType)
			}
			m.DeliverTxHookWasmGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeliverTxHookWasmGasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDynamicBaseFeeUpwardAdjustment", wireType)
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
			if err := m.MaxDynamicBaseFeeUpwardAdjustment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDynamicBaseFeeDownwardAdjustment", wireType)
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
			if err := m.MaxDynamicBaseFeeDownwardAdjustment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetGasUsedPerBlock", wireType)
			}
			m.TargetGasUsedPerBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetGasUsedPerBlock |= uint64(b&0x7F) << shift
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
