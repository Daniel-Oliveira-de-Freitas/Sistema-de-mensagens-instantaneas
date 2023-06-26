package main

import (
	"bufio"         // Ler a entrada do usuário.
	"bytes"         // Tabalhar com buffers de bytes.
	"crypto/tls"    //Configuração de TLS.
	"crypto/x509"   // Trabalhar com certificados X.509
	"encoding/json" // Utilizado para JSON marshaling e unmarshaling.
	"encoding/pem"  // Decodificar dados de certificados codificados por PEM.
	"fmt"           // imprimir mensagens
	"io"            // Operações de entrada/saída
	"log"           // Utilizado para registrar mensagens
	"net/http"      // Solicitações HTTP e manipular respostas
	"os"            // Acessar o ambiente do sistema operacional
	"path"          // caminhos de arquivo
	"time"          // Medir e manipular o tempo

	"github.com/lucas-clemente/quic-go/http3" // implementação de ida e volta HTTP/3.
)

// Obtém o diretório de trabalho atual
func main() {
	for {
		currentPath, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		// Cria um novo RoundTripper para comunicação HTTP/3
		roundTripper := &http3.RoundTripper{
			TLSClientConfig: &tls.Config{
				// Definir os certificados CA raiz para verificação TLS
				RootCAs: getRootCA(currentPath),
			},
		}
		// Fecha o RoundTripper ao final da função principal
		defer roundTripper.Close()

		// Cria um cliente HTTP com o RoundTripper customizado
		client := &http.Client{
			Transport: roundTripper,
		}

		// Definição das URLs de destino
		addr := "https://localhost:8080/"          // URL para solicitação GET
		addr2 := "https://localhost:8080/mensagem" // URL para solicitação POST

		// Envia uma requisição GET para o servidor
		rsp, err := client.Get(addr)
		if err != nil {
			panic(err)
		}
		defer rsp.Body.Close()

		// Lê o corpo da resposta em um buffer
		body := &bytes.Buffer{}
		_, err = io.Copy(body, rsp.Body)
		if err != nil {
			panic(err)
		}

		// Imprime o comprimento e o conteúdo do corpo da resposta
		log.Printf("Body length: %d bytes\n", body.Len())
		log.Printf("Response body: %s\n", body.Bytes())

		// Solicita ao usuário que digite um texto de mensagem
		var inputText string
		fmt.Print("Enter the message text: ")
		// Cria um novo leitor para ler a entrada da entrada padrão (teclado)
		reader := bufio.NewReader(os.Stdin)
		// Lê a entrada até que um caractere de nova linha ('\n') seja encontrado
		// A entrada é armazenada na variável 'inputText', e qualquer erro que ocorra é atribuído a 'err'
		inputText, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Remove o caractere de nova linha ('\n')
		inputText = inputText[:len(inputText)-1]

		if inputText == "sair" || inputText == "Sair" || inputText == "SAIR" {
			break // Encerra o laço de repetição se o usuário digitar "sair"
		}

		// Cria uma mensagem JSON com o texto digitado
		message := map[string]string{"texto": inputText}
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			log.Fatal(err)
		}

		// Envia uma requisição POST com a mensagem JSON
		start := time.Now()
		resp, err := client.Post(addr2, "application/json", bytes.NewBuffer(jsonMessage))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// Imprime o status da resposta e o tempo de ida e volta
		log.Printf("Response status: %s", resp.Status)
		log.Printf("RTT: %v\n", time.Since(start).Nanoseconds())
	}
}

// Função para carregar os certificados de CA raiz de um arquivo PEM
func getRootCA(certPath string) *x509.CertPool {
	caCertPath := path.Join(certPath, "ca.pem")
	caCertRaw, err := os.ReadFile(caCertPath)
	if err != nil {
		panic(err)
	}

	// Decodifica o certificado codificado por PEM
	p, _ := pem.Decode(caCertRaw)
	if p.Type != "CERTIFICATE" {
		panic("expected a certificate")
	}

	// ParseCertificate analisa o certificado do bloco PEM decodificado e retorna um objeto de certificado
	caCert, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		panic(err) // erro ao analisar o certificado, entre em pânico e interrompa a execução
	}

	// Cria um novo pool de certificados para manter os certificados de CA raiz
	certPool := x509.NewCertPool()
	// AddCert adiciona um certificado ao pool
	certPool.AddCert(caCert)

	// Retorna o pool de certificados contendo o certificado CA raiz
	return certPool
}
