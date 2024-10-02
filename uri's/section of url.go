// Assignment
// Let's use the url.URL struct again. This time, we'll parse a URL string and add all the different parts to a custom ParsedURL struct. We'll learn more about each part later.

// Complete the newParsedURL function. It should return a ParsedURL with all the parts of a URL. For example, given this URL:

// http://testuser:testpass@testdomain.com:8080/testpath?testsearch=testvalue#testhash

// newParsedURL should return a struct with these fields:

// return ParsedURL{
// 	protocol: "http",
// 	username: "testuser",
// 	password: "testpass",
// 	hostname: "testdomain.com",
// 	port:     "8080",
// 	pathname: "/testpath",
// 	search:   "testsearch=testvalue",
// 	hash:     "testhash",
// }
// Copy icon
// Use these fields and methods of the URL struct:

// Scheme
// User.Username()
// User.Password() - this method returns two values: the password if set and whether it was set
// Hostname()
// Port()
// Path
// RawQuery
// Fragment


package main

import (
	"net/url"
)

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	password := ""
	if pw, hasPassword := parsedUrl.User.Password(); hasPassword {
		password = pw
	}

	return ParsedURL{
		protocol: parsedUrl.Scheme,
		username: parsedUrl.User.Username(),
		password: password,
		hostname: parsedUrl.Hostname(),
		port:     parsedUrl.Port(),
		pathname: parsedUrl.Path,
		search:   parsedUrl.RawQuery,
		hash:     parsedUrl.Fragment,
	}
}
