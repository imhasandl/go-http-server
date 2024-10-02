// Assignment
// Complete the getContentType function. It takes a pointer to a http.Response as input and should return the Content-Type header.

// Use the .Get() method on the Response struct's Header field to get what you need.

// Note: even though we've emphasized the need to close the response body, this is a mock response without a body. Trying to .Close() it will cause a panic.



package main

import (
	"net/http"
)

func getContentType(res *http.Response) string {
	header := res.Header.Get("Content-Type")

	return header
}
