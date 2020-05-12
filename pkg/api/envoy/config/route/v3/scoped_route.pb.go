// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/route/v3/scoped_route.proto

package envoy_config_route_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

// Specifies a routing scope, which associates a
// :ref:`Key<envoy_api_msg_config.route.v3.ScopedRouteConfiguration.Key>` to a
// :ref:`envoy_api_msg_config.route.v3.RouteConfiguration` (identified by its resource name).
//
// The HTTP connection manager builds up a table consisting of these Key to
// RouteConfiguration mappings, and looks up the RouteConfiguration to use per
// request according to the algorithm specified in the
// :ref:`scope_key_builder<envoy_api_field_extensions.filters.network.http_connection_manager.v3.ScopedRoutes.scope_key_builder>`
// assigned to the HttpConnectionManager.
//
// For example, with the following configurations (in YAML):
//
// HttpConnectionManager config:
//
// .. code::
//
//   ...
//   scoped_routes:
//     name: foo-scoped-routes
//     scope_key_builder:
//       fragments:
//         - header_value_extractor:
//             name: X-Route-Selector
//             element_separator: ,
//             element:
//               separator: =
//               key: vip
//
// ScopedRouteConfiguration resources (specified statically via
// :ref:`scoped_route_configurations_list<envoy_api_field_extensions.filters.network.http_connection_manager.v3.ScopedRoutes.scoped_route_configurations_list>`
// or obtained dynamically via SRDS):
//
// .. code::
//
//  (1)
//   name: route-scope1
//   route_configuration_name: route-config1
//   key:
//      fragments:
//        - string_key: 172.10.10.20
//
//  (2)
//   name: route-scope2
//   route_configuration_name: route-config2
//   key:
//     fragments:
//       - string_key: 172.20.20.30
//
// A request from a client such as:
//
// .. code::
//
//     GET / HTTP/1.1
//     Host: foo.com
//     X-Route-Selector: vip=172.10.10.20
//
// would result in the routing table defined by the `route-config1`
// RouteConfiguration being assigned to the HTTP request/stream.
//
type ScopedRouteConfiguration struct {
	// The name assigned to the routing scope.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource name to use for a :ref:`envoy_api_msg_service.discovery.v3.DiscoveryRequest` to an
	// RDS server to fetch the :ref:`envoy_api_msg_config.route.v3.RouteConfiguration` associated
	// with this scope.
	RouteConfigurationName string `protobuf:"bytes,2,opt,name=route_configuration_name,json=routeConfigurationName,proto3" json:"route_configuration_name,omitempty"`
	// The key to match against.
	Key                  *ScopedRouteConfiguration_Key `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ScopedRouteConfiguration) Reset()         { *m = ScopedRouteConfiguration{} }
func (m *ScopedRouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*ScopedRouteConfiguration) ProtoMessage()    {}
func (*ScopedRouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_3339c04c2fb0ae6d, []int{0}
}
func (m *ScopedRouteConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScopedRouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScopedRouteConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScopedRouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScopedRouteConfiguration.Merge(m, src)
}
func (m *ScopedRouteConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *ScopedRouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ScopedRouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ScopedRouteConfiguration proto.InternalMessageInfo

func (m *ScopedRouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScopedRouteConfiguration) GetRouteConfigurationName() string {
	if m != nil {
		return m.RouteConfigurationName
	}
	return ""
}

func (m *ScopedRouteConfiguration) GetKey() *ScopedRouteConfiguration_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

// Specifies a key which is matched against the output of the
// :ref:`scope_key_builder<envoy_api_field_extensions.filters.network.http_connection_manager.v3.ScopedRoutes.scope_key_builder>`
// specified in the HttpConnectionManager. The matching is done per HTTP
// request and is dependent on the order of the fragments contained in the
// Key.
type ScopedRouteConfiguration_Key struct {
	// The ordered set of fragments to match against. The order must match the
	// fragments in the corresponding
	// :ref:`scope_key_builder<envoy_api_field_extensions.filters.network.http_connection_manager.v3.ScopedRoutes.scope_key_builder>`.
	Fragments            []*ScopedRouteConfiguration_Key_Fragment `protobuf:"bytes,1,rep,name=fragments,proto3" json:"fragments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *ScopedRouteConfiguration_Key) Reset()         { *m = ScopedRouteConfiguration_Key{} }
func (m *ScopedRouteConfiguration_Key) String() string { return proto.CompactTextString(m) }
func (*ScopedRouteConfiguration_Key) ProtoMessage()    {}
func (*ScopedRouteConfiguration_Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_3339c04c2fb0ae6d, []int{0, 0}
}
func (m *ScopedRouteConfiguration_Key) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScopedRouteConfiguration_Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScopedRouteConfiguration_Key.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScopedRouteConfiguration_Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScopedRouteConfiguration_Key.Merge(m, src)
}
func (m *ScopedRouteConfiguration_Key) XXX_Size() int {
	return m.Size()
}
func (m *ScopedRouteConfiguration_Key) XXX_DiscardUnknown() {
	xxx_messageInfo_ScopedRouteConfiguration_Key.DiscardUnknown(m)
}

