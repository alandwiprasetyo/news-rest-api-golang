package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"strconv"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/dto"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/services"
)

func TestTopicUpdateService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestTopicUpdateService Test Suite")
}

var _ = Describe("Test TestTopicUpdateService", func() {

	var _ = BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = AfterEach(func() {
		database.DropTable()
	})

	Describe("Test func UpdateTopic", func() {
		It("should return validation error", func() {
			dto := &dto.TopicDTO{}

			service := services.TopicUpdateService{}
			res := service.UpdateTopic("id", dto)

			Expect(res.ErrorValidation).To(Not(BeNil()))
			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeFalse())
		})

		It("should return not found", func() {
			dto := &dto.TopicDTO{
				Name:        "Name",
				Description: "This is description",
			}

			service := services.TopicUpdateService{}
			res := service.UpdateTopic("id", dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeFalse())
		})

		It("should return topic", func() {
			topic := tables.Topic{
				Name:        "Name",
				Description: "This is description",
			}
			database.GetDatabase().Create(&topic)

			dto := &dto.TopicDTO{
				Name:        "Name",
				Description: "This is description",
			}

			service := services.TopicUpdateService{}
			res := service.UpdateTopic(strconv.Itoa(topic.ID), dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeTrue())
		})
	})
})
