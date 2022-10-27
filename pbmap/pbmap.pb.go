// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: pbmap.proto

package pbmap

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MethodMap gives the parameters how to Map a GRPC Service Method to an REST Endpoint
type MethodMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// "GET", "PUT", "POST", "DELETE"
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// Format "a/b/c/:d/:e" - a/b/c are fixed in the path, :d and :e are paramaters mapped to fields D, E in the underlying GRPC Request
	// Has to be compatible with Julian Schmidt's HTTP Router
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Format "key1=:var1&key=:var2" They keys will be mapped to the variables var1, var2 of the unterlying GRPC Request
	QueryString string `protobuf:"bytes,3,opt,name=query_string,json=queryString,proto3" json:"query_string,omitempty"`
	Body        string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *MethodMap) Reset() {
	*x = MethodMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodMap) ProtoMessage() {}

func (x *MethodMap) ProtoReflect() protoreflect.Message {
	mi := &file_pbmap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodMap.ProtoReflect.Descriptor instead.
func (*MethodMap) Descriptor() ([]byte, []int) {
	return file_pbmap_proto_rawDescGZIP(), []int{0}
}

func (x *MethodMap) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *MethodMap) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MethodMap) GetQueryString() string {
	if x != nil {
		return x.QueryString
	}
	return ""
}

func (x *MethodMap) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

// If an Submessage of an Request-Message in Protobuf has a field (directly at the moment, deeper nesting is not supported at the moment), then additional parameters defined in query (similar to the querystring syntax of your rest-framework) will be added to the queryString documentation and mapped automatically to the objects
type QueryMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *QueryMap) Reset() {
	*x = QueryMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMap) ProtoMessage() {}

func (x *QueryMap) ProtoReflect() protoreflect.Message {
	mi := &file_pbmap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMap.ProtoReflect.Descriptor instead.
func (*QueryMap) Descriptor() ([]byte, []int) {
	return file_pbmap_proto_rawDescGZIP(), []int{1}
}

func (x *QueryMap) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// Page provides information from the query string which you can use to get a larger collection of objects pagewise
type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_pbmap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_pbmap_proto_rawDescGZIP(), []int{2}
}

func (x *Page) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Page) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=Query,proto3" json:"Query,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbmap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_pbmap_proto_rawDescGZIP(), []int{3}
}

func (x *Filter) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// Sort provides information from the query string on how to sort a given collection of objects
type Sort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=Field,proto3" json:"Field,omitempty"`
	Desc  bool   `protobuf:"varint,2,opt,name=Desc,proto3" json:"Desc,omitempty"`
}

func (x *Sort) Reset() {
	*x = Sort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmap_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_pbmap_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_pbmap_proto_rawDescGZIP(), []int{4}
}

func (x *Sort) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Sort) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

// ServiceMap provides information for how to map an GRPC Service to an REST Service
type ServiceMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version       string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	BaseUri       string `protobuf:"bytes,2,opt,name=base_uri,json=baseUri,proto3" json:"base_uri,omitempty"`
	TargetPackage string `protobuf:"bytes,3,opt,name=target_package,json=targetPackage,proto3" json:"target_package,omitempty"`
}

func (x *ServiceMap) Reset() {
	*x = ServiceMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmap_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMap) ProtoMessage() {}

func (x *ServiceMap) ProtoReflect() protoreflect.Message {
	mi := &file_pbmap_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMap.ProtoReflect.Descriptor instead.
func (*ServiceMap) Descriptor() ([]byte, []int) {
	return file_pbmap_proto_rawDescGZIP(), []int{5}
}

func (x *ServiceMap) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServiceMap) GetBaseUri() string {
	if x != nil {
		return x.BaseUri
	}
	return ""
}

func (x *ServiceMap) GetTargetPackage() string {
	if x != nil {
		return x.TargetPackage
	}
	return ""
}

var file_pbmap_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodMap)(nil),
		Field:         87654321,
		Name:          "pbmap.method_map",
		Tag:           "bytes,87654321,opt,name=method_map",
		Filename:      "pbmap.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*QueryMap)(nil),
		Field:         9813242,
		Name:          "pbmap.query_map",
		Tag:           "bytes,9813242,opt,name=query_map",
		Filename:      "pbmap.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*ServiceMap)(nil),
		Field:         87654323,
		Name:          "pbmap.service_map",
		Tag:           "bytes,87654323,opt,name=service_map",
		Filename:      "pbmap.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         97454431,
		Name:          "pbmap.dart_package",
		Tag:           "bytes,97454431,opt,name=dart_package",
		Filename:      "pbmap.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional pbmap.MethodMap method_map = 87654321;
	E_MethodMap = &file_pbmap_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional pbmap.QueryMap query_map = 9813242;
	E_QueryMap = &file_pbmap_proto_extTypes[1]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional pbmap.ServiceMap service_map = 87654323;
	E_ServiceMap = &file_pbmap_proto_extTypes[2]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// optional string dart_package = 97454431;
	E_DartPackage = &file_pbmap_proto_extTypes[3]
)

