package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"masasdani.com/lambda-learn/pkg/handler"
)

func main() {
	lambda.Start(handler.DefaultHandler)
}
