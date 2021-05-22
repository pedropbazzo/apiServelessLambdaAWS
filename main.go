// main.go

package main

import (
  "net/http"

  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

// handleCreateUser

// handleGetUsers

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  if req.Path == "/users" {
    if req.HTTPMethod == "GET" {
      return handleGetUsers(req)
    }
    if req.HTTPMethod == "POST" {
      return handleCreateUser(req)
    }
  }

  return events.APIGatewayProxyResponse{
    StatusCode: http.StatusMethodNotAllowed,
    Body:       http.StatusText(http.StatusMethodNotAllowed),
  }, nil
}

func main() {
  lambda.Start(router)
}

func handleGetUsers(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	users, err := GetUsers()
	if err != nil {
	  return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body: http.StatusText(http.StatusInternalServerError),
	  }, nil
	}
  
	js, err := json.Marshal(users)
	if err != nil {
	  return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body: http.StatusText(http.StatusInternalServerError),
	  }, nil
	}
  
	return events.APIGatewayProxyResponse{
	  StatusCode: http.StatusOK,
	  Body: string(js),
	}, nil
  }
  
  func handleCreateUser(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  
	var user User
	err := json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
	  return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       err.Error(),
	  }, nil
	}
  
	err = CreateUser(user)
	if err != nil {
	  return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       err.Error(),
	  }, nil
	}
  
	return events.APIGatewayProxyResponse{
	  StatusCode: http.StatusCreated,
	  Body:       "Created",
	}, nil
  }