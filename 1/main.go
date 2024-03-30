package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func main() {

	//Abrinco conexão mysql
	db, err := sql.Open("mysql", "root:172983456@/teste")
	if err != nil {
		panic(err)
	}

	//Garantindo fechamento da conexão
	defer db.Close()

	//Criando um produto e recebendo o ponteiro dele
	p := NewProduct("Cadeira Gamer", 20.50)

	//Inserindo o produto no banco
	err = insertProduct(db, p)
	if err != nil {
		panic(err)
	}

}

// Função que cria um produto recebendo os parâmetros
// e retornando o ponteiro do produto criado.
func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.NewString(),
		Name:  name,
		Price: price,
	}
}

// Função que insere o produto no banco de dados
func insertProduct(db *sql.DB, p *Product) error {

	//Preparando a criação do statement
	stmt, err := db.Prepare("insert into products(id,name,price) values (?,?,?)")
	if err != nil {
		return err
	}

	//Garantindo o fechamento do statement
	defer stmt.Close()

	//Executando o statement
	_, err = stmt.Exec(p.ID, p.Name, p.Price)
	if err != nil {
		return err
	}

	//Se der tudo certo, sem erros.
	return nil
}
