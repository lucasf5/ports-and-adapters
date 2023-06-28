package main

import (
	repositories "github.com/lucasf5/app/Domain/Repositories"
)

type ProductService struct {
	Persistence repositories.ProductPersistenceInterface
}

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

func main() {
}
