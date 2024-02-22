// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akila/revenue/v1/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	"akila/x/revenue/v1/types"
	cosmossdk_io_math "cosmossdk.io/math"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// V2GenesisState defines the module's genesis state.
type V2GenesisState struct {
	// V2Params are the revenue module parameters
	V2Params V2Params `protobuf:"bytes,1,opt,name=V2Params,proto3" json:"V2Params"`
	// revenues is a slice of active registered contracts for fee distribution
	Revenues []types.Revenue `protobuf:"bytes,2,rep,name=revenues,proto3" json:"revenues"`
}

func (m *V2GenesisState) Reset()         { *m = V2GenesisState{} }
func (m *V2GenesisState) String() string { return proto.CompactTextString(m) }
func (*V2GenesisState) ProtoMessage()    {}
func (*V2GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_649d64d9c3438055, []int{0}
}

func (m *V2GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *V2GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_V2GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *V2GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V2GenesisState.Merge(m, src)
}

func (m *V2GenesisState) XXX_Size() int {
	return m.Size()
}

func (m *V2GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_V2GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_V2GenesisState proto.InternalMessageInfo

func (m *V2GenesisState) GetV2Params() V2Params {
	if m != nil {
		return m.V2Params
	}
	return V2Params{}
}

func (m *V2GenesisState) GetRevenues() []types.Revenue {
	if m != nil {
		return m.Revenues
	}
	return nil
}

// V2Params defines the revenue module V2Params
type V2Params struct {
	// enable_revenue defines a parameter to enable the revenue module
	EnableRevenue bool `protobuf:"varint,1,opt,name=enable_revenue,json=enableRevenue,proto3" json:"enable_revenue,omitempty"`
	// developer_shares defines the proportion of the transaction fees to be
	// distributed to the registered contract owner
	DeveloperShares cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=developer_shares,json=developerShares,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"developer_shares"`
	// addr_derivation_cost_create defines the cost of address derivation for
	// verifying the contract deployer at fee registration
	AddrDerivationCostCreate uint64 `protobuf:"varint,3,opt,name=addr_derivation_cost_create,json=addrDerivationCostCreate,proto3" json:"addr_derivation_cost_create,omitempty"`
}

func (m *V2Params) Reset()         { *m = V2Params{} }
func (m *V2Params) String() string { return proto.CompactTextString(m) }
func (*V2Params) ProtoMessage()    {}
func (*V2Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_649d64d9c3438055, []int{1}
}

func (m *V2Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *V2Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_V2Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *V2Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V2Params.Merge(m, src)
}

func (m *V2Params) XXX_Size() int {
	return m.Size()
}

func (m *V2Params) XXX_DiscardUnknown() {
	xxx_messageInfo_V2Params.DiscardUnknown(m)
}

var xxx_messageInfo_V2Params proto.InternalMessageInfo

func (m *V2Params) GetEnableRevenue() bool {
	if m != nil {
		return m.EnableRevenue
	}
	return false
}

func (m *V2Params) GetAddrDerivationCostCreate() uint64 {
	if m != nil {
		return m.AddrDerivationCostCreate
	}
	return 0
}

func init() {
	proto.RegisterType((*V2GenesisState)(nil), "akila.revenue.v1.V2GenesisState")
	proto.RegisterType((*V2Params)(nil), "akila.revenue.v1.V2Params")
}

func init() { proto.RegisterFile("akila/revenue/v1/genesis.proto", fileDescriptor_649d64d9c3438055) }

var fileDescriptor_649d64d9c3438055 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xb3, 0x6f, 0x4b, 0xe9, 0xbb, 0xf5, 0x4f, 0x09, 0x1e, 0x62, 0x85, 0xb4, 0x14, 0x94,
	0x20, 0xb8, 0xb1, 0x15, 0xbc, 0x88, 0x97, 0xb6, 0xe0, 0x55, 0xd2, 0x93, 0x5e, 0xc2, 0x36, 0x19,
	0xd2, 0x60, 0x9b, 0x0d, 0xbb, 0xdb, 0xa0, 0x67, 0xbf, 0x80, 0x1f, 0xab, 0xde, 0x7a, 0x14, 0x0f,
	0x45, 0xda, 0x2f, 0x22, 0xd9, 0x8d, 0x41, 0xec, 0x25, 0x19, 0xe6, 0x79, 0x9e, 0xdf, 0x0c, 0x3b,
	0xd8, 0x86, 0x6c, 0xce, 0x84, 0xcb, 0x21, 0x83, 0x64, 0x01, 0x6e, 0xd6, 0x73, 0x23, 0x48, 0x40,
	0xc4, 0x82, 0xa4, 0x9c, 0x49, 0x66, 0x36, 0x95, 0x4e, 0x0a, 0x9d, 0x64, 0xbd, 0xd6, 0x6e, 0xe2,
	0x47, 0x54, 0x89, 0xd6, 0x51, 0xc4, 0x22, 0xa6, 0x4a, 0x37, 0xaf, 0x74, 0xb7, 0xfb, 0x8a, 0xf0,
	0xde, 0x9d, 0x26, 0x8f, 0x25, 0x95, 0x60, 0x5e, 0xe3, 0x5a, 0x4a, 0x39, 0x9d, 0x0b, 0x0b, 0x75,
	0x90, 0xd3, 0xe8, 0x5b, 0xe4, 0xef, 0x24, 0x72, 0xaf, 0xf4, 0x41, 0x75, 0xb9, 0x6e, 0x1b, 0x5e,
	0xe1, 0x36, 0x6f, 0x70, 0xbd, 0xb0, 0x08, 0xeb, 0x5f, 0xa7, 0xe2, 0x34, 0xfa, 0xc7, 0xbb, 0x49,
	0x4f, 0x97, 0x45, 0xb4, 0x0c, 0x74, 0xdf, 0x11, 0xae, 0x69, 0xaa, 0x79, 0x8a, 0x0f, 0x20, 0xa1,
	0x93, 0x19, 0xf8, 0x85, 0xaa, 0xf6, 0xa8, 0x7b, 0xfb, 0xba, 0x5b, 0x10, 0xcc, 0x07, 0xdc, 0x0c,
	0x21, 0x83, 0x19, 0x4b, 0x81, 0xfb, 0x62, 0x4a, 0xb9, 0x1a, 0x8b, 0x9c, 0xff, 0x03, 0x92, 0xb3,
	0x3f, 0xd7, 0xed, 0xb3, 0x28, 0x96, 0xd3, 0xc5, 0x84, 0x04, 0x6c, 0xee, 0x06, 0x4c, 0xe4, 0x6f,
	0xa3, 0x7f, 0x17, 0x22, 0x7c, 0x72, 0xe5, 0x4b, 0x0a, 0x82, 0x8c, 0x20, 0xf0, 0x0e, 0x4b, 0xce,
	0x58, 0x61, 0xcc, 0x5b, 0x7c, 0x42, 0xc3, 0x90, 0xfb, 0x21, 0xf0, 0x38, 0xa3, 0x32, 0x66, 0x89,
	0x1f, 0x30, 0x21, 0xfd, 0x80, 0x03, 0x95, 0x60, 0x55, 0x3a, 0xc8, 0xa9, 0x7a, 0x56, 0x6e, 0x19,
	0x95, 0x8e, 0x21, 0x13, 0x72, 0xa8, 0xf4, 0xc1, 0x68, 0xb9, 0xb1, 0xd1, 0x6a, 0x63, 0xa3, 0xaf,
	0x8d, 0x8d, 0xde, 0xb6, 0xb6, 0xb1, 0xda, 0xda, 0xc6, 0xc7, 0xd6, 0x36, 0x1e, 0xcf, 0x7f, 0x6d,
	0xa4, 0x8f, 0xa5, 0xbf, 0x59, 0xef, 0xd2, 0x7d, 0x2e, 0x0f, 0xa7, 0x36, 0x9b, 0xd4, 0xd4, 0x79,
	0xae, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x16, 0xe7, 0x21, 0x0b, 0x08, 0x02, 0x00, 0x00,
}

func (m *V2GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *V2GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *V2GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Revenues) > 0 {
		for iNdEx := len(m.Revenues) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Revenues[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
		size, err := m.V2Params.MarshalToSizedBuffer(dAtA[:i])
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

func (m *V2Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *V2Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *V2Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AddrDerivationCostCreate != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AddrDerivationCostCreate))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.DeveloperShares.Size()
		i -= size
		if _, err := m.DeveloperShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.EnableRevenue {
		i--
		if m.EnableRevenue {
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

func (m *V2GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.V2Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Revenues) > 0 {
		for _, e := range m.Revenues {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *V2Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableRevenue {
		n += 2
	}
	l = m.DeveloperShares.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.AddrDerivationCostCreate != 0 {
		n += 1 + sovGenesis(uint64(m.AddrDerivationCostCreate))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *V2GenesisState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: V2GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: V2GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2Params", wireType)
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
			if err := m.V2Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revenues", wireType)
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
			m.Revenues = append(m.Revenues, types.Revenue{})
			if err := m.Revenues[len(m.Revenues)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

func (m *V2Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: V2Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: V2Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableRevenue", wireType)
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
			m.EnableRevenue = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeveloperShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeveloperShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddrDerivationCostCreate", wireType)
			}
			m.AddrDerivationCostCreate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AddrDerivationCostCreate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
