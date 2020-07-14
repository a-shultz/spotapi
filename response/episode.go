package response

import (
	"net/url"
	"time"
)

type EpisodeFull struct {
	audioPreviewURL url.URL
	description string
	duration time.Duration
	explicit bool
	externalURLs map[string]url.URL
	href url.URL
	images []Image
	isExternallyHosted bool
	isPlayable bool
	languages []string
	name string
	releaseDate time.Time
	releaseDatePrecision string
	resumePoint ResumePoint
	show SimpleShow
	uri URI
}

type EpisodeSimple struct {
	audioPreviewURL url.URL
	description string
	duration time.Duration
	explicit bool
	externalURLs map[string]url.URL
	href url.URL
	images []Image
	isExternallyHosted bool
	isPlayable bool
	languages []string
	name string
	releaseDate time.Time
	releaseDatePrecision string
	resumePoint ResumePoint
	uri URI
}
