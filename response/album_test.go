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

	JustBeforeEach(func() {
		album, err = response.NewAlbumFullResponse(json)
	})

	Context("when a Full Album Response is created", func() {
		Context("with valid json", func() {
			BeforeEach(func() {
				json = ReadJSON("./testJSON/album/full/valid/valid.json")
			})

			Specify("a AlbumFullResponse pointer is returned", func() {
				Expect(album).ToNot(BeNil())
				Expect(album).To(BeAssignableToTypeOf(&response.AlbumFullResponse{}))
			})

			Specify("no error occurs", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
		Context("with invalid json", func() {
			BeforeEach(func() {
				json = ReadJSON("./testJSON/album/full/invalid/invalid.json")
			})

			Specify("a nil pointer is returned", func() {
				Expect(album).To(BeNil())
			})

			Specify("an error occurs", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

var _ = Describe("Full Album", func() {
	var (
		json          []byte
		albumResponse *response.AlbumFullResponse
		album         *response.AlbumFull
		err           error
	)

	BeforeEach(func() {
		json = ReadJSON("./testJSON/album/full/valid/valid.json")
	})

	JustBeforeEach(func() {
		albumResponse, err = response.NewAlbumFullResponse(json)
		Expect(albumResponse).ToNot(BeNil())
		Expect(err).ToNot(HaveOccurred())
		album = response.NewAlbumFull(albumResponse)
	})

	Context("when a Full Album is created", func() {
		Specify("a AlbumFull pointer is returned", func() {
			Expect(album).ToNot(BeNil())
			Expect(album).To(BeAssignableToTypeOf(&response.AlbumFull{}))
		})

		Specify("no error occurs", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
