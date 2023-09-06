package showrepository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"time"
)

type ShowDO struct {
	PK      int    `dynamodbav:"PK"`
	DATE    string `dynamodbav:"DATE"`
	ADDRESS string `dynamodbav:"ADDRESS"`
	VENUE   string `dynamodbav:"VENUE"`
}

func GetShows(showRunner *TableRunner) ([]*ShowDO, error) {
	var err error
	var response *dynamodb.QueryOutput
	var shows []*ShowDO

	key := expression.KeyAnd(
		expression.Key("PK").Equal(expression.Value(1)),
		expression.Key("DATE").GreaterThanEqual(expression.Value(time.Now().Format("2006-01-02T15:04:05"))))
	expr, err := expression.NewBuilder().WithKeyCondition(key).Build()
	if err != nil {
		log.Printf("Couldn't build expression for query. Here's why: %v\n", err)
	} else {
		response, err = showRunner.DynamoDbClient.Query(context.TODO(), &dynamodb.QueryInput{
			KeyConditionExpression:    expr.KeyCondition(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			TableName:                 aws.String(showRunner.TableName),
		})
		if err != nil {
			log.Printf("Couldn't query for shows. Here's why: %v\n", err)
		} else {
			err = attributevalue.UnmarshalListOfMaps(response.Items, &shows)
			if err != nil {
				log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
			}
		}
	}
	return shows, err
}
