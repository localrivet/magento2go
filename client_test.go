package magento2go

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {

	expectErr := func(c *struct {
		accessToken string
		host        string
		path        string
		scheme      string
		debug       bool
	}, errSubstring string) {
		_, err := NewClient(c)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).Should(ContainSubstring(errSubstring))
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

	Describe("Config Params", func() {

		It("should return error if empty config", func() {
			expectErr(nil, "config cannot be empty")
		})

		It("should return error if empty config accessToken", func() {
			expectErr(&struct {
				accessToken string
				host        string
				path        string
				scheme      string
				debug       bool
			}{
				accessToken: "",
				host:        "",
				path:        "",
				scheme:      "",
				debug:       false,
			}, "accessToken cannot be empty")
		})

		It("should return error if empty config host", func() {
			expectErr(&struct {
				accessToken string
				host        string
				path        string
				scheme      string
				debug       bool
			}{
				accessToken: "test",
				host:        "",
				path:        "",
				scheme:      "",
				debug:       false,
			}, "host cannot be empty")
		})

		It("should return error if empty config path", func() {
			expectErr(&struct {
				accessToken string
				host        string
				path        string
				scheme      string
				debug       bool
			}{
				accessToken: "test",
				host:        "test",
				path:        "",
				scheme:      "",
				debug:       false,
			}, "path cannot be empty")
		})

		It("should return error if empty config scheme", func() {
			expectErr(&struct {
				accessToken string
				host        string
				path        string
				scheme      string
				debug       bool
			}{
				accessToken: "test",
				host:        "test",
				path:        "test",
				scheme:      "",
				debug:       false,
			}, "scheme cannot be empty")
		})

		It("should not error if filled config", func() {
			expectNilErr(&struct {
				accessToken string
				host        string
				path        string
				scheme      string
				debug       bool
			}{
				accessToken: "test",
				host:        "test",
				path:        "test",
				scheme:      "https",
				debug:       false,
			})
		})

	})

	Describe("Runtime Client", func() {
		It("should not error if filled config", func() {
			expectNilErr(&struct {
				accessToken string
				host        string
				path        string
				scheme      string
				debug       bool
			}{
				accessToken: "test",
				host:        "test",
				path:        "test",
				scheme:      "https",
				debug:       false,
			})
		})
	})
})
