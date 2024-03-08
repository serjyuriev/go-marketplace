package models

import (
	"errors"
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	mperr "productsvc/internal/pkg/errors"
)

// Product представляет собой товарную единицу маркетплейса.
type Product struct {
	Id          uuid.UUID
	Name        string
	Description string
	Quantity    int
	Price       *money.Money
}

// Validate выполняет валидацию полей Name и Description.
func (p *Product) Validate() error {
	if p.Name == "" {
		return mperr.NewProductValidationError("Name", errors.New("field cannot be empty"))
	}

	if p.Description == "" {
		return mperr.NewProductValidationError("Description", errors.New("field cannot be empty"))
	}

	return nil
}
