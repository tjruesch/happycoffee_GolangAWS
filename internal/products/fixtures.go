package main

import "github.com/truesch/happycoffee_GolangAWS/internal/products/domain"

func loadFixtures(repo domain.Repository) {
	err := initProducts(repo)
	if err != nil {
		panic("could not load fixtures")
	}
}

func initProducts(repo domain.Repository) error {
	products := []*domain.Product{
		domain.NewProduct("Cappuccino", 3.5, 1),
		domain.NewProduct("Esspresso", 2, 4),
		domain.NewProduct("Cold Brew", 4, 0),
		domain.NewProduct("Americano", 3.5, 3),
	}

	for _, p := range products {
		err := repo.SaveProduct(p)
		if err != nil {
			return err
		}
	}

	return nil
}
