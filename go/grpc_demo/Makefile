.PHONY: proto clean

PROTO_DIR := proto
OUT_DIR := proto/gen
GENERATOR_CMD := go run ./cmd/gen/main.go

clean:
ifeq ($(OS),Windows_NT)
	@cmd /c if exist $(subst /,\,$(OUT_DIR)) rmdir /Q /S $(subst /,\,$(OUT_DIR))
else
	@rm -rf $(OUT_DIR)
endif
	@echo "Clean"


proto: clean
ifeq ($(OS),Windows_NT)
	@cmd /c "mkdir $(subst /,\,$(OUT_DIR))"
else
	@mkdir -p $(OUT_DIR)
endif
	@protoc \
       --proto_path=$(PROTO_DIR) \
       --descriptor_set_out=$(OUT_DIR)/protos.desc --include_imports $(PROTO_DIR)/*.proto \
       --go_out=$(PROTO_DIR) \
       --go-grpc_out=$(PROTO_DIR) \
       --go-vtproto_out=$(PROTO_DIR) --go-vtproto_opt=features=marshal+unmarshal+size \
       $(PROTO_DIR)/*.proto

	@echo "Proto files generated in the 'proto' directory with vtprotobuf."
