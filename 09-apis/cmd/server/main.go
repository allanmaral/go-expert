package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/allanmaral/go-expert/09-apis/configs"
	_ "github.com/allanmaral/go-expert/09-apis/docs"
	"github.com/allanmaral/go-expert/09-apis/internal/entity"
	"github.com/allanmaral/go-expert/09-apis/internal/infra/database"
	"github.com/allanmaral/go-expert/09-apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title			Go Expert API Example
// @version			1.0
// @description		Product API with authentication
// @termsOfService	http://swagger.io/terms/
//
// @contact.name	Allan Ribeiro
// @contact.url		https://github.com/allanmaral
// @contact.email	allanmaralr@gmail.com
//
// @license.name	MIT
// @license.url		https://github.com/allanmaral/go-expert/LICENCE.txt
//
// @host			localhost:8000
// @basePath		/
// @securityDefinitions.apiKey	ApiKeyAuth
// @in				header
// @name			Authorization
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
	r.Use(middleware.Recoverer)
	r.Use(LogRequest)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/", productHandler.GetProducts)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/auth/login", userHandler.Login)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
