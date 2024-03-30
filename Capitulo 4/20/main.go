package main

import (
	"io"
	"net/http"
)

// http post
func main() {

	//Criando o client para realizar requisição
	client := http.Client{}

	//Criando a request mas sem executar ainda.
	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	//Colocando o header json na requisição
	req.Header.Set("Accept", "application/json")

	//Dizendo para o client realizar a requisição
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	//Garantindo fechamento do corpo da resposta
	defer resp.Body.Close()

	//Lendo as informações do corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//Monstrando a resposta em formato string no console
	println(string(body))

}
