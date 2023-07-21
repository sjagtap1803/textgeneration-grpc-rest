package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gw "grpc-gateway/grpc_go"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")

	// http to gRPC status code mapping
	errorMap = map[string]string{
		"StatusOK":                  "0",
		"StatusBadRequest":          "3",
		"StatusInternalServerError": "13",
		"StatusServiceUnavailable":  "14",
	}
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterTextGenerationHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Handle REST paths for uploading files
	err = mux.HandlePath("POST", "/v1/textgeneration/single-upload", UploadSingleFile)
	if err != nil {
		return err
	}

	err = mux.HandlePath("POST", "/v1/textgeneration/multi-upload", UploadMultipleFiles)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	fmt.Println("grpc gateway listening on port 50052")
	return http.ListenAndServe(":50052", mux)
}

func createErrorResponse(status string, err error) string {
	return fmt.Sprintf("{\"code\":%s, \"message\":%s, \"details\": []}", errorMap[status], strconv.Quote(err.Error()))
}

func UploadSingleFile(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	// Parse Form from the request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, createErrorResponse("StatusBadRequest", err), http.StatusBadRequest)
		return
	}

	// Read file components
	multipartFile, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, createErrorResponse("StatusBadRequest", err), http.StatusBadRequest)
		return
	}
	defer multipartFile.Close()

	// Create buffer and read file contents
	buffer := &bytes.Buffer{}
	_, err = io.Copy(buffer, multipartFile)
	if err != nil {
		http.Error(w, createErrorResponse("StatusBadRequest", err), http.StatusBadRequest)
		return
	}

	// Convert Buffer to byte slice
	filedata := buffer.Bytes()

	// Setup grpc client stub
	conn, err := grpc.Dial(*grpcServerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, createErrorResponse("StatusServiceUnavailable", err), http.StatusServiceUnavailable)
		return
	}
	defer conn.Close()
	client := gw.NewTextGenerationClient(conn)

	// Call SingleFileUpload rpc
	ctx := context.Background()
	response, err := client.SingleFileUpload(ctx, &gw.FileUpload{UploadOneof: &gw.FileUpload_Filedata{filedata}})
	if err != nil {
		http.Error(w, createErrorResponse("StatusInternalServerError", err), http.StatusInternalServerError)
		return
	}

	// Write response to client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"text\": " + strconv.Quote(response.Text) + "}"))
}

func UploadMultipleFiles(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	// Parse Form from the request
	err := r.ParseMultipartForm(20 << 32)
	if err != nil {
		if err != nil {
			http.Error(w, createErrorResponse("StatusBadRequest", err), http.StatusBadRequest)
			return
		}
	}

	// Setup grpc client stub
	conn, err := grpc.Dial(*grpcServerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, createErrorResponse("StatusServiceUnavailable", err), http.StatusServiceUnavailable)
		return
	}
	defer conn.Close()
	client := gw.NewTextGenerationClient(conn)

	// Call MultiFileUpload rpc and create channel
	ctx := context.Background()
	stream, err := client.MultiFileUpload(ctx)
	if err != nil {
		http.Error(w, createErrorResponse("StatusInternalServerError", err), http.StatusInternalServerError)
		return
	}
	waitc := make(chan struct{})

	// Iterate over uploaded files
	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		http.Error(w, createErrorResponse("StatusBadRequest", err), http.StatusBadRequest)
		return
	}

	buffer := &bytes.Buffer{}
	responsebody := ""

	// receive stream of responses from server via go routine and collect all responses in a string
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				http.Error(w, createErrorResponse("StatusInternalServerError", err), http.StatusInternalServerError)
				return
			}
			responsebody = responsebody + "{\"result\": {\"text\": " + strconv.Quote(in.Text) + "}}\n"
		}
	}()

	// send stream of requests to server via channel
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			http.Error(w, createErrorResponse("StatusBadRequest", err), http.StatusBadRequest)
			return
		}
		defer f.Close()
		_, err = io.Copy(buffer, f)
		if err != nil {
			http.Error(w, createErrorResponse("StatusBadRequest", err), http.StatusBadRequest)
			return
		}

		filedata := buffer.Bytes()
		err = stream.Send(&gw.FileUpload{UploadOneof: &gw.FileUpload_Filedata{filedata}})
		if err != nil {
			http.Error(w, createErrorResponse("StatusInternalServerError", err), http.StatusInternalServerError)
			return
		}
		buffer.Reset()
	}
	stream.CloseSend()
	<-waitc
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responsebody))
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
