package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Your code here
	lambda.Start(HandleRequest)
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Request struct {
	Name string `json:"name"`
}

func HandleRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Received request: ", event.Body)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, " + event.Body,
	}

	return response, nil
}
