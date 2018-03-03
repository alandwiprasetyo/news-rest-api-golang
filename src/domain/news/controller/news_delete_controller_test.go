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

func TestDeleteNews(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "NewsDeleteController Test Suite")
}

var _ = ginkgo.Describe("Test Delete News", func() {
	var router *gin.Engine

	var _ = ginkgo.BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
		router = routes.GetRouter()
	})

	var _ = ginkgo.AfterEach(func() {
		database.DropTable()
	})

	ginkgo.It("should success to create news with valid payload", func() {
		payload := map[string]interface{}{

			"headline":    "headline",
			"title":       "title",
			"description": "this is description",
			"tags":        "tag, tes",
			"status":      "draft",
		}

		body, _ := json.Marshal(payload)
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/news", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		gomega.Expect(w.Code).To(gomega.Equal(http.StatusCreated))

		latestNews := models.News{}
		database.GetDatabase().Last(&latestNews)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("DELETE", fmt.Sprintf("/news/%d", latestNews.ID), nil)
		router.ServeHTTP(w, req)
		gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))

		latestNewsAfterDelete := models.News{}
		database.GetDatabase().Last(&latestNewsAfterDelete)

		gomega.Expect(latestNews.ID).To(gomega.Not(gomega.Equal(latestNewsAfterDelete.ID)))
	})
})
