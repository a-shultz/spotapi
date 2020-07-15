package response_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"spotapi/response"
)

var _ = Describe("Copyright Response", func() {
	var (
		json      []byte
		copyright *response.CopyrightResponse
		err       error
	)

	Context("when a Copyright Response is created", func() {
		Context("with valid JSON", func() {
			BeforeEach(func() {
				json = ReadJSON("./testJSON/copyright/valid/typeP.json")
			})

			JustBeforeEach(func() {
				copyright, err = response.NewCopyrightResponse(json)
			})

			Specify("a Copyright Response pointer is returned", func() {
				Expect(copyright).ToNot(BeNil())
				Expect(copyright).To(BeAssignableToTypeOf(&response.CopyrightResponse{}))
			})

			Specify("no error occurs", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
		Context("with invalid JSON", func() {
			BeforeEach(func() {
				json = ReadJSON("./testJSON/copyright/invalid/invalid.json")
			})

			JustBeforeEach(func() {
				copyright, err = response.NewCopyrightResponse(json)
			})

			Specify("a nil pointer is returned", func() {
				Expect(copyright).To(BeNil())
			})

			Specify("an error occurs", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

var _ = Describe("Copyright", func() {
	var (
		json      []byte
		copyright *response.Copyright
	)

	BeforeEach(func() {
		json = ReadJSON("./testJSON/copyright/valid/typeP.json")
	})

	JustBeforeEach(func() {
		copyrightResponse, err := response.NewCopyrightResponse(json)
		Expect(copyrightResponse).ToNot(BeNil())
		Expect(err).ToNot(HaveOccurred())
		copyright = response.NewCopyright(copyrightResponse)
	})

	Context("when a Copyright is created", func() {
		Specify("a Copyright pointer is returned", func() {
			Expect(copyright).ToNot(BeNil())
			Expect(copyright).To(BeAssignableToTypeOf(&response.Copyright{}))
		})

		It("has the copyright text", func() {
			Expect(copyright.Text()).To(Equal("(P) 2019 Zelig Records, LLC"))
		})

		It("has the copyright type", func() {
			Expect(copyright.Type()).To(Equal("P"))
		})

		It("implements the stringer interface using the copyright text", func() {
			Expect(copyright.String()).To(Equal("(P) 2019 Zelig Records, LLC"))
		})
	})

	Context("that is type P", func() {
		It("is not type C", func() {
			Expect(copyright.IsTypeC()).To(BeFalse())
		})

		It("is type P", func() {
			Expect(copyright.IsTypeP()).To(BeTrue())
		})
	})

	Context("that is type C", func() {
		BeforeEach(func() {
			json = ReadJSON("./testJSON/copyright/valid/typeC.json")
		})

		It("is type C", func() {
			Expect(copyright.IsTypeC()).To(BeTrue())
		})

		It("is not type P", func() {
			Expect(copyright.IsTypeP()).To(BeFalse())
		})
	})
})
