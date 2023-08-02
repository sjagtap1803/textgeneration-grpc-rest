## Reverse Proxy Gateway
This Go module lets the user create a RESTful reverse proxy for the gRPC server. In a new terminal, cd into this directory and run the following command to start the reverse proxy.
```
$ source gateway_config.sh && go run .
```
The reverse proxy runs at `localhost:50052` by default. This can be changed by editing `gateway_config.sh`.