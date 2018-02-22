package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/dto"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/services"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models"
	"strconv"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
)

func TestProductUpdateService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductUpdateService Test Suite")
}

var _ = Describe("Test ProductUpdateService", func() {

	var _ = BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = AfterEach(func() {
		database.DropTable()
	})

	Describe("Test func UpdateProduct", func() {
		It("should return validation error", func() {
			dto := &dto.NewsDTO{}

			service := services.NewsUpdateService{}
			res := service.UpdateNews("id", dto)

			Expect(res.ErrorValidation).To(Not(BeNil()))
			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeFalse())
		})

		It("should return not found", func() {
			dto := &dto.NewsDTO{
				Headline:    "Headline 2",
				Title:       "Title",
				Status:      "draft",
				Description: "This is description",
				Tags:        "Tags 1, Tag 2",
			}

			service := services.NewsUpdateService{}
			res := service.UpdateNews("id", dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeFalse())
		})

		It("should return product", func() {
			news := models.News{
				Headline:    "Headline 2",
				Title:       "Title",
				Status:      "draft",
				Description: "This is description",
				Tags:        "Tags 1, Tag 2",
			}
			database.GetDatabase().Create(&news)

			dto := &dto.NewsDTO{
				Headline:    "Headline 2",
				Title:       "Title",
				Status:      "draft",
				Description: "This is description",
				Tags:        "Tags 1, Tag 2",
			}

			service := services.NewsUpdateService{}
			res := service.UpdateNews(strconv.Itoa(news.ID), dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeTrue())
		})
	})
})
