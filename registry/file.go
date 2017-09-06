package registry

import (
	"fmt"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Files []*File

func (fs *Files) Add(f *File) {
	for _, c := range *fs {
		if c.Name == f.Name {
			return
		}
	}
	(*fs) = append((*fs), f)
}

func (fs *Files) Get(key string) (*File, bool) {
	for _, c := range *fs {
		if c.Name == key {
			return c, true
		}
	}
	return nil, false
}

type File struct {
	File     *descriptor.FileDescriptorProto
	Name     string
	Messages []*Message
	Package  string
	Registry *Registry
}

func NewFile(f *descriptor.FileDescriptorProto, r *Registry) *File {
	file := &File{
		File:     f,
		Name:     *f.Name,
		Package:  f.GetPackage(),
		Registry: r,
	}

	isExtension := false
	messages := make([]*Message, 0, len(f.MessageType))

	for j, m := range f.MessageType {
		for _, e := range f.Extension {
			if *e.TypeName == fmt.Sprintf(".%s.%s", file.Package, *m.Name) {
				isExtension = true

				break
			}
		}

		if !isExtension {
			messages = append(messages, NewMessage(m, file, j))
		}

		isExtension = false
	}

	file.Messages = messages

	return file
}
