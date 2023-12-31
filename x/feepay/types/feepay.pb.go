// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: juno/feepay/v1/feepay.proto

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

// This defines the address, balance, and wallet limit
// of a fee pay contract.
type FeePayContract struct {
	// The address of the contract.
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// The ledger balance of the contract.
	Balance uint64 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	// The number of times a wallet may interact with the contract.
	WalletLimit uint64 `protobuf:"varint,3,opt,name=wallet_limit,json=walletLimit,proto3" json:"wallet_limit,omitempty"`
}

func (m *FeePayContract) Reset()         { *m = FeePayContract{} }
func (m *FeePayContract) String() string { return proto.CompactTextString(m) }
func (*FeePayContract) ProtoMessage()    {}
func (*FeePayContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ea6771eacbfed1, []int{0}
}
func (m *FeePayContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeePayContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeePayContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeePayContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeePayContract.Merge(m, src)
}
func (m *FeePayContract) XXX_Size() int {
	return m.Size()
}
func (m *FeePayContract) XXX_DiscardUnknown() {
	xxx_messageInfo_FeePayContract.DiscardUnknown(m)
}

var xxx_messageInfo_FeePayContract proto.InternalMessageInfo

func (m *FeePayContract) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *FeePayContract) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *FeePayContract) GetWalletLimit() uint64 {
	if m != nil {
		return m.WalletLimit
	}
	return 0
}

