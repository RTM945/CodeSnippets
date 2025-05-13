.PHONY: proto

proto:
	@protoc \
		--proto_path=proto \
		--go_out=proto \
		--go-grpc_out=proto \
		--go-vtproto_out=proto --go-vtproto_opt=features=marshal+unmarshal+size \
		proto/*.proto
	@echo "Proto files generated in the 'proto' directory with vtprotobuf."
