// Package emitters generates go-specific code.
package emitters

import (
	"github.com/doozer-de/restgen/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

// GetConverterName returns the name of the converter used to
// convert the values given in the parameter and/or query string.
func GetConverterName(f *registry.FieldMetadatas) string {
	return protoconverters[(*f)[len(*f)-1].ProtoKind]
}

var protoconverters = map[descriptorpb.FieldDescriptorProto_Type]string{
	descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:   "ToFloat64",
	descriptorpb.FieldDescriptorProto_TYPE_FLOAT:    "ToFloat32",
	descriptorpb.FieldDescriptorProto_TYPE_INT64:    "ToInt64",
	descriptorpb.FieldDescriptorProto_TYPE_UINT64:   "ToUint64",
	descriptorpb.FieldDescriptorProto_TYPE_INT32:    "ToInt32",
	descriptorpb.FieldDescriptorProto_TYPE_FIXED64:  "ToUint64",
	descriptorpb.FieldDescriptorProto_TYPE_FIXED32:  "ToUint32",
	descriptorpb.FieldDescriptorProto_TYPE_BOOL:     "ToBool",
	descriptorpb.FieldDescriptorProto_TYPE_STRING:   "ToString",
	descriptorpb.FieldDescriptorProto_TYPE_BYTES:    "ToBytes",
	descriptorpb.FieldDescriptorProto_TYPE_UINT32:   "ToUint32",
	descriptorpb.FieldDescriptorProto_TYPE_SINT32:   "ToInt32",
	descriptorpb.FieldDescriptorProto_TYPE_SINT64:   "ToInt64",
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: "ToInt32",
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: "ToInt64",
}
