import logging

import grpc
import textgeneration_pb2
import textgeneration_pb2_grpc


def guide_process_text_gen(stub, prompt):
    response = stub.ProcessTextGen(textgeneration_pb2.TextGenPrompt(prompt=prompt))
    print(f'Generated text: {response.text}')


def guide_process_text_gen_cleanup(stub, prompt, cleanup=False):
    response = stub.ProcessTextGen(textgeneration_pb2.TextGenPrompt(prompt=prompt, cleanup=cleanup))
    print(f'Generated text: {response.text}')


def guide_process_text_gen_multi_stream(stub, prompt):
    responses = stub.ProcessTextGenMultiStream(textgeneration_pb2.TextGenPrompt(prompt=prompt))
    for response in responses:
        print(f'Generated text: {response.text}')


def guide_process_text_gen_multi_stream_cleanup(stub, prompt, cleanup=False):
    responses = stub.ProcessTextGenMultiStream(textgeneration_pb2.TextGenPrompt(prompt=prompt, cleanup=cleanup))
    for response in responses:
        print(f'Generated text: {response.text}')


def generate_prompts(prompts, cleanup=False):
    for prompt in prompts:
        yield textgeneration_pb2.TextGenPrompt(prompt=prompt, cleanup=cleanup)


def guide_process_text_gen_stream_request(stub, prompts):
    response = stub.ProcessTextGenStreamRequest(generate_prompts(prompts))
    print(f'Generated text: {response.text}')


def guide_process_text_gen_stream_request_cleanup(stub, prompts, cleanup=False):
    response = stub.ProcessTextGenStreamRequest(generate_prompts(prompts, cleanup=cleanup))
    print(f'Generated text: {response.text}')


def guide_process_text_gen_stream(stub, prompts):
    responses = stub.ProcessTextGenStream(generate_prompts(prompts))
    for response in responses:
        print(f'Generated text: {response.text}')


def guide_process_text_gen_stream_cleanup(stub, prompts, cleanup=False):
    responses = stub.ProcessTextGenStream(generate_prompts(prompts, cleanup=cleanup))
    for response in responses:
        print(f'Generated text: {response.text}')


def guide_single_file_upload_filename(stub, filename):
    response = stub.SingleFileUpload(textgeneration_pb2.FileUpload(filename=filename))
    print(f'Generated text: {response.text}')


def guide_single_file_upload_filedata(stub, filedata):
    response = stub.SingleFileUpload(textgeneration_pb2.FileUpload(filedata=filedata))
    print(f'Generated text: {response.text}')


def stream_filenames(files):
    for filename in files:
        yield textgeneration_pb2.FileUpload(filename=filename)


def stream_filedata(files):
    for file in files:
        with open(file, "rb") as f:
            filedata = f.read()
            yield textgeneration_pb2.FileUpload(filedata=filedata)


def guide_multi_file_upload_filename(stub, files):
    responses = stub.MultiFileUpload(stream_filenames(files))
    for response in responses:
        print(f'Generated text: {response.text}')


def guide_multi_file_upload_filedata(stub, files):
    responses = stub.MultiFileUpload(stream_filedata(files))
    for response in responses:
        print(f'Generated text: {response.text}')


def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = textgeneration_pb2_grpc.TextGenerationStub(channel)
        
        print("Testing Unary RPC with single prompt...")
        guide_process_text_gen(stub, "Far far away there was a")
        print("Testing Unary RPC with clean-up...")
        guide_process_text_gen_cleanup(stub, "Far far away there was a", cleanup=True)
        print("Success\n")

        print("Testing Unary-Stream RPC...")
        guide_process_text_gen_multi_stream(stub, "He was born on")
        print("Testing Unary-Stream RPC with clean-up...")
        guide_process_text_gen_multi_stream_cleanup(stub, "He was born on", cleanup=True)
        print("Success\n")

        print("Testing Stream-Unary RPC...")
        guide_process_text_gen_stream_request(stub, ["Far far away there was a", "Once upon a time", "In the land of Oz"])
        print("Testing Stream-Unary RPC with clean-up...")
        guide_process_text_gen_stream_request_cleanup(stub, ["Far far away there was a", "Once upon a time", "In the land of Oz"], cleanup=True)
        print("Success")

        print("Testing bi-directional Stream RPC...")
        guide_process_text_gen_stream(stub, ["Far far away there was a", "Once upon a time", "In the land of Oz"])
        print("Testing bi-directional Stream RPC with clean-up...")
        guide_process_text_gen_stream_cleanup(stub, ["Far far away there was a", "Once upon a time", "In the land of Oz"], cleanup=True)
        print("Success")

        print("Testing single file upload RPC with file path")
        guide_single_file_upload_filename(stub, "./data/data1.txt")
        print("Testing single file upload RPC with binary data")
        with open("./data/data1.txt", "rb") as data:
            guide_single_file_upload_filedata(stub, data.read())
        print("Success")

        print("Testing multi file upload RPC with file path")
        guide_multi_file_upload_filename(stub, ["./data/data1.txt", "./data/data2.txt", "./data/data3.txt"])
        print("Testing multi file upload RPC with binary data")
        guide_multi_file_upload_filedata(stub, ["./data/data1.txt", "./data/data2.txt", "./data/data3.txt"])
        print("Success")


if __name__ == "__main__":
    logging.basicConfig()
    run()