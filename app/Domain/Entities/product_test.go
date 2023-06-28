package app_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	app "github.com/lucasf5/app/Domain/Entities"
)

const (
	ENABLED     = "enabled"
	DISABLED    = "disabled"
	Description = "Salada de frutas"
)

func TestProductIsValid(t *testing.T) {

	Salada, err := app.NewProduct("Salada", 25.00, Description, ENABLED)

	if err != nil {
		panic(err)
	}
	assert.Equal(t, Salada.Name, "Salada")
	assert.Equal(t, Salada.Price, 25.00)
	assert.Equal(t, Salada.Description, Description)
	assert.Equal(t, Salada.Status, ENABLED)
	require.NotEmpty(t, Salada.Id)
}

func TestProductEnable(t *testing.T) {

	Salada, err := app.NewProduct("Salada", 25.00, Description, DISABLED)

	if err != nil {
		panic(err)
	}

	err = Salada.Enable()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, Salada.Status, ENABLED)
}

func TestInvalidProductName(t *testing.T) {

	_, err := app.NewProduct("", 25.00, Description, ENABLED)

	if err != nil {
		assert.Equal(t, err.Error(), "Product name is required")
	}
}

func TestInvalidProductPrice(t *testing.T) {

	_, err := app.NewProduct("Salada", 0, Description, ENABLED)

	if err != nil {
		assert.Equal(t, err.Error(), "Product price is required")
	}
}

func TestInvalidProductStatus(t *testing.T) {

	_, err := app.NewProduct("Salada", 25.00, Description, "")

	if err != nil {
		assert.Equal(t, err.Error(), "Product status is required")
	}
}

func TestDisabledWhenPriceIsNotZero(t *testing.T) {
	Salada, err := app.NewProduct("Salada", 1, Description, ENABLED)

	if err != nil {
		panic(err)
	}

	err = Salada.Disable()

	if err != nil {
		assert.Equal(t, err.Error(), "Product price need to be zero")
	}

}
