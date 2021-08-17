
docker run --volume "$(PWD):/workspace" -v "C:\Users\khorevaa\go\bin\bin\protoc-gen-go.exe:/bin/protoc-gen-go" --workdir /workspace bufbuild/buf generate


