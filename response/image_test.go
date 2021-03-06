package response_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"net/url"

	"spotapi/response"
)

var _ = Describe("Image Response", func() {
	var (
		data  []byte
		image *response.ImageResponse
		err   error
	)

	JustBeforeEach(func() {
		image, err = response.NewImageResponse(data)
	})

	Context("when an Image Response is created", func() {
		Context("with valid JSON", func() {
			BeforeEach(func() {
				data = ReadJSON("./testJSON/image/valid/valid.json")
			})

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
				data = ReadJSON("./testJSON/image/invalid/invalid.json")
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

var _ = Describe("Image", func() {
	var (
		data          []byte
		imageResponse *response.ImageResponse
		responseErr   error
		image         *response.Image
		err           error
	)

	Context("when an Image is created", func() {
		Context("with a valid url", func() {
			BeforeEach(func() {
				data = ReadJSON("./testJSON/image/valid/valid.json")
			})

			JustBeforeEach(func() {
				imageResponse, responseErr = response.NewImageResponse(data)
				Expect(imageResponse).ToNot(BeNil())
				Expect(responseErr).ToNot(HaveOccurred())
				image, err = response.NewImage(imageResponse)
			})

			Specify("an Image pointer is returned", func() {
				Expect(image).ToNot(BeNil())
				Expect(image).To(BeAssignableToTypeOf(&response.Image{}))
			})

			Specify("an error doesn't occur", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("has a height", func() {
				Expect(image.Height()).To(Equal(640))
			})

			It("has a width", func() {
				Expect(image.Width()).To(Equal(640))
			})

			It("has a URL", func() {
				testURL, err := url.Parse("https://i.scdn.co/image/ab67616d0000b273f142205e336ec0af3e1a4eb0")
				Expect(err).ToNot(HaveOccurred())
				Expect(image.URL()).To(BeEquivalentTo(testURL))
			})

			It("has a RawURL", func() {
				Expect(image.RawURL()).To(Equal("https://i.scdn.co/image/ab67616d0000b273f142205e336ec0af3e1a4eb0"))
			})

			It("implements the Stringer interface", func() {
				Expect(image.String()).To(Equal("https://i.scdn.co/image/ab67616d0000b273f142205e336ec0af3e1a4eb0 (640 x 640)"))
			})
		})

		DescribeTable("with invalid data",
			func(data []byte) {
				imageResponse, responseErr = response.NewImageResponse(data)
				Expect(imageResponse).ToNot(BeNil())
				Expect(responseErr).ToNot(HaveOccurred())

				image, err = response.NewImage(imageResponse)
				Expect(image).To(BeNil())
				Expect(err).To(HaveOccurred())
			},
			Entry("with an invalid url", ReadJSON("./testJSON/image/invalid/urlParse.json")),
			Entry("with no scheme", ReadJSON("./testJSON/image/invalid/noSchema.json")),
			Entry("with no host", ReadJSON("./testJSON/image/invalid/noHost.json")),
		)
	})
})
