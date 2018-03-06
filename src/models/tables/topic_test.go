package tables

import (
	"testing"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
	"github.com/alandwiprasetyo/rest-api/src/database"
)

func TestTopic(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Topics Model Test Suite")
}

var _ = ginkgo.Describe("Test Topics Model", func() {
	var _ = ginkgo.BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
	})

	var _ = ginkgo.AfterEach(func() {
		database.DropTable()
	})
	ginkgo.Describe("Test Create Topics Model", func() {

		ginkgo.It("Should create Topics models", func() {
			topic := Topic{Name: "Topic name", Description: "This is description"}
			gomega.Expect(topic.Name).To(gomega.Equal("Topic name"))
			gomega.Expect(topic.Description).To(gomega.Equal("This is description"))
		})
	})
})
