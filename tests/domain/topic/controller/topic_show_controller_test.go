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
	"fmt"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

func TestShowTopic(test *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(test, "TopicShowController Test Suite")
}

var _ = ginkgo.Describe("Test Show Topic", func() {
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

	ginkgo.It("should return show topic by ID", func() {
		product := tables.Topic{
			Name:    "Topic name",
			Description: "This is description",
		}
		database.GetDatabase().Create(&product)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/topics/%d", product.ID), nil)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		gomega.Expect(w.Code).To(gomega.Equal(http.StatusOK))
	})

})
