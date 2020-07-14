package response

import (
	"net/url"
	"time"
)

type SimplifiedTrackResponse struct {
	Artists          []SimplifiedArtistResponse `json:"artists"`
	AvailableMarkets []string                   `json:"available_markets"`
	DiscNumber       int                        `json:"disc_number"`
	DurationMS       int                        `json:"duration_ms"`
	Explicit         bool                       `json:"explicit"`
	ExternalURLs     map[string]string          `json:"external_urls"`
	Href             string                     `json:"href"`
	ID               string                     `json:"id"`
	IsPlayable       bool                       `json:"is_playable"`
	LinkedFrom       LinkedTrackResponse        `json:"linked_from"`
	Restrictions     map[string]string          `json:"restrictions"`
	Name             string                     `json:"name"`
	PreviewURL       string                     `json:"preview_url"`
	TrackNumber      int                        `json:"track_number"`
	Type             string                     `json:"type"`
	URI              string                     `json:"uri"`
	IsLocal          bool                       `json:"is_local"`
}

type LinkedTrackResponse struct {
	ExternalURLs map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type TuneableTrack struct {
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

type AudioFeatures struct {
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

type PlaylistTrack struct {
	addedAt time.Time
	addedBy PublicUser
	isLocal bool
	track   interface{}
}

type SavedTrack struct {
	addedAt time.Time
	track   TrackFull
}

type TrackFull struct {
	album            FullAlbum
	artists          []SimpleArtist
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

type TrackSimple struct {
	artists          []SimpleArtist
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

type TrackLink struct {
	externalURLs map[string]*url.URL
	href         url.URL
	uri          URI
}
