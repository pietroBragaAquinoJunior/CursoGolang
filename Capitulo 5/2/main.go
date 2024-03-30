package main

import (
	"log"
	"net/http"
	"time"
)

// Context
func main() {

	//Criando uma resposta para o get na rota padrão
	http.HandleFunc("/", handler)

	// Subindo servidor
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

// Função que resolve a rota padrão
func handler(w http.ResponseWriter, r *http.Request) {

	//Pegando o contexto da request
	ctx := r.Context()

	//log
	log.Println("Request iniciada.")
	defer log.Println("Request finalizada.")

	//Select para esperar as respostas
	select {
	// 5 segundos é o tempo que leva para completar.
	case <-time.After(time.Second * 5):
		log.Println("Request processada com sucesso")
		w.Write([]byte("Request processada com sucesso"))
	//Caso o cliente cancele, resposta conflict.
	case <-ctx.Done():
		log.Println("Request cancelada pelo client.")
	}

}
