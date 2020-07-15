package response

import (
	"net/url"
)

type Page struct {
	href     url.URL
	items    []interface{}
	limit    int
	next     url.URL
	offset   int
	previous url.URL
	total    int
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
