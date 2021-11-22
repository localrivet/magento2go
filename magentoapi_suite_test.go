package magento2go_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMagentoApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Magento2go Suite")
}
