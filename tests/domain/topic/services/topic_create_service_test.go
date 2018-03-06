package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/dto"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/services"
)

func TestTopicCreateService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TopicCreateService Test Suite")
}

var _ = Describe("Test TopicCreateService", func() {
	var _ = BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = AfterEach(func() {
		database.DropTable()
	})

	Describe("Test func CreateTopic", func() {
		It("should return error validation", func() {
			dto := &dto.TopicDTO{}

			service := services.TopicCreateService{}
			res := service.CreateTopic(dto)

			Expect(res.ErrorValidation).To(Not(BeNil()))
			Expect(res.Error).To(Equal(""))
			Expect(res.Topic.ID).To(Equal(0))
		})

		It("should return topic", func() {
			dto := &dto.TopicDTO{
				Name:        "Topic Name",
				Description: "This is description",
			}

			service := services.TopicCreateService{}
			res := service.CreateTopic(dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.Topic.ID).To(Not(BeNil()))
		})
	})
})