var xxx_messageInfo_ScopedRouteConfiguration_Key proto.InternalMessageInfo

func (m *ScopedRouteConfiguration_Key) GetFragments() []*ScopedRouteConfiguration_Key_Fragment {
	if m != nil {
		return m.Fragments
	}
	return nil
}

type ScopedRouteConfiguration_Key_Fragment struct {
	// Types that are valid to be assigned to Type:
	//	*ScopedRouteConfiguration_Key_Fragment_StringKey
	Type                 isScopedRouteConfiguration_Key_Fragment_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *ScopedRouteConfiguration_Key_Fragment) Reset()         { *m = ScopedRouteConfiguration_Key_Fragment{} }
func (m *ScopedRouteConfiguration_Key_Fragment) String() string { return proto.CompactTextString(m) }
func (*ScopedRouteConfiguration_Key_Fragment) ProtoMessage()    {}
func (*ScopedRouteConfiguration_Key_Fragment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3339c04c2fb0ae6d, []int{0, 0, 0}
}
func (m *ScopedRouteConfiguration_Key_Fragment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScopedRouteConfiguration_Key_Fragment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScopedRouteConfiguration_Key_Fragment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScopedRouteConfiguration_Key_Fragment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScopedRouteConfiguration_Key_Fragment.Merge(m, src)
}
func (m *ScopedRouteConfiguration_Key_Fragment) XXX_Size() int {
	return m.Size()
}
func (m *ScopedRouteConfiguration_Key_Fragment) XXX_DiscardUnknown() {
	xxx_messageInfo_ScopedRouteConfiguration_Key_Fragment.DiscardUnknown(m)
}

var xxx_messageInfo_ScopedRouteConfiguration_Key_Fragment proto.InternalMessageInfo

type isScopedRouteConfiguration_Key_Fragment_Type interface {
	isScopedRouteConfiguration_Key_Fragment_Type()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ScopedRouteConfiguration_Key_Fragment_StringKey struct {
	StringKey string `protobuf:"bytes,1,opt,name=string_key,json=stringKey,proto3,oneof" json:"string_key,omitempty"`
}

func (*ScopedRouteConfiguration_Key_Fragment_StringKey) isScopedRouteConfiguration_Key_Fragment_Type() {
}

func (m *ScopedRouteConfiguration_Key_Fragment) GetType() isScopedRouteConfiguration_Key_Fragment_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ScopedRouteConfiguration_Key_Fragment) GetStringKey() string {
	if x, ok := m.GetType().(*ScopedRouteConfiguration_Key_Fragment_StringKey); ok {
		return x.StringKey
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ScopedRouteConfiguration_Key_Fragment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ScopedRouteConfiguration_Key_Fragment_StringKey)(nil),
	}
}

func init() {
	proto.RegisterType((*ScopedRouteConfiguration)(nil), "envoy.config.route.v3.ScopedRouteConfiguration")
	proto.RegisterType((*ScopedRouteConfiguration_Key)(nil), "envoy.config.route.v3.ScopedRouteConfiguration.Key")
	proto.RegisterType((*ScopedRouteConfiguration_Key_Fragment)(nil), "envoy.config.route.v3.ScopedRouteConfiguration.Key.Fragment")
}

func init() {
	proto.RegisterFile("envoy/config/route/v3/scoped_route.proto", fileDescriptor_3339c04c2fb0ae6d)
}

