package response

import (
	"net/url"
)

type Category struct {
	href url.URL
	icons []Image
	id string
	name string
}
