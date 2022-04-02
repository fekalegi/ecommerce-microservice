package order_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestOrderService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestOrder Suite")
}
