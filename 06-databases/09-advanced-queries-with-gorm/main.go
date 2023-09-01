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

	// select with limit and offset
	var products []Product
	db.Limit(2).Offset(2).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

	fmt.Println("--------------")

	db.Where("price < ?", 300).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

	fmt.Println("--------------")

	db.Where("name LIKE ?", "%book%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
