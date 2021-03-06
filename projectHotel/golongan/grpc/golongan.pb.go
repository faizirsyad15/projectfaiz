// Code generated by protoc-gen-go. DO NOT EDIT.
// source: golongan.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	golongan.proto

It has these top-level messages:
	AddGolonganReq
	ReadGolonganByNamaReq
	ReadGolonganByNamaResp
	ReadGolonganResp
	UpdateGolonganReq
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddGolonganReq struct {
	IdGolongan   string `protobuf:"bytes,1,opt,name=idGolongan" json:"idGolongan,omitempty"`
	NamaGolongan string `protobuf:"bytes,2,opt,name=namaGolongan" json:"namaGolongan,omitempty"`
	Status       string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *AddGolonganReq) Reset()                    { *m = AddGolonganReq{} }
func (m *AddGolonganReq) String() string            { return proto.CompactTextString(m) }
func (*AddGolonganReq) ProtoMessage()               {}
func (*AddGolonganReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddGolonganReq) GetIdGolongan() string {
	if m != nil {
		return m.IdGolongan
	}
	return ""
}

func (m *AddGolonganReq) GetNamaGolongan() string {
	if m != nil {
		return m.NamaGolongan
	}
	return ""
}

func (m *AddGolonganReq) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ReadGolonganByNamaReq struct {
	NamaGolongan string `protobuf:"bytes,1,opt,name=namaGolongan" json:"namaGolongan,omitempty"`
}

func (m *ReadGolonganByNamaReq) Reset()                    { *m = ReadGolonganByNamaReq{} }
func (m *ReadGolonganByNamaReq) String() string            { return proto.CompactTextString(m) }
func (*ReadGolonganByNamaReq) ProtoMessage()               {}
func (*ReadGolonganByNamaReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadGolonganByNamaReq) GetNamaGolongan() string {
	if m != nil {
		return m.NamaGolongan
	}
	return ""
}

type ReadGolonganByNamaResp struct {
	IdGolongan   string `protobuf:"bytes,1,opt,name=idGolongan" json:"idGolongan,omitempty"`
	NamaGolongan string `protobuf:"bytes,2,opt,name=namaGolongan" json:"namaGolongan,omitempty"`
	Status       string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *ReadGolonganByNamaResp) Reset()                    { *m = ReadGolonganByNamaResp{} }
func (m *ReadGolonganByNamaResp) String() string            { return proto.CompactTextString(m) }
func (*ReadGolonganByNamaResp) ProtoMessage()               {}
func (*ReadGolonganByNamaResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadGolonganByNamaResp) GetIdGolongan() string {
	if m != nil {
		return m.IdGolongan
	}
	return ""
}

func (m *ReadGolonganByNamaResp) GetNamaGolongan() string {
	if m != nil {
		return m.NamaGolongan
	}
	return ""
}

func (m *ReadGolonganByNamaResp) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ReadGolonganResp struct {
	AllGolongan []*ReadGolonganByNamaResp `protobuf:"bytes,1,rep,name=allGolongan" json:"allGolongan,omitempty"`
}

func (m *ReadGolonganResp) Reset()                    { *m = ReadGolonganResp{} }
func (m *ReadGolonganResp) String() string            { return proto.CompactTextString(m) }
func (*ReadGolonganResp) ProtoMessage()               {}
func (*ReadGolonganResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadGolonganResp) GetAllGolongan() []*ReadGolonganByNamaResp {
	if m != nil {
		return m.AllGolongan
	}
	return nil
}

type UpdateGolonganReq struct {
	IdGolongan   string `protobuf:"bytes,1,opt,name=idGolongan" json:"idGolongan,omitempty"`
	NamaGolongan string `protobuf:"bytes,2,opt,name=namaGolongan" json:"namaGolongan,omitempty"`
	Status       string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *UpdateGolonganReq) Reset()                    { *m = UpdateGolonganReq{} }
func (m *UpdateGolonganReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateGolonganReq) ProtoMessage()               {}
func (*UpdateGolonganReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateGolonganReq) GetIdGolongan() string {
	if m != nil {
		return m.IdGolongan
	}
	return ""
}

