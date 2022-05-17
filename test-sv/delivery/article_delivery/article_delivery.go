package article_delivery

import "github.com/gin-gonic/gin"

type ArticleDelivery interface {
	GetAllArticle(*gin.Context)
	GetArticleByLimit(*gin.Context)
	GetArticleById(*gin.Context)
	CreateArticle(*gin.Context)
	DeleteArticleById(*gin.Context)
	UpdateArticle(*gin.Context)
}
