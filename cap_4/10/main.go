package main

import "net/http"

// Servidor Mux
func main() {

	//Criando o multiplexer
	mux := http.NewServeMux()

	//Função que resolve a rota padrão
	mux.HandleFunc("/", firstHandler)

	// espera o tipo Handler, a interface Handler diz que espera
	//uma struct que tenha o método ServeHTTP implementado.
	mux.Handle("/home", homeStruct{title: "Olá"})

	//Subindo servidor
	http.ListenAndServe(":8080", mux)

}

// Função que desempenha papel de Handler
func firstHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá"))
}

// Criei uma estrutura que vai virar um Handler pois para uma estruct
// virar Handler ele só precisa implementar o método ServeHTTP.
type homeStruct struct {
	title string
}

// Função que a struct homeStruct precisa implementar para virar um http.Handler
func (h homeStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.title))
}
