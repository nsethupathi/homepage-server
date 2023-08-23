package AwsService

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type PartiQLRunner struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

type Show struct {
	PK      int    `dynamodbav:"PK"`
	DATE    string `dynamodbav:"DATE"`
	ADDRESS string `dynamodbav:"ADDRESS"`
	VENUE   string `dynamodbav:"VENUE"`
}

func BuildRunner() *PartiQLRunner {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	return &PartiQLRunner{
		DynamoDbClient: dynamodb.NewFromConfig(cfg),
		TableName:      "SHOWS",
	}
}

func (runner PartiQLRunner) GetShows() ([]*Show, error) {
	var shows []*Show
	response, err := runner.DynamoDbClient.ExecuteStatement(
		context.TODO(),
		&dynamodb.ExecuteStatementInput{
			Statement: aws.String(fmt.Sprintf("SELECT * FROM \"%v\" WHERE PK = 1 AND DATE >= DATEADD(HOUR, -5, UTCNOW())", runner.TableName)),
		})
	if err != nil {
		log.Printf("Couldn't get info. Here's why: %v\n", err)
	} else {
		err = attributevalue.UnmarshalListOfMaps(response.Items, &shows)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}
	return shows, err
}
