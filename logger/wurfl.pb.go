// Code generated by protoc-gen-go.
// source: wurfl.proto
// DO NOT EDIT!

package logger

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Wurfl struct {
}

func (m *Wurfl) Reset()                    { *m = Wurfl{} }
func (m *Wurfl) String() string            { return proto.CompactTextString(m) }
func (*Wurfl) ProtoMessage()               {}
func (*Wurfl) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func init() {
	proto.RegisterType((*Wurfl)(nil), "logger.Wurfl")
}

func init() { proto.RegisterFile("wurfl.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 57 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2f, 0x2d, 0x4a,
	0xcb, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0xc9, 0x4f, 0x4f, 0x4f, 0x2d, 0x52,
	0x62, 0xe7, 0x62, 0x0d, 0x07, 0x09, 0x27, 0xb1, 0x81, 0xc5, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x89, 0xb7, 0x98, 0x20, 0x26, 0x00, 0x00, 0x00,
}
