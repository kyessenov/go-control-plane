// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/sds.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TlsParameters_TlsProtocol int32

const (
	TlsParameters_TLS_AUTO TlsParameters_TlsProtocol = 0
	TlsParameters_TLSv1_0  TlsParameters_TlsProtocol = 1
	TlsParameters_TLSv1_1  TlsParameters_TlsProtocol = 2
	TlsParameters_TLSv1_2  TlsParameters_TlsProtocol = 3
	TlsParameters_TLSv1_3  TlsParameters_TlsProtocol = 4
)

var TlsParameters_TlsProtocol_name = map[int32]string{
	0: "TLS_AUTO",
	1: "TLSv1_0",
	2: "TLSv1_1",
	3: "TLSv1_2",
	4: "TLSv1_3",
}
var TlsParameters_TlsProtocol_value = map[string]int32{
	"TLS_AUTO": 0,
	"TLSv1_0":  1,
	"TLSv1_1":  2,
	"TLSv1_2":  3,
	"TLSv1_3":  4,
}

func (x TlsParameters_TlsProtocol) String() string {
	return proto.EnumName(TlsParameters_TlsProtocol_name, int32(x))
}
func (TlsParameters_TlsProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor13, []int{1, 0}
}

type DataSource struct {
	// Types that are valid to be assigned to Specifier:
	//	*DataSource_Filename
	//	*DataSource_Inline
	Specifier isDataSource_Specifier `protobuf_oneof:"specifier"`
}

func (m *DataSource) Reset()                    { *m = DataSource{} }
func (m *DataSource) String() string            { return proto.CompactTextString(m) }
func (*DataSource) ProtoMessage()               {}
func (*DataSource) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

type isDataSource_Specifier interface {
	isDataSource_Specifier()
}

type DataSource_Filename struct {
	Filename string `protobuf:"bytes,1,opt,name=filename,oneof"`
}
type DataSource_Inline struct {
	Inline []byte `protobuf:"bytes,2,opt,name=inline,proto3,oneof"`
}

func (*DataSource_Filename) isDataSource_Specifier() {}
func (*DataSource_Inline) isDataSource_Specifier()   {}

func (m *DataSource) GetSpecifier() isDataSource_Specifier {
	if m != nil {
		return m.Specifier
	}
	return nil
}

func (m *DataSource) GetFilename() string {
	if x, ok := m.GetSpecifier().(*DataSource_Filename); ok {
		return x.Filename
	}
	return ""
}

func (m *DataSource) GetInline() []byte {
	if x, ok := m.GetSpecifier().(*DataSource_Inline); ok {
		return x.Inline
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DataSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DataSource_OneofMarshaler, _DataSource_OneofUnmarshaler, _DataSource_OneofSizer, []interface{}{
		(*DataSource_Filename)(nil),
		(*DataSource_Inline)(nil),
	}
}

func _DataSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DataSource)
	// specifier
	switch x := m.Specifier.(type) {
	case *DataSource_Filename:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Filename)
	case *DataSource_Inline:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Inline)
	case nil:
	default:
		return fmt.Errorf("DataSource.Specifier has unexpected type %T", x)
	}
	return nil
}

func _DataSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DataSource)
	switch tag {
	case 1: // specifier.filename
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Specifier = &DataSource_Filename{x}
		return true, err
	case 2: // specifier.inline
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Specifier = &DataSource_Inline{x}
		return true, err
	default:
		return false, nil
	}
}

func _DataSource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DataSource)
	// specifier
	switch x := m.Specifier.(type) {
	case *DataSource_Filename:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Filename)))
		n += len(x.Filename)
	case *DataSource_Inline:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Inline)))
		n += len(x.Inline)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TlsParameters struct {
	TlsMinimumProtocolVersion TlsParameters_TlsProtocol `protobuf:"varint,1,opt,name=tls_minimum_protocol_version,json=tlsMinimumProtocolVersion,enum=envoy.api.v2.TlsParameters_TlsProtocol" json:"tls_minimum_protocol_version,omitempty"`
	TlsMaximumProtocolVersion TlsParameters_TlsProtocol `protobuf:"varint,2,opt,name=tls_maximum_protocol_version,json=tlsMaximumProtocolVersion,enum=envoy.api.v2.TlsParameters_TlsProtocol" json:"tls_maximum_protocol_version,omitempty"`
	CipherSuites              []string                  `protobuf:"bytes,3,rep,name=cipher_suites,json=cipherSuites" json:"cipher_suites,omitempty"`
	EcdhCurves                []string                  `protobuf:"bytes,4,rep,name=ecdh_curves,json=ecdhCurves" json:"ecdh_curves,omitempty"`
}

func (m *TlsParameters) Reset()                    { *m = TlsParameters{} }
func (m *TlsParameters) String() string            { return proto.CompactTextString(m) }
func (*TlsParameters) ProtoMessage()               {}
func (*TlsParameters) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *TlsParameters) GetTlsMinimumProtocolVersion() TlsParameters_TlsProtocol {
	if m != nil {
		return m.TlsMinimumProtocolVersion
	}
	return TlsParameters_TLS_AUTO
}

func (m *TlsParameters) GetTlsMaximumProtocolVersion() TlsParameters_TlsProtocol {
	if m != nil {
		return m.TlsMaximumProtocolVersion
	}
	return TlsParameters_TLS_AUTO
}

