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
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	Id         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
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
	db.AutoMigrate(&Product{}, &Category{})

	// // create category
	// electronicsCategory := Category{
	// 	Name: "Electronics",
	// }
	// db.Create(&electronicsCategory)
	//
	// kitchenCategory := Category{
	// 	Name: "Kitchen",
	// }
	// db.Create(&kitchenCategory)
	//
	// // create product
	// db.Create(&Product{
	// 	Name:       "Notebook",
	// 	Price:      2000.00,
	// 	Categories: []Category{electronicsCategory, kitchenCategory},
	// })

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Printf("Category %s:\n", category.Name)
		for _, product := range category.Products {
			fmt.Printf("  - %s\n", product.Name)
		}
	}
}
