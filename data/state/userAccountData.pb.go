// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: userAccountData.proto

package state

import (
	bytes "bytes"
	fmt "fmt"
	github_com_ElrondNetwork_elrond_go_data "github.com/ElrondNetwork/elrond-go/data"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type UserAccountData struct {
	Nonce           uint64        `protobuf:"varint,1,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Balance         *math_big.Int `protobuf:"bytes,2,opt,name=Balance,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"Balance,omitempty"`
	CodeHash        []byte        `protobuf:"bytes,3,opt,name=CodeHash,proto3" json:"CodeHash,omitempty"`
	RootHash        []byte        `protobuf:"bytes,4,opt,name=RootHash,proto3" json:"RootHash,omitempty"`
	Address         []byte        `protobuf:"bytes,5,opt,name=Address,proto3" json:"Address,omitempty"`
	DeveloperReward *math_big.Int `protobuf:"bytes,6,opt,name=DeveloperReward,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"DeveloperReward,omitempty"`
	OwnerAddress    []byte        `protobuf:"bytes,7,opt,name=OwnerAddress,proto3" json:"OwnerAddress,omitempty"`
}

func (m *UserAccountData) Reset()      { *m = UserAccountData{} }
func (*UserAccountData) ProtoMessage() {}
func (*UserAccountData) Descriptor() ([]byte, []int) {
	return fileDescriptor_275d64df7d722770, []int{0}
}
func (m *UserAccountData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserAccountData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *UserAccountData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAccountData.Merge(m, src)
}
func (m *UserAccountData) XXX_Size() int {
	return m.Size()
}
func (m *UserAccountData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAccountData.DiscardUnknown(m)
}

var xxx_messageInfo_UserAccountData proto.InternalMessageInfo

func (m *UserAccountData) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *UserAccountData) GetBalance() *math_big.Int {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *UserAccountData) GetCodeHash() []byte {
	if m != nil {
		return m.CodeHash
	}
	return nil
}

func (m *UserAccountData) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *UserAccountData) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *UserAccountData) GetDeveloperReward() *math_big.Int {
	if m != nil {
		return m.DeveloperReward
	}
	return nil
}

func (m *UserAccountData) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*UserAccountData)(nil), "proto.UserAccountData")
}

func init() { proto.RegisterFile("userAccountData.proto", fileDescriptor_275d64df7d722770) }

