package main

import (
	"test-sv/config"
	"test-sv/delivery/article_delivery"
	"test-sv/repository/article_repository"
	"test-sv/routes"
	"test-sv/usecase/article_usecase"

	"github.com/go-playground/validator/v10"
)

func main() {
	connection := config.Connect()
	validate := validator.New()
	articleRepository := article_repository.NewArticleRepository(connection)
	articleUsecase := article_usecase.NewArticleUsecase(articleRepository, validate)
	articleDelivery := article_delivery.NewArticleDelivery(articleUsecase)
	router := routes.NewRouter(articleDelivery)
	router.Run("localhost:8080")
}
