.PHONY: test all proto get-deps

all: test

get-deps:
	go get -t -u ./...

test: proto
	go test -v ./...

proto: service.pb.go

service.pb.go:
	protoc -I.. -I. service.proto --go_out=plugins=grpc:.
