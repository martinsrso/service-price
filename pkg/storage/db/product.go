package db

import (
	"github.com/martinsrso/service-price/pkg/models"
	"github.com/martinsrso/service-price/pkg/storage"

	"github.com/jmoiron/sqlx"
)

type defaultProductRepository struct {
	db sqlx.Ext
}

func NewProductRepository(db sqlx.Ext) storage.ProductRepository {
	return defaultProductRepository{
		db: db,
	}
}

func (repo defaultProductRepository) Create(product models.Product) error {
	return nil
}

func (repo defaultProductRepository) FindByID(id string) (*models.Product, error) {
	return nil, nil
}

func (repo defaultProductRepository) Update(product models.Product) error {
	return nil
}
