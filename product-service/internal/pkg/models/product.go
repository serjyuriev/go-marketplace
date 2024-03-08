package models

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

// Product представляет собой товарную единицу маркетплейса.
type Product struct {
	Id          uuid.UUID
	Name        string
	Description string
	Quantity    int
	Price       *money.Money
}
