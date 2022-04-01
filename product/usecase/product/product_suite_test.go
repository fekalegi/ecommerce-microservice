package product_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestProductService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestProduct Suite")
}
