package dto

import (
	"testing"
	"github.com/onsi/gomega"
	"github.com/onsi/ginkgo"
	"github.com/alandwiprasetyo/rest-api/src/common"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
	"github.com/alandwiprasetyo/rest-api/src/database"
	dto2 "github.com/alandwiprasetyo/rest-api/src/domain/news/dto"
)

func TestNewsDTO(test *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(test, "NewsDTO Test Suite")
}

var _ = ginkgo.Describe("Test NewsDTO", func() {
	var _ = ginkgo.BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = ginkgo.AfterEach(func() {
		database.DropTable()
	})

	ginkgo.It("Should return error validation", func() {
		dto := dto2.NewsDTO{}
		error := common.GetValidator().Struct(dto)
		gomega.Expect(error).To(gomega.Not(gomega.BeNil()))
	})

	ginkgo.It("should success validation", func() {
		dto := dto2.NewsDTO{
			Headline:    "Headline",
			Title:       "Title",
			Status:      "draft",
			Description: "This is description",
			Tags:        "Tags 1, Tag 2",
		}
		error := common.GetValidator().Struct(dto)
		gomega.Expect(error).To(gomega.BeNil())
	})
})
