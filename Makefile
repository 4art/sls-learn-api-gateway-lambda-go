build:
	go get github.com/aws/aws-lambda-go/lambda
	go get github.com/aws/aws-lambda-go/events
	go get github.com/aws/aws-sdk-go/service/dynamodb
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/hello.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/helloParam hello/helloParam.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/listCities hello/listCities.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/putCity hello/putCity.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/getCity hello/getCity.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/deleteCity hello/deleteCity.go
