package main

import (
	"fmt"
	"net/http"
	"os"
)

// Servidor Http
func main() {

	//Redireciona para função correspondente
	http.HandleFunc("/", BuscaCepHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro ao subir o servidor.")
	}

}

// Função chamada ao fazer get na rota padrão
func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	//Qualquer coisa além do / vai para not found
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Pegando parâmetro cep da url get, se não vier bad request.
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Coloca no Header retorno tipo json
	w.Header().Set("Content-Type", "application/json")

	//Coloca no header status 200
	w.WriteHeader(http.StatusOK)

	//Retorna mensagem simples
	_, _ = w.Write([]byte("Hello World!"))

}
