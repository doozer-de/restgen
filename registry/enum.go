package registry

import (
	"fmt"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Enum struct {
	Type *descriptor.EnumDescriptorProto

	File     *File
	Registry *Registry

	Filename string
	Package  string
	Comment  string
	Index    int
	Name     string

	Values EnumValues
}

func (e *Enum) String() string {
	return fmt.Sprintf(".%s.%s", e.Package, e.Type.GetName())
}

func NewEnum(d *descriptor.EnumDescriptorProto, f *File, index int) *Enum {
	e := &Enum{
		Type:     d,
		File:     f,
		Registry: f.Registry,
		Package:  f.Package,
		Index:    index,
	}

	return e
}

type EnumValue struct {
	Type *descriptor.EnumValueDescriptorProto
	Enum *Enum

	Registry *Registry
}

type Enums []*Enum

func (es *Enums) Add(ne *Enum) {
	for _, e := range *es {
		if e.String() == ne.String() {
			return
		}
	}
	*es = append(*es, ne)
}

type EnumValues []*EnumValue

func NewEnumValue(d *descriptor.EnumValueDescriptorProto, e *Enum, r *Registry) *EnumValue {
	return &EnumValue{
		Type:     d,
		Enum:     e,
		Registry: r,
	}
}
