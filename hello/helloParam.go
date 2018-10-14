package main

import (
	"./model"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func HandlerParameter(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pathVariables := request.PathParameters
	name := pathVariables["name"]
	log.Println("Path parameters: ", pathVariables)
	if name == "" {
		name = "from API-Gateway lambda"
	}
	response := model.Response{Message: "Hello " + name}
	b_res, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	return events.APIGatewayProxyResponse{Body: string(b_res), StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandlerParameter)
}
