.PHONY: generate

generate:
	rm -rf examplepb
	mkdir -p examplepb
	protoc --proto_path=. \
	       --go_out=examplepb --go-grpc_out=examplepb \
	       --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
	       example.proto
