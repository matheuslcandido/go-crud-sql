package main

import (
	"bytes"
	"fmt"
	"go-curd-sql/database"
	"go-curd-sql/entity"
	"go-curd-sql/repositories"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "go-crud-sql: ", log.Lshortfile)
	)
	defer fmt.Print(&buf)

	logger.Print("Application started!")

	db, err := database.Connect(logger)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	produto := entity.NewProduct("insert", 100)
	err = repositories.InsertProdutc(logger, db, *produto)
	if err != nil {
		panic(err)
	}

	produto.Name = "updated"
	produto.Amount = 200

	err = repositories.UpdateProduct(logger, db, *produto)
	if err != nil {
		panic(err)
	}

	selectedProduct, err := repositories.SelectOneProduct(logger, db, produto.Id)
	if err != nil {
		panic(err)
	}

	produto2 := entity.NewProduct("insert", 100)
	err = repositories.InsertProdutc(logger, db, *produto2)
	if err != nil {
		panic(err)
	}
	logger.Printf("Product data from selectOne: %v", selectedProduct)

	selectedProducts, err := repositories.SelectAllProducts(logger, db)
	if err != nil {
		panic(err)
	}

	for _, selectedProduct := range selectedProducts {
		logger.Printf("Product data from selectAll: %v", selectedProduct)
	}

	logger.Print("Application finished!")
}
