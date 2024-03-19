package handlers

import "github.com/aws/aws-lambda-go/events"

func DefaultHandler(event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       "Hello, World!",
	}, nil
}
