package main

import (
	"fmt"
	"net/http"

	"github.com/allanmaral/go-expert/09-apis/configs"
	"github.com/allanmaral/go-expert/09-apis/internal/entity"
	"github.com/allanmaral/go-expert/09-apis/internal/infra/database"
	"github.com/allanmaral/go-expert/09-apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	fmt.Println(config)

	productRepository := database.NewGormProductRepository(db)
	productHandler := handlers.NewProductHandler(productRepository)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
