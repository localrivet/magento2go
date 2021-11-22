package magento2go_test

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/localrivet/magento2go"
	"github.com/localrivet/magento2go/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MagentoApi", func() {

	getConfig := func() magento2go.Config {
		return magento2go.Config{
			AccessToken: os.Getenv("MAGENTO_ACCESS_TOKEN"),
			Host:        os.Getenv("MAGENTO_STORE_HOSTNAME"),
			Path:        "/rest/default",
			Scheme:      os.Getenv("MAGENTO_STORE_SCHEME"),
			Debug:       false,
		}
	}

	validSku := "valid-sku"
	invalidSku := "invalid-sku"

	var timeout int64 = 300

	getEnv := func() {
		err := godotenv.Load()
		Expect(err).To(BeNil())
	}

	expectNilErr := func(c magento2go.Config) {
		_, err := magento2go.NewCommunityClient(c)
		Expect(err).To(BeNil())
	}

	BeforeEach(func() {
		getEnv()
	})

	Describe("Connection", func() {
		It("should return valid connection", func() {
			expectNilErr(getConfig())
		})
	})

	Describe("Api", func() {
		var mc *client.MagentoCommunity
		var err error
		var api *magento2go.MagentoApi

		BeforeEach(func() {
			mc, err = magento2go.NewCommunityClient(getConfig())
			Expect(err).To(BeNil())
		})

		It("should return a valid MagentoApi", func() {
			api = magento2go.NewMagentoApi(mc, timeout)
			Expect(api).To(Not(BeNil()))
		})

		Describe("GetProductBySku", func() {
			BeforeEach(func() {
				api = magento2go.NewMagentoApi(mc, timeout)
			})

			It("should return an error", func() {
				_, err := api.GetProductBySku(invalidSku)
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).Should(ContainSubstring("The product that was requested doesn't exist. Verify the product and try again."))
			})

			It("should return a valid product by sku", func() {
				product, err := api.GetProductBySku(validSku)
				Expect(err).To(BeNil())
				Expect(product).To(Not(BeNil()))
				Expect(*product.Sku).To(BeIdenticalTo(validSku))
			})
		})

		Describe("GetAllProducts", func() {
			var currentPage int64 = 0
			var pageSize int64 = 300

			BeforeEach(func() {
				api = magento2go.NewMagentoApi(mc, timeout)
			})

			It("should return an error", func() {
				currentPage = 0 // doesn't seem to throw any error
				pageSize = 301  // too many, the max is 300
				_, err := api.GetAllProducts(currentPage, pageSize)
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).Should(ContainSubstring("Maximum SearchCriteria pageSize is 300"))
			})

			It("should return a valid product by sku", func() {
				testSku := "bee-pure-gut-health"
				product, err := api.GetProductBySku(testSku)
				Expect(err).To(BeNil())
				Expect(product).To(Not(BeNil()))
				Expect(*product.Sku).To(BeIdenticalTo(testSku))
			})
		})
	})
})
