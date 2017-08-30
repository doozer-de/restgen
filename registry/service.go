package registry

import (
	"fmt"
	"log"

	"github.com/doozer-de/restgen/pbmap"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Service wraps a ServiceDescriptorProto with additions for code Generation
type Service struct {
	Type     *descriptor.ServiceDescriptorProto
	Registry *Registry
	File     *descriptor.FileDescriptorProto

	Package       string
	GoPackage     string
	Name          string
	Imports       []string
	BaseURI       string
	Methods       map[string]*Method
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

	if inputType, ok := s.Registry.Messages[*method.InputType]; ok {
		m.InputType = inputType.Type
	}

	if outputType, ok := s.Registry.Messages[*method.OutputType]; ok {
		m.OutputType = outputType.Type
	}

	m.setMethodMapExtension()
	err := m.setQueryMaps()

	if err != nil {
		log.Fatalf("Could not read the query map options: %s", err)
	}

	s.Methods[m.Name] = m
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

func (s *Service) getServiceMapExtension() (*pbmap.ServiceMap, error) {
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
