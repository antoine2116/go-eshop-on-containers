package models

import (
	"time"
)

type Brand struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Brand     string    `json:"brand"`
	CreatedAt time.Time `json:"createdAt"` // https://gorm.io/docs/models.html#gorm-Model
	UpdatedAt time.Time `json:"updatedAt"` // https://gorm.io/docs/models.html#gorm-Model
}
