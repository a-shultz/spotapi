package response

import (
	"net/url"
)

type Private struct {
	country      string
	displayName  string
	email        string
	externalURLs map[string]*url.URL
	followers    Followers
	href         url.URL
	images       []Image
	product      string
	uri          URI
}

type PublicUser struct {
	displayName  string
	externalURLs map[string]*url.URL
	followers    Followers
	href         url.URL
	images       []Image
	uri          URI
}
