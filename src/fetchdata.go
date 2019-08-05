package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	resty "github.com/go-resty/resty"
)

/**
 * Fetch the posts and save to SQS
 * @todo - does too much. break this up.
 **/
func fetchAndSave(pageToFetch int, siteURL string) int {

	wpSiteURL := siteURL + "?page=" + strconv.Itoa(pageToFetch)
	fmt.Println("page to fetch: ", wpSiteURL)

	// Set up the client.
	client := resty.New()

	// Fetch the data from wp.
	posts := []Post{}
	resp, err := client.R().SetResult(&posts).Get(wpSiteURL)

	if err != nil {
		fmt.Println("Error", err)
		return 0
	}

	if resp.StatusCode() != 200 {
		fmt.Println("Error", "Status Code not 200")
		return 0
	}

	// For each entry save it after transforming it into the new payload
	// and turning it into a JSON string.
	for i := 0; i < len(posts); i++ {

		media := Media{
			FeaturedMedia: posts[i].FeaturedMedia,
			WPMediaLink:   posts[i].Links.Attachments,
		}

		payload := Payload{
			ID:           posts[i].ID,
			Date:         posts[i].Date,
			LastModified: posts[i].Modified,
			Title:        posts[i].Title.Rendered,
			Content:      posts[i].Content.Rendered,
			Media:        media,
		}

		var jsonData []byte
		jsonData, err := json.Marshal(payload)

		if err != nil {
			fmt.Println("Error", err)
		} else {
			// Send over to SQS
			saveToSQS(string(jsonData))
		}
	}

	// Get the number of pages we need to iterate over
	headers := resp.Header()
	totalPages, err := strconv.Atoi(headers["X-Wp-Totalpages"][0])

	return totalPages
}
