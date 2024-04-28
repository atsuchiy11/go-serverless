package domain

import "time"

type ProductModel struct {
	ID          uint64
	Name        string
	Price       int
	ReleaseDate time.Time
}

func NewProductModel(id uint64, name string, price int, releaseDate time.Time) *ProductModel {
	return &ProductModel{
		ID:          id,
		Name:        name,
		Price:       price,
		ReleaseDate: releaseDate,
	}
}
