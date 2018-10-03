package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received path: ", request.Path)
	response := Response{Message: "Hello from API-Gateway lambda"}
	b_res, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	return events.APIGatewayProxyResponse{Body: string(b_res), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
