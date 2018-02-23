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
)

func TestListTopic(test *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(test, "TopicListController Test Suite")
}

var _ = ginkgo.Describe("Test List Topic", func() {
	var router *gin.Engine

	var _ = ginkgo.BeforeEach(func() {
		migrations.Migration()
		seeders.Seeder()
		router = routes.GetRouter()
	})

	var _ = ginkgo.AfterEach(func() {
		database.DropTable()
	})

	ginkgo.It("should return product list", func() {
		product := models.Topic{
			Name:       "Topic Name",
			Description: "This is description",
		}
		database.GetDatabase().Create(&product)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/topics", nil)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))
	})

})
