// Package emitters generates go-specific code.
package emitters

import (
	"github.com/doozer-de/restgen/registry"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// GetConverterName returns the name of the converter used to
// convert the values given in the parameter and/or query string.
func GetConverterName(f *registry.FieldMetadatas) string {
	return protoconverters[(*f)[len(*f)-1].ProtoKind]
}

var protoconverters = map[descriptor.FieldDescriptorProto_Type]string{
	descriptor.FieldDescriptorProto_TYPE_DOUBLE:   "ToFloat64",
	descriptor.FieldDescriptorProto_TYPE_FLOAT:    "ToFloat32",
	descriptor.FieldDescriptorProto_TYPE_INT64:    "ToInt64",
	descriptor.FieldDescriptorProto_TYPE_UINT64:   "ToUint64",
	descriptor.FieldDescriptorProto_TYPE_INT32:    "ToInt32",
	descriptor.FieldDescriptorProto_TYPE_FIXED64:  "ToUint64",
	descriptor.FieldDescriptorProto_TYPE_FIXED32:  "ToUint32",
	descriptor.FieldDescriptorProto_TYPE_BOOL:     "ToBool",
	descriptor.FieldDescriptorProto_TYPE_STRING:   "ToString",
	descriptor.FieldDescriptorProto_TYPE_BYTES:    "ToBytes",
	descriptor.FieldDescriptorProto_TYPE_UINT32:   "ToUint32",
	descriptor.FieldDescriptorProto_TYPE_SINT32:   "ToInt32",
	descriptor.FieldDescriptorProto_TYPE_SINT64:   "ToInt64",
	descriptor.FieldDescriptorProto_TYPE_SFIXED32: "ToInt32",
	descriptor.FieldDescriptorProto_TYPE_SFIXED64: "ToInt64",
}
