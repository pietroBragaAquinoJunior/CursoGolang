package main

import (
	"fmt"
	"net/http"
	"os"
)

// Servidor Http
func main() {

	//Redireciona para função correspondente
	http.HandleFunc("/", BuscaCep)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro ao subir o servidor.")
	}

}

// Função chamada ao fazer get na rota padrão
func BuscaCep(w http.ResponseWriter, r *http.Request) {

	//Retorna mensagem simples
	_, _ = w.Write([]byte("Hello World!"))

}
