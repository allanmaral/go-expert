//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/allanmaral/go-expert/19-dependency-injection/product"
	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepositorySQL,
	wire.Bind(new(product.ProductRepository), new(*product.ProductRepositorySQL)),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
