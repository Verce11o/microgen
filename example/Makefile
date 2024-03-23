gen-proto:
	mkdir -p protos/gen/go
	protoc -I protos protos/*.proto --go_out=protos/gen/go/ --go_opt=paths=source_relative --go-grpc_out=protos/gen/go/ --go-grpc_opt=paths=source_relative
