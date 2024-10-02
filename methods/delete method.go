// Assignment
// Complete the deleteUser function. It takes a baseURL, id and apiKey as inputs. You'll need to:

// Create a new request using http.NewRequest and use the provided fullURL.
// Modify the request headers
// Set the X-API-Key header, with apiKey as its value
// Make the request using the http.Client's Do method. You can create a new http.Client or use the http.DefaultClient
// Check the response status code. If the status code indicates a non-successful response, return an error. Otherwise return nil.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-API-KEY", apiKey)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	res.Body.Close()

	status := res.StatusCode
	if status >= 400 || status >= 500 {
		return nil
	}

	return nil
}

// don't touch below this line

func getUsers(url, apiKey string) ([]User, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
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

func logUsers(users []User) {
	for _, user := range users {
		fmt.Printf("Character name: %s, Class: %s, Level: %d, User: %s\n", user.CharacterName, user.Class, user.Level, user.User.Name)
	}
}

type User struct {
	Id            string `json:"id"`
	CharacterName string `json:"characterName"`
	Class         string `json:"class"`
	Level         int    `json:"level"`
	User          struct {
		Name string `json:"name"`
	} `json:"user"`
}
