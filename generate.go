//go:generate rm -f ./proto/*.pb.go
//go:generate protoc --proto_path=./proto --go_opt=paths=source_relative --go_out=./proto ./proto/config.proto
package main
