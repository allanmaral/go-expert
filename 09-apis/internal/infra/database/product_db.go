package database

import (
	"github.com/allanmaral/go-expert/09-apis/internal/entity"
	"gorm.io/gorm"
)

type GormProductRepository struct {
	DB *gorm.DB
}

var _ ProductRepository = (*GormProductRepository)(nil)

func NewGormProductRepository(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{DB: db}
}

func (p *GormProductRepository) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *GormProductRepository) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	clause := p.DB
	if page != 0 && limit != 0 {
		clause = clause.Limit(limit).Offset((page - 1) * limit)
	}

	err := clause.Order("created_at " + sort).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, err
}

func (p *GormProductRepository) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *GormProductRepository) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *GormProductRepository) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
