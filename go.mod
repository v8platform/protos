module github.com/v8platform/protos

go 1.17

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/spf13/cast v1.4.1
	github.com/v8platform/encoder v0.0.1
	github.com/v8platform/protoc-gen-go-ras v0.0.0-20210902121924-0bba4413ddc3
	go.buf.build/v8platform/go-gen-ras/v8platform/rasapis v1.2.1
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

//replace github.com/v8platform/encoder v0.0.1 => ../encoder

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	go.buf.build/v8platform/go-gen-ras/v8platform/encodingapis v1.2.2 // indirect
	go.buf.build/v8platform/go-gen-ras/v8platform/serializeapis v1.2.1 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.6 // indirect
)
