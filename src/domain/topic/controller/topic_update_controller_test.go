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
	"github.com/alandwiprasetyo/rest-api/src/models"
	"net/http/httptest"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"bytes"
)

func TestUpdateTopic(test *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(test, "TopicShowController Test Suite")
}

var _ = ginkgo.Describe("Test Update Topic", func() {
	var router *gin.Engine

	var _ = ginkgo.BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
		router = routes.GetRouter()
	})

	var _ = ginkgo.AfterEach(func() {
		database.DropTable()
	})

	ginkgo.It("should return not found", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/topics/12313-1231"), nil)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		gomega.Expect(w.Code).To(gomega.Equal(http.StatusNotFound))
	})

	ginkgo.It("should update topic by ID", func() {
		product := models.Topic{
			Name:    "Topic Name example",
			Description: "This is description",
		}
		database.GetDatabase().Create(&product)

		payload := map[string]interface{}{
			"name":       "Topic Name example updated",
			"description": "this is description update",
		}
		body, _ := json.Marshal(payload)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/topics/%d", product.ID), bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		database.GetDatabase().Last(&product)

		gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))
		gomega.Expect(payload["name"].(string)).To(gomega.Equal(product.Name))
		gomega.Expect(payload["description"].(string)).To(gomega.Equal(product.Description))

	})

})
