package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	protogen.Options{}.
		Run(func(gen *protogen.Plugin) error {
			for _, f := range gen.Files {
				if f.Generate && shouldProcess(f) {
					GenerateFile(gen, f)
				}
			}
			gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
			return nil
		})
}

func shouldProcess(file *protogen.File) bool {
	ignoredFiles := []string{"graphql/graphql.proto", "graphql.proto", "google/protobuf/descriptor.proto", "google/protobuf/wrappers.proto", "google/protobuf/timestamp.proto", "github.com/kitt-technology/protoc-gen-graphql/graphql/graphql.proto"}
	for _, ignored := range ignoredFiles {
		if *file.Proto.Name == ignored {
			return false
		}
	}
	// if proto.HasExtension(file.Proto.Options, graphql.E_Disabled) {
	// 	return !proto.GetExtension(file.Proto.Options, graphql.E_Disabled).(bool)
	// }
	return true
}
