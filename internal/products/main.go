package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/adapters"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/ports"
)

func main() {
	// initialize Data
	repo := adapters.NewProjectDynamoDBRepository()
	repo.Init("happycoffe_products")

	go loadFixtures(repo)

	// initialize HTTP server
	productsServer := ports.NewProductsServer(repo)
	productsServer.SetMiddlewares()
	productsServer.SetRoutes()

	logrus.Info("Starting HTTP server")

	http.ListenAndServe(":9090", &productsServer.Router)
}
