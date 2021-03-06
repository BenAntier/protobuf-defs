// Code generated by protoc-gen-go.
// source: device.proto
// DO NOT EDIT!

package logger

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Type represents the type of the device ID
type Device_Type int32

const (
	// UNKNOWN is the default device ID and should be considered invalid.
	Device_UNKNOWN Device_Type = 0
	// DESKTOP_COOKIE represents a cookie ID that was sent from a desktop
	// browser.
	Device_DESKTOP_COOKIE Device_Type = 100
	// MOBILE_COOKIE represents a cookie ID that was sent from a mobile phone
	// browser.
	Device_MOBILE_COOKIE Device_Type = 101
	// IDFA represents an identifier for advertisers and is available only on
	// iOS devices.
	Device_IDFA Device_Type = 200
	// AAID represents an identifier for advertisers and is available only on
	// Android devices.
	Device_AAID Device_Type = 201
)

var Device_Type_name = map[int32]string{
	0:   "UNKNOWN",
	100: "DESKTOP_COOKIE",
	101: "MOBILE_COOKIE",
	200: "IDFA",
	201: "AAID",
}
var Device_Type_value = map[string]int32{
	"UNKNOWN":        0,
	"DESKTOP_COOKIE": 100,
	"MOBILE_COOKIE":  101,
	"IDFA":           200,
	"AAID":           201,
}

func (x Device_Type) String() string {
	return proto.EnumName(Device_Type_name, int32(x))
}
func (Device_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// Device represents a device ID or a cookie ID depending on where the event
// was sent from.
type Device struct {
	// Type represents the type of the cookie or device id recorded in the id
	// below.
	Type Device_Type `protobuf:"varint,1,opt,name=type,enum=logger.Device_Type" json:"type,omitempty"`
	// id can be our cookie id, IDFA or AAID.
	Id string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*Device)(nil), "logger.Device")
	proto.RegisterEnum("logger.Device_Type", Device_Type_name, Device_Type_value)
}

func init() { proto.RegisterFile("device.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x49, 0x2d, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0xc9, 0x4f, 0x4f, 0x4f, 0x2d,
	0x52, 0x9a, 0xcc, 0xc8, 0xc5, 0xe6, 0x02, 0x96, 0x10, 0x52, 0xe7, 0x62, 0x29, 0xa9, 0x2c, 0x48,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0x12, 0xd6, 0x83, 0xa8, 0xd0, 0x83, 0xc8, 0xea, 0x85,
	0x00, 0xa5, 0x82, 0xc0, 0x0a, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x98, 0x80, 0xca, 0x38,
	0x83, 0x80, 0x2c, 0xa5, 0x00, 0x2e, 0x16, 0x90, 0xac, 0x10, 0x37, 0x17, 0x7b, 0xa8, 0x9f, 0xb7,
	0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0x83, 0x90, 0x10, 0x17, 0x9f, 0x8b, 0x6b, 0xb0, 0x77, 0x88, 0x7f,
	0x40, 0xbc, 0xb3, 0xbf, 0xbf, 0xb7, 0xa7, 0xab, 0x40, 0x8a, 0x90, 0x20, 0x17, 0xaf, 0xaf, 0xbf,
	0x93, 0xa7, 0x8f, 0x2b, 0x4c, 0x28, 0x55, 0x88, 0x93, 0x8b, 0xc5, 0xd3, 0xc5, 0xcd, 0x51, 0xe0,
	0x04, 0x23, 0x88, 0xe9, 0xe8, 0xe8, 0xe9, 0x22, 0x70, 0x92, 0x31, 0x89, 0x0d, 0xec, 0x48, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xf7, 0x8b, 0x62, 0xb4, 0x00, 0x00, 0x00,
}
