package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {

	// Conexão com Mysql usando Gorm
	dsn := "root:172983456@tcp(127.0.0.1:3306)/teste?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//Criando Tabela Product no mysql
	err = db.AutoMigrate(&Product{})
	if err != nil {
		return
	}

	//Adicionando linha na tabela Product
	//db.Create(&Product{Name: "Notebook", Price: 1000.00})

	//Criando vários produtos
	//products := []Product{
	//	{Name: "Notebook", Price: 1000.00},
	//	{Name: "Mouse", Price: 50.00},
	//	{Name: "Keyboard", Price: 100.00},
	//}
	//db.Create(products)

	//Selecionando um produto com condição
	//var product Product
	//db.First(&product, 3)
	//fmt.Println(product)
	//db.First(&product, "name = ?", "Mouse")
	//fmt.Println(product)

	//Selecionando vários produtos
	//var products []Product
	//db.Find(&products)
	//for _, p := range products {
	//	fmt.Println(p)
	//}
}
