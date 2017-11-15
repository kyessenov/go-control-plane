// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/filter/http/http_connection_manager.proto

package http

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v21 "github.com/envoyproxy/go-control-plane/api"
import envoy_api_v26 "github.com/envoyproxy/go-control-plane/api"
import envoy_api_v25 "github.com/envoyproxy/go-control-plane/api"
import envoy_api_v2_filter1 "github.com/envoyproxy/go-control-plane/api/filter"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"
import google_protobuf4 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HttpConnectionManager_CodecType int32

const (
	HttpConnectionManager_AUTO  HttpConnectionManager_CodecType = 0
	HttpConnectionManager_HTTP1 HttpConnectionManager_CodecType = 1
	HttpConnectionManager_HTTP2 HttpConnectionManager_CodecType = 2
)

var HttpConnectionManager_CodecType_name = map[int32]string{
	0: "AUTO",
	1: "HTTP1",
	2: "HTTP2",
}
var HttpConnectionManager_CodecType_value = map[string]int32{
	"AUTO":  0,
	"HTTP1": 1,
	"HTTP2": 2,
}

func (x HttpConnectionManager_CodecType) String() string {
	return proto.EnumName(HttpConnectionManager_CodecType_name, int32(x))
}
func (HttpConnectionManager_CodecType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{2, 0}
}

type HttpConnectionManager_ForwardClientCertDetails int32

const (
	HttpConnectionManager_SANITIZE            HttpConnectionManager_ForwardClientCertDetails = 0
	HttpConnectionManager_FORWARD_ONLY        HttpConnectionManager_ForwardClientCertDetails = 1
	HttpConnectionManager_APPEND_FORWARD      HttpConnectionManager_ForwardClientCertDetails = 2
	HttpConnectionManager_SANITIZE_SET        HttpConnectionManager_ForwardClientCertDetails = 3
	HttpConnectionManager_ALWAYS_FORWARD_ONLY HttpConnectionManager_ForwardClientCertDetails = 4
)

var HttpConnectionManager_ForwardClientCertDetails_name = map[int32]string{
	0: "SANITIZE",
	1: "FORWARD_ONLY",
	2: "APPEND_FORWARD",
	3: "SANITIZE_SET",
	4: "ALWAYS_FORWARD_ONLY",
}
var HttpConnectionManager_ForwardClientCertDetails_value = map[string]int32{
	"SANITIZE":            0,
	"FORWARD_ONLY":        1,
	"APPEND_FORWARD":      2,
	"SANITIZE_SET":        3,
	"ALWAYS_FORWARD_ONLY": 4,
}

func (x HttpConnectionManager_ForwardClientCertDetails) String() string {
	return proto.EnumName(HttpConnectionManager_ForwardClientCertDetails_name, int32(x))
}
func (HttpConnectionManager_ForwardClientCertDetails) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{2, 1}
}

type HttpConnectionManager_Tracing_OperationName int32

const (
	HttpConnectionManager_Tracing_INGRESS HttpConnectionManager_Tracing_OperationName = 0
	HttpConnectionManager_Tracing_EGRESS  HttpConnectionManager_Tracing_OperationName = 1
)

var HttpConnectionManager_Tracing_OperationName_name = map[int32]string{
	0: "INGRESS",
	1: "EGRESS",
}
var HttpConnectionManager_Tracing_OperationName_value = map[string]int32{
	"INGRESS": 0,
	"EGRESS":  1,
}

func (x HttpConnectionManager_Tracing_OperationName) String() string {
	return proto.EnumName(HttpConnectionManager_Tracing_OperationName_name, int32(x))
}
func (HttpConnectionManager_Tracing_OperationName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{2, 0, 0}
}

type Rds struct {
	ConfigSource    *envoy_api_v21.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource" json:"config_source,omitempty"`
	RouteConfigName string                      `protobuf:"bytes,2,opt,name=route_config_name,json=routeConfigName" json:"route_config_name,omitempty"`
}

