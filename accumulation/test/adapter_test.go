// test/adapter_test.go
package test

import (
	"accumulation/internal/adapter"
	"accumulation/internal/domain"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDynamoDBRepository struct {
	mock.Mock
}

func (m *MockDynamoDBRepository) GetPointByID(id string) (*domain.Point, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Point), args.Error(1)
}

func (m *MockDynamoDBRepository) CreatePoint(Point *domain.Point) error {
	args := m.Called(Point)
	return args.Error(0)
}

func TestDynamoDBRepository_GetPointByID(t *testing.T) {
	mockDynamoDB := &MockDynamoDBRepository{}
	repository := adapter.NewDynamoDBRepository(mockDynamoDB, "Points")

	// Implement your mock data and expectations
	mockDynamoDB.On("GetItem", mock.Anything).Return(&dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"id":   {S: aws.String("1")},
			"name": {S: aws.String("Compra TC")},
		},
	}, nil)

	Point, err := repository.GetPointByID("1")

	assert.NoError(t, err)
	assert.NotNil(t, Point)
	assert.Equal(t, "1", Point.ID)
	assert.Equal(t, "Compra TC", Point.Name)
}

func TestDynamoDBRepository_CreatePoint(t *testing.T) {
	mockDynamoDB := &MockDynamoDBRepository{}
	repository := adapter.NewDynamoDBRepository(mockDynamoDB, "Points")

	// Implement your mock data and expectations
	mockDynamoDB.On("PutItem", mock.Anything).Return(nil)

	err := repository.CreatePoint(&domain.Point{ID: "1", Name: "Jane Doe"})

	assert.NoError(t, err)
}
