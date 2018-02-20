package pbmap

//go:generate sh -c "protoc -I. -I`go env GOPATH`/src --go_out=. *.proto"

// NOTE fix invalid import pathes in generated file
//go:generate sh -c "sed -i'.bak' -e'/^import google_protobuf/s,.*,import google_protobuf \"github.com/golang/protobuf/protoc-gen-go/descriptor\",' pbmap.pb.go"
