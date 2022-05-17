package article_usecase

import "test-sv/models/dto"

type ArticleUsecase interface {
	GetAllArticle() dto.Response
	GetArticleByLimit(limit int, offset int) dto.Response
	GetArticleById(id int32) dto.Response
	CreateArticle(body dto.ArticleCreateRequest) dto.Response
	DeleteArticleById(id int32) dto.Response
	UpdateArticle(body dto.ArticleCreateRequest, id int32) dto.Response
}
