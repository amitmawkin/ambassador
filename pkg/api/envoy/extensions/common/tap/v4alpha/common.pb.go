// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/common/tap/v4alpha/common.proto

package envoy_extensions_common_tap_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha "github.com/datawire/ambassador/pkg/api/envoy/config/core/v4alpha"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/config/tap/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Common configuration for all tap extensions.
type CommonExtensionConfig struct {
	// Types that are valid to be assigned to ConfigType:
	//	*CommonExtensionConfig_AdminConfig
	//	*CommonExtensionConfig_StaticConfig
	//	*CommonExtensionConfig_TapdsConfig
	ConfigType           isCommonExtensionConfig_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CommonExtensionConfig) Reset()         { *m = CommonExtensionConfig{} }
func (m *CommonExtensionConfig) String() string { return proto.CompactTextString(m) }
func (*CommonExtensionConfig) ProtoMessage()    {}
func (*CommonExtensionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a673673e8648a9b, []int{0}
}
func (m *CommonExtensionConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonExtensionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommonExtensionConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommonExtensionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonExtensionConfig.Merge(m, src)
}
func (m *CommonExtensionConfig) XXX_Size() int {
	return m.Size()
}
func (m *CommonExtensionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonExtensionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonExtensionConfig proto.InternalMessageInfo

type isCommonExtensionConfig_ConfigType interface {
	isCommonExtensionConfig_ConfigType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type CommonExtensionConfig_AdminConfig struct {
	AdminConfig *AdminConfig `protobuf:"bytes,1,opt,name=admin_config,json=adminConfig,proto3,oneof" json:"admin_config,omitempty"`
}
type CommonExtensionConfig_StaticConfig struct {
	StaticConfig *v3.TapConfig `protobuf:"bytes,2,opt,name=static_config,json=staticConfig,proto3,oneof" json:"static_config,omitempty"`
}
type CommonExtensionConfig_TapdsConfig struct {
	TapdsConfig *CommonExtensionConfig_TapDSConfig `protobuf:"bytes,3,opt,name=tapds_config,json=tapdsConfig,proto3,oneof" json:"tapds_config,omitempty"`
}

func (*CommonExtensionConfig_AdminConfig) isCommonExtensionConfig_ConfigType()  {}
func (*CommonExtensionConfig_StaticConfig) isCommonExtensionConfig_ConfigType() {}
func (*CommonExtensionConfig_TapdsConfig) isCommonExtensionConfig_ConfigType()  {}

func (m *CommonExtensionConfig) GetConfigType() isCommonExtensionConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *CommonExtensionConfig) GetAdminConfig() *AdminConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_AdminConfig); ok {
		return x.AdminConfig
	}
	return nil
}

func (m *CommonExtensionConfig) GetStaticConfig() *v3.TapConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_StaticConfig); ok {
		return x.StaticConfig
	}
	return nil
}

func (m *CommonExtensionConfig) GetTapdsConfig() *CommonExtensionConfig_TapDSConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_TapdsConfig); ok {
		return x.TapdsConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CommonExtensionConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CommonExtensionConfig_AdminConfig)(nil),
		(*CommonExtensionConfig_StaticConfig)(nil),
		(*CommonExtensionConfig_TapdsConfig)(nil),
	}
}

