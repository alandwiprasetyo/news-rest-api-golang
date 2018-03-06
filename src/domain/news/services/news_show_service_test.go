package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"strconv"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

func TestNewsShowService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NewsShowService Test Suite")
}

var _ = Describe("Test NewsShowService", func() {
	var _ = BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = AfterEach(func() {
		database.DropTable()
	})
	Describe("Test func ShowNews", func() {
		It("should return not found", func() {
			service := NewsShowService{}
			res := service.ShowNews("1232-12")

			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeFalse())
		})

		It("should return news", func() {
			news := tables.News{
				Headline:    "Headline",
				Title:       "Title",
				Status:      "draft",
				Description: "This is description",
				Tags:        "Tags 1, Tag 2",
			}
			database.GetDatabase().Create(&news)

			service := NewsShowService{}
			res := service.ShowNews(strconv.Itoa(news.ID))

			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeTrue())
		})
	})
})
