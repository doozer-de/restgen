package registry

import (
	"fmt"
	"log"

	"github.com/doozer-de/restgen/pbmap"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Methods []*Method

func (ms *Methods) Add(m *Method) {
	for _, c := range *ms {
		if c.Name == m.Name {
			return
		}
	}
	(*ms) = append((*ms), m)

}

func (ms *Methods) Get(key string) (*Method, bool) {
	for _, c := range *ms {
		if c.Name == key {
			return c, true
		}
	}
	return nil, false
}

// Service wraps a ServiceDescriptorProto with additions for code Generation
type Service struct {
	Type     *descriptor.ServiceDescriptorProto
	Registry *Registry
	File     *File

	Package       string
	GoPackage     string
	Name          string
	Imports       []string
	BaseURI       string
	Methods       Methods
	Comment       string
	Index         int
	TargetPackage string
	Version       string
}

func (s *Service) RegisterMethod(method *descriptor.MethodDescriptorProto) {
	m := &Method{
		Type:     method,
		Name:     *method.Name,
		Package:  s.Package,
		Registry: s.Registry,
	}

	if inputType, ok := s.Registry.Messages.Get(*method.InputType); ok {
		m.InputType = inputType.Type
	}

	if outputType, ok := s.Registry.Messages.Get(*method.OutputType); ok {
		m.OutputType = outputType.Type
	}

	m.setMethodMapExtension()
	err := m.setQueryMaps()

	if err != nil {
		log.Fatalf("Could not read the query map options: %s", err)
	}

	s.Methods.Add(m)
}

// GetMappedMethods returns all the Methods of the currenct service that have a MethodMap extension
// Only those with a MethodMap extension will be exposed on the REST interface
func (s *Service) GetMappedMethods() []*Method {
	methods := make([]*Method, 0, len(s.Methods))

	for _, m := range s.Methods {
		if m.HasMethodMap() {
			methods = append(methods, m)
		}
	}

	return methods
}

// ServiceType returns the name of the GRPC service, which will be used by the REST endpoints.
func (s *Service) ServiceType() string {
	return fmt.Sprintf("%s.%sServer", s.Package, s.Name)
}

func (s *Service) HasServiceMapExtension() bool {
	opt := s.Type.GetOptions()
	if opt == nil {
		return false
	}
	if !proto.HasExtension(opt, pbmap.E_ServiceMap) {
		return false
	}

	return true
}

func (s *Service) getServiceMapExtension() (*pbmap.ServiceMap, error) {
	if !s.HasServiceMapExtension() {
		return nil, fmt.Errorf("%s does not have a ServiceMap Extension", s.Name)
	}
	opt := s.Type.GetOptions()
	if opt == nil {
		return nil, fmt.Errorf("No Options on Service")
	}
	if !proto.HasExtension(opt, pbmap.E_ServiceMap) {
		return nil, fmt.Errorf("ServiceMap Extension not present")
	}
	ext, err := proto.GetExtension(opt, pbmap.E_ServiceMap)
	if err != nil {
		return nil, fmt.Errorf("Error getting ServiceMap extension")
	}
	sm, ok := ext.(*pbmap.ServiceMap)
	if !ok {
		return nil, fmt.Errorf("Error getting ServiceMap extension, wrong type")
	}
	return sm, nil
}
