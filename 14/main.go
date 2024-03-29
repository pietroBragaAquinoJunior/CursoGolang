package main

import (
	"html/template"
	"os"
)

// Struct para carregar informações do curso
type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

// Html Dinâmico com template Must
func main() {

	//Criando os cursos
	cursos := Cursos{
		{"Go", 40},
		{"Java", 50},
	}

	//Criando o template
	tmp := template.New("template.html")

	//Html do template usando Parse
	tmp, err := tmp.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}

	//Criando template Must.
	t := template.Must(tmp, err)

	//Executando template passando a variável
	err = t.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}

}
