package response_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"spotapi/response"
)

var _ = Describe("URI", func() {
	var (
		resourceType string
		id           string
		uri          *response.URI
		err          error
	)

	BeforeEach(func() {
		resourceType = "track"
		id = "6rqhFgbbKwnb9MLmUQDhG6"
	})

	JustBeforeEach(func() {
		uri, err = response.NewURI(resourceType, id)
	})

	Context("when a URI is created", func() {
		Context("with valid input", func() {
			Specify("a URI pointer is returned", func() {
				Expect(uri).ToNot(BeNil())
				Expect(uri).To(BeAssignableToTypeOf(&response.URI{}))
			})

			Specify("no error occurs", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("has a Spotify ID", func() {
				Expect(uri.ID()).To(Equal("6rqhFgbbKwnb9MLmUQDhG6"))
			})

			It("has a resource type", func() {
				Expect(uri.ResourceType()).To(Equal("track"))
			})

			It("has a Spotify URI", func() {
				Expect(uri.URI()).To(Equal("spotify:track:6rqhFgbbKwnb9MLmUQDhG6"))
			})

			It("has a Spotify API URL", func() {
				Expect(uri.APIURL()).To(Equal("https://api.spotify.com/v1/tracks/6rqhFgbbKwnb9MLmUQDhG6"))
			})

			It("has a Spotify Open URL", func() {
				Expect(uri.OpenURL()).To(Equal("https://open.spotify.com/track/6rqhFgbbKwnb9MLmUQDhG6"))
			})

			It("has a Spotify Play URL", func() {
				Expect(uri.PlayURL()).To(Equal("https://play.spotify.com/track/6rqhFgbbKwnb9MLmUQDhG6"))
			})
		})

		Context("with invalid input", func() {
			Context("with an ID that is not 22 characters", func() {
				BeforeEach(func() {
					id = "aaBB00"
				})

				Specify("a nil pointer is returned", func() {
					Expect(uri).To(BeNil())
				})

				Specify("an error occurs", func() {
					Expect(err).To(HaveOccurred())
				})
			})

			Context("with an ID that is not in base62 encoding", func() {
				BeforeEach(func() {
					id = "aaBB00"
				})

				Specify("a nil pointer is returned", func() {
					Expect(uri).To(BeNil())
				})

				Specify("an error occurs", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
