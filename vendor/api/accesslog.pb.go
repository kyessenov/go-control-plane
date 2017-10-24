// Code generated by protoc-gen-go.
// source: api/accesslog.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/accesslog.proto
	api/address.proto
	api/base.proto
	api/bootstrap.proto
	api/cds.proto
	api/discovery.proto
	api/eds.proto
	api/hds.proto
	api/health_check.proto
	api/lds.proto
	api/protocol.proto
	api/rds.proto
	api/rls.proto
	api/sds.proto

It has these top-level messages:
	EnvoyAccessLog
	Pipe
	SocketAddress
	BindConfig
	Address
	Locality
	Node
	Endpoint
	Metadata
	RuntimeUInt32
	HeaderValue
	HeaderValueOption
	ApiConfigSource
	AggregatedConfigSource
	ConfigSource
	LightstepConfig
	ZipkinConfig
	Tracing
	Admin
	ClusterManager
	StatsdSink
	StatsSink
	TagSpecifier
	StatsConfig
	Watchdog
	Runtime
	RateLimitServiceConfig
	Bootstrap
	UpstreamBindConfig
	CircuitBreakers
	Cluster
	DiscoveryRequest
	DiscoveryResponse
	LbEndpoint
	LocalityLbEndpoints
	EndpointLoadMetricStats
	UpstreamLocalityStats
	ClusterStats
	LoadStatsRequest
	ClusterLoadAssignment
	LoadStatsResponse
	Capability
	HealthCheckRequest
	EndpointHealth
	EndpointHealthResponse
	HealthCheckRequestOrEndpointHealthResponse
	LocalityEndpoints
	ClusterHealthCheck
	HealthCheckSpecifier
	HealthCheck
	Filter
	FilterChainMatch
	FilterChain
	Listener
	TcpProtocolOptions
	Http1ProtocolOptions
	Http2ProtocolOptions
	GrpcProtocolOptions
	WeightedCluster
	RouteMatch
	CorsPolicy
	RouteAction
	RedirectAction
	Decorator
	Route
	VirtualCluster
	RateLimit
	HeaderMatcher
	VirtualHost
	RouteConfiguration
	RateLimitRequest
	RateLimitDescriptor
	RateLimitResponse
	DataSource
	TlsParameters
	TlsCertificate
	TlsSessionTicketKeys
	CertificateValidationContext
	CommonTlsContext
	UpstreamTlsContext
	DownstreamTlsContext
	SdsSecretConfig
	Secret
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf2 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EnvoyAccessLog_Protocol int32

const (
	EnvoyAccessLog_PROTOCOL_UNSPECIFIED EnvoyAccessLog_Protocol = 0
	EnvoyAccessLog_HTTP10               EnvoyAccessLog_Protocol = 1
	EnvoyAccessLog_HTTP11               EnvoyAccessLog_Protocol = 2
	EnvoyAccessLog_HTTP2                EnvoyAccessLog_Protocol = 3
)

var EnvoyAccessLog_Protocol_name = map[int32]string{
	0: "PROTOCOL_UNSPECIFIED",
	1: "HTTP10",
	2: "HTTP11",
	3: "HTTP2",
}
var EnvoyAccessLog_Protocol_value = map[string]int32{
	"PROTOCOL_UNSPECIFIED": 0,
	"HTTP10":               1,
	"HTTP11":               2,
	"HTTP2":                3,
}

