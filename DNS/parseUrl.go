// Assignment
// Complete the getDomainNameFromURL function. Given a full URL, it should return the domain (or host) name. Simply return any potential errors.

package main

import (
	"net/url"
	"fmt"
)

func getDomainNameFromURL(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid address") 
	}

	return parsedUrl.Hostname(), nil 
}
