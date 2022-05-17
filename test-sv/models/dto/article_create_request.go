package dto

type ArticleCreateRequest struct {
	Title    string `validate:"required,min=1,max=200" json:"title"`
	Content  string `validate:"required,min=1" json:"content"`
	Category string `validate:"required,min=1,max=200" json:"category"`
	Status   string `validate:"required,min=1,max=100" json:"status"`
}