// This object is used to store the number of times a wallet has
// interacted with a contract.
type FeePayWalletUsage struct {
	// The contract address.
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// The wallet address.
	WalletAddress string `protobuf:"bytes,2,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	// The number of uses corresponding to a wallet.
	Uses uint64 `protobuf:"varint,3,opt,name=uses,proto3" json:"uses,omitempty"`
}

func (m *FeePayWalletUsage) Reset()         { *m = FeePayWalletUsage{} }
func (m *FeePayWalletUsage) String() string { return proto.CompactTextString(m) }
func (*FeePayWalletUsage) ProtoMessage()    {}
func (*FeePayWalletUsage) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ea6771eacbfed1, []int{1}
}
func (m *FeePayWalletUsage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeePayWalletUsage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeePayWalletUsage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeePayWalletUsage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeePayWalletUsage.Merge(m, src)
}
func (m *FeePayWalletUsage) XXX_Size() int {
	return m.Size()
}
func (m *FeePayWalletUsage) XXX_DiscardUnknown() {
	xxx_messageInfo_FeePayWalletUsage.DiscardUnknown(m)
}

var xxx_messageInfo_FeePayWalletUsage proto.InternalMessageInfo

func (m *FeePayWalletUsage) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *FeePayWalletUsage) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *FeePayWalletUsage) GetUses() uint64 {
	if m != nil {
		return m.Uses
	}
	return 0
}

func init() {
	proto.RegisterType((*FeePayContract)(nil), "juno.feepay.v1.FeePayContract")
	proto.RegisterType((*FeePayWalletUsage)(nil), "juno.feepay.v1.FeePayWalletUsage")
}

func init() { proto.RegisterFile("juno/feepay/v1/feepay.proto", fileDescriptor_14ea6771eacbfed1) }

var fileDescriptor_14ea6771eacbfed1 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x2a, 0xcd, 0xcb,
	0xd7, 0x4f, 0x4b, 0x4d, 0x2d, 0x48, 0xac, 0xd4, 0x2f, 0x33, 0x84, 0xb2, 0xf4, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0xf8, 0x40, 0x92, 0x7a, 0x50, 0xa1, 0x32, 0x43, 0xa5, 0x0a, 0x2e, 0x3e, 0xb7,
	0xd4, 0xd4, 0x80, 0xc4, 0x4a, 0xe7, 0xfc, 0xbc, 0x92, 0xa2, 0xc4, 0xe4, 0x12, 0x21, 0x4d, 0x2e,
	0x81, 0x64, 0x28, 0x3b, 0x3e, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x88, 0x1f, 0x26, 0xee, 0x08, 0x11, 0x16, 0x92, 0xe0, 0x62, 0x4f, 0x4a, 0xcc, 0x49,
	0xcc, 0x4b, 0x4e, 0x95, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x09, 0x82, 0x71, 0x85, 0x14, 0xb9, 0x78,
	0xca, 0x13, 0x73, 0x72, 0x52, 0x4b, 0xe2, 0x73, 0x32, 0x73, 0x33, 0x4b, 0x24, 0x98, 0xc1, 0xd2,
	0xdc, 0x10, 0x31, 0x1f, 0x90, 0x90, 0x52, 0x25, 0x97, 0x20, 0xc4, 0xe6, 0x70, 0xb0, 0x60, 0x68,
	0x71, 0x62, 0x7a, 0x2a, 0x29, 0x96, 0xab, 0x72, 0xf1, 0x41, 0xad, 0x80, 0x29, 0x64, 0x02, 0x2b,
	0xe4, 0x85, 0x88, 0xc2, 0x94, 0x09, 0x71, 0xb1, 0x94, 0x16, 0xa7, 0x16, 0x43, 0x5d, 0x00, 0x66,
	0x3b, 0x79, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x5e, 0x7a, 0x66,
	0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x73, 0x7e, 0x71, 0x6e, 0x7e, 0x31, 0x2c,
	0x5c, 0x8a, 0xf5, 0xc1, 0xc1, 0x5a, 0x01, 0x0b, 0xd8, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36,
	0x70, 0xa8, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x06, 0xb9, 0x8c, 0x74, 0x01, 0x00,
	0x00,
}

func (m *FeePayContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeePayContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeePayContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WalletLimit != 0 {
		i = encodeVarintFeepay(dAtA, i, uint64(m.WalletLimit))
		i--
		dAtA[i] = 0x18
	}
	if m.Balance != 0 {
		i = encodeVarintFeepay(dAtA, i, uint64(m.Balance))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintFeepay(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FeePayWalletUsage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeePayWalletUsage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeePayWalletUsage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Uses != 0 {
		i = encodeVarintFeepay(dAtA, i, uint64(m.Uses))
		i--
		dAtA[i] = 0x18
	}
	if len(m.WalletAddress) > 0 {
		i -= len(m.WalletAddress)
		copy(dAtA[i:], m.WalletAddress)
		i = encodeVarintFeepay(dAtA, i, uint64(len(m.WalletAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintFeepay(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFeepay(dAtA []byte, offset int, v uint64) int {
	offset -= sovFeepay(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FeePayContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovFeepay(uint64(l))
	}
	if m.Balance != 0 {
		n += 1 + sovFeepay(uint64(m.Balance))
	}
	if m.WalletLimit != 0 {
		n += 1 + sovFeepay(uint64(m.WalletLimit))
	}
	return n
}

func (m *FeePayWalletUsage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovFeepay(uint64(l))
	}
	l = len(m.WalletAddress)
	if l > 0 {
		n += 1 + l + sovFeepay(uint64(l))
	}
	if m.Uses != 0 {
		n += 1 + sovFeepay(uint64(m.Uses))
	}
	return n
}

func sovFeepay(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFeepay(x uint64) (n int) {
	return sovFeepay(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FeePayContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeepay
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
			return fmt.Errorf("proto: FeePayContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeePayContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeepay
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
				return ErrInvalidLengthFeepay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeepay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			m.Balance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeepay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Balance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletLimit", wireType)
			}
			m.WalletLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeepay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WalletLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFeepay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeepay
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
func (m *FeePayWalletUsage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeepay
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
			return fmt.Errorf("proto: FeePayWalletUsage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeePayWalletUsage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeepay
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
				return ErrInvalidLengthFeepay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeepay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeepay
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
				return ErrInvalidLengthFeepay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeepay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WalletAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uses", wireType)
			}
			m.Uses = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeepay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uses |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFeepay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeepay
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
func skipFeepay(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeepay
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
					return 0, ErrIntOverflowFeepay
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
					return 0, ErrIntOverflowFeepay
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
				return 0, ErrInvalidLengthFeepay
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFeepay
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFeepay
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFeepay        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeepay          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFeepay = fmt.Errorf("proto: unexpected end of group")
)
