package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset-utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.0,
	// })

	var p Product
	db.First(&p, 3)
	p.Name = "New notebook name"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 4)
	db.Delete(&p2)
}
