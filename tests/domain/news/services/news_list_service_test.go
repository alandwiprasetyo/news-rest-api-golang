package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/services"
)

func TestProductListService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NewsListService Test Suite")
}

var _ = Describe("Test NewsListService", func() {
	var _ = BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = AfterEach(func() {
		database.DropTable()
	})

	Describe("Test func ListNews", func() {
		It("should return news list", func() {
			product := tables.News{
				Headline:    "Headline",
				Title:       "Title",
				Status:      "draft",
				Description: "This is description",
				Tags:        "Tags 1, Tag 2",
			}
			database.GetDatabase().Create(&product)

			service := services.NewsListService{}
			res := service.ListNews("", "")

			Expect(res.Error).To(Equal(""))
		})
	})
})
