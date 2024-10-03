// getResources
// getResources takes a url string and returns a slice of maps []map[string]interface{} and an error.

// Decode the response body into a slice of maps []map[string]interface{} and return it.
// logResources
// logResources takes a slice of maps []map[string]interface{} and prints its keys and values to the console. Because maps are unsorted we will be adding formatted strings to a slice of strings []string which is then sorted.

// Iterate over the slice of map[string]interface{}
// For each map[string]interface{} get its keys and values using range and append it to formattedStrings as Key: %s - Value: %v, where %s is the key and %v is the value.


package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resources)
	if err != nil {
		return []map[string]any{}, err
	}

	return resources, nil
}

func logResources(resources []map[string]any) {
	var formattedStrings []string

	for _, resource := range resources {
		for key, val := range resource {
			formattedStrings = append(formattedStrings, fmt.Sprintf("Key: %s - Value: %v", key, val))
		}
	}

	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}
