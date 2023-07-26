import logging
import yaml
from yaml.loader import SafeLoader
from concurrent import futures
from transformers import pipeline


import grpc
import textgeneration_pb2
import textgeneration_pb2_grpc


class TextGenerationServicer(textgeneration_pb2_grpc.TextGenerationServicer):
    def __init__(self, model):
        super().__init__()
        self.pipeline = pipeline(model=model)
    
    def ProcessTextGen(self, request, context):
        output_txt = self.pipeline(request.prompt, clean_up_tokenization_spaces=request.cleanup)[0]["generated_text"]
        return textgeneration_pb2.TextGenResponse(text=output_txt)
    
    def ProcessTextGenMultiStream(self, request, context):
        for i in range(3):
            output_txt = self.pipeline(request.prompt, clean_up_tokenization_spaces=request.cleanup)[0]["generated_text"]
            yield textgeneration_pb2.TextGenResponse(text=output_txt)
    
    def ProcessTextGenStreamRequest(self, request_iterator, context):
        output_txt = ""
        for request in request_iterator:
            output_txt += self.pipeline(request.prompt, clean_up_tokenization_spaces=request.cleanup)[0]["generated_text"]
            output_txt += " "
        return textgeneration_pb2.TextGenResponse(text=output_txt[:-1])
    
    def ProcessTextGenStream(self, request_iterator, context):
        for request in request_iterator:
            output_txt = self.pipeline(request.prompt, clean_up_tokenization_spaces=request.cleanup)[0]["generated_text"]
            yield textgeneration_pb2.TextGenResponse(text=output_txt)
    
    def SingleFileUpload(self, request, context):
        if request.filename:
            with open(request.filename, "rb") as f:
                input_text = f.read()
        else:
            input_text = request.filedata
        output_txt = self.pipeline(input_text.decode().strip(), clean_up_tokenization_spaces=True)[0]["generated_text"]
        return textgeneration_pb2.TextGenResponse(text=output_txt)

    def MultiFileUpload(self, request_iterator, context):
        for request in request_iterator:
            if request.filename:
                with open(request.filename, "rb") as f:
                    input_text = f.read()
            else:
                input_text = request.filedata
            output_txt = self.pipeline(input_text.decode().strip(), clean_up_tokenization_spaces=True)[0]["generated_text"]
            yield textgeneration_pb2.TextGenResponse(text=output_txt)
        

def serve():
    cfg_file = "./server_config.yaml"
    with open(cfg_file) as c_info_file:
        server_cfg = yaml.load(c_info_file, Loader=SafeLoader)

    port = server_cfg["server_info"]["server_port"]
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    textgeneration_pb2_grpc.add_TextGenerationServicer_to_server(TextGenerationServicer(model=server_cfg["model_info"]["model"]), server)
    server.add_insecure_port('[::]:' + port)
    server.start()
    print(f'grpc server is listening on port {port}')
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()