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
)

func TestCreateNews(test *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(test, "NewsCreateController Test Suite")
}

var _ = ginkgo.Describe("Test Create News", func() {
	var router *gin.Engine

	var _ = ginkgo.BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
		router = routes.Setup()
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
	})

	//ginkgo.It("should error validation with empty payload", func() {
	//
	//	payload := map[string]interface{}{
	//		"headline":    "headline",
	//	}
	//	body, _ := json.Marshal(payload)
	//	w := httptest.NewRecorder()
	//	request, _ := http.NewRequest("POST", "/news", bytes.NewReader(body))
	//	request.Header.Add("Content-Type", "application/json")
	//	router.ServeHTTP(w, request)
	//	gomega.Expect(w.Code).To(gomega.Equal(http.StatusBadRequest))
	//})

})
