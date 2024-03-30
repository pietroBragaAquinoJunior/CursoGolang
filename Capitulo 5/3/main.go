package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

// Cancelar o contexto do lado do client
func main() {

	// Criando context com tempo de 5s
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Second*10)

	//Garantindo fechamento do context
	defer cancel()

	//Criando requisição com contexto
	req, err := http.NewRequestWithContext(
		ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	//Realizando a request com client padrão
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	//Garantir fechamento do corpo da requisição
	defer res.Body.Close()

	//Pegando o corpo da requisição e jogando no console
	io.Copy(os.Stdout, res.Body)

}
