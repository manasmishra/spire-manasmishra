// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workload_attestor.proto

/*
Package node_agent_proto is a generated protocol buffer package.

It is generated from these files:
	workload_attestor.proto

It has these top-level messages:
	AttestRequest
	AttestResponse
*/
package node_agent_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import node_agent_proto1 "github.com/spiffe/sri/node_agent/plugins/common/proto"

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

// ConfigureRequest from public import github.com/spiffe/sri/node_agent/plugins/common/proto/common.proto
type ConfigureRequest node_agent_proto1.ConfigureRequest

func (m *ConfigureRequest) Reset()         { (*node_agent_proto1.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string { return (*node_agent_proto1.ConfigureRequest)(m).String() }
func (*ConfigureRequest) ProtoMessage()    {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*node_agent_proto1.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/sri/node_agent/plugins/common/proto/common.proto
type ConfigureResponse node_agent_proto1.ConfigureResponse

func (m *ConfigureResponse) Reset()         { (*node_agent_proto1.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string { return (*node_agent_proto1.ConfigureResponse)(m).String() }
func (*ConfigureResponse) ProtoMessage()    {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*node_agent_proto1.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/sri/node_agent/plugins/common/proto/common.proto
type GetPluginInfoRequest node_agent_proto1.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset() { (*node_agent_proto1.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string {
	return (*node_agent_proto1.GetPluginInfoRequest)(m).String()
}
func (*GetPluginInfoRequest) ProtoMessage() {}

// GetPluginInfoResponse from public import github.com/spiffe/sri/node_agent/plugins/common/proto/common.proto
type GetPluginInfoResponse node_agent_proto1.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset() { (*node_agent_proto1.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string {
	return (*node_agent_proto1.GetPluginInfoResponse)(m).String()
}
func (*GetPluginInfoResponse) ProtoMessage() {}
func (m *GetPluginInfoResponse) GetName() string {
	return (*node_agent_proto1.GetPluginInfoResponse)(m).GetName()
}
func (m *GetPluginInfoResponse) GetCategory() string {
	return (*node_agent_proto1.GetPluginInfoResponse)(m).GetCategory()
}
func (m *GetPluginInfoResponse) GetType() string {
	return (*node_agent_proto1.GetPluginInfoResponse)(m).GetType()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*node_agent_proto1.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*node_agent_proto1.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*node_agent_proto1.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*node_agent_proto1.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*node_agent_proto1.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*node_agent_proto1.GetPluginInfoResponse)(m).GetCompany()
}

// * Represents the workload PID.
type AttestRequest struct {
	Pid int32 `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
}

func (m *AttestRequest) Reset()                    { *m = AttestRequest{} }
func (m *AttestRequest) String() string            { return proto.CompactTextString(m) }
func (*AttestRequest) ProtoMessage()               {}
func (*AttestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AttestRequest) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

// * Represents a list of selectors resolved for a given PID.
type AttestResponse struct {
	Selectors []string `protobuf:"bytes,1,rep,name=selectors" json:"selectors,omitempty"`
}

func (m *AttestResponse) Reset()                    { *m = AttestResponse{} }
func (m *AttestResponse) String() string            { return proto.CompactTextString(m) }
func (*AttestResponse) ProtoMessage()               {}
func (*AttestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AttestResponse) GetSelectors() []string {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func init() {
	proto.RegisterType((*AttestRequest)(nil), "node_agent_proto.AttestRequest")
	proto.RegisterType((*AttestResponse)(nil), "node_agent_proto.AttestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WorkloadAttestor service

type WorkloadAttestorClient interface {
	// / Returns a list of selectors resolved for a given PID
	Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error)
	// / Applies the plugin configuration and returns configuration errors
	Configure(ctx context.Context, in *node_agent_proto1.ConfigureRequest, opts ...grpc.CallOption) (*node_agent_proto1.ConfigureResponse, error)
	// / Returns the version and related metadata of the plugin
	GetPluginInfo(ctx context.Context, in *node_agent_proto1.GetPluginInfoRequest, opts ...grpc.CallOption) (*node_agent_proto1.GetPluginInfoResponse, error)
}

type workloadAttestorClient struct {
	cc *grpc.ClientConn
}

func NewWorkloadAttestorClient(cc *grpc.ClientConn) WorkloadAttestorClient {
	return &workloadAttestorClient{cc}
}

func (c *workloadAttestorClient) Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error) {
	out := new(AttestResponse)
	err := grpc.Invoke(ctx, "/node_agent_proto.WorkloadAttestor/Attest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadAttestorClient) Configure(ctx context.Context, in *node_agent_proto1.ConfigureRequest, opts ...grpc.CallOption) (*node_agent_proto1.ConfigureResponse, error) {
	out := new(node_agent_proto1.ConfigureResponse)
	err := grpc.Invoke(ctx, "/node_agent_proto.WorkloadAttestor/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadAttestorClient) GetPluginInfo(ctx context.Context, in *node_agent_proto1.GetPluginInfoRequest, opts ...grpc.CallOption) (*node_agent_proto1.GetPluginInfoResponse, error) {
	out := new(node_agent_proto1.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/node_agent_proto.WorkloadAttestor/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkloadAttestor service

type WorkloadAttestorServer interface {
	// / Returns a list of selectors resolved for a given PID
	Attest(context.Context, *AttestRequest) (*AttestResponse, error)
	// / Applies the plugin configuration and returns configuration errors
	Configure(context.Context, *node_agent_proto1.ConfigureRequest) (*node_agent_proto1.ConfigureResponse, error)
	// / Returns the version and related metadata of the plugin
	GetPluginInfo(context.Context, *node_agent_proto1.GetPluginInfoRequest) (*node_agent_proto1.GetPluginInfoResponse, error)
}

func RegisterWorkloadAttestorServer(s *grpc.Server, srv WorkloadAttestorServer) {
	s.RegisterService(&_WorkloadAttestor_serviceDesc, srv)
}

func _WorkloadAttestor_Attest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadAttestorServer).Attest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node_agent_proto.WorkloadAttestor/Attest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadAttestorServer).Attest(ctx, req.(*AttestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkloadAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(node_agent_proto1.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node_agent_proto.WorkloadAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadAttestorServer).Configure(ctx, req.(*node_agent_proto1.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkloadAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(node_agent_proto1.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node_agent_proto.WorkloadAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadAttestorServer).GetPluginInfo(ctx, req.(*node_agent_proto1.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkloadAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "node_agent_proto.WorkloadAttestor",
	HandlerType: (*WorkloadAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Attest",
			Handler:    _WorkloadAttestor_Attest_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _WorkloadAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _WorkloadAttestor_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workload_attestor.proto",
}

func init() { proto.RegisterFile("workload_attestor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xad, 0x8b, 0x0b, 0x0d, 0xac, 0x94, 0x5c, 0x5c, 0x16, 0xc1, 0x5a, 0x41, 0xf7, 0x94,
	0x80, 0x3e, 0xc1, 0xea, 0x41, 0xc4, 0xcb, 0x52, 0x04, 0x8f, 0xb5, 0xdb, 0x4e, 0x6b, 0xb0, 0xcd,
	0xc4, 0x24, 0xc5, 0x47, 0xf1, 0x75, 0xc5, 0x34, 0x75, 0x59, 0x8b, 0xee, 0x6d, 0x32, 0xf3, 0xcd,
	0xff, 0xe7, 0x1f, 0x72, 0xf2, 0x81, 0xfa, 0xad, 0xc1, 0xbc, 0xcc, 0x72, 0x6b, 0xc1, 0x58, 0xd4,
	0x4c, 0x69, 0xb4, 0x48, 0x23, 0x89, 0x25, 0x64, 0x79, 0x0d, 0xd2, 0x66, 0xae, 0xb3, 0xb8, 0xad,
	0x85, 0x7d, 0xed, 0x36, 0xac, 0xc0, 0x96, 0x1b, 0x25, 0xaa, 0x0a, 0xb8, 0xd1, 0x82, 0x6f, 0x39,
	0xae, 0x9a, 0xae, 0x16, 0xd2, 0xf0, 0x02, 0xdb, 0x16, 0x25, 0x77, 0x6b, 0xfe, 0xd1, 0xab, 0x26,
	0xe7, 0x64, 0xb6, 0x72, 0x3e, 0x29, 0xbc, 0x77, 0x60, 0x2c, 0x8d, 0xc8, 0x44, 0x89, 0x72, 0x1e,
	0xc4, 0xc1, 0xf2, 0x28, 0xfd, 0x2e, 0x13, 0x46, 0x8e, 0x07, 0xc4, 0x28, 0x94, 0x06, 0xe8, 0x29,
	0x09, 0x0d, 0x34, 0x50, 0x58, 0xd4, 0x66, 0x1e, 0xc4, 0x93, 0x65, 0x98, 0x6e, 0x1b, 0xd7, 0x9f,
	0x87, 0x24, 0x7a, 0xf6, 0x21, 0x56, 0x3e, 0x03, 0x7d, 0x24, 0xd3, 0xbe, 0xa6, 0x67, 0xec, 0x77,
	0x10, 0xb6, 0xf3, 0x83, 0x45, 0xfc, 0x37, 0xe0, 0xfd, 0x9f, 0x48, 0x78, 0x87, 0xb2, 0x12, 0x75,
	0xa7, 0x81, 0x26, 0x63, 0xfc, 0x67, 0x38, 0x48, 0x5e, 0xfc, 0xcb, 0x78, 0xd5, 0x17, 0x32, 0xbb,
	0x07, 0xbb, 0x76, 0x27, 0x7b, 0x90, 0x15, 0xd2, 0xcb, 0xf1, 0xd6, 0x0e, 0x30, 0xa8, 0x5f, 0xed,
	0xe5, 0x7a, 0x87, 0xf5, 0xc1, 0x66, 0xea, 0xc6, 0x37, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8a,
	0x09, 0xab, 0x51, 0xe6, 0x01, 0x00, 0x00,
}
