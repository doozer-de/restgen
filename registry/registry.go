package registry

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

const restmapimport = "github.com/doozer-de/restgen/pbmap"

// Messages is a map from the messages name ot the Message
type Messages map[string]*Message

// The Registry is the root for registering the Types found in the protobuf structures from the compiler
type Registry struct {
	Files    map[string]*File
	FilesI   map[int]*File // To build a path hierarchy to get the source locations
	Service  *Service
	Messages Messages
	Package  string
}

func (r *Registry) registerMessageProto(pkg string, d *descriptor.DescriptorProto) {
	key := fmt.Sprintf(".%s.%s", pkg, d.GetName())
	var fields Fields

	m := &Message{
		Package:  pkg,
		Type:     d,
		Fields:   fields,
		Registry: r,
	}

	for _, field := range d.GetField() {
		fields = append(fields, NewField(field, m, r))
	}

	m.Fields = fields

	r.Messages[key] = m
}

func (r *Registry) registerServiceProto(pkg string, gopkg string, d *descriptor.ServiceDescriptorProto, f *descriptor.FileDescriptorProto) {
	s := &Service{
		Package:   pkg,
		GoPackage: gopkg,
		Name:      d.GetName(),
		Type:      d,
		Imports:   []string{},
		Methods:   make(map[string]*Method),
		Registry:  r,
		File:      NewFile(f, r),
	}

	r.Service = s

	ext, err := r.Service.getServiceMapExtension()

	if err != nil {
		return
	}

	r.Service.BaseURI = ext.BaseUri
	r.Service.TargetPackage = ext.TargetPackage
	r.Service.Version = ext.Version

	for _, method := range d.GetMethod() {
		r.Service.RegisterMethod(method)
	}
}

// New createsa  new Registry that will read all the information neede from the given CodeGeneratorRequest
func New(r *plugin.CodeGeneratorRequest) *Registry {
	reg := &Registry{
		Files:    make(map[string]*File),
		Messages: make(map[string]*Message),
	}

	files := r.GetProtoFile()

	// Register all Messages
	for _, f := range files {
		reg.Files[f.GetName()] = &File{File: f, Name: f.GetName()}
		pkg := f.GetPackage()

		for _, m := range f.GetMessageType() {
			reg.registerMessageProto(pkg, m)
		}
	}

	// The last file is the service file we want to generate code for (the imports come first)
	serviceFile := files[len(files)-1]
	svcs := serviceFile.GetService()

	if len(svcs) > 0 {
		svc := svcs[0]
		if reg.Service != nil {
			log.Fatal("Only one service in file supported")
		}

		gopkg := serviceFile.GetOptions().GetGoPackage()

		if gopkg == "" {
			log.Fatal("In the service-file the option go_package is needed with the full path of the package to generate correct pathes.")
		}

		reg.registerServiceProto(serviceFile.GetPackage(), gopkg, svc, serviceFile)
	}

	return reg
}

// QueryStringParam is provides all the information nessesary for the code generation to map on query string parameter to a go field
type QueryStringParam struct {
	Converter string
	Field     string // can also be more deep:
}
