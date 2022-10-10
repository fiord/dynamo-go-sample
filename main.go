package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Sample struct {
	ID   string `dynamo:"id,hash"`
	Name string `dynamo:"name"`
}

func main() {
	disableSSL := true
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String("localhost:8000"),
		DisableSSL:  &disableSSL,
		Credentials: credentials.NewStaticCredentials("id", "secret", "token"),
	})
	if err != nil {
		fmt.Printf("[error] %v\n", err)
		return
	}

	svc := dynamodb.New(sess)
	put := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("foo"),
			},
			"Name": {
				S: aws.String("bar"),
			},
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String("example"),
	}
	putResult, err := svc.PutItem(put)
	if err != nil {
		fmt.Printf("[error] %v\n", err)
		return
	}
	fmt.Printf("Put Result: %v\n", putResult)

	get := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("foo"),
			},
		},
		TableName: aws.String("example"),
	}
	getResult, err := svc.GetItem(get)
	if err != nil {
		fmt.Printf("[error] %v\n", err)
		return
	}
	fmt.Printf("Get Result: %v\n", getResult)
}
