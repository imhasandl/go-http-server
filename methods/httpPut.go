// Assignment
// Complete the updateUser and getUserById functions.

// The updateUser function takes a baseURL, id and apiKey string, and data User as input. It returns a User and error. It should:

// Encode the user data as JSON using json.Marshal
// Create a new request using http.NewRequest
// Set the method as PUT
// Set the url to fullURL
// Set the body as a bytes.Buffer containing the encoded JSON data using bytes.NewBuffer
// Modify the request headers
// Set the Content-Type header, with application/json as its value
// Set the X-API-Key header, with apiKey as its value
// Make the request using the http.Client's Do method. You can create a new http.Client or use the http.DefaultClient
// Decode and return the response's JSON body (which is also a User)
// getUserById
// getUserById takes a baseURL, id and apiKey string. It returns a User and error. It should:

// Create a new request using http.NewRequest
// Set the method as GET
// Set the url to fullURL
// Set the body as nil
// Modify the request headers
// Set the X-API-Key header, with apiKey as its value
// Make the request using the http.Client's Do method. You can create a new http.Client or use the http.DefaultClient
// Decode and return the response's JSON body (which is also a User)
// We've included the fullURL creation logic for you in both functions, we'll be talking more about URL building in the next chapter.


package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	userData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(userData))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("X-API-KEY", apiKey)
	
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var updUser User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&updUser)
	if err != nil {
		return User{}, err
	}

	return updUser, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Set("X-API-KEY", apiKey)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// don't touch below this line

type User struct {
	CharacterName string `json:"characterName"`
	Class         string `json:"class"`
	ID            string `json:"id"`
	Level         int    `json:"level"`
	PvpEnabled    bool   `json:"pvpEnabled"`
	User          struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func logUser(user User) {
	fmt.Printf("User uuid: %s, Character Name: %s, Class: %s, Level: %d, PVP Status: %t, User name: %s\n",
		user.ID, user.CharacterName, user.Class, user.Level, user.PvpEnabled, user.User.Name)
}
