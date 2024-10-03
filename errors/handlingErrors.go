// Assignment
// Complete the fetchData function. It needs to handle network errors and non-OK responses. It should return the status code, an int, and an error.

// Make a http.Get request to the url
// If a network error occurs, return a 0 status code, and the following error where %v is the error value.
// network error: %v
// Copy icon
// If the response does not have status code 200 return the status code, and the following error, where %s is the .Status field of the response.
// non-OK HTTP status: %s
// Copy icon
// If no errors occurred simply return the .StatusCode and nil.
// Note: the fmt package is already imported: you may find it helpful.

package main

import (
	"fmt"
	"net/http"
)

func fetchData(url string) (int, error) {
	req, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("network error: %v", err)
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return req.StatusCode, fmt.Errorf("non-OK HTTP status: %v", req.Status)
		
	}

	return req.StatusCode, nil
	
}
