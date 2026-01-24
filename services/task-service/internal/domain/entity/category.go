package entity

import "github.com/shopspring/decimal"

type Category struct {
	ID    string
	Title string
	Desc  string
	Price decimal.Decimal
}
