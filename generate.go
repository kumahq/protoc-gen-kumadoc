//go:generate rm -f ./proto/generated/*
//go:generate protoc --proto_path=./proto --go_opt=paths=source_relative --go_out=./proto/generated ./proto/config.proto

package main
