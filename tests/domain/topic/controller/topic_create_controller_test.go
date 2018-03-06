package controller

import (
	"testing"
	"github.com/onsi/gomega"
	"github.com/onsi/ginkgo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"net/http/httptest"
	"net/http"
	"bytes"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
	"github.com/alandwiprasetyo/rest-api/src/routes"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

func TestCreateTopic(test *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(test, "TopicCreateController Test Suite")
}

var _ = ginkgo.Describe("Test Create Topic", func() {
	var router *gin.Engine

	var _ = ginkgo.BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
		router = routes.GetRouter()
	})

	var _ = ginkgo.AfterEach(func() {
		database.DropTable()
	})

	ginkgo.It("should success to create topic with valid payload", func() {
		topic := tables.Topic{}
		database.GetDatabase().Last(&topic)
		payload := map[string]interface{}{
			"name":    "Topic name",
			"description": "this is description",
		}

		body, _ := json.Marshal(payload)
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/topics", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		newData := tables.Topic{}
		database.GetDatabase().Last(&newData)

		gomega.Expect(w.Code).To(gomega.Equal(http.StatusCreated))
		gomega.Expect(newData.ID).To(gomega.Not(gomega.Equal(topic.ID)))
	})

	ginkgo.It("should error validation with empty payload", func() {
		topic := tables.Topic{}
		database.GetDatabase().Last(&topic)

		payload := map[string]interface{}{
			"headline": "headline",
		}
		body, _ := json.Marshal(payload)
		w := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/topics", bytes.NewReader(body))
		request.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, request)

		newData := tables.Topic{}
		database.GetDatabase().Last(&newData)

		gomega.Expect(w.Code).To(gomega.Equal(http.StatusBadRequest))
		gomega.Expect(newData.ID).To(gomega.Equal(topic.ID))
	})
})
