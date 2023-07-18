import logging

import grpc
import textgeneration_pb2
import textgeneration_pb2_grpc


class GrpcClient():
    def __init__(self, address, port):
        self.address = address
        self.port = port
        self.channel = grpc.insecure_channel(f'{self.address}:{self.port}')
        self.stub = textgeneration_pb2_grpc.TextGenerationStub(self.channel)

    def run_process_text_gen(self, prompt):
        response = self.stub.ProcessTextGen(textgeneration_pb2.TextGenPrompt(prompt=prompt))
        print(repr(f'Generated text: {response.text}'))

    def run_process_text_gen_cleanup(self, prompt, cleanup=False):
        response = self.stub.ProcessTextGen(textgeneration_pb2.TextGenPrompt(prompt=prompt, cleanup=cleanup))
        print(repr(f'Generated text: {response.text}'))

    def run_process_text_gen_multi_stream(self, prompt):
        responses = self.stub.ProcessTextGenMultiStream(textgeneration_pb2.TextGenPrompt(prompt=prompt))
        for response in responses:
            print(repr(f'Generated text: {response.text}'))

    def run_process_text_gen_multi_stream_cleanup(self, prompt, cleanup=False):
        responses = self.stub.ProcessTextGenMultiStream(textgeneration_pb2.TextGenPrompt(prompt=prompt, cleanup=cleanup))
        for response in responses:
            print(repr(f'Generated text: {response.text}'))

    def stream_prompts(self, prompts, cleanup=False):
        for prompt in prompts:
            yield textgeneration_pb2.TextGenPrompt(prompt=prompt, cleanup=cleanup)

    def run_process_text_gen_stream_request(self, prompts):
        response = self.stub.ProcessTextGenStreamRequest(self.stream_prompts(prompts))
        print(repr(f'Generated text: {response.text}'))

    def run_process_text_gen_stream_request_cleanup(self, prompts, cleanup=False):
        response = self.stub.ProcessTextGenStreamRequest(self.stream_prompts(prompts, cleanup=cleanup))
        print(repr(f'Generated text: {response.text}'))

    def run_process_text_gen_stream(self, prompts):
        responses = self.stub.ProcessTextGenStream(self.stream_prompts(prompts))
        for response in responses:
            print(repr(f'Generated text: {response.text}'))

    def run_process_text_gen_stream_cleanup(self, prompts, cleanup=False):
        responses = self.stub.ProcessTextGenStream(self.stream_prompts(prompts, cleanup=cleanup))
        for response in responses:
            print(repr(f'Generated text: {response.text}'))

    def run_single_file_upload_filename(self, filename):
        response = self.stub.SingleFileUpload(textgeneration_pb2.FileUpload(filename=filename))
        print(repr(f'Generated text: {response.text}'))

    def run_single_file_upload_filedata(self, filedata):
        response = self.stub.SingleFileUpload(textgeneration_pb2.FileUpload(filedata=filedata))
        print(repr(f'Generated text: {response.text}'))

    def stream_filenames(self, files):
        for filename in files:
            yield textgeneration_pb2.FileUpload(filename=filename)

    def stream_filedata(self, files):
        for file in files:
            with open(file, "rb") as f:
                filedata = f.read()
                yield textgeneration_pb2.FileUpload(filedata=filedata)

    def run_multi_file_upload_filename(self, files):
        responses = self.stub.MultiFileUpload(self.stream_filenames(files))
        for response in responses:
            print(repr(f'Generated text: {response.text}'))

    def run_multi_file_upload_filedata(self, files):
        responses = self.stub.MultiFileUpload(self.stream_filedata(files))
        for response in responses:
            print(repr(f'Generated text: {response.text}'))


def run():
    client = GrpcClient("localhost", "50051")

    print("Running Unary RPC with single prompt...")
    client.run_process_text_gen("Far far away there was a")
    print("Running Unary RPC with clean-up...")
    client.run_process_text_gen_cleanup("Far far away there was a", cleanup=True)
    print("Success\n")

    print("Running Unary-Stream RPC...")
    client.run_process_text_gen_multi_stream("He was born on")
    print("Running Unary-Stream RPC with clean-up...")
    client.run_process_text_gen_multi_stream_cleanup("He was born on", cleanup=True)
    print("Success\n")

    print("Running Stream-Unary RPC...")
    client.run_process_text_gen_stream_request(["Far far away there was a", "Once upon a time", "In the land of Oz"])
    print("Running Stream-Unary RPC with clean-up...")
    client.run_process_text_gen_stream_request_cleanup(["Far far away there was a", "Once upon a time", "In the land of Oz"], cleanup=True)
    print("Success\n")

    print("Running bi-directional Stream RPC...")
    client.run_process_text_gen_stream(["Far far away there was a", "Once upon a time", "In the land of Oz"])
    print("Running bi-directional Stream RPC with clean-up...")
    client.run_process_text_gen_stream_cleanup(["Far far away there was a", "Once upon a time", "In the land of Oz"], cleanup=True)
    print("Success\n")

    print("Running single file upload RPC with file path")
    client.run_single_file_upload_filename("./data/data1.txt")
    print("Running single file upload RPC with binary data")
    with open("./data/data1.txt", "rb") as data:
        client.run_single_file_upload_filedata(data.read())
    print("Success\n")

    print("Running multi file upload RPC with file path")
    client.run_multi_file_upload_filename(["./data/data1.txt", "./data/data2.txt", "./data/data3.txt"])
    print("Running multi file upload RPC with binary data")
    client.run_multi_file_upload_filedata(["./data/data1.txt", "./data/data2.txt", "./data/data3.txt"])
    print("Success")


if __name__ == "__main__":
    logging.basicConfig()
    run()