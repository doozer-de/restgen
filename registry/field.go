package registry

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Fields is a slice of Fields of a protobuf message
type Fields []*Field

// Field represents the information about a field in a message
// at the moment just a simple wrapper for easier access and a reference to the Registry
type Field struct {
	Type     *descriptor.FieldDescriptorProto
	Registry *Registry

	Name string
}

// NewField instantiates a new Field
func NewField(r *Registry, f *descriptor.FieldDescriptorProto) *Field {
	return &Field{
		Type: f,
		Name: (*f).GetName(),
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