func (m *Rds) Reset()                    { *m = Rds{} }
func (m *Rds) String() string            { return proto.CompactTextString(m) }
func (*Rds) ProtoMessage()               {}
func (*Rds) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Rds) GetConfigSource() *envoy_api_v21.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func (m *Rds) GetRouteConfigName() string {
	if m != nil {
		return m.RouteConfigName
	}
	return ""
}

type HttpFilter struct {
	Name         string                   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Config       *google_protobuf4.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
	DeprecatedV1 *HttpFilter_DeprecatedV1 `protobuf:"bytes,3,opt,name=deprecated_v1,json=deprecatedV1" json:"deprecated_v1,omitempty"`
}

func (m *HttpFilter) Reset()                    { *m = HttpFilter{} }
func (m *HttpFilter) String() string            { return proto.CompactTextString(m) }
func (*HttpFilter) ProtoMessage()               {}
func (*HttpFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *HttpFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HttpFilter) GetConfig() *google_protobuf4.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *HttpFilter) GetDeprecatedV1() *HttpFilter_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

type HttpFilter_DeprecatedV1 struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *HttpFilter_DeprecatedV1) Reset()                    { *m = HttpFilter_DeprecatedV1{} }
func (m *HttpFilter_DeprecatedV1) String() string            { return proto.CompactTextString(m) }
func (*HttpFilter_DeprecatedV1) ProtoMessage()               {}
func (*HttpFilter_DeprecatedV1) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1, 0} }

func (m *HttpFilter_DeprecatedV1) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type HttpConnectionManager struct {
	CodecType  HttpConnectionManager_CodecType `protobuf:"varint,1,opt,name=codec_type,json=codecType,enum=envoy.api.v2.filter.http.HttpConnectionManager_CodecType" json:"codec_type,omitempty"`
	StatPrefix string                          `protobuf:"bytes,2,opt,name=stat_prefix,json=statPrefix" json:"stat_prefix,omitempty"`
	// Types that are valid to be assigned to RouteSpecifier:
	//	*HttpConnectionManager_Rds
	//	*HttpConnectionManager_RouteConfig
	RouteSpecifier              isHttpConnectionManager_RouteSpecifier             `protobuf_oneof:"route_specifier"`
	HttpFilters                 []*HttpFilter                                      `protobuf:"bytes,5,rep,name=http_filters,json=httpFilters" json:"http_filters,omitempty"`
	AddUserAgent                *google_protobuf1.BoolValue                        `protobuf:"bytes,6,opt,name=add_user_agent,json=addUserAgent" json:"add_user_agent,omitempty"`
	Tracing                     *HttpConnectionManager_Tracing                     `protobuf:"bytes,7,opt,name=tracing" json:"tracing,omitempty"`
	HttpProtocolOptions         *envoy_api_v26.Http1ProtocolOptions                `protobuf:"bytes,8,opt,name=http_protocol_options,json=httpProtocolOptions" json:"http_protocol_options,omitempty"`
	Http2ProtocolOptions        *envoy_api_v26.Http2ProtocolOptions                `protobuf:"bytes,9,opt,name=http2_protocol_options,json=http2ProtocolOptions" json:"http2_protocol_options,omitempty"`
	ServerName                  string                                             `protobuf:"bytes,10,opt,name=server_name,json=serverName" json:"server_name,omitempty"`
	IdleTimeout                 *google_protobuf.Duration                          `protobuf:"bytes,11,opt,name=idle_timeout,json=idleTimeout" json:"idle_timeout,omitempty"`
	DrainTimeout                *google_protobuf.Duration                          `protobuf:"bytes,12,opt,name=drain_timeout,json=drainTimeout" json:"drain_timeout,omitempty"`
	AccessLog                   []*envoy_api_v2_filter1.AccessLog                  `protobuf:"bytes,13,rep,name=access_log,json=accessLog" json:"access_log,omitempty"`
	UseRemoteAddress            *google_protobuf1.BoolValue                        `protobuf:"bytes,14,opt,name=use_remote_address,json=useRemoteAddress" json:"use_remote_address,omitempty"`
	GenerateRequestId           *google_protobuf1.BoolValue                        `protobuf:"bytes,15,opt,name=generate_request_id,json=generateRequestId" json:"generate_request_id,omitempty"`
	ForwardClientCertDetails    HttpConnectionManager_ForwardClientCertDetails     `protobuf:"varint,16,opt,name=forward_client_cert_details,json=forwardClientCertDetails,enum=envoy.api.v2.filter.http.HttpConnectionManager_ForwardClientCertDetails" json:"forward_client_cert_details,omitempty"`
	SetCurrentClientCertDetails *HttpConnectionManager_SetCurrentClientCertDetails `protobuf:"bytes,17,opt,name=set_current_client_cert_details,json=setCurrentClientCertDetails" json:"set_current_client_cert_details,omitempty"`
}

