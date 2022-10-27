// Package protoc-gen-bff implements a plugin to generate go code for REST endpoints from protocol buffer files.
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/doozer-de/restgen/protoc-gen-bff/emitters"
	"github.com/doozer-de/restgen/registry"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

//go:embed bff.tmpl
var bffTmpl string

func main() {
	input, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	req := new(pluginpb.CodeGeneratorRequest)

	if err = proto.Unmarshal(input, req); err != nil {
		panic(err)
	}

	out, err := generate(req)

	if err != nil {
		s := err.Error()
		emit(&pluginpb.CodeGeneratorResponse{Error: &s})
		return
	}

	emit(&pluginpb.CodeGeneratorResponse{File: out})
}

func ignore() []*pluginpb.CodeGeneratorResponse_File {
	s := "// +build ignore \n\n// Nothing to see here, just a dummy"
	name := "noservice_gen.go"
	f := &pluginpb.CodeGeneratorResponse_File{}
	f.Content = &s
	f.Name = &name
	return []*pluginpb.CodeGeneratorResponse_File{f}
}

func generate(r *pluginpb.CodeGeneratorRequest) ([]*pluginpb.CodeGeneratorResponse_File, error) {
	reg := registry.New(r)

	f := pluginpb.CodeGeneratorResponse_File{}

	if reg.Service == nil || !reg.Service.HasServiceMapExtension() {
		return ignore(), nil
	}

	out := &bytes.Buffer{}

	funcMap := template.FuncMap{
		"GetConverterName": emitters.GetConverterName,
	}

	name := fmt.Sprintf("%v_gen.go", reg.Service.TargetPackage)
	f.Name = &name
	t := template.Must(template.New("handlers").Funcs(funcMap).Parse(bffTmpl))
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
	return []*pluginpb.CodeGeneratorResponse_File{&f}, nil
}

func emit(r *pluginpb.CodeGeneratorResponse) {
	buf, err := proto.Marshal(r)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stdout.Write(buf); err != nil {
		log.Fatal(err)
	}
}
