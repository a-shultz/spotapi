package response

import (
	"net/url"
)

type UserPrivateResponse struct{}

type UserPrivate struct {
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

type UserPublicResponse struct{}

type UserPublic struct {
	displayName  string
	externalURLs map[string]*url.URL
	followers    Followers
	href         url.URL
	images       []Image
	uri          URI
}
