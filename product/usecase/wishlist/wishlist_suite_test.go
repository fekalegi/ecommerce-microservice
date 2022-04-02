package wishlist_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestWishlistService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestWishlist Suite")
}
