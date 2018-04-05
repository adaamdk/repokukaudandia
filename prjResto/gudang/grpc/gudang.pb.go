// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gudang.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	gudang.proto

It has these top-level messages:
	AddGudangReq
	UpdateGudangReq
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

// 5.
type AddGudangReq struct {
	ID           string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	NamaGudang   string `protobuf:"bytes,2,opt,name=namaGudang" json:"namaGudang,omitempty"`
	AlamatGudang string `protobuf:"bytes,3,opt,name=alamatGudang" json:"alamatGudang,omitempty"`
	LuasGudang   string `protobuf:"bytes,4,opt,name=luasGudang" json:"luasGudang,omitempty"`
	Status       int32  `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
	CreatedBy    string `protobuf:"bytes,6,opt,name=createdBy" json:"createdBy,omitempty"`
	CreatedOn    string `protobuf:"bytes,7,opt,name=createdOn" json:"createdOn,omitempty"`
	UpdatedBy    string `protobuf:"bytes,8,opt,name=updatedBy" json:"updatedBy,omitempty"`
	UpdatedOn    string `protobuf:"bytes,9,opt,name=updatedOn" json:"updatedOn,omitempty"`
}

func (m *AddGudangReq) Reset()                    { *m = AddGudangReq{} }
func (m *AddGudangReq) String() string            { return proto.CompactTextString(m) }
func (*AddGudangReq) ProtoMessage()               {}
func (*AddGudangReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddGudangReq) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AddGudangReq) GetNamaGudang() string {
	if m != nil {
		return m.NamaGudang
	}
	return ""
}

func (m *AddGudangReq) GetAlamatGudang() string {
	if m != nil {
		return m.AlamatGudang
	}
	return ""
}

func (m *AddGudangReq) GetLuasGudang() string {
	if m != nil {
		return m.LuasGudang
	}
	return ""
}

func (m *AddGudangReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddGudangReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *AddGudangReq) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *AddGudangReq) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func (m *AddGudangReq) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

type UpdateGudangReq struct {
	ID           string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	NamaGudang   string `protobuf:"bytes,2,opt,name=namaGudang" json:"namaGudang,omitempty"`
	AlamatGudang string `protobuf:"bytes,3,opt,name=alamatGudang" json:"alamatGudang,omitempty"`
	LuasGudang   string `protobuf:"bytes,4,opt,name=luasGudang" json:"luasGudang,omitempty"`
	Status       int32  `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
	CreatedBy    string `protobuf:"bytes,6,opt,name=createdBy" json:"createdBy,omitempty"`
	CreatedOn    string `protobuf:"bytes,7,opt,name=createdOn" json:"createdOn,omitempty"`
	UpdatedBy    string `protobuf:"bytes,8,opt,name=updatedBy" json:"updatedBy,omitempty"`
	UpdatedOn    string `protobuf:"bytes,9,opt,name=updatedOn" json:"updatedOn,omitempty"`
}

func (m *UpdateGudangReq) Reset()                    { *m = UpdateGudangReq{} }
func (m *UpdateGudangReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateGudangReq) ProtoMessage()               {}
func (*UpdateGudangReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpdateGudangReq) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UpdateGudangReq) GetNamaGudang() string {
	if m != nil {
		return m.NamaGudang
	}
	return ""
}

func (m *UpdateGudangReq) GetAlamatGudang() string {
	if m != nil {
		return m.AlamatGudang
	}
	return ""
}

func (m *UpdateGudangReq) GetLuasGudang() string {
	if m != nil {
		return m.LuasGudang
	}
	return ""
}

