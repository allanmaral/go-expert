package main

import (
	"fmt"
	"net/http"

	"github.com/allanmaral/go-expert/09-apis/configs"
	"github.com/allanmaral/go-expert/09-apis/internal/entity"
	"github.com/allanmaral/go-expert/09-apis/internal/infra/database"
	"github.com/allanmaral/go-expert/09-apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	fmt.Println(cfg)

	productRepository := database.NewGormProductRepository(db)
	productHandler := handlers.NewProductHandler(productRepository)

	userRepository := database.NewGormUserRepository(db)
	userHandler := handlers.NewUsersHandler(userRepository, cfg.TokenAuth, cfg.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Post("/products", productHandler.CreateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/users", userHandler.CreateUser)
	r.Post("/auth/login", userHandler.Login)

	http.ListenAndServe(":8000", r)
}
