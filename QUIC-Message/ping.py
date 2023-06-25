from ping3 import ping, verbose_ping

# Medição básica de RTT
rtt = ping('https://localhost:5555')
print(f"RTT: {rtt} ms")

