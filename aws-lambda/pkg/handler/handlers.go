package handler

import "github.com/aws/aws-lambda-go/events"

func DefaultHandler(event events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	path := event.RequestContext.HTTP.Path
	switch path {
	case "/test":
		return PingHandler(event)
	case "/":
		return ok("Hello, World!")
	default:
		return apiResponse(404, "Not Found")
	}
}

func PingHandler(event events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	return ok("up")
}
