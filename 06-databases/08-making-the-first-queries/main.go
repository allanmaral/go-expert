package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Name  string
	ID    int `gorm:"primaryKey"`
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// select one by id
	var product Product
	db.First(&product, 4)
	fmt.Println(product)

	// select where name = Mouse
	var productWhereNameEqualsMouse Product
	db.First(&productWhereNameEqualsMouse, "name = ?", "Mouse")
	fmt.Println(productWhereNameEqualsMouse)

	// select all
	var producs []Product
	db.Find(&producs)
	fmt.Println("-- select all")
	fmt.Println("| ID\t | Name \t | Price \t|")
	for _, p := range producs {
		fmt.Printf("| %d\t | %s\t | %.2f \t|\n", p.ID, p.Name, p.Price)
	}
}
