package response

import "net/url"

type Followers struct {
	href url.URL
	total int
}