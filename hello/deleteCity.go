package main

import (
	"./controller"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(controller.DeleteCityHandler)
}
