package service

import (
	context "context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

type Show struct {
	PK      int    `dynamodbav:"PK"`
	DATE    string `dynamodbav:"DATE"`
	ADDRESS string `dynamodbav:"ADDRESS"`
	VENUE   string `dynamodbav:"VENUE"`
	STATUS  string `dynamodbav:"STATUS"`
}

func (runner PartiQLRunner) GetShows() (Show, error) {
	var show Show
	params, err := attributevalue.MarshalList([]interface{}{})
	if err != nil {
		panic(err)
	}
	response, err := runner.DynamoDbClient.ExecuteStatement(context.TODO(), &dynamodb.ExecuteStatementInput{
		Statement: aws.String(
			fmt.Sprintf("SELECT * FROM mytable WHERE PK = 1 AND STATUS = 'UPCOMING' ORDER BY DATE",
				runner.TableName)),
		Parameters: params,
	})
	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", title, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Items[0], &show)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}
	return movie, err
}