func (m *TlsParameters) GetCipherSuites() []string {
	if m != nil {
		return m.CipherSuites
	}
	return nil
}

func (m *TlsParameters) GetEcdhCurves() []string {
	if m != nil {
		return m.EcdhCurves
	}
	return nil
}

type TlsCertificate struct {
	CertificateChain           *DataSource   `protobuf:"bytes,1,opt,name=certificate_chain,json=certificateChain" json:"certificate_chain,omitempty"`
	PrivateKey                 *DataSource   `protobuf:"bytes,2,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	Password                   *DataSource   `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	OcspStaple                 *DataSource   `protobuf:"bytes,4,opt,name=ocsp_staple,json=ocspStaple" json:"ocsp_staple,omitempty"`
	SignedCertificateTimestamp []*DataSource `protobuf:"bytes,5,rep,name=signed_certificate_timestamp,json=signedCertificateTimestamp" json:"signed_certificate_timestamp,omitempty"`
}

func (m *TlsCertificate) Reset()                    { *m = TlsCertificate{} }
func (m *TlsCertificate) String() string            { return proto.CompactTextString(m) }
func (*TlsCertificate) ProtoMessage()               {}
func (*TlsCertificate) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *TlsCertificate) GetCertificateChain() *DataSource {
	if m != nil {
		return m.CertificateChain
	}
	return nil
}

func (m *TlsCertificate) GetPrivateKey() *DataSource {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *TlsCertificate) GetPassword() *DataSource {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *TlsCertificate) GetOcspStaple() *DataSource {
	if m != nil {
		return m.OcspStaple
	}
	return nil
}

func (m *TlsCertificate) GetSignedCertificateTimestamp() []*DataSource {
	if m != nil {
		return m.SignedCertificateTimestamp
	}
	return nil
}

