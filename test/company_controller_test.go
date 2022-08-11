package test

import (
	"bytes"
	"encoding/json"
	"testcompany/config"
	"testcompany/controllers"
	"testcompany/database"
	"testcompany/models"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestShowAll(t *testing.T) {

	config.Init()
	database.StartDatabase()

	r := gin.Default()

	r.GET("/company/", controllers.ShowAll)
	req, err := http.NewRequest("GET", "/company/", nil)

	if err != nil {
		t.Fatalf("Not able to create request %s/n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var companies []models.Company
	json.Unmarshal(w.Body.Bytes(), &companies)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestCreate(t *testing.T) {

	config.Init()
	database.StartDatabase()

	r := gin.Default()
	r.POST("/company/", controllers.Create)

	company := models.Company{
		Name:     "Demo Company",
		ID:       6,
		Location: "Blr",
	}
	jsonValue, _ := json.Marshal(company)
	req, _ := http.NewRequest("POST", "/company/", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
