package app

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Product struct {
	Id          string  `valid:"uuidv4"`
	Name        string  `valid:"required"`
	Price       float64 `valid:"float,optional"`
	Description string  `valid:"required"`
	Status      string  `valid:"required"`
}

func NewProduct(name string, price float64, description string, status string) (*Product, error) {
	product := &Product{
		Id:          uuid.New().String(),
		Name:        name,
		Price:       price,
		Description: description,
		Status:      status,
	}

	_, err := product.IsValid()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) IsValid() (bool, error) {
	if p.Id == "" {
		return false, errors.New("Product id is required")
	}

	if p.Name == "" {
		return false, errors.New("Product name is required")
	}

	if p.Price <= 0 {
		return false, errors.New("Product price is required")
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("Product status is required")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("Product price is required")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New("Product price need to be zero")
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetDescription() string {
	return p.Description
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetAllInfos() {
	fmt.Println("Product id: ", p.Id, "\nProduct name: ", p.Name, "\nProduct price: ", p.Price, "\nProduct description: ", p.Description, "\nProduct status: ", p.Status)
}
