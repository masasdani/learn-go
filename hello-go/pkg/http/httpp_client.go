package http

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// CreateHTTPClient creates a new HTTP client with custom settings
func CreateHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 10, // Set a timeout of 10 seconds
	}
	return client
}

// ReadResponseBody reads the response body and parses the JSON
func ReadResponseBody(res *http.Response) (map[string]interface{}, error) {
	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Close the response body
	defer res.Body.Close()

	// Create a map to hold the parsed JSON
	var result map[string]interface{}

	// Parse the JSON
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
