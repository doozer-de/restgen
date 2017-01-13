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

	"bitbucket.org/doozer-de/restgen/protoc-gen-bff/emitters"
	"bitbucket.org/doozer-de/restgen/registry"
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
		log.Fatalf("Could not generate: %s", err)
	}

	emit(&plugin.CodeGeneratorResponse{File: out})
}

func generate(r *plugin.CodeGeneratorRequest) ([]*plugin.CodeGeneratorResponse_File, error) {
	reg := registry.New(r)

	out := &bytes.Buffer{}

	f := plugin.CodeGeneratorResponse_File{}

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
