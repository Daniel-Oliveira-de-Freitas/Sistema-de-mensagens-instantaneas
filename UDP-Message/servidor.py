import socket

HOST = '127.0.0.1'  # Endereço IP do servidor
PORT = 12345  # Porta para comunicação

# Criação do socket UDP
server_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
server_socket.bind((HOST, PORT))

print(f"Servidor de chat iniciado em {HOST}:{PORT}")

clients = set()  # Conjunto para rastrear os endereços dos clientes

while True:
    # Recebe mensagem do cliente
    data, address = server_socket.recvfrom(1024)

    if address not in clients:
        # Novo cliente conectado
        clients.add(address)
        print(f"Novo cliente conectado: {address}")

    print(f"Recebido do cliente {address}: {data.decode()}")

    # Envia a mensagem recebida para todos os outros clientes conectados
    for client_address in clients:
        if client_address != address:
            server_socket.sendto(data, client_address)
