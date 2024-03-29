package main

import (
	"html/template"
	"net/http"
)

// Struct para carregar informações do curso
type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

// Usando Mux pro servidor e Must template para a página.
func main() {

	//Criando Mux
	mux := http.NewServeMux()

	//Subindo Servidor http
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {

			//Criando os cursos
			cursos := Cursos{
				{"Go", 40},
				{"Java", 50},
			}

			//Criando o template
			tmp := template.New("content.html")

			//Html do template usando Parse
			tmp, err := tmp.ParseFiles("content.html")
			if err != nil {
				panic(err)
			}

			//Criando template Must.
			t := template.Must(tmp, err)

			//Executando o template com a saída writer do mux
			//e passando os cursos.
			err = t.Execute(w, cursos)
			if err != nil {
				panic(err)
			}

		})

	// Subindo o servidor
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}
