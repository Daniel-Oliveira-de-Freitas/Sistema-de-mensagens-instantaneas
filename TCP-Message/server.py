import socket
import threading

class ChatServer:
    def __init__(self):
        self.connections = []
        self.server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        self.server_socket.bind(("0.0.0.0", 5555))
        self.server_socket.listen(5)

    def start(self):
        while True:
            conn, addr = self.server_socket.accept()
            self.connections.append(conn)
            threading.Thread(target=self.handle_client, args=(conn, addr)).start()

    def handle_client(self, conn, addr):
        while True:
            try:
                message = conn.recv(1024).decode().strip()
                if not message:
                    continue
                print(f"New message: {message}")
                for connection in self.connections:
                    if connection is not conn:
                        connection.send(f"{addr[0]}: {message}".encode())
            except:
                self.connections.remove(conn)
                conn.close()
                return

if __name__ == "__main__":
    server = ChatServer()
    server.start()
