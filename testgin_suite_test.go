package testgo

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testgo Suite")
}
