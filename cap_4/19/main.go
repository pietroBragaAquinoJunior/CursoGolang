package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

// http post
func main() {

	//Criando o client para realizar requisição
	client := http.Client{}

	//Preparar o conteúdo do body para ser enviado no post
	jsonVar := bytes.NewBuffer([]byte(`{"name":"teste"}`))

	//Fazendo a requisição post.
	res, err := client.Post("http://www.google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}

	//Garantir fechamento do corpo da requisição
	defer res.Body.Close()

	//Mostrar saída no console
	io.CopyBuffer(os.Stdout, res.Body, nil)

}
