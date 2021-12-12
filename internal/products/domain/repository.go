package domain

import (
	"context"
)

type Repository interface {
	GetAllProducts(ctx context.Context) []*Product

	AddNewProduct(ctx context.Context, product *Product)
	DeleteProduct(ctx context.Context, productName string)
}
