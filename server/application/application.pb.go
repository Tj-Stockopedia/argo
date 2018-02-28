// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/application/application.proto

/*
Package application is a generated protocol buffer package.

Application Service

Application Service API performs CRUD actions against application resources

It is generated from these files:
	server/application/application.proto

It has these top-level messages:
	ApplicationQuery
	ApplicationResponse
	ApplicationSyncRequest
	ApplicationSyncResult
	ResourceDetails
*/
package application

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import k8s_io_api_core_v1 "k8s.io/api/core/v1"
import github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"

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

// ApplicationQuery is a query for application resources
type ApplicationQuery struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ApplicationQuery) Reset()                    { *m = ApplicationQuery{} }
func (m *ApplicationQuery) String() string            { return proto.CompactTextString(m) }
func (*ApplicationQuery) ProtoMessage()               {}
func (*ApplicationQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ApplicationQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplicationResponse struct {
}

func (m *ApplicationResponse) Reset()                    { *m = ApplicationResponse{} }
func (m *ApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*ApplicationResponse) ProtoMessage()               {}
func (*ApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// ApplicationSyncRequest is a request to apply the config state to live state
type ApplicationSyncRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DryRun bool   `protobuf:"varint,2,opt,name=dryRun" json:"dryRun,omitempty"`
}

func (m *ApplicationSyncRequest) Reset()                    { *m = ApplicationSyncRequest{} }
func (m *ApplicationSyncRequest) String() string            { return proto.CompactTextString(m) }
func (*ApplicationSyncRequest) ProtoMessage()               {}
func (*ApplicationSyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ApplicationSyncRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApplicationSyncRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

// ApplicationSyncResult is a result of a sync requeswt
type ApplicationSyncResult struct {
	Message   string             `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Resources []*ResourceDetails `protobuf:"bytes,2,rep,name=resources" json:"resources,omitempty"`
}

func (m *ApplicationSyncResult) Reset()                    { *m = ApplicationSyncResult{} }
func (m *ApplicationSyncResult) String() string            { return proto.CompactTextString(m) }
func (*ApplicationSyncResult) ProtoMessage()               {}
func (*ApplicationSyncResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ApplicationSyncResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ApplicationSyncResult) GetResources() []*ResourceDetails {
	if m != nil {
		return m.Resources
	}
	return nil
}

type ResourceDetails struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Kind      string `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Message   string `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
}

func (m *ResourceDetails) Reset()                    { *m = ResourceDetails{} }
func (m *ResourceDetails) String() string            { return proto.CompactTextString(m) }
func (*ResourceDetails) ProtoMessage()               {}
func (*ResourceDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ResourceDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceDetails) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *ResourceDetails) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ResourceDetails) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ApplicationQuery)(nil), "application.ApplicationQuery")
	proto.RegisterType((*ApplicationResponse)(nil), "application.ApplicationResponse")
	proto.RegisterType((*ApplicationSyncRequest)(nil), "application.ApplicationSyncRequest")
	proto.RegisterType((*ApplicationSyncResult)(nil), "application.ApplicationSyncResult")
	proto.RegisterType((*ResourceDetails)(nil), "application.ResourceDetails")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ApplicationService service

type ApplicationServiceClient interface {
	// List returns list of applications
	List(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationList, error)
	// Watch returns stream of application change events.
	Watch(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (ApplicationService_WatchClient, error)
	// Create creates an application
	Create(ctx context.Context, in *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error)
	// Get returns an application by name
	Get(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error)
	// Update updates an application
	Update(ctx context.Context, in *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error)
	// Delete deletes an application
	Delete(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (*ApplicationResponse, error)
	// ListPods returns pods in an application
	ListPods(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (*k8s_io_api_core_v1.PodList, error)
	// Sync syncs an application to its target state
	Sync(ctx context.Context, in *ApplicationSyncRequest, opts ...grpc.CallOption) (*ApplicationSyncResult, error)
}

type applicationServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServiceClient(cc *grpc.ClientConn) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) List(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationList, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationList)
	err := grpc.Invoke(ctx, "/application.ApplicationService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) Watch(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (ApplicationService_WatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ApplicationService_serviceDesc.Streams[0], c.cc, "/application.ApplicationService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApplicationService_WatchClient interface {
	Recv() (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationWatchEvent, error)
	grpc.ClientStream
}

type applicationServiceWatchClient struct {
	grpc.ClientStream
}

func (x *applicationServiceWatchClient) Recv() (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationWatchEvent, error) {
	m := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationWatchEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *applicationServiceClient) Create(ctx context.Context, in *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application)
	err := grpc.Invoke(ctx, "/application.ApplicationService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) Get(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application)
	err := grpc.Invoke(ctx, "/application.ApplicationService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) Update(ctx context.Context, in *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, opts ...grpc.CallOption) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error) {
	out := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application)
	err := grpc.Invoke(ctx, "/application.ApplicationService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) Delete(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (*ApplicationResponse, error) {
	out := new(ApplicationResponse)
	err := grpc.Invoke(ctx, "/application.ApplicationService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) ListPods(ctx context.Context, in *ApplicationQuery, opts ...grpc.CallOption) (*k8s_io_api_core_v1.PodList, error) {
	out := new(k8s_io_api_core_v1.PodList)
	err := grpc.Invoke(ctx, "/application.ApplicationService/ListPods", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) Sync(ctx context.Context, in *ApplicationSyncRequest, opts ...grpc.CallOption) (*ApplicationSyncResult, error) {
	out := new(ApplicationSyncResult)
	err := grpc.Invoke(ctx, "/application.ApplicationService/Sync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApplicationService service

type ApplicationServiceServer interface {
	// List returns list of applications
	List(context.Context, *ApplicationQuery) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationList, error)
	// Watch returns stream of application change events.
	Watch(*ApplicationQuery, ApplicationService_WatchServer) error
	// Create creates an application
	Create(context.Context, *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error)
	// Get returns an application by name
	Get(context.Context, *ApplicationQuery) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error)
	// Update updates an application
	Update(context.Context, *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application) (*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application, error)
	// Delete deletes an application
	Delete(context.Context, *ApplicationQuery) (*ApplicationResponse, error)
	// ListPods returns pods in an application
	ListPods(context.Context, *ApplicationQuery) (*k8s_io_api_core_v1.PodList, error)
	// Sync syncs an application to its target state
	Sync(context.Context, *ApplicationSyncRequest) (*ApplicationSyncResult, error)
}

func RegisterApplicationServiceServer(s *grpc.Server, srv ApplicationServiceServer) {
	s.RegisterService(&_ApplicationService_serviceDesc, srv)
}

func _ApplicationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).List(ctx, req.(*ApplicationQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ApplicationQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApplicationServiceServer).Watch(m, &applicationServiceWatchServer{stream})
}

type ApplicationService_WatchServer interface {
	Send(*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationWatchEvent) error
	grpc.ServerStream
}

type applicationServiceWatchServer struct {
	grpc.ServerStream
}

func (x *applicationServiceWatchServer) Send(m *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.ApplicationWatchEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ApplicationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Create(ctx, req.(*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Get(ctx, req.(*ApplicationQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Update(ctx, req.(*github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Application))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Delete(ctx, req.(*ApplicationQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_ListPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).ListPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/ListPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).ListPods(ctx, req.(*ApplicationQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).Sync(ctx, req.(*ApplicationSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "application.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ApplicationService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ApplicationService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ApplicationService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ApplicationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ApplicationService_Delete_Handler,
		},
		{
			MethodName: "ListPods",
			Handler:    _ApplicationService_ListPods_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _ApplicationService_Sync_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _ApplicationService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server/application/application.proto",
}

func init() { proto.RegisterFile("server/application/application.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xc1, 0x6e, 0x13, 0x3d,
	0x10, 0xc7, 0xb5, 0x6d, 0xbe, 0x7c, 0x8d, 0x7b, 0x00, 0x99, 0xb6, 0x0a, 0xdb, 0x54, 0x44, 0x6e,
	0x85, 0x42, 0x25, 0xbc, 0x4d, 0xb9, 0xa0, 0xde, 0x80, 0x02, 0x02, 0x71, 0x28, 0x8b, 0x10, 0x12,
	0x17, 0xe4, 0xee, 0x8e, 0xb6, 0x4b, 0x76, 0x6d, 0xd7, 0xf6, 0xae, 0x14, 0x21, 0x2e, 0xbc, 0x00,
	0x42, 0x3c, 0x00, 0xcf, 0xc1, 0x91, 0x23, 0x67, 0x5e, 0x81, 0x07, 0x41, 0x76, 0xb3, 0x64, 0xd3,
	0xa4, 0xc9, 0x25, 0x07, 0x6e, 0xb3, 0x9e, 0xd9, 0xf9, 0xff, 0x3c, 0x63, 0x8f, 0xd1, 0x9e, 0x06,
	0x55, 0x82, 0x0a, 0x98, 0x94, 0x59, 0x1a, 0x31, 0x93, 0x0a, 0x5e, 0xb7, 0xa9, 0x54, 0xc2, 0x08,
	0xbc, 0x5e, 0x5b, 0xf2, 0x37, 0x12, 0x91, 0x08, 0xb7, 0x1e, 0x58, 0xeb, 0x22, 0xc4, 0xef, 0x24,
	0x42, 0x24, 0x19, 0x04, 0x4c, 0xa6, 0x01, 0xe3, 0x5c, 0x18, 0x17, 0xac, 0x47, 0x5e, 0x32, 0xb8,
	0xaf, 0x69, 0x2a, 0x9c, 0x37, 0x12, 0x0a, 0x82, 0xb2, 0x1f, 0x24, 0xc0, 0x41, 0x31, 0x03, 0xf1,
	0x28, 0xe6, 0x59, 0x92, 0x9a, 0xb3, 0xe2, 0x94, 0x46, 0x22, 0x0f, 0x98, 0x72, 0x12, 0xef, 0x9d,
	0x71, 0x37, 0x8a, 0x03, 0x39, 0x48, 0xec, 0xcf, 0x7a, 0x02, 0xb4, 0xec, 0xb3, 0x4c, 0x9e, 0xb1,
	0xa9, 0x54, 0xe4, 0x36, 0xba, 0xfe, 0x60, 0x1c, 0xf7, 0xb2, 0x00, 0x35, 0xc4, 0x18, 0x35, 0x38,
	0xcb, 0xa1, 0xed, 0x75, 0xbd, 0x5e, 0x2b, 0x74, 0x36, 0xd9, 0x44, 0x37, 0x6a, 0x71, 0x21, 0x68,
	0x29, 0xb8, 0x06, 0x72, 0x8c, 0xb6, 0x6a, 0xcb, 0xaf, 0x86, 0x3c, 0x0a, 0xe1, 0xbc, 0x00, 0x6d,
	0x66, 0x25, 0xc1, 0x5b, 0xa8, 0x19, 0xab, 0x61, 0x58, 0xf0, 0xf6, 0x4a, 0xd7, 0xeb, 0xad, 0x85,
	0xa3, 0x2f, 0x92, 0xa3, 0xcd, 0xa9, 0x2c, 0xba, 0xc8, 0x0c, 0x6e, 0xa3, 0xff, 0x73, 0xd0, 0x9a,
	0x25, 0x55, 0x9e, 0xea, 0x13, 0x1f, 0xa1, 0x96, 0x02, 0x2d, 0x0a, 0x15, 0x81, 0x6e, 0xaf, 0x74,
	0x57, 0x7b, 0xeb, 0x87, 0x1d, 0x5a, 0x6f, 0x47, 0x38, 0xf2, 0x1e, 0x83, 0x61, 0x69, 0xa6, 0xc3,
	0x71, 0x38, 0x39, 0x47, 0xd7, 0x2e, 0x79, 0x67, 0xd2, 0x62, 0xd4, 0x18, 0xa4, 0x3c, 0x76, 0xac,
	0xad, 0xd0, 0xd9, 0xb8, 0x83, 0x5a, 0xd6, 0xa7, 0x25, 0x8b, 0xa0, 0xbd, 0xea, 0x1c, 0xe3, 0x85,
	0x3a, 0x6e, 0x63, 0x02, 0xf7, 0xf0, 0x47, 0x0b, 0xe1, 0xfa, 0x16, 0x41, 0x95, 0x69, 0x04, 0xf8,
	0xb3, 0x87, 0x1a, 0x2f, 0x52, 0x6d, 0xf0, 0xce, 0x04, 0xfb, 0xe5, 0x8e, 0xf8, 0xcf, 0xe9, 0xb8,
	0xe3, 0xb4, 0xea, 0xb8, 0x33, 0xde, 0x45, 0x31, 0x95, 0x83, 0x84, 0xda, 0x8e, 0x4f, 0xe4, 0xa8,
	0x3a, 0x5e, 0x4f, 0x66, 0xa5, 0x48, 0xe7, 0xd3, 0xaf, 0xdf, 0x5f, 0x57, 0xb6, 0xf0, 0x86, 0x3b,
	0x62, 0x65, 0xbf, 0x7e, 0x4e, 0x34, 0xfe, 0xe6, 0xa1, 0xff, 0xde, 0x30, 0x13, 0x9d, 0x2d, 0x42,
	0x3a, 0x59, 0x0e, 0x92, 0xd3, 0x7a, 0x5c, 0x02, 0x37, 0x64, 0xd7, 0x81, 0xed, 0xe0, 0xed, 0x0a,
	0x4c, 0x1b, 0x05, 0x2c, 0x9f, 0xe0, 0x3b, 0xf0, 0xf0, 0x77, 0x0f, 0x35, 0x1f, 0x29, 0x60, 0x06,
	0xf0, 0x93, 0xe5, 0x30, 0xf8, 0x4b, 0xca, 0x43, 0x6e, 0xb9, 0x1d, 0xdc, 0x24, 0x33, 0x4b, 0x7b,
	0xe4, 0xed, 0xe3, 0x2f, 0x1e, 0x5a, 0x7d, 0x0a, 0x0b, 0xdb, 0xbd, 0x2c, 0x9e, 0xa9, 0x8a, 0xd6,
	0x79, 0x82, 0x0f, 0xf6, 0xe0, 0x7e, 0xc4, 0x3f, 0x3d, 0xd4, 0x7c, 0x2d, 0xe3, 0x7f, 0xb1, 0x9e,
	0x81, 0xe3, 0xbf, 0xe3, 0xef, 0xcd, 0xe6, 0xcf, 0xc1, 0xb0, 0x98, 0x19, 0x46, 0xdd, 0x46, 0x6c,
	0x7d, 0x39, 0x6a, 0x1e, 0x43, 0x06, 0x06, 0x16, 0x55, 0xb8, 0x7b, 0x95, 0xfb, 0xef, 0x64, 0x1b,
	0xd5, 0x6e, 0x7f, 0x6e, 0xed, 0x24, 0x5a, 0xb3, 0x77, 0xea, 0x44, 0xc4, 0x7a, 0x91, 0xe2, 0x36,
	0xbd, 0x18, 0xec, 0x76, 0xeb, 0xd4, 0x0e, 0x76, 0x5a, 0xf6, 0xe9, 0x89, 0x88, 0xdd, 0x9d, 0xec,
	0x39, 0x31, 0x82, 0xbb, 0x73, 0xc4, 0x02, 0x69, 0x55, 0x86, 0xa8, 0x61, 0xe7, 0x23, 0xde, 0xbd,
	0x4a, 0xad, 0x36, 0x83, 0x7d, 0x32, 0x3f, 0xc8, 0x8e, 0xd8, 0x4a, 0x9a, 0xcc, 0x95, 0xd6, 0x43,
	0x1e, 0x3d, 0x3c, 0x78, 0x4b, 0xe7, 0xbd, 0x3b, 0xd3, 0xcf, 0xe3, 0x69, 0xd3, 0xbd, 0x31, 0xf7,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x3d, 0x28, 0x52, 0x3b, 0x07, 0x00, 0x00,
}
