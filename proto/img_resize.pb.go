// Code generated by protoc-gen-go.
// source: img_resize.proto
// DO NOT EDIT!

/*
Package imgresizer is a generated protocol buffer package.

It is generated from these files:
	img_resize.proto

It has these top-level messages:
	Request
*/
package imgresizer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Request struct {
	Url    string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Width  int32  `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	Height int32  `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Request)(nil), "imgresizer.Request")
}

var fileDescriptor0 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xcc, 0x4d, 0x8f,
	0x2f, 0x4a, 0x2d, 0xce, 0xac, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x8a,
	0x40, 0x04, 0x8a, 0x94, 0x3c, 0xb9, 0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04,
	0xb8, 0x98, 0x4b, 0x8b, 0x72, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x11,
	0x2e, 0xd6, 0xf2, 0xcc, 0x94, 0x92, 0x0c, 0x09, 0x26, 0xa0, 0x18, 0x6b, 0x10, 0x84, 0x23, 0x24,
	0xc6, 0xc5, 0x96, 0x91, 0x9a, 0x99, 0x9e, 0x51, 0x22, 0xc1, 0x0c, 0x16, 0x86, 0xf2, 0x92, 0xd8,
	0xc0, 0xa6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x89, 0x1d, 0x09, 0x71, 0x00, 0x00,
	0x00,
}
