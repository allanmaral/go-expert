package main

import (
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

	// ctx := db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.00,
	// })
	products := []Product{
		{Name: "Notebook", Price: 1000.0},
		{Name: "Mouse", Price: 50.0},
		{Name: "Keyboard", Price: 100.0},
	}
	res := db.Create(&products)
	if res.Error != nil {
		panic(res.Error)
	}
}
