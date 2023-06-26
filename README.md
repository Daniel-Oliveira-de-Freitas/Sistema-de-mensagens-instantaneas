# Sistema-de-mensagens-instantaneas

## Go instalação
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
