package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
	"strings"
)

func main() {
	protogen.Options{
		ParamFunc: func(name, value string) error {
			if name != "custom" {
				return fmt.Errorf("unknown option %s", name)
			}

			for _, value := range strings.Split(value, ",") {
				value = strings.TrimSpace(value)

				kv := strings.Split(value, "=")
				if len(kv) != 2 {
					return fmt.Errorf("invalid (-1)")
				}
				name := strings.TrimSpace(kv[0])
				kv = strings.Split(kv[1], ":")
				if len(kv) != 2 {
					return fmt.Errorf("invalid (-2)")
				}
				pointer := false
				if strings.HasPrefix(kv[0], "*") {
					pointer = true
					kv[0] = kv[0][1:]
				}
				as := FullToIdent(kv[0])
				var convertFuncName protogen.GoIdent
				if strings.HasPrefix(kv[1], ".") {
					convertFuncName = protogen.GoIdent{
						GoName:       kv[1][1:],
						GoImportPath: as.GoImportPath,
					}
				} else {
					convertFuncName = FullToIdent(kv[1])
				}

				CustomMap[name] = &CustomOverrideProvider{name: as, pointer: pointer, convertFunc: convertFuncName}
			}

			return nil
		},
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				GenerateFile(gen, f)
			}
		}
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return nil
	})
}
