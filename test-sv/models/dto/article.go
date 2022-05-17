package dto

import "time"

type Article struct {
	Id          int32     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	Status      string    `json:"status"`
	CreatedDate time.Time `json:"createdDate"`
}
