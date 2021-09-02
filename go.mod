module github.com/v8platform/protos

go 1.17

require go.buf.build/v8platform/go-gen-ras/v8platform/rasapis v1.2.1

replace github.com/v8platform/encoder v0.0.1 => ../encoder

require (
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/v8platform/encoder v0.0.1 // indirect
	go.buf.build/v8platform/go-gen-ras/v8platform/encodingapis v1.2.2 // indirect
	go.buf.build/v8platform/go-gen-ras/v8platform/serializeapis v1.2.1 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
