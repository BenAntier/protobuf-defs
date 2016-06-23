// Code generated by protoc-gen-go.
// source: publisher.proto
// DO NOT EDIT!

package logger

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Publisher represents the Publisher (or customer) of Publica.
type Publisher struct {
	// uuid represents the id of the publisher as recorded in Publica's database.
	Uuid string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	// script_id represents a custom id that the publisher can set through the
	// javascript bootloader. It allows the publisher to do things like A/B
	// testing, per-page metrics, per-site metrics, etc..
	ScriptId string `protobuf:"bytes,2,opt,name=script_id,json=scriptId" json:"script_id,omitempty"`
	// subscription is an instance of Subscription defined in subscription.proto.
	Subscription *Subscription `protobuf:"bytes,3,opt,name=subscription" json:"subscription,omitempty"`
}

func (m *Publisher) Reset()                    { *m = Publisher{} }
func (m *Publisher) String() string            { return proto.CompactTextString(m) }
func (*Publisher) ProtoMessage()               {}
func (*Publisher) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Publisher) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func init() {
	proto.RegisterType((*Publisher)(nil), "logger.Publisher")
}

func init() { proto.RegisterFile("publisher.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0x4d, 0xca,
	0xc9, 0x2c, 0xce, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0xc9, 0x4f,
	0x4f, 0x4f, 0x2d, 0x92, 0x12, 0x2a, 0x2e, 0x4d, 0x2a, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc,
	0xcf, 0x83, 0xc8, 0x29, 0x95, 0x71, 0x71, 0x06, 0xc0, 0x94, 0x0b, 0x09, 0x71, 0xb1, 0x94, 0x96,
	0x66, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0xd2, 0x5c, 0x9c, 0x10,
	0x3d, 0xf1, 0x40, 0x09, 0x26, 0xb0, 0x04, 0x07, 0x44, 0xc0, 0x33, 0x45, 0xc8, 0x82, 0x8b, 0x07,
	0xd9, 0x4c, 0x09, 0x66, 0xa0, 0x3c, 0xb7, 0x91, 0x88, 0x1e, 0xc4, 0x42, 0xbd, 0x60, 0x24, 0xb9,
	0x20, 0x14, 0x95, 0x49, 0x6c, 0x60, 0xeb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x25,
	0xed, 0x3c, 0xad, 0x00, 0x00, 0x00,
}
