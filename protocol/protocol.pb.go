// Copyright 2020 gorse Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: protocol.proto

package protocol

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

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model []byte `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *Model) GetModel() []byte {
	if x != nil {
		return x.Model
	}
	return nil
}

type Partition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Begin string `protobuf:"bytes,1,opt,name=begin,proto3" json:"begin,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Partition) Reset() {
	*x = Partition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partition) ProtoMessage() {}

func (x *Partition) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partition.ProtoReflect.Descriptor instead.
func (*Partition) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *Partition) GetBegin() string {
	if x != nil {
		return x.Begin
	}
	return ""
}

func (x *Partition) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_protocol_proto protoreflect.FileDescriptor

var file_protocol_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x1d, 0x0a, 0x05, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x33, 0x0a, 0x09, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x20,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0x95, 0x02, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12,
	0x37, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x16, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x16, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x17, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x65, 0x6e, 0x67, 0x68, 0x61, 0x6f, 0x7a,
	0x2f, 0x67, 0x6f, 0x72, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_proto_rawDescOnce sync.Once
	file_protocol_proto_rawDescData = file_protocol_proto_rawDesc
)

func file_protocol_proto_rawDescGZIP() []byte {
	file_protocol_proto_rawDescOnce.Do(func() {
		file_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_proto_rawDescData)
	})
	return file_protocol_proto_rawDescData
}

var file_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protocol_proto_goTypes = []interface{}{
	(*Model)(nil),     // 0: protocol.Model
	(*Partition)(nil), // 1: protocol.Partition
	(*Response)(nil),  // 2: protocol.Response
}
var file_protocol_proto_depIdxs = []int32{
	0, // 0: protocol.RouteGuide.BroadcastModel:input_type -> protocol.Model
	1, // 1: protocol.RouteGuide.BroadcastUserPartition:input_type -> protocol.Partition
	1, // 2: protocol.RouteGuide.BroadcastItemPartition:input_type -> protocol.Partition
	1, // 3: protocol.RouteGuide.BroadcastLabelPartition:input_type -> protocol.Partition
	2, // 4: protocol.RouteGuide.BroadcastModel:output_type -> protocol.Response
	2, // 5: protocol.RouteGuide.BroadcastUserPartition:output_type -> protocol.Response
	2, // 6: protocol.RouteGuide.BroadcastItemPartition:output_type -> protocol.Response
	2, // 7: protocol.RouteGuide.BroadcastLabelPartition:output_type -> protocol.Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocol_proto_init() }
func file_protocol_proto_init() {
	if File_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Partition); i {
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
		file_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_proto_goTypes,
		DependencyIndexes: file_protocol_proto_depIdxs,
		MessageInfos:      file_protocol_proto_msgTypes,
	}.Build()
	File_protocol_proto = out.File
	file_protocol_proto_rawDesc = nil
	file_protocol_proto_goTypes = nil
	file_protocol_proto_depIdxs = nil
}