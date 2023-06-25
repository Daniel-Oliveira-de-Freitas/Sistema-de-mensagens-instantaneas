import asyncio
import h3

class InstantMessagingClient:
    def __init__(self, host, port):
        self.host = host
        self.port = port
        self.loop = asyncio.get_event_loop()

    async def send_message(self, message):
        async with h3.connect(self.host, self.port) as conn:
            headers = [('content-length', str(len(message)))]
            response = await conn.request('POST', '/', headers=headers, body=message.encode())
            response_text = await response.read()
            print(f'Resposta do servidor: {response_text.decode()}')

    def run(self):
        message = input('Digite uma mensagem: ')
        self.loop.run_until_complete(self.send_message(message))

client = InstantMessagingClient('localhost', 8000)
client.run()
