package registry

import (
	"fmt"
	"log"

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

	Values []*EnumValue
}

func (e *Enum) String() string {
	return fmt.Sprintf(".%s.%s", e.Package, e.Name)
}

func NewEnum(d *descriptor.EnumDescriptorProto, f *File, index int) *Enum {
	e := &Enum{
		Type:     d,
		File:     f,
		Registry: f.Registry,
		Package:  f.Package,
		Index:    index,
		Name:     d.GetName(),
	}

	valuesMap := map[int]*descriptor.EnumValueDescriptorProto{}

	for _, v := range d.Value {
		valuesMap[int(*v.Number)] = v
	}

	for i := 0; i < len(valuesMap); i++ {
		if v, ok := valuesMap[i]; ok {
			e.Values = append(e.Values, NewEnumValue(v, e, f.Registry))
		} else {
			log.Fatalf("error on enum %s: Values from 0..%d should be present. Value %d not found ", e.Name, len(valuesMap), i)
		}
	}

	return e
}

type EnumValue struct {
	Type *descriptor.EnumValueDescriptorProto
	Enum *Enum

	Name   string
	Number int32

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

func (es *Enums) Get(name string) (*Enum, bool) {
	for _, e := range *es {
		if e.Name == name || e.String() == name {
			return e, true
		}
	}

	return nil, false
}

func NewEnumValue(d *descriptor.EnumValueDescriptorProto, e *Enum, r *Registry) *EnumValue {
	return &EnumValue{
		Enum:     e,
		Registry: r,
		Type:     d,
		Name:     *d.Name,
		Number:   *d.Number,
	}
}
