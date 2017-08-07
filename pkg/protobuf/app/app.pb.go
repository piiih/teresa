// Code generated by protoc-gen-go.
// source: pkg/protobuf/app/app.proto
// DO NOT EDIT!

/*
Package app is a generated protocol buffer package.

It is generated from these files:
	pkg/protobuf/app/app.proto

It has these top-level messages:
	CreateRequest
	ListResponse
	LogsRequest
	LogsResponse
	InfoRequest
	InfoResponse
	SetEnvRequest
	UnsetEnvRequest
	Empty
*/
package app

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

type CreateRequest struct {
	Name        string                   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Team        string                   `protobuf:"bytes,2,opt,name=team" json:"team,omitempty"`
	ProcessType string                   `protobuf:"bytes,3,opt,name=process_type,json=processType" json:"process_type,omitempty"`
	Limits      *CreateRequest_Limits    `protobuf:"bytes,4,opt,name=limits" json:"limits,omitempty"`
	AutoScale   *CreateRequest_AutoScale `protobuf:"bytes,5,opt,name=auto_scale,json=autoScale" json:"auto_scale,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *CreateRequest) GetProcessType() string {
	if m != nil {
		return m.ProcessType
	}
	return ""
}

func (m *CreateRequest) GetLimits() *CreateRequest_Limits {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *CreateRequest) GetAutoScale() *CreateRequest_AutoScale {
	if m != nil {
		return m.AutoScale
	}
	return nil
}

type CreateRequest_Limits struct {
	Default        []*CreateRequest_Limits_LimitRangeQuantity `protobuf:"bytes,1,rep,name=default" json:"default,omitempty"`
	DefaultRequest []*CreateRequest_Limits_LimitRangeQuantity `protobuf:"bytes,2,rep,name=default_request,json=defaultRequest" json:"default_request,omitempty"`
}

func (m *CreateRequest_Limits) Reset()                    { *m = CreateRequest_Limits{} }
func (m *CreateRequest_Limits) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest_Limits) ProtoMessage()               {}
func (*CreateRequest_Limits) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CreateRequest_Limits) GetDefault() []*CreateRequest_Limits_LimitRangeQuantity {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *CreateRequest_Limits) GetDefaultRequest() []*CreateRequest_Limits_LimitRangeQuantity {
	if m != nil {
		return m.DefaultRequest
	}
	return nil
}

type CreateRequest_Limits_LimitRangeQuantity struct {
	Quantity string `protobuf:"bytes,1,opt,name=quantity" json:"quantity,omitempty"`
	Resource string `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
}

func (m *CreateRequest_Limits_LimitRangeQuantity) Reset() {
	*m = CreateRequest_Limits_LimitRangeQuantity{}
}
func (m *CreateRequest_Limits_LimitRangeQuantity) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_Limits_LimitRangeQuantity) ProtoMessage()    {}
func (*CreateRequest_Limits_LimitRangeQuantity) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

func (m *CreateRequest_Limits_LimitRangeQuantity) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *CreateRequest_Limits_LimitRangeQuantity) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type CreateRequest_AutoScale struct {
	CpuTargetUtilization int32 `protobuf:"varint,1,opt,name=cpu_target_utilization,json=cpuTargetUtilization" json:"cpu_target_utilization,omitempty"`
	Max                  int32 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	Min                  int32 `protobuf:"varint,3,opt,name=min" json:"min,omitempty"`
}

func (m *CreateRequest_AutoScale) Reset()                    { *m = CreateRequest_AutoScale{} }
func (m *CreateRequest_AutoScale) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest_AutoScale) ProtoMessage()               {}
func (*CreateRequest_AutoScale) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *CreateRequest_AutoScale) GetCpuTargetUtilization() int32 {
	if m != nil {
		return m.CpuTargetUtilization
	}
	return 0
}

func (m *CreateRequest_AutoScale) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *CreateRequest_AutoScale) GetMin() int32 {
	if m != nil {
		return m.Min
	}
	return 0
}

