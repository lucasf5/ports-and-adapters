package repositories

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetPrice() float64
	GetDescription() string
	GetStatus() string
	GetAllInfos()
}

type ProductReaderInterface interface {
	GetAll() ([]ProductInterface, error)
	Get(id string) (ProductInterface, error)
}

type ProductWriterInterface interface {
	Save(name string, price float64, description string, status string) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReaderInterface
	ProductWriterInterface
}

type ProductServiceInterface interface {
	ProductPersistenceInterface
}
