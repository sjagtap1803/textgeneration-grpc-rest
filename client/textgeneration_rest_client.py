import logging
import requests
import yaml
from yaml.loader import SafeLoader


class RestClient():
    def __init__(self, address, port, rest_route):
        self.rest_url = "http://" + address
        self.rest_port = port
        self.rest_route = rest_route
        self.connection_url = f"{self.rest_url}:{self.rest_port}/{self.rest_route}"
        self.timeout = 60
        self.headers = {'Content-Type': 'application/x-www-form-urlencoded'}

    def run_process_text_gen(self, prompt):
        connect_textgen = f"{self.connection_url}/unary-unary"
        data = '{"prompt": ' + f'"{prompt}"' + '}'
        req = requests.post(url=connect_textgen, headers=self.headers, data=data, timeout=self.timeout)
        return req

    def run_process_text_gen_multi_stream(self, prompt):
        connect_textgen = f"{self.connection_url}/unary-stream"
        data = '{"prompt": ' + f'"{prompt}"' + '}'
        req = requests.post(url=connect_textgen, headers=self.headers, data=data, timeout=self.timeout)
        return req

    def run_process_text_gen_stream_request(self, prompts: list):
        connect_textgen = f"{self.connection_url}/stream-unary"
        data = ""
        for prompt in prompts:
            data += '{"prompt": ' + f'"{prompt}"' + '}'
        req = requests.post(url=connect_textgen, headers=self.headers, data=data, timeout=self.timeout)
        return req

    def run_process_text_gen_stream(self, prompts: list):
        connect_textgen = f"{self.connection_url}/stream-stream"
        data = ""
        for prompt in prompts:
            data += '{"prompt": ' + f'"{prompt}"' + '}'
        req = requests.post(url=connect_textgen, headers=self.headers, data=data, timeout=self.timeout)
        return req

    def run_single_file_upload(self, file):
        connect_textgen = f"{self.connection_url}/single-upload"
        files = {
            'file': open(file, 'rb')
        }
        req = requests.post(url=connect_textgen, files=files)
        return req

    def run_multi_file_upload(self, files: list):
        connect_textgen = f"{self.connection_url}/multi-upload"
        files = [('files', open(file, 'rb')) for file in files]
        req = requests.post(url=connect_textgen, files=files)
        return req


def run():
    cfg_file = "./client_config.yaml"
    with open(cfg_file) as c_info_file:
        client_cfg = yaml.load(c_info_file, Loader=SafeLoader)

    rest_cfg = client_cfg["rest_info"]
    client = RestClient(rest_cfg["address"], rest_cfg["port"], rest_cfg["route"])

    print("Running Unary RPC with single prompt...")
    response = client.run_process_text_gen("Far far away there was a")
    print(repr(response.content.decode()))
    print("Success\n")

    print("Running Unary-Stream RPC...")
    response = client.run_process_text_gen_multi_stream("Far far away there was a")
    print(repr(response.content.decode()))
    print("Success\n")

    print("Running Stream-Unary RPC...")
    response = client.run_process_text_gen_stream_request(["Far far away", "There was a", "name was John Doe"])
    print(repr(response.content.decode()))
    print("Success\n")

    print("Running Stream-Stream RPC...")
    response = client.run_process_text_gen_stream(["Far far away", "There was a", "name was John Doe"])
    print(repr(response.content.decode()))
    print("Success\n")

    print("Running single file upload RPC...")
    response = client.run_single_file_upload("./data/data1.txt")
    print(repr(response.content.decode()))
    print("Success\n")

    print("Running multiple file upload RPC...")
    response = client.run_multi_file_upload(["./data/data1.txt", "./data/data2.txt", "./data/data3.txt"])
    print(repr(response.content.decode()))
    print("Success")
    

if __name__ == "__main__":
    logging.basicConfig()
    run()