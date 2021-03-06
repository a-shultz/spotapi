package response

import (
	"encoding/json"
	"net/url"
)

type ArtistFullResponse struct{}

type ArtistFull struct {
	externalURLs map[string]*url.URL
	followers    Followers
	genres       []string
	href         url.URL
	images       []Image
	name         string
	popularity   uint8
	uri          URI
}

type ArtistSimpleResponse struct {
	ExternalURLs map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

func NewArtistSimpleResponse(data []byte) (*ArtistSimpleResponse, error) {
	artist := &ArtistSimpleResponse{}
	err := json.Unmarshal(data, &artist)
	if err != nil {
		return nil, err
	}
	return artist, nil
}

type ArtistSimple struct {
	externalURLs map[string]*url.URL
	href         url.URL
	name         string
	uri          URI
}

type ArtistPageResponse struct{}

type ArtistPage struct{}
