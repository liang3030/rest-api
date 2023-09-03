package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}
	switch request.HTTPMethod {
	case "GET":
		name := request.QueryStringParameters["name"]
		if name != "" {
			ApiResponse = events.APIGatewayProxyResponse{Body: "hey " + name + "welcome!", StatusCode: 200}
		} else {
			ApiResponse = events.APIGatewayProxyResponse{Body: "Error: Query Parametter name missing", StatusCode: 500}

		}
	}

	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}
