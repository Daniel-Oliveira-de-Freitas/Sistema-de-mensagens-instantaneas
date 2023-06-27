<p align="center">
  <h1 align="center">📧 Sistema de Mensagens Instantâneas</h1>
</p>

## Descrição
A implementação presente neste repositório trata-se de uma aplicação desenvolvida em cima da arquitetura client/server, cujo objetivo do client é enviar mensagens para o server, onde o mesmo capta/escuta as requisições, por meio do protocolo de transporte QUIC.
<hr>

### TCP-Message & UDP-Message
O mesmo se replica para o diretório [TCP-Message](https://github.com/Daniel-Oliveira-de-Freitas/Sistema-de-mensagens-instantaneas/tree/main/TCP-Message), que por meio serviu de cenário exemplificativo para o desenvolvimento do [QUIC-Message-GO](https://github.com/Daniel-Oliveira-de-Freitas/Sistema-de-mensagens-instantaneas/tree/main/QUIC-Message-GO).

### QUIC-CONNECTION
Neste diretório [QUIC-CONNECTION](https://github.com/Daniel-Oliveira-de-Freitas/Sistema-de-mensagens-instantaneas/tree/main/QUIC-CONNECTION), os mantenedores fizeram uma tentativa do cenário citado, porém no decorrer do progresso aparecerão diversos problemas (dependências, importações), contudo obtivemos sucesso com o arquivo [connection.py](https://github.com/Daniel-Oliveira-de-Freitas/Sistema-de-mensagens-instantaneas/blob/main/QUIC-CONNECTION/connection.py) este faz uma conexão com HTTP3 por meio da biblioteca aioquic.
<hr>

### QUIC-Message-GO
Em suma, escolhemos elaborar/desenvolver com a biblioteca Quic-go version 0.31.1 com o Golang version 1.19 no complemento do Gin Web framework baseado na arquitetura client/server.

#### Pré-requesito
| Sistema Operacional | Versão do Linux | Go   | Quic-go   |
|---------------------|-----------------|------|-----------|
| Ubuntu              | 22.04.2         | 1.19 | 0.31.1    |

#### Go instalação
Instala nos dois diretórios client e server
```
wget https://golang.org/dl/go1.19.linux-amd64.tar.gz
```
```
tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz
```
```
export PATH=$PATH:/usr/local/go/bin
```

#### [ALTERNATIVA] Go instalação via terminal Linux
```
sudo apt install golang-go
```

#### Gerar os certs
##### Pré-requisito
- OpenSSL: Certifique-se de que o OpenSSL esteja instalado em seu sistema
	```
	./generate_cert.sh
	```
- Após o término do script, você terá os seguintes arquivos:

    - ca.key: Arquivo da chave da CA.
    - ca.pem: Arquivo do certificado da CA.
    - cert.csr: Arquivo da solicitação de assinatura de certificado.
    - private.key: Arquivo da chave privada.
    - cert.pem: Arquivo do certificado assinado.
#### Documentação
###### Classes
- A classe `main.go` serve para criar e iniciar um servidor HTTP/3 com o protocolo QUIC que lida com requisições em rotas específicas e registra as mensagens recebidas.
- A classe `client.go` serve para se comunicar com um servidor usando HTTP/3 e QUIC. Ele realiza solicitações GET e POST, exibe as respostas no log e permite ao usuário digitar mensagens para enviar ao servidor.

###### Métodos
- *Método `main` da classe `main.go`*
  1. Obtém o diretório de trabalho atual.
   ```go
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
   ```
  2. Cria um novo roteador Gin.
   ```go
	router := gin.New()
   ```
  3. Define uma rota para a raiz ("/") que retorna um JSON com a mensagem "Oi".
   ```go
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Oi": "OK"})
	})
   ```  
  4. Define uma rota para "/mensagem" que recebe uma requisição POST contendo um JSON com um campo "texto". A função associada a essa rota lê o JSON e registra a mensagem recebida.
   ```go
	router.POST("/mensagem", func(c *gin.Context) {
		var mensagem struct {
			Texto string `json:"texto"`
		}
		if err := c.ShouldBindJSON(&mensagem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println("Mensagem recebida do cliente:", mensagem.Texto)
		c.JSON(http.StatusOK, gin.H{"status": "Mensagem recebida"})
	})
   ```
  5. Cria um novo servidor HTTP/3 usando a estrutura http3.Server. Define o endereço do servidor como "127.0.0.1:8080" e o roteador Gin como o manipulador de solicitações.
   ```go
	server := http3.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}
    ```
  6. Inicia o servidor HTTP/3 para escutar na porta especificada (8080) e usar o certificado e chave fornecidos.
    ```go
	err = server.ListenAndServeTLS(path.Join(currentPath, "cert.pem"), path.Join(currentPath, "private.key"))
	if err != nil {
		log.Printf("Server error: %v", err)
	}
    ```
###### Subprogramas

Define uma rota para a raiz ("/") que retorna um JSON com a mensagem "Oi"
```go
router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Oi": "OK"})
	})
```
###### Variáveis

#### Executar a aplicação
```
cd server
```
```
go build main.go
```
```
./main
```

```
cd client
```
```
go build client.go
```
```
./client
```
