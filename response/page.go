package response

import (
	"net/url"
)

type TracksPageResponse struct {
	Href string                     `json:"href"`
	Items []SimplifiedTrackResponse `json:"items"`
	Limit int                       `json:"limit"`
	Next string                     `json:"next"`
	Offset int                      `json:"offset"`
	Previous string                 `json:"previous"`
	Total int                       `json:"total"`
}

type Page struct {
	href url.URL
	items []interface{}
	limit int
	next url.URL
	offset int
	previous url.URL
	total int
}

type CursorPage struct {
	href    url.URL
	items   []interface{}
	limit   int
	next    url.URL
	cursors Cursor
	total   int
}

type Cursor struct {
	after string
}
