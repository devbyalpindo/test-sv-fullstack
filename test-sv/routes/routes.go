package routes

import (
	"test-sv/delivery/article_delivery"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(articleDelivery article_delivery.ArticleDelivery) *gin.Engine {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	router.GET("/article", articleDelivery.GetAllArticle)
	router.GET("/article/:limit", articleDelivery.GetArticleById)
	router.GET("/article/:limit/:offset", articleDelivery.GetArticleByLimit)
	router.POST("/article", articleDelivery.CreateArticle)
	router.DELETE("/article/:id", articleDelivery.DeleteArticleById)
	router.PUT("/article/:id", articleDelivery.UpdateArticle)

	return router
}
