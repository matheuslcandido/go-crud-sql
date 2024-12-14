package repositories

import (
	"database/sql"
	"go-curd-sql/entity"
	"log"
)

func InsertProdutc(logger *log.Logger, db *sql.DB, product entity.Product) error {
	logger.Printf("Starting insert product: %v on DB", product.Id)

	statement, err := db.Prepare("INSERT INTO produtos(id, name, amount) VALUES($1, $2, $3)")
	if err != nil {
		logger.Print(err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(product.Id, product.Name, product.Amount)
	if err != nil {
		logger.Print(err)
		return err
	}

	logger.Printf("Product: %v inserted on DB with succes", product.Id)

	return nil
}

func UpdateProduct(logger *log.Logger, db *sql.DB, product entity.Product) error {
	logger.Printf("Starting update product: %v on DB", product.Id)

	statement, err := db.Prepare("UPDATE produtos SET name=$2, amount=$3 WHERE id = $1")
	if err != nil {
		logger.Print(err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(product.Id, product.Name, product.Amount)
	if err != nil {
		logger.Print(err)
		return err
	}

	logger.Printf("Product: %v updated with succes", product.Id)

	return nil
}

func DeleteProduct(logger *log.Logger, db *sql.DB, id string) error {
	logger.Printf("Starting delete product: %v from DB", id)

	statement, err := db.Prepare("DELETE FROM produtos WHERE id = $1")
	if err != nil {
		logger.Print(err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		logger.Print(err)
		return err
	}

	logger.Printf("Product: %v deleted from DB", id)

	return nil
}

func SelectOneProduct(logger *log.Logger, db *sql.DB, id string) (*entity.Product, error) {
	logger.Printf("Starting select product: %v on DB", id)

	var selectedProduct entity.Product

	row := db.QueryRow("SELECT id, name, amount from produtos where id = $1", id)
	err := row.Scan(&selectedProduct.Id, &selectedProduct.Name, &selectedProduct.Amount)
	if err != nil {
		logger.Print(err)
		return nil, err
	}

	logger.Printf("Product: %v selected on DB", id)

	return &selectedProduct, nil
}

func SelectAllProducts(logger *log.Logger, db *sql.DB) ([]entity.Product, error) {
	logger.Print("Starting select all products from DB")

	var selectedProducts []entity.Product
	rows, err := db.Query("SELECT id, name, amount FROM produtos")
	if err != nil {
		logger.Print(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product

		rows.Scan(&product.Id, &product.Name, &product.Amount)

		selectedProducts = append(selectedProducts, product)
	}

	return selectedProducts, nil
}
