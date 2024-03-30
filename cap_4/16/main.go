package main

import (
	"html/template"
	"os"
)

// Struct Curso
type Curso struct {
	Nome         string
	CargaHoraria int
}

// Array de Curso
type Cursos []Curso

// Composição de templates
func main() {

	// array de strings com os nomes dos templates
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	// Criar o must pegando o template content que importa
	// o header e o footer.
	t := template.Must(
		template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 70},
	})
	if err != nil {
		panic(err)
	}
}
