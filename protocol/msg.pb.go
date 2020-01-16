// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package protocol

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

type Request struct {
	ServicePath          []byte            `protobuf:"bytes,2,opt,name=service_path,json=servicePath,proto3" json:"service_path,omitempty"`
	Metadata             map[string][]byte `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Payload              []byte            `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetServicePath() []byte {
	if m != nil {
		return m.ServicePath
	}
	return nil
}

func (m *Request) GetMetadata() map[string][]byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Request) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Response struct {
	RetCode              uint32            `protobuf:"varint,1,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
	RetMsg               []byte            `protobuf:"bytes,2,opt,name=ret_msg,json=retMsg,proto3" json:"ret_msg,omitempty"`
	Metadata             map[string][]byte `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Payload              []byte            `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetRetCode() uint32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *Response) GetRetMsg() []byte {
	if m != nil {
		return m.RetMsg
	}
	return nil
}

func (m *Response) GetMetadata() map[string][]byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Response) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "protocol.Request")
	proto.RegisterMapType((map[string][]byte)(nil), "protocol.Request.MetadataEntry")
	proto.RegisterType((*Response)(nil), "protocol.Response")
	proto.RegisterMapType((map[string][]byte)(nil), "protocol.Response.MetadataEntry")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x14, 0x84, 0xd9, 0x56, 0x9b, 0xf4, 0xb5, 0x05, 0x59, 0x04, 0xa3, 0x17, 0x63, 0x45, 0xc8, 0x29,
	0x01, 0xbd, 0x88, 0xf5, 0xa4, 0x78, 0x2c, 0xc8, 0x1e, 0xbd, 0x84, 0x4d, 0xf2, 0x48, 0xaa, 0x49,
	0x76, 0xdd, 0x7d, 0x29, 0xe4, 0xdf, 0xf9, 0x1f, 0xfc, 0x43, 0xd2, 0x34, 0x11, 0xc5, 0xbb, 0xa7,
	0x9d, 0x59, 0xde, 0x7c, 0xcc, 0xc0, 0xb4, 0xb2, 0x79, 0xa8, 0x8d, 0x22, 0xc5, 0xdd, 0xee, 0x49,
	0x55, 0xb9, 0xfc, 0x60, 0xe0, 0x08, 0x7c, 0x6f, 0xd0, 0x12, 0xbf, 0x80, 0xb9, 0x45, 0xb3, 0xdd,
	0xa4, 0x18, 0x6b, 0x49, 0x85, 0x37, 0xf2, 0x59, 0x30, 0x17, 0xb3, 0xfe, 0xef, 0x59, 0x52, 0xc1,
	0x57, 0xe0, 0x56, 0x48, 0x32, 0x93, 0x24, 0xbd, 0xb1, 0x3f, 0x0e, 0x66, 0xd7, 0xe7, 0xe1, 0xc0,
	0x0a, 0x7b, 0x4e, 0xb8, 0xee, 0x2f, 0x9e, 0x6a, 0x32, 0xad, 0xf8, 0x0e, 0x70, 0x0f, 0x1c, 0x2d,
	0xdb, 0x52, 0xc9, 0xcc, 0x3b, 0xe8, 0xd0, 0x83, 0x3d, 0x5b, 0xc1, 0xe2, 0x57, 0x88, 0x1f, 0xc1,
	0xf8, 0x0d, 0x5b, 0x8f, 0xf9, 0x2c, 0x98, 0x8a, 0x9d, 0xe4, 0xc7, 0x70, 0xb8, 0x95, 0x65, 0x83,
	0x7d, 0xab, 0xbd, 0xb9, 0x1b, 0xdd, 0xb2, 0xe5, 0x27, 0x03, 0x57, 0xa0, 0xd5, 0xaa, 0xb6, 0xc8,
	0x4f, 0xc1, 0x35, 0x48, 0x71, 0xaa, 0x32, 0xec, 0xd2, 0x0b, 0xe1, 0x18, 0xa4, 0x47, 0x95, 0x21,
	0x3f, 0x81, 0x9d, 0x8c, 0x2b, 0x9b, 0xf7, 0x8c, 0x89, 0x41, 0x5a, 0xdb, 0x9c, 0xdf, 0xff, 0x19,
	0xe5, 0xff, 0x1c, 0xb5, 0x27, 0xff, 0xf3, 0xaa, 0x87, 0xab, 0x97, 0xcb, 0x7c, 0x43, 0x45, 0x93,
	0x84, 0xa9, 0xaa, 0xa2, 0xb2, 0x49, 0x64, 0xad, 0x8d, 0x7a, 0x8d, 0x72, 0x65, 0x74, 0x1a, 0x0d,
	0xf5, 0x92, 0x49, 0xa7, 0x6e, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x5a, 0x2e, 0x6a, 0xdd,
	0x01, 0x00, 0x00,
}
