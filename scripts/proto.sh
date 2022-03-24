#!/bin/bash
set -e

readonly service="$1"

protoc -I ./api/protobuf \
   --go_out ./internal/common/genproto/tag --go_opt paths=source_relative \
   --go-grpc_out ./internal/common/genproto/tag --go-grpc_opt paths=source_relative \
   --grpc-gateway_out=internal/common/genproto/tag --grpc-gateway_opt paths=source_relative \
  ./api/protobuf/tag/tag.proto
