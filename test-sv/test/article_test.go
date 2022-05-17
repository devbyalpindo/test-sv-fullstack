package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"test-sv/delivery/article_delivery"
	"test-sv/helpers"
	"test-sv/models/entity"
	"test-sv/repository/article_repository"
	"test-sv/routes"
	"test-sv/usecase/article_usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func setupTestDB() *gorm.DB {
	dsn := "root:root@tcp(mariadb:3306)/article_test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helpers.PanicIfError(err)
	DB.AutoMigrate(&entity.Article{})
	return DB

}

func setupRouter(db *gorm.DB) *gin.Engine {
	connection := setupTestDB()
	validate := validator.New()
	articleRepository := article_repository.NewArticleRepository(connection)
	articleUsecase := article_usecase.NewArticleUsecase(articleRepository, validate)
	articleDelivery := article_delivery.NewArticleDelivery(articleUsecase)
	router := routes.NewRouter(articleDelivery)
	return router
}

func truncateArticle(db *gorm.DB) {
	db.Exec("TRUNCATE articles")
}

func TestCreateArticleSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	connection := setupTestDB()
	truncateArticle(connection)
	routes := setupRouter(connection)

	requestBody := strings.NewReader(`{"title" : "testing article", "content" : "berhasil menambah data article", "category" : "testing", "status" : "Draft"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/article", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	routes.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 201, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, int(responseBody["statusCode"].(float64)))
	assert.Equal(t, nil, responseBody["error"])

}

func TestCreateArticleFailed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	connection := setupTestDB()
	routes := setupRouter(connection)

	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/article", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	routes.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 403, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 403, int(responseBody["statusCode"].(float64)))
	assert.NotEqual(t, nil, responseBody["error"])

}

func TestGetArticleSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	connection := setupTestDB()
	routes := setupRouter(connection)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/article", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	routes.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["statusCode"].(float64)))
	assert.Equal(t, nil, responseBody["error"])

}

func TestUpdateArticleSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	connection := setupTestDB()
	routes := setupRouter(connection)

	requestBody := strings.NewReader(`{"title" : "merubah testing article", "content" : "berhasil merubah data article", "category" : "testing", "status" : "Publish"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/article/1", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	routes.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["statusCode"].(float64)))
	assert.Equal(t, nil, responseBody["error"])

}

func TestUpdateArticleFailed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	connection := setupTestDB()
	routes := setupRouter(connection)

	requestBody := strings.NewReader(`{"title" : "merubah testing article", "content" : "berhasil merubah data article", "category" : "testing", "status" : "Publish"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/article/1000", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	routes.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["statusCode"].(float64)))
	assert.NotEqual(t, nil, responseBody["error"])

}

func TestDeleteArticleSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	connection := setupTestDB()
	routes := setupRouter(connection)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/article/1", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	routes.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["statusCode"].(float64)))
	assert.Equal(t, nil, responseBody["error"])

}

func TestDeleteArticleFailed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	connection := setupTestDB()
	routes := setupRouter(connection)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/article/1000", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	routes.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["statusCode"].(float64)))
	assert.NotEqual(t, nil, responseBody["error"])

}
