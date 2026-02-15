package main

import (
	"fmt"
	"net/url"
)

/*

In Go, The net/url package provides functionalities to parse, construct, and manipulate URLs (Uniform Resource Locators). URLs are used to identify resources on the internet, such as web pages, images, and files.
Here's a beginner-friendly explanation of some key concepts in the net/url package:
Parsing URLs: The url. Parse function is used to parse a string into a URL object (url. URL struct). This allows you to break down a URL into its individual components, such as scheme, host, path, and query parameters.
Accessing URL Components: Once you have a URL object, you can access its components using various fields:
Scheme: Indicates the protocol used (e.g., "http", "https").
Host: Specifies the domain name and optionally the port number.
D Path: Represents the path component of the URL, which specifies the resource's location on the server.
RawQuery: Contains the raw query string, including query parameters.
Query Parameters: Query parameters are key-value pairs appended to the end of a URL, usually starting with a ? and separated by &. The url. Values type is used to represent query parameters as a map. You can access and manipulate query
parameters using methods like Get, Set, and Add.

*/

func main() {
	fmt.Println("Learning URL")
	myURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("Type of URL: %T\n", myURL)

	parsedURL, err := url.Parse(myURL)

	if err != nil {
		fmt.Println("Can't parse URL ", err)
		return
	}
	fmt.Printf("Type of URL: %T\n", parsedURL)

	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)

}
