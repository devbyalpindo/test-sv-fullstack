package article_usecase

import (
	"errors"
	"test-sv/helpers"
	"test-sv/models/dto"
	"test-sv/models/entity"
	"test-sv/repository/article_repository"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ArticleUsecaseImpl struct {
	ArticleRepository article_repository.ArticleRepository
	Validate          *validator.Validate
}

func NewArticleUsecase(articleRepository article_repository.ArticleRepository, validate *validator.Validate) ArticleUsecase {
	return &ArticleUsecaseImpl{
		ArticleRepository: articleRepository,
		Validate:          validate,
	}
}

func (usecase *ArticleUsecaseImpl) GetAllArticle() dto.Response {
	articleList, err := usecase.ArticleRepository.GetAllArticle()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}
	helpers.PanicIfError(err)
	response := []dto.Article{}
	for _, article := range articleList {
		responseData := dto.Article{
			Id:          article.Id,
			Title:       article.Title,
			Content:     article.Content,
			Category:    article.Category,
			Status:      article.Status,
			CreatedDate: article.CreatedDate,
		}
		response = append(response, responseData)
	}

	return helpers.ResponseSuccess("ok", nil, response, 200)
}

func (usecase *ArticleUsecaseImpl) GetArticleByLimit(limit int, offset int) dto.Response {
	articleList, err := usecase.ArticleRepository.GetArticleByLimit(limit, offset)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}
	helpers.PanicIfError(err)
	response := []dto.Article{}
	for _, article := range articleList.Article {
		responseData := dto.Article{
			Id:          article.Id,
			Title:       article.Title,
			Content:     article.Content,
			Category:    article.Category,
			Status:      article.Status,
			CreatedDate: article.CreatedDate,
		}
		response = append(response, responseData)
	}

	var result map[string]any = make(map[string]any)
	result["total"] = articleList.Total
	result["article"] = response

	return helpers.ResponseSuccess("ok", nil, result, 200)
}

func (usecase *ArticleUsecaseImpl) GetArticleById(id int32) dto.Response {
	article, err := usecase.ArticleRepository.GetArticleById(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", "Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}
	helpers.PanicIfError(err)

	responseData := dto.Article{
		Id:          article.Id,
		Title:       article.Title,
		Content:     article.Content,
		Category:    article.Category,
		Status:      article.Status,
		CreatedDate: article.CreatedDate,
	}

	return helpers.ResponseSuccess("ok", nil, responseData, 200)
}

func (usecase *ArticleUsecaseImpl) CreateArticle(body dto.ArticleCreateRequest) dto.Response {
	err := usecase.Validate.Struct(body)

	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]dto.CustomMessageError, len(ve))
			for i, fe := range ve {
				out[i] = dto.CustomMessageError{
					Field:    fe.Field(),
					Messsage: helpers.MessageError(fe.Tag()),
				}
			}
			return helpers.ResponseError("Bad Request", out, 403)
		}

	}

	if status := helpers.CheckStatus(body.Status); !status {
		return helpers.ResponseError("Bad Request", "Status Unknow", 403)
	}

	payloadArticle := entity.Article{
		Title:       body.Title,
		Content:     body.Content,
		Category:    body.Category,
		Status:      body.Status,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}

	article, err := usecase.ArticleRepository.CreateArticle(payloadArticle)
	helpers.PanicIfError(err)

	return helpers.ResponseSuccess("ok", nil, article, 201)

}

func (usecase *ArticleUsecaseImpl) DeleteArticleById(id int32) dto.Response {
	err := usecase.ArticleRepository.DeleteArticleById(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err.Error(), 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err.Error(), 500)
	}
	return helpers.ResponseSuccess("ok", nil, nil, 200)

}

func (usecase *ArticleUsecaseImpl) UpdateArticle(body dto.ArticleCreateRequest, id int32) dto.Response {
	err := usecase.Validate.Struct(body)

	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]dto.CustomMessageError, len(ve))
			for i, fe := range ve {
				out[i] = dto.CustomMessageError{
					Field:    fe.Field(),
					Messsage: helpers.MessageError(fe.Tag()),
				}
			}
			return helpers.ResponseError("Bad Request", out, 403)
		}

	}

	if status := helpers.CheckStatus(body.Status); !status {
		return helpers.ResponseError("Bad Request", "Status Unknow", 403)
	}

	payloadArticle := entity.Article{
		Title:       body.Title,
		Content:     body.Content,
		Category:    body.Category,
		Status:      body.Status,
		UpdatedDate: time.Now(),
	}

	article, err := usecase.ArticleRepository.UpdateArticle(payloadArticle, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Article not found", err.Error(), 404)
	}
	helpers.PanicIfError(err)

	return helpers.ResponseSuccess("ok", nil, article, 200)

}
