#!/bin/bash
set -e

proto_dir=./schema

protoc --gogoopsee_out=plugins=grpc+graphql,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:${proto_dir} --proto_path=/gopath/src:${proto_dir} ${proto_dir}/*.proto

GOPATH="$GOPATH:/build"
go get -t ./... && \
  go test -v ./...
