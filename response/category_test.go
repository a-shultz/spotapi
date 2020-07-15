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
		json = ReadJSON("./testJSON/category/valid/valid.json")
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
				json = ReadJSON("./testJSON/category/invalid/invalid.json")
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
