// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: receipt.proto

package receipt

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_ElrondNetwork_elrond_go_core_data "github.com/multiversx/mx-chain-core-go/data"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// Receipt holds all the data needed for a transaction receipt
type Receipt struct {
	Value   *math_big.Int `protobuf:"bytes,1,opt,name=Value,proto3,casttypewith=math/big.Int;github.com/multiversx/mx-chain-core-go/data.BigIntCaster" json:"value"`
	SndAddr []byte        `protobuf:"bytes,2,opt,name=SndAddr,proto3" json:"sender"`
	Data    []byte        `protobuf:"bytes,3,opt,name=Data,proto3" json:"data,omitempty"`
	TxHash  []byte        `protobuf:"bytes,4,opt,name=TxHash,proto3" json:"txHash"`
}

func (m *Receipt) Reset()      { *m = Receipt{} }
func (*Receipt) ProtoMessage() {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_ace1d6eb38fad2c8, []int{0}
}
func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(m, src)
}
func (m *Receipt) XXX_Size() int {
	return m.Size()
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetValue() *math_big.Int {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Receipt) GetSndAddr() []byte {
	if m != nil {
		return m.SndAddr
	}
	return nil
}

func (m *Receipt) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Receipt) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func init() {
	proto.RegisterType((*Receipt)(nil), "proto.Receipt")
}

func init() { proto.RegisterFile("receipt.proto", fileDescriptor_ace1d6eb38fad2c8) }

var fileDescriptor_ace1d6eb38fad2c8 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0xbd, 0x4e, 0x33, 0x31,
	0x10, 0x3c, 0x7f, 0x5f, 0x7e, 0x24, 0x0b, 0x28, 0x5c, 0x9d, 0x28, 0xf6, 0x50, 0x84, 0x10, 0x05,
	0xb9, 0x2b, 0x28, 0xa9, 0x12, 0x7e, 0x44, 0x1a, 0x8a, 0x0b, 0xa2, 0xa0, 0xf3, 0xc5, 0xc6, 0xb1,
	0xc8, 0x9d, 0x23, 0x67, 0xc3, 0x4f, 0xc7, 0x23, 0xf0, 0x18, 0x88, 0x27, 0xa1, 0x4c, 0x99, 0x2a,
	0x10, 0x47, 0x42, 0x28, 0x55, 0x1e, 0x01, 0xc5, 0x09, 0x82, 0x6a, 0x77, 0x66, 0x76, 0x67, 0xa4,
	0xa1, 0x9b, 0x56, 0x76, 0xa4, 0xee, 0x63, 0xdc, 0xb7, 0x06, 0x0d, 0x2b, 0xfb, 0xb1, 0x5d, 0x57,
	0x1a, 0xbb, 0xc3, 0x2c, 0xee, 0x98, 0x3c, 0x51, 0x46, 0x99, 0xc4, 0xd3, 0xd9, 0xf0, 0xc6, 0x23,
	0x0f, 0xfc, 0xb6, 0xfa, 0xaa, 0x7d, 0x12, 0x5a, 0x4d, 0x57, 0x3e, 0x4c, 0xd3, 0xf2, 0x15, 0xef,
	0x0d, 0x65, 0x48, 0x76, 0xc8, 0xfe, 0x46, 0xb3, 0x3d, 0x9f, 0x44, 0xe5, 0xbb, 0x25, 0xf1, 0xfa,
	0x1e, 0x9d, 0xe5, 0x1c, 0xbb, 0x49, 0xa6, 0x55, 0xdc, 0x2a, 0xf0, 0xe8, 0x4f, 0xc6, 0x69, 0xcf,
	0x9a, 0x42, 0x5c, 0x48, 0xbc, 0x37, 0xf6, 0x36, 0x91, 0x1e, 0xd5, 0x95, 0xa9, 0x77, 0x8c, 0x95,
	0x89, 0xe0, 0xc8, 0xe3, 0xa6, 0x56, 0xad, 0x02, 0x8f, 0xf9, 0x00, 0xa5, 0x4d, 0x57, 0x09, 0x6c,
	0x97, 0x56, 0xdb, 0x85, 0x68, 0x08, 0x61, 0xc3, 0x7f, 0x3e, 0x8c, 0xce, 0x27, 0x51, 0x65, 0x20,
	0x0b, 0x21, 0x6d, 0xfa, 0x23, 0xb1, 0x3d, 0x5a, 0x3a, 0xe1, 0xc8, 0xc3, 0xff, 0xfe, 0x84, 0xcd,
	0x27, 0xd1, 0xd6, 0xd2, 0xf1, 0xc0, 0xe4, 0x1a, 0x65, 0xde, 0xc7, 0xc7, 0xd4, 0xeb, 0xac, 0x46,
	0x2b, 0x97, 0x0f, 0xe7, 0x7c, 0xd0, 0x0d, 0x4b, 0xbf, 0x66, 0xe8, 0x99, 0x74, 0xad, 0x34, 0x1b,
	0xa3, 0x29, 0x04, 0xe3, 0x29, 0x04, 0x8b, 0x29, 0x90, 0x27, 0x07, 0xe4, 0xc5, 0x01, 0x79, 0x73,
	0x40, 0x46, 0x0e, 0xc8, 0xd8, 0x01, 0xf9, 0x70, 0x40, 0xbe, 0x1c, 0x04, 0x0b, 0x07, 0xe4, 0x79,
	0x06, 0xc1, 0x68, 0x06, 0xc1, 0x78, 0x06, 0xc1, 0x75, 0x75, 0xdd, 0x73, 0x56, 0xf1, 0x95, 0x1d,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x65, 0xec, 0xc7, 0x7c, 0x79, 0x01, 0x00, 0x00,
}

