// Code generated by protoc-gen-go.
// source: geoip.proto
// DO NOT EDIT!

package logger

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// GeoIp represents geo-location information about the User.
type GeoIp struct {
	// ip represents the IP of the user in v4 or v6.
	//
	// Types that are valid to be assigned to Ip:
	//	*GeoIp_V4
	//	*GeoIp_V6
	Ip isGeoIp_Ip `protobuf_oneof:"ip"`
	// isp represents the ISP as defined in the message Isp above.
	Isp *GeoIp_Isp `protobuf:"bytes,3,opt,name=isp" json:"isp,omitempty"`
	// country is a two character country code as defined in ISO 3167-1.
	Country string `protobuf:"bytes,4,opt,name=country" json:"country,omitempty"`
	// city is the full city name.
	City string `protobuf:"bytes,5,opt,name=city" json:"city,omitempty"`
	// postalCode is the postal code.
	PostalCode string `protobuf:"bytes,6,opt,name=postal_code,json=postalCode" json:"postal_code,omitempty"`
	// the time_zone represents the TZ as defined on
	// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	TimeZone string `protobuf:"bytes,7,opt,name=time_zone,json=timeZone" json:"time_zone,omitempty"`
	// metro_code represents the Phone Area code.
	MetroCode uint32 `protobuf:"varint,8,opt,name=metro_code,json=metroCode" json:"metro_code,omitempty"`
}

func (m *GeoIp) Reset()                    { *m = GeoIp{} }
func (m *GeoIp) String() string            { return proto.CompactTextString(m) }
func (*GeoIp) ProtoMessage()               {}
func (*GeoIp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type isGeoIp_Ip interface {
	isGeoIp_Ip()
}

type GeoIp_V4 struct {
	V4 string `protobuf:"bytes,1,opt,name=v4,oneof"`
}
type GeoIp_V6 struct {
	V6 string `protobuf:"bytes,2,opt,name=v6,oneof"`
}

func (*GeoIp_V4) isGeoIp_Ip() {}
func (*GeoIp_V6) isGeoIp_Ip() {}

func (m *GeoIp) GetIp() isGeoIp_Ip {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *GeoIp) GetV4() string {
	if x, ok := m.GetIp().(*GeoIp_V4); ok {
		return x.V4
	}
	return ""
}

func (m *GeoIp) GetV6() string {
	if x, ok := m.GetIp().(*GeoIp_V6); ok {
		return x.V6
	}
	return ""
}

func (m *GeoIp) GetIsp() *GeoIp_Isp {
	if m != nil {
		return m.Isp
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GeoIp) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GeoIp_OneofMarshaler, _GeoIp_OneofUnmarshaler, _GeoIp_OneofSizer, []interface{}{
		(*GeoIp_V4)(nil),
		(*GeoIp_V6)(nil),
	}
}

func _GeoIp_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GeoIp)
	// ip
	switch x := m.Ip.(type) {
	case *GeoIp_V4:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.V4)
	case *GeoIp_V6:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.V6)
	case nil:
	default:
		return fmt.Errorf("GeoIp.Ip has unexpected type %T", x)
	}
	return nil
}

func _GeoIp_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GeoIp)
	switch tag {
	case 1: // ip.v4
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Ip = &GeoIp_V4{x}
		return true, err
	case 2: // ip.v6
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Ip = &GeoIp_V6{x}
		return true, err
	default:
		return false, nil
	}
}

func _GeoIp_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GeoIp)
	// ip
	switch x := m.Ip.(type) {
	case *GeoIp_V4:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.V4)))
		n += len(x.V4)
	case *GeoIp_V6:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.V6)))
		n += len(x.V6)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Isp represents the Isp of the user.
type GeoIp_Isp struct {
	// autonomous_system_number or ASN (as defined on Wikipedia) is a
	// collection of connected Internet Protocol (IP) routing prefixes under
	// the control of one or more network operators on behalf of a single
	// administrative entity or domain that presents a common, clearly
	// defined routing policy to the Internet.
	AutonomousSystemNumber uint32 `protobuf:"varint,1,opt,name=autonomous_system_number,json=autonomousSystemNumber" json:"autonomous_system_number,omitempty"`
	// autonomous_system_organization is the organisation (TODO: define this)
	AutonomousSystemOrganization string `protobuf:"bytes,2,opt,name=autonomous_system_organization,json=autonomousSystemOrganization" json:"autonomous_system_organization,omitempty"`
	// name represents the ISP name.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// organization represents the name of the organisation.
	Organization string `protobuf:"bytes,4,opt,name=organization" json:"organization,omitempty"`
}

