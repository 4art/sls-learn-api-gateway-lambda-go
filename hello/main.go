package main

import (
	"./model"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Received path: ", request.Path)
	response := model.Response{Message: "Hello from API-Gateway lambda"}
	b_res, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	return events.APIGatewayProxyResponse{Body: string(b_res), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
