package usecases

import "github.com/martinsrso/service-price/pkg/models"

type ProductUserCase interface {
	GetPrices(id string) (*models.Product, error)
	Register(product models.Product) (*models.Product, error)
}

type CurrencyUserCase interface {
	Create(currency models.Currency) (*models.Currency, error)
}
