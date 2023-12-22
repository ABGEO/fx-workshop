package repository

import (
	"github.com/abgeo/fx-workshop/model"
	"gorm.io/gorm"
)

type IProductRepository interface {
	Create(product *model.Product) error
	FindAll() ([]model.Product, error)
	FindByID(id uint) (*model.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) (*ProductRepository, error) {
	return &ProductRepository{
		db: db,
	}, nil
}

func (repo *ProductRepository) Create(product *model.Product) error {
	return repo.
		db.
		Create(product).
		Error
}

func (repo *ProductRepository) FindAll() ([]model.Product, error) {
	var products []model.Product

	return products, repo.
		db.
		Find(&products).
		Error
}

func (repo *ProductRepository) FindByID(id uint) (*model.Product, error) {
	var product *model.Product

	return product, repo.
		db.
		First(&product, id).
		Error
}
