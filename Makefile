PROTO_SRC_DIR := pkg/protos
PROTO_OUT_DIR := pkg/pb

.PHONY: protos
protos:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/protos/user/user.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/protos/dbmanager/dbmanager.proto