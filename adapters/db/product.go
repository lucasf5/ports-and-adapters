package db

import (
	"database/sql"

	product "github.com/lucasf5/app/Domain/Entities"
	app "github.com/lucasf5/app/Domain/Repositories"
)

type ProductDB struct {
	db *sql.DB
}

func (p *ProductDB) GetProduct(id string) (app.ProductInterface, error) {
	var product product.Product

	return &product, nil
}
