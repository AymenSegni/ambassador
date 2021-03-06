// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/grpc_web/v2/grpc_web.proto

package envoy_config_filter_http_grpc_web_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

// gRPC Web filter config.
type GrpcWeb struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcWeb) Reset()         { *m = GrpcWeb{} }
func (m *GrpcWeb) String() string { return proto.CompactTextString(m) }
func (*GrpcWeb) ProtoMessage()    {}
func (*GrpcWeb) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fe810f98c50c033, []int{0}
}
func (m *GrpcWeb) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrpcWeb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrpcWeb.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GrpcWeb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcWeb.Merge(m, src)
}
func (m *GrpcWeb) XXX_Size() int {
	return m.Size()
}
func (m *GrpcWeb) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcWeb.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcWeb proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GrpcWeb)(nil), "envoy.config.filter.http.grpc_web.v2.GrpcWeb")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/grpc_web/v2/grpc_web.proto", fileDescriptor_8fe810f98c50c033)
}

var fileDescriptor_8fe810f98c50c033 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x2f, 0x4f, 0x4d, 0xd2, 0x2f, 0x33,
	0x82, 0xb3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0xc0, 0x9a, 0xf4, 0x20, 0x9a, 0xf4,
	0x20, 0x9a, 0xf4, 0x40, 0x9a, 0xf4, 0xe0, 0x0a, 0xcb, 0x8c, 0xa4, 0xe4, 0x4a, 0x53, 0x0a, 0x12,
	0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0xf5, 0x73, 0x33, 0xd3,
	0x8b, 0x12, 0x4b, 0x52, 0x21, 0xa6, 0x48, 0xc9, 0x62, 0xc8, 0x17, 0x97, 0x24, 0x96, 0x94, 0x16,
	0x43, 0xa4, 0x95, 0x38, 0xb9, 0xd8, 0xdd, 0x8b, 0x0a, 0x92, 0xc3, 0x53, 0x93, 0x9c, 0x3a, 0x19,
	0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x4f, 0x33, 0xfe,
	0xf5, 0xb3, 0x6a, 0x0b, 0x69, 0x42, 0x1c, 0x91, 0x5a, 0x51, 0x92, 0x9a, 0x57, 0x0c, 0xd2, 0x0e,
	0x75, 0x48, 0x31, 0xba, 0x4b, 0x8c, 0x77, 0x35, 0x9c, 0xb8, 0xc8, 0xc6, 0x24, 0xc0, 0xc0, 0x65,
	0x94, 0x99, 0xaf, 0x07, 0xd6, 0x55, 0x50, 0x94, 0x5f, 0x51, 0xa9, 0x47, 0x8c, 0x2f, 0x9c, 0x78,
	0xa0, 0x8e, 0x08, 0x00, 0x39, 0x2a, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x3a, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x01, 0x30, 0x3d, 0x33, 0x39, 0x01, 0x00, 0x00,
}

func (m *GrpcWeb) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcWeb) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrpcWeb) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintGrpcWeb(dAtA []byte, offset int, v uint64) int {
	offset -= sovGrpcWeb(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GrpcWeb) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGrpcWeb(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGrpcWeb(x uint64) (n int) {
	return sovGrpcWeb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GrpcWeb) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpcWeb
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
			return fmt.Errorf("proto: GrpcWeb: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrpcWeb: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGrpcWeb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpcWeb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpcWeb
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
func skipGrpcWeb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGrpcWeb
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
					return 0, ErrIntOverflowGrpcWeb
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
					return 0, ErrIntOverflowGrpcWeb
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
				return 0, ErrInvalidLengthGrpcWeb
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGrpcWeb
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGrpcWeb
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGrpcWeb        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGrpcWeb          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGrpcWeb = fmt.Errorf("proto: unexpected end of group")
)
