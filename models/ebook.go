package models

import "time"

type Ebook struct {
	ID            int64     `gorm:"column:id;primary_key"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
	Title         string    `gorm:"column:title"`
	Author        string    `gorm:"column:author"`
	Price         float64   `gorm:"column:price"`
	CoverImageURL string    `gorm:"column:cover_image_url"`
	ASIN          string    `gorm:"column:asin"`
	Description   string    `gorm:"column:description"`
}

func (Ebook) TableName() string {
	return "ebook"
}
