package ports

import (
	"encoding/json"
	"errors"
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

type ProductOutput struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	HappyDay string `json:"happy_day"`
	Discount int    `json:"discount"`
}

type ProductInput struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	HappyDay string  `json:"happy_day"`
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
	productIn := ProductInput{}

	err := json.NewDecoder(r.Body).Decode(&productIn)
	if err != nil {
		BadRequest("Invalid input", err, w, r)
	}

	p, err := s.inputToDomainProduct(&productIn)
	if err != nil {
		BadRequest("Invalid weekday", err, w, r)
	}

	err = s.repo.SaveProduct(p)
	if err != nil {
		InternalError("Error saving product to database", err, w, r)
	}
}

func (s server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	NotImplementedError("Endpoint not implemented", nil, w, r)
}

func (s server) domainProductsToOutput(products []*domain.Product) []*ProductOutput {
	out := []*ProductOutput{}
	for _, p := range products {
		out = append(out, s.domainProductToOutput(p))
	}

	return out
}

func (s server) domainProductToOutput(p *domain.Product) *ProductOutput {

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

	return &ProductOutput{
		Name:     p.Name(),
		Price:    fmt.Sprintf("%.2f", p.Price()) + " EUR",
		HappyDay: day,
		Discount: p.DiscountInPercent(),
	}
}

func (s server) inputToDomainProduct(in *ProductInput) (*domain.Product, error) {
	var day int
	switch in.HappyDay {
	case "Sunday":
		day = 0
	case "Monday":
		day = 1
	case "Tuesday":
		day = 2
	case "Wednesday":
		day = 3
	case "Thursday":
		day = 4
	case "Friday":
		day = 5
	case "Saturday":
		day = 6
	default:
		day = -1
	}

	if day == -1 {
		return &domain.Product{}, errors.New("invalid day")
	}

	return domain.NewProduct(in.Name, in.Price, day), nil
}
