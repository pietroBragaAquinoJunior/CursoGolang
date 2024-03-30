package main

import (
	"context"
	"fmt"
	"time"
)

// Context
func main() {

	//Iniciando o contexto
	ctx := context.Background()

	//Colocando regra de cancelamento no context
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)

	//Boa prática cancelar o contexto de qualquer jeito.
	defer cancel()

	//Chamando a função
	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	// O select com case esperam respostas.
	select {
	// caso ele receba resposta do ctx.Done(), foi cancelado.
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	// caso ele receba a resposta do processamento
	// com 5 segundos vai cancelar primeiro...
	// ele printa que deu tudo certo.
	case <-time.After(time.Second * 5):
		fmt.Println("Hotel booked.")
	}
}
