// Assignment
// Our game designer wants us to build a "treasure chest"! We don't want to award players with too many items though, so let's limit the number of items our players receive by using the limit=x query parameter, where x is the number of records we are requesting from the server. We'll also continue to sort the items so that less valuable items are looted first.

// Complete the lootTreasure function. It should add two query parameters to the URL passed to getItems:

// sort
// limit
// Sort by the quality property.

// Set the limit based on the chestRarity.

// Common treasure = 1 item
// Rare treasure = 3 items
// Legendary treasure = 5 items

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type Item struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Quality int    `json:"quality"`
}

func getItems(url string) []Item {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	var items []Item
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}

	return items
}

func logItems(items []Item) string {
	log := ""
	for _, item := range items {
		log += fmt.Sprintf("- %s with quality score: %d\n", item.Name, item.Quality)
	}
	return log
}

func Test(t *testing.T) {
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/items"

	type testCase struct {
		rarity   string
		expected []Item
	}

	tests := []testCase{
		{
			rarity: "Common",
			expected: []Item{
				{Name: "Light Leather", Quality: 1},
			},
		},
		{
			rarity: "Rare",
			expected: []Item{
				{Name: "Light Leather", Quality: 1},
				{Name: "Gold Ore", Quality: 1},
				{Name: "Healing Potion", Quality: 4},
			},
		},
	}

	if withSubmit {
		tests = append(tests, testCase{
			rarity: "Legendary",
			expected: []Item{
				{Name: "Light Leather", Quality: 1},
				{Name: "Gold Ore", Quality: 1},
				{Name: "Healing Potion", Quality: 4},
				{Name: "Padded Leather", Quality: 6},
				{Name: "Copper Ore", Quality: 6},
			},
		})
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		items := lootTreasure(url, test.rarity)

		passed := true
		for i, expectedItem := range test.expected {
			if i >= len(items) {
				passed = false
				break
			}

			if expectedItem.Name != items[i].Name || expectedItem.Quality != items[i].Quality {
				passed = false
				break
			}
		}

		if passed {
			passCount++
			fmt.Printf(`---------------------------------
Rarity:     %s
Expected:
%v
Actual:
%v
PASS
`, test.rarity, logItems(test.expected), logItems(items))
		} else {
			failCount++
			t.Errorf(`---------------------------------
Rarity:     %s
Expected:
%v
Actual:
%v
FAIL`, test.rarity, logItems(test.expected), logItems(items))
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

func lootTreasure(baseURL, chestRarity string) []Item {
	sort := "?sort=quality"
	var limit string

	if chestRarity == "Common" {
		limit = "&limit=1"
	} else if chestRarity == "Rare" {
		limit = "&limit=3"
	} else {
		limit = "&limit=5"
	}

	fullURL := baseURL + sort + limit
	return getItems(fullURL)
}
