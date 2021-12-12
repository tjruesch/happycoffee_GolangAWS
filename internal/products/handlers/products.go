package handlers

import (
	"fmt"
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
		out = append(out, domainProductToReadModel(p))
	}

	return out
}

func domainProductToReadModel(p *domain.Product) *ProductReadModel {

	var day string
	switch p.HappyDay() {
	case 0:
		day = "Sunday"
	case 1:
		day = "Monday"
	case 2:
		day = "Tuesday"
	case 3:
		day = "Wednesday"
	case 4:
		day = "Thursday"
	case 5:
		day = "Friday"
	case 6:
		day = "Saturday"
	}

	return &ProductReadModel{
		Name:     p.Name(),
		Price:    fmt.Sprintf("%.2f", p.Price()) + " EUR",
		HappyDay: day,
		Discount: p.DiscountInPercent(),
	}
}
