// Code generated by protoc-gen-go.
// source: page.proto
// DO NOT EDIT!

package logger

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Scheme defines a struct for the scheme.
type Page_Scheme int32

const (
	Page_UNKNOWN Page_Scheme = 0
	Page_HTTP    Page_Scheme = 1
	Page_HTTPS   Page_Scheme = 2
)

var Page_Scheme_name = map[int32]string{
	0: "UNKNOWN",
	1: "HTTP",
	2: "HTTPS",
}
var Page_Scheme_value = map[string]int32{
	"UNKNOWN": 0,
	"HTTP":    1,
	"HTTPS":   2,
}

func (x Page_Scheme) String() string {
	return proto.EnumName(Page_Scheme_name, int32(x))
}
func (Page_Scheme) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0, 0} }

// Page defines a struct to collect all the information about a Page like the
// URL, the path, the referral etc...
type Page struct {
	// url is the full URL including the scheme, the domain, the path and the
	// params.
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	// scheme is either http or https and is defined in the Scheme above.
	Scheme Page_Scheme `protobuf:"varint,2,opt,name=scheme,enum=logger.Page_Scheme" json:"scheme,omitempty"`
	// domain is the domain (including any sub-domain) for the page.
	Domain string `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	// path is the path of the page (without the domain or the query params).
	Path string `protobuf:"bytes,4,opt,name=path" json:"path,omitempty"`
	// params is a map<string, string> that represents the query params.
	Params map[string]string `protobuf:"bytes,5,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// referrer is the referrer for the page.
	Referrer string `protobuf:"bytes,6,opt,name=referrer" json:"referrer,omitempty"`
	// load_time is the duration calculated between 'navigationStart' and
	// 'loadEventEnd' available on the 'window.performance.timing' browser
	// JavaScript object.
	LoadTime *google_protobuf1.Duration `protobuf:"bytes,7,opt,name=load_time,json=loadTime" json:"load_time,omitempty"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Page) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Page) GetLoadTime() *google_protobuf1.Duration {
	if m != nil {
		return m.LoadTime
	}
	return nil
}

func init() {
	proto.RegisterType((*Page)(nil), "logger.Page")
	proto.RegisterEnum("logger.Page_Scheme", Page_Scheme_name, Page_Scheme_value)
}

func init() { proto.RegisterFile("page.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8f, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xff, 0xed, 0xba, 0x6c, 0xbd, 0x85, 0x3f, 0xe5, 0x2a, 0x12, 0xf7, 0x20, 0x63, 0x4f,
	0x43, 0x21, 0x93, 0x09, 0xa2, 0x3e, 0x2b, 0x08, 0xc2, 0x1c, 0x5d, 0xc5, 0x47, 0xc9, 0x6c, 0xd6,
	0x15, 0xdb, 0xa6, 0x64, 0xad, 0xb0, 0x2f, 0xeb, 0x67, 0x31, 0x4d, 0x3a, 0xd1, 0xb7, 0x73, 0xef,
	0x39, 0x39, 0xb9, 0x3f, 0x80, 0x8a, 0xa7, 0x82, 0x55, 0x4a, 0xd6, 0x12, 0x49, 0x2e, 0xd3, 0x54,
	0xa8, 0xd1, 0x59, 0x2a, 0x65, 0x9a, 0x8b, 0x99, 0xd9, 0xae, 0x9b, 0xcd, 0x2c, 0x69, 0x14, 0xaf,
	0x33, 0x59, 0xda, 0xdc, 0xe4, 0xcb, 0x05, 0x6f, 0xa9, 0x9f, 0x61, 0x08, 0xbd, 0x46, 0xe5, 0xd4,
	0x19, 0x3b, 0x53, 0x3f, 0x6a, 0x25, 0x5e, 0x00, 0xd9, 0xbd, 0x6f, 0x45, 0x21, 0xa8, 0xab, 0x97,
	0xff, 0xe7, 0x47, 0xcc, 0x76, 0xb2, 0x36, 0xcf, 0x56, 0xc6, 0x8a, 0xba, 0x08, 0x9e, 0x00, 0x49,
	0x64, 0xc1, 0xb3, 0x92, 0xf6, 0x4c, 0x43, 0x37, 0x21, 0x82, 0x57, 0xf1, 0x7a, 0x4b, 0x3d, 0xb3,
	0x35, 0x1a, 0x2f, 0x81, 0x54, 0x5c, 0xf1, 0x62, 0x47, 0xfb, 0xe3, 0xde, 0x34, 0x98, 0xd3, 0x3f,
	0xc5, 0x4b, 0x63, 0x3d, 0x94, 0xb5, 0xda, 0x47, 0x5d, 0x0e, 0x47, 0x30, 0x54, 0x62, 0x23, 0x94,
	0x12, 0x8a, 0x12, 0xd3, 0xf4, 0x33, 0xe3, 0x35, 0xf8, 0xb9, 0xe4, 0xc9, 0x5b, 0x9d, 0xe9, 0x4b,
	0x07, 0xda, 0x0c, 0xe6, 0xa7, 0xcc, 0x52, 0xb3, 0x03, 0x35, 0xbb, 0xef, 0xa8, 0xa3, 0x61, 0x9b,
	0x8d, 0x75, 0x74, 0x74, 0x0b, 0xc1, 0xaf, 0xaf, 0x5a, 0xfe, 0x0f, 0xb1, 0x3f, 0xf0, 0x6b, 0x89,
	0xc7, 0xd0, 0xff, 0xe4, 0x79, 0x63, 0xf1, 0xfd, 0xc8, 0x0e, 0x77, 0xee, 0x8d, 0x33, 0x39, 0x07,
	0x62, 0xf1, 0x31, 0x80, 0xc1, 0xcb, 0xe2, 0x69, 0xf1, 0xfc, 0xba, 0x08, 0xff, 0xe1, 0x10, 0xbc,
	0xc7, 0x38, 0x5e, 0x86, 0x0e, 0xfa, 0xd0, 0x6f, 0xd5, 0x2a, 0x74, 0xd7, 0xc4, 0xdc, 0x70, 0xf5,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x1a, 0x9b, 0x52, 0x9d, 0x01, 0x00, 0x00,
}
