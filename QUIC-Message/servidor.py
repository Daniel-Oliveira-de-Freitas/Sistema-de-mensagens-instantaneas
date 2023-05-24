import asyncio
import argparse
from aioquic.asyncio import serve
from aioquic.quic.connection import QuicConnectionProtocol
from aioquic.quic.configuration import QuicConfiguration
from aioquic.quic.events import DatagramFrameReceived, QuicEvent, QuicReceiveStreamDataAvailable, StreamData
from typing import Dict

class ChatServerProtocol(QuicConnectionProtocol):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.clients: Dict[str, QuicConnectionProtocol] = {}

    async def handle_event(self, event: QuicEvent):
        if isinstance(event, DatagramFrameReceived):
            message = event.data.decode()
            self.broadcast(message)
        elif isinstance(event, QuicReceiveStreamDataAvailable):
            stream_data = await self._quic_stream_receive_data(event.stream_id, event.max_data)
            message = stream_data.data.decode()
            self.broadcast(message)

    def broadcast(self, message: str):
        for client_id, client in self.clients.items():
            try:
                client.send_datagram(message.encode())
            except Exception:
                self.clients.pop(client_id)

    def quic_accept(self, cid: bytes):
        super().quic_accept(cid)
        self.clients[self._quic_address[0]] = self

async def main():
    parser = argparse.ArgumentParser(description="Chat server using aioquic")
    parser.add_argument("bind", type=str, help="Bind to IP address")
    parser.add_argument("port", type=int, help="Port to listen on")
    args = parser.parse_args()

    config = QuicConfiguration(is_server=True)
    config.load_cert_chain("cert.pem", "key.pem")

    server = await serve(
        args.bind,
        args.port,
        None,
        None,
        create_protocol=ChatServerProtocol,
        configuration=config,
    )

if __name__ == "__main__":
    asyncio.run(main())
