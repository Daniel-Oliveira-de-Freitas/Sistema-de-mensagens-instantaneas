package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/lucas-clemente/quic-go/http3"
)

func main() {
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			RootCAs: getRootCA(currentPath),
		},
	}
	defer roundTripper.Close()

	client := &http.Client{
		Transport: roundTripper,
	}

	addr := "https://localhost:8080/"
	addr2 := "https://localhost:8080/mensagem"
	rsp, err := client.Get(addr)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	body := &bytes.Buffer{}
	_, err = io.Copy(body, rsp.Body)
	if err != nil {
		panic(err)
	}

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

	message := map[string]string{"texto": inputText}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	resp, err := client.Post(addr2, "application/json", bytes.NewBuffer(jsonMessage))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Printf("Response status: %s", resp.Status)
	log.Printf("RTT: %v\n", time.Since(start).Nanoseconds())
}

func getRootCA(certPath string) *x509.CertPool {
	caCertPath := path.Join(certPath, "ca.pem")
	caCertRaw, err := os.ReadFile(caCertPath)
	if err != nil {
		panic(err)
	}

	p, _ := pem.Decode(caCertRaw)
	if p.Type != "CERTIFICATE" {
		panic("expected a certificate")
	}

	caCert, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(caCert)

	return certPool
}
