package magento2go

import (
	"magento2go/client"
	"os"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MagentoApi", func() {

	var timeout int64 = 300

	getEnv := func() {
		err := godotenv.Load()
		Expect(err).To(BeNil())
	}

	expectNilErr := func(c *struct {
		accessToken string
		host        string
		path        string
		scheme      string
		debug       bool
	}) {
		_, err := NewClient(c)
		Expect(err).To(BeNil())
	}

	BeforeEach(func() {
		getEnv()
	})

	Describe("Connection", func() {
		It("should return valid connection", func() {
			expectNilErr(&struct {
				accessToken string
				host        string
				path        string
				scheme      string
				debug       bool
			}{
				accessToken: os.Getenv("MAGENTO_ACCESS_TOKEN"),
				host:        os.Getenv("MAGENTO_STORE_HOSTNAME"),
				path:        "/rest/default",
				scheme:      os.Getenv("MAGENTO_STORE_SCHEME"),
				debug:       false,
			})
		})
	})

	Describe("Api", func() {

		var mc *client.MagentoCommunity
		var api *magentoApi

		BeforeEach(func() {
			mc, _ = NewClient(&struct {
				accessToken string
				host        string
				path        string
				scheme      string
				debug       bool
			}{
				accessToken: os.Getenv("MAGENTO_ACCESS_TOKEN"),
				host:        os.Getenv("MAGENTO_STORE_HOSTNAME"),
				path:        "/rest/default",
				scheme:      os.Getenv("MAGENTO_STORE_SCHEME"),
				debug:       false,
			})
		})

		It("should return a valid MagentoApi", func() {
			api = NewMagentoApi(mc, timeout)
			Expect(api).To(Not(BeNil()))
		})

		Describe("GetProductBySku", func() {
			BeforeEach(func() {
				api = NewMagentoApi(mc, timeout)
			})

			It("should return an error", func() {
				invalidSku := "yo-this-is-invalid"
				_, err := api.Product.GetProductBySku(invalidSku)
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).Should(ContainSubstring("The product that was requested doesn't exist. Verify the product and try again."))

			})

			It("should return a valid product by sku", func() {
				testSku := "bee-pure-gut-health"
				product, err := api.Product.GetProductBySku(testSku)
				Expect(err).To(BeNil())
				Expect(product).To(Not(BeNil()))
				Expect(*product.Sku).To(BeIdenticalTo(testSku))
			})
		})

		Describe("GetAllProducts", func() {
			var currentPage int64 = 0
			var pageSize int64 = 300

			BeforeEach(func() {
				api = NewMagentoApi(mc, timeout)
			})

			It("should return an error", func() {
				currentPage = 0 // doesn't seem to throw any error
				pageSize = 301  // too many, the max is 300
				_, err := api.Product.GetAllProducts(currentPage, pageSize)
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).Should(ContainSubstring("Maximum SearchCriteria pageSize is 300"))
			})

			It("should return a valid product by sku", func() {
				testSku := "bee-pure-gut-health"
				product, err := api.Product.GetProductBySku(testSku)
				Expect(err).To(BeNil())
				Expect(product).To(Not(BeNil()))
				Expect(*product.Sku).To(BeIdenticalTo(testSku))
			})
		})
	})
})
