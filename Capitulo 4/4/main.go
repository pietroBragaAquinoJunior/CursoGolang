package main

import (
	"encoding/json"
	"os"
)

// Definição de Conta
type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

// Manipulação de Json
func main() {

	//Criando conta
	conta := Conta{Numero: 1, Saldo: 100}

	//Transformando em Json
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	//Log
	println(string(res))

	// Criando Encoder para cuspir resposta no console
	err2 := json.NewEncoder(os.Stdout).Encode(conta)
	if err2 != nil {
		return
	}

	//Criando Json Puro
	jsonPuro := []byte(`{"n":2,"s":200}`)

	//Criando nova Conta
	var conta2 Conta

	//Jogando dados do jsonPuro na conta2
	err3 := json.Unmarshal(jsonPuro, &conta2)
	if err3 != nil {
		return
	}

	//Log
	println(conta2.Saldo)

}
