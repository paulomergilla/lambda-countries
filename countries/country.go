package country

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var dynamoClient *dynamodb.Client

type Country struct {
	CountryID   string `json:"CountryID"`
	CountryName string `json:"CountryName"`
	Continent   string `json:"Continent"`
}

func init() {
	InitializeDynamo()
}

func InitializeDynamo() {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("ap-southeast-2"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	dynamoClient = dynamodb.NewFromConfig(cfg)
}

func FetchCountries() ([]Country, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Countries"),
	}

	result, err := dynamoClient.Scan(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var messages []Country
	err = attributevalue.UnmarshalListOfMaps(result.Items, &messages)
	return messages, err
}
