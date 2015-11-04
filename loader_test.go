package pobin_test

import (
	"github.com/flosch/pongo2"
	. "github.com/sharpner/pobin"
	"github.com/sharpner/pobin/assets"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test the loader", func() {
	Context("No Fallback loading works", func() {
		var (
			loader      pongo2.TemplateLoader
			templateSet *pongo2.TemplateSet
		)

		BeforeEach(func() {
			loader = NewMemoryTemplateLoader(assets_test.Asset)
			templateSet = pongo2.NewSet("testing", loader)
		})

		It("should be validly parsed", func() {
			loginTemplate, err := templateSet.FromFile("templates/sites/index.tpl")
			Expect(err).ToNot(HaveOccurred())

			out, err := loginTemplate.Execute(pongo2.Context{
				"action": "/login",
			})

			Expect(err).ToNot(HaveOccurred())
			Expect(string(out)).To(ContainSubstring("Example login page"))
			Expect(string(out)).To(ContainSubstring("<title>"))
			Expect(string(out)).To(ContainSubstring("/login"))
		})

		It("should be validly parsed when calling from cache", func() {
			loginTemplate, err := templateSet.FromCache("templates/sites/index.tpl")
			Expect(err).ToNot(HaveOccurred())

			out, err := loginTemplate.Execute(pongo2.Context{
				"action": "/login",
			})

			Expect(err).ToNot(HaveOccurred())
			Expect(string(out)).To(ContainSubstring("Example login page"))
			Expect(string(out)).To(ContainSubstring("<title>"))
			Expect(string(out)).To(ContainSubstring("/login"))
		})
	})
})
