package main

import (
	"html/template"
	"os"
	"strings"
)

// Struct Curso
type Curso struct {
	Nome         string
	CargaHoraria int
}

// Array de Curso
type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Usar Função no template
func main() {

	// array de strings com os nomes dos templates
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	// Criando e colocando mapa de funções no template
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})

	// Criar o must pegando o template content que possui
	// o header e o footer.
	t = template.Must(t.ParseFiles(templates...))

	// Executar o must e redirecionar o resultado para
	// o console e passar o array de Curso como variável
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 70},
	})
	if err != nil {
		panic(err)
	}

}
