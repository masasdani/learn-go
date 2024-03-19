package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"masasdani.com/lambda-learn/internal/handlers"
)

func main() {
	lambda.Start(handlers.DefaultHandler)
}
