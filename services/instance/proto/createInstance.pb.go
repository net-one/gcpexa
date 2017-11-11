// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createInstance.proto

/*
Package createInstance is a generated protocol buffer package.

It is generated from these files:
	createInstance.proto

It has these top-level messages:
	CreateInstanceRequest
	CreateInstanceResponse
*/
package createInstance

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type CreateInstanceRequest struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instanceName" json:"instanceName,omitempty"`
}

func (m *CreateInstanceRequest) Reset()                    { *m = CreateInstanceRequest{} }
func (m *CreateInstanceRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateInstanceRequest) ProtoMessage()               {}
func (*CreateInstanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateInstanceRequest) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type CreateInstanceResponse struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instanceName" json:"instanceName,omitempty"`
}

func (m *CreateInstanceResponse) Reset()                    { *m = CreateInstanceResponse{} }
func (m *CreateInstanceResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateInstanceResponse) ProtoMessage()               {}
func (*CreateInstanceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateInstanceResponse) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateInstanceRequest)(nil), "CreateInstanceRequest")
	proto.RegisterType((*CreateInstanceResponse)(nil), "CreateInstanceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Instance service

type InstanceClient interface {
	CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...client.CallOption) (*CreateInstanceResponse, error)
}

type instanceClient struct {
	c           client.Client
	serviceName string
}

func NewInstanceClient(serviceName string, c client.Client) InstanceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "instance"
	}
	return &instanceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *instanceClient) CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...client.CallOption) (*CreateInstanceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Instance.CreateInstance", in)
	out := new(CreateInstanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Instance service

type InstanceHandler interface {
	CreateInstance(context.Context, *CreateInstanceRequest, *CreateInstanceResponse) error
}

func RegisterInstanceHandler(s server.Server, hdlr InstanceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Instance{hdlr}, opts...))
}

type Instance struct {
	InstanceHandler
}

func (h *Instance) CreateInstance(ctx context.Context, in *CreateInstanceRequest, out *CreateInstanceResponse) error {
	return h.InstanceHandler.CreateInstance(ctx, in, out)
}

func init() { proto.RegisterFile("createInstance.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2e, 0x4a, 0x4d,
	0x2c, 0x49, 0xf5, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x57, 0xb2, 0xe6, 0x12, 0x75, 0x46, 0x11, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52,
	0xe2, 0xe2, 0xc9, 0x84, 0x0a, 0xf9, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xa1, 0x88, 0x29, 0xd9, 0x70, 0x89, 0xa1, 0x6b, 0x2e, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x25, 0x46,
	0xb7, 0x91, 0x3f, 0x17, 0x07, 0x4c, 0x9f, 0x90, 0x33, 0x17, 0x1f, 0xaa, 0x49, 0x42, 0x62, 0x7a,
	0x58, 0xdd, 0x25, 0x25, 0xae, 0x87, 0xdd, 0x4a, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x97, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x5a, 0xf8, 0x81, 0xea, 0x00, 0x00, 0x00,
}