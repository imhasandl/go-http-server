// Assignment
// In previous lessons, we've converted response into slices of bytes, and then strings. Now, decode the response data directly into a slice of items and return that instead.

// Create a nil slice of items []Item.
// Create a new *json.Decoder using json.NewDecoder.
// Decode the response body using the decoder's Decode method.
// Return the slice of items.
// If any error occurs return a nil slice and the error.


package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getItems(url string) ([]Item, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var urls []Item
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&urls); err != nil {
		return nil, fmt.Errorf("Something went wrong: %w", err)
	}

	return urls, nil
}




