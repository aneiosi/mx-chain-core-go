// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validatorStatus.proto

package validator

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
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

// ValidatorStatus holds information about a validator
type ValidatorStatus struct {
	TempRating                         float32 `protobuf:"fixed32,1,opt,name=TempRating,proto3" json:"tempRating"`
	NumLeaderSuccess                   uint32  `protobuf:"varint,2,opt,name=NumLeaderSuccess,proto3" json:"numLeaderSuccess"`
	NumLeaderFailure                   uint32  `protobuf:"varint,3,opt,name=NumLeaderFailure,proto3" json:"numLeaderFailure"`
	NumValidatorSuccess                uint32  `protobuf:"varint,4,opt,name=NumValidatorSuccess,proto3" json:"numValidatorSuccess"`
	NumValidatorFailure                uint32  `protobuf:"varint,5,opt,name=NumValidatorFailure,proto3" json:"numValidatorFailure"`
	NumValidatorIgnoredSignatures      uint32  `protobuf:"varint,6,opt,name=NumValidatorIgnoredSignatures,proto3" json:"numValidatorIgnoredSignatures"`
	Rating                             float32 `protobuf:"fixed32,7,opt,name=Rating,proto3" json:"rating"`
	RatingModifier                     float32 `protobuf:"fixed32,8,opt,name=RatingModifier,proto3" json:"ratingModifier"`
	TotalNumLeaderSuccess              uint32  `protobuf:"varint,9,opt,name=TotalNumLeaderSuccess,proto3" json:"totalNumLeaderSuccess"`
	TotalNumLeaderFailure              uint32  `protobuf:"varint,10,opt,name=TotalNumLeaderFailure,proto3" json:"totalNumLeaderFailure"`
	TotalNumValidatorSuccess           uint32  `protobuf:"varint,11,opt,name=TotalNumValidatorSuccess,proto3" json:"totalNumValidatorSuccess"`
	TotalNumValidatorFailure           uint32  `protobuf:"varint,12,opt,name=TotalNumValidatorFailure,proto3" json:"totalNumValidatorFailure"`
	TotalNumValidatorIgnoredSignatures uint32  `protobuf:"varint,13,opt,name=TotalNumValidatorIgnoredSignatures,proto3" json:"totalNumValidatorIgnoredSignatures"`
	ShardId                            uint32  `protobuf:"varint,14,opt,name=ShardId,proto3" json:"shardId"`
	ValidatorStatus                    string  `protobuf:"bytes,15,opt,name=ValidatorStatus,proto3" json:"validatorStatus,omitempty"`
}

func (m *ValidatorStatus) Reset()      { *m = ValidatorStatus{} }
func (*ValidatorStatus) ProtoMessage() {}
func (*ValidatorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_42eb597ab5bfcbec, []int{0}
}
func (m *ValidatorStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ValidatorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorStatus.Merge(m, src)
}
func (m *ValidatorStatus) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorStatus proto.InternalMessageInfo

func (m *ValidatorStatus) GetTempRating() float32 {
	if m != nil {
		return m.TempRating
	}
	return 0
}

func (m *ValidatorStatus) GetNumLeaderSuccess() uint32 {
	if m != nil {
		return m.NumLeaderSuccess
	}
	return 0
}

func (m *ValidatorStatus) GetNumLeaderFailure() uint32 {
	if m != nil {
		return m.NumLeaderFailure
	}
	return 0
}

func (m *ValidatorStatus) GetNumValidatorSuccess() uint32 {
	if m != nil {
		return m.NumValidatorSuccess
	}
	return 0
}

func (m *ValidatorStatus) GetNumValidatorFailure() uint32 {
	if m != nil {
		return m.NumValidatorFailure
	}
	return 0
}

func (m *ValidatorStatus) GetNumValidatorIgnoredSignatures() uint32 {
	if m != nil {
		return m.NumValidatorIgnoredSignatures
	}
	return 0
}

func (m *ValidatorStatus) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *ValidatorStatus) GetRatingModifier() float32 {
	if m != nil {
		return m.RatingModifier
	}
	return 0
}

func (m *ValidatorStatus) GetTotalNumLeaderSuccess() uint32 {
	if m != nil {
		return m.TotalNumLeaderSuccess
	}
	return 0
}

