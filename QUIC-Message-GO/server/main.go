package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-clemente/quic-go/http3"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Oi": "OK"})
	})

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

	err = server.ListenAndServeTLS(path.Join(currentPath, "cert.pem"), path.Join(currentPath, "private.key"))
	if err != nil {
		log.Printf("Server error: %v", err)
	}
}
