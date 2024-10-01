// Assignment
// Take a look at the getItemData function that I've provided in http.go. It retrieves information about items from Fantasy Quest's servers via HTTP as a slice of bytes []byte.

// In main.go do the following:

// Convert the slice of bytes to a string with the built-in string() type conversion.
// Print the string representation of the bytes to the console.
// It should look like a big ugly string of text.


package main

import (
	"fmt"
	"io"
	"net/http"
	"log"
)

func getItemData() ([]byte, error) {
	res, err := http.Get("https://api.boot.dev/v1/courses_rest_api/learn-http/items")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	return data, err
}


func main() {
	items, err := getItemData()
	if err != nil {
		log.Fatalf("error getting item data: %v", err)
	}

	// Don't edit above this 

	stringifiedData := string(items)
	fmt.Println(stringifiedData)
	
}
