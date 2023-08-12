package service

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type PartiQLRunner struct {
	DynamoDbClient *dynamodb.DynamoDB
	TableName      string
}

func Build() *PartiQLRunner {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	return &PartiQLRunner{
		DynamoDbClient: svc,
		TableName:      "SHOWS",
	}
}
