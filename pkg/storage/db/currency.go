package db

import (
	"github.com/martinsrso/service-price/pkg/models"
	"github.com/martinsrso/service-price/pkg/storage"

	"github.com/jmoiron/sqlx"
)

type defaultCurrencyRepository struct {
	db sqlx.Ext
}

func NewCurrencyRepository(db sqlx.Ext) storage.CurrencyRepository {
	return defaultCurrencyRepository{
		db: db,
	}
}

func (repo defaultCurrencyRepository) Create(product models.Currency) error {
	return nil
}

func (repo defaultCurrencyRepository) FindByID(id string) (*models.Currency, error) {
	return nil, nil
}

func (repo defaultCurrencyRepository) Update(product models.Currency) error {
	return nil
}

