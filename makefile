all: proto

proto:
	protoc --go_out=paths=source_relative:pkg/service --go-grpc_out=paths=source_relative:pkg/service proto/item.proto


	