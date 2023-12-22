package main

import (
	"accumulation/internal/adapter"
	"accumulation/internal/app"
	"accumulation/internal/domain"
	"accumulation/internal/usecase"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaHandler struct {
	myApp app.MyApp
}

func NewLambdaHandler(pointRepository domain.PointRepository) *LambdaHandler {
	pointUsecase := usecase.NewPointUsecase(pointRepository)
	myApp := app.NewMyApp(*pointUsecase)
	return &LambdaHandler{
		myApp: *myApp,
	}
}

func (h *LambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	b := []byte(request.Body)
	var body domain.Point
	err := json.Unmarshal(b, &body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	err = h.myApp.HandleRequest(&body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("point: success create"),
	}, nil
}

func main() {
	tableName := "Point"
	pointRepository := adapter.NewDynamoDBRepository(tableName)
	handler := NewLambdaHandler(pointRepository)
	lambda.Start(handler.HandleRequest)
}
