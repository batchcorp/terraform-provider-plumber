// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_args_nats_jetstream.proto

package args

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

type NatsJetstreamTLSOptions struct {
	// @gotags: kong:"help='CA file (only needed if addr is tls://*)',env='PLUMBER_RELAY_NATS_JETSTREAM_TLS_CA_CERT'"
	TlsCaCert []byte `protobuf:"bytes,1,opt,name=tls_ca_cert,json=tlsCaCert,proto3" json:"tls_ca_cert,omitempty" kong:"help='CA file (only needed if addr is tls://*)',env='PLUMBER_RELAY_NATS_JETSTREAM_TLS_CA_CERT'"`
	// @gotags: kong:"help='Client cert file (only needed if addr is tls://*)',env='PLUMBER_RELAY_NATS_JETSTREAM_TLS_CLIENT_CERT'"
	TlsClientCert []byte `protobuf:"bytes,2,opt,name=tls_client_cert,json=tlsClientCert,proto3" json:"tls_client_cert,omitempty" kong:"help='Client cert file (only needed if addr is tls://*)',env='PLUMBER_RELAY_NATS_JETSTREAM_TLS_CLIENT_CERT'"`
	// @gotags: kong:"help='Client key file (only needed if addr is tls://*)',env='PLUMBER_RELAY_NATS_JETSTREAM_TLS_CLIENT_KEY'"
	TlsClientKey []byte `protobuf:"bytes,3,opt,name=tls_client_key,json=tlsClientKey,proto3" json:"tls_client_key,omitempty" kong:"help='Client key file (only needed if addr is tls://*)',env='PLUMBER_RELAY_NATS_JETSTREAM_TLS_CLIENT_KEY'"`
	// @gotags: kong:"help='Whether to verify server certificate',env='PLUMBER_RELAY_NATS_JETSTREAM_SKIP_VERIFY_TLS'"
	TlsSkipVerify        bool     `protobuf:"varint,4,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty" kong:"help='Whether to verify server certificate',env='PLUMBER_RELAY_NATS_JETSTREAM_SKIP_VERIFY_TLS'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatsJetstreamTLSOptions) Reset()         { *m = NatsJetstreamTLSOptions{} }
func (m *NatsJetstreamTLSOptions) String() string { return proto.CompactTextString(m) }
func (*NatsJetstreamTLSOptions) ProtoMessage()    {}
func (*NatsJetstreamTLSOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1868c7e9e1b078c, []int{0}
}

func (m *NatsJetstreamTLSOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatsJetstreamTLSOptions.Unmarshal(m, b)
}
func (m *NatsJetstreamTLSOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatsJetstreamTLSOptions.Marshal(b, m, deterministic)
}
func (m *NatsJetstreamTLSOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatsJetstreamTLSOptions.Merge(m, src)
}
func (m *NatsJetstreamTLSOptions) XXX_Size() int {
	return xxx_messageInfo_NatsJetstreamTLSOptions.Size(m)
}
func (m *NatsJetstreamTLSOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_NatsJetstreamTLSOptions.DiscardUnknown(m)
}

var xxx_messageInfo_NatsJetstreamTLSOptions proto.InternalMessageInfo

func (m *NatsJetstreamTLSOptions) GetTlsCaCert() []byte {
	if m != nil {
		return m.TlsCaCert
	}
	return nil
}

func (m *NatsJetstreamTLSOptions) GetTlsClientCert() []byte {
	if m != nil {
		return m.TlsClientCert
	}
	return nil
}

func (m *NatsJetstreamTLSOptions) GetTlsClientKey() []byte {
	if m != nil {
		return m.TlsClientKey
	}
	return nil
}

func (m *NatsJetstreamTLSOptions) GetTlsSkipVerify() bool {
	if m != nil {
		return m.TlsSkipVerify
	}
	return false
}

type NatsJetstreamConn struct {
	// @gotags: kong:"help='Dial string for NATS server. Ex: nats://localhost:4222',default='nats://localhost:4222',env='PLUMBER_RELAY_NATS_JETSTREAM_DSN'"
	Dsn string `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty" kong:"help='Dial string for NATS server. Ex: nats://localhost:4222',default='nats://localhost:4222',env='PLUMBER_RELAY_NATS_JETSTREAM_DSN'"`
	// @gotags: kong:"help='NATS .creds file containing authentication credentials',env='PLUMBER_RELAY_NATS_JETSTREAM_CREDENTIALS'"
	UserCredentials []byte `protobuf:"bytes,2,opt,name=user_credentials,json=userCredentials,proto3" json:"user_credentials,omitempty" kong:"help='NATS .creds file containing authentication credentials',env='PLUMBER_RELAY_NATS_JETSTREAM_CREDENTIALS'"`
	// @gotags: kong:"help='User specified client ID to connect with',default=plumber,env='PLUMBER_RELAY_NATS_JETSTREAM_CLIENT_ID'"
	ClientId string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" kong:"help='User specified client ID to connect with',default=plumber,env='PLUMBER_RELAY_NATS_JETSTREAM_CLIENT_ID'"`
	// @gotags: kong:"embed"
	TlsOptions           *NatsJetstreamTLSOptions `protobuf:"bytes,4,opt,name=tls_options,json=tlsOptions,proto3" json:"tls_options,omitempty" kong:"embed"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *NatsJetstreamConn) Reset()         { *m = NatsJetstreamConn{} }
func (m *NatsJetstreamConn) String() string { return proto.CompactTextString(m) }
func (*NatsJetstreamConn) ProtoMessage()    {}
func (*NatsJetstreamConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1868c7e9e1b078c, []int{1}
}

func (m *NatsJetstreamConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatsJetstreamConn.Unmarshal(m, b)
}
func (m *NatsJetstreamConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatsJetstreamConn.Marshal(b, m, deterministic)
}
func (m *NatsJetstreamConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatsJetstreamConn.Merge(m, src)
}
func (m *NatsJetstreamConn) XXX_Size() int {
	return xxx_messageInfo_NatsJetstreamConn.Size(m)
}
func (m *NatsJetstreamConn) XXX_DiscardUnknown() {
	xxx_messageInfo_NatsJetstreamConn.DiscardUnknown(m)
}

var xxx_messageInfo_NatsJetstreamConn proto.InternalMessageInfo

func (m *NatsJetstreamConn) GetDsn() string {
	if m != nil {
		return m.Dsn
	}
	return ""
}

func (m *NatsJetstreamConn) GetUserCredentials() []byte {
	if m != nil {
		return m.UserCredentials
	}
	return nil
}

func (m *NatsJetstreamConn) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *NatsJetstreamConn) GetTlsOptions() *NatsJetstreamTLSOptions {
	if m != nil {
		return m.TlsOptions
	}
	return nil
}

type NatsJetstreamReadArgs struct {
	// @gotags: kong:"help='NATS JetStream stream name. Ex: orders.>',env='PLUMBER_RELAY_NATS_JETSTREAM_STREAM'"
	Stream               string   `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty" kong:"help='NATS JetStream stream name. Ex: orders.>',env='PLUMBER_RELAY_NATS_JETSTREAM_STREAM'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatsJetstreamReadArgs) Reset()         { *m = NatsJetstreamReadArgs{} }
