import requests


class RestClient():
    def __init__(self, address, port, rest_route):
        self.rest_url = "http://" + address
        self.rest_port = port
        self.rest_route = rest_route
        self.connection_url = f"{self.rest_url}:{self.rest_port}/{self.rest_route}"
        self.timeout = 60
        self.headers = {'Content-Type': 'application/x-www-form-urlencoded'}

    def guide_process_text_gen(self, prompt):
        connect_textgen = f"{self.connection_url}/unary-unary"
        data = '{"prompt": ' + f'"{prompt}"' + '}'
        req = requests.post(url=connect_textgen, headers=self.headers, data=data, timeout=self.timeout)
        return req

    def guide_process_text_gen_multi_stream(self, prompt):
        connect_textgen = f"{self.connection_url}/unary-stream"
        data = '{"prompt": ' + f'"{prompt}"' + '}'
        req = requests.post(url=connect_textgen, headers=self.headers, data=data, timeout=self.timeout)
        return req

    def guide_process_text_gen_stream_request(self, prompts: list):
        connect_textgen = f"{self.connection_url}/stream-unary"
        data = ""
        for prompt in prompts:
            data += '{"prompt": ' + f'"{prompt}"' + '}'
        req = requests.post(url=connect_textgen, headers=self.headers, data=data, timeout=self.timeout)
        return req

    def guide_process_text_gen_stream(self, prompts: list):
        connect_textgen = f"{self.connection_url}/stream-stream"
        data = ""
        for prompt in prompts:
            data += '{"prompt": ' + f'"{prompt}"' + '}'
        req = requests.post(url=connect_textgen, headers=self.headers, data=data, timeout=self.timeout)
        return req


def run():
    client = RestClient("localhost", "50052", "v1/textgeneration")

    print("Testing Unary RPC with single prompt...")
    response = client.guide_process_text_gen("Far far away there was a")
    print(response.content.decode())
    print("Success")

    print("Testing Unary-Stream RPC...")
    response = client.guide_process_text_gen_multi_stream("Far far away there was a")
    print(response.content.decode())
    print("Success")

    print("Testing Stream-Unary RPC...")
    response = client.guide_process_text_gen_stream_request(["Far far away", "There was a", "name was John Doe"])
    print(response.content.decode())
    print("Success")

    print("Testing Stream-Stream RPC...")
    response = client.guide_process_text_gen_stream(["Far far away", "There was a", "name was John Doe"])
    print(response.content.decode())
    print("Success")
    

if __name__ == "__main__":
    run()