package main

import (
	"database/sql"
	"fmt"
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
	p := newProduct("Cadeira Gamer", 20.50)

	//Inserindo o produto no banco
	err = insertProduct(db, p)
	if err != nil {
		panic(err)
	}

	//Fazendo update do preço do produto
	p.Price = 300
	err = updateProduct(db, p)
	if err != nil {
		panic(err)
	}

	//Testando select de um produto
	p2, err := selectOneProduct(db, p.ID)
	fmt.Printf("Product: %v, possui o preço de %.2f", p2.Name, p2.Price)

	//Testando select de todos os produtos
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f \n", p.Name, p.Price)
	}

}

// Função que cria um produto recebendo os parâmetros
// e retornando o ponteiro do produto criado.
func newProduct(name string, price float64) *Product {
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

func updateProduct(db *sql.DB, p *Product) error {

	//Preparando a criação do statement
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}

	//Garantindo o fechamento do statement
	defer stmt.Close()

	//Executando o statement
	_, err = stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		return err
	}

	//Se der tudo certo, sem erros.
	return nil

}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {

	//Preparando a criação do statement
	stmt, err := db.Prepare("select id,name,price from products where id = ?")
	if err != nil {
		return nil, err
	}

	//Garantindo o fechamento do statement
	defer stmt.Close()

	//Criando produto e colocando nele o resultado da linha
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	//Retornando o produto sem erros
	return &p, nil

}

func selectAllProducts(db *sql.DB) ([]Product, error) {

	//Executando a query sem statement, pois não possui argumento
	rows, err := db.Query("select id,name,price from products")
	if err != nil {
		return nil, err
	}

	//Garantindo o fechamento
	defer rows.Close()

	// Percorrendo as linhas e inserindo no array de produtos
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	//Retornando produtos sem erro
	return products, nil

}
