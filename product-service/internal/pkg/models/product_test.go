package models

import (
	"errors"
	"github.com/stretchr/testify/assert"
	mperr "productsvc/internal/pkg/errors"
	"testing"
)

func TestProduct_Validate(t *testing.T) {
	tests := []struct {
		name    string
		product *Product
		err     error
	}{
		{
			name: "product with Name and Description",
			product: &Product{
				Name:        "SuperProduct",
				Description: "Sample description text!",
			},
			err: nil,
		},
		{
			name: "product with Name and without Description",
			product: &Product{
				Name:        "SomethingIsMissing",
				Description: "",
			},
			err: mperr.NewProductValidationError("Description", errors.New("field cannot be empty")),
		},
		{
			name: "product without Name and with Description",
			product: &Product{
				Name:        "",
				Description: "Sample description text!",
			},
			err: mperr.NewProductValidationError("Name", errors.New("field cannot be empty")),
		},
		{
			name: "product without Name and without Description",
			product: &Product{
				Name:        "",
				Description: "",
			},
			err: mperr.NewProductValidationError("Name", errors.New("field cannot be empty")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.product.Validate()

			assert.Equal(t, tt.err, actual)
		})
	}
}