func (m *ValidatorStatus) GetTotalNumLeaderFailure() uint32 {
	if m != nil {
		return m.TotalNumLeaderFailure
	}
	return 0
}

func (m *ValidatorStatus) GetTotalNumValidatorSuccess() uint32 {
	if m != nil {
		return m.TotalNumValidatorSuccess
	}
	return 0
}

func (m *ValidatorStatus) GetTotalNumValidatorFailure() uint32 {
	if m != nil {
		return m.TotalNumValidatorFailure
	}
	return 0
}

func (m *ValidatorStatus) GetTotalNumValidatorIgnoredSignatures() uint32 {
	if m != nil {
		return m.TotalNumValidatorIgnoredSignatures
	}
	return 0
}

func (m *ValidatorStatus) GetShardId() uint32 {
	if m != nil {
		return m.ShardId
	}
	return 0
}

func (m *ValidatorStatus) GetValidatorStatus() string {
	if m != nil {
		return m.ValidatorStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*ValidatorStatus)(nil), "proto.ValidatorStatus")
}

func init() { proto.RegisterFile("validatorStatus.proto", fileDescriptor_42eb597ab5bfcbec) }

var fileDescriptor_42eb597ab5bfcbec = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xe3, 0xc1, 0xda, 0xd5, 0x63, 0xdd, 0xe4, 0x51, 0x91, 0x21, 0x6a, 0x97, 0x4a, 0xa0,
	0x1e, 0xa0, 0x3b, 0x70, 0xe3, 0x84, 0x8a, 0x04, 0xaa, 0x04, 0x43, 0x72, 0x27, 0x84, 0xb8, 0xa5,
	0x4d, 0x96, 0x45, 0x6a, 0xe2, 0xc9, 0x71, 0x26, 0x71, 0xe3, 0x11, 0x78, 0x0c, 0x1e, 0x85, 0x63,
	0x8f, 0x3d, 0x59, 0xd4, 0xbd, 0x20, 0x9f, 0xf6, 0x08, 0x08, 0xb7, 0x19, 0x6d, 0x9c, 0x16, 0x4e,
	0x49, 0xbe, 0xff, 0xff, 0xff, 0xcb, 0x97, 0x7c, 0xd6, 0x07, 0x1b, 0xd7, 0xde, 0x38, 0xf2, 0x3d,
	0xc1, 0xf8, 0x40, 0x78, 0x22, 0x4b, 0xbb, 0x57, 0x9c, 0x09, 0x86, 0x76, 0xcd, 0xe5, 0xe1, 0xf3,
	0x30, 0x12, 0x97, 0xd9, 0xb0, 0x3b, 0x62, 0xf1, 0x69, 0xc8, 0x42, 0x76, 0x6a, 0xca, 0xc3, 0xec,
	0xc2, 0x3c, 0x99, 0x07, 0x73, 0xb7, 0x48, 0xb5, 0xd5, 0x1e, 0x3c, 0xfc, 0xb8, 0xce, 0x43, 0x5d,
	0x08, 0xcf, 0x83, 0xf8, 0x8a, 0x7a, 0x22, 0x4a, 0x42, 0x17, 0xb4, 0x40, 0x67, 0xa7, 0x57, 0xd7,
	0x92, 0x40, 0x71, 0x5b, 0xa5, 0x2b, 0x0e, 0xf4, 0x0a, 0x1e, 0x9d, 0x65, 0xf1, 0xbb, 0xc0, 0xf3,
	0x03, 0x3e, 0xc8, 0x46, 0xa3, 0x20, 0x4d, 0xdd, 0x9d, 0x16, 0xe8, 0x1c, 0xf4, 0xee, 0x6b, 0x49,
	0x8e, 0x92, 0x82, 0x46, 0x2d, 0xf7, 0x1a, 0xe1, 0x8d, 0x17, 0x8d, 0x33, 0x1e, 0xb8, 0x77, 0x4a,
	0x08, 0x4b, 0x8d, 0x5a, 0x6e, 0xd4, 0x87, 0xc7, 0x67, 0x59, 0xfc, 0xf7, 0x4b, 0x96, 0x6d, 0xdc,
	0x35, 0x90, 0x07, 0x5a, 0x92, 0xe3, 0xc4, 0x96, 0x69, 0x59, 0xa6, 0x88, 0xca, 0xfb, 0xd9, 0x2d,
	0x47, 0xe5, 0x2d, 0x95, 0x65, 0x50, 0x08, 0x9b, 0xab, 0xe5, 0x7e, 0x98, 0x30, 0x1e, 0xf8, 0x83,
	0x28, 0x4c, 0x3c, 0x91, 0xf1, 0x20, 0x75, 0x2b, 0x06, 0xfa, 0x58, 0x4b, 0xd2, 0x4c, 0xb6, 0x19,
	0xe9, 0x76, 0x0e, 0x6a, 0xc3, 0xca, 0x72, 0x5c, 0x55, 0x33, 0x2e, 0xa8, 0x25, 0xa9, 0xf0, 0xc5,
	0xa8, 0x96, 0x0a, 0x7a, 0x09, 0xeb, 0x8b, 0xbb, 0xf7, 0xcc, 0x8f, 0x2e, 0xa2, 0x80, 0xbb, 0x7b,
	0xc6, 0x8b, 0xb4, 0x24, 0x75, 0xbe, 0xa6, 0xd0, 0x82, 0x13, 0x7d, 0x80, 0x8d, 0x73, 0x26, 0xbc,
	0xb1, 0x35, 0xe7, 0x9a, 0xf9, 0x80, 0x13, 0x2d, 0x49, 0x43, 0x94, 0x19, 0x68, 0x79, 0xce, 0x06,
	0xe6, 0xbf, 0x19, 0x6e, 0x02, 0xe6, 0x3f, 0xba, 0x3c, 0x87, 0x3e, 0x41, 0x37, 0x17, 0xac, 0x53,
	0xb0, 0x6f, 0x98, 0x8f, 0xb4, 0x24, 0xae, 0xd8, 0xe0, 0xa1, 0x1b, 0xd3, 0xa5, 0xe4, 0xbc, 0xdb,
	0x7b, 0x5b, 0xc8, 0x79, 0xc3, 0x1b, 0xd3, 0xe8, 0x1a, 0xb6, 0x2d, 0xcd, 0x3e, 0x23, 0x07, 0xe6,
	0x1d, 0x4f, 0xb5, 0x24, 0x6d, 0xf1, 0x4f, 0x37, 0xfd, 0x0f, 0x22, 0x7a, 0x02, 0xab, 0x83, 0x4b,
	0x8f, 0xfb, 0x7d, 0xdf, 0xad, 0x1b, 0xf8, 0xbe, 0x96, 0xa4, 0x9a, 0x2e, 0x4a, 0x34, 0xd7, 0xd0,
	0x5b, 0x6b, 0x35, 0xb8, 0x87, 0x2d, 0xd0, 0xa9, 0xf5, 0x9a, 0x5a, 0x92, 0x93, 0xc2, 0x16, 0x7a,
	0xc6, 0xe2, 0xe8, 0xcf, 0x7e, 0x10, 0x5f, 0x68, 0x31, 0xd5, 0x7b, 0x3d, 0x99, 0x61, 0x67, 0x3a,
	0xc3, 0xce, 0xcd, 0x0c, 0x83, 0xaf, 0x0a, 0x83, 0xef, 0x0a, 0x83, 0x1f, 0x0a, 0x83, 0x89, 0xc2,
	0x60, 0xaa, 0x30, 0xf8, 0xa9, 0x30, 0xf8, 0xa5, 0xb0, 0x73, 0xa3, 0x30, 0xf8, 0x36, 0xc7, 0xce,
	0x64, 0x8e, 0x9d, 0xe9, 0x1c, 0x3b, 0x9f, 0x6b, 0xb7, 0x6f, 0x19, 0x56, 0xcc, 0xc2, 0x7a, 0xf1,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x3e, 0x4e, 0xe7, 0xff, 0x04, 0x00, 0x00,
}

