// Package browser provides utilities for interacting with users' browsers.
package browser

import "fmt"

// OpenURLs opens the given URLs in the default web browser.
func OpenURLs(urls ...string) bool {
	for _, url := range urls {
		fmt.Println(url)
	}

	return true
}
