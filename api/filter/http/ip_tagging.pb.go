// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/filter/http/ip_tagging.proto

package http

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2 "github.com/envoyproxy/go-control-plane/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type IPTagging_RequestType int32

const (
	IPTagging_BOTH     IPTagging_RequestType = 0
	IPTagging_INTERNAL IPTagging_RequestType = 1
	IPTagging_EXTERNAL IPTagging_RequestType = 2
)

var IPTagging_RequestType_name = map[int32]string{
	0: "BOTH",
	1: "INTERNAL",
	2: "EXTERNAL",
}
var IPTagging_RequestType_value = map[string]int32{
	"BOTH":     0,
	"INTERNAL": 1,
	"EXTERNAL": 2,
}

func (x IPTagging_RequestType) String() string {
	return proto.EnumName(IPTagging_RequestType_name, int32(x))
}
func (IPTagging_RequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

type IPTagging struct {
	RequestType IPTagging_RequestType `protobuf:"varint,1,opt,name=request_type,json=requestType,enum=envoy.api.v2.filter.http.IPTagging_RequestType" json:"request_type,omitempty"`
	IpTags      []*IPTagging_IPTag    `protobuf:"bytes,2,rep,name=ip_tags,json=ipTags" json:"ip_tags,omitempty"`
}

func (m *IPTagging) Reset()                    { *m = IPTagging{} }
func (m *IPTagging) String() string            { return proto.CompactTextString(m) }
func (*IPTagging) ProtoMessage()               {}
func (*IPTagging) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *IPTagging) GetRequestType() IPTagging_RequestType {
	if m != nil {
		return m.RequestType
	}
	return IPTagging_BOTH
}

func (m *IPTagging) GetIpTags() []*IPTagging_IPTag {
	if m != nil {
		return m.IpTags
	}
	return nil
}

type IPTagging_IPTag struct {
	IpTagName string                    `protobuf:"bytes,1,opt,name=ip_tag_name,json=ipTagName" json:"ip_tag_name,omitempty"`
	IpList    []*envoy_api_v2.CidrRange `protobuf:"bytes,2,rep,name=ip_list,json=ipList" json:"ip_list,omitempty"`
}

func (m *IPTagging_IPTag) Reset()                    { *m = IPTagging_IPTag{} }
func (m *IPTagging_IPTag) String() string            { return proto.CompactTextString(m) }
func (*IPTagging_IPTag) ProtoMessage()               {}
func (*IPTagging_IPTag) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

func (m *IPTagging_IPTag) GetIpTagName() string {
	if m != nil {
		return m.IpTagName
	}
	return ""
}

func (m *IPTagging_IPTag) GetIpList() []*envoy_api_v2.CidrRange {
	if m != nil {
		return m.IpList
	}
	return nil
}

func init() {
	proto.RegisterType((*IPTagging)(nil), "envoy.api.v2.filter.http.IPTagging")
	proto.RegisterType((*IPTagging_IPTag)(nil), "envoy.api.v2.filter.http.IPTagging.IPTag")
	proto.RegisterEnum("envoy.api.v2.filter.http.IPTagging_RequestType", IPTagging_RequestType_name, IPTagging_RequestType_value)
}

func init() { proto.RegisterFile("api/filter/http/ip_tagging.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x6d, 0xd5, 0xb9, 0xa6, 0x43, 0x66, 0x2e, 0x96, 0x1d, 0xa4, 0xec, 0x34, 0x2f, 0xa9,
	0x74, 0x9f, 0xc0, 0xc9, 0xc0, 0xc1, 0xa8, 0x12, 0x7a, 0xd0, 0x53, 0x89, 0x34, 0xc6, 0x17, 0xb6,
	0x36, 0x26, 0xaf, 0x83, 0x7e, 0x2e, 0xbf, 0xa0, 0x34, 0x19, 0xfe, 0x39, 0x08, 0xbb, 0xe5, 0x09,
	0xcf, 0xf3, 0x7b, 0x1e, 0x5e, 0x92, 0x0a, 0x0d, 0xd9, 0x2b, 0x6c, 0x50, 0x9a, 0xec, 0x0d, 0x51,
	0x67, 0xa0, 0x2b, 0x14, 0x4a, 0x41, 0xa3, 0x98, 0x36, 0x2d, 0xb6, 0x34, 0x91, 0xcd, 0xae, 0xed,
	0x98, 0xd0, 0xc0, 0x76, 0x39, 0xf3, 0x56, 0xd6, 0x5b, 0x27, 0x17, 0x7d, 0x56, 0xd4, 0xb5, 0x91,
	0xd6, 0x7a, 0xf3, 0xf4, 0x33, 0x24, 0xd1, 0xea, 0xb1, 0xf4, 0x00, 0xca, 0xc9, 0xc8, 0xc8, 0xf7,
	0x0f, 0x69, 0xb1, 0xc2, 0x4e, 0xcb, 0x24, 0x48, 0x83, 0xd9, 0x79, 0x9e, 0xb1, 0xff, 0x88, 0xec,
	0x3b, 0xca, 0xb8, 0xcf, 0x95, 0x9d, 0x96, 0x3c, 0x36, 0x3f, 0x82, 0x2e, 0xc8, 0x99, 0x9f, 0x68,
	0x93, 0x30, 0x3d, 0x9e, 0xc5, 0xf9, 0xf5, 0x21, 0x38, 0xf7, 0xe2, 0x03, 0xd0, 0xa5, 0x50, 0x76,
	0xf2, 0x4c, 0x4e, 0xdd, 0x07, 0xbd, 0x22, 0xb1, 0x87, 0x55, 0x8d, 0xd8, 0xfa, 0x7d, 0x11, 0x8f,
	0x9c, 0xab, 0x10, 0x5b, 0x49, 0x6f, 0x5c, 0xd9, 0x06, 0x2c, 0xee, 0xcb, 0x2e, 0xff, 0x96, 0xdd,
	0x41, 0x6d, 0xb8, 0x68, 0x94, 0xec, 0xd1, 0x6b, 0xb0, 0x38, 0x9d, 0x93, 0xf8, 0xd7, 0x74, 0x3a,
	0x24, 0x27, 0x8b, 0x87, 0xf2, 0x7e, 0x7c, 0x44, 0x47, 0x64, 0xb8, 0x2a, 0xca, 0x25, 0x2f, 0x6e,
	0xd7, 0xe3, 0xa0, 0x57, 0xcb, 0xa7, 0xbd, 0x0a, 0x5f, 0x06, 0xee, 0x78, 0xf3, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xdb, 0x11, 0xca, 0xd8, 0x8d, 0x01, 0x00, 0x00,
}
