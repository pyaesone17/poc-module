package dynamodb

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// NewRepository is factory
func NewRepository() *Repository {
	config := aws.NewConfig()
	sess, _ := session.NewSession(config)
	sess = session.Must(session.NewSession(config))
	client := dynamodb.New(sess)

	return &Repository{
		client: client,
	}
}

// Repository is struct
type Repository struct {
	client dynamodbiface.DynamoDBAPI
}

// Put is to create item
func (repository *Repository) Put(id string, data interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tableAttributes := map[string]interface{}{
		"id":   id,
		"data": data,
	}

	item, _ := dynamodbattribute.MarshalMap(tableAttributes)

	params := &dynamodb.PutItemInput{
		TableName: aws.String("demo"),
		Item:      item,
	}

	repository.client.PutItemWithContext(ctx, params)

	fmt.Println("Putting data using dynamodb database")
	fmt.Printf("Id: %s , data %v \n", id, data)
}

// Find is to get item
func (repository *Repository) Find(id string) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	data, _ := repository.client.GetItemWithContext(ctx, params)

	fmt.Println("Finding data using dynamodb database")
	fmt.Printf("Id: %s \n", id)

	return data
}