func (m *HttpConnectionManager) Reset()                    { *m = HttpConnectionManager{} }
func (m *HttpConnectionManager) String() string            { return proto.CompactTextString(m) }
func (*HttpConnectionManager) ProtoMessage()               {}
func (*HttpConnectionManager) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type isHttpConnectionManager_RouteSpecifier interface {
	isHttpConnectionManager_RouteSpecifier()
}

type HttpConnectionManager_Rds struct {
	Rds *Rds `protobuf:"bytes,3,opt,name=rds,oneof"`
}
type HttpConnectionManager_RouteConfig struct {
	RouteConfig *envoy_api_v25.RouteConfiguration `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,oneof"`
}

func (*HttpConnectionManager_Rds) isHttpConnectionManager_RouteSpecifier()         {}
func (*HttpConnectionManager_RouteConfig) isHttpConnectionManager_RouteSpecifier() {}

func (m *HttpConnectionManager) GetRouteSpecifier() isHttpConnectionManager_RouteSpecifier {
	if m != nil {
		return m.RouteSpecifier
	}
	return nil
}

func (m *HttpConnectionManager) GetCodecType() HttpConnectionManager_CodecType {
	if m != nil {
		return m.CodecType
	}
	return HttpConnectionManager_AUTO
}

func (m *HttpConnectionManager) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *HttpConnectionManager) GetRds() *Rds {
	if x, ok := m.GetRouteSpecifier().(*HttpConnectionManager_Rds); ok {
		return x.Rds
	}
	return nil
}

func (m *HttpConnectionManager) GetRouteConfig() *envoy_api_v25.RouteConfiguration {
	if x, ok := m.GetRouteSpecifier().(*HttpConnectionManager_RouteConfig); ok {
		return x.RouteConfig
	}
	return nil
}

func (m *HttpConnectionManager) GetHttpFilters() []*HttpFilter {
	if m != nil {
		return m.HttpFilters
	}
	return nil
}

func (m *HttpConnectionManager) GetAddUserAgent() *google_protobuf1.BoolValue {
	if m != nil {
		return m.AddUserAgent
	}
	return nil
}

func (m *HttpConnectionManager) GetTracing() *HttpConnectionManager_Tracing {
	if m != nil {
		return m.Tracing
	}
	return nil
}

func (m *HttpConnectionManager) GetHttpProtocolOptions() *envoy_api_v26.Http1ProtocolOptions {
	if m != nil {
		return m.HttpProtocolOptions
	}
	return nil
}

func (m *HttpConnectionManager) GetHttp2ProtocolOptions() *envoy_api_v26.Http2ProtocolOptions {
	if m != nil {
		return m.Http2ProtocolOptions
	}
	return nil
}

func (m *HttpConnectionManager) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *HttpConnectionManager) GetIdleTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *HttpConnectionManager) GetDrainTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.DrainTimeout
	}
	return nil
}

func (m *HttpConnectionManager) GetAccessLog() []*envoy_api_v2_filter1.AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

func (m *HttpConnectionManager) GetUseRemoteAddress() *google_protobuf1.BoolValue {
	if m != nil {
		return m.UseRemoteAddress
	}
	return nil
}

func (m *HttpConnectionManager) GetGenerateRequestId() *google_protobuf1.BoolValue {
	if m != nil {
		return m.GenerateRequestId
	}
	return nil
}

