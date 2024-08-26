package main

import (
	country "countries/countries"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Initializing DynamoDB b0011")

	country.InitializeDynamo()

	body, err := country.FetchCountries()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	jbody, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	log.Println("saving chat successful, sending response")
	log.Println(body)
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jbody),
	}

	return response, nil
}
