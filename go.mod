module github.com/v8platform/protos

go 1.17

//replace github.com/v8platform/encoder v0.0.3 => ../../khorevaa/encoder

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/spf13/cast v1.4.1
	github.com/v8platform/encoder v0.0.4
	github.com/v8platform/protoc-gen-go-ras v0.2.1
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210903162649-d08c68adba83 // indirect
)
