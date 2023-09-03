package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	Id           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	dsn := "root:root@tcp(localhost:3306)/goexpert?charset-utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// // create category
	// db.Create(&Category{
	// 	Name: "Electronics",
	// })
	//
	// // create product
	// db.Create(&Product{
	// 	Name:       "Notebook",
	// 	Price:      2000.00,
	// 	CategoryID: 1,
	// })
	//
	// // create serial nuber
	// db.Create(&SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: 1,
	// })

	var categories []Category
	err = db.
		Model(&Category{}).
		Preload("Products").
		Preload("Products.SerialNumber"). // when loading nested properties, need to use the full path
		Find(&categories).
		Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Printf("Category %s:\n", category.Name)
		for _, product := range category.Products {
			fmt.Printf("  - %s (%s)\n", product.Name, product.SerialNumber.Number)
		}
	}
}
