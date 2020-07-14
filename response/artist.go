package response

import (
	"net/url"
)

type FullArtist struct {
	externalURLs map[string]*url.URL
	followers    Followers
	genres       []string
	href         url.URL
	images       []Image
	name         string
	popularity   uint8
	uri          URI
}

type SimplifiedArtistResponse struct {
	ExternalURLs map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type SimpleArtist struct {
	externalURLs map[string]*url.URL
	href         url.URL
	name         string
	uri          URI
}
