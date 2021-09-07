module github.com/v8platform/protos

go 1.17

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/v8platform/encoder v0.0.1
	github.com/v8platform/protoc-gen-go-ras v0.0.0-20210902165457-013367855358
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/v8platform/encoder v0.0.1 => ../encoder

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	golang.org/x/net v0.0.0-20210610132358-84b48f89b13b // indirect
	golang.org/x/sys v0.0.0-20210611083646-a4fc73990273 // indirect
	golang.org/x/text v0.3.6 // indirect
)
