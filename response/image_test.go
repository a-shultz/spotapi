package response_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"spotapi/response"
)

var _ = Describe("Image Response", func() {
	var (
		data  []byte
		image *response.ImageResponse
		err   error
	)

	BeforeEach(func() {
		data = []byte(`{
			"height": 640,
			"url": "https://i.scdn.co/image/ab67616d0000b273f142205e336ec0af3e1a4eb0",
			"width": 640
		}`)
	})

	JustBeforeEach(func() {
		image, err = response.NewImageResponse(data)
	})

	Context("when an Image Response is created", func() {
		Context("with valid JSON", func() {
			Specify("an Image Response pointer is returned", func() {
				Expect(image).ToNot(BeNil())
				Expect(image).To(BeAssignableToTypeOf(&response.ImageResponse{}))
			})

			Specify("an error doesn't occur", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("with invalid JSON", func() {
			BeforeEach(func() {
				data = []byte(`{
					"height"": "kinda average",
					"url": "https://i.scdn.co/image/ab67616d0000b273f142205e336ec0af3e1a4eb0",
					"width": 640
				}`)
			})

			Specify("a nil pointer is returned", func() {
				Expect(image).To(BeNil())
			})

			Specify("an error occurs", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
