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

	Describe("calling NewURI()", func() {
		Context("with valid input", func() {
			Specify("return a URI pointer", func() {
				Expect(uri).ToNot(BeNil())
				Expect(uri).To(BeAssignableToTypeOf(&response.URI{}))
			})

			Specify("does not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("with invalid input", func() {
			Context("with an ID that is not 22 characters", func() {
				BeforeEach(func() {
					id = "aaBB00"
				})

				Specify("return nil instead of the URI pointer", func() {
					Expect(uri).To(BeNil())
				})

				Specify("return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})

			Context("with an ID that is not in base62 encoding", func() {
				BeforeEach(func() {
					id = "!@#$ abcdefghijklmnopq"
				})

				Specify("return nil instead of the URI pointer", func() {
					Expect(uri).To(BeNil())
				})

				Specify("return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("calling uri.ID()", func() {
		It("returns the Spotify ID", func() {
			Expect(uri.ID()).To(Equal("6rqhFgbbKwnb9MLmUQDhG6"))
		})
	})

	Describe("calling uri.ResourceType()", func() {
		It("returns the resource type", func() {
			Expect(uri.ResourceType()).To(Equal("track"))
		})
	})

	Describe("calling uri.URI()", func() {
		It("returns the Spotify ID for the resource", func() {
			Expect(uri.URI()).To(Equal("spotify:track:6rqhFgbbKwnb9MLmUQDhG6"))
		})
	})

	Describe("calling uri.APIURL()", func() {
		It("returns the URL referencing the location of the resource using the Spotify API", func() {
			Expect(uri.APIURL()).To(Equal("https://api.spotify.com/v1/tracks/6rqhFgbbKwnb9MLmUQDhG6"))
		})
	})

	Describe("calling uri.OpenURL()", func() {
		It("returns the URL for opening the resource on a web browser", func() {
			Expect(uri.OpenURL()).To(Equal("https://open.spotify.com/track/6rqhFgbbKwnb9MLmUQDhG6"))
		})
	})

	Describe("calling uri.PlayURL()", func() {
		It("returns the URL for playing the resource on a web browser", func() {
			Expect(uri.PlayURL()).To(Equal("https://play.spotify.com/track/6rqhFgbbKwnb9MLmUQDhG6"))
		})
	})
})
