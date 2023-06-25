import argparse
import asyncio
from aioquic.asyncio import connect
from aioquic.quic.connection import QuicConnectionProtocol
from aioquic.quic.configuration import QuicConfiguration
from aioquic.h3.connection import H3_ALPN

async def chat_client():
    parser = argparse.ArgumentParser(description="Chat client using aioquic")
    parser.add_argument("host", type=str, help="Server host name")
    parser.add_argument("port", type=int, help="Server port number")
    args = parser.parse_args()

    config = QuicConfiguration(
        is_client=True,
        alpn_protocols=["chat"],
        verify_mode="none",
    )

    async with connect(
        args.host,
        args.port,
        configuration=config,
        create_protocol=QuicConnectionProtocol,
    ) as client:
        while True:
            message = input("Enter message: ")
            await client.send_stream_data(1, message.encode())

if __name__ == "__main__":
    asyncio.run(chat_client())
