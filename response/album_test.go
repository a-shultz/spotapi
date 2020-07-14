package response_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"spotapi/response"
)

var _ = Describe("Full Album Response", func() {
	var (
		json  []byte
		album *response.FullAlbumResponse
		err   error
	)

	BeforeEach(func() {
		json = ReadJSON("./testJSON/validAlbum.json")
	})

	JustBeforeEach(func() {
		album, err = response.NewFullAlbumResponse(json)
	})

	Context("when a Full Album Response is created", func() {
		Context("with valid json", func() {
			Specify("a FullAlbumResponse pointer is returned", func() {
				Expect(album).ToNot(BeNil())
				Expect(album).To(BeAssignableToTypeOf(&response.FullAlbumResponse{}))
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
		album *response.FullAlbum
		err   error
	)

	BeforeEach(func() {
		json = ReadJSON("./testJSON/validAlbum.json")
	})

	JustBeforeEach(func() {
		album, err = response.NewFullAlbum(json)
	})

	Context("when a Full Album is created", func() {
		Context("with valid JSON", func() {
			Specify("a FullAlbum pointer is returned", func() {
				Expect(album).ToNot(BeNil())
				Expect(album).To(BeAssignableToTypeOf(&response.FullAlbum{}))
			})

			Specify("no error occurs", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
