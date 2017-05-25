// Code generated by protoc-gen-gogo.
// source: api.proto
// DO NOT EDIT!

/*
	Package protogen is a generated protocol buffer package.

	It is generated from these files:
		api.proto

	It has these top-level messages:
		SessionChunk
		Chunks
*/
package protogen

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SessionChunk is a chunk to be posted in the context of the session
type SessionChunk struct {
	// Namespace is session namespace
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,json=namespace,proto3" json:"Namespace,omitempty"`
	// SessionID is session ID
	SessionID string `protobuf:"bytes,2,opt,name=SessionID,json=sessionID,proto3" json:"SessionID,omitempty"`
	// Chunk is a session chunk
	Chunk []byte `protobuf:"bytes,3,opt,name=Chunk,json=chunk,proto3" json:"Chunk,omitempty"`
}

func (m *SessionChunk) Reset()                    { *m = SessionChunk{} }
func (m *SessionChunk) String() string            { return proto.CompactTextString(m) }
func (*SessionChunk) ProtoMessage()               {}
func (*SessionChunk) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{0} }

type Chunks struct {
	Chunks []*SessionChunk `protobuf:"bytes,1,rep,name=Chunks,json=chunks" json:"Chunks,omitempty"`
}

func (m *Chunks) Reset()                    { *m = Chunks{} }
func (m *Chunks) String() string            { return proto.CompactTextString(m) }
func (*Chunks) ProtoMessage()               {}
func (*Chunks) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{1} }

func (m *Chunks) GetChunks() []*SessionChunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func init() {
	proto.RegisterType((*SessionChunk)(nil), "protogen.SessionChunk")
	proto.RegisterType((*Chunks)(nil), "protogen.Chunks")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Audit service

type AuditClient interface {
	SessionChunks(ctx context.Context, opts ...grpc.CallOption) (Audit_SessionChunksClient, error)
}

type auditClient struct {
	cc *grpc.ClientConn
}

func NewAuditClient(cc *grpc.ClientConn) AuditClient {
	return &auditClient{cc}
}

func (c *auditClient) SessionChunks(ctx context.Context, opts ...grpc.CallOption) (Audit_SessionChunksClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Audit_serviceDesc.Streams[0], c.cc, "/protogen.Audit/SessionChunks", opts...)
	if err != nil {
		return nil, err
	}
	x := &auditSessionChunksClient{stream}
	return x, nil
}

type Audit_SessionChunksClient interface {
	Send(*Chunks) error
	CloseAndRecv() (*google_protobuf1.Empty, error)
	grpc.ClientStream
}

type auditSessionChunksClient struct {
	grpc.ClientStream
}

func (x *auditSessionChunksClient) Send(m *Chunks) error {
	return x.ClientStream.SendMsg(m)
}

func (x *auditSessionChunksClient) CloseAndRecv() (*google_protobuf1.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(google_protobuf1.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Audit service

type AuditServer interface {
	SessionChunks(Audit_SessionChunksServer) error
}

func RegisterAuditServer(s *grpc.Server, srv AuditServer) {
	s.RegisterService(&_Audit_serviceDesc, srv)
}

func _Audit_SessionChunks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AuditServer).SessionChunks(&auditSessionChunksServer{stream})
}

type Audit_SessionChunksServer interface {
	SendAndClose(*google_protobuf1.Empty) error
	Recv() (*Chunks, error)
	grpc.ServerStream
}

type auditSessionChunksServer struct {
	grpc.ServerStream
}

func (x *auditSessionChunksServer) SendAndClose(m *google_protobuf1.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *auditSessionChunksServer) Recv() (*Chunks, error) {
	m := new(Chunks)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Audit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protogen.Audit",
	HandlerType: (*AuditServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SessionChunks",
			Handler:       _Audit_SessionChunks_Handler,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptorApi,
}

func (m *SessionChunk) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SessionChunk) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintApi(data, i, uint64(len(m.Namespace)))
		i += copy(data[i:], m.Namespace)
	}
	if len(m.SessionID) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintApi(data, i, uint64(len(m.SessionID)))
		i += copy(data[i:], m.SessionID)
	}
	if len(m.Chunk) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintApi(data, i, uint64(len(m.Chunk)))
		i += copy(data[i:], m.Chunk)
	}
	return i, nil
}

