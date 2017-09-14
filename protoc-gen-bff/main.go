// Package protoc-gen-bff implements a plugin to generate go code for REST endpoints from protocol buffer files.
//
// Usage: TODO
package main

//go:generate sh -c "go-bindata *.tmpl"

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/doozer-de/restgen/protoc-gen-bff/emitters"
	"github.com/doozer-de/restgen/registry"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	req := new(plugin.CodeGeneratorRequest)

	if err = proto.Unmarshal(input, req); err != nil {
		panic(err)
	}

	out, err := generate(req)

	if err != nil {
		s := err.Error()
		emit(&plugin.CodeGeneratorResponse{Error: &s})
		return
	}

	emit(&plugin.CodeGeneratorResponse{File: out})
}

func ignore() []*plugin.CodeGeneratorResponse_File {
	s := "// +build ignore \n\n// Nothing to see here, just a dummy"
	name := "noservice_gen.go"
	f := &plugin.CodeGeneratorResponse_File{}
	f.Content = &s
	f.Name = &name
	return []*plugin.CodeGeneratorResponse_File{f}
}

func generate(r *plugin.CodeGeneratorRequest) ([]*plugin.CodeGeneratorResponse_File, error) {
	reg := registry.New(r)

	f := plugin.CodeGeneratorResponse_File{}

	if reg.Service == nil || !reg.Service.HasServiceMapExtension() {
		return ignore(), nil
	}

	out := &bytes.Buffer{}

	funcMap := template.FuncMap{
		"GetConverterName": emitters.GetConverterName,
	}

	name := fmt.Sprintf("%v_gen.go", reg.Service.TargetPackage)
	f.Name = &name
	t := template.Must(template.New("handlers").Funcs(funcMap).Parse(string(MustAsset("bff.tmpl"))))
	s := reg.Service

	err := t.Execute(out, s)

	if err != nil {
		log.Fatal(err)
	}

	outformatted, err := format.Source(out.Bytes())

	if err != nil {
		log.Print("Error formatting GoCode, using unformated one")
		outformatted = out.Bytes()
	}

	outString := string(outformatted)
	f.Content = &outString
	return []*plugin.CodeGeneratorResponse_File{&f}, nil
}

func emit(r *plugin.CodeGeneratorResponse) {
	buf, err := proto.Marshal(r)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stdout.Write(buf); err != nil {
		log.Fatal(err)
	}
}