var fileDescriptor_3339c04c2fb0ae6d = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0xeb, 0xd3, 0x30,
	0x18, 0xc6, 0x4d, 0x5b, 0xff, 0x6e, 0xd9, 0x65, 0x14, 0xd4, 0x52, 0xb1, 0x56, 0x45, 0xac, 0x28,
	0xa9, 0xb4, 0x5e, 0x1c, 0x5e, 0xac, 0x20, 0xc2, 0x50, 0x47, 0xfd, 0x00, 0x25, 0xae, 0x59, 0x09,
	0xba, 0xa4, 0x24, 0x69, 0xb1, 0x37, 0x8f, 0xe2, 0xc1, 0xc3, 0x8e, 0x7e, 0x14, 0xef, 0xc2, 0x8e,
	0xfa, 0x0d, 0x64, 0x9f, 0x42, 0x76, 0x92, 0xa4, 0x13, 0x27, 0x9b, 0x30, 0xbc, 0xb5, 0x79, 0x9f,
	0xe7, 0xf7, 0x3c, 0x79, 0x09, 0x8c, 0x08, 0x6b, 0x79, 0x17, 0xcf, 0x39, 0x5b, 0xd0, 0x2a, 0x16,
	0xbc, 0x51, 0x24, 0x6e, 0xd3, 0x58, 0xce, 0x79, 0x4d, 0xca, 0xc2, 0xfc, 0xa3, 0x5a, 0x70, 0xc5,
	0xdd, 0x8b, 0x46, 0x89, 0x7a, 0x25, 0xea, 0x27, 0x6d, 0xea, 0x5f, 0x6d, 0xca, 0x1a, 0xc7, 0x98,
	0x31, 0xae, 0xb0, 0xa2, 0x9c, 0xc9, 0x58, 0x2a, 0xac, 0x1a, 0xd9, 0xbb, 0xfc, 0xeb, 0x07, 0xe3,
	0x96, 0x08, 0x49, 0x39, 0xa3, 0xac, 0xda, 0x49, 0x2e, 0xb7, 0xf8, 0x2d, 0x2d, 0xb1, 0x0e, 0xde,
	0x7d, 0xf4, 0x83, 0x1b, 0x9f, 0x1c, 0xe8, 0xbd, 0x32, 0x45, 0x72, 0x9d, 0xf6, 0xc4, 0x24, 0x37,
	0xc2, 0x70, 0xdc, 0x2b, 0xd0, 0x61, 0x78, 0x49, 0x3c, 0x10, 0x82, 0x68, 0x98, 0x5d, 0xd8, 0x66,
	0x8e, 0xb0, 0x42, 0x90, 0x9b, 0x43, 0xf7, 0x31, 0xf4, 0x4c, 0xc1, 0x62, 0xbe, 0xef, 0x29, 0x8c,
	0xc1, 0xfa, 0xdb, 0x70, 0x49, 0x1c, 0xb0, 0x5f, 0x68, 0xc4, 0x4b, 0x68, 0xbf, 0x21, 0x9d, 0x67,
	0x87, 0x20, 0x1a, 0x25, 0x29, 0x3a, 0x7a, 0x79, 0xf4, 0xaf, 0x76, 0x68, 0x4a, 0xba, 0x6c, 0xb0,
	0xcd, 0xce, 0x7f, 0x04, 0xd6, 0x18, 0xe4, 0x9a, 0xe4, 0xaf, 0x2c, 0x68, 0x4f, 0x49, 0xe7, 0x96,
	0x70, 0xb8, 0x10, 0xb8, 0x5a, 0x12, 0xa6, 0xa4, 0x07, 0x42, 0x3b, 0x1a, 0x25, 0x8f, 0xfe, 0x03,
	0x8f, 0x9e, 0xee, 0x20, 0x26, 0x67, 0x05, 0xac, 0x01, 0xc8, 0xff, 0x80, 0x7d, 0x09, 0x07, 0xbf,
	0x05, 0xee, 0x35, 0x08, 0xa5, 0x12, 0x94, 0x55, 0x85, 0xbe, 0x91, 0x59, 0xd8, 0xb3, 0x73, 0xf9,
	0xb0, 0x3f, 0x9b, 0x92, 0x6e, 0xf2, 0xf0, 0xf3, 0xd7, 0x0f, 0xc1, 0x03, 0x98, 0xf4, 0x2d, 0x70,
	0x4d, 0x51, 0x9b, 0x9c, 0x18, 0x3e, 0x82, 0x8e, 0xea, 0x6a, 0xe2, 0xda, 0x3f, 0x33, 0x30, 0xb9,
	0xaf, 0x39, 0x77, 0xe1, 0x9d, 0x93, 0x39, 0x93, 0x7b, 0xda, 0x71, 0x1b, 0xde, 0x3a, 0xc9, 0x91,
	0x3d, 0x5f, 0x6f, 0x02, 0xf0, 0x6d, 0x13, 0x80, 0x1f, 0x9b, 0x00, 0x7c, 0x79, 0xbf, 0xfe, 0x7e,
	0x66, 0x8d, 0x01, 0xbc, 0x49, 0x79, 0xbf, 0xbf, 0x5a, 0xf0, 0x77, 0xdd, 0xf1, 0x55, 0x66, 0xe3,
	0x3d, 0xe8, 0x4c, 0xbf, 0xae, 0x19, 0x78, 0x7d, 0x66, 0x9e, 0x59, 0xfa, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x1b, 0x53, 0x03, 0xf8, 0x04, 0x03, 0x00, 0x00,
}

