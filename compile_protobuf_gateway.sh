PROTO_DIR=./protos
PROTO_FILE_GO=${PROTO_DIR}/textgeneration_go.proto
PROTO_FILE_GATEWAY=${PROTO_DIR}/textgeneration_gateway.proto

GATEWAY_DIR=./gateway
OUT_PATH=${GATEWAY_DIR}/grpc_go
OPENAPI_V2_PATH=./openapiv2

# Generate grpc code in go
protoc -I ${PROTO_DIR} \
    --go_out ${OUT_PATH} \
    --go_opt paths=source_relative \
    --go-grpc_out ${OUT_PATH} \
    --go-grpc_opt paths=source_relative \
    ${PROTO_FILE_GO}

# Generate grpc code in go for gateway
protoc -I ${PROTO_DIR} \
    --grpc-gateway_out ${OUT_PATH} \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    ${PROTO_FILE_GATEWAY}

# Generate swagger docs
protoc -I ${PROTO_DIR} \
    --openapiv2_out ${OPENAPI_V2_PATH} \
    --openapiv2_opt logtostderr=true \
    ${PROTO_FILE_GATEWAY}