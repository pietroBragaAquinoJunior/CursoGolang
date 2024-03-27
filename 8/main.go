package main

import (
	json2 "encoding/json"
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

// Servidor Http
func main() {

	//Redireciona para função correspondente
	http.HandleFunc("/", BuscaCepHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro ao subir o servidor.")
	}

}

// Função chamada ao fazer get na rota padrão
func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	//Qualquer coisa além do / vai para not found
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Pegando parâmetro cep da url get, se não vier bad request.
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Coloca no Header retorno tipo json
	w.Header().Set("Content-Type", "application/json")

	//Coloca no header status 200
	w.WriteHeader(http.StatusOK)

	//Retorna mensagem simples
	_, _ = w.Write([]byte("Hello World!"))

}

func BuscaCep(cep string) (*ViaCep, error) {

	//Requisição com o cep para o viacep.com.br
	res, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}

	//Garantir fechamento do body
	defer res.Body.Close()

	//Pegar informações do body
	json, error := io.ReadAll(res.Body)
	if error != nil {
		return nil, error
	}

	//Transformar json em struct ViaCep
	var viacep ViaCep
	error = json2.Unmarshal(json, &viacep)
	if error != nil {
		return nil, error
	}

	//retornar struct ViaCep
	return &viacep, nil

}
