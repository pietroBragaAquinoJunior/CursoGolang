package main

import (
	"bufio"
	"fmt"
	"os"
)

// Manipulando Arquivos
func main() {

	//Criando Arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//Escrevendo no Arquivo
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo!"))
	if err != nil {
		panic(err)
	}

	//Log
	fmt.Printf("Arquivo criado com sucesso! Tamanho %d \n", tamanho)

	//Fechando o arquivo
	f.Close()

	// Lendo o arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//Log
	fmt.Println(string(arquivo))

	//Abrindo o arquivo
	arquivo2, err2 := os.Open("arquivo.txt")
	if err2 != nil {
		panic(err2)
	}

	//Leitura de pouco em pouco do arquivo
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	//Fechando o arquivo
	arquivo2.Close()

	//Apagando arquivo
	err3 := os.Remove("arquivo.txt")
	if err3 != nil {
		panic(err3)
	}

}
