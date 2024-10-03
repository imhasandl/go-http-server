// Assignment
// The backend team building Fantasy Quest's server has added support for query parameters! The server offers a sort parameter that can be optionally added by appending ?sort=fieldName to the end of the URL, where fieldName is the property of the resource you want the response records sorted by.

// We are creating a leaderboard and want to sort our user records by their level.

// Update the getUsers function's fullURL with a query parameter that will sort the user records by level.

package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	fullURL := url + "?sort=level"
	
	res, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var users []User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
