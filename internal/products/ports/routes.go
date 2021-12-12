package ports

import "github.com/go-chi/chi"

func (s *server) SetRoutes() {
	s.Router.Route("/v1", func(r chi.Router) {
		r.Route("/products", func(r chi.Router) {
			r.Get("/", s.GetProducts)
			r.Post("/", s.AddNewProduct)
			r.Delete("/{productID}", s.DeleteProduct)
		})
	})
}
