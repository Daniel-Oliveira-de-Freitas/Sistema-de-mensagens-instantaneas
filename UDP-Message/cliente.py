import socket

HOST = '127.0.0.1'  # Endereço IP do servidor
PORT = 12345  # Porta para comunicação

# Criação do socket UDP
client_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

while True:
    # Leitura da mensagem do usuário
    message = input("Digite uma mensagem (ou 'sair' para encerrar): ")

    if message.lower() == 'sair':
        break

    # Envio da mensagem ao servidor
    client_socket.sendto(message.encode(), (HOST, PORT))

    # Recebe a resposta do servidor
    data, _ = client_socket.recvfrom(1024)
    print(f"Resposta do servidor: {data.decode()}")

# Encerramento do socket do cliente
client_socket.close()
