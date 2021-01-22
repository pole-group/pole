// Copyright (c) 2020, Conf-Group. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: raft.proto

package proto

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

type EntryMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    int64     `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Type    EntryType `protobuf:"varint,2,opt,name=type,proto3,enum=proto.EntryType" json:"type,omitempty"`
	Peers   []string  `protobuf:"bytes,3,rep,name=peers,proto3" json:"peers,omitempty"`
	DataLen int64     `protobuf:"varint,4,opt,name=dataLen,proto3" json:"dataLen,omitempty"`
	// Don't change field id of `old_peers' in the consideration of backward
	// compatibility
	OldPeers []string `protobuf:"bytes,5,rep,name=oldPeers,proto3" json:"oldPeers,omitempty"`
	// Checksum fot this log entry, since 1.2.6, added by boyan@antfin.com
	Checksum    int64    `protobuf:"varint,6,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Learners    []string `protobuf:"bytes,7,rep,name=learners,proto3" json:"learners,omitempty"`
	OldLearners []string `protobuf:"bytes,8,rep,name=oldLearners,proto3" json:"oldLearners,omitempty"`
}

func (x *EntryMeta) Reset() {
	*x = EntryMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryMeta) ProtoMessage() {}

func (x *EntryMeta) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryMeta.ProtoReflect.Descriptor instead.
func (*EntryMeta) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{0}
}

func (x *EntryMeta) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *EntryMeta) GetType() EntryType {
	if x != nil {
		return x.Type
	}
	return EntryType_EntryTypeUnKnow
}

func (x *EntryMeta) GetPeers() []string {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *EntryMeta) GetDataLen() int64 {
	if x != nil {
		return x.DataLen
	}
	return 0
}

func (x *EntryMeta) GetOldPeers() []string {
	if x != nil {
		return x.OldPeers
	}
	return nil
}

func (x *EntryMeta) GetChecksum() int64 {
	if x != nil {
		return x.Checksum
	}
	return 0
}

func (x *EntryMeta) GetLearners() []string {
	if x != nil {
		return x.Learners
	}
	return nil
}

func (x *EntryMeta) GetOldLearners() []string {
	if x != nil {
		return x.OldLearners
	}
	return nil
}

type SnapshotMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastIncludedIndex int64    `protobuf:"varint,1,opt,name=lastIncludedIndex,proto3" json:"lastIncludedIndex,omitempty"`
	LastIncludedTerm  int64    `protobuf:"varint,2,opt,name=lastIncludedTerm,proto3" json:"lastIncludedTerm,omitempty"`
	Peers             []string `protobuf:"bytes,3,rep,name=peers,proto3" json:"peers,omitempty"`
	OldPeers          []string `protobuf:"bytes,4,rep,name=oldPeers,proto3" json:"oldPeers,omitempty"`
	Learners          []string `protobuf:"bytes,5,rep,name=learners,proto3" json:"learners,omitempty"`
	OldLearners       []string `protobuf:"bytes,6,rep,name=oldLearners,proto3" json:"oldLearners,omitempty"`
}

func (x *SnapshotMeta) Reset() {
	*x = SnapshotMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotMeta) ProtoMessage() {}

func (x *SnapshotMeta) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotMeta.ProtoReflect.Descriptor instead.
func (*SnapshotMeta) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotMeta) GetLastIncludedIndex() int64 {
	if x != nil {
		return x.LastIncludedIndex
	}
	return 0
}

func (x *SnapshotMeta) GetLastIncludedTerm() int64 {
	if x != nil {
		return x.LastIncludedTerm
	}
	return 0
}

func (x *SnapshotMeta) GetPeers() []string {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *SnapshotMeta) GetOldPeers() []string {
	if x != nil {
		return x.OldPeers
	}
	return nil
}

func (x *SnapshotMeta) GetLearners() []string {
	if x != nil {
		return x.Learners
	}
	return nil
}

func (x *SnapshotMeta) GetOldLearners() []string {
	if x != nil {
		return x.OldLearners
	}
	return nil
}

var File_raft_proto protoreflect.FileDescriptor

var file_raft_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f,
	0x72, 0x65, 0x1a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea,
	0x01, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x61, 0x74, 0x61, 0x4c, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x61,
	0x74, 0x61, 0x4c, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x50, 0x65, 0x65, 0x72,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x50, 0x65, 0x65, 0x72,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x6c, 0x64,
	0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x6c, 0x64, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x22, 0xd8, 0x01, 0x0a, 0x0c,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x11,
	0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x61,
	0x73, 0x74, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x64, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x6c, 0x64, 0x50, 0x65, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x6c, 0x64, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x6c, 0x64, 0x4c, 0x65, 0x61, 0x72, 0x6e,
	0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x4c, 0x65,
	0x61, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raft_proto_rawDescOnce sync.Once
	file_raft_proto_rawDescData = file_raft_proto_rawDesc
)

func file_raft_proto_rawDescGZIP() []byte {
	file_raft_proto_rawDescOnce.Do(func() {
		file_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_raft_proto_rawDescData)
	})
	return file_raft_proto_rawDescData
}

var file_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_raft_proto_goTypes = []interface{}{
	(*EntryMeta)(nil),    // 0: proto.EntryMeta
	(*SnapshotMeta)(nil), // 1: proto.SnapshotMeta
	(EntryType)(0),       // 2: proto.EntryType
}
var file_raft_proto_depIdxs = []int32{
	2, // 0: proto.EntryMeta.type:type_name -> proto.EntryType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_raft_proto_init() }
func file_raft_proto_init() {
	if File_raft_proto != nil {
		return
	}
	file_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_raft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryMeta); i {
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
		file_raft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotMeta); i {
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
			RawDescriptor: file_raft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_raft_proto_goTypes,
		DependencyIndexes: file_raft_proto_depIdxs,
		MessageInfos:      file_raft_proto_msgTypes,
	}.Build()
	File_raft_proto = out.File
	file_raft_proto_rawDesc = nil
	file_raft_proto_goTypes = nil
	file_raft_proto_depIdxs = nil
}
