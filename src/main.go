package main

import (
	"fmt"
	"os"
)

const (
	wpSiteURL     = "http://www.armando.ws/wp-json/wp/v2/posts"
	fetchAllPages = false // TURNING OFF TO SAVE ON AWS COST - More pages!
)

/**
 * Main driver
 **/
func main() {

	// Initial request to get the first page and check if we need to get more pages.
	totalPages := fetchAndSave(1, wpSiteURL)

	if totalPages == 1 {
		fmt.Println("No more pages found.  Done")
		os.Exit(3) // Done.
	}

	if fetchAllPages {
		for currentPage := 2; currentPage <= totalPages; currentPage++ {
			fetchAndSave(currentPage, wpSiteURL)
		}
	}

	fmt.Println("Done!")
}
