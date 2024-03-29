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

// Html Dinâmico com template Must
func main() {

	//Criando o curso
	curso := Curso{"Go", 40}

	//Criando o template
	tmp := template.New("CursoTemplate")

	//Html do template usando Parse
	tmp, err := tmp.Parse("Curso: {{.Nome}} - Carga Horária {{.CargaHoraria}}")
	if err != nil {
		panic(err)
	}

	//Criando template Must.
	t := template.Must(tmp, err)

	//Executando template passando a variável
	err = t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
