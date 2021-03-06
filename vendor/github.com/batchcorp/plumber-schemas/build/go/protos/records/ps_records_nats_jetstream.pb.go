// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/ps_records_nats_jetstream.proto

package records

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NatsJetstream struct {
	Stream               string   `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatsJetstream) Reset()         { *m = NatsJetstream{} }
func (m *NatsJetstream) String() string { return proto.CompactTextString(m) }
func (*NatsJetstream) ProtoMessage()    {}
func (*NatsJetstream) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb39dca56fdde41, []int{0}
}

func (m *NatsJetstream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatsJetstream.Unmarshal(m, b)
}
func (m *NatsJetstream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatsJetstream.Marshal(b, m, deterministic)
}
func (m *NatsJetstream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatsJetstream.Merge(m, src)
}
func (m *NatsJetstream) XXX_Size() int {
	return xxx_messageInfo_NatsJetstream.Size(m)
}
func (m *NatsJetstream) XXX_DiscardUnknown() {
	xxx_messageInfo_NatsJetstream.DiscardUnknown(m)
}

var xxx_messageInfo_NatsJetstream proto.InternalMessageInfo

func (m *NatsJetstream) GetStream() string {
	if m != nil {
		return m.Stream
	}
	return ""
}

func (m *NatsJetstream) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*NatsJetstream)(nil), "protos.records.NatsJetstream")
}

func init() {
	proto.RegisterFile("records/ps_records_nats_jetstream.proto", fileDescriptor_2fb39dca56fdde41)
}

var fileDescriptor_2fb39dca56fdde41 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8d, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x46, 0xa9, 0x60, 0xc1, 0xa0, 0x0e, 0x45, 0xa4, 0x63, 0x71, 0xb1, 0x8b, 0xbd, 0xc1, 0x55,
	0x1d, 0x1c, 0x1d, 0x1c, 0x3a, 0xba, 0x94, 0x24, 0x0d, 0x6d, 0xa5, 0x31, 0xe1, 0xee, 0xe2, 0xef,
	0x17, 0xda, 0x38, 0x7d, 0xdf, 0x83, 0x07, 0x4f, 0x1c, 0xd1, 0x68, 0x87, 0x2d, 0x81, 0xa7, 0x26,
	0xde, 0xe6, 0x23, 0x99, 0x9a, 0xb7, 0x61, 0x62, 0x34, 0xd2, 0x56, 0x1e, 0x1d, 0xbb, 0x6c, 0x3b,
	0x0d, 0x55, 0x51, 0x3a, 0x5c, 0xc5, 0xe6, 0x29, 0x99, 0x1e, 0x7f, 0x2d, 0xdb, 0x8b, 0x74, 0x7e,
	0x79, 0x52, 0x24, 0xe5, 0xaa, 0x8e, 0x94, 0xed, 0xc4, 0xf2, 0x2b, 0xc7, 0x60, 0xf2, 0x45, 0x91,
	0x94, 0xeb, 0x7a, 0x86, 0xfb, 0xed, 0x75, 0xe9, 0x06, 0xee, 0x83, 0xaa, 0xb4, 0xb3, 0xa0, 0x24,
	0xeb, 0x5e, 0x3b, 0xf4, 0xe0, 0xc7, 0x60, 0x95, 0xc1, 0x13, 0xe9, 0xde, 0x58, 0x49, 0xa0, 0xc2,
	0x30, 0xb6, 0xd0, 0x39, 0x98, 0xf3, 0x10, 0xf3, 0x2a, 0x9d, 0xf8, 0xfc, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0xcf, 0x9a, 0xc6, 0x40, 0xc0, 0x00, 0x00, 0x00,
}
