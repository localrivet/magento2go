package magento2go_test

import (
	"github.com/localrivet/magento2go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {

	config := magento2go.Config{}
	config.Debug = false

	expectErr := func(c magento2go.Config, errSubstring string) {
		_, err := magento2go.NewCommunityClient(c)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).Should(ContainSubstring(errSubstring))
	}

	expectNilErr := func(c magento2go.Config) {
		_, err := magento2go.NewCommunityClient(c)
		Expect(err).To(BeNil())
	}

	Describe("Config Params", func() {

		It("should return error if empty config", func() {
			expectErr((magento2go.Config{}), "config cannot be empty")
		})

		It("should return error if empty config accessToken", func() {
			config.Debug = false
			expectErr(config, "accessToken cannot be empty")
		})

		It("should return error if empty config host", func() {
			config.AccessToken = "test"
			expectErr(config, "host cannot be empty")
		})

		It("should return error if empty config path", func() {
			config.AccessToken = "test"
			config.Host = "test"
			expectErr(config, "path cannot be empty")
		})

		It("should return error if empty config scheme", func() {
			config.AccessToken = "test"
			config.Host = "test"
			config.Path = "test"
			expectErr(config, "scheme cannot be empty")
		})

		It("should not error if filled config", func() {
			config.AccessToken = "test"
			config.Host = "test"
			config.Path = "test"
			config.Scheme = "https"
			config.Debug = false
			expectNilErr(config)
		})

	})
})
