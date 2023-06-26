from ping3 import ping

# Medição básica de RTT
rtt = ping('https://localhost:5555')
print(f"RTT: {rtt} ms")

