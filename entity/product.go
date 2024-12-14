package entity

import "github.com/google/uuid"

type Product struct {
	Id     string
	Name   string
	Amount int
}

func NewProduct(name string, amount int) *Product {
	return &Product{
		Id:     uuid.New().String(),
		Name:   name,
		Amount: amount,
	}
}
