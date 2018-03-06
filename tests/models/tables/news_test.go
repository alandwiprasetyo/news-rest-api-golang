package tables

import (
	"testing"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

func TestNews(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "News Model Test Suite")
}

var _ = ginkgo.Describe("Test News Model", func() {
	var _ = ginkgo.BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = ginkgo.AfterEach(func() {
		database.DropTable()
	})
	ginkgo.Describe("Test Create News Model", func() {

		ginkgo.It("Should create news models", func() {
			news := tables.News{Headline: "Headline", Title: "Title", Status: "draft", Description: "This is description", Tags: "Tags 1, Tag 2"}

			gomega.Expect(news.Headline).To(gomega.Equal("Headline"))
			gomega.Expect(news.Title).To(gomega.Equal("Title"))
			gomega.Expect(news.Status).To(gomega.Equal("draft"))
			gomega.Expect(news.Description).To(gomega.Equal("This is description"))
			gomega.Expect(news.Tags).To(gomega.Equal("Tags 1, Tag 2"))
		})
	})
})
