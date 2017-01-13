package samplesvc

//go:generate sh -c "protoc -Ipb -I`go env GOPATH`/src --go_out=plugins=grpc:`go env GOPATH` pb/*.proto"
//go:generate sh -c "protoc -Ipb -I`go env GOPATH`/src --bff_out=. pb/*.proto"