func (m *HttpConnectionManager) GetForwardClientCertDetails() HttpConnectionManager_ForwardClientCertDetails {
	if m != nil {
		return m.ForwardClientCertDetails
	}
	return HttpConnectionManager_SANITIZE
}

func (m *HttpConnectionManager) GetSetCurrentClientCertDetails() *HttpConnectionManager_SetCurrentClientCertDetails {
	if m != nil {
		return m.SetCurrentClientCertDetails
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HttpConnectionManager) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HttpConnectionManager_OneofMarshaler, _HttpConnectionManager_OneofUnmarshaler, _HttpConnectionManager_OneofSizer, []interface{}{
		(*HttpConnectionManager_Rds)(nil),
		(*HttpConnectionManager_RouteConfig)(nil),
	}
}

func _HttpConnectionManager_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HttpConnectionManager)
	// route_specifier
	switch x := m.RouteSpecifier.(type) {
	case *HttpConnectionManager_Rds:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rds); err != nil {
			return err
		}
	case *HttpConnectionManager_RouteConfig:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RouteConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HttpConnectionManager.RouteSpecifier has unexpected type %T", x)
	}
	return nil
}

func _HttpConnectionManager_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HttpConnectionManager)
	switch tag {
	case 3: // route_specifier.rds
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Rds)
		err := b.DecodeMessage(msg)
		m.RouteSpecifier = &HttpConnectionManager_Rds{msg}
		return true, err
	case 4: // route_specifier.route_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(envoy_api_v25.RouteConfiguration)
		err := b.DecodeMessage(msg)
		m.RouteSpecifier = &HttpConnectionManager_RouteConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HttpConnectionManager_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HttpConnectionManager)
	// route_specifier
	switch x := m.RouteSpecifier.(type) {
	case *HttpConnectionManager_Rds:
		s := proto.Size(x.Rds)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HttpConnectionManager_RouteConfig:
		s := proto.Size(x.RouteConfig)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HttpConnectionManager_Tracing struct {
	OperationName         HttpConnectionManager_Tracing_OperationName `protobuf:"varint,1,opt,name=operation_name,json=operationName,enum=envoy.api.v2.filter.http.HttpConnectionManager_Tracing_OperationName" json:"operation_name,omitempty"`
	RequestHeadersForTags []string                                    `protobuf:"bytes,2,rep,name=request_headers_for_tags,json=requestHeadersForTags" json:"request_headers_for_tags,omitempty"`
}

func (m *HttpConnectionManager_Tracing) Reset()         { *m = HttpConnectionManager_Tracing{} }
func (m *HttpConnectionManager_Tracing) String() string { return proto.CompactTextString(m) }
func (*HttpConnectionManager_Tracing) ProtoMessage()    {}
func (*HttpConnectionManager_Tracing) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{2, 0}
}

func (m *HttpConnectionManager_Tracing) GetOperationName() HttpConnectionManager_Tracing_OperationName {
	if m != nil {
		return m.OperationName
	}
	return HttpConnectionManager_Tracing_INGRESS
}

func (m *HttpConnectionManager_Tracing) GetRequestHeadersForTags() []string {
	if m != nil {
		return m.RequestHeadersForTags
	}
	return nil
}

