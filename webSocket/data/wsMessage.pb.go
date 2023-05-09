// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wsMessage.proto

package data

import (
	bytes "bytes"
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

// WsMessage contains all the information needed for a WebSocket message
type WsMessage struct {
	WithAcknowledge bool   `protobuf:"varint,1,opt,name=WithAcknowledge,proto3" json:"withAcknowledge,omitempty"`
	Counter         uint64 `protobuf:"varint,2,opt,name=Counter,proto3" json:"counter,omitempty"`
	Type            int32  `protobuf:"varint,3,opt,name=Type,proto3" json:"type,omitempty"`
	Payload         []byte `protobuf:"bytes,4,opt,name=Payload,proto3" json:"payload,omitempty"`
	Topic           string `protobuf:"bytes,5,opt,name=Topic,proto3" json:"topic,omitempty"`
}

func (m *WsMessage) Reset()      { *m = WsMessage{} }
func (*WsMessage) ProtoMessage() {}
func (*WsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e88e8c2dafbb96c, []int{0}
}
func (m *WsMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *WsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WsMessage.Merge(m, src)
}
func (m *WsMessage) XXX_Size() int {
	return m.Size()
}
func (m *WsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_WsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_WsMessage proto.InternalMessageInfo

func (m *WsMessage) GetWithAcknowledge() bool {
	if m != nil {
		return m.WithAcknowledge
	}
	return false
}

func (m *WsMessage) GetCounter() uint64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *WsMessage) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *WsMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *WsMessage) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*WsMessage)(nil), "proto.WsMessage")
}

func init() { proto.RegisterFile("wsMessage.proto", fileDescriptor_5e88e8c2dafbb96c) }

var fileDescriptor_5e88e8c2dafbb96c = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x3d, 0x4f, 0x83, 0x40,
	0x18, 0xc7, 0x79, 0x6a, 0x51, 0x4b, 0x8c, 0x8d, 0x18, 0x13, 0x34, 0xf1, 0x81, 0x38, 0x18, 0x4c,
	0xa4, 0x0c, 0x6e, 0x3a, 0x89, 0x83, 0x93, 0x89, 0xd1, 0x26, 0x4d, 0xdc, 0x80, 0x9e, 0x94, 0xb4,
	0xf4, 0x08, 0x1c, 0x52, 0x36, 0x57, 0x37, 0x3f, 0x86, 0x1f, 0xc5, 0xb1, 0x63, 0x27, 0x62, 0xaf,
	0x8b, 0x61, 0xea, 0x47, 0x30, 0x05, 0x9b, 0xa0, 0xcb, 0xdd, 0x3d, 0xff, 0x97, 0xdf, 0x25, 0x77,
	0x52, 0x3b, 0x8d, 0xef, 0x48, 0x1c, 0xdb, 0x1e, 0xe9, 0x84, 0x11, 0x65, 0x54, 0x16, 0xcb, 0xed,
	0xc8, 0xf0, 0x7c, 0x36, 0x48, 0x9c, 0x8e, 0x4b, 0x03, 0xd3, 0xa3, 0x1e, 0x35, 0x4b, 0xd9, 0x49,
	0x9e, 0xcb, 0xa9, 0x1c, 0xca, 0x53, 0xd5, 0x3a, 0x79, 0x6b, 0x48, 0xad, 0xde, 0x9a, 0x24, 0xdf,
	0x4a, 0xed, 0x9e, 0xcf, 0x06, 0xd7, 0xee, 0x70, 0x4c, 0xd3, 0x11, 0xe9, 0x7b, 0x44, 0x01, 0x0d,
	0xf4, 0x6d, 0xeb, 0xb8, 0xc8, 0xd5, 0xc3, 0xf4, 0xaf, 0x75, 0x4e, 0x03, 0x9f, 0x91, 0x20, 0x64,
	0xd9, 0xc3, 0xff, 0x96, 0x6c, 0x4a, 0x5b, 0x37, 0x34, 0x19, 0x33, 0x12, 0x29, 0x0d, 0x0d, 0xf4,
	0xa6, 0x75, 0x50, 0xe4, 0xea, 0x9e, 0x5b, 0x49, 0xb5, 0xe2, 0x3a, 0x25, 0x9f, 0x4a, 0xcd, 0x6e,
	0x16, 0x12, 0x65, 0x43, 0x03, 0x5d, 0xb4, 0xe4, 0x22, 0x57, 0x77, 0x59, 0x16, 0xd6, 0xef, 0x28,
	0xfd, 0x15, 0xf8, 0xde, 0xce, 0x46, 0xd4, 0xee, 0x2b, 0x4d, 0x0d, 0xf4, 0x9d, 0x0a, 0x1c, 0x56,
	0x52, 0x1d, 0xfc, 0x9b, 0x92, 0xcf, 0x24, 0xb1, 0x4b, 0x43, 0xdf, 0x55, 0x44, 0x0d, 0xf4, 0x96,
	0xb5, 0x5f, 0xe4, 0x6a, 0x9b, 0xad, 0x84, 0x5a, 0xb8, 0x4a, 0x58, 0x6c, 0x3a, 0x47, 0x61, 0x36,
	0x47, 0x61, 0x39, 0x47, 0x78, 0xe5, 0x08, 0x1f, 0x1c, 0xe1, 0x93, 0x23, 0x4c, 0x39, 0xc2, 0x8c,
	0x23, 0x7c, 0x71, 0x84, 0x6f, 0x8e, 0xc2, 0x92, 0x23, 0xbc, 0x2f, 0x50, 0x98, 0x2e, 0x50, 0x98,
	0x2d, 0x50, 0x78, 0xba, 0xac, 0x3d, 0x7a, 0x90, 0x8c, 0x98, 0xff, 0x42, 0xa2, 0x78, 0x62, 0x06,
	0x13, 0xc3, 0x1d, 0xd8, 0xfe, 0xd8, 0x70, 0x69, 0x44, 0x0c, 0x8f, 0x9a, 0x29, 0x71, 0x1e, 0xa9,
	0x3b, 0x24, 0xcc, 0xec, 0xdb, 0xcc, 0xbe, 0x5a, 0x2d, 0xce, 0x66, 0xf9, 0x11, 0x17, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xfb, 0xed, 0x8e, 0x00, 0xd1, 0x01, 0x00, 0x00,
}

