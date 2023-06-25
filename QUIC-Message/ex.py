import asyncio
from aioquic.asyncio.client import connect
from aioquic.asyncio.protocol import QuicConnectionProtocol
from aioquic.quic.configuration import QuicConfiguration
from aioquic.h3.connection import H3_ALPN


class MyQuicClientProtocol(QuicConnectionProtocol):
    
    def data_received(self, data):
        for event in self._quic._handle_event():
            if isinstance(event, self.QuicEventConnected):
                # A conexão QUIC foi estabelecida
                self._quic.send_h3_request(
                    method="GET",
                    authority="example.com",
                    path="/",
                    headers=[(b"user-agent", b"MyQUICClient/1.0")],
                )
            # se foi recebidos cabeçalhos de resposta
            elif isinstance(event, self.QuicEventHeadersReceived):
                # Cabeçalhos de resposta HTTP/3 foram recebidos
                headers = event.headers
                print(f"Received headers: {headers}")
            # se foi recebidos dados de resposta
            elif isinstance(event, self.QuicEventDataReceived):
                # Dados de resposta HTTP/3 foram recebidos
                data = event.data
                print(f"Received data: {data.decode()}")


async def run_quic_client():
    configuration = QuicConfiguration(is_client=True)
    configuration.alpn_protocols = H3_ALPN
    async with connect("example.com", 443, configuration=configuration) as protocol:
        await protocol.wait_connected()
        print('conectou')
        print(configuration.server_name)
        await asyncio.sleep(1)  # Aguarda um segundo para que a conexão seja estabelecida completamente


asyncio.run(run_quic_client())