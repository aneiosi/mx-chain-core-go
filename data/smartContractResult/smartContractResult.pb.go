// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: smartContractResult.proto

package smartContractResult

import (
	bytes "bytes"
	fmt "fmt"
	github_com_ElrondNetwork_elrond_go_core_vmcommon "github.com/ElrondNetwork/elrond-go/core/vmcommon"
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

type SmartContractResult struct {
	Nonce          uint64                                                    `protobuf:"varint,1,opt,name=Nonce,proto3" json:"nonce"`
	Value          *math_big.Int                                             `protobuf:"bytes,2,opt,name=Value,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"value"`
	RcvAddr        []byte                                                    `protobuf:"bytes,3,opt,name=RcvAddr,proto3" json:"receiver"`
	SndAddr        []byte                                                    `protobuf:"bytes,4,opt,name=SndAddr,proto3" json:"sender"`
	RelayerAddr    []byte                                                    `protobuf:"bytes,5,opt,name=RelayerAddr,proto3" json:"relayer"`
	RelayedValue   *math_big.Int                                             `protobuf:"bytes,6,opt,name=RelayedValue,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"relayedValue"`
	Code           []byte                                                    `protobuf:"bytes,7,opt,name=Code,proto3" json:"code,omitempty"`
	Data           []byte                                                    `protobuf:"bytes,8,opt,name=Data,proto3" json:"data,omitempty"`
	PrevTxHash     []byte                                                    `protobuf:"bytes,9,opt,name=PrevTxHash,proto3" json:"prevTxHash"`
	OriginalTxHash []byte                                                    `protobuf:"bytes,10,opt,name=OriginalTxHash,proto3" json:"originalTxHash"`
	GasLimit       uint64                                                    `protobuf:"varint,11,opt,name=GasLimit,proto3" json:"gasLimit"`
	GasPrice       uint64                                                    `protobuf:"varint,12,opt,name=GasPrice,proto3" json:"gasPrice"`
	CallType       github_com_ElrondNetwork_elrond_go_core_vmcommon.CallType `protobuf:"varint,13,opt,name=CallType,proto3,casttype=github.com/ElrondNetwork/elrond-go/core/vmcommon.CallType" json:"callType"`
	CodeMetadata   []byte                                                    `protobuf:"bytes,14,opt,name=CodeMetadata,proto3" json:"codeMetadata,omitempty"`
	ReturnMessage  []byte                                                    `protobuf:"bytes,15,opt,name=ReturnMessage,proto3" json:"returnMessage,omitempty"`
	OriginalSender []byte                                                    `protobuf:"bytes,16,opt,name=OriginalSender,proto3" json:"originalSender,omitempty"`
}

func (m *SmartContractResult) Reset()      { *m = SmartContractResult{} }
func (*SmartContractResult) ProtoMessage() {}
func (*SmartContractResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc1605de0d3d805, []int{0}
}
func (m *SmartContractResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SmartContractResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SmartContractResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartContractResult.Merge(m, src)
}
func (m *SmartContractResult) XXX_Size() int {
	return m.Size()
}
func (m *SmartContractResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartContractResult.DiscardUnknown(m)
}

var xxx_messageInfo_SmartContractResult proto.InternalMessageInfo

func (m *SmartContractResult) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *SmartContractResult) GetValue() *math_big.Int {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SmartContractResult) GetRcvAddr() []byte {
	if m != nil {
		return m.RcvAddr
	}
	return nil
}

func (m *SmartContractResult) GetSndAddr() []byte {
	if m != nil {
		return m.SndAddr
	}
	return nil
}

func (m *SmartContractResult) GetRelayerAddr() []byte {
	if m != nil {
		return m.RelayerAddr
	}
	return nil
}

func (m *SmartContractResult) GetRelayedValue() *math_big.Int {
	if m != nil {
		return m.RelayedValue
	}
	return nil
}

func (m *SmartContractResult) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *SmartContractResult) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SmartContractResult) GetPrevTxHash() []byte {
	if m != nil {
		return m.PrevTxHash
	}
	return nil
}

func (m *SmartContractResult) GetOriginalTxHash() []byte {
	if m != nil {
		return m.OriginalTxHash
	}
	return nil
}

func (m *SmartContractResult) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *SmartContractResult) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *SmartContractResult) GetCallType() github_com_ElrondNetwork_elrond_go_core_vmcommon.CallType {
	if m != nil {
		return m.CallType
	}
	return 0
}

func (m *SmartContractResult) GetCodeMetadata() []byte {
	if m != nil {
		return m.CodeMetadata
	}
	return nil
}

