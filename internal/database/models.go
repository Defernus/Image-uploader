package database

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Original  string         `json:"original"`
	Thumbnail string         `json:"thumbnail"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
