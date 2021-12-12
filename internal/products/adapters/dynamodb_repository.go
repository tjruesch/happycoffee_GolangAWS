package adapters

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/sirupsen/logrus"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/domain"
)

type ProductsDynamoDBRepository struct {
	DynamoClient *dynamodb.DynamoDB
	TableName    string
}

type ProductModel struct {
	Name     string
	Price    float32
	HappyDay int
}

// Note that you need AWS credentials for these operations. In this setup
// you can store your credentials in the .env file
func NewProjectDynamoDBRepository() *ProductsDynamoDBRepository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String("eu-central-1")},
	}))

	return &ProductsDynamoDBRepository{DynamoClient: dynamodb.New(sess)}
}

// Initialize DynamoDB
func (r *ProductsDynamoDBRepository) Init(tablename string) error {

	r.TableName = tablename

	err := r.CreateTable(true)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceInUseException:
				return nil
			default:
				logrus.Error(aerr.Error())
				return err
			}
		}
		return err
	}

	return nil
}

// Create a new Table
func (r ProductsDynamoDBRepository) CreateTable(waitfortable bool) error {
	in := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Name"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Name"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(r.TableName),
	}

	_, err := r.DynamoClient.CreateTable(in)

	// stop and return if error
	if err != nil {
		return err
	}

	// check if table is created
	if waitfortable {
		err = r.waitForTable()
		if err != nil {
			return err
		}
	}

	logrus.Info("created table ", r.TableName)

	return nil
}

// Assert that the table is actually created
func (r ProductsDynamoDBRepository) waitForTable() error {
	err := r.DynamoClient.WaitUntilTableExists(
		&dynamodb.DescribeTableInput{
			TableName: aws.String(r.TableName),
		},
	)

	return err
}

// Save
func (r ProductsDynamoDBRepository) SaveProduct(domainProduct *domain.Product) error {
	av, err := r.marshalProject(domainProduct)
	if err != nil {
		return err
	}

	in := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(r.TableName),
	}

	_, err = r.DynamoClient.PutItem(in)

	return err
}

// Get single product
func (r ProductsDynamoDBRepository) GetProduct(name string) (*domain.Product, error) {

	in := &dynamodb.GetItemInput{
		TableName: &r.TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(name),
			},
		},
	}

	obj, err := r.DynamoClient.GetItem(in)
	if err != nil {
		logrus.Error("unable to get item from DynamDB")
		return &domain.Product{}, err
	}

	projectModel, err := r.unmarshalProject(obj.Item)

	return projectModel, err
}

// Get all products
func (r ProductsDynamoDBRepository) GetAllProducts() ([]*domain.Product, error) {

	proj := expression.NamesList(expression.Name("Name"), expression.Name("HappyDay"), expression.Name("Price"))

	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		logrus.Error("Got error building expression: ", err)
		return []*domain.Product{}, err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(r.TableName),
	}

	result, err := r.DynamoClient.Scan(params)
	if err != nil {
		logrus.Error("Query API call failed: ", err)
		return []*domain.Product{}, err
	}

	var projects []*domain.Product
	for _, i := range result.Items {

		item, err := r.unmarshalProject(i)
		if err != nil {
			logrus.Error("Got error unmarshalling: ", err)
			return []*domain.Product{}, err
		}

		projects = append(projects, item)
	}

	return projects, err
}

// Delete product
func (r ProductsDynamoDBRepository) DeleteProduct(name string) error {

	in := &dynamodb.DeleteItemInput{
		TableName:    &r.TableName,
		ReturnValues: aws.String("ALL_OLD"),
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(name),
			},
		},
	}

	out, err := r.DynamoClient.DeleteItem(in)
	if err != nil {
		logrus.Error("unable to delete item in dynamodb")
		return err
	}

	// DeleteItem will not fail if no item with the name is found
	// but DeleteItemOutput.Attributes will be nil
	if out.Attributes == nil {
		return errors.New("item not found")
	}

	return err
}

func (r ProductsDynamoDBRepository) marshalProject(p *domain.Product) (map[string]*dynamodb.AttributeValue, error) {
	pm := ProductModel{
		Name:     p.Name(),
		Price:    p.Price(),
		HappyDay: p.HappyDay(),
	}

	putInput, err := dynamodbattribute.MarshalMap(pm)
	if err != nil {
		return map[string]*dynamodb.AttributeValue{}, err
	}

	return putInput, err
}

func (r ProductsDynamoDBRepository) unmarshalProject(attribute map[string]*dynamodb.AttributeValue) (*domain.Product, error) {
	pm := &ProductModel{}
	err := dynamodbattribute.UnmarshalMap(attribute, pm)
	if err != nil {
		logrus.Error("could not unmarshal project from dynamodb")
		return &domain.Product{}, err
	}

	return domain.NewProduct(pm.Name, pm.Price, pm.HappyDay), nil
}
