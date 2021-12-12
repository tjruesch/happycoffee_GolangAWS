package domain

type Repository interface {
	GetAllProducts() ([]*Product, error)

	SaveProduct(product *Product) error
	DeleteProduct(productName string) error
}
