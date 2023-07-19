# Text Generation Service using gRPC-Gateway
This repository hosts code for a text generation service that supports both gRPC as well as RESTful endpoints. The server utilzes a light-weight HuggingFace model to perform text generation and has been implemented using gRPC generated Python code. For exposing the gRPC sever as a set of REST endpoints, the gRPC-Gateway (https://github.com/grpc-ecosystem/grpc-gateway/tree/main) Go module has been used to set up a reverse proxy for handling HTTP methods.

## Repository Structure
The directories in this repository are as follows:
* `server`: Contains server implementation with gRPC generated Python code.
* `gateway`: Contains a Go module for running the reverse proxy (gRPC-Gateway).
* `client`: Contains code for both gRPC and REST clients.
* `protos`: Contains protocol buffer definitions.
* `openapiv2`: Contains a `.swagger.json` file for generating API documentation.

The repo also contains scripts to re-generate gRPC code in both Python (`compile_protobuf_client_server.sh`) and Go (`compile_protobuf_gateway.sh`) in case the protocol buffer definitions change.

## Instructions
### Python
1. Create a virtual environment and activate it. This can be done using Anaconda or venv.
2. Install the required packages
```
$ pip install -r requirements.txt
```
### Go
1. Download and install Go (https://go.dev/doc/install)
2. Install the protocol compiler.
```
$ apt-get install -y protobuf-compiler
```
3. Download some dependencies and update `PATH`.
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
$ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
$ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
$ export PATH="$PATH:$(go env GOPATH)/bin"

The compile scripts should work after following the above instructions. Further instructions can be found in individual directories.