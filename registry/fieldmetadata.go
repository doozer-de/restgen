package registry

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// FieldMetadatas are not part of the Registry. They represent a path to a field in the object tree.
// like r.User.Name.Firstname in a Protobuf Request.
type FieldMetadatas []FieldMetadata

// FieldMetadata holds Information needed to generate the code to copy a defined QueryString parameters to a actually field
type FieldMetadata struct {
	Type      string
	ProtoKind descriptor.FieldDescriptorProto_Type
	Name      string
}

// GetType gets the type of the final field.
// Protobuf has a leading . which will be removed for being a proper go identifier
func (f *FieldMetadata) GetType() string {
	return strings.TrimLeft(f.Type, ".")
}

// Generate returns the go code to set a field value in an struct tree
// Since protobuf works with pointer intermediate structs have to be instantiated
func (fmds *FieldMetadatas) Generate() string {
	s := "req"
	var builder func(fmds []FieldMetadata) string
	builder = func(fmds []FieldMetadata) string {
		if len(fmds) == 0 {
			return ""
		}
		fmd := fmds[0]

		s = fmt.Sprintf("%s.%s", s, fmd.Name)
		pre := fmt.Sprintf("if %s == nil {\n    %s = &%s{}\n", s, s, fmd.GetType())

		var mid string
		if len(fmds) > 1 {
			mid = builder(fmds[1:])
		}

		post := "}"

		return fmt.Sprintf("%s%s%s", pre, mid, post)
	}

	return builder((*fmds)[:len(*fmds)-1])
}

// GetPath returns the path to write out to set a value
// For example r.User.Name.Firstname
func (fmds *FieldMetadatas) GetPath() string {
	var path string

	for _, fmd := range *fmds {
		// Write the first letter in upper case, because the go strcuts and the items in
		// them created by the protobuf compiler are written with the first letter in
		// upper case
		// TODO(0x434d53): More Testcases
		name := fmt.Sprintf("%s%s", strings.ToUpper(fmd.Name[0:1]), fmd.Name[1:])

		path = fmt.Sprintf("%s.%s", path, name)
	}

	return path
}
