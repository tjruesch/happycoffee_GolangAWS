package ports_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/adapters"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/ports"
)

func TestAddProduct(t *testing.T) {
	repo := adapters.NewProjectDynamoDBRepository()
	repo.Init("happycoffee_products_test")
	server := ports.NewProductsServer(repo)

	buf := bytes.NewBuffer(nil)
	proIn := newProductIn()

	err := json.NewEncoder(buf).Encode(proIn)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/products", buf)

	server.AddNewProduct(w, r)

	require.Equal(t, w.Result().StatusCode, http.StatusOK)
}

func TestGetProducts(t *testing.T) {
	repo := adapters.NewProjectDynamoDBRepository()
	repo.TableName = "happycoffee_products_test"
	server := ports.NewProductsServer(repo)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/products", nil)

	server.GetProducts(w, r)

	t.Log(w.Body.String())
	require.Equal(t, w.Result().StatusCode, http.StatusOK)
}

func TestDeleteProduct(t *testing.T) {
	repo := adapters.NewProjectDynamoDBRepository()
	repo.Init("happycoffee_products_test")
	server := ports.NewProductsServer(repo)
	server.SetMiddlewares()
	server.SetRoutes()

	// Add product to delete
	buf := bytes.NewBuffer(nil)
	proIn := newProductIn()

	err := json.NewEncoder(buf).Encode(proIn)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/products", buf)

	server.AddNewProduct(w, r)

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodDelete, "/v1/products/Some-Coffee", nil)

	// chi middleware is needed to parse the url params
	server.Router.ServeHTTP(w, r)
	require.Equal(t, w.Result().StatusCode, http.StatusOK)
}

func newProductIn() *ports.ProductInput {
	return &ports.ProductInput{
		Name:     "Some Coffee",
		Price:    3.45,
		HappyDay: "Monday",
	}
}
