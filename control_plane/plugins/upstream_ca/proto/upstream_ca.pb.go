// Code generated by protoc-gen-go. DO NOT EDIT.
// source: upstream_ca.proto

/*
Package control_plane_proto is a generated protocol buffer package.

It is generated from these files:
	upstream_ca.proto

It has these top-level messages:
	SubmitCSRRequest
	SubmitCSRResponse
*/
package control_plane_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import control_plane_proto1 "github.com/spiffe/sri/control_plane/plugins/common/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// ConfigureRequest from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type ConfigureRequest control_plane_proto1.ConfigureRequest

func (m *ConfigureRequest) Reset() { (*control_plane_proto1.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string {
	return (*control_plane_proto1.ConfigureRequest)(m).String()
}
func (*ConfigureRequest) ProtoMessage() {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*control_plane_proto1.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type ConfigureResponse control_plane_proto1.ConfigureResponse

func (m *ConfigureResponse) Reset() { (*control_plane_proto1.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string {
	return (*control_plane_proto1.ConfigureResponse)(m).String()
}
func (*ConfigureResponse) ProtoMessage() {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*control_plane_proto1.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type GetPluginInfoRequest control_plane_proto1.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset() { (*control_plane_proto1.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string {
	return (*control_plane_proto1.GetPluginInfoRequest)(m).String()
}
func (*GetPluginInfoRequest) ProtoMessage() {}

// GetPluginInfoResponse from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type GetPluginInfoResponse control_plane_proto1.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset() { (*control_plane_proto1.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).String()
}
func (*GetPluginInfoResponse) ProtoMessage() {}
func (m *GetPluginInfoResponse) GetName() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetName()
}
func (m *GetPluginInfoResponse) GetCategory() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetCategory()
}
func (m *GetPluginInfoResponse) GetType() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetType()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetCompany()
}

type SubmitCSRRequest struct {
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (m *SubmitCSRRequest) Reset()                    { *m = SubmitCSRRequest{} }
func (m *SubmitCSRRequest) String() string            { return proto.CompactTextString(m) }
func (*SubmitCSRRequest) ProtoMessage()               {}
func (*SubmitCSRRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SubmitCSRRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

type SubmitCSRResponse struct {
	Cert                []byte `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	UpstreamTrustBundle []byte `protobuf:"bytes,2,opt,name=upstreamTrustBundle,proto3" json:"upstreamTrustBundle,omitempty"`
}

func (m *SubmitCSRResponse) Reset()                    { *m = SubmitCSRResponse{} }
func (m *SubmitCSRResponse) String() string            { return proto.CompactTextString(m) }
func (*SubmitCSRResponse) ProtoMessage()               {}
func (*SubmitCSRResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SubmitCSRResponse) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *SubmitCSRResponse) GetUpstreamTrustBundle() []byte {
	if m != nil {
		return m.UpstreamTrustBundle
	}
	return nil
}

func init() {
	proto.RegisterType((*SubmitCSRRequest)(nil), "control_plane_proto.SubmitCSRRequest")
	proto.RegisterType((*SubmitCSRResponse)(nil), "control_plane_proto.SubmitCSRResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UpstreamCA service

type UpstreamCAClient interface {
	// *Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *control_plane_proto1.ConfigureRequest, opts ...grpc.CallOption) (*control_plane_proto1.ConfigureResponse, error)
	// *Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *control_plane_proto1.GetPluginInfoRequest, opts ...grpc.CallOption) (*control_plane_proto1.GetPluginInfoResponse, error)
	// *Will take in a CSR and submit it to the upstream CA for signing
	// (“upstream” CA can be local self-signed root in simple case).
	SubmitCSR(ctx context.Context, in *SubmitCSRRequest, opts ...grpc.CallOption) (*SubmitCSRResponse, error)
}

type upstreamCAClient struct {
	cc *grpc.ClientConn
}

func NewUpstreamCAClient(cc *grpc.ClientConn) UpstreamCAClient {
	return &upstreamCAClient{cc}
}

func (c *upstreamCAClient) Configure(ctx context.Context, in *control_plane_proto1.ConfigureRequest, opts ...grpc.CallOption) (*control_plane_proto1.ConfigureResponse, error) {
	out := new(control_plane_proto1.ConfigureResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.UpstreamCA/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamCAClient) GetPluginInfo(ctx context.Context, in *control_plane_proto1.GetPluginInfoRequest, opts ...grpc.CallOption) (*control_plane_proto1.GetPluginInfoResponse, error) {
	out := new(control_plane_proto1.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.UpstreamCA/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamCAClient) SubmitCSR(ctx context.Context, in *SubmitCSRRequest, opts ...grpc.CallOption) (*SubmitCSRResponse, error) {
	out := new(SubmitCSRResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.UpstreamCA/SubmitCSR", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UpstreamCA service

type UpstreamCAServer interface {
	// *Responsible for configuration of the plugin.
	Configure(context.Context, *control_plane_proto1.ConfigureRequest) (*control_plane_proto1.ConfigureResponse, error)
	// *Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *control_plane_proto1.GetPluginInfoRequest) (*control_plane_proto1.GetPluginInfoResponse, error)
	// *Will take in a CSR and submit it to the upstream CA for signing
	// (“upstream” CA can be local self-signed root in simple case).
	SubmitCSR(context.Context, *SubmitCSRRequest) (*SubmitCSRResponse, error)
}

func RegisterUpstreamCAServer(s *grpc.Server, srv UpstreamCAServer) {
	s.RegisterService(&_UpstreamCA_serviceDesc, srv)
}

func _UpstreamCA_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(control_plane_proto1.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamCAServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.UpstreamCA/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamCAServer).Configure(ctx, req.(*control_plane_proto1.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpstreamCA_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(control_plane_proto1.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamCAServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.UpstreamCA/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamCAServer).GetPluginInfo(ctx, req.(*control_plane_proto1.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpstreamCA_SubmitCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitCSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamCAServer).SubmitCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.UpstreamCA/SubmitCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamCAServer).SubmitCSR(ctx, req.(*SubmitCSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UpstreamCA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "control_plane_proto.UpstreamCA",
	HandlerType: (*UpstreamCAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _UpstreamCA_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _UpstreamCA_GetPluginInfo_Handler,
		},
		{
			MethodName: "SubmitCSR",
			Handler:    _UpstreamCA_SubmitCSR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upstream_ca.proto",
}

func init() { proto.RegisterFile("upstream_ca.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x6d, 0x15, 0xa1, 0x83, 0x42, 0xbb, 0xbd, 0x94, 0x9e, 0xa4, 0xa8, 0xa8, 0x87, 0xac,
	0xe8, 0x13, 0x68, 0x10, 0xf1, 0x56, 0x52, 0x3d, 0xd8, 0x4b, 0x48, 0xd6, 0x49, 0x5c, 0x48, 0x76,
	0xd6, 0xfd, 0xf3, 0x56, 0x3e, 0xa4, 0xb8, 0x6e, 0x8b, 0xd5, 0x14, 0xbd, 0xcd, 0xce, 0xfc, 0x66,
	0xe7, 0xfb, 0x3e, 0x18, 0x79, 0x6d, 0x9d, 0xc1, 0xa2, 0xcd, 0x45, 0x91, 0x68, 0x43, 0x8e, 0xd8,
	0x58, 0x90, 0x72, 0x86, 0x9a, 0x5c, 0x37, 0x85, 0xc2, 0x3c, 0x34, 0xa7, 0x77, 0xb5, 0x74, 0xaf,
	0xbe, 0x4c, 0x04, 0xb5, 0xdc, 0x6a, 0x59, 0x55, 0xc8, 0xad, 0x91, 0x7c, 0x03, 0xe5, 0xba, 0xf1,
	0xb5, 0x54, 0x96, 0x0b, 0x6a, 0x5b, 0x52, 0x3c, 0x6c, 0xc6, 0xc7, 0xd7, 0xdf, 0xb3, 0x63, 0x18,
	0x2e, 0x7c, 0xd9, 0x4a, 0x97, 0x2e, 0xb2, 0x0c, 0xdf, 0x3c, 0x5a, 0xc7, 0x86, 0xb0, 0x2b, 0xac,
	0x99, 0xf4, 0x8e, 0x7a, 0x67, 0x07, 0xd9, 0x67, 0x39, 0x7b, 0x86, 0xd1, 0x37, 0xca, 0x6a, 0x52,
	0x16, 0x19, 0x83, 0x3d, 0x81, 0xc6, 0x45, 0x2e, 0xd4, 0xec, 0x12, 0xc6, 0x2b, 0xfd, 0x8f, 0xc6,
	0x5b, 0x77, 0xeb, 0xd5, 0x4b, 0x83, 0x93, 0x7e, 0x40, 0xba, 0x46, 0x57, 0xef, 0x7d, 0x80, 0xa7,
	0xd8, 0x4f, 0x6f, 0xd8, 0x12, 0x06, 0x29, 0xa9, 0x4a, 0xd6, 0xde, 0x20, 0x3b, 0x49, 0x3a, 0x9c,
	0x27, 0xeb, 0x79, 0xd4, 0x3b, 0x3d, 0xfd, 0x0b, 0x8b, 0x82, 0x2b, 0x38, 0xbc, 0x47, 0x37, 0x0f,
	0x99, 0x3c, 0xa8, 0x8a, 0xd8, 0x79, 0xe7, 0xe2, 0x06, 0xb3, 0xba, 0x71, 0xf1, 0x1f, 0x34, 0xde,
	0x59, 0xc2, 0x60, 0x9d, 0xd6, 0x16, 0x0f, 0x3f, 0x33, 0xdf, 0xe2, 0xe1, 0x57, 0xe8, 0xf3, 0x9d,
	0x72, 0x3f, 0x8c, 0xae, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x27, 0x6b, 0xa0, 0xc7, 0x29, 0x02,
	0x00, 0x00,
}
