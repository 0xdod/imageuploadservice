PROTO_OUT_DIR = internal/grpc

hello:
	echo "hello world"

gen: 
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(PROTO_OUT_DIR)/image_uploader.proto

run-server: 
	./run-server

test:
	go test -coverprofile=cover.out -short ./...

cover:
	go tool cover -html=cover.out

build:
	docker build -t image-uploader .