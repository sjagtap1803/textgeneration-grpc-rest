## Text Generation Client
This directory contains example scripts for gRPC as well as REST clients. In a new terminal, cd into this directory and run the following command from your virtual environment depending on the type of client:
* gRPC
```
$ python textgeneration_client.py
```
* REST
```
$ python textgeneration_rest_client.py
```
Alternatively, the REST endpoints can be tested using Postman or curl commands. Here are some curl commands to test the service.
1. `curl -X POST -d '{"prompt": "Far far away"}' http://localhost:50052/v1/textgeneration/unary-unary`
2. `curl -X POST -d '{"prompt": "Far far away"}' http://localhost:50052/v1/textgeneration/unary-stream`
3. `curl -X POST -d '{"prompt": "Far far away"}{"prompt": "Once upon a time"}' http://localhost:50052/v1/textgeneration/stream-unary`
4. `curl -X POST -d '{"prompt": "Far far away"}{"prompt": "Once upon a time"}' http://localhost:50052/v1/textgeneration/stream-stream`
5. `curl -X POST -F "file=@/path/to/textgeneration-grpc-rest/client/data/data1.txt" http://localhost:50052/v1/textgeneration/single-upload`
6. `curl -X POST -F "files=@/path/to/textgeneration-grpc-rest/client/data/data1.txt" -F "files=@/path/to/repo/client/data/data2.txt" http://localhost:50052/v1/textgeneration/multi-upload`