func (m *Chunks) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Chunks) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Chunks) > 0 {
		for _, msg := range m.Chunks {
			data[i] = 0xa
			i++
			i = encodeVarintApi(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Api(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Api(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintApi(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *SessionChunk) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Chunk)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	return n
}

func (m *Chunks) Size() (n int) {
	var l int
	_ = l
	if len(m.Chunks) > 0 {
		for _, e := range m.Chunks {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	return n
}

func sovApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SessionChunk) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SessionChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunk = append(m.Chunk[:0], data[iNdEx:postIndex]...)
			if m.Chunk == nil {
				m.Chunk = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func (m *Chunks) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Chunks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chunks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunks = append(m.Chunks, &SessionChunk{})
			if err := m.Chunks[len(m.Chunks)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func skipApi(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipApi(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api.proto", fileDescriptorApi) }

var fileDescriptorApi = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0xbf, 0x4a, 0xc4, 0x30,
	0x1c, 0xc7, 0x2f, 0x1e, 0x57, 0x6c, 0x3c, 0xe1, 0x08, 0x72, 0x94, 0x2a, 0xa5, 0x74, 0xea, 0x62,
	0x02, 0xe7, 0xe2, 0xe2, 0xe0, 0x5f, 0x70, 0x71, 0xa8, 0x9b, 0x8b, 0xa6, 0xbd, 0x98, 0x06, 0x6d,
	0x12, 0x9a, 0x14, 0xe9, 0x9b, 0xf8, 0x48, 0x8e, 0x3e, 0x82, 0xd4, 0x17, 0x91, 0xa6, 0x2d, 0xde,
	0x94, 0xdf, 0xf7, 0xcf, 0x2f, 0xc9, 0x07, 0xfa, 0x54, 0x0b, 0xac, 0x6b, 0x65, 0x15, 0xda, 0x77,
	0x07, 0x67, 0x32, 0x7c, 0xe2, 0xc2, 0x96, 0x4d, 0x8e, 0x0b, 0x55, 0x11, 0x5e, 0xeb, 0xe2, 0x94,
	0x15, 0xca, 0xb4, 0xc6, 0xb2, 0x51, 0x72, 0x6a, 0xd9, 0x07, 0x6d, 0x89, 0x2d, 0x45, 0xbd, 0x7d,
	0xd6, 0xb4, 0xb6, 0x2d, 0xe1, 0x4a, 0xf1, 0x77, 0x46, 0xb5, 0x30, 0xe3, 0x48, 0xa8, 0x16, 0x84,
	0x4a, 0xa9, 0x2c, 0xb5, 0x42, 0x49, 0x33, 0xbc, 0x12, 0x1e, 0x8f, 0xa9, 0x53, 0x79, 0xf3, 0x4a,
	0x58, 0xa5, 0x6d, 0x3b, 0x84, 0xc9, 0x0b, 0x5c, 0x3e, 0x32, 0x63, 0x84, 0x92, 0xd7, 0x65, 0x23,
	0xdf, 0xd0, 0x09, 0xf4, 0x1f, 0x68, 0xc5, 0x8c, 0xa6, 0x05, 0x0b, 0x40, 0x0c, 0x52, 0x3f, 0xf3,
	0xe5, 0x64, 0xf4, 0xe9, 0xd8, 0xbe, 0xbf, 0x09, 0xf6, 0x86, 0xd4, 0x4c, 0x06, 0x3a, 0x82, 0x0b,
	0x77, 0x49, 0x30, 0x8f, 0x41, 0xba, 0xcc, 0x16, 0x45, 0x2f, 0x92, 0x73, 0xe8, 0x39, 0xd7, 0x20,
	0x3c, 0x4d, 0x01, 0x88, 0xe7, 0xe9, 0xc1, 0x66, 0x8d, 0x27, 0x7e, 0xbc, 0xfb, 0x87, 0xcc, 0x73,
	0x8b, 0x66, 0x73, 0x07, 0x17, 0x97, 0xcd, 0x56, 0x58, 0x74, 0x01, 0x0f, 0x77, 0x0b, 0x06, 0xad,
	0xfe, 0x37, 0x07, 0x27, 0x5c, 0xe3, 0x81, 0x12, 0x4f, 0x94, 0xf8, 0xb6, 0xa7, 0x4c, 0x66, 0x29,
	0xb8, 0x5a, 0x7d, 0x75, 0x11, 0xf8, 0xee, 0x22, 0xf0, 0xd3, 0x45, 0xe0, 0xf3, 0x37, 0x9a, 0xe5,
	0x9e, 0x6b, 0x9d, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x0b, 0x96, 0x6d, 0x8c, 0x01, 0x00,
	0x00,
}
