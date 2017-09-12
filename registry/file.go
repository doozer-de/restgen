package registry

import (
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
	Options  map[string]string
}

func NewFile(f *descriptor.FileDescriptorProto, r *Registry) *File {
	file := &File{
		File:     f,
		Name:     *f.Name,
		Package:  f.GetPackage(),
		Registry: r,
	}

	var messages []*Message

	for j, m := range f.MessageType {
		messages = append(messages, NewMessage(m, file, j))
	}

	file.Messages = messages

	return file
}
