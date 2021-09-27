module github.com/v8platform/protos

go 1.17

//replace github.com/v8platform/encoder v0.0.3 => ../../khorevaa/encoder

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/v8platform/encoder v0.0.4
	github.com/v8platform/protoc-gen-go-ras v0.2.1
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/kr/pretty v0.3.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
)
