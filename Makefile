SHELL := /bin/bash

.PHONY: run-gateway run-shorten

loadenv:
	source .env

run-gateway:
	go run apigateway/*.go

run-shorten:
	go run services/shortener/*.go

run-retriever:
	go run services/retriever/*.go


gen:
	@protoc \
		--proto_path=proto "proto/shortenUrl.proto" \
		--go_out=proto/genproto/shortenpb --go_opt=paths=source_relative \
  		--go-grpc_out=proto/genproto/shortenpb --go-grpc_opt=paths=source_relative
	@protoc \
		--proto_path=proto "proto/retrieveUrl.proto" \
		--go_out=proto/genproto/retrievepb --go_opt=paths=source_relative \
  		--go-grpc_out=proto/genproto/retrievepb --go-grpc_opt=paths=source_relative