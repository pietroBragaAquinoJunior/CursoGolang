package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Trabalhando com Timeout
func main() {

	//Definindo o timeout do http.client como 1 segundo
	c := http.Client{
		Timeout: time.Second,
	}

	//Fazendo requisição para o google.com
	resp, err := c.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}

	//Garantir fechamento do corpo da requisição
	defer resp.Body.Close()

	//Lendo corpo da requisição
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, string(body))

}
