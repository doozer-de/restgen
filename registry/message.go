package registry

import (
	"fmt"
	"log"
	"strings"

	"github.com/doozer-de/restgen/pbmap"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Message wraps a protobuf message with a little extra information for code generation
type Message struct {
	Type *descriptor.DescriptorProto

	File     *File
	Registry *Registry

	Filename string
	Package  string
	Comment  string
	Index    int
	Fields   Fields
}

func NewMessage(d *descriptor.DescriptorProto, f *File, index int) *Message {
	m := &Message{
		Type:     d,
		File:     f,
		Registry: f.Registry,
		Package:  f.Package,
		Index:    index,
	}

	fields := make(Fields, 0, len(d.GetField()))

	for _, field := range d.GetField() {
		fields = append(fields, NewField(field, m, f.Registry))
	}

	m.Fields = fields

	return m
}

// hasQueryMap returns true if m has a QueryMap defined
func (m *Message) hasQueryMap() bool {
	opt := m.Type.GetOptions()

	if opt == nil {
		return false
	}

	if proto.HasExtension(opt, pbmap.E_QueryMap) {
		return true
	}

	return false
}

func (m *Message) getQueryMap() (*pbmap.QueryMap, error) {
	if !m.hasQueryMap() {
		return nil, fmt.Errorf("message does not have a querymap")
	}
	opt := m.Type.GetOptions()
	ext, err := proto.GetExtension(opt, pbmap.E_QueryMap)

	if err != nil {
		return nil, fmt.Errorf("error getting QueryMap extension")
	}

	qm, ok := ext.(*pbmap.QueryMap)
	// Should never happen
	if !ok {
		return nil, fmt.Errorf("error getting QueryMap, wrong type")
	}

	return qm, nil
}

func (m *Message) getFieldType(path string) (*FieldMetadatas, error) {
	pathfields := strings.Split(path, ".")

	for _, definedField := range m.Type.Field {
		if strings.Compare(*definedField.Name, pathfields[0]) == 0 {
			if len(pathfields) == 1 {
				if *definedField.Type != descriptor.FieldDescriptorProto_TYPE_MESSAGE {
					res := []FieldMetadata{
						{
							// TypeName not needed, because we progress further and just have to collect information we need for the initialization
							ProtoKind: *definedField.Type,
							Name:      *definedField.Name,
						},
					}
					f := FieldMetadatas(res)

					return &f, nil
				}
			} else if *definedField.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
				if definedField.TypeName == nil {
					panic(*definedField)
				}
				if mNew, ok := m.Registry.Messages.Get(*definedField.TypeName); ok {
					mds, err := mNew.getFieldType(strings.Join(pathfields[1:], "."))
					if err != nil {
						return nil, fmt.Errorf("%s.%s", pathfields[0], err)
					}
					mt := []FieldMetadata{
						{
							//Success: We found the type value to set for the path in the object tree
							Type:      *definedField.TypeName,
							ProtoKind: *definedField.Type,
							Name:      *definedField.Name,
						},
					}

					mt = append(mt, *mds...)
					f := FieldMetadatas(mt)
					return &f, nil
				}
			}
		}
	}

	// Fall though for all other cases: Wrong type
	return nil, fmt.Errorf("'%s' not found in '%s' or wrong type", path, m.String())
}

func (m *Message) String() string {
	return fmt.Sprintf(".%s.%s", m.Package, m.Type.GetName())
}

// GetFieldType returns the type of a field in the struct tree below the message
func (m *Message) GetFieldType(path string) *FieldMetadatas {
	res, err := m.getFieldType(path)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

// Name returns the Name of the message
func (m *Message) Name() string {
	return m.Type.GetName()
}
