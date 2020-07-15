package response

import (
	"net/url"
)

type CategoryResponse struct{}

type Category struct {
	href  url.URL
	icons []Image
	id    string
	name  string
}

type CategoryPageResponse struct{}

type CategoryPage struct{}
