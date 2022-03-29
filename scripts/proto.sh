#!/bin/bash
set -e

readonly service="$1"

# protoc -I . --grpc-gateway_out ./gen/go \
#     --grpc-gateway_opt logtostderr=true \
#     --grpc-gateway_opt paths=source_relative \
#     --grpc-gateway_opt grpc_api_configuration=path/to/config.yaml \
#     --grpc-gateway_opt standalone=true \
#     your/service/v1/your_service.proto

# protoc -I ./api/protobuf \
#    --go_out ./internal/common/genproto --go_opt paths=source_relative \
#    --go-grpc_out ./internal/common/genproto --go-grpc_opt paths=source_relative \
#    --grpc-gateway_out ./internal/common/genproto --grpc-gateway_opt paths=source_relative \
#    --grpc-gateway_opt grpc_api_configuration=path/to/config.yaml \
#    ./api/protobuf/$1