// [#not-implemented-hide:]
type CommonExtensionConfig_TapDSConfig struct {
	// Configuration for the source of TapDS updates for this Cluster.
	ConfigSource *v4alpha.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	// Tap config to request from XDS server.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonExtensionConfig_TapDSConfig) Reset()         { *m = CommonExtensionConfig_TapDSConfig{} }
func (m *CommonExtensionConfig_TapDSConfig) String() string { return proto.CompactTextString(m) }
func (*CommonExtensionConfig_TapDSConfig) ProtoMessage()    {}
func (*CommonExtensionConfig_TapDSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a673673e8648a9b, []int{0, 0}
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Merge(m, src)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Size() int {
	return m.Size()
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonExtensionConfig_TapDSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonExtensionConfig_TapDSConfig proto.InternalMessageInfo

func (m *CommonExtensionConfig_TapDSConfig) GetConfigSource() *v4alpha.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func (m *CommonExtensionConfig_TapDSConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Configuration for the admin handler. See :ref:`here <config_http_filters_tap_admin_handler>` for
// more information.
type AdminConfig struct {
	// Opaque configuration ID. When requests are made to the admin handler, the passed opaque ID is
	// matched to the configured filter opaque ID to determine which filter to configure.
	ConfigId             string   `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdminConfig) Reset()         { *m = AdminConfig{} }
func (m *AdminConfig) String() string { return proto.CompactTextString(m) }
func (*AdminConfig) ProtoMessage()    {}
func (*AdminConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a673673e8648a9b, []int{1}
}
func (m *AdminConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AdminConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AdminConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AdminConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminConfig.Merge(m, src)
}
func (m *AdminConfig) XXX_Size() int {
	return m.Size()
}
func (m *AdminConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AdminConfig proto.InternalMessageInfo

func (m *AdminConfig) GetConfigId() string {
	if m != nil {
		return m.ConfigId
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonExtensionConfig)(nil), "envoy.extensions.common.tap.v4alpha.CommonExtensionConfig")
	proto.RegisterType((*CommonExtensionConfig_TapDSConfig)(nil), "envoy.extensions.common.tap.v4alpha.CommonExtensionConfig.TapDSConfig")
	proto.RegisterType((*AdminConfig)(nil), "envoy.extensions.common.tap.v4alpha.AdminConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/common/tap/v4alpha/common.proto", fileDescriptor_8a673673e8648a9b)
}

var fileDescriptor_8a673673e8648a9b = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4d, 0x6b, 0xd4, 0x40,
	0x18, 0xc7, 0x9d, 0x74, 0xad, 0xed, 0x24, 0x85, 0x32, 0x20, 0x4a, 0xc4, 0xb0, 0x56, 0x41, 0x11,
	0x4d, 0x5a, 0x53, 0x3c, 0xac, 0x17, 0x9d, 0x5a, 0x5f, 0x6e, 0x65, 0xab, 0xbd, 0x96, 0x31, 0x99,
	0xd6, 0xc1, 0x66, 0x66, 0x48, 0x66, 0x43, 0xf7, 0xe6, 0x51, 0xfc, 0x08, 0x7e, 0x14, 0x6f, 0x1e,
	0x84, 0x82, 0x17, 0xfd, 0x06, 0xb2, 0x9f, 0x42, 0xf6, 0x24, 0xf3, 0x12, 0x37, 0x41, 0x71, 0x97,
	0xde, 0x32, 0xcf, 0xcb, 0xef, 0xf9, 0xff, 0x9f, 0xc9, 0xc0, 0x4d, 0xca, 0x6b, 0x31, 0x4e, 0xe8,
	0xa9, 0xa2, 0xbc, 0x62, 0x82, 0x57, 0x49, 0x26, 0x8a, 0x42, 0xf0, 0x44, 0x11, 0x99, 0xd4, 0xdb,
	0xe4, 0x44, 0xbe, 0x25, 0x2e, 0x14, 0xcb, 0x52, 0x28, 0x81, 0x6e, 0x9a, 0x8e, 0x78, 0xd6, 0x11,
	0xbb, 0xb4, 0x22, 0x32, 0x76, 0x1d, 0xe1, 0x7d, 0x8b, 0xcd, 0x04, 0x3f, 0x62, 0xc7, 0x49, 0x26,
	0x4a, 0xda, 0x82, 0xe9, 0xd8, 0x61, 0x25, 0x46, 0x65, 0x46, 0x2d, 0x33, 0xec, 0x77, 0xca, 0xcd,
	0xe8, 0xb4, 0x33, 0x35, 0xbc, 0x3e, 0xca, 0x25, 0x49, 0x08, 0xe7, 0x42, 0x11, 0x65, 0x74, 0x56,
	0x8a, 0xa8, 0x51, 0xe5, 0xd2, 0x37, 0xfe, 0x4a, 0xd7, 0xb4, 0xd4, 0xea, 0x18, 0x3f, 0x76, 0x25,
	0x57, 0x6a, 0x72, 0xc2, 0x72, 0xa2, 0x68, 0xd2, 0x7c, 0xd8, 0xc4, 0xc6, 0x97, 0x1e, 0xbc, 0xbc,
	0x63, 0x66, 0xed, 0x36, 0x9e, 0x76, 0x8c, 0x10, 0xf4, 0x1a, 0x06, 0x24, 0x2f, 0x18, 0x3f, 0xb4,
	0xc2, 0xae, 0x82, 0x3e, 0xb8, 0xe3, 0x3f, 0xd8, 0x8c, 0x17, 0xd8, 0x40, 0xfc, 0x44, 0x37, 0x5a,
	0xce, 0x8b, 0x0b, 0x43, 0x9f, 0xcc, 0x8e, 0x68, 0x17, 0xae, 0x69, 0xf1, 0x2c, 0x6b, 0xb8, 0x9e,
	0xe1, 0x46, 0x8e, 0x6b, 0x83, 0x16, 0x96, 0xc6, 0xaf, 0x88, 0xfc, 0x43, 0x09, 0x6c, 0x9b, 0xc3,
	0xbc, 0x83, 0x81, 0x22, 0x32, 0xaf, 0x1a, 0xca, 0x92, 0xa1, 0x3c, 0x5b, 0x48, 0xdd, 0x3f, 0xfd,
	0xea, 0x59, 0x4f, 0xf7, 0x67, 0x9a, 0x0d, 0xdd, 0x1e, 0xc3, 0x6f, 0x00, 0xfa, 0xad, 0x34, 0x3a,
	0x80, 0x6b, 0x9d, 0x8b, 0x74, 0xbb, 0xb9, 0xdd, 0xf5, 0xa0, 0x2f, 0xbe, 0x35, 0x53, 0xc7, 0xf6,
	0x4d, 0x39, 0x5e, 0x99, 0xe2, 0x8b, 0x1f, 0x81, 0xb7, 0x0e, 0x86, 0x41, 0xd6, 0x8a, 0xa3, 0x6b,
	0xb0, 0xc7, 0x49, 0x41, 0xcd, 0x4a, 0x56, 0xf1, 0xa5, 0x29, 0xee, 0x95, 0x5e, 0x1f, 0x0c, 0x4d,
	0x70, 0xf0, 0xfc, 0xd3, 0xd7, 0x0f, 0x11, 0x86, 0x8f, 0xff, 0xeb, 0x30, 0x9d, 0x6f, 0x6e, 0xf0,
	0x48, 0x83, 0x1e, 0xc2, 0xed, 0xf3, 0x80, 0x30, 0x82, 0xbe, 0xb3, 0xae, 0xc6, 0x92, 0xa2, 0xa5,
	0x5f, 0x18, 0x6c, 0x1c, 0x41, 0xbf, 0x75, 0xe1, 0xe8, 0x16, 0x5c, 0x75, 0x25, 0x2c, 0x37, 0x9b,
	0x69, 0x59, 0x59, 0xb1, 0x99, 0x97, 0xf9, 0x60, 0x4b, 0xab, 0xb8, 0x07, 0xef, 0xce, 0x51, 0xd1,
	0x02, 0xe3, 0x83, 0xb3, 0x49, 0x04, 0xbe, 0x4f, 0x22, 0xf0, 0x73, 0x12, 0x81, 0xcf, 0xef, 0xcf,
	0x7e, 0x2c, 0x7b, 0xeb, 0x1e, 0xdc, 0x62, 0xc2, 0xee, 0x5d, 0x96, 0xe2, 0x74, 0xbc, 0xc8, 0x0f,
	0x80, 0x7d, 0xeb, 0x6d, 0x4f, 0xbf, 0x80, 0x3d, 0xf0, 0x66, 0xd9, 0x3c, 0x85, 0xf4, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1b, 0x9f, 0xbb, 0x7b, 0x0f, 0x04, 0x00, 0x00,
}

func (m *CommonExtensionConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonExtensionConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonExtensionConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ConfigType != nil {
		{
			size := m.ConfigType.Size()
			i -= size
			if _, err := m.ConfigType.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *CommonExtensionConfig_AdminConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonExtensionConfig_AdminConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AdminConfig != nil {
		{
			size, err := m.AdminConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *CommonExtensionConfig_StaticConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonExtensionConfig_StaticConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StaticConfig != nil {
		{
			size, err := m.StaticConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *CommonExtensionConfig_TapdsConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonExtensionConfig_TapdsConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.TapdsConfig != nil {
		{
			size, err := m.TapdsConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *CommonExtensionConfig_TapDSConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonExtensionConfig_TapDSConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonExtensionConfig_TapDSConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.ConfigSource != nil {
		{
			size, err := m.ConfigSource.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AdminConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdminConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AdminConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ConfigId) > 0 {
		i -= len(m.ConfigId)
		copy(dAtA[i:], m.ConfigId)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.ConfigId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommonExtensionConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConfigType != nil {
		n += m.ConfigType.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CommonExtensionConfig_AdminConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AdminConfig != nil {
		l = m.AdminConfig.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}
func (m *CommonExtensionConfig_StaticConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StaticConfig != nil {
		l = m.StaticConfig.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}
func (m *CommonExtensionConfig_TapdsConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TapdsConfig != nil {
		l = m.TapdsConfig.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}
func (m *CommonExtensionConfig_TapDSConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConfigSource != nil {
		l = m.ConfigSource.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AdminConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConfigId)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommonExtensionConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommonExtensionConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonExtensionConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AdminConfig{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ConfigType = &CommonExtensionConfig_AdminConfig{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StaticConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v3.TapConfig{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ConfigType = &CommonExtensionConfig_StaticConfig{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TapdsConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &CommonExtensionConfig_TapDSConfig{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ConfigType = &CommonExtensionConfig_TapdsConfig{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CommonExtensionConfig_TapDSConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TapDSConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TapDSConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfigSource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConfigSource == nil {
				m.ConfigSource = &v4alpha.ConfigSource{}
			}
			if err := m.ConfigSource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AdminConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdminConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdminConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfigId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConfigId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)