func (m *GeoIp_Isp) Reset()                    { *m = GeoIp_Isp{} }
func (m *GeoIp_Isp) String() string            { return proto.CompactTextString(m) }
func (*GeoIp_Isp) ProtoMessage()               {}
func (*GeoIp_Isp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

func init() {
	proto.RegisterType((*GeoIp)(nil), "logger.GeoIp")
	proto.RegisterType((*GeoIp_Isp)(nil), "logger.GeoIp.Isp")
}

func init() { proto.RegisterFile("geoip.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x6d, 0xd3, 0xcf, 0xa9, 0x05, 0x9d, 0x83, 0x2c, 0x7e, 0x53, 0x2f, 0x9e, 0x72, 0x50,
	0x29, 0x9e, 0x55, 0xd0, 0x5e, 0x14, 0xe2, 0xcd, 0x4b, 0x48, 0xeb, 0x12, 0x16, 0x9a, 0x9d, 0x65,
	0x77, 0x23, 0xb4, 0x6f, 0xe3, 0x83, 0xf8, 0x6e, 0xee, 0x4e, 0xf0, 0xa3, 0x7a, 0x9b, 0xf9, 0xfd,
	0xf2, 0x1f, 0x32, 0xb3, 0x30, 0x2a, 0x25, 0x29, 0x93, 0x1a, 0x4b, 0x9e, 0xb0, 0xb7, 0xa4, 0xb2,
	0x94, 0x76, 0xf2, 0x9e, 0x40, 0xf7, 0x5e, 0xd2, 0xcc, 0xe0, 0x0e, 0xb4, 0xdf, 0xae, 0x44, 0xeb,
	0xb4, 0x75, 0x3e, 0x7c, 0xd8, 0xca, 0x42, 0xcd, 0x64, 0x2a, 0xda, 0xdf, 0x64, 0x8a, 0x67, 0x90,
	0x28, 0x67, 0x44, 0x12, 0xd0, 0xe8, 0x62, 0x37, 0x6d, 0x66, 0xa4, 0x9c, 0x4f, 0x67, 0xce, 0x64,
	0xd1, 0xa2, 0x80, 0xfe, 0x82, 0x6a, 0xed, 0xed, 0x4a, 0x74, 0x62, 0x36, 0xfb, 0x6a, 0x11, 0xa1,
	0xb3, 0x50, 0x7e, 0x25, 0xba, 0x8c, 0xb9, 0xc6, 0x13, 0x18, 0x19, 0x72, 0xbe, 0x58, 0xe6, 0x0b,
	0x7a, 0x95, 0xa2, 0xc7, 0x0a, 0x1a, 0x74, 0x1b, 0x08, 0x1e, 0xc0, 0xd0, 0xab, 0x4a, 0xe6, 0x6b,
	0xd2, 0x52, 0xf4, 0x59, 0x0f, 0x22, 0x78, 0x09, 0x3d, 0x1e, 0x01, 0x54, 0xd2, 0x5b, 0x6a, 0xc2,
	0x83, 0x60, 0xc7, 0xd9, 0x90, 0x49, 0xcc, 0xee, 0x7f, 0xb4, 0x20, 0x09, 0xff, 0x85, 0xd7, 0x20,
	0x8a, 0xda, 0x93, 0xa6, 0x8a, 0x6a, 0x97, 0xbb, 0x95, 0xf3, 0xb2, 0xca, 0x75, 0x5d, 0xcd, 0xa5,
	0xe5, 0x8d, 0xc7, 0xd9, 0xde, 0x8f, 0x7f, 0x66, 0xfd, 0xc8, 0x16, 0xef, 0xe0, 0xf8, 0x7f, 0x92,
	0x6c, 0x59, 0x68, 0xb5, 0x2e, 0xbc, 0x22, 0xdd, 0xdc, 0x27, 0x3b, 0xfc, 0x9b, 0x7f, 0xfa, 0xf5,
	0x4d, 0x5c, 0x5c, 0x17, 0x95, 0xe4, 0xc3, 0x85, 0xc5, 0x63, 0x8d, 0x13, 0xd8, 0xde, 0x98, 0xd3,
	0xdc, 0x6a, 0x83, 0xdd, 0x74, 0xa0, 0xad, 0xcc, 0xbc, 0xc7, 0x4f, 0x76, 0xf9, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x0d, 0x46, 0x14, 0xaa, 0xc1, 0x01, 0x00, 0x00,
}
