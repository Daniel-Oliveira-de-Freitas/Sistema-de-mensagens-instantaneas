import asyncio
from aioquic.asyncio import serve
from aioquic.quic.configuration import QuicConfiguration

HOST = '127.0.0.1'  # Endereço IP do servidor
PORT = 12345  # Porta para comunicação

clients = set()  # Conjunto para rastrear os endereços dos clientes

async def handle_client(reader, writer):
    address = writer.get_extra_info('peername')
    if address not in clients:
        # Novo cliente conectado
        clients.add(address)
        print(f"Novo cliente conectado: {address}")

    while True:
        # Recebe mensagem do cliente
        data = await reader.read(1024)
        if not data:
            break

        print(f"Recebido do cliente {address}: {data.decode()}")

        # Envia a mensagem recebida para todos os outros clientes conectados
        for client_address in clients:
            if client_address != address:
                writer.write(data)

        await writer.drain()

    # Cliente desconectado
    clients.remove(address)
    print(f"Cliente desconectado: {address}")
    writer.close()

async def run_server():
    # Configuração do QUIC
    configuration = QuicConfiguration(is_client=False)

    # Inicia o servidor QUIC
    await serve(HOST, PORT, configuration=configuration, create_protocol=handle_client)

print(f"Servidor de chat iniciado em {HOST}:{PORT}")

asyncio.run(run_server())
