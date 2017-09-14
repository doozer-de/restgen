package registry

import (
	"github.com/doozer-de/restgen/pbmap"
	"github.com/golang/protobuf/proto"
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
	Type     *descriptor.FileDescriptorProto
	Name     string
	Messages []*Message
	Package  string
	Registry *Registry
	// Options holds additional options. For example our dart_package options
	Options map[string]string
}

func NewFile(f *descriptor.FileDescriptorProto, r *Registry) *File {
	file := &File{
		Type:     f,
		Name:     *f.Name,
		Package:  f.GetPackage(),
		Registry: r,
		Options:  getAdditionalOptions(f),
	}

	var messages []*Message

	for j, m := range f.MessageType {
		messages = append(messages, NewMessage(m, file, j))
	}

	file.Messages = messages

	return file
}

func getAdditionalOptions(f *descriptor.FileDescriptorProto) map[string]string {
	m := make(map[string]string)

	if dp, ok := dartPackageOption(f); ok {
		m["DartPackage"] = dp
	}

	return m
}

func dartPackageOption(f *descriptor.FileDescriptorProto) (string, bool) {
	opt := f.GetOptions()
	if opt == nil {
		return "", false
	}

	if !proto.HasExtension(opt, pbmap.E_DartPackage) {
		return "", false
	}
	ext, err := proto.GetExtension(opt, pbmap.E_DartPackage)
	if err != nil {
		return "", false
	}
	sm, ok := ext.(*string)
	if !ok {
		return "", false
	}

	return *sm, true
}
