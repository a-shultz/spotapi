package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

type ImageResponse struct {
	Height uint   `json:"height,int"`
	URL    string `json:"url"`
	Width  uint   `json:"width"`
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
	height  uint
	address *url.URL
	width   uint
}

func NewImage(image *ImageResponse) (*Image, error) {
	address, err := url.Parse(image.URL)
	if err != nil {
		return nil, err
	} else if address.Scheme == "" || address.Host == "" {
		return nil, errors.New("invalid URL")
	}
	return &Image{
		height:  image.Height,
		address: address,
		width:   image.Width,
	}, nil
}

func (image Image) Height() int {
	return int(image.height)
}

func (image Image) Width() int {
	return int(image.width)
}

func (image Image) URL() *url.URL {
	return image.address
}

func (image Image) RawURL() string {
	return image.address.String()
}

func (image Image) String() string {
	return fmt.Sprintf("%s (%d x %d)", image.RawURL(), image.Width(), image.Height())
}
