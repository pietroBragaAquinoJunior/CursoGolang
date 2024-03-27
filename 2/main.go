package main

import (
	"io"
	"net/http"
)

// Requisições Http
func main() {

	//Chamada http para o google
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	//Garantir fechamento do corpo da requisição
	defer req.Body.Close()

	// Leitura do io.Reader (req.Body)
	res, err2 := io.ReadAll(req.Body)
	if err2 != nil {
		panic(err2)
	}

	//Log
	println(string(res))

}
