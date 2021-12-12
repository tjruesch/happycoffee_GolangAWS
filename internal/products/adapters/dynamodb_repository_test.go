package adapters_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/adapters"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/domain"
)

// Please note that these Tests are coupled and depend on each other.
// In a more prodcution-like environment you would want to structure your
// tests differntly

const testTable string = "products_test"

func TestInit(t *testing.T) {
	dynamoRepo := adapters.NewProjectDynamoDBRepository()
	dynamoRepo.Init(testTable)
}

func TestSaveProduct(t *testing.T) {
	dynamoRepo := adapters.NewProjectDynamoDBRepository()
	dynamoRepo.TableName = testTable

	testDomainProject := domain.NewProduct("Black Coffee", 1.50, 0)
	testDomainProject2 := domain.NewProduct("Cappuccino", 1.50, 0)

	err := dynamoRepo.SaveProduct(testDomainProject)
	require.NoError(t, err)

	err = dynamoRepo.SaveProduct(testDomainProject2)
	require.NoError(t, err)
}

func TestGetAllProducts(t *testing.T) {
	dynamoRepo := adapters.NewProjectDynamoDBRepository()
	dynamoRepo.TableName = testTable

	prods, err := dynamoRepo.GetAllProducts()
	require.NoError(t, err)

	for _, p := range prods {
		t.Log(*p)
	}

	exp1 := domain.NewProduct("Black Coffee", 1.50, 0)
	exp2 := domain.NewProduct("Cappuccino", 1.50, 0)
	expected := []*domain.Product{exp1, exp2}

	require.Equal(t, prods, expected)

}

func TestGetProduct(t *testing.T) {
	dynamoRepo := adapters.NewProjectDynamoDBRepository()
	dynamoRepo.TableName = testTable

	prod, err := dynamoRepo.GetProduct("Black Coffee")
	require.NoError(t, err)

	require.Equal(t, prod, domain.NewProduct("Black Coffee", 1.50, 0))
}

func TestDeleteProduct(t *testing.T) {
	dynamoRepo := adapters.NewProjectDynamoDBRepository()
	dynamoRepo.TableName = testTable

	err := dynamoRepo.DeleteProduct("Black Coffee")
	require.NoError(t, err)
}