func (this *Receipt) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Receipt)
	if !ok {
		that2, ok := that.(Receipt)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_core_data.BigIntCaster{}
		if !__caster.Equal(this.Value, that1.Value) {
			return false
		}
	}
	if !bytes.Equal(this.SndAddr, that1.SndAddr) {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if !bytes.Equal(this.TxHash, that1.TxHash) {
		return false
	}
	return true
}
func (this *Receipt) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&receipt.Receipt{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "SndAddr: "+fmt.Sprintf("%#v", this.SndAddr)+",\n")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "TxHash: "+fmt.Sprintf("%#v", this.TxHash)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringReceipt(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Receipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Receipt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Receipt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SndAddr) > 0 {
		i -= len(m.SndAddr)
		copy(dAtA[i:], m.SndAddr)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.SndAddr)))
		i--
		dAtA[i] = 0x12
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_core_data.BigIntCaster{}
		size := __caster.Size(m.Value)
		i -= size
		if _, err := __caster.MarshalTo(m.Value, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintReceipt(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintReceipt(dAtA []byte, offset int, v uint64) int {
	offset -= sovReceipt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Receipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	{
		__caster := &github_com_ElrondNetwork_elrond_go_core_data.BigIntCaster{}
		l = __caster.Size(m.Value)
		n += 1 + l + sovReceipt(uint64(l))
	}
	l = len(m.SndAddr)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	return n
}

func sovReceipt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReceipt(x uint64) (n int) {
	return sovReceipt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Receipt) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Receipt{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`SndAddr:` + fmt.Sprintf("%v", this.SndAddr) + `,`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`TxHash:` + fmt.Sprintf("%v", this.TxHash) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringReceipt(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Receipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReceipt
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
			return fmt.Errorf("proto: Receipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Receipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_core_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Value = tmp
				}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SndAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SndAddr = append(m.SndAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.SndAddr == nil {
				m.SndAddr = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReceipt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReceipt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReceipt
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
func skipReceipt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReceipt
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
					return 0, ErrIntOverflowReceipt
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
					return 0, ErrIntOverflowReceipt
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
				return 0, ErrInvalidLengthReceipt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReceipt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReceipt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReceipt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReceipt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReceipt = fmt.Errorf("proto: unexpected end of group")
)
