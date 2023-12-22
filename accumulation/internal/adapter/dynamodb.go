package adapter

import (
	"accumulation/internal/domain"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBRepository struct {
	dynamoDBClient *dynamodb.DynamoDB
	tableName      string
}

func (r *DynamoDBRepository) GetPointByID(id string) (*domain.Point, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := r.dynamoDBClient.GetItem(input)
	if err != nil {
		return nil, err
	}

	point := &domain.Point{
		ID:   id,
		Name: aws.StringValue(result.Item["name"].S),
	}

	return point, nil
}

func NewDynamoDBRepository(tableName string) *DynamoDBRepository {
	sess := session.Must(session.NewSession())
	return &DynamoDBRepository{
		dynamoDBClient: dynamodb.New(sess),
		tableName:      tableName,
	}
}

func (r *DynamoDBRepository) CreatePoint(point *domain.Point) error {
	totalAsString := fmt.Sprintf("%.2f", point.Total)

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(point.ID),
			},
			"user": {
				S: aws.String(point.User),
			},
			"name": {
				S: aws.String(point.Name),
			},
			"total": {
				N: aws.String(totalAsString),
			},
		},
	}

	_, err := r.dynamoDBClient.PutItem(input)
	return err
}
