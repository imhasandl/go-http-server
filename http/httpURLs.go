// Assignment
// I've updated the getItemData() function to be a bit more flexible. It now takes a URL as a parameter.

// Try running the code in its current state. You should notice an error because the URL we're using is invalid.

// Fix the code so that the call to getItemData function uses the provided itemURL.

// This time the printed data won't be as ugly, I added a prettify function that adds some formatting.


package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"log"
)

func getItemData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return data, nil
}

func prettify(data string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(data), "", "  ")
	if err != nil {
		return "", fmt.Errorf("error indenting JSON: %w", err)
	}
	return prettyJSON.String(), nil
}

const itemURL = "https://api.boot.dev/v1/courses_rest_api/learn-http/items"

func main() {
	items, err := getItemData(itemURL)
	if err != nil {
		log.Fatalf("error getting item data: %v", err)
	}
	prettyData, err := prettify(string(items))
	if err != nil {
		log.Fatalf("error prettifying data: %v", err)
	}
	fmt.Println(prettyData)
}