type TlsSessionTicketKeys struct {
	Keys []*DataSource `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
}

func (m *TlsSessionTicketKeys) Reset()                    { *m = TlsSessionTicketKeys{} }
func (m *TlsSessionTicketKeys) String() string            { return proto.CompactTextString(m) }
func (*TlsSessionTicketKeys) ProtoMessage()               {}
func (*TlsSessionTicketKeys) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *TlsSessionTicketKeys) GetKeys() []*DataSource {
	if m != nil {
		return m.Keys
	}
	return nil
}

type CertificateValidationContext struct {
	TrustedCa                         *DataSource                `protobuf:"bytes,1,opt,name=trusted_ca,json=trustedCa" json:"trusted_ca,omitempty"`
	VerifyCertificateHash             []string                   `protobuf:"bytes,2,rep,name=verify_certificate_hash,json=verifyCertificateHash" json:"verify_certificate_hash,omitempty"`
	VerifySpkiSha256                  []string                   `protobuf:"bytes,3,rep,name=verify_spki_sha256,json=verifySpkiSha256" json:"verify_spki_sha256,omitempty"`
	VerifySubjectAltName              []string                   `protobuf:"bytes,4,rep,name=verify_subject_alt_name,json=verifySubjectAltName" json:"verify_subject_alt_name,omitempty"`
	RequireOcspStaple                 *google_protobuf.BoolValue `protobuf:"bytes,5,opt,name=require_ocsp_staple,json=requireOcspStaple" json:"require_ocsp_staple,omitempty"`
	RequireSignedCertificateTimestamp *google_protobuf.BoolValue `protobuf:"bytes,6,opt,name=require_signed_certificate_timestamp,json=requireSignedCertificateTimestamp" json:"require_signed_certificate_timestamp,omitempty"`
}

func (m *CertificateValidationContext) Reset()                    { *m = CertificateValidationContext{} }
func (m *CertificateValidationContext) String() string            { return proto.CompactTextString(m) }
func (*CertificateValidationContext) ProtoMessage()               {}
func (*CertificateValidationContext) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func (m *CertificateValidationContext) GetTrustedCa() *DataSource {
	if m != nil {
		return m.TrustedCa
	}
	return nil
}

func (m *CertificateValidationContext) GetVerifyCertificateHash() []string {
	if m != nil {
		return m.VerifyCertificateHash
	}
	return nil
}

func (m *CertificateValidationContext) GetVerifySpkiSha256() []string {
	if m != nil {
		return m.VerifySpkiSha256
	}
	return nil
}

func (m *CertificateValidationContext) GetVerifySubjectAltName() []string {
	if m != nil {
		return m.VerifySubjectAltName
	}
	return nil
}

func (m *CertificateValidationContext) GetRequireOcspStaple() *google_protobuf.BoolValue {
	if m != nil {
		return m.RequireOcspStaple
	}
	return nil
}

func (m *CertificateValidationContext) GetRequireSignedCertificateTimestamp() *google_protobuf.BoolValue {
	if m != nil {
		return m.RequireSignedCertificateTimestamp
	}
	return nil
}

type CommonTlsContext struct {
	TlsParams                      *TlsParameters                 `protobuf:"bytes,1,opt,name=tls_params,json=tlsParams" json:"tls_params,omitempty"`
	TlsCertificates                []*TlsCertificate              `protobuf:"bytes,2,rep,name=tls_certificates,json=tlsCertificates" json:"tls_certificates,omitempty"`
	TlsCertificateSdsSecretConfigs []*SdsSecretConfig             `protobuf:"bytes,6,rep,name=tls_certificate_sds_secret_configs,json=tlsCertificateSdsSecretConfigs" json:"tls_certificate_sds_secret_configs,omitempty"`
	ValidationContext              *CertificateValidationContext  `protobuf:"bytes,3,opt,name=validation_context,json=validationContext" json:"validation_context,omitempty"`
	AlpnProtocols                  []string                       `protobuf:"bytes,4,rep,name=alpn_protocols,json=alpnProtocols" json:"alpn_protocols,omitempty"`
	DeprecatedV1                   *CommonTlsContext_DeprecatedV1 `protobuf:"bytes,5,opt,name=deprecated_v1,json=deprecatedV1" json:"deprecated_v1,omitempty"`
}

func (m *CommonTlsContext) Reset()                    { *m = CommonTlsContext{} }
func (m *CommonTlsContext) String() string            { return proto.CompactTextString(m) }
func (*CommonTlsContext) ProtoMessage()               {}
func (*CommonTlsContext) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

func (m *CommonTlsContext) GetTlsParams() *TlsParameters {
	if m != nil {
		return m.TlsParams
	}
	return nil
}

func (m *CommonTlsContext) GetTlsCertificates() []*TlsCertificate {
	if m != nil {
		return m.TlsCertificates
	}
	return nil
}

func (m *CommonTlsContext) GetTlsCertificateSdsSecretConfigs() []*SdsSecretConfig {
	if m != nil {
		return m.TlsCertificateSdsSecretConfigs
	}
	return nil
}

func (m *CommonTlsContext) GetValidationContext() *CertificateValidationContext {
	if m != nil {
		return m.ValidationContext
	}
	return nil
}

func (m *CommonTlsContext) GetAlpnProtocols() []string {
	if m != nil {
		return m.AlpnProtocols
	}
	return nil
}

func (m *CommonTlsContext) GetDeprecatedV1() *CommonTlsContext_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

type CommonTlsContext_DeprecatedV1 struct {
	AltAlpnProtocols string `protobuf:"bytes,1,opt,name=alt_alpn_protocols,json=altAlpnProtocols" json:"alt_alpn_protocols,omitempty"`
}

func (m *CommonTlsContext_DeprecatedV1) Reset()         { *m = CommonTlsContext_DeprecatedV1{} }
func (m *CommonTlsContext_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*CommonTlsContext_DeprecatedV1) ProtoMessage()    {}
func (*CommonTlsContext_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{5, 0}
}

func (m *CommonTlsContext_DeprecatedV1) GetAltAlpnProtocols() string {
	if m != nil {
		return m.AltAlpnProtocols
	}
	return ""
}

type UpstreamTlsContext struct {
	CommonTlsContext *CommonTlsContext `protobuf:"bytes,1,opt,name=common_tls_context,json=commonTlsContext" json:"common_tls_context,omitempty"`
	Sni              string            `protobuf:"bytes,2,opt,name=sni" json:"sni,omitempty"`
}

func (m *UpstreamTlsContext) Reset()                    { *m = UpstreamTlsContext{} }
func (m *UpstreamTlsContext) String() string            { return proto.CompactTextString(m) }
func (*UpstreamTlsContext) ProtoMessage()               {}
func (*UpstreamTlsContext) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{6} }

func (m *UpstreamTlsContext) GetCommonTlsContext() *CommonTlsContext {
	if m != nil {
		return m.CommonTlsContext
	}
	return nil
}

func (m *UpstreamTlsContext) GetSni() string {
	if m != nil {
		return m.Sni
	}
	return ""
}

type DownstreamTlsContext struct {
	CommonTlsContext         *CommonTlsContext          `protobuf:"bytes,1,opt,name=common_tls_context,json=commonTlsContext" json:"common_tls_context,omitempty"`
	RequireClientCertificate *google_protobuf.BoolValue `protobuf:"bytes,2,opt,name=require_client_certificate,json=requireClientCertificate" json:"require_client_certificate,omitempty"`
	RequireSni               *google_protobuf.BoolValue `protobuf:"bytes,3,opt,name=require_sni,json=requireSni" json:"require_sni,omitempty"`
	// Types that are valid to be assigned to SessionTicketKeysType:
	//	*DownstreamTlsContext_SessionTicketKeys
	//	*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig
	SessionTicketKeysType isDownstreamTlsContext_SessionTicketKeysType `protobuf_oneof:"session_ticket_keys_type"`
}

func (m *DownstreamTlsContext) Reset()                    { *m = DownstreamTlsContext{} }
func (m *DownstreamTlsContext) String() string            { return proto.CompactTextString(m) }
func (*DownstreamTlsContext) ProtoMessage()               {}
func (*DownstreamTlsContext) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{7} }

type isDownstreamTlsContext_SessionTicketKeysType interface {
	isDownstreamTlsContext_SessionTicketKeysType()
}

type DownstreamTlsContext_SessionTicketKeys struct {
	SessionTicketKeys *TlsSessionTicketKeys `protobuf:"bytes,4,opt,name=session_ticket_keys,json=sessionTicketKeys,oneof"`
}
type DownstreamTlsContext_SessionTicketKeysSdsSecretConfig struct {
	SessionTicketKeysSdsSecretConfig *SdsSecretConfig `protobuf:"bytes,5,opt,name=session_ticket_keys_sds_secret_config,json=sessionTicketKeysSdsSecretConfig,oneof"`
}

func (*DownstreamTlsContext_SessionTicketKeys) isDownstreamTlsContext_SessionTicketKeysType() {}
func (*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig) isDownstreamTlsContext_SessionTicketKeysType() {
}

func (m *DownstreamTlsContext) GetSessionTicketKeysType() isDownstreamTlsContext_SessionTicketKeysType {
	if m != nil {
		return m.SessionTicketKeysType
	}
	return nil
}

func (m *DownstreamTlsContext) GetCommonTlsContext() *CommonTlsContext {
	if m != nil {
		return m.CommonTlsContext
	}
	return nil
}

func (m *DownstreamTlsContext) GetRequireClientCertificate() *google_protobuf.BoolValue {
	if m != nil {
		return m.RequireClientCertificate
	}
	return nil
}

func (m *DownstreamTlsContext) GetRequireSni() *google_protobuf.BoolValue {
	if m != nil {
		return m.RequireSni
	}
	return nil
}

func (m *DownstreamTlsContext) GetSessionTicketKeys() *TlsSessionTicketKeys {
	if x, ok := m.GetSessionTicketKeysType().(*DownstreamTlsContext_SessionTicketKeys); ok {
		return x.SessionTicketKeys
	}
	return nil
}

func (m *DownstreamTlsContext) GetSessionTicketKeysSdsSecretConfig() *SdsSecretConfig {
	if x, ok := m.GetSessionTicketKeysType().(*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig); ok {
		return x.SessionTicketKeysSdsSecretConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DownstreamTlsContext) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DownstreamTlsContext_OneofMarshaler, _DownstreamTlsContext_OneofUnmarshaler, _DownstreamTlsContext_OneofSizer, []interface{}{
		(*DownstreamTlsContext_SessionTicketKeys)(nil),
		(*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig)(nil),
	}
}

func _DownstreamTlsContext_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DownstreamTlsContext)
	// session_ticket_keys_type
	switch x := m.SessionTicketKeysType.(type) {
	case *DownstreamTlsContext_SessionTicketKeys:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SessionTicketKeys); err != nil {
			return err
		}
	case *DownstreamTlsContext_SessionTicketKeysSdsSecretConfig:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SessionTicketKeysSdsSecretConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DownstreamTlsContext.SessionTicketKeysType has unexpected type %T", x)
	}
	return nil
}

func _DownstreamTlsContext_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DownstreamTlsContext)
	switch tag {
	case 4: // session_ticket_keys_type.session_ticket_keys
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TlsSessionTicketKeys)
		err := b.DecodeMessage(msg)
		m.SessionTicketKeysType = &DownstreamTlsContext_SessionTicketKeys{msg}
		return true, err
	case 5: // session_ticket_keys_type.session_ticket_keys_sds_secret_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SdsSecretConfig)
		err := b.DecodeMessage(msg)
		m.SessionTicketKeysType = &DownstreamTlsContext_SessionTicketKeysSdsSecretConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DownstreamTlsContext_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DownstreamTlsContext)
	// session_ticket_keys_type
	switch x := m.SessionTicketKeysType.(type) {
	case *DownstreamTlsContext_SessionTicketKeys:
		s := proto.Size(x.SessionTicketKeys)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DownstreamTlsContext_SessionTicketKeysSdsSecretConfig:
		s := proto.Size(x.SessionTicketKeysSdsSecretConfig)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SdsSecretConfig struct {
	Name      string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SdsConfig *ConfigSource `protobuf:"bytes,2,opt,name=sds_config,json=sdsConfig" json:"sds_config,omitempty"`
}

func (m *SdsSecretConfig) Reset()                    { *m = SdsSecretConfig{} }
func (m *SdsSecretConfig) String() string            { return proto.CompactTextString(m) }
func (*SdsSecretConfig) ProtoMessage()               {}
func (*SdsSecretConfig) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{8} }

func (m *SdsSecretConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SdsSecretConfig) GetSdsConfig() *ConfigSource {
	if m != nil {
		return m.SdsConfig
	}
	return nil
}

type Secret struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*Secret_TlsCertificate
	//	*Secret_SessionTicketKeys
	Type isSecret_Type `protobuf_oneof:"type"`
}

func (m *Secret) Reset()                    { *m = Secret{} }
func (m *Secret) String() string            { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()               {}
func (*Secret) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{9} }

type isSecret_Type interface {
	isSecret_Type()
}

type Secret_TlsCertificate struct {
	TlsCertificate *TlsCertificate `protobuf:"bytes,2,opt,name=tls_certificate,json=tlsCertificate,oneof"`
}
type Secret_SessionTicketKeys struct {
	SessionTicketKeys *TlsSessionTicketKeys `protobuf:"bytes,3,opt,name=session_ticket_keys,json=sessionTicketKeys,oneof"`
}

func (*Secret_TlsCertificate) isSecret_Type()    {}
func (*Secret_SessionTicketKeys) isSecret_Type() {}

func (m *Secret) GetType() isSecret_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Secret) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Secret) GetTlsCertificate() *TlsCertificate {
	if x, ok := m.GetType().(*Secret_TlsCertificate); ok {
		return x.TlsCertificate
	}
	return nil
}

func (m *Secret) GetSessionTicketKeys() *TlsSessionTicketKeys {
	if x, ok := m.GetType().(*Secret_SessionTicketKeys); ok {
		return x.SessionTicketKeys
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Secret) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Secret_OneofMarshaler, _Secret_OneofUnmarshaler, _Secret_OneofSizer, []interface{}{
		(*Secret_TlsCertificate)(nil),
		(*Secret_SessionTicketKeys)(nil),
	}
}

func _Secret_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Secret)
	// type
	switch x := m.Type.(type) {
	case *Secret_TlsCertificate:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TlsCertificate); err != nil {
			return err
		}
	case *Secret_SessionTicketKeys:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SessionTicketKeys); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Secret.Type has unexpected type %T", x)
	}
	return nil
}

func _Secret_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Secret)
	switch tag {
	case 2: // type.tls_certificate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TlsCertificate)
		err := b.DecodeMessage(msg)
		m.Type = &Secret_TlsCertificate{msg}
		return true, err
	case 3: // type.session_ticket_keys
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TlsSessionTicketKeys)
		err := b.DecodeMessage(msg)
		m.Type = &Secret_SessionTicketKeys{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Secret_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Secret)
	// type
	switch x := m.Type.(type) {
	case *Secret_TlsCertificate:
		s := proto.Size(x.TlsCertificate)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Secret_SessionTicketKeys:
		s := proto.Size(x.SessionTicketKeys)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*DataSource)(nil), "envoy.api.v2.DataSource")
	proto.RegisterType((*TlsParameters)(nil), "envoy.api.v2.TlsParameters")
	proto.RegisterType((*TlsCertificate)(nil), "envoy.api.v2.TlsCertificate")
	proto.RegisterType((*TlsSessionTicketKeys)(nil), "envoy.api.v2.TlsSessionTicketKeys")
	proto.RegisterType((*CertificateValidationContext)(nil), "envoy.api.v2.CertificateValidationContext")
	proto.RegisterType((*CommonTlsContext)(nil), "envoy.api.v2.CommonTlsContext")
	proto.RegisterType((*CommonTlsContext_DeprecatedV1)(nil), "envoy.api.v2.CommonTlsContext.DeprecatedV1")
	proto.RegisterType((*UpstreamTlsContext)(nil), "envoy.api.v2.UpstreamTlsContext")
	proto.RegisterType((*DownstreamTlsContext)(nil), "envoy.api.v2.DownstreamTlsContext")
	proto.RegisterType((*SdsSecretConfig)(nil), "envoy.api.v2.SdsSecretConfig")
	proto.RegisterType((*Secret)(nil), "envoy.api.v2.Secret")
	proto.RegisterEnum("envoy.api.v2.TlsParameters_TlsProtocol", TlsParameters_TlsProtocol_name, TlsParameters_TlsProtocol_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SecretDiscoveryService service

type SecretDiscoveryServiceClient interface {
	StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error)
	FetchSecrets(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type secretDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecretDiscoveryServiceClient(cc *grpc.ClientConn) SecretDiscoveryServiceClient {
	return &secretDiscoveryServiceClient{cc}
}

func (c *secretDiscoveryServiceClient) StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[0], c.cc, "/envoy.api.v2.SecretDiscoveryService/StreamSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceStreamSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_StreamSecretsClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceStreamSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceStreamSecretsClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) FetchSecrets(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/envoy.api.v2.SecretDiscoveryService/FetchSecrets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SecretDiscoveryService service

type SecretDiscoveryServiceServer interface {
	StreamSecrets(SecretDiscoveryService_StreamSecretsServer) error
	FetchSecrets(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterSecretDiscoveryServiceServer(s *grpc.Server, srv SecretDiscoveryServiceServer) {
	s.RegisterService(&_SecretDiscoveryService_serviceDesc, srv)
}

func _SecretDiscoveryService_StreamSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).StreamSecrets(&secretDiscoveryServiceStreamSecretsServer{stream})
}

type SecretDiscoveryService_StreamSecretsServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceStreamSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceStreamSecretsServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_FetchSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.SecretDiscoveryService/FetchSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.SecretDiscoveryService",
	HandlerType: (*SecretDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchSecrets",
			Handler:    _SecretDiscoveryService_FetchSecrets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSecrets",
			Handler:       _SecretDiscoveryService_StreamSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/sds.proto",
}

func init() { proto.RegisterFile("api/sds.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 1175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x76, 0xdb, 0x44,
	0x10, 0xf6, 0x5f, 0x43, 0x33, 0x76, 0x52, 0x67, 0xdb, 0x82, 0x30, 0xa1, 0x0d, 0x82, 0x1e, 0x72,
	0x4a, 0x8f, 0xd3, 0xb8, 0xb4, 0x9c, 0x16, 0x6e, 0x1a, 0x07, 0x9a, 0x43, 0x0b, 0x2d, 0x92, 0x9b,
	0x03, 0xdc, 0x88, 0x8d, 0x3c, 0x89, 0x17, 0xcb, 0x92, 0xaa, 0x59, 0xab, 0xf5, 0x2d, 0xaf, 0x00,
	0xaf, 0xc1, 0x53, 0xf0, 0x08, 0x3c, 0x40, 0x2f, 0xe0, 0x29, 0xb8, 0xe2, 0xec, 0x6a, 0x6d, 0xcb,
	0x76, 0xe2, 0x9c, 0xc3, 0xcf, 0x9d, 0xe6, 0xef, 0x9b, 0xd9, 0x9d, 0x99, 0x6f, 0x05, 0x6b, 0x3c,
	0x16, 0x3b, 0xd4, 0xa5, 0x66, 0x9c, 0x44, 0x32, 0x62, 0x35, 0x0c, 0xd3, 0x68, 0xd4, 0xe4, 0xb1,
	0x68, 0xa6, 0xad, 0xc6, 0xba, 0x32, 0x1e, 0x71, 0xc2, 0xcc, 0xda, 0xb8, 0xac, 0xe4, 0xae, 0x20,
	0x3f, 0x4a, 0x31, 0x19, 0x19, 0xe5, 0xe6, 0x49, 0x14, 0x9d, 0x04, 0xb8, 0xa3, 0x6c, 0x3c, 0x0c,
	0x23, 0xc9, 0xa5, 0x88, 0x42, 0x03, 0xd8, 0xb8, 0x66, 0xac, 0x5a, 0x3a, 0x1a, 0x1e, 0xef, 0xbc,
	0x4c, 0x78, 0x1c, 0x63, 0x62, 0xec, 0xf6, 0x37, 0x00, 0xfb, 0x5c, 0x72, 0x37, 0x1a, 0x26, 0x3e,
	0xb2, 0x4d, 0xb8, 0x78, 0x2c, 0x02, 0x0c, 0xf9, 0x00, 0xad, 0xe2, 0x56, 0x71, 0x7b, 0xf5, 0xa0,
	0xe0, 0x4c, 0x34, 0xcc, 0x82, 0x15, 0x11, 0x06, 0x22, 0x44, 0xab, 0xb4, 0x55, 0xdc, 0xae, 0x1d,
	0x14, 0x1c, 0x23, 0xef, 0x55, 0x61, 0x95, 0x62, 0xf4, 0xc5, 0xb1, 0xc0, 0xc4, 0xfe, 0xab, 0x04,
	0x6b, 0x9d, 0x80, 0x9e, 0xf1, 0x84, 0x0f, 0x50, 0x62, 0x42, 0xac, 0x07, 0x9b, 0x32, 0x20, 0x6f,
	0x20, 0x42, 0x31, 0x18, 0x0e, 0x3c, 0x9d, 0xd9, 0x8f, 0x02, 0x2f, 0xc5, 0x84, 0x44, 0x14, 0xea,
	0x54, 0xeb, 0xad, 0x0f, 0x9b, 0xf9, 0xc3, 0x37, 0x67, 0x20, 0xb4, 0x64, 0xc2, 0x9c, 0xb7, 0x65,
	0x40, 0x5f, 0x65, 0x58, 0x63, 0xdd, 0x61, 0x86, 0x34, 0xc9, 0xc4, 0x5f, 0x9d, 0x9e, 0xa9, 0xf4,
	0x0f, 0x32, 0x65, 0x58, 0xf3, 0x99, 0xde, 0x87, 0x35, 0x5f, 0xc4, 0x3d, 0x4c, 0x3c, 0x1a, 0x0a,
	0x89, 0x64, 0x95, 0xb7, 0xca, 0xdb, 0xab, 0x4e, 0x2d, 0x53, 0xba, 0x5a, 0xc7, 0xae, 0x43, 0x15,
	0xfd, 0x6e, 0xcf, 0xf3, 0x87, 0x49, 0x8a, 0x64, 0x55, 0xb4, 0x0b, 0x28, 0x55, 0x5b, 0x6b, 0xec,
	0xa7, 0x50, 0xcd, 0xe5, 0x63, 0x35, 0xb8, 0xd8, 0x79, 0xe2, 0x7a, 0x0f, 0x9f, 0x77, 0x9e, 0xd6,
	0x0b, 0xac, 0x0a, 0x6f, 0x74, 0x9e, 0xb8, 0xe9, 0xae, 0x77, 0xbb, 0x5e, 0x9c, 0x0a, 0xbb, 0xf5,
	0xd2, 0x54, 0x68, 0xd5, 0xcb, 0x53, 0xe1, 0x4e, 0xbd, 0x62, 0xbf, 0x2e, 0xc1, 0x7a, 0x27, 0xa0,
	0x36, 0x26, 0x52, 0x1c, 0x0b, 0x9f, 0x4b, 0x64, 0x9f, 0xc3, 0x86, 0x3f, 0x15, 0x3d, 0xbf, 0xc7,
	0x45, 0x76, 0xe5, 0xd5, 0x96, 0x35, 0x7b, 0x11, 0xd3, 0x49, 0x70, 0xea, 0xb9, 0x90, 0xb6, 0x8a,
	0x60, 0xf7, 0xa1, 0x1a, 0x27, 0x22, 0x55, 0x10, 0x7d, 0x1c, 0xe9, 0x9b, 0x5c, 0x06, 0x00, 0xc6,
	0xf9, 0x31, 0x8e, 0xd8, 0xc7, 0x70, 0x31, 0xe6, 0x44, 0x2f, 0xa3, 0xa4, 0x6b, 0x95, 0xcf, 0x89,
	0x9b, 0x78, 0xaa, 0x84, 0x91, 0x4f, 0xb1, 0x47, 0x92, 0xc7, 0x01, 0x5a, 0x95, 0xf3, 0x12, 0x2a,
	0x67, 0x57, 0xfb, 0xb2, 0xef, 0x61, 0x93, 0xc4, 0x49, 0x88, 0x5d, 0x2f, 0x7f, 0x72, 0x29, 0x06,
	0x48, 0x92, 0x0f, 0x62, 0xeb, 0xc2, 0x56, 0x79, 0x29, 0x56, 0x23, 0x8b, 0xce, 0xdd, 0x62, 0x67,
	0x1c, 0x6b, 0xef, 0xc3, 0x95, 0x4e, 0x40, 0x2e, 0x92, 0x1a, 0x83, 0x8e, 0xf0, 0xfb, 0x28, 0x1f,
	0xe3, 0x88, 0xd8, 0x2d, 0xa8, 0xf4, 0x71, 0x44, 0x56, 0xf1, 0x1c, 0x6c, 0xed, 0x65, 0xff, 0x5a,
	0x86, 0xcd, 0x1c, 0xfc, 0x21, 0x0f, 0x44, 0x57, 0x2f, 0x6e, 0x3b, 0x0a, 0x25, 0xbe, 0x92, 0xec,
	0x13, 0x00, 0x99, 0x0c, 0x49, 0xaa, 0x33, 0xf0, 0x73, 0xdb, 0xb5, 0x6a, 0x7c, 0xdb, 0x9c, 0xdd,
	0x83, 0xb7, 0x52, 0x4c, 0xc4, 0xf1, 0x68, 0xe6, 0xec, 0x3d, 0x4e, 0x3d, 0xab, 0xa4, 0xe7, 0xef,
	0x6a, 0x66, 0xce, 0x65, 0x3f, 0xe0, 0xd4, 0x63, 0xb7, 0x80, 0x99, 0x38, 0x8a, 0xfb, 0xc2, 0xa3,
	0x1e, 0x6f, 0xdd, 0xbd, 0x67, 0xa6, 0xba, 0x9e, 0x59, 0xdc, 0xb8, 0x2f, 0x5c, 0xad, 0x67, 0x77,
	0x27, 0x59, 0x68, 0x78, 0xf4, 0x23, 0xfa, 0xd2, 0xe3, 0x81, 0xf4, 0x34, 0x71, 0x64, 0x53, 0x7e,
	0xc5, 0x84, 0x64, 0xd6, 0x87, 0x81, 0xfc, 0x5a, 0x51, 0xc8, 0x97, 0x70, 0x39, 0xc1, 0x17, 0x43,
	0x91, 0xa0, 0x97, 0xef, 0xed, 0x05, 0x7d, 0xbc, 0x46, 0x33, 0x23, 0xab, 0xe6, 0x98, 0xac, 0x9a,
	0x7b, 0x51, 0x14, 0x1c, 0xf2, 0x60, 0x88, 0xce, 0x86, 0x09, 0x7b, 0x3a, 0x6d, 0x72, 0x1f, 0x3e,
	0x18, 0x63, 0x2d, 0x6d, 0xf6, 0xca, 0xb9, 0xe0, 0xef, 0x19, 0x1c, 0xf7, 0xec, 0xae, 0xff, 0x52,
	0x81, 0x7a, 0x3b, 0x1a, 0x0c, 0xa2, 0x50, 0x6d, 0x97, 0xe9, 0xd1, 0x03, 0x00, 0xc5, 0x36, 0xb1,
	0x22, 0x0f, 0x32, 0x3d, 0x7a, 0x67, 0x09, 0xb7, 0x38, 0xab, 0xd2, 0x88, 0xc4, 0x1e, 0x41, 0x5d,
	0xc5, 0xe6, 0x4a, 0x26, 0xdd, 0x9f, 0x6a, 0x6b, 0x73, 0x01, 0x21, 0x57, 0x91, 0x73, 0x49, 0xce,
	0xc8, 0xc4, 0x04, 0xd8, 0x73, 0x40, 0x1e, 0x75, 0xc9, 0x23, 0xf4, 0x13, 0x94, 0x9e, 0x1f, 0x85,
	0xc7, 0xe2, 0x84, 0xac, 0x15, 0x0d, 0xfd, 0xee, 0x2c, 0xb4, 0xdb, 0x25, 0x57, 0xbb, 0xb5, 0xb5,
	0x97, 0x73, 0x6d, 0x16, 0x7b, 0xce, 0x4c, 0xec, 0x3b, 0x60, 0xe9, 0x64, 0x50, 0x15, 0xb4, 0xba,
	0x05, 0xb3, 0xd1, 0x37, 0x67, 0xa1, 0x97, 0xcd, 0xb6, 0xb3, 0x91, 0x2e, 0x8c, 0xfb, 0x0d, 0x58,
	0xe7, 0x41, 0x1c, 0x4e, 0x18, 0x7b, 0x4c, 0x96, 0x6b, 0x4a, 0x3b, 0xe6, 0x47, 0x62, 0xcf, 0x60,
	0xad, 0x8b, 0x71, 0x82, 0x0a, 0xb7, 0xeb, 0xa5, 0xbb, 0x66, 0x72, 0x3e, 0x9a, 0x4b, 0x3e, 0xd7,
	0xa8, 0xe6, 0xfe, 0x24, 0xe6, 0x70, 0xd7, 0xa9, 0x75, 0x73, 0x52, 0xe3, 0x33, 0xa8, 0xe5, 0xad,
	0x6a, 0x0d, 0xd4, 0x24, 0xcf, 0x15, 0xa3, 0x1f, 0x43, 0xa7, 0xce, 0x03, 0xf9, 0x30, 0x5f, 0x8f,
	0x2d, 0x81, 0x3d, 0x8f, 0x49, 0x26, 0xc8, 0x07, 0xb9, 0xb9, 0x78, 0x02, 0xcc, 0xd7, 0x25, 0x78,
	0xba, 0x33, 0xe6, 0x9e, 0xb2, 0xf9, 0xb8, 0xb6, 0xbc, 0x54, 0xa7, 0xee, 0xcf, 0x4f, 0x59, 0x1d,
	0xca, 0x14, 0x0a, 0x4d, 0xb8, 0xab, 0x8e, 0xfa, 0xb4, 0x5f, 0x97, 0xe1, 0xca, 0x7e, 0xf4, 0x32,
	0xfc, 0x9f, 0x13, 0x7f, 0x0b, 0x8d, 0xf1, 0x82, 0xf9, 0x81, 0xc0, 0x50, 0xe6, 0x87, 0xcc, 0x3c,
	0x00, 0xcb, 0xd6, 0xca, 0x32, 0xd1, 0x6d, 0x1d, 0x9c, 0x7f, 0x92, 0x3e, 0x85, 0xea, 0x64, 0x75,
	0x43, 0x61, 0x26, 0x68, 0x19, 0x14, 0x8c, 0x37, 0x34, 0x14, 0xac, 0x03, 0x97, 0x29, 0x63, 0x5f,
	0x4f, 0x6a, 0xfa, 0xf5, 0x34, 0xef, 0x66, 0xef, 0x83, 0xbd, 0xb0, 0x3c, 0x0b, 0x4c, 0x7d, 0x50,
	0x70, 0x36, 0x68, 0x81, 0xbe, 0x63, 0xb8, 0x71, 0x0a, 0xea, 0xe2, 0x2a, 0x99, 0x89, 0x5b, 0xbe,
	0x49, 0x07, 0x05, 0x67, 0x6b, 0x21, 0xc5, 0x9c, 0xcf, 0x5e, 0x03, 0xac, 0xd3, 0x32, 0xca, 0x51,
	0x8c, 0xf6, 0x0f, 0x70, 0x69, 0xce, 0x9d, 0x31, 0xa8, 0x4c, 0xff, 0xcb, 0x1c, 0xfd, 0xcd, 0xee,
	0x03, 0xa8, 0x02, 0x4d, 0x65, 0xe3, 0x8e, 0xcc, 0xf5, 0x59, 0xd9, 0xc6, 0xcf, 0x04, 0x75, 0x29,
	0x53, 0xd8, 0xbf, 0x15, 0x61, 0x25, 0xc3, 0x3f, 0x15, 0xf9, 0x11, 0x5c, 0x9a, 0x63, 0x15, 0x03,
	0xbf, 0x94, 0x9d, 0x0e, 0x0a, 0xce, 0xfa, 0x2c, 0x87, 0x9c, 0xd5, 0xad, 0xf2, 0xbf, 0xea, 0xd6,
	0xde, 0x0a, 0x54, 0xd4, 0x3d, 0xb5, 0xfe, 0x28, 0xc2, 0x9b, 0xd9, 0x29, 0xf6, 0xc7, 0xbf, 0xc5,
	0x2e, 0x26, 0xa9, 0xf0, 0x91, 0x1d, 0xc2, 0x9a, 0xab, 0xf7, 0x23, 0xb3, 0x13, 0x9b, 0x5b, 0x80,
	0x49, 0x80, 0x83, 0x2f, 0x86, 0x48, 0xb2, 0x71, 0xfd, 0x4c, 0x3b, 0xc5, 0x51, 0x48, 0x68, 0x17,
	0xb6, 0x8b, 0xb7, 0x8b, 0xec, 0x05, 0xd4, 0xbe, 0x40, 0xe9, 0xf7, 0xfe, 0x33, 0xd8, 0xad, 0x9f,
	0x7e, 0xff, 0xf3, 0xe7, 0x52, 0xc3, 0xbe, 0xba, 0x93, 0xb6, 0xa6, 0x7f, 0xf7, 0x0f, 0xb2, 0xa9,
	0xa3, 0x07, 0xc5, 0x9b, 0x47, 0x2b, 0x7a, 0x23, 0xee, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x4d,
	0xe3, 0xb1, 0x66, 0x2d, 0x0c, 0x00, 0x00,
}
