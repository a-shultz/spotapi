package response

import (
	"encoding/json"
	"net/url"
	"time"
)

type AlbumFullResponse struct {
	AlbumType            string                 `json:"album_type"`
	Artists              []ArtistSimpleResponse `json:"artists"`
	AvailableMarkets     []string               `json:"available_markets"`
	Copyrights           []CopyrightResponse    `json:"copyrights"`
	ExternalIDs          map[string]string      `json:"external_ids"`
	ExternalURLs         map[string]string      `json:"external_urls"`
	Genres               []string               `json:"genres"`
	Href                 string                 `json:"href"`
	ID                   string                 `json:"id"`
	Images               []ImageResponse        `json:"images"`
	Label                string                 `json:"label"`
	Name                 string                 `json:"name"`
	Popularity           int                    `json:"popularity"`
	ReleaseDate          string                 `json:"release_date"`
	ReleaseDatePrecision string                 `json:"release_date_precision"`
	Restrictions         map[string]string      `json:"restrictions"`
	Tracks               TrackPageResponse      `json:"tracks"`
	Type                 string                 `json:"type"`
	URI                  string                 `json:"uri"`
}

func NewAlbumFullResponse(data []byte) (*AlbumFullResponse, error) {
	var fullAlbumResponse = new(AlbumFullResponse)
	err := json.Unmarshal(data, &fullAlbumResponse)
	if err != nil {
		return nil, err
	}
	return fullAlbumResponse, nil
}

type AlbumFull struct {
	albumType        string
	artists          []ArtistSimple
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

func NewAlbumFull(response *AlbumFullResponse) *AlbumFull {
	return &AlbumFull{
		albumType: response.AlbumType,
	}
}

type AlbumSimpleResponse struct {
	AlbumGroup           string                 `json:"album_group"`
	AlbumType            string                 `json:"album_type"`
	Artists              []ArtistSimpleResponse `json:"artists"`
	AvailableMarkets     []string               `json:"availableMarkets"`
	ExternalURLs         map[string]string      `json:"external_urls"`
	Href                 string                 `json:"href"`
	ID                   string                 `json:"id"`
	Images               []ImageResponse        `json:"images"`
	Name                 string                 `json:"name"`
	ReleaseDate          string                 `json:"release_date"`
	ReleaseDatePrecision string                 `json:"release_date_precision"`
	Restrictions         map[string]string      `json:"restrictions"`
	Type                 string                 `json:"type"`
	URI                  string                 `json:"uri"`
}

func NewAlbumSimpleResponse(data []byte) (*AlbumSimpleResponse, error) {
	album := &AlbumSimpleResponse{}
	err := json.Unmarshal(data, &album)
	if err != nil {
		return nil, err
	}
	return album, nil
}

type AlbumSimple struct {
	albumGroup       string
	albumType        string
	artists          []ArtistSimple
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

type AlbumPageResponse struct{}

type AlbumPage struct{}

type AlbumSavedResponse struct{}

type AlbumSaved struct {
	addedAt time.Time
	album   AlbumFull
}

type AlbumSavedPageResponse struct{}

type AlbumSavedPage struct{}
