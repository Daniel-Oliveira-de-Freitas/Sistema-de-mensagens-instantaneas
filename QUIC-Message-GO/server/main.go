package main

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/lucas-clemente/quic-go/http3"
)

// Definicação da função main
func main() {
	// Obtém o diretório de trabalho atual
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Cria um novo roteador Gin
	router := gin.New()

	// Define uma rota para a raiz ("/") que retorna um JSON com a mensagem "Oi".
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Oi": "OK"})
	})

	// Define uma rota para "/mensagem" que recebe uma requisição POST contendo um JSON
	// com um campo "texto". A função associdada a essa rota lê o JSON e registra a
	// mensagem recebida.
	router.POST("/mensagem", func(c *gin.Context) {
		// define um campo Texto do tipo string serializar ou desserializar a estrutura para JSON
		var mensagem struct {
			Texto string `json:"texto"`
		}
		// c é um objeto do contexto do GIN, ShouldBindJSON à variável err. Se ocorrer um erro durante o bind, um JSON de resposta com um status 400
		if err := c.ShouldBindJSON(&mensagem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// O valor do campo Texto da estrutura mensagem é impresso no log.
		log.Println("Mensagem recebida do cliente:", mensagem.Texto)
		c.JSON(http.StatusOK, gin.H{"status": "Mensagem recebida"})
	})

	// Cria um novo servidor HTTP/3 usando a estrutura http3.Server. Define o endereço do servidor como "127.0.0.1:8080" e o roteador Gin como o manipulador de solicitações.
	server := http3.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	// go func() {
	// 	err = server.ListenAndServeTLS(path.Join(currentPath, "cert.pem"), path.Join(currentPath, "private.key"))
	// 	if err != nil {
	// 		log.Printf("Server error: %v", err)
	// 	}
	// }()

	// select {}

	// Inicia o servidor HTTP/3 para escutar na porta especificada (8080) e usar o certificado e chave fornecidos.
	// o caminho do certificado TLS/SSL e o caminho da chave privada correspondente ao certificado
	err = server.ListenAndServeTLS(path.Join(currentPath, "cert.pem"), path.Join(currentPath, "private.key"))
	if err != nil {
		// Registra qualquer erro ocorrido durante a execução do servidor.
		log.Printf("Server error: %v", err)
	}
}
