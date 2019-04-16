package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var dynamoLocalEndpoint string

// DynamoDb creates DynamoDb Object.
func DynamoDb() *dynamo.DB {
	config := &aws.Config{Region: aws.String("ap-northeast-1")}
	if dynamoLocalEndpoint != "" {
		config.Endpoint = aws.String(dynamoLocalEndpoint)
	}
	return dynamo.New(session.New(), config)
}

// Table create Table Object.
func Table(table string) dynamo.Table {
	db := DynamoDb()
	return db.Table(table)
}
