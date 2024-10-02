// Assignment
// Update the getUserCode function to return the status code of the response.

// you should get a 404 response when requesting a user that doesn't exist!

package main

import (
	"net/http"
)

func getUserCode(url string) int {
	res, err := http.Get(url)
	if err != nil {
		return 0
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return  404
	}
	return res.StatusCode
}
