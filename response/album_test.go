package response_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"spotapi/response"
)

var _ = Describe("Full Album Response", func() {
	var (
		json  []byte
		album *response.AlbumFullResponse
		err   error
	)

	BeforeEach(func() {
		json = ReadJSON("./testJSON/validAlbum.json")
	})

	JustBeforeEach(func() {
		album, err = response.NewAlbumFullResponse(json)
	})

	Context("when a Full Album Response is created", func() {
		Context("with valid json", func() {
			Specify("a AlbumFullResponse pointer is returned", func() {
				Expect(album).ToNot(BeNil())
				Expect(album).To(BeAssignableToTypeOf(&response.AlbumFullResponse{}))
			})

			Specify("no error occurs", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})

var _ = Describe("Full Album", func() {
	var (
		json  []byte
		album *response.AlbumFull
		err   error
	)

	BeforeEach(func() {
		json = ReadJSON("./testJSON/validAlbum.json")
	})

	JustBeforeEach(func() {
		album, err = response.NewAlbumFull(json)
	})

	Context("when a Full Album is created", func() {
		Context("with valid JSON", func() {
			Specify("a AlbumFull pointer is returned", func() {
				Expect(album).ToNot(BeNil())
				Expect(album).To(BeAssignableToTypeOf(&response.AlbumFull{}))
			})

			Specify("no error occurs", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
