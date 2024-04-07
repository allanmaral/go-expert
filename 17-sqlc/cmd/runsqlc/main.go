package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/allanmaral/go-expert/17-sqlc/internal/db"
	// "github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queries := db.New(conn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Sample description", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	//
	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "edd2dbe9-b819-43e9-a883-a837e719cab4",
		Name:        "Backend with Go",
		Description: sql.NullString{String: "Build a better backend!", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	category, err := queries.GetCategory(ctx, "edd2dbe9-b819-43e9-a883-a837e719cab4")
	if err != nil {
		panic(err)
	}

	fmt.Printf("ID: %s, Name: %s, Description: %s\n", category.ID, category.Name, category.Description.String)
}
