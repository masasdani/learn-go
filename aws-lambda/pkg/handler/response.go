package handler

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{}) (*events.APIGatewayV2HTTPResponse, error) {
	resp := events.APIGatewayV2HTTPResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}

func ok(body interface{}) (*events.APIGatewayV2HTTPResponse, error) {
	return apiResponse(200, body)
}
