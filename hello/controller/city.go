package controller

import (
	"../model"
	"../repo"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func PutCityHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := request.Body
	var city model.City
	e := json.Unmarshal([]byte(body), &city)
	if e != nil {
		bytes, _ := json.Marshal(model.Response{Message: "Cannot parse body as city"})
		return events.APIGatewayProxyResponse{Body: string(bytes), StatusCode: 400}, nil
	}
	repo.PutCity(&city)
	responseStr, _ := json.Marshal(model.Response{Message: "City was added"})
	return events.APIGatewayProxyResponse{Body: string(responseStr), StatusCode: 200}, nil
}

func ListCitiesHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cities := repo.ListCities()
	responseStr, _ := json.Marshal(cities)
	return events.APIGatewayProxyResponse{Body: string(responseStr), StatusCode: 200}, nil
}

func GetCityHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pathVariables := request.PathParameters
	id := pathVariables["id"]
	city := repo.FindCityById(id)
	if city.Id == "" {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 200}, nil
	}
	responseStr, _ := json.Marshal(city)
	return events.APIGatewayProxyResponse{Body: string(responseStr), StatusCode: 200}, nil
}

func DeleteCityHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pathVariables := request.PathParameters
	id := pathVariables["id"]
	repo.DeleteCity(id)
	bytes, _ := json.Marshal(model.Response{Message: fmt.Sprintf("city with id %s was deleted", id)})
	return events.APIGatewayProxyResponse{Body: string(bytes), StatusCode: 200}, nil
}
