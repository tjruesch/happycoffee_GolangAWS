package adapters

import "github.com/truesch/happycoffee_GolangAWS/internal/products/domain"

func GetAllProducts() []*domain.Product {
	products := []*domain.Product{}

	products = append(products, domain.NewProduct("Cappuchino", 3.5, 1))
	products = append(products, domain.NewProduct("Espresso", 2, 4))
	products = append(products, domain.NewProduct("Cold Brew", 4, 2))

	return products
}