var File_pbmap_proto protoreflect.FileDescriptor

var file_pbmap_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x62, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x62, 0x6d, 0x61, 0x70, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4d, 0x61, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x21, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x20, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d,
	0x61, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x62, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x3a, 0x2c,
	0xd2, 0xcf, 0xb7, 0x25, 0x27, 0x0a, 0x25, 0x70, 0x61, 0x67, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x3d, 0x3c, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x3e, 0x26, 0x70, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x3d, 0x3c, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3e, 0x22, 0x30, 0x0a, 0x06,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x10, 0xd2, 0xcf,
	0xb7, 0x25, 0x0b, 0x0a, 0x09, 0x71, 0x3d, 0x3c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x3e, 0x22, 0x51,
	0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x44, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x65, 0x73, 0x63,
	0x3a, 0x1f, 0xd2, 0xcf, 0xb7, 0x25, 0x1a, 0x0a, 0x18, 0x73, 0x6f, 0x72, 0x74, 0x3d, 0x3c, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x3e, 0x26, 0x64, 0x65, 0x73, 0x63, 0x3d, 0x3c, 0x44, 0x65, 0x73, 0x63,
	0x3e, 0x22, 0x68, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73,
	0x65, 0x55, 0x72, 0x69, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x3a, 0x52, 0x0a, 0x0a, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb1, 0xff, 0xe5, 0x29, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x6d, 0x61, 0x70, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x4d, 0x61, 0x70, 0x52, 0x09, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4d, 0x61, 0x70, 0x3a,
	0x50, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x70, 0x12, 0x1f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfa, 0xf9,
	0xd6, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x6d, 0x61, 0x70, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x70, 0x52, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61,
	0x70, 0x3a, 0x56, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x70,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xb3, 0xff, 0xe5, 0x29, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x6d,
	0x61, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x3a, 0x42, 0x0a, 0x0c, 0x64, 0x61, 0x72,
	0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdf, 0x92, 0xbc, 0x2e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x61, 0x72, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x42, 0x24, 0x5a,
	0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x6f, 0x7a,
	0x65, 0x72, 0x2d, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62,
	0x6d, 0x61, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbmap_proto_rawDescOnce sync.Once
	file_pbmap_proto_rawDescData = file_pbmap_proto_rawDesc
)

func file_pbmap_proto_rawDescGZIP() []byte {
	file_pbmap_proto_rawDescOnce.Do(func() {
		file_pbmap_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmap_proto_rawDescData)
	})
	return file_pbmap_proto_rawDescData
}

var file_pbmap_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pbmap_proto_goTypes = []interface{}{
	(*MethodMap)(nil),                   // 0: pbmap.MethodMap
	(*QueryMap)(nil),                    // 1: pbmap.QueryMap
	(*Page)(nil),                        // 2: pbmap.Page
	(*Filter)(nil),                      // 3: pbmap.Filter
	(*Sort)(nil),                        // 4: pbmap.Sort
	(*ServiceMap)(nil),                  // 5: pbmap.ServiceMap
	(*descriptorpb.MethodOptions)(nil),  // 6: google.protobuf.MethodOptions
	(*descriptorpb.MessageOptions)(nil), // 7: google.protobuf.MessageOptions
	(*descriptorpb.ServiceOptions)(nil), // 8: google.protobuf.ServiceOptions
	(*descriptorpb.FileOptions)(nil),    // 9: google.protobuf.FileOptions
}
var file_pbmap_proto_depIdxs = []int32{
	6, // 0: pbmap.method_map:extendee -> google.protobuf.MethodOptions
	7, // 1: pbmap.query_map:extendee -> google.protobuf.MessageOptions
	8, // 2: pbmap.service_map:extendee -> google.protobuf.ServiceOptions
	9, // 3: pbmap.dart_package:extendee -> google.protobuf.FileOptions
	0, // 4: pbmap.method_map:type_name -> pbmap.MethodMap
	1, // 5: pbmap.query_map:type_name -> pbmap.QueryMap
	5, // 6: pbmap.service_map:type_name -> pbmap.ServiceMap
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	4, // [4:7] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbmap_proto_init() }
func file_pbmap_proto_init() {
	if File_pbmap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbmap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodMap); i {
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
		file_pbmap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMap); i {
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
		file_pbmap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_pbmap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_pbmap_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sort); i {
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
		file_pbmap_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMap); i {
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
			RawDescriptor: file_pbmap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_pbmap_proto_goTypes,
		DependencyIndexes: file_pbmap_proto_depIdxs,
		MessageInfos:      file_pbmap_proto_msgTypes,
		ExtensionInfos:    file_pbmap_proto_extTypes,
	}.Build()
	File_pbmap_proto = out.File
	file_pbmap_proto_rawDesc = nil
	file_pbmap_proto_goTypes = nil
	file_pbmap_proto_depIdxs = nil
}
