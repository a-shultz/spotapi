package response_test

import (
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestResponse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Response Suite")
}

func ReadJSON(filename string) []byte {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return buf
}