func (m *SmartContractResult) GetReturnMessage() []byte {
	if m != nil {
		return m.ReturnMessage
	}
	return nil
}

func (m *SmartContractResult) GetOriginalSender() []byte {
	if m != nil {
		return m.OriginalSender
	}
	return nil
}

func init() {
	proto.RegisterType((*SmartContractResult)(nil), "proto.SmartContractResult")
}

func init() { proto.RegisterFile("smartContractResult.proto", fileDescriptor_edc1605de0d3d805) }

var fileDescriptor_edc1605de0d3d805 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x31, 0x6f, 0x13, 0x3f,
	0x18, 0xc6, 0xe3, 0x7f, 0x9b, 0x26, 0x75, 0xd3, 0xfc, 0x91, 0x2b, 0x81, 0x29, 0xc8, 0xae, 0x10,
	0xaa, 0x32, 0xd0, 0x64, 0x60, 0x02, 0x24, 0xa4, 0x26, 0x45, 0x50, 0x89, 0x96, 0xca, 0x2d, 0x0c,
	0x6c, 0xce, 0xdd, 0xcb, 0xf5, 0xc4, 0xdd, 0x39, 0xf2, 0x39, 0x81, 0x6c, 0x7c, 0x04, 0x3e, 0x06,
	0xe2, 0x93, 0x30, 0x76, 0x60, 0xe8, 0x74, 0xd0, 0xeb, 0x82, 0x6e, 0xea, 0xcc, 0x84, 0xce, 0x49,
	0x93, 0xbb, 0xc2, 0xc0, 0xc0, 0x94, 0xbc, 0xcf, 0xf3, 0x7b, 0x9f, 0x37, 0x7e, 0x63, 0x19, 0xdf,
	0x8c, 0x43, 0xa9, 0x4d, 0x4f, 0x45, 0x46, 0x4b, 0xc7, 0x08, 0x88, 0x87, 0x81, 0x69, 0x0f, 0xb4,
	0x32, 0x8a, 0x54, 0xed, 0xc7, 0xfa, 0x96, 0xe7, 0x9b, 0xe3, 0x61, 0xbf, 0xed, 0xa8, 0xb0, 0xe3,
	0x29, 0x4f, 0x75, 0xac, 0xdc, 0x1f, 0xbe, 0xb1, 0x95, 0x2d, 0xec, 0xb7, 0x49, 0xd7, 0x9d, 0xaf,
	0x35, 0xbc, 0x76, 0xf8, 0x7b, 0x26, 0xe1, 0xb8, 0xba, 0xaf, 0x22, 0x07, 0x28, 0xda, 0x40, 0xad,
	0xc5, 0xee, 0x72, 0x96, 0xf0, 0x6a, 0x94, 0x0b, 0x62, 0xa2, 0x13, 0x17, 0x57, 0x5f, 0xc9, 0x60,
	0x08, 0xf4, 0xbf, 0x0d, 0xd4, 0x6a, 0x74, 0xf7, 0x73, 0x60, 0x94, 0x0b, 0x9f, 0xbf, 0xf1, 0xed,
	0x50, 0x9a, 0xe3, 0x4e, 0xdf, 0xf7, 0xda, 0xbb, 0x91, 0x79, 0x54, 0xf8, 0x41, 0x4f, 0x02, 0xad,
	0x22, 0x77, 0x1f, 0xcc, 0x3b, 0xa5, 0xdf, 0x76, 0xc0, 0x56, 0x5b, 0x9e, 0xea, 0xb8, 0xd2, 0xc8,
	0x76, 0xd7, 0xf7, 0x76, 0x23, 0xd3, 0x93, 0xb1, 0x01, 0x2d, 0x26, 0xe1, 0x64, 0x13, 0xd7, 0x84,
	0x33, 0xda, 0x76, 0x5d, 0x4d, 0x17, 0xec, 0x9c, 0x46, 0x96, 0xf0, 0xba, 0x06, 0x07, 0xfc, 0x11,
	0x68, 0x71, 0x69, 0x92, 0xbb, 0xb8, 0x76, 0x18, 0xb9, 0x96, 0x5b, 0xb4, 0x1c, 0xce, 0x12, 0xbe,
	0x14, 0x43, 0xe4, 0xe6, 0xd4, 0xd4, 0x22, 0x5b, 0x78, 0x45, 0x40, 0x20, 0xc7, 0xa0, 0x2d, 0x59,
	0xb5, 0xe4, 0x4a, 0x96, 0xf0, 0x9a, 0x9e, 0xc8, 0xa2, 0xe8, 0x93, 0x31, 0x6e, 0x4c, 0x4a, 0x77,
	0x72, 0xd2, 0x25, 0xcb, 0xbf, 0xcc, 0x12, 0xde, 0xd0, 0x05, 0xfd, 0xdf, 0x1c, 0xb8, 0x34, 0x8a,
	0x6c, 0xe2, 0xc5, 0x9e, 0x72, 0x81, 0xd6, 0xec, 0x48, 0x92, 0x25, 0xbc, 0xe9, 0x28, 0x17, 0xee,
	0xa9, 0xd0, 0x37, 0x10, 0x0e, 0xcc, 0x58, 0x58, 0x3f, 0xe7, 0x76, 0xa4, 0x91, 0xb4, 0x3e, 0xe7,
	0xf2, 0xe8, 0x22, 0x97, 0xfb, 0xa4, 0x8d, 0xf1, 0x81, 0x86, 0xd1, 0xd1, 0xfb, 0x67, 0x32, 0x3e,
	0xa6, 0xcb, 0x96, 0x6e, 0x66, 0x09, 0xc7, 0x83, 0x99, 0x2a, 0x0a, 0x04, 0x79, 0x88, 0x9b, 0x2f,
	0xb4, 0xef, 0xf9, 0x91, 0x0c, 0xa6, 0x3d, 0x78, 0x3e, 0x41, 0x95, 0x1c, 0x71, 0x85, 0x24, 0x2d,
	0x5c, 0x7f, 0x2a, 0xe3, 0xe7, 0x7e, 0xe8, 0x1b, 0xba, 0x62, 0x6f, 0x8f, 0xfd, 0xd3, 0xbc, 0xa9,
	0x26, 0x66, 0xee, 0x94, 0x3c, 0xd0, 0xbe, 0x03, 0xb4, 0x51, 0x22, 0xad, 0x26, 0x66, 0x2e, 0x01,
	0x5c, 0xef, 0xc9, 0x20, 0x38, 0x1a, 0x0f, 0x80, 0xae, 0x6e, 0xa0, 0xd6, 0x42, 0x77, 0x37, 0x27,
	0x9d, 0xa9, 0xf6, 0x33, 0xe1, 0x0f, 0xfe, 0x62, 0xeb, 0x8e, 0xd2, 0xd0, 0x19, 0x85, 0x8e, 0x0a,
	0x43, 0x15, 0xb5, 0x2f, 0x03, 0xc5, 0x2c, 0x9a, 0x3c, 0xc6, 0x8d, 0x7c, 0xad, 0x7b, 0x60, 0x64,
	0xbe, 0x46, 0xda, 0xb4, 0x87, 0x5e, 0xcf, 0x12, 0x7e, 0xdd, 0x29, 0xe8, 0x85, 0xf5, 0x96, 0x78,
	0xb2, 0x8d, 0x57, 0x05, 0x98, 0xa1, 0x8e, 0xf6, 0x20, 0x8e, 0xa5, 0x07, 0xf4, 0x7f, 0x1b, 0x70,
	0x2b, 0x4b, 0xf8, 0x0d, 0x5d, 0x34, 0x0a, 0x09, 0xe5, 0x0e, 0xb2, 0x33, 0xdf, 0xfc, 0xa1, 0xbd,
	0xbe, 0xf4, 0x9a, 0xcd, 0xb8, 0x9d, 0x25, 0x9c, 0xaa, 0x92, 0x53, 0x08, 0xb9, 0xd2, 0xd3, 0xdd,
	0x3b, 0x39, 0x63, 0x95, 0xd3, 0x33, 0x56, 0xb9, 0x38, 0x63, 0xe8, 0x43, 0xca, 0xd0, 0xa7, 0x94,
	0xa1, 0x2f, 0x29, 0x43, 0x27, 0x29, 0x43, 0xa7, 0x29, 0x43, 0xdf, 0x53, 0x86, 0x7e, 0xa4, 0xac,
	0x72, 0x91, 0x32, 0xf4, 0xf1, 0x9c, 0x55, 0x4e, 0xce, 0x59, 0xe5, 0xf4, 0x9c, 0x55, 0x5e, 0xaf,
	0xfd, 0xe1, 0x85, 0xe9, 0x2f, 0xd9, 0xc7, 0xe2, 0xfe, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86,
	0xbf, 0x80, 0xfc, 0x7f, 0x04, 0x00, 0x00,
}

