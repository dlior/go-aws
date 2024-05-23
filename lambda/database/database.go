package database

import (
	"fmt"
	"lambda/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
	tableName     string
}

func NewDynamoDBClient() DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return DynamoDBClient{
		databaseStore: db,
		tableName:     "StocksPrice",
	}
}

func (db DynamoDBClient) AddStocksPrice(stocksPrice types.StocksPrice) error {
	data, err := dynamodbattribute.MarshalMap(stocksPrice)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	item := &dynamodb.PutItemInput{
		TableName: aws.String(db.tableName),
		Item:      data,
	}

	if _, err := db.databaseStore.PutItem(item); err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
