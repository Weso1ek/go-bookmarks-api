package dto

import "time"

type Category struct {
	Id              int32 `gorm:"primarykey"`
	Title           string
	Body            *string
	MetaTitle       string `gorm:"column:meta_title"`
	MetaDescription string `gorm:"column:meta_description"`
	Logo            *string
	Path            *string
	Slug            *string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

func (Category) TableName() string {
	return "categories"
}