type ListResponse struct {
	Apps []*ListResponse_App `protobuf:"bytes,1,rep,name=apps" json:"apps,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListResponse) GetApps() []*ListResponse_App {
	if m != nil {
		return m.Apps
	}
	return nil
}

type ListResponse_App struct {
	Team string   `protobuf:"bytes,1,opt,name=team" json:"team,omitempty"`
	Name string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Urls []string `protobuf:"bytes,3,rep,name=urls" json:"urls,omitempty"`
}

func (m *ListResponse_App) Reset()                    { *m = ListResponse_App{} }
func (m *ListResponse_App) String() string            { return proto.CompactTextString(m) }
func (*ListResponse_App) ProtoMessage()               {}
func (*ListResponse_App) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *ListResponse_App) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *ListResponse_App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListResponse_App) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

type LogsRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Lines  int64  `protobuf:"varint,2,opt,name=lines" json:"lines,omitempty"`
	Follow bool   `protobuf:"varint,3,opt,name=follow" json:"follow,omitempty"`
}

func (m *LogsRequest) Reset()                    { *m = LogsRequest{} }
func (m *LogsRequest) String() string            { return proto.CompactTextString(m) }
func (*LogsRequest) ProtoMessage()               {}
func (*LogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogsRequest) GetLines() int64 {
	if m != nil {
		return m.Lines
	}
	return 0
}

func (m *LogsRequest) GetFollow() bool {
	if m != nil {
		return m.Follow
	}
	return false
}

type LogsResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *LogsResponse) Reset()                    { *m = LogsResponse{} }
func (m *LogsResponse) String() string            { return proto.CompactTextString(m) }
func (*LogsResponse) ProtoMessage()               {}
func (*LogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LogsResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type InfoRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *InfoRequest) Reset()                    { *m = InfoRequest{} }
func (m *InfoRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()               {}
func (*InfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type InfoResponse struct {
	Team      string                  `protobuf:"bytes,1,opt,name=team" json:"team,omitempty"`
	Addresses []*InfoResponse_Address `protobuf:"bytes,2,rep,name=addresses" json:"addresses,omitempty"`
	EnvVars   []*InfoResponse_EnvVar  `protobuf:"bytes,3,rep,name=env_vars,json=envVars" json:"env_vars,omitempty"`
	Status    *InfoResponse_Status    `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	AutoScale *InfoResponse_AutoScale `protobuf:"bytes,5,opt,name=auto_scale,json=autoScale" json:"auto_scale,omitempty"`
	Limits    *InfoResponse_Limits    `protobuf:"bytes,6,opt,name=limits" json:"limits,omitempty"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *InfoResponse) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *InfoResponse) GetAddresses() []*InfoResponse_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *InfoResponse) GetEnvVars() []*InfoResponse_EnvVar {
	if m != nil {
		return m.EnvVars
	}
	return nil
}

func (m *InfoResponse) GetStatus() *InfoResponse_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *InfoResponse) GetAutoScale() *InfoResponse_AutoScale {
	if m != nil {
		return m.AutoScale
	}
	return nil
}

func (m *InfoResponse) GetLimits() *InfoResponse_Limits {
	if m != nil {
		return m.Limits
	}
	return nil
}

type InfoResponse_Address struct {
	Hostname string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
}

func (m *InfoResponse_Address) Reset()                    { *m = InfoResponse_Address{} }
func (m *InfoResponse_Address) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_Address) ProtoMessage()               {}
func (*InfoResponse_Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *InfoResponse_Address) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type InfoResponse_EnvVar struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *InfoResponse_EnvVar) Reset()                    { *m = InfoResponse_EnvVar{} }
func (m *InfoResponse_EnvVar) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_EnvVar) ProtoMessage()               {}
func (*InfoResponse_EnvVar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 1} }

func (m *InfoResponse_EnvVar) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *InfoResponse_EnvVar) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type InfoResponse_Status struct {
	Cpu  int32                      `protobuf:"varint,1,opt,name=cpu" json:"cpu,omitempty"`
	Pods []*InfoResponse_Status_Pod `protobuf:"bytes,3,rep,name=pods" json:"pods,omitempty"`
}

func (m *InfoResponse_Status) Reset()                    { *m = InfoResponse_Status{} }
func (m *InfoResponse_Status) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_Status) ProtoMessage()               {}
func (*InfoResponse_Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 2} }

func (m *InfoResponse_Status) GetCpu() int32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *InfoResponse_Status) GetPods() []*InfoResponse_Status_Pod {
	if m != nil {
		return m.Pods
	}
	return nil
}

type InfoResponse_Status_Pod struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	State string `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
}