func (this *SmartContractResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SmartContractResult)
	if !ok {
		that2, ok := that.(SmartContractResult)
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
		if !__caster.Equal(this.Value, that1.Value) {
			return false
		}
	}
	if !bytes.Equal(this.RcvAddr, that1.RcvAddr) {
		return false
	}
	if !bytes.Equal(this.SndAddr, that1.SndAddr) {
		return false
	}
	if !bytes.Equal(this.RelayerAddr, that1.RelayerAddr) {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.RelayedValue, that1.RelayedValue) {
			return false
		}
	}
	if !bytes.Equal(this.Code, that1.Code) {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if !bytes.Equal(this.PrevTxHash, that1.PrevTxHash) {
		return false
	}
	if !bytes.Equal(this.OriginalTxHash, that1.OriginalTxHash) {
		return false
	}
	if this.GasLimit != that1.GasLimit {
		return false
	}
	if this.GasPrice != that1.GasPrice {
		return false
	}
	if this.CallType != that1.CallType {
		return false
	}
	if !bytes.Equal(this.CodeMetadata, that1.CodeMetadata) {
		return false
	}
	if !bytes.Equal(this.ReturnMessage, that1.ReturnMessage) {
		return false
	}
	if !bytes.Equal(this.OriginalSender, that1.OriginalSender) {
		return false
	}
	return true
}
func (this *SmartContractResult) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 20)
	s = append(s, "&smartContractResult.SmartContractResult{")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "RcvAddr: "+fmt.Sprintf("%#v", this.RcvAddr)+",\n")
	s = append(s, "SndAddr: "+fmt.Sprintf("%#v", this.SndAddr)+",\n")
	s = append(s, "RelayerAddr: "+fmt.Sprintf("%#v", this.RelayerAddr)+",\n")
	s = append(s, "RelayedValue: "+fmt.Sprintf("%#v", this.RelayedValue)+",\n")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "PrevTxHash: "+fmt.Sprintf("%#v", this.PrevTxHash)+",\n")
	s = append(s, "OriginalTxHash: "+fmt.Sprintf("%#v", this.OriginalTxHash)+",\n")
	s = append(s, "GasLimit: "+fmt.Sprintf("%#v", this.GasLimit)+",\n")
	s = append(s, "GasPrice: "+fmt.Sprintf("%#v", this.GasPrice)+",\n")
	s = append(s, "CallType: "+fmt.Sprintf("%#v", this.CallType)+",\n")
	s = append(s, "CodeMetadata: "+fmt.Sprintf("%#v", this.CodeMetadata)+",\n")
	s = append(s, "ReturnMessage: "+fmt.Sprintf("%#v", this.ReturnMessage)+",\n")
	s = append(s, "OriginalSender: "+fmt.Sprintf("%#v", this.OriginalSender)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSmartContractResult(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SmartContractResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SmartContractResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SmartContractResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OriginalSender) > 0 {
		i -= len(m.OriginalSender)
		copy(dAtA[i:], m.OriginalSender)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.OriginalSender)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.ReturnMessage) > 0 {
		i -= len(m.ReturnMessage)
		copy(dAtA[i:], m.ReturnMessage)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.ReturnMessage)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.CodeMetadata) > 0 {
		i -= len(m.CodeMetadata)
		copy(dAtA[i:], m.CodeMetadata)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.CodeMetadata)))
		i--
		dAtA[i] = 0x72
	}
	if m.CallType != 0 {
		i = encodeVarintSmartContractResult(dAtA, i, uint64(m.CallType))
		i--
		dAtA[i] = 0x68
	}
	if m.GasPrice != 0 {
		i = encodeVarintSmartContractResult(dAtA, i, uint64(m.GasPrice))
		i--
		dAtA[i] = 0x60
	}
	if m.GasLimit != 0 {
		i = encodeVarintSmartContractResult(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x58
	}
	if len(m.OriginalTxHash) > 0 {
		i -= len(m.OriginalTxHash)
		copy(dAtA[i:], m.OriginalTxHash)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.OriginalTxHash)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.PrevTxHash) > 0 {
		i -= len(m.PrevTxHash)
		copy(dAtA[i:], m.PrevTxHash)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.PrevTxHash)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0x3a
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.RelayedValue)
		i -= size
		if _, err := __caster.MarshalTo(m.RelayedValue, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSmartContractResult(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.RelayerAddr) > 0 {
		i -= len(m.RelayerAddr)
		copy(dAtA[i:], m.RelayerAddr)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.RelayerAddr)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SndAddr) > 0 {
		i -= len(m.SndAddr)
		copy(dAtA[i:], m.SndAddr)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.SndAddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RcvAddr) > 0 {
		i -= len(m.RcvAddr)
		copy(dAtA[i:], m.RcvAddr)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.RcvAddr)))
		i--
		dAtA[i] = 0x1a
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.Value)
		i -= size
		if _, err := __caster.MarshalTo(m.Value, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSmartContractResult(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Nonce != 0 {
		i = encodeVarintSmartContractResult(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSmartContractResult(dAtA []byte, offset int, v uint64) int {
	offset -= sovSmartContractResult(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SmartContractResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovSmartContractResult(uint64(m.Nonce))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.Value)
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.RcvAddr)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.SndAddr)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.RelayerAddr)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.RelayedValue)
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.PrevTxHash)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.OriginalTxHash)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	if m.GasLimit != 0 {
		n += 1 + sovSmartContractResult(uint64(m.GasLimit))
	}
	if m.GasPrice != 0 {
		n += 1 + sovSmartContractResult(uint64(m.GasPrice))
	}
	if m.CallType != 0 {
		n += 1 + sovSmartContractResult(uint64(m.CallType))
	}
	l = len(m.CodeMetadata)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.ReturnMessage)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.OriginalSender)
	if l > 0 {
		n += 2 + l + sovSmartContractResult(uint64(l))
	}
	return n
}

