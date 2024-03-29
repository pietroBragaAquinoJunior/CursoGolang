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

// Html Dinâmico
func main() {

	//Criando o curso
	curso := Curso{"Go", 40}

	//Criando o template
	tmp := template.New("CursoTemplate")

	//Html do template usando Parse
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária {{.CargaHoraria}}")

	//Executando template passando a variável
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
