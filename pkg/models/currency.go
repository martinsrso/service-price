package models

// Currency represents a single currency
type Currency struct {
	ID            string  `db:"id" json:"id"`
	Name          string  `db':"name" json:"name"`
	PriceInDollar float64 `db:"price_in_dollar" json:"price_in_dollar"`
}

// NewCurrency is for registering a new currency
type NewCurrency struct {
	ID            string  `json:"id"`
	PriceInDollar float64 `json:"price_in_dollar"`
}
