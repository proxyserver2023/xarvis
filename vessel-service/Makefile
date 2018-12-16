build:
	protoc -I. --go_out=plugins=grpc:$(HOME)/github/gapi \
		pkg/models/proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o dist/api cmd/api/main.go 
	chmod +x dist/api
	docker build -t gapi -f ./build/docker/Dockerfile .

run:
	docker run -p 50051:50051 gapi