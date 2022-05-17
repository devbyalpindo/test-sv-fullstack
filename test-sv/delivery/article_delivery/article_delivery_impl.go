package article_delivery

import (
	"net/http"
	"strconv"
	"test-sv/helpers"
	"test-sv/models/dto"
	"test-sv/usecase/article_usecase"

	"github.com/gin-gonic/gin"
)

type ArticleDeliveryImpl struct {
	usecase article_usecase.ArticleUsecase
}

func NewArticleDelivery(articleUsecase article_usecase.ArticleUsecase) ArticleDelivery {
	return &ArticleDeliveryImpl{articleUsecase}
}

func (res *ArticleDeliveryImpl) GetAllArticle(c *gin.Context) {

	response := res.usecase.GetAllArticle()
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *ArticleDeliveryImpl) GetArticleByLimit(c *gin.Context) {
	limit := c.Param("limit")
	offset := c.Param("offset")
	limits, err := strconv.Atoi(limit)
	helpers.PanicIfError(err)
	offsets, err := strconv.Atoi(offset)
	helpers.PanicIfError(err)

	response := res.usecase.GetArticleByLimit(limits, offsets)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *ArticleDeliveryImpl) GetArticleById(c *gin.Context) {
	articleId := c.Param("limit")
	id, err := strconv.Atoi(articleId)
	helpers.PanicIfError(err)

	response := res.usecase.GetArticleById(int32(id))
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *ArticleDeliveryImpl) CreateArticle(c *gin.Context) {
	articleCreateRequest := dto.ArticleCreateRequest{}
	if err := c.ShouldBindJSON(&articleCreateRequest); err != nil {
		errorRes := helpers.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.CreateArticle(articleCreateRequest)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (res *ArticleDeliveryImpl) DeleteArticleById(c *gin.Context) {
	id := c.Param("id")
	articleId, err := strconv.Atoi(id)
	helpers.PanicIfError(err)
	response := res.usecase.DeleteArticleById(int32(articleId))

	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (res *ArticleDeliveryImpl) UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	articleId, err := strconv.Atoi(id)
	helpers.PanicIfError(err)

	articleCreateRequest := dto.ArticleCreateRequest{}
	if err := c.ShouldBindJSON(&articleCreateRequest); err != nil {
		errorRes := helpers.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.UpdateArticle(articleCreateRequest, int32(articleId))
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