type HttpConnectionManager_SetCurrentClientCertDetails struct {
	Subject *google_protobuf1.BoolValue `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
	San     *google_protobuf1.BoolValue `protobuf:"bytes,2,opt,name=san" json:"san,omitempty"`
}

func (m *HttpConnectionManager_SetCurrentClientCertDetails) Reset() {
	*m = HttpConnectionManager_SetCurrentClientCertDetails{}
}
func (m *HttpConnectionManager_SetCurrentClientCertDetails) String() string {
	return proto.CompactTextString(m)
}
func (*HttpConnectionManager_SetCurrentClientCertDetails) ProtoMessage() {}
func (*HttpConnectionManager_SetCurrentClientCertDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{2, 1}
}

func (m *HttpConnectionManager_SetCurrentClientCertDetails) GetSubject() *google_protobuf1.BoolValue {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *HttpConnectionManager_SetCurrentClientCertDetails) GetSan() *google_protobuf1.BoolValue {
	if m != nil {
		return m.San
	}
	return nil
}

func init() {
	proto.RegisterType((*Rds)(nil), "envoy.api.v2.filter.http.Rds")
	proto.RegisterType((*HttpFilter)(nil), "envoy.api.v2.filter.http.HttpFilter")
	proto.RegisterType((*HttpFilter_DeprecatedV1)(nil), "envoy.api.v2.filter.http.HttpFilter.DeprecatedV1")
	proto.RegisterType((*HttpConnectionManager)(nil), "envoy.api.v2.filter.http.HttpConnectionManager")
	proto.RegisterType((*HttpConnectionManager_Tracing)(nil), "envoy.api.v2.filter.http.HttpConnectionManager.Tracing")
	proto.RegisterType((*HttpConnectionManager_SetCurrentClientCertDetails)(nil), "envoy.api.v2.filter.http.HttpConnectionManager.SetCurrentClientCertDetails")
	proto.RegisterEnum("envoy.api.v2.filter.http.HttpConnectionManager_CodecType", HttpConnectionManager_CodecType_name, HttpConnectionManager_CodecType_value)
	proto.RegisterEnum("envoy.api.v2.filter.http.HttpConnectionManager_ForwardClientCertDetails", HttpConnectionManager_ForwardClientCertDetails_name, HttpConnectionManager_ForwardClientCertDetails_value)
	proto.RegisterEnum("envoy.api.v2.filter.http.HttpConnectionManager_Tracing_OperationName", HttpConnectionManager_Tracing_OperationName_name, HttpConnectionManager_Tracing_OperationName_value)
}

func init() { proto.RegisterFile("api/filter/http/http_connection_manager.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 1037 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x4b, 0x6f, 0x1b, 0x37,
	0x10, 0xf6, 0x5a, 0x8e, 0x1d, 0x8d, 0x1e, 0x91, 0x99, 0xa6, 0xd9, 0xca, 0x6d, 0x62, 0x08, 0x3d,
	0x18, 0x7d, 0xac, 0x21, 0xb5, 0x40, 0x50, 0xa0, 0x2f, 0x59, 0x96, 0x23, 0xb7, 0xae, 0xed, 0x52,
	0x8a, 0xf3, 0xb8, 0x10, 0xf4, 0xee, 0x68, 0xbd, 0x85, 0xbc, 0xdc, 0x92, 0x5c, 0xa7, 0x3e, 0xf6,
	0xd4, 0x73, 0xff, 0x56, 0x81, 0xfe, 0x81, 0xfe, 0x9a, 0x82, 0xdc, 0x5d, 0x45, 0x7e, 0xc8, 0xae,
	0x7b, 0x11, 0x86, 0x33, 0xdf, 0xf7, 0x71, 0x86, 0x33, 0x5c, 0x0a, 0x3e, 0xe7, 0x49, 0xb4, 0x39,
	0x8e, 0x26, 0x1a, 0xe5, 0xe6, 0x89, 0xd6, 0x89, 0xfd, 0x61, 0xbe, 0x88, 0x63, 0xf4, 0x75, 0x24,
	0x62, 0x76, 0xca, 0x63, 0x1e, 0xa2, 0xf4, 0x12, 0x29, 0xb4, 0x20, 0x2e, 0xc6, 0x67, 0xe2, 0xdc,
	0xe3, 0x49, 0xe4, 0x9d, 0x75, 0xbc, 0x8c, 0xe7, 0x19, 0x4a, 0xb3, 0x6e, 0x84, 0x8e, 0xb9, 0xc2,
	0x0c, 0xd9, 0x24, 0x66, 0x6d, 0x4d, 0x5f, 0x4c, 0x72, 0x5f, 0xcd, 0xf8, 0x64, 0xa0, 0xf2, 0x65,
	0x73, 0x66, 0x6f, 0xee, 0xfb, 0xa8, 0xd4, 0x44, 0x84, 0x79, 0xec, 0x49, 0x28, 0x44, 0x38, 0xc1,
	0x4c, 0xe1, 0x38, 0x1d, 0x6f, 0x06, 0xa9, 0xe4, 0x26, 0xa1, 0x3c, 0xfe, 0xe1, 0xe5, 0xb8, 0xd2,
	0x32, 0xf5, 0xf5, 0x3c, 0xf6, 0x5b, 0xc9, 0x93, 0x04, 0x65, 0xbe, 0x73, 0x4b, 0x42, 0x89, 0x06,
	0x8a, 0x7c, 0x07, 0x35, 0x5f, 0xc4, 0xe3, 0x28, 0x64, 0x4a, 0xa4, 0xd2, 0x47, 0xd7, 0x59, 0x77,
	0x36, 0x2a, 0x9d, 0xa6, 0x77, 0xa1, 0xca, 0x9e, 0x85, 0x0c, 0x2d, 0x82, 0x56, 0xfd, 0x99, 0x15,
	0xf9, 0x04, 0x56, 0xa5, 0x48, 0x35, 0xb2, 0x5c, 0x26, 0xe6, 0xa7, 0xe8, 0x2e, 0xae, 0x3b, 0x1b,
	0x65, 0xfa, 0xc0, 0x06, 0x32, 0xee, 0x3e, 0x3f, 0xc5, 0xd6, 0xdf, 0x0e, 0xc0, 0x40, 0xeb, 0x64,
	0xc7, 0x16, 0x4c, 0x08, 0x2c, 0x59, 0xb4, 0x63, 0xd1, 0xd6, 0x26, 0x9b, 0xb0, 0x9c, 0x09, 0x59,
	0x8d, 0x4a, 0xe7, 0xb1, 0x97, 0xd5, 0xe1, 0x15, 0x75, 0x78, 0x43, 0x5b, 0x25, 0xcd, 0x61, 0xe4,
	0x08, 0x6a, 0x01, 0x26, 0x12, 0x7d, 0xae, 0x31, 0x60, 0x67, 0x6d, 0xb7, 0x64, 0x79, 0x6d, 0x6f,
	0x5e, 0x9b, 0xbc, 0x77, 0x19, 0x78, 0xdb, 0x53, 0xe6, 0x51, 0x9b, 0x56, 0x83, 0x99, 0x55, 0xb3,
	0x05, 0xd5, 0xd9, 0xa8, 0x49, 0x56, 0x9f, 0x27, 0xd3, 0x64, 0x8d, 0xdd, 0xfa, 0xab, 0x0e, 0x8f,
	0x8c, 0x5a, 0x6f, 0x3a, 0x2b, 0x3f, 0x65, 0xa3, 0x42, 0x5e, 0x01, 0xf8, 0x22, 0x40, 0x9f, 0x4d,
	0x39, 0xf5, 0xce, 0x57, 0x37, 0xa7, 0x74, 0x45, 0xc4, 0xeb, 0x19, 0x85, 0xd1, 0x79, 0x82, 0xb4,
	0xec, 0x17, 0x26, 0x79, 0x0a, 0x15, 0xa5, 0xb9, 0x66, 0x89, 0xc4, 0x71, 0xf4, 0x5b, 0x7e, 0xd2,
	0x60, 0x5c, 0x87, 0xd6, 0x43, 0xda, 0x50, 0x92, 0x81, 0xca, 0x8f, 0xe1, 0xa3, 0xf9, 0x7b, 0xd2,
	0x40, 0x0d, 0x16, 0xa8, 0xc1, 0x92, 0x3e, 0x54, 0x67, 0x7b, 0xe8, 0x2e, 0x59, 0xee, 0xfa, 0x45,
	0x2e, 0x7d, 0xd7, 0xcc, 0x7c, 0x0e, 0x07, 0x0b, 0xb4, 0x32, 0xd3, 0x62, 0xf2, 0x1c, 0xaa, 0xf6,
	0xea, 0x64, 0xbb, 0x28, 0xf7, 0xde, 0x7a, 0x69, 0xa3, 0xd2, 0xf9, 0xf8, 0xbf, 0x74, 0x82, 0x56,
	0x4e, 0xa6, 0xb6, 0x22, 0xdf, 0x43, 0x9d, 0x07, 0x01, 0x4b, 0x15, 0x4a, 0xc6, 0x43, 0x8c, 0xb5,
	0xbb, 0x9c, 0x4f, 0xe5, 0xe5, 0x61, 0xd8, 0x12, 0x62, 0x72, 0xc4, 0x27, 0x29, 0xd2, 0x2a, 0x0f,
	0x82, 0x17, 0x0a, 0x65, 0xd7, 0xe0, 0xc9, 0xcf, 0xb0, 0xa2, 0x25, 0xf7, 0xa3, 0x38, 0x74, 0x57,
	0x2c, 0xf5, 0xd9, 0x5d, 0x0f, 0x7f, 0x94, 0xd1, 0x69, 0xa1, 0x43, 0x8e, 0xe0, 0x91, 0xad, 0xae,
	0xb8, 0xd0, 0x4c, 0x24, 0x06, 0xaf, 0xdc, 0xfb, 0x76, 0x83, 0xd6, 0xc5, 0x0d, 0x8c, 0x68, 0xfb,
	0x30, 0x87, 0x1e, 0x64, 0x48, 0xfa, 0xd0, 0x08, 0x5c, 0x72, 0x92, 0x57, 0xf0, 0xbe, 0x71, 0x77,
	0xae, 0x0a, 0x97, 0xe7, 0x09, 0x77, 0x2e, 0x0b, 0xbf, 0x77, 0x72, 0x8d, 0xd7, 0x8e, 0x0a, 0xca,
	0x33, 0x94, 0xd9, 0xa5, 0x84, 0x7c, 0x54, 0xac, 0xcb, 0xdc, 0x47, 0xf2, 0x35, 0x54, 0xa3, 0x60,
	0x82, 0x4c, 0x47, 0xa7, 0x28, 0x52, 0xed, 0x56, 0xec, 0x86, 0x1f, 0x5c, 0x39, 0xe5, 0xed, 0xbc,
	0xe1, 0xb4, 0x62, 0xe0, 0xa3, 0x0c, 0x4d, 0xbe, 0x85, 0x5a, 0x20, 0x79, 0x14, 0x4f, 0xe9, 0xd5,
	0xdb, 0xe8, 0x55, 0x8b, 0x2f, 0xf8, 0xdf, 0x00, 0x64, 0x9f, 0x3c, 0x36, 0x11, 0xa1, 0x5b, 0xb3,
	0xc3, 0xf2, 0xe4, 0xda, 0x36, 0x75, 0x2d, 0x6c, 0x4f, 0x84, 0xb4, 0xcc, 0x0b, 0x93, 0x0c, 0x80,
	0xa4, 0x0a, 0x99, 0xc4, 0x53, 0xa1, 0x91, 0xf1, 0x20, 0x90, 0xa8, 0x94, 0x5b, 0xbf, 0x75, 0x50,
	0x1a, 0xa9, 0x42, 0x6a, 0x49, 0xdd, 0x8c, 0x43, 0x7e, 0x80, 0x87, 0x21, 0xc6, 0x28, 0xb9, 0x36,
	0x72, 0xbf, 0xa6, 0xa8, 0x34, 0x8b, 0x02, 0xf7, 0xc1, 0xad, 0x52, 0xab, 0x05, 0x8d, 0x66, 0xac,
	0xdd, 0x80, 0xfc, 0xe1, 0xc0, 0xda, 0x58, 0xc8, 0xb7, 0x5c, 0x06, 0xcc, 0x9f, 0x44, 0x18, 0x6b,
	0xe6, 0xa3, 0xd4, 0x2c, 0x40, 0xcd, 0xa3, 0x89, 0x72, 0x1b, 0xf6, 0x53, 0x30, 0xb8, 0xeb, 0x34,
	0xee, 0x64, 0x92, 0x3d, 0xab, 0xd8, 0x43, 0xa9, 0xb7, 0x33, 0x3d, 0xea, 0x8e, 0xe7, 0x44, 0xc8,
	0x9f, 0x0e, 0x3c, 0x55, 0xa8, 0x99, 0x9f, 0x4a, 0x69, 0xd3, 0xb8, 0x26, 0x9b, 0x55, 0x5b, 0xe2,
	0x8f, 0x77, 0xcd, 0x66, 0x88, 0xba, 0x97, 0xa9, 0x5e, 0x4d, 0x68, 0x4d, 0xcd, 0x0f, 0x36, 0xff,
	0x71, 0x60, 0x25, 0xbf, 0x58, 0x64, 0x02, 0x75, 0x91, 0x60, 0x36, 0x19, 0x6c, 0xfa, 0x0e, 0xd4,
	0x3b, 0xfd, 0xff, 0x79, 0x53, 0xbd, 0x83, 0x42, 0xcd, 0xcc, 0x36, 0xad, 0x89, 0xd9, 0x25, 0x79,
	0x06, 0x6e, 0xd1, 0xda, 0x13, 0xe4, 0x01, 0x4a, 0xc5, 0xc6, 0x42, 0x32, 0xcd, 0x43, 0xe5, 0x2e,
	0xae, 0x97, 0x36, 0xca, 0xf4, 0x51, 0x1e, 0x1f, 0x64, 0xe1, 0x1d, 0x21, 0x47, 0x3c, 0x54, 0xad,
	0x0d, 0xa8, 0x5d, 0x10, 0x26, 0x15, 0x58, 0xd9, 0xdd, 0x7f, 0x4e, 0xfb, 0xc3, 0x61, 0x63, 0x81,
	0x00, 0x2c, 0xf7, 0x33, 0xdb, 0x69, 0xfe, 0xee, 0xc0, 0xda, 0x0d, 0x27, 0x43, 0xbe, 0x84, 0x15,
	0x95, 0x1e, 0xff, 0x82, 0xbe, 0x9e, 0x3e, 0xb2, 0xf3, 0x47, 0xab, 0x80, 0x92, 0xcf, 0xa0, 0xa4,
	0x78, 0x9c, 0xbf, 0x86, 0x37, 0x31, 0x0c, 0xac, 0xf5, 0x29, 0x94, 0xa7, 0xaf, 0x06, 0xb9, 0x0f,
	0x4b, 0xdd, 0x17, 0xa3, 0x83, 0xc6, 0x02, 0x29, 0xc3, 0xbd, 0xc1, 0x68, 0x74, 0xd8, 0x6e, 0x38,
	0x85, 0xd9, 0x69, 0x2c, 0xb6, 0xce, 0xc1, 0x9d, 0x37, 0x57, 0xa4, 0x0a, 0xf7, 0x87, 0xdd, 0xfd,
	0xdd, 0xd1, 0xee, 0x9b, 0x7e, 0x63, 0x81, 0x34, 0xa0, 0xba, 0x73, 0x40, 0x5f, 0x76, 0xe9, 0x36,
	0x3b, 0xd8, 0xdf, 0x7b, 0xdd, 0x70, 0x08, 0x81, 0x7a, 0xf7, 0xf0, 0xb0, 0xbf, 0xbf, 0xcd, 0xf2,
	0x40, 0x63, 0xd1, 0xa0, 0x0a, 0x0e, 0x1b, 0xf6, 0x47, 0x8d, 0x12, 0x79, 0x0c, 0x0f, 0xbb, 0x7b,
	0x2f, 0xbb, 0xaf, 0x87, 0xec, 0x02, 0x7d, 0x69, 0x6b, 0x15, 0xb2, 0x3f, 0x07, 0x4c, 0x25, 0xe8,
	0x47, 0xe3, 0x08, 0xe5, 0xd6, 0xf2, 0x9b, 0x25, 0xd3, 0xe4, 0xe3, 0x65, 0x5b, 0xdb, 0x17, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x61, 0x32, 0x35, 0x97, 0x09, 0x00, 0x00,
}