func (this *ValidatorStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorStatus)
	if !ok {
		that2, ok := that.(ValidatorStatus)
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
	if this.TempRating != that1.TempRating {
		return false
	}
	if this.NumLeaderSuccess != that1.NumLeaderSuccess {
		return false
	}
	if this.NumLeaderFailure != that1.NumLeaderFailure {
		return false
	}
	if this.NumValidatorSuccess != that1.NumValidatorSuccess {
		return false
	}
	if this.NumValidatorFailure != that1.NumValidatorFailure {
		return false
	}
	if this.NumValidatorIgnoredSignatures != that1.NumValidatorIgnoredSignatures {
		return false
	}
	if this.Rating != that1.Rating {
		return false
	}
	if this.RatingModifier != that1.RatingModifier {
		return false
	}
	if this.TotalNumLeaderSuccess != that1.TotalNumLeaderSuccess {
		return false
	}
	if this.TotalNumLeaderFailure != that1.TotalNumLeaderFailure {
		return false
	}
	if this.TotalNumValidatorSuccess != that1.TotalNumValidatorSuccess {
		return false
	}
	if this.TotalNumValidatorFailure != that1.TotalNumValidatorFailure {
		return false
	}
	if this.TotalNumValidatorIgnoredSignatures != that1.TotalNumValidatorIgnoredSignatures {
		return false
	}
	if this.ShardId != that1.ShardId {
		return false
	}
	if this.ValidatorStatus != that1.ValidatorStatus {
		return false
	}
	return true
}
func (this *ValidatorStatus) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 19)
	s = append(s, "&validator.ValidatorStatus{")
	s = append(s, "TempRating: "+fmt.Sprintf("%#v", this.TempRating)+",\n")
	s = append(s, "NumLeaderSuccess: "+fmt.Sprintf("%#v", this.NumLeaderSuccess)+",\n")
	s = append(s, "NumLeaderFailure: "+fmt.Sprintf("%#v", this.NumLeaderFailure)+",\n")
	s = append(s, "NumValidatorSuccess: "+fmt.Sprintf("%#v", this.NumValidatorSuccess)+",\n")
	s = append(s, "NumValidatorFailure: "+fmt.Sprintf("%#v", this.NumValidatorFailure)+",\n")
	s = append(s, "NumValidatorIgnoredSignatures: "+fmt.Sprintf("%#v", this.NumValidatorIgnoredSignatures)+",\n")
	s = append(s, "Rating: "+fmt.Sprintf("%#v", this.Rating)+",\n")
	s = append(s, "RatingModifier: "+fmt.Sprintf("%#v", this.RatingModifier)+",\n")
	s = append(s, "TotalNumLeaderSuccess: "+fmt.Sprintf("%#v", this.TotalNumLeaderSuccess)+",\n")
	s = append(s, "TotalNumLeaderFailure: "+fmt.Sprintf("%#v", this.TotalNumLeaderFailure)+",\n")
	s = append(s, "TotalNumValidatorSuccess: "+fmt.Sprintf("%#v", this.TotalNumValidatorSuccess)+",\n")
	s = append(s, "TotalNumValidatorFailure: "+fmt.Sprintf("%#v", this.TotalNumValidatorFailure)+",\n")
	s = append(s, "TotalNumValidatorIgnoredSignatures: "+fmt.Sprintf("%#v", this.TotalNumValidatorIgnoredSignatures)+",\n")
	s = append(s, "ShardId: "+fmt.Sprintf("%#v", this.ShardId)+",\n")
	s = append(s, "ValidatorStatus: "+fmt.Sprintf("%#v", this.ValidatorStatus)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringValidatorStatus(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ValidatorStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorStatus) > 0 {
		i -= len(m.ValidatorStatus)
		copy(dAtA[i:], m.ValidatorStatus)
		i = encodeVarintValidatorStatus(dAtA, i, uint64(len(m.ValidatorStatus)))
		i--
		dAtA[i] = 0x7a
	}
	if m.ShardId != 0 {
		i = encodeVarintValidatorStatus(dAtA, i, uint64(m.ShardId))
		i--
		dAtA[i] = 0x70
	}
	if m.TotalNumValidatorIgnoredSignatures != 0 {
		i = encodeVarintValidatorStatus(dAtA, i, uint64(m.TotalNumValidatorIgnoredSignatures))
		i--
		dAtA[i] = 0x68
	}
	if m.TotalNumValidatorFailure != 0 {
		i = encodeVarintValidatorStatus(dAtA, i, uint64(m.TotalNumValidatorFailure))
		i--
		dAtA[i] = 0x60
	}
	if m.TotalNumValidatorSuccess != 0 {
		i = encodeVarintValidatorStatus(dAtA, i, uint64(m.TotalNumValidatorSuccess))
		i--
		dAtA[i] = 0x58
	}
	if m.TotalNumLeaderFailure != 0 {
		i = encodeVarintValidatorStatus(dAtA, i, uint64(m.TotalNumLeaderFailure))
		i--
		dAtA[i] = 0x50
	}
	if m.TotalNumLeaderSuccess != 0 {
		i = encodeVarintValidatorStatus(dAtA, i, uint64(m.TotalNumLeaderSuccess))
		i--
		dAtA[i] = 0x48
	}
	if m.RatingModifier != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.RatingModifier))))
		i--
		dAtA[i] = 0x45
	}
	if m.Rating != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Rating))))
		i--
		dAtA[i] = 0x3d
	}
	if m.NumValidatorIgnoredSignatures != 0 {
		i = encodeVarintValidatorStatus(dAtA, i, uint64(m.NumValidatorIgnoredSignatures))
		i--
		dAtA[i] = 0x30
	}
	if m.NumValidatorFailure != 0 {
		i = encodeVarintValidatorStatus(dAtA, i, uint64(m.NumValidatorFailure))
		i--
		dAtA[i] = 0x28
	}
	if m.NumValidatorSuccess != 0 {
		i = encodeVarintValidatorStatus(dAtA, i, uint64(m.NumValidatorSuccess))
		i--
		dAtA[i] = 0x20
	}
	if m.NumLeaderFailure != 0 {
		i = encodeVarintValidatorStatus(dAtA, i, uint64(m.NumLeaderFailure))
		i--
		dAtA[i] = 0x18
	}
	if m.NumLeaderSuccess != 0 {
		i = encodeVarintValidatorStatus(dAtA, i, uint64(m.NumLeaderSuccess))
		i--
		dAtA[i] = 0x10
	}
	if m.TempRating != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.TempRating))))
		i--
		dAtA[i] = 0xd
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidatorStatus(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidatorStatus(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TempRating != 0 {
		n += 5
	}
	if m.NumLeaderSuccess != 0 {
		n += 1 + sovValidatorStatus(uint64(m.NumLeaderSuccess))
	}
	if m.NumLeaderFailure != 0 {
		n += 1 + sovValidatorStatus(uint64(m.NumLeaderFailure))
	}
	if m.NumValidatorSuccess != 0 {
		n += 1 + sovValidatorStatus(uint64(m.NumValidatorSuccess))
	}
	if m.NumValidatorFailure != 0 {
		n += 1 + sovValidatorStatus(uint64(m.NumValidatorFailure))
	}
	if m.NumValidatorIgnoredSignatures != 0 {
		n += 1 + sovValidatorStatus(uint64(m.NumValidatorIgnoredSignatures))
	}
	if m.Rating != 0 {
		n += 5
	}
	if m.RatingModifier != 0 {
		n += 5
	}
	if m.TotalNumLeaderSuccess != 0 {
		n += 1 + sovValidatorStatus(uint64(m.TotalNumLeaderSuccess))
	}
	if m.TotalNumLeaderFailure != 0 {
		n += 1 + sovValidatorStatus(uint64(m.TotalNumLeaderFailure))
	}
	if m.TotalNumValidatorSuccess != 0 {
		n += 1 + sovValidatorStatus(uint64(m.TotalNumValidatorSuccess))
	}
	if m.TotalNumValidatorFailure != 0 {
		n += 1 + sovValidatorStatus(uint64(m.TotalNumValidatorFailure))
	}
	if m.TotalNumValidatorIgnoredSignatures != 0 {
		n += 1 + sovValidatorStatus(uint64(m.TotalNumValidatorIgnoredSignatures))
	}
	if m.ShardId != 0 {
		n += 1 + sovValidatorStatus(uint64(m.ShardId))
	}
	l = len(m.ValidatorStatus)
	if l > 0 {
		n += 1 + l + sovValidatorStatus(uint64(l))
	}
	return n
}

func sovValidatorStatus(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidatorStatus(x uint64) (n int) {
	return sovValidatorStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ValidatorStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ValidatorStatus{`,
		`TempRating:` + fmt.Sprintf("%v", this.TempRating) + `,`,
		`NumLeaderSuccess:` + fmt.Sprintf("%v", this.NumLeaderSuccess) + `,`,
		`NumLeaderFailure:` + fmt.Sprintf("%v", this.NumLeaderFailure) + `,`,
		`NumValidatorSuccess:` + fmt.Sprintf("%v", this.NumValidatorSuccess) + `,`,
		`NumValidatorFailure:` + fmt.Sprintf("%v", this.NumValidatorFailure) + `,`,
		`NumValidatorIgnoredSignatures:` + fmt.Sprintf("%v", this.NumValidatorIgnoredSignatures) + `,`,
		`Rating:` + fmt.Sprintf("%v", this.Rating) + `,`,
		`RatingModifier:` + fmt.Sprintf("%v", this.RatingModifier) + `,`,
		`TotalNumLeaderSuccess:` + fmt.Sprintf("%v", this.TotalNumLeaderSuccess) + `,`,
		`TotalNumLeaderFailure:` + fmt.Sprintf("%v", this.TotalNumLeaderFailure) + `,`,
		`TotalNumValidatorSuccess:` + fmt.Sprintf("%v", this.TotalNumValidatorSuccess) + `,`,
		`TotalNumValidatorFailure:` + fmt.Sprintf("%v", this.TotalNumValidatorFailure) + `,`,
		`TotalNumValidatorIgnoredSignatures:` + fmt.Sprintf("%v", this.TotalNumValidatorIgnoredSignatures) + `,`,
		`ShardId:` + fmt.Sprintf("%v", this.ShardId) + `,`,
		`ValidatorStatus:` + fmt.Sprintf("%v", this.ValidatorStatus) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringValidatorStatus(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ValidatorStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidatorStatus
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
			return fmt.Errorf("proto: ValidatorStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field TempRating", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.TempRating = float32(math.Float32frombits(v))
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumLeaderSuccess", wireType)
			}
			m.NumLeaderSuccess = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumLeaderSuccess |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumLeaderFailure", wireType)
			}
			m.NumLeaderFailure = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumLeaderFailure |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumValidatorSuccess", wireType)
			}
			m.NumValidatorSuccess = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumValidatorSuccess |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumValidatorFailure", wireType)
			}
			m.NumValidatorFailure = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumValidatorFailure |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumValidatorIgnoredSignatures", wireType)
			}
			m.NumValidatorIgnoredSignatures = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumValidatorIgnoredSignatures |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rating", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Rating = float32(math.Float32frombits(v))
		case 8:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field RatingModifier", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.RatingModifier = float32(math.Float32frombits(v))
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalNumLeaderSuccess", wireType)
			}
			m.TotalNumLeaderSuccess = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalNumLeaderSuccess |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalNumLeaderFailure", wireType)
			}
			m.TotalNumLeaderFailure = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalNumLeaderFailure |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalNumValidatorSuccess", wireType)
			}
			m.TotalNumValidatorSuccess = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalNumValidatorSuccess |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalNumValidatorFailure", wireType)
			}
			m.TotalNumValidatorFailure = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalNumValidatorFailure |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalNumValidatorIgnoredSignatures", wireType)
			}
			m.TotalNumValidatorIgnoredSignatures = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalNumValidatorIgnoredSignatures |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardId", wireType)
			}
			m.ShardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShardId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorStatus
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
				return ErrInvalidLengthValidatorStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidatorStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidatorStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthValidatorStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthValidatorStatus
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
func skipValidatorStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidatorStatus
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
					return 0, ErrIntOverflowValidatorStatus
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
					return 0, ErrIntOverflowValidatorStatus
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
				return 0, ErrInvalidLengthValidatorStatus
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidatorStatus
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidatorStatus
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidatorStatus        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidatorStatus          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidatorStatus = fmt.Errorf("proto: unexpected end of group")
)
