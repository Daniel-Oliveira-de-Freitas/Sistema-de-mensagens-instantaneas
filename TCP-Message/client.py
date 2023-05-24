import socket
import threading

class ChatClient:
    def __init__(self, username, host, port):
        self.username = username
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.socket.connect((host, port))
        threading.Thread(target=self.receive_messages).start()

    def receive_messages(self):
        while True:
            message = self.socket.recv(1024).decode().strip()
            if not message:
                continue
            print(message)

    def send_message(self, message):
        self.socket.send(f"{self.username}: {message}".encode())

if __name__ == "__main__":
    username = input("Enter your username: ")
    host = input("Enter the server IP address: ")
    port = 5555
    client = ChatClient(username, host, port)

    while True:
        message = input()
        client.send_message(message)
