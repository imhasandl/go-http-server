// Assignment
// Complete the createUser function. It should:

// Take a URL and apiKey string, and User data as parameters
// Encode the user data using json.Marshal
// Create a new POST request using http.NewRequest. Use a bytes.NewBuffer to create a io.Reader from the JSON data.
// Modify the request headers:
// Set the Content-Type header, with application/json as its value
// Set the X-API-Key header, with apiKey as its value
// Make the request using the http.Client's Do method. You can create a new http.Client or use the http.DefaultClient
// Decode and return the response's JSON body (which is also a User)
// Don't copy paste from the code above. Type it out and understand what each line does.


package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
	userData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(userData))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer req.Body.Close()

	var newUser User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&newUser)
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

// Don't touch below this line

type User struct {
	CharacterName string `json:"characterName"`
	Class         string `json:"class"`
	Level         int    `json:"level"`
	PvpEnabled    bool   `json:"pvpEnabled"`
	User          struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func getUsers(url, apiKey string) ([]User, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []User{}, err
	}

	req.Header.Set("X-API-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []User{}, err
	}
	defer res.Body.Close()

	var users []User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&users)
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

func logUsers(users []User) {
	for _, user := range users {
		fmt.Printf("Character name: %s, Class: %s, Level: %d, User: %s\n", user.CharacterName, user.Class, user.Level, user.User.Name)
	}
}
