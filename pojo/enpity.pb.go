// Copyright (c) 2020, pole-group. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: enpity.proto

package pojo

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CheckType int32

const (
	CheckType_TCP       CheckType = 0
	CheckType_HTTP      CheckType = 1
	CheckType_CUSTOMER  CheckType = 2
	CheckType_HeartBeat CheckType = 3
)

// Enum value maps for CheckType.
var (
	CheckType_name = map[int32]string{
		0: "TCP",
		1: "HTTP",
		2: "CUSTOMER",
		3: "HeartBeat",
	}
	CheckType_value = map[string]int32{
		"TCP":       0,
		"HTTP":      1,
		"CUSTOMER":  2,
		"HeartBeat": 3,
	}
)

func (x CheckType) Enum() *CheckType {
	p := new(CheckType)
	*p = x
	return p
}

func (x CheckType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckType) Descriptor() protoreflect.EnumDescriptor {
	return file_enpity_proto_enumTypes[0].Descriptor()
}

func (CheckType) Type() protoreflect.EnumType {
	return &file_enpity_proto_enumTypes[0]
}

func (x CheckType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckType.Descriptor instead.
func (CheckType) EnumDescriptor() ([]byte, []int) {
	return file_enpity_proto_rawDescGZIP(), []int{0}
}

type FileType int32

const (
	FileType_Yaml       FileType = 0
	FileType_Json       FileType = 1
	FileType_Properties FileType = 2
	FileType_Toml       FileType = 3
	FileType_Xml        FileType = 4
	FileType_Conf       FileType = 5
)

// Enum value maps for FileType.
var (
	FileType_name = map[int32]string{
		0: "Yaml",
		1: "Json",
		2: "Properties",
		3: "Toml",
		4: "Xml",
		5: "Conf",
	}
	FileType_value = map[string]int32{
		"Yaml":       0,
		"Json":       1,
		"Properties": 2,
		"Toml":       3,
		"Xml":        4,
		"Conf":       5,
	}
)

func (x FileType) Enum() *FileType {
	p := new(FileType)
	*p = x
	return p
}

func (x FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_enpity_proto_enumTypes[1].Descriptor()
}

func (FileType) Type() protoreflect.EnumType {
	return &file_enpity_proto_enumTypes[1]
}

func (x FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileType.Descriptor instead.
func (FileType) EnumDescriptor() ([]byte, []int) {
	return file_enpity_proto_rawDescGZIP(), []int{1}
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName      string           `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Group            string           `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Metadata         *ServiceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ProtectThreshold float64          `protobuf:"fixed64,5,opt,name=protectThreshold,proto3" json:"protectThreshold,omitempty"`
	Selector         string           `protobuf:"bytes,6,opt,name=selector,proto3" json:"selector,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enpity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_enpity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_enpity_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Service) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Service) GetMetadata() *ServiceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Service) GetProtectThreshold() float64 {
	if x != nil {
		return x.ProtectThreshold
	}
	return 0
}

func (x *Service) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

type ServiceMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ServiceMetadata) Reset() {
	*x = ServiceMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enpity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMetadata) ProtoMessage() {}

func (x *ServiceMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_enpity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMetadata.ProtoReflect.Descriptor instead.
func (*ServiceMetadata) Descriptor() ([]byte, []int) {
	return file_enpity_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceMetadata) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string           `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Group       string           `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	ClusterName string           `protobuf:"bytes,3,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Metadata    *ClusterMetadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enpity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_enpity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_enpity_proto_rawDescGZIP(), []int{2}
}

func (x *Cluster) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Cluster) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Cluster) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *Cluster) GetMetadata() *ClusterMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ClusterMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ClusterMetadata) Reset() {
	*x = ClusterMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enpity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterMetadata) ProtoMessage() {}

