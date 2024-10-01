// Assignment
// Update the getItemData function in http.go.

// Change the return signature to return []Item instead of []byte.
// Because the function will now return a decoded slice of items, change the name from getItemData to getItems.
// Get the data from the response body using io.ReadAll, creating a slice of bytes []byte.
// Create a nil slice of items []Item.
// Use json.Unmarshal on the data to get the JSON data.
// Return the items.


package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getItems(url string) ([]Item, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong: %w", err)
	}

	var items []Item
	if err = json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	return items, nil
}


