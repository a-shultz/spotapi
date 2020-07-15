package response

import (
	"net/url"
	"time"
)

type TrackFullResponse struct{}

type TrackFull struct {
	album            AlbumFull
	artists          []ArtistSimple
	availableMarkets []string
	discNumber       int
	duration         time.Duration
	explicit         bool
	externalIDs      map[string]string
	href             url.URL
	isPlayable       bool
	linkedFrom       TrackLink
	restrictions     map[string]string
	name             string
	popularity       uint8
	previewURL       url.URL
	trackNumber      int
	uri              URI
	isLocal          bool
}

type TrackSimpleResponse struct {
	Artists          []ArtistSimpleResponse `json:"artists"`
	AvailableMarkets []string               `json:"available_markets"`
	DiscNumber       int                    `json:"disc_number"`
	DurationMS       int                    `json:"duration_ms"`
	Explicit         bool                   `json:"explicit"`
	ExternalURLs     map[string]string      `json:"external_urls"`
	Href             string                 `json:"href"`
	ID               string                 `json:"id"`
	IsPlayable       bool                   `json:"is_playable"`
	LinkedFrom       TrackLinkResponse      `json:"linked_from"`
	Restrictions     map[string]string      `json:"restrictions"`
	Name             string                 `json:"name"`
	PreviewURL       string                 `json:"preview_url"`
	TrackNumber      int                    `json:"track_number"`
	Type             string                 `json:"type"`
	URI              string                 `json:"uri"`
	IsLocal          bool                   `json:"is_local"`
}

type TrackSimple struct {
	artists          []ArtistSimple
	availableMarkets []string
	discNumber       int
	duration         time.Duration
	explicit         bool
	externalIDs      map[string]string
	href             url.URL
	isPlayable       bool
	linkedFrom       TrackLink
	restrictions     map[string]string
	name             string
	previewURL       url.URL
	trackNumber      int
	uri              URI
	isLocal          bool
}

type TrackPageResponse struct {
	Href     string                `json:"href"`
	Items    []TrackSimpleResponse `json:"items"`
	Limit    int                   `json:"limit"`
	Next     string                `json:"next"`
	Offset   int                   `json:"offset"`
	Previous string                `json:"previous"`
	Total    int                   `json:"total"`
}

type TrackPage struct{}

type TrackSavedResponse struct{}

type TrackSaved struct {
	addedAt time.Time
	track   TrackFull
}

type TrackSavedPageResponse struct{}

type TrackSavedPage struct{}

type TrackPlaylistResponse struct{}

type TrackPlaylist struct {
	addedAt time.Time
	addedBy UserPublic
	isLocal bool
	track   interface{}
}

type TrackLinkResponse struct {
	ExternalURLs map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type TrackLink struct {
	externalURLs map[string]*url.URL
	href         url.URL
	uri          URI
}

type TrackAudioFeaturesResponse struct{}

type TrackAudioFeatures struct {
	acousticness     float32
	analysisURL      url.URL
	danceability     float32
	duration         time.Duration
	energy           float32
	instrumentalness float32
	key              int
	liveness         float32
	loudness         float32
	mode             int
	speechiness      float32
	tempo            float32
	timeSignature    int
	href             url.URL
	uri              URI
	valence          float32
}

type TrackTuneableResponse struct{}

type TrackTuneable struct {
	acousticness     float32
	danceability     float32
	duration         time.Duration
	energy           float32
	instrumentalness float32
	key              int
	liveness         float32
	loudness         float32
	mode             int
	popularity       float32
	speechiness      float32
	tempo            float32
	timeSignature    int
	valence          float32
}
