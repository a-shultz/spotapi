package response

import "net/url"

type PlaylistFullResponse struct{}

type PlaylistFull struct {
	collaborative bool
	description   string
	externalURLs  map[string]url.URL
	followers     Followers
	href          url.URL
	images        []Image
	name          string
	owner         UserPublic
	public        bool
	snapshotId    string
	tracks        Page
	uri           URI
}

type PlaylistSimpleResponse struct{}

type PlaylistSimple struct {
	collaborative bool
	description   string
	externalURLs  map[string]url.URL
	href          url.URL
	images        []Image
	name          string
	owner         UserPublic
	public        bool
	snapshotId    string
	tracks        Page
	uri           URI
}

type PlaylistPageResponse struct{}

type PlaylistPage struct{}

type PlaylistItemResponse struct{}

type PlaylistItem struct{}

type PlaylistItemPageResponse struct{}

type PlaylistItemPage struct{}
