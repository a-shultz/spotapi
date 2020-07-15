package response

import (
	"encoding/json"
	"net/url"
)

type ImageResponse struct {
	Height int    `json:"height,int"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

func NewImageResponse(data []byte) (*ImageResponse, error) {
	image := &ImageResponse{}
	err := json.Unmarshal(data, &image)
	if err != nil {
		return nil, err
	}
	return image, nil
}

type Image struct {
	height int
	url    url.URL
	width  int
}