func (m *ScopedRouteConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScopedRouteConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScopedRouteConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Key != nil {
		{
			size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintScopedRoute(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RouteConfigurationName) > 0 {
		i -= len(m.RouteConfigurationName)
		copy(dAtA[i:], m.RouteConfigurationName)
		i = encodeVarintScopedRoute(dAtA, i, uint64(len(m.RouteConfigurationName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintScopedRoute(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ScopedRouteConfiguration_Key) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScopedRouteConfiguration_Key) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScopedRouteConfiguration_Key) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Fragments) > 0 {
		for iNdEx := len(m.Fragments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fragments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintScopedRoute(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ScopedRouteConfiguration_Key_Fragment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScopedRouteConfiguration_Key_Fragment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScopedRouteConfiguration_Key_Fragment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Type != nil {
		{
			size := m.Type.Size()
			i -= size
			if _, err := m.Type.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ScopedRouteConfiguration_Key_Fragment_StringKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScopedRouteConfiguration_Key_Fragment_StringKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.StringKey)
	copy(dAtA[i:], m.StringKey)
	i = encodeVarintScopedRoute(dAtA, i, uint64(len(m.StringKey)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func encodeVarintScopedRoute(dAtA []byte, offset int, v uint64) int {
	offset -= sovScopedRoute(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ScopedRouteConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovScopedRoute(uint64(l))
	}
	l = len(m.RouteConfigurationName)
	if l > 0 {
		n += 1 + l + sovScopedRoute(uint64(l))
	}
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovScopedRoute(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ScopedRouteConfiguration_Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Fragments) > 0 {
		for _, e := range m.Fragments {
			l = e.Size()
			n += 1 + l + sovScopedRoute(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ScopedRouteConfiguration_Key_Fragment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != nil {
		n += m.Type.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ScopedRouteConfiguration_Key_Fragment_StringKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StringKey)
	n += 1 + l + sovScopedRoute(uint64(l))
	return n
}

func sovScopedRoute(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScopedRoute(x uint64) (n int) {
	return sovScopedRoute(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ScopedRouteConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScopedRoute
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
			return fmt.Errorf("proto: ScopedRouteConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScopedRouteConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScopedRoute
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
				return ErrInvalidLengthScopedRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScopedRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteConfigurationName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScopedRoute
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
				return ErrInvalidLengthScopedRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScopedRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RouteConfigurationName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScopedRoute
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
				return ErrInvalidLengthScopedRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScopedRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Key == nil {
				m.Key = &ScopedRouteConfiguration_Key{}
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScopedRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScopedRoute
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScopedRoute
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
func (m *ScopedRouteConfiguration_Key) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScopedRoute
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
			return fmt.Errorf("proto: Key: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Key: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fragments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScopedRoute
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
				return ErrInvalidLengthScopedRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScopedRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fragments = append(m.Fragments, &ScopedRouteConfiguration_Key_Fragment{})
			if err := m.Fragments[len(m.Fragments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScopedRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScopedRoute
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScopedRoute
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
func (m *ScopedRouteConfiguration_Key_Fragment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScopedRoute
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
			return fmt.Errorf("proto: Fragment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fragment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScopedRoute
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
				return ErrInvalidLengthScopedRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScopedRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = &ScopedRouteConfiguration_Key_Fragment_StringKey{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScopedRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScopedRoute
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScopedRoute
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
func skipScopedRoute(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScopedRoute
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
					return 0, ErrIntOverflowScopedRoute
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
					return 0, ErrIntOverflowScopedRoute
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
				return 0, ErrInvalidLengthScopedRoute
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScopedRoute
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScopedRoute
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScopedRoute        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScopedRoute          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScopedRoute = fmt.Errorf("proto: unexpected end of group")
)
