//go:generate rm -f ./proto/*.pb.go
//go:generate protoc --proto_path=./proto --go_opt=paths=source_relative --go_out=./proto ./proto/config.proto
//go:generate go build -o /tmp/kumadoc .
//go:generate protoc -I=./tmp -I=. --plugin=protoc-gen-kumadoc=/tmp/kumadoc --kumadoc_out=./tmp/generated ./tmp/foo.proto
package main
