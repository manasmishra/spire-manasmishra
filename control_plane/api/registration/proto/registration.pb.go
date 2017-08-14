// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registration.proto

/*
Package control_plane_proto is a generated protocol buffer package.

It is generated from these files:
	registration.proto

It has these top-level messages:
	RegisteredEntry
	FederatedEntry
	FederatedBundle
	CreateFederatedEntryRequest
	CreateFederatedEntryResponse
	CreateFederatedBundleRequest
	CreateFederatedBundleResponse
	ListFederatedBundlesRequest
	ListFederatedBundlesResponse
	UpdateFederatedBundleRequest
	UpdateFederatedBundleResponse
	DeleteFederatedBundleRequest
	DeleteFederatedBundleResponse
	CreateEntryRequest
	CreateEntryResponse
	ListAttestorEntriesRequest
	ListAttestorEntriesResponse
	ListSelectorEntriesRequest
	ListSelectorEntriesResponse
	ListSpiffeEntriesRequest
	ListSpiffeEntriesResponse
	DeleteEntryRequest
	DeleteEntryResponse
*/
package control_plane_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type RegisteredEntry struct {
	SelectorType string `protobuf:"bytes,1,opt,name=selectorType" json:"selectorType,omitempty"`
	Selector     string `protobuf:"bytes,2,opt,name=selector" json:"selector,omitempty"`
	Attestor     string `protobuf:"bytes,3,opt,name=attestor" json:"attestor,omitempty"`
	SpiffeId     string `protobuf:"bytes,4,opt,name=spiffeId" json:"spiffeId,omitempty"`
	Ttl          int32  `protobuf:"varint,5,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *RegisteredEntry) Reset()                    { *m = RegisteredEntry{} }
func (m *RegisteredEntry) String() string            { return proto.CompactTextString(m) }
func (*RegisteredEntry) ProtoMessage()               {}
func (*RegisteredEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisteredEntry) GetSelectorType() string {
	if m != nil {
		return m.SelectorType
	}
	return ""
}

func (m *RegisteredEntry) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *RegisteredEntry) GetAttestor() string {
	if m != nil {
		return m.Attestor
	}
	return ""
}

func (m *RegisteredEntry) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *RegisteredEntry) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type FederatedEntry struct {
	RegisteredEntry            *RegisteredEntry `protobuf:"bytes,1,opt,name=registeredEntry" json:"registeredEntry,omitempty"`
	FederateBundleSpiffeIdList []string         `protobuf:"bytes,2,rep,name=federateBundleSpiffeIdList" json:"federateBundleSpiffeIdList,omitempty"`
}

func (m *FederatedEntry) Reset()                    { *m = FederatedEntry{} }
func (m *FederatedEntry) String() string            { return proto.CompactTextString(m) }
func (*FederatedEntry) ProtoMessage()               {}
func (*FederatedEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FederatedEntry) GetRegisteredEntry() *RegisteredEntry {
	if m != nil {
		return m.RegisteredEntry
	}
	return nil
}

func (m *FederatedEntry) GetFederateBundleSpiffeIdList() []string {
	if m != nil {
		return m.FederateBundleSpiffeIdList
	}
	return nil
}

type FederatedBundle struct {
	FederateBundleSpiffeId string `protobuf:"bytes,1,opt,name=federateBundleSpiffeId" json:"federateBundleSpiffeId,omitempty"`
	FederatedBundle        []byte `protobuf:"bytes,2,opt,name=federatedBundle,proto3" json:"federatedBundle,omitempty"`
	Ttl                    int32  `protobuf:"varint,3,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *FederatedBundle) Reset()                    { *m = FederatedBundle{} }
func (m *FederatedBundle) String() string            { return proto.CompactTextString(m) }
func (*FederatedBundle) ProtoMessage()               {}
func (*FederatedBundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FederatedBundle) GetFederateBundleSpiffeId() string {
	if m != nil {
		return m.FederateBundleSpiffeId
	}
	return ""
}

func (m *FederatedBundle) GetFederatedBundle() []byte {
	if m != nil {
		return m.FederatedBundle
	}
	return nil
}

func (m *FederatedBundle) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type CreateFederatedEntryRequest struct {
	FederatedEntry *FederatedEntry `protobuf:"bytes,1,opt,name=federatedEntry" json:"federatedEntry,omitempty"`
}

func (m *CreateFederatedEntryRequest) Reset()                    { *m = CreateFederatedEntryRequest{} }
func (m *CreateFederatedEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateFederatedEntryRequest) ProtoMessage()               {}
func (*CreateFederatedEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateFederatedEntryRequest) GetFederatedEntry() *FederatedEntry {
	if m != nil {
		return m.FederatedEntry
	}
	return nil
}

type CreateFederatedEntryResponse struct {
}

func (m *CreateFederatedEntryResponse) Reset()                    { *m = CreateFederatedEntryResponse{} }
func (m *CreateFederatedEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateFederatedEntryResponse) ProtoMessage()               {}
func (*CreateFederatedEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type CreateFederatedBundleRequest struct {
	FederatedBundle *FederatedBundle `protobuf:"bytes,1,opt,name=federatedBundle" json:"federatedBundle,omitempty"`
}

func (m *CreateFederatedBundleRequest) Reset()                    { *m = CreateFederatedBundleRequest{} }
func (m *CreateFederatedBundleRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateFederatedBundleRequest) ProtoMessage()               {}
func (*CreateFederatedBundleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateFederatedBundleRequest) GetFederatedBundle() *FederatedBundle {
	if m != nil {
		return m.FederatedBundle
	}
	return nil
}

type CreateFederatedBundleResponse struct {
}

func (m *CreateFederatedBundleResponse) Reset()                    { *m = CreateFederatedBundleResponse{} }
func (m *CreateFederatedBundleResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateFederatedBundleResponse) ProtoMessage()               {}
func (*CreateFederatedBundleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ListFederatedBundlesRequest struct {
}

func (m *ListFederatedBundlesRequest) Reset()                    { *m = ListFederatedBundlesRequest{} }
func (m *ListFederatedBundlesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListFederatedBundlesRequest) ProtoMessage()               {}
func (*ListFederatedBundlesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ListFederatedBundlesResponse struct {
	FederateBundleSpiffeIdList []string `protobuf:"bytes,1,rep,name=federateBundleSpiffeIdList" json:"federateBundleSpiffeIdList,omitempty"`
}

func (m *ListFederatedBundlesResponse) Reset()                    { *m = ListFederatedBundlesResponse{} }
func (m *ListFederatedBundlesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListFederatedBundlesResponse) ProtoMessage()               {}
func (*ListFederatedBundlesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListFederatedBundlesResponse) GetFederateBundleSpiffeIdList() []string {
	if m != nil {
		return m.FederateBundleSpiffeIdList
	}
	return nil
}

type UpdateFederatedBundleRequest struct {
	FederatedBundle *FederatedBundle `protobuf:"bytes,1,opt,name=federatedBundle" json:"federatedBundle,omitempty"`
}

func (m *UpdateFederatedBundleRequest) Reset()                    { *m = UpdateFederatedBundleRequest{} }
func (m *UpdateFederatedBundleRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateFederatedBundleRequest) ProtoMessage()               {}
func (*UpdateFederatedBundleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UpdateFederatedBundleRequest) GetFederatedBundle() *FederatedBundle {
	if m != nil {
		return m.FederatedBundle
	}
	return nil
}

type UpdateFederatedBundleResponse struct {
}

func (m *UpdateFederatedBundleResponse) Reset()                    { *m = UpdateFederatedBundleResponse{} }
func (m *UpdateFederatedBundleResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateFederatedBundleResponse) ProtoMessage()               {}
func (*UpdateFederatedBundleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type DeleteFederatedBundleRequest struct {
	FederateBundleSpiffeId string `protobuf:"bytes,1,opt,name=federateBundleSpiffeId" json:"federateBundleSpiffeId,omitempty"`
}

func (m *DeleteFederatedBundleRequest) Reset()                    { *m = DeleteFederatedBundleRequest{} }
func (m *DeleteFederatedBundleRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteFederatedBundleRequest) ProtoMessage()               {}
func (*DeleteFederatedBundleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeleteFederatedBundleRequest) GetFederateBundleSpiffeId() string {
	if m != nil {
		return m.FederateBundleSpiffeId
	}
	return ""
}

type DeleteFederatedBundleResponse struct {
}

func (m *DeleteFederatedBundleResponse) Reset()                    { *m = DeleteFederatedBundleResponse{} }
func (m *DeleteFederatedBundleResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteFederatedBundleResponse) ProtoMessage()               {}
func (*DeleteFederatedBundleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type CreateEntryRequest struct {
	RegisteredEntry *RegisteredEntry `protobuf:"bytes,1,opt,name=registeredEntry" json:"registeredEntry,omitempty"`
}

func (m *CreateEntryRequest) Reset()                    { *m = CreateEntryRequest{} }
func (m *CreateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateEntryRequest) ProtoMessage()               {}
func (*CreateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *CreateEntryRequest) GetRegisteredEntry() *RegisteredEntry {
	if m != nil {
		return m.RegisteredEntry
	}
	return nil
}

type CreateEntryResponse struct {
}

func (m *CreateEntryResponse) Reset()                    { *m = CreateEntryResponse{} }
func (m *CreateEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateEntryResponse) ProtoMessage()               {}
func (*CreateEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type ListAttestorEntriesRequest struct {
	Attestor string `protobuf:"bytes,1,opt,name=attestor" json:"attestor,omitempty"`
}

func (m *ListAttestorEntriesRequest) Reset()                    { *m = ListAttestorEntriesRequest{} }
func (m *ListAttestorEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListAttestorEntriesRequest) ProtoMessage()               {}
func (*ListAttestorEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ListAttestorEntriesRequest) GetAttestor() string {
	if m != nil {
		return m.Attestor
	}
	return ""
}

type ListAttestorEntriesResponse struct {
	FederatedEntryList []*FederatedEntry `protobuf:"bytes,1,rep,name=federatedEntryList" json:"federatedEntryList,omitempty"`
}

func (m *ListAttestorEntriesResponse) Reset()                    { *m = ListAttestorEntriesResponse{} }
func (m *ListAttestorEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListAttestorEntriesResponse) ProtoMessage()               {}
func (*ListAttestorEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *ListAttestorEntriesResponse) GetFederatedEntryList() []*FederatedEntry {
	if m != nil {
		return m.FederatedEntryList
	}
	return nil
}

type ListSelectorEntriesRequest struct {
	SelectorType string `protobuf:"bytes,1,opt,name=selectorType" json:"selectorType,omitempty"`
	Selector     string `protobuf:"bytes,2,opt,name=selector" json:"selector,omitempty"`
}

func (m *ListSelectorEntriesRequest) Reset()                    { *m = ListSelectorEntriesRequest{} }
func (m *ListSelectorEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSelectorEntriesRequest) ProtoMessage()               {}
func (*ListSelectorEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *ListSelectorEntriesRequest) GetSelectorType() string {
	if m != nil {
		return m.SelectorType
	}
	return ""
}

func (m *ListSelectorEntriesRequest) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

type ListSelectorEntriesResponse struct {
	FederatedEntryList []*FederatedEntry `protobuf:"bytes,1,rep,name=federatedEntryList" json:"federatedEntryList,omitempty"`
}

func (m *ListSelectorEntriesResponse) Reset()                    { *m = ListSelectorEntriesResponse{} }
func (m *ListSelectorEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSelectorEntriesResponse) ProtoMessage()               {}
func (*ListSelectorEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *ListSelectorEntriesResponse) GetFederatedEntryList() []*FederatedEntry {
	if m != nil {
		return m.FederatedEntryList
	}
	return nil
}

type ListSpiffeEntriesRequest struct {
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffeId" json:"spiffeId,omitempty"`
}

func (m *ListSpiffeEntriesRequest) Reset()                    { *m = ListSpiffeEntriesRequest{} }
func (m *ListSpiffeEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSpiffeEntriesRequest) ProtoMessage()               {}
func (*ListSpiffeEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *ListSpiffeEntriesRequest) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

type ListSpiffeEntriesResponse struct {
	FederatedEntryList []*FederatedEntry `protobuf:"bytes,1,rep,name=federatedEntryList" json:"federatedEntryList,omitempty"`
}

func (m *ListSpiffeEntriesResponse) Reset()                    { *m = ListSpiffeEntriesResponse{} }
func (m *ListSpiffeEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSpiffeEntriesResponse) ProtoMessage()               {}
func (*ListSpiffeEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *ListSpiffeEntriesResponse) GetFederatedEntryList() []*FederatedEntry {
	if m != nil {
		return m.FederatedEntryList
	}
	return nil
}

type DeleteEntryRequest struct {
	SelectorType string `protobuf:"bytes,1,opt,name=selectorType" json:"selectorType,omitempty"`
	Selector     string `protobuf:"bytes,2,opt,name=selector" json:"selector,omitempty"`
}

func (m *DeleteEntryRequest) Reset()                    { *m = DeleteEntryRequest{} }
func (m *DeleteEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteEntryRequest) ProtoMessage()               {}
func (*DeleteEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *DeleteEntryRequest) GetSelectorType() string {
	if m != nil {
		return m.SelectorType
	}
	return ""
}

func (m *DeleteEntryRequest) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

type DeleteEntryResponse struct {
	FederatedEntryList []*FederatedEntry `protobuf:"bytes,1,rep,name=federatedEntryList" json:"federatedEntryList,omitempty"`
}

func (m *DeleteEntryResponse) Reset()                    { *m = DeleteEntryResponse{} }
func (m *DeleteEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteEntryResponse) ProtoMessage()               {}
func (*DeleteEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *DeleteEntryResponse) GetFederatedEntryList() []*FederatedEntry {
	if m != nil {
		return m.FederatedEntryList
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisteredEntry)(nil), "control_plane_proto.RegisteredEntry")
	proto.RegisterType((*FederatedEntry)(nil), "control_plane_proto.FederatedEntry")
	proto.RegisterType((*FederatedBundle)(nil), "control_plane_proto.FederatedBundle")
	proto.RegisterType((*CreateFederatedEntryRequest)(nil), "control_plane_proto.CreateFederatedEntryRequest")
	proto.RegisterType((*CreateFederatedEntryResponse)(nil), "control_plane_proto.CreateFederatedEntryResponse")
	proto.RegisterType((*CreateFederatedBundleRequest)(nil), "control_plane_proto.CreateFederatedBundleRequest")
	proto.RegisterType((*CreateFederatedBundleResponse)(nil), "control_plane_proto.CreateFederatedBundleResponse")
	proto.RegisterType((*ListFederatedBundlesRequest)(nil), "control_plane_proto.ListFederatedBundlesRequest")
	proto.RegisterType((*ListFederatedBundlesResponse)(nil), "control_plane_proto.ListFederatedBundlesResponse")
	proto.RegisterType((*UpdateFederatedBundleRequest)(nil), "control_plane_proto.UpdateFederatedBundleRequest")
	proto.RegisterType((*UpdateFederatedBundleResponse)(nil), "control_plane_proto.UpdateFederatedBundleResponse")
	proto.RegisterType((*DeleteFederatedBundleRequest)(nil), "control_plane_proto.DeleteFederatedBundleRequest")
	proto.RegisterType((*DeleteFederatedBundleResponse)(nil), "control_plane_proto.DeleteFederatedBundleResponse")
	proto.RegisterType((*CreateEntryRequest)(nil), "control_plane_proto.CreateEntryRequest")
	proto.RegisterType((*CreateEntryResponse)(nil), "control_plane_proto.CreateEntryResponse")
	proto.RegisterType((*ListAttestorEntriesRequest)(nil), "control_plane_proto.ListAttestorEntriesRequest")
	proto.RegisterType((*ListAttestorEntriesResponse)(nil), "control_plane_proto.ListAttestorEntriesResponse")
	proto.RegisterType((*ListSelectorEntriesRequest)(nil), "control_plane_proto.ListSelectorEntriesRequest")
	proto.RegisterType((*ListSelectorEntriesResponse)(nil), "control_plane_proto.ListSelectorEntriesResponse")
	proto.RegisterType((*ListSpiffeEntriesRequest)(nil), "control_plane_proto.ListSpiffeEntriesRequest")
	proto.RegisterType((*ListSpiffeEntriesResponse)(nil), "control_plane_proto.ListSpiffeEntriesResponse")
	proto.RegisterType((*DeleteEntryRequest)(nil), "control_plane_proto.DeleteEntryRequest")
	proto.RegisterType((*DeleteEntryResponse)(nil), "control_plane_proto.DeleteEntryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	CreateFederatedEntry(ctx context.Context, in *CreateFederatedEntryRequest, opts ...grpc.CallOption) (*CreateFederatedEntryResponse, error)
	CreateFederatedBundle(ctx context.Context, in *CreateFederatedBundleRequest, opts ...grpc.CallOption) (*CreateFederatedBundleResponse, error)
	ListFederatedBundles(ctx context.Context, in *ListFederatedBundlesRequest, opts ...grpc.CallOption) (*ListFederatedBundlesResponse, error)
	UpdateFederatedBundle(ctx context.Context, in *UpdateFederatedBundleRequest, opts ...grpc.CallOption) (*UpdateFederatedBundleResponse, error)
	DeleteFederatedBundle(ctx context.Context, in *DeleteFederatedBundleRequest, opts ...grpc.CallOption) (*DeleteFederatedBundleResponse, error)
	CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*CreateEntryResponse, error)
	ListAttestorEntries(ctx context.Context, in *ListAttestorEntriesRequest, opts ...grpc.CallOption) (*ListAttestorEntriesResponse, error)
	ListSelectorEntries(ctx context.Context, in *ListSelectorEntriesRequest, opts ...grpc.CallOption) (*ListSelectorEntriesResponse, error)
	ListSpiffeEntries(ctx context.Context, in *ListSpiffeEntriesRequest, opts ...grpc.CallOption) (*ListSpiffeEntriesResponse, error)
	DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*DeleteEntryResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) CreateFederatedEntry(ctx context.Context, in *CreateFederatedEntryRequest, opts ...grpc.CallOption) (*CreateFederatedEntryResponse, error) {
	out := new(CreateFederatedEntryResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.node/CreateFederatedEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) CreateFederatedBundle(ctx context.Context, in *CreateFederatedBundleRequest, opts ...grpc.CallOption) (*CreateFederatedBundleResponse, error) {
	out := new(CreateFederatedBundleResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.node/CreateFederatedBundle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ListFederatedBundles(ctx context.Context, in *ListFederatedBundlesRequest, opts ...grpc.CallOption) (*ListFederatedBundlesResponse, error) {
	out := new(ListFederatedBundlesResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.node/ListFederatedBundles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) UpdateFederatedBundle(ctx context.Context, in *UpdateFederatedBundleRequest, opts ...grpc.CallOption) (*UpdateFederatedBundleResponse, error) {
	out := new(UpdateFederatedBundleResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.node/UpdateFederatedBundle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) DeleteFederatedBundle(ctx context.Context, in *DeleteFederatedBundleRequest, opts ...grpc.CallOption) (*DeleteFederatedBundleResponse, error) {
	out := new(DeleteFederatedBundleResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.node/DeleteFederatedBundle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*CreateEntryResponse, error) {
	out := new(CreateEntryResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.node/CreateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ListAttestorEntries(ctx context.Context, in *ListAttestorEntriesRequest, opts ...grpc.CallOption) (*ListAttestorEntriesResponse, error) {
	out := new(ListAttestorEntriesResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.node/ListAttestorEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ListSelectorEntries(ctx context.Context, in *ListSelectorEntriesRequest, opts ...grpc.CallOption) (*ListSelectorEntriesResponse, error) {
	out := new(ListSelectorEntriesResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.node/ListSelectorEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ListSpiffeEntries(ctx context.Context, in *ListSpiffeEntriesRequest, opts ...grpc.CallOption) (*ListSpiffeEntriesResponse, error) {
	out := new(ListSpiffeEntriesResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.node/ListSpiffeEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*DeleteEntryResponse, error) {
	out := new(DeleteEntryResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.node/DeleteEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	CreateFederatedEntry(context.Context, *CreateFederatedEntryRequest) (*CreateFederatedEntryResponse, error)
	CreateFederatedBundle(context.Context, *CreateFederatedBundleRequest) (*CreateFederatedBundleResponse, error)
	ListFederatedBundles(context.Context, *ListFederatedBundlesRequest) (*ListFederatedBundlesResponse, error)
	UpdateFederatedBundle(context.Context, *UpdateFederatedBundleRequest) (*UpdateFederatedBundleResponse, error)
	DeleteFederatedBundle(context.Context, *DeleteFederatedBundleRequest) (*DeleteFederatedBundleResponse, error)
	CreateEntry(context.Context, *CreateEntryRequest) (*CreateEntryResponse, error)
	ListAttestorEntries(context.Context, *ListAttestorEntriesRequest) (*ListAttestorEntriesResponse, error)
	ListSelectorEntries(context.Context, *ListSelectorEntriesRequest) (*ListSelectorEntriesResponse, error)
	ListSpiffeEntries(context.Context, *ListSpiffeEntriesRequest) (*ListSpiffeEntriesResponse, error)
	DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_CreateFederatedEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFederatedEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).CreateFederatedEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.node/CreateFederatedEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).CreateFederatedEntry(ctx, req.(*CreateFederatedEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_CreateFederatedBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFederatedBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).CreateFederatedBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.node/CreateFederatedBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).CreateFederatedBundle(ctx, req.(*CreateFederatedBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ListFederatedBundles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFederatedBundlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ListFederatedBundles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.node/ListFederatedBundles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ListFederatedBundles(ctx, req.(*ListFederatedBundlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_UpdateFederatedBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFederatedBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).UpdateFederatedBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.node/UpdateFederatedBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).UpdateFederatedBundle(ctx, req.(*UpdateFederatedBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_DeleteFederatedBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFederatedBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).DeleteFederatedBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.node/DeleteFederatedBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).DeleteFederatedBundle(ctx, req.(*DeleteFederatedBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.node/CreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).CreateEntry(ctx, req.(*CreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ListAttestorEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttestorEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ListAttestorEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.node/ListAttestorEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ListAttestorEntries(ctx, req.(*ListAttestorEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ListSelectorEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSelectorEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ListSelectorEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.node/ListSelectorEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ListSelectorEntries(ctx, req.(*ListSelectorEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ListSpiffeEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSpiffeEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ListSpiffeEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.node/ListSpiffeEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ListSpiffeEntries(ctx, req.(*ListSpiffeEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.node/DeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).DeleteEntry(ctx, req.(*DeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "control_plane_proto.node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFederatedEntry",
			Handler:    _Node_CreateFederatedEntry_Handler,
		},
		{
			MethodName: "CreateFederatedBundle",
			Handler:    _Node_CreateFederatedBundle_Handler,
		},
		{
			MethodName: "ListFederatedBundles",
			Handler:    _Node_ListFederatedBundles_Handler,
		},
		{
			MethodName: "UpdateFederatedBundle",
			Handler:    _Node_UpdateFederatedBundle_Handler,
		},
		{
			MethodName: "DeleteFederatedBundle",
			Handler:    _Node_DeleteFederatedBundle_Handler,
		},
		{
			MethodName: "CreateEntry",
			Handler:    _Node_CreateEntry_Handler,
		},
		{
			MethodName: "ListAttestorEntries",
			Handler:    _Node_ListAttestorEntries_Handler,
		},
		{
			MethodName: "ListSelectorEntries",
			Handler:    _Node_ListSelectorEntries_Handler,
		},
		{
			MethodName: "ListSpiffeEntries",
			Handler:    _Node_ListSpiffeEntries_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _Node_DeleteEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registration.proto",
}

func init() { proto.RegisterFile("registration.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0x16, 0x91, 0x69, 0xd5, 0xc0, 0xa6, 0x41, 0xc6, 0x4d, 0x21, 0x5a, 0x90, 0xc8,
	0x85, 0x50, 0x82, 0x54, 0x71, 0x42, 0xe2, 0x53, 0x42, 0xa0, 0x1e, 0x9c, 0xc2, 0x09, 0x51, 0x42,
	0x3d, 0x41, 0xa9, 0x2c, 0xdb, 0xac, 0xb7, 0x12, 0x55, 0xa5, 0xde, 0xf8, 0x15, 0x5c, 0xf8, 0x01,
	0xfc, 0x48, 0x64, 0x7b, 0xed, 0x76, 0xd7, 0xbb, 0x9b, 0x94, 0x92, 0xde, 0xb2, 0x99, 0x9d, 0x79,
	0x6f, 0x66, 0x9e, 0xdf, 0x02, 0x61, 0xf8, 0x6d, 0x9a, 0x72, 0x36, 0xe6, 0xd3, 0x38, 0x1a, 0x24,
	0x2c, 0xe6, 0x31, 0x69, 0x1f, 0xc4, 0x11, 0x67, 0x71, 0xb8, 0x9f, 0x84, 0xe3, 0x08, 0xf7, 0xf3,
	0x3f, 0xe9, 0x2f, 0x07, 0x5a, 0x7e, 0x7e, 0x17, 0x19, 0x06, 0xaf, 0x23, 0xce, 0x8e, 0x09, 0x85,
	0xb5, 0x14, 0x43, 0x3c, 0xe0, 0x31, 0xdb, 0x3b, 0x4e, 0xd0, 0x75, 0x7a, 0x4e, 0xbf, 0xe9, 0x4b,
	0xff, 0x11, 0x0f, 0xae, 0x97, 0x67, 0x77, 0x29, 0x8f, 0x57, 0xe7, 0x2c, 0x36, 0xe6, 0x1c, 0xd3,
	0x2c, 0xd6, 0x28, 0x62, 0xe5, 0x39, 0xcf, 0x4b, 0xa6, 0x93, 0x09, 0xbe, 0x0d, 0xdc, 0x65, 0x91,
	0x27, 0xce, 0xe4, 0x06, 0x34, 0x38, 0x0f, 0xdd, 0x95, 0x9e, 0xd3, 0x5f, 0xf1, 0xb3, 0x9f, 0xf4,
	0xb7, 0x03, 0xeb, 0x6f, 0x30, 0x40, 0x36, 0xe6, 0x25, 0xb9, 0x5d, 0x68, 0x31, 0x99, 0x6f, 0xce,
	0x6f, 0x75, 0x78, 0x7f, 0xa0, 0xe9, 0x6f, 0xa0, 0xf4, 0xe6, 0xab, 0xc9, 0xe4, 0x19, 0x78, 0x13,
	0x81, 0xf0, 0xe2, 0x28, 0x0a, 0x42, 0x1c, 0x09, 0x3a, 0xef, 0xa7, 0x29, 0x77, 0x97, 0x7a, 0x8d,
	0x7e, 0xd3, 0xb7, 0xdc, 0xa0, 0x3f, 0x1d, 0x68, 0x55, 0x14, 0x8b, 0x38, 0xd9, 0x81, 0x5b, 0xfa,
	0x0c, 0x31, 0x4a, 0x43, 0x94, 0xf4, 0xa1, 0x35, 0x91, 0x4b, 0xe5, 0xb3, 0x5d, 0xf3, 0xd5, 0xbf,
	0xcb, 0x51, 0x35, 0xce, 0x46, 0x75, 0x08, 0x9b, 0x2f, 0x19, 0x8e, 0x39, 0xca, 0xf3, 0xf2, 0xf1,
	0xfb, 0x11, 0xa6, 0x9c, 0xbc, 0x83, 0xf5, 0x89, 0x14, 0x10, 0x53, 0xbb, 0xa7, 0x9d, 0x9a, 0x52,
	0x43, 0x49, 0xa5, 0x77, 0xa0, 0xab, 0xc7, 0x4a, 0x93, 0x38, 0x4a, 0x91, 0x46, 0xb5, 0x78, 0x41,
	0xbb, 0x24, 0xb3, 0x5b, 0xef, 0xd3, 0xb6, 0x43, 0xb5, 0x8a, 0x9a, 0x4c, 0xef, 0xc2, 0x96, 0x01,
	0x4f, 0x10, 0xda, 0x82, 0xcd, 0x6c, 0x59, 0x4a, 0x38, 0x15, 0x7c, 0xe8, 0x67, 0xe8, 0xea, 0xc3,
	0x45, 0xfa, 0x0c, 0x8d, 0x38, 0x33, 0x35, 0x12, 0x41, 0xf7, 0x43, 0x12, 0x5c, 0xe9, 0x3c, 0x0c,
	0x78, 0x62, 0x1e, 0x1f, 0xa1, 0xfb, 0x0a, 0x43, 0x34, 0x12, 0xfa, 0x47, 0x01, 0x67, 0xc0, 0x86,
	0xba, 0x02, 0x38, 0x00, 0x52, 0x6c, 0x4a, 0x12, 0xe7, 0x7f, 0xfe, 0xa6, 0x69, 0x07, 0xda, 0x12,
	0x8a, 0x00, 0x7f, 0x0a, 0x5e, 0xb6, 0x8e, 0xe7, 0xc2, 0x8b, 0xb2, 0xe0, 0xb4, 0x12, 0x81, 0xe4,
	0x5a, 0x8e, 0xec, 0x5a, 0x94, 0x15, 0xfa, 0xa9, 0x65, 0x0a, 0x7d, 0x8c, 0x80, 0xc8, 0x5f, 0x48,
	0xa5, 0x8b, 0x39, 0x3f, 0x30, 0x4d, 0x3a, 0xfd, 0x54, 0xb0, 0x1d, 0x09, 0x57, 0x55, 0xd8, 0x5e,
	0xd2, 0xa3, 0xcb, 0x8e, 0x6a, 0xd5, 0x17, 0xd9, 0xd1, 0x0e, 0xb8, 0x39, 0x66, 0xae, 0x96, 0xfa,
	0xf4, 0x53, 0x59, 0x63, 0xd5, 0x99, 0x26, 0x70, 0x5b, 0x93, 0xb7, 0x48, 0xa6, 0x7b, 0x40, 0x0a,
	0x1d, 0x4b, 0x32, 0xbd, 0xec, 0xcc, 0x0f, 0xa1, 0x2d, 0x55, 0x5d, 0x60, 0x07, 0xc3, 0x3f, 0x4d,
	0x58, 0x8e, 0xe2, 0x00, 0xc9, 0x09, 0x6c, 0xe8, 0xbc, 0x9a, 0x6c, 0x6b, 0x2b, 0x5b, 0x9e, 0x10,
	0xef, 0xf1, 0x05, 0x32, 0x44, 0x6b, 0xa7, 0xd0, 0xd1, 0x1a, 0x33, 0x99, 0xab, 0x96, 0xe4, 0x49,
	0xde, 0xf0, 0x22, 0x29, 0x02, 0xff, 0x04, 0x36, 0x74, 0xc6, 0x6e, 0x68, 0xde, 0xf2, 0x44, 0x18,
	0x9a, 0xb7, 0xbe, 0x1a, 0xa7, 0xd0, 0xd1, 0xba, 0xb0, 0xa1, 0x79, 0xdb, 0x0b, 0x61, 0x68, 0xde,
	0x6a, 0xf2, 0x19, 0xbe, 0xd6, 0x8c, 0x0d, 0xf8, 0xb6, 0x07, 0xc1, 0x80, 0x6f, 0xf5, 0x7a, 0xf2,
	0x05, 0x56, 0xcf, 0xb9, 0x30, 0x79, 0x60, 0xd9, 0x9f, 0xa4, 0xb3, 0xfe, 0xec, 0x8b, 0x02, 0xe1,
	0x07, 0xb4, 0x35, 0xb6, 0x4c, 0x1e, 0x19, 0x77, 0xa5, 0xb7, 0x7e, 0x6f, 0x7b, 0xfe, 0x04, 0x19,
	0x59, 0xb1, 0x4f, 0x0b, 0xb2, 0xde, 0xc6, 0x2d, 0xc8, 0x26, 0x67, 0x66, 0x70, 0xb3, 0x66, 0x86,
	0xe4, 0xa1, 0xb9, 0x8c, 0xc6, 0x6c, 0xbd, 0xc1, 0xbc, 0xd7, 0xcf, 0x36, 0x79, 0xce, 0xb8, 0x0c,
	0x9b, 0xac, 0x1b, 0xa6, 0x61, 0x93, 0x1a, 0x0f, 0xfc, 0x7a, 0x2d, 0x0f, 0x3d, 0xf9, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x8c, 0x89, 0x61, 0xc2, 0xb8, 0x0c, 0x00, 0x00,
}
