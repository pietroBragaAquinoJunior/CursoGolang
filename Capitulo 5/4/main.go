package main

import (
	"context"
	"fmt"
)

// Context with value
func main() {

	//Criando o meu context
	ctx := context.Background()

	//Criando context com valor
	ctx = context.WithValue(ctx, "token", "senha")

	//Chamando a função passando o context
	bookHotel(ctx)

}

// Printando o valor token do context
func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
