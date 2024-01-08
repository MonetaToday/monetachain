// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: monetachain/tokenfactory/denom.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

type Denom struct {
	Denom              string                `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Description        string                `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Ticker             string                `protobuf:"bytes,3,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Precision          int32                 `protobuf:"varint,4,opt,name=precision,proto3" json:"precision,omitempty"`
	Url                string                `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	MaxSupply          cosmossdk_io_math.Int `protobuf:"bytes,6,opt,name=maxSupply,proto3,customtype=cosmossdk.io/math.Int" json:"maxSupply"`
	Supply             cosmossdk_io_math.Int `protobuf:"bytes,7,opt,name=supply,proto3,customtype=cosmossdk.io/math.Int" json:"supply"`
	CanChangeMaxSupply bool                  `protobuf:"varint,8,opt,name=canChangeMaxSupply,proto3" json:"canChangeMaxSupply,omitempty"`
	Owner              string                `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *Denom) Reset()         { *m = Denom{} }
func (m *Denom) String() string { return proto.CompactTextString(m) }
func (*Denom) ProtoMessage()    {}
func (*Denom) Descriptor() ([]byte, []int) {
	return fileDescriptor_57bb3f4a9abe49ed, []int{0}
}
func (m *Denom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Denom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Denom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Denom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Denom.Merge(m, src)
}
func (m *Denom) XXX_Size() int {
	return m.Size()
}
func (m *Denom) XXX_DiscardUnknown() {
	xxx_messageInfo_Denom.DiscardUnknown(m)
}

var xxx_messageInfo_Denom proto.InternalMessageInfo

func (m *Denom) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Denom) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Denom) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *Denom) GetPrecision() int32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *Denom) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Denom) GetCanChangeMaxSupply() bool {
	if m != nil {
		return m.CanChangeMaxSupply
	}
	return false
}

func (m *Denom) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*Denom)(nil), "monetachain.tokenfactory.Denom")
}

func init() {
	proto.RegisterFile("monetachain/tokenfactory/denom.proto", fileDescriptor_57bb3f4a9abe49ed)
}

var fileDescriptor_57bb3f4a9abe49ed = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x8d, 0x5b, 0x12, 0x1a, 0xb3, 0x80, 0x55, 0x50, 0xa8, 0x50, 0x1a, 0x21, 0x86, 0x0a, 0x89,
	0x04, 0x89, 0x8d, 0xb1, 0x30, 0xd0, 0x01, 0x86, 0xb0, 0xb1, 0x20, 0x93, 0x9a, 0x36, 0x6a, 0xe3,
	0x8b, 0x62, 0x57, 0xb4, 0x7f, 0xc1, 0x67, 0x30, 0x32, 0xf0, 0x11, 0x1d, 0x2b, 0x26, 0xc4, 0x50,
	0xa1, 0x76, 0xe8, 0x6f, 0xa0, 0xd8, 0x11, 0x0d, 0x12, 0x13, 0x4b, 0x74, 0xef, 0xbd, 0xbb, 0x77,
	0xb9, 0xf3, 0xe1, 0xa3, 0x04, 0x38, 0x93, 0x34, 0xea, 0xd3, 0x98, 0x07, 0x12, 0x06, 0x8c, 0x3f,
	0xd2, 0x48, 0x42, 0x36, 0x09, 0xba, 0x8c, 0x43, 0xe2, 0xa7, 0x19, 0x48, 0x20, 0x4e, 0x29, 0xcb,
	0x2f, 0x67, 0x35, 0x76, 0x68, 0x12, 0x73, 0x08, 0xd4, 0x57, 0x27, 0x37, 0xf6, 0x23, 0x10, 0x09,
	0x88, 0x7b, 0x85, 0x02, 0x0d, 0x0a, 0xa9, 0xde, 0x83, 0x1e, 0x68, 0x3e, 0x8f, 0x34, 0x7b, 0xb8,
	0xaa, 0x60, 0xf3, 0x32, 0xef, 0x46, 0xea, 0xd8, 0x54, 0x6d, 0x1d, 0xe4, 0xa1, 0x96, 0x1d, 0x6a,
	0x40, 0x3c, 0xbc, 0xd5, 0x65, 0x22, 0xca, 0xe2, 0x54, 0xc6, 0xc0, 0x9d, 0x8a, 0xd2, 0xca, 0x14,
	0xd9, 0xc3, 0x96, 0x8c, 0xa3, 0x01, 0xcb, 0x9c, 0xaa, 0x12, 0x0b, 0x44, 0x0e, 0xb0, 0x9d, 0x66,
	0x2c, 0x8a, 0x45, 0x5e, 0xb7, 0xe1, 0xa1, 0x96, 0x19, 0xae, 0x09, 0xb2, 0x8d, 0xab, 0xa3, 0x6c,
	0xe8, 0x98, 0xaa, 0x24, 0x0f, 0xc9, 0x0d, 0xb6, 0x13, 0x3a, 0xbe, 0x1d, 0xa5, 0xe9, 0x70, 0xe2,
	0x58, 0x39, 0xdf, 0x3e, 0x9d, 0xce, 0x9b, 0xc6, 0xe7, 0xbc, 0xb9, 0xab, 0x07, 0x11, 0xdd, 0x81,
	0x1f, 0x43, 0x90, 0x50, 0xd9, 0xf7, 0x3b, 0x5c, 0xbe, 0xbf, 0x9d, 0xe0, 0x62, 0xc2, 0x0e, 0x97,
	0x2f, 0xab, 0xd7, 0x63, 0x14, 0xae, 0x2d, 0xc8, 0x15, 0xb6, 0x84, 0x36, 0xdb, 0xfc, 0xa7, 0x59,
	0x51, 0x4f, 0x7c, 0x4c, 0x22, 0xca, 0x2f, 0xfa, 0x94, 0xf7, 0xd8, 0xf5, 0xcf, 0x2f, 0xd6, 0x3c,
	0xd4, 0xaa, 0x85, 0x7f, 0x28, 0xf9, 0x26, 0xe1, 0x89, 0xb3, 0xcc, 0xb1, 0xf5, 0x26, 0x15, 0x68,
	0x9f, 0x4f, 0x17, 0x2e, 0x9a, 0x2d, 0x5c, 0xf4, 0xb5, 0x70, 0xd1, 0xf3, 0xd2, 0x35, 0x66, 0x4b,
	0xd7, 0xf8, 0x58, 0xba, 0xc6, 0x9d, 0x57, 0xbe, 0x83, 0xf1, 0xef, 0x4b, 0x90, 0x93, 0x94, 0x89,
	0x07, 0x4b, 0x3d, 0xd6, 0xd9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x8f, 0x0b, 0xb5, 0x32,
	0x02, 0x00, 0x00,
}

func (m *Denom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Denom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Denom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x4a
	}
	if m.CanChangeMaxSupply {
		i--
		if m.CanChangeMaxSupply {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	{
		size := m.Supply.Size()
		i -= size
		if _, err := m.Supply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDenom(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.MaxSupply.Size()
		i -= size
		if _, err := m.MaxSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDenom(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Precision != 0 {
		i = encodeVarintDenom(dAtA, i, uint64(m.Precision))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Ticker) > 0 {
		i -= len(m.Ticker)
		copy(dAtA[i:], m.Ticker)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Ticker)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDenom(dAtA []byte, offset int, v uint64) int {
	offset -= sovDenom(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Denom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	l = len(m.Ticker)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	if m.Precision != 0 {
		n += 1 + sovDenom(uint64(m.Precision))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	l = m.MaxSupply.Size()
	n += 1 + l + sovDenom(uint64(l))
	l = m.Supply.Size()
	n += 1 + l + sovDenom(uint64(l))
	if m.CanChangeMaxSupply {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	return n
}

func sovDenom(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDenom(x uint64) (n int) {
	return sovDenom(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Denom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDenom
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
			return fmt.Errorf("proto: Denom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Denom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ticker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
			}
			m.Precision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Precision |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Supply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanChangeMaxSupply", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
			m.CanChangeMaxSupply = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDenom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDenom
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
func skipDenom(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDenom
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
					return 0, ErrIntOverflowDenom
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
					return 0, ErrIntOverflowDenom
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
				return 0, ErrInvalidLengthDenom
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDenom
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDenom
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDenom        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDenom          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDenom = fmt.Errorf("proto: unexpected end of group")
)