func (m *NatsJetstreamReadArgs) String() string { return proto.CompactTextString(m) }
func (*NatsJetstreamReadArgs) ProtoMessage()    {}
func (*NatsJetstreamReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1868c7e9e1b078c, []int{2}
}

func (m *NatsJetstreamReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatsJetstreamReadArgs.Unmarshal(m, b)
}
func (m *NatsJetstreamReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatsJetstreamReadArgs.Marshal(b, m, deterministic)
}
func (m *NatsJetstreamReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatsJetstreamReadArgs.Merge(m, src)
}
func (m *NatsJetstreamReadArgs) XXX_Size() int {
	return xxx_messageInfo_NatsJetstreamReadArgs.Size(m)
}
func (m *NatsJetstreamReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_NatsJetstreamReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_NatsJetstreamReadArgs proto.InternalMessageInfo

func (m *NatsJetstreamReadArgs) GetStream() string {
	if m != nil {
		return m.Stream
	}
	return ""
}

type NatsJetstreamWriteArgs struct {
	// @gotags: kong:"help='NATS JetStream stream name. Ex: orders.>'"
	Stream               string   `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty" kong:"help='NATS JetStream stream name. Ex: orders.>'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatsJetstreamWriteArgs) Reset()         { *m = NatsJetstreamWriteArgs{} }
func (m *NatsJetstreamWriteArgs) String() string { return proto.CompactTextString(m) }
func (*NatsJetstreamWriteArgs) ProtoMessage()    {}
func (*NatsJetstreamWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1868c7e9e1b078c, []int{3}
}

func (m *NatsJetstreamWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatsJetstreamWriteArgs.Unmarshal(m, b)
}
func (m *NatsJetstreamWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatsJetstreamWriteArgs.Marshal(b, m, deterministic)
}
func (m *NatsJetstreamWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatsJetstreamWriteArgs.Merge(m, src)
}
func (m *NatsJetstreamWriteArgs) XXX_Size() int {
	return xxx_messageInfo_NatsJetstreamWriteArgs.Size(m)
}
func (m *NatsJetstreamWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_NatsJetstreamWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_NatsJetstreamWriteArgs proto.InternalMessageInfo

func (m *NatsJetstreamWriteArgs) GetStream() string {
	if m != nil {
		return m.Stream
	}
	return ""
}

func init() {
	proto.RegisterType((*NatsJetstreamTLSOptions)(nil), "protos.args.NatsJetstreamTLSOptions")
	proto.RegisterType((*NatsJetstreamConn)(nil), "protos.args.NatsJetstreamConn")
	proto.RegisterType((*NatsJetstreamReadArgs)(nil), "protos.args.NatsJetstreamReadArgs")
	proto.RegisterType((*NatsJetstreamWriteArgs)(nil), "protos.args.NatsJetstreamWriteArgs")
}

func init() { proto.RegisterFile("ps_args_nats_jetstream.proto", fileDescriptor_a1868c7e9e1b078c) }

var fileDescriptor_a1868c7e9e1b078c = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0x51, 0x5d, 0x8c, 0xb5, 0x76, 0x6b, 0x77, 0xa1, 0xae, 0xa1, 0xa5, 0x18, 0x61, 0x8a,
	0x7b, 0xa8, 0x54, 0x92, 0x53, 0xc8, 0x29, 0x11, 0x39, 0xe4, 0x0f, 0x09, 0xc8, 0x21, 0x81, 0x5c,
	0xc4, 0x4a, 0x9a, 0xc8, 0x1b, 0x4b, 0x5a, 0xb1, 0x33, 0x0a, 0xf8, 0x53, 0xe5, 0x92, 0x0f, 0x18,
	0xb4, 0x92, 0x13, 0xfb, 0x90, 0x9c, 0x76, 0xe6, 0xed, 0x6f, 0x86, 0x37, 0x8f, 0xfd, 0x2a, 0x31,
	0x14, 0x3a, 0xc5, 0xb0, 0x10, 0x84, 0xe1, 0x03, 0x10, 0x92, 0x06, 0x91, 0xbb, 0xa5, 0x56, 0xa4,
	0x78, 0xdf, 0x3c, 0xe8, 0xd6, 0x84, 0xf3, 0x64, 0xb1, 0x1f, 0x97, 0x82, 0xf0, 0x6c, 0x03, 0x5d,
	0x5f, 0x2c, 0xae, 0x4a, 0x92, 0xaa, 0x40, 0xfe, 0x9b, 0xf5, 0x29, 0xc3, 0x30, 0x16, 0x61, 0x0c,
	0x9a, 0x26, 0xd6, 0xd4, 0x9a, 0x0f, 0x02, 0x9b, 0x32, 0xf4, 0x85, 0x0f, 0x9a, 0xf8, 0x1f, 0x36,
	0x34, 0xff, 0x99, 0x84, 0x82, 0x1a, 0xe6, 0x93, 0x61, 0xbe, 0xd4, 0x8c, 0x51, 0x0d, 0x37, 0x63,
	0x5f, 0xb7, 0xb8, 0x15, 0xac, 0x27, 0x1d, 0x83, 0x0d, 0x5e, 0xb1, 0x73, 0x58, 0x6f, 0xb6, 0xe1,
	0x4a, 0x96, 0xe1, 0x23, 0x68, 0x79, 0xbf, 0x9e, 0x7c, 0x9e, 0x5a, 0xf3, 0x9e, 0xd9, 0xb6, 0x58,
	0xc9, 0xf2, 0xc6, 0x88, 0xce, 0xb3, 0xc5, 0xbe, 0xed, 0x38, 0xf6, 0x55, 0x51, 0xf0, 0x11, 0xeb,
	0x24, 0x58, 0x18, 0x8f, 0x76, 0x50, 0x97, 0xfc, 0x2f, 0x1b, 0x55, 0x08, 0x3a, 0x8c, 0x35, 0x24,
	0x50, 0x90, 0x14, 0x19, 0xb6, 0xf6, 0x86, 0xb5, 0xee, 0xbf, 0xc9, 0xfc, 0x27, 0xb3, 0x5b, 0x73,
	0x32, 0x31, 0xde, 0xec, 0xa0, 0xd7, 0x08, 0xa7, 0x09, 0x3f, 0x69, 0x52, 0x50, 0x4d, 0x28, 0xc6,
	0x53, 0x7f, 0x6f, 0xe6, 0x6e, 0x85, 0xe8, 0xbe, 0x13, 0x60, 0xc0, 0x28, 0xc3, 0xb6, 0x76, 0x3c,
	0xf6, 0x7d, 0x07, 0x0b, 0x40, 0x24, 0x47, 0x3a, 0x45, 0x3e, 0x66, 0xdd, 0x46, 0x69, 0xcd, 0xb7,
	0x9d, 0xf3, 0x9f, 0x8d, 0x77, 0x06, 0x6e, 0xb5, 0x24, 0xf8, 0x68, 0xe2, 0xf8, 0xf0, 0xee, 0x20,
	0x95, 0xb4, 0xac, 0x22, 0x37, 0x56, 0xb9, 0x17, 0x09, 0x8a, 0x97, 0xb1, 0xd2, 0xa5, 0x57, 0x66,
	0x55, 0x1e, 0x81, 0xfe, 0x87, 0xf1, 0x12, 0x72, 0x81, 0x5e, 0x54, 0xc9, 0x2c, 0xf1, 0x52, 0xe5,
	0x35, 0x37, 0x78, 0xf5, 0x0d, 0x51, 0xd7, 0x34, 0xfb, 0x2f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7e,
	0x6a, 0x97, 0x13, 0x3c, 0x02, 0x00, 0x00,
}
