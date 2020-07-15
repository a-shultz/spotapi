package response

import (
	"net/url"
	"time"
)

type EpisodeFullResponse struct{}

type EpisodeFull struct {
	audioPreviewURL      url.URL
	description          string
	duration             time.Duration
	explicit             bool
	externalURLs         map[string]url.URL
	href                 url.URL
	images               []Image
	isExternallyHosted   bool
	isPlayable           bool
	languages            []string
	name                 string
	releaseDate          time.Time
	releaseDatePrecision string
	resumePoint          PlayResumePoint
	show                 ShowSimple
	uri                  URI
}

type EpisodeSimpleResponse struct{}

type EpisodeSimple struct {
	audioPreviewURL      url.URL
	description          string
	duration             time.Duration
	explicit             bool
	externalURLs         map[string]url.URL
	href                 url.URL
	images               []Image
	isExternallyHosted   bool
	isPlayable           bool
	languages            []string
	name                 string
	releaseDate          time.Time
	releaseDatePrecision string
	resumePoint          PlayResumePoint
	uri                  URI
}

type EpisodePageResponse struct{}

type EpisodePage struct{}
