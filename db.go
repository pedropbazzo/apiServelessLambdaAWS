// db.go

package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
)
const AWS_REGION = "sa-east-1"
const TABLE_NAME = "go-serverless-api"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(AWS_REGION))

// GetUsers retrieves all the users from the DB
func GetUsers() ([]User, error) {

	input := &dynamodb.ScanInput{
	  TableName: aws.String(tableName),
	}
	result, err := db.Scan(input)
	if err != nil {
	  return []User{}, err
	}
	if len(result.Items) == 0 {
	  return []User{}, nil
	}
  
	var users[]User
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
	  return []User{}, err
	}
  
	return users, nil
  }
  
  // CreateUser inserts a new User item to the table.
  func CreateUser(user User) error {
  
	// Generates a new random ID
	uuid, err := uuid.NewV4()
	if err != nil {
	  return err
	}
  
	// Creates the item that's going to be inserted
	input := &dynamodb.PutItemInput{
	  TableName: aws.String(tableName),
	  Item: map[string]*dynamodb.AttributeValue{
		"id": {
		  S: aws.String(fmt.Sprintf("%v", uuid)),
		},
		"name": {
		  S: aws.String(user.Name),
		},
	  },
	}
  
	_, err = db.PutItem(input)
	return err
  }