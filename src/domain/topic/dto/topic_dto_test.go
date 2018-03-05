package dto

import (
	"testing"
	"github.com/onsi/gomega"
	"github.com/onsi/ginkgo"
	dto2 "github.com/alandwiprasetyo/rest-api/src/domain/topic/dto"
	"github.com/alandwiprasetyo/rest-api/src/common"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
	"github.com/alandwiprasetyo/rest-api/src/database"
)

func TestTopicDTO(test *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(test, "TopicDTO Test Suite")
}

var _ = ginkgo.Describe("Test TopicDTO", func() {
	var _ = ginkgo.BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = ginkgo.AfterEach(func() {
		database.DropTable()
	})
	ginkgo.It("Should return error validation", func() {
		dto := dto2.TopicDTO{}
		error := common.GetValidator().Struct(dto)
		gomega.Expect(error).To(gomega.Not(gomega.BeNil()))
	})

	ginkgo.It("should success validation", func() {
		dto := dto2.TopicDTO{
			Name:        "Topic Name",
			Description: "This is description",
		}
		error := common.GetValidator().Struct(dto)
		gomega.Expect(error).To(gomega.BeNil())
	})
})
