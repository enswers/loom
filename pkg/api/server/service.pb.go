// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package server is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Job
	WorkerInfo
	EmptyResponse
*/
package server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Job struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Topic string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Tasks string `protobuf:"bytes,3,opt,name=tasks" json:"tasks,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Job) GetTasks() string {
	if m != nil {
		return m.Tasks
	}
	return ""
}

type WorkerInfo struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Topic string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
}

func (m *WorkerInfo) Reset()                    { *m = WorkerInfo{} }
func (m *WorkerInfo) String() string            { return proto.CompactTextString(m) }
func (*WorkerInfo) ProtoMessage()               {}
func (*WorkerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WorkerInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WorkerInfo) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()                    { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()               {}
func (*EmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Job)(nil), "loom.server.Job")
	proto.RegisterType((*WorkerInfo)(nil), "loom.server.WorkerInfo")
	proto.RegisterType((*EmptyResponse)(nil), "loom.server.EmptyResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0xc9, 0xcf, 0xcf, 0xd5,
	0x03, 0x89, 0xa5, 0x16, 0x29, 0x39, 0x72, 0x31, 0x7b, 0xe5, 0x27, 0x09, 0xf1, 0x71, 0x31, 0x65,
	0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x89, 0x70, 0xb1, 0x96,
	0xe4, 0x17, 0x64, 0x26, 0x4b, 0x30, 0x81, 0x85, 0x20, 0x1c, 0xb0, 0x68, 0x62, 0x71, 0x76, 0xb1,
	0x04, 0x33, 0x54, 0x14, 0xc4, 0x51, 0x32, 0xe2, 0xe2, 0x0a, 0xcf, 0x2f, 0xca, 0x4e, 0x2d, 0xf2,
	0xcc, 0x4b, 0xcb, 0x27, 0xce, 0x24, 0x25, 0x7e, 0x2e, 0x5e, 0xd7, 0xdc, 0x82, 0x92, 0xca, 0xa0,
	0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa3, 0xdd, 0x8c, 0x5c, 0x2c, 0x3e, 0xf9, 0xf9, 0xb9,
	0x42, 0x96, 0x5c, 0x3c, 0xc1, 0xa5, 0x49, 0xc5, 0xc9, 0x45, 0x99, 0x49, 0xa9, 0x20, 0x97, 0x89,
	0xeb, 0x21, 0x39, 0x57, 0x0f, 0x61, 0x91, 0x94, 0x00, 0x8a, 0x04, 0x48, 0xa9, 0x25, 0x17, 0x67,
	0x50, 0x6a, 0x41, 0x7e, 0x51, 0x09, 0x88, 0x83, 0x21, 0x2d, 0x25, 0x85, 0x22, 0x82, 0x62, 0xbd,
	0x90, 0x2d, 0x17, 0x2f, 0x5c, 0xab, 0x4b, 0x7e, 0x5e, 0x2a, 0x69, 0xda, 0x9d, 0x38, 0xa2, 0xd8,
	0x20, 0xe2, 0x49, 0x6c, 0xe0, 0x30, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x28, 0xea,
	0x25, 0x74, 0x01, 0x00, 0x00,
}
