package product

import "database/sql"

type ProductRepository interface {
	GetProduct(int) (Product, error)
}

type ProductRepositorySQL struct {
	db *sql.DB
}

func NewProductRepositorySQL(db *sql.DB) *ProductRepositorySQL {
	return &ProductRepositorySQL{db}
}

func (r *ProductRepositorySQL) GetProduct(id int) (Product, error) {
	return Product{
		ID:   id,
		Name: "Product Name",
	}, nil
}

type ProductRepositoryTXT struct {
}

func NewProductRepositoryTXT(db *sql.DB) *ProductRepositoryTXT {
	return &ProductRepositoryTXT{}
}

func (r *ProductRepositoryTXT) GetProduct(id int) (Product, error) {
	return Product{
		ID:   id,
		Name: "Product Name",
	}, nil
}
