package models

// Product represents a single product
type Product struct {
	ID            string  `db:"id" json:"id"`
	PriceInDollar float64 `db:"price_in_dollar" json:"price_in_dollar"`
}

// NewProduct is for registering a new product
type NewProduct struct {
	ID            string  `json:"id"`
	PriceInDollar float64 `json:"price_in_dollar"`
}
