package main

import "net/http"

// File server
func main() {

	//Criando o File Server Handle
	fileServer := http.FileServer(http.Dir("./public"))

	//Criando o Mux Handler
	mux := http.NewServeMux()

	//Colocando o fileHandler para resolver a rota base no MuxHandler
	mux.Handle("/", fileServer)
	mux.HandleFunc("/teste", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Testando!"))
	})

	//Ativando servidor com o muxHandler no topo
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}

}
