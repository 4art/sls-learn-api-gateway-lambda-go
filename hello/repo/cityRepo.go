package repo

import (
	"../model"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
	"os/exec"
	"strings"
)

var svc = createService()
var table = os.Getenv("DYNAMODB_TABLE")

func PutCity(city *model.City) {
	city.Id = generateId()
	values, e := dynamodbattribute.MarshalMap(city)
	if e != nil {
		panic(e)
	}
	_, err := svc.PutItem(&dynamodb.PutItemInput{
		Item:      values,
		TableName: &table,
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func ListCities() []model.City {
	scanOutput, e := svc.Scan(&dynamodb.ScanInput{TableName: &table})
	if e != nil {
		panic(e)
	}
	var cities []model.City
	var city model.City
	items := scanOutput.Items
	for _, item := range items {
		dynamodbattribute.UnmarshalMap(item, &city)
		cities = append(cities, city)
	}
	return cities
}

func FindCityById(id string) model.City {
	key := map[string]*dynamodb.AttributeValue{
		"id": {S: &id},
	}
	itemOutput, e := svc.GetItem(&dynamodb.GetItemInput{
		Key:       key,
		TableName: &table,
	})
	if e != nil {
		panic(e)
	}
	var city model.City
	dynamodbattribute.UnmarshalMap(itemOutput.Item, &city)
	return city
}

func DeleteCity(id string) {
	key := map[string]*dynamodb.AttributeValue{
		"id": {S: &id},
	}
	svc.DeleteItem(&dynamodb.DeleteItemInput{
		Key:       key,
		TableName: &table,
	})
}

//TODO add update
func createService() *dynamodb.DynamoDB {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	return dynamodb.New(sess)
}

func generateId() string { //TODO delete \n in the end
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic(err)
	}
	id := string(out)
	id = strings.Replace(id, "\n", "", -1)
	return id
}
