# Sistema-de-mensagens-instantaneas
## Descrição
A implementação presente neste repostitório trata-se de uma aplicação desenvolvida em cima da arquitetura client/server, cujo objetivo do client é enviar mensagens para o server, onde o mesmo capta/escuta as requisições, por meio do protocolo QUIC.

### TCP-MESSAGE & UDP-Message
O mesmo se replica para o diretório [TCP-Message](https://github.com/Daniel-Oliveira-de-Freitas/Sistema-de-mensagens-instantaneas/tree/main/TCP-Message), que por meio serviu de cenário exemplificativo para o desenvolvimento do [QUIC-Message-GO](https://github.com/Daniel-Oliveira-de-Freitas/Sistema-de-mensagens-instantaneas/tree/main/QUIC-Message-GO).

### QUIC-CONNECTION
Neste, diretório [QUIC-CONNECTION](https://github.com/Daniel-Oliveira-de-Freitas/Sistema-de-mensagens-instantaneas/tree/main/QUIC-CONNECTION), os mantenedores fizeram uma tentativa do cenário citado, porém no decorrer do progresso aparecerão diversos problemas (dependências, importações), contudo obtivemos sucesso com o arquivo [connection.py]([caminho/do/arquivo.py](https://github.com/Daniel-Oliveira-de-Freitas/Sistema-de-mensagens-instantaneas/blob/main/QUIC-CONNECTION/connection.py)) este faz uma conexão com HTTP3 por meio da biblioteca aioquic.

### QUIC-Message-GO

#### Go instalação
Para linux Ubuntu 22.04.2

Go 1.19

Quic-go 0.31.1

Instala nos dois diretórios client e server
- wget https://golang.org/dl/go1.19.linux-amd64.tar.gz
- tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz
- export PATH=$PATH:/usr/local/go/bin

## Go instalação via terminal Linux

sudo apt install golang-go

## Gerar os certs
./generate_cert.sh


## Executar
- cd server
- go build main.go
- ./main


- cd client
- go build client.go
- ./client
