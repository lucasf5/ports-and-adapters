package app

import (
	product "github.com/lucasf5/app/Domain/Entities"
	repositories "github.com/lucasf5/app/Domain/Repositories"
)

type ProductService struct {
	Persistence repositories.ProductPersistenceInterface
}

func (s *ProductService) GetAll() ([]repositories.ProductInterface, error) {
	products, err := s.Persistence.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) Get(id string) (repositories.ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Save(name string, price float64, description string, status string) (repositories.ProductInterface, error) {
	newProduct, err := product.NewProduct(name, price, description, status)
	if err != nil {
		return nil, err
	}

	product, err := s.Persistence.Save(newProduct.Name, newProduct.Price, newProduct.Description, newProduct.Status)
	if err != nil {
		return nil, err
	}

	return product, nil

}
