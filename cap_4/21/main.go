package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

// Context
func main() {

	//Criando o contexto
	ctx := context.Background()

	//Dizendo que esse contexto possui timeout
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	//Garantindo que o context vai ser fechado de qualquer jeito
	defer cancel()

	//Criando a requisição
	req, err := http.NewRequestWithContext(ctx, "GET", "http://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	//Fazendo a requisição
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	//Garantindo fechamento do body
	defer res.Body.Close()

	//Lendo e mostrando no console o body
	body, err := io.ReadAll(res.Body)
	println(string(body))

}
