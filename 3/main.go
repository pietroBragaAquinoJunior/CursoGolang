package main

import "fmt"

// Defer
func main() {

	defer fmt.Println("Primeira Linha")
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")

}
