package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/dto"
	"github.com/alandwiprasetyo/rest-api/src/domain/news/services"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
	"github.com/alandwiprasetyo/rest-api/src/database"
)

func TestNewsCreateService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NewsCreateService Test Suite")
}

var _ = Describe("Test NewsCreateService", func() {
	var _ = BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = AfterEach(func() {
		database.DropTable()
	})

	Describe("Test func CreateNews", func() {
		It("should return error validation", func() {
			dto := &dto.NewsDTO{}

			service := services.NewsCreateService{}
			res := service.CreateNews(dto)

			Expect(res.ErrorValidation).To(Not(BeNil()))
			Expect(res.Error).To(Equal(""))
			Expect(res.News.ID).To(Equal(0))
		})

		It("should return product", func() {
			dto := &dto.NewsDTO{
				Headline:    "Headline",
				Title:       "Title",
				Status:      "draft",
				Description: "This is description",
				Tags:        "Tags 1, Tag 2",
			}

			service := services.NewsCreateService{}
			res := service.CreateNews(dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.News.ID).To(Not(BeNil()))
		})
	})
})
