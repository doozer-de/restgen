package registry

import (
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Fields is a slice of Fields of a protobuf message
type Fields []*Field

// Field represents the information about a field in a message
// at the moment just a simple wrapper for easier access and a reference to the Registry
type Field struct {
	Type *descriptor.FieldDescriptorProto
	// Is used for maps. A map is a nested type in protobuf. This nested type is only accessible via the descriptor.DescriptorProto
	Message   *Message
	Registry  *Registry
	Name      string
	IsComplex bool
}

// NewField instantiates a new Field
func NewField(f *descriptor.FieldDescriptorProto, m *Message, r *Registry) *Field {
	return &Field{
		Type:      f,
		Message:   m,
		Registry:  r,
		Name:      (*f).GetName(),
		IsComplex: false,
	}
}

// IsRepeated returns true, if the protobuf field is repated
func (f *Field) IsRepeated() bool {
	return (*f.Type).GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED
}

// IsBytes returns true if the protobuf field is of types bytes
func (f *Field) IsBytes() bool {
	return (*f.Type).GetType() == descriptor.FieldDescriptorProto_TYPE_BYTES
}

// IsMessage returns true if the protobuf field is a protobuf message
func (f *Field) IsMessage() bool {
	return (*f.Type).GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

func (f *Field) IsIntermediateMap() bool {
	return strings.HasSuffix((*f.Type).GetTypeName(), "Entry")
}