func (this *WsMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WsMessage)
	if !ok {
		that2, ok := that.(WsMessage)
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
	if this.WithAcknowledge != that1.WithAcknowledge {
		return false
	}
	if this.Counter != that1.Counter {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if !bytes.Equal(this.Payload, that1.Payload) {
		return false
	}
	if this.Topic != that1.Topic {
		return false
	}
	return true
}
func (this *WsMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&data.WsMessage{")
	s = append(s, "WithAcknowledge: "+fmt.Sprintf("%#v", this.WithAcknowledge)+",\n")
	s = append(s, "Counter: "+fmt.Sprintf("%#v", this.Counter)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "Topic: "+fmt.Sprintf("%#v", this.Topic)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringWsMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *WsMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WsMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WsMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Topic) > 0 {
		i -= len(m.Topic)
		copy(dAtA[i:], m.Topic)
		i = encodeVarintWsMessage(dAtA, i, uint64(len(m.Topic)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintWsMessage(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if m.Type != 0 {
		i = encodeVarintWsMessage(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x18
	}
	if m.Counter != 0 {
		i = encodeVarintWsMessage(dAtA, i, uint64(m.Counter))
		i--
		dAtA[i] = 0x10
	}
	if m.WithAcknowledge {
		i--
		if m.WithAcknowledge {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintWsMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovWsMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WsMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.WithAcknowledge {
		n += 2
	}
	if m.Counter != 0 {
		n += 1 + sovWsMessage(uint64(m.Counter))
	}
	if m.Type != 0 {
		n += 1 + sovWsMessage(uint64(m.Type))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovWsMessage(uint64(l))
	}
	l = len(m.Topic)
	if l > 0 {
		n += 1 + l + sovWsMessage(uint64(l))
	}
	return n
}

func sovWsMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWsMessage(x uint64) (n int) {
	return sovWsMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *WsMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WsMessage{`,
		`WithAcknowledge:` + fmt.Sprintf("%v", this.WithAcknowledge) + `,`,
		`Counter:` + fmt.Sprintf("%v", this.Counter) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`Topic:` + fmt.Sprintf("%v", this.Topic) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringWsMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *WsMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWsMessage
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
			return fmt.Errorf("proto: WsMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WsMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithAcknowledge", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWsMessage
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
			m.WithAcknowledge = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWsMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWsMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWsMessage
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
				return ErrInvalidLengthWsMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWsMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWsMessage
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
				return ErrInvalidLengthWsMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWsMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWsMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWsMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWsMessage
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
func skipWsMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWsMessage
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
					return 0, ErrIntOverflowWsMessage
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
					return 0, ErrIntOverflowWsMessage
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
				return 0, ErrInvalidLengthWsMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWsMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWsMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWsMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWsMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWsMessage = fmt.Errorf("proto: unexpected end of group")
)
