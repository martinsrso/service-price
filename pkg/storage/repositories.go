package storage

import "github.com/martinsrso/service-price/pkg/models"

// ProductRepository is for interacting with product
type ProductRepository interface {
	FindByID(productID string) (*models.Product, error)
	Create(product models.Product) error
	Update(product models.Product) error
}

// CurrencyRepository is for interacting with currency
type CurrencyRepository interface {
	FindByID(currencyID string) (*models.Currency, error)
	Create(currency models.Currency) error
	Update(currency models.Currency) error
}
