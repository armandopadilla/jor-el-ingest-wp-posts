package main

// Title - title.rendered
type Title struct {
	Rendered string `json:rendered`
}

// Content - content.rendered
type Content struct {
	Rendered string `json:rendered`
}

// Link -_links.wp:attachments.href
type Link struct {
	Attachments []struct {
		Href string `json:"href"`
	} `json:"wp:attachment"`
}

// Post - object for WP Object
type Post struct {
	ID            int     `json:"id"`
	Title         Title   `json:"title`
	Content       Content `json:content`
	Date          string  `json:"date"`
	Modified      string  `json:"modified"`
	FeaturedMedia int     `json:"featured_media"`
	Links         Link    `json:"_links"`
}

// Media - Used in new JSON
type Media struct {
	FeaturedMedia int
	WPMediaLink   []struct {
		Href string `json:"href"`
	}
}

// Payload - New payload created from the items we only care about.
type Payload struct {
	ID           int
	Date         string
	LastModified string
	Title        string
	Content      string
	Media        Media
}
