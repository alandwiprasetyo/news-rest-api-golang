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
	"github.com/alandwiprasetyo/rest-api/src/domain/topic/services"
)

func TestTopicShowService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TopicShowService Test Suite")
}

var _ = Describe("Test TopicShowService", func() {
	var _ = BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = AfterEach(func() {
		database.DropTable()
	})
	Describe("Test func ShowTopic", func() {
		It("should return not found", func() {
			service := services.TopicShowService{}
			res := service.ShowTopic("1232-12")

			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeFalse())
		})

		It("should return topic", func() {
			topic := tables.Topic{
				Name:        "Topic Name",
				Description: "This is description",
			}
			database.GetDatabase().Create(&topic)

			service := services.TopicShowService{}
			res := service.ShowTopic(strconv.Itoa(topic.ID))

			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeTrue())
		})
	})
})
