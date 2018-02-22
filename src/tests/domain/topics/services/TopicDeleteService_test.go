package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"strconv"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/services"
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
)

func TestTopicDeleteService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TopicDeleteService Test Suite")
}

var _ = Describe("Test TopicDeleteService", func() {
	var _ = BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = AfterEach(func() {
		database.DropTable()
	})

	Describe("Test func DeleteTopic", func() {
		It("should return not found", func() {
			service := services.TopicDeleteService{}
			res := service.DeleteTopic("2121212121")

			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeTrue())
		})

		It("should return topic", func() {
			topic := models.Topic{
				Name:    "Topic Name",
				Description: "This is description",
			}
			database.GetDatabase().Create(&topic)

			service := services.TopicDeleteService{}
			res := service.DeleteTopic(strconv.Itoa(topic.ID))

			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeTrue())
		})
	})
})