var fileDescriptor_275d64df7d722770 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x31, 0x4f, 0xc2, 0x40,
	0x1c, 0xc5, 0x7b, 0x48, 0xc1, 0x5c, 0x48, 0x48, 0x2e, 0x9a, 0x34, 0x0c, 0x7f, 0x09, 0x13, 0x0b,
	0x74, 0x70, 0x74, 0x30, 0x14, 0x48, 0x64, 0xc1, 0xa4, 0x89, 0x8b, 0x8b, 0xb9, 0xb6, 0x67, 0x21,
	0x42, 0xff, 0xe4, 0x7a, 0x95, 0xd5, 0x8f, 0xe0, 0xc7, 0x30, 0x7e, 0x12, 0x47, 0x46, 0x36, 0xe5,
	0x58, 0xdc, 0xe4, 0x23, 0x98, 0x1e, 0xa9, 0x11, 0x67, 0xa7, 0xbb, 0xdf, 0x7b, 0xb9, 0xf7, 0x5e,
	0x72, 0xf4, 0x34, 0x4b, 0x85, 0xec, 0x85, 0x21, 0x66, 0x89, 0x1a, 0x70, 0xc5, 0xbb, 0x0b, 0x89,
	0x0a, 0x99, 0x6d, 0x8e, 0x46, 0x27, 0x9e, 0xaa, 0x49, 0x16, 0x74, 0x43, 0x9c, 0xbb, 0x31, 0xc6,
	0xe8, 0x1a, 0x39, 0xc8, 0xee, 0x0d, 0x19, 0x30, 0xb7, 0xfd, 0xab, 0xd6, 0x57, 0x89, 0xd6, 0x6f,
	0x0e, 0xf3, 0xd8, 0x09, 0xb5, 0xc7, 0x98, 0x84, 0xc2, 0x21, 0x4d, 0xd2, 0x2e, 0xfb, 0x7b, 0x60,
	0x77, 0xb4, 0xea, 0xf1, 0x19, 0xcf, 0xf5, 0x52, 0x93, 0xb4, 0x6b, 0xde, 0xf0, 0xf5, 0xfd, 0xac,
	0x37, 0xe7, 0x6a, 0xe2, 0x06, 0xd3, 0xb8, 0x3b, 0x4a, 0xd4, 0xc5, 0xaf, 0xea, 0xe1, 0x4c, 0x62,
	0x12, 0x8d, 0x85, 0x5a, 0xa2, 0x7c, 0x70, 0x85, 0xa1, 0x4e, 0x8c, 0x6e, 0x94, 0x0f, 0xf6, 0xa6,
	0xf1, 0x28, 0x51, 0x7d, 0x9e, 0x2a, 0x21, 0xfd, 0x22, 0x95, 0x35, 0xe8, 0x71, 0x1f, 0x23, 0x71,
	0xc5, 0xd3, 0x89, 0x73, 0x94, 0x37, 0xf8, 0x3f, 0x9c, 0x7b, 0x3e, 0xa2, 0x32, 0x5e, 0x79, 0xef,
	0x15, 0xcc, 0x1c, 0x5a, 0xed, 0x45, 0x91, 0x14, 0x69, 0xea, 0xd8, 0xc6, 0x2a, 0x90, 0x21, 0xad,
	0x0f, 0xc4, 0xa3, 0x98, 0xe1, 0x42, 0x48, 0x5f, 0x2c, 0xb9, 0x8c, 0x9c, 0xca, 0x7f, 0x4e, 0xff,
	0x9b, 0xce, 0x5a, 0xb4, 0x76, 0xbd, 0x4c, 0x84, 0x2c, 0xf6, 0x54, 0xcd, 0x9e, 0x03, 0xcd, 0xbb,
	0x5c, 0x6d, 0xc0, 0x5a, 0x6f, 0xc0, 0xda, 0x6d, 0x80, 0x3c, 0x69, 0x20, 0x2f, 0x1a, 0xc8, 0x9b,
	0x06, 0xb2, 0xd2, 0x40, 0xd6, 0x1a, 0xc8, 0x87, 0x06, 0xf2, 0xa9, 0xc1, 0xda, 0x69, 0x20, 0xcf,
	0x5b, 0xb0, 0x56, 0x5b, 0xb0, 0xd6, 0x5b, 0xb0, 0x6e, 0xed, 0x54, 0x71, 0x25, 0x82, 0x8a, 0xf9,
	0xb9, 0xf3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x75, 0x36, 0x4d, 0x08, 0x02, 0x00, 0x00,
}

