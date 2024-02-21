// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/erc20/v1/genesis.proto

package v3types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	"akila/x/erc20/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// V3GenesisState defines the module's genesis state.
type V3GenesisState struct {
	// V3Params are the erc20 module parameters at genesis
	V3Params V3Params `protobuf:"bytes,1,opt,name=V3Params,proto3" json:"V3Params"`
	// token_pairs is a slice of the registered token pairs at genesis
	TokenPairs []types.TokenPair `protobuf:"bytes,2,rep,name=token_pairs,json=tokenPairs,proto3" json:"token_pairs"`
}

func (m *V3GenesisState) Reset()         { *m = V3GenesisState{} }
func (m *V3GenesisState) String() string { return proto.CompactTextString(m) }
func (*V3GenesisState) ProtoMessage()    {}
func (*V3GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f4674601b0d6987, []int{0}
}

func (m *V3GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *V3GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_V3GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *V3GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V3GenesisState.Merge(m, src)
}

func (m *V3GenesisState) XXX_Size() int {
	return m.Size()
}

func (m *V3GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_V3GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_V3GenesisState proto.InternalMessageInfo

func (m *V3GenesisState) GetV3Params() V3Params {
	if m != nil {
		return m.V3Params
	}
	return V3Params{}
}

func (m *V3GenesisState) GetTokenPairs() []types.TokenPair {
	if m != nil {
		return m.TokenPairs
	}
	return nil
}

// V3Params defines the erc20 module V3Params
type V3Params struct {
	// enable_erc20 is the parameter to enable the conversion of Cosmos coins <--> ERC20 tokens.
	EnableErc20 bool `protobuf:"varint,1,opt,name=enable_erc20,json=enableErc20,proto3" json:"enable_erc20,omitempty"`
	// enable_evm_hook is the parameter to enable the EVM hook that converts an ERC20 token to a Cosmos
	// Coin by transferring the Tokens through a MsgEthereumTx to the ModuleAddress Ethereum address.
	EnableEVMHook bool `protobuf:"varint,2,opt,name=enable_evm_hook,json=enableEvmHook,proto3" json:"enable_evm_hook,omitempty"`
}

func (m *V3Params) Reset()         { *m = V3Params{} }
func (m *V3Params) String() string { return proto.CompactTextString(m) }
func (*V3Params) ProtoMessage()    {}
func (*V3Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f4674601b0d6987, []int{1}
}

func (m *V3Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *V3Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_V3Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *V3Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V3Params.Merge(m, src)
}

func (m *V3Params) XXX_Size() int {
	return m.Size()
}

func (m *V3Params) XXX_DiscardUnknown() {
	xxx_messageInfo_V3Params.DiscardUnknown(m)
}

var xxx_messageInfo_V3Params proto.InternalMessageInfo

func (m *V3Params) GetEnableErc20() bool {
	if m != nil {
		return m.EnableErc20
	}
	return false
}

func (m *V3Params) GetEnableEVMHook() bool {
	if m != nil {
		return m.EnableEVMHook
	}
	return false
}

func init() {
	proto.RegisterType((*V3GenesisState)(nil), "akila.erc20.v1.V3GenesisState")
	proto.RegisterType((*V3Params)(nil), "akila.erc20.v1.V3Params")
}

func init() { proto.RegisterFile("akila/erc20/v1/genesis.proto", fileDescriptor_2f4674601b0d6987) }

var fileDescriptor_2f4674601b0d6987 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x2d, 0xcb, 0xcd,
	0x2f, 0xd6, 0x4f, 0x2d, 0x4a, 0x36, 0x32, 0xd0, 0x2f, 0x33, 0xd4, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d,
	0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0xcb, 0xea, 0x81, 0x65, 0xf5,
	0xca, 0x0c, 0xa5, 0xa4, 0xd0, 0x54, 0x43, 0x24, 0xc0, 0x6a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3,
	0xc1, 0x4c, 0x7d, 0x10, 0x0b, 0x22, 0xaa, 0xd4, 0xc6, 0xc8, 0xc5, 0xe3, 0x0e, 0x31, 0x33, 0xb8,
	0x24, 0xb1, 0x24, 0x55, 0xc8, 0x84, 0x8b, 0xad, 0x20, 0xb1, 0x28, 0x31, 0xb7, 0x58, 0x82, 0x51,
	0x81, 0x51, 0x83, 0xdb, 0x48, 0x4c, 0x0f, 0xd5, 0x0e, 0xbd, 0x00, 0xb0, 0xac, 0x13, 0xcb, 0x89,
	0x7b, 0xf2, 0x0c, 0x41, 0x50, 0xb5, 0x42, 0x0e, 0x5c, 0xdc, 0x25, 0xf9, 0xd9, 0xa9, 0x79, 0xf1,
	0x05, 0x89, 0x99, 0x45, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x92, 0xe8, 0x5a, 0x43,
	0x40, 0x4a, 0x02, 0x12, 0x33, 0x8b, 0xa0, 0xba, 0xb9, 0x4a, 0x60, 0x02, 0xc5, 0x4a, 0x69, 0x5c,
	0x6c, 0x10, 0x93, 0x85, 0x14, 0xb9, 0x78, 0x52, 0xf3, 0x12, 0x93, 0x72, 0x52, 0xe3, 0xc1, 0x1a,
	0xc1, 0xee, 0xe0, 0x08, 0xe2, 0x86, 0x88, 0xb9, 0x82, 0x84, 0x84, 0x2c, 0xb9, 0xf8, 0x61, 0x4a,
	0xca, 0x72, 0xe3, 0x33, 0xf2, 0xf3, 0xb3, 0x25, 0x98, 0x40, 0xaa, 0x9c, 0x04, 0x1f, 0xdd, 0x93,
	0xe7, 0x75, 0x85, 0xa8, 0x0c, 0xf3, 0xf5, 0xc8, 0xcf, 0xcf, 0x0e, 0xe2, 0x85, 0x6a, 0x2c, 0xcb,
	0x05, 0x71, 0x9d, 0x9c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x23,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x1a, 0x8e, 0x60, 0xb2, 0xcc,
	0xd0, 0x40, 0xbf, 0x02, 0x1a, 0xa6, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xb0, 0x33,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xc5, 0x52, 0x8e, 0x9d, 0x01, 0x00, 0x00,
}

func (m *V3GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *V3GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *V3GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenPairs) > 0 {
		for iNdEx := len(m.TokenPairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenPairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.V3Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *V3Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *V3Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *V3Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EnableEVMHook {
		i--
		if m.EnableEVMHook {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.EnableErc20 {
		i--
		if m.EnableErc20 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *V3GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.V3Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.TokenPairs) > 0 {
		for _, e := range m.TokenPairs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *V3Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableErc20 {
		n += 2
	}
	if m.EnableEVMHook {
		n += 2
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *V3GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: V3GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: V3GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V3Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.V3Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenPairs = append(m.TokenPairs, types.TokenPair{})
			if err := m.TokenPairs[len(m.TokenPairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func (m *V3Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: V3Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: V3Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableErc20", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.EnableErc20 = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableEVMHook", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.EnableEVMHook = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