func (x EnvoyAccessLog_Protocol) String() string {
	return proto.EnumName(EnvoyAccessLog_Protocol_name, int32(x))
}
func (EnvoyAccessLog_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type EnvoyAccessLog_ResponseFlag int32

const (
	EnvoyAccessLog_FAILED_LOCAL_HEALTHCHECK        EnvoyAccessLog_ResponseFlag = 0
	EnvoyAccessLog_NO_HEALTHY_UPSTREAM             EnvoyAccessLog_ResponseFlag = 1
	EnvoyAccessLog_UPSTREAM_REQUEST_TIMEOUT        EnvoyAccessLog_ResponseFlag = 2
	EnvoyAccessLog_LOCAL_RESET                     EnvoyAccessLog_ResponseFlag = 3
	EnvoyAccessLog_UPSTREAM_REMOTE_RESET           EnvoyAccessLog_ResponseFlag = 4
	EnvoyAccessLog_UPSTREAM_CONNECTION_FAILURE     EnvoyAccessLog_ResponseFlag = 5
	EnvoyAccessLog_UPSTREAM_CONNECTION_TERMINATION EnvoyAccessLog_ResponseFlag = 6
	EnvoyAccessLog_UPSTREAM_OVERFLOW               EnvoyAccessLog_ResponseFlag = 7
	EnvoyAccessLog_NO_ROUTE_FOUND                  EnvoyAccessLog_ResponseFlag = 8
	EnvoyAccessLog_DELAY_INJECTED                  EnvoyAccessLog_ResponseFlag = 9
	EnvoyAccessLog_FAULT_INJECTED                  EnvoyAccessLog_ResponseFlag = 10
	EnvoyAccessLog_RATE_LIMITED                    EnvoyAccessLog_ResponseFlag = 11
)

var EnvoyAccessLog_ResponseFlag_name = map[int32]string{
	0:  "FAILED_LOCAL_HEALTHCHECK",
	1:  "NO_HEALTHY_UPSTREAM",
	2:  "UPSTREAM_REQUEST_TIMEOUT",
	3:  "LOCAL_RESET",
	4:  "UPSTREAM_REMOTE_RESET",
	5:  "UPSTREAM_CONNECTION_FAILURE",
	6:  "UPSTREAM_CONNECTION_TERMINATION",
	7:  "UPSTREAM_OVERFLOW",
	8:  "NO_ROUTE_FOUND",
	9:  "DELAY_INJECTED",
	10: "FAULT_INJECTED",
	11: "RATE_LIMITED",
}
var EnvoyAccessLog_ResponseFlag_value = map[string]int32{
	"FAILED_LOCAL_HEALTHCHECK":        0,
	"NO_HEALTHY_UPSTREAM":             1,
	"UPSTREAM_REQUEST_TIMEOUT":        2,
	"LOCAL_RESET":                     3,
	"UPSTREAM_REMOTE_RESET":           4,
	"UPSTREAM_CONNECTION_FAILURE":     5,
	"UPSTREAM_CONNECTION_TERMINATION": 6,
	"UPSTREAM_OVERFLOW":               7,
	"NO_ROUTE_FOUND":                  8,
	"DELAY_INJECTED":                  9,
	"FAULT_INJECTED":                  10,
	"RATE_LIMITED":                    11,
}

func (x EnvoyAccessLog_ResponseFlag) String() string {
	return proto.EnumName(EnvoyAccessLog_ResponseFlag_name, int32(x))
}
func (EnvoyAccessLog_ResponseFlag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1}
}

type EnvoyAccessLog_TLSVersion int32

const (
	EnvoyAccessLog_VERSION_UNSPECIFIED EnvoyAccessLog_TLSVersion = 0
	EnvoyAccessLog_TLSv1               EnvoyAccessLog_TLSVersion = 1
	EnvoyAccessLog_TLSv1_1             EnvoyAccessLog_TLSVersion = 2
	EnvoyAccessLog_TLSv1_2             EnvoyAccessLog_TLSVersion = 3
	EnvoyAccessLog_TLSv1_3             EnvoyAccessLog_TLSVersion = 4
)

var EnvoyAccessLog_TLSVersion_name = map[int32]string{
	0: "VERSION_UNSPECIFIED",
	1: "TLSv1",
	2: "TLSv1_1",
	3: "TLSv1_2",
	4: "TLSv1_3",
}
var EnvoyAccessLog_TLSVersion_value = map[string]int32{
	"VERSION_UNSPECIFIED": 0,
	"TLSv1":               1,
	"TLSv1_1":             2,
	"TLSv1_2":             3,
	"TLSv1_3":             4,
}

func (x EnvoyAccessLog_TLSVersion) String() string {
	return proto.EnumName(EnvoyAccessLog_TLSVersion_name, int32(x))
}
func (EnvoyAccessLog_TLSVersion) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

