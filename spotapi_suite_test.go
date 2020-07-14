package spotapi_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSpotapi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spotapi Suite")
}
