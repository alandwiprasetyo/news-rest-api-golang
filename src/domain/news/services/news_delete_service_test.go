package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/services"
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"strconv"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
)

func TestProductDeleteService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductDeleteService Test Suite")
}

var _ = Describe("Test ProductDeleteService", func() {
	var _ = BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = AfterEach(func() {
		database.DropTable()
	})

	Describe("Test func DeleteProduct", func() {
		It("should return not found", func() {
			service := services.NewsDeleteService{}
			res := service.DeleteNews("12-21")

			Expect(res.Error).To(Equal("record not found"))
			Expect(res.IsFound).To(BeFalse())
		})

		It("should return product", func() {
			product := models.News{
				Headline:    "Headline",
				Title:       "Title",
				Status:      "draft",
				Description: "This is description",
				Tags:        "Tags 1, Tag 2",
			}
			database.GetDatabase().Create(&product)

			service := services.NewsDeleteService{}
			res := service.DeleteNews(strconv.Itoa(product.ID))

			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeTrue())
		})
	})
})
