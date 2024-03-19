package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func createHTTPClient() *http.Client {
	// Create a new HTTP client with custom settings
	client := &http.Client{
		Timeout: time.Second * 10, // Set a timeout of 10 seconds
	}

	return client
}

func main() {
	// Create an HTTP client
	client := createHTTPClient()
	res, err := client.Get("https://dns.google/resolve?name=masasdani.com&type=A")
	if err != nil {
		// Handle error
		log.Fatalf("Error: %v", err)
	}

	result, err := readResponseBody(res)
	if err != nil {
		// Handle error
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(result["Answer"].([]interface{})[0].(map[string]interface{})["data"])

	// Use the client to make HTTP requests
	// ...
}

func readResponseBody(res *http.Response) (map[string]interface{}, error) {
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
