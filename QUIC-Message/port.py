import socket

def check_port_availability(port):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        try:
            s.bind(('localhost', port))
            return True
        except OSError:
            return False

# Exemplo de uso
port = 4433
if check_port_availability(port):
    print(f"A porta {port} está disponível.")
else:
    print(f"A porta {port} não está disponível.")
