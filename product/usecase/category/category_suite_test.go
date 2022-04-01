package category_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCategoryService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestCategory Suite")
}
