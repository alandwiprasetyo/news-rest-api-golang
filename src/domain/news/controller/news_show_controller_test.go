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
)

func TestShowNews(test *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(test, "NewsShowController Test Suite")
}

var _ = ginkgo.Describe("Test Show News", func() {
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
		req, _ := http.NewRequest("GET", fmt.Sprintf("/news/12313-1231"), nil)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		gomega.Expect(w.Code).To(gomega.Equal(http.StatusNotFound))
	})

	ginkgo.It("should return show news by ID", func() {
		product := models.News{
			Headline:    "Headline",
			Title:       "Title",
			Description: "This is description",
			Status:      "draft",
			Tags:        "Tag1, Tag 2",
		}
		database.GetDatabase().Create(&product)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/news/%d", product.ID), nil)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))
	})

})
