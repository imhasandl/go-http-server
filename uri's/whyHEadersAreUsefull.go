// Assignment
// Because we don't want any of our users to accidentally overwrite another user's saved data, our backend team has required that we use the X-API-Key header for all requests to Fantasy-Quest's server. When we use different values, we're telling the server that we are a different person.

// There's a bug in our main function! It should:

// Get a Location from our game server.
// Log that location to the console.
// Update the location and send those changes back to the server.
// Get the location struct again and log the updated location to the console.
// Run the code in its current state to see the return values.
// Notice that the two structs that are logged to the console are the same. That's because the X-API-Key they are using is different, the update isn't being applied to the same account that we're checking in the last getLocationResponse call.

// Use the same apiKey value in the last call to getLocationResponse.


package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"fmt"
)


func main() {
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/locations/52fdfc07-2182-454f-963f-5f0f9a621d72"
	apiKey := generateKey()

	oldLocation, err := getLocationResponse(apiKey, url)
	if err != nil {
		fmt.Println("Error getting old location:", err)
		return
	}
	fmt.Println("Got old location:")
	fmt.Printf("- name: %s\n", oldLocation.Name)
	fmt.Printf("- recommendedLevel: %d\n", oldLocation.RecommendedLevel)
	fmt.Println("--------------------------------")

	newLocationData := Location{
		Discovered:       false,
		ID:               "52fdfc07-2182-454f-963f-5f0f9a621d72",
		Name:             "Bloodstone Swamp",
		RecommendedLevel: 10,
	}

	if err := putLocation(apiKey, url, newLocationData); err != nil {
		fmt.Println("Error updating location:", err)
		return
	}
	fmt.Println("Location updated!")
	fmt.Println("---")

	newLocation, err := getLocationResponse(apiKey, url)
	if err != nil {
		fmt.Println("Error getting new location:", err)
		return
	}
	fmt.Println("Got new location:")
	fmt.Printf("- name: %s\n", newLocation.Name)
	fmt.Printf("- recommendedLevel: %d\n", newLocation.RecommendedLevel)
	fmt.Println("--------------------------------")
}


type Location struct {
	Discovered       bool   `json:"discovered"`
	ID               string `json:"id"`
	Name             string `json:"name"`
	RecommendedLevel int    `json:"recommendedLevel"`
}

func generateKey() string {
	const characters = "ABCDEF0123456789"
	result := ""
	rand.New(rand.NewSource(0))
	for i := 0; i < 16; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}
	return result
}

func getLocationResponse(apiKey, url string) (Location, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return Location{}, err
	}

	return location, nil
}

func putLocation(apiKey, url string, location Location) error {
	client := &http.Client{}
	data, err := json.Marshal(location)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
