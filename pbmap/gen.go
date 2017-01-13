package pbmap

//go:generate sh -c "protoc -I. -I`go env GOPATH`/src --go_out=. *.proto"

// NOTE fix invalid import pathes in generated file
//go:generate sh -c "sed -ie '/^import google_protobuf/cimport google_protobuf \"github.com/golang/protobuf/protoc-gen-go/descriptor\"' map.pb.go"
