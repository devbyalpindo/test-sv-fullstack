package article_repository

import (
	"test-sv/models/dto"
	"test-sv/models/entity"
)

type ArticleRepository interface {
	GetAllArticle() ([]entity.Article, error)
	GetArticleByLimit(limit int, offset int) (dto.ArticleWithTotal, error)
	GetArticleById(id int32) (*entity.Article, error)
	CreateArticle(entity.Article) (*entity.Article, error)
	DeleteArticleById(id int32) error
	UpdateArticle(article entity.Article, id int32) (*entity.Article, error)
}