func (m *UpdateGolonganReq) GetNamaGolongan() string {
	if m != nil {
		return m.NamaGolongan
	}
	return ""
}

func (m *UpdateGolonganReq) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*AddGolonganReq)(nil), "grpc.AddGolonganReq")
	proto.RegisterType((*ReadGolonganByNamaReq)(nil), "grpc.ReadGolonganByNamaReq")
	proto.RegisterType((*ReadGolonganByNamaResp)(nil), "grpc.ReadGolonganByNamaResp")
	proto.RegisterType((*ReadGolonganResp)(nil), "grpc.ReadGolonganResp")
	proto.RegisterType((*UpdateGolonganReq)(nil), "grpc.UpdateGolonganReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for GolonganService service

type GolonganServiceClient interface {
	AddGolongan(ctx context.Context, in *AddGolonganReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadGolongan(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadGolonganResp, error)
	UpdateGolongan(ctx context.Context, in *UpdateGolonganReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadGolonganByNama(ctx context.Context, in *ReadGolonganByNamaReq, opts ...grpc1.CallOption) (*ReadGolonganByNamaResp, error)
}

type golonganServiceClient struct {
	cc *grpc1.ClientConn
}

func NewGolonganServiceClient(cc *grpc1.ClientConn) GolonganServiceClient {
	return &golonganServiceClient{cc}
}

func (c *golonganServiceClient) AddGolongan(ctx context.Context, in *AddGolonganReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.GolonganService/AddGolongan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *golonganServiceClient) ReadGolongan(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadGolonganResp, error) {
	out := new(ReadGolonganResp)
	err := grpc1.Invoke(ctx, "/grpc.GolonganService/ReadGolongan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *golonganServiceClient) UpdateGolongan(ctx context.Context, in *UpdateGolonganReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.GolonganService/UpdateGolongan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *golonganServiceClient) ReadGolonganByNama(ctx context.Context, in *ReadGolonganByNamaReq, opts ...grpc1.CallOption) (*ReadGolonganByNamaResp, error) {
	out := new(ReadGolonganByNamaResp)
	err := grpc1.Invoke(ctx, "/grpc.GolonganService/ReadGolonganByNama", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GolonganService service

type GolonganServiceServer interface {
	AddGolongan(context.Context, *AddGolonganReq) (*google_protobuf.Empty, error)
	ReadGolongan(context.Context, *google_protobuf.Empty) (*ReadGolonganResp, error)
	UpdateGolongan(context.Context, *UpdateGolonganReq) (*google_protobuf.Empty, error)
	ReadGolonganByNama(context.Context, *ReadGolonganByNamaReq) (*ReadGolonganByNamaResp, error)
}

func RegisterGolonganServiceServer(s *grpc1.Server, srv GolonganServiceServer) {
	s.RegisterService(&_GolonganService_serviceDesc, srv)
}

func _GolonganService_AddGolongan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGolonganReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolonganServiceServer).AddGolongan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GolonganService/AddGolongan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolonganServiceServer).AddGolongan(ctx, req.(*AddGolonganReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GolonganService_ReadGolongan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolonganServiceServer).ReadGolongan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GolonganService/ReadGolongan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolonganServiceServer).ReadGolongan(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GolonganService_UpdateGolongan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGolonganReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolonganServiceServer).UpdateGolongan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GolonganService/UpdateGolongan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolonganServiceServer).UpdateGolongan(ctx, req.(*UpdateGolonganReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GolonganService_ReadGolonganByNama_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadGolonganByNamaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolonganServiceServer).ReadGolonganByNama(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GolonganService/ReadGolonganByNama",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolonganServiceServer).ReadGolonganByNama(ctx, req.(*ReadGolonganByNamaReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GolonganService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.GolonganService",
	HandlerType: (*GolonganServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddGolongan",
			Handler:    _GolonganService_AddGolongan_Handler,
		},
		{
			MethodName: "ReadGolongan",
			Handler:    _GolonganService_ReadGolongan_Handler,
		},
		{
			MethodName: "UpdateGolongan",
			Handler:    _GolonganService_UpdateGolongan_Handler,
		},
		{
			MethodName: "ReadGolonganByNama",
			Handler:    _GolonganService_ReadGolonganByNama_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "golongan.proto",
}

func init() { proto.RegisterFile("golongan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0x4c, 0xda, 0x1f, 0x85, 0xdf, 0x4b, 0x89, 0xfa, 0xd0, 0x58, 0x52, 0x11, 0xd9, 0x53, 0x4f,
	0x5b, 0xa8, 0x47, 0x51, 0xfc, 0x83, 0x78, 0x13, 0x8c, 0xf8, 0x01, 0xb6, 0xcd, 0x1a, 0x0a, 0x49,
	0x76, 0x9b, 0x6c, 0x85, 0x7e, 0x25, 0x3f, 0xa5, 0xec, 0xa6, 0x0d, 0x9b, 0xa6, 0xed, 0x4d, 0x8f,
	0x79, 0x99, 0x37, 0x33, 0x3b, 0x6f, 0xc0, 0x4f, 0x44, 0x2a, 0xf2, 0x84, 0xe5, 0x54, 0x16, 0x42,
	0x09, 0xfc, 0x97, 0x14, 0x72, 0x16, 0x0e, 0x13, 0x21, 0x92, 0x94, 0x8f, 0xcd, 0x6c, 0xba, 0xfc,
	0x1c, 0xf3, 0x4c, 0xaa, 0x55, 0x05, 0x21, 0x29, 0xf8, 0x0f, 0x71, 0xfc, 0xb2, 0xde, 0x8b, 0xf8,
	0x02, 0x2f, 0x01, 0xe6, 0xf5, 0x60, 0xe0, 0x5e, 0xb9, 0xa3, 0xff, 0x91, 0x35, 0x41, 0x02, 0xfd,
	0x9c, 0x65, 0xac, 0x46, 0x74, 0x0c, 0xa2, 0x31, 0xc3, 0x00, 0x7a, 0xa5, 0x62, 0x6a, 0x59, 0x0e,
	0xba, 0xe6, 0xef, 0xfa, 0x8b, 0xdc, 0xc0, 0x59, 0xc4, 0x59, 0xcd, 0xf5, 0xb8, 0x7a, 0x65, 0x19,
	0xd3, 0xa2, 0xdb, 0xa4, 0x6e, 0x9b, 0x94, 0x28, 0x08, 0x76, 0x2d, 0x97, 0xf2, 0x57, 0x2d, 0x47,
	0x70, 0x6c, 0xab, 0x1a, 0xbd, 0x3b, 0xf0, 0x58, 0x9a, 0x5a, 0x82, 0xdd, 0x91, 0x37, 0xb9, 0xa0,
	0x3a, 0x6d, 0xba, 0xdb, 0x62, 0x64, 0x2f, 0x10, 0x01, 0x27, 0x1f, 0x32, 0x66, 0x8a, 0xff, 0x51,
	0xee, 0x93, 0xef, 0x0e, 0x1c, 0x6d, 0x40, 0xef, 0xbc, 0xf8, 0x9a, 0xcf, 0x38, 0xde, 0x82, 0x67,
	0x5d, 0x1e, 0x4f, 0x2b, 0xfb, 0xcd, 0x32, 0x84, 0x01, 0xad, 0xca, 0x43, 0x37, 0xe5, 0xa1, 0xcf,
	0xba, 0x3c, 0xc4, 0xc1, 0x7b, 0xe8, 0xdb, 0x4f, 0xc5, 0x3d, 0x48, 0xcd, 0xb0, 0x1d, 0x8b, 0x0e,
	0x84, 0x38, 0xf8, 0x04, 0x7e, 0x33, 0x05, 0x3c, 0xaf, 0xb0, 0xad, 0x6c, 0x0e, 0xd8, 0x78, 0x03,
	0x6c, 0x27, 0x8e, 0xc3, 0xfd, 0xb7, 0x58, 0x84, 0x07, 0x0f, 0x45, 0x9c, 0x69, 0xcf, 0x88, 0x5c,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xa2, 0xab, 0xd0, 0x4e, 0x03, 0x00, 0x00,
}
