package response

import (
	"errors"
	"fmt"
)

// A URI represents an identifier and pointer to a Spotify resource.
type URI struct {
	resourceType string
	id           string
}

// NewURI returns a new URI pointer given a resourceType and id.
func NewURI(resourceType string, id string) (*URI, error) {
	if !isSpotifyID(id) {
		return nil, errors.New(fmt.Sprintf("URI creation failed, \"%s\" is not a valid Spotify ID", id))
	}
	return &URI{resourceType, id}, nil
}

// ResourceType returns the resource type for the Spotify resource that the URI
// references.
func (uri URI) ResourceType() string {
	return uri.resourceType
}

// ID returns the Spotify ID for the Spotify resource that the URI
// references. The Spotify ID does not contain the encoded information
// that could identify the resource type.
func (uri URI) ID() string {
	return uri.id
}

// URI returns the formatted Spotify URI for identifying a Spotify resource.
// URI returns the same output as String
func (uri URI) URI() string {
	return fmt.Sprintf("spotify:%s:%s", uri.resourceType, uri.id)
}

// APIURL returns the API Endpoint URL for retrieving the referenced Spotify
// resource.
func (uri URI) APIURL() string {
	return fmt.Sprintf("https://api.spotify.com/v1/%ss/%s", uri.resourceType, uri.id)
}

// OpenURL returns the URL for opening the Spotify resource in a web browser.
func (uri URI) OpenURL() string {
	return fmt.Sprintf("https://open.spotify.com/%s/%s", uri.resourceType, uri.id)
}

// PlayURL returns the URL for playing the Spotify resource in a web browser.
func (uri URI) PlayURL() string {
	return fmt.Sprintf("https://play.spotify.com/%s/%s", uri.resourceType, uri.id)
}

// String returns the formatted Spotify URI for identifying a Spotify resource.
// String returns the same output as URI.
func (uri URI) String() string {
	return uri.URI()
}

func isSpotifyID(id string) bool {
	return len(id) == 22 && isBase62String(id)
}

func isBase62String(s string) bool {
	for _, r := range s {
		if !isBase62Rune(r) {
			return false
		}
	}
	return true
}

func isBase62Rune(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9')
}
