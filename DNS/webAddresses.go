// Assignment
// Cloudflare is a tech company that provides a cool HTTP server that we can use to look up the IP address of any domain.

// I've provided a getIPAddress function that makes a request to Cloudflare. The function takes a domain name as input and returns the IP address associated with that domain.

// The function currently prints a string representation of the entire response we receive from Cloudflare.

// Run the code to see the structure of the response and what fields it contains.
// Import the "encoding/json" package and unmarshal the response as you have done before.
// Update the code to return just the first IP address found within (if it exists).
// If there is no IP address, return an empty string and an error.
// I've provided a DNSResponse struct in dns.go you might find useful.


package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	DNSResponse struct {
		Status   int        `json:"Status"`
		Tc       bool       `json:"TC"`
		Rd       bool       `json:"RD"`
		Ra       bool       `json:"RA"`
		Ad       bool       `json:"AD"`
		Cd       bool       `json:"CD"`
		Question []Question `json:"Question"`
		Answer   []Answer   `json:"Answer"`
	}
	Question struct {
		Name string `json:"name"`
		Type int    `json:"type"`
	}
	Answer struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	}
)



func getIPAddress(domain string) (string, error) {
	url := fmt.Sprintf("https://cloudflare-dns.com/dns-query?name=%s&type=A", domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var dnsRes DNSResponse
	if err := json.Unmarshal(body, &dnsRes); err != nil {
		return "", fmt.Errorf("error unmarshalling json: %w", err)
	}

	if len(dnsRes.Answer) == 0 {
		return "", fmt.Errorf("no answer found")
	}

	return dnsRes.Answer[0].Data, nil
	
}
