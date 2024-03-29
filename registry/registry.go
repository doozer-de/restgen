package registry

import (
	"log"

	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const restmapimport = "github.com/doozer-de/restgen/pbmap"

// Messages is a map from the messages name ot the Message
type Messages []*Message

func (ms *Messages) Add(m *Message) {
	for _, c := range *ms {
		if c.String() == m.String() {
			return
		}
	}
	*ms = append(*ms, m)
}

func (ms *Messages) Get(key string) (*Message, bool) {
	for _, c := range *ms {
		if c.String() == key {
			return c, true
		}
	}
	return nil, false
}

// The Registry is the root for registering the Types found in the protobuf structures from the compiler
type Registry struct {
	Files    Files
	Service  *Service
	RootFile *File
	Messages Messages
	Enums    Enums
	Package  string
}

func (r *Registry) registerMessageProto(pkg string, d *descriptorpb.DescriptorProto) {
	var fields Fields

	m := &Message{
		Package:  pkg,
		Type:     d,
		Registry: r,
	}

	for _, field := range d.GetField() {
		fields = append(fields, NewField(field, m, r))
	}

	m.Fields = fields

	r.Messages.Add(m)
}

func (r *Registry) registerEnumProto(d *descriptorpb.EnumDescriptorProto, f *File, index int) {
	e := NewEnum(d, f, index)

	r.Enums.Add(e)
}

func (r *Registry) registerServiceProto(file *File) {
	svcs := file.Type.GetService()

	if len(svcs) == 0 {
		return
	}
	svc := svcs[0]
	if r.Service != nil {
		log.Fatal("Only one service in file supported")
	}

	gopkg := *file.Type.GetOptions().GoPackage

	s := &Service{
		Package:   file.Package,
		GoPackage: gopkg,
		Name:      svc.GetName(),
		Type:      svc,
		Imports:   []string{},
		Methods:   []*Method{},
		Registry:  r,
		File:      file,
	}

	r.Service = s
	r.RootFile = s.File

	ext, err := r.Service.getServiceMapExtension()

	if err != nil {
		return
	}

	r.Service.BaseURI = ext.BaseUri
	r.Service.TargetPackage = ext.TargetPackage
	r.Service.Version = ext.Version

	for _, method := range svc.GetMethod() {
		r.Service.RegisterMethod(method)
	}
}

// New createsa  new Registry that will read all the information neede from the given CodeGeneratorRequest
func New(r *pluginpb.CodeGeneratorRequest) *Registry {
	reg := &Registry{
		Files:    []*File{},
		Messages: []*Message{},
	}

	files := r.GetProtoFile()

	// Register all Messages, Enums
	for _, f := range files {
		file := NewFile(f, reg)
		reg.Files.Add(file)

		pkg := f.GetPackage()

		for _, m := range f.GetMessageType() {
			reg.registerMessageProto(pkg, m)
		}

		for _, e := range f.GetEnumType() {
			reg.registerEnumProto(e, file, 0) // TOOO(cs) find correct index
		}

	}

	// The last file is the service file we want to generate code for (the imports come first)
	rootFile := reg.Files[len(reg.Files)-1]
	reg.RootFile = rootFile
	reg.registerServiceProto(reg.RootFile)

	return reg
}

// QueryStringParam is provides all the information nessesary for the code generation to map on query string parameter to a go field
type QueryStringParam struct {
	Converter string
	Field     string // can also be more deep:
}
