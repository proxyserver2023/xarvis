# gapi [WIP]

Go rest api bootstrap template with SOLID and Clean Architecture using Microservices Best Practices.

Microservices
1. golang
2. mongodb
3. grpc
4. docker
5. Google Cloud
6. Kubernetes
7. NATS
8. CircleCI
9. Terraform
10. go-micro

```bash
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Features We have covered here...

1. Service Discovery

## Mongodb with go api

 Production Quality Mongodb in docker
 1. data should be spearated from the docker container
 2. custom config
 3. admin tasks

```bash
docker run --name MONGO --restart=always -d -p 27017:27017 mongo mongod --auth
docker exec -it MONGO bash
```