func (m *InfoResponse_Status_Pod) Reset()                    { *m = InfoResponse_Status_Pod{} }
func (m *InfoResponse_Status_Pod) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_Status_Pod) ProtoMessage()               {}
func (*InfoResponse_Status_Pod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 2, 0} }

func (m *InfoResponse_Status_Pod) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InfoResponse_Status_Pod) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type InfoResponse_AutoScale struct {
	CpuTargetUtilization int32 `protobuf:"varint,1,opt,name=cpu_target_utilization,json=cpuTargetUtilization" json:"cpu_target_utilization,omitempty"`
	Max                  int32 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	Min                  int32 `protobuf:"varint,3,opt,name=min" json:"min,omitempty"`
}

func (m *InfoResponse_AutoScale) Reset()                    { *m = InfoResponse_AutoScale{} }
func (m *InfoResponse_AutoScale) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_AutoScale) ProtoMessage()               {}
func (*InfoResponse_AutoScale) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 3} }

func (m *InfoResponse_AutoScale) GetCpuTargetUtilization() int32 {
	if m != nil {
		return m.CpuTargetUtilization
	}
	return 0
}

func (m *InfoResponse_AutoScale) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *InfoResponse_AutoScale) GetMin() int32 {
	if m != nil {
		return m.Min
	}
	return 0
}

type InfoResponse_Limits struct {
	Default        []*InfoResponse_Limits_LimitRangeQuantity `protobuf:"bytes,1,rep,name=default" json:"default,omitempty"`
	DefaultRequest []*InfoResponse_Limits_LimitRangeQuantity `protobuf:"bytes,2,rep,name=default_request,json=defaultRequest" json:"default_request,omitempty"`
}

func (m *InfoResponse_Limits) Reset()                    { *m = InfoResponse_Limits{} }
func (m *InfoResponse_Limits) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_Limits) ProtoMessage()               {}
func (*InfoResponse_Limits) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 4} }

func (m *InfoResponse_Limits) GetDefault() []*InfoResponse_Limits_LimitRangeQuantity {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *InfoResponse_Limits) GetDefaultRequest() []*InfoResponse_Limits_LimitRangeQuantity {
	if m != nil {
		return m.DefaultRequest
	}
	return nil
}

