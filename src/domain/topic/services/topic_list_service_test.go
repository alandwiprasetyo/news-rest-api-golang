package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/alandwiprasetyo/rest-api/src/models"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/services"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
)

func TestTopicListService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TopicListService Test Suite")
}

var _ = Describe("Test TopicListService", func() {
	var _ = BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = AfterEach(func() {
		database.DropTable()
	})

	Describe("Test func ListTopics", func() {
		It("should return topic list", func() {
			topic := models.Topic{
				Name:    "Topic Name",
				Description: "This is description",
			}
			database.GetDatabase().Create(&topic)

			service := services.TopicListService{}
			res := service.ListTopics()

			Expect(res.Error).To(Equal(""))
		})
	})
})
