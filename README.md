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
    
<hr>

- *Método `main` da classe `client.go`*
  1. Obtém o diretório de trabalho atual.
   ```go
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
   ```
  2. Cria um novo RoundTripper para comunicação HTTP/3
   ```go
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			RootCAs: getRootCA(currentPath), // Definir os certificados CA raiz para verificação TLS
		},
	}
	defer roundTripper.Close() // Fecha o RoundTripper ao final da função principal
   ```
  3. Cria um cliente HTTP com o RoundTripper customizado
   ```go
	client := &http.Client{
		Transport: roundTripper,
	}
   ```
  4. Definição das URLs de destino
   ```go
	addr := "https://localhost:8080/"          // URL para solicitação GET
	addr2 := "https://localhost:8080/mensagem" // URL para solicitação POST
   ```
  5. Envia uma requisição GET para o servidor
   ```go
	rsp, err := client.Get(addr)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
   ```
  6. Lê o corpo da resposta em um buffer
   ```go
		body := &bytes.Buffer{}
		_, err = io.Copy(body, rsp.Body)
		if err != nil {
			panic(err)
		}
   ```
  7. Imprime o comprimento e o conteúdo do corpo da resposta
   ```go
		log.Printf("Body length: %d bytes\n", body.Len())
		log.Printf("Response body: %s\n", body.Bytes())
   ```
  8. Solicita ao usuário que digite um texto de mensagem
   ```go
	 	var inputText string
		fmt.Print("Enter the message text: ")
		reader := bufio.NewReader(os.Stdin) // Cria um novo leitor para ler a entrada da entrada padrão (teclado)
		inputText, err = reader.ReadString('\n') // Lê a entrada até que um caractere de nova linha ('\n') seja encontrado
		if err != nil {
			log.Fatal(err)
		}
   ```
  9. Remove o caractere de nova linha ('\n')
  ```go
	 	inputText = inputText[:len(inputText)-1]

		if inputText == "sair" || inputText == "Sair" || inputText == "SAIR" {
			break // Encerra o laço de repetição se o usuário digitar "sair"
		}
  ```
  10. Cria uma mensagem JSON com o texto digitado
    ```go
	message := map[string]string{"texto": inputText}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}
    ```
  11. Envia uma requisição POST com a mensagem JSON
    ```go
		start := time.Now()
		resp, err := client.Post(addr2, "application/json", bytes.NewBuffer(jsonMessage))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
    ```
  12. Imprime o status da resposta e o tempo de ida e volta
    ```go
		log.Printf("Response status: %s", resp.Status)
		log.Printf("RTT: %v\n", time.Since(start).Nanoseconds())
    ```
- *Método `getRootCA(certPath string) *x509.CertPool` da classe `client.go`*
  13. Lê o arquivo PEM contendo o certificado de CA raiz
    ```go
	caCertPath := path.Join(certPath, "ca.pem")
	caCertRaw, err := os.ReadFile(caCertPath)
	if err != nil {
		panic(err)
	}
    ```
  14. Decodifica o certificado codificado por PEM
    ```go
	p, _ := pem.Decode(caCertRaw)
	if p.Type != "CERTIFICATE" {
		panic("expected a certificate")
	}
    ```
  15. ParseCertificate analisa o certificado do bloco PEM decodificado e retorna um objeto de certificado
    ```go
	caCert, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		panic(err)
	}
    ```
  16. Cria um novo pool de certificados para manter os certificados de CA raiz
    ```go
	certPool := x509.NewCertPool()
	certPool.AddCert(caCert)
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
