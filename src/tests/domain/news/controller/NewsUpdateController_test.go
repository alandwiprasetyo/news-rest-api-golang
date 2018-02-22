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

func TestUpdateNews(test *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(test, "NewsShowController Test Suite")
}

var _ = ginkgo.Describe("Test Update News", func() {
	var router *gin.Engine

	var _ = ginkgo.BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
		router = routes.Setup()
	})

	var _ = ginkgo.AfterEach(func() {
		database.DropTable()
	})

	ginkgo.It("should return not found", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/news/12313-1231"), nil)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		gomega.Expect(w.Code).To(gomega.Equal(http.StatusNotFound))
	})

	ginkgo.It("should update news by ID", func() {
		product := models.News{
			Headline:    "Headline",
			Title:       "Title",
			Description: "This is description",
			Status:      "draft",
			Tags:        "Tag1, Tag 2",
		}
		database.GetDatabase().Create(&product)

		payload := map[string]interface{}{
			"headline":    "headline update",
			"title":       "title update",
			"description": "this is description update",
			"tags":        "tag, tes update",
			"status":      "publish",
		}
		body, _ := json.Marshal(payload)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/news/%d", product.ID), bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		database.GetDatabase().Last(&product)

		gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))
		gomega.Expect(payload["headline"].(string)).To(gomega.Equal(product.Headline))
		gomega.Expect(payload["title"].(string)).To(gomega.Equal(product.Title))
		gomega.Expect(payload["description"].(string)).To(gomega.Equal(product.Description))
		gomega.Expect(payload["tags"].(string)).To(gomega.Equal(product.Tags))
		gomega.Expect(payload["status"].(string)).To(gomega.Equal(product.Status))

	})

})
