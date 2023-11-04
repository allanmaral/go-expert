package database

import (
	"fmt"
	"testing"

	"github.com/allanmaral/go-expert/09-apis/internal/entity"
	generators "github.com/allanmaral/go-expert/09-apis/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func makeProductRepository(t *testing.T) (ProductRepository, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	userRepository := NewGormProductRepository(db)

	return userRepository, db
}

func makeProduct() *entity.Product {
	name := generators.RandonName()
	price := generators.RandonPrice()
	product, _ := entity.NewProduct(name, price)

	return product
}

func TestCreateProduct(t *testing.T) {
	product := makeProduct()
	sut, db := makeProductRepository(t)

	err := sut.Create(product)

	assert.Nil(t, err)
	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindAllProducts(t *testing.T) {
	sut, db := makeProductRepository(t)

	for i := 1; i <= 25; i++ {
		product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i), generators.RandonPrice())
		err := db.Create(product).Error
		assert.NoError(t, err)
	}

	products, err := sut.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = sut.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = sut.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 25", products[4].Name)

	products, err = sut.FindAll(0, 0, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 25)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 25", products[24].Name)
}

func TestFindProductById(t *testing.T) {
	sut, db := makeProductRepository(t)
	product, _ := entity.NewProduct("A specific product", generators.RandonPrice())
	db.Create(product)

	result, err := sut.FindByID(product.ID.String())

	assert.NoError(t, err)
	assert.Equal(t, product.Name, result.Name)
	assert.Equal(t, product.Price, result.Price)
}

func TestUpdateProduct(t *testing.T) {
	sut, db := makeProductRepository(t)
	product, _ := entity.NewProduct("A specific product", generators.RandonPrice())
	db.Create(product)

	product.Name = "A new product name"
	err := sut.Update(product)

	var result entity.Product
	db.First(&result, "id = ?", product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "A new product name", result.Name)
	assert.Equal(t, product.Price, result.Price)
}

func TestDeleteProduct(t *testing.T) {
	sut, db := makeProductRepository(t)
	product := makeProduct()
	db.Create(product)

	err := sut.Delete(product.ID.String())

	assert.NoError(t, err)
	var result entity.Product
	err = db.First(&result, "id = ?", product.ID.String()).Error
	assert.Error(t, err)
}
