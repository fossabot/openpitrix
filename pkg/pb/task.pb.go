// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/golang/protobuf/ptypes/wrappers"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateTaskRequest struct {
	X          *google_protobuf2.StringValue `protobuf:"bytes,1,opt,name=_" json:"_,omitempty"`
	JobId      *google_protobuf2.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	NodeId     *google_protobuf2.StringValue `protobuf:"bytes,3,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Target     *google_protobuf2.StringValue `protobuf:"bytes,4,opt,name=target" json:"target,omitempty"`
	TaskAction *google_protobuf2.StringValue `protobuf:"bytes,5,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Directive  *google_protobuf2.StringValue `protobuf:"bytes,6,opt,name=directive" json:"directive,omitempty"`
}

func (m *CreateTaskRequest) Reset()                    { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()               {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *CreateTaskRequest) GetX() *google_protobuf2.StringValue {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *CreateTaskRequest) GetJobId() *google_protobuf2.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *CreateTaskRequest) GetNodeId() *google_protobuf2.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *CreateTaskRequest) GetTarget() *google_protobuf2.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *CreateTaskRequest) GetTaskAction() *google_protobuf2.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *CreateTaskRequest) GetDirective() *google_protobuf2.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

type CreateTaskResponse struct {
	TaskId *google_protobuf2.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId  *google_protobuf2.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *CreateTaskResponse) Reset()                    { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()               {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *CreateTaskResponse) GetTaskId() *google_protobuf2.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *CreateTaskResponse) GetJobId() *google_protobuf2.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

type Task struct {
	TaskId     *google_protobuf2.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId      *google_protobuf2.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	TaskAction *google_protobuf2.StringValue `protobuf:"bytes,3,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Status     *google_protobuf2.StringValue `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	ErrorCode  *google_protobuf2.UInt32Value `protobuf:"bytes,5,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	Directive  *google_protobuf2.StringValue `protobuf:"bytes,6,opt,name=directive" json:"directive,omitempty"`
	Executor   *google_protobuf2.StringValue `protobuf:"bytes,7,opt,name=executor" json:"executor,omitempty"`
	Owner      *google_protobuf2.StringValue `protobuf:"bytes,8,opt,name=owner" json:"owner,omitempty"`
	Target     *google_protobuf2.StringValue `protobuf:"bytes,9,opt,name=target" json:"target,omitempty"`
	NodeId     *google_protobuf2.StringValue `protobuf:"bytes,10,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	CreateTime *google_protobuf3.Timestamp   `protobuf:"bytes,11,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime *google_protobuf3.Timestamp   `protobuf:"bytes,12,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *Task) GetTaskId() *google_protobuf2.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *Task) GetJobId() *google_protobuf2.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *Task) GetTaskAction() *google_protobuf2.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *Task) GetStatus() *google_protobuf2.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Task) GetErrorCode() *google_protobuf2.UInt32Value {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func (m *Task) GetDirective() *google_protobuf2.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *Task) GetExecutor() *google_protobuf2.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *Task) GetOwner() *google_protobuf2.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Task) GetTarget() *google_protobuf2.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Task) GetNodeId() *google_protobuf2.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Task) GetCreateTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Task) GetStatusTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type DescribeTasksRequest struct {
	TaskId   []string                      `protobuf:"bytes,1,rep,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId    []string                      `protobuf:"bytes,2,rep,name=job_id,json=jobId" json:"job_id,omitempty"`
	Executor *google_protobuf2.StringValue `protobuf:"bytes,3,opt,name=executor" json:"executor,omitempty"`
	Target   *google_protobuf2.StringValue `protobuf:"bytes,4,opt,name=target" json:"target,omitempty"`
	Status   []string                      `protobuf:"bytes,5,rep,name=status" json:"status,omitempty"`
	// default is 20, max value is 200
	Limit uint32 `protobuf:"varint,6,opt,name=limit" json:"limit,omitempty"`
	// default is 0
	Offset uint32 `protobuf:"varint,7,opt,name=offset" json:"offset,omitempty"`
}

