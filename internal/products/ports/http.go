package ports

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/domain"
)

type server struct {
	repo   domain.Repository
	Router chi.Mux
}

type ProductReadModel struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	HappyDay string `json:"happy_day"`
	Discount int    `json:"discount"`
}

func NewProductsServer(repo domain.Repository) *server {
	return &server{
		repo:   repo,
		Router: *chi.NewRouter(),
	}
}

func (s server) GetProducts(w http.ResponseWriter, r *http.Request) {
	productsInDB, err := s.repo.GetAllProducts()
	if err != nil {
		InternalError("Error getting products from database", err, w, r)
	}
	products := s.domainProductsToOutput(productsInDB)

	render.Respond(w, r, products)
}

func (s server) AddNewProduct(w http.ResponseWriter, r *http.Request) {
	NotImplementedError("Endpoint not implemented", nil, w, r)
}

func (s server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	NotImplementedError("Endpoint not implemented", nil, w, r)
}

func (s server) domainProductsToOutput(products []*domain.Product) []*ProductReadModel {
	out := []*ProductReadModel{}
	for _, p := range products {
		out = append(out, s.domainProductToReadModel(p))
	}

	return out
}

func (s server) domainProductToReadModel(p *domain.Product) *ProductReadModel {

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
