// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/allanmaral/go-expert/19-dependency-injection/product"
	"github.com/google/wire"

)

// Injectors from wire.go:

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	productRepositorySQL := product.NewProductRepositorySQL(db)
	productUseCase := product.NewProductUseCase(productRepositorySQL)
	return productUseCase
}

// wire.go:

var setRepositoryDependency = wire.NewSet(product.NewProductRepositorySQL, wire.Bind(new(product.ProductRepository), new(*product.ProductRepositorySQL)))
