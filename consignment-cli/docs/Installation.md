# Microservices with Golang

## Install gRPC and Protocol Buffer

```bash
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

build docker image

```bash
docker build -t gapi -f ./build/docker/Dockerfile .
```