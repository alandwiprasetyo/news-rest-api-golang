package controller

import (
	"testing"
	"github.com/onsi/gomega"
	"github.com/onsi/ginkgo"
	"github.com/gin-gonic/gin"
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/migrations"
	"github.com/alandwiprasetyo/rest-api/src/models/seeders"
	"github.com/alandwiprasetyo/rest-api/src/routes"
	"net/http/httptest"
	"net/http"
	"bytes"
	"github.com/alandwiprasetyo/rest-api/src/models"
	"encoding/json"
	"fmt"
)

func TestDeleteTopic(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "TopicDeleteController Test Suite")
}

var _ = ginkgo.Describe("Test Delete Topic", func() {
	var router *gin.Engine

	var _ = ginkgo.BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
		router = routes.GetRouter()
	})

	var _ = ginkgo.AfterEach(func() {
		database.DropTable()
	})

	ginkgo.It("should success to delete topic", func() {

		payload := map[string]interface{}{
			"name":    "Topic name",
			"description": "this is description",
		}

		body, _ := json.Marshal(payload)
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/topics", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		gomega.Expect(w.Code).To(gomega.Equal(http.StatusCreated))

		latestTopic := models.Topic{}
		database.GetDatabase().Last(&latestTopic)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("DELETE", fmt.Sprintf("/topics/%d", latestTopic.ID), nil)
		router.ServeHTTP(w, req)
		gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

		latestTopicAfterDelete := models.Topic{}
		database.GetDatabase().Last(&latestTopicAfterDelete)

		gomega.Expect(latestTopic.ID).To(gomega.Not(gomega.Equal(latestTopicAfterDelete.ID)))
	})
})
