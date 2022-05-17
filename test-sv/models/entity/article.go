package entity

import (
	"time"
)

type Article struct {
	Id          int32     `gorm:"type:int;primary_key;auto_increment;not_null"`
	Title       string    `gorm:"type:varchar(200)"`
	Content     string    `gorm:"type:text"`
	Category    string    `gorm:"type:varchar(100)"`
	Status      string    `gorm:"type:varchar(100)"`
	CreatedDate time.Time `gorm:"type:timestamp"`
	UpdatedDate time.Time `gorm:"type:timestamp"`
}
