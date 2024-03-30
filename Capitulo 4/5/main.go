package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Struct Json que vem como resposta do site https://viacep.com.br
type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// Requisições Http para consumir Api e salvar informações em arquivo.
func main() {

	//Criar arquivo
	file, err := os.Create("cidade.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro ao criar arquivo cidade.txt")
	}

	//Para cada argumento do comando go run...
	for _, cep := range os.Args[1:] {

		//Fazer requisição com o argumento
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}

		//Garantir fechamento do corpo da requisição
		defer req.Body.Close()

		//Lendo corpo da requisição
		res, err := io.ReadAll(req.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		//Criar Objeto ViaCep
		var data ViaCep

		//Salvar json em data
		err = json.Unmarshal(res, &data)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Erro ao fazer a transformação json em struct: %v\n", err)
		}

		//Escrever no arquivo
		_, err = file.WriteString(fmt.Sprintf("Cep: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))

	}

	//Fechamento do arquivo
	file.Close()

	//Apagando arquivo
	err = os.Remove("cidade.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro ao apagar o arquivo.")
	}

}
