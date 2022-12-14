// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: buchain/goldoracle/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgSendReqGoldPriceRequest struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Port      string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelID string `protobuf:"bytes,3,opt,name=channelID,proto3" json:"channelID,omitempty"`
	Timeout   uint64 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (m *MsgSendReqGoldPriceRequest) Reset()         { *m = MsgSendReqGoldPriceRequest{} }
func (m *MsgSendReqGoldPriceRequest) String() string { return proto.CompactTextString(m) }
func (*MsgSendReqGoldPriceRequest) ProtoMessage()    {}
func (*MsgSendReqGoldPriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ca5f8bd582fa64a, []int{0}
}
func (m *MsgSendReqGoldPriceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendReqGoldPriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendReqGoldPriceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendReqGoldPriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendReqGoldPriceRequest.Merge(m, src)
}
func (m *MsgSendReqGoldPriceRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendReqGoldPriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendReqGoldPriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendReqGoldPriceRequest proto.InternalMessageInfo

func (m *MsgSendReqGoldPriceRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSendReqGoldPriceRequest) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgSendReqGoldPriceRequest) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *MsgSendReqGoldPriceRequest) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type MsgSendReqGoldPriceResponse struct {
}

func (m *MsgSendReqGoldPriceResponse) Reset()         { *m = MsgSendReqGoldPriceResponse{} }
func (m *MsgSendReqGoldPriceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendReqGoldPriceResponse) ProtoMessage()    {}
func (*MsgSendReqGoldPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ca5f8bd582fa64a, []int{1}
}
func (m *MsgSendReqGoldPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendReqGoldPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendReqGoldPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendReqGoldPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendReqGoldPriceResponse.Merge(m, src)
}
func (m *MsgSendReqGoldPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendReqGoldPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendReqGoldPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendReqGoldPriceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSendReqGoldPriceRequest)(nil), "buchain.goldoracle.MsgSendReqGoldPriceRequest")
	proto.RegisterType((*MsgSendReqGoldPriceResponse)(nil), "buchain.goldoracle.MsgSendReqGoldPriceResponse")
}

func init() { proto.RegisterFile("buchain/goldoracle/tx.proto", fileDescriptor_3ca5f8bd582fa64a) }

var fileDescriptor_3ca5f8bd582fa64a = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x2a, 0x4d, 0xce,
	0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0xcf, 0xcf, 0x49, 0xc9, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x2f,
	0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0x4a, 0xea, 0x21, 0x24, 0x95, 0x9a,
	0x18, 0xb9, 0xa4, 0x7c, 0x8b, 0xd3, 0x83, 0x53, 0xf3, 0x52, 0x82, 0x52, 0x0b, 0xdd, 0xf3, 0x73,
	0x52, 0x02, 0x8a, 0x32, 0x93, 0x53, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8,
	0xd8, 0x93, 0x8b, 0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60,
	0x5c, 0x21, 0x21, 0x2e, 0x96, 0x82, 0xfc, 0xa2, 0x12, 0x09, 0x26, 0xb0, 0x30, 0x98, 0x2d, 0x24,
	0xc3, 0xc5, 0x99, 0x9c, 0x91, 0x98, 0x97, 0x97, 0x9a, 0xe3, 0xe9, 0x22, 0xc1, 0x0c, 0x96, 0x40,
	0x08, 0x80, 0xcc, 0x2a, 0xc9, 0xcc, 0x4d, 0xcd, 0x2f, 0x2d, 0x91, 0x60, 0x51, 0x60, 0xd4, 0x60,
	0x09, 0x82, 0x71, 0x95, 0x64, 0xb9, 0xa4, 0xb1, 0xba, 0xa1, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0xd5,
	0xa8, 0x8a, 0x8b, 0xd9, 0xb7, 0x38, 0x5d, 0xa8, 0x98, 0x4b, 0x00, 0x5d, 0x89, 0x90, 0x9e, 0x1e,
	0xa6, 0x9f, 0xf4, 0x70, 0xfb, 0x47, 0x4a, 0x9f, 0x68, 0xf5, 0x10, 0xbb, 0x9d, 0x4c, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x3a, 0xa9, 0x54, 0x17, 0x12, 0xd6, 0x15,
	0x28, 0xa1, 0x5d, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x71, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe0, 0xa9, 0xc4, 0x76, 0x90, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SendReqGoldPrice(ctx context.Context, in *MsgSendReqGoldPriceRequest, opts ...grpc.CallOption) (*MsgSendReqGoldPriceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendReqGoldPrice(ctx context.Context, in *MsgSendReqGoldPriceRequest, opts ...grpc.CallOption) (*MsgSendReqGoldPriceResponse, error) {
	out := new(MsgSendReqGoldPriceResponse)
	err := c.cc.Invoke(ctx, "/buchain.goldoracle.Msg/SendReqGoldPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SendReqGoldPrice(context.Context, *MsgSendReqGoldPriceRequest) (*MsgSendReqGoldPriceResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendReqGoldPrice(ctx context.Context, req *MsgSendReqGoldPriceRequest) (*MsgSendReqGoldPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReqGoldPrice not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendReqGoldPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendReqGoldPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendReqGoldPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buchain.goldoracle.Msg/SendReqGoldPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendReqGoldPrice(ctx, req.(*MsgSendReqGoldPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buchain.goldoracle.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendReqGoldPrice",
			Handler:    _Msg_SendReqGoldPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buchain/goldoracle/tx.proto",
}

func (m *MsgSendReqGoldPriceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendReqGoldPriceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendReqGoldPriceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timeout != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendReqGoldPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendReqGoldPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendReqGoldPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSendReqGoldPriceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovTx(uint64(m.Timeout))
	}
	return n
}

func (m *MsgSendReqGoldPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSendReqGoldPriceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendReqGoldPriceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendReqGoldPriceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendReqGoldPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendReqGoldPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendReqGoldPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
