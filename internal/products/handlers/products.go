package handlers

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/adapters"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/domain"
)

type ProductReadModel struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	HappyDay string `json:"happy_day"`
	Discount int    `json:"discount"`
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	productsInDB := adapters.GetAllProducts()
	products := domainProductsToOutput(productsInDB)

	render.Respond(w, r, products)
}

func AddNewProduct(w http.ResponseWriter, r *http.Request) {
	NotImplementedError("Endpoint not implemented", nil, w, r)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	NotImplementedError("Endpoint not implemented", nil, w, r)
}

func domainProductsToOutput(products []*domain.Product) []*ProductReadModel {
	out := []*ProductReadModel{}
	for _, p := range products {
		out = append(out, domainProductToOutput(p))
	}

	return out
}

func domainProductToOutput(p *domain.Product) *ProductReadModel {
	return &ProductReadModel{
		Name:     p.Name(),
		Price:    p.Price(),
		HappyDay: p.HappyDay(),
		Discount: p.DiscountInPercent(),
	}
}
