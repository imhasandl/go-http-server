// Assignment
// I've updated the getItems function from before. Now it accepts just a domain as input. It's convenient this way because it means if the API we're using ever changes its domain, we only need to update one variable.

// Problem is, there is a bug. The API isn't hosted on boot.dev, it's hosted on the "api" subdomain! Fix the bug.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

type Item struct {
	Name string
}

func getItems(domain string) ([]Item, error) {
	res, err := http.Get("https://" + domain + "/v1/courses_rest_api/learn-http/items")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var items []Item
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func logItems(items []Item) {
	for _, item := range items {
		fmt.Println(item.Name)
	}
}

const domain = "api.boot.dev"


func main() {
	items, err := getItems(domain)
	if err != nil {
		log.Fatalf("error getting item data: %v", err)
	}
	logItems(items)
}
