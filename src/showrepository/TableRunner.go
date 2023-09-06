package showrepository

import (
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"golang.org/x/net/context"
	"log"
)

type TableRunner struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func BuildRunner(tableName string) *TableRunner {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	return &TableRunner{
		DynamoDbClient: dynamodb.NewFromConfig(cfg),
		TableName:      tableName,
	}
}
