package repository

import (
	"fmt"

	"github.com/abgeo/fx-workshop/config"
	"github.com/abgeo/fx-workshop/database"
	"github.com/abgeo/fx-workshop/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository() (*ProductRepository, error) {
	conf, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize config: %w", err)
	}

	db, err := database.New(conf.Database.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

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
