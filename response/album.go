package response

import (
	"encoding/json"
	"net/url"
	"time"
)

type FullAlbumResponse struct {
	AlbumType            string                     `json:"album_type"`
	Artists              []SimplifiedArtistResponse `json:"artists"`
	AvailableMarkets     []string                   `json:"available_markets"`
	Copyrights           []CopyrightResponse        `json:"copyrights"`
	ExternalIDs          map[string]string          `json:"external_ids"`
	ExternalURLs         map[string]string          `json:"external_urls"`
	Genres               []string                   `json:"genres"`
	Href                 string                     `json:"href"`
	ID                   string                     `json:"id"`
	Images               []ImageResponse            `json:"images"`
	Label                string                     `json:"label"`
	Name                 string                     `json:"name"`
	Popularity           int                        `json:"popularity"`
	ReleaseDate          string                     `json:"release_date"`
	ReleaseDatePrecision string                     `json:"release_date_precision"`
	Restrictions         map[string]string          `json:"restrictions"`
	Tracks               TracksPageResponse         `json:"tracks"`
	Type                 string                     `json:"type"`
	URI                  string                     `json:"uri"`
}

type FullAlbum struct {
	albumType        string
	artists          []SimpleArtist
	availableMarkets []string
	copyrights       []Copyright
	externalID       map[string]string
	externalURLs     map[string]*url.URL
	genres           []string
	href             url.URL
	images           []Image
	label            string
	name             string
	popularity       uint8
	releaseDate      time.Time
	restrictions     map[string]string
	tracks           Page
	uri              URI
}

func NewFullAlbumResponse(data []byte) (*FullAlbumResponse, error) {
	var fullAlbumResponse = new(FullAlbumResponse)
	err := json.Unmarshal(data, &fullAlbumResponse)
	if err != nil {
		return nil, err
	}
	return fullAlbumResponse, nil
}

func NewFullAlbum(data []byte) (*FullAlbum, error) {
	response, err := NewFullAlbumResponse(data)
	if err != nil {
		return nil, err
	}
	return &FullAlbum{
		albumType: response.AlbumType,
	}, nil
}

type SimplifiedAlbumResponse struct {
	AlbumGroup string                  `json:"album_group"`
	AlbumType string                   `json:"album_type"`
	Artists []SimplifiedArtistResponse `json:"artists"`
	AvailableMarkets []string          `json:"availableMarkets"`
	ExternalURLs map[string]string     `json:"external_urls"`
	Href string                        `json:"href"`
	ID string                          `json:"id"`
	Images []ImageResponse             `json:"images"`
	Name string                        `json:"name"`
	ReleaseDate string                 `json:"release_date"`
	ReleaseDatePrecision string        `json:"release_date_precision"`
	Restrictions map[string]string     `json:"restrictions"`
	Type string                        `json:"type"`
	URI string                         `json:"uri"`
}

type SimplifiedAlbum struct {
	albumGroup       string
	albumType        string
	artists          []SimpleArtist
	availableMarkets []string
	externalURLs     map[string]*url.URL
	genres           []string
	href             url.URL
	images           []Image
	name             string
	releaseDate      time.Time
	restrictions     map[string]string
	uri              URI
}

type SavedAlbum struct {
	addedAt time.Time
	album   FullAlbum
}
