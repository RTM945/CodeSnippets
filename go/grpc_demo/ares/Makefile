PROTOC      := protoc/bin/protoc.exe
PROTOC_PATH := protoc/include
PROTO_DIR   := proto
OUT_DIR     := proto/gen

proto:
	$(PROTOC) \
		-I$(PROTOC_PATH) \
		-I$(PROTO_DIR) \
		--descriptor_set_out=$(OUT_DIR)/protos.desc \
		--include_imports \
		--go_out=paths=source_relative:$(OUT_DIR) \
		--go-grpc_out=paths=source_relative:$(OUT_DIR) \
		--go-vtproto_out=paths=source_relative:$(OUT_DIR) --go-vtproto_opt=features=marshal+unmarshal+size \
		$(PROTO_DIR)/*.proto

	@echo "Proto files generated in the 'proto' directory with vtprotobuf."

.PHONY: proto