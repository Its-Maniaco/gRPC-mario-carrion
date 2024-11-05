PROTO_FILES := $(shell find proto -name "*.proto")

gen:
	protoc --proto_path=proto $(PROTO_FILES) --go_out=. --go-grpc_out=.

# protoc --proto_path=proto $(find proto -name "*.proto") --go_out=. --go-grpc_out=.

go_server:
	go run ./server/*.go

go_client:
	go run ./client/main.go