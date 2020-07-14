package response

import "net/url"

type FullPlaylist struct {
	collaborative bool
	description string
	externalURLs map[string]url.URL
	followers Followers
	href url.URL
	images []Image
	name string
	owner PublicUser
	public bool
	snapshotId string
	tracks Page
	uri URI
}

type SimplePlaylist struct {
	collaborative bool
	description string
	externalURLs map[string]url.URL
	href url.URL
	images []Image
	name string
	owner PublicUser
	public bool
	snapshotId string
	tracks Page
	uri URI
}
