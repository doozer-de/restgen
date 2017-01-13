package registry

import (
	"fmt"
	"log"
	"strings"

	"bitbucket.org/doozer-de/restgen/pbmap"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Method is wrapper around the MethodDescriptorProto with some additional/extracted information useful to generate the Gateway
type Method struct {
	Type            *descriptor.MethodDescriptorProto
	Package         string
	Name            string
	RESTMethod      string
	RESTPath        string
	RESTPathVars    PathParams
	RESTQueryString QueryStringParams
	RESTBody        string
	comment         string
	methodMapParsed bool
	hasMethodMap    bool
	InputType       *descriptor.DescriptorProto
	OutputType      *descriptor.DescriptorProto
	Registry        *Registry
}

// getQueryMaps returns all the querymaps of the  m or it's submessages
func (m *Method) setQueryMaps() error {
	var foundQueryMap bool
	for _, fieldProto := range m.InputType.GetField() {
		if fieldProto.GetType() != descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			// Only Messages can potentially have QueryMaps
			continue
		}

		msg, ok := m.Registry.Messages[*fieldProto.TypeName]
		if !ok {
			// FIXME(ds) add propper error handling
			panic(fmt.Errorf("Message %v not found", *fieldProto.TypeName))
		}

		if !msg.hasQueryMap() {
			// Message has no QueryMap
			continue
		}
		qm, err := msg.getQueryMap()

		if err != nil {
			return fmt.Errorf("Error getting querymap of %v: %v", msg.String(), err)
		}

		qs, err := NewQueryStringParams(qm.Query)

		if err != nil {
			return fmt.Errorf("error parsing query from %s: %s", qm.String(), err)
		}

		for _, qsp := range *qs {
			fieldName := Sanitize(fieldProto.GetName())
			// In the messages that have a QueryMap we define only a Field in that message. To assign the actual value in the generated code we need the full path to the field.
			md := msg.GetFieldType(qsp.Field)
			ownMetaData := []FieldMetadata{
				FieldMetadata{
					typeString: fieldProto.GetTypeName(),
					ProtoKind:  *fieldProto.Type,
					Name:       *fieldProto.Name,
				},
			}

			ownMetaData = append(ownMetaData, (*md)...)
			fmds := FieldMetadatas(ownMetaData)

			qsp.Metadata = &fmds
			qsp.Field = fmt.Sprintf("%s.%s", fieldName, qsp.Field)
			m.RESTQueryString = append(m.RESTQueryString, qsp)
			foundQueryMap = true
		}
	}

	if foundQueryMap {
		m.Registry.Service.Imports = append(m.Registry.Service.Imports, restmapimport)
	}

	return nil
}

// InputTypeName retrurns the mame if the method's input type
func (m *Method) InputTypeName() string {
	return strings.TrimLeft(*m.Type.InputType, ".")
}

// OutputTypeName returns the name of the method's output type
func (m *Method) OutputTypeName() string {
	return strings.TrimLeft(*m.Type.OutputType, ".")
}

// HasMethodMap returns true, if the method has a MethodMap options
func (m *Method) HasMethodMap() bool {
	return m.hasMethodMap
}

// HasBodyMapping returns true, if in the REST Mapping the requests body is mapped
func (m *Method) HasBodyMapping() bool {
	return m.RESTBody != ""
}

// HasQueryStringMapping returns true, if the Method has a Query String Mapping
func (m *Method) HasQueryStringMapping() bool {
	return len(m.RESTQueryString) != 0
}

// HasPathParamsMapping returns true, if url-path parameters are mapped
func (m *Method) HasPathParamsMapping() bool {
	return len(m.RESTPathVars) != 0
}

// HarmonizedRESTPath returns a cleaned-up version of Mapped REST-Path so it can be registered without conflicts to the router
func (m *Method) HarmonizedRESTPath() string {
	return harmonizePathVars(m.RESTPath)
}

func (m *Method) setMethodMapExtension() {
	opt := m.Type.GetOptions()
	if opt == nil {
		return
	}
	if !proto.HasExtension(opt, pbmap.E_MethodMap) {
		return
	}
	ext, err := proto.GetExtension(opt, pbmap.E_MethodMap)
	if err != nil {
		return
	}
	mm, ok := ext.(*pbmap.MethodMap)

	if !ok {
		return
	}
	m.hasMethodMap = true
	m.RESTMethod = mm.Method
	m.RESTBody = mm.Body
	m.RESTPath = mm.Path

	qs, err := NewQueryStringParams(mm.QueryString)
	if err != nil {
		log.Fatalf("Invid Query string definition: %s", err)
	}
	m.RESTQueryString = *qs

	message := m.Registry.Messages[*m.Type.InputType]
	for i, param := range m.RESTQueryString {
		m.RESTQueryString[i].Metadata = message.GetFieldType(param.Field)
	}

	for i, v := range ParsePath(mm.Path) {
		san := Sanitize(v)

		fmds, err := message.getFieldType(san)
		if err != nil {
			log.Fatalf("Invalid path definition. Field '%s' (sanitized: %s) not found: %s", v, san, err)
		}

		m.RESTPathVars = append(m.RESTPathVars, PathParam{
			N:              i,
			FieldRaw:       v,
			FieldSanitized: san,
			Metadata:       fmds,
		})
	}

	m.methodMapParsed = true
}
