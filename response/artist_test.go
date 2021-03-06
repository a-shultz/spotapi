package response_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"spotapi/response"
)

var _ = Describe("Artist Simple Response", func() {
	var (
		json   []byte
		artist *response.ArtistSimpleResponse
		err    error
	)

	JustBeforeEach(func() {
		artist, err = response.NewArtistSimpleResponse(json)
	})

	Context("when an Artist Simple Response is created", func() {
		Context("with valid JSON", func() {
			BeforeEach(func() {
				json = ReadJSON("./testJSON/artist/simple/valid/valid.json")
			})

			Specify("an Artist Simple Response pointer is returned", func() {
				Expect(artist).ToNot(BeNil())
				Expect(artist).To(BeAssignableToTypeOf(&response.ArtistSimpleResponse{}))
			})

			Specify("no error occurs", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("with invalid JSON", func() {
			BeforeEach(func() {
				json = ReadJSON("./testJSON/artist/simple/invalid/invalid.json")
			})

			Specify("a nil pointer is returned", func() {
				Expect(artist).To(BeNil())
			})

			Specify("an error occurs", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
