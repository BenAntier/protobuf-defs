.PHONY: test all proto get-deps lint

all: proto rpc.swagger.json rpc.swagger.pb.go test

get-deps:
	go get -u github.com/gengo/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/gengo/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/ckaznocha/protoc-gen-lint
	go get -u github.com/jteeuwen/go-bindata/go-bindata
	go get -u ./...

test: proto
	go test -v ./...

# TODO: The googleapis paths below should be switched to upstream once
# https://github.com/googleapis/googleapis/pull/40 is merged.
proto: lint
	protoc -I/usr/local/include \
		-I. -I.. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
		--go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
		*.proto
	protoc -I/usr/local/include \
		-I. -I.. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		*.proto

rpc.swagger.json: proto
	protoc -I/usr/local/include \
		-I. -I.. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		rpc.proto

rpc.swagger.pb.go: rpc.swagger.json
	go-bindata -o rpc.swagger.pb.go -pkg logger rpc.swagger.json

swagger: rpc.swagger.pb.go

lint:
	protoc -I/usr/local/include \
		-I. -I.. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
		--lint_out=. \
		*.proto