func (x *ClusterMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_enpity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterMetadata.ProtoReflect.Descriptor instead.
func (*ClusterMetadata) Descriptor() ([]byte, []int) {
	return file_enpity_proto_rawDescGZIP(), []int{3}
}

func (x *ClusterMetadata) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName     string            `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Group           string            `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Ip              string            `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port            int64             `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	ClusterName     string            `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Weight          float64           `protobuf:"fixed64,6,opt,name=weight,proto3" json:"weight,omitempty"`
	Metadata        *InstanceMetadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Ephemeral       bool              `protobuf:"varint,8,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
	Enabled         bool              `protobuf:"varint,9,opt,name=enabled,proto3" json:"enabled,omitempty"`
	HealthCheckType CheckType         `protobuf:"varint,10,opt,name=healthCheckType,proto3,enum=pojo.CheckType" json:"healthCheckType,omitempty"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enpity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_enpity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_enpity_proto_rawDescGZIP(), []int{4}
}

func (x *Instance) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Instance) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Instance) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Instance) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Instance) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *Instance) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Instance) GetMetadata() *InstanceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Instance) GetEphemeral() bool {
	if x != nil {
		return x.Ephemeral
	}
	return false
}

func (x *Instance) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Instance) GetHealthCheckType() CheckType {
	if x != nil {
		return x.HealthCheckType
	}
	return CheckType_TCP
}

type InstanceMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InstanceMetadata) Reset() {
	*x = InstanceMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enpity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceMetadata) ProtoMessage() {}

func (x *InstanceMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_enpity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceMetadata.ProtoReflect.Descriptor instead.
func (*InstanceMetadata) Descriptor() ([]byte, []int) {
	return file_enpity_proto_rawDescGZIP(), []int{5}
}

func (x *InstanceMetadata) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ConfigFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group    string   `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	FileName string   `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FileType FileType `protobuf:"varint,3,opt,name=fileType,proto3,enum=pojo.FileType" json:"fileType,omitempty"`
	Content  string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ConfigFile) Reset() {
	*x = ConfigFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enpity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigFile) ProtoMessage() {}

func (x *ConfigFile) ProtoReflect() protoreflect.Message {
	mi := &file_enpity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigFile.ProtoReflect.Descriptor instead.
func (*ConfigFile) Descriptor() ([]byte, []int) {
	return file_enpity_proto_rawDescGZIP(), []int{6}
}

func (x *ConfigFile) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *ConfigFile) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ConfigFile) GetFileType() FileType {
	if x != nil {
		return x.FileType
	}
	return FileType_Yaml
}

func (x *ConfigFile) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_enpity_proto protoreflect.FileDescriptor

var file_enpity_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x6e, 0x70, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x70, 0x6f, 0x6a, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x31, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6f, 0x6a,
	0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x10, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x6f, 0x6a, 0x6f,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x96, 0x01, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x6f, 0x6a, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8f,
	0x01, 0x0a, 0x0f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x3f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x6f, 0x6a, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xc7, 0x02, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6f, 0x6a, 0x6f, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d,
	0x65, 0x72, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x70, 0x68, 0x65,
	0x6d, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x39, 0x0a, 0x0f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x6a, 0x6f, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x10, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x70, 0x6f, 0x6a, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x84,
	0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x6a, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0x3b, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x43, 0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48,
	0x54, 0x54, 0x50, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45,
	0x52, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74,
	0x10, 0x03, 0x2a, 0x4b, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x59, 0x61, 0x6d, 0x6c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x73, 0x6f, 0x6e,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x6f, 0x6d, 0x6c, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03,
	0x58, 0x6d, 0x6c, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x10, 0x05, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enpity_proto_rawDescOnce sync.Once
	file_enpity_proto_rawDescData = file_enpity_proto_rawDesc
)

func file_enpity_proto_rawDescGZIP() []byte {
	file_enpity_proto_rawDescOnce.Do(func() {
		file_enpity_proto_rawDescData = protoimpl.X.CompressGZIP(file_enpity_proto_rawDescData)
	})
	return file_enpity_proto_rawDescData
}

var file_enpity_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_enpity_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_enpity_proto_goTypes = []interface{}{
	(CheckType)(0),           // 0: pojo.CheckType
	(FileType)(0),            // 1: pojo.FileType
	(*Service)(nil),          // 2: pojo.Service
	(*ServiceMetadata)(nil),  // 3: pojo.ServiceMetadata
	(*Cluster)(nil),          // 4: pojo.Cluster
	(*ClusterMetadata)(nil),  // 5: pojo.ClusterMetadata
	(*Instance)(nil),         // 6: pojo.Instance
	(*InstanceMetadata)(nil), // 7: pojo.InstanceMetadata
	(*ConfigFile)(nil),       // 8: pojo.ConfigFile
	nil,                      // 9: pojo.ServiceMetadata.MetadataEntry
	nil,                      // 10: pojo.ClusterMetadata.MetadataEntry
	nil,                      // 11: pojo.InstanceMetadata.MetadataEntry
}
var file_enpity_proto_depIdxs = []int32{
	3,  // 0: pojo.Service.metadata:type_name -> pojo.ServiceMetadata
	9,  // 1: pojo.ServiceMetadata.metadata:type_name -> pojo.ServiceMetadata.MetadataEntry
	5,  // 2: pojo.Cluster.metadata:type_name -> pojo.ClusterMetadata
	10, // 3: pojo.ClusterMetadata.metadata:type_name -> pojo.ClusterMetadata.MetadataEntry
	7,  // 4: pojo.Instance.metadata:type_name -> pojo.InstanceMetadata
	0,  // 5: pojo.Instance.healthCheckType:type_name -> pojo.CheckType
	11, // 6: pojo.InstanceMetadata.metadata:type_name -> pojo.InstanceMetadata.MetadataEntry
	1,  // 7: pojo.ConfigFile.fileType:type_name -> pojo.FileType
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_enpity_proto_init() }
func file_enpity_proto_init() {
	if File_enpity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enpity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_enpity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_enpity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_enpity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_enpity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_enpity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_enpity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigFile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enpity_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enpity_proto_goTypes,
		DependencyIndexes: file_enpity_proto_depIdxs,
		EnumInfos:         file_enpity_proto_enumTypes,
		MessageInfos:      file_enpity_proto_msgTypes,
	}.Build()
	File_enpity_proto = out.File
	file_enpity_proto_rawDesc = nil
	file_enpity_proto_goTypes = nil
	file_enpity_proto_depIdxs = nil
}
