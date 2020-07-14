package response

import "net/url"

type ImageResponse struct {
	Height int `json:"height"`
	URL string `json:"url"`
	Width int `json:"width"`
}

type Image struct {
	height int
	url url.URL
	width int
}
