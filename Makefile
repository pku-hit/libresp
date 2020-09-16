go:
	protoc --proto_path=$(GOPATH)/src:. --gofast_out=plugin=grpc:$(GOPATH)/src *.proto

all: go