type EnvoyAccessLog struct {
	RequestMethod           RequestMethod                 `protobuf:"varint,1,opt,name=request_method,json=requestMethod,enum=envoy.api.v2.RequestMethod" json:"request_method,omitempty"`
	StartTime               *google_protobuf3.Timestamp   `protobuf:"bytes,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	ProtocolVariant         EnvoyAccessLog_Protocol       `protobuf:"varint,3,opt,name=protocol_variant,json=protocolVariant,enum=envoy.api.v2.EnvoyAccessLog_Protocol" json:"protocol_variant,omitempty"`
	ResponseFlags           []EnvoyAccessLog_ResponseFlag `protobuf:"varint,4,rep,packed,name=response_flags,json=responseFlags,enum=envoy.api.v2.EnvoyAccessLog_ResponseFlag" json:"response_flags,omitempty"`
	UpstreamHost            string                        `protobuf:"bytes,5,opt,name=upstream_host,json=upstreamHost" json:"upstream_host,omitempty"`
	UpstreamCluster         string                        `protobuf:"bytes,6,opt,name=upstream_cluster,json=upstreamCluster" json:"upstream_cluster,omitempty"`
	DestinationHost         string                        `protobuf:"bytes,7,opt,name=destination_host,json=destinationHost" json:"destination_host,omitempty"`
	RequestBodyBytes        *google_protobuf2.UInt64Value `protobuf:"bytes,8,opt,name=request_body_bytes,json=requestBodyBytes" json:"request_body_bytes,omitempty"`
	ResponseBodyBytes       *google_protobuf2.UInt64Value `protobuf:"bytes,9,opt,name=response_body_bytes,json=responseBodyBytes" json:"response_body_bytes,omitempty"`
	RequestHeadersBytes     *google_protobuf2.UInt64Value `protobuf:"bytes,10,opt,name=request_headers_bytes,json=requestHeadersBytes" json:"request_headers_bytes,omitempty"`
	ResponseHeadersBytes    *google_protobuf2.UInt64Value `protobuf:"bytes,11,opt,name=response_headers_bytes,json=responseHeadersBytes" json:"response_headers_bytes,omitempty"`
	Secure                  *google_protobuf2.BoolValue   `protobuf:"bytes,12,opt,name=secure" json:"secure,omitempty"`
	HealthCheck             *google_protobuf2.BoolValue   `protobuf:"bytes,13,opt,name=health_check,json=healthCheck" json:"health_check,omitempty"`
	ResponseCode            *google_protobuf2.UInt32Value `protobuf:"bytes,14,opt,name=response_code,json=responseCode" json:"response_code,omitempty"`
	UserAgent               string                        `protobuf:"bytes,15,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	Path                    string                        `protobuf:"bytes,17,opt,name=path" json:"path,omitempty"`
	Referer                 string                        `protobuf:"bytes,18,opt,name=referer" json:"referer,omitempty"`
	ForwardedFor            string                        `protobuf:"bytes,19,opt,name=forwarded_for,json=forwardedFor" json:"forwarded_for,omitempty"`
	RequestId               string                        `protobuf:"bytes,20,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	Authority               string                        `protobuf:"bytes,21,opt,name=authority" json:"authority,omitempty"`
	ResponseDuration        *google_protobuf.Duration     `protobuf:"bytes,22,opt,name=response_duration,json=responseDuration" json:"response_duration,omitempty"`
	UpstreamServiceDuration *google_protobuf.Duration     `protobuf:"bytes,23,opt,name=upstream_service_duration,json=upstreamServiceDuration" json:"upstream_service_duration,omitempty"`
	OriginalPath            string                        `protobuf:"bytes,24,opt,name=original_path,json=originalPath" json:"original_path,omitempty"`
	Metadata                *google_protobuf1.Struct      `protobuf:"bytes,25,opt,name=metadata" json:"metadata,omitempty"`
	RequestHeaders          []*HeaderValue                `protobuf:"bytes,26,rep,name=request_headers,json=requestHeaders" json:"request_headers,omitempty"`
	ResponseHeaders         []*HeaderValue                `protobuf:"bytes,27,rep,name=response_headers,json=responseHeaders" json:"response_headers,omitempty"`
	TlsSniHostname          string                        `protobuf:"bytes,28,opt,name=tls_sni_hostname,json=tlsSniHostname" json:"tls_sni_hostname,omitempty"`
	TlsVersion              EnvoyAccessLog_TLSVersion     `protobuf:"varint,29,opt,name=tls_version,json=tlsVersion,enum=envoy.api.v2.EnvoyAccessLog_TLSVersion" json:"tls_version,omitempty"`
	TlsCipherSuite          *google_protobuf2.UInt32Value `protobuf:"bytes,30,opt,name=tls_cipher_suite,json=tlsCipherSuite" json:"tls_cipher_suite,omitempty"`
}

func (m *EnvoyAccessLog) Reset()                    { *m = EnvoyAccessLog{} }
func (m *EnvoyAccessLog) String() string            { return proto.CompactTextString(m) }
func (*EnvoyAccessLog) ProtoMessage()               {}
func (*EnvoyAccessLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EnvoyAccessLog) GetRequestMethod() RequestMethod {
	if m != nil {
		return m.RequestMethod
	}
	return RequestMethod_METHOD_UNSPECIFIED
}

func (m *EnvoyAccessLog) GetStartTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *EnvoyAccessLog) GetProtocolVariant() EnvoyAccessLog_Protocol {
	if m != nil {
		return m.ProtocolVariant
	}
	return EnvoyAccessLog_PROTOCOL_UNSPECIFIED
}

func (m *EnvoyAccessLog) GetResponseFlags() []EnvoyAccessLog_ResponseFlag {
	if m != nil {
		return m.ResponseFlags
	}
	return nil
}

func (m *EnvoyAccessLog) GetUpstreamHost() string {
	if m != nil {
		return m.UpstreamHost
	}
	return ""
}

func (m *EnvoyAccessLog) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *EnvoyAccessLog) GetDestinationHost() string {
	if m != nil {
		return m.DestinationHost
	}
	return ""
}

func (m *EnvoyAccessLog) GetRequestBodyBytes() *google_protobuf2.UInt64Value {
	if m != nil {
		return m.RequestBodyBytes
	}
	return nil
}

func (m *EnvoyAccessLog) GetResponseBodyBytes() *google_protobuf2.UInt64Value {
	if m != nil {
		return m.ResponseBodyBytes
	}
	return nil
}

func (m *EnvoyAccessLog) GetRequestHeadersBytes() *google_protobuf2.UInt64Value {
	if m != nil {
		return m.RequestHeadersBytes
	}
	return nil
}

func (m *EnvoyAccessLog) GetResponseHeadersBytes() *google_protobuf2.UInt64Value {
	if m != nil {
		return m.ResponseHeadersBytes
	}
	return nil
}

func (m *EnvoyAccessLog) GetSecure() *google_protobuf2.BoolValue {
	if m != nil {
		return m.Secure
	}
	return nil
}

func (m *EnvoyAccessLog) GetHealthCheck() *google_protobuf2.BoolValue {
	if m != nil {
		return m.HealthCheck
	}
	return nil
}

func (m *EnvoyAccessLog) GetResponseCode() *google_protobuf2.UInt32Value {
	if m != nil {
		return m.ResponseCode
	}
	return nil
}

func (m *EnvoyAccessLog) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *EnvoyAccessLog) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *EnvoyAccessLog) GetReferer() string {
	if m != nil {
		return m.Referer
	}
	return ""
}

func (m *EnvoyAccessLog) GetForwardedFor() string {
	if m != nil {
		return m.ForwardedFor
	}
	return ""
}

func (m *EnvoyAccessLog) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EnvoyAccessLog) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *EnvoyAccessLog) GetResponseDuration() *google_protobuf.Duration {
	if m != nil {
		return m.ResponseDuration
	}
	return nil
}

func (m *EnvoyAccessLog) GetUpstreamServiceDuration() *google_protobuf.Duration {
	if m != nil {
		return m.UpstreamServiceDuration
	}
	return nil
}

func (m *EnvoyAccessLog) GetOriginalPath() string {
	if m != nil {
		return m.OriginalPath
	}
	return ""
}

func (m *EnvoyAccessLog) GetMetadata() *google_protobuf1.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *EnvoyAccessLog) GetRequestHeaders() []*HeaderValue {
	if m != nil {
		return m.RequestHeaders
	}
	return nil
}

func (m *EnvoyAccessLog) GetResponseHeaders() []*HeaderValue {
	if m != nil {
		return m.ResponseHeaders
	}
	return nil
}

func (m *EnvoyAccessLog) GetTlsSniHostname() string {
	if m != nil {
		return m.TlsSniHostname
	}
	return ""
}

func (m *EnvoyAccessLog) GetTlsVersion() EnvoyAccessLog_TLSVersion {
	if m != nil {
		return m.TlsVersion
	}
	return EnvoyAccessLog_VERSION_UNSPECIFIED
}

func (m *EnvoyAccessLog) GetTlsCipherSuite() *google_protobuf2.UInt32Value {
	if m != nil {
		return m.TlsCipherSuite
	}
	return nil
}

func init() {
	proto.RegisterType((*EnvoyAccessLog)(nil), "envoy.api.v2.EnvoyAccessLog")
	proto.RegisterEnum("envoy.api.v2.EnvoyAccessLog_Protocol", EnvoyAccessLog_Protocol_name, EnvoyAccessLog_Protocol_value)
	proto.RegisterEnum("envoy.api.v2.EnvoyAccessLog_ResponseFlag", EnvoyAccessLog_ResponseFlag_name, EnvoyAccessLog_ResponseFlag_value)
	proto.RegisterEnum("envoy.api.v2.EnvoyAccessLog_TLSVersion", EnvoyAccessLog_TLSVersion_name, EnvoyAccessLog_TLSVersion_value)
}

func init() { proto.RegisterFile("api/accesslog.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1078 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x73, 0x1a, 0x37,
	0x14, 0x0d, 0xc6, 0x5f, 0x5c, 0x30, 0x6c, 0xe4, 0x38, 0x96, 0x1d, 0x27, 0xf1, 0xd8, 0xd3, 0xa9,
	0xf3, 0x82, 0x6b, 0xdc, 0xe9, 0x4c, 0x1f, 0xfa, 0x80, 0x97, 0xa5, 0x90, 0x2e, 0x2c, 0x5d, 0x16,
	0x67, 0xf2, 0xa4, 0x91, 0x59, 0x19, 0x76, 0xba, 0xb0, 0x54, 0x12, 0x64, 0xfc, 0xd7, 0xfa, 0xaf,
	0xfa, 0x0f, 0x3a, 0xd2, 0x7e, 0x18, 0xec, 0xc6, 0xf6, 0x9b, 0x74, 0xee, 0x39, 0x67, 0x8f, 0xa4,
	0xbb, 0x17, 0x76, 0xe9, 0x2c, 0x38, 0xa7, 0xc3, 0x21, 0x13, 0x22, 0x8c, 0x46, 0xd5, 0x19, 0x8f,
	0x64, 0x84, 0x4a, 0x6c, 0xba, 0x88, 0xee, 0xaa, 0x74, 0x16, 0x54, 0x17, 0xb5, 0xc3, 0xb2, 0xa2,
	0xdc, 0x50, 0xc1, 0xe2, 0xea, 0xe1, 0x87, 0x51, 0x14, 0x8d, 0x42, 0x76, 0xae, 0x77, 0x37, 0xf3,
	0xdb, 0x73, 0x7f, 0xce, 0xa9, 0x0c, 0xa2, 0x69, 0x52, 0x3f, 0x7a, 0x58, 0x17, 0x92, 0xcf, 0x87,
	0x32, 0xa9, 0x7e, 0x7c, 0x58, 0x95, 0xc1, 0x84, 0x09, 0x49, 0x27, 0xb3, 0xef, 0xd9, 0x7f, 0xe3,
	0x74, 0x36, 0x63, 0x5c, 0xc4, 0xf5, 0x93, 0x7f, 0x0d, 0x28, 0x5b, 0x2a, 0x5f, 0x5d, 0xa7, 0xb6,
	0xa3, 0x11, 0xba, 0x82, 0x32, 0x67, 0x7f, 0xcf, 0x99, 0x90, 0x64, 0xc2, 0xe4, 0x38, 0xf2, 0x71,
	0xee, 0x38, 0x77, 0x56, 0xae, 0xbd, 0xab, 0x2e, 0x1f, 0xa4, 0xea, 0xc6, 0x9c, 0x8e, 0xa6, 0xb8,
	0x3b, 0x7c, 0x79, 0x8b, 0x7e, 0x05, 0x10, 0x92, 0x72, 0x49, 0x54, 0x1e, 0xbc, 0x76, 0x9c, 0x3b,
	0x2b, 0xd6, 0x0e, 0xab, 0x71, 0x96, 0x6a, 0x9a, 0xa5, 0xea, 0xa5, 0x61, 0xdd, 0x82, 0x66, 0xab,
	0x3d, 0xea, 0x81, 0xa1, 0x09, 0xc3, 0x28, 0x24, 0x0b, 0xca, 0x03, 0x3a, 0x95, 0x38, 0xaf, 0x03,
	0xfc, 0xb0, 0x1a, 0x60, 0x35, 0x76, 0xb5, 0x97, 0x88, 0xdc, 0x4a, 0x2a, 0xbf, 0x8e, 0xd5, 0xa8,
	0xa7, 0x0e, 0x24, 0x66, 0xd1, 0x54, 0x30, 0x72, 0x1b, 0xd2, 0x91, 0xc0, 0xeb, 0xc7, 0xf9, 0xb3,
	0x72, 0xed, 0xd3, 0x93, 0x7e, 0x6e, 0x22, 0x69, 0x86, 0x74, 0xa4, 0x8e, 0x77, 0xbf, 0x13, 0xe8,
	0x14, 0x76, 0xe6, 0x33, 0x21, 0x39, 0xa3, 0x13, 0x32, 0x8e, 0x84, 0xc4, 0x1b, 0xc7, 0xb9, 0xb3,
	0x82, 0x5b, 0x4a, 0xc1, 0x56, 0x24, 0x24, 0xfa, 0x04, 0x46, 0x46, 0x1a, 0x86, 0x73, 0x21, 0x19,
	0xc7, 0x9b, 0x9a, 0x57, 0x49, 0x71, 0x33, 0x86, 0x15, 0xd5, 0x67, 0x42, 0x06, 0x53, 0xfd, 0xf2,
	0xb1, 0xe5, 0x56, 0x4c, 0x5d, 0xc2, 0xb5, 0xeb, 0x67, 0x40, 0xe9, 0xeb, 0xdc, 0x44, 0xfe, 0x1d,
	0xb9, 0xb9, 0x93, 0x4c, 0xe0, 0x6d, 0x7d, 0xc3, 0x47, 0x8f, 0x6e, 0x78, 0xd0, 0x9e, 0xca, 0x5f,
	0x7e, 0xbe, 0xa6, 0xe1, 0x9c, 0xb9, 0x46, 0xa2, 0xbb, 0x8a, 0xfc, 0xbb, 0x2b, 0xa5, 0x42, 0x36,
	0xec, 0x66, 0x17, 0xb3, 0x64, 0x56, 0x78, 0x81, 0xd9, 0xeb, 0x54, 0x78, 0xef, 0xd6, 0x83, 0xbd,
	0x34, 0xd9, 0x98, 0x51, 0x9f, 0x71, 0x91, 0xf8, 0xc1, 0x0b, 0xfc, 0x76, 0x13, 0x69, 0x2b, 0x56,
	0xc6, 0x8e, 0x2e, 0xbc, 0xcd, 0xf2, 0xad, 0x5a, 0x16, 0x5f, 0x60, 0xf9, 0x26, 0xd5, 0xae, 0x78,
	0xd6, 0x60, 0x53, 0xb0, 0xe1, 0x9c, 0x33, 0x5c, 0xfa, 0x4e, 0x57, 0x5e, 0x45, 0xaa, 0x75, 0x94,
	0x43, 0xc2, 0x44, 0xbf, 0x41, 0x69, 0xcc, 0x68, 0x28, 0xc7, 0x64, 0x38, 0x66, 0xc3, 0xbf, 0xf0,
	0xce, 0xb3, 0xca, 0x62, 0xcc, 0x37, 0x15, 0x1d, 0xd5, 0x21, 0x6b, 0x1f, 0x32, 0x8c, 0x7c, 0x86,
	0xcb, 0x4f, 0xa4, 0xbf, 0xac, 0xc5, 0x0e, 0xa5, 0x54, 0x62, 0x46, 0x3e, 0x43, 0xef, 0x01, 0xe6,
	0x82, 0x71, 0x42, 0x47, 0x6c, 0x2a, 0x71, 0x45, 0xb7, 0x46, 0x41, 0x21, 0x75, 0x05, 0x20, 0x04,
	0xeb, 0x33, 0x2a, 0xc7, 0xf8, 0xb5, 0x2e, 0xe8, 0x35, 0xc2, 0xb0, 0xc5, 0xd9, 0x2d, 0xe3, 0x8c,
	0x63, 0xa4, 0xe1, 0x74, 0xab, 0xba, 0xf7, 0x36, 0xe2, 0xdf, 0x28, 0xf7, 0x99, 0x4f, 0x6e, 0x23,
	0x8e, 0x77, 0xe3, 0xee, 0xcd, 0xc0, 0x66, 0xc4, 0xd5, 0x17, 0xd3, 0xd7, 0x0c, 0x7c, 0xfc, 0x26,
	0xfe, 0x62, 0x82, 0xb4, 0x7d, 0x74, 0x04, 0x05, 0x3a, 0x97, 0xe3, 0x88, 0x07, 0xf2, 0x0e, 0xef,
	0xc5, 0xd5, 0x0c, 0x40, 0x4d, 0xc8, 0xfa, 0x83, 0xa4, 0xf3, 0x0c, 0xbf, 0xd5, 0xa7, 0x3e, 0x78,
	0x74, 0xea, 0x46, 0x42, 0x50, 0x0d, 0x1a, 0x6b, 0x52, 0x04, 0x0d, 0xe0, 0x20, 0xfb, 0x85, 0x04,
	0xe3, 0x8b, 0x60, 0xb8, 0xe4, 0xb7, 0xff, 0x9c, 0xdf, 0x7e, 0xaa, 0xed, 0xc7, 0xd2, 0xcc, 0xf6,
	0x14, 0x76, 0x22, 0x1e, 0x8c, 0x82, 0x29, 0x0d, 0x89, 0xbe, 0x37, 0x1c, 0x5f, 0x40, 0x0a, 0xf6,
	0xd4, 0xfd, 0x5d, 0xc2, 0xf6, 0x84, 0x49, 0xea, 0x53, 0x49, 0xf1, 0x81, 0xfe, 0xd4, 0xfe, 0xa3,
	0x4f, 0xf5, 0xf5, 0x2c, 0x76, 0x33, 0x22, 0xba, 0x82, 0xca, 0x83, 0x7f, 0x00, 0x1f, 0x1e, 0xe7,
	0x75, 0xcc, 0x95, 0x59, 0x13, 0xb7, 0x64, 0xfc, 0xd2, 0xe5, 0xd5, 0xd6, 0x47, 0x0d, 0x30, 0x1e,
	0x76, 0x3d, 0x7e, 0xf7, 0x9c, 0x49, 0xe5, 0x41, 0xb3, 0xa3, 0x33, 0x30, 0x64, 0x28, 0x88, 0x98,
	0x06, 0x7a, 0x9c, 0x4c, 0xe9, 0x84, 0xe1, 0x23, 0x7d, 0xcc, 0xb2, 0x0c, 0x45, 0x7f, 0x1a, 0xb4,
	0x12, 0x14, 0xb5, 0xa0, 0xa8, 0x98, 0x0b, 0xc6, 0x85, 0xba, 0xd6, 0xf7, 0x7a, 0xd6, 0xfe, 0xf8,
	0xe4, 0x6c, 0xf4, 0xec, 0xfe, 0x75, 0x4c, 0x77, 0x41, 0x86, 0x22, 0x59, 0xa3, 0x66, 0xfc, 0xcd,
	0x61, 0x30, 0x1b, 0x33, 0x4e, 0xc4, 0x3c, 0x90, 0x0c, 0x7f, 0x78, 0x41, 0xaf, 0xab, 0x44, 0xa6,
	0x16, 0xf5, 0x95, 0xe6, 0xe4, 0x77, 0xd8, 0x4e, 0xa7, 0x39, 0xc2, 0xf0, 0xa6, 0xe7, 0x3a, 0x9e,
	0x63, 0x3a, 0x36, 0x19, 0x74, 0xfb, 0x3d, 0xcb, 0x6c, 0x37, 0xdb, 0x56, 0xc3, 0x78, 0x85, 0x00,
	0x36, 0x5b, 0x9e, 0xd7, 0xbb, 0xf8, 0xc9, 0xc8, 0x65, 0xeb, 0x0b, 0x63, 0x0d, 0x15, 0x60, 0x43,
	0xad, 0x6b, 0x46, 0xfe, 0xe4, 0x9f, 0x35, 0x28, 0x2d, 0xcf, 0x71, 0x74, 0x04, 0xb8, 0x59, 0x6f,
	0xdb, 0x56, 0x83, 0xd8, 0x8e, 0x59, 0xb7, 0x49, 0xcb, 0xaa, 0xdb, 0x5e, 0xcb, 0x6c, 0x59, 0xe6,
	0x1f, 0xc6, 0x2b, 0xb4, 0x0f, 0xbb, 0x5d, 0x27, 0xc1, 0xbe, 0x92, 0x41, 0xaf, 0xef, 0xb9, 0x56,
	0xbd, 0x63, 0xe4, 0x94, 0x2c, 0xdd, 0x11, 0xd7, 0xfa, 0x73, 0x60, 0xf5, 0x3d, 0xe2, 0xb5, 0x3b,
	0x96, 0x33, 0xf0, 0x8c, 0x35, 0x54, 0x81, 0x62, 0xec, 0xe6, 0x5a, 0x7d, 0xcb, 0x33, 0xf2, 0xe8,
	0x00, 0xf6, 0x96, 0xe8, 0x1d, 0xc7, 0xb3, 0x92, 0xd2, 0x3a, 0xfa, 0x08, 0xef, 0xb2, 0x92, 0xe9,
	0x74, 0xbb, 0x96, 0xe9, 0xb5, 0x9d, 0x2e, 0x51, 0xa1, 0x06, 0xae, 0x65, 0x6c, 0xa0, 0x53, 0xf8,
	0xf8, 0x7f, 0x04, 0xcf, 0x72, 0x3b, 0xed, 0x6e, 0x5d, 0xad, 0x8d, 0x4d, 0xb4, 0x07, 0xaf, 0x33,
	0x92, 0x73, 0x6d, 0xb9, 0x4d, 0xdb, 0xf9, 0x62, 0x6c, 0x21, 0x04, 0xe5, 0xae, 0x43, 0x5c, 0x67,
	0xe0, 0x59, 0xa4, 0xe9, 0x0c, 0xba, 0x0d, 0x63, 0x5b, 0x61, 0x0d, 0xcb, 0xae, 0x7f, 0x25, 0xed,
	0xee, 0x67, 0xcb, 0xf4, 0xac, 0x86, 0x51, 0x50, 0x58, 0xb3, 0x3e, 0xb0, 0xbd, 0x7b, 0x0c, 0x90,
	0x01, 0x25, 0xb7, 0xee, 0x59, 0xc4, 0x6e, 0x77, 0xda, 0x0a, 0x29, 0x9e, 0x7c, 0x01, 0xb8, 0x7f,
	0x67, 0x75, 0x37, 0xd7, 0x96, 0xdb, 0x57, 0x59, 0x56, 0x9f, 0xa1, 0x00, 0x1b, 0x9e, 0xdd, 0x5f,
	0x5c, 0x18, 0x39, 0x54, 0x84, 0x2d, 0xbd, 0x24, 0xea, 0x19, 0xb2, 0x4d, 0xcd, 0xc8, 0xdf, 0x6f,
	0x2e, 0x8d, 0xf5, 0x9b, 0x4d, 0xdd, 0x04, 0x97, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x77, 0xc6,
	0xae, 0x61, 0x2e, 0x09, 0x00, 0x00,
}
