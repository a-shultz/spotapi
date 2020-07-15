package response

import (
	"net/url"
)

type ContextResponse struct{}

type Context struct {
	objectType   string
	href         url.URL
	externalURLs map[string]*url.URL
	uri          URI
}
