package response_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"spotapi/response"
)

var _ = Describe("Album", func() {

	Describe("FullAlbum", func() {
		var(
			json []byte
		)

		BeforeEach(func() {
			json = ReadJSON("./testJSON/validAlbum.json")
		})

		Describe("FullAlbumResponse", func() {
			Describe("calling NewFullAlbumResponse()", func() {
				Context("with valid JSON", func() {
					var (
						fullAlbumResponse *response.FullAlbumResponse
						err error
					)

					BeforeEach(func() {
						fullAlbumResponse, err = response.NewFullAlbumResponse(json)
					})

					Specify("returns a FullAlbumResponse pointer", func() {
						Expect(fullAlbumResponse).ToNot(BeNil())
						Expect(fullAlbumResponse).To(BeAssignableToTypeOf(&response.FullAlbumResponse{}))
					})

					Specify("does not error", func() {
						Expect(err).ToNot(HaveOccurred())
					})
				})
			})
		})

		Describe("calling NewFullAlbum()", func() {
			Context("with valid JSON", func() {
				var (
					fullAlbum *response.FullAlbum
					err error
				)

				BeforeEach(func() {
					fullAlbum, err = response.NewFullAlbum(json)
				})

				Specify("returns a FullAlbum pointer", func() {
					Expect(fullAlbum).ToNot(BeNil())
					Expect(fullAlbum).To(BeAssignableToTypeOf(&response.FullAlbum{}))
				})

				Specify("does not return an error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})
		})
	})

	Describe("SimplifiedAlbum", func() {

	})
})
