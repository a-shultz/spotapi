package response_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"spotapi/response"
)

var _ = Describe("Category", func() {
	var (
		json     []byte
		category *response.CategoryResponse
		err      error
	)

	BeforeEach(func() {
		json = []byte(`{
			"href": "https://api.spotify.com/v1/browse/categories/hiphop",
			"icons": [
			  {
				"height": 274,
				"url": "https://t.scdn.co/media/original/hip-274_0a661854d61e29eace5fe63f73495e68_274x274.jpg",
				"width": 274
			  }
			],
			"id": "hiphop",
			"name": "Hip-Hop"
		  }`)
	})

	Context("when a Copyright Response is created", func() {
		Context("with valid JSON", func() {
			JustBeforeEach(func() {
				category, err = response.NewCategoryResponse(json)
			})

			Specify("a Category Response pointer is returned", func() {
				Expect(category).ToNot(BeNil())
				Expect(category).To(BeAssignableToTypeOf(&response.CategoryResponse{}))
			})

			Specify("no error occurs", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
		Context("with invalid JSON", func() {
			BeforeEach(func() {
				json = []byte(`{
					"href": ""https://api.spotify.com/v1/browse/categories/hiphop",
					"icons": [
					  {
						"height": 274,
						"url": "https://t.scdn.co/media/original/hip-274_0a661854d61e29eace5fe63f73495e68_274x274.jpg",
						"width": 274
					  }
					],
					"id": "hiphop",
					"name": "Hip-Hop"
				}`)
			})

			JustBeforeEach(func() {
				category, err = response.NewCategoryResponse(json)
			})

			Specify("a nil pointer is returned", func() {
				Expect(category).To(BeNil())
			})

			Specify("an error occurs", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