func (this *UserAccountData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserAccountData)
	if !ok {
		that2, ok := that.(UserAccountData)
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
	if this.Nonce != that1.Nonce {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.Balance, that1.Balance) {
			return false
		}
	}
	if !bytes.Equal(this.CodeHash, that1.CodeHash) {
		return false
	}
	if !bytes.Equal(this.RootHash, that1.RootHash) {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.DeveloperReward, that1.DeveloperReward) {
			return false
		}
	}
	if !bytes.Equal(this.OwnerAddress, that1.OwnerAddress) {
		return false
	}
	return true
}
func (this *UserAccountData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&state.UserAccountData{")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "Balance: "+fmt.Sprintf("%#v", this.Balance)+",\n")
	s = append(s, "CodeHash: "+fmt.Sprintf("%#v", this.CodeHash)+",\n")
	s = append(s, "RootHash: "+fmt.Sprintf("%#v", this.RootHash)+",\n")
	s = append(s, "Address: "+fmt.Sprintf("%#v", this.Address)+",\n")
	s = append(s, "DeveloperReward: "+fmt.Sprintf("%#v", this.DeveloperReward)+",\n")
	s = append(s, "OwnerAddress: "+fmt.Sprintf("%#v", this.OwnerAddress)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringUserAccountData(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *UserAccountData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserAccountData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserAccountData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintUserAccountData(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x3a
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.DeveloperReward)
		i -= size
		if _, err := __caster.MarshalTo(m.DeveloperReward, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintUserAccountData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintUserAccountData(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RootHash) > 0 {
		i -= len(m.RootHash)
		copy(dAtA[i:], m.RootHash)
		i = encodeVarintUserAccountData(dAtA, i, uint64(len(m.RootHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CodeHash) > 0 {
		i -= len(m.CodeHash)
		copy(dAtA[i:], m.CodeHash)
		i = encodeVarintUserAccountData(dAtA, i, uint64(len(m.CodeHash)))
		i--
		dAtA[i] = 0x1a
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.Balance)
		i -= size
		if _, err := __caster.MarshalTo(m.Balance, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintUserAccountData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Nonce != 0 {
		i = encodeVarintUserAccountData(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintUserAccountData(dAtA []byte, offset int, v uint64) int {
	offset -= sovUserAccountData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserAccountData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovUserAccountData(uint64(m.Nonce))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.Balance)
		n += 1 + l + sovUserAccountData(uint64(l))
	}
	l = len(m.CodeHash)
	if l > 0 {
		n += 1 + l + sovUserAccountData(uint64(l))
	}
	l = len(m.RootHash)
	if l > 0 {
		n += 1 + l + sovUserAccountData(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovUserAccountData(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.DeveloperReward)
		n += 1 + l + sovUserAccountData(uint64(l))
	}
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovUserAccountData(uint64(l))
	}
	return n
}

func sovUserAccountData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUserAccountData(x uint64) (n int) {
	return sovUserAccountData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UserAccountData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UserAccountData{`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Balance:` + fmt.Sprintf("%v", this.Balance) + `,`,
		`CodeHash:` + fmt.Sprintf("%v", this.CodeHash) + `,`,
		`RootHash:` + fmt.Sprintf("%v", this.RootHash) + `,`,
		`Address:` + fmt.Sprintf("%v", this.Address) + `,`,
		`DeveloperReward:` + fmt.Sprintf("%v", this.DeveloperReward) + `,`,
		`OwnerAddress:` + fmt.Sprintf("%v", this.OwnerAddress) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringUserAccountData(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UserAccountData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserAccountData
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
			return fmt.Errorf("proto: UserAccountData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserAccountData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserAccountData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserAccountData
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
				return ErrInvalidLengthUserAccountData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUserAccountData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Balance = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserAccountData
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
				return ErrInvalidLengthUserAccountData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUserAccountData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeHash = append(m.CodeHash[:0], dAtA[iNdEx:postIndex]...)
			if m.CodeHash == nil {
				m.CodeHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserAccountData
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
				return ErrInvalidLengthUserAccountData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUserAccountData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootHash = append(m.RootHash[:0], dAtA[iNdEx:postIndex]...)
			if m.RootHash == nil {
				m.RootHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserAccountData
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
				return ErrInvalidLengthUserAccountData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUserAccountData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeveloperReward", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserAccountData
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
				return ErrInvalidLengthUserAccountData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUserAccountData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.DeveloperReward = tmp
				}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserAccountData
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
				return ErrInvalidLengthUserAccountData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUserAccountData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = append(m.OwnerAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.OwnerAddress == nil {
				m.OwnerAddress = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserAccountData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUserAccountData
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUserAccountData
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
func skipUserAccountData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserAccountData
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
					return 0, ErrIntOverflowUserAccountData
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
					return 0, ErrIntOverflowUserAccountData
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
				return 0, ErrInvalidLengthUserAccountData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUserAccountData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUserAccountData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUserAccountData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserAccountData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUserAccountData = fmt.Errorf("proto: unexpected end of group")
)