func (m *DescribeTasksRequest) Reset()                    { *m = DescribeTasksRequest{} }
func (m *DescribeTasksRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeTasksRequest) ProtoMessage()               {}
func (*DescribeTasksRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *DescribeTasksRequest) GetTaskId() []string {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *DescribeTasksRequest) GetJobId() []string {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *DescribeTasksRequest) GetExecutor() *google_protobuf2.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *DescribeTasksRequest) GetTarget() *google_protobuf2.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *DescribeTasksRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeTasksRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeTasksRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type DescribeTasksResponse struct {
	TotalCount uint32  `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	TaskSet    []*Task `protobuf:"bytes,2,rep,name=task_set,json=taskSet" json:"task_set,omitempty"`
}

func (m *DescribeTasksResponse) Reset()                    { *m = DescribeTasksResponse{} }
func (m *DescribeTasksResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeTasksResponse) ProtoMessage()               {}
func (*DescribeTasksResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *DescribeTasksResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeTasksResponse) GetTaskSet() []*Task {
	if m != nil {
		return m.TaskSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTaskRequest)(nil), "openpitrix.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "openpitrix.CreateTaskResponse")
	proto.RegisterType((*Task)(nil), "openpitrix.Task")
	proto.RegisterType((*DescribeTasksRequest)(nil), "openpitrix.DescribeTasksRequest")
	proto.RegisterType((*DescribeTasksResponse)(nil), "openpitrix.DescribeTasksResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskManager service

type TaskManagerClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error)
}

type taskManagerClient struct {
	cc *grpc.ClientConn
}

func NewTaskManagerClient(cc *grpc.ClientConn) TaskManagerClient {
	return &taskManagerClient{cc}
}

func (c *taskManagerClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/CreateTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error) {
	out := new(DescribeTasksResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/DescribeTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskManager service

type TaskManagerServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	DescribeTasks(context.Context, *DescribeTasksRequest) (*DescribeTasksResponse, error)
}

func RegisterTaskManagerServer(s *grpc.Server, srv TaskManagerServer) {
	s.RegisterService(&_TaskManager_serviceDesc, srv)
}

func _TaskManager_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_DescribeTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).DescribeTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/DescribeTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).DescribeTasks(ctx, req.(*DescribeTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.TaskManager",
	HandlerType: (*TaskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskManager_CreateTask_Handler,
		},
		{
			MethodName: "DescribeTasks",
			Handler:    _TaskManager_DescribeTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}

func init() { proto.RegisterFile("task.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0xd2, 0xb8, 0xcd, 0x98, 0x48, 0xb0, 0x6a, 0xc1, 0x8a, 0x4a, 0x1b, 0x72, 0x2a,
	0x82, 0x26, 0x90, 0x16, 0x09, 0x51, 0x71, 0x28, 0xe5, 0x12, 0x21, 0x2e, 0x69, 0xe1, 0xc0, 0x25,
	0xda, 0xd8, 0x13, 0xb3, 0x6d, 0xea, 0x35, 0xbb, 0x93, 0xa6, 0x47, 0xc4, 0x81, 0x07, 0x80, 0x67,
	0xe1, 0x49, 0x38, 0x72, 0x84, 0xc7, 0xe0, 0x80, 0x76, 0xed, 0xb4, 0xa6, 0xa5, 0xe0, 0x14, 0x89,
	0x93, 0xe5, 0xd9, 0xef, 0xf7, 0xee, 0xce, 0x3f, 0x33, 0x06, 0x20, 0xae, 0x0f, 0x5a, 0x89, 0x92,
	0x24, 0x19, 0xc8, 0x04, 0xe3, 0x44, 0x90, 0x12, 0xc7, 0xf5, 0x95, 0x48, 0xca, 0x68, 0x84, 0x6d,
	0xbb, 0x32, 0x18, 0x0f, 0xdb, 0x13, 0xc5, 0x93, 0x04, 0x95, 0x4e, 0xd9, 0xfa, 0xea, 0xd9, 0x75,
	0x12, 0x87, 0xa8, 0x89, 0x1f, 0x26, 0x19, 0xb0, 0x9c, 0x01, 0x3c, 0x11, 0x6d, 0x1e, 0xc7, 0x92,
	0x38, 0x09, 0x19, 0x4f, 0xe5, 0xf7, 0xec, 0x23, 0x58, 0x8f, 0x30, 0x5e, 0xd7, 0x13, 0x1e, 0x45,
	0xa8, 0xda, 0x32, 0xb1, 0xc4, 0x79, 0xba, 0xf9, 0xb5, 0x04, 0xd7, 0x77, 0x14, 0x72, 0xc2, 0x3d,
	0xae, 0x0f, 0x7a, 0xf8, 0x76, 0x8c, 0x9a, 0xd8, 0x1d, 0x70, 0xfa, 0xbe, 0xd3, 0x70, 0xd6, 0xbc,
	0xce, 0x72, 0x2b, 0xdd, 0xad, 0x35, 0x3d, 0x4e, 0x6b, 0x97, 0x94, 0x88, 0xa3, 0x57, 0x7c, 0x34,
	0xc6, 0xde, 0x15, 0xb6, 0x01, 0xee, 0xbe, 0x1c, 0xf4, 0x45, 0xe8, 0x97, 0x0a, 0xf0, 0x95, 0x7d,
	0x39, 0xe8, 0x86, 0xec, 0x21, 0xcc, 0xc7, 0x32, 0x44, 0xa3, 0x2a, 0x17, 0x50, 0xb9, 0x06, 0xee,
	0x86, 0x6c, 0x13, 0x5c, 0xe2, 0x2a, 0x42, 0xf2, 0xe7, 0x8a, 0xa8, 0x52, 0x96, 0x3d, 0x01, 0xcf,
	0x38, 0xd1, 0xe7, 0x81, 0xb9, 0xb8, 0x5f, 0x29, 0x20, 0xb5, 0xd6, 0x6d, 0x5b, 0x9e, 0x3d, 0x86,
	0x6a, 0x28, 0x14, 0x06, 0x24, 0x8e, 0xd0, 0x77, 0x0b, 0x88, 0x4f, 0xf1, 0xe6, 0x3b, 0x07, 0x58,
	0x3e, 0xbb, 0x3a, 0x91, 0xb1, 0x46, 0x73, 0x7d, 0x7b, 0x22, 0x11, 0x16, 0x4a, 0xb2, 0x6b, 0xe0,
	0x6e, 0x78, 0xa9, 0x54, 0x37, 0x3f, 0x57, 0x60, 0xce, 0x6c, 0xfe, 0x3f, 0x37, 0x3d, 0x9b, 0xf2,
	0xf2, 0x8c, 0x29, 0xdf, 0x04, 0x57, 0x13, 0xa7, 0xb1, 0x2e, 0xe6, 0x73, 0xca, 0xb2, 0x2d, 0x00,
	0x54, 0x4a, 0xaa, 0x7e, 0x20, 0x43, 0xbc, 0xd0, 0xe6, 0x97, 0xdd, 0x98, 0x36, 0x3a, 0x99, 0x53,
	0x96, 0xdf, 0x91, 0x21, 0xfe, 0x8b, 0xcb, 0xec, 0x11, 0x2c, 0xe0, 0x31, 0x06, 0x63, 0x92, 0xca,
	0x9f, 0x2f, 0x20, 0x3d, 0xa1, 0x59, 0x07, 0x2a, 0x72, 0x12, 0xa3, 0xf2, 0x17, 0x8a, 0xe4, 0xd6,
	0xa2, 0xb9, 0x26, 0xa8, 0xce, 0xd0, 0x04, 0xb9, 0x8e, 0x83, 0x19, 0x3a, 0x6e, 0x0b, 0xbc, 0xc0,
	0xd6, 0x6f, 0xdf, 0x0c, 0x21, 0xdf, 0xb3, 0xd2, 0xfa, 0x39, 0xe9, 0xde, 0x74, 0x42, 0xf5, 0x20,
	0xc5, 0x4d, 0xc0, 0x88, 0x53, 0x6b, 0x52, 0xf1, 0xd5, 0xbf, 0x8b, 0x53, 0xdc, 0x04, 0x9a, 0x3f,
	0x1c, 0x58, 0x7c, 0x86, 0x3a, 0x50, 0x62, 0x60, 0x9b, 0x47, 0x4f, 0x67, 0xd3, 0xcd, 0x7c, 0x1d,
	0x97, 0xd7, 0xaa, 0x27, 0x95, 0xba, 0x94, 0xab, 0x54, 0x13, 0xcf, 0x6a, 0x31, 0xef, 0x4e, 0x79,
	0x26, 0x77, 0x2e, 0x37, 0x6e, 0x6e, 0x9c, 0x14, 0x6f, 0x25, 0x3d, 0x5e, 0x56, 0x9e, 0x8b, 0x50,
	0x19, 0x89, 0x43, 0x41, 0xb6, 0xba, 0x6a, 0xbd, 0xf4, 0xc5, 0xd0, 0x72, 0x38, 0xd4, 0x48, 0xb6,
	0x72, 0x6a, 0xbd, 0xec, 0xad, 0x89, 0xb0, 0x74, 0xe6, 0xf6, 0xd9, 0xec, 0x58, 0x05, 0x8f, 0x24,
	0xf1, 0x51, 0x3f, 0x90, 0xe3, 0x98, 0x6c, 0x2b, 0xd7, 0x7a, 0x60, 0x43, 0x3b, 0x26, 0xc2, 0xee,
	0xc2, 0x82, 0xcd, 0x8f, 0xf9, 0xa6, 0x49, 0x84, 0xd7, 0xb9, 0xd6, 0x3a, 0xfd, 0xfb, 0xb4, 0xec,
	0x20, 0xb2, 0x19, 0xdc, 0x45, 0xea, 0x7c, 0x73, 0xc0, 0x33, 0x91, 0x17, 0x3c, 0xe6, 0x11, 0x2a,
	0xf6, 0x1c, 0xe0, 0x74, 0x5e, 0xb1, 0x5b, 0x79, 0xe1, 0xb9, 0xbf, 0x44, 0x7d, 0xe5, 0xa2, 0xe5,
	0xec, 0xa8, 0x1f, 0x1c, 0xa8, 0xfd, 0x72, 0x09, 0xd6, 0xc8, 0x2b, 0x7e, 0xe7, 0x6e, 0xfd, 0xf6,
	0x1f, 0x88, 0xf4, 0xb3, 0xcd, 0xfb, 0x1f, 0xb7, 0x97, 0x59, 0x3d, 0xcc, 0xd6, 0x1a, 0xe6, 0x2a,
	0xba, 0x31, 0x11, 0xf4, 0xa6, 0x31, 0x14, 0x23, 0x42, 0xf5, 0xfe, 0xcb, 0xf7, 0x4f, 0x25, 0x8f,
	0x55, 0xdb, 0x47, 0x0f, 0xda, 0x76, 0xf1, 0xe9, 0xdc, 0xeb, 0x52, 0x32, 0x18, 0xb8, 0xd6, 0xb6,
	0x8d, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x5d, 0x40, 0xdc, 0x98, 0x07, 0x00, 0x00,
}
