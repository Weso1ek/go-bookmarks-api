package dto

import "time"

type Category struct {
	Id              int32   `gorm:"primarykey" json:"id"`
	Title           string  `json:"title"`
	Body            *string `json:"body"`
	MetaTitle       string  `gorm:"column:meta_title" json:"metaTitle"`
	MetaDescription string  `gorm:"column:meta_description" json:"metaDescription"`
	Logo            *string `json:"logo"`
	Path            *string `json:"path"`
	Slug            *string `json:"slug"`
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

func (Category) TableName() string {
	return "categories"
}
