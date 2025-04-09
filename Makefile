PROTO_SRC_DIR := pkg/protos
PROTO_OUT_DIR := pkg/pb

.PHONY: protos
protos:
	@rm -rf $(PROTO_OUT_DIR) && mkdir -p $(PROTO_OUT_DIR)
	@find $(PROTO_SRC_DIR) -name "*.proto" -exec protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative {} \;
	@cp -r $(PROTO_SRC_DIR)/* $(PROTO_OUT_DIR)
	@find $(PROTO_SRC_DIR) -name "*.go" -exec rm {} \;
	@find $(PROTO_OUT_DIR) -name "*.proto" -exec rm {} \;