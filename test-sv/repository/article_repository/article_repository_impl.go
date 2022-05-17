package article_repository

import (
	"test-sv/helpers"
	"test-sv/models/dto"
	"test-sv/models/entity"

	"gorm.io/gorm"
)

type ArticleRepositoryImpl struct {
	DB *gorm.DB
}

func NewArticleRepository(DB *gorm.DB) ArticleRepository {
	return &ArticleRepositoryImpl{DB}
}

func (repository *ArticleRepositoryImpl) GetAllArticle() ([]entity.Article, error) {
	article := []entity.Article{}
	err := repository.DB.Model(&entity.Article{}).Scan(&article).Error
	helpers.PanicIfError(err)
	if len(article) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return article, nil
}

func (repository *ArticleRepositoryImpl) GetArticleByLimit(limit int, offset int) (dto.ArticleWithTotal, error) {
	article := []entity.Article{}
	var totalData int64
	err := repository.DB.Model(&entity.Article{}).Where("status = ?", "Published").Count(&totalData).Limit(limit).Offset(offset).Scan(&article).Error
	helpers.PanicIfError(err)
	if len(article) <= 0 {
		resultError := dto.ArticleWithTotal{
			Total:   0,
			Article: nil,
		}
		return resultError, gorm.ErrRecordNotFound
	}

	result := dto.ArticleWithTotal{
		Total:   totalData,
		Article: article,
	}

	return result, nil
}

func (repository *ArticleRepositoryImpl) GetArticleById(id int32) (*entity.Article, error) {
	article := entity.Article{}
	result := repository.DB.Where("id = ?", id).Find(&article)

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &article, nil
}

func (repository *ArticleRepositoryImpl) CreateArticle(article entity.Article) (*entity.Article, error) {
	if err := repository.DB.Create(&article).Error; err != nil {
		return nil, err
	}

	return &article, nil
}

func (repository *ArticleRepositoryImpl) DeleteArticleById(id int32) error {
	result := repository.DB.Where("id = ?", id).Delete(&entity.Article{})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repository *ArticleRepositoryImpl) UpdateArticle(article entity.Article, id int32) (*entity.Article, error) {
	result := repository.DB.Model(&article).Where("id = ?", id).Updates(entity.Article{Title: article.Title, Content: article.Content, Category: article.Category, Status: article.Status, UpdatedDate: article.UpdatedDate})

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &article, nil
}
