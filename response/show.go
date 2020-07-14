package response

import (
	"net/url"
	"time"
)

type FullShow struct {
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

type SimpleShow struct {
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

type SavedShow struct {
	addedAt time.Time
	show    FullShow
}
