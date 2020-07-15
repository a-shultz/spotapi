package response

import (
	"net/url"
	"time"
)

type ShowFullResponse struct{}

type ShowFull struct {
	availableMarkets   []string
	copyrights         []Copyright
	description        string
	explicit           bool
	episodes           Page
	externalURLs       map[string]*url.URL
	href               url.URL
	images             []Image
	isExternallyHosted bool
	languages          []string
	mediaType          string
	name               string
	publisher          string
	uri                URI
}

type ShowSimpleResponse struct{}

type ShowSimple struct {
	availableMarkets   []string
	copyrights         []Copyright
	description        string
	explicit           bool
	externalURLs       map[string]*url.URL
	href               url.URL
	images             []Image
	isExternallyHosted bool
	languages          []string
	mediaType          string
	name               string
	publisher          string
	uri                URI
}

type ShowSavedResponse struct{}

type ShowSaved struct {
	addedAt time.Time
	show    ShowFull
}

type ShowSavedPageResponse struct{}

type ShowSavedPage struct{}