func (m *UpdateGudangReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateGudangReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *UpdateGudangReq) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *UpdateGudangReq) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func (m *UpdateGudangReq) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func init() {
	proto.RegisterType((*AddGudangReq)(nil), "grpc.AddGudangReq")
	proto.RegisterType((*UpdateGudangReq)(nil), "grpc.UpdateGudangReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for GudangService service

type GudangServiceClient interface {
	AddGudang(ctx context.Context, in *AddGudangReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	UpdateGudang(ctx context.Context, in *UpdateGudangReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
}

type gudangServiceClient struct {
	cc *grpc1.ClientConn
}

func NewGudangServiceClient(cc *grpc1.ClientConn) GudangServiceClient {
	return &gudangServiceClient{cc}
}

func (c *gudangServiceClient) AddGudang(ctx context.Context, in *AddGudangReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.GudangService/AddGudang", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gudangServiceClient) UpdateGudang(ctx context.Context, in *UpdateGudangReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.GudangService/UpdateGudang", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GudangService service

type GudangServiceServer interface {
	AddGudang(context.Context, *AddGudangReq) (*google_protobuf.Empty, error)
	UpdateGudang(context.Context, *UpdateGudangReq) (*google_protobuf.Empty, error)
}

func RegisterGudangServiceServer(s *grpc1.Server, srv GudangServiceServer) {
	s.RegisterService(&_GudangService_serviceDesc, srv)
}

func _GudangService_AddGudang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGudangReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GudangServiceServer).AddGudang(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GudangService/AddGudang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GudangServiceServer).AddGudang(ctx, req.(*AddGudangReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GudangService_UpdateGudang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGudangReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GudangServiceServer).UpdateGudang(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GudangService/UpdateGudang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GudangServiceServer).UpdateGudang(ctx, req.(*UpdateGudangReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GudangService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.GudangService",
	HandlerType: (*GudangServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddGudang",
			Handler:    _GudangService_AddGudang_Handler,
		},
		{
			MethodName: "UpdateGudang",
			Handler:    _GudangService_UpdateGudang_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "gudang.proto",
}

func init() { proto.RegisterFile("gudang.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x92, 0xd1, 0x4a, 0xc3, 0x30,
	0x18, 0x85, 0x6d, 0xdd, 0xaa, 0xfd, 0xa9, 0x0a, 0x3f, 0x38, 0xc2, 0x14, 0x19, 0xbd, 0xda, 0x55,
	0x06, 0x7a, 0xe5, 0x95, 0x28, 0x13, 0xd9, 0x55, 0xa1, 0xe2, 0x03, 0x64, 0x6d, 0x0c, 0x42, 0x9b,
	0xd6, 0x36, 0x11, 0xf6, 0x0c, 0xde, 0x09, 0xbe, 0xaf, 0x34, 0xa9, 0x5b, 0x2a, 0xf8, 0x06, 0xbb,
	0xcc, 0x77, 0xce, 0xf9, 0x09, 0x87, 0x03, 0x91, 0xd0, 0x39, 0x93, 0x82, 0xd6, 0x4d, 0xa5, 0x2a,
	0x1c, 0x89, 0xa6, 0xce, 0xa6, 0x17, 0xa2, 0xaa, 0x44, 0xc1, 0x17, 0x86, 0xad, 0xf5, 0xeb, 0x82,
	0x97, 0xb5, 0xda, 0x58, 0x4b, 0xfc, 0xe5, 0x43, 0x74, 0x9f, 0xe7, 0x4f, 0x26, 0x96, 0xf2, 0x77,
	0x3c, 0x05, 0x7f, 0xb5, 0x24, 0xde, 0xcc, 0x9b, 0x87, 0xa9, 0xbf, 0x5a, 0xe2, 0x15, 0x80, 0x64,
	0x25, 0xb3, 0x06, 0xe2, 0x1b, 0xee, 0x10, 0x8c, 0x21, 0x62, 0x05, 0x2b, 0x99, 0xea, 0x1d, 0x87,
	0xc6, 0x31, 0x60, 0xdd, 0x8d, 0x42, 0xb3, 0xb6, 0x77, 0x8c, 0xec, 0x8d, 0x1d, 0xc1, 0x09, 0x04,
	0xad, 0x62, 0x4a, 0xb7, 0x64, 0x3c, 0xf3, 0xe6, 0xe3, 0xb4, 0x7f, 0xe1, 0x25, 0x84, 0x59, 0xc3,
	0x99, 0xe2, 0xf9, 0xc3, 0x86, 0x04, 0x26, 0xb6, 0x03, 0x8e, 0x9a, 0x48, 0x72, 0x34, 0x50, 0x13,
	0xd9, 0xa9, 0xba, 0xce, 0xfb, 0xec, 0xb1, 0x55, 0xb7, 0xc0, 0x51, 0x13, 0x49, 0xc2, 0x81, 0x9a,
	0xc8, 0xf8, 0xdb, 0x87, 0xb3, 0x17, 0xf3, 0xda, 0xf7, 0xe2, 0xf4, 0x72, 0xfd, 0xe9, 0xc1, 0x89,
	0xfd, 0xda, 0x33, 0x6f, 0x3e, 0xde, 0x32, 0x8e, 0xb7, 0x10, 0x6e, 0xd7, 0x83, 0x48, 0xbb, 0xbd,
	0x51, 0x77, 0x4e, 0xd3, 0x09, 0xb5, 0xeb, 0xa3, 0xbf, 0xeb, 0xa3, 0x8f, 0xdd, 0xfa, 0xe2, 0x03,
	0xbc, 0x83, 0xc8, 0xed, 0x18, 0xcf, 0x6d, 0xfa, 0x4f, 0xef, 0xff, 0x1f, 0x58, 0x07, 0x86, 0xdc,
	0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x90, 0x25, 0xe1, 0xf4, 0x02, 0x00, 0x00,
}