type InfoResponse_Limits_LimitRangeQuantity struct {
	Quantity string `protobuf:"bytes,1,opt,name=quantity" json:"quantity,omitempty"`
	Resource string `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
}

func (m *InfoResponse_Limits_LimitRangeQuantity) Reset() {
	*m = InfoResponse_Limits_LimitRangeQuantity{}
}
func (m *InfoResponse_Limits_LimitRangeQuantity) String() string { return proto.CompactTextString(m) }
func (*InfoResponse_Limits_LimitRangeQuantity) ProtoMessage()    {}
func (*InfoResponse_Limits_LimitRangeQuantity) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 4, 0}
}

func (m *InfoResponse_Limits_LimitRangeQuantity) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *InfoResponse_Limits_LimitRangeQuantity) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type SetEnvRequest struct {
	Name    string                  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	EnvVars []*SetEnvRequest_EnvVar `protobuf:"bytes,2,rep,name=env_vars,json=envVars" json:"env_vars,omitempty"`
}

func (m *SetEnvRequest) Reset()                    { *m = SetEnvRequest{} }
func (m *SetEnvRequest) String() string            { return proto.CompactTextString(m) }
func (*SetEnvRequest) ProtoMessage()               {}
func (*SetEnvRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SetEnvRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetEnvRequest) GetEnvVars() []*SetEnvRequest_EnvVar {
	if m != nil {
		return m.EnvVars
	}
	return nil
}

type SetEnvRequest_EnvVar struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *SetEnvRequest_EnvVar) Reset()                    { *m = SetEnvRequest_EnvVar{} }
func (m *SetEnvRequest_EnvVar) String() string            { return proto.CompactTextString(m) }
func (*SetEnvRequest_EnvVar) ProtoMessage()               {}
func (*SetEnvRequest_EnvVar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *SetEnvRequest_EnvVar) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetEnvRequest_EnvVar) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UnsetEnvRequest struct {
	Name    string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	EnvVars []string `protobuf:"bytes,2,rep,name=env_vars,json=envVars" json:"env_vars,omitempty"`
}

func (m *UnsetEnvRequest) Reset()                    { *m = UnsetEnvRequest{} }
func (m *UnsetEnvRequest) String() string            { return proto.CompactTextString(m) }
func (*UnsetEnvRequest) ProtoMessage()               {}
func (*UnsetEnvRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UnsetEnvRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UnsetEnvRequest) GetEnvVars() []string {
	if m != nil {
		return m.EnvVars
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*CreateRequest)(nil), "app.CreateRequest")
	proto.RegisterType((*CreateRequest_Limits)(nil), "app.CreateRequest.Limits")
	proto.RegisterType((*CreateRequest_Limits_LimitRangeQuantity)(nil), "app.CreateRequest.Limits.LimitRangeQuantity")
	proto.RegisterType((*CreateRequest_AutoScale)(nil), "app.CreateRequest.AutoScale")
	proto.RegisterType((*ListResponse)(nil), "app.ListResponse")
	proto.RegisterType((*ListResponse_App)(nil), "app.ListResponse.App")
	proto.RegisterType((*LogsRequest)(nil), "app.LogsRequest")
	proto.RegisterType((*LogsResponse)(nil), "app.LogsResponse")
	proto.RegisterType((*InfoRequest)(nil), "app.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "app.InfoResponse")
	proto.RegisterType((*InfoResponse_Address)(nil), "app.InfoResponse.Address")
	proto.RegisterType((*InfoResponse_EnvVar)(nil), "app.InfoResponse.EnvVar")
	proto.RegisterType((*InfoResponse_Status)(nil), "app.InfoResponse.Status")
	proto.RegisterType((*InfoResponse_Status_Pod)(nil), "app.InfoResponse.Status.Pod")
	proto.RegisterType((*InfoResponse_AutoScale)(nil), "app.InfoResponse.AutoScale")
	proto.RegisterType((*InfoResponse_Limits)(nil), "app.InfoResponse.Limits")
	proto.RegisterType((*InfoResponse_Limits_LimitRangeQuantity)(nil), "app.InfoResponse.Limits.LimitRangeQuantity")
	proto.RegisterType((*SetEnvRequest)(nil), "app.SetEnvRequest")
	proto.RegisterType((*SetEnvRequest_EnvVar)(nil), "app.SetEnvRequest.EnvVar")
	proto.RegisterType((*UnsetEnvRequest)(nil), "app.UnsetEnvRequest")
	proto.RegisterType((*Empty)(nil), "app.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for App service

type AppClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error)
	Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (App_LogsClient, error)
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	SetEnv(ctx context.Context, in *SetEnvRequest, opts ...grpc.CallOption) (*Empty, error)
	UnsetEnv(ctx context.Context, in *UnsetEnvRequest, opts ...grpc.CallOption) (*Empty, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListResponse, error)
}

type appClient struct {
	cc *grpc.ClientConn
}

func NewAppClient(cc *grpc.ClientConn) AppClient {
	return &appClient{cc}
}

func (c *appClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/app.App/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (App_LogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_App_serviceDesc.Streams[0], c.cc, "/app.App/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &appLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type App_LogsClient interface {
	Recv() (*LogsResponse, error)
	grpc.ClientStream
}

type appLogsClient struct {
	grpc.ClientStream
}

func (x *appLogsClient) Recv() (*LogsResponse, error) {
	m := new(LogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/app.App/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) SetEnv(ctx context.Context, in *SetEnvRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/app.App/SetEnv", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) UnsetEnv(ctx context.Context, in *UnsetEnvRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/app.App/UnsetEnv", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/app.App/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for App service

type AppServer interface {
	Create(context.Context, *CreateRequest) (*Empty, error)
	Logs(*LogsRequest, App_LogsServer) error
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	SetEnv(context.Context, *SetEnvRequest) (*Empty, error)
	UnsetEnv(context.Context, *UnsetEnvRequest) (*Empty, error)
	List(context.Context, *Empty) (*ListResponse, error)
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AppServer).Logs(m, &appLogsServer{stream})
}

type App_LogsServer interface {
	Send(*LogsResponse) error
	grpc.ServerStream
}

type appLogsServer struct {
	grpc.ServerStream
}

func (x *appLogsServer) Send(m *LogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _App_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_SetEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).SetEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/SetEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).SetEnv(ctx, req.(*SetEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_UnsetEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsetEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).UnsetEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/UnsetEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).UnsetEnv(ctx, req.(*UnsetEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _App_Create_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _App_Info_Handler,
		},
		{
			MethodName: "SetEnv",
			Handler:    _App_SetEnv_Handler,
		},
		{
			MethodName: "UnsetEnv",
			Handler:    _App_UnsetEnv_Handler,
		},
		{
			MethodName: "List",
			Handler:    _App_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _App_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/protobuf/app/app.proto",
}

func init() { proto.RegisterFile("pkg/protobuf/app/app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 803 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x96, 0x63, 0xc7, 0x49, 0x4e, 0x52, 0x5a, 0x46, 0x4b, 0xe5, 0x9a, 0x5e, 0xa4, 0x96, 0x2a,
	0x05, 0xb5, 0x64, 0x43, 0x5a, 0x09, 0x09, 0x6e, 0x88, 0x50, 0x90, 0x90, 0x22, 0x51, 0x66, 0x77,
	0xb9, 0x8d, 0xa6, 0xc9, 0x24, 0x58, 0x75, 0x3c, 0x53, 0xcf, 0x38, 0x6c, 0x10, 0xbc, 0x00, 0x3c,
	0x09, 0x0f, 0xc6, 0x13, 0x20, 0xee, 0xd1, 0xfc, 0xd8, 0xb1, 0xb3, 0x49, 0x50, 0x91, 0xe8, 0x45,
	0xe4, 0x33, 0x67, 0xbe, 0xf3, 0x33, 0xdf, 0x9c, 0x73, 0x26, 0x10, 0xf2, 0x37, 0xeb, 0x4b, 0x9e,
	0x31, 0xc9, 0x5e, 0xe7, 0xab, 0x4b, 0xc2, 0xb9, 0xfa, 0x0d, 0xb5, 0x02, 0xb9, 0x84, 0xf3, 0xe8,
	0x0f, 0x0f, 0xee, 0x7d, 0x9d, 0x51, 0x22, 0x29, 0xa6, 0x6f, 0x73, 0x2a, 0x24, 0x42, 0xe0, 0xa5,
	0x64, 0x43, 0x03, 0xa7, 0xef, 0x0c, 0x3a, 0x58, 0xcb, 0x4a, 0x27, 0x29, 0xd9, 0x04, 0x0d, 0xa3,
	0x53, 0x32, 0x7a, 0x02, 0x3d, 0x9e, 0xb1, 0x05, 0x15, 0x62, 0x2e, 0x77, 0x9c, 0x06, 0xae, 0xde,
	0xeb, 0x5a, 0xdd, 0xf5, 0x8e, 0x53, 0xf4, 0x19, 0xf8, 0x49, 0xbc, 0x89, 0xa5, 0x08, 0xbc, 0xbe,
	0x33, 0xe8, 0x8e, 0x1f, 0x0d, 0x55, 0xf4, 0x5a, 0xb8, 0xe1, 0x4c, 0x03, 0xb0, 0x05, 0xa2, 0x2f,
	0x01, 0x48, 0x2e, 0xd9, 0x5c, 0x2c, 0x48, 0x42, 0x83, 0xa6, 0x36, 0x7b, 0x7c, 0xc4, 0x6c, 0x92,
	0x4b, 0x76, 0xa5, 0x30, 0xb8, 0x43, 0x0a, 0x31, 0xfc, 0xdb, 0x01, 0xdf, 0xf8, 0x43, 0xdf, 0x40,
	0x6b, 0x49, 0x57, 0x24, 0x4f, 0x64, 0xe0, 0xf4, 0xdd, 0x41, 0x77, 0xfc, 0xfc, 0x64, 0x6c, 0xf3,
	0xc1, 0x24, 0x5d, 0xd3, 0xef, 0x73, 0x92, 0xca, 0x58, 0xee, 0x70, 0x61, 0x8c, 0x6e, 0xe0, 0xbe,
	0x15, 0xe7, 0x99, 0xb1, 0x0a, 0x1a, 0xff, 0xc1, 0xdf, 0x07, 0xd6, 0x89, 0x45, 0x86, 0x33, 0x40,
	0x77, 0x51, 0x28, 0x84, 0xf6, 0x5b, 0x2b, 0x5b, 0xfa, 0xcb, 0xb5, 0xda, 0xcb, 0xa8, 0x60, 0x79,
	0xb6, 0xa0, 0xf6, 0x1a, 0xca, 0x75, 0x48, 0xa1, 0x53, 0xf2, 0x81, 0x5e, 0xc2, 0xc3, 0x05, 0xcf,
	0xe7, 0x92, 0x64, 0x6b, 0x2a, 0xe7, 0xb9, 0x8c, 0x93, 0xf8, 0x67, 0x22, 0x63, 0x96, 0x6a, 0x97,
	0x4d, 0x7c, 0xb1, 0xe0, 0xf9, 0xb5, 0xde, 0xbc, 0xd9, 0xef, 0xa1, 0x07, 0xe0, 0x6e, 0xc8, 0xad,
	0xf6, 0xdc, 0xc4, 0x4a, 0xd4, 0x9a, 0x38, 0xd5, 0xd7, 0xaa, 0x34, 0x71, 0x1a, 0xfd, 0x02, 0xbd,
	0x59, 0x2c, 0x24, 0xa6, 0x82, 0xb3, 0x54, 0x50, 0xf4, 0x09, 0x78, 0x84, 0x73, 0x61, 0x09, 0xfe,
	0x48, 0x13, 0x52, 0x05, 0x0c, 0x27, 0x9c, 0x63, 0x0d, 0x09, 0x27, 0xe0, 0x4e, 0x38, 0x2f, 0xeb,
	0xc8, 0xa9, 0xd4, 0x51, 0x51, 0x6f, 0x8d, 0x7a, 0xbd, 0xe5, 0x59, 0x22, 0x02, 0xb7, 0xef, 0x2a,
	0x9d, 0x92, 0xa3, 0xef, 0xa0, 0x3b, 0x63, 0x6b, 0x71, 0xae, 0x4c, 0x2f, 0xa0, 0x99, 0xc4, 0x29,
	0x15, 0xda, 0x97, 0x8b, 0xcd, 0x02, 0x3d, 0x04, 0x7f, 0xc5, 0x92, 0x84, 0xfd, 0xa4, 0xcf, 0xd2,
	0xc6, 0x76, 0x15, 0x45, 0xd0, 0x33, 0x0e, 0xed, 0x71, 0x74, 0x72, 0xb7, 0x72, 0x9f, 0xdc, 0xad,
	0x8c, 0x9e, 0x40, 0xf7, 0xdb, 0x74, 0xc5, 0xce, 0x04, 0x8d, 0xfe, 0xf4, 0xa1, 0x67, 0x30, 0x55,
	0x3f, 0x07, 0x87, 0xfc, 0x1c, 0x3a, 0x64, 0xb9, 0xcc, 0xa8, 0x10, 0x3a, 0x3b, 0xb7, 0x6c, 0x86,
	0xaa, 0xe5, 0x70, 0x62, 0x20, 0x78, 0x8f, 0x45, 0x2f, 0xa0, 0x4d, 0xd3, 0xed, 0x7c, 0x4b, 0x32,
	0xc3, 0x46, 0x77, 0x1c, 0xdc, 0xb5, 0x9b, 0xa6, 0xdb, 0x1f, 0x48, 0x86, 0x5b, 0x54, 0x7f, 0x05,
	0x1a, 0x81, 0x2f, 0x24, 0x91, 0x79, 0xd1, 0x77, 0x47, 0x4c, 0xae, 0xf4, 0x3e, 0xb6, 0x38, 0xf4,
	0xc5, 0x91, 0xb6, 0xfb, 0xf8, 0x48, 0x82, 0x47, 0xba, 0x4e, 0x45, 0xb3, 0x5d, 0xee, 0x9f, 0x8a,
	0x56, 0x6f, 0xf2, 0xf0, 0x29, 0xb4, 0xec, 0x51, 0x55, 0x59, 0xff, 0xc8, 0x84, 0xac, 0xb0, 0x5a,
	0xae, 0xc3, 0x11, 0xf8, 0xe6, 0x64, 0xaa, 0x16, 0xdf, 0xd0, 0xa2, 0x27, 0x94, 0xa8, 0xae, 0x7a,
	0x4b, 0x92, 0xbc, 0x28, 0x1b, 0xb3, 0x08, 0x7f, 0x05, 0xdf, 0x1c, 0x4c, 0x59, 0x2c, 0x78, 0x6e,
	0x4b, 0x5e, 0x89, 0x68, 0x04, 0x1e, 0x67, 0xcb, 0x82, 0xc5, 0xc7, 0xa7, 0x28, 0x19, 0xbe, 0x62,
	0x4b, 0xac, 0x91, 0xe1, 0x25, 0xb8, 0xaf, 0xd8, 0xf2, 0x54, 0xa5, 0x29, 0xe6, 0xca, 0xf0, 0x7a,
	0xf1, 0x9e, 0xfa, 0x30, 0xfc, 0x6b, 0x3f, 0xe6, 0xa6, 0x87, 0x63, 0xee, 0xd9, 0x29, 0xf2, 0xcf,
	0x4e, 0xb9, 0xeb, 0x53, 0x53, 0xee, 0x9d, 0xdc, 0xfd, 0xaf, 0x43, 0x2e, 0xfa, 0xdd, 0x81, 0x7b,
	0x57, 0x54, 0x4e, 0xd3, 0xed, 0xb9, 0x11, 0xf0, 0xb2, 0xd2, 0x2f, 0xd5, 0x3e, 0xab, 0x59, 0x1e,
	0x36, 0xcc, 0xbb, 0x57, 0x5a, 0xf4, 0x15, 0xdc, 0xbf, 0x49, 0xc5, 0xbf, 0xa6, 0xf3, 0xe8, 0x20,
	0x9d, 0x4e, 0x19, 0x33, 0x6a, 0x41, 0x73, 0xba, 0xe1, 0x72, 0x37, 0xfe, 0xad, 0x61, 0x86, 0xe3,
	0x00, 0x7c, 0xf3, 0x9c, 0x20, 0x74, 0xf7, 0x6d, 0x09, 0x41, 0xeb, 0xb4, 0x05, 0xfa, 0x14, 0x3c,
	0x35, 0xb9, 0xd0, 0x03, 0x33, 0x72, 0xf7, 0x53, 0x31, 0xfc, 0xb0, 0xa2, 0x31, 0xf7, 0x35, 0x72,
	0xd0, 0x33, 0xf0, 0xd4, 0x0d, 0x5a, 0x78, 0x65, 0x9e, 0x59, 0x78, 0x6d, 0x7a, 0x0d, 0xc0, 0x37,
	0x5c, 0xd9, 0x2c, 0x6a, 0xc4, 0xd5, 0xb2, 0x78, 0x0e, 0xed, 0x82, 0x02, 0x74, 0xa1, 0xf5, 0x07,
	0x8c, 0xd4, 0xd0, 0x4f, 0xc1, 0x53, 0x6f, 0x03, 0xaa, 0xe8, 0x8a, 0x6c, 0x2b, 0x4f, 0xc6, 0x6b,
	0x5f, 0xff, 0x37, 0x79, 0xf1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x16, 0x9f, 0xb7, 0xb9,
	0x08, 0x00, 0x00,
}
