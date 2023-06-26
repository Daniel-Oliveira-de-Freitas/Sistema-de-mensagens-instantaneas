import asyncio
from aioquic.asyncio import connect
from aioquic.asyncio.protocol import QuicConnectionProtocol

SERVER_HOST = '127.0.0.1'  # Endereço IP do servidor
SERVER_PORT = 12345  # Porta para comunicação

class ClientProtocol(QuicConnectionProtocol):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.transport = None

    def quic_event_received(self, event):
        if isinstance(event, asyncio.BaseTransport):
            self.transport = event

async def run_client():
    async with connect(SERVER_HOST, SERVER_PORT, create_protocol=ClientProtocol) as protocol:
        while True:
            # Leitura da mensagem do usuário
            message = input("Digite uma mensagem (ou 'sair' para encerrar): ")

            if message.lower() == 'sair':
                break

            # Envio da mensagem ao servidor
            protocol._quic.send_stream_data(0, message.encode())

            # Recebe a resposta do servidor
            data = await protocol._quic.receive_data()
            print(f"Resposta do servidor: {data.decode()}")

loop = asyncio.get_event_loop()
loop.run_until_complete(run_client())
loop.close()
