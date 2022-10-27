package samplesvc

//go:generate protoc -I. -I.. --go_out=. --go_opt=paths=source_relative pb/samplesvc.proto
//go:generate protoc -I. -I.. --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/samplesvc.proto
//go:generate protoc -I. -I.. --bff_out=. pb/samplesvc.proto
