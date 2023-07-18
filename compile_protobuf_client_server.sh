PROTO_DIR=./protos
PROTO_FILE_PYTHON=${PROTO_DIR}/textgeneration.proto

OUT_PATH_CLIENT=./client
OUT_PATH_SERVER=./server

# Generate grpc code in python for client
python -m grpc_tools.protoc -I ${PROTO_DIR} \
    --python_out ${OUT_PATH_CLIENT} \
    --pyi_out ${OUT_PATH_CLIENT} \
    --grpc_python_out ${OUT_PATH_CLIENT} \
    ${PROTO_FILE_PYTHON}

# Generate grpc code in python for server
python -m grpc_tools.protoc -I ${PROTO_DIR} \
    --python_out ${OUT_PATH_SERVER} \
    --pyi_out ${OUT_PATH_SERVER} \
    --grpc_python_out ${OUT_PATH_SERVER} \
    ${PROTO_FILE_PYTHON}