func sovSmartContractResult(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSmartContractResult(x uint64) (n int) {
	return sovSmartContractResult(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SmartContractResult) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SmartContractResult{`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`RcvAddr:` + fmt.Sprintf("%v", this.RcvAddr) + `,`,
		`SndAddr:` + fmt.Sprintf("%v", this.SndAddr) + `,`,
		`RelayerAddr:` + fmt.Sprintf("%v", this.RelayerAddr) + `,`,
		`RelayedValue:` + fmt.Sprintf("%v", this.RelayedValue) + `,`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`PrevTxHash:` + fmt.Sprintf("%v", this.PrevTxHash) + `,`,
		`OriginalTxHash:` + fmt.Sprintf("%v", this.OriginalTxHash) + `,`,
		`GasLimit:` + fmt.Sprintf("%v", this.GasLimit) + `,`,
		`GasPrice:` + fmt.Sprintf("%v", this.GasPrice) + `,`,
		`CallType:` + fmt.Sprintf("%v", this.CallType) + `,`,
		`CodeMetadata:` + fmt.Sprintf("%v", this.CodeMetadata) + `,`,
		`ReturnMessage:` + fmt.Sprintf("%v", this.ReturnMessage) + `,`,
		`OriginalSender:` + fmt.Sprintf("%v", this.OriginalSender) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSmartContractResult(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SmartContractResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSmartContractResult
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
			return fmt.Errorf("proto: SmartContractResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SmartContractResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Value = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RcvAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RcvAddr = append(m.RcvAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.RcvAddr == nil {
				m.RcvAddr = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SndAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SndAddr = append(m.SndAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.SndAddr == nil {
				m.SndAddr = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayerAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RelayerAddr = append(m.RelayerAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.RelayerAddr == nil {
				m.RelayerAddr = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayedValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.RelayedValue = tmp
				}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = append(m.Code[:0], dAtA[iNdEx:postIndex]...)
			if m.Code == nil {
				m.Code = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevTxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevTxHash = append(m.PrevTxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PrevTxHash == nil {
				m.PrevTxHash = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalTxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalTxHash = append(m.OriginalTxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.OriginalTxHash == nil {
				m.OriginalTxHash = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			m.GasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallType", wireType)
			}
			m.CallType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CallType |= github_com_ElrondNetwork_elrond_go_core_vmcommon.CallType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeMetadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeMetadata = append(m.CodeMetadata[:0], dAtA[iNdEx:postIndex]...)
			if m.CodeMetadata == nil {
				m.CodeMetadata = []byte{}
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnMessage", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReturnMessage = append(m.ReturnMessage[:0], dAtA[iNdEx:postIndex]...)
			if m.ReturnMessage == nil {
				m.ReturnMessage = []byte{}
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalSender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
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
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalSender = append(m.OriginalSender[:0], dAtA[iNdEx:postIndex]...)
			if m.OriginalSender == nil {
				m.OriginalSender = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSmartContractResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSmartContractResult
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
func skipSmartContractResult(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSmartContractResult
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
					return 0, ErrIntOverflowSmartContractResult
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
					return 0, ErrIntOverflowSmartContractResult
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
				return 0, ErrInvalidLengthSmartContractResult
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSmartContractResult
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSmartContractResult
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSmartContractResult        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSmartContractResult          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSmartContractResult = fmt.Errorf("proto: unexpected end of group